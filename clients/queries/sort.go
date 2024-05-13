package queries

type SortFields []string

type Sort interface {
	GetSortFields() SortFields
}
