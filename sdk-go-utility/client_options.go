package sdkgoutility

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/sdk-go/auth"
	"github.com/databricks/sdk-go/options/client"
)

// ClientOptions returns the options that point a modular SDK client, such as
// the one modelregistry.NewClient builds, at the workspace of cfg, the config
// of an existing classic workspace client. A client built from them:
//
//   - sends its calls to cfg.Host, with cfg.WorkspaceID in the
//     X-Databricks-Workspace-Id header when it is set, and carries
//     cfg.AccountID;
//   - signs its requests with credentials and sends them over transport;
//   - reads neither a config file nor the environment, so DATABRICKS_HOST,
//     DATABRICKS_TOKEN, DATABRICKS_CONFIG_PROFILE and ~/.databrickscfg change
//     nothing;
//   - sets no timeout of its own, so a call ends when ctx ends or its
//     retrier stops retrying, as a classic call does: the classic client
//     looks at its retry timeout only between attempts, and runs each attempt
//     and its wait for the rate limiter on the context of the caller. The
//     retrier of Retrier stops after cfg.RetryTimeoutSeconds when it is
//     positive, after 5 minutes when it is 0, or 20 minutes for a transport
//     whose SkipRetryOnIO is true, for which the classic client clears its
//     retry timeout, and never when it is negative, which the provider sets
//     unless the user sets retry_timeout_seconds. Without a retrier a call is
//     one attempt.
//
// ClientOptions returns an error when cfg, credentials or transport is nil,
// including a nil pointer in the interface, or when cfg.Host is not the
// scheme://host that EnsureResolved leaves there. It only reads cfg, and of
// cfg only Host, WorkspaceID and AccountID.
//
// It leaves to the caller:
//
//   - passing the config of a workspace client that already exists, since
//     building one with databricks.NewWorkspaceClient resolves and
//     authenticates its config. A config that is not resolved can lack the
//     scheme of its host and hold the workspace ID in its query, and
//     ClientOptions tells it only by such a host. A config that has not
//     authenticated rewrites its host on its first call, under a lock that
//     ClientOptions can't take, so reading the host then is a data race;
//   - credentials and transport. ClientOptions doesn't derive them from cfg
//     and doesn't wrap them, so authentication, the User-Agent, TLS settings,
//     request logging and cfg.HTTPTimeoutSeconds are theirs to bring;
//   - what a single call needs: a retrier, such as the one Retrier returns,
//     without which a modular client doesn't retry, a rate limiter and, if
//     one is wanted, a timeout of its own;
//   - building the client, with further options such as a logger.
func ClientOptions(cfg *config.Config, credentials auth.Credentials, transport http.RoundTripper) ([]client.Option, error) {
	switch {
	case cfg == nil:
		return nil, errors.New("modular SDK client options: config is nil")
	case clientOptionsIsNil(credentials):
		return nil, errors.New("modular SDK client options: credentials are nil")
	case clientOptionsIsNil(transport):
		return nil, errors.New("modular SDK client options: transport is nil")
	}
	if err := clientOptionsCheckResolvedHost(cfg.Host); err != nil {
		return nil, err
	}
	return []client.Option{
		client.WithHost(cfg.Host),
		client.WithWorkspaceID(cfg.WorkspaceID),
		client.WithAccountID(cfg.AccountID),
		client.WithCredentials(credentials),
		client.WithHTTPClient(&http.Client{Transport: transport}),
		client.WithoutConfigFile(),
		client.WithoutEnv(),
		// Zero turns off the 5 minutes that a modular client otherwise gives
		// every call, which would cut an attempt in flight.
		client.WithTimeout(0),
	}, nil
}

// clientOptionsIsNil reports whether value is nil or a nil pointer, map, slice,
// func or channel in an interface. A comparison with nil misses the latter, and
// the modular client would then panic on its first call instead of failing
// here.
func clientOptionsIsNil(value any) bool {
	if value == nil {
		return true
	}
	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
		return v.IsNil()
	}
	return false
}

// clientOptionsCheckResolvedHost returns an error unless host is scheme://host
// with an optional port, the form in which EnsureResolved stores the host of a
// classic config. The modular client uses the host as it is: without a scheme
// every call fails, and it drops a path and a query, where an unresolved
// classic config can still carry the workspace ID.
func clientOptionsCheckResolvedHost(host string) error {
	if host == "" {
		return errors.New("modular SDK client options: config has no host")
	}
	parsed, err := url.Parse(host)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" ||
		(&url.URL{Scheme: parsed.Scheme, Host: parsed.Host}).String() != host {
		return fmt.Errorf("modular SDK client options: config host %q is not of the form scheme://host, resolve the config first", host)
	}
	return nil
}
