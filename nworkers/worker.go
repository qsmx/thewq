package nworkers

import (
	"reflect"
)

// Worker worker
type Worker interface {
	Unsubscribe()
	ReSubscribe()
}

type worker struct {
	key string
	ptr uintptr

	typ   reflect.Type
	value reflect.Value
}

func (w *worker) Unsubscribe() {
	_C.unsubscribe(w.key, w.ptr)
}

func (w *worker) ReSubscribe() {
	w.subscribe()
}

func (w *worker) subscribe() {
	_C.subscribe(w.key, w.ptr, w.value)
}
