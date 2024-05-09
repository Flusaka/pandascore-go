package clients

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RequestMethod string

const (
	MethodGet RequestMethod = "GET"
)

type Request struct {
	client   *baseClient
	endpoint string
}

func (r Request) Get(value interface{}) error {
	request, err := buildHttpRequest(&r, MethodGet)
	if err != nil {
		return err
	}

	httpResponse, err := r.client.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()
	if httpResponse.StatusCode < http.StatusOK || httpResponse.StatusCode >= http.StatusMultipleChoices {
		return errors.New("request failed")
	}

	decoder := json.NewDecoder(httpResponse.Body)
	return decoder.Decode(value)
}

func buildHttpRequest(request *Request, method RequestMethod) (*http.Request, error) {
	requestUrl := request.client.baseUrl.JoinPath(request.endpoint)

	httpRequest, err := http.NewRequest(string(method), requestUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(request.client.accessToken) > 0 {
		httpRequest.Header.Add("Authorization", "Bearer "+request.client.accessToken)
	}
	return httpRequest, nil
}
