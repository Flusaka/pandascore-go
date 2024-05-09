package clients

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestNewBaseClient(t *testing.T) {
	accessToken := "testAccessToken"
	baseClient := newBaseClient(GameDota2, accessToken)

	assert.NotNil(t, baseClient)
	assert.Equal(t, baseClient.accessToken, accessToken)
	assert.Equal(t, baseClient.baseUrl, &url.URL{
		Scheme: "https",
		Host:   "api.pandascore.co/dota2",
	})
}

func TestBaseClient_GetUpcomingTournaments(t *testing.T) {
	accessToken := "8FG9WnjcQBp9FkS8PA6bTQAEKYQefsBhWBjOG_hC7VYu4vWLxNM"
	baseClient := newBaseClient(GameDota2, accessToken)

	upcomingTournaments, err := baseClient.GetUpcomingTournaments()
	assert.Nil(t, err)
	assert.NotNil(t, upcomingTournaments)
}
