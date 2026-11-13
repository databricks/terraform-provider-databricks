package common

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/traceparent"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/sdk-go/auth"
	"golang.org/x/time/rate"
)

// Defaults of httpclient.NewApiClient.
const (
	classicDefaultHTTPTimeout  = 60 * time.Second
	classicDefaultRetryTimeout = 5 * time.Minute
)

// sdkGoBridge holds the classic client settings that sdk-go clients reuse.
type sdkGoBridge struct {
	httpConfig  httpclient.ClientConfig
	retry       retries.RetryConfig
	httpTimeout time.Duration
	limiter     *rate.Limiter
	httpClient  *http.Client
}

// newSdkGoBridge reads the settings that api, the classic API client built from cfg,
// derives from cfg, so both clients follow retry_timeout_seconds, rate_limit,
// http_timeout_seconds and the TLS and proxy settings, and share api's rate limiter.
func newSdkGoBridge(cfg *config.Config, api *httpclient.ApiClient) (*sdkGoBridge, error) {
	httpConfig, err := config.HTTPClientConfigFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	limiter, err := classicRateLimiter(api)
	if err != nil {
		return nil, err
	}
	httpTimeout := httpConfig.HTTPTimeout
	if httpTimeout == 0 {
		httpTimeout = classicDefaultHTTPTimeout
	}
	retryTimeout := httpConfig.RetryTimeout
	if retryTimeout == 0 {
		retryTimeout = classicDefaultRetryTimeout
	}
	base := classicTransport(httpConfig)
	// httpclient.NewApiClient makes the same exception for test fixture transports.
	skippable, ok := base.(interface{ SkipRetryOnIO() bool })
	if ok && skippable.SkipRetryOnIO() {
		retryTimeout = 0
	}
	b := &sdkGoBridge{
		httpConfig:  httpConfig,
		httpTimeout: httpTimeout,
		limiter:     limiter,
		httpClient: &http.Client{Transport: &sdkGoTransport{
			base:        base,
			visitors:    httpConfig.Visitors,
			httpTimeout: httpTimeout,
		}},
	}
	retries.WithTimeout(retryTimeout)(&b.retry)
	return b, nil
}

// classicRateLimiter returns the rate limiter of api. httpclient.ApiClient doesn't
// export it, but sharing it is the only way to keep sdk-go and classic calls within one
// rate_limit budget.
func classicRateLimiter(api *httpclient.ApiClient) (*rate.Limiter, error) {
	field := reflect.ValueOf(api).Elem().FieldByName("rateLimiter")
	if !field.IsValid() || field.Type() != reflect.TypeFor[*rate.Limiter]() {
		return nil, fmt.Errorf("cannot share the rate limiter of %T: "+
			"its rateLimiter field is missing or has another type", api)
	}
	return (*rate.Limiter)(field.UnsafePointer()), nil
}

// workspaceApiClient returns the API client of w. databricks.WorkspaceClient doesn't
// export it, but sdk-go calls must share its rate limiter, as w's own calls do.
func workspaceApiClient(w *databricks.WorkspaceClient) (*httpclient.ApiClient, error) {
	field := reflect.ValueOf(w).Elem().FieldByName("apiClient")
	if !field.IsValid() || field.Type() != reflect.TypeFor[*httpclient.ApiClient]() {
		return nil, fmt.Errorf("cannot share the API client of %T: "+
			"its apiClient field is missing or has another type", w)
	}
	api := (*httpclient.ApiClient)(field.UnsafePointer())
	if api == nil {
		// Mock workspace clients have no API client.
		return nil, fmt.Errorf("%T has no API client to share", w)
	}
	return api, nil
}

var (
	classicSharedTransport   = newClassicTransport(false)
	classicInsecureTransport = newClassicTransport(true)
)

// classicTransport mirrors httpclient.ClientConfig.httpTransport, which is unexported.
func classicTransport(cfg httpclient.ClientConfig) http.RoundTripper {
	if cfg.Transport != nil {
		return cfg.Transport
	}
	if cfg.InsecureSkipVerify {
		return classicInsecureTransport
	}
	return classicSharedTransport
}

// newClassicTransport mirrors the unexported httpclient.makeDefaultTransport.
func newClassicTransport(insecureSkipVerify bool) *http.Transport {
	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		IdleConnTimeout:       180 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	if insecureSkipVerify {
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return t
}

// classicCredentials asks the classic config to authenticate every request, so token
// refresh and extra authentication headers keep working. Like the classic client, it
// authenticates under the attempt's timeout, which it hands to sdkGoTransport.
type classicCredentials struct {
	cfg         *config.Config
	httpTimeout time.Duration
}

// Name is only used by sdk-go's own User-Agent, which sdkGoTransport replaces.
func (classicCredentials) Name() string {
	return "classic-config"
}

func (c classicCredentials) AuthHeaders(ctx context.Context) ([]auth.Header, error) {
	timeout := startAttemptTimeout(ctx, c.httpTimeout)
	req, err := http.NewRequestWithContext(timeout.ctx, http.MethodGet, c.cfg.Host, nil)
	if err == nil {
		if authErr := c.cfg.Authenticate(req); authErr != nil {
			err = &visitorError{authErr}
		}
	}
	attempt, ok := ctx.Value(sdkGoAttemptKey{}).(*sdkGoAttempt)
	if err != nil || !ok {
		timeout.end()
		if err != nil {
			return nil, err
		}
	} else {
		attempt.keepTimeout(timeout)
	}
	var headers []auth.Header
	for key, values := range req.Header {
		for _, value := range values {
			headers = append(headers, auth.Header{Key: key, Value: value})
		}
	}
	return headers, nil
}

// sdkGoTransport applies the classic client's request visitors, tracing header,
// User-Agent and per-attempt inactivity timeout to sdk-go requests.
type sdkGoTransport struct {
	base        http.RoundTripper
	visitors    []httpclient.RequestVisitor
	httpTimeout time.Duration
}

func (t *sdkGoTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	attempt, _ := req.Context().Value(sdkGoAttemptKey{}).(*sdkGoAttempt)
	var timeout *attemptTimeout
	if attempt != nil {
		timeout = attempt.takeTimeout()
	}
	if timeout == nil {
		timeout = startAttemptTimeout(req.Context(), t.httpTimeout)
	}
	// Like the classic client, the visitors already run under the attempt's timeout.
	r := req.Clone(timeout.ctx)
	for _, visit := range t.visitors {
		if err := visit(r); err != nil {
			timeout.end()
			return nil, &visitorError{err}
		}
	}
	traceparent.AddTraceparent(r)
	r.Header.Set("User-Agent", sdkGoUserAgent(r.Context(), req.Header.Get("User-Agent")))
	if attempt != nil {
		attempt.request = r
	}
	return t.send(req.Context(), r, timeout)
}

// sdkGoUserAgent keeps the classic User-Agent, which usage analytics parse, and appends
// the sdk-go module segments, such as sdk-go-scim/0.6.0, as the new SDK's marker.
func sdkGoUserAgent(ctx context.Context, sdkGoUserAgent string) string {
	segments := []string{useragent.FromContext(ctx)}
	for _, segment := range strings.Fields(sdkGoUserAgent) {
		if strings.HasPrefix(segment, "sdk-go-") {
			segments = append(segments, segment)
		}
	}
	return strings.Join(segments, " ")
}

// send sends r, whose context ends with timeout, and reports an attempt that timed out
// with the classic client's message unless parent ended first.
func (t *sdkGoTransport) send(parent context.Context, r *http.Request,
	timeout *attemptTimeout) (*http.Response, error) {
	if r.Body != nil && r.Body != http.NoBody {
		r.Body = &inactivityBody{ReadCloser: r.Body, timer: timeout.timer}
	}
	resp, err := t.base.RoundTrip(r)
	if err != nil {
		timeout.end()
		if timeout.timer.expired() && parent.Err() == nil {
			return nil, fmt.Errorf("request timed out after %s of inactivity",
				t.httpTimeout)
		}
		return nil, err
	}
	resp.Body = &inactivityBody{ReadCloser: resp.Body, timer: timeout.timer,
		cancel: timeout.cancel}
	return resp, nil
}

// attemptTimeout cancels one attempt after its timeout passes without progress. It
// starts before authentication and lasts until the response body is closed.
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

// inactivityBody extends the timer on every read. A response body also ends the attempt
// when closed.
type inactivityBody struct {
	io.ReadCloser
	timer  *inactivityTimer
	cancel context.CancelFunc
}

func (b *inactivityBody) Read(p []byte) (int, error) {
	n, err := b.ReadCloser.Read(p)
	b.timer.extend()
	return n, err
}

func (b *inactivityBody) Close() error {
	err := b.ReadCloser.Close()
	if b.cancel != nil {
		b.timer.stop()
		b.cancel()
	}
	return err
}
