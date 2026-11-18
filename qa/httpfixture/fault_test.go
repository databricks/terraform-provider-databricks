package httpfixture

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// faultServer serves calls through a real connection, which faults require.
func faultServer(t *testing.T, calls ...Call) (*Handler, string) {
	t.Helper()
	h := NewHandler()
	require.NoError(t, h.StartStep("step 1 (faults)", calls))
	server := httptest.NewServer(h)
	t.Cleanup(server.Close)
	return h, server.URL
}

func textCall(path, body string) Call {
	return Call{
		Request:  Request{Method: "GET", Path: path},
		Response: Response{Status: 200, Body: &Body{Text: &body}},
	}
}

func faulted(c Call, fault Fault) Call {
	c.Response.Fault = fault
	if fault == FaultReset || fault == FaultClose {
		c.Response.Status, c.Response.Body = 0, nil
	}
	return c
}

func TestDelayWaitsBeforeResponding(t *testing.T) {
	c := textCall("/slow", "done")
	c.Response.Delay = "100ms"
	h, url := faultServer(t, c)
	start := time.Now()
	resp, err := http.Get(url + "/slow")
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "done", string(body))
	require.GreaterOrEqual(t, time.Since(start), 100*time.Millisecond)
	require.NoError(t, h.Verify())
}

func TestDelayEndsWhenClientGivesUp(t *testing.T) {
	c := textCall("/slow", "never read")
	c.Response.Delay = "1h"
	h := NewHandler()
	require.NoError(t, h.StartStep("step 1 (faults)", []Call{c}))
	server := httptest.NewServer(h)
	client := &http.Client{Timeout: 100 * time.Millisecond}
	_, err := client.Get(server.URL + "/slow")
	require.ErrorContains(t, err, "Client.Timeout exceeded while awaiting headers")
	// Close waits for active handlers, so it returns only once the delay has ended.
	closed := make(chan struct{})
	go func() {
		server.Close()
		close(closed)
	}()
	select {
	case <-closed:
	case <-time.After(10 * time.Second):
		t.Fatal("the delayed handler kept running after the client gave up")
	}
	require.NoError(t, h.Verify(), "the timed-out call is consumed")
}

func TestDelayDoesNotBlockConcurrentRequests(t *testing.T) {
	slow := textCall("/slow", "slow")
	slow.Response.Delay = "1h"
	h, url := faultServer(t, slow, textCall("/fast", "fast"))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	slowDone := make(chan error, 1)
	go func() {
		req, _ := http.NewRequestWithContext(ctx, "GET", url+"/slow", nil)
		_, err := http.DefaultClient.Do(req)
		slowDone <- err
	}()
	require.Eventually(t, func() bool {
		h.mu.Lock()
		defer h.mu.Unlock()
		return h.calls[0].used
	}, 5*time.Second, time.Millisecond, "the slow call holds its delay")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url + "/fast")
	require.NoError(t, err)
	resp.Body.Close()
	require.Equal(t, 200, resp.StatusCode)
	cancel()
	require.ErrorIs(t, <-slowDone, context.Canceled)
	require.NoError(t, h.Verify())
}

func TestFaultReset(t *testing.T) {
	h, url := faultServer(t, faulted(textCall("/reset", "unused"), FaultReset))
	_, err := http.Get(url + "/reset")
	require.ErrorContains(t, err, "connection reset by peer")
	require.NoError(t, h.Verify())
}

func TestFaultClose(t *testing.T) {
	h, url := faultServer(t, faulted(textCall("/close", "unused"), FaultClose))
	_, err := http.Get(url + "/close")
	require.ErrorIs(t, err, io.EOF)
	require.True(t, strings.HasSuffix(err.Error(), `/close": EOF`), err.Error())
	require.NoError(t, h.Verify())
}

func TestFaultTruncateBody(t *testing.T) {
	h, url := faultServer(t, faulted(textCall("/truncate", "0123456789"), FaultTruncateBody))
	resp, err := http.Get(url + "/truncate")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, int64(10), resp.ContentLength, "headers declare the whole body")
	body, err := io.ReadAll(resp.Body)
	require.ErrorIs(t, err, io.ErrUnexpectedEOF)
	require.Equal(t, "01234", string(body))
	require.NoError(t, h.Verify())
}

func TestFaultStallBody(t *testing.T) {
	h, url := faultServer(t, faulted(textCall("/stall", "0123456789"), FaultStallBody))
	client := &http.Client{Timeout: 200 * time.Millisecond}
	resp, err := client.Get(url + "/stall")
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.ErrorContains(t, err, "Client.Timeout or context cancellation while reading body")
	require.Equal(t, "01234", string(body))
	require.NoError(t, h.Verify())
}

func TestStatusLineWithoutReasonPhrase(t *testing.T) {
	c := call("GET", "/missing", nil, "")
	c.Response = Response{Status: 404, StatusLine: "HTTP/1.1 404", Body: &Body{JSON: map[string]any{
		"error_code": "NOT_FOUND",
	}}}
	h, url := faultServer(t, c)
	resp, err := http.Get(url + "/missing")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, 404, resp.StatusCode)
	require.Equal(t, "404", resp.Status, "no reason phrase on the wire")
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.JSONEq(t, `{"error_code":"NOT_FOUND"}`, string(body))
	require.NoError(t, h.Verify())
}

// Go's client resends an idempotent request that fails on a reused connection.
// The handler closes every connection, so a fault after another call still fails.
func TestFaultAfterAnotherCallIsNotRetriedByTheClient(t *testing.T) {
	h, url := faultServer(t, textCall("/first", "ok"), faulted(textCall("/close", "unused"), FaultClose))
	resp, err := http.Get(url + "/first")
	require.NoError(t, err)
	_, _ = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	_, err = http.Get(url + "/close")
	require.ErrorIs(t, err, io.EOF)
	require.NoError(t, h.Verify())
}

func TestFaultedCallsAreConsumedInOrder(t *testing.T) {
	reset := faulted(textCall("/retry", "unused"), FaultReset)
	h, url := faultServer(t, reset, textCall("/retry", "second attempt"))
	h.SetEnforceCallOrder(true)
	_, err := http.Get(url + "/retry")
	require.ErrorContains(t, err, "connection reset by peer")
	resp, err := http.Get(url + "/retry")
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "second attempt", string(body))
	require.NoError(t, h.Verify())

	// Without the retry, the second fixture stays unused and fails the step.
	h, url = faultServer(t, reset, textCall("/retry", "second attempt"))
	_, err = http.Get(url + "/retry")
	require.ErrorContains(t, err, "connection reset by peer")
	require.ErrorContains(t, h.Verify(), "unconsumed call 2: GET /retry")
}

func TestFaultRequiresHijacking(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("step 1 (faults)", []Call{faulted(textCall("/close", "unused"), FaultClose)}))
	send(h, "GET", "/close", "")
	require.ErrorContains(t, h.Verify(), "step 1 (faults): GET /close: fault and status_line require a connection that supports hijacking")
}

func TestInvalidResponseWireFields(t *testing.T) {
	text := func(value string) *Body { return &Body{Text: &value} }
	for _, tc := range []struct {
		name     string
		response Response
		message  string
	}{
		{"malformed delay", Response{Status: 200, Delay: "soon"}, ` response: delay "soon" must be a positive duration`},
		{"number delay", Response{Status: 200, Delay: "5"}, ` response: delay "5" must be a positive duration`},
		{"zero delay", Response{Status: 200, Delay: "0s"}, ` response: delay "0s" must be a positive duration`},
		{"negative delay", Response{Status: 200, Delay: "-1s"}, ` response: delay "-1s" must be a positive duration`},
		{"unknown fault", Response{Fault: "timeout"}, ` response: unknown fault "timeout"; expected reset, close, truncate_body or stall_body`},
		{"reset with status", Response{Status: 200, Fault: FaultReset}, " response: fault reset writes no response; omit status"},
		{"close with body", Response{Body: text("x"), Fault: FaultClose}, " response: fault close writes no response"},
		{"close with headers", Response{Headers: map[string]string{"X": "y"}, Fault: FaultClose}, " response: fault close writes no response"},
		{"reset with status line", Response{StatusLine: "HTTP/1.1 404", Fault: FaultReset}, " response: fault reset writes no response"},
		{"truncate without status", Response{Body: text("x"), Fault: FaultTruncateBody}, ": invalid response status 0"},
		{"truncate without body", Response{Status: 200, Fault: FaultTruncateBody}, " response: fault truncate_body requires a nonempty response body"},
		{"stall with empty body", Response{Status: 200, Body: text(""), Fault: FaultStallBody}, " response: fault stall_body requires a nonempty response body"},
		{"status line code", Response{Status: 404, StatusLine: "HTTP/1.1 400"}, ` response: status_line "HTTP/1.1 400" must be HTTP/1.1 404`},
		{"status line protocol", Response{Status: 404, StatusLine: "HTTP/2 404"}, ` response: status_line "HTTP/2 404" must be HTTP/1.1 404`},
		{"status line header injection", Response{Status: 404, StatusLine: "HTTP/1.1 404\r\nX-Extra: 1"}, ` response: status_line "HTTP/1.1 404\r\nX-Extra: 1" must be HTTP/1.1 404`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := Call{Request: Request{Method: "GET", Path: "/wire"}, Response: tc.response}
			err := NewHandler().StartStep("step 2 (update)", []Call{textCall("/ok", "ok"), c})
			require.ErrorContains(t, err, "step 2 (update): call 2"+tc.message)
		})
	}
}

func TestValidResponseWireFields(t *testing.T) {
	text := "x"
	for _, response := range []Response{
		{Fault: FaultReset, Delay: "10ms"},
		{Fault: FaultClose},
		{Status: 200, Body: &Body{Text: &text}, Fault: FaultTruncateBody, StatusLine: "HTTP/1.1 200"},
		{Status: 404, StatusLine: "HTTP/1.0 404 "},
		{Status: 404, StatusLine: "HTTP/1.1 404 Volume Not Found"},
	} {
		c := Call{Request: Request{Method: "GET", Path: "/wire"}, Response: response}
		require.NoError(t, NewHandler().StartStep("valid", []Call{c}))
	}
}
