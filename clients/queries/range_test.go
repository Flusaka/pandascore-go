package queries

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TestRange struct {
	BeginAt *ValueRange[time.Time] `key:"begin_at"`
	EndAt   *ValueRange[time.Time] `key:"end_at"`
}

func TestGetRangeQueryKeyValues(t *testing.T) {
	r := TestRange{BeginAt: &ValueRange[time.Time]{
		Lower: time.Now(),
		Upper: time.Now().Add(time.Minute),
	}}
	keyValues := GetRangeQueryKeyValues(r)
	assert.Greater(t, len(keyValues["begin_at"]), 0)
	assert.Equal(t, len(keyValues["end_at"]), 0)
}
