package nodetypecloud_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/qa/replay/clienttest"
)

func TestNodeTypeCloudDetection(t *testing.T) {
	for _, test := range []struct {
		name  string
		cloud string
	}{
		{"TestAccCluster_ListNodeTypesAWS", "aws"},
		{"TestAccCluster_ListNodeTypesAWSUcws", "aws"},
		{"TestAccClusterAPI_ListNodeTypesAzure", "azure"},
		{"TestAccClusterAPI_ListNodeTypesAzureUcws", "azure"},
		{"TestAccClusterAPI_ListNodeTypesGcp", "gcp"},
		{"TestAccClusterAPI_ListNodeTypesGcpUcws", "gcp"},
	} {
		t.Run(test.name, func(t *testing.T) {
			client := clienttest.WorkspaceClient(t, "testdata/"+test.cloud+".yaml")
			nodeTypes, err := client.Clusters.ListNodeTypes(context.Background())
			require.NoError(t, err)
			require.NotEmpty(t, nodeTypes.NodeTypes)
			assert.Equal(t, test.cloud == "aws", clusters.IsAws(nodeTypes))
			assert.Equal(t, test.cloud == "azure", clusters.IsAzure(nodeTypes))
			assert.Equal(t, test.cloud == "gcp", clusters.IsGcp(nodeTypes))
		})
	}
}
