package httpfixture

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnforcedOrderPreservesFIFOAndReusableFallbacks(t *testing.T) {
	h := NewHandler()
	h.SetEnforceCallOrder(true)
	require.NoError(t, h.SetReusableFixtures([]Call{call("GET", "/metadata", nil, "metadata")}))
	require.NoError(t, h.StartStep("step 1 (create)", []Call{
		call("POST", "/groups", map[string]any{"name": "a"}, "created"),
		call("GET", "/groups/1", nil, "first read"),
		call("GET", "/groups/1", nil, "second read"),
	}))
	require.Equal(t, `"metadata"`, send(h, "GET", "/metadata", "").Body.String())
	require.Equal(t, `"created"`, send(h, "POST", "/groups", `{"name":"a"}`).Body.String())
	require.Equal(t, `"first read"`, send(h, "GET", "/groups/1", "").Body.String())
	require.Equal(t, `"metadata"`, send(h, "GET", "/metadata", "").Body.String())
	require.Equal(t, `"second read"`, send(h, "GET", "/groups/1", "").Body.String())
	require.NoError(t, h.Verify())
	require.NoError(t, h.StartStep("step 2 (destroy)", []Call{call("DELETE", "/groups/1", nil, "deleted")}))
	require.Equal(t, `"deleted"`, send(h, "DELETE", "/groups/1", "").Body.String())
	require.NoError(t, h.Verify())
}

func TestOutOfOrderCallFailsWithoutConsumptionOrReusableBypass(t *testing.T) {
	h := NewHandler()
	h.SetEnforceCallOrder(true)
	require.NoError(t, h.SetReusableFixtures([]Call{call("GET", "/groups/1", nil, "fallback")}))
	require.NoError(t, h.StartStep("step 1 (create)", []Call{
		call("POST", "/groups", nil, "created"),
		call("GET", "/groups/1", nil, "read"),
	}))
	response := send(h, "GET", "/groups/1", "")
	require.Equal(t, 400, response.Code)
	require.Contains(t, response.Body.String(), "step 1 (create): out-of-order HTTP call")
	require.Contains(t, response.Body.String(), "matched call 2; next expected call 1: POST /groups")
	require.ErrorContains(t, h.Verify(), "unconsumed call 2")
	require.Equal(t, `"created"`, send(h, "POST", "/groups", "").Body.String())
	require.Equal(t, `"read"`, send(h, "GET", "/groups/1", "").Body.String())
	require.ErrorContains(t, h.Verify(), "out-of-order")
	// Cleanup still installs its own ordered transcript after a failed step.
	require.ErrorContains(t, h.StartStep("destroy", []Call{call("DELETE", "/groups/1", nil, "deleted")}), "out-of-order")
	require.Equal(t, `"deleted"`, send(h, "DELETE", "/groups/1", "").Body.String())
	require.NoError(t, h.Verify())
}

func TestOrderSetterPreservesConsumedCallsAndFailures(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("update", []Call{
		call("GET", "/first", nil, "first"),
		call("GET", "/second", nil, "second"),
		call("GET", "/third", nil, "third"),
	}))
	require.Equal(t, 200, send(h, "GET", "/second", "").Code)
	h.SetEnforceCallOrder(true)
	require.Equal(t, 400, send(h, "GET", "/third", "").Code)
	h.SetEnforceCallOrder(false)
	require.Equal(t, `"third"`, send(h, "GET", "/third", "").Body.String())
	require.Equal(t, `"first"`, send(h, "GET", "/first", "").Body.String())
	require.ErrorContains(t, h.Verify(), "out-of-order")
}

func TestOrderedMatchingDistinguishesBodies(t *testing.T) {
	h := NewHandler()
	h.SetEnforceCallOrder(true)
	require.NoError(t, h.StartStep("create", []Call{
		call("POST", "/groups", map[string]any{"name": "a"}, "a"),
		call("POST", "/groups", map[string]any{"name": "b"}, "b"),
	}))
	require.Equal(t, 400, send(h, "POST", "/groups", `{"name":"b"}`).Code)
	require.Equal(t, `"a"`, send(h, "POST", "/groups", `{"name":"a"}`).Body.String())
	require.Equal(t, `"b"`, send(h, "POST", "/groups", `{"name":"b"}`).Body.String())
	require.ErrorContains(t, h.Verify(), "out-of-order")
}
