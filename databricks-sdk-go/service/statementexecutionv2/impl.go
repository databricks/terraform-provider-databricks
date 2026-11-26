// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package statementexecutionv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just StatementExecutionV2 API methods
type statementExecutionV2Impl struct {
	client *client.DatabricksClient
}

func (a *statementExecutionV2Impl) CancelStatement(ctx context.Context, request CancelStatementRequest) (*CancelStatementResponse, error) {
	var cancelStatementResponse CancelStatementResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/statements/%v/cancel", request.WarehouseId, request.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelStatementResponse)
	return &cancelStatementResponse, err
}

func (a *statementExecutionV2Impl) ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/statements", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &statementResponse)
	return &statementResponse, err
}

func (a *statementExecutionV2Impl) GetResultData(ctx context.Context, request GetResultDataRequest) (*ResultData, error) {
	var resultData ResultData
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/statements/%v/result/chunks/%v", request.WarehouseId, request.StatementId, request.ChunkIndex)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &resultData)
	return &resultData, err
}

func (a *statementExecutionV2Impl) GetStatementResult(ctx context.Context, request GetStatementResultRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/statements/%v", request.WarehouseId, request.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &statementResponse)
	return &statementResponse, err
}
