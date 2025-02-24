package multicache

import (
	"errors"
	"sync"
)

type CacheKeyType interface {
	CacheKey() string
}

type Store[T any] interface {
	From(string, any) error
	To() (*T, error)
	CacheKey() string
}

type NewStoreFunc[T any] func(Serializer[T]) Store[T]

type StringStore[T any] struct {
	key        string
	value      T
	serializer Serializer[T]
}

func (s *StringStore[T]) From(key string, data any) error {
	var v T
	err := s.serializer.Unmarshal([]byte(data.(string)), &v)
	if err != nil {
		return err
	}
	s.key = key
	s.value = v
	return nil
}

func (s *StringStore[T]) To() (*T, error) {
	return &s.value, nil
}

func (s *StringStore[T]) CacheKey() string {
	if v, ok := any(s.value).(CacheKeyType); ok {
		return v.CacheKey()
	}
	return s.key
}

func NewStringStore[T any](serializer Serializer[T]) Store[T] {
	return &StringStore[T]{
		serializer: serializer,
	}
}

type LazyStore[T any] struct {
	key   string
	data  any
	store Store[T]
	once  sync.Once
}

func (s *LazyStore[T]) From(key string, data any) error {
	s.data = data
	s.key = key
	return nil
}

func (s *LazyStore[T]) To() (*T, error) {
	s.once.Do(func() {
		s.store.From(s.key, s.data)
		s.data = nil
		s.key = ""
	})
	return s.store.To()
}

func (s *LazyStore[T]) CacheKey() string {
	return s.store.CacheKey()
}

func NewStringLazyStore[T any](serializer Serializer[T]) Store[T] {
	return &LazyStore[T]{
		store: NewStringStore[T](serializer),
	}
}

type HashStoreI[T any] interface {
	From(any) error
	To() (map[string]*T, error)
}

type NewHashStoreFunc[T any] func(Serializer[T]) HashStoreI[T]

type HashStore[T any] struct {
	value      map[string]*T
	serializer Serializer[T]
}

func (s *HashStore[T]) From(input any) error {
	data, ok := input.(map[string]string)
	if !ok {
		return errors.New("input is not a map[string]string")
	}

	object := make(map[string]*T)
	for field, value := range data {
		var v T
		if err := s.serializer.Unmarshal([]byte(value), &v); err != nil {
			return err
		}
		object[field] = &v
	}
	s.value = object
	return nil
}

func (s *HashStore[T]) To() (map[string]*T, error) {
	return s.value, nil
}

func NewHashStore[T any](serializer Serializer[T]) HashStoreI[T] {
	return &HashStore[T]{
		serializer: serializer,
	}
}

type HashLazyStore[T any] struct {
	data  any
	store HashStoreI[T]
	once  sync.Once
}

func (s *HashLazyStore[T]) From(data any) error {
	s.data = data
	return nil
}

func (s *HashLazyStore[T]) To() (map[string]*T, error) {
	s.once.Do(func() {
		s.store.From(s.data)
		s.data = nil
	})
	return s.store.To()
}

func NewHashLazyStore[T any](serializer Serializer[T]) HashStoreI[T] {
	return &HashLazyStore[T]{
		store: NewHashStore[T](serializer),
	}
}
