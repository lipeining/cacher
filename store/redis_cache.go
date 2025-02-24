package multicache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCacher[T any] struct {
	Cache      redis.Cmdable
	Serializer Serializer[T]
	NewStore   NewStoreFunc[T]
}

func (r *RedisCacher[T]) TTL(ctx context.Context, key string) (time.Duration, error) {
	value, err := r.Cache.TTL(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return -1, nil
		}
		return -2, err
	}

	return value, err
}

func (r *RedisCacher[T]) Del(ctx context.Context, key string) (bool, error) {
	num, err := r.Cache.Del(ctx, key).Result()

	if err != nil {
		return num > 0, err
	}

	return num > 0, err
}

func (r *RedisCacher[T]) Get(ctx context.Context, key string) (Store[T], error) {
	value, err := r.Cache.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	store := r.NewStore(r.Serializer)
	err = store.From(key, value)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (r *RedisCacher[T]) Set(ctx context.Context, key string, value Store[T], d time.Duration) error {
	obj, err := value.To()
	if err != nil {
		return err
	}

	v, err := r.Serializer.Marshal(obj)
	if err != nil {
		return err
	}

	_, err = r.Cache.Set(ctx, key, v, d).Result()

	return err
}

func (r *RedisCacher[T]) WriteBack(ctx context.Context, key string, value Store[T], err error) {
	if err != nil {
		// 有错误时，根据错误类型回写
		// TODO
	} else {
		// TODO 如何确定 expire 时间
		setErr := r.Set(ctx, key, value, 0)
		if setErr != nil {
			// TODO
		}
	}
}

func NewRedisCache[T any](client redis.Cmdable, serializer Serializer[T], newStore NewStoreFunc[T]) internalCacher[T] {
	cacher := RedisCacher[T]{Cache: client, Serializer: serializer, NewStore: newStore}
	return &cacher
}
