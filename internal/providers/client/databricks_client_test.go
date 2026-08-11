package client

import (
	"context"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// noopLoader is a no-op config.Loader installed in tests to suppress the SDK's
// default loader chain (env-var attributes + ~/.databrickscfg). The SDK injects
// its defaults only when Loaders is empty (see config.go EnsureResolved), so
// providing this single no-op keeps EnsureResolved from reading ambient state
// — protecting the test from a developer's DATABRICKS_* env vars or a
// ~/.databrickscfg that would otherwise trigger
// "more than one authorization method configured" when combined with our
// explicit Token, and from any other unrelated fields the loaders might fill.
// Mirrors the unexported noopLoader the SDK already defines for its own
// internal API client.
type noopLoader struct{}

func (noopLoader) Name() string                   { return "noop" }
func (noopLoader) Configure(*config.Config) error { return nil }

// TestPrepareDatabricksClient_NormalizesNoneWorkspaceID verifies that the
// "none" sentinel that the Databricks CLI writes to ~/.databrickscfg for
// account-level profiles is normalized to an empty string at provider configure
// time. Without this normalization, downstream parseWorkspaceID call sites fail
// with a strconv.ParseInt error.
//
// The test is hermetic against ambient developer state: Loaders is set to a
// single no-op to suppress env-var and databrickscfg reads, HostMetadataResolver
// is stubbed so EnsureResolved does not fetch /.well-known/databricks-config
// or mutate config fields from discovery, and the host is a fake .invalid TLD.
func TestPrepareDatabricksClient_NormalizesNoneWorkspaceID(t *testing.T) {
	tests := []struct {
		name                string
		workspaceID         string
		expectedWorkspaceID string
	}{
		{
			name:                "none sentinel is normalized to empty",
			workspaceID:         "none",
			expectedWorkspaceID: "",
		},
		{
			name:                "valid numeric workspace_id is preserved",
			workspaceID:         "1234567890",
			expectedWorkspaceID: "1234567890",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &config.Config{
				Host:        "https://test.invalid",
				Token:       "test-token",
				WorkspaceID: tc.workspaceID,
				Loaders:     []config.Loader{noopLoader{}},
				HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) {
					return nil, nil
				},
			}
			pc, err := PrepareDatabricksClient(context.Background(), cfg, nil)
			require.NoError(t, err)
			require.NotNil(t, pc)
			assert.Equal(t, tc.expectedWorkspaceID, pc.Config.WorkspaceID)
		})
	}
}

// workspaceHostMetadataResolver returns a resolver that reports a workspace host
// advertising the given workspace_id via /.well-known/databricks-config.
func workspaceHostMetadataResolver(workspaceID string) config.HostMetadataResolver {
	return func(context.Context, string) (*config.HostMetadata, error) {
		return &config.HostMetadata{
			WorkspaceID: workspaceID,
			HostType:    config.WorkspaceHost,
		}, nil
	}
}

// prepareForWorkspaceHost builds a hermetic client for a workspace host with the
// given user-supplied workspace_id and a metadata resolver advertising
// hostWorkspaceID. Ambient loaders are suppressed so the test is deterministic.
func prepareForWorkspaceHost(t *testing.T, userWorkspaceID, hostWorkspaceID string) (*common.DatabricksClient, error) {
	t.Helper()
	cfg := &config.Config{
		Host:                 "https://test.invalid",
		Token:                "test-token",
		WorkspaceID:          userWorkspaceID,
		Loaders:              []config.Loader{noopLoader{}},
		HostMetadataResolver: workspaceHostMetadataResolver(hostWorkspaceID),
	}
	return PrepareDatabricksClient(context.Background(), cfg, nil)
}

// TestPrepareDatabricksClient_WorkspaceIDFromHostMetadata verifies that for
// workspace hosts the workspace_id is resolved/validated from the host's
// /.well-known/databricks-config metadata and the cache is seeded so downstream
// CurrentWorkspaceID is a hit that never calls SCIM /Me. A cache hit is provable
// here because these clients have no usable workspace client, so any /Me attempt
// would error.
func TestPrepareDatabricksClient_WorkspaceIDFromHostMetadata(t *testing.T) {
	t.Run("user empty is back-filled and seeded", func(t *testing.T) {
		pc, err := prepareForWorkspaceHost(t, "", "12345")
		require.NoError(t, err)
		assert.Equal(t, "12345", pc.Config.WorkspaceID)
		id, err := pc.CurrentWorkspaceID(context.Background())
		require.NoError(t, err)
		assert.Equal(t, int64(12345), id)
	})

	t.Run("user matches host is seeded", func(t *testing.T) {
		pc, err := prepareForWorkspaceHost(t, "12345", "12345")
		require.NoError(t, err)
		id, err := pc.CurrentWorkspaceID(context.Background())
		require.NoError(t, err)
		assert.Equal(t, int64(12345), id)
	})

	t.Run("numeric mismatch fails fast at configure", func(t *testing.T) {
		_, err := prepareForWorkspaceHost(t, "99999", "12345")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "workspace_id mismatch")
	})

	t.Run("none sentinel is cleared and seeded from host", func(t *testing.T) {
		pc, err := prepareForWorkspaceHost(t, "none", "12345")
		require.NoError(t, err)
		assert.Equal(t, "", pc.Config.WorkspaceID)
		id, err := pc.CurrentWorkspaceID(context.Background())
		require.NoError(t, err)
		assert.Equal(t, int64(12345), id)
	})

	t.Run("connection-id user with numeric host does not error and is seeded", func(t *testing.T) {
		pc, err := prepareForWorkspaceHost(t, "my-connection-id", "12345")
		require.NoError(t, err)
		id, err := pc.CurrentWorkspaceID(context.Background())
		require.NoError(t, err)
		assert.Equal(t, int64(12345), id)
	})
}

// TestPrepareDatabricksClient_HostMetadataUntouchedCases verifies the paths that
// must NOT seed or fail: account/unified hosts, and hosts whose metadata omits
// workspace_id (which preserves the lazy SCIM /Me fallback).
func TestPrepareDatabricksClient_HostMetadataUntouchedCases(t *testing.T) {
	t.Run("account host is untouched", func(t *testing.T) {
		cfg := &config.Config{
			Host:      "https://accounts.test.invalid",
			AccountID: "abc",
			Token:     "test-token",
			Loaders:   []config.Loader{noopLoader{}},
			HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) {
				return &config.HostMetadata{HostType: config.AccountHost}, nil
			},
		}
		pc, err := PrepareDatabricksClient(context.Background(), cfg, nil)
		require.NoError(t, err)
		assert.Equal(t, config.AccountHost, pc.HostTypeForTerraform())
	})

	t.Run("unified host with metadata workspace_id is skipped", func(t *testing.T) {
		cfg := &config.Config{
			Host:        "https://test.invalid",
			Token:       "test-token",
			WorkspaceID: "99999",
			Loaders:     []config.Loader{noopLoader{}},
			HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) {
				return &config.HostMetadata{WorkspaceID: "12345", HostType: config.UnifiedHost}, nil
			},
		}
		// Unified host: the numeric mismatch check is skipped, so no fail-fast.
		pc, err := PrepareDatabricksClient(context.Background(), cfg, nil)
		require.NoError(t, err)
		assert.Equal(t, config.UnifiedHost, pc.HostTypeForTerraform())
	})

	t.Run("host omitting workspace_id does not seed", func(t *testing.T) {
		pc, err := prepareForWorkspaceHost(t, "", "")
		require.NoError(t, err)
		// No seed happened, so CurrentWorkspaceID would attempt a lookup. We only
		// assert the config was left empty (no back-fill from empty metadata).
		assert.Equal(t, "", pc.Config.WorkspaceID)
	})
}

// TestPrepareDatabricksClient_SingleMetadataFetch guards the hard constraint that
// wrapping the resolver adds no /.well-known/databricks-config request: the
// resolver must be invoked exactly once across PrepareDatabricksClient and a
// subsequent CurrentWorkspaceID (which is a seeded cache hit).
func TestPrepareDatabricksClient_SingleMetadataFetch(t *testing.T) {
	var calls atomic.Int32
	cfg := &config.Config{
		Host:    "https://test.invalid",
		Token:   "test-token",
		Loaders: []config.Loader{noopLoader{}},
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) {
			calls.Add(1)
			return &config.HostMetadata{WorkspaceID: "12345", HostType: config.WorkspaceHost}, nil
		},
	}
	pc, err := PrepareDatabricksClient(context.Background(), cfg, nil)
	require.NoError(t, err)
	_, err = pc.CurrentWorkspaceID(context.Background())
	require.NoError(t, err)
	assert.Equal(t, int32(1), calls.Load(), "expected exactly one .well-known/databricks-config fetch")
}

// TestPrepareDatabricksClient_HonorsHostMetadataResolverFactory verifies that the
// resolver wrapper installed during configuration honors
// config.DefaultHostMetadataResolverFactory. The wrapper always sets
// cfg.HostMetadataResolver, so the SDK never consults the factory itself; the
// wrapper must replicate the SDK's precedence and consult it. A resolver returning
// the factory's metadata is proven by the config being back-filled from it. Without
// the fix the wrapper falls through to the built-in HTTP fetch (against the fake
// host), which yields no metadata and no back-fill.
func TestPrepareDatabricksClient_HonorsHostMetadataResolverFactory(t *testing.T) {
	var factoryCalls atomic.Int32
	prevFactory := config.DefaultHostMetadataResolverFactory
	t.Cleanup(func() { config.DefaultHostMetadataResolverFactory = prevFactory })
	config.DefaultHostMetadataResolverFactory = func(*config.Config) config.HostMetadataResolver {
		return func(context.Context, string) (*config.HostMetadata, error) {
			factoryCalls.Add(1)
			return &config.HostMetadata{WorkspaceID: "778899", HostType: config.WorkspaceHost}, nil
		}
	}

	cfg := &config.Config{
		Host:    "https://test.invalid",
		Token:   "test-token",
		Loaders: []config.Loader{noopLoader{}},
		// HostMetadataResolver deliberately left nil so the factory is the source.
	}
	pc, err := PrepareDatabricksClient(context.Background(), cfg, nil)
	require.NoError(t, err)
	assert.Equal(t, int32(1), factoryCalls.Load(), "factory resolver must be consulted exactly once")
	// Back-filled from the factory's metadata (empty on main, where the factory is bypassed).
	assert.Equal(t, "778899", pc.Config.WorkspaceID)
	id, err := pc.CurrentWorkspaceID(context.Background())
	require.NoError(t, err)
	assert.Equal(t, int64(778899), id)
}
