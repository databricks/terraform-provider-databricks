package service

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestLibraryCreate(t *testing.T) {
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
	defer func() {
		err := client.Clusters().PermanentDelete(clusterInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	clusterID := clusterInfo.ClusterID

	err = client.Clusters().WaitForClusterRunning(clusterID, 10, 20)
	assert.NoError(t, err, err)

	libraries := []model.Library{
		{
			Pypi: &model.PyPi{
				Package: "networkx",
			},
		},
		{
			Maven: &model.Maven{
				Coordinates: "com.crealytics:spark-excel_2.12:0.13.1",
			},
		},
	}

	err = client.Libraries().Create(clusterID, libraries)
	assert.NoError(t, err, err)

	defer func() {
		err = client.Libraries().Delete(clusterID, libraries)
		assert.NoError(t, err, err)
	}()

	libraryStatusList, err := client.Libraries().List(clusterID)
	assert.NoError(t, err, err)
	assert.Equal(t, len(libraryStatusList), len(libraries))
}
