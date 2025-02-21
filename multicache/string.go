package multicache

import (
	"context"

	"github.com/lipeining/cache/localcache"
	"github.com/lipeining/cache/rdbcache"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
)

type StringMultiCache[T any] struct {
	rdbCache   rdbcache.RedisCache
	serializer Serializer[T]
	newStore   func(Serializer[T]) Store[T]
	storeCache *localcache.LocalCache[Store[T]]
	sf         singleflight.Group
	loader     *Loader
	logger     *log.Helper

	scanRule        string
	localCacheClose bool
}

func NewStringMultiCache[T any](rdbCache rdbcache.RedisCache, localCache *localcache.ShardedCache, config *Config, serializer Serializer[T], newStore func(Serializer[T]) Store[T], logger log.Logger) *StringMultiCache[T] {
	m := &StringMultiCache[T]{
		scanRule:        config.ScanRule,
		localCacheClose: config.LocalCacheClose,
		rdbCache:        rdbCache,
		serializer:      serializer,
		newStore:        newStore,
		storeCache:      localcache.NewLocalCache[Store[T]](localCache, &localcache.Config{ReloadTickerSec: config.ReloadTickerSec}, logger),
		logger:          log.NewHelper(logger),
	}
	m.loader = NewLoader(m.getKeys, m.fetchAndCacheKey, config, logger)
	return m
}

func (m *StringMultiCache[T]) StartLoad() {
	if !m.localCacheClose {
		m.loader.StartLoad()
	}
}

func (m *StringMultiCache[T]) getKeys(ctx context.Context) ([]string, error) {
	allKeys, err := m.rdbCache.Scan(ctx, m.scanRule, RedisKeyTypeString)
	if err != nil {
		m.logger.Errorf("scan redis duffle's all keys fail, scan_rule: %s, scan_type: %s, err: %v", m.scanRule, RedisKeyTypeString, err)
		return nil, err
	}
	return allKeys, nil
}

func (m *StringMultiCache[T]) fetchAndCacheKey(ctx context.Context, key string) error {
	data, err := m.rdbCache.Get(ctx, key)
	if err != nil {
		m.logger.Errorf("load redis data fail, key: %s, err: %v", key, err)
		return err
	}

	_, err = m.setToLocalCache(key, data)
	if err != nil {
		m.logger.Errorf("set to local cache fail, key: %s, err: %v", key, err)
		return err
	}

	return nil
}

func (m *StringMultiCache[T]) setToLocalCache(key string, data any) (Store[T], error) {
	store := m.newStore(m.serializer)
	err := store.From(key, data)
	if err != nil {
		return nil, err
	}
	if !m.localCacheClose {
		m.storeCache.Set(&store)
	}
	return store, nil
}

func (m *StringMultiCache[T]) Get(ctx context.Context, key string) (*T, error) {
	if !m.localCacheClose {
		val, err := m.storeCache.Get(key)
		if err != nil {
			return nil, err
		}
		if val != nil {
			return (*val).To()
		}

		m.logger.Warnf("multi-cache key miss, key: %s", key)
	}

	ret, err, _ := m.sf.Do(key, func() (interface{}, error) {
		return m.rdbCache.Get(ctx, key)
	})
	if err != nil {
		return nil, err
	}

	store, err := m.setToLocalCache(key, ret)
	if err != nil {
		m.logger.Errorf("set to local cache fail, key: %s, err: %v", key, err)
		return nil, err
	}

	return store.To()
}
