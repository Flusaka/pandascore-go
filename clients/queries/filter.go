package queries

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Filterer interface {
	GetFilterString() string
}

type Filter interface {
	GetFilterQuery() map[string]string
}

type FilterValue[T any] struct {
	Value T
}

func (v FilterValue[T]) GetFilterString() string {
	return fmt.Sprintf("%v", v.Value)
}

type BoolFilter = FilterValue[bool]

type FilterArray[T any] struct {
	Values []T
}

func (v FilterArray[T]) GetFilterString() string {
	var values []string
	for _, v := range v.Values {
		values = append(values, fmt.Sprintf("%v", v))
	}
	return strings.Join(values, ",")
}

type DateFilterArray FilterArray[time.Time]

func (d DateFilterArray) GetFilterString() string {
	var values []string
	for _, v := range d.Values {
		values = append(values, fmt.Sprintf("%v", v.UTC().Format(time.RFC3339)))
	}
	return strings.Join(values, ",")
}

func GetFilterQueryKeyValues(f Filter) map[string]string {
	return getReflectedKeyValues(f, func(value reflect.Value) string {
		if value.IsNil() {
			return ""
		}

		val, isValid := value.Interface().(Filterer)
		if !isValid {
			return ""
		}
		return val.GetFilterString()
	})
}
