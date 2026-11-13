package common

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/retries"
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
		httpClient: &http.Client{
			Transport: newSdkGoTransport(httpConfig, base, httpTimeout),
			// sdkGoTransport follows redirects itself, like the classic client.
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
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
