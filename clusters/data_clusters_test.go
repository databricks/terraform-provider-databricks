package clusters

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClustersDataSource(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",

				Response: ClusterList{
					Clusters: []ClusterInfo{
						{
							ClusterID: "b",
						},
						{
							ClusterID: "a",
						},
					},
				},
			},
		},
		Resource:    DataSourceClusters(),
		NonWritable: true,
		Read:        true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestClustersDataSourceContainsName(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",
				Response: ClusterList{
					Clusters: []ClusterInfo{
						{
							ClusterID:   "b",
							ClusterName: "THIS NAME",
						},
						{
							ClusterID:   "a",
							ClusterName: "that name",
						},
					},
				},
			},
		},
		Resource:    DataSourceClusters(),
		NonWritable: true,
		Read:        true,
		ID:          "_",
		HCL:         `cluster_name_contains = "this"`,
	}.Apply(t)
	require.NoError(t, err)
	ids := d.Get("ids").(*schema.Set)
	assert.True(t, ids.Contains("b"))
	assert.Equal(t, 1, ids.Len())
}

func TestClustersDataSourceErrorsOut(t *testing.T) {
	client, _ := client.New(&config.Config{
		Host:                ".",
		Token:               ".",
		RetryTimeoutSeconds: 1,
		HTTPTimeoutSeconds:  1,
	})
	diag := DataSourceClusters().ToResource().ReadContext(context.Background(), nil, &common.DatabricksClient{
		DatabricksClient: client,
	})
	assert.NotNil(t, diag)
	assert.True(t, diag.HasError())
}
