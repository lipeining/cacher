package canal

import (
	"sync"

	"github.com/go-mysql-org/go-mysql/canal"
)

type Canal struct {
	*canal.Canal
	cfg     *canal.Config
	handler *CanalEventHandler
	mu      sync.Mutex
}

func NewCanal(name string, opts ...Option) *Canal {
	cfg, err := canal.NewConfigWithFile(name)
	if err != nil {
		panic(err)
	}

	for _, opt := range opts {
		opt(cfg)
	}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}

	handler := NewCanalEventHandler()
	c.SetEventHandler(handler)

	return &Canal{Canal: c, cfg: cfg, handler: handler}
}

func (c *Canal) Start() error {
	// should we check all listeners are set?
	return c.Canal.Run()
}

func (c *Canal) ReStart(where string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cfg.Dump.Where = where

	newCanal, err := canal.NewCanal(c.cfg)
	if err != nil {
		return err
	}

	newCanal.SetEventHandler(c.handler)
	c.Canal = newCanal
	return c.Start()
}
