package compute

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTinyClusterInCommonPoolPossiblyReused(t *testing.T) {
	defer common.CleanupEnvironment()()
	common.ResetCommonEnvironmentClient()
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Response: InstancePoolList{
				InstancePools: []InstancePoolAndStats{},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-pools/create",
			ExpectedRequest: InstancePool{
				AwsAttributes: &InstancePoolAwsAttributes{
					Availability: "SPOT",
				},
				NodeTypeID:                         "m4.large",
				IdleInstanceAutoTerminationMinutes: 15,
				DiskSpec: &InstancePoolDiskSpec{
					DiskCount: 1,
					DiskSize:  32,
					DiskType: &InstancePoolDiskType{
						EbsVolumeType: "GENERAL_PURPOSE_SSD",
					},
				},
				PreloadedSparkVersions: []string{"6.6.x-scala2.11"},
				InstancePoolName:       "Terraform Integration Test by test",
				MaxCapacity:            10,
			},
			Response: InstancePoolAndStats{
				InstancePoolID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Response: map[string]interface{}{},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/create",
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
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Response: InstancePoolList{
				InstancePools: []InstancePoolAndStats{
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
	// trick common env client
	os.Setenv("DATABRICKS_HOST", client.Host)
	os.Setenv("DATABRICKS_TOKEN", client.Token)
	os.Setenv("USER", "test")

	c, err := NewTinyClusterInCommonPool()
	require.NoError(t, err)
	assert.NotNil(t, c)
}
