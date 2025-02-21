package multicache

import (
	"fmt"
	"testing"
)

type TestObject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (t TestObject) CacheKey() string {
	return fmt.Sprintf("test:%d", t.ID)
}

func TestStringStoreByObject(t *testing.T) {
	serializer := NewJSONSerializer[TestObject]()
	store := NewStringStore[TestObject](serializer)
	err := store.From("test", `{"id": 1, "name": "test"}`)
	if err != nil {
		t.Errorf("failed to create store: %v", err)
	}
	object, err := store.To()
	if err != nil {
		t.Errorf("failed to get object: %v", err)
	}
	if object.ID != 1 || object.Name != "test" {
		t.Errorf("object is not correct: %v", object)
	}
	cacheKey := store.CacheKey()
	if cacheKey != "test:1" {
		t.Errorf("cache key is not correct: %s", cacheKey)
	}
}

func TestStringLazyStoreByObject(t *testing.T) {
	serializer := NewJSONSerializer[TestObject]()
	store := NewStringLazyStore[TestObject](serializer)
	err := store.From("test", `{"id": 1, "name": "test"}`)
	if err != nil {
		t.Errorf("failed to create store: %v", err)
	}
	object, err := store.To()
	if err != nil {
		t.Errorf("failed to get object: %v", err)
	}
	if object.ID != 1 || object.Name != "test" {
		t.Errorf("object is not correct: %v", object)
	}
	cacheKey := store.CacheKey()
	if cacheKey != "test:1" {
		t.Errorf("cache key is not correct: %s", cacheKey)
	}
}

func TestStringByString(t *testing.T) {
	serializer := NewStringSerializer()
	store := NewStringStore[string](serializer)
	err := store.From("test", "test-string")
	if err != nil {
		t.Errorf("failed to create store: %v", err)
	}
	object, err := store.To()
	if err != nil {
		t.Errorf("failed to get object: %v", err)
	}
	if *object != "test-string" {
		t.Errorf("object is not correct: %s", *object)
	}
	cacheKey := store.CacheKey()
	if cacheKey != "test" {
		t.Errorf("cache key is not correct: %s", cacheKey)
	}
}

func TestHashStoreByObject(t *testing.T) {
	serializer := NewJSONSerializer[TestObject]()
	store := NewHashStore[TestObject](serializer)
	err := store.From(map[string]string{"test": `{"id": 1, "name": "test"}`})
	if err != nil {
		t.Errorf("failed to create store: %v", err)
	}
	object, err := store.To()
	if err != nil {
		t.Errorf("failed to get object: %v", err)
	}
	if object["test"].ID != 1 || object["test"].Name != "test" {
		t.Errorf("object is not correct: %v", object)
	}
}

func TestHashLazyStoreByObject(t *testing.T) {
	serializer := NewJSONSerializer[TestObject]()
	store := NewHashLazyStore[TestObject](serializer)
	err := store.From(map[string]string{"test": `{"id": 1, "name": "test"}`})
	if err != nil {
		t.Errorf("failed to create store: %v", err)
	}
	object, err := store.To()
	if err != nil {
		t.Errorf("failed to get object: %v", err)
	}
	if object["test"].ID != 1 || object["test"].Name != "test" {
		t.Errorf("object is not correct: %v", object)
	}
}
