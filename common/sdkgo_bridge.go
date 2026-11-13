package common

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

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
	classicDefaultRateLimit    = 15
)

// sdkGoBridge holds the classic client settings that sdk-go clients reuse.
type sdkGoBridge struct {
	httpConfig httpclient.ClientConfig
	retry      retries.RetryConfig
	limiter    *rate.Limiter
	httpClient *http.Client
}

// newSdkGoBridge reads the settings the classic client derives from cfg, so both clients
// follow retry_timeout_seconds, rate_limit, http_timeout_seconds and the TLS and proxy
// settings.
func newSdkGoBridge(cfg *config.Config) (*sdkGoBridge, error) {
	httpConfig, err := config.HTTPClientConfigFromConfig(cfg)
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
	rateLimit := httpConfig.RateLimitPerSecond
	if rateLimit == 0 {
		rateLimit = classicDefaultRateLimit
	}
	limit := rate.Limit(rateLimit)
	base := classicTransport(httpConfig)
	// httpclient.NewApiClient makes the same exception for test fixture transports.
	skippable, ok := base.(interface{ SkipRetryOnIO() bool })
	if ok && skippable.SkipRetryOnIO() {
		limit = rate.Inf
		retryTimeout = 0
	}
	b := &sdkGoBridge{
		httpConfig: httpConfig,
		limiter:    rate.NewLimiter(limit, rateLimit),
		httpClient: &http.Client{Transport: &sdkGoTransport{
			base:        base,
			visitors:    httpConfig.Visitors,
			httpTimeout: httpTimeout,
		}},
	}
	retries.WithTimeout(retryTimeout)(&b.retry)
	return b, nil
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
// refresh and extra authentication headers keep working.
type classicCredentials struct {
	cfg *config.Config
}

// Name is only used by sdk-go's own User-Agent, which sdkGoTransport replaces.
func (classicCredentials) Name() string {
	return "classic-config"
}

func (c classicCredentials) AuthHeaders(ctx context.Context) ([]auth.Header, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.cfg.Host, nil)
	if err != nil {
		return nil, err
	}
	if err := c.cfg.Authenticate(req); err != nil {
		return nil, fmt.Errorf("failed during request visitor: %w", err)
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
	r := req.Clone(req.Context())
	for _, visit := range t.visitors {
		if err := visit(r); err != nil {
			return nil, fmt.Errorf("failed during request visitor: %w", err)
		}
	}
	traceparent.AddTraceparent(r)
	r.Header.Set("User-Agent", sdkGoUserAgent(r.Context(), req.Header.Get("User-Agent")))
	if attempt, ok := r.Context().Value(sdkGoAttemptKey{}).(*sdkGoAttempt); ok {
		attempt.request = r
	}
	return t.roundTripWithTimeout(r)
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

// roundTripWithTimeout cancels an attempt after httpTimeout without progress, like the
// classic client, and reports it with the classic client's message.
func (t *sdkGoTransport) roundTripWithTimeout(r *http.Request) (*http.Response, error) {
	ctx, cancel := context.WithCancel(r.Context())
	timer := newInactivityTimer(t.httpTimeout, cancel)
	attempt := r.WithContext(ctx)
	if attempt.Body != nil && attempt.Body != http.NoBody {
		attempt.Body = &inactivityBody{ReadCloser: attempt.Body, timer: timer}
	}
	resp, err := t.base.RoundTrip(attempt)
	if err != nil {
		timer.stop()
		cancel()
		if timer.expired() && r.Context().Err() == nil {
			return nil, fmt.Errorf("request timed out after %s of inactivity",
				t.httpTimeout)
		}
		return nil, err
	}
	resp.Body = &inactivityBody{ReadCloser: resp.Body, timer: timer, cancel: cancel}
	return resp, nil
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
