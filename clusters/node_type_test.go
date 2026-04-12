package clusters_test

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAccCluster_ListNodeTypesAWS validates that an AWS workspace's ListNodeTypes
// response is detected as AWS and not Azure or GCP.
func TestAccCluster_ListNodeTypesAWS(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if !acceptance.IsAws(t) {
		t.Skip("Skipping test because it requires AWS")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	assert.True(t, clusters.IsAws(nodeTypes), "expected IsAws to be true for AWS workspace")
	assert.False(t, clusters.IsAzure(nodeTypes), "expected IsAzure to be false for AWS workspace")
	assert.False(t, clusters.IsGcp(nodeTypes), "expected IsGcp to be false for AWS workspace")
}

// TestAccCluster_ListNodeTypesAWSUcws validates that an AWS workspace's ListNodeTypes
// response is detected as AWS and not Azure or GCP.
func TestAccCluster_ListNodeTypesAWSUcws(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	if !acceptance.IsAws(t) {
		t.Skip("Skipping test because it requires AWS")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	assert.True(t, clusters.IsAws(nodeTypes), "expected IsAws to be true for AWS workspace")
	assert.False(t, clusters.IsAzure(nodeTypes), "expected IsAzure to be false for AWS workspace")
	assert.False(t, clusters.IsGcp(nodeTypes), "expected IsGcp to be false for AWS workspace")
}

// TestAccClusterAPI_ListNodeTypesAzure validates that an Azure workspace's ListNodeTypes
// response is detected as Azure and not AWS or GCP.
func TestAccClusterAPI_ListNodeTypesAzure(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if !acceptance.IsAzure(t) {
		t.Skip("Skipping test because it requires Azure")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	assert.False(t, clusters.IsAws(nodeTypes), "expected IsAws to be false for Azure workspace")
	assert.True(t, clusters.IsAzure(nodeTypes), "expected IsAzure to be true for Azure workspace")
	assert.False(t, clusters.IsGcp(nodeTypes), "expected IsGcp to be false for Azure workspace")
}

// TestAccClusterAPI_ListNodeTypesAzureUcws validates that an Azure workspace's ListNodeTypes
// response is detected as Azure and not AWS or GCP.
func TestAccClusterAPI_ListNodeTypesAzureUcws(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	if !acceptance.IsAzure(t) {
		t.Skip("Skipping test because it requires Azure")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	assert.False(t, clusters.IsAws(nodeTypes), "expected IsAws to be false for Azure workspace")
	assert.True(t, clusters.IsAzure(nodeTypes), "expected IsAzure to be true for Azure workspace")
	assert.False(t, clusters.IsGcp(nodeTypes), "expected IsGcp to be false for Azure workspace")
}

// TestAccClusterAPI_ListNodeTypesGcp validates that a GCP workspace's ListNodeTypes
// response is detected as GCP and not AWS or Azure.
func TestAccClusterAPI_ListNodeTypesGcp(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if !acceptance.IsGcp(t) {
		t.Skip("Skipping test because it requires GCP")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	assert.False(t, clusters.IsAws(nodeTypes), "expected IsAws to be false for GCP workspace")
	assert.False(t, clusters.IsAzure(nodeTypes), "expected IsAzure to be false for GCP workspace")
	assert.True(t, clusters.IsGcp(nodeTypes), "expected IsGcp to be true for GCP workspace")
}

// TestAccClusterAPI_ListNodeTypesGcpUcws validates that a GCP workspace's ListNodeTypes
// response is detected as GCP and not AWS or Azure.
func TestAccClusterAPI_ListNodeTypesGcpUcws(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	if !acceptance.IsGcp(t) {
		t.Skip("Skipping test because it requires GCP")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	assert.False(t, clusters.IsAws(nodeTypes), "expected IsAws to be false for GCP workspace")
	assert.False(t, clusters.IsAzure(nodeTypes), "expected IsAzure to be false for GCP workspace")
	assert.True(t, clusters.IsGcp(nodeTypes), "expected IsGcp to be true for GCP workspace")
}
