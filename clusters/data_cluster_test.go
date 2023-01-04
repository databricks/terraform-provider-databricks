package clusters

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClusterDataByID(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=abc",
				Response: ClusterInfo{
					ClusterID:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  ClusterStateRunning,
					AutoScale: &AutoScale{
						MaxWorkers: 4,
					},
				},
			},
		},
		Resource:    DataSourceCluster(),
		HCL:         `cluster_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "abc",
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, 15, d.Get("cluster_info.0.autotermination_minutes"))
	assert.Equal(t, "Shared Autoscaling", d.Get("cluster_info.0.cluster_name"))
	assert.Equal(t, "i3.xlarge", d.Get("cluster_info.0.node_type_id"))
	assert.Equal(t, 4, d.Get("cluster_info.0.autoscale.0.max_workers"))
	assert.Equal(t, "RUNNING", d.Get("cluster_info.0.state"))

	for k, v := range d.State().Attributes {
		fmt.Printf("assert.Equal(t, %#v, d.Get(%#v))\n", v, k)
	}
}

func TestClusterDataByName(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",

				Response: ClusterList{
					Clusters: []ClusterInfo{{
						ClusterID:              "abc",
						NumWorkers:             100,
						ClusterName:            "Shared Autoscaling",
						SparkVersion:           "7.1-scala12",
						NodeTypeID:             "i3.xlarge",
						AutoterminationMinutes: 15,
						State:                  ClusterStateRunning,
						AutoScale: &AutoScale{
							MaxWorkers: 4,
						},
					}},
				},
			},
		},
		Resource:    DataSourceCluster(),
		HCL:         `name = "Shared Autoscaling"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, 15, d.Get("cluster_info.0.autotermination_minutes"))
	assert.Equal(t, "Shared Autoscaling", d.Get("cluster_info.0.cluster_name"))
	assert.Equal(t, "i3.xlarge", d.Get("cluster_info.0.node_type_id"))
	assert.Equal(t, 4, d.Get("cluster_info.0.autoscale.0.max_workers"))
	assert.Equal(t, "RUNNING", d.Get("cluster_info.0.state"))

	for k, v := range d.State().Attributes {
		fmt.Printf("assert.Equal(t, %#v, d.Get(%#v))\n", v, k)
	}
}

func TestClusterDataByName_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",

				Response: ClusterList{
					Clusters: []ClusterInfo{},
				},
			},
		},
		Resource:    DataSourceCluster(),
		HCL:         `name = "Unknown"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "there is no cluster with name 'Unknown'")
}

func TestClusterDataByName_ListError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceCluster(),
		HCL:         `name = "Unknown"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}

func TestClusterData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceCluster(),
		Read:        true,
		NonWritable: true,
		HCL:         `cluster_id = "abc"`,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}

func TestClusterData_ErrorNoParams(t *testing.T) {
	qa.ResourceFixture{
		Resource:    DataSourceCluster(),
		Read:        true,
		NonWritable: true,
		HCL:         "",
		ID:          "_",
	}.ExpectError(t, "you need to specify either `name` or `cluster_id`")
}
