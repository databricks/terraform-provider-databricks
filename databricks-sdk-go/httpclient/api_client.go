package httpclient

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/httpclient/traceparent"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/logger/httplog"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
)

type RequestVisitor func(*http.Request) error

type ClientConfig struct {
	AuthVisitor RequestVisitor
	Visitors    []RequestVisitor

	AccountID string
	Host      string

	RetryTimeout       time.Duration
	HTTPTimeout        time.Duration
	InsecureSkipVerify bool
	DebugHeaders       bool
	DebugTruncateBytes int
	RateLimitPerSecond int

	ErrorMapper     func(ctx context.Context, resp common.ResponseWrapper) error
	ErrorRetriable  func(ctx context.Context, err error) bool
	TransientErrors []string

	Transport http.RoundTripper
}

type DoOption struct {
	in           RequestVisitor
	out          func(body *common.ResponseWrapper) error
	body         any
	contentType  string
	isAuthOption bool
	queryParams  map[string]any
}

type ApiClient struct {
	config      ClientConfig
	rateLimiter *rate.Limiter
	httpClient  *http.Client
}

// IsAccountClient returns true if the client is configured for Accounts API.
func (apic *ApiClient) IsAccountClient() bool {
	if apic.config.AccountID == "" {
		return false
	}
	if strings.HasPrefix(apic.config.Host, "https://accounts.") {
		return true
	}
	if strings.HasPrefix(apic.config.Host, "https://accounts-dod.") {
		return true
	}
	return false
}

var defaultTransport = makeDefaultTransport()

func makeDefaultTransport() *http.Transport {
	return &http.Transport{
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
}

func (cfg ClientConfig) httpTransport() http.RoundTripper {
	if cfg.Transport != nil {
		return cfg.Transport
	}
	if cfg.InsecureSkipVerify {
		t := makeDefaultTransport()
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		return t
	}
	return defaultTransport
}

const (
	defaultHTTPTimeout         = 60 * time.Second
	defaultRetryTimeout        = 5 * time.Minute
	defaultDebugTruncateBytes  = 96
	defaultRateLimitPerSeconds = 15
)

func setDefaults(cfg *ClientConfig) {
	if cfg.HTTPTimeout == 0 {
		cfg.HTTPTimeout = defaultHTTPTimeout
	}
	if cfg.RetryTimeout == 0 {
		cfg.RetryTimeout = defaultRetryTimeout
	}
	if cfg.DebugTruncateBytes == 0 {
		cfg.DebugTruncateBytes = defaultDebugTruncateBytes
	}
	if cfg.ErrorMapper == nil {
		cfg.ErrorMapper = DefaultErrorMapper
	}
	if cfg.ErrorRetriable == nil {
		cfg.ErrorRetriable = DefaultErrorRetriable
	}
	if cfg.RateLimitPerSecond == 0 {
		cfg.RateLimitPerSecond = defaultRateLimitPerSeconds
	}
}

func NewApiClient(cfg ClientConfig) *ApiClient {
	setDefaults(&cfg)

	transport := cfg.httpTransport()
	rateLimit := rate.Limit(cfg.RateLimitPerSecond)

	// Depend on the HTTP fixture interface to prevent any coupling.
	//
	// TODO: Evaluate whether this can be implemented in a cleaner way.
	if skippable, ok := transport.(interface {
		SkipRetryOnIO() bool
	}); ok && skippable.SkipRetryOnIO() {
		rateLimit = rate.Inf
		cfg.RetryTimeout = 0
	}

	return &ApiClient{
		config:      cfg,
		rateLimiter: rate.NewLimiter(rateLimit, cfg.RateLimitPerSecond),
		httpClient: &http.Client{
			// We deal with request timeouts ourselves such that we do not
			// time out during request or response body reads that make
			// progress (e.g. on a slower network connection).
			Timeout:   0,
			Transport: transport,
		},
	}
}

// Do sends an HTTP request against path.
func (c *ApiClient) Do(ctx context.Context, method, path string, opts ...DoOption) error {
	var authVisitor RequestVisitor
	var explicitQueryParams map[string]any
	visitors := c.config.Visitors[:]
	for _, o := range opts {
		if o.queryParams != nil {
			if explicitQueryParams != nil {
				return fmt.Errorf("only one set of query params is allowed")
			}
			explicitQueryParams = o.queryParams
		}
		if o.in == nil {
			continue
		}
		if o.isAuthOption {
			if authVisitor != nil {
				return fmt.Errorf("only one auth visitor is allowed")
			}
			authVisitor = o.in
		} else {
			// merge client-wide and request-specific visitors
			visitors = append(visitors, o.in)
		}

	}
	// Use default AuthVisitor if none is provided
	if authVisitor == nil {
		authVisitor = c.config.AuthVisitor
	}
	if authVisitor != nil {
		// AuthVisitors must be the first visitors because other visitors depend on auth being configured
		visitors = append([]RequestVisitor{authVisitor}, visitors...)
	}

	var data any
	var contentType string
	for _, o := range opts {
		if o.body == nil {
			continue
		}
		data = o.body
		contentType = o.contentType
	}
	requestBody, err := makeRequestBody(method, &path, data, contentType, explicitQueryParams)
	if err != nil {
		return fmt.Errorf("request marshal: %w", err)
	}
	responseBody, err := c.perform(ctx, method, path, requestBody, visitors...)
	if err != nil {
		return err
	}
	for _, o := range opts {
		if o.out == nil {
			continue
		}
		err = o.out(responseBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ApiClient) isRetriable(ctx context.Context, err error) bool {
	if c.config.ErrorRetriable(ctx, err) {
		return true
	}
	if isRetriableUrlError(err) {
		// all IO errors are retriable
		logger.Debugf(ctx, "Attempting retry because of IO error: %s", err)
		return true
	}
	message := err.Error()
	// Handle transient errors for retries
	for _, substring := range c.config.TransientErrors {
		if strings.Contains(message, substring) {
			logger.Debugf(ctx, "Attempting retry because of %#v", substring)
			return true
		}
	}
	// some API's recommend retries on HTTP 500, but we'll add that later
	return false
}

// Common error-handling logic for all responses that may need to be retried.
//
// If the error is retriable, return a retries.Err to retry the request. However, as the request body will have been consumed
// by the first attempt, the body must be reset before retrying. If the body cannot be reset, return a retries.Err to halt.
//
// Always returns nil for the first parameter as there is no meaningful response body to return in the error case.
//
// If it is certain that an error should not be retried, use failRequest() instead.
func (c *ApiClient) handleError(ctx context.Context, err error, body common.RequestBody) (*common.ResponseWrapper, *retries.Err) {
	if !c.isRetriable(ctx, err) {
		return nil, retries.Halt(err)
	}
	if resetErr := body.Reset(); resetErr != nil {
		return nil, retries.Halt(errors.Join(err, resetErr)) // join to not lose the original error
	}
	return nil, retries.Continue(err)
}

// Fails the request with a retries.Err to halt future retries.
func (c *ApiClient) failRequest(msg string, err error) (*common.ResponseWrapper, *retries.Err) {
	err = fmt.Errorf("%s: %w", msg, err)
	return nil, retries.Halt(err)
}

func (c *ApiClient) attempt(
	ctx context.Context,
	method string,
	requestURL string,
	requestBody common.RequestBody,
	visitors ...RequestVisitor,
) func() (*common.ResponseWrapper, *retries.Err) {
	return func() (*common.ResponseWrapper, *retries.Err) {
		err := c.rateLimiter.Wait(ctx)
		if err != nil {
			return c.failRequest("failed in rate limiter", err)
		}

		pctx := ctx

		// This timeout context enables us to extend the request timeout
		// while the request or response body is being read.
		// It exists because the net/http package uses a fixed timeout regardless of payload size.
		ctx, ticker := newTimeoutContext(pctx, c.config.HTTPTimeout)
		request, err := http.NewRequestWithContext(ctx, method, requestURL, requestBody.Reader)
		if err != nil {
			return c.failRequest("failed creating new request", err)
		}
		for _, requestVisitor := range visitors {
			err = requestVisitor(request)
			if err != nil {
				return c.failRequest("failed during request visitor", err)
			}
		}
		// Set traceparent for distributed tracing.
		// This must be done after all visitors have run, as they may modify the request.
		traceparent.AddTraceparent(request)

		// request.Context() holds context potentially enhanced by visitors
		request.Header.Set("User-Agent", useragent.FromContext(request.Context()))
		if request.Header.Get("Content-Type") == "" && requestBody.ContentType != "" {
			request.Header.Set("Content-Type", requestBody.ContentType)
		}
		// If there is a request body, wrap it to extend the request timeout while it is being read.
		// Note: we do not wrap the request body earlier, because [http.NewRequestWithContext] performs
		// type probing on the body variable to determine the content length.
		if request.Body != nil && request.Body != http.NoBody {
			request.Body = newRequestBodyTicker(ticker, request.Body)
		}

		// attempt the actual request
		response, err := c.httpClient.Do(request)

		// After this point, the request body has (probably) been consumed. handleError() must be called to reset it if
		// possible.
		if uerr, ok := err.(*url.Error); ok {
			// If the timeout context has been canceled but the parent context hasn't, then the request has timed out.
			if pctx.Err() == nil && uerr.Err == context.Canceled {
				uerr.Err = fmt.Errorf("request timed out after %s of inactivity", c.config.HTTPTimeout)
			}
		}

		// If there is a response body, wrap it to extend the request timeout while it is being read.
		if response != nil && response.Body != nil {
			response.Body = newResponseBodyTicker(ticker, response.Body)
		} else {
			// If there is no response body or an error is returned, the request
			// has completed and there is no need to extend the timeout. Cancel
			// the context to clean up the underlying goroutine.
			ticker.Cancel()
		}

		// By this point, the request body has certainly been consumed.
		var responseWrapper common.ResponseWrapper
		if err == nil {
			responseWrapper, err = common.NewResponseWrapper(response, requestBody)
		}
		if err == nil {
			err = c.config.ErrorMapper(ctx, responseWrapper)
		}

		c.recordRequestLog(ctx, request, response, err, requestBody.DebugBytes, responseWrapper.DebugBytes)

		if err == nil {
			return &responseWrapper, nil
		}

		return c.handleError(ctx, err, requestBody)
	}
}

func (c *ApiClient) recordRequestLog(
	ctx context.Context,
	request *http.Request,
	response *http.Response,
	err error,
	requestBody, responseBody []byte,
) {
	// Don't compute expensive debug message if debug logging is not enabled.
	if !logger.Get(ctx).Enabled(ctx, logger.LevelDebug) {
		return
	}
	message := httplog.RoundTripStringer{
		Request:                  request,
		Response:                 response,
		Err:                      err,
		RequestBody:              requestBody,
		ResponseBody:             responseBody,
		DebugHeaders:             c.config.DebugHeaders,
		DebugTruncateBytes:       c.config.DebugTruncateBytes,
		DebugAuthorizationHeader: true,
	}.String()
	logger.Debugf(ctx, "%s", message)
}

// RoundTrip implements http.RoundTripper to integrate with golang.org/x/oauth2.
//
// Warning: This implementation violates the standard http.RoundTripper contract
// by retrying failed requests and interpreting HTTP status codes as errors.
// Specifically, retry logic should live at the application level rather than
// the transport level.
func (c *ApiClient) RoundTrip(request *http.Request) (*http.Response, error) {
	ctx := request.Context()
	requestURL := request.URL.String()
	resp, err := retries.Poll(ctx, c.config.RetryTimeout,
		c.attempt(ctx, request.Method, requestURL, common.RequestBody{
			Reader: request.Body,
			// DO NOT DECODE BODY, because it may contain sensitive payload,
			// like Azure Service Principal in a multipart/form-data body.
			DebugBytes: []byte("<http.RoundTripper>"),
		}, func(r *http.Request) error {
			r.Header = request.Header
			return nil
		}))
	if err != nil {
		return nil, err
	}
	// here we assume only successful responses, as HTTP 4XX and 5XX are mapped
	// to Go's error implementations.
	return resp.Response, nil
}

// InContextForOAuth2 returns a context with a custom *http.Client to be used
// for only for token acquisition through golang.org/x/oauth2 package
func (c *ApiClient) InContextForOAuth2(ctx context.Context) context.Context {
	return context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
		Timeout:   c.config.HTTPTimeout,
		Transport: c,
	})
}

func (c *ApiClient) perform(
	ctx context.Context,
	method,
	requestURL string,
	requestBody common.RequestBody,
	visitors ...RequestVisitor,
) (*common.ResponseWrapper, error) {
	resp, err := retries.Poll(ctx, c.config.RetryTimeout,
		c.attempt(ctx, method, requestURL, requestBody, visitors...))
	var timedOut *retries.ErrTimedOut
	if errors.As(err, &timedOut) {
		// TODO: check if we want to unwrap this error here
		return nil, timedOut.Unwrap()
	} else if err != nil {
		// Don't re-wrap, as upper layers may depend on handling apierr.APIError.
		return nil, err
	}
	return resp, nil
}
