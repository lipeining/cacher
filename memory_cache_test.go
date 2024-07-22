package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/coocood/freecache"
)

func TestMemoryCache(t *testing.T) {
	type StoreInfo struct {
		StoreId string `json:"storeId"`
		Handle  string `json:"handle"`
	}

	storeInfoMemoryCacher := memoryHolder(freecache.NewCache(1000), NewJSONSerializer[StoreInfo]())

	aStore := StoreInfo{StoreId: "a", Handle: "a"}
	bStore := StoreInfo{StoreId: "b", Handle: "b"}

	if err := storeInfoMemoryCacher.Set(context.Background(), "a", &aStore, time.Hour); err != nil {
		t.Fatal(err)
	}

	if err := storeInfoMemoryCacher.Set(context.Background(), "b", &bStore, time.Hour); err != nil {
		t.Fatal(err)
	}

	if aCache, err := storeInfoMemoryCacher.Get(context.Background(), "a"); err != nil {
		t.Fatal(err)
	} else if aCache.StoreId != aStore.StoreId {
		t.Fatal(aCache, aStore)
	}

	if bCache, err := storeInfoMemoryCacher.Get(context.Background(), "b"); err != nil {
		t.Fatal(err)
	} else if bCache.StoreId != bStore.StoreId {
		t.Fatal(bCache, bStore)
	}
}

func TestCounterCache(t *testing.T) {
	counterMemoryCacher := memoryHolder(freecache.NewCache(100000), NewJSONSerializer[int32]())
	counter := int32(1)
	if err := counterMemoryCacher.Set(context.Background(), "c", &counter, time.Hour); err != nil {
		t.Fatal(err)
	}

	if cCache, err := counterMemoryCacher.Get(context.Background(), "c"); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(*cCache)
	}
}
