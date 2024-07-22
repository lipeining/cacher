package cache

import (
	"context"
	"time"
)

func NewMultiCache[T any](cacher ...internalCacher[T]) *MultiCache[T] {
	if len(cacher) == 0 {
		panic("no cacher")
	}

	return &MultiCache[T]{Cacher: cacher}
}

type MultiCache[T any] struct {
	Cacher []internalCacher[T]
}

type writeBackItem struct {
	index int
	err   error
}

func (m *MultiCache[T]) Get(ctx context.Context, key string) (*T, error) {
	writeBackList := make([]writeBackItem, 0)

	for index, cacher := range m.Cacher {
		v, err := cacher.Get(ctx, key)
		if err == nil && v != nil {
			// 获取到 cache，根据实际回写
			if len(writeBackList) != 0 {
				for _, writeBack := range writeBackList {
					go func(pos int, err error) {
						m.Cacher[pos].WriteBack(context.Background(), key, v, err)
					}(writeBack.index, writeBack.err)
				}
			}

			return v, nil
		} else {
			writeBackList = append(writeBackList, writeBackItem{index: index, err: err})
		}
	}

	// 没有值
	return nil, nil
}

func (m *MultiCache[T]) Del(ctx context.Context, key string) error {
	errors := make([]error, 0)

	for _, cacher := range m.Cacher {
		_, err := cacher.Del(ctx, key)
		if err != nil {
			errors = append(errors, err)
		}
	}

	// TODO 合并 errors
	return errors[0]
}

func (m *MultiCache[T]) Set(ctx context.Context, key string, value *T, d time.Duration) error {
	errors := make([]error, 0)

	for _, cacher := range m.Cacher {
		err := cacher.Set(ctx, key, value, d)
		if err != nil {
			errors = append(errors, err)
		}
	}

	// TODO 合并 errors
	return errors[0]
}
