package queries

import (
	"fmt"
	"reflect"
)

type Ranger interface {
	GetRangeString() string
}

type ValueRange[T any] struct {
	Lower T
	Upper T
}

func (v ValueRange[T]) GetRangeString() string {
	return fmt.Sprintf("%v,%v", v.Lower, v.Upper)
}

type Range interface {
	GetRangeQuery() map[string]string
}

func GetRangeQueryKeyValues(r Range) map[string]string {
	return getReflectedKeyValues(r, func(value reflect.Value) string {
		if value.IsNil() {
			return ""
		}

		val, isValid := value.Interface().(Ranger)
		if !isValid {
			return ""
		}
		return val.GetRangeString()
	})
}
