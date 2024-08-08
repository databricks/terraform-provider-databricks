package clusters

import (
	"context"
	"errors"
	"fmt"

	// "reflect"

	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetOrCreateRunningCluster_AzureAuth(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Response: map[string]any{},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.1/clusters/spark-versions",
			Response: compute.GetSparkVersionsResponse{
				Versions: []compute.SparkVersion{
					{
						Key:  "7.1.x-cpu-ml-scala2.12",
						Name: "7.1 ML (includes Apache Spark 3.0.0, Scala 2.12)",
					},
					{
						Key:  "apache-spark-2.4.x-scala2.11",
						Name: "Light 2.4 (includes Apache Spark 2.4, Scala 2.11)",
					},
					{
						Key:  "7.3.x-scala2.12",
						Name: "7.3 LTS (includes Apache Spark 3.0.1, Scala 2.12)",
					},
					{
						Key:  "6.4.x-scala2.11",
						Name: "6.4 (includes Apache Spark 2.4.5, Scala 2.11)",
					},
				},
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.1/clusters/list-node-types",
			Response: compute.ListNodeTypesResponse{
				NodeTypes: []compute.NodeType{
					{
						NodeTypeId:     "Standard_F4s",
						InstanceTypeId: "Standard_F4s",
						MemoryMb:       8192,
						NumCores:       4,
						NodeInstanceType: &compute.NodeInstanceType{
							LocalDisks:      1,
							InstanceTypeId:  "Standard_F4s",
							LocalDiskSizeGb: 16,
							LocalNvmeDisks:  0,
						},
					},
					{
						NodeTypeId:     "Standard_L80s_v2",
						InstanceTypeId: "Standard_L80s_v2",
						MemoryMb:       655360,
						NumCores:       80,
						NodeInstanceType: &compute.NodeInstanceType{
							LocalDisks:      2,
							InstanceTypeId:  "Standard_L80s_v2",
							LocalDiskSizeGb: 160,
							LocalNvmeDisks:  1,
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
				SparkVersion:           "7.3.x-scala2.12",
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

	client.Config.AzureResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).GetOrCreateRunningCluster("mount")
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
						SparkVersion:           "7.3.x-scala2.12",
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

	client.Config.AzureResourceID = "/a/b/c"

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).GetOrCreateRunningCluster("mount")
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

	client.Config.AzureResourceID = "/a/b/c"

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).waitForClusterStatus("abc", ClusterStateRunning)
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

	ctx := context.Background()
	_, err = NewClustersAPI(ctx, client).waitForClusterStatus("abc", ClusterStateRunning)
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
				TerminationReason: &TerminationReason{Code: "unknown", Type: "broken",
					Parameters: map[string]string{"abc": "def"}},
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	client.Config.AzureResourceID = "/a/b/c"

	ctx := context.Background()
	_, err = NewClustersAPI(ctx, client).waitForClusterStatus("abc", ClusterStateRunning)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "abc is not able to transition from UNKNOWN to RUNNING: Something strange is going on")
	assert.Contains(t, err.Error(), "code: unknown, type: broken")
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

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).waitForClusterStatus("abc", ClusterStateRunning)
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

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).Edit(Cluster{
		ClusterID:   "abc",
		ClusterName: "Morty",
	})
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestResizeCluster_FailsForNonRunningCluster(t *testing.T) {
	clusterStates := []ClusterState{ClusterStateUnknown,
		ClusterStateError,
		ClusterStatePending,
		ClusterStateRestarting,
		ClusterStateResizing,
		ClusterStateTerminating,
		ClusterStateTerminated,
	}
	for _, clusterState := range clusterStates {
		t.Run(fmt.Sprintf("CLUSTER STATE %s", clusterState), func(t *testing.T) {
			client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
				{
					Method:   "GET",
					Resource: "/api/2.0/clusters/get?cluster_id=abc",
					Response: ClusterInfo{
						State:     clusterState,
						ClusterID: "abc",
					},
				},
			})
			require.NoError(t, err)

			ctx := context.Background()
			_, err = NewClustersAPI(ctx, client).Resize(ResizeRequest{
				ClusterID:  "abc",
				NumWorkers: 10,
			})
			require.Error(t, err)
			assert.Contains(t, err.Error(), "resize: Cluster abc is in "+clusterState+" state. RUNNING state required to use resize API")
			server.Close()
		})
	}
}

func TestResizeCluster_NormalRun(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:      ClusterStateRunning,
				ClusterID:  "abc",
				NumWorkers: 4,
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/resize",
			ExpectedRequest: ResizeRequest{
				ClusterID:  "abc",
				NumWorkers: 10,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:      ClusterStateResizing,
				ClusterID:  "abc",
				NumWorkers: 10,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:      ClusterStateRunning,
				ClusterID:  "abc",
				NumWorkers: 10,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).Resize(ResizeRequest{
		ClusterID:  "abc",
		NumWorkers: 10,
	})
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
	assert.Equal(t, 10, int(clusterInfo.NumWorkers))
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

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).Edit(Cluster{
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

	ctx := context.Background()
	_, err = NewClustersAPI(ctx, client).Edit(Cluster{
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

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).StartAndGetInfo("abc")
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestStartAndGetInfo_InitScript(t *testing.T) {
	TestInitScripts := []InitScriptStorageInfo{
		{
			Dbfs: &DbfsStorageInfo{
				Destination: "dbfs:/my_init_script.sh",
			},
		},
		{
			Gcs: &GcsStorageInfo{
				Destination: "gs://my_bucket/my_init_script.sh",
			},
		},
		{
			S3: &S3StorageInfo{
				Destination: "s3://my_bucket/my_init_script.sh",
			},
		},
		{
			Abfss: &AbfssStorageInfo{
				Destination: "abfss://my_bucket/my_init_script.sh",
			},
		},
		{
			File: &LocalFileInfo{
				Destination: "/my_init_script.sh",
			},
		},
	}

	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				State:       ClusterStateRunning,
				ClusterID:   "abc",
				InitScripts: TestInitScripts,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).StartAndGetInfo("abc")
	require.NoError(t, err)
	assert.Equal(t, TestInitScripts, clusterInfo.InitScripts)
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

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).StartAndGetInfo("abc")
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

	ctx := context.Background()
	clusterInfo, err := NewClustersAPI(ctx, client).StartAndGetInfo("abc")
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

	ctx := context.Background()
	_, err = NewClustersAPI(ctx, client).StartAndGetInfo("abc")
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

	ctx := context.Background()
	err = NewClustersAPI(ctx, client).PermanentDelete("abc")
	require.NoError(t, err)
}

func TestEventsSinglePage(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: EventsRequest{
				ClusterID: "abc",
			},
			Response: EventsResponse{
				Events: []ClusterEvent{
					{
						ClusterID: "abc",
						Timestamp: int64(123),
						Type:      EvTypeRunning,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
				},
				TotalCount: 1,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	clusterEvents, err := NewClustersAPI(ctx, client).Events(EventsRequest{ClusterID: "abc"})
	require.NoError(t, err)
	assert.Equal(t, len(clusterEvents), 1)
	assert.Equal(t, clusterEvents[0].ClusterID, "abc")
	assert.Equal(t, clusterEvents[0].Timestamp, int64(123))
	assert.Equal(t, clusterEvents[0].Type, EvTypeRunning)
	assert.Equal(t, clusterEvents[0].Details.CurrentNumWorkers, int32(2))
	assert.Equal(t, clusterEvents[0].Details.TargetNumWorkers, int32(2))
}

func TestEventsTwoPages(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: EventsRequest{
				ClusterID: "abc",
			},
			Response: EventsResponse{
				Events: []ClusterEvent{
					{
						ClusterID: "abc",
						Timestamp: int64(123),
						Type:      EvTypeRunning,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
				},
				TotalCount: 2,
				NextPage: &EventsRequest{
					ClusterID: "abc",
					Offset:    1,
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: EventsRequest{
				ClusterID: "abc",
				Offset:    1,
			},
			Response: EventsResponse{
				Events: []ClusterEvent{
					{
						ClusterID: "abc",
						Timestamp: int64(124),
						Type:      EvTypeTerminating,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
				},
				TotalCount: 2,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	clusterEvents, err := NewClustersAPI(ctx, client).Events(EventsRequest{ClusterID: "abc"})
	require.NoError(t, err)
	assert.Equal(t, len(clusterEvents), 2)
	assert.Equal(t, clusterEvents[0].ClusterID, "abc")
	assert.Equal(t, clusterEvents[0].Timestamp, int64(123))
	assert.Equal(t, clusterEvents[0].Type, EvTypeRunning)
	assert.Equal(t, clusterEvents[0].Details.CurrentNumWorkers, int32(2))
	assert.Equal(t, clusterEvents[0].Details.TargetNumWorkers, int32(2))
	assert.Equal(t, clusterEvents[1].ClusterID, "abc")
	assert.Equal(t, clusterEvents[1].Timestamp, int64(124))
	assert.Equal(t, clusterEvents[1].Type, EvTypeTerminating)
	assert.Equal(t, clusterEvents[1].Details.CurrentNumWorkers, int32(2))
	assert.Equal(t, clusterEvents[1].Details.TargetNumWorkers, int32(2))
}

func TestEventsTwoPagesMaxItems(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: EventsRequest{
				ClusterID: "abc",
				Limit:     1,
			},
			Response: EventsResponse{
				Events: []ClusterEvent{
					{
						ClusterID: "abc",
						Timestamp: int64(123),
						Type:      EvTypeRunning,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
				},
				TotalCount: 2,
				NextPage: &EventsRequest{
					ClusterID: "abc",
					Offset:    1,
				},
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	clusterEvents, err := NewClustersAPI(ctx, client).Events(EventsRequest{ClusterID: "abc", MaxItems: 1, Limit: 1})
	require.NoError(t, err)
	assert.Equal(t, len(clusterEvents), 1)
	assert.Equal(t, clusterEvents[0].ClusterID, "abc")
	assert.Equal(t, clusterEvents[0].Timestamp, int64(123))
	assert.Equal(t, clusterEvents[0].Type, EvTypeRunning)
	assert.Equal(t, clusterEvents[0].Details.CurrentNumWorkers, int32(2))
	assert.Equal(t, clusterEvents[0].Details.TargetNumWorkers, int32(2))
}

func TestEventsTwoPagesMaxThreeItems(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: EventsRequest{
				ClusterID: "abc",
				Limit:     2,
			},
			Response: EventsResponse{
				Events: []ClusterEvent{
					{
						ClusterID: "abc",
						Timestamp: int64(123),
						Type:      EvTypeRunning,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
					{
						ClusterID: "abc",
						Timestamp: int64(124),
						Type:      EvTypeRunning,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
				},
				TotalCount: 4,
				NextPage: &EventsRequest{
					ClusterID: "abc",
					Offset:    2,
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: EventsRequest{
				ClusterID: "abc",
				Offset:    2,
			},
			Response: EventsResponse{
				Events: []ClusterEvent{
					{
						ClusterID: "abc",
						Timestamp: int64(125),
						Type:      EvTypeTerminating,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
					{
						ClusterID: "abc",
						Timestamp: int64(126),
						Type:      EvTypeRunning,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
				},
				TotalCount: 4,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	clusterEvents, err := NewClustersAPI(ctx, client).Events(EventsRequest{ClusterID: "abc", MaxItems: 3, Limit: 2})
	require.NoError(t, err)
	assert.Equal(t, len(clusterEvents), 3)
	assert.Equal(t, clusterEvents[0].ClusterID, "abc")
	assert.Equal(t, clusterEvents[0].Timestamp, int64(123))
	assert.Equal(t, clusterEvents[0].Type, EvTypeRunning)
	assert.Equal(t, clusterEvents[0].Details.CurrentNumWorkers, int32(2))
	assert.Equal(t, clusterEvents[0].Details.TargetNumWorkers, int32(2))
	assert.Equal(t, clusterEvents[1].Timestamp, int64(124))
	assert.Equal(t, clusterEvents[2].Timestamp, int64(125))
}

func TestEventsTwoPagesNoNextPage(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: EventsRequest{
				ClusterID: "abc",
				Limit:     2,
			},
			Response: EventsResponse{
				Events: []ClusterEvent{
					{
						ClusterID: "abc",
						Timestamp: int64(123),
						Type:      EvTypeRunning,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
					{
						ClusterID: "abc",
						Timestamp: int64(124),
						Type:      EvTypeRunning,
						Details: EventDetails{
							CurrentNumWorkers: int32(2),
							TargetNumWorkers:  int32(2),
						},
					},
				},
				TotalCount: 4,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	clusterEvents, err := NewClustersAPI(ctx, client).Events(EventsRequest{ClusterID: "abc", MaxItems: 3, Limit: 2})
	require.NoError(t, err)
	assert.Equal(t, len(clusterEvents), 2)
	assert.Equal(t, clusterEvents[0].ClusterID, "abc")
	assert.Equal(t, clusterEvents[0].Timestamp, int64(123))
	assert.Equal(t, clusterEvents[0].Type, EvTypeRunning)
	assert.Equal(t, clusterEvents[0].Details.CurrentNumWorkers, int32(2))
	assert.Equal(t, clusterEvents[0].Details.TargetNumWorkers, int32(2))
	assert.Equal(t, clusterEvents[1].Timestamp, int64(124))
}

func TestEventsEmptyResult(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: EventsRequest{
				ClusterID: "abc",
				Limit:     2,
			},
			Response: EventsResponse{
				Events:     []ClusterEvent{},
				TotalCount: 0,
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	clusterEvents, err := NewClustersAPI(ctx, client).Events(EventsRequest{ClusterID: "abc", MaxItems: 3, Limit: 2})
	require.NoError(t, err)
	assert.Equal(t, len(clusterEvents), 0)
}

func TestClusterState_CanReach(t *testing.T) {
	tests := []struct {
		from ClusterState
		to   ClusterState
		want bool
	}{
		{ClusterStatePending, ClusterStatePending, true},
		{ClusterStatePending, ClusterStateRunning, true},
		{ClusterStatePending, ClusterStateRestarting, true},
		{ClusterStatePending, ClusterStateResizing, true},
		{ClusterStatePending, ClusterStateTerminating, true},
		{ClusterStatePending, ClusterStateTerminated, true},
		{ClusterStatePending, ClusterStateError, false},
		{ClusterStatePending, ClusterStateUnknown, false},

		{ClusterStateRunning, ClusterStatePending, false},
		{ClusterStateRunning, ClusterStateRunning, true},
		{ClusterStateRunning, ClusterStateRestarting, true},
		{ClusterStateRunning, ClusterStateResizing, true},
		{ClusterStateRunning, ClusterStateTerminating, true},
		{ClusterStateRunning, ClusterStateTerminated, true},
		{ClusterStateRunning, ClusterStateError, false},
		{ClusterStateRunning, ClusterStateUnknown, false},

		{ClusterStateRestarting, ClusterStatePending, false},
		{ClusterStateRestarting, ClusterStateRunning, true},
		{ClusterStateRestarting, ClusterStateRestarting, true},
		{ClusterStateRestarting, ClusterStateResizing, true},
		{ClusterStateRestarting, ClusterStateTerminating, true},
		{ClusterStateRestarting, ClusterStateTerminated, true},
		{ClusterStateRestarting, ClusterStateError, false},
		{ClusterStateRestarting, ClusterStateUnknown, false},

		{ClusterStateResizing, ClusterStatePending, false},
		{ClusterStateResizing, ClusterStateRunning, true},
		{ClusterStateResizing, ClusterStateRestarting, true},
		{ClusterStateResizing, ClusterStateResizing, true},
		{ClusterStateResizing, ClusterStateTerminating, true},
		{ClusterStateResizing, ClusterStateTerminated, true},
		{ClusterStateResizing, ClusterStateError, false},
		{ClusterStateResizing, ClusterStateUnknown, false},

		{ClusterStateTerminating, ClusterStatePending, false},
		{ClusterStateTerminating, ClusterStateRunning, false},
		{ClusterStateTerminating, ClusterStateRestarting, false},
		{ClusterStateTerminating, ClusterStateResizing, false},
		{ClusterStateTerminating, ClusterStateTerminating, true},
		{ClusterStateTerminating, ClusterStateTerminated, true},
		{ClusterStateTerminating, ClusterStateError, false},
		{ClusterStateTerminating, ClusterStateUnknown, false},

		{ClusterStateTerminated, ClusterStatePending, false},
		{ClusterStateTerminated, ClusterStateRunning, false},
		{ClusterStateTerminated, ClusterStateRestarting, false},
		{ClusterStateTerminated, ClusterStateResizing, false},
		{ClusterStateTerminated, ClusterStateTerminating, false},
		{ClusterStateTerminated, ClusterStateTerminated, true},
		{ClusterStateTerminated, ClusterStateError, false},
		{ClusterStateTerminated, ClusterStateUnknown, false},

		{ClusterStateError, ClusterStatePending, false},
		{ClusterStateError, ClusterStateRunning, false},
		{ClusterStateError, ClusterStateRestarting, false},
		{ClusterStateError, ClusterStateResizing, false},
		{ClusterStateError, ClusterStateTerminating, false},
		{ClusterStateError, ClusterStateTerminated, false},
		{ClusterStateError, ClusterStateError, true},
		{ClusterStateError, ClusterStateUnknown, false},

		{ClusterStateUnknown, ClusterStatePending, false},
		{ClusterStateUnknown, ClusterStateRunning, false},
		{ClusterStateUnknown, ClusterStateRestarting, false},
		{ClusterStateUnknown, ClusterStateResizing, false},
		{ClusterStateUnknown, ClusterStateTerminating, false},
		{ClusterStateUnknown, ClusterStateTerminated, false},
		{ClusterStateUnknown, ClusterStateError, false},
		{ClusterStateUnknown, ClusterStateUnknown, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s to %s", tt.from, tt.to), func(t *testing.T) {
			if got := tt.from.CanReach(tt.to); got != tt.want {
				t.Errorf("ClusterState.CanReach() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFailureOfPermanentDeleteOnCreateFailure(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/create",
			Response: Cluster{
				ClusterID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Status:   418,
			Response: apierr.APIError{
				ErrorCode: "TEST",
				Message:   "nothing",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/delete",
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
			Status:   418,
			Response: apierr.APIError{
				ErrorCode: "TEST",
				Message:   "You should unpin the cluster first",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/unpin",
			Status:   418,
			Response: apierr.NotFound("missing"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewClustersAPI(ctx, client)
		_, err := a.Create(Cluster{})
		assert.EqualError(t, err, "missing")
	})
}

func TestWrapMissingClusterError(t *testing.T) {
	assert.EqualError(t, wrapMissingClusterError(fmt.Errorf("x"), "abc"), "x")
	assert.EqualError(t, wrapMissingClusterError(&apierr.APIError{
		Message: "Cluster abc does not exist",
	}, "abc"), "Cluster abc does not exist")
}

func TestExpiredClusterAssumedAsRemoved(t *testing.T) {
	err := wrapMissingClusterError(&apierr.APIError{
		ErrorCode: "INVALID_STATE",
		Message:   "Cannot access cluster X that was terminated or unpinned more than Y days ago.",
	}, "X")
	var ae *apierr.APIError
	assert.True(t, errors.As(err, &ae))
	assert.Equal(t, 404, ae.StatusCode)
}
