package multicache

import (
	"context"
	"errors"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type MemoryCacher[T any] struct {
	Cache *gocache.Cache
}

func (m *MemoryCacher[T]) TTL(ctx context.Context, key string) (time.Duration, error) {
	_, ttl, ok := m.Cache.GetWithExpiration(key)
	if !ok {
		return -1, nil
	}

	return time.Until(ttl), nil
}

func (m *MemoryCacher[T]) Del(ctx context.Context, key string) (bool, error) {
	m.Cache.Delete(key)
	return true, nil
}

func (m *MemoryCacher[T]) Get(ctx context.Context, key string) (Store[T], error) {
	value, ok := m.Cache.Get(key)
	if !ok {
		return nil, errors.New("not found")
	}

	ret, ok := value.(Store[T])
	if !ok {
		return nil, errors.New("invalid value")
	}

	return ret, nil
}

func (m *MemoryCacher[T]) Set(ctx context.Context, key string, value Store[T], d time.Duration) error {
	m.Cache.Set(key, value, d)
	return nil
}

func (m *MemoryCacher[T]) WriteBack(ctx context.Context, key string, value Store[T], err error) {
	if err != nil {
		// 有错误时，根据错误类型回写
		// TODO
	} else {
		// TODO 如何确定 expire 时间
		setErr := m.Set(ctx, key, value, 0)
		if setErr != nil {
			// TODO
		}
	}
}

func NewMemoryCache[T any](cache *gocache.Cache) internalCacher[T] {
	cacher := MemoryCacher[T]{Cache: cache}
	return &cacher
}
