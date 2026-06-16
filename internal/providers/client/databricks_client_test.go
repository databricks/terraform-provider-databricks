package client

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
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
