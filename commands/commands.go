package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Command is the struct that contains what the 1.2 api returns for the commands api
type Command struct {
	ID      string                 `json:"id,omitempty"`
	Status  string                 `json:"status,omitempty"`
	Results *common.CommandResults `json:"results,omitempty"`
}

// NewCommandsAPI creates CommandsAPI instance from provider meta
func NewCommandsAPI(ctx context.Context, m any) CommandsAPI {
	return CommandsAPI{
		client:  m.(*common.DatabricksClient),
		context: context.WithValue(ctx, common.Api, common.API_1_2),
	}
}

// CommandsAPI exposes the Context & Commands API
type CommandsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Execute creates a spark context and executes a command and then closes context
// Any leading whitespace is trimmed
func (a CommandsAPI) Execute(clusterID, language, commandStr string) common.CommandResults {
	// this is the place, where API version propagation through context looks strange
	ctx := context.WithValue(a.context, common.Api, common.API_2_0)
	cluster, err := clusters.NewClustersAPI(ctx, a.client).Get(clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	if !cluster.IsRunningOrResizing() {
		return common.CommandResults{
			ResultType: "error",
			Summary:    fmt.Sprintf("Cluster %s has to be running or resizing, but is %s", clusterID, cluster.State),
		}
	}
	commandStr = TrimLeadingWhitespace(commandStr)
	log.Printf("[INFO] Executing %s command on %s:\n%s", language, clusterID, commandStr)
	context, err := a.createContext(language, clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	err = a.waitForContextReady(context, clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	commandID, err := a.createCommand(context, clusterID, language, commandStr)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	// TODO: merge getCommand and waitForCommandFinished to "waitForCommandResults"
	err = a.waitForCommandFinished(commandID, context, clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	command, err := a.getCommand(commandID, context, clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	err = a.deleteContext(context, clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	if command.Results == nil {
		log.Printf("[ERROR] Command has no results: %#v", command)
		return common.CommandResults{
			ResultType: "error",
			Summary:    "Command has no results",
		}
	}
	return *command.Results
}

type genericCommandRequest struct {
	CommandID string `json:"commandId,omitempty" url:"commandId,omitempty"`
	Language  string `json:"language,omitempty" url:"language,omitempty"`
	ClusterID string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	ContextID string `json:"contextId,omitempty" url:"contextId,omitempty"`
	Command   string `json:"command,omitempty" url:"command,omitempty"`
}

func (a CommandsAPI) createCommand(contextID, clusterID, language, commandStr string) (string, error) {
	var command Command
	err := a.client.Post(a.context, "/commands/execute", genericCommandRequest{
		Language:  language,
		ClusterID: clusterID,
		ContextID: contextID,
		Command:   commandStr,
	}, &command)
	return command.ID, err
}

func (a CommandsAPI) getCommand(commandID, contextID, clusterID string) (Command, error) {
	var commandResp Command
	err := a.client.Get(a.context, "/commands/status", genericCommandRequest{
		CommandID: commandID,
		ContextID: contextID,
		ClusterID: clusterID,
	}, &commandResp)
	return commandResp, err
}

func (a CommandsAPI) deleteContext(contextID, clusterID string) error {
	return a.client.Post(a.context, "/contexts/destroy", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, nil)
}

func (a CommandsAPI) getContext(contextID, clusterID string) (string, error) {
	var contextStatus Command // internal hack, yes
	err := a.client.Get(a.context, "/contexts/status", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, &contextStatus)
	return contextStatus.Status, err
}

func (a CommandsAPI) createContext(language, clusterID string) (string, error) {
	var context Command // internal hack, yes
	err := a.client.Post(a.context, "/contexts/create", genericCommandRequest{
		Language:  language,
		ClusterID: clusterID,
	}, &context)
	return context.ID, err
}

func (a CommandsAPI) waitForCommandFinished(commandID, contextID, clusterID string) error {
	return resource.RetryContext(a.context, 10*time.Minute, func() *resource.RetryError {
		commandInfo, err := a.getCommand(commandID, contextID, clusterID)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		switch commandInfo.Status {
		case "Cancelling", "Cancelled", "Error":
			return resource.NonRetryableError(fmt.Errorf("Command cannot finish: %s", commandInfo.Status))
		case "Finished":
			return nil
		}
		log.Printf("[DEBUG] Command is in %s state", commandInfo.Status)
		return resource.RetryableError(errors.New(commandInfo.Status))
	})
}

func (a CommandsAPI) waitForContextReady(contextID, clusterID string) error {
	return resource.RetryContext(a.context, 10*time.Minute, func() *resource.RetryError {
		status, err := a.getContext(contextID, clusterID)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if status == "Error" {
			return resource.NonRetryableError(errors.New(status))
		}
		if status == "Running" {
			return nil
		}
		return resource.RetryableError(errors.New(status))
	})
}
