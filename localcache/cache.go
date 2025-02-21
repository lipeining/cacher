package localcache

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	gocache "github.com/patrickmn/go-cache"
)

type LocalCacheType interface {
	CacheKey() string
}

type LoadAllFunc[T LocalCacheType] func() ([]*T, error)
type LoadUpdateFunc[T LocalCacheType] func(lastLoadTime time.Time) ([]*T, error)

type Config struct {
	ReloadTickerSec int
}

// LocalCache 本地缓存：支持泛型，支持全量加载，定时加载机制
type LocalCache[T LocalCacheType] struct {
	cache  *ShardedCache
	logger *log.Helper

	loadAll    LoadAllFunc[T]
	lastLoad   time.Time
	loadUpdate LoadUpdateFunc[T]
	config     *Config
}

func NewLocalCache[T LocalCacheType](cache *ShardedCache, config *Config, logger log.Logger) *LocalCache[T] {
	return &LocalCache[T]{cache: cache, logger: log.NewHelper(logger), config: config}
}

func (c *LocalCache[T]) Init(values []*T) error {
	return c.MSet(values)
}

func (c *LocalCache[T]) StartLoad(loadAll LoadAllFunc[T], loadUpdate LoadUpdateFunc[T]) {
	c.loadAll = loadAll
	c.loadUpdate = loadUpdate

	all, err := c.loadAll()
	if err != nil {
		c.logger.Errorf("localcache load all data failed, err: %v", err)
		return
	}
	c.MSet(all)
	c.lastLoad = time.Now()

	go c.startTickLoad()
}

func (c *LocalCache[T]) startTickLoad() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ticker := time.NewTicker(time.Duration(c.config.ReloadTickerSec) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		all, err := c.loadUpdate(c.lastLoad)
		if err != nil {
			c.logger.Errorf("localcache load all data failed, err: %v", err)
			continue
		}
		c.MSet(all)
		c.lastLoad = time.Now()
	}
}

// 获取 key 对应的对象，key 不存在时，返回 nil, nil
func (c *LocalCache[T]) Get(key string) (*T, error) {
	var ret *T
	val, ok := c.cache.Get(key)
	if !ok {
		return nil, nil
	}

	ret, ok = val.(*T)
	if !ok {
		return nil, fmt.Errorf("key %s is not of type %T", key, ret)
	}
	return ret, nil
}

func (c *LocalCache[T]) Set(value *T) error {
	c.cache.Set((*value).CacheKey(), value, gocache.NoExpiration)
	return nil
}

func (c *LocalCache[T]) Delete(key string) error {
	c.cache.Delete(key)
	return nil
}

func (c *LocalCache[T]) MSet(values []*T) error {
	for _, value := range values {
		c.cache.Set((*value).CacheKey(), value, gocache.NoExpiration)
	}
	return nil
}
