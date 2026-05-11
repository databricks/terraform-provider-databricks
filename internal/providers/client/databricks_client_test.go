package client

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPrepareDatabricksClient_NormalizesNoneWorkspaceID verifies that the
// "none" sentinel that the Databricks CLI writes to ~/.databrickscfg for
// account-level profiles is normalized to an empty string at provider configure
// time. Without this normalization, downstream parseWorkspaceID call sites fail
// with a strconv.ParseInt error.
//
// Both subtests set WorkspaceID to a non-zero value, so the SDK's ConfigAttributes
// env loader (which only fills zero fields) cannot pollute the assertion from an
// ambient DATABRICKS_WORKSPACE_ID. The HostMetadataResolver stub suppresses the
// /.well-known/databricks-config HTTP fetch; the host is a fake .invalid TLD as
// defense in depth.
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
