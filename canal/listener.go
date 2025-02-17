package canal

import (
	"fmt"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/go-mysql-org/go-mysql/schema"
)

type Listener interface {
	Scan(table *schema.Table, rows []any) any
	Insert(row any, header *replication.EventHeader) error
	Update(before any, after any, header *replication.EventHeader) error
	Delete(row any, header *replication.EventHeader) error
}

type TypedListener[T any] interface {
	Scan(table *schema.Table, rows []any) *T
	Insert(row *T, header *replication.EventHeader) error
	Update(before *T, after *T, header *replication.EventHeader) error
	Delete(row *T, header *replication.EventHeader) error
}

type WrapListener[T any] struct {
	TypedListener[T]
}

func NewWrapListener[T any](listener TypedListener[T]) Listener {
	return &WrapListener[T]{TypedListener: listener}
}

func (l *WrapListener[T]) Scan(table *schema.Table, rows []any) any {
	return l.TypedListener.Scan(table, rows)
}

func (l *WrapListener[T]) Insert(row any, header *replication.EventHeader) error {
	rowT, ok := row.(*T)
	if !ok {
		return fmt.Errorf("row is not of type %T", rowT)
	}
	return l.TypedListener.Insert(rowT, header)
}

func (l *WrapListener[T]) Update(before any, after any, header *replication.EventHeader) error {
	beforeT, ok := before.(*T)
	if !ok {
		return fmt.Errorf("before is not of type %T", beforeT)
	}
	afterT, ok := after.(*T)
	if !ok {
		return fmt.Errorf("after is not of type %T", afterT)
	}
	return l.TypedListener.Update(beforeT, afterT, header)
}

func (l *WrapListener[T]) Delete(row any, header *replication.EventHeader) error {
	rowT, ok := row.(*T)
	if !ok {
		return fmt.Errorf("row is not of type %T", rowT)
	}
	return l.TypedListener.Delete(rowT, header)
}

func (c *Canal) AddListener(table string, listener Listener) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.handler.listeners[table]; ok {
		return fmt.Errorf("listener for table %s already exists", table)
	}
	c.handler.listeners[table] = listener
	return nil
}

func (c *Canal) SetListener(table string, listener Listener) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.handler.listeners[table] = listener
}

func (c *Canal) DelListener(table string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.handler.listeners, table)
}

type CanalEventHandler struct {
	canal.DummyEventHandler
	listeners map[string]Listener
}

func NewCanalEventHandler() *CanalEventHandler {
	return &CanalEventHandler{listeners: make(map[string]Listener)}
}

func (h *CanalEventHandler) String() string {
	return "CanalEventHandler"
}

func (h *CanalEventHandler) OnRow(e *canal.RowsEvent) error {
	table := e.Table.Name
	listener, ok := h.listeners[table]
	if !ok {
		return nil
	}

	switch e.Action {
	case canal.InsertAction:
		row := listener.Scan(e.Table, e.Rows[0])
		return listener.Insert(row, e.Header)
	case canal.UpdateAction:
		before := listener.Scan(e.Table, e.Rows[0])
		after := listener.Scan(e.Table, e.Rows[1])
		return listener.Update(before, after, e.Header)
	case canal.DeleteAction:
		row := listener.Scan(e.Table, e.Rows[0])
		return listener.Delete(row, e.Header)
	}

	return nil
}
