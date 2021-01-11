package httpclientmock

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// HTTPClientMock manages the HTTPClient.Do responses expected for tests and stores the
// received request and current request count to use for test validation.
type HTTPClientMock struct {
	DoMock DoMock
}

// DoMock manages the HTTPClient.Do responses expected for tests and stores the
// received request and current request count to use for test validation.
type DoMock struct {
	Calls     int
	Responses []HTTPResponseMock
	Requests  []*http.Request
}

// HTTPResponseMock represents how the http response data should be mocked for test.
type HTTPResponseMock struct {
	StatusCode int
	Body       string
	HasError   bool
}

// Do mocks HTTPClient.Do method; it stores the http.Request it received on the Mock instance, increments the request count total and returns the mocked response expected for the http call.
func (m *HTTPClientMock) Do(request *http.Request) (*http.Response, error) {
	calls := m.DoMock.Calls
	m.DoMock.Requests = append(m.DoMock.Requests, request)

	if calls >= len(m.DoMock.Responses) {
		return nil, fmt.Errorf("Test HTTP Error: a response was not mocked for '%s' http method call to '%s'", request.Method, request.URL.String())
	}

	var resp *http.Response
	var err error

	if m.DoMock.Responses[calls].HasError {
		err = errors.New("Error")
	} else {
		resp = &http.Response{
			StatusCode: m.DoMock.Responses[calls].StatusCode,
			Body:       ioutil.NopCloser(strings.NewReader(m.DoMock.Responses[calls].Body)),
		}
	}

	m.DoMock.Calls = calls + 1
	return resp, err
}
