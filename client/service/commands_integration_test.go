package service

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext(t *testing.T) {
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

	clusterId := clusterInfo.ClusterID

	err = client.Clusters().WaitForClusterRunning(clusterId, 10, 20)
	assert.NoError(t, err, err)

	context, err := client.Commands().createContext("python", clusterId)
	assert.NoError(t, err, err)
	t.Log(context)

	err = client.Commands().waitForContextReady(context, clusterId, 1, 1)
	assert.NoError(t, err, err)

	status, err := client.Commands().getContext(context, clusterId)
	assert.NoError(t, err, err)
	assert.True(t, status == "Running")
	t.Log(status)

	commandId, err := client.Commands().createCommand(context, clusterId, "python", "print('hello world')")
	assert.NoError(t, err, err)

	err = client.Commands().waitForCommandFinished(commandId, context, clusterId, 5, 20)
	assert.NoError(t, err, err)

	resp, err := client.Commands().getCommand(commandId, context, clusterId)
	assert.NoError(t, err, err)
	assert.NotNil(t, resp.Results.Data)

	// Testing the public api Execute
	command, err := client.Commands().Execute(clusterId, "python", "print('hello world')")
	assert.NoError(t, err, err)
	assert.NotNil(t, command.Results.Data)
}
