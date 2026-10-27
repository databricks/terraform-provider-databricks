package qa

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPFixtureHandlerReplacesConsumedPhases(t *testing.T) {
	handler := NewHTTPFixtureHandler(t)
	server := httptest.NewServer(handler)
	defer server.Close()

	handler.SetFixtures([]HTTPFixture{{
		Method:   http.MethodGet,
		Resource: "/first",
		Response: "first",
	}})
	assertFixtureResponse(t, server, "/first", "first")

	handler.SetFixtures([]HTTPFixture{{
		Method:   http.MethodGet,
		Resource: "/second",
		Response: "second",
	}})
	assertFixtureResponse(t, server, "/second", "second")
	handler.AssertAllFixturesConsumed()
}

func TestHTTPFixtureHandlerReusesExplicitFixture(t *testing.T) {
	handler := NewHTTPFixtureHandler(t)
	server := httptest.NewServer(handler)
	defer server.Close()
	handler.SetFixtures([]HTTPFixture{{
		Method:       http.MethodGet,
		Resource:     "/reusable",
		Response:     "response",
		ReuseRequest: true,
	}})

	for range 2 {
		response, err := server.Client().Get(server.URL + "/reusable")
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		require.NoError(t, response.Body.Close())
	}
	handler.AssertAllFixturesConsumed()
}

func assertFixtureResponse(t *testing.T, server *httptest.Server, path, expected string) {
	t.Helper()
	response, err := server.Client().Get(server.URL + path)
	require.NoError(t, err)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, expected, string(body))
}
