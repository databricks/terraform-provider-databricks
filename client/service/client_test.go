package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ForceErrorServer(t *testing.T, response string, responseStatus int, apiCall func(client DBApiClient) (interface{}, error)) (interface{}, error) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(responseStatus)
		_, err := rw.Write([]byte(response))
		assert.NoError(t, err, err)
	}))
	// Close the server when test finishes
	defer server.Close()
	var config DBApiClientConfig
	config.Host = server.URL
	config.Setup()

	var dbClient DBApiClient
	dbClient.SetConfig(&config)

	return apiCall(dbClient)
}

func TestClient_HandleErrors(t *testing.T) {
	tests := []struct {
		name               string
		response           string
		responseStatus     int
		expectedErrorCode  string
		expectedMessage    string
		expectedResource   string
		expectedStatusCode int
		apiCall            func(client DBApiClient) (interface{}, error)
	}{
		{
			name: "Status 404",
			response: `{
							"error_code": "RESOURCE_DOES_NOT_EXIST",
							"message": "Token ... does not exist!"
						}`,
			responseStatus:     http.StatusNotFound,
			expectedErrorCode:  "RESOURCE_DOES_NOT_EXIST",
			expectedMessage:    "Token ... does not exist!",
			expectedResource:   "/api/2.0/token/create",
			expectedStatusCode: 404,
			apiCall: func(client DBApiClient) (interface{}, error) {
				return client.Tokens().Create(10, "USERS")
			},
		},
		{
			name:               "HTML Status 404",
			response:           `<pre> Hello world </pre>`,
			responseStatus:     http.StatusNotFound,
			expectedErrorCode:  "NOT_FOUND",
			expectedMessage:    "Hello world",
			expectedResource:   "/api/2.0/token/create",
			expectedStatusCode: 404,
			apiCall: func(client DBApiClient) (interface{}, error) {
				return client.Tokens().Create(10, "USERS")
			},
		},
		{
			name:               "Invalid HTML Status 404",
			response:           `<html> Hello world </html>`,
			responseStatus:     http.StatusNotFound,
			expectedErrorCode:  "NOT_FOUND",
			expectedMessage:    "Response from server (404) <html> Hello world </html>: invalid character '<' looking for beginning of value",
			expectedResource:   "/api/2.0/token/create",
			expectedStatusCode: 404,
			apiCall: func(client DBApiClient) (interface{}, error) {
				return client.Tokens().Create(10, "USERS")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ForceErrorServer(t, tt.response, tt.responseStatus, tt.apiCall)
			t.Log(err)
			assert.IsType(t, APIError{}, err)
			assert.Equal(t, tt.expectedErrorCode, err.(APIError).ErrorCode, "error code is not the same")
			assert.Equal(t, tt.expectedMessage, err.(APIError).Message, "message is not the same")
			assert.Equal(t, tt.expectedResource, err.(APIError).Resource, "resource is not the same")
			assert.Equal(t, tt.expectedStatusCode, err.(APIError).StatusCode, "status code is not the same")
		})
	}
}
