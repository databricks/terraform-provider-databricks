package compute

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/pools"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommonInstancePoolID_Existing(t *testing.T) {
	defer common.CleanupEnvironment()()
	common.ResetCommonEnvironmentClient()
	oncePool = sync.Once{}
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Response: pools.InstancePoolList{
				InstancePools: []pools.InstancePoolAndStats{
					{
						InstancePoolID:   "abc",
						InstancePoolName: "Terraform Integration Test by test",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		os.Setenv("DATABRICKS_HOST", client.Host)
		os.Setenv("DATABRICKS_TOKEN", client.Token)
		os.Setenv("USER", "test")

		id := CommonInstancePoolID()
		assert.Equal(t, "abc", id)
	})
}

func TestCommonInstancePoolID_Panic(t *testing.T) {
	defer common.CleanupEnvironment()()
	defer func() { recover() }()
	common.ResetCommonEnvironmentClient()
	oncePool = sync.Once{}
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Status:   404,
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		os.Setenv("DATABRICKS_HOST", client.Host)
		os.Setenv("DATABRICKS_TOKEN", client.Token)

		CommonInstancePoolID()
	})
}

var sparkVersions = clusters.SparkVersionsList{
	SparkVersions: []clusters.SparkVersion{
		{
			Version:     "7.1.x-cpu-ml-scala2.12",
			Description: "7.1 ML (includes Apache Spark 3.0.0, Scala 2.12)",
		},
		{
			Version:     "apache-spark-2.4.x-scala2.11",
			Description: "Light 2.4 (includes Apache Spark 2.4, Scala 2.11)",
		},
		{
			Version:     "7.3.x-scala2.12",
			Description: "7.3 LTS (includes Apache Spark 3.0.1, Scala 2.12)",
		},
		{
			Version:     "6.4.x-scala2.11",
			Description: "6.4 (includes Apache Spark 2.4.5, Scala 2.11)",
		},
	},
}

var nodeTypes = clusters.NodeTypeList{
	NodeTypes: []clusters.NodeType{
		{
			NodeTypeID:     "m4.large",
			InstanceTypeID: "m4.large",
			NodeInstanceType: &clusters.NodeInstanceType{
				LocalDisks:     1,
				InstanceTypeID: "m4.large",
			},
		},
	},
}

func TestNewTinyClusterInCommonPoolPossiblyReused(t *testing.T) {
	defer common.CleanupEnvironment()()
	common.ResetCommonEnvironmentClient()
	oncePool = sync.Once{}
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/spark-versions",
			Response:     sparkVersions,
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Response: pools.InstancePoolList{
				InstancePools: []pools.InstancePoolAndStats{},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-pools/create",
			ExpectedRequest: pools.InstancePool{
				AwsAttributes: &pools.InstancePoolAwsAttributes{
					Availability: "SPOT",
				},
				NodeTypeID:                         "m4.large",
				IdleInstanceAutoTerminationMinutes: 15,
				PreloadedSparkVersions:             []string{"7.3.x-scala2.12"},
				InstancePoolName:                   "Terraform Integration Test by test",
				MaxCapacity:                        10,
			},
			Response: pools.InstancePoolAndStats{
				InstancePoolID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Response: map[string]interface{}{},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/list-node-types",
			Response:     nodeTypes,
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/create",
			Response: clusters.ClusterID{
				ClusterID: "bcd",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=bcd",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)
	// trick common env client
	err = os.Setenv("DATABRICKS_HOST", client.Host)
	assert.NoError(t, err)
	err = os.Setenv("DATABRICKS_TOKEN", client.Token)
	assert.NoError(t, err)
	err = os.Setenv("USER", "test")
	assert.NoError(t, err)

	c := NewTinyClusterInCommonPoolPossiblyReused()
	assert.NotNil(t, c)
}

func TestNewTinyClusterInCommonPool(t *testing.T) {
	defer common.CleanupEnvironment()()
	common.ResetCommonEnvironmentClient()
	oncePool = sync.Once{}
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/spark-versions",
			Response:     sparkVersions,
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Response: pools.InstancePoolList{
				InstancePools: []pools.InstancePoolAndStats{
					{
						InstancePoolName: "Terraform Integration Test by test",
						InstancePoolID:   "abc",
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/create",
			Response: clusters.ClusterID{
				ClusterID: "bcd",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=bcd",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)
	// trick common env client
	os.Setenv("DATABRICKS_HOST", client.Host)
	os.Setenv("DATABRICKS_TOKEN", client.Token)
	os.Setenv("USER", "test")

	c, err := NewTinyClusterInCommonPool()
	require.NoError(t, err)
	assert.NotNil(t, c)

	client = CommonEnvironmentClientWithRealCommandExecutor()
	client.CommandExecutor(context.Background())
}
