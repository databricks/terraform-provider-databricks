package sdkgoutility

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/sdk-go/auth"
	"github.com/databricks/sdk-go/options/client"
)

// clientOptionsClassicRetryTimeout is how long the classic client keeps
// retrying a call when Config.RetryTimeoutSeconds is zero.
const clientOptionsClassicRetryTimeout = 5 * time.Minute

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
//   - bounds a call, retries included, by cfg.RetryTimeoutSeconds as the
//     classic client does: 5 minutes when it is zero and no bound when it is
//     negative, which is what the provider sets unless the user chooses one.
//
// ClientOptions returns an error when cfg, credentials or transport is nil,
// including a nil pointer in the interface, or when cfg.Host is not the
// scheme://host that EnsureResolved leaves there. It only reads cfg, and of
// cfg only Host, WorkspaceID, AccountID and RetryTimeoutSeconds.
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
//   - what a single call needs: a retrier, without which a modular client
//     doesn't retry, a rate limiter and a timeout of its own;
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
		client.WithTimeout(clientOptionsCallTimeout(cfg.RetryTimeoutSeconds)),
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

// clientOptionsCallTimeout turns Config.RetryTimeoutSeconds into the value for
// client.WithTimeout. The two read zero differently: the classic client takes
// it as unset and retries for 5 minutes, and client.WithTimeout takes it as no
// timeout. The classic client sets no deadline for a negative value, while the
// modular client would start every call with a deadline that has passed.
func clientOptionsCallTimeout(retryTimeoutSeconds int) time.Duration {
	switch {
	case retryTimeoutSeconds == 0:
		return clientOptionsClassicRetryTimeout
	case retryTimeoutSeconds < 0:
		return 0
	}
	return time.Duration(retryTimeoutSeconds) * time.Second
}
