package nworkers

import "errors"

var (
	ErrIsntFunc  = errors.New("it's not func")
	ErrValidFunc = errors.New("it's valid type")

	ErrNotFoundSubscribe = errors.New("not found subscribe")
)
