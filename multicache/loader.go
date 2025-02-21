package multicache

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type GetKeysFunc func(ctx context.Context) ([]string, error)
type FetchAndCacheKeyFunc func(ctx context.Context, key string) error

type Loader struct {
	GetKeys          GetKeysFunc
	FetchAndCacheKey FetchAndCacheKeyFunc

	maxConcurrency   int
	reloadShuffleSec int
	reloadTickerSec  int
	reloadCount      int
	reloadMaxCount   int
	logger           *log.Helper
}

func NewLoader(getKeys GetKeysFunc, fetchAndCacheKey FetchAndCacheKeyFunc, config *Config, logger log.Logger) *Loader {
	return &Loader{
		GetKeys:          getKeys,
		FetchAndCacheKey: fetchAndCacheKey,

		maxConcurrency:   config.MaxConcurrency,
		reloadShuffleSec: config.ReloadShuffleSec,
		reloadTickerSec:  config.ReloadTickerSec,
		reloadCount:      0,
		reloadMaxCount:   config.ReloadMaxCount,
		logger:           log.NewHelper(logger),
	}
}

func (l *Loader) sleepRandom() {
	// sleep a random time to avoid all pods reload in same time
	// reduce reload stress of redis
	randomDurationSec := l.reloadShuffleSec
	if randomDurationSec < l.reloadTickerSec/2 {
		randomDurationSec = l.reloadTickerSec / 2
	}
	time.Sleep(time.Duration(rand.Intn(randomDurationSec)) * time.Second)
}

func (l *Loader) StartLoad() {
	// 第一次加载必须成功
	l.load(true)
	go l.startTickLoad()
}

func (l *Loader) startTickLoad() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			// start again
			l.reload()
		}
	}()

	l.sleepRandom()
	ticker := time.NewTicker(time.Duration(l.reloadTickerSec) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		l.load(false)
	}
}

// 重试 reload，有最大次数限制
func (l *Loader) reload() {
	l.reloadCount++
	// 限制重启次数，避免无限 goroutine 情况
	if l.reloadCount > l.reloadMaxCount {
		l.logger.Errorf("reload times too many, reloadCount: %d", l.reloadCount)
		return
	}

	go l.startTickLoad()

	l.logger.Infof("start reload redis cache, reloadCount: %d", l.reloadCount)
}

func (l *Loader) load(must bool) {
	// 扫描所有需要拉取的keys
	start := time.Now()
	allKeys, err := l.GetKeys(context.TODO())
	if err != nil {
		if must {
			panic(err)
		}
		return
	}
	scanCost := time.Since(start)

	// 并发拉取所有key的数据缓存到本地
	var (
		wg       sync.WaitGroup
		workers  = make(chan struct{}, l.maxConcurrency)
		loadfail int
	)
	defer close(workers)

	start = time.Now()
	for _, key := range allKeys {
		workers <- struct{}{}

		wg.Add(1)
		go func(k string) {
			defer func() {
				wg.Done()
				<-workers
			}()

			if err := l.FetchAndCacheKey(context.TODO(), k); err != nil {
				loadfail++
			}
		}(key)
	}
	wg.Wait()
	loadCost := time.Since(start)

	key := ""
	if len(allKeys) > 0 {
		key = allKeys[0]
	}

	l.logger.Infof("load redis cache success, keys[0]: %s, scan_cost: %dms, load_cost: %dms, scan_keys: %d, fail: %d", key, scanCost.Milliseconds(), loadCost.Milliseconds(), len(allKeys), loadfail)
}
