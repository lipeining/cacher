package cache

import (
	"context"
	"time"
)

type internalCacher[T any] interface {
	Get(ctx context.Context, key string) (*T, error)
	TTL(ctx context.Context, key string) (time.Duration, error)
	Set(ctx context.Context, key string, value *T, d time.Duration) error
	Del(ctx context.Context, key string) (bool, error)
	WriteBack(ctx context.Context, key string, value *T, err error)
}
