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

type Params[R queries.Range] struct {
	Range R
}

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

func (c *baseClient) GetUpcomingMatches() ([]types.BaseMatch, error) {
	var matches []types.BaseMatch
	err := c.Request(EndpointUpcomingMatches).Get(&matches)
	return matches, err
}

func (c *baseClient) GetUpcomingMatchesWithParams(params Params[queries.MatchRange]) ([]types.BaseMatch, error) {
	var matches []types.BaseMatch
	err := c.Request(EndpointUpcomingMatches).WithRange(params.Range).Get(&matches)
	return matches, err
}
