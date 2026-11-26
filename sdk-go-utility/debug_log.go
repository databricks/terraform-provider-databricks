package sdkgoutility

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/logger/httplog"
)

// DebugLog returns a transport that sends each request through transport and
// writes it to the debug log as the classic client does: one message for
// every attempt, with the method and the URL, the headers when debug_headers
// is set, the request body, and the status and the body of the response or
// the error of the attempt. A modular client writes no such message, so a
// resource that moves to one keeps its request log by sending its requests
// through this transport.
//
// The message is the one of the classic client's recordRequestLog. It is
// built by httplog.RoundTripStringer, which redacts secrets and cuts long
// values, with debug_headers and debug_truncate_bytes of cfg, and written
// with logger.Debugf of the classic SDK, which the provider sends to the
// Terraform log. It is written when the classic client writes it:
//
//   - once the caller has read the response body to its end, with the body,
//     or at once for a response without a body;
//   - when reading the body fails, with that error and without the body;
//   - when the attempt fails before a response arrives, with the error as
//     http.Client reports it to the classic client;
//   - at once, with <Streaming response> for the body, for a response that the
//     classic client hands to its caller unread: one that is not JSON, to a
//     request whose Accept header asks for something other than JSON.
//
// The body of a request is logged when it is JSON that net/http can read
// again, as the bodies that a modular client marshals are. Any other body is
// one that the caller streams, and it is logged as <io.Reader> without being
// read, as the classic client logs a body that it is given as a reader. A
// caller that closes the response body before its end gets the message
// without the body, because a part of a body can't be redacted, and one that
// neither reads the body to its end nor closes it gets no message.
//
// When debug logging is off for the context of a request, as
// logger.Get(ctx).Enabled(ctx, logger.LevelDebug) tells, DebugLog only calls
// transport: it reads no body and keeps no copy. The caller always gets what
// transport returns, the response with its status, headers and body bytes or
// the error, and its request is never modified. The returned transport is
// safe for concurrent requests.
//
// DebugLog leaves the place of this transport among the others of a client to
// its caller. It logs the request as it receives it and the outcome that the
// transport below it returns, so it belongs:
//
//   - inside every transport that changes the request, such as one that sets
//     a header, so that the log shows the headers that are sent;
//   - outside every transport that can end an attempt itself, such as one
//     that times an attempt out, so that the log shows the error that ended
//     the attempt.
//
// A transport outside DebugLog that fails without calling it leaves no
// message, as the classic client writes none for a request that its rate
// limiter or a request visitor refuses. DebugLog also leaves retries to its
// caller: it never sends a request twice, and every request that reaches it
// is an attempt with a message of its own, also each request of a redirect
// that http.Client follows.
//
// DebugLog returns an error when cfg or transport is nil, or when cfg can't
// be resolved.
func DebugLog(cfg *config.Config, transport http.RoundTripper) (http.RoundTripper, error) {
	if cfg == nil {
		return nil, errors.New("DebugLog: cfg is nil")
	}
	if transport == nil {
		return nil, errors.New("DebugLog: transport is nil")
	}
	clientConfig, err := config.HTTPClientConfigFromConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("DebugLog: %w", err)
	}
	truncateBytes := clientConfig.DebugTruncateBytes
	if truncateBytes == 0 {
		truncateBytes = debugLogDefaultTruncateBytes
	}
	return &debugLogTransport{
		transport:     transport,
		headers:       clientConfig.DebugHeaders,
		truncateBytes: truncateBytes,
		errorMapper:   clientConfig.ErrorMapper,
	}, nil
}

// debugLogDefaultTruncateBytes is what httpclient.NewApiClient uses for a
// debug_truncate_bytes that is not set.
const debugLogDefaultTruncateBytes = 96

type debugLogTransport struct {
	transport     http.RoundTripper
	headers       bool
	truncateBytes int
	errorMapper   func(context.Context, common.ResponseWrapper) error
}

func (t *debugLogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	// The message can't be built without a URL, and http.Client sends no
	// request without one.
	if req.URL == nil || !logger.Get(ctx).Enabled(ctx, logger.LevelDebug) {
		return t.transport.RoundTrip(req)
	}
	attempt := &debugLogAttempt{transport: t, request: req, requestBody: debugLogRequestBody(req)}
	resp, err := t.transport.RoundTrip(req)
	switch {
	case err != nil:
		attempt.write(nil, nil, debugLogClientError(req, err))
	case resp == nil:
		// Not a valid outcome. http.Client reports it to the caller.
	case resp.Body == nil || resp.Body == http.NoBody:
		// There is nothing that the caller has to read first.
		attempt.writeRead(resp, nil)
	case debugLogIsStream(req, resp):
		attempt.write(resp, []byte("<Streaming response>"), nil)
	default:
		resp.Body = &debugLogBody{ReadCloser: resp.Body, attempt: attempt, response: resp}
	}
	return resp, err
}

// debugLogRequestBody returns what the classic client logs as the body of
// req: the bytes of a JSON body, and <io.Reader> for a body that it is given
// as a reader. A JSON body is read from GetBody, which net/http sets for a
// body that is held in memory, so the body that transport sends stays
// untouched.
func debugLogRequestBody(req *http.Request) []byte {
	if req.Body == nil || req.Body == http.NoBody {
		return nil
	}
	if req.GetBody != nil && debugLogIsJSON(req.Header) {
		if body, err := req.GetBody(); err == nil {
			defer body.Close()
			if bs, err := io.ReadAll(body); err == nil {
				return bs
			}
		}
	}
	return []byte("<io.Reader>")
}

func debugLogIsJSON(header http.Header) bool {
	return strings.HasPrefix(header.Get("Content-Type"), "application/json")
}

// debugLogIsStream tells whether resp is one that the classic client hands
// to its caller without reading it, see common.NewResponseWrapper. The
// classic client sets an Accept header on every request and reads the
// response when that asks for JSON. A modular client sets the header only
// when it asks for something else, so a request without one counts as asking
// for JSON. The body of a connection that switches protocols is not a
// response body, and it is left as it is as well.
func debugLogIsStream(req *http.Request, resp *http.Response) bool {
	if resp.StatusCode == http.StatusSwitchingProtocols {
		return true
	}
	accept := req.Header.Get("Accept")
	return accept != "" && accept != "application/json" && !debugLogIsJSON(resp.Header)
}

// debugLogClientError returns the error that http.Client makes of err, the
// error of a transport, which is what the classic client logs for an attempt
// that fails before a response arrives.
func debugLogClientError(req *http.Request, err error) error {
	if tlsErr, ok := err.(tls.RecordHeaderError); ok && string(tlsErr.RecordHeader[:]) == "HTTP/" {
		err = http.ErrSchemeMismatch
	}
	op := "Get"
	if req.Method != "" {
		op = req.Method[:1] + strings.ToLower(req.Method[1:])
	}
	requestURL := req.URL.String()
	if _, hasPassword := req.URL.User.Password(); hasPassword {
		requestURL = strings.Replace(requestURL, req.URL.User.String()+"@", req.URL.User.Username()+":***@", 1)
	}
	return &url.Error{Op: op, URL: requestURL, Err: err}
}

// debugLogAttempt is one request on its way through the transport.
type debugLogAttempt struct {
	transport   *debugLogTransport
	request     *http.Request
	requestBody []byte
}

// write writes the message of the attempt. A nil resp is an attempt that
// failed before a response arrived.
func (a *debugLogAttempt) write(resp *http.Response, body []byte, err error) {
	message := httplog.RoundTripStringer{
		Request:                  a.request,
		Response:                 resp,
		Err:                      err,
		RequestBody:              a.requestBody,
		ResponseBody:             body,
		DebugHeaders:             a.transport.headers,
		DebugTruncateBytes:       a.transport.truncateBytes,
		DebugAuthorizationHeader: true,
	}.String()
	logger.Debugf(a.request.Context(), "%s", message)
}

// writeRead writes the message for a response whose body has been read to
// its end. For an empty body the classic client adds the error that its error
// mapper makes of the response, so the mapper of the config makes it here.
func (a *debugLogAttempt) writeRead(resp *http.Response, body []byte) {
	var err error
	if len(body) == 0 {
		// The mapper reads the request of the response, which a transport
		// other than http.Transport may not set.
		mapped := *resp
		mapped.Request = a.request
		mapped.Body = http.NoBody
		err = a.transport.errorMapper(a.request.Context(), common.ResponseWrapper{
			Response:   &mapped,
			ReadCloser: http.NoBody,
		})
	}
	a.write(resp, body, err)
}

// debugLogBody is a response body that keeps a copy of what its caller reads
// and writes the message of the attempt once, when the caller is done with it.
type debugLogBody struct {
	io.ReadCloser
	attempt  *debugLogAttempt
	response *http.Response

	// mu guards the fields below, for a caller that closes the body while
	// another goroutine reads it.
	mu      sync.Mutex
	read    bytes.Buffer
	written bool
}

func (b *debugLogBody) Read(p []byte) (int, error) {
	n, err := b.ReadCloser.Read(p)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.written {
		return n, err
	}
	b.read.Write(p[:n])
	switch {
	case err == io.EOF:
		b.written = true
		b.attempt.writeRead(b.response, b.read.Bytes())
	case err != nil:
		b.written = true
		// The classic client logs neither the part that it read nor
		// anything else of a body that it failed to read.
		b.attempt.write(b.response, nil, fmt.Errorf("response body: %w", err))
	}
	return n, err
}

func (b *debugLogBody) Close() error {
	b.mu.Lock()
	if !b.written {
		b.written = true
		b.attempt.write(b.response, nil, nil)
	}
	b.mu.Unlock()
	return b.ReadCloser.Close()
}
