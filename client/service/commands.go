package service

import (
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

var (
	// IPython's output prefixes
	outRE = regexp.MustCompile(`Out\[[\d\s]+\]:\s`)
	// HTML tags
	tagRE = regexp.MustCompile(`<[^>]*>`)
	// just exception content without exception name
	exceptionRE = regexp.MustCompile(`.*Exception: (.*)`)
	// usual error message explanation is hidden in this key
	errorMessageRE = regexp.MustCompile(`ErrorMessage=(.*)\n`)
)

// CommandsAPI exposes the Context & Commands API
type CommandsAPI struct {
	client *DatabricksClient
}

// CommandMock mocks the execution of command
type CommandMock func(commandStr string) (string, error)

// CommandExecutorMock simplifies command testing
type commandExecutorMock struct {
	mock CommandMock
}

// Execute mock command with given mock function
func (c commandExecutorMock) Execute(clusterID, language, commandStr string) (string, error) {
	return c.mock(commandStr)
}

// CommandExecutor creates a spark context and executes a command and then closes context
type CommandExecutor interface {
	Execute(clusterID, language, commandStr string) (string, error)
}

// TrimLeadingWhitespace removes leading whitespace
func TrimLeadingWhitespace(commandStr string) (newCommand string) {
	lines := strings.Split(commandStr, "\n")
	leadingWhitespace := 1<<31 - 1
	for _, line := range lines {
		for pos, char := range line {
			if char == ' ' || char == '\t' {
				continue
			}
			// first non-whitespace character
			if pos < leadingWhitespace {
				leadingWhitespace = pos
			}
			// is not needed further
			break
		}
	}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		if len(lines[i]) < leadingWhitespace {
			newCommand += lines[i] + "\n" // or not..
		} else {
			newCommand += lines[i][leadingWhitespace:] + "\n"
		}
	}
	return
}

// Execute creates a spark context and executes a command and then closes context
// Any leading whitespace is trimmed
func (a CommandsAPI) Execute(clusterID, language, commandStr string) (result string, err error) {
	cluster, err := a.client.Clusters().Get(clusterID)
	if err != nil {
		return
	}
	if !cluster.IsRunningOrResizing() {
		err = fmt.Errorf("Cluster %s has to be running or resizing, but is %s", clusterID, cluster.State)
		return
	}
	commandStr = TrimLeadingWhitespace(commandStr)
	log.Printf("[INFO] Executing %s command on %s:\n%s", language, clusterID, commandStr)
	context, err := a.createContext(language, clusterID)
	if err != nil {
		return
	}
	err = a.waitForContextReady(context, clusterID, 1, 10)
	if err != nil {
		return
	}
	commandID, err := a.createCommand(context, clusterID, language, commandStr)
	if err != nil {
		return
	}
	err = a.waitForCommandFinished(commandID, context, clusterID, 5, 10)
	if err != nil {
		return
	}
	command, err := a.getCommand(commandID, context, clusterID)
	if err != nil {
		return
	}
	err = a.deleteContext(context, clusterID)
	if err != nil {
		return
	}
	return a.parseCommandResults(command)
}

func (a CommandsAPI) parseCommandResults(command model.Command) (result string, err error) {
	switch command.Results.ResultType {
	case "text":
		result = outRE.ReplaceAllLiteralString(command.Results.Data.(string), "")
		return
	case "error":
		log.Printf("[DEBUG] error caused by command: %s", command.Results.Cause)
		summary := tagRE.ReplaceAllLiteralString(command.Results.Summary, "")
		summary = html.UnescapeString(summary)
		exceptionMatches := exceptionRE.FindStringSubmatch(summary)
		if len(exceptionMatches) == 2 {
			summary = strings.ReplaceAll(exceptionMatches[1], "; nested exception is:", "")
			summary = strings.TrimRight(summary, " ")
			err = errors.New(summary)
			return
		}
		errorMessageMatches := errorMessageRE.FindStringSubmatch(command.Results.Cause)
		if len(errorMessageMatches) == 2 {
			err = errors.New(errorMessageMatches[1])
			return
		}
		err = errors.New(summary)
		return
	}
	err = fmt.Errorf("Unknown result type %s: %v", command.Results.ResultType, command.Results.Data)
	return
}

type genericCommandRequest struct {
	CommandID string `json:"commandId,omitempty" url:"commandId,omitempty"`
	Language  string `json:"language,omitempty" url:"language,omitempty"`
	ClusterID string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	ContextID string `json:"contextId,omitempty" url:"contextId,omitempty"`
	Command   string `json:"command,omitempty" url:"command,omitempty"`
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
	if a.client.authVisitor == nil {
		return fmt.Errorf("Authentication not initialized")
	}
	body, err := a.client.genericQuery(http.MethodPost, path, request,
		a.client.authVisitor, a.api12)
	if err != nil {
		return err
	}
	return a.client.unmarshall(path, body, &response)
}

func (a CommandsAPI) get(path string, request interface{}, response interface{}) error {
	if a.client.authVisitor == nil {
		return fmt.Errorf("Authentication not initialized")
	}
	body, err := a.client.genericQuery(http.MethodGet, path, request,
		a.client.authVisitor, a.api12)
	if err != nil {
		return err
	}
	return a.client.unmarshall(path, body, &response)
}

func (a CommandsAPI) api12(r *http.Request) error {
	if r.URL == nil {
		return fmt.Errorf("No URL found in request")
	}
	r.URL.Path = fmt.Sprintf("/api/1.2%s", r.URL.Path)
	r.Header.Set("Content-Type", "application/json")

	url, err := url.Parse(a.client.Host)
	if err != nil {
		return err
	}
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme

	return nil
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
			log.Println("[DEBUG] Waiting for context to go to running, current state is pending.")
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
