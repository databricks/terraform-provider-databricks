package common

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestClientForReconcile() *DatabricksClient {
	return &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{},
		},
	}
}

func TestReconcileWorkspaceIDFromHostMetadata(t *testing.T) {
	tests := []struct {
		name            string
		hostType        config.HostType
		userWorkspaceID string
		hostWorkspaceID string
		wantErr         string
		wantCachedID    int64 // 0 means "not seeded"
	}{
		{
			name:            "workspace host, user empty, seeds from host",
			hostType:        config.WorkspaceHost,
			userWorkspaceID: "",
			hostWorkspaceID: "12345",
			wantCachedID:    12345,
		},
		{
			name:            "workspace host, user matches host, seeds",
			hostType:        config.WorkspaceHost,
			userWorkspaceID: "12345",
			hostWorkspaceID: "12345",
			wantCachedID:    12345,
		},
		{
			name:            "workspace host, numeric mismatch, fails fast",
			hostType:        config.WorkspaceHost,
			userWorkspaceID: "99999",
			hostWorkspaceID: "12345",
			wantErr:         "workspace_id mismatch",
			wantCachedID:    0,
		},
		{
			name:            "workspace host, none sentinel, seeds from host",
			hostType:        config.WorkspaceHost,
			userWorkspaceID: "none",
			hostWorkspaceID: "12345",
			wantCachedID:    12345,
		},
		{
			name:            "workspace host, connection id user, numeric host, no error and seeds",
			hostType:        config.WorkspaceHost,
			userWorkspaceID: "some-connection-id",
			hostWorkspaceID: "12345",
			wantCachedID:    12345,
		},
		{
			name:            "workspace host, host omits workspace_id, no seed no error",
			hostType:        config.WorkspaceHost,
			userWorkspaceID: "12345",
			hostWorkspaceID: "",
			wantCachedID:    0,
		},
		{
			name:            "account host, untouched",
			hostType:        config.AccountHost,
			userWorkspaceID: "",
			hostWorkspaceID: "12345",
			wantCachedID:    0,
		},
		{
			name:            "unified host, skipped even with host workspace_id",
			hostType:        config.UnifiedHost,
			userWorkspaceID: "99999",
			hostWorkspaceID: "12345",
			wantCachedID:    0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClientForReconcile()
			err := c.ReconcileWorkspaceIDFromHostMetadata(tc.hostType, tc.userWorkspaceID, tc.hostWorkspaceID)
			if tc.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.wantErr)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tc.wantCachedID, c.cachedWorkspaceID)
		})
	}
}

// TestReconcileWorkspaceIDFromHostMetadataSeedsCacheHit proves that after seeding,
// CurrentWorkspaceID is a cache hit that never reaches WorkspaceClient()/SCIM /Me.
// The client has no usable credentials or workspace client, so a cache miss would
// error; returning the seeded value with no error demonstrates the SCIM call is
// skipped.
func TestReconcileWorkspaceIDFromHostMetadataSeedsCacheHit(t *testing.T) {
	c := newTestClientForReconcile()
	require.NoError(t, c.ReconcileWorkspaceIDFromHostMetadata(config.WorkspaceHost, "", "678910"))

	id, err := c.CurrentWorkspaceID(context.Background())
	require.NoError(t, err)
	assert.Equal(t, int64(678910), id)
}
