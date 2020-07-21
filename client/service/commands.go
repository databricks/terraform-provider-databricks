package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// CommandsAPI exposes the Context & Commands API
type CommandsAPI struct {
	client *DatabricksClient
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

type genericCommandRequest struct {
	CommandID string `json:"commandId,omitempty" url:"commandId,omitempty"`
	Language  string `json:"language,omitempty"`
	ClusterID string `json:"clusterId,omitempty"`
	ContextID string `json:"contextId,omitempty"`
	Command   string `json:"command,omitempty"`
}

func (a CommandsAPI) createCommand(contextID, clusterID, language, commandStr string) (string, error) {
	var command model.Command
	err := a.post("/commands/execute", genericCommandRequest{
		Language:  language,
		ClusterID: clusterID,
		ContextID: contextID,
		Command:   commandStr,
	}, &command)
	return command.ID, err
}

func (a CommandsAPI) getCommand(commandID, contextID, clusterID string) (model.Command, error) {
	var commandResp model.Command
	err := a.get("/commands/status", genericCommandRequest{
		CommandID: commandID,
		ContextID: contextID,
		ClusterID: clusterID,
	}, &commandResp)
	return commandResp, err
}

func (a CommandsAPI) deleteContext(contextID, clusterID string) error {
	return a.post("/contexts/destroy", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, nil)
}

func (a CommandsAPI) getContext(contextID, clusterID string) (string, error) {
	var contextStatus model.Command // internal hack, yes
	err := a.get("/contexts/status", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, &contextStatus)
	return contextStatus.Status, err
}

func (a CommandsAPI) createContext(language, clusterID string) (string, error) {
	var context model.Command // internal hack, yes
	err := a.post("/contexts/create", genericCommandRequest{
		Language:  language,
		ClusterID: clusterID,
	}, &context)
	return context.ID, err
}

func (a CommandsAPI) post(path string, request interface{}, response interface{}) error {
	if a.client.auth == nil {
		return fmt.Errorf("Authentication not initialized")
	}
	body, err := a.client.genericQuery2(http.MethodPost, path, request,
		a.client.auth, a.api12)
	if err != nil {
		return err
	}
	return a.client.unmarshall(path, body, &response)
}

func (a CommandsAPI) get(path string, request interface{}, response interface{}) error {
	if a.client.auth == nil {
		return fmt.Errorf("Authentication not initialized")
	}
	body, err := a.client.genericQuery2(http.MethodGet, path, request,
		a.client.auth, a.api12)
	if err != nil {
		return err
	}
	return a.client.unmarshall(path, body, &response)
}

func (a CommandsAPI) api12(r *http.Request) (*http.Request, error) {
	if r.URL == nil {
		return nil, fmt.Errorf("No URL found in request")
	}
	r.URL.Path = fmt.Sprintf("/api/1.2%s", r.URL.Path)
	r.Header.Set("Content-Type", "application/json")

	url, err := url.Parse(a.client.Host)
	if err != nil {
		return nil, err
	}
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme

	return r, nil
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
