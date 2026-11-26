package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// StaticHeaders returns a value for [Config.Headers] that sets the given HTTP
// headers on every request. Use it for headers that do not change from request
// to request; for headers that vary per request, set [Config.Headers] to a
// function directly.
func StaticHeaders(headers map[string]string) func(*http.Request) error {
	return func(r *http.Request) error {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
		return nil
	}
}

func HTTPClientConfigFromConfig(cfg *Config) (httpclient.ClientConfig, error) {
	if skippable, ok := cfg.HTTPTransport.(interface {
		SkipRetryOnIO() bool
	}); ok && skippable.SkipRetryOnIO() {
		cfg.Loaders = []Loader{noopLoader{}}
		cfg.Credentials = noopAuth{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return httpclient.ClientConfig{}, err
	}

	visitors := []httpclient.RequestVisitor{
		func(r *http.Request) error {
			if r.URL == nil {
				return fmt.Errorf("no URL found in request")
			}
			url, err := url.Parse(cfg.Host)
			if err != nil {
				return err
			}
			r.URL.Host = url.Host
			r.URL.Scheme = url.Scheme
			return nil
		},
		authInUserAgentVisitor(cfg),
		func(r *http.Request) error {
			// Detect if we are running in a CI/CD environment
			provider := useragent.CiCdProvider()
			if provider == "" {
				return nil
			}
			// Add the detected CI/CD provider to the user agent
			ctx := useragent.InContext(r.Context(), useragent.CicdKey, provider)
			*r = *r.WithContext(ctx) // replace request
			return nil
		},
		func(r *http.Request) error {
			// Detect if the SDK is being run in a Databricks Runtime.
			v := useragent.Runtime()
			if v == "" {
				return nil
			}
			// Add the detected Databricks Runtime version to the user agent
			ctx := useragent.InContext(r.Context(), useragent.RuntimeKey, v)
			*r = *r.WithContext(ctx) // replace request
			return nil
		},
		func(r *http.Request) error {
			// Detect if we are running inside a known AI coding agent.
			provider := useragent.AgentProvider()
			if provider == "" {
				return nil
			}
			// Add the detected agent to the user agent.
			ctx := useragent.InContext(r.Context(), useragent.AgentKey, provider)
			*r = *r.WithContext(ctx) // replace request
			return nil
		},
		func(r *http.Request) error {
			// Detect if we are running inside a known agent meta-harness.
			provider := useragent.MetaHarnessProvider()
			if provider == "" {
				return nil
			}
			// Add the detected meta-harness to the user agent.
			ctx := useragent.InContext(r.Context(), useragent.MetaHarnessKey, provider)
			*r = *r.WithContext(ctx) // replace request
			return nil
		},
	}

	// Apply the caller-provided header hook on every request. This runs after
	// AuthVisitor, so headers are added on top of normal authentication.
	if cfg.Headers != nil {
		visitors = append(visitors, cfg.Headers)
	}

	return httpclient.ClientConfig{
		AccountID:          cfg.AccountID,
		Host:               cfg.Host,
		RetryTimeout:       time.Duration(cfg.RetryTimeoutSeconds) * time.Second,
		HTTPTimeout:        time.Duration(cfg.HTTPTimeoutSeconds) * time.Second,
		RateLimitPerSecond: cfg.RateLimitPerSecond,
		DebugHeaders:       cfg.DebugHeaders,
		DebugTruncateBytes: cfg.DebugTruncateBytes,
		InsecureSkipVerify: cfg.InsecureSkipVerify,
		Transport:          cfg.HTTPTransport,
		AuthVisitor:        cfg.Authenticate,
		Visitors:           visitors,
		TransientErrors: []string{
			// This is temporary workaround for SCIM API returning 500.
			// TODO: Remove when it's fixed.
			"REQUEST_LIMIT_EXCEEDED",
		},
		ErrorMapper: apierr.GetAPIError,
		ErrorRetriable: func(ctx context.Context, err error) bool {
			var apiErr *apierr.APIError
			if errors.As(err, &apiErr) {
				return apiErr.IsRetriable(ctx)
			}
			return false
		},
	}, nil
}

// noopLoader skips configuration loading
type noopLoader struct{}

func (noopLoader) Name() string                { return "noop" }
func (noopLoader) Configure(cfg *Config) error { return nil }

// noopAuth skips authentication
type noopAuth struct{}

func (noopAuth) Name() string { return "noop" }
func (noopAuth) Configure(context.Context, *Config) (credentials.CredentialsProvider, error) {
	visitor := func(r *http.Request) error { return nil }
	return credentials.CredentialsProviderFn(visitor), nil
}

// Deprecated: use [HTTPClientConfigFromConfig] with [httpclient.NewApiClient].
func (c *Config) NewApiClient() (*httpclient.ApiClient, error) {
	cfg, err := HTTPClientConfigFromConfig(c)
	if err != nil {
		return nil, err
	}
	return httpclient.NewApiClient(cfg), nil
}
