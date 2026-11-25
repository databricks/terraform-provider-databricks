package sdkgoutility

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/sdk-go/auth"
	"github.com/databricks/sdk-go/options/client"
)

// NewClient builds a modular SDK client that behaves like the classic client
// of cfg. cfg is the config of an existing classic client, such as w.Config
// of the client that WorkspaceClientUnifiedProvider returns, and construct is
// the constructor of the modular client:
//
//	registry, err := sdkgoutility.NewClient(ctx, w.Config, modelregistry.NewClient)
//
// It puts the helpers of this package together: the client authenticates
// with Credentials(cfg), sends its requests over Transport(cfg) inside
// AttemptTimeout inside DebugLog, and is built from the ClientOptions of
// these. So the client calls the workspace of cfg with the headers of the
// classic client, ends a stalled attempt with the error of the classic
// client, writes each attempt to the debug log as the classic client does,
// and reads neither the environment nor a config file.
//
// DebugLog belongs inside Transport, which sets the headers, and Transport
// takes no transport to send over. So DebugLog logs the request before
// Transport has changed it and the outcome that Transport returns:
//
//   - with debug_headers the log shows the Host and the User-Agent of the
//     modular client and neither Accept, Traceparent nor a header of the
//     cfg.Headers hook, which is not what is sent. Every other line of a
//     call that the workspace answers is the one of the classic client;
//   - a request that a request visitor refuses is logged, where the classic
//     client logs no attempt;
//   - the private link login page is logged as the error response that
//     Transport returns in its place.
//
// NewClient returns the zero value of C and an error when cfg or construct
// is nil, and when a step fails. The error names the step, which is the
// helper or "constructor", and wraps what the step returned. construct is
// not called after a step has failed.
//
// NewClient keeps no state, and every call builds a new client. For a cfg
// that is resolved, which the config of a classic client is, a call sends no
// request and starts no goroutine: it allocates the transports, the options
// and the client, about a microsecond of work. So a resource can build a
// client for every operation. The clients of one cfg share what is not
// theirs: the token, which only cfg caches and refreshes, and the
// connections, which are those of cfg.HTTPTransport when it is set and
// otherwise a pool that Transport keeps for the modular clients of every
// config: one for the configs that verify TLS certificates and one for
// those with cfg.InsecureSkipVerify.
//
// NewClient leaves to the caller:
//
//   - passing the config of a classic client that has authenticated, as
//     ClientOptions asks. The first call of a modular client authenticates a
//     config that has not, but that authentication writes cfg.Host, which
//     NewClient reads, so the two must not run at the same time;
//   - what a single call needs and no client option can set: a retrier,
//     without which the client doesn't retry, and a rate limiter;
//   - the error of a failed call, which is the one of the modular SDK and
//     not the apierr.APIError of the classic client.
func NewClient[C any](ctx context.Context, cfg *config.Config, construct func(context.Context, ...client.Option) (C, error)) (C, error) {
	return newClientBuild(ctx, cfg, construct, newClientHelpers)
}

// newClientSteps are the helpers that NewClient puts together. They are
// fields so that a test can make each of them fail: of the real helpers,
// AttemptTimeout and DebugLog can't fail once Transport has resolved the
// config.
type newClientSteps struct {
	credentials    func(*config.Config) auth.Credentials
	transport      func(*config.Config) (http.RoundTripper, error)
	attemptTimeout func(*config.Config, http.RoundTripper) (http.RoundTripper, error)
	debugLog       func(*config.Config, http.RoundTripper) (http.RoundTripper, error)
	clientOptions  func(*config.Config, auth.Credentials, http.RoundTripper) ([]client.Option, error)
}

// newClientHelpers are the steps of NewClient.
var newClientHelpers = newClientSteps{
	credentials:    Credentials,
	transport:      func(cfg *config.Config) (http.RoundTripper, error) { return Transport(cfg, nil) },
	attemptTimeout: AttemptTimeout,
	debugLog:       DebugLog,
	clientOptions:  ClientOptions,
}

// newClientBuild is NewClient with the given steps. Their order is fixed by
// what each needs: Transport resolves cfg, which the other helpers read,
// AttemptTimeout goes around Transport, so that the wait for progress covers
// the whole attempt, and DebugLog goes around AttemptTimeout, so that the log
// shows the error that ended an attempt.
func newClientBuild[C any](ctx context.Context, cfg *config.Config, construct func(context.Context, ...client.Option) (C, error), steps newClientSteps) (C, error) {
	var zero C
	switch {
	case cfg == nil:
		return zero, errors.New("modular SDK client: config is nil")
	case construct == nil:
		return zero, errors.New("modular SDK client: constructor is nil")
	}
	transport, err := steps.transport(cfg)
	if err != nil {
		return zero, fmt.Errorf("modular SDK client: Transport: %w", err)
	}
	transport, err = steps.attemptTimeout(cfg, transport)
	if err != nil {
		return zero, fmt.Errorf("modular SDK client: AttemptTimeout: %w", err)
	}
	transport, err = steps.debugLog(cfg, transport)
	if err != nil {
		return zero, fmt.Errorf("modular SDK client: DebugLog: %w", err)
	}
	options, err := steps.clientOptions(cfg, steps.credentials(cfg), transport)
	if err != nil {
		return zero, fmt.Errorf("modular SDK client: ClientOptions: %w", err)
	}
	built, err := construct(ctx, options...)
	if err != nil {
		return zero, fmt.Errorf("modular SDK client: constructor: %w", err)
	}
	return built, nil
}
