package httpfixture

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func repeated(c Call, times Times, minGap, maxGap string) Call {
	c.Times, c.MinGap, c.MaxGap = times, minGap, maxGap
	return c
}

func TestTimesMatchesACallNTimesInARow(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("step 1 (create)", []Call{repeated(call("GET", "/retry", nil, "busy"), Exactly(3), "", "")}))
	for range 3 {
		require.Equal(t, `"busy"`, send(h, "GET", "/retry", "").Body.String())
	}
	require.NoError(t, h.Verify())
	require.Equal(t, 400, send(h, "GET", "/retry", "").Code)
	require.ErrorContains(t, h.Verify(), "step 1 (create): unexpected HTTP call: GET /retry")
}

func TestTimesBelowMinimumNamesTheCount(t *testing.T) {
	for _, tc := range []struct {
		times   Times
		message string
	}{
		{Exactly(3), "step 1 (create): unconsumed call 1: GET /retry: expected 3 times, matched 1"},
		{Times{Min: 2, Max: 3}, "step 1 (create): unconsumed call 1: GET /retry: expected 2 to 3 times, matched 1"},
	} {
		h := NewHandler()
		require.NoError(t, h.StartStep("step 1 (create)", []Call{repeated(call("GET", "/retry", nil, "busy"), tc.times, "", "")}))
		require.Equal(t, 200, send(h, "GET", "/retry", "").Code)
		err := h.Verify()
		require.Error(t, err)
		require.Equal(t, tc.message, err.Error())
	}
}

func TestBoundedTimesPassesAnywhereInTheRange(t *testing.T) {
	for _, sent := range []int{2, 3} {
		h := NewHandler()
		require.NoError(t, h.StartStep("step 1 (create)", []Call{repeated(call("GET", "/retry", nil, "busy"), Times{Min: 2, Max: 3}, "", "")}))
		for range sent {
			require.Equal(t, 200, send(h, "GET", "/retry", "").Code)
		}
		require.NoError(t, h.Verify(), "%d calls", sent)
	}
	h := NewHandler()
	require.NoError(t, h.StartStep("step 1 (optional)", []Call{repeated(call("GET", "/retry", nil, "busy"), Times{Min: 0, Max: 1}, "", "")}))
	require.NoError(t, h.Verify(), "a minimum of 0 makes the call optional")
}

// A range is greedy: an identical signature listed after it answers only once
// the range reaches its maximum.
func TestBoundedTimesFillsItsMaximumBeforeAnIdenticalCall(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("step 1 (create)", []Call{
		repeated(call("GET", "/retry", nil, "busy"), Times{Min: 1, Max: 2}, "", ""),
		call("GET", "/retry", nil, "done"),
	}))
	for _, expected := range []string{`"busy"`, `"busy"`, `"done"`} {
		require.Equal(t, expected, send(h, "GET", "/retry", "").Body.String())
	}
	require.NoError(t, h.Verify())
}

func TestOrderedTimesMustFinishBeforeTheNextCall(t *testing.T) {
	h := NewHandler()
	h.SetEnforceCallOrder(true)
	require.NoError(t, h.StartStep("step 1 (create)", []Call{
		repeated(call("POST", "/groups", nil, "busy"), Exactly(2), "", ""),
		call("GET", "/groups/1", nil, "read"),
	}))
	require.Equal(t, 200, send(h, "POST", "/groups", "").Code)
	response := send(h, "GET", "/groups/1", "")
	require.Equal(t, 400, response.Code)
	require.Contains(t, response.Body.String(), "matched call 2; next expected call 1: POST /groups")
	require.Equal(t, 200, send(h, "POST", "/groups", "").Code)
	require.Equal(t, `"read"`, send(h, "GET", "/groups/1", "").Body.String())
	require.ErrorContains(t, h.Verify(), "out-of-order")
}

// Once ordered replay moves past a satisfied range, the range cannot repeat.
func TestOrderedBoundedTimesClosesWhenTheNextCallArrives(t *testing.T) {
	h := NewHandler()
	h.SetEnforceCallOrder(true)
	require.NoError(t, h.StartStep("step 1 (create)", []Call{
		repeated(call("POST", "/groups", nil, "busy"), Times{Min: 1, Max: 3}, "", ""),
		call("GET", "/groups/1", nil, "read"),
	}))
	require.Equal(t, 200, send(h, "POST", "/groups", "").Code)
	require.Equal(t, 200, send(h, "POST", "/groups", "").Code)
	require.Equal(t, `"read"`, send(h, "GET", "/groups/1", "").Body.String())
	require.NoError(t, h.Verify())
	require.Equal(t, 400, send(h, "POST", "/groups", "").Code)
	require.ErrorContains(t, h.Verify(), "unexpected HTTP call: POST /groups")
}

// sendAfter waits on synctest's fake clock, so gaps are exact.
func sendAfter(h *Handler, wait time.Duration, path string) int {
	time.Sleep(wait)
	return send(h, "GET", path, "").Code
}

func TestGapBoundsMeasureFromThePreviousMatchedCall(t *testing.T) {
	for _, tc := range []struct {
		name, minGap, maxGap string
		wait                 time.Duration
		message              string
	}{
		{"within", "1s", "2s", 1500 * time.Millisecond, ""},
		{"below", "1s", "2s", 500 * time.Millisecond,
			"step 1 (retry): call 2: GET /second arrived 500ms after call 1 (GET /first), below min_gap 1s"},
		{"above", "", "2s", 3 * time.Second,
			"step 1 (retry): call 2: GET /second arrived 3s after call 1 (GET /first), above max_gap 2s"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				h := NewHandler()
				require.NoError(t, h.SetReusableFixtures([]Call{call("GET", "/metadata", nil, "metadata")}))
				require.NoError(t, h.StartStep("step 1 (retry)", []Call{
					call("GET", "/first", nil, "first"),
					repeated(call("GET", "/second", nil, "second"), Times{}, tc.minGap, tc.maxGap),
				}))
				require.Equal(t, 200, sendAfter(h, 0, "/first"))
				// Reusable calls are not matched step calls, so they do not move the reference.
				require.Equal(t, 200, sendAfter(h, tc.wait/2, "/metadata"))
				require.Equal(t, 200, sendAfter(h, tc.wait-tc.wait/2, "/second"), "an out-of-bounds call is still answered")
				err := h.Verify()
				if tc.message == "" {
					require.NoError(t, err)
					return
				}
				require.Error(t, err)
				require.Equal(t, tc.message, err.Error())
			})
		})
	}
}

// Repetitions measure from their own previous repetition, even when another
// call arrives in between, and the first repetition is not bounded.
func TestGapBoundsWithTimesMeasureBetweenRepetitions(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := NewHandler()
		require.NoError(t, h.StartStep("step 1 (retry)", []Call{
			repeated(call("GET", "/retry", nil, "busy"), Exactly(3), "1s", "2s"),
			call("GET", "/other", nil, "other"),
		}))
		require.Equal(t, 200, sendAfter(h, 0, "/retry"))
		require.Equal(t, 200, sendAfter(h, 900*time.Millisecond, "/other"))
		require.Equal(t, 200, sendAfter(h, 200*time.Millisecond, "/retry"))
		require.Equal(t, 200, sendAfter(h, 500*time.Millisecond, "/retry"))
		err := h.Verify()
		require.Error(t, err)
		require.Equal(t, "step 1 (retry): call 1: GET /retry repetition 3 arrived 500ms after repetition 2, below min_gap 1s", err.Error())
	})
}

// slowBody takes wait to deliver its bytes, as a large upload would.
type slowBody struct {
	wait time.Duration
	body io.Reader
}

func (s *slowBody) Read(p []byte) (int, error) {
	time.Sleep(s.wait)
	s.wait = 0
	return s.body.Read(p)
}

// A gap ends when a request arrives, so time spent reading its body neither
// breaks max_gap nor hides a min_gap violation.
func TestGapBoundsMeasureArrivalBeforeTheBody(t *testing.T) {
	for _, tc := range []struct {
		name, minGap, maxGap, message string
	}{
		{"within max_gap", "", "1s", ""},
		{"below min_gap", "1s", "", "step 1 (retry): call 2: POST /second arrived 500ms after call 1 (GET /first), below min_gap 1s"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				h := NewHandler()
				require.NoError(t, h.StartStep("step 1 (retry)", []Call{
					call("GET", "/first", nil, "first"),
					repeated(call("POST", "/second", map[string]any{"name": "a"}, "second"), Times{}, tc.minGap, tc.maxGap),
				}))
				require.Equal(t, 200, sendAfter(h, 0, "/first"))
				time.Sleep(500 * time.Millisecond)
				w := httptest.NewRecorder()
				h.ServeHTTP(w, httptest.NewRequest("POST", "/second", &slowBody{2 * time.Second, strings.NewReader(`{"name":"a"}`)}))
				require.Equal(t, 200, w.Code)
				err := h.Verify()
				if tc.message == "" {
					require.NoError(t, err)
					return
				}
				require.Error(t, err)
				require.Equal(t, tc.message, err.Error())
			})
		})
	}
}

// Concurrent requests can take the lock out of arrival order. The latest
// arrival stays the reference, so a later call is not measured from an earlier one.
func TestGapBoundsMeasureFromTheLatestArrival(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("step 1 (create)", []Call{
		call("GET", "/a", nil, "a"),
		call("GET", "/b", nil, "b"),
		repeated(call("GET", "/c", nil, "c"), Times{}, "1s", ""),
	}))
	start := time.Now()
	for _, arrival := range []struct {
		path  string
		after time.Duration
	}{{"/b", time.Second}, {"/a", 0}, {"/c", 1500 * time.Millisecond}} {
		w := httptest.NewRecorder()
		e, _, _ := h.consume(w, httptest.NewRequest("GET", arrival.path, nil), start.Add(arrival.after))
		require.NotNil(t, e, arrival.path)
	}
	err := h.Verify()
	require.Error(t, err)
	require.Equal(t, "step 1 (create): call 3: GET /c arrived 500ms after call 2 (GET /b), below min_gap 1s", err.Error())
}

func TestGapBoundsNeedAnEarlierCallInTheStep(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("step 1 (create)", []Call{call("GET", "/first", nil, "first")}))
	require.Equal(t, 200, send(h, "GET", "/first", "").Code)
	require.NoError(t, h.StartStep("step 2 (update)", []Call{repeated(call("GET", "/second", nil, "second"), Times{}, "1s", "")}))
	require.Equal(t, 200, send(h, "GET", "/second", "").Code)
	err := h.Verify()
	require.Error(t, err)
	require.Equal(t, "step 2 (update): call 1: GET /second arrived first in its step, so min_gap and max_gap have no earlier call to measure from", err.Error())
}

func TestRepeatFieldsValidateOnLoad(t *testing.T) {
	for _, tc := range []struct{ fields, message string }{
		{"times: 0", `call 1: times "0" must be a positive count or {min, max}`},
		{"times: often", `call 1: times "often" must be a positive count or {min, max}`},
		{"times: [1, 2]", "call 1: times must be a positive count or {min, max}"},
		{"times: {min: 2}", "call 1: times range requires both min and max"},
		{"times: {min: 1, max: 2, step: 1}", `call 1: unknown times field "step"; expected min and max`},
		{"times: {min: 1, min: 2, max: 3}", "call 1: times.min is set twice"},
		{"times: {min: one, max: 2}", `call 1: times.min "one" must be an integer`},
		{"times: {min: 3, max: 2}", "call 1: times {min: 3, max: 2} needs 0 <= min <= max and max >= 1"},
		{"times: {min: 0, max: 0}", "call 1: times {min: 0, max: 0} needs 0 <= min <= max and max >= 1"},
		{"min_gap: soon", `call 1: min_gap "soon" must be a positive duration, such as 1s`},
		{"max_gap: -1s", `call 1: max_gap "-1s" must be a positive duration, such as 1s`},
		{"min_gap: 2s, max_gap: 1s", "call 1: min_gap 2s exceeds max_gap 1s"},
		{"times: 1, min_gap: 1s", "call 1: with times, min_gap and max_gap bound the gaps between repetitions, so times needs a max of at least 2"},
	} {
		t.Run(tc.fields, func(t *testing.T) {
			var calls []Call
			require.NoError(t, yaml.Unmarshal([]byte("- {request: {method: GET, path: /retry}, response: {status: 429}, "+tc.fields+"}"), &calls))
			err := NewHandler().StartStep("step 1 (create)", calls)
			require.Error(t, err)
			require.Equal(t, "step 1 (create): "+tc.message, err.Error())
		})
	}
}

// Exactly(0) must not become the zero Times, which expects the call once.
func TestExactlyRejectsCountsBelowOne(t *testing.T) {
	for _, n := range []int{0, -2} {
		err := NewHandler().StartStep("step 1 (create)", []Call{repeated(call("GET", "/retry", nil, "busy"), Exactly(n), "", "")})
		require.Error(t, err)
		require.Equal(t, fmt.Sprintf("step 1 (create): call 1: times %d must be a positive count", n), err.Error())
	}
}

func TestReusableFixturesRejectRepeatFields(t *testing.T) {
	h := NewHandler()
	require.EqualError(t, h.SetReusableFixtures([]Call{repeated(call("GET", "/metadata", nil, "m"), Exactly(2), "", "")}),
		"reusable call 1: times, min_gap and max_gap apply only to step calls")
	require.EqualError(t, h.SetReusableFixtures([]Call{repeated(call("GET", "/metadata", nil, "m"), Times{}, "", "1s")}),
		"reusable call 1: times, min_gap and max_gap apply only to step calls")
}

func TestTimesYAMLRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		times   Times
		encoded string
	}{
		{Times{}, "{}\n"},
		{Exactly(3), "times: 3\n"},
		{Times{Min: 2, Max: 3}, "times:\n    min: 2\n    max: 3\n"},
	} {
		encoded, err := yaml.Marshal(struct {
			Times Times `yaml:"times,omitempty"`
		}{tc.times})
		require.NoError(t, err)
		require.Equal(t, tc.encoded, string(encoded))
		var decoded struct {
			Times Times `yaml:"times,omitempty"`
		}
		require.NoError(t, yaml.Unmarshal(encoded, &decoded))
		require.Equal(t, tc.times, decoded.Times)
	}
	var aliased struct {
		Base  Times `yaml:"base"`
		Times Times `yaml:"times"`
	}
	require.NoError(t, yaml.Unmarshal([]byte("base: &retries {min: 2, max: 3}\ntimes: *retries\n"), &aliased))
	require.Equal(t, Times{Min: 2, Max: 3}, aliased.Times, "journeys reuse values through YAML anchors")
}
