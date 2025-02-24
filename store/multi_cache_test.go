package multicache

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	gocache "github.com/patrickmn/go-cache"
)

func TestMultiCache(t *testing.T) {
	memoryCache := NewMemoryCache[StoreInfo](gocache.New(time.Hour, time.Minute*5))
	redisCache := NewRedisCache[StoreInfo](redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}), NewJSONSerializer[StoreInfo](), NewStringStore[StoreInfo])

	c := NewMultiCache[StoreInfo](memoryCache, redisCache)

	store := NewStringStore[StoreInfo](NewJSONSerializer[StoreInfo]())
	store.From("m", `{"storeId":"m","handle":"m"}`)

	if err := c.Set(context.Background(), "m", store, time.Hour); err != nil {
		t.Fatal(err)
	}

	cCache, err := c.Get(context.Background(), "m")
	if err != nil {
		t.Fatal(err)
	}

	s, _ := cCache.To()
	if s.StoreId != "m" {
		t.Fatal(cCache)
	}

	// del memory cache, will get from redis
	if _, err := memoryCache.Del(context.Background(), "m"); err != nil {
		t.Fatal(err)
	}

	cCache, err = c.Get(context.Background(), "m")
	if err != nil {
		t.Fatal(err)
	}

	s, _ = cCache.To()
	if s.StoreId != "m" {
		t.Fatal(cCache)
	}
}
