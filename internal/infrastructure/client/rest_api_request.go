package client

import (
	"bytes"
	"encoding/json"
	"io"
	domain2 "myApiController/internal/domain"
	"myApiController/internal/domain/model"
	"net/http"
	"regexp"
	"time"
)

const INNER_VARAIBLE = `\{(\w+)\}`

type RestApi struct {
	url        string
	HttpClient *http.Client
	headers    http.Header
	httpMethod string
}

func NewRestApi(path, httpMethod string, headers map[string]string, httpClient http.Client) domain2.DataRowClient {
	header := http.Header{}
	for key, value := range headers {
		header.Add(key, value)
	}

	return &RestApi{
		url:        path,
		HttpClient: &httpClient,
		headers:    header,
		httpMethod: httpMethod,
	}
}

func (e *RestApi) DoRequest(params map[string]string, bodyStr string) (domain2.DataExchange, error) {
	var (
		err error
		req *http.Request
	)

	path := e.generatePath(params)
	if bodyStr == "" {
		req, err = http.NewRequest(e.httpMethod, path, nil)
	} else {
		body := bytes.NewBuffer([]byte(bodyStr))
		req, err = http.NewRequest(e.httpMethod, path, body)
	}

	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()
	req.Header = e.headers

	var response model.ApiResponse
	startTime := time.Now()
	resp, err := e.HttpClient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return response, err
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Seconds()
	response["api_consumer_elapsed_time"] = elapsedTime

	return response, nil
}

func (e *RestApi) generatePath(params map[string]string) string {
	re := regexp.MustCompile(INNER_VARAIBLE)
	if !re.MatchString(e.url) {
		return e.url
	}
	return re.ReplaceAllStringFunc(e.url, func(match string) string {
		key := match[1 : len(match)-1]
		value, ok := params[key]
		if ok {
			return value
		}
		return match
	})
}
