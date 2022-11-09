package acceptance

import (
	"context"
	"reflect"
	"testing"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccListClustersIntegration(t *testing.T) {
	qa.RequireAnyCloudEnv(t)
	t.Parallel()
	client := common.CommonEnvironmentClient()
	ctx := context.Background()
	clustersAPI := clusters.NewClustersAPI(ctx, client)
	randomName := qa.RandomName()

	cluster := clusters.Cluster{
		NumWorkers:  1,
		ClusterName: "Terraform Integration Test " + randomName,
		SparkVersion: clustersAPI.LatestSparkVersionOrDefault(
			clusters.SparkVersionRequest{
				Latest:          true,
			}),
		InstancePoolID:         qa.GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
		IdempotencyToken:       "acc-list-" + randomName,
		AutoterminationMinutes: 15,
	}
	clusterReadInfo, err := clustersAPI.Create(cluster)
	require.NoError(t, err, err)
	assert.True(t, clusterReadInfo.NumWorkers == cluster.NumWorkers)
	assert.True(t, clusterReadInfo.ClusterName == cluster.ClusterName)
	assert.True(t, reflect.DeepEqual(clusterReadInfo.SparkEnvVars, cluster.SparkEnvVars))
	assert.True(t, clusterReadInfo.SparkVersion == cluster.SparkVersion)
	assert.True(t, clusterReadInfo.AutoterminationMinutes == cluster.AutoterminationMinutes)
	assert.True(t, clusterReadInfo.State == clusters.ClusterStateRunning)

	defer func() {
		err = clustersAPI.Terminate(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		clusterReadInfo, err = clustersAPI.Get(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
		assert.True(t, clusterReadInfo.State == clusters.ClusterStateTerminated)

		err = clustersAPI.Unpin(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		err = clustersAPI.PermanentDelete(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	err = clustersAPI.Pin(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)

	clusterReadInfo, err = clustersAPI.Get(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.State == clusters.ClusterStateRunning)
}

func TestAccListClustersResizeIntegrationTest(t *testing.T) {
	qa.RequireAnyCloudEnv(t)
	t.Parallel()
	client := common.CommonEnvironmentClient()
	ctx := context.Background()
	clustersAPI := clusters.NewClustersAPI(ctx, client)
	randomName := qa.RandomName()

	cluster := clusters.Cluster{
		NumWorkers:  1,
		ClusterName: "Terraform Integration Test " + randomName,
		SparkVersion: clustersAPI.LatestSparkVersionOrDefault(
			clusters.SparkVersionRequest{
				Latest:          true,
			}),
		InstancePoolID:         qa.GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
		IdempotencyToken:       "acc-list-" + randomName,
		AutoterminationMinutes: 15,
	}

	clusterReadInfo, err := clustersAPI.Create(cluster)
	require.NoError(t, err, err)
	assert.True(t, clusterReadInfo.NumWorkers == cluster.NumWorkers)
	assert.True(t, clusterReadInfo.ClusterName == cluster.ClusterName)
	assert.True(t, reflect.DeepEqual(clusterReadInfo.SparkEnvVars, cluster.SparkEnvVars))
	assert.True(t, clusterReadInfo.SparkVersion == cluster.SparkVersion)
	assert.True(t, clusterReadInfo.AutoterminationMinutes == cluster.AutoterminationMinutes)
	assert.True(t, clusterReadInfo.State == clusters.ClusterStateRunning)

	// Resize num workers
	clusterReadInfo, err = clustersAPI.Resize(clusters.ResizeRequest{ClusterID: clusterReadInfo.ClusterID, NumWorkers: 2})
	require.NoError(t, err, err)
	assert.True(t, clusterReadInfo.NumWorkers == 2)
	assert.True(t, clusterReadInfo.State == clusters.ClusterStateRunning)

	// Resize cluster to become autoscaling
	clusterReadInfo, err = clustersAPI.Resize(clusters.ResizeRequest{ClusterID: clusterReadInfo.ClusterID, AutoScale: &clusters.AutoScale{
		MinWorkers: 1,
		MaxWorkers: 2,
	}})
	require.NoError(t, err, err)
	assert.True(t, clusterReadInfo.AutoScale.MinWorkers == 1)
	assert.True(t, clusterReadInfo.AutoScale.MaxWorkers == 2)
	assert.True(t, clusterReadInfo.State == clusters.ClusterStateRunning)

	// cleanup
	err = clustersAPI.Terminate(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)

	clusterReadInfo, err = clustersAPI.Get(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.State == clusters.ClusterStateTerminated)

	err = clustersAPI.Unpin(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)

	err = clustersAPI.PermanentDelete(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)
}
