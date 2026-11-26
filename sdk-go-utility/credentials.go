package sdkgoutility

import (
	"context"
	"fmt"
	"maps"
	"net/http"
	"slices"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/sdk-go/auth"
)

// credentialsUnknownName is the name of credentials whose config has no
// authentication type: it has not authenticated yet, or its credentials
// strategy has an empty name. A modular client puts the name in its
// User-Agent and fails every request when the name is empty.
const credentialsUnknownName = "unknown"

// Credentials returns the credentials of a modular SDK client that
// authenticates the way the classic client of cfg does. Pass them to the
// client with client.WithCredentials, so that the client doesn't resolve
// credentials of its own from the environment.
//
// AuthHeaders calls cfg.Authenticate on a request that carries the context of
// the caller, and returns every header that the call set, ordered by header
// name and, for one name, in the order the values were added. A failure is
// returned wrapped, behind the message the classic client gives to a failed
// authentication, "failed during request visitor: ". Name returns
// cfg.AuthType, the authentication type that the config resolved, and
// "unknown" while that is empty.
//
// Nothing is copied from cfg or cached. Resolving the config, choosing the
// credentials strategy and caching and refreshing tokens stay with cfg, so
// the modular client and the classic clients of the same config share one
// token.
//
// Credentials leaves to the caller:
//
//   - a cfg that is not nil. Resolve it with cfg.EnsureResolved before
//     several goroutines use the credentials, as databricks.NewWorkspaceClient
//     does for the classic client: the config doesn't synchronize its first
//     resolution.
//   - everything of the request other than its authentication headers: the
//     host, the workspace routing headers, the User-Agent and the headers of
//     cfg.Headers.
//   - cancellation. A strategy that ignores the context of the request, as
//     personal access tokens, basic authentication and OAuth M2M do, returns
//     its headers for a cancelled context too. The modular client then fails
//     on sending the request.
func Credentials(cfg *config.Config) auth.Credentials {
	return credentialsFromConfig{cfg: cfg}
}

// credentialsFromConfig is an auth.Credentials on top of a classic config.
type credentialsFromConfig struct {
	cfg *config.Config
}

func (c credentialsFromConfig) Name() string {
	if c.cfg.AuthType == "" {
		return credentialsUnknownName
	}
	return c.cfg.AuthType
}

// AuthHeaders authenticates a request for the path "/" and no host, which is
// what the classic client hands to cfg.Authenticate as well: it sets the host
// only after authentication. Reading cfg.Host here instead would race with
// the first authentication of the config, which rewrites the host under a
// lock that only the config can take.
func (c credentialsFromConfig) AuthHeaders(ctx context.Context) ([]auth.Header, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating new request: %w", err)
	}
	if err := c.cfg.Authenticate(req); err != nil {
		return nil, fmt.Errorf("failed during request visitor: %w", err)
	}
	headers := make([]auth.Header, 0, len(req.Header))
	for _, name := range slices.Sorted(maps.Keys(req.Header)) {
		for _, value := range req.Header[name] {
			headers = append(headers, auth.Header{Key: name, Value: value})
		}
	}
	return headers, nil
}
