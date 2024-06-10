package queries

type TournamentSortFieldKey string

const (
	TournamentSortBeginAt TournamentSortFieldKey = "begin_at"
)

type TournamentRange struct {
	BeginAt       *DateRange   `key:"begin_at"`
	DetailedStats *BoolRange   `key:"detailed_stats"`
	EndAt         *DateRange   `key:"end_at"`
	HasBracket    *BoolRange   `key:"has_bracket"`
	ID            *IntRange    `key:"id"`
	ModifiedAt    *DateRange   `key:"modified_at"`
	Name          *StringRange `key:"name"`
	PrizePool     *IntRange    `key:"prizepool"`
	SerieID       *IntRange    `key:"serie_id"`
	Slug          *StringRange `key:"slug"`
	Tier          *StringRange `key:"tier"`
}

type TournamentSortField = SortFieldBase[TournamentSortFieldKey]

type TournamentSort = SortBase[TournamentSortField]

func NewTournamentSort(sortFields []TournamentSortField) TournamentSort {
	sort := TournamentSort{
		sortFields: sortFields,
	}
	return sort
}

// region TournamentRange implementation

func (r TournamentRange) GetRangeQuery() map[string]string {
	return GetRangeQueryKeyValues(r)
}

// endregion
