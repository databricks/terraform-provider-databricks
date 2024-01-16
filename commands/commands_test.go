package commands

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/terraform-provider-databricks/qa"
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
			Response: clusters.ClusterInfo{
				ClusterID:              "abc",
				NumWorkers:             100,
				ClusterName:            "Shared Autoscaling",
				SparkVersion:           "7.1-scala12",
				NodeTypeID:             "i3.xlarge",
				AutoterminationMinutes: 15,
				State:                  clusters.ClusterStateRunning,
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			ExpectedRequest: map[string]any{
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

func TestCommandsAPIExecute_FailGettingCluster(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		commands := NewCommandsAPI(ctx, client)
		cr := commands.Execute("abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_StoppedCluster(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "TERMINATED",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		commands := NewCommandsAPI(ctx, client)
		cr := commands.Execute("abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Cluster abc has to be running or resizing, but is TERMINATED")
	})
}

func TestCommandsAPIExecute_FailToCreateContext(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		commands := NewCommandsAPI(ctx, client)
		cr := commands.Execute("abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToWaitForContext(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		commands := NewCommandsAPI(ctx, client)
		cr := commands.Execute("abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToCreateCommand(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		commands := NewCommandsAPI(ctx, client)
		cr := commands.Execute("abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToWaitForCommand(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		commands := NewCommandsAPI(ctx, client)
		cr := commands.Execute("abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToGetCommand(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Response: Command{
				Status: "Finished",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		commands := NewCommandsAPI(ctx, client)
		cr := commands.Execute("abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToDeleteContext(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Response: Command{
				Status: "Finished",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		commands := NewCommandsAPI(ctx, client)
		cr := commands.Execute("abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_NoCommandResults(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Response: Command{
				Status: "Finished",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		commands := NewCommandsAPI(ctx, client)
		cr := commands.Execute("abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Command has no results")
	})
}
