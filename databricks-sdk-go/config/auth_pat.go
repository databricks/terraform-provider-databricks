package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	authcred "github.com/databricks/databricks-sdk-go/config/experimental/auth/credentials"
)

type PatCredentials struct{}

func (c PatCredentials) Name() string {
	return "pat"
}

func (c PatCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if err := unsupportedGroupRoleAssumption(cfg, c.Name()); err != nil {
		return nil, err
	}
	// Host inference is not supported for PAT authentication. This is
	// arguably a redundant check as requests will fail anyway if no
	// host is provided. This check exists for backward compatibility
	// with the previous PAT credentials implementation.
	if cfg.Host == "" {
		return nil, fmt.Errorf("host is required for PAT authentication")
	}
	creds, err := authcred.NewPATCredentials(cfg.Token)
	if err != nil {
		return nil, err
	}
	return credentials.FromCredentials(creds), nil
}
