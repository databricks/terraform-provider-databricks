package compute

import (
	"context"
	"errors"
	"fmt"
	"html"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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

// NewCommandsAPI creates CommandsAPI instance from provider meta
func NewCommandsAPI(ctx context.Context, m interface{}) CommandsAPI {
	return CommandsAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// CommandsAPI exposes the Context & Commands API
type CommandsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Execute creates a spark context and executes a command and then closes context
// Any leading whitespace is trimmed
func (a CommandsAPI) Execute(clusterID, language, commandStr string) (result string, err error) {
	cluster, err := NewClustersAPI(a.context, a.client).Get(clusterID)
	if err != nil {
		return
	}
	if !cluster.IsRunningOrResizing() {
		err = fmt.Errorf("Cluster %s has to be running or resizing, but is %s", clusterID, cluster.State)
		return
	}
	commandStr = internal.TrimLeadingWhitespace(commandStr)
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
	err = a.waitForCommandFinished(commandID, context, clusterID)
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

func (a CommandsAPI) parseCommandResults(command Command) (result string, err error) {
	if command.Results == nil {
		err = fmt.Errorf("Command has no results: %#v", command)
		return
	}
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
	var command Command
	err := a.client.OldAPI(a.context, "POST", "/commands/execute", genericCommandRequest{
		Language:  language,
		ClusterID: clusterID,
		ContextID: contextID,
		Command:   commandStr,
	}, &command)
	return command.ID, err
}

func (a CommandsAPI) getCommand(commandID, contextID, clusterID string) (Command, error) {
	var commandResp Command
	err := a.client.OldAPI(a.context, "GET", "/commands/status", genericCommandRequest{
		CommandID: commandID,
		ContextID: contextID,
		ClusterID: clusterID,
	}, &commandResp)
	return commandResp, err
}

func (a CommandsAPI) deleteContext(contextID, clusterID string) error {
	return a.client.OldAPI(a.context, "POST", "/contexts/destroy", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, nil)
}

func (a CommandsAPI) getContext(contextID, clusterID string) (string, error) {
	var contextStatus Command // internal hack, yes
	err := a.client.OldAPI(a.context, "GET", "/contexts/status", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, &contextStatus)
	return contextStatus.Status, err
}

func (a CommandsAPI) createContext(language, clusterID string) (string, error) {
	var context Command // internal hack, yes
	err := a.client.OldAPI(a.context, "POST", "/contexts/create", genericCommandRequest{
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
		return resource.RetryableError(fmt.Errorf(commandInfo.Status))
	})
}

func (a CommandsAPI) waitForContextReady(contextID, clusterID string) error {
	return resource.RetryContext(a.context, 10*time.Minute, func() *resource.RetryError {
		status, err := a.getContext(contextID, clusterID)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if status == "Error" {
			return resource.NonRetryableError(fmt.Errorf(status))
		}
		if status == "Running" {
			return nil
		}
		return resource.RetryableError(fmt.Errorf(status))
	})
}
