package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCacher[T any] struct {
	Cache      redis.Cmdable
	Serializer Serializer[T]
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

func (r *RedisCacher[T]) Get(ctx context.Context, key string) (*T, error) {
	var a T

	value, err := r.Cache.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return &a, nil
		}
		return &a, err
	}

	err = r.Serializer.Unmarshal([]byte(value), &a)
	return &a, err
}

func (r *RedisCacher[T]) Set(ctx context.Context, key string, value *T, d time.Duration) error {
	v, err := r.Serializer.Marshal(value)
	if err != nil {
		return err
	}

	_, err = r.Cache.Set(ctx, key, v, d).Result()

	return err
}

func (r *RedisCacher[T]) WriteBack(ctx context.Context, key string, value *T, err error) {
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

func redisHolder[T any](client redis.Cmdable) internalCacher[T] {
	cacher := RedisCacher[T]{Cache: client, Serializer: NewJSONSerializer[T]()}
	return &cacher
}
