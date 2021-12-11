package clusters

import (
	"context"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClustersDataSource(t *testing.T) {
	d, err := qa.ResourceFixture{
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
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, []interface{}{"a", "b"}, d.Get("ids"))
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
	assert.Equal(t, []interface{}{"b"}, d.Get("ids"))
}

func TestClustersDataSourceErrorsOut(t *testing.T) {
	diag := DataSourceClusters().ReadContext(context.Background(), nil, &common.DatabricksClient{
		Host: ".", Token: "."})
	assert.NotNil(t, diag)
	assert.True(t, diag.HasError())
}
