package compute

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type CommandExecutorV2 struct {
	clustersAPI  *ClustersAPI
	executionAPI *CommandExecutionAPI
	language     Language
	clusterID    string
	contextID    string
}

type commandExecutionAPIUtilities interface {
	Start(ctx context.Context, clusterID string, language Language) (*CommandExecutorV2, error)
}

// Start the command execution context on a cluster and ensure it transitions to a running state
func (a *CommandExecutionAPI) Start(ctx context.Context, clusterID string, language Language) (*CommandExecutorV2, error) {
	// re-initializing clusters API here, so that we have less IF's in the generator for WorkspaceClient
	clustersAPI := NewClusters(a.commandExecutionImpl.client)
	err := clustersAPI.EnsureClusterIsRunning(ctx, clusterID)
	if err != nil {
		return nil, err
	}
	wait, err := a.Create(ctx, CreateContext{
		ClusterId: clusterID,
		Language:  language,
	})
	if err != nil {
		return nil, err
	}
	commandContext, err := wait.Get()
	if err != nil {
		return nil, err
	}
	return &CommandExecutorV2{
		clustersAPI:  clustersAPI,
		executionAPI: a,
		language:     language,
		clusterID:    clusterID,
		contextID:    commandContext.Id,
	}, nil
}

func (c *CommandExecutorV2) Destroy(ctx context.Context) error {
	return c.executionAPI.Destroy(ctx, DestroyContext{
		ClusterId: c.clusterID,
		ContextId: c.contextID,
	})
}

// Execute runs given command in the running cluster and context and waits for the results
func (c *CommandExecutorV2) Execute(ctx context.Context, command string) (*Results, error) {
	wait, err := c.executionAPI.Execute(ctx, Command{
		Command:   TrimLeadingWhitespace(command),
		ClusterId: c.clusterID,
		ContextId: c.contextID,
		Language:  c.language,
	})
	if err != nil {
		return nil, err
	}
	status, err := wait.Get()
	if err != nil {
		return nil, err
	}
	return status.Results, nil
}

// CommandExecutor creates a spark context and executes a command and then closes context
type CommandExecutor interface {
	Execute(ctx context.Context, clusterID, language, commandStr string) Results
}

// CommandMock mocks the execution of command
type CommandMock func(commandStr string) Results

func (m CommandMock) Execute(_ context.Context, _, _, commandStr string) Results {
	return m(commandStr)
}

func NewCommandExecutor(client *client.DatabricksClient) CommandExecutor {
	return &CommandsHighLevelAPI{
		clusters:  NewClusters(client),
		execution: NewCommandExecution(client),
	}
}

// CommandsHighLevelAPI exposes more friendly wrapper over command execution
type CommandsHighLevelAPI struct {
	clusters  *ClustersAPI
	execution *CommandExecutionAPI
}

// Execute creates a spark context and executes a command and then closes context
// Any leading whitespace is trimmed
//
// Deprecated: please switch to CommandExecutorV2
func (a *CommandsHighLevelAPI) Execute(ctx context.Context, clusterID, language, commandStr string) Results {
	ctx = useragent.InContext(ctx, "sdk-feature", "command-execution")
	cluster, err := a.clusters.Get(ctx, GetClusterRequest{
		ClusterId: clusterID,
	})
	if err != nil {
		return Results{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	if !cluster.IsRunningOrResizing() {
		return Results{
			ResultType: "error",
			Summary:    fmt.Sprintf("Cluster %s has to be running or resizing, but is %s", clusterID, cluster.State),
		}
	}
	commandStr = TrimLeadingWhitespace(commandStr)
	logger.Debugf(ctx, "Executing %s command on %s:\n%s", language, clusterID, commandStr)
	context, err := a.execution.CreateAndWait(ctx, CreateContext{
		ClusterId: clusterID,
		Language:  Language(language),
	})
	if err != nil {
		return Results{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	defer a.execution.Destroy(ctx, DestroyContext{
		ClusterId: clusterID,
		ContextId: context.Id,
	})
	command, err := a.execution.ExecuteAndWait(ctx, Command{
		ClusterId: clusterID,
		ContextId: context.Id,
		Language:  Language(language),
		Command:   commandStr,
	})
	if err != nil {
		return Results{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	if command.Results == nil {
		logger.Warnf(ctx, "Command has no results: %#v", command)
		return Results{
			ResultType: "error",
			Summary:    "Command has no results",
		}
	}
	return *command.Results
}
