package client

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	domain2 "myApiController/internal/domain"
	"net/http"
	"strings"
	"testing"
)

const (
	EXAMPLE_SIMPLE_URL      = "http://test.com/GET"
	EXAMPLE_WITH_PARAMS     = "http://test.com/GET/{param1}/{param2}"
	EXAMPLE_POST_SIMPLE_URL = "http://test.com/POST"
)

type restApiTestScenario struct {
	test        *testing.T
	client      domain2.DataRowClient
	mockClient  http.Client
	params      map[string]string
	bodyStr     string
	result      domain2.DataExchange
	err         error
	expectedErr error
}

/*-- Test --*/
func TestGetRequestOk(t *testing.T) {
	s := startRestApiTestScenario(t)
	headers := map[string]string{"hello": "world"}
	params := map[string]string{"lorem": "ipsum"}
	s.givenAServerOk(EXAMPLE_SIMPLE_URL, headers, http.MethodGet)
	s.andSomeParams(params)
	s.whenDoingRequest()
	s.thenThereIsNoError()
}

func TestPostRequestOk(t *testing.T) {
	s := startRestApiTestScenario(t)
	headers := map[string]string{"hello": "world"}
	s.givenAServerOk(EXAMPLE_POST_SIMPLE_URL, headers, http.MethodPost)
	s.andSomeBody(`{"hello":"world"}`)
	s.whenDoingRequest()
	s.thenThereIsNoError()
}

func TestGetRequestOkWithParams(t *testing.T) {
	s := startRestApiTestScenario(t)
	headers := map[string]string{"hello": "world"}
	params := map[string]string{"lorem": "ipsum", "param1": "12345", "param2": "sample"}
	s.givenAServerOk(EXAMPLE_WITH_PARAMS, headers, http.MethodGet)
	s.andSomeParams(params)
	s.whenDoingRequest()
	s.thenThereIsNoError()
}

func TestGetRequestOkWithParamsNoProvided(t *testing.T) {
	s := startRestApiTestScenario(t)
	headers := map[string]string{"hello": "world"}
	params := map[string]string{}
	s.givenAServerOk(EXAMPLE_WITH_PARAMS, headers, http.MethodGet)
	s.andSomeParams(params)
	s.whenDoingRequest()
	s.thenThereIsNoError()
}

func TestGetRequestWithBadResponse(t *testing.T) {
	s := startRestApiTestScenario(t)
	headers := map[string]string{"hello": "world"}
	params := map[string]string{"lorem": "ipsum"}
	s.givenAServerBadResponse(EXAMPLE_SIMPLE_URL, headers, http.MethodGet)
	s.andSomeParams(params)
	s.andAnExpectedError(fmt.Errorf("invalid character 'H' looking for beginning of value"))
	s.whenDoingRequest()
	s.thenThereIsAnError()
}

func TestErrorProcessingGetRequest(t *testing.T) {
	s := startRestApiTestScenario(t)
	headers := map[string]string{"hello": "world"}
	params := map[string]string{"lorem": "ipsum"}
	s.givenAServerWithError(EXAMPLE_SIMPLE_URL, headers, http.MethodGet)
	s.andSomeParams(params)
	s.andAnExpectedError(fmt.Errorf(`Get "http://test.com/GET?lorem=ipsum": something went wrong`))
	s.whenDoingRequest()
	s.thenThereIsAnError()
}

func TestErrorEOFResponseOnGetRequest(t *testing.T) {
	s := startRestApiTestScenario(t)
	headers := map[string]string{"hello": "world"}
	params := map[string]string{"lorem": "ipsum"}
	s.givenAServerWithErrorResponseEOF(EXAMPLE_SIMPLE_URL, headers, http.MethodGet)
	s.andSomeParams(params)
	s.andAnExpectedError(fmt.Errorf("invalid character 'H' looking for beginning of value"))
	s.whenDoingRequest()
	s.thenThereIsAnError()
}

/*-- steps --*/
func startRestApiTestScenario(t *testing.T) *restApiTestScenario {
	t.Parallel()
	return &restApiTestScenario{
		test: t,
	}
}

func (e *restApiTestScenario) givenAServerOk(path string, headers map[string]string, method string) {
	e.mockClient = http.Client{Transport: &mockClientOk{}}
	e.client = NewRestApi(path, method, headers, e.mockClient)
}

func (e *restApiTestScenario) givenAServerBadResponse(path string, headers map[string]string, method string) {
	e.mockClient = http.Client{Transport: &mockClientBadResponse{}}
	e.client = NewRestApi(path, method, headers, e.mockClient)
}

func (e *restApiTestScenario) givenAServerWithError(path string, headers map[string]string, method string) {
	e.mockClient = http.Client{Transport: &mockClientError{}}
	e.client = NewRestApi(path, method, headers, e.mockClient)
}

func (e *restApiTestScenario) givenAServerWithErrorResponseEOF(path string, headers map[string]string, method string) {
	e.mockClient = http.Client{Transport: &mockClientResponseEOF{}}
	e.client = NewRestApi(path, method, headers, e.mockClient)
}

func (e *restApiTestScenario) andSomeParams(params map[string]string) {
	e.params = params
}

func (e *restApiTestScenario) andSomeBody(body string) {
	e.bodyStr = body
}

func (e *restApiTestScenario) andAnExpectedError(err error) {
	e.expectedErr = err
}

func (e *restApiTestScenario) whenDoingRequest() {
	e.result, e.err = e.client.DoRequest(e.params, e.bodyStr)
}

func (e *restApiTestScenario) thenThereIsNoError() {
	assert.NoError(e.test, e.err)
}

func (e *restApiTestScenario) thenThereIsAnError() {
	assert.Error(e.test, e.err)
	assert.Equal(e.test, e.expectedErr.Error(), e.err.Error())
}

/*-- mocks --*/
type mockClientOk struct{}

func (t *mockClientOk) RoundTrip(req *http.Request) (*http.Response, error) {
	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("{}")),
	}
	return res, nil
}

type mockClientBadResponse struct{}

func (t *mockClientBadResponse) RoundTrip(req *http.Request) (*http.Response, error) {
	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("Hello, world")),
	}
	return res, nil
}

type mockClientError struct{}

func (t *mockClientError) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("something went wrong")
}

type mockClientResponseEOF struct{}

func (t *mockClientResponseEOF) RoundTrip(req *http.Request) (*http.Response, error) {
	body := "Hello, world"
	limit := int64(len(body)) - 1
	reader := io.LimitReader(strings.NewReader(body), limit)
	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(reader),
	}
	return res, nil
}
