package queries

type MatchSortFieldKey string

const (
	MatchSortBeginAt      MatchSortFieldKey = "begin_at"
	MatchSortTournamentId MatchSortFieldKey = "tournament_id"
)

type MatchFilter struct {
}

type MatchRange struct {
	BeginAt *DateRange `key:"begin_at"`
}

type MatchSortField struct {
	FieldName  MatchSortFieldKey
	Descending bool
}

type MatchSort struct {
	sortFields []MatchSortField
}

func NewMatchSort(sortFields []MatchSortField) MatchSort {
	sort := MatchSort{
		sortFields: sortFields,
	}
	return sort
}

// region MatchFilter implementation
// endregion

// region MatchRange implementation

func (r MatchRange) GetRangeQuery() map[string]string {
	return GetRangeQueryKeyValues(r)
}

// endregion

// region MatchSort implementation

func (f MatchSortField) GetFieldName() string {
	return string(f.FieldName)
}

func (f MatchSortField) IsDescending() bool {
	return f.Descending
}

func (s MatchSort) GetSortFields() []SortField {
	sortFields := make([]SortField, len(s.sortFields))
	for i, msf := range s.sortFields {
		sortFields[i] = msf
	}
	return sortFields
}

// endregion
