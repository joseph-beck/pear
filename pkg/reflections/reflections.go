package reflections

import (
	"fmt"
	"reflect"
)

func ExpectType[T any](r any) T {
	e := reflect.TypeOf((*T)(nil)).Elem()
	v := reflect.TypeOf(r)

	if e == v {
		return r.(T)
	}

	panic(fmt.Sprintf("Expected %T but received %T instead", e, v))
}
