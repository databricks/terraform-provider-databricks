package compute

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
)

// NewCommandsAPI creates CommandsAPI instance from provider meta
func NewCommandsAPI(m interface{}) CommandsAPI {
	return CommandsAPI{client: m.(*common.DatabricksClient)}
}

// CommandsAPI exposes the Context & Commands API
type CommandsAPI struct {
	client *common.DatabricksClient
}

// Execute creates a spark context and executes a command and then closes context
// Any leading whitespace is trimmed
func (a CommandsAPI) Execute(clusterID, language, commandStr string) common.CommandResults {
	cluster, err := NewClustersAPI(a.client).Get(clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary: err.Error(),
		}
	}
	if !cluster.IsRunningOrResizing() {
		return common.CommandResults{
			ResultType: "error",
			Summary: fmt.Sprintf("Cluster %s has to be running or resizing, but is %s", clusterID, cluster.State),
		}
	}
	commandStr = internal.TrimLeadingWhitespace(commandStr)
	log.Printf("[INFO] Executing %s command on %s:\n%s", language, clusterID, commandStr)
	context, err := a.createContext(language, clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary: err.Error(),
		}
	}
	err = a.waitForContextReady(context, clusterID, 1, 10)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary: err.Error(),
		}
	}
	commandID, err := a.createCommand(context, clusterID, language, commandStr)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary: err.Error(),
		}
	}
	err = a.waitForCommandFinished(commandID, context, clusterID, 5, 10)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary: err.Error(),
		}
	}
	command, err := a.getCommand(commandID, context, clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary: err.Error(),
		}
	}
	err = a.deleteContext(context, clusterID)
	if err != nil {
		return common.CommandResults{
			ResultType: "error",
			Summary: err.Error(),
		}
	}
	if command.Results == nil {
		return common.CommandResults{
			ResultType: "error",
			Summary: fmt.Sprintf("Command has no results: %#v", command),
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
	err := a.client.OldAPI("POST", "/commands/execute", genericCommandRequest{
		Language:  language,
		ClusterID: clusterID,
		ContextID: contextID,
		Command:   commandStr,
	}, &command)
	return command.ID, err
}

func (a CommandsAPI) getCommand(commandID, contextID, clusterID string) (Command, error) {
	var commandResp Command
	err := a.client.OldAPI("GET", "/commands/status", genericCommandRequest{
		CommandID: commandID,
		ContextID: contextID,
		ClusterID: clusterID,
	}, &commandResp)
	return commandResp, err
}

func (a CommandsAPI) deleteContext(contextID, clusterID string) error {
	return a.client.OldAPI("POST", "/contexts/destroy", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, nil)
}

func (a CommandsAPI) getContext(contextID, clusterID string) (string, error) {
	var contextStatus Command // internal hack, yes
	err := a.client.OldAPI("GET", "/contexts/status", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, &contextStatus)
	return contextStatus.Status, err
}

func (a CommandsAPI) createContext(language, clusterID string) (string, error) {
	var context Command // internal hack, yes
	err := a.client.OldAPI("POST", "/contexts/create", genericCommandRequest{
		Language:  language,
		ClusterID: clusterID,
	}, &context)
	return context.ID, err
}

func (a CommandsAPI) waitForCommandFinished(commandID, contextID, clusterID string, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		for {
			commandInfo, err := a.getCommand(commandID, contextID, clusterID)
			if err != nil {
				errChan <- err
				return
			}
			if commandInfo.Status == "Finished" {
				errChan <- nil
				return
			} else if commandInfo.Status == "Cancelling" || commandInfo.Status == "Cancelled" || commandInfo.Status == "Error" {
				errChan <- fmt.Errorf("context is in a failure state: %s", commandInfo.Status)
				return
			}
			log.Println(fmt.Sprintf("[DEBUG] Waiting for command to finish, current state is: %s.", commandInfo.Status))
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(timeoutDurationMinutes * time.Minute):
		return errors.New("Timed out context has not reached running state")
	}
}

func (a CommandsAPI) waitForContextReady(contextID, clusterID string,
	sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		for {
			status, err := a.getContext(contextID, clusterID)
			if err != nil {
				errChan <- err
				return
			}
			if status == "Running" {
				errChan <- nil
				return
			} else if status == "Error" {
				errChan <- errors.New("context is in a errored state")
				return
			}

			log.Printf("[DEBUG] Waiting for context to go to running, current state is %s", status)
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(timeoutDurationMinutes * time.Minute):
		return errors.New("Timed out context has not reached running state")
	}
}
