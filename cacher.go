package cache

import (
	"context"
	"fmt"
	"time"
)

type Cacher[T any] struct {
	key     string
	handler func() (*T, bool)
	get     func(ctx context.Context, key string) (*T, bool)
	set     func(ctx context.Context, key string, val *T, d time.Duration) error
}

func (c *Cacher[T]) Value(ctx context.Context) (*T, bool) {
	data, ok := c.get(ctx, c.key)
	if ok {
		return data, ok
	}

	data, ok = c.handler()
	if !ok {
		return data, ok
	}

	go func() {
		err := c.set(context.TODO(), c.key, data, 0)
		if err != nil {
			// TODO
			fmt.Println(err)
		}
	}()

	return data, ok
}

func (c *Cacher[T]) Get(ctx context.Context, key string) (*T, bool) {
	data, ok := c.get(ctx, c.key)
	if ok {
		return data, ok
	}

	data, ok = c.handler()
	if !ok {
		return data, ok
	}

	go func() {
		err := c.set(context.TODO(), c.key, data, 0)
		if err != nil {
			// TODO
			fmt.Println(err)
		}
	}()

	return data, ok
}

func (c *Cacher[T]) Set(ctx context.Context, key string, value *T, d time.Duration) {

}
