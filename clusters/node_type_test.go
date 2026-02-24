package clusters_test

import (
	"context"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAccCluster_ListNodeTypesAWS validates that an AWS workspace's ListNodeTypes
// response contains node types matching the AWS detection pattern ("." followed by "large")
// and does not contain node types matching the Azure detection pattern ("Standard_").
func TestAccCluster_ListNodeTypesAWS(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if !acceptance.IsAws(t) {
		t.Skip("Skipping test because it requires AWS")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	// Validate no node types match the Azure pattern
	for _, nt := range nodeTypes.NodeTypes {
		assert.False(t, strings.Contains(nt.NodeTypeId, "Standard_"),
			"node type %s should not match the Azure pattern (Standard_)", nt.NodeTypeId)
	}

	// Validate at least one node type matches the AWS pattern ("." followed by "large")
	hasAwsPattern := false
	for _, nt := range nodeTypes.NodeTypes {
		dotIdx := strings.Index(nt.NodeTypeId, ".")
		if dotIdx >= 0 && strings.Contains(nt.NodeTypeId[dotIdx:], "large") {
			hasAwsPattern = true
			break
		}
	}
	assert.True(t, hasAwsPattern,
		"expected at least one node type matching the AWS pattern (dot followed by large)")
}

// TestAccClusterAPI_ListNodeTypesAzure validates that an Azure workspace's ListNodeTypes
// response contains node types matching the Azure detection pattern ("Standard_")
// and does not contain node types matching the AWS detection pattern ("." followed by "large").
func TestAccClusterAPI_ListNodeTypesAzure(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if !acceptance.IsAzure(t) {
		t.Skip("Skipping test because it requires Azure")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	// Validate no node types match the AWS pattern
	for _, nt := range nodeTypes.NodeTypes {
		dotIdx := strings.Index(nt.NodeTypeId, ".")
		assert.False(t, dotIdx >= 0 && strings.Contains(nt.NodeTypeId[dotIdx:], "large"),
			"node type %s should not match the AWS pattern (dot followed by large)", nt.NodeTypeId)
	}

	// Validate at least one node type matches the Azure pattern ("Standard_")
	hasAzurePattern := false
	for _, nt := range nodeTypes.NodeTypes {
		if strings.Contains(nt.NodeTypeId, "Standard_") {
			hasAzurePattern = true
			break
		}
	}
	assert.True(t, hasAzurePattern,
		"expected at least one node type matching the Azure pattern (Standard_)")
}

// TestAccClusterAPI_ListNodeTypesGcp validates that a GCP workspace's ListNodeTypes
// response does not contain node types matching the AWS detection pattern ("." followed by "large")
// or the Azure detection pattern ("Standard_"). GCP is detected as neither AWS nor Azure.
func TestAccClusterAPI_ListNodeTypesGcp(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if !acceptance.IsGcp(t) {
		t.Skip("Skipping test because it requires GCP")
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	nodeTypes, err := w.Clusters.ListNodeTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, nodeTypes.NodeTypes)

	// Validate no node types match the AWS pattern
	for _, nt := range nodeTypes.NodeTypes {
		dotIdx := strings.Index(nt.NodeTypeId, ".")
		assert.False(t, dotIdx >= 0 && strings.Contains(nt.NodeTypeId[dotIdx:], "large"),
			"node type %s should not match the AWS pattern (dot followed by large)", nt.NodeTypeId)
	}

	// Validate no node types match the Azure pattern
	for _, nt := range nodeTypes.NodeTypes {
		assert.False(t, strings.Contains(nt.NodeTypeId, "Standard_"),
			"node type %s should not match the Azure pattern (Standard_)", nt.NodeTypeId)
	}
}
