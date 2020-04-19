package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"log"
	"net/http"
	"time"
)

// TokensAPI exposes the Secrets API
type CommandsAPI struct {
	Client DBApiClient
}

func (a CommandsAPI) Execute(clusterId, langauge, commandStr string) (model.Command, error) {
	var resp model.Command
	context, err := a.createContext(langauge, clusterId)
	if err != nil {
		return resp, err
	}
	err = a.waitForContextReady(context, clusterId, 1, 10)
	if err != nil {
		return resp, err
	}
	commandId, err := a.createCommand(context, clusterId, langauge, commandStr)
	if err != nil {
		return resp, err
	}
	err = a.waitForCommandFinished(commandId, context, clusterId, 5, 10)
	if err != nil {
		return resp, err
	}
	command, err := a.getCommand(commandId, context, clusterId)
	if err != nil {
		return resp, err
	}
	err = a.deleteContext(context, clusterId)
	return command, err
}

func (a CommandsAPI) createCommand(contextId, clusterId, language, commandStr string) (string, error) {
	var command struct {
		Id string `json:"id,omitempty"`
	}
	commandRequest := struct {
		Language  string `json:"language,omitempty"`
		ClusterId string `json:"clusterId,omitempty"`
		ContextId string `json:"contextId,omitempty"`
		Command   string `json:"command,omitempty"`
	}{
		Language:  language,
		ClusterId: clusterId,
		ContextId: contextId,
		Command:   commandStr,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/commands/execute", "1.2", nil, commandRequest, nil)
	if err != nil {
		return command.Id, nil
	}
	log.Println(string(resp))
	err = json.Unmarshal(resp, &command)
	return command.Id, err
}

func (a CommandsAPI) getCommand(commandId, contextId, clusterId string) (model.Command, error) {
	var commandResp model.Command
	contextGetRequest := struct {
		CommandId string `json:"commandId,omitempty" url:"commandId,omitempty"`
		ContextId string `json:"contextId,omitempty" url:"contextId,omitempty"`
		ClusterId string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	}{
		CommandId: commandId,
		ContextId: contextId,
		ClusterId: clusterId,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/commands/status", "1.2", nil, contextGetRequest, nil)
	if err != nil {
		return commandResp, err
	}
	log.Println(string(resp))
	err = json.Unmarshal(resp, &commandResp)
	return commandResp, err
}

func (a CommandsAPI) waitForCommandFinished(commandId, contextId, clusterID string, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		for {
			commandInfo, err := a.getCommand(commandId, contextId, clusterID)
			if err != nil {
				errChan <- err
			}
			if commandInfo.Status == "Finished" {
				errChan <- nil
			} else if commandInfo.Status == "Cancelling" || commandInfo.Status == "Cancelled" || commandInfo.Status == "Error" {
				errChan <- errors.New(fmt.Sprintf("Context is in a failure state: %s.", commandInfo.Status))
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

func (a CommandsAPI) deleteCommand(commandId, contextId, clusterId string) error {
	contextDeleteRequest := struct {
		CommandId string `json:"commandId,omitempty" url:"commandId,omitempty"`
		ContextId string `json:"contextId,omitempty" url:"contextId,omitempty"`
		ClusterId string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	}{
		CommandId: commandId,
		ContextId: contextId,
		ClusterId: clusterId,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/commands/cancel", "1.2", nil, contextDeleteRequest, nil)
	return err
}

func (a CommandsAPI) deleteContext(contextId, clusterId string) error {
	contextDeleteRequest := struct {
		ContextId string `json:"contextId,omitempty" url:"contextId,omitempty"`
		ClusterId string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	}{
		ContextId: contextId,
		ClusterId: clusterId,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/contexts/destroy", "1.2", nil, contextDeleteRequest, nil)
	return err
}

func (a CommandsAPI) waitForContextReady(contextId, clusterID string, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		for {
			status, err := a.getContext(contextId, clusterID)
			if err != nil {
				errChan <- err
			}
			if status == "Running" {
				errChan <- nil
			} else if status == "Error" {
				errChan <- errors.New("Context is in a errored state.")
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

func (a CommandsAPI) getContext(contextId, clusterId string) (string, error) {
	var contextStatus struct {
		Id     string `json:"id,omitempty"`
		Status string `json:"status,omitempty"`
	}
	contextGetRequest := struct {
		ContextId string `json:"contextId,omitempty" url:"contextId,omitempty"`
		ClusterId string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	}{
		ContextId: contextId,
		ClusterId: clusterId,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/contexts/status", "1.2", nil, contextGetRequest, nil)
	if err != nil {
		return contextStatus.Status, err
	}
	err = json.Unmarshal(resp, &contextStatus)
	return contextStatus.Status, err
}

func (a CommandsAPI) createContext(language, clusterId string) (string, error) {
	var context struct {
		Id string `json:"id,omitempty"`
	}
	contextRequest := struct {
		Language  string `json:"language,omitempty"`
		ClusterId string `json:"clusterId,omitempty"`
	}{
		Language:  language,
		ClusterId: clusterId,
	}

	resp, err := a.Client.performQuery(http.MethodPost, "/contexts/create", "1.2", nil, contextRequest, nil)
	if err != nil {
		return context.Id, nil
	}
	err = json.Unmarshal(resp, &context)
	return context.Id, err
}
