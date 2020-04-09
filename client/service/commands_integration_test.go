package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()
	clusterId := "0408-231224-veeps362"
	context, err := client.Commands().createContext("python", clusterId)
	assert.NoError(t, err, err)
	t.Log(context)

	defer func() {
		err := client.Commands().deleteContext(context, clusterId)
		assert.NoError(t, err, err)
	}()

	err = client.Commands().waitForContextReady(context, clusterId, 1, 1)
	assert.NoError(t, err, err)

	status, err := client.Commands().getContext(context, clusterId)
	assert.NoError(t, err, err)
	t.Log(status)

	commandId, err := client.Commands().createCommand(context, clusterId, "python", "print('hello world')")
	assert.NoError(t, err, err)
	t.Log(commandId)

	err = client.Commands().waitForCommandFinished(commandId, context, clusterId, 5, 20)
	assert.NoError(t, err, err)

	resp, err := client.Commands().getCommand(commandId, context, clusterId)
	assert.NoError(t, err, err)
	t.Log(resp.Results.Data)
}

func TestListDirectory(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()
	status, err := client.DBFS().List("/mnt/sri_test_mount_test123/", false)
	assert.NoError(t, err, err)
	t.Log(status)
}

func TestMnt(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()
	clusterId := "0408-231224-veeps362"
	langauge := "python"
	command := `AWS_BUCKET_NAME = "dredge-notebook-data-123"
MOUNT_NAME = "sri_test_mount_test123"
dbutils.fs.mount("s3a://%s" % AWS_BUCKET_NAME, "/mnt/%s" % MOUNT_NAME)
dbutils.fs.refreshMounts()
`
	commandResp, err := client.Commands().Execute(clusterId, langauge, command)
	assert.NoError(t, err, err)
	t.Log(commandResp.Results.ResultType) //error on failure, table on success
	t.Log(commandResp.Results.Data)
}

func TestUnmount(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()
	clusterId := "0408-231224-veeps362"
	langauge := "python"
	command := `AWS_BUCKET_NAME = "dredge-notebook-data-123"
MOUNT_NAME = "sri_test_mount_test123"
#dbutils.fs.unmount("/mnt/%s" % MOUNT_NAME)
dbutils.fs.refreshMounts()
dbutils.notebook.exit("success")
`
	commandResp, err := client.Commands().Execute(clusterId, langauge, command)
	assert.NoError(t, err, err)
	t.Log(commandResp.Results.ResultType) //error on failure, table/text on success
	t.Log(commandResp.Results.Data)
}
