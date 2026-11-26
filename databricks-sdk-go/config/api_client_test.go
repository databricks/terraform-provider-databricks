package config

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStaticHeaders(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost", nil)
	require.NoError(t, err)
	err = StaticHeaders(map[string]string{
		"x-databricks-traffic-id": "testenv://liteswap/dms",
	})(req)
	require.NoError(t, err)
	assert.Equal(t, "testenv://liteswap/dms", req.Header.Get("x-databricks-traffic-id"))
}

// TestHeaders_Dynamic verifies the Headers hook is invoked on every request, so
// a dynamic function produces a fresh value each time.
func TestHeaders_Dynamic(t *testing.T) {
	var n int
	cfg := &Config{
		Host:  "http://localhost",
		Token: "x",
		Headers: func(r *http.Request) error {
			n++
			r.Header.Set("x-request-id", fmt.Sprintf("req-%d", n))
			return nil
		},
		Loaders:       []Loader{mockLoader(func(*Config) error { return nil })},
		HTTPTransport: metadataNotFoundTransport,
	}
	clientCfg, err := HTTPClientConfigFromConfig(cfg)
	require.NoError(t, err)

	apply := func() *http.Request {
		req, err := http.NewRequest(http.MethodGet, cfg.Host, nil)
		require.NoError(t, err)
		// AuthVisitor runs first in the real request path; later visitors
		// depend on it (e.g. AuthType for the user agent).
		require.NoError(t, clientCfg.AuthVisitor(req))
		for _, v := range clientCfg.Visitors {
			require.NoError(t, v(req))
		}
		return req
	}

	req1, req2 := apply(), apply()
	assert.Equal(t, "req-1", req1.Header.Get("x-request-id"))
	assert.Equal(t, "req-2", req2.Header.Get("x-request-id"))
}
