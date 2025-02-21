package localcache

import (
	"runtime"
	"time"

	"github.com/cespare/xxhash/v2"
	gocache "github.com/patrickmn/go-cache"
)

const (
	// default expired keys cleaned up duration(1min)
	DefaultCleanUp time.Duration = time.Minute
)

type ShardedCache struct {
	*shardedCache
}

type shardedCache struct {
	segmentSize2N int
	segmentMark   uint64
	defaultExpire time.Duration
	initItemSize  int

	caches  []*gocache.Cache
	janitor *shardedJanitor
}

func (sc *shardedCache) segment(k string) *gocache.Cache {
	return sc.caches[xxhash.Sum64([]byte(k))&sc.segmentMark]
}

func (sc *shardedCache) Set(k string, x interface{}, d time.Duration) {
	sc.segment(k).Set(k, x, d)
}

func (sc *shardedCache) Add(k string, x interface{}, d time.Duration) error {
	return sc.segment(k).Add(k, x, d)
}

func (sc *shardedCache) Replace(k string, x interface{}, d time.Duration) error {
	return sc.segment(k).Replace(k, x, d)
}

func (sc *shardedCache) Get(k string) (interface{}, bool) {
	return sc.segment(k).Get(k)
}

func (sc *shardedCache) Increment(k string, n int64) error {
	return sc.segment(k).Increment(k, n)
}

func (sc *shardedCache) IncrementFloat(k string, n float64) error {
	return sc.segment(k).IncrementFloat(k, n)
}

func (sc *shardedCache) Decrement(k string, n int64) error {
	return sc.segment(k).Decrement(k, n)
}

func (sc *shardedCache) DecrementFloat(k string, n float64) error {
	return sc.segment(k).DecrementFloat(k, n)
}

func (sc *shardedCache) Delete(k string) {
	sc.segment(k).Delete(k)
}

func (sc *shardedCache) DeleteExpired() {
	for _, v := range sc.caches {
		v.DeleteExpired()
	}
}

func (sc *shardedCache) Items() []map[string]gocache.Item {
	res := make([]map[string]gocache.Item, len(sc.caches))
	for i, v := range sc.caches {
		res[i] = v.Items()
	}
	return res
}

func (sc *shardedCache) Keys() [][]string {
	res := make([][]string, len(sc.caches))
	for i, v := range sc.caches {
		items := v.Items()
		keys := make([]string, 0, len(items))
		for k := range items {
			keys = append(keys, k)
		}

		res[i] = keys
	}

	return res
}

func (sc *shardedCache) Flush() {
	for _, v := range sc.caches {
		v.Flush()
	}
}

type shardedJanitor struct {
	Interval time.Duration
	stop     chan bool
}

func (j *shardedJanitor) Run(sc *shardedCache) {
	tick := time.NewTicker(j.Interval)
	for {
		select {
		case <-tick.C:
			sc.DeleteExpired()
		case <-j.stop:
			return
		}
	}
}

func stopShardedJanitor(sc *ShardedCache) {
	sc.janitor.stop <- true
}

func newShardedJanitor(sc *shardedCache, ci time.Duration) *shardedJanitor {
	j := &shardedJanitor{
		Interval: ci,
		stop:     make(chan bool),
	}
	go j.Run(sc)

	return j
}

func newShardedCache(segmentSize2N int, defaultExpire time.Duration, initItemSize int) *shardedCache {
	if initItemSize < 0 {
		initItemSize = 0
	}
	if defaultExpire <= 0 {
		defaultExpire = gocache.NoExpiration
	}
	cs := &shardedCache{
		segmentSize2N: segmentSize2N,
		segmentMark:   uint64(segmentSize2N) - 1,
		defaultExpire: defaultExpire,
		initItemSize:  initItemSize,
	}

	cs.caches = make([]*gocache.Cache, 0, segmentSize2N)
	for i := 0; i < segmentSize2N; i++ {
		items := make(map[string]gocache.Item, initItemSize)
		cs.caches = append(cs.caches, gocache.NewFrom(defaultExpire, 0, items))
	}

	return cs
}

func NewShardedCache(segmentSize2N int, defaultExpire, cleanUp time.Duration, initItemSize int) *ShardedCache {
	sc := newShardedCache(segmentSize2N, defaultExpire, initItemSize)
	ret := &ShardedCache{shardedCache: sc}

	if cleanUp <= 0 {
		cleanUp = DefaultCleanUp
	}
	ret.janitor = newShardedJanitor(sc, cleanUp)
	runtime.SetFinalizer(ret, stopShardedJanitor)

	return ret
}
