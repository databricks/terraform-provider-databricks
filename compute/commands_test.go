package compute

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test interface compliance
var _ common.CommandExecutor = (*CommandsAPI)(nil)

func TestAccContext(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.CommonEnvironmentClient()
	clusterInfo := NewTinyClusterInCommonPoolPossiblyReused()
	clusterID := clusterInfo.ClusterID
	c := NewCommandsAPI(client)

	result, err := c.Execute(clusterID, "python", `print('hello world')`)
	require.NoError(t, err)
	assert.Equal(t, "hello world", result)

	// exceptions are regexed away for readability
	result, err = c.Execute(clusterID, "python", `raise Exception("Not Found")`)
	qa.AssertErrorStartsWith(t, err, "Not Found")
	assert.Equal(t, "", result)

	// but errors are not
	result, err = c.Execute(clusterID, "python", `raise KeyError("foo")`)
	qa.AssertErrorStartsWith(t, err, "KeyError: 'foo'")
	assert.Equal(t, "", result)

	// so it is more clear to read and debug
	result, err = c.Execute(clusterID, "python", `return 'hello world'`)
	qa.AssertErrorStartsWith(t, err, "SyntaxError: 'return' outside function")
	assert.Equal(t, "", result)

	result, err = c.Execute(clusterID, "python", `"Hello World!"`)
	assert.NoError(t, err)
	assert.Equal(t, "'Hello World!'", result)

	result, err = c.Execute(clusterID, "python", `
		print("Hello World!")
		dbutils.notebook.exit("success")`)
	assert.NoError(t, err)
	assert.Equal(t, "success", result)
}
