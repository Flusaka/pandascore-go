package pandascore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	accessToken := "testAccessToken"
	client := NewClient(accessToken)

	assert.NotNil(t, client)
	assert.NotNil(t, client.Dota2)
}
