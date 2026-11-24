// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lrotesting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just LroTesting API methods
type lroTestingImpl struct {
	client *client.DatabricksClient
}

func (a *lroTestingImpl) CancelOperation(ctx context.Context, request CancelOperationRequest) error {
	path := fmt.Sprintf("/api/2.0/lro-testing/operations/%v/cancel", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, nil)
	return err
}

func (a *lroTestingImpl) CreateTestResource(ctx context.Context, request CreateTestResourceRequest) (*Operation, error) {
	var operation Operation
	path := "/api/2.0/lro-testing/resources"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Resource, &operation)
	return &operation, err
}

func (a *lroTestingImpl) DeleteTestResource(ctx context.Context, request DeleteTestResourceRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/lro-testing/resources/%v", request.ResourceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *lroTestingImpl) GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/lro-testing/operations/%v", request.Name)
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

func (a *lroTestingImpl) GetTestResource(ctx context.Context, request GetTestResourceRequest) (*TestResource, error) {
	var testResource TestResource
	path := fmt.Sprintf("/api/2.0/lro-testing/resources/%v", request.ResourceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &testResource)
	return &testResource, err
}
