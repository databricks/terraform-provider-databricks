package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// CommandsAPI exposes the Context & Commands API
type CommandsAPI struct {
	Client *DBApiClient
}

// CommandExecutor creates a spark context and executes a command and then closes context
type CommandExecutor interface {
	Execute(clusterID, language, commandStr string) (model.Command, error)
}

// Execute creates a spark context and executes a command and then closes context
func (a CommandsAPI) Execute(clusterID, language, commandStr string) (model.Command, error) {
	var resp model.Command
	context, err := a.createContext(language, clusterID)
	if err != nil {
		return resp, err
	}
	err = a.waitForContextReady(context, clusterID, 1, 10)
	if err != nil {
		return resp, err
	}
	commandID, err := a.createCommand(context, clusterID, language, commandStr)
	if err != nil {
		return resp, err
	}
	err = a.waitForCommandFinished(commandID, context, clusterID, 5, 10)
	if err != nil {
		return resp, err
	}
	command, err := a.getCommand(commandID, context, clusterID)
	if err != nil {
		return resp, err
	}
	err = a.deleteContext(context, clusterID)
	return command, err
}

func (a CommandsAPI) createCommand(contextID, clusterID, language, commandStr string) (string, error) {
	var command struct {
		ID string `json:"id,omitempty"`
	}
	commandRequest := struct {
		Language  string `json:"language,omitempty"`
		ClusterID string `json:"clusterId,omitempty"`
		ContextID string `json:"contextId,omitempty"`
		Command   string `json:"command,omitempty"`
	}{
		Language:  language,
		ClusterID: clusterID,
		ContextID: contextID,
		Command:   commandStr,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/commands/execute", "1.2", nil, commandRequest, nil)
	if err != nil {
		return command.ID, err
	}
	err = json.Unmarshal(resp, &command)
	return command.ID, err
}

func (a CommandsAPI) getCommand(commandID, contextID, clusterID string) (model.Command, error) {
	var commandResp model.Command
	contextGetRequest := struct {
		CommandID string `json:"commandId,omitempty" url:"commandId,omitempty"`
		ContextID string `json:"contextId,omitempty" url:"contextId,omitempty"`
		ClusterID string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	}{
		CommandID: commandID,
		ContextID: contextID,
		ClusterID: clusterID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/commands/status", "1.2", nil, contextGetRequest, nil)
	if err != nil {
		return commandResp, err
	}
	err = json.Unmarshal(resp, &commandResp)
	return commandResp, err
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
			log.Println(fmt.Sprintf("Waiting for command to finish, current state is: %s.", commandInfo.Status))
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

// Commented till this will be used
//func (a CommandsAPI) deleteCommand(commandID, contextID, clusterID string) error {
//	contextDeleteRequest := struct {
//		CommandID string `json:"commandId,omitempty" url:"commandId,omitempty"`
//		ContextID string `json:"contextId,omitempty" url:"contextId,omitempty"`
//		ClusterID string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
//	}{
//		CommandID: commandID,
//		ContextID: contextID,
//		ClusterID: clusterID,
//	}
//	_, err := a.Client.performQuery(http.MethodPost, "/commands/cancel", "1.2", nil, contextDeleteRequest, nil)
//	return err
//}

func (a CommandsAPI) deleteContext(contextID, clusterID string) error {
	contextDeleteRequest := struct {
		ContextID string `json:"contextId,omitempty" url:"contextId,omitempty"`
		ClusterID string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	}{
		ContextID: contextID,
		ClusterID: clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/contexts/destroy", "1.2", nil, contextDeleteRequest, nil)
	return err
}

func (a CommandsAPI) waitForContextReady(contextID, clusterID string, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
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
			log.Println("Waiting for context to go to running, current state is pending.")
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

func (a CommandsAPI) getContext(contextID, clusterID string) (string, error) {
	var contextStatus struct {
		ID     string `json:"id,omitempty"`
		Status string `json:"status,omitempty"`
	}
	contextGetRequest := struct {
		ContextID string `json:"contextId,omitempty" url:"contextId,omitempty"`
		ClusterID string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	}{
		ContextID: contextID,
		ClusterID: clusterID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/contexts/status", "1.2", nil, contextGetRequest, nil)
	if err != nil {
		return contextStatus.Status, err
	}
	err = json.Unmarshal(resp, &contextStatus)
	return contextStatus.Status, err
}

func (a CommandsAPI) createContext(language, clusterID string) (string, error) {
	var context struct {
		ID string `json:"id,omitempty"`
	}
	contextRequest := struct {
		Language  string `json:"language,omitempty"`
		ClusterID string `json:"clusterId,omitempty"`
	}{
		Language:  language,
		ClusterID: clusterID,
	}

	resp, err := a.Client.performQuery(http.MethodPost, "/contexts/create", "1.2", nil, contextRequest, nil)
	if err != nil {
		return context.ID, err
	}
	err = json.Unmarshal(resp, &context)
	return context.ID, err
}
