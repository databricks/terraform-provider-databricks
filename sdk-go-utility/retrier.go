package sdkgoutility

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/sdk-go/core/ops"
)

// Retrier returns a retrier that makes one modular SDK call retry by the rules
// of the classic client, github.com/databricks/databricks-sdk-go, instead of
// those of the modular SDK. The rules are the ones that the classic
// httpclient.ApiClient applies to a request of a client built from cfg:
//
//   - What is retried. The retrier passes the error of an attempt to classic,
//     which returns the error that the classic client has for the same
//     failure, and retries when the classic client retries that error: when
//     ErrorRetriable of config.HTTPClientConfigFromConfig(cfg) says so (a 429,
//     a 503 and the messages that the classic apierr package calls
//     transient), when the error is a *url.Error with one of the network
//     messages in retrierNetworkErrors, or when its message contains one of
//     the TransientErrors of that configuration, such as
//     REQUEST_LIMIT_EXCEEDED. It doesn't retry an error that classic turns
//     into nil. It applies these rules to every error of an attempt, also to
//     one from before the request is sent, such as a failed authentication,
//     which the classic client never retries.
//   - How long it waits. The wait before retry n is the classic backoff,
//     retries.RetryConfig.Backoff(n). A Retry-After header doesn't change the
//     wait, because the classic client doesn't read that header.
//   - When it stops. The classic retry timeout is cfg.RetryTimeoutSeconds, or
//     5 minutes when that is 0, and it runs from the moment Retrier returns.
//     A negative cfg.RetryTimeoutSeconds, which the provider sets when the
//     user sets none, is no timeout. The retrier stops when the timeout or
//     the deadline of ctx would pass during the next wait, and when ctx has
//     ended. The call then returns the error of its last attempt, as a
//     classic call does. Only a ctx that is cancelled during a wait ends the
//     call differently: the modular SDK returns the error of ctx then, and
//     the classic client the error of the last attempt.
//
// The retrier writes the lines that the classic client logs about a retry, to
// the logger of the classic SDK. Retrier returns an error when cfg or classic
// is nil, and the error of config.HTTPClientConfigFromConfig when cfg doesn't
// resolve.
//
// Retrier leaves to the caller:
//
//   - Converting errors. classic is the only thing that knows how a modular
//     SDK error looks to the classic client, and it must return an error that
//     isn't an API error as it is.
//   - Building one retrier for each call, right before the call, and passing
//     it with call.WithRetrier(func() ops.Retrier { return retrier }). A
//     retrier counts the attempts of its call and isn't safe for concurrent
//     use. An iterator such as ListRegisteredModelsIter passes its options to
//     the call of every page, so its pages would share one count and one
//     timeout.
//   - The timeout of the call. A modular client ends a call after 5 minutes
//     unless call.WithTimeout or client.WithTimeout says otherwise, whatever
//     the retrier answers.
//   - Rate limiting, which a call takes through call.WithLimiter.
func Retrier(ctx context.Context, cfg *config.Config, classic func(error) error) (ops.Retrier, error) {
	if cfg == nil {
		return nil, errors.New("sdkgoutility.Retrier: cfg is nil")
	}
	if classic == nil {
		return nil, errors.New("sdkgoutility.Retrier: classic is nil")
	}
	httpConfig, err := config.HTTPClientConfigFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	r := &retrierForCall{
		ctx:             ctx,
		classic:         classic,
		errorRetriable:  httpConfig.ErrorRetriable,
		transientErrors: httpConfig.TransientErrors,
		retryConfig:     retrierClassicRetryConfig(httpConfig),
	}
	if timeout := r.retryConfig.Timeout(); timeout > 0 {
		r.deadline = time.Now().Add(timeout)
	}
	if deadline, ok := ctx.Deadline(); ok && (r.deadline.IsZero() || deadline.Before(r.deadline)) {
		r.deadline = deadline
	}
	return r, nil
}

// retrierClassicRetryConfig returns the retry configuration that the classic
// client gives a request, by the steps that the client takes from its HTTP
// client configuration to retries.Poll.
func retrierClassicRetryConfig(httpConfig httpclient.ClientConfig) retries.RetryConfig {
	timeout := httpConfig.RetryTimeout
	if timeout == 0 {
		// defaultRetryTimeout of the classic httpclient package.
		timeout = 5 * time.Minute
	}
	// httpclient.NewApiClient clears the timeout for the fixture transports of
	// the classic SDK, and the retries package then applies its own default.
	if fixture, ok := httpConfig.Transport.(interface{ SkipRetryOnIO() bool }); ok && fixture.SkipRetryOnIO() {
		timeout = 0
	}
	var retryConfig retries.RetryConfig
	retries.WithTimeout(timeout)(&retryConfig)
	return retryConfig
}

// retrierForCall is the ops.Retrier that Retrier returns.
type retrierForCall struct {
	// ctx is the context of the call. The retrier lives as long as that call.
	ctx     context.Context
	classic func(error) error
	// errorRetriable and transientErrors are ErrorRetriable and
	// TransientErrors of the classic HTTP client configuration.
	errorRetriable  func(context.Context, error) bool
	transientErrors []string
	retryConfig     retries.RetryConfig
	// deadline is when the retries stop. It is zero when only ctx stops them.
	deadline time.Time
	// attempt is the number of attempts that have failed with a retriable
	// error.
	attempt int
}

// IsRetriable does for a failed attempt what Run of the classic retries
// package does: it picks the wait by the number of the attempt, and gives up
// when the retries would end during that wait.
func (r *retrierForCall) IsRetriable(err error) (time.Duration, bool) {
	err = r.classic(err)
	if err == nil {
		return 0, false
	}
	if !r.isRetriable(err) {
		logger.Debugf(r.ctx, "non-retriable error: %s", err)
		return 0, false
	}
	r.attempt++
	wait := r.retryConfig.Backoff(r.attempt)
	logger.Tracef(r.ctx, "%s. Sleeping %s",
		strings.TrimSuffix(err.Error(), "."),
		wait.Round(time.Millisecond))
	if r.ctx.Err() != nil || !r.deadline.IsZero() && !time.Now().Add(wait).Before(r.deadline) {
		return 0, false
	}
	return wait, true
}

// isRetriable is isRetriable of the classic httpclient.ApiClient, with the
// debug messages that it logs.
func (r *retrierForCall) isRetriable(err error) bool {
	if r.errorRetriable(r.ctx, err) {
		return true
	}
	if retrierIsNetworkError(err) {
		logger.Debugf(r.ctx, "Attempting retry because of IO error: %s", err)
		return true
	}
	message := err.Error()
	for _, substring := range r.transientErrors {
		if strings.Contains(message, substring) {
			logger.Debugf(r.ctx, "Attempting retry because of %#v", substring)
			return true
		}
	}
	return false
}

// retrierNetworkErrors are the messages of the network errors that the
// classic client retries. They are a copy of urlErrorTransientErrorMessages of
// the classic httpclient package, which doesn't export the list.
var retrierNetworkErrors = []string{
	"connection reset by peer",
	"TLS handshake timeout",
	"connection refused",
	"Unexpected error",
	"i/o timeout",
}

// retrierIsNetworkError is isRetriableUrlError of the classic httpclient
// package.
func retrierIsNetworkError(err error) bool {
	var urlError *url.Error
	if !errors.As(err, &urlError) {
		return false
	}
	message := err.Error()
	for _, substring := range retrierNetworkErrors {
		if strings.Contains(message, substring) {
			return true
		}
	}
	return false
}
