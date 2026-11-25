package config

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/config/credentials"
)

type BasicCredentials struct {
}

func (c BasicCredentials) Name() string {
	return "basic"
}

func (c BasicCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if err := unsupportedGroupRoleAssumption(cfg, c.Name()); err != nil {
		return nil, err
	}
	if cfg.Username == "" || cfg.Password == "" || cfg.Host == "" {
		return nil, nil
	}
	tokenUnB64 := fmt.Sprintf("%s:%s", cfg.Username, cfg.Password)
	b64 := base64.StdEncoding.EncodeToString([]byte(tokenUnB64))
	visitor := func(r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("Basic %s", b64))
		return nil
	}
	return credentials.CredentialsProviderFn(visitor), nil
}
