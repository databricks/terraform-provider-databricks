package clusters_test

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAccCluster_ListNodeTypesAWS validates that an AWS workspace's ListNodeTypes
// response contains the expected AWS default node types (used by defaultSmallestNodeType)
// and does not contain Azure or GCP node types.
func TestAccCluster_ListNodeTypesAWS(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if !acceptance.IsAws(t) {
		t.Skip("Skipping test because it requires AWS")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	azureAndGcpNodeTypes := []string{
		"Standard_D4pds_v6",
		"Standard_D4ds_v5",
		"n1-standard-4",
	}
	for _, nt := range nodeTypes.NodeTypes {
		for _, excluded := range azureAndGcpNodeTypes {
			assert.NotEqual(t, excluded, nt.NodeTypeId,
				"node type %s should not be present in the response", excluded)
		}
	}

	awsNodeTypes := []string{
		"rgd-fleet.xlarge",
		"m6g.xlarge",
		"md-fleet.xlarge",
		"i3.xlarge",
	}
	nodeTypeIds := make(map[string]bool)
	for _, nt := range nodeTypes.NodeTypes {
		nodeTypeIds[nt.NodeTypeId] = true
	}
	for _, expected := range awsNodeTypes {
		assert.True(t, nodeTypeIds[expected],
			"expected AWS node type %s to be present in the response", expected)
	}
}

// TestAccClusterAPI_ListNodeTypesAzure validates that an Azure workspace's ListNodeTypes
// response contains the expected Azure default node types (used by defaultSmallestNodeType)
// and does not contain AWS or GCP node types.
func TestAccClusterAPI_ListNodeTypesAzure(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	awsAndGcpNodeTypes := []string{
		"rgd-fleet.xlarge",
		"m6g.xlarge",
		"md-fleet.xlarge",
		"i3.xlarge",
		"n1-standard-4",
	}
	for _, nt := range nodeTypes.NodeTypes {
		for _, excluded := range awsAndGcpNodeTypes {
			assert.NotEqual(t, excluded, nt.NodeTypeId,
				"node type %s should not be present in the response", excluded)
		}
	}

	azureNodeTypes := []string{
		"Standard_D4pds_v6",
		"Standard_D4ds_v5",
	}
	nodeTypeIds := make(map[string]bool)
	for _, nt := range nodeTypes.NodeTypes {
		nodeTypeIds[nt.NodeTypeId] = true
	}
	for _, expected := range azureNodeTypes {
		assert.True(t, nodeTypeIds[expected],
			"expected Azure node type %s to be present in the response", expected)
	}
}

// TestAccClusterAPI_ListNodeTypesGcp validates that a GCP workspace's ListNodeTypes
// response contains the expected GCP default node types (used by defaultSmallestNodeType)
// and does not contain AWS or Azure node types.
func TestAccClusterAPI_ListNodeTypesGcp(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	awsAndAzureNodeTypes := []string{
		"rgd-fleet.xlarge",
		"m6g.xlarge",
		"md-fleet.xlarge",
		"i3.xlarge",
		"Standard_D4pds_v6",
		"Standard_D4ds_v5",
	}
	for _, nt := range nodeTypes.NodeTypes {
		for _, excluded := range awsAndAzureNodeTypes {
			assert.NotEqual(t, excluded, nt.NodeTypeId,
				"node type %s should not be present in the response", excluded)
		}
	}

	gcpNodeTypes := []string{
		"n1-standard-4",
	}
	nodeTypeIds := make(map[string]bool)
	for _, nt := range nodeTypes.NodeTypes {
		nodeTypeIds[nt.NodeTypeId] = true
	}
	for _, expected := range gcpNodeTypes {
		assert.True(t, nodeTypeIds[expected],
			"expected GCP node type %s to be present in the response", expected)
	}
}
