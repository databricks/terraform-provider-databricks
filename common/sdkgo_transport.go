package common

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	sdkcommon "github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/traceparent"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/logger/httplog"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// Default of httpclient.ClientConfig.DebugTruncateBytes.
const classicDefaultDebugTruncateBytes = 96

// sdkGoTransport sends each sdk-go attempt the way httpclient.ApiClient sends one: it
// sends the classic request, runs the request visitors under the attempt's timeout,
// follows redirects, reads the response, rejects successful responses the classic error
// mapper rejects, and writes the classic debug log.
type sdkGoTransport struct {
	// client follows redirects over the classic transport, as the classic client does.
	client             *http.Client
	visitors           []httpclient.RequestVisitor
	errorMapper        func(context.Context, sdkcommon.ResponseWrapper) error
	httpTimeout        time.Duration
	debugHeaders       bool
	debugTruncateBytes int
}

func newSdkGoTransport(cfg httpclient.ClientConfig, base http.RoundTripper,
	httpTimeout time.Duration) *sdkGoTransport {
	truncate := cfg.DebugTruncateBytes
	if truncate == 0 {
		truncate = classicDefaultDebugTruncateBytes
	}
	return &sdkGoTransport{
		client:             &http.Client{Transport: base},
		visitors:           cfg.Visitors,
		errorMapper:        cfg.ErrorMapper,
		httpTimeout:        httpTimeout,
		debugHeaders:       cfg.DebugHeaders,
		debugTruncateBytes: truncate,
	}
}

func (t *sdkGoTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	attempt, ok := req.Context().Value(sdkGoAttemptKey{}).(*sdkGoAttempt)
	if !ok {
		return nil, errors.New("sdk-go request sent without callSdkGo")
	}
	timeout := attempt.takeTimeout()
	if timeout == nil {
		timeout = startAttemptTimeout(req.Context(), t.httpTimeout)
	}
	// Like the classic client, the visitors already run under the attempt's timeout.
	r := req.Clone(timeout.ctx)
	if err := attempt.useClassicRequest(r); err != nil {
		timeout.end()
		return nil, err
	}
	for _, visit := range t.visitors {
		if err := visit(r); err != nil {
			timeout.end()
			return nil, &visitorError{err}
		}
	}
	attempt.setClassicHeaders(r)
	traceparent.AddTraceparent(r)
	r.Header.Set("User-Agent", sdkGoUserAgent(r.Context(), req.Header.Get("User-Agent")))
	requestBody := debugRequestBody(r)
	resp, responseBody, err := t.exchange(req.Context(), r, timeout)
	attempt.request, attempt.response = r, nil
	if resp != nil {
		attempt.request, attempt.response = resp.Request, resp
	}
	if err == nil && resp.StatusCode < http.StatusBadRequest {
		// A Private Link redirect ends in a successful login page.
		err = t.errorMapper(r.Context(), sdkcommon.ResponseWrapper{
			Response:   resp,
			ReadCloser: io.NopCloser(bytes.NewReader(responseBody)),
			DebugBytes: responseBody,
		})
	}
	t.log(r, resp, err, requestBody, responseBody)
	if err == nil {
		err = attempt.decodeLikeClassicClient(resp, requestBody, responseBody)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// decodeLikeClassicClient decodes a successful response into the classic type of the
// call, as httpclient.WithResponseUnmarshal does, and gives sdk-go that value's encoding.
// So sdk-go reads the values the classic client reads, such as "123" for a string field
// that the response sets to the number 123.
func (a *sdkGoAttempt) decodeLikeClassicClient(resp *http.Response,
	requestBody, responseBody []byte) error {
	if a.classic.response == nil || resp.StatusCode >= http.StatusBadRequest {
		return nil
	}
	response := a.classic.response()
	// Like the classic client, an empty body leaves the response empty.
	if len(responseBody) > 0 {
		if err := json.Unmarshal(responseBody, &response); err != nil {
			return &responseDecodeError{
				classicDecodeError(err, resp, requestBody, responseBody)}
		}
	}
	body, err := json.Marshal(response)
	if err != nil {
		return &responseDecodeError{err}
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	return nil
}

// Mirrors the unexported httpclient.makeUnexpectedResponse.
const classicDecodeErrorRequestLog = "This is likely a bug in the Databricks SDK for " +
	"Go or the underlying REST API. Please report this issue with the following " +
	"debugging information to the SDK issue tracker at " +
	"https://github.com/databricks/databricks-sdk-go/issues. Request log:\n```\n%s\n```"

// classicDecodeError returns the error of httpclient.WithResponseUnmarshal for err, the
// failure to decode the successful response resp.
func classicDecodeError(err error, resp *http.Response,
	requestBody, responseBody []byte) error {
	// The classic client asserts the error's type rather than unwrapping it.
	if _, ok := err.(*json.SyntaxError); ok && bytes.HasPrefix(responseBody, []byte("<")) {
		return httpclient.ErrHTMLContent
	}
	requestLog := httplog.RoundTripStringer{
		Request:            resp.Request,
		Response:           resp,
		RequestBody:        requestBody,
		ResponseBody:       responseBody,
		DebugHeaders:       true,
		DebugTruncateBytes: 10 * 1024,
	}
	return fmt.Errorf("failed to unmarshal response body: %w. "+
		classicDecodeErrorRequestLog, err, requestLog.String())
}

// exchange sends r, whose context ends with timeout, follows redirects and reads the
// response body. Like the classic client, it returns the url.Error of the request it
// sent, with the timeout message if the attempt timed out before parent ended.
func (t *sdkGoTransport) exchange(parent context.Context, r *http.Request,
	timeout *attemptTimeout) (*http.Response, []byte, error) {
	defer timeout.end()
	if r.Body != nil && r.Body != http.NoBody {
		r.Body = &inactivityBody{ReadCloser: r.Body, timer: timeout.timer}
	}
	resp, err := t.client.Do(r)
	if err != nil {
		urlErr, ok := err.(*url.Error)
		if ok && timeout.timer.expired() && parent.Err() == nil {
			urlErr.Err = fmt.Errorf("request timed out after %s of inactivity",
				t.httpTimeout)
		}
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(&inactivityBody{ReadCloser: resp.Body, timer: timeout.timer})
	if err != nil {
		return nil, nil, &responseBodyError{err}
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	if resp.Request == nil {
		resp.Request = r
	}
	return resp, body, nil
}

// debugRequestBody returns a copy of the request body for the debug log.
func debugRequestBody(r *http.Request) []byte {
	if r.GetBody == nil {
		return nil
	}
	body, err := r.GetBody()
	if err != nil {
		return nil
	}
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		return nil
	}
	return data
}

// log writes the debug line that httpclient.ApiClient.recordRequestLog writes.
func (t *sdkGoTransport) log(r *http.Request, resp *http.Response, err error,
	requestBody, responseBody []byte) {
	ctx := r.Context()
	if !logger.Get(ctx).Enabled(ctx, logger.LevelDebug) {
		return
	}
	message := httplog.RoundTripStringer{
		Request:                  r,
		Response:                 resp,
		Err:                      err,
		RequestBody:              requestBody,
		ResponseBody:             responseBody,
		DebugHeaders:             t.debugHeaders,
		DebugTruncateBytes:       t.debugTruncateBytes,
		DebugAuthorizationHeader: true,
	}.String()
	logger.Debugf(ctx, "%s", message)
}

// sdkGoUserAgent keeps the classic User-Agent, which usage analytics parse, and appends
// the sdk-go module segments, such as sdk-go-uc-volumes/0.8.0, as the new SDK's marker.
func sdkGoUserAgent(ctx context.Context, sdkGoUserAgent string) string {
	segments := []string{useragent.FromContext(ctx)}
	for _, segment := range strings.Fields(sdkGoUserAgent) {
		if strings.HasPrefix(segment, "sdk-go-") {
			segments = append(segments, segment)
		}
	}
	return strings.Join(segments, " ")
}

// responseBodyError is a failure to read a response body, which the classic client
// reports with this message and doesn't retry.
type responseBodyError struct {
	err error
}

func (e *responseBodyError) Error() string {
	return "response body: " + e.err.Error()
}

func (e *responseBodyError) Unwrap() error {
	return e.err
}

// attemptTimeout cancels one attempt after its timeout passes without progress. It
// starts before authentication and lasts until the response body is read.
type attemptTimeout struct {
	ctx    context.Context
	cancel context.CancelFunc
	timer  *inactivityTimer
}

func startAttemptTimeout(parent context.Context, timeout time.Duration) *attemptTimeout {
	ctx, cancel := context.WithCancel(parent)
	timer := newInactivityTimer(timeout, cancel)
	return &attemptTimeout{ctx: ctx, cancel: cancel, timer: timer}
}

func (t *attemptTimeout) end() {
	t.timer.stop()
	t.cancel()
}

type inactivityTimer struct {
	timeout time.Duration
	timer   *time.Timer
	fired   atomic.Bool
}

func newInactivityTimer(timeout time.Duration,
	cancel context.CancelFunc) *inactivityTimer {
	t := &inactivityTimer{timeout: timeout}
	t.timer = time.AfterFunc(timeout, func() {
		t.fired.Store(true)
		cancel()
	})
	return t
}

func (t *inactivityTimer) extend() {
	if !t.fired.Load() {
		t.timer.Reset(t.timeout)
	}
}

func (t *inactivityTimer) stop() {
	t.timer.Stop()
}

func (t *inactivityTimer) expired() bool {
	return t.fired.Load()
}

// inactivityBody extends the timer on every read of a request or response body.
type inactivityBody struct {
	io.ReadCloser
	timer *inactivityTimer
}

func (b *inactivityBody) Read(p []byte) (int, error) {
	n, err := b.ReadCloser.Read(p)
	b.timer.extend()
	return n, err
}
