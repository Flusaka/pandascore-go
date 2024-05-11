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

type Params struct {
	Range queries.Range
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

func (c *baseClient) RequestWithParams(endpoint Endpoint, params Params) *Request {
	return &Request{
		client:     c,
		endpoint:   endpoint,
		rangeQuery: params.Range,
	}
}

func (c *baseClient) GetUpcomingMatches() ([]types.BaseMatch, error) {
	var matches []types.BaseMatch
	err := c.Request(EndpointUpcomingMatches).Get(&matches)
	return matches, err
}

func (c *baseClient) GetUpcomingMatchesWithParams(params Params) ([]types.BaseMatch, error) {
	var matches []types.BaseMatch
	err := c.RequestWithParams(EndpointUpcomingMatches, params).Get(&matches)
	return matches, err
}
