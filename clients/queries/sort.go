package queries

import "fmt"

type SortFieldValue[T fmt.Stringer] struct {
	FieldName  T
	Descending bool
}

func (f SortFieldValue[T]) GetFieldName() string {
	return f.FieldName.String()
}

func (f SortFieldValue[T]) IsDescending() bool {
	return f.Descending
}

type SortField interface {
	GetFieldName() string
	IsDescending() bool
}

type Sort interface {
	GetSortFields() []SortField
}
