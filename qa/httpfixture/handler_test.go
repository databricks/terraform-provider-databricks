package httpfixture

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func call(method, path string, body any, result string) Call {
	c := Call{Request: Request{Method: method, Path: path}, Response: Response{Status: 200, Body: &Body{JSON: result}}}
	if body != nil {
		c.Request.Body = &Body{JSON: body}
	}
	return c
}

func send(h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(method, path, strings.NewReader(body)))
	return w
}

func TestMatchesWholeSignatureAndConsumesFIFO(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("step 1 (update)", []Call{
		call("POST", "/groups", map[string]any{"name": "a"}, "a-first"),
		call("POST", "/groups", map[string]any{"name": "a"}, "a-second"),
		call("POST", "/groups", map[string]any{"name": "b"}, "b"),
	}))
	require.Equal(t, `"b"`, send(h, "POST", "/groups", `{"name":"b"}`).Body.String())
	require.Equal(t, `"a-first"`, send(h, "POST", "/groups", `{"name":"a"}`).Body.String())
	require.Equal(t, `"a-second"`, send(h, "POST", "/groups", `{"name":"a"}`).Body.String())
	require.NoError(t, h.Verify())
}

func TestSemanticJSONAndQueryMatching(t *testing.T) {
	h := NewHandler()
	c := call("POST", "/groups", map[string]any{"number": 1, "values": []any{true, "a"}}, "ok")
	c.Request.Query = map[string][]string{"a": {"one"}, "b": {"two"}}
	c.Request.Headers = map[string]string{"X-Test": "value"}
	require.NoError(t, h.StartStep("create", []Call{c}))
	req := httptest.NewRequest("POST", "/groups?b=two&a=one", strings.NewReader(" {\"values\": [true, \"a\"], \"number\": 1.0} "))
	req.Header.Set("X-Test", "value")
	req.Header.Set("X-Unasserted", "ignored")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	require.NoError(t, h.Verify())
}

func TestUnexpectedRequestsDoNotConsumeFixtures(t *testing.T) {
	for _, body := range []string{`{"values":[2,1]}`, `null`, "", "not-json", `{"values":[1,2]} {}`} {
		t.Run(body, func(t *testing.T) {
			h := NewHandler()
			require.NoError(t, h.StartStep("step 2 (update)", []Call{call("POST", "/groups", map[string]any{"values": []int{1, 2}}, "ok")}))
			require.Equal(t, 400, send(h, "POST", "/groups", body).Code)
			require.ErrorContains(t, h.Verify(), "step 2 (update)")
			require.ErrorContains(t, h.Verify(), "unconsumed call 1")
		})
	}
}

func TestAbsentBodyDiffersFromJSONNull(t *testing.T) {
	h := NewHandler()
	nullCall := call("POST", "/groups", nil, "null")
	nullCall.Request.Body = &Body{JSON: nil}
	require.NoError(t, h.StartStep("bodies", []Call{call("POST", "/groups", nil, "absent"), nullCall}))
	require.Equal(t, `"null"`, send(h, "POST", "/groups", "null").Body.String())
	require.Equal(t, `"absent"`, send(h, "POST", "/groups", "").Body.String())
	require.NoError(t, h.Verify())
}

func TestReusableMetadataAndStepBoundaries(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.SetReusableFixtures([]Call{call("GET", "/metadata", nil, "metadata")}))
	require.NoError(t, h.StartStep("create", []Call{call("GET", "/groups/1", nil, "old")}))
	require.ErrorContains(t, h.Verify(), "create: unconsumed")
	require.Equal(t, `"old"`, send(h, "GET", "/groups/1", "").Body.String())
	require.NoError(t, h.StartStep("update", []Call{call("GET", "/groups/1", nil, "new")}))
	for range 3 {
		require.Equal(t, `"metadata"`, send(h, "GET", "/metadata", "").Body.String())
	}
	require.Equal(t, `"new"`, send(h, "GET", "/groups/1", "").Body.String())
	require.NoError(t, h.Verify())
}

func TestFailedStepCanInstallDestroyFixtures(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("create", []Call{call("POST", "/groups", nil, "id")}))
	require.ErrorContains(t, h.StartStep("destroy", []Call{call("DELETE", "/groups/id", nil, "deleted")}), "create: unconsumed")
	require.Equal(t, `"deleted"`, send(h, "DELETE", "/groups/id", "").Body.String())
	require.NoError(t, h.Verify())
}

func TestConcurrentRequestsConsumeExactlyOnce(t *testing.T) {
	h := NewHandler()
	calls := make([]Call, 32)
	for i := range calls {
		calls[i] = call("GET", "/groups/1", nil, fmt.Sprint(i))
	}
	require.NoError(t, h.StartStep("concurrent", calls))
	responses := make(chan string, len(calls))
	var workers sync.WaitGroup
	for range calls {
		workers.Go(func() { responses <- send(h, "GET", "/groups/1", "").Body.String() })
	}
	workers.Wait()
	close(responses)
	seen := map[string]bool{}
	for response := range responses {
		require.False(t, seen[response], "fixture consumed twice")
		seen[response] = true
	}
	require.Len(t, seen, len(calls))
	require.NoError(t, h.Verify())
}

func TestRejectInvalidFixtures(t *testing.T) {
	for _, c := range []Call{
		{Request: Request{Path: "/groups"}, Response: Response{Status: 200}},
		{Request: Request{Method: "GET", Path: "groups"}, Response: Response{Status: 200}},
		{Request: Request{Method: "GET", Path: "/groups?a=b"}, Response: Response{Status: 200}},
		{Request: Request{Method: "GET", Path: "/groups"}},
	} {
		h := NewHandler()
		require.Error(t, h.SetReusableFixtures([]Call{c}))
		require.Error(t, h.StartStep("invalid", []Call{c}))
	}
}

func TestJSONNumbersPreserveLargeIntegers(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("numbers", []Call{
		call("POST", "/numbers", json.Number("9007199254740993"), "large"),
		call("POST", "/numbers", json.Number("9007199254740992"), "neighbor"),
		call("POST", "/numbers", 1000, "exponent"),
	}))
	require.Equal(t, `"neighbor"`, send(h, "POST", "/numbers", "9007199254740992").Body.String())
	require.Equal(t, `"large"`, send(h, "POST", "/numbers", "9007199254740993").Body.String())
	require.Equal(t, `"exponent"`, send(h, "POST", "/numbers", "1e3").Body.String())
	require.NoError(t, h.Verify())
}

func TestEscapedPathsRemainDistinct(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("paths", []Call{
		call("GET", "/resources/a%2Fb", nil, "escaped"),
		call("GET", "/resources/a/b", nil, "nested"),
	}))
	require.Equal(t, `"nested"`, send(h, "GET", "/resources/a/b", "").Body.String())
	require.Equal(t, `"escaped"`, send(h, "GET", "/resources/a%2Fb", "").Body.String())
	require.NoError(t, h.Verify())
}

func TestReusableFixturesCanBeReplacedWithoutResettingStep(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.SetReusableFixtures([]Call{call("GET", "/metadata", nil, "first")}))
	require.NoError(t, h.StartStep("create", []Call{call("GET", "/metadata", nil, "step")}))
	require.Equal(t, `"step"`, send(h, "GET", "/metadata", "").Body.String())
	require.Equal(t, `"first"`, send(h, "GET", "/metadata", "").Body.String())
	require.NoError(t, h.SetReusableFixtures([]Call{call("GET", "/metadata", nil, "replacement")}))
	require.Equal(t, `"replacement"`, send(h, "GET", "/metadata", "").Body.String())
	require.NoError(t, h.Verify())
	require.NoError(t, h.SetReusableFixtures(nil))
	require.Equal(t, 400, send(h, "GET", "/metadata", "").Code)
	require.ErrorContains(t, h.Verify(), "create: unexpected HTTP call")
}

func TestInvalidReusableReplacementKeepsPreviousFixturesAndFailures(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.SetReusableFixtures([]Call{call("GET", "/metadata", nil, "original")}))
	require.NoError(t, h.StartStep("create", nil))
	require.Equal(t, 400, send(h, "GET", "/unexpected", "").Code)
	invalid := call("GET", "/metadata", nil, "invalid")
	invalid.Response.Status = 0
	require.Error(t, h.SetReusableFixtures([]Call{call("GET", "/metadata", nil, "partial"), invalid}))
	require.Equal(t, `"original"`, send(h, "GET", "/metadata", "").Body.String())
	require.NoError(t, h.SetReusableFixtures([]Call{call("GET", "/metadata", nil, "valid")}))
	require.Equal(t, `"valid"`, send(h, "GET", "/metadata", "").Body.String())
	require.ErrorContains(t, h.Verify(), "create: unexpected HTTP call")
}
