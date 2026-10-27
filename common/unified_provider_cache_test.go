package common

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnifiedProviderConcurrentClientCache(t *testing.T) {
	for _, workspaces := range []int{1, 4} {
		t.Run(fmt.Sprintf("%d_workspaces", workspaces), func(t *testing.T) {
			cfg := unifiedHostConfig(t, "https://unified.cloud.databricks.com")
			inner, err := client.New(cfg)
			require.NoError(t, err)
			parent := &DatabricksClient{DatabricksClient: inner}
			const workers = 32
			clients := make([]*DatabricksClient, workers)
			errors := make([]error, workers)
			start := make(chan struct{})
			var group sync.WaitGroup
			for i := range workers {
				group.Go(func() {
					<-start
					// Connection IDs use direct routing and need no account lookup.
					id := fmt.Sprintf("connection-%d", i%workspaces)
					clients[i], errors[i] = parent.getDatabricksClientForUnifiedProvider(context.Background(), id)
				})
			}
			close(start)
			group.Wait()
			for i, result := range clients {
				require.NoError(t, errors[i])
				require.NotNil(t, result)
				id := fmt.Sprintf("connection-%d", i%workspaces)
				assert.Equal(t, id, result.Config.WorkspaceID)
				assert.Same(t, clients[i%workspaces].DatabricksClient, result.DatabricksClient)
				workspace, err := result.WorkspaceClient()
				require.NoError(t, err)
				assert.Same(t, parent.cachedWorkspaceClients[id], workspace)
			}
		})
	}
}

func TestUnifiedProviderConcurrentWorkspaceCacheUpdates(t *testing.T) {
	cfg := unifiedHostConfig(t, "https://unified.cloud.databricks.com")
	inner, err := client.New(cfg)
	require.NoError(t, err)
	workspace, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	parent := &DatabricksClient{DatabricksClient: inner}
	parent.SetWorkspaceClientForWorkspace("connection-1", workspace)
	_, err = parent.getDatabricksClientForUnifiedProvider(context.Background(), "connection-1")
	require.NoError(t, err)
	var group sync.WaitGroup
	for range 16 {
		group.Go(func() {
			for range 100 {
				parent.SetWorkspaceClientForWorkspace("connection-1", workspace)
			}
		})
		group.Go(func() {
			for range 100 {
				result, err := parent.getDatabricksClientForUnifiedProvider(context.Background(), "connection-1")
				if assert.NoError(t, err) {
					assert.Same(t, workspace, result.cachedWorkspaceClient)
				}
			}
		})
	}
	group.Wait()
}

func TestUnifiedProviderClientCacheRetriesAfterInvalidWorkspace(t *testing.T) {
	cfg := unifiedHostConfig(t, "https://unified.cloud.databricks.com")
	inner, err := client.New(cfg)
	require.NoError(t, err)
	parent := &DatabricksClient{DatabricksClient: inner}
	_, err = parent.getDatabricksClientForUnifiedProvider(context.Background(), "")
	require.ErrorContains(t, err, "workspace_id must not be empty")
	result, err := parent.getDatabricksClientForUnifiedProvider(context.Background(), "connection-1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, config.UnifiedHost, result.Config.HostType())
}
