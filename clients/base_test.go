package clients

import (
	"net/url"
	"testing"
	"time"

	"github.com/flusaka/pandascore-go/clients/queries"
	"github.com/stretchr/testify/assert"
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

	upcomingMatches, err := baseClient.GetUpcomingMatchesWithParams(MatchParams{
		Range: queries.MatchRange{
			BeginAt: &queries.DateRange{
				Lower: time.Now().Truncate(time.Hour * 24),
				Upper: time.Now().Truncate(time.Hour * 24).Add(time.Hour * 24),
			},
		},
		Sort: queries.NewMatchSort([]queries.MatchSortField{
			{
				FieldName:  queries.MatchSortBeginAt,
				Descending: true,
			},
			{
				FieldName:  queries.MatchSortTournamentId,
				Descending: false,
			},
		}),
	})
	assert.Nil(t, err)
	assert.NotNil(t, upcomingMatches)
}

func TestBaseClient_GetUpcomingTournaments(t *testing.T) {
	accessToken := "8FG9WnjcQBp9FkS8PA6bTQAEKYQefsBhWBjOG_hC7VYu4vWLxNM"
	baseClient := newBaseClient(GameDota2, accessToken)

	upcomingTournaments, err := baseClient.GetUpcomingTournaments()
	assert.Nil(t, err)
	assert.NotNil(t, upcomingTournaments)
}

func TestBaseClient_GetRunningTournaments(t *testing.T) {
	accessToken := "8FG9WnjcQBp9FkS8PA6bTQAEKYQefsBhWBjOG_hC7VYu4vWLxNM"
	baseClient := newBaseClient(GameDota2, accessToken)

	runningTournaments, err := baseClient.GetRunningTournaments()
	assert.Nil(t, err)
	assert.NotNil(t, runningTournaments)
}

func TestBaseClient_GetUpcomingTournamentsWithParams(t *testing.T) {
	accessToken := "8FG9WnjcQBp9FkS8PA6bTQAEKYQefsBhWBjOG_hC7VYu4vWLxNM"
	baseClient := newBaseClient(GameDota2, accessToken)

	upcomingTournaments, err := baseClient.GetUpcomingTournamentsWithParams(TournamentParams{
		Range: queries.TournamentRange{},
		Sort: queries.NewTournamentSort([]queries.TournamentSortField{
			{
				FieldName:  queries.TournamentSortBeginAt,
				Descending: false,
			},
		}),
	})
	assert.Nil(t, err)
	assert.NotNil(t, upcomingTournaments)
}
