package queries

import (
	"time"
)

type MatchFilter struct {
}

type MatchRange struct {
	BeginAt *ValueRange[time.Time] `key:"begin_at"`
}

type MatchSort struct {
}

// region MatchFilter implementation
// endregion

// region MatchRange implementation

func (r MatchRange) GetRangeQuery() map[string]string {
	return GetRangeQueryKeyValues(r)
}

// endregion
