package compute

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test interface compliance
var _ common.CommandExecutor = (*CommandsAPI)(nil)

func TestSomeCommands(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/get?cluster_id=abc",
			Response: ClusterInfo{
				ClusterID:              "abc",
				NumWorkers:             100,
				ClusterName:            "Shared Autoscaling",
				SparkVersion:           "7.1-scala12",
				NodeTypeID:             "i3.xlarge",
				AutoterminationMinutes: 15,
				State:                  ClusterStateRunning,
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			ExpectedRequest: map[string]interface{}{
				"clusterId": "abc",
				"language":  "python",
			},
			Response: Command{
				ID: "123",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/1.2/contexts/status?clusterId=abc&contextId=123",
			ReuseRequest: true,
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			ExpectedRequest: genericCommandRequest{
				Language:  "python",
				ClusterID: "abc",
				ContextID: "123",
				Command:   "print(\"done\")\n",
			},
			Response: Command{
				ID: "234",
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/1.2/commands/status?clusterId=abc&commandId=234&contextId=123",
			Response: Command{
				Status: "Finished",
				Results: &CommandResults{
					ResultType: "text",
					Data:       "done",
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
			ExpectedRequest: genericCommandRequest{
				ClusterID: "abc",
				ContextID: "123",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)
	ctx := context.Background()
	commands := NewCommandsAPI(ctx, client)

	result, err := commands.Execute("abc", "python", `print("done")`)
	require.NoError(t, err)
	assert.Equal(t, "done", result)
}

func TestAccContext(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.CommonEnvironmentClient()
	clusterInfo := NewTinyClusterInCommonPoolPossiblyReused()
	clusterID := clusterInfo.ClusterID
	ctx := context.Background()
	c := NewCommandsAPI(ctx, client)

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
