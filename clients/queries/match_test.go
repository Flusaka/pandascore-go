package queries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchSort_GetSortFields(t *testing.T) {
	sort := NewMatchSort(MatchSortFields{
		MatchSortBeginAt:      true,
		MatchSortTournamentId: false,
	})

	fields := sort.GetSortFields()
	assert.Len(t, fields, 2)
	assert.Equal(t, "-begin_at", fields[0])
	assert.Equal(t, "tournament_id", fields[1])
}
