// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package supervisoragents

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just SupervisorAgents API methods
type supervisorAgentsImpl struct {
	client *client.DatabricksClient
}

func (a *supervisorAgentsImpl) CreateExample(ctx context.Context, request CreateExampleRequest) (*Example, error) {
	var example Example
	path := fmt.Sprintf("/api/2.1/%v/examples", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Example, &example)
	return &example, err
}

func (a *supervisorAgentsImpl) CreateSupervisorAgent(ctx context.Context, request CreateSupervisorAgentRequest) (*SupervisorAgent, error) {
	var supervisorAgent SupervisorAgent
	path := "/api/2.1/supervisor-agents"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.SupervisorAgent, &supervisorAgent)
	return &supervisorAgent, err
}

func (a *supervisorAgentsImpl) CreateTool(ctx context.Context, request CreateToolRequest) (*Tool, error) {
	var tool Tool
	path := fmt.Sprintf("/api/2.1/%v/tools", request.Parent)
	queryParams := make(map[string]any)

	if request.ToolId != "" {
		queryParams["tool_id"] = request.ToolId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Tool, &tool)
	return &tool, err
}

func (a *supervisorAgentsImpl) DeleteExample(ctx context.Context, request DeleteExampleRequest) error {
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *supervisorAgentsImpl) DeleteSupervisorAgent(ctx context.Context, request DeleteSupervisorAgentRequest) error {
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *supervisorAgentsImpl) DeleteTool(ctx context.Context, request DeleteToolRequest) error {
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *supervisorAgentsImpl) GetExample(ctx context.Context, request GetExampleRequest) (*Example, error) {
	var example Example
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &example)
	return &example, err
}

func (a *supervisorAgentsImpl) GetPermissionLevels(ctx context.Context, request GetSupervisorAgentPermissionLevelsRequest) (*GetSupervisorAgentPermissionLevelsResponse, error) {
	var getSupervisorAgentPermissionLevelsResponse GetSupervisorAgentPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/supervisor-agents/%v/permissionLevels", request.SupervisorAgentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getSupervisorAgentPermissionLevelsResponse)
	return &getSupervisorAgentPermissionLevelsResponse, err
}

func (a *supervisorAgentsImpl) GetPermissions(ctx context.Context, request GetSupervisorAgentPermissionsRequest) (*SupervisorAgentPermissions, error) {
	var supervisorAgentPermissions SupervisorAgentPermissions
	path := fmt.Sprintf("/api/2.0/permissions/supervisor-agents/%v", request.SupervisorAgentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &supervisorAgentPermissions)
	return &supervisorAgentPermissions, err
}

func (a *supervisorAgentsImpl) GetSupervisorAgent(ctx context.Context, request GetSupervisorAgentRequest) (*SupervisorAgent, error) {
	var supervisorAgent SupervisorAgent
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &supervisorAgent)
	return &supervisorAgent, err
}

func (a *supervisorAgentsImpl) GetTool(ctx context.Context, request GetToolRequest) (*Tool, error) {
	var tool Tool
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &tool)
	return &tool, err
}

// Lists examples under a Supervisor Agent.
func (a *supervisorAgentsImpl) ListExamples(ctx context.Context, request ListExamplesRequest) listing.Iterator[Example] {

	getNextPage := func(ctx context.Context, req ListExamplesRequest) (*ListExamplesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExamples(ctx, req)
	}
	getItems := func(resp *ListExamplesResponse) []Example {
		return resp.Examples
	}
	getNextReq := func(resp *ListExamplesResponse) *ListExamplesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists examples under a Supervisor Agent.
func (a *supervisorAgentsImpl) ListExamplesAll(ctx context.Context, request ListExamplesRequest) ([]Example, error) {
	iterator := a.ListExamples(ctx, request)
	return listing.ToSlice[Example](ctx, iterator)
}

func (a *supervisorAgentsImpl) internalListExamples(ctx context.Context, request ListExamplesRequest) (*ListExamplesResponse, error) {
	var listExamplesResponse ListExamplesResponse
	path := fmt.Sprintf("/api/2.1/%v/examples", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExamplesResponse)
	return &listExamplesResponse, err
}

// Lists the Supervisor Agents in the workspace that the caller has permission
// to view.
func (a *supervisorAgentsImpl) ListSupervisorAgents(ctx context.Context, request ListSupervisorAgentsRequest) listing.Iterator[SupervisorAgent] {

	getNextPage := func(ctx context.Context, req ListSupervisorAgentsRequest) (*ListSupervisorAgentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSupervisorAgents(ctx, req)
	}
	getItems := func(resp *ListSupervisorAgentsResponse) []SupervisorAgent {
		return resp.SupervisorAgents
	}
	getNextReq := func(resp *ListSupervisorAgentsResponse) *ListSupervisorAgentsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the Supervisor Agents in the workspace that the caller has permission
// to view.
func (a *supervisorAgentsImpl) ListSupervisorAgentsAll(ctx context.Context, request ListSupervisorAgentsRequest) ([]SupervisorAgent, error) {
	iterator := a.ListSupervisorAgents(ctx, request)
	return listing.ToSlice[SupervisorAgent](ctx, iterator)
}

func (a *supervisorAgentsImpl) internalListSupervisorAgents(ctx context.Context, request ListSupervisorAgentsRequest) (*ListSupervisorAgentsResponse, error) {
	var listSupervisorAgentsResponse ListSupervisorAgentsResponse
	path := "/api/2.1/supervisor-agents"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSupervisorAgentsResponse)
	return &listSupervisorAgentsResponse, err
}

// Lists Tools under a Supervisor Agent.
func (a *supervisorAgentsImpl) ListTools(ctx context.Context, request ListToolsRequest) listing.Iterator[Tool] {

	getNextPage := func(ctx context.Context, req ListToolsRequest) (*ListToolsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListTools(ctx, req)
	}
	getItems := func(resp *ListToolsResponse) []Tool {
		return resp.Tools
	}
	getNextReq := func(resp *ListToolsResponse) *ListToolsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists Tools under a Supervisor Agent.
func (a *supervisorAgentsImpl) ListToolsAll(ctx context.Context, request ListToolsRequest) ([]Tool, error) {
	iterator := a.ListTools(ctx, request)
	return listing.ToSlice[Tool](ctx, iterator)
}

func (a *supervisorAgentsImpl) internalListTools(ctx context.Context, request ListToolsRequest) (*ListToolsResponse, error) {
	var listToolsResponse ListToolsResponse
	path := fmt.Sprintf("/api/2.1/%v/tools", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listToolsResponse)
	return &listToolsResponse, err
}

func (a *supervisorAgentsImpl) SetPermissions(ctx context.Context, request SupervisorAgentPermissionsRequest) (*SupervisorAgentPermissions, error) {
	var supervisorAgentPermissions SupervisorAgentPermissions
	path := fmt.Sprintf("/api/2.0/permissions/supervisor-agents/%v", request.SupervisorAgentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &supervisorAgentPermissions)
	return &supervisorAgentPermissions, err
}

func (a *supervisorAgentsImpl) UpdateExample(ctx context.Context, request UpdateExampleRequest) (*Example, error) {
	var example Example
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Example, &example)
	return &example, err
}

func (a *supervisorAgentsImpl) UpdatePermissions(ctx context.Context, request SupervisorAgentPermissionsRequest) (*SupervisorAgentPermissions, error) {
	var supervisorAgentPermissions SupervisorAgentPermissions
	path := fmt.Sprintf("/api/2.0/permissions/supervisor-agents/%v", request.SupervisorAgentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &supervisorAgentPermissions)
	return &supervisorAgentPermissions, err
}

func (a *supervisorAgentsImpl) UpdateSupervisorAgent(ctx context.Context, request UpdateSupervisorAgentRequest) (*SupervisorAgent, error) {
	var supervisorAgent SupervisorAgent
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.SupervisorAgent, &supervisorAgent)
	return &supervisorAgent, err
}

func (a *supervisorAgentsImpl) UpdateTool(ctx context.Context, request UpdateToolRequest) (*Tool, error) {
	var tool Tool
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Tool, &tool)
	return &tool, err
}
