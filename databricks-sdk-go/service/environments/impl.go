// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package environments

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Environments API methods
type environmentsImpl struct {
	client *client.DatabricksClient
}

func (a *environmentsImpl) CreateWorkspaceBaseEnvironment(ctx context.Context, request CreateWorkspaceBaseEnvironmentRequest) (*Operation, error) {
	var operation Operation
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := "/api/environments/v1/workspace-base-environments"
	queryParams := make(map[string]any)

	if request.RequestId != "" || slices.Contains(request.ForceSendFields, "RequestId") {
		queryParams["request_id"] = request.RequestId
	}

	if request.WorkspaceBaseEnvironmentId != "" || slices.Contains(request.ForceSendFields, "WorkspaceBaseEnvironmentId") {
		queryParams["workspace_base_environment_id"] = request.WorkspaceBaseEnvironmentId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceBaseEnvironment, &operation)
	return &operation, err
}

func (a *environmentsImpl) DeleteWorkspaceBaseEnvironment(ctx context.Context, request DeleteWorkspaceBaseEnvironmentRequest) error {
	path := fmt.Sprintf("/api/environments/v1/%v", request.Name)
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

func (a *environmentsImpl) GetDefaultWorkspaceBaseEnvironment(ctx context.Context, request GetDefaultWorkspaceBaseEnvironmentRequest) (*DefaultWorkspaceBaseEnvironment, error) {
	var defaultWorkspaceBaseEnvironment DefaultWorkspaceBaseEnvironment
	path := fmt.Sprintf("/api/environments/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &defaultWorkspaceBaseEnvironment)
	return &defaultWorkspaceBaseEnvironment, err
}

func (a *environmentsImpl) GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/environments/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *environmentsImpl) GetWorkspaceBaseEnvironment(ctx context.Context, request GetWorkspaceBaseEnvironmentRequest) (*WorkspaceBaseEnvironment, error) {
	var workspaceBaseEnvironment WorkspaceBaseEnvironment
	path := fmt.Sprintf("/api/environments/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceBaseEnvironment)
	return &workspaceBaseEnvironment, err
}

// Lists all WorkspaceBaseEnvironments in the workspace.
//
// Databricks provides the following base environments:
//
// - `workspace-base-environments/databricks_ai_...`: includes popular AI and
// deep learning packages for serverless GPU compute.
//
// - `workspace-base-environments/databricks_ml_...`: includes popular ML
// packages for serverless compute.
//
// Databricks-provided base environments are versioned. For example,
// `workspace-base-environments/databricks_ml_v5` corresponds to the ML
// environment built on environment version 5.
func (a *environmentsImpl) ListWorkspaceBaseEnvironments(ctx context.Context, request ListWorkspaceBaseEnvironmentsRequest) listing.Iterator[WorkspaceBaseEnvironment] {

	getNextPage := func(ctx context.Context, req ListWorkspaceBaseEnvironmentsRequest) (*ListWorkspaceBaseEnvironmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListWorkspaceBaseEnvironments(ctx, req)
	}
	getItems := func(resp *ListWorkspaceBaseEnvironmentsResponse) []WorkspaceBaseEnvironment {
		return resp.WorkspaceBaseEnvironments
	}
	getNextReq := func(resp *ListWorkspaceBaseEnvironmentsResponse) *ListWorkspaceBaseEnvironmentsRequest {
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

// Lists all WorkspaceBaseEnvironments in the workspace.
//
// Databricks provides the following base environments:
//
// - `workspace-base-environments/databricks_ai_...`: includes popular AI and
// deep learning packages for serverless GPU compute.
//
// - `workspace-base-environments/databricks_ml_...`: includes popular ML
// packages for serverless compute.
//
// Databricks-provided base environments are versioned. For example,
// `workspace-base-environments/databricks_ml_v5` corresponds to the ML
// environment built on environment version 5.
func (a *environmentsImpl) ListWorkspaceBaseEnvironmentsAll(ctx context.Context, request ListWorkspaceBaseEnvironmentsRequest) ([]WorkspaceBaseEnvironment, error) {
	iterator := a.ListWorkspaceBaseEnvironments(ctx, request)
	return listing.ToSlice[WorkspaceBaseEnvironment](ctx, iterator)
}

func (a *environmentsImpl) internalListWorkspaceBaseEnvironments(ctx context.Context, request ListWorkspaceBaseEnvironmentsRequest) (*ListWorkspaceBaseEnvironmentsResponse, error) {
	var listWorkspaceBaseEnvironmentsResponse ListWorkspaceBaseEnvironmentsResponse
	path := "/api/environments/v1/workspace-base-environments"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceBaseEnvironmentsResponse)
	return &listWorkspaceBaseEnvironmentsResponse, err
}

func (a *environmentsImpl) RefreshWorkspaceBaseEnvironment(ctx context.Context, request RefreshWorkspaceBaseEnvironmentRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/environments/v1/%v/refresh", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *environmentsImpl) UpdateDefaultWorkspaceBaseEnvironment(ctx context.Context, request UpdateDefaultWorkspaceBaseEnvironmentRequest) (*DefaultWorkspaceBaseEnvironment, error) {
	var defaultWorkspaceBaseEnvironment DefaultWorkspaceBaseEnvironment
	path := fmt.Sprintf("/api/environments/v1/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.DefaultWorkspaceBaseEnvironment, &defaultWorkspaceBaseEnvironment)
	return &defaultWorkspaceBaseEnvironment, err
}

func (a *environmentsImpl) UpdateWorkspaceBaseEnvironment(ctx context.Context, request UpdateWorkspaceBaseEnvironmentRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/environments/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceBaseEnvironment, &operation)
	return &operation, err
}
