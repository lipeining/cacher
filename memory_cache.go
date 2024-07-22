package cache

import (
	"context"
	"errors"
	"time"

	"github.com/coocood/freecache"
)

type MemoryCacher[T any] struct {
	Cache      *freecache.Cache
	Serializer Serializer[T]
}

func (m *MemoryCacher[T]) TTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := m.Cache.TTL([]byte(key))
	if err != nil {
		if errors.Is(err, freecache.ErrNotFound) {
			return -1, nil
		}
		return -1, err
	}

	return time.Second * time.Duration(ttl), err
}

func (m *MemoryCacher[T]) Del(ctx context.Context, key string) (bool, error) {
	ok := m.Cache.Del([]byte(key))
	return ok, nil
}

func (m *MemoryCacher[T]) Get(ctx context.Context, key string) (*T, error) {
	var a T
	value, err := m.Cache.Get([]byte(key))
	if err != nil {
		if errors.Is(err, freecache.ErrNotFound) {
			return &a, nil
		}
		return &a, err
	}

	err = m.Serializer.Unmarshal(value, &a)
	return &a, err
}

func (m *MemoryCacher[T]) Set(ctx context.Context, key string, value *T, d time.Duration) error {
	v, err := m.Serializer.Marshal(value)
	if err != nil {
		return err
	}

	err = m.Cache.Set([]byte(key), []byte(v), int(d.Seconds()))
	return err
}

func (m *MemoryCacher[T]) WriteBack(ctx context.Context, key string, value *T, err error) {
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

func memoryHolder[T any](client *freecache.Cache, serializer Serializer[T]) internalCacher[T] {
	cacher := MemoryCacher[T]{Cache: client, Serializer: serializer}
	return &cacher
}
