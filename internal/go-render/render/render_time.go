package render

import (
	"reflect"
	"time"
)

var timeType = reflect.TypeOf(time.Time{})

func renderTime(value reflect.Value) (rendered string, ok bool) {
	if value.Type() != timeType {
		return "", false
	}

	defer func() {
		// If the value is a private field of a containing struct, calling value.Interface() will panic.
		if r := recover(); r != nil {
			rendered = ""
			ok = false
		}
	}()
	return value.Interface().(interface{ String() string }).String(), true
}
