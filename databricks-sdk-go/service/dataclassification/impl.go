// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dataclassification

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just DataClassification API methods
type dataClassificationImpl struct {
	client *client.DatabricksClient
}

func (a *dataClassificationImpl) CreateCatalogConfig(ctx context.Context, request CreateCatalogConfigRequest) (*CatalogConfig, error) {
	var catalogConfig CatalogConfig
	path := fmt.Sprintf("/api/data-classification/v1/%v/config", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.CatalogConfig, &catalogConfig)
	return &catalogConfig, err
}

func (a *dataClassificationImpl) DeleteCatalogConfig(ctx context.Context, request DeleteCatalogConfigRequest) error {
	path := fmt.Sprintf("/api/data-classification/v1/%v", request.Name)
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

func (a *dataClassificationImpl) GetCatalogConfig(ctx context.Context, request GetCatalogConfigRequest) (*CatalogConfig, error) {
	var catalogConfig CatalogConfig
	path := fmt.Sprintf("/api/data-classification/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &catalogConfig)
	return &catalogConfig, err
}

func (a *dataClassificationImpl) RescanCatalog(ctx context.Context, request RescanCatalogRequest) (*RescanCatalogResponse, error) {
	var rescanCatalogResponse RescanCatalogResponse
	path := fmt.Sprintf("/api/data-classification/v1/%v/rescan", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &rescanCatalogResponse)
	return &rescanCatalogResponse, err
}

func (a *dataClassificationImpl) UpdateCatalogConfig(ctx context.Context, request UpdateCatalogConfigRequest) (*CatalogConfig, error) {
	var catalogConfig CatalogConfig
	path := fmt.Sprintf("/api/data-classification/v1/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.CatalogConfig, &catalogConfig)
	return &catalogConfig, err
}
