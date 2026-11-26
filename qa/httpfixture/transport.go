package httpfixture

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// NewClientTransport returns the transport for clients of a fixture server, so
// faults reach them. Go's client silently resends a replayable request that fails
// on a reused connection, so it opens one connection per request.
func NewClientTransport() http.RoundTripper {
	return bufferedBodyTransport{base: &http.Transport{DisableKeepAlives: true}}
}

// bufferedBodyTransport sends a copy of the request body. net/http can return a
// response that closes the connection before it has finished reading the body,
// and the Databricks SDK rewinds that same reader for a retry, which could then
// go out empty.
type bufferedBodyTransport struct{ base http.RoundTripper }

func (t bufferedBodyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body == nil || req.Body == http.NoBody {
		return t.base.RoundTrip(req)
	}
	body, err := readBody(req)
	if err != nil {
		return nil, err
	}
	// The length stays as declared, so an unknown one is still sent chunked, and
	// a wrong one fails as net/http fails it.
	if req.ContentLength > 0 && req.ContentLength != int64(len(body)) {
		return nil, fmt.Errorf("http: ContentLength=%d with Body length %d", req.ContentLength, len(body))
	}
	clone := req.Clone(req.Context())
	clone.Body, clone.GetBody = http.NoBody, func() (io.ReadCloser, error) { return http.NoBody, nil }
	if len(body) > 0 {
		clone.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(body)), nil }
		clone.Body, _ = clone.GetBody()
	}
	return t.base.RoundTrip(clone)
}

// readBody reads and closes the body. When the request is canceled first, it
// closes the body to unblock a stalled read and returns the cause unwrapped, as
// net/http does; the SDK compares it with context.Canceled to report a timeout.
func readBody(req *http.Request) ([]byte, error) {
	var once sync.Once
	closeBody := func() (err error) {
		once.Do(func() { err = req.Body.Close() })
		return err
	}
	type result struct {
		body []byte
		err  error
	}
	read := make(chan result, 1)
	go func() {
		body, err := io.ReadAll(req.Body)
		read <- result{body, errors.Join(err, closeBody())}
	}()
	select {
	case r := <-read:
		return r.body, r.err
	case <-req.Context().Done():
		_ = closeBody()
		return nil, context.Cause(req.Context())
	}
}
