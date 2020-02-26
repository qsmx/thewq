package nworkers

import (
	"reflect"
	"sync"

	"github.com/qsmx/thewq"
)

var _C = &container{
	c:   map[string][]uintptr{},
	all: map[uintptr]reflect.Value{},
}

type container struct {
	l sync.RWMutex
	c map[string][]uintptr

	la  sync.RWMutex
	all map[uintptr]reflect.Value

	wg sync.WaitGroup
}

func (c *container) unsubscribe(key string, ptr uintptr) {
	c.l.Lock()
	list, ok := c.c[key]
	if !ok {
		c.l.Unlock()
		return
	}

	var idx int
	var v uintptr

	for idx, v = range list {
		if v == ptr {
			if len(list) == 1 {
				delete(c.c, key)
			} else {
				list[idx] = list[len(list)-1]
				c.c[key] = list[:len(list)-1]
			}

			break
		}
	}

	c.l.Unlock()

	c.la.Lock()
	delete(c.all, ptr)
	c.la.Unlock()
}

func (c *container) subscribe(key string, ptr uintptr, value reflect.Value) {
	c.l.Lock()
	if _, ok := c.c[key]; ok {
		c.c[key] = append(c.c[key], ptr)
	} else {
		c.c[key] = []uintptr{ptr}
	}
	c.l.Unlock()

	c.la.Lock()
	c.all[ptr] = value
	c.la.Unlock()
}

func (c *container) publish(key string, args ...interface{}) error {
	c.l.RLock()
	list, ok := c.c[key]
	if !ok || len(list) == 0 {
		c.l.RUnlock()
		return ErrNotFoundSubscribe
	}
	c.l.RUnlock()

	l := len(args)
	a := make([]reflect.Value, l)
	for i, v := range args {
		a[i] = reflect.ValueOf(v)
	}

	c.la.RLock()
	for _, ptr := range list {
		value, ok := c.all[ptr]
		if !ok {
			continue
		}

		c.wg.Add(1)
		go func() {
			defer func() {
				c.wg.Done()

				if err := recover(); err != nil {
					thewq.Errorf("Panic, [ERR]: %s", err)
				}
			}()

			value.Call(a)
		}()
	}
	c.la.RUnlock()

	return nil
}

func (c *container) wait() {
	c.wg.Wait()
}
