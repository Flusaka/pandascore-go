package queries

type MatchSortFieldKey string

const (
	MatchSortBeginAt      MatchSortFieldKey = "begin_at"
	MatchSortTournamentId MatchSortFieldKey = "tournament_id"
)

type MatchSortFields map[MatchSortFieldKey]bool

type MatchFilter struct {
}

type MatchRange struct {
	BeginAt *DateRange `key:"begin_at"`
}

type MatchSort struct {
	sortFields MatchSortFields
}

func NewMatchSort(sortFields MatchSortFields) MatchSort {
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

func (r MatchSort) GetSortFields() SortFields {
	sortFields := make(SortFields, 0, len(r.sortFields))
	for key, value := range r.sortFields {
		query := ""
		if value {
			query += "-"
		}
		query += string(key)
		sortFields = append(sortFields, query)
	}
	return sortFields
}

// endregion
