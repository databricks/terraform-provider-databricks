package httpfixture

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

// Rewinding the caller's reader after the round trip, as the SDK does for a
// retry, must leave the body that net/http may still be reading intact.
func TestClientTransportSendsACopyOfTheBody(t *testing.T) {
	var sent *http.Request
	transport := bufferedBodyTransport{base: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		sent = req
		return &http.Response{StatusCode: 204, Body: http.NoBody, Request: req}, nil
	})}
	original := bytes.NewReader([]byte(`{"name":"volume"}`))
	req, err := http.NewRequest("PATCH", "http://fixture.test/volumes", original)
	require.NoError(t, err)
	_, err = transport.RoundTrip(req)
	require.NoError(t, err)
	_, err = original.Seek(0, io.SeekStart)
	require.NoError(t, err)
	late, err := io.ReadAll(sent.Body)
	require.NoError(t, err)
	require.Equal(t, `{"name":"volume"}`, string(late))
	require.Equal(t, int64(len(late)), sent.ContentLength)
	require.Equal(t, original.Size(), int64(original.Len()), "the late read leaves the rewound reader alone")
	again, err := sent.GetBody()
	require.NoError(t, err)
	retried, err := io.ReadAll(again)
	require.NoError(t, err)
	require.Equal(t, string(late), string(retried))

	req, err = http.NewRequest("POST", "http://fixture.test/empty", io.NopCloser(strings.NewReader("")))
	require.NoError(t, err)
	_, err = transport.RoundTrip(req)
	require.NoError(t, err)
	require.Equal(t, http.NoBody, sent.Body)
	require.Zero(t, sent.ContentLength)
}

// The copy keeps the request's framing: an unknown length is still sent chunked,
// and a declared length that differs from the body fails as net/http fails it.
func TestClientTransportKeepsTheDeclaredLength(t *testing.T) {
	var framing []string
	var length int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		framing, length = r.TransferEncoding, r.ContentLength
		w.WriteHeader(204)
	}))
	defer server.Close()
	client := &http.Client{Transport: NewClientTransport()}
	resp, err := client.Post(server.URL, "text/plain", io.NopCloser(strings.NewReader("upload")))
	require.NoError(t, err)
	resp.Body.Close()
	require.Equal(t, []string{"chunked"}, framing, "an unknown length stays unknown")
	resp, err = client.Post(server.URL, "text/plain", strings.NewReader("upload"))
	require.NoError(t, err)
	resp.Body.Close()
	require.Nil(t, framing)
	require.Equal(t, int64(6), length)

	req, err := http.NewRequest("POST", server.URL, io.NopCloser(strings.NewReader("upload")))
	require.NoError(t, err)
	req.ContentLength = 3
	_, err = client.Do(req)
	require.ErrorContains(t, err, "http: ContentLength=3 with Body length 6")
}

// roundTripWithin fails the test instead of hanging when RoundTrip blocks.
func roundTripWithin(t *testing.T, transport http.RoundTripper, req *http.Request) error {
	t.Helper()
	done := make(chan error, 1)
	go func() {
		_, err := transport.RoundTrip(req)
		done <- err
	}()
	select {
	case err := <-done:
		return err
	case <-time.After(10 * time.Second):
		t.Fatal("RoundTrip kept reading a stalled body after the request was canceled")
		return nil
	}
}

// A stalled streaming body must not hang the client past its cancellation.
func TestClientTransportStopsReadingWhenCanceled(t *testing.T) {
	transport := bufferedBodyTransport{base: roundTripFunc(func(*http.Request) (*http.Response, error) {
		t.Error("a canceled request must not reach the base transport")
		return nil, nil
	})}
	stalled, writer := io.Pipe()
	ctx, cancel := context.WithCancel(context.Background())
	req, err := http.NewRequestWithContext(ctx, "POST", "http://fixture.test/upload", stalled)
	require.NoError(t, err)
	time.AfterFunc(50*time.Millisecond, cancel)
	err = roundTripWithin(t, transport, req)
	require.Equal(t, context.Canceled, err, "the SDK compares the error with context.Canceled")
	_, err = writer.Write([]byte("late"))
	require.ErrorIs(t, err, io.ErrClosedPipe, "the source body is closed")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		t.Error("a timed-out request must not reach the server")
	}))
	defer server.Close()
	stalled, _ = io.Pipe()
	client := &http.Client{Transport: NewClientTransport(), Timeout: 100 * time.Millisecond}
	timedOut := make(chan error, 1)
	go func() {
		_, err := client.Post(server.URL, "text/plain", stalled)
		timedOut <- err
	}()
	select {
	case err := <-timedOut:
		// net/http adds its Client.Timeout annotation only once its timer has fired.
		require.ErrorIs(t, err, context.DeadlineExceeded)
	case <-time.After(10 * time.Second):
		t.Fatal("the client timeout did not stop reading a stalled body")
	}
}
