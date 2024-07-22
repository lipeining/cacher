package cache

import (
	"context"
	"fmt"
	"time"
)

type BatchCacher[T any] struct {
	handler func(ctx context.Context, keys []string) (map[string]*T, error)
	get     func(ctx context.Context, key string) (*T, bool)
	set     func(ctx context.Context, key string, val *T, d time.Duration) error
}

func (c *BatchCacher[T]) Value(ctx context.Context, keys []string) (map[string]*T, bool) {
	mapper := make(map[string]*T)
	miss := make([]string, 0)

	for _, key := range keys {
		data, ok := c.get(ctx, key)
		if ok {
			mapper[key] = data
		} else {
			miss = append(miss, key)
		}
	}

	if len(miss) == 0 {
		return mapper, true
	}

	list, err := c.handler(ctx, miss)
	if err != nil {
		// TODO
		return mapper, false
	}

	for key, val := range list {
		mapper[key] = val
	}

	go func() {
		for key, val := range list {
			err := c.set(context.TODO(), key, val, 0)
			if err != nil {
				// TODO
				fmt.Println(err)
			}
		}

	}()

	return mapper, true
}
