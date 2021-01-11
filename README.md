# http-test-helpers
HTTP Test Helpers to Support Testing HTTP in GoLang.

- `HttpClientMock` - struct which defines the HTTPClient mock functionality.

- `HttpClientMock.Do` - manages the HTTPClient.Do responses expected for tests and stores the received HTTP Request and total request count to use in test validation.

- `DoMock` - struct which defines the HTTPClient.Do responses expected for tests and stores the received HTTP requests and total request count to use in test validation.
  - Structure: Calls (int), Requests ([]*http.Requests), Responses ([]HTTPResponseMock)

- `HTTPResponseMock` - the mocked HTTP Response the test expects for a particular request.
  - Structure: StatusCode (int), Body (string), HasError (bool)