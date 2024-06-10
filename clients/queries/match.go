package queries

type MatchSortFieldKey string

const (
	MatchSortBeginAt      MatchSortFieldKey = "begin_at"
	MatchSortTournamentId MatchSortFieldKey = "tournament_id"
	MatchSortScheduledAt  MatchSortFieldKey = "scheduled_at"
)

type MatchFilter struct {
}

type MatchRange struct {
	BeginAt     *DateRange `key:"begin_at"`
	ScheduledAt *DateRange `key:"scheduled_at"`
}

type MatchSortField = SortFieldBase[MatchSortFieldKey]

type MatchSort = SortBase[MatchSortField]

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
