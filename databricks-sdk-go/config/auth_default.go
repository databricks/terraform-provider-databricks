package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
)

const authDocURL = "https://docs.databricks.com/en/dev-tools/auth.html#databricks-client-unified-authentication"

// ErrCannotConfigureDefault indicates that no credentials strategy in the
// chain was able to provide valid credentials.
var ErrCannotConfigureDefault = fmt.Errorf("cannot configure default credentials, please check %s to configure credentials for your preferred authentication method", authDocURL)

// NewCredentialsChain returns a new CredentialsStrategy that tries the given
// strategies in order and returns the first one that succeeds. If
// [Config.AuthType] is set, only the strategy with a matching name is tried.
//
// INTERNAL: This function is not part of the public API and is subject to
// change. Users are encouraged to use an explicit credentials strategy rather
// than relying on a custom credentials chain.
func NewCredentialsChain(strategies ...CredentialsStrategy) CredentialsStrategy {
	return &credentialsChain{strategies: strategies}
}

type credentialsChain struct {
	strategies []CredentialsStrategy
	name       string
}

func (c *credentialsChain) Name() string {
	return c.name
}

func (c *credentialsChain) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	// If an auth type is specified, try to configure the credentials for that
	// specific auth type. If an error is encountered, return it.
	if cfg.AuthType != "" {
		for _, s := range c.strategies {
			if s.Name() == cfg.AuthType {
				logger.Tracef(ctx, "Attempting to configure auth: %q", s.Name())
				c.name = s.Name()
				return s.Configure(ctx, cfg)
			}
		}
		return nil, fmt.Errorf("auth type %q not found, please check %s for a list of supported auth types", cfg.AuthType, authDocURL)
	}

	// If no auth type is specified, try the strategies in order. If a strategy
	// succeeds, returns the credentials provider. If a strategy fails, swallow
	// the error and try the next strategy.
	for _, s := range c.strategies {
		logger.Tracef(ctx, "Attempting to configure auth: %q", s.Name())
		cp, err := s.Configure(ctx, cfg)
		if err != nil || cp == nil {
			logger.Debugf(ctx, "Failed to configure auth %q: %v", s.Name(), err)
			continue
		}
		// Many credentials providers can only be truly validated after a
		// request is made (e.g. because they need to exercise some hand-shake
		// or verify that tokens exist in the cache). We perform a dry run to
		// validate the credentials provider.
		if err := dryRun(cp); err != nil {
			logger.Debugf(ctx, "Failed to configure auth %q: %v", s.Name(), err)
			continue
		}
		c.name = s.Name()
		return cp, nil // success
	}

	return nil, ErrCannotConfigureDefault
}

func dryRun(cp credentials.CredentialsProvider) error {
	req := &http.Request{Header: make(http.Header)} // dummy
	err := cp.SetHeaders(req)
	return err
}

// DefaultCredentials implements the default credentials strategy. It tries
// each known authentication method in priority order and selects the first
// one that succeeds.
type DefaultCredentials struct {
	name string
}

func (c *DefaultCredentials) Name() string {
	if c.name == "" {
		return "default"
	}
	return c.name
}

func (c *DefaultCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	chain := NewCredentialsChain(
		PatCredentials{},
		BasicCredentials{},
		M2mCredentials{},
		u2mCredentials{},
		MetadataServiceCredentials{},
		// OIDC Strategies.
		GitHubOIDCCredentials{},
		AzureDevOpsOIDCCredentials{},
		EnvOIDCCredentials{},
		FileOIDCCredentials{},
		// Azure strategies.
		AzureGithubOIDCCredentials{},
		AzureMsiCredentials{},
		AzureClientSecretCredentials{},
		AzureCliCredentials{},
		// Google strategies.
		GoogleCredentials{},
		GoogleDefaultCredentials{},
	)
	cp, err := chain.Configure(ctx, cfg)
	if err != nil {
		return nil, err
	}

	c.name = chain.Name()
	if c.name == "" {
		c.name = "default"
	}

	return cp, nil
}
