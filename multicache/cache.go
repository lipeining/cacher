package multicache

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"duffle-cache/internal/data/localcache"
	"duffle-cache/internal/data/rdbcache"
	"duffle-cache/internal/util"

	"github.com/go-kratos/kratos/v2/log"
	gocache "github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
)

const (
	RedisKeyTypeHash    = "hash"
	RedisKeyTypeString  = "string"
	RedisKeyTypeSet     = "set"
	RedisMaxConcurrency = 5  // 最大并发打包数量
	ReloadShuffleSec    = 30 // 随机延迟reload因子，用于各实例reload打散
	ReloadMaxCount      = 5  // 最大重试次数
)

type Config struct {
	ScanRule         string
	ScanType         string
	MaxConcurrency   int
	ReloadShuffleSec int
	ReloadMaxCount   int
	ReloadTickerSec  int
	LocalCacheClose  bool
}

type MultiCache[T any] struct {
	serializer Serializer[T]
	logger     *log.Helper
	rdbCache   rdbcache.RedisCache
	localCache *localcache.ShardedCache
	sf         singleflight.Group

	scanRule         string
	scanType         string
	maxConcurrency   int
	reloadShuffleSec int
	reloadTickerSec  int
	localCacheClose  bool
	reloadCount      int
	reloadMaxCount   int
}

func NewMultiCache[T any](rdbCache rdbcache.RedisCache, localCache *localcache.ShardedCache, config *Config, serializer Serializer[T], logger log.Logger) *MultiCache[T] {
	cli := &MultiCache[T]{
		localCache:       localCache,
		rdbCache:         rdbCache,
		scanRule:         config.ScanRule,
		scanType:         config.ScanType,
		maxConcurrency:   config.MaxConcurrency,
		reloadShuffleSec: config.ReloadShuffleSec,
		reloadTickerSec:  config.ReloadTickerSec,
		localCacheClose:  config.LocalCacheClose,
		reloadCount:      0,
		reloadMaxCount:   config.ReloadMaxCount,
		serializer:       serializer,
		logger:           log.NewHelper(logger),
	}
	return cli
}

func (m *MultiCache[T]) sleepRandom() {
	// sleep a random time to avoid all pods reload in same time
	// reduce reload stress of redis
	randomDurationSec := m.reloadShuffleSec
	if randomDurationSec < m.reloadTickerSec/2 {
		randomDurationSec = m.reloadTickerSec / 2
	}
	time.Sleep(time.Duration(rand.Intn(randomDurationSec)) * time.Second)
}

func (m *MultiCache[T]) StartLoad() {
	if m.localCacheClose {
		return
	}

	// 第一次加载必须成功
	m.load(true)
	go m.startTickLoad()
}

func (m *MultiCache[T]) startTickLoad() {
	defer func() {
		if err := recover(); err != nil {
			util.PrintPanicStack(err, m.logger)
			// start again
			m.reload()
		}
	}()

	m.sleepRandom()
	ticker := time.NewTicker(time.Duration(m.reloadTickerSec) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		m.load(false)
	}
}

// 重试 reload，有最大次数限制
func (m *MultiCache[T]) reload() {
	m.reloadCount++
	// 限制重启次数，避免无限 goroutine 情况
	if m.reloadCount > m.reloadMaxCount {
		m.logger.Errorf("reload times too many, reloadCount: %d", m.reloadCount)
		return
	}

	go m.startTickLoad()

	m.logger.Infof("start reload redis cache, reloadCount: %d", m.reloadCount)
}

func (m *MultiCache[T]) load(must bool) {
	// 扫描所有需要拉取的keys
	start := time.Now()
	allKeys, err := m.rdbCache.Scan(context.TODO(), m.scanRule, m.scanType)
	if err != nil {
		if must {
			panic(err)
		}
		m.logger.Errorf("scan redis duffle's all keys fail, scan_rule: %s, scan_type: %s, err: %v", m.scanRule, m.scanType, err)
		return
	}
	scanCost := time.Since(start)

	// 并发拉取所有key的数据缓存到本地
	var (
		wg       sync.WaitGroup
		workers  = make(chan struct{}, m.maxConcurrency)
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

			if err := m.fetchAndCacheKey(context.TODO(), k); err != nil {
				loadfail++
			}
		}(key)
	}
	wg.Wait()
	loadCost := time.Since(start)

	m.logger.Infof("load redis cache success, scan_rule: %s, scan_type: %s, scan_cost: %dms, load_cost: %dms, scan_keys: %d, fail: %d", m.scanRule, m.scanType, scanCost.Milliseconds(), loadCost.Milliseconds(), len(allKeys), loadfail)
}

func (m *MultiCache[T]) fetchAndCacheKey(ctx context.Context, key string) error {
	if m.scanType == RedisKeyTypeHash {
		return m.loadHash(ctx, key)
	} else if m.scanType == RedisKeyTypeString {
		return m.loadString(ctx, key)
	}

	m.logger.Errorf("unsupported scan type, scan_type: %s", m.scanType)
	return errors.Errorf("unsupported scan type: %s", m.scanType)
}

func (m *MultiCache[T]) loadHash(ctx context.Context, key string) error {
	data, err := m.rdbCache.HGetAllByBatch(ctx, key)
	if err != nil {
		m.logger.Errorf("load redis data fail, key: %s, err: %v", key, err)
		return err
	}

	object := make(map[string]*T)
	for field, value := range data {
		var v T
		if err := m.serializer.Unmarshal([]byte(value), &v); err != nil {
			m.logger.Errorf("unmarshal redis data fail, key: %s, field: %s, err: %v", key, field, err)
			return err
		}
		object[field] = &v
	}

	m.localCache.Set(key, object, gocache.NoExpiration)
	return nil
}

func (m *MultiCache[T]) loadString(ctx context.Context, key string) error {
	data, err := m.rdbCache.Get(ctx, key)
	if err != nil {
		m.logger.Errorf("load redis data fail, key: %s, err: %v", key, err)
		return err
	}

	var v T
	if err := m.serializer.Unmarshal([]byte(data), &v); err != nil {
		m.logger.Errorf("unmarshal redis data fail, key: %s,  err: %v", key, err)
		return err
	}

	m.localCache.Set(key, &v, gocache.NoExpiration)
	return nil
}

func (m *MultiCache[T]) GetMap(ctx context.Context, key string) (map[string]*T, error) {
	if !m.localCacheClose {
		val, ok := m.localCache.Get(key)
		if ok {
			if valMap, ok := val.(map[string]*T); ok {
				return valMap, nil
			} else {
				// key type not match hash
				m.logger.Errorf("multi-cache key type not hash, key: %s", key)
				return nil, errors.Errorf("key type is not hash, key: %s", key)
			}
		}

		m.logger.Warnf("multi-cache hash key miss, key: %s", key)
	}

	// penetrate local cache, get data from redis
	ret, err, _ := m.sf.Do(key, func() (interface{}, error) {
		return m.rdbCache.HGetAll(ctx, key)
	})
	if err != nil {
		return nil, err
	}

	object := make(map[string]*T)
	if retMap, ok := ret.(map[string]string); ok {
		for field, value := range retMap {
			var v T
			if err := m.serializer.Unmarshal([]byte(value), &v); err != nil {
				m.logger.Errorf("unmarshal redis data fail, key: %s, field: %s, err: %v", key, field, err)
				return nil, err
			}
			object[field] = &v
		}
	}

	if !m.localCacheClose {
		m.localCache.Set(key, object, gocache.NoExpiration)
	}

	return object, nil
}

func (m *MultiCache[T]) GetMapMul(ctx context.Context, key string, fields ...string) ([]*T, error) {
	if !m.localCacheClose {
		val, ok := m.localCache.Get(key)
		if ok {
			if valMap, ok := val.(map[string]*T); ok {
				ret := make([]*T, 0, len(fields))
				for _, f := range fields {
					if v, ok := valMap[f]; ok {
						ret = append(ret, v)
					} else {
						var tmp *T
						ret = append(ret, tmp)
					}
				}
				return ret, nil
			} else {
				// key type not match hash
				m.logger.Errorf("multi-cache key type not hash, key: %s", key)
				return nil, errors.Errorf("key type is not hash, key: %s", key)
			}
		}
		m.logger.Warnf("multi-cache hash key miss, key: %s", key)
	}

	ret, err, _ := m.sf.Do(key, func() (interface{}, error) {
		return m.rdbCache.HMGet(ctx, key, fields...)
	})
	if err != nil {
		return nil, err
	}

	object := make([]*T, 0, len(fields))
	if retMap, ok := ret.([]string); ok {
		for i, value := range retMap {
			var v T
			if err := m.serializer.Unmarshal([]byte(value), &v); err != nil {
				m.logger.Errorf("unmarshal redis data fail, key: %s, field: %s, err: %v", key, fields[i], err)
				return nil, err
			}
			object = append(object, &v)
		}
	}

	return object, nil
}

func (m *MultiCache[T]) Get(ctx context.Context, key string) (*T, error) {
	if !m.localCacheClose {
		val, ok := m.localCache.Get(key)
		if ok {
			if v, ok := val.(*T); ok {
				return v, nil
			} else {
				m.logger.Errorf("multi-cache key type not match, key: %s", key)
				return nil, errors.Errorf("key type not match, key: %s", key)
			}
		}
		m.logger.Warnf("multi-cache key miss, key: %s", key)
	}

	ret, err, _ := m.sf.Do(key, func() (interface{}, error) {
		return m.rdbCache.Get(ctx, key)
	})
	if err != nil {
		return nil, err
	}

	if data, ok := ret.(string); ok {
		var v T
		if err := m.serializer.Unmarshal([]byte(data), &v); err != nil {
			m.logger.Errorf("unmarshal redis data fail, key: %s, err: %v", key, err)
			return nil, err
		}
		if !m.localCacheClose {
			m.localCache.Set(key, &v, gocache.NoExpiration)
		}
		return &v, nil
	}

	return nil, errors.Errorf("multi-cache get key fail, key: %s", key)
}
