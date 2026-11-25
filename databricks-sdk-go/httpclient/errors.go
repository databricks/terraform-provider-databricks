package httpclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/databricks/databricks-sdk-go/common"
)

type HttpError struct {
	*http.Response
	Message string
	err     error
}

func (r *HttpError) Unwrap() error {
	return r.err
}

// HTTPStatusCode returns the HTTP status code of the error response.
func (r *HttpError) HTTPStatusCode() int {
	return r.Response.StatusCode
}

// Header returns the HTTP headers of the error response.
func (r *HttpError) Header() http.Header {
	return r.Response.Header
}

func (r *HttpError) Error() string {
	return fmt.Sprintf("http %d: %s", r.StatusCode, r.Message)
}

// DefaultErrorMapper returns *HttpError
func DefaultErrorMapper(ctx context.Context, resp common.ResponseWrapper) error {
	if resp.Response.StatusCode < 400 {
		return nil
	}
	raw, err := io.ReadAll(resp.ReadCloser)
	if err != nil {
		return &HttpError{
			Response: resp.Response,
			Message:  "failed to read response",
			err:      err,
		}
	}
	return &HttpError{
		Response: resp.Response,
		Message:  string(raw),
	}
}

func DefaultErrorRetriable(ctx context.Context, err error) bool {
	var httpError *HttpError
	if errors.As(err, &httpError) {
		if httpError.StatusCode == http.StatusTooManyRequests {
			return true
		}
		if httpError.StatusCode == http.StatusGatewayTimeout {
			return true
		}
	}
	return false
}

var urlErrorTransientErrorMessages = []string{
	"connection reset by peer",
	"TLS handshake timeout",
	"connection refused",
	"Unexpected error",
	"i/o timeout",
}

func isRetriableUrlError(err error) bool {
	var urlError *url.Error
	if !errors.As(err, &urlError) {
		return false
	}
	for _, msg := range urlErrorTransientErrorMessages {
		if strings.Contains(err.Error(), msg) {
			return true
		}
	}
	return false
}
