package clients

import (
	"encoding/json"
	"fmt"
	"github.com/flusaka/pandascore-go/clients/queries"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestMethod string
type Endpoint string

const (
	MethodGet RequestMethod = "GET"

	EndpointUpcomingMatches Endpoint = "matches/upcoming"
)

type Request struct {
	client   *baseClient
	endpoint Endpoint
	//Filter   Filter
	rangeQuery queries.Range
	sortQuery  queries.Sort
	//Search   Search
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
	addQueryParameters("range", request.rangeQuery.GetRangeQuery(), query)
	addSortParameters(request.sortQuery.GetSortFields(), query)
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
