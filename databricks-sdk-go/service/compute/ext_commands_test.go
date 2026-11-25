package compute

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/qa"
)

func commonFixtureWithStatusResponse(response CommandStatusResponse) qa.HTTPFixtures {
	return []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.1/clusters/get?cluster_id=abc",
			Response: ClusterDetails{
				ClusterId:              "abc",
				NumWorkers:             100,
				ClusterName:            "Shared Autoscaling",
				SparkVersion:           "7.1-scala12",
				NodeTypeId:             "i3.xlarge",
				AutoterminationMinutes: 15,
				State:                  StateRunning,
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			ExpectedRequest: map[string]interface{}{
				"clusterId": "abc",
				"language":  "python",
			},
			Response: Created{
				Id: "123",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/1.2/contexts/status?clusterId=abc&contextId=123",
			ReuseRequest: true,
			Response: ContextStatusResponse{
				Id:     "123",
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			ExpectedRequest: Command{
				Language:  "python",
				ClusterId: "abc",
				ContextId: "123",
				Command:   "print(\"done\")\n",
			},
			Response: Created{
				Id: "234",
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
			ExpectedRequest: DestroyContext{
				ClusterId: "abc",
				ContextId: "123",
			},
		},
	}
}

func TestCommandWithExecutionError(t *testing.T) {
	client, server := commonFixtureWithStatusResponse(CommandStatusResponse{
		Status: "Finished",
		Results: &Results{
			ResultType: "error",
			Cause: `
---
ExecutionError: An error occurred
StatusCode=400
StatusDescription=BadRequest
---
			`,
		},
	}).Client(t)
	defer server.Close()
	ctx := context.Background()
	commands := NewCommandExecutor(client)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "An error occurred\nStatusCode=400\nStatusDescription=BadRequest", result.Error())
}

func TestCommandWithEmptyErrorMessageUsesSummary(t *testing.T) {
	client, server := commonFixtureWithStatusResponse(CommandStatusResponse{
		Status: "Finished",
		Results: &Results{
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
	}).Client(t)
	defer server.Close()
	ctx := context.Background()
	commands := NewCommandExecutor(client)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "Proper error", result.Error())
}

func TestCommandWithErrorMessage(t *testing.T) {
	client, server := commonFixtureWithStatusResponse(CommandStatusResponse{
		Status: "Finished",
		Results: &Results{
			ResultType: "error",
			Cause: `
---
ErrorCode=
ErrorMessage=An error occurred
---
			`,
		},
	}).Client(t)
	defer server.Close()
	ctx := context.Background()
	commands := NewCommandExecutor(client)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "An error occurred", result.Error())
}

func TestCommandWithExceptionMessage(t *testing.T) {
	client, server := commonFixtureWithStatusResponse(CommandStatusResponse{
		Status: "Finished",
		Results: &Results{
			ResultType: "error",
			Summary:    "Exception: An error occurred",
		},
	}).Client(t)
	defer server.Close()
	ctx := context.Background()
	commands := NewCommandExecutor(client)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "An error occurred", result.Error())
}

func TestSomeCommands(t *testing.T) {
	client, server := commonFixtureWithStatusResponse(CommandStatusResponse{
		Status: "Finished",
		Results: &Results{
			ResultType: "text",
			Data:       "done",
		},
	}).Client(t)
	defer server.Close()
	ctx := context.Background()
	commands := NewCommandExecutor(client)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, false, result.Failed())
	assert.Equal(t, "done", result.Text())
}

func TestCommandsAPIExecute_FailGettingCluster(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=abc",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := NewCommandExecutor(client)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_StoppedCluster(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=abc",
			Response: ClusterDetails{
				State: "TERMINATED",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := NewCommandExecutor(client)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Cluster abc has to be running or resizing, but is TERMINATED")
	})
}

func TestCommandsAPIExecute_FailToCreateContext(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=abc",
			Response: ClusterDetails{
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
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := NewCommandExecutor(client)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToWaitForContext(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=abc",
			Response: ClusterDetails{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Created{
				Id: "abc",
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
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := NewCommandExecutor(client)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToCreateCommand(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=abc",
			Response: ClusterDetails{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Created{
				Id: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: CommandStatusResponse{
				Id:     "abc",
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
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
			ExpectedRequest: DestroyContext{
				ClusterId: "abc",
				ContextId: "abc",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := NewCommandExecutor(client)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToWaitForCommand(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=abc",
			Response: ClusterDetails{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Created{
				Id: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: CommandStatusResponse{
				Id:     "abc",
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Created{
				Id: "abc",
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
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
			ExpectedRequest: DestroyContext{
				ClusterId: "abc",
				ContextId: "abc",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := NewCommandExecutor(client)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToGetCommand(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=abc",
			Response: ClusterDetails{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Created{
				Id: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: CommandStatusResponse{
				Id:     "abc",
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Created{
				Id: "abc",
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
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
			ExpectedRequest: DestroyContext{
				ClusterId: "abc",
				ContextId: "abc",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := NewCommandExecutor(client)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_NoResults(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=abc",
			Response: ClusterDetails{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Created{
				Id: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: CommandStatusResponse{
				Id:     "abc",
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Created{
				Id: "abc",
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Response: CommandStatusResponse{
				Id:     "abc",
				Status: "Finished",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := NewCommandExecutor(client)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Command has no results")
	})
}
