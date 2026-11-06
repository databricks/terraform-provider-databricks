package httpfixture

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRefreshPreservesMatchersAcrossInterleavingAndRepeatedCalls(t *testing.T) {
	old := Call{Request: Request{Method: "POST", Path: "/items", Query: map[string][]string{"request_id": {"old"}},
		Body: &Body{JSON: map[string]any{"name": "old"}}, Matchers: Matchers{
			Query: map[string]Matcher{"request_id": AnyString}, Headers: map[string]Matcher{"X-Request-ID": AnyString},
			JSON: []JSONMatcher{{Path: []any{"name"}, Type: AnyString}},
		}}, Response: Response{Status: 200}}
	newCall := func() Call {
		return Call{Request: Request{Method: "POST", Path: "/items", Query: map[string][]string{"request_id": {"new"}},
			Body: &Body{JSON: map[string]any{"name": "new"}}, Headers: map[string]string{"X-Request-Id": "new-id", "User-Agent": "sdk"}},
			Response: Response{Status: 201}}
	}
	calls := []Call{{Request: Request{Method: "GET", Path: "/other"}}, newCall(), newCall(), newCall()}
	require.NoError(t, PreserveMatchers(calls, []Call{old, old}))
	for _, call := range calls[1:] {
		require.Equal(t, old.Request.Matchers, call.Request.Matchers)
		require.Equal(t, "new-id", call.Request.Headers["X-Request-Id"])
		require.NotContains(t, call.Request.Headers, "User-Agent")
		require.Equal(t, "new", call.Request.Body.JSON.(map[string]any)["name"])
	}
	require.Empty(t, calls[0].Request.Matchers)
	require.ErrorContains(t, PreserveMatchers(calls[:1], []Call{old}), "could not preserve")
}
