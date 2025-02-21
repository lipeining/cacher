package multicache

import (
	"sync"
)

type LoadFunc[O any, T any] func(O) (*T, error)

type LazyValue[O any, T any] struct {
	item   *T
	raw    O
	load   LoadFunc[O, T]
	mu     sync.RWMutex
	loaded bool // 标记是否已加载
}

func NewLazyValue[O any, T any](raw O, load LoadFunc[O, T]) *LazyValue[O, T] {
	return &LazyValue[O, T]{
		raw:  raw,
		load: load,
	}
}

func (l *LazyValue[O, T]) Get() (*T, error) {
	l.mu.RLock()

	if l.loaded {
		l.mu.RUnlock()
		return l.item, nil
	}

	l.mu.RUnlock()
	l.mu.Lock()
	defer l.mu.Unlock()

	item, err := l.load(l.raw)
	if err != nil {
		return nil, err
	}

	l.item = item
	l.loaded = true
	return item, nil
}

func (l *LazyValue[O, T]) BackgroundLoad() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.loaded {
		return
	}

	item, err := l.load(l.raw)
	if err != nil {
		return
	}

	l.item = item
	l.loaded = true
}
