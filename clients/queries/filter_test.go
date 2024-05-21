package queries

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFilterArrayInt_GetFilterString(t *testing.T) {
	a := FilterArray[int]{
		Values: []int{0, 1, 2, 3},
	}

	filter := a.GetFilterString()
	assert.Equal(t, filter, "0,1,2,3")
}

func TestDateFilterArray_GetFilterString(t *testing.T) {
	d1 := time.Date(2024, 05, 20, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(2024, 05, 21, 0, 0, 0, 0, time.UTC)
	dfa := DateFilterArray{
		Values: []time.Time{d1, d2},
	}
	filter := dfa.GetFilterString()
	assert.Equal(t, filter, "2024-05-20T00:00:00Z,2024-05-21T00:00:00Z")
}
