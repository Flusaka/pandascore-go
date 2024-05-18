package queries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchSort_GetSortFields(t *testing.T) {
	sort := NewMatchSort(
		[]MatchSortField{
			{
				FieldName:  "begin_at",
				Descending: false,
			},
			{
				FieldName:  "tournament_id",
				Descending: true,
			},
		},
	)

	fields := sort.GetSortFields()
	assert.Len(t, fields, 2)
	assert.Equal(t, "begin_at", fields[0].GetFieldName())
	assert.Equal(t, false, fields[0].IsDescending())
	assert.Equal(t, "tournament_id", fields[1].GetFieldName())
	assert.Equal(t, true, fields[1].IsDescending())
}
