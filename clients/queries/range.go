package queries

import (
	"fmt"
	"reflect"
	"time"
)

type Ranger interface {
	GetRangeString() string
}

type Range interface {
	GetRangeQuery() map[string]string
}

type ValueRange[T any] struct {
	Lower T
	Upper T
}

func (v ValueRange[T]) GetRangeString() string {
	return fmt.Sprintf("%v,%v", v.Lower, v.Upper)
}

type StringRange = ValueRange[string]
type BoolRange = ValueRange[bool]
type IntRange = ValueRange[int]
type DateRange ValueRange[time.Time]

func (d DateRange) GetRangeString() string {
	// Date range format has to be in UTC and RFC3339/ISO8601 format
	return fmt.Sprintf("%v,%v", d.Lower.UTC().Format(time.RFC3339), d.Upper.UTC().Format(time.RFC3339))
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
