package multicache

import (
	"context"
	"testing"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type StoreInfo struct {
	StoreId string `json:"storeId"`
	Handle  string `json:"handle"`
}

func TestMemoryCache(t *testing.T) {
	c := NewMemoryCache[StoreInfo](gocache.New(time.Hour, time.Minute*5))

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

func TestCounterCache(t *testing.T) {
	c := NewMemoryCache[int](gocache.New(time.Hour, time.Minute*5))
	storeCounter := NewStringStore[int](NewIntSerializer())
	storeCounter.From("c", `1`)

	if err := c.Set(context.Background(), "c", storeCounter, time.Hour); err != nil {
		t.Fatal(err)
	}

	if cCache, err := c.Get(context.Background(), "c"); err != nil {
		t.Fatal(err)
	} else if s, _ := cCache.To(); *s != 1 {
		t.Fatal(cCache)
	}
}

func TestLazyCache(t *testing.T) {
	c := NewMemoryCache[StoreInfo](gocache.New(time.Hour, time.Minute*5))
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
