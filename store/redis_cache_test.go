package multicache

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestRedisCache(t *testing.T) {
	c := NewRedisCache[StoreInfo](redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}), NewJSONSerializer[StoreInfo](), NewStringStore[StoreInfo])

	aStore := StoreInfo{StoreId: "a", Handle: "a"}
	bStore := StoreInfo{StoreId: "b", Handle: "b"}

	storeA := NewStringStore[StoreInfo](NewJSONSerializer[StoreInfo]())
	storeA.From("a", `{"storeId":"a","handle":"a"}`)

	if err := c.Set(context.Background(), "a", storeA, time.Hour); err != nil {
		t.Fatal(err)
	}

	storeB := NewStringStore[StoreInfo](NewJSONSerializer[StoreInfo]())
	storeB.From("b", `{"storeId":"b","handle":"b"}`)

	if err := c.Set(context.Background(), "b", storeB, time.Hour); err != nil {
		t.Fatal(err)
	}

	if aCache, err := c.Get(context.Background(), "a"); err != nil {
		t.Fatal(err)
	} else if s, _ := aCache.To(); s.StoreId != aStore.StoreId {
		t.Fatal(aCache, aStore)
	}

	if bCache, err := c.Get(context.Background(), "b"); err != nil {
		t.Fatal(err)
	} else if s, _ := bCache.To(); s.StoreId != bStore.StoreId {
		t.Fatal(bCache, bStore)
	}
}

func TestRedisCounterCache(t *testing.T) {
	c := NewRedisCache[int](redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}), NewIntSerializer(), NewStringStore[int])
	storeCounter := NewStringStore[int](NewIntSerializer())
	storeCounter.From("counter", `1`)

	if err := c.Set(context.Background(), "counter", storeCounter, time.Hour); err != nil {
		t.Fatal(err)
	}

	if cCache, err := c.Get(context.Background(), "counter"); err != nil {
		t.Fatal(err)
	} else if s, _ := cCache.To(); *s != 1 {
		t.Fatal(cCache)
	}
}

func TestRedisLazyCache(t *testing.T) {
	c := NewRedisCache[StoreInfo](redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}), NewJSONSerializer[StoreInfo](), NewStringLazyStore[StoreInfo])

	store := NewStringLazyStore[StoreInfo](NewJSONSerializer[StoreInfo]())
	store.From("c", `{"storeId":"c","handle":"c"}`)

	if err := c.Set(context.Background(), "c", store, time.Hour); err != nil {
		t.Fatal(err)
	}

	if cCache, err := c.Get(context.Background(), "c"); err != nil {
		t.Fatal(err)
	} else if s, _ := cCache.To(); s.StoreId != "c" {
		t.Fatal(cCache)
	}
}
