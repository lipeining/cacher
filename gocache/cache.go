package gocache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	ristretto_store "github.com/eko/gocache/store/ristretto/v4"
	redis "github.com/redis/go-redis/v9"
)

type HashCache[T any] struct {
	cache *cache.ChainCache[map[string]T]
}

func NewHashCache[T any](localStore store.StoreInterface, redisStore store.StoreInterface) *HashCache[T] {
	return &HashCache[T]{
		cache: cache.NewChain[map[string]T](
			cache.New[map[string]T](localStore),
			cache.New[map[string]T](redisStore),
		),
	}
}

func (c *HashCache[T]) Get(ctx context.Context, key any) (map[string]T, error) {
	return c.cache.Get(ctx, key)
}

func (c *HashCache[T]) Set(ctx context.Context, key any, value map[string]T) error {
	return c.cache.Set(ctx, key, value)
}

func (c *HashCache[T]) Delete(ctx context.Context, key any) error {
	return c.cache.Delete(ctx, key)
}

func (c *HashCache[T]) Invalidate(ctx context.Context, options ...store.InvalidateOption) error {
	return c.cache.Invalidate(ctx, options...)
}

func (c *HashCache[T]) Clear(ctx context.Context) error {
	return c.cache.Clear(ctx)
}

func (c *HashCache[T]) GetType() string {
	return "HashCache"
}

func NewRistrettoStore(cache *ristretto.Cache) store.StoreInterface {
	return ristretto_store.NewRistretto(cache)
}

// RedisClientInterface represents a go-redis/redis client
type RedisClientInterface interface {
	HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd
	HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	FlushAll(ctx context.Context) *redis.StatusCmd
}

type RedisStore[T any] struct {
	client RedisClientInterface
}

func NewRedisStore[T any](client *redis.Client) store.StoreInterface {
	return &RedisStore[T]{client: client}
}

// return map[string]string
func (s *RedisStore[T]) Get(ctx context.Context, key any) (any, error) {
	val, err := s.client.HGetAll(ctx, key.(string)).Result()
	if err != nil {
		return nil, err
	}

	// convert to map[string]T
	object := make(map[string]T)
	for k, v := range val {
		var value T
		if err := json.Unmarshal([]byte(v), &value); err != nil {
			return nil, err
		}
		object[k] = value
	}

	return object, nil
}

// return map[string]string
func (s *RedisStore[T]) GetWithTTL(ctx context.Context, key any) (any, time.Duration, error) {
	val, err := s.Get(ctx, key)
	return val, 0, err
}

// set value of map[string]T
func (s *RedisStore[T]) Set(ctx context.Context, key any, value any, options ...store.Option) error {
	objects := make(map[string]string)
	for k, v := range value.(map[string]T) {
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		objects[k] = string(data)
	}

	return s.client.HMSet(ctx, key.(string), objects).Err()
}

func (s *RedisStore[T]) Delete(ctx context.Context, key any) error {
	return s.client.Del(ctx, key.(string)).Err()
}

func (s *RedisStore[T]) Invalidate(ctx context.Context, options ...store.InvalidateOption) error {
	return nil
}

func (s *RedisStore[T]) Clear(ctx context.Context) error {
	return s.client.FlushAll(ctx).Err()
}

func (s *RedisStore[T]) GetType() string {
	return "RedisStore"
}

type TestObject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func my_func() {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{NumCounters: 1000, MaxCost: 100, BufferItems: 64})
	if err != nil {
		panic(err)
	}
	localStore := NewRistrettoStore(ristrettoCache)

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	redisStore := NewRedisStore[TestObject](redisClient)

	hashCache := NewHashCache[TestObject](localStore, redisStore)

	hashCache.Set(context.Background(), "test", map[string]TestObject{
		"test": {ID: "1", Name: "test"},
	})

	val, err := hashCache.Get(context.Background(), "test")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
