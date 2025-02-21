package multicache

import (
	"context"
	"fmt"

	"github.com/lipeining/cache/localcache"
	"github.com/lipeining/cache/rdbcache"

	"github.com/go-kratos/kratos/v2/log"
	gocache "github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
)

type HashMultiCache[T any] struct {
	rdbCache   rdbcache.RedisCache
	localCache *localcache.ShardedCache
	serializer Serializer[T]
	newStore   NewHashStoreFunc[T]
	sf         singleflight.Group
	loader     *Loader
	logger     *log.Helper

	scanRule        string
	localCacheClose bool
}

func NewHashMultiCache[T any](rdbCache rdbcache.RedisCache, localCache *localcache.ShardedCache, config *Config, serializer Serializer[T], newStore NewHashStoreFunc[T], logger log.Logger) *HashMultiCache[T] {
	m := &HashMultiCache[T]{
		scanRule:        config.ScanRule,
		localCacheClose: config.LocalCacheClose,
		rdbCache:        rdbCache,
		localCache:      localCache,
		serializer:      serializer,
		newStore:        newStore,
		logger:          log.NewHelper(logger),
	}
	m.loader = NewLoader(m.getKeys, m.fetchAndCacheKey, config, logger)
	return m
}

func (m *HashMultiCache[T]) StartLoad() {
	if !m.localCacheClose {
		m.loader.StartLoad()
	}
}

func (m *HashMultiCache[T]) getKeys(ctx context.Context) ([]string, error) {
	allKeys, err := m.rdbCache.Scan(ctx, m.scanRule, RedisKeyTypeHash)
	if err != nil {
		m.logger.Errorf("[HashMultiCache] [GetKeys] scan redis duffle's all keys fail, scan_rule: %s, scan_type: %s, err: %v", m.scanRule, RedisKeyTypeHash, err)
		return nil, err
	}
	return allKeys, nil
}

func (m *HashMultiCache[T]) fetchAndCacheKey(ctx context.Context, key string) error {
	data, err := m.rdbCache.HGetAllByBatch(ctx, key)
	if err != nil {
		m.logger.Errorf("[HashMultiCache] [FetchAndCacheKey] get redis duffle's all fields fail, key: %s, err: %v", key, err)
		return err
	}

	_, err = m.setToLocalCache(key, data)
	if err != nil {
		m.logger.Errorf("[HashMultiCache] [FetchAndCacheKey] set to local cache fail, key: %s, err: %v", key, err)
		return err
	}

	return nil
}

func (m *HashMultiCache[T]) setToLocalCache(key string, data any) (HashStoreI[T], error) {
	store := m.newStore(m.serializer)
	err := store.From(data)
	if err != nil {
		return nil, err
	}
	if !m.localCacheClose {
		m.localCache.Set(key, store, gocache.NoExpiration)
	}
	return store, nil
}

func (m *HashMultiCache[T]) getFromLocalCache(key string) (map[string]*T, error) {
	val, ok := m.localCache.Get(key)
	if !ok {
		return nil, nil
	}

	ret, ok := val.(HashStoreI[T])
	if !ok {
		return nil, fmt.Errorf("key %s is not of type HashStoreI[T]", key)
	}
	return ret.To()
}

func (m *HashMultiCache[T]) GetMap(ctx context.Context, key string) (map[string]*T, error) {
	if !m.localCacheClose {
		result, err := m.getFromLocalCache(key)
		if err != nil {
			return nil, err
		}
		if result != nil {
			return result, nil
		}
		m.logger.Warnf("[HashMultiCache] [GetMap] get from local cache key miss, key: %s", key)
	}

	// penetrate local cache, get data from redis
	ret, err, _ := m.sf.Do(key, func() (interface{}, error) {
		return m.rdbCache.HGetAll(ctx, key)
	})
	if err != nil {
		return nil, err
	}

	store, err := m.setToLocalCache(key, ret)
	if err != nil {
		m.logger.Errorf("[HashMultiCache] [GetMap] set to local cache fail, key: %s, err: %v", key, err)
		return nil, err
	}

	return store.To()
}

func (m *HashMultiCache[T]) GetMapMul(ctx context.Context, key string, fields ...string) ([]*T, error) {
	if !m.localCacheClose {
		result, err := m.getFromLocalCache(key)
		if err != nil {
			return nil, err
		}
		if result != nil {
			ret := make([]*T, 0, len(fields))
			for _, f := range fields {
				if v, ok := result[f]; ok {
					ret = append(ret, v)
				} else {
					var tmp *T
					ret = append(ret, tmp)
				}
			}
			return ret, nil
		}
		m.logger.Warnf("[HashMultiCache] [GetMapMul] get from local cache key miss, key: %s", key)
	}

	ret, err, _ := m.sf.Do(key, func() (interface{}, error) {
		return m.rdbCache.HMGet(ctx, key, fields...)
	})
	if err != nil {
		return nil, err
	}

	object := make([]*T, 0, len(fields))
	if retMap, ok := ret.([]interface{}); ok {
		for i, value := range retMap {
			var v T

			if value == nil {
				object = append(object, &v)
				continue
			}
			data, ok := value.(string)
			if !ok {
				object = append(object, &v)
				continue
			}

			if err := m.serializer.Unmarshal([]byte(data), &v); err != nil {
				m.logger.Errorf("[HashMultiCache] unmarshal redis data fail, key: %s, field: %s, err: %v", key, fields[i], err)
				return nil, err
			}
			object = append(object, &v)
		}
	}

	return object, nil
}

func (m *HashMultiCache[T]) UpdateHashField(ctx context.Context, key string, field string) error {
	if m.localCacheClose {
		return nil
	}

	result, err := m.getFromLocalCache(key)
	if err != nil {
		return err
	}
	if result == nil {
		return errors.Errorf("[HashMultiCache] key miss, key: %s", key)
	}

	ret, err := m.rdbCache.HMGet(ctx, key, field)
	if err != nil {
		m.logger.Errorf("[HashMultiCache] [UpdateHashField] get redis data fail, key: %s, field: %s, err: %v", key, field, err)
		return err
	}

	for _, value := range ret {
		if value == nil {
			m.logger.Warnf("[HashMultiCache] [UpdateHashField] got nil value, key: %s, field: %s", key, field)
			continue
		}

		data, ok := value.(string)
		if !ok {
			m.logger.Warnf("[HashMultiCache] [UpdateHashField] got not string value, key: %s, field: %s", key, field)
			continue
		}

		var v T
		if err := m.serializer.Unmarshal([]byte(data), &v); err != nil {
			m.logger.Errorf("[HashMultiCache] [UpdateHashField] unmarshal redis data fail, key: %s, field: %s, err: %v", key, field, err)
			return err
		}
		result[field] = &v
	}

	m.logger.Infof("[HashMultiCache] [UpdateHashField] update hash field success, key: %s, field: %s", key, field)
	return nil
}

func (m *HashMultiCache[T]) DeleteHashField(ctx context.Context, key string, field string) error {
	if m.localCacheClose {
		return nil
	}

	result, err := m.getFromLocalCache(key)
	if err != nil {
		return err
	}
	if result == nil {
		m.logger.Warnf("[HashMultiCache] [DeleteHashField] hash key miss, key: %s", key)
		return nil
	}

	delete(result, field)
	m.logger.Infof("[HashMultiCache] [DeleteHashField] delete hash field success, key: %s, field: %s", key, field)
	return nil
}

func (m *HashMultiCache[T]) UpdatePrefix(ctx context.Context, prefix string) error {
	// 扫描所有需要拉取的keys
	allKeys, err := m.rdbCache.Scan(ctx, prefix, RedisKeyTypeHash)
	if err != nil {
		m.logger.Errorf("[HashMultiCache] [UpdatePrefix] scan redis duffle's all keys fail, scan_rule: %s, scan_type: %s, err: %v", prefix, RedisKeyTypeHash, err)
		return err
	}

	for _, key := range allKeys {
		if err := m.fetchAndCacheKey(ctx, key); err != nil {
			m.logger.Errorf("[HashMultiCache] [UpdatePrefix] fetch and cache key fail, key: %s, err: %v", key, err)
			return err
		}
	}

	m.logger.Infof("[HashMultiCache] [UpdatePrefix] update data success, prefix: %s", prefix)
	return nil
}
