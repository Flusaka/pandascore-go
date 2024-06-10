package clients

import (
	"github.com/flusaka/pandascore-go/clients/queries"
	"github.com/flusaka/pandascore-go/types"
	"net/http"
	"net/url"
)

type Game string

const (
	BaseUrl = "api.pandascore.co"

	GameDota2 Game = "dota2"
)

type baseClient struct {
	baseUrl     *url.URL
	httpClient  *http.Client
	accessToken string
}

type Params[R queries.Range, So queries.Sort] struct {
	Range R
	Sort  So
}

type MatchParams = Params[queries.MatchRange, queries.MatchSort]
type TournamentParams = Params[queries.TournamentRange, queries.TournamentSort]

func newBaseClient(game Game, accessToken string) *baseClient {
	return &baseClient{
		baseUrl:     &url.URL{Scheme: "https", Host: BaseUrl, Path: string(game)},
		httpClient:  http.DefaultClient,
		accessToken: accessToken,
	}
}

func (c *baseClient) Request(endpoint Endpoint) *Request {
	return &Request{
		client:   c,
		endpoint: endpoint,
	}
}

func (c *baseClient) GetUpcomingMatches() ([]types.Match, error) {
	var matches []types.Match
	err := c.Request(EndpointUpcomingMatches).Get(&matches)
	return matches, err
}

func (c *baseClient) GetUpcomingMatchesWithParams(params MatchParams) ([]types.Match, error) {
	var matches []types.Match
	err := c.Request(EndpointUpcomingMatches).WithRange(params.Range).WithSort(params.Sort).Get(&matches)
	return matches, err
}

func (c *baseClient) GetUpcomingTournaments() ([]types.Tournament, error) {
	var tournaments []types.Tournament
	err := c.Request(EndpointUpcomingTournaments).Get(&tournaments)
	return tournaments, err
}

func (c *baseClient) GetUpcomingTournamentsWithParams(params TournamentParams) ([]types.Tournament, error) {
	var tournaments []types.Tournament
	err := c.Request(EndpointUpcomingTournaments).WithRange(params.Range).WithSort(params.Sort).Get(&tournaments)
	return tournaments, err
}

func (c *baseClient) GetRunningTournaments() ([]types.Tournament, error) {
	var tournaments []types.Tournament
	err := c.Request(EndpointRunningTournaments).Get(&tournaments)
	return tournaments, err
}

func (c *baseClient) GetRunningTournamentsWithParams(params TournamentParams) ([]types.Tournament, error) {
	var tournaments []types.Tournament
	err := c.Request(EndpointRunningTournaments).WithRange(params.Range).WithSort(params.Sort).Get(&tournaments)
	return tournaments, err
}
