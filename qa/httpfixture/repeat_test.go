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

// retry is a GET /retry call answered "busy".
func retry(times Times, minGap, maxGap string) Call {
	return repeated(call("GET", "/retry", nil, "busy"), times, minGap, maxGap)
}

func startStep(t *testing.T, h *Handler, label string, calls ...Call) {
	t.Helper()
	require.NoError(t, h.StartStep(label, calls))
}

func requireFailure(t *testing.T, h *Handler, message string) {
	t.Helper()
	err := h.Verify()
	require.Error(t, err)
	require.Equal(t, message, err.Error())
}

func TestTimesMatchesACallNTimesInARow(t *testing.T) {
	h := NewHandler()
	startStep(t, h, "step 1 (create)", retry(Exactly(3), "", ""))
	for range 3 {
		require.Equal(t, `"busy"`, send(h, "GET", "/retry", "").Body.String())
	}
	require.NoError(t, h.Verify())
	require.Equal(t, 400, send(h, "GET", "/retry", "").Code)
	require.ErrorContains(t, h.Verify(),
		"step 1 (create): unexpected HTTP call: GET /retry")
}

func TestTimesBelowMinimumNamesTheCount(t *testing.T) {
	const unconsumed = "step 1 (create): unconsumed call 1: GET /retry: "
	for _, tc := range []struct {
		times   Times
		message string
	}{
		{Exactly(3), unconsumed + "expected 3 times, matched 1"},
		{Times{Min: 2, Max: 3}, unconsumed + "expected 2 to 3 times, matched 1"},
	} {
		h := NewHandler()
		startStep(t, h, "step 1 (create)", retry(tc.times, "", ""))
		require.Equal(t, 200, send(h, "GET", "/retry", "").Code)
		requireFailure(t, h, tc.message)
	}
}

func TestBoundedTimesPassesAnywhereInTheRange(t *testing.T) {
	for _, sent := range []int{2, 3} {
		h := NewHandler()
		startStep(t, h, "step 1 (create)", retry(Times{Min: 2, Max: 3}, "", ""))
		for range sent {
			require.Equal(t, 200, send(h, "GET", "/retry", "").Code)
		}
		require.NoError(t, h.Verify(), "%d calls", sent)
	}
	h := NewHandler()
	startStep(t, h, "step 1 (optional)", retry(Times{Min: 0, Max: 1}, "", ""))
	require.NoError(t, h.Verify(), "a minimum of 0 makes the call optional")
}

// A range is greedy: an identical signature listed after it answers only once
// the range reaches its maximum.
func TestBoundedTimesFillsItsMaximumBeforeAnIdenticalCall(t *testing.T) {
	h := NewHandler()
	startStep(t, h, "step 1 (create)",
		retry(Times{Min: 1, Max: 2}, "", ""), call("GET", "/retry", nil, "done"))
	for _, expected := range []string{`"busy"`, `"busy"`, `"done"`} {
		require.Equal(t, expected, send(h, "GET", "/retry", "").Body.String())
	}
	require.NoError(t, h.Verify())
}

func TestOrderedTimesMustFinishBeforeTheNextCall(t *testing.T) {
	h := NewHandler()
	h.SetEnforceCallOrder(true)
	startStep(t, h, "step 1 (create)",
		repeated(call("POST", "/groups", nil, "busy"), Exactly(2), "", ""),
		call("GET", "/groups/1", nil, "read"))
	require.Equal(t, 200, send(h, "POST", "/groups", "").Code)
	response := send(h, "GET", "/groups/1", "")
	require.Equal(t, 400, response.Code)
	require.Contains(t, response.Body.String(),
		"matched call 2; next expected call 1: POST /groups")
	require.Equal(t, 200, send(h, "POST", "/groups", "").Code)
	require.Equal(t, `"read"`, send(h, "GET", "/groups/1", "").Body.String())
	require.ErrorContains(t, h.Verify(), "out-of-order")
}

// Once ordered replay moves past a satisfied range, the range cannot repeat.
func TestOrderedBoundedTimesClosesWhenTheNextCallArrives(t *testing.T) {
	h := NewHandler()
	h.SetEnforceCallOrder(true)
	startStep(t, h, "step 1 (create)",
		repeated(call("POST", "/groups", nil, "busy"), Times{Min: 1, Max: 3}, "", ""),
		call("GET", "/groups/1", nil, "read"))
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
	const second = "step 1 (retry): call 2: GET /second arrived "
	for _, tc := range []struct {
		name, minGap, maxGap string
		wait                 time.Duration
		message              string
	}{
		{"within", "1s", "2s", 1500 * time.Millisecond, ""},
		{"below", "1s", "2s", 500 * time.Millisecond,
			second + "500ms after call 1 (GET /first), below min_gap 1s"},
		{"above", "", "2s", 3 * time.Second,
			second + "3s after call 1 (GET /first), above max_gap 2s"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				h := NewHandler()
				metadata := call("GET", "/metadata", nil, "metadata")
				require.NoError(t, h.SetReusableFixtures([]Call{metadata}))
				second := call("GET", "/second", nil, "second")
				startStep(t, h, "step 1 (retry)", call("GET", "/first", nil, "first"),
					repeated(second, Times{}, tc.minGap, tc.maxGap))
				require.Equal(t, 200, sendAfter(h, 0, "/first"))
				// Reusable calls are not matched step calls, so they leave the
				// reference alone.
				require.Equal(t, 200, sendAfter(h, tc.wait/2, "/metadata"))
				require.Equal(t, 200, sendAfter(h, tc.wait-tc.wait/2, "/second"),
					"an out-of-bounds call is still answered")
				if tc.message == "" {
					require.NoError(t, h.Verify())
					return
				}
				requireFailure(t, h, tc.message)
			})
		})
	}
}

// Repetitions measure from their own previous repetition, even when another
// call arrives in between, and the first repetition is not bounded.
func TestGapBoundsWithTimesMeasureBetweenRepetitions(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := NewHandler()
		startStep(t, h, "step 1 (retry)",
			retry(Exactly(3), "1s", "2s"), call("GET", "/other", nil, "other"))
		require.Equal(t, 200, sendAfter(h, 0, "/retry"))
		require.Equal(t, 200, sendAfter(h, 900*time.Millisecond, "/other"))
		require.Equal(t, 200, sendAfter(h, 200*time.Millisecond, "/retry"))
		require.Equal(t, 200, sendAfter(h, 500*time.Millisecond, "/retry"))
		requireFailure(t, h, "step 1 (retry): call 1: GET /retry repetition 3 "+
			"arrived 500ms after repetition 2, below min_gap 1s")
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
		{"below min_gap", "1s", "", "step 1 (retry): call 2: POST /second " +
			"arrived 500ms after call 1 (GET /first), below min_gap 1s"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				h := NewHandler()
				second := call("POST", "/second", map[string]any{"name": "a"}, "second")
				startStep(t, h, "step 1 (retry)", call("GET", "/first", nil, "first"),
					repeated(second, Times{}, tc.minGap, tc.maxGap))
				require.Equal(t, 200, sendAfter(h, 0, "/first"))
				time.Sleep(500 * time.Millisecond)
				w := httptest.NewRecorder()
				body := &slowBody{2 * time.Second, strings.NewReader(`{"name":"a"}`)}
				h.ServeHTTP(w, httptest.NewRequest("POST", "/second", body))
				require.Equal(t, 200, w.Code)
				if tc.message == "" {
					require.NoError(t, h.Verify())
					return
				}
				requireFailure(t, h, tc.message)
			})
		})
	}
}

// Concurrent requests can take the lock out of arrival order after reading
// their bodies. Gaps follow arrival order, so a later request that locks first
// neither hides a gap nor becomes an earlier call's reference.
func TestGapBoundsFollowArrivalOrderNotLockOrder(t *testing.T) {
	for _, tc := range []struct {
		name     string
		calls    []Call
		arrivals []string
		after    []time.Duration
		message  string
	}{
		{"repetitions", []Call{retry(Exactly(2), "", "1s")},
			[]string{"/retry", "/retry"}, []time.Duration{2 * time.Second, 0},
			"step 1 (create): call 1: GET /retry repetition 2 arrived 2s after " +
				"repetition 1, above max_gap 1s"},
		{"previous matched call", []Call{
			call("GET", "/a", nil, "a"), call("GET", "/b", nil, "b"),
			repeated(call("GET", "/c", nil, "c"), Times{}, "1s", ""),
		}, []string{"/b", "/a", "/c"},
			[]time.Duration{time.Second, 0, 1500 * time.Millisecond},
			"step 1 (create): call 3: GET /c arrived 500ms after call 2 (GET /b), " +
				"below min_gap 1s"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler()
			startStep(t, h, "step 1 (create)", tc.calls...)
			start := time.Now()
			for i, path := range tc.arrivals {
				req := httptest.NewRequest("GET", path, nil)
				e, _, _ := h.consume(httptest.NewRecorder(), req, start.Add(tc.after[i]))
				require.NotNil(t, e, path)
			}
			requireFailure(t, h, tc.message)
		})
	}
}

func TestGapBoundsNeedAnEarlierCallInTheStep(t *testing.T) {
	h := NewHandler()
	startStep(t, h, "step 1 (create)", call("GET", "/first", nil, "first"))
	require.Equal(t, 200, send(h, "GET", "/first", "").Code)
	startStep(t, h, "step 2 (update)",
		repeated(call("GET", "/second", nil, "second"), Times{}, "1s", ""))
	require.Equal(t, 200, send(h, "GET", "/second", "").Code)
	requireFailure(t, h, "step 2 (update): call 1: GET /second arrived first in its "+
		"step, so min_gap and max_gap have no earlier call to measure from")
}

func TestRepeatFieldsValidateOnLoad(t *testing.T) {
	const (
		count    = "must be a positive count or {min, max}"
		duration = "must be a positive duration, such as 1s"
		bounds   = "needs 0 <= min <= max and max >= 1"
	)
	for _, tc := range []struct{ fields, message string }{
		{"times: 0", `call 1: times "0" ` + count},
		{"times: often", `call 1: times "often" ` + count},
		{"times: [1, 2]", "call 1: times " + count},
		{"times: {min: 2}", "call 1: times range requires both min and max"},
		{"times: {min: 1, max: 2, step: 1}",
			`call 1: unknown times field "step"; expected min and max`},
		{"times: {min: 1, min: 2, max: 3}", "call 1: times.min is set twice"},
		{"times: {min: one, max: 2}", `call 1: times.min "one" must be an integer`},
		{"times: {min: 3, max: 2}", "call 1: times {min: 3, max: 2} " + bounds},
		{"times: {min: 0, max: 0}", "call 1: times {min: 0, max: 0} " + bounds},
		{"min_gap: soon", `call 1: min_gap "soon" ` + duration},
		{"max_gap: -1s", `call 1: max_gap "-1s" ` + duration},
		{"min_gap: 2s, max_gap: 1s", "call 1: min_gap 2s exceeds max_gap 1s"},
		{"times: 1, min_gap: 1s", "call 1: with times, min_gap and max_gap bound the " +
			"gaps between repetitions, so times needs a max of at least 2"},
	} {
		t.Run(tc.fields, func(t *testing.T) {
			var calls []Call
			fixture := "- {request: {method: GET, path: /retry}, " +
				"response: {status: 429}, " + tc.fields + "}"
			require.NoError(t, yaml.Unmarshal([]byte(fixture), &calls))
			err := NewHandler().StartStep("step 1 (create)", calls)
			require.Error(t, err)
			require.Equal(t, "step 1 (create): "+tc.message, err.Error())
		})
	}
}

// Exactly(0) must not become the zero Times, which expects the call once.
func TestExactlyRejectsCountsBelowOne(t *testing.T) {
	for _, n := range []int{0, -2} {
		calls := []Call{retry(Exactly(n), "", "")}
		err := NewHandler().StartStep("step 1 (create)", calls)
		require.Error(t, err)
		require.Equal(t,
			fmt.Sprintf("step 1 (create): call 1: times %d must be a positive count", n),
			err.Error())
	}
}

func TestReusableFixturesRejectRepeatFields(t *testing.T) {
	const message = "reusable call 1: times, min_gap and max_gap apply only to step calls"
	metadata := call("GET", "/metadata", nil, "m")
	h := NewHandler()
	require.EqualError(t,
		h.SetReusableFixtures([]Call{repeated(metadata, Exactly(2), "", "")}), message)
	require.EqualError(t,
		h.SetReusableFixtures([]Call{repeated(metadata, Times{}, "", "1s")}), message)
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
	anchored := "base: &retries {min: 2, max: 3}\ntimes: *retries\n"
	require.NoError(t, yaml.Unmarshal([]byte(anchored), &aliased))
	require.Equal(t, Times{Min: 2, Max: 3}, aliased.Times,
		"journeys reuse values through YAML anchors")
	var bounds struct {
		Times Times `yaml:"times"`
	}
	require.NoError(t, yaml.Unmarshal([]byte("times: {min: &n 2, max: *n}\n"), &bounds))
	require.Equal(t, Exactly(2), bounds.Times, "a bound can be an alias")
}
