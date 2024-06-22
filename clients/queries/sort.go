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

type SortFieldBase[K ~string] struct {
	FieldName  K
	Descending bool
}

func (s SortFieldBase[K]) GetFieldName() string {
	return string(s.FieldName)
}

func (s SortFieldBase[K]) IsDescending() bool {
	return s.Descending
}

type SortBase[SF SortField] struct {
	sortFields []SF
}

func (s SortBase[SF]) GetSortFields() []SortField {
	sortFields := make([]SortField, len(s.sortFields))
	for i, tsf := range s.sortFields {
		sortFields[i] = tsf
	}
	return sortFields
}
