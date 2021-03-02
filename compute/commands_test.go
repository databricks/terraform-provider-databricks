package compute

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test interface compliance
var _ common.CommandExecutor = (*CommandsAPI)(nil)

func commonFixtureWithStatusResponse(response Command) []qa.HTTPFixture {
	return []qa.HTTPFixture{
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
			Response:     response,
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
			ExpectedRequest: genericCommandRequest{
				ClusterID: "abc",
				ContextID: "123",
			},
		},
	}
}

func TestCommandWithExecutionError(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &common.CommandResults{
			ResultType: "error",
			Cause: `
---
ExecutionError: An error occurred
StatusCode=400
StatusDescription=BadRequest
---
			`,
		},
	}))
	defer server.Close()
	require.NoError(t, err)
	ctx := context.Background()
	commands := NewCommandsAPI(ctx, client)

	result := commands.Execute("abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "An error occurred\nStatusCode=400\nStatusDescription=BadRequest", result.Error())
}

func TestCommandWithEmptyErrorMessageUsesSummary(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &common.CommandResults{
			ResultType: "error",
			Cause: `
---
ErrorCode=
ErrorMessage=
    other text
---
			`,
			Summary: "Proper error",
		},
	}))
	defer server.Close()
	require.NoError(t, err)
	ctx := context.Background()
	commands := NewCommandsAPI(ctx, client)

	result := commands.Execute("abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "Proper error", result.Error())
}

func TestCommandWithErrorMessage(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &common.CommandResults{
			ResultType: "error",
			Cause: `
---
ErrorCode=
ErrorMessage=An error occurred
---
			`,
		},
	}))
	defer server.Close()
	require.NoError(t, err)
	ctx := context.Background()
	commands := NewCommandsAPI(ctx, client)

	result := commands.Execute("abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "An error occurred", result.Error())
}

func TestCommandWithExceptionMessage(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &common.CommandResults{
			ResultType: "error",
			Summary:    "Exception: An error occurred",
		},
	}))
	defer server.Close()
	require.NoError(t, err)
	ctx := context.Background()
	commands := NewCommandsAPI(ctx, client)

	result := commands.Execute("abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "An error occurred", result.Error())
}

func TestSomeCommands(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &common.CommandResults{
			ResultType: "text",
			Data:       "done",
		},
	}))
	defer server.Close()
	require.NoError(t, err)
	ctx := context.Background()
	commands := NewCommandsAPI(ctx, client)

	result := commands.Execute("abc", "python", `print("done")`)
	assert.Equal(t, false, result.Failed())
	assert.Equal(t, "done", result.Text())
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

	result := c.Execute(clusterID, "python", `print('hello world')`)
	assert.Equal(t, "hello world", result.Text())

	// exceptions are regexed away for readability
	result = c.Execute(clusterID, "python", `raise Exception("Not Found")`)
	qa.AssertErrorStartsWith(t, result.Err(), "Not Found")

	// but errors are not
	result = c.Execute(clusterID, "python", `raise KeyError("foo")`)
	qa.AssertErrorStartsWith(t, result.Err(), "KeyError: 'foo'")

	// so it is more clear to read and debug
	result = c.Execute(clusterID, "python", `return 'hello world'`)
	qa.AssertErrorStartsWith(t, result.Err(), "SyntaxError: 'return' outside function")

	result = c.Execute(clusterID, "python", `"Hello World!"`)
	assert.Equal(t, "'Hello World!'", result.Text())

	result = c.Execute(clusterID, "python", `
 		print("Hello World!")
 		dbutils.notebook.exit("success")`)
	assert.Equal(t, "success", result.Text())
}
