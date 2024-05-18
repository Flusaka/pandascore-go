package queries

type SortField interface {
	GetFieldName() string
	IsDescending() bool
}

type Sort interface {
	GetSortFields() []SortField
}
