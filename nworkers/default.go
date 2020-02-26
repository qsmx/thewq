// subscribe

package nworkers

import (
	"reflect"
	"unsafe"
)

type FuncHandler interface{}

// Publish publish message
func Publish(key string, args ...interface{}) error {
	return _C.publish(key, args...)
}

// Subscribe subscribe message
func Subscribe(key string, fnc FuncHandler) (Worker, error) {
	typ := reflect.TypeOf(fnc)
	if typ.Kind() != reflect.Func {
		return nil, ErrIsntFunc
	}

	return subscribe(key, uintptr(unsafe.Pointer(&fnc)), typ, reflect.ValueOf(fnc))
}

func subscribe(key string, ptr uintptr, typ reflect.Type, value reflect.Value) (Worker, error) {
	if value.IsNil() {
		return nil, ErrValidFunc
	}

	w := &worker{
		key:   key,
		ptr:   ptr,
		typ:   typ,
		value: value,
	}

	w.subscribe()
	return w, nil
}

func Stop() {
	_C.wait()
}
