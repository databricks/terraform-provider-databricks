package service

import (
	"reflect"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestListClustersIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()

	cluster := model.Cluster{
		NumWorkers:  1,
		ClusterName: "my-cluster",
		SparkEnvVars: map[string]string{
			"PYSPARK_PYTHON": "/databricks/python3/bin/python3",
		},
		AwsAttributes: &model.AwsAttributes{
			EbsVolumeType:  model.EbsVolumeTypeGeneralPurposeSsd,
			EbsVolumeCount: 1,
			EbsVolumeSize:  32,
		},
		SparkVersion:           "6.2.x-scala2.11",
		NodeTypeID:             GetCloudInstanceType(client),
		DriverNodeTypeID:       GetCloudInstanceType(client),
		IdempotencyToken:       "my-cluster",
		AutoterminationMinutes: 20,
	}

	clusterInfo, err := client.Clusters().Create(cluster)
	assert.NoError(t, err, err)

	clusterReadInfo, err := client.Clusters().Get(clusterInfo.ClusterID)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.NumWorkers == cluster.NumWorkers)
	assert.True(t, clusterReadInfo.ClusterName == cluster.ClusterName)
	assert.True(t, reflect.DeepEqual(clusterReadInfo.SparkEnvVars, cluster.SparkEnvVars))
	assert.True(t, clusterReadInfo.SparkVersion == cluster.SparkVersion)
	assert.True(t, clusterReadInfo.NodeTypeID == cluster.NodeTypeID)
	assert.True(t, clusterReadInfo.DriverNodeTypeID == cluster.DriverNodeTypeID)
	assert.True(t, clusterReadInfo.AutoterminationMinutes == cluster.AutoterminationMinutes)

	defer func() {
		err = client.Clusters().Delete(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		err = client.Clusters().WaitForClusterTerminated(clusterReadInfo.ClusterID, 10, 20)
		assert.NoError(t, err, err)

		clusterReadInfo, err = client.Clusters().Get(clusterInfo.ClusterID)
		assert.NoError(t, err, err)
		assert.True(t, clusterReadInfo.State == model.ClusterStateTerminated)

		err = client.Clusters().Unpin(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		err = client.Clusters().PermanentDelete(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	err = client.Clusters().Pin(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)

	err = client.Clusters().WaitForClusterRunning(clusterReadInfo.ClusterID, 10, 20)
	assert.NoError(t, err, err)

	clusterReadInfo, err = client.Clusters().Get(clusterInfo.ClusterID)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.State == model.ClusterStateRunning)
}
