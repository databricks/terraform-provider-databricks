package provisioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAzureResourceId_AzureWorkspace(t *testing.T) {
	w := Workspace{
		WorkspaceName: "test",
		AzureWorkspaceInfo: &AzureWorkspaceInfo{
			SubscriptionId: "sub",
			ResourceGroup:  "rg",
		},
	}
	assert.Equal(t, "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Databricks/workspaces/test", w.AzureResourceId())
}

func TestAzureResourceId_NonAzureWorkspace(t *testing.T) {
	w := Workspace{
		WorkspaceName: "test",
	}
	assert.Equal(t, "", w.AzureResourceId())
}
