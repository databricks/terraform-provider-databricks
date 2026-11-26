package httplog

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoHeadersNoBody(t *testing.T) {
	res := RoundTripStringer{
		Request: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path:     "/",
				RawQuery: "foo=bar&baz=qux",
			},
		},
		Response: &http.Response{
			Status: "200 OK",
			Proto:  "HTTP/1.1",
		},
	}.String()
	assert.Equal(t, `GET /?foo=bar&baz=qux
< HTTP/1.1 200 OK`, res)
}

func TestRequestAndResponseHaveHeadersAndBody(t *testing.T) {
	res := RoundTripStringer{
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{Path: "/"},
			Header: http.Header{
				"Foo": []string{"bar"},
				"Bar": []string{"baz"},
			},
		},
		Response: &http.Response{
			Status: "200 OK",
			Proto:  "HTTP/1.1",
			Header: http.Header{
				"Response-Foo": []string{"bar"},
				"Response-Bar": []string{"baz"},
			},
		},
		RequestBody:        []byte("request-hello"),
		ResponseBody:       []byte("response-hello"),
		DebugHeaders:       true,
		DebugTruncateBytes: 100,
	}.String()
	assert.Equal(t, `GET /
> * Host: 
> * Bar: baz
> * Foo: bar
> request-hello
< HTTP/1.1 200 OK
< * Response-Bar: baz
< * Response-Foo: bar
< response-hello`, res)
}

func TestDoNotPrintHeadersWhenNotConfigured(t *testing.T) {
	res := RoundTripStringer{
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{Path: "/"},
			Header: http.Header{
				"Foo": []string{"bar"},
				"Bar": []string{"baz"},
			},
		},
		Response: &http.Response{
			Status: "200 OK",
			Proto:  "HTTP/1.1",
			Header: http.Header{
				"Response-Foo": []string{"bar"},
				"Response-Bar": []string{"baz"},
			},
		},
		RequestBody:        []byte("request-hello"),
		ResponseBody:       []byte("response-hello"),
		DebugHeaders:       false,
		DebugTruncateBytes: 100,
	}.String()
	assert.Equal(t, `GET /
> request-hello
< HTTP/1.1 200 OK
< response-hello`, res)
}

func TestHideAuthorizationHeaderWhenConfigured(t *testing.T) {
	res := RoundTripStringer{
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{Path: "/"},
			Header: http.Header{
				"Foo":                                    []string{"bar"},
				"Authorization":                          []string{"baz"},
				"X-Databricks-Azure-SP-Management-Token": []string{"open sesame"},
				"X-Databricks-GCP-SA-Access-Token":       []string{"alohamora"},
			},
		},
		Response: &http.Response{
			Status: "200 OK",
			Proto:  "HTTP/1.1",
		},
		RequestBody:              []byte("request-hello"),
		ResponseBody:             []byte("response-hello"),
		DebugHeaders:             true,
		DebugTruncateBytes:       100,
		DebugAuthorizationHeader: false,
	}.String()
	assert.Equal(t, `GET /
> * Host: 
> * Authorization: REDACTED
> * Foo: bar
> * X-Databricks-Azure-SP-Management-Token: REDACTED
> * X-Databricks-GCP-SA-Access-Token: REDACTED
> request-hello
< HTTP/1.1 200 OK
< response-hello`, res)
}

func TestNilResponse(t *testing.T) {
	res := RoundTripStringer{
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{Path: "/"},
		},
		Err: &url.Error{
			Op:  "Get",
			URL: "http://example.com",
			Err: errors.New("request timed out after 1m0s of inactivity"),
		},
	}.String()
	assert.Equal(t, `GET /
< Error: Get "http://example.com": request timed out after 1m0s of inactivity`, res)
}

func TestFailureToConsumeResponse(t *testing.T) {
	res := RoundTripStringer{
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{Path: "/"},
		},
		Response: &http.Response{
			Status: "200 OK",
			Proto:  "HTTP/1.1",
			Header: http.Header{
				"Foo": []string{"bar"},
			},
		},
		Err:          errors.New("failed to read response body"),
		DebugHeaders: true,
	}.String()
	assert.Equal(t, `GET /
> * Host: 
< HTTP/1.1 200 OK (Error: failed to read response body)
< * Foo: ... (3 more bytes)`, res)
}
