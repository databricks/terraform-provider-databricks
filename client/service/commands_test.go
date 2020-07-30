package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test interface compliance
var _ CommandExecutor = (*CommandsAPI)(nil)

func TestAccContext(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := CommonEnvironmentClient()
	clusterInfo, err := client.Clusters().GetOrCreateRunningCluster("Terraform Integration Test")
	assert.NoError(t, err)

	clusterID := clusterInfo.ClusterID
	c := client.Commands()

	result, err := c.Execute(clusterID, "python", `print('hello world')`)
	assert.NoError(t, err)
	assert.Equal(t, "hello world", result)

	// exceptions are regexed away for readability
	result, err = c.Execute(clusterID, "python", `raise Exception("Not Found")`)
	assert.EqualError(t, err, "Not Found")
	assert.Equal(t, "", result)

	// but errors are not
	result, err = c.Execute(clusterID, "python", `raise KeyError("foo")`)
	assert.EqualError(t, err, "KeyError: 'foo'")
	assert.Equal(t, "", result)

	// so it is more clear to read and debug
	result, err = c.Execute(clusterID, "python", `return 'hello world'`)
	assert.EqualError(t, err, "SyntaxError: 'return' outside function")
	assert.Equal(t, "", result)

	result, err = c.Execute(clusterID, "python", `"Hello World!"`)
	assert.NoError(t, err)
	assert.Equal(t, "'Hello World!'", result)

	result, err = c.Execute(clusterID, "python", `
		print("Hello World!")
		dbutils.notebook.exit("success")`)
	assert.NoError(t, err)
	assert.Equal(t, "success", result)

	result, err = c.Execute(clusterID, "python", `dbutils.fs.ls("/mnt/qwertyui")`)
	assert.EqualError(t, err, "File/1617350743259444/mnt/qwertyui does not exist")
	assert.Equal(t, "", result)
}
