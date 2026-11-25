// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package agentbricks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AgentBricks API methods
type agentBricksImpl struct {
	client *client.DatabricksClient
}

func (a *agentBricksImpl) CancelOptimize(ctx context.Context, request CancelCustomLlmOptimizationRunRequest) error {
	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize/cancel", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, nil)
	return err
}

func (a *agentBricksImpl) CreateCustomLlm(ctx context.Context, request CreateCustomLlmRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := "/api/2.0/custom-llms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}

func (a *agentBricksImpl) DeleteCustomLlm(ctx context.Context, request DeleteCustomLlmRequest) error {
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
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

func (a *agentBricksImpl) GetCustomLlm(ctx context.Context, request GetCustomLlmRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}

func (a *agentBricksImpl) StartOptimize(ctx context.Context, request StartCustomLlmOptimizationRunRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}

func (a *agentBricksImpl) UpdateCustomLlm(ctx context.Context, request UpdateCustomLlmRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}
