package queries

type MatchSortFieldKey string

const (
	MatchSortBeginAt      MatchSortFieldKey = "begin_at"
	MatchSortTournamentId MatchSortFieldKey = "tournament_id"
)

func (k MatchSortFieldKey) String() string { return string(k) }

type MatchFilter struct {
}

type MatchRange struct {
	BeginAt     *DateRange `key:"begin_at"`
	ScheduledAt *DateRange `key:"scheduled_at"`
}

type MatchSortField = SortFieldValue[MatchSortFieldKey]

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

func (f MatchFilter) GetFilterQuery() map[string]string {
	return GetFilterQueryKeyValues(f)
}

// endregion

// region MatchRange implementation

func (r MatchRange) GetRangeQuery() map[string]string {
	return GetRangeQueryKeyValues(r)
}

// endregion

// region MatchSort implementation

func (s MatchSort) GetSortFields() []SortField {
	sortFields := make([]SortField, len(s.sortFields))
	for i, msf := range s.sortFields {
		sortFields[i] = msf
	}
	return sortFields
}

// endregion
