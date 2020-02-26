package nworkers

import (
	"fmt"
	"testing"
)

func TestNWorkers_Subscribe(t *testing.T) {
	v := func(a int, b ...int) { fmt.Println("worker_test_1", a, b) }
	v1 := func(a ...int) { fmt.Println("worker_test_2", a) }
	f0, err := Subscribe("v", v)
	f1, err := Subscribe("v", v1)
	t.Log(err)

	Publish("v", 10, 12)
	Publish("v", 14)
	Publish("v", 15, 1, 2, 3, 4, 5)

	f0.Unsubscribe()
	_ = f1

	Publish("v", 15, 10, 123, 0x88382)
	Publish("v", 8, 0123)

	Stop()
}

func TestConvertibleTo(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(a[1:])

	t.Log(a[0:3], a[4:], a[7:])
}
