package clusters

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartClusterAndGetInfo_Pending(t *testing.T) {
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
	w, err := client.WorkspaceClient()
	if err != nil {
		panic(err)
	}
	clusterInfo, err := StartClusterAndGetInfo(ctx, w, "abc")
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestStartClusterAndGetInfo_Terminating(t *testing.T) {
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
	w, err := client.WorkspaceClient()
	if err != nil {
		panic(err)
	}
	clusterInfo, err := StartClusterAndGetInfo(ctx, w, "abc")
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestStartClusterAndGetInfo_Error(t *testing.T) {
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
	w, err := client.WorkspaceClient()
	if err != nil {
		panic(err)
	}
	clusterInfo, err := StartClusterAndGetInfo(ctx, w, "abc")
	require.NoError(t, err)
	assert.Equal(t, ClusterStateRunning, string(clusterInfo.State))
}

func TestStartClusterAndGetInfo_StartingError(t *testing.T) {
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
	w, err := client.WorkspaceClient()
	if err != nil {
		panic(err)
	}
	_, err = StartClusterAndGetInfo(ctx, w, "abc")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "I am a teapot")
}
