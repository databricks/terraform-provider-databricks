package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
)

// errReader(true) will also fail on Close
type errReader bool

func (errReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("test error")
}

func (i errReader) Close() error {
	if i {
		return fmt.Errorf("test error")
	}
	return nil
}

type hc func(r *http.Request) (*http.Response, error)

func (cb hc) RoundTrip(r *http.Request) (*http.Response, error) {
	return cb(r)
}

func TestNew(t *testing.T) {
	c := NewApiClient(ClientConfig{})

	require.Equal(t, 96, c.config.DebugTruncateBytes)
	require.Equal(t, 5*time.Minute, c.config.RetryTimeout)
}

func TestSimpleRequestFailsURLError(t *testing.T) {
	c := NewApiClient(ClientConfig{
		RetryTimeout: 1 * time.Millisecond,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "GET", r.Method)
			require.Equal(t, "/a/b", r.URL.Path)
			require.Equal(t, "c=d", r.URL.RawQuery)
			require.Equal(t, "f", r.Header.Get("e"))
			return nil, fmt.Errorf("nope")
		}),
	})
	err := c.Do(context.Background(), "GET", "/a/b",
		WithRequestHeaders(map[string]string{
			"e": "f",
		}), WithRequestData(map[string]string{
			"c": "d",
		}))
	require.EqualError(t, err, `Get "/a/b?c=d": nope`)
}

func TestQueryParameters(t *testing.T) {
	c := NewApiClient(ClientConfig{
		RetryTimeout: 1 * time.Millisecond,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "POST", r.Method)
			require.Equal(t, "/a/b", r.URL.Path)
			require.Equal(t, "c=d&e=1", r.URL.RawQuery)
			return nil, fmt.Errorf("nope")
		}),
	})
	err := c.Do(context.Background(), "POST", "/a/b",
		WithQueryParameters(map[string]any{
			"c": "d",
			"e": 1,
		}))
	require.EqualError(t, err, `Post "/a/b?c=d&e=1": nope`)
}

func TestSimpleRequestFailsAPIError(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "GET", r.Method)
			require.Equal(t, "/a/b", r.URL.Path)
			require.Equal(t, "c=d", r.URL.RawQuery)
			require.Equal(t, "f", r.Header.Get("e"))
			return &http.Response{
				StatusCode: 400,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader(`nope`)),
			}, nil
		}),
	})
	err := c.Do(context.Background(), "GET", "/a/b",
		WithRequestHeaders(map[string]string{
			"e": "f",
		}), WithRequestData(map[string]string{
			"c": "d",
		}))
	require.EqualError(t, err, "http 400: nope")
}

func TestSimpleRequestSucceeds(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	var resp Dummy
	err := c.Do(context.Background(), "POST", "/c",
		WithRequestData(Dummy{1}),
		WithResponseUnmarshal(&resp))
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
}

func TestSimpleRequestRetried(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	var retried [1]bool
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			if !retried[0] {
				retried[0] = true
				return nil, &url.Error{
					Op:  "open",
					URL: "/a/b",
					Err: fmt.Errorf("connection refused"),
				}
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	var resp Dummy
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestData(Dummy{1}),
		WithResponseUnmarshal(&resp))
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
	require.True(t, retried[0], "request was not retried")
}

func TestSimpleRequestNotRetried(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	var tried, retried bool
	transportErr := &url.Error{
		Op:  "open",
		URL: "/a/b",
		Err: fmt.Errorf("some other reason"),
	}

	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			if !tried {
				tried = true
				return nil, transportErr
			}
			retried = true
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	var resp Dummy
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestData(Dummy{1}),
		WithResponseUnmarshal(&resp))
	require.Error(t, err)
	require.ErrorIs(t, err, transportErr)
	require.True(t, tried, "request was not tried")
	require.False(t, retried, "request was retried")
}

func TestHaltAttemptForLimit(t *testing.T) {
	ctx := context.Background()
	c := &ApiClient{
		config:      ClientConfig{},
		rateLimiter: &rate.Limiter{},
	}
	req, err := common.NewRequestBody([]byte{})
	require.NoError(t, err)
	_, rerr := c.attempt(ctx, "GET", "foo", req)()
	require.NotNil(t, rerr)
	require.Equal(t, true, rerr.Halt)
	require.EqualError(t, rerr.Err, "failed in rate limiter: rate: Wait(n=1) exceeds limiter's burst 0")
}

func TestHaltAttemptForNewRequest(t *testing.T) {
	ctx := context.Background()
	c := NewApiClient(ClientConfig{})
	req, err := common.NewRequestBody([]byte{})
	require.NoError(t, err)
	_, rerr := c.attempt(ctx, "🥱", "/", req)()
	require.NotNil(t, rerr)
	require.Equal(t, true, rerr.Halt)
	require.EqualError(t, rerr.Err, `failed creating new request: net/http: invalid method "🥱"`)
}

func TestHaltAttemptForVisitor(t *testing.T) {
	ctx := context.Background()
	c := NewApiClient(ClientConfig{})
	req, err := common.NewRequestBody([]byte{})
	require.NoError(t, err)
	_, rerr := c.attempt(ctx, "GET", "/", req,
		func(r *http.Request) error {
			return fmt.Errorf("🥱")
		})()
	require.NotNil(t, rerr)
	require.Equal(t, true, rerr.Halt)
	require.EqualError(t, rerr.Err, "failed during request visitor: 🥱")
}

func TestFailPerformChannel(t *testing.T) {
	ctx := context.Background()
	c := &ApiClient{
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	err := c.Do(ctx, "GET", "/", WithRequestData(true))
	require.EqualError(t, err, "request marshal: unsupported query string data: true")
}

func TestSimpleRequestErrReaderBodyStreamResponse(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(false),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/octet-stream"}
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{}))
	require.NoError(t, err, "streaming response bodies are not read")
}

func TestSimpleRequestErrReaderCloseBody(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(true),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{}))
	require.EqualError(t, err, "response body: test error")
}

func TestSimpleRequestErrReaderCloseBody_StreamResponse(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(true),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/octet-stream"}
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{}))
	require.NoError(t, err, "response body should not be closed for streaming responses")
}

func timeoutTransport(r *http.Request) (*http.Response, error) {
	select {
	case <-r.Context().Done():
		return nil, r.Context().Err()
	case <-time.After(50 * time.Millisecond):
		return nil, fmt.Errorf("test timeout")
	}
}

func TestSimpleRequestContextCancel(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Cancel outer context after 10ms
	go func() {
		defer cancel()
		time.Sleep(10 * time.Millisecond)
	}()

	c := NewApiClient(ClientConfig{
		Transport: hc(timeoutTransport),
	})
	err := c.Do(ctx, "GET", "/a", WithRequestData(map[string]any{}))
	require.ErrorContains(t, err, "context canceled")
}

func TestSimpleRequestContextDeadline(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Millisecond))
	defer cancel()

	c := NewApiClient(ClientConfig{
		Transport: hc(timeoutTransport),
	})
	err := c.Do(ctx, "GET", "/a", WithRequestData(map[string]any{}))
	require.ErrorContains(t, err, "context deadline exceeded")
}

func TestSimpleRequestTimeout(t *testing.T) {
	ctx := context.Background()

	c := NewApiClient(ClientConfig{
		HTTPTimeout: 10 * time.Millisecond,
		Transport:   hc(timeoutTransport),
	})
	err := c.Do(ctx, "GET", "/a", WithRequestData(map[string]any{}))
	require.ErrorContains(t, err, "request timed out after 10ms of inactivity")
}

type BufferLogger struct {
	strings.Builder
}

func (l *BufferLogger) Enabled(_ context.Context, level logger.Level) bool {
	return true
}

func (l *BufferLogger) Tracef(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[TRACE] "+format+"\n", v...))
}

func (l *BufferLogger) Debugf(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[DEBUG] "+format+"\n", v...))
}

func (l *BufferLogger) Infof(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[INFO] "+format+"\n", v...))
}

func (l *BufferLogger) Warnf(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[WARN] "+format+"\n", v...))
}

func (l *BufferLogger) Errorf(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[ERROR] "+format+"\n", v...))
}

func configureBufferedLogger(t *testing.T) *BufferLogger {
	prevLogger := logger.DefaultLogger
	bufLogger := &BufferLogger{}
	logger.DefaultLogger = bufLogger
	t.Cleanup(func() {
		logger.DefaultLogger = prevLogger
	})
	return bufLogger
}

func TestSimpleResponseRedaction(t *testing.T) {
	bufLogger := configureBufferedLogger(t)

	c := NewApiClient(ClientConfig{
		DebugTruncateBytes: 16,
		DebugHeaders:       true,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Proto:      "HTTP/3.4",
				Status:     "200 Fine",
				Body: io.NopCloser(strings.NewReader(`{
					"string_value": "__SENSITIVE01__",
					"inner": {
					  "token_value": "__SENSITIVE02__",
					  "content": "__SENSITIVE03__"
					},
					"list": [
					  {
					    "token_value": "__SENSITIVE04__"
					  }
					],
					"longer": "12345678901234567890qwerty"
				}`)),
				Request: r,
			}, nil
		}),
	})
	err := c.Do(context.Background(), "GET", "/a",
		WithRequestData(map[string]any{
			"b": 0,
			"a": 3,
			"c": 23,
		}))
	require.NoError(t, err)
	// not testing for exact logged lines, as header order is not deterministic
	require.NotContains(t, bufLogger.String(), "__SENSITIVE01__")
	require.NotContains(t, bufLogger.String(), "__SENSITIVE02__")
	require.NotContains(t, bufLogger.String(), "__SENSITIVE03__")
	require.NotContains(t, bufLogger.String(), "__SENSITIVE04__")
	require.NotContains(t, bufLogger.String(), "12345678901234567890qwerty")
}

func TestInlineArrayDebugging(t *testing.T) {
	bufLogger := configureBufferedLogger(t)

	c := NewApiClient(ClientConfig{
		DebugTruncateBytes: 2048,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(strings.NewReader(`[
					{"foo": "bar"}
				]`)),
				Request: r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "GET", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{
			"b": 0,
			"a": 3,
			"c": 23,
		}))
	require.NoError(t, err)

	require.Equal(t, `[DEBUG] GET /a?a=3&b=0&c=23
<  
< [
<   {
<     "foo": "bar"
<   }
< ]
`, bufLogger.String())
}

func TestInlineArrayDebugging_StreamResponse(t *testing.T) {
	bufLogger := configureBufferedLogger(t)

	c := NewApiClient(ClientConfig{
		DebugTruncateBytes: 2048,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`lots of bytes`)),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/octet-stream"}
	err := c.Do(context.Background(), "GET", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{
			"b": 0,
			"a": 3,
			"c": 23,
		}))
	require.NoError(t, err)

	require.Equal(t, `[DEBUG] GET /a?a=3&b=0&c=23
<  
< <Streaming response>
`, bufLogger.String())
}

func TestLogQueryParametersWithPercent(t *testing.T) {
	bufLogger := configureBufferedLogger(t)

	c := NewApiClient(ClientConfig{
		DebugTruncateBytes: 2048,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": "bar"}`)),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "GET", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{
			"filter": "name like '%'",
		}))
	assert.NoError(t, err)
	assert.Equal(t, `[DEBUG] GET /a?filter=name like '%'
<  
< {
<   "foo": "bar"
< }
`, bufLogger.String())
}

func TestLogCancelledRequest(t *testing.T) {
	bufLogger := configureBufferedLogger(t)

	ctx, cancel := context.WithCancel(context.Background())
	c := NewApiClient(ClientConfig{
		DebugTruncateBytes: 2048,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			cancel()
			return nil, ctx.Err()
		}),
	})
	err := c.Do(context.Background(), "GET", "/a")
	assert.Error(t, err)
	assert.Equal(t, `[DEBUG] GET /a
< Error: Get "/a": request timed out after 1m0s of inactivity
[DEBUG] non-retriable error: Get "/a": request timed out after 1m0s of inactivity
`, bufLogger.String())
}

func TestStreamRequestFromFileWithReset(t *testing.T) {
	// make a temporary file with some content
	f, err := os.CreateTemp("", "databricks-client-test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	_, err = f.WriteString("hello world")
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// Make a reader that reads this file
	r, err := os.Open(f.Name())
	require.NoError(t, err)
	defer r.Close()

	succeed := false
	handler := func(req *http.Request) (*http.Response, error) {
		bytes, err := io.ReadAll(req.Body)
		require.NoError(t, err)
		require.Equal(t, "hello world", string(bytes))
		if succeed {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader("succeeded")),
				Request:    req,
			}, nil
		}
		succeed = true
		return &http.Response{
			StatusCode: 429,
			Body:       io.NopCloser(strings.NewReader("failed")),
			Request:    req,
		}, nil
	}

	client := NewApiClient(ClientConfig{
		Transport: hc(handler),
	})

	respBytes := bytes.Buffer{}
	err = client.Do(context.Background(), "POST", "/a",
		WithRequestData(r),
		WithResponseUnmarshal(&respBytes))
	require.NoError(t, err)
	require.Equal(t, "succeeded", respBytes.String())
	require.True(t, succeed)
}

type customReader struct{}

func (c customReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func TestCannotRetryArbitraryReader(t *testing.T) {
	client := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 429,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		}),
	})
	err := client.Do(context.Background(), "POST", "/a",
		WithRequestData(customReader{}))
	require.ErrorContains(t, err, "cannot reset reader of type httpclient.customReader")
}

func TestRetryGetRequest(t *testing.T) {
	// This test was added in response to https://github.com/databricks/terraform-provider-databricks/issues/2675.
	succeed := false
	handler := func(req *http.Request) (*http.Response, error) {
		require.Nil(t, req.Body)

		if succeed {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader("succeeded")),
				Request:    req,
			}, nil
		}

		succeed = true
		return &http.Response{
			StatusCode: 429,
			Body:       io.NopCloser(strings.NewReader("failed")),
			Request:    req,
		}, nil
	}

	client := NewApiClient(ClientConfig{
		Transport: hc(handler),
	})

	respBytes := bytes.Buffer{}
	err := client.Do(context.Background(), "GET", "/a",
		WithResponseUnmarshal(&respBytes))
	require.NoError(t, err)
	require.Equal(t, "succeeded", respBytes.String())
	require.True(t, succeed)
}

func TestOAuth2Integration(t *testing.T) {
	inner := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 204,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		}),
	})

	ctx := context.Background()
	ctx = inner.InContextForOAuth2(ctx)

	outer := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: "abc",
	}))
	res, err := outer.Get("abc")
	require.NoError(t, err)
	require.Equal(t, 204, res.StatusCode)
}

func TestErrorOnMultipleAuthVisitor(t *testing.T) {
	defaultAuthVisitor := func(r *http.Request) error {
		r.Header.Set("X-Auth", "abc")
		return nil
	}
	customAuthVisitor := func(r *http.Request) error {
		r.Header.Set("X-Auth-custom", "def")
		return nil
	}
	authOption := DoOption{
		in:           customAuthVisitor,
		isAuthOption: true,
	}
	c := NewApiClient(ClientConfig{
		AuthVisitor: defaultAuthVisitor,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "", r.Header.Get("X-Auth"))
			require.Equal(t, "def", r.Header.Get("X-Auth-custom"))
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	err := c.Do(context.Background(), "GET", "/a/b", WithRequestData(map[string]any{}), authOption, authOption)
	require.Error(t, err, "only one auth visitor is allowed")
}

func TestCustomAuthVisitor(t *testing.T) {
	defaultAuthVisitor := func(r *http.Request) error {
		r.Header.Set("X-Auth", "abc")
		return nil
	}
	customAuthVisitor := func(r *http.Request) error {
		r.Header.Set("X-Auth-custom", "def")
		return nil
	}
	authOption := DoOption{
		in:           customAuthVisitor,
		isAuthOption: true,
	}
	c := NewApiClient(ClientConfig{
		AuthVisitor: defaultAuthVisitor,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "", r.Header.Get("X-Auth"))
			require.Equal(t, "def", r.Header.Get("X-Auth-custom"))
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	err := c.Do(context.Background(), "GET", "/a/b", WithRequestData(map[string]any{}), authOption)
	require.NoError(t, err)
}

func TestDefaultAuthVisitor(t *testing.T) {
	defaultAuthVisitor := func(r *http.Request) error {
		r.Header.Set("X-Auth", "abc")
		return nil
	}
	anotherVisitor := func(r *http.Request) error {
		r.Header.Set("X-Unrelated", "def")
		return nil
	}
	authOption := DoOption{
		in: anotherVisitor,
	}
	c := NewApiClient(ClientConfig{
		AuthVisitor: defaultAuthVisitor,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "abc", r.Header.Get("X-Auth"))
			require.Equal(t, "def", r.Header.Get("X-Unrelated"))
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	err := c.Do(context.Background(), "GET", "/a/b", WithRequestData(map[string]any{}), authOption)
	require.NoError(t, err)
}

func TestTraceparentHeader(t *testing.T) {
	seenTraceparents := []string{}
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			tp := r.Header.Get("Traceparent")
			assert.NotEmpty(t, tp)
			assert.NotContains(t, seenTraceparents, tp)
			seenTraceparents = append(seenTraceparents, tp)
			return &http.Response{
				StatusCode: 200,
				Request:    r,
			}, nil
		}),
	})

	for i := 0; i < 10; i++ {
		err := c.Do(context.Background(), "GET", "/a/b")
		assert.NoError(t, err)
	}
}

func TestTraceparentHeaderDoesNotOverrideUserHeader(t *testing.T) {
	userTraceparent := "00-thetraceid-theparentid-00"
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			tp := r.Header.Get("Traceparent")
			assert.NotEmpty(t, tp)
			assert.Equal(t, userTraceparent, tp)
			return &http.Response{
				StatusCode: 200,
				Request:    r,
			}, nil
		}),
	})

	err := c.Do(context.Background(), "GET", "/a/b", WithRequestHeaders(map[string]string{
		"Traceparent": userTraceparent,
	}))
	assert.NoError(t, err)
}
