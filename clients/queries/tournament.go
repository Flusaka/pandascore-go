package queries

import "github.com/flusaka/pandascore-go/types"

type TournamentSortFieldKey string

const (
	TournamentSortBeginAt TournamentSortFieldKey = "begin_at"
)

func (k TournamentSortFieldKey) String() string {
	return string(k)
}

type TournamentFilter struct {
	BeginAt       *DateFilterArray         `key:"begin_at"`
	DetailedStats *BoolFilter              `key:"detailed_stats"`
	EndAt         *DateFilterArray         `key:"end_at"`
	HasBracket    *BoolFilter              `key:"has_bracket"`
	ID            *FilterArray[int]        `key:"id"`
	LiveSupported *BoolFilter              `key:"live_supported"`
	ModifiedAt    *DateFilterArray         `key:"modified_at"`
	Name          *FilterArray[string]     `key:"name"`
	SerieID       *FilterArray[int]        `key:"serie_id"`
	Tier          *FilterArray[types.Tier] `key:"tier"`
}

type TournamentRange struct {
	BeginAt       *DateRange `key:"begin_at"`
	DetailedStats *BoolRange `key:"detailed_stats"`
	EndAt         *DateRange `key:"end_at"`
	HasBracket    *BoolRange `key:"has_bracket"`
	ID            *IntRange  `key:"id"`
	ModifiedAt    *DateRange `key:"modified_at"`
	SerieID       *IntRange  `key:"serie_id"`
	Tier          *ValueRange[types.Tier]
}

type TournamentSortField = SortFieldValue[TournamentSortFieldKey]

type TournamentSort = SortBase[TournamentSortField]

func NewTournamentSort(sortFields []TournamentSortField) TournamentSort {
	return TournamentSort{
		sortFields: sortFields,
	}
}

// region TournamentFilter implementation

func (f TournamentFilter) GetFilterQuery() map[string]string {
	return GetFilterQueryKeyValues(f)
}

// endregion

// region TournamentRange implementation

func (r TournamentRange) GetRangeQuery() map[string]string {
	return GetRangeQueryKeyValues(r)
}

// endregion
