package service

import (
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"github.com/stretchr/testify/assert"
	"testing"
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

	defer func() {
		err = client.Clusters().WaitForClusterTerminated(clusterReadInfo.ClusterID, 10, 20)
		assert.NoError(t, err, err)

		err = client.Clusters().Unpin(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		err = client.Clusters().PermanentDelete(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	err = client.Clusters().Pin(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)

	err = client.Clusters().WaitForClusterRunning(clusterReadInfo.ClusterID, 10, 20)
	assert.NoError(t, err, err)

	err = client.Clusters().Delete(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)

}
