package common

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_HandleErrors(t *testing.T) {
	tests := []struct {
		name               string
		response           string
		responseStatus     int
		expectedErrorCode  string
		expectedMessage    string
		expectedResource   string
		expectedStatusCode int
		apiCall            func(client DatabricksClient) error
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
			apiCall: func(client DatabricksClient) error {
				return client.Post("/token/create", map[string]string{
					"foo": "bar",
				}, nil)
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
			apiCall: func(client DatabricksClient) error {
				return client.Post("/token/create", map[string]string{
					"foo": "bar",
				}, nil)
			},
		},
		{
			name:               "Invalid HTML Status 404",
			response:           `<html> Hello world </html>`,
			responseStatus:     http.StatusNotFound,
			expectedErrorCode:  "NOT_FOUND",
			expectedMessage:    "Response from server (404 Not Found) <html> Hello world </html>: invalid character '<' looking for beginning of value",
			expectedResource:   "/api/2.0/token/create",
			expectedStatusCode: 404,
			apiCall: func(client DatabricksClient) error {
				return client.Post("/token/create", map[string]string{
					"foo": "bar",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(tt.responseStatus)
				_, err := rw.Write([]byte(tt.response))
				assert.NoError(t, err, err)
			}))
			// Close the server when test finishes
			defer server.Close()
			client := DatabricksClient{
				Host:  server.URL,
				Token: "...",
			}
			err := client.ConfigureWithAuthentication()
			assert.NoError(t, err)

			err = tt.apiCall(client)
			t.Log(err)
			assert.IsType(t, APIError{}, err)
			assert.Equal(t, tt.expectedErrorCode, err.(APIError).ErrorCode, "error code is not the same")
			assert.Equal(t, tt.expectedMessage, err.(APIError).Message, "message is not the same")
			assert.Equal(t, tt.expectedResource, err.(APIError).Resource, "resource is not the same")
			assert.Equal(t, tt.expectedStatusCode, err.(APIError).StatusCode, "status code is not the same")
		})
	}
}
