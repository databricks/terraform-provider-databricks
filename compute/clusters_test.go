package compute

import (
	"os"
	"reflect"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetOrCreateRunningCluster_AzureAuth(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Response: map[string]interface{}{},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/list-node-types",
			Response: NodeTypeList{
				[]NodeType{
					{
						NodeTypeID:     "Standard_F4s",
						InstanceTypeID: "Standard_F4s",
						MemoryMB:       8192,
						NumCores:       4,
						NodeInstanceType: &NodeInstanceType{
							LocalDisks:      1,
							InstanceTypeID:  "Standard_F4s",
							LocalDiskSizeGB: 16,
							LocalNVMeDisks:  0,
						},
					},
					{
						NodeTypeID:     "Standard_L80s_v2",
						InstanceTypeID: "Standard_L80s_v2",
						MemoryMB:       655360,
						NumCores:       80,
						NodeInstanceType: &NodeInstanceType{
							LocalDisks:      2,
							InstanceTypeID:  "Standard_L80s_v2",
							LocalDiskSizeGB: 160,
							LocalNVMeDisks:  1,
						},
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/create",
			ExpectedRequest: Cluster{
				AutoterminationMinutes: 10,
				ClusterName:            "mount",
				NodeTypeID:             "Standard_F4s",
				NumWorkers:             1,
				SparkVersion:           CommonRuntimeVersion(),
			},
			Response: ClusterID{
				ClusterID: "bcd",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=bcd",
			Response: ClusterInfo{
				State: "RUNNING",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	client.AzureAuth.ResourceID = "/a/b/c"

	clusterInfo, err := NewClustersAPI(client).GetOrCreateRunningCluster("mount")
	require.NoError(t, err)

	assert.NotNil(t, clusterInfo)
}

func TestGetOrCreateRunningCluster_Existing_AzureAuth(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Response: ClusterList{
				Clusters: []ClusterInfo{
					{
						ClusterID:              "abc",
						State:                  "TERMINATED",
						AutoterminationMinutes: 10,
						ClusterName:            "mount",
						NodeTypeID:             "Standard_F4s",
						NumWorkers:             1,
						SparkVersion:           CommonRuntimeVersion(),
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State: "TERMINATED",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/start",
			ExpectedRequest: ClusterID{
				ClusterID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State: "RUNNING",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	client.AzureAuth.ResourceID = "/a/b/c"

	clusterInfo, err := NewClustersAPI(client).GetOrCreateRunningCluster("mount")
	require.NoError(t, err)

	assert.NotNil(t, clusterInfo)
}

func TestWaitForClusterStatus_RetryOnNotFound(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: common.APIErrorBody{
				Message: "Nope",
			},
			Status: 404,
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State: "RUNNING",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	client.AzureAuth.ResourceID = "/a/b/c"

	clusterInfo, err := NewClustersAPI(client).waitForClusterStatus("abc", ClusterStateRunning)
	require.NoError(t, err)

	assert.NotNil(t, clusterInfo)
}

func TestWaitForClusterStatus_StopRetryingEarly(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: common.APIErrorBody{
				Message: "I am a teapot",
			},
			Status: 418,
		},
	})
	defer server.Close()
	require.NoError(t, err)

	_, err = NewClustersAPI(client).waitForClusterStatus("abc", ClusterStateRunning)
	require.Error(t, err)
	require.Contains(t, err.Error(), "I am a teapot")
}

func TestWaitForClusterStatus_NotReachable(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:        ClusterStateUnknown,
				StateMessage: "Something strange is going on",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	client.AzureAuth.ResourceID = "/a/b/c"

	_, err = NewClustersAPI(client).waitForClusterStatus("abc", ClusterStateRunning)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "abc is not able to transition from UNKNOWN to RUNNING: Something strange is going on.")
}

func TestWaitForClusterStatus_NormalRetry(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State: ClusterStatePending,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State: ClusterStateRunning,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	clusterInfo, err := NewClustersAPI(client).waitForClusterStatus("abc", ClusterStateRunning)
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestEditCluster_Pending(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStatePending,
				ClusterID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStateRunning,
				ClusterID: "abc",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/edit",
			Response: Cluster{
				ClusterID:   "abc",
				ClusterName: "Morty",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State: ClusterStateRunning,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	clusterInfo, err := NewClustersAPI(client).Edit(Cluster{
		ClusterID:   "abc",
		ClusterName: "Morty",
	})
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestEditCluster_Terminating(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStateTerminating,
				ClusterID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStateTerminated,
				ClusterID: "abc",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/edit",
			Response: Cluster{
				ClusterID:   "abc",
				ClusterName: "Morty",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State: ClusterStateRunning,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	clusterInfo, err := NewClustersAPI(client).Edit(Cluster{
		ClusterID:   "abc",
		ClusterName: "Morty",
	})
	require.NoError(t, err)
	assert.Equal(t, ClusterStateTerminated, string(clusterInfo.State))
}

func TestEditCluster_Error(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:        ClusterStateError,
				ClusterID:    "abc",
				StateMessage: "I am a teapot",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	_, err = NewClustersAPI(client).Edit(Cluster{
		ClusterID:   "abc",
		ClusterName: "Morty",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "I am a teapot")
}

func TestStartAndGetInfo_Pending(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStatePending,
				ClusterID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStateRunning,
				ClusterID: "abc",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	clusterInfo, err := NewClustersAPI(client).StartAndGetInfo("abc")
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestStartAndGetInfo_Terminating(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStateTerminating,
				ClusterID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStateTerminated,
				ClusterID: "abc",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/start",
			ExpectedRequest: ClusterID{
				ClusterID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStateRunning,
				ClusterID: "abc",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	clusterInfo, err := NewClustersAPI(client).StartAndGetInfo("abc")
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestStartAndGetInfo_Error(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:        ClusterStateError,
				StateMessage: "I am a teapot",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/start",
			ExpectedRequest: ClusterID{
				ClusterID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:     ClusterStateRunning,
				ClusterID: "abc",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	clusterInfo, err := NewClustersAPI(client).StartAndGetInfo("abc")
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestStartAndGetInfo_StartingError(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:        ClusterStateError,
				StateMessage: "I am a teapot",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/start",
			ExpectedRequest: ClusterID{
				ClusterID: "abc",
			},
			Response: common.APIErrorBody{
				Message: "I am a teapot!",
			},
			Status: 418,
		},
	})
	defer server.Close()
	require.NoError(t, err)

	_, err = NewClustersAPI(client).StartAndGetInfo("abc")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "I am a teapot")
}

func TestPermanentDelete_Pinned(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/delete",
			ExpectedRequest: ClusterID{
				ClusterID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State: ClusterStateTerminated,
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/permanent-delete",
			ExpectedRequest: ClusterID{
				ClusterID: "abc",
			},
			Response: common.APIErrorBody{
				Message: "unpin the cluster first",
			},
			Status: 400,
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/unpin",
			ExpectedRequest: ClusterID{
				ClusterID: "abc",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/permanent-delete",
			ExpectedRequest: ClusterID{
				ClusterID: "abc",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	err = NewClustersAPI(client).PermanentDelete("abc")
	require.NoError(t, err)
}

func TestAccListClustersIntegration(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.CommonEnvironmentClient()
	randomName := qa.RandomName()

	cluster := Cluster{
		NumWorkers:             1,
		ClusterName:            "Terraform Integration Test " + randomName,
		SparkVersion:           CommonRuntimeVersion(),
		InstancePoolID:         CommonInstancePoolID(),
		IdempotencyToken:       "acc-list-" + randomName,
		AutoterminationMinutes: 15,
	}
	clusterReadInfo, err := NewClustersAPI(client).Create(cluster)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.NumWorkers == cluster.NumWorkers)
	assert.True(t, clusterReadInfo.ClusterName == cluster.ClusterName)
	assert.True(t, reflect.DeepEqual(clusterReadInfo.SparkEnvVars, cluster.SparkEnvVars))
	assert.True(t, clusterReadInfo.SparkVersion == cluster.SparkVersion)
	assert.True(t, clusterReadInfo.AutoterminationMinutes == cluster.AutoterminationMinutes)
	assert.True(t, clusterReadInfo.State == ClusterStateRunning)

	defer func() {
		err = NewClustersAPI(client).Terminate(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		clusterReadInfo, err = NewClustersAPI(client).Get(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
		assert.True(t, clusterReadInfo.State == ClusterStateTerminated)

		err = NewClustersAPI(client).Unpin(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		err = NewClustersAPI(client).PermanentDelete(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	err = NewClustersAPI(client).Pin(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)

	clusterReadInfo, err = NewClustersAPI(client).Get(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.State == ClusterStateRunning)
}

func TestClusters_SortNodeTypes_Deprecated(t *testing.T) {
	nodeTypes := []NodeType{
		{
			IsDeprecated: true,
			NodeTypeID:   "deprecated1",
		},
		{
			IsDeprecated: false,
			NodeTypeID:   "not deprecated",
		},
		{
			IsDeprecated: true,
			NodeTypeID:   "deprecated2",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "not deprecated", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_Memory(t *testing.T) {
	nodeTypes := []NodeType{
		{
			MemoryMB:   3,
			NodeTypeID: "3",
		},
		{
			MemoryMB:   1,
			NodeTypeID: "1",
		},
		{
			MemoryMB:   2,
			NodeTypeID: "2",
		},
		{
			MemoryMB:   2,
			NodeTypeID: "another 2",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_CPU(t *testing.T) {
	nodeTypes := []NodeType{
		{
			NumCores:   3,
			NodeTypeID: "3",
		},
		{
			NumCores:   1,
			NodeTypeID: "1",
		},
		{
			NumCores:   2,
			NodeTypeID: "2",
		},
		{
			NumCores:   1,
			NodeTypeID: "another 1",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_GPU(t *testing.T) {
	nodeTypes := []NodeType{
		{
			NumGPUs:    3,
			NodeTypeID: "3",
		},
		{
			NumGPUs:    1,
			NodeTypeID: "1",
		},
		{
			NumGPUs:    2,
			NodeTypeID: "2",
		},
		{
			NumGPUs:    1,
			NodeTypeID: "another 1",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_CPU_Deprecated(t *testing.T) {
	nodeTypes := []NodeType{
		{
			NumCores:     3,
			IsDeprecated: false,
			NodeTypeID:   "3 not deprecated",
		},
		{
			NumCores:     1,
			IsDeprecated: true,
			NodeTypeID:   "1 deprecated",
		},
		{
			NumCores:     2,
			IsDeprecated: false,
			NodeTypeID:   "2 not deprecated",
		},
		{
			NumCores:     1,
			IsDeprecated: false,
			NodeTypeID:   "1 not deprecated",
		},
		{
			NumCores:     2,
			IsDeprecated: true,
			NodeTypeID:   "2 deprecated",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1 not deprecated", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_LocalDisks(t *testing.T) {
	nodeTypes := []NodeType{
		{
			NodeInstanceType: &NodeInstanceType{
				LocalDisks: 3,
			},
			NodeTypeID: "3",
		},
		{
			NodeInstanceType: &NodeInstanceType{
				LocalDisks: 1,
			},
			NodeTypeID: "1",
		},
		{
			NodeInstanceType: &NodeInstanceType{
				LocalDisks: 2,
			},
			NodeTypeID: "2",
		},
		{
			NodeInstanceType: &NodeInstanceType{
				LocalDisks: 3,
			},
			NodeTypeID: "another 3",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1", smallestNodeType.NodeTypeID)
}

func TestAwsAccSmallestNodeType(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.CommonEnvironmentClient()
	nodeType := NewClustersAPI(client).GetSmallestNodeTypeWithStorage()
	assert.Equal(t, "m5d.large", nodeType)
}
