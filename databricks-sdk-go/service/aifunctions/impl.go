// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aifunctions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AiFunctions API methods
type aiFunctionsImpl struct {
	client *client.DatabricksClient
}

func (a *aiFunctionsImpl) AiClassify(ctx context.Context, request AiClassifyRequest) (*AiClassifyResponse, error) {
	var aiClassifyResponse AiClassifyResponse
	path := "/api/2.0/ai-functions/ai-classify"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &aiClassifyResponse)
	return &aiClassifyResponse, err
}

func (a *aiFunctionsImpl) AiDecide(ctx context.Context, request AiDecideRequest) (*AiDecideResponse, error) {
	var aiDecideResponse AiDecideResponse
	path := "/api/2.0/ai-functions/ai-decide"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &aiDecideResponse)
	return &aiDecideResponse, err
}

func (a *aiFunctionsImpl) AiExtract(ctx context.Context, request AiExtractRequest) (*AiExtractResponse, error) {
	var aiExtractResponse AiExtractResponse
	path := "/api/2.0/ai-functions/ai-extract"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &aiExtractResponse)
	return &aiExtractResponse, err
}

func (a *aiFunctionsImpl) AiParseDocument(ctx context.Context, request AiParseDocumentRequest) (*AiParseDocumentResponse, error) {
	var aiParseDocumentResponse AiParseDocumentResponse
	path := "/api/2.0/ai-functions/ai-parse-document"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &aiParseDocumentResponse)
	return &aiParseDocumentResponse, err
}

func (a *aiFunctionsImpl) CancelBatch(ctx context.Context, request CancelBatchRequest) (*CancelBatchResponse, error) {
	var cancelBatchResponse CancelBatchResponse
	path := fmt.Sprintf("/api/2.0/ai-functions/batches/%v/cancel", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelBatchResponse)
	return &cancelBatchResponse, err
}

func (a *aiFunctionsImpl) CreateBatch(ctx context.Context, request CreateBatchRequest) (*Batch, error) {
	var batch Batch
	path := "/api/2.0/ai-functions/batches"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Batch, &batch)
	return &batch, err
}

func (a *aiFunctionsImpl) GetBatch(ctx context.Context, request GetBatchRequest) (*Batch, error) {
	var batch Batch
	path := fmt.Sprintf("/api/2.0/ai-functions/batches/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &batch)
	return &batch, err
}

// Lists the batches that the caller created in this workspace.
func (a *aiFunctionsImpl) ListBatches(ctx context.Context, request ListBatchesRequest) listing.Iterator[Batch] {

	getNextPage := func(ctx context.Context, req ListBatchesRequest) (*ListBatchesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListBatches(ctx, req)
	}
	getItems := func(resp *ListBatchesResponse) []Batch {
		return resp.Batches
	}
	getNextReq := func(resp *ListBatchesResponse) *ListBatchesRequest {
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

// Lists the batches that the caller created in this workspace.
func (a *aiFunctionsImpl) ListBatchesAll(ctx context.Context, request ListBatchesRequest) ([]Batch, error) {
	iterator := a.ListBatches(ctx, request)
	return listing.ToSlice[Batch](ctx, iterator)
}

func (a *aiFunctionsImpl) internalListBatches(ctx context.Context, request ListBatchesRequest) (*ListBatchesResponse, error) {
	var listBatchesResponse ListBatchesResponse
	path := "/api/2.0/ai-functions/batches"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBatchesResponse)
	return &listBatchesResponse, err
}
