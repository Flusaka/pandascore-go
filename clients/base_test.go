package clients

import (
	"github.com/flusaka/pandascore-go/clients/queries"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
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

func TestBaseClient_GetUpcomingMatches(t *testing.T) {
	accessToken := "8FG9WnjcQBp9FkS8PA6bTQAEKYQefsBhWBjOG_hC7VYu4vWLxNM"
	baseClient := newBaseClient(GameDota2, accessToken)

	upcomingMatches, err := baseClient.GetUpcomingMatches()
	assert.Nil(t, err)
	assert.NotNil(t, upcomingMatches)
}

func TestBaseClient_GetUpcomingMatchesWithParams(t *testing.T) {
	accessToken := "8FG9WnjcQBp9FkS8PA6bTQAEKYQefsBhWBjOG_hC7VYu4vWLxNM"
	baseClient := newBaseClient(GameDota2, accessToken)

	upcomingMatches, err := baseClient.GetUpcomingMatchesWithParams(Params[queries.MatchRange]{
		Range: queries.MatchRange{
			BeginAt: &queries.DateRange{
				Lower: time.Now().Truncate(time.Hour * 24),
				Upper: time.Now().Truncate(time.Hour * 24).Add(time.Hour * 24),
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, upcomingMatches)
}
