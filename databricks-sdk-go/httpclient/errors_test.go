package httpclient

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleRequestAPIError(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
				Body: io.NopCloser(strings.NewReader(`{
					"error_code": "NOT_FOUND",
					"message": "Something was not found"
				}`)),
				Request: r,
			}, nil
		}),
	})
	err := c.Do(context.Background(), "PATCH", "/a", WithRequestData(map[string]any{}))
	var httpErr *HttpError
	if assert.ErrorAs(t, err, &httpErr) {
		require.Equal(t, 400, httpErr.StatusCode)
	}
}

func TestSimpleRequestErrReaderBody(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(false),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "PATCH", "/a", WithRequestHeaders(headers), WithRequestData(map[string]any{}))
	require.EqualError(t, err, "response body: test error")
}
