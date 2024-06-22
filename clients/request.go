package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/flusaka/pandascore-go/clients/queries"
)

type RequestMethod string
type Endpoint string

const (
	MethodGet       RequestMethod = "GET"
	EndpointMatches Endpoint      = "matches"

	EndpointUpcomingMatches Endpoint = "matches/upcoming"
	// EndpointUpcomingTournaments is the endpoint to retrieve all upcoming tournaments
	EndpointUpcomingTournaments Endpoint = "tournaments/upcoming"
	// EndpointRunningTournaments is the endpoint to retrieve all currently running tournaments
	EndpointRunningTournaments Endpoint = "tournaments/running"
)

type Request struct {
	client     *baseClient
	endpoint   Endpoint
	filter     queries.Filter
	rangeQuery queries.Range
	sortQuery  queries.Sort
	//Search   queries.Search
}

func (r *Request) Get(value interface{}) error {
	request, err := buildHttpRequest(r, MethodGet)
	if err != nil {
		return err
	}

	httpResponse, err := r.client.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()
	if httpResponse.StatusCode < http.StatusOK || httpResponse.StatusCode >= http.StatusMultipleChoices {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		return fmt.Errorf("request failed: %s", string(errorResponse))
	}

	decoder := json.NewDecoder(httpResponse.Body)
	return decoder.Decode(value)
}

func buildHttpRequest(request *Request, method RequestMethod) (*http.Request, error) {
	requestUrl := request.client.baseUrl.JoinPath(string(request.endpoint))
	requestUrl.RawQuery = getQueryString(request, requestUrl.Query())

	httpRequest, err := http.NewRequest(string(method), requestUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(request.client.accessToken) > 0 {
		httpRequest.Header.Add("Authorization", "Bearer "+request.client.accessToken)
	}
	return httpRequest, nil
}

func getQueryString(request *Request, query url.Values) string {
	if request.filter != nil {
		addQueryParameters("filter", request.filter.GetFilterQuery(), query)
	}
	if request.rangeQuery != nil {
		addQueryParameters("range", request.rangeQuery.GetRangeQuery(), query)
	}
	if request.sortQuery != nil {
		addSortParameters(request.sortQuery.GetSortFields(), query)
	}
	return query.Encode()
}

func addQueryParameters(parameter string, parameters map[string]string, query url.Values) {
	for key, value := range parameters {
		query.Set(fmt.Sprintf("%s[%s]", parameter, key), value)
	}
}

func addSortParameters(fields []queries.SortField, query url.Values) {
	if len(fields) > 0 {
		sortFields := make([]string, len(fields))
		for i, sf := range fields {
			fieldString := ""
			if sf.IsDescending() {
				fieldString += "-"
			}
			fieldString += sf.GetFieldName()
			sortFields[i] = fieldString
		}
		sortQuery := strings.Join(sortFields, ",")
		query.Set("sort", sortQuery)
	}
}
