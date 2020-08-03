package service

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test interface compliance
var _ CommandExecutor = (*CommandsAPI)(nil)

func TestAccContext(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := CommonEnvironmentClient()
	clusterInfo := NewTinyClusterInCommonPoolPossiblyReused()
	clusterID := clusterInfo.ClusterID
	c := client.Commands()

	result, err := c.Execute(clusterID, "python", `print('hello world')`)
	require.NoError(t, err)
	assert.Equal(t, "hello world", result)

	// exceptions are regexed away for readability
	result, err = c.Execute(clusterID, "python", `raise Exception("Not Found")`)
	assertErrorStartsWith(t, err, "Not Found")
	assert.Equal(t, "", result)

	// but errors are not
	result, err = c.Execute(clusterID, "python", `raise KeyError("foo")`)
	assertErrorStartsWith(t, err, "KeyError: 'foo'")
	assert.Equal(t, "", result)

	// so it is more clear to read and debug
	result, err = c.Execute(clusterID, "python", `return 'hello world'`)
	assertErrorStartsWith(t, err, "SyntaxError: 'return' outside function")
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
	assert.True(t, strings.Contains(err.Error(), "/mnt/qwertyui does not exist"), err.Error())
	assert.Equal(t, "", result)
}
