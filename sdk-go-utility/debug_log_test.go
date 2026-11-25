package sdkgoutility

import (
	"bytes"
	"cmp"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"testing"
	"testing/iotest"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/auth/credentials"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const debugLogURL = "http://u.test/api/2.0/x"

// debugLogRecorder is a logger of the classic SDK that keeps the debug
// messages that it is given, also when it says that debug logging is off.
type debugLogRecorder struct {
	off      bool
	mu       sync.Mutex
	asked    []logger.Level
	messages []string
}

func (r *debugLogRecorder) Enabled(_ context.Context, level logger.Level) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.asked = append(r.asked, level)
	return !r.off && level >= logger.LevelDebug
}

func (r *debugLogRecorder) Debugf(_ context.Context, format string, v ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.messages = append(r.messages, fmt.Sprintf(format, v...))
}

func (r *debugLogRecorder) Tracef(context.Context, string, ...any) {}
func (r *debugLogRecorder) Infof(context.Context, string, ...any)  {}
func (r *debugLogRecorder) Warnf(context.Context, string, ...any)  {}
func (r *debugLogRecorder) Errorf(context.Context, string, ...any) {}

func (r *debugLogRecorder) logged() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return slices.Clone(r.messages)
}

// request returns a request whose context logs to the recorder.
func (r *debugLogRecorder) request(t *testing.T, method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequestWithContext(logger.NewContext(context.Background(), r), method, url, body)
	require.NoError(t, err)
	return req
}

type debugLogRoundTrip func(*http.Request) (*http.Response, error)

func (f debugLogRoundTrip) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

type debugLogLoader struct{ err error }

func (l debugLogLoader) Name() string                   { return "debug-log-test" }
func (l debugLogLoader) Configure(*config.Config) error { return l.err }

// debugLogConfig returns a config that reads neither the environment nor a
// file, and that asks its host for nothing.
func debugLogConfig(host string, headers bool, truncateBytes int) *config.Config {
	return &config.Config{
		Host:                 host,
		Token:                "dapi-debug-log",
		AuthType:             "pat",
		DebugHeaders:         headers,
		DebugTruncateBytes:   truncateBytes,
		RetryTimeoutSeconds:  30,
		Loaders:              []config.Loader{debugLogLoader{}},
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) { return nil, nil },
	}
}

// debugLogOver returns DebugLog over transport, for a config without a server.
func debugLogOver(t *testing.T, headers bool, truncateBytes int, transport http.RoundTripper) http.RoundTripper {
	logged, err := DebugLog(debugLogConfig("http://u.test", headers, truncateBytes), transport)
	require.NoError(t, err)
	return logged
}

// debugLogResponse returns a response as a transport other than
// http.Transport may build it: without its request, or without a body.
func debugLogResponse(status int, contentType string, body io.ReadCloser) *http.Response {
	resp := &http.Response{
		Proto:      "HTTP/1.1",
		StatusCode: status,
		Status:     fmt.Sprintf("%d %s", status, http.StatusText(status)),
		Header:     http.Header{},
		Body:       body,
	}
	if contentType != "" {
		resp.Header.Set("Content-Type", contentType)
	}
	return resp
}

func debugLogText(body string) io.ReadCloser { return io.NopCloser(strings.NewReader(body)) }

func TestDebugLog_Errors(t *testing.T) {
	unresolved := debugLogConfig("http://u.test", false, 0)
	unresolved.Loaders = []config.Loader{debugLogLoader{err: errors.New("no config")}}
	for name, c := range map[string]struct {
		cfg       *config.Config
		transport http.RoundTripper
		want      string
	}{
		"nil cfg":                    {nil, http.DefaultTransport, "DebugLog: cfg is nil"},
		"nil transport":              {debugLogConfig("http://u.test", false, 0), nil, "DebugLog: transport is nil"},
		"cfg that can't be resolved": {unresolved, http.DefaultTransport, "DebugLog: resolve: no config"},
	} {
		t.Run(name, func(t *testing.T) {
			transport, err := DebugLog(c.cfg, c.transport)
			assert.Nil(t, transport)
			assert.ErrorContains(t, err, c.want)
		})
	}
}

// debugLogExchange is one request that a caller sends through DebugLog over
// a plain transport, reading the response to its end and closing it.
type debugLogExchange struct {
	name        string
	method, url string
	header      map[string]string
	body        io.Reader
	prepare     func(*http.Request)
	response    *http.Response
	err         error
	headers     bool
	truncate    int
	// off turns debug logging off for the context of the request.
	off bool
	// direct is set when the caller gets the body of the transport itself.
	direct bool
	// want is the message, and empty when there is none.
	want string
}

func (e debugLogExchange) run(t *testing.T) {
	recorder := &debugLogRecorder{off: e.off}
	req := recorder.request(t, cmp.Or(e.method, "GET"), cmp.Or(e.url, debugLogURL), e.body)
	for name, value := range e.header {
		req.Header.Set(name, value)
	}
	if e.prepare != nil {
		e.prepare(req)
	}
	sentBody, body := req.Body, cmp.Or(e.response, &http.Response{}).Body
	base := debugLogRoundTrip(func(sent *http.Request) (*http.Response, error) {
		assert.True(t, req == sent && sentBody == sent.Body, "the transport gets the request of the caller")
		return e.response, e.err
	})
	resp, err := debugLogOver(t, e.headers, e.truncate, base).RoundTrip(req)
	assert.Same(t, e.response, resp)
	assert.Equal(t, e.err, err)
	if resp != nil && resp.Body != nil {
		assert.Equal(t, e.direct, resp.Body == body)
		_, _ = io.Copy(io.Discard, resp.Body)
		assert.NoError(t, resp.Body.Close())
	}
	assert.Equal(t, e.want, strings.Join(recorder.logged(), "\n---\n"))
	assert.Subset(t, []logger.Level{logger.LevelDebug}, recorder.asked, "asked as the classic client asks")
}

func TestDebugLog_Message(t *testing.T) {
	json := map[string]string{"Content-Type": "application/json; charset=utf-8"}
	stream := map[string]string{"Accept": "application/octet-stream"}
	boom := errors.New("boom")
	failing := io.NopCloser(iotest.ErrReader(boom))
	unread := func(req *http.Request) {
		req.GetBody = func() (io.ReadCloser, error) { t.Error("the request body is read"); return nil, boom }
	}
	for _, e := range []debugLogExchange{{
		name: "query and JSON response", url: debugLogURL + "?name=a%20b&n=1",
		response: debugLogResponse(200, "application/json", debugLogText(`{"b":[1,"two"],"a":{"token":"s3cret"}}`)),
		want:     "GET /api/2.0/x?name=a b&n=1\n< HTTP/1.1 200 OK\n< {\n<   \"a\": {\n<     \"token\": \"**REDACTED**\"\n<   },\n<   \"b\": [\n<     1,\n<     \"two\"\n<   ]\n< }",
	}, {
		name: "JSON request body", method: "POST", header: json, direct: true,
		body:     strings.NewReader(`{"name":"m","password":"s3cret"}`),
		response: debugLogResponse(200, "", http.NoBody),
		want:     "POST /api/2.0/x\n> {\n>   \"name\": \"m\",\n>   \"password\": \"**REDACTED**\"\n> }\n< HTTP/1.1 200 OK",
	}, {
		name: "request body that is a stream", method: "PUT", header: json, direct: true,
		body:     io.NopCloser(strings.NewReader(`{"password":"s3cret"}`)),
		response: debugLogResponse(200, "", http.NoBody),
		want:     "PUT /api/2.0/x\n> <io.Reader>\n< HTTP/1.1 200 OK",
	}, {
		name: "request body that is not JSON", method: "POST", header: map[string]string{"Content-Type": "text/plain"}, direct: true,
		body:     strings.NewReader("password=s3cret"),
		response: debugLogResponse(200, "", http.NoBody),
		want:     "POST /api/2.0/x\n> <io.Reader>\n< HTTP/1.1 200 OK",
	}, {
		name: "request body that can't be read again", method: "POST", header: json, direct: true,
		body:     strings.NewReader(`{"password":"s3cret"}`),
		prepare:  func(req *http.Request) { req.GetBody = func() (io.ReadCloser, error) { return nil, boom } },
		response: debugLogResponse(200, "", http.NoBody),
		want:     "POST /api/2.0/x\n> <io.Reader>\n< HTTP/1.1 200 OK",
	}, {
		name: "request body that fails when it is read again", method: "POST", header: json, direct: true,
		body:     strings.NewReader(`{"password":"s3cret"}`),
		prepare:  func(req *http.Request) { req.GetBody = func() (io.ReadCloser, error) { return failing, nil } },
		response: debugLogResponse(200, "", http.NoBody),
		want:     "POST /api/2.0/x\n> <io.Reader>\n< HTTP/1.1 200 OK",
	}, {
		name: "empty request body", method: "POST", header: json, direct: true,
		body:     bytes.NewReader(nil),
		response: debugLogResponse(204, "", http.NoBody),
		want:     "POST /api/2.0/x\n< HTTP/1.1 204 No Content",
	}, {
		name:     "error response with an empty body",
		response: debugLogResponse(404, "", debugLogText("")),
		want:     "GET /api/2.0/x\n< HTTP/1.1 404 Not Found (Error: Not Found)",
	}, {
		name:     "error response without a body",
		response: debugLogResponse(500, "", nil),
		want:     "GET /api/2.0/x\n< HTTP/1.1 500 Internal Server Error (Error: Internal Server Error)",
	}, {
		name:     "response body that is not JSON, at the default truncation",
		response: debugLogResponse(502, "text/html", debugLogText(strings.Repeat("x", 100))),
		want:     "GET /api/2.0/x\n< HTTP/1.1 502 Bad Gateway\n< " + strings.Repeat("x", 96) + "... (4 more bytes)",
	}, {
		name:     "response body that fails to be read",
		response: debugLogResponse(200, "application/json", io.NopCloser(io.MultiReader(strings.NewReader(`{"token":"s3cret"`), iotest.ErrReader(boom)))),
		want:     "GET /api/2.0/x\n< HTTP/1.1 200 OK (Error: response body: boom)",
	}, {
		name: "response that the request asks for as a stream", header: stream, direct: true,
		response: debugLogResponse(200, "text/plain", debugLogText("the file")),
		want:     "GET /api/2.0/x\n< HTTP/1.1 200 OK\n< <Streaming response>",
	}, {
		name: "JSON response to a request for a stream", header: stream,
		response: debugLogResponse(404, "application/json", debugLogText(`{"error_code":"NOT_FOUND"}`)),
		want:     "GET /api/2.0/x\n< HTTP/1.1 404 Not Found\n< {\n<   \"error_code\": \"NOT_FOUND\"\n< }",
	}, {
		name: "connection that switches protocols", direct: true,
		response: debugLogResponse(101, "", debugLogText("not HTTP")),
		want:     "GET /api/2.0/x\n< HTTP/1.1 101 Switching Protocols\n< <Streaming response>",
	}, {
		name: "attempt that fails", method: "POST", url: "http://user:s3cret@u.test/p", err: boom,
		want: "POST /p\n< Error: Post \"http://user:***@u.test/p\": boom",
	}, {
		name: "attempt that fails with a response", response: debugLogResponse(200, "", debugLogText("ignored")), err: boom, direct: true,
		want: "GET /api/2.0/x\n< Error: Get \"" + debugLogURL + "\": boom",
	}, {
		name: "request without a method", prepare: func(req *http.Request) { req.Method = "" }, err: boom,
		want: " /api/2.0/x\n< Error: Get \"" + debugLogURL + "\": boom",
	}, {
		name: "HTTP server behind an HTTPS URL", url: "https://u.test/p",
		err:  tls.RecordHeaderError{Msg: "not TLS", RecordHeader: [5]byte{'H', 'T', 'T', 'P', '/'}},
		want: "GET /p\n< Error: Get \"https://u.test/p\": http: server gave HTTP response to HTTPS client",
	}, {
		name: "headers", headers: true, header: map[string]string{"Authorization": "Bearer dapi-123", "X-B": "2"},
		response: debugLogResponse(200, "application/json", debugLogText(`{"a":1}`)),
		want:     "GET /api/2.0/x\n> * Host: u.test\n> * Authorization: Bearer dapi-123\n> * X-B: 2\n< HTTP/1.1 200 OK\n< * Content-Type: application/json\n< {\n<   \"a\": 1\n< }",
	}, {
		name: "truncation", headers: true, truncate: 10, method: "POST", header: map[string]string{"Content-Type": "application/json"},
		body:     strings.NewReader(`{"name":"0123456789abcdef"}`),
		response: debugLogResponse(200, "text/plain", debugLogText("0123456789abcdef")),
		want: "POST /api/2.0/x\n> * Host: u.test\n> * Content-Type: applicatio... (6 more bytes)\n" +
			"> {\n>   \"name\": \"0123456789... (6 more bytes)\"\n> }\n< HTTP/1.1 200 OK\n< * Content-Type: text/plain\n< 0123456789... (6 more bytes)",
	}, {
		name: "debug logging is off", off: true, method: "POST", header: json, direct: true,
		body: strings.NewReader(`{"b":2}`), prepare: unread,
		response: debugLogResponse(200, "application/json", debugLogText(`{"a":1}`)),
	}, {
		name: "request without a URL", prepare: func(req *http.Request) { req.URL = nil }, direct: true,
		response: debugLogResponse(200, "application/json", debugLogText(`{"a":1}`)),
	}, {
		name: "transport that returns nothing",
	}} {
		t.Run(e.name, e.run)
	}
}

// debugLogCountedBody counts how often it is closed.
type debugLogCountedBody struct {
	io.Reader
	closes int
}

func (b *debugLogCountedBody) Close() error { b.closes++; return nil }

func TestDebugLog_RequestAndResponseAreUnchanged(t *testing.T) {
	const sent, served, all = `{"name":"m"}`, `{"name":"m","token":"s3cret"}`, 1 << 20
	const unread = "POST /api/2.0/x?a=b\n> {\n>   \"name\": \"m\"\n> }\n< HTTP/1.1 200 OK"
	const read = unread + "\n< {\n<   \"name\": \"m\",\n<   \"token\": \"**REDACTED**\"\n< }"
	for _, c := range []struct {
		name   string
		limit  int64 // how much of the body the caller reads
		parts  bool
		closes int
		want   string
	}{
		{"read at once", all, false, 1, read},
		{"read in parts", all, true, 1, read},
		{"closed early", 20, false, 1, unread},
		{"closed unread", 0, false, 1, unread},
		{"closed twice", all, false, 2, read},
		{"never closed", all, false, 0, read},
		{"neither read to its end nor closed", 20, false, 0, ""},
	} {
		t.Run(c.name, func(t *testing.T) {
			recorder := &debugLogRecorder{}
			req := recorder.request(t, "POST", debugLogURL+"?a=b", strings.NewReader(sent))
			req.Header.Set("Content-Type", "application/json")
			before, body := req.Clone(req.Context()), req.Body
			counted := &debugLogCountedBody{Reader: strings.NewReader(served)}
			response := debugLogResponse(200, "application/json", counted)
			header := response.Header.Clone()
			base := debugLogRoundTrip(func(got *http.Request) (*http.Response, error) {
				assert.Same(t, req, got)
				body, err := io.ReadAll(got.Body)
				assert.Equal(t, []any{sent, nil}, []any{string(body), err}, "the transport sends the whole body")
				return response, nil
			})
			resp, err := debugLogOver(t, false, 0, base).RoundTrip(req)
			require.NoError(t, err)
			assert.Same(t, response, resp)
			assert.Equal(t, []any{200, header}, []any{resp.StatusCode, resp.Header})
			reader := io.LimitReader(resp.Body, c.limit)
			if c.parts {
				reader = iotest.OneByteReader(reader)
			}
			got, err := io.ReadAll(reader)
			assert.Equal(t, []any{served[:min(c.limit, int64(len(served)))], nil}, []any{string(got), err})
			for range c.closes {
				assert.NoError(t, resp.Body.Close())
				// A body that is read after it is closed is not logged again.
				_, _ = io.Copy(io.Discard, resp.Body)
			}
			assert.Equal(t, c.closes, counted.closes)
			assert.Equal(t, c.want, strings.Join(recorder.logged(), "\n---\n"))
			assert.True(t, body == req.Body, "the request keeps its body")
			assert.Equal(t, []any{before.Method, before.URL, before.Header, before.Host, before.ContentLength},
				[]any{req.Method, req.URL, req.Header, req.Host, req.ContentLength})
		})
	}
}

func TestDebugLog_Concurrent(t *testing.T) {
	recorder := &debugLogRecorder{}
	echo := debugLogRoundTrip(func(req *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(req.Body)
		return debugLogResponse(200, "application/json", io.NopCloser(bytes.NewReader(body))), err
	})
	transport := debugLogOver(t, false, 0, echo)
	var want []string
	var group sync.WaitGroup
	for i := range 50 {
		body := fmt.Sprintf(`{"n":%d}`, i)
		want = append(want, fmt.Sprintf("POST /api/2.0/x/%d\n> {\n>   \"n\": %d\n> }\n< HTTP/1.1 200 OK\n< {\n<   \"n\": %d\n< }", i, i, i))
		group.Go(func() {
			req := recorder.request(t, "POST", fmt.Sprintf("%s/%d", debugLogURL, i), strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			resp, err := transport.RoundTrip(req)
			assert.NoError(t, err)
			got, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, body, string(got))
		})
	}
	group.Wait()
	assert.ElementsMatch(t, want, recorder.logged())

	// A caller may close a body while another goroutine reads it.
	recorder = &debugLogRecorder{}
	reader, writer := io.Pipe()
	piped := debugLogRoundTrip(func(*http.Request) (*http.Response, error) { return debugLogResponse(200, "", reader), nil })
	resp, err := debugLogOver(t, false, 0, piped).RoundTrip(recorder.request(t, "GET", debugLogURL, nil))
	require.NoError(t, err)
	group.Go(func() { _, _ = io.Copy(io.Discard, resp.Body) })
	group.Go(func() { _, _ = writer.Write([]byte("abc")); resp.Body.Close() })
	group.Wait()
	assert.Equal(t, []string{"GET /api/2.0/x\n< HTTP/1.1 200 OK"}, recorder.logged())
}

// The order that the doc comment of DebugLog asks for, and what another loses.
func TestDebugLog_Order(t *testing.T) {
	served := debugLogRoundTrip(func(*http.Request) (*http.Response, error) {
		return debugLogResponse(200, "", debugLogText("ok")), nil
	})
	setsHeader := func(next http.RoundTripper) http.RoundTripper {
		return debugLogRoundTrip(func(req *http.Request) (*http.Response, error) {
			req = req.Clone(req.Context())
			req.Header.Set("X-Set-By", "a transport")
			return next.RoundTrip(req)
		})
	}
	endsAttempt := func(http.RoundTripper) http.RoundTripper {
		return debugLogRoundTrip(func(*http.Request) (*http.Response, error) {
			return nil, errors.New("request timed out after 1m0s of inactivity")
		})
	}
	logs := func(next http.RoundTripper) http.RoundTripper { return debugLogOver(t, true, 0, next) }
	for _, c := range []struct {
		name      string
		transport http.RoundTripper
		want      []string
	}{
		{"inside the transport that sets a header and outside the one that ends the attempt", setsHeader(logs(endsAttempt(served))),
			[]string{"GET /api/2.0/x\n> * Host: u.test\n> * X-Set-By: a transport\n< Error: Get \"" + debugLogURL + "\": request timed out after 1m0s of inactivity"}},
		{"outside the transport that sets a header", logs(setsHeader(served)), []string{"GET /api/2.0/x\n> * Host: u.test\n< HTTP/1.1 200 OK\n< ok"}},
		{"inside the transport that ends the attempt", endsAttempt(logs(served)), nil},
	} {
		t.Run(c.name, func(t *testing.T) {
			recorder := &debugLogRecorder{}
			if resp, err := c.transport.RoundTrip(recorder.request(t, "GET", debugLogURL, nil)); err == nil {
				_, _ = io.Copy(io.Discard, resp.Body)
			}
			assert.Equal(t, c.want, recorder.logged())
		})
	}
}

// debugLogServed is what one httptest.Server answers to a call of the
// classic client and to the same call of a modular client over DebugLog.
type debugLogServed struct {
	name, call        string
	status            int
	contentType, body string
	// fail is "close" for a connection that is closed without a response and
	// "cut" for a body that ends before its announced length.
	fail string
	// shows is a part of the message at the default debug_truncate_bytes.
	shows string
}

func (s debugLogServed) serve(w http.ResponseWriter, r *http.Request) {
	_, _ = io.Copy(io.Discard, r.Body)
	if s.fail == "close" {
		conn, _, _ := w.(http.Hijacker).Hijack()
		conn.Close()
		return
	}
	w.Header().Set("Date", "Mon, 02 Jan 2006 15:04:05 GMT")
	if s.contentType != "" {
		w.Header().Set("Content-Type", s.contentType)
	}
	if s.fail == "cut" {
		w.Header().Set("Content-Length", strconv.Itoa(len(s.body)+10))
	}
	w.WriteHeader(s.status)
	_, _ = io.WriteString(w, s.body)
}

// classic sends the call through the classic client.
func (s debugLogServed) classic(ctx context.Context, w *databricks.WorkspaceClient) {
	switch s.call {
	case "GET":
		_, _ = w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: "a model"})
	case "POST":
		_, _ = w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{Name: "m", Description: strings.Repeat("long ", 30)})
	case "PATCH":
		_, _ = w.ModelRegistry.UpdateModel(ctx, ml.UpdateModelRequest{Name: "m", Description: "d"})
	case "DELETE":
		_ = w.ModelRegistry.DeleteModel(ctx, ml.DeleteModelRequest{Name: "m"})
	}
}

// modular sends the same request as classic through the modular client.
func (s debugLogServed) modular(ctx context.Context, c *modelregistry.Client) {
	switch s.call {
	case "GET":
		_, _ = c.GetRegisteredModelDatabricks(ctx, modelregistry.GetRegisteredModelDatabricksRequest{Name: new("a model")})
	case "POST":
		_, _ = c.CreateRegisteredModel(ctx, modelregistry.CreateRegisteredModelRequest{Name: new("m"), Description: new(strings.Repeat("long ", 30))})
	case "PATCH":
		_, _ = c.UpdateRegisteredModel(ctx, modelregistry.UpdateRegisteredModelRequest{Name: new("m"), Description: new("d")})
	case "DELETE":
		_, _ = c.DeleteRegisteredModel(ctx, modelregistry.DeleteRegisteredModelRequest{Name: new("m")})
	}
}

// logs returns the debug messages of the classic client and of the modular
// client for the call, both against one server.
func (s debugLogServed) logs(t *testing.T, headers bool, truncate int) (classic, modular []string) {
	server := httptest.NewServer(http.HandlerFunc(s.serve))
	t.Cleanup(server.Close)
	recorder := &debugLogRecorder{}
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(debugLogConfig(server.URL, headers, truncate)))
	require.NoError(t, err)
	s.classic(logger.NewContext(context.Background(), recorder), w)
	classic = recorder.logged()
	recorder = &debugLogRecorder{}
	creds, err := credentials.NewPATCredentials("dapi-debug-log")
	require.NoError(t, err)
	transport, err := DebugLog(debugLogConfig(server.URL, headers, truncate), http.DefaultTransport)
	require.NoError(t, err)
	c, err := modelregistry.NewClient(context.Background(), client.WithHost(server.URL), client.WithCredentials(creds),
		client.WithHTTPClient(&http.Client{Transport: transport}), client.WithoutConfigFile(), client.WithoutEnv())
	require.NoError(t, err)
	s.modular(logger.NewContext(context.Background(), recorder), c)
	return classic, recorder.logged()
}

// debugLogDifferingHeader matches the request header lines that differ
// between the two clients, see TestDebugLog_DiffersFromClassicClient.
var debugLogDifferingHeader = regexp.MustCompile(`(?m)^> \* (Host|Accept|Traceparent|User-Agent): .*\n`)

func TestDebugLog_MatchesClassicClient(t *testing.T) {
	const model = `{"registered_model_databricks":{"name":"a model","description":"d","id":"1"}}`
	long := `{"registered_model_databricks":{"name":"a model","description":"` + strings.Repeat("long ", 30) + `"}}`
	list := `{"registered_model_databricks":{"latest_versions":[` + strings.Repeat(`{"description":"`+strings.Repeat("v", 90)+`"},`, 15) + `{}]}}`
	for _, s := range []debugLogServed{
		{name: "GET, 200 with a JSON body", call: "GET", status: 200, contentType: "application/json", body: model, shows: "GET /api/2.0/mlflow/databricks/registered-models/get?name=a model\n"},
		{name: "POST with a JSON body", call: "POST", status: 200, contentType: "application/json", body: model, shows: "> {\n>   \"description\": \"long long "},
		{name: "PATCH", call: "PATCH", status: 200, contentType: "application/json", body: model, shows: "> {\n>   \"description\": \"d\",\n>   \"name\": \"m\"\n> }\n< HTTP/1.1 200 OK"},
		{name: "DELETE", call: "DELETE", status: 200, contentType: "application/json", body: "{}", shows: "DELETE /api/2.0/mlflow/registered-models/delete?name=m\n"},
		{name: "200 with an empty body", call: "DELETE", status: 200, shows: "< HTTP/1.1 200 OK"},
		{name: "200 with an empty body to a GET", call: "GET", status: 200, contentType: "application/json", shows: "< HTTP/1.1 200 OK"},
		{name: "200 with a string longer than the truncation", call: "GET", status: 200, contentType: "application/json", body: long, shows: "long l... (54 more bytes)\""},
		{name: "200 with a list longer than the truncation", call: "GET", status: 200, contentType: "application/json", body: list, shows: "additional elements)\""},
		{name: "200 with a body that is not JSON", call: "PATCH", status: 200, contentType: "text/plain", body: strings.Repeat("plain ", 40) + "\nsecond line", shows: "< plain plain "},
		{name: "200 with secrets", call: "POST", status: 200, contentType: "application/json", body: `{"registered_model":{"name":"m"},"token":"s3cret","password":"s3cret"}`, shows: "\"token\": \"**REDACTED**\""},
		{name: "error response", call: "GET", status: 404, contentType: "application/json", body: `{"error_code":"RESOURCE_DOES_NOT_EXIST","message":"no model"}`, shows: "< HTTP/1.1 404 Not Found\n"},
		{name: "error response to a POST", call: "POST", status: 400, contentType: "application/json", body: `{"error_code":"INVALID_PARAMETER_VALUE","message":"bad"}`, shows: "< HTTP/1.1 400 Bad Request\n"},
		{name: "error response with an empty body", call: "PATCH", status: 403, shows: "< HTTP/1.1 403 Forbidden (Error: Forbidden)"},
		{name: "error response that is not JSON", call: "DELETE", status: 500, contentType: "text/html", body: "<html><pre>it broke</pre></html>", shows: "< <html><pre>it broke</pre></html>"},
		{name: "request that fails before a response arrives", call: "GET", fail: "close", shows: "\n< Error: Get \"http://127.0.0.1:"},
		{name: "POST that fails before a response arrives", call: "POST", fail: "close", shows: "\n< Error: Post \"http://127.0.0.1:"},
		{name: "response body that ends early", call: "GET", status: 200, contentType: "application/json", body: model, fail: "cut", shows: "< HTTP/1.1 200 OK (Error: response body: unexpected EOF)"},
	} {
		for _, settings := range []struct {
			headers  bool
			truncate int
		}{{false, 0}, {true, 0}, {true, 20}, {false, 2000}} {
			t.Run(fmt.Sprintf("%s/debug_headers=%t/debug_truncate_bytes=%d", s.name, settings.headers, settings.truncate), func(t *testing.T) {
				classic, modular := s.logs(t, settings.headers, settings.truncate)
				require.Len(t, modular, 1)
				assert.True(t, settings.truncate != 0 || strings.Contains(modular[0], s.shows), modular[0])
				assert.NotContains(t, modular[0], "s3cret")
				// The retry loop of the classic client adds a message of its
				// own for an attempt that failed. DebugLog doesn't retry.
				if s.status >= 400 || s.fail != "" {
					require.Len(t, classic, 2)
					assert.Contains(t, classic[1], "non-retriable error: ")
				} else {
					require.Len(t, classic, 1)
				}
				assert.Equal(t, debugLogDifferingHeader.ReplaceAllString(classic[0], ""), debugLogDifferingHeader.ReplaceAllString(modular[0], ""))
				assert.Equal(t, settings.headers, classic[0] != modular[0], "the messages differ only with debug_headers")
			})
		}
	}
}

// The request header lines that differ between the two clients. Every other
// line of the two messages is the same.
func TestDebugLog_DiffersFromClassicClient(t *testing.T) {
	headersOnlyIn := func(message, other string) (names []string) {
		for _, line := range strings.SplitAfter(message, "\n") {
			if !strings.Contains(other, line) {
				names = append(names, debugLogDifferingHeader.ReplaceAllString(line, "$1"))
			}
		}
		return names
	}
	s := debugLogServed{call: "POST", status: 200, contentType: "application/json", body: "{}"}
	classic, modular := s.logs(t, true, 0)
	require.Equal(t, []int{1, 1}, []int{len(classic), len(modular)})
	assert.Equal(t, []string{"Host", "Accept", "Traceparent", "User-Agent"}, headersOnlyIn(classic[0], modular[0]))
	assert.Equal(t, []string{"Host", "User-Agent"}, headersOnlyIn(modular[0], classic[0]))
	assert.Contains(t, classic[0], "\n> * Host: \n> * Accept: application/json\n> * Authorization: Bearer dapi-debug-log\n> * Content-Type: application/json\n> * Traceparent: 00-")
	assert.Regexp(t, `\n> \* Host: 127\.0\.0\.1:\d+\n> \* Authorization: Bearer dapi-debug-log\n> \* Content-Type: application/json\n> \* User-Agent: `, modular[0])
}
