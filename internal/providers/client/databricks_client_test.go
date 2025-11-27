package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHttpTransportWrapper_RoundTrip_Headers(t *testing.T) {
	// Create a test server to capture the request
	var capturedHeaders http.Header
	var capturedPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		capturedPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create the transport wrapper with custom headers
	wrapper := &httpTransportWrapper{
		base: http.DefaultTransport,
		headers: map[string]string{
			"X-Custom-Header": "custom-value",
			"X-Actor-Id":      "actor-123",
		},
	}

	// Create a request and send it
	req, err := http.NewRequest("GET", server.URL+"/api/test", nil)
	require.NoError(t, err)

	client := &http.Client{Transport: wrapper}
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Verify headers were added
	assert.Equal(t, "custom-value", capturedHeaders.Get("X-Custom-Header"))
	assert.Equal(t, "actor-123", capturedHeaders.Get("X-Actor-Id"))
	assert.Equal(t, "/api/test", capturedPath)
}

func TestHttpTransportWrapper_RoundTrip_PathPrefix(t *testing.T) {
	// Create a test server to capture the request
	var capturedPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create the transport wrapper with path prefix
	wrapper := &httpTransportWrapper{
		base:       http.DefaultTransport,
		pathPrefix: "/proxy/env-01",
	}

	// Create a request and send it
	req, err := http.NewRequest("GET", server.URL+"/api/test", nil)
	require.NoError(t, err)

	client := &http.Client{Transport: wrapper}
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Verify path prefix was added
	assert.Equal(t, "/proxy/env-01/api/test", capturedPath)
}

func TestHttpTransportWrapper_RoundTrip_PathPrefixWithTrailingSlash(t *testing.T) {
	// Create a test server to capture the request
	var capturedPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create the transport wrapper with path prefix that has trailing slash
	wrapper := &httpTransportWrapper{
		base:       http.DefaultTransport,
		pathPrefix: "/proxy/env-01/",
	}

	// Create a request and send it
	req, err := http.NewRequest("GET", server.URL+"/api/test", nil)
	require.NoError(t, err)

	client := &http.Client{Transport: wrapper}
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Verify trailing slash is trimmed from prefix
	assert.Equal(t, "/proxy/env-01/api/test", capturedPath)
}

func TestHttpTransportWrapper_RoundTrip_HeadersAndPathPrefix(t *testing.T) {
	// Create a test server to capture the request
	var capturedHeaders http.Header
	var capturedPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		capturedPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create the transport wrapper with both headers and path prefix
	wrapper := &httpTransportWrapper{
		base: http.DefaultTransport,
		headers: map[string]string{
			"Authorization": "Bearer token123",
		},
		pathPrefix: "/environments/staging",
	}

	// Create a request and send it
	req, err := http.NewRequest("GET", server.URL+"/api/2.0/clusters/list", nil)
	require.NoError(t, err)

	client := &http.Client{Transport: wrapper}
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Verify both headers and path prefix were applied
	assert.Equal(t, "Bearer token123", capturedHeaders.Get("Authorization"))
	assert.Equal(t, "/environments/staging/api/2.0/clusters/list", capturedPath)
}

func TestHttpTransportWrapper_RoundTrip_EmptyConfig(t *testing.T) {
	// Create a test server to capture the request
	var capturedPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create the transport wrapper with no customizations
	wrapper := &httpTransportWrapper{
		base:    http.DefaultTransport,
		headers: nil,
	}

	// Create a request and send it
	req, err := http.NewRequest("GET", server.URL+"/api/test", nil)
	require.NoError(t, err)

	client := &http.Client{Transport: wrapper}
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Verify no modifications
	assert.Equal(t, "/api/test", capturedPath)
}
