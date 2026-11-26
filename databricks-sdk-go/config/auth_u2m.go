package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
)

type u2mCredentials struct{}

func (u u2mCredentials) Name() string {
	return "databricks-cli"
}

func (u u2mCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if err := unsupportedGroupRoleAssumption(cfg, u.Name()); err != nil {
		return nil, err
	}
	if cfg.Host == "" {
		return nil, fmt.Errorf("host is required")
	}

	// We only partially support custom scopes with databricks-cli auth.
	// Users can specify their scopes when logging in with `databricks auth login`,
	// but not when using `databricks-cli` auth in the SDKs.
	// The token store is keyed by profile name (when an explicit profile is set)
	// or by host (legacy/implicit default). In either case, scopes are not part
	// of the cache key, so custom scopes would be silently ignored.
	// TODO: remove this validation once the token store can identify scopes based on their permissions. This will
	// allow users to specify scopes explicitly in the SDKs.
	if err := validateCliScopes(cfg); err != nil {
		return nil, err
	}

	ts, err := NewCliTokenSource(ctx, cfg)
	if err != nil {
		return nil, err
	}
	_, err = ts.Token(ctx)
	if err != nil {
		return nil, err
	}
	return credentials.NewOAuthCredentialsProviderFromTokenSource(auth.NewCachedTokenSource(ts, cacheOptions(cfg)...)), nil
}

// ErrCustomScopesNotSupported is returned when custom scopes are specified
// with databricks-cli auth, which is not supported because the CLI's token
// cache is keyed by host, not by scopes.
const ErrCustomScopesNotSupported = "custom scopes are not supported with databricks-cli auth; " +
	"scopes are determined by what was last used when logging in with `databricks auth login`"

// validateCliScopes returns an error if custom scopes are
// specified with databricks-cli auth. The CLI's token cache is keyed by host,
// not by scopes, so custom scopes would be silently ignored otherwise. Custom scopes
// from config files are allowed since `databricks auth login` writes them there.
func validateCliScopes(cfg *Config) error {
	if len(cfg.Scopes) == 0 {
		return nil
	}
	for _, attr := range ConfigAttributes {
		if attr.Name != "scopes" {
			continue
		}
		if cfg.getSource(&attr).Type == SourceFile {
			return nil
		}
		return errors.New(ErrCustomScopesNotSupported)
	}
	return nil
}

var DatabricksCliCredentials = u2mCredentials{}
