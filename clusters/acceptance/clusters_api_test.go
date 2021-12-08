package acceptance

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/compute"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccListClustersIntegration(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

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
				LongTermSupport: true,
			}),
		InstancePoolID:         compute.CommonInstancePoolID(),
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
