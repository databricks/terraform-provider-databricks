// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aisearch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just AiSearch API methods
type aiSearchImpl struct {
	client *client.DatabricksClient
}

func (a *aiSearchImpl) CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Endpoint, error) {
	var endpoint Endpoint
	path := fmt.Sprintf("/api/2.0/ai-search/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)

	if request.EndpointId != "" || slices.Contains(request.ForceSendFields, "EndpointId") {
		queryParams["endpoint_id"] = request.EndpointId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Endpoint, &endpoint)
	return &endpoint, err
}

func (a *aiSearchImpl) CreateIndex(ctx context.Context, request CreateIndexRequest) (*Index, error) {
	var index Index
	path := fmt.Sprintf("/api/2.0/ai-search/%v/indexes", request.Parent)
	queryParams := make(map[string]any)

	if request.IndexId != "" || slices.Contains(request.ForceSendFields, "IndexId") {
		queryParams["index_id"] = request.IndexId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Index, &index)
	return &index, err
}

func (a *aiSearchImpl) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error {
	path := fmt.Sprintf("/api/2.0/ai-search/%v", request.Name)
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

func (a *aiSearchImpl) DeleteIndex(ctx context.Context, request DeleteIndexRequest) error {
	path := fmt.Sprintf("/api/2.0/ai-search/%v", request.Name)
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

func (a *aiSearchImpl) GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error) {
	var endpoint Endpoint
	path := fmt.Sprintf("/api/2.0/ai-search/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &endpoint)
	return &endpoint, err
}

func (a *aiSearchImpl) GetIndex(ctx context.Context, request GetIndexRequest) (*Index, error) {
	var index Index
	path := fmt.Sprintf("/api/2.0/ai-search/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &index)
	return &index, err
}

// List AI Search endpoints in a workspace.
func (a *aiSearchImpl) ListEndpoints(ctx context.Context, request ListEndpointsRequest) listing.Iterator[Endpoint] {

	getNextPage := func(ctx context.Context, req ListEndpointsRequest) (*ListEndpointsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListEndpoints(ctx, req)
	}
	getItems := func(resp *ListEndpointsResponse) []Endpoint {
		return resp.Endpoints
	}
	getNextReq := func(resp *ListEndpointsResponse) *ListEndpointsRequest {
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

// List AI Search endpoints in a workspace.
func (a *aiSearchImpl) ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]Endpoint, error) {
	iterator := a.ListEndpoints(ctx, request)
	return listing.ToSlice[Endpoint](ctx, iterator)
}

func (a *aiSearchImpl) internalListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error) {
	var listEndpointsResponse ListEndpointsResponse
	path := fmt.Sprintf("/api/2.0/ai-search/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listEndpointsResponse)
	return &listEndpointsResponse, err
}

// List AI Search indexes on an endpoint.
func (a *aiSearchImpl) ListIndexes(ctx context.Context, request ListIndexesRequest) listing.Iterator[Index] {

	getNextPage := func(ctx context.Context, req ListIndexesRequest) (*ListIndexesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListIndexes(ctx, req)
	}
	getItems := func(resp *ListIndexesResponse) []Index {
		return resp.Indexes
	}
	getNextReq := func(resp *ListIndexesResponse) *ListIndexesRequest {
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

// List AI Search indexes on an endpoint.
func (a *aiSearchImpl) ListIndexesAll(ctx context.Context, request ListIndexesRequest) ([]Index, error) {
	iterator := a.ListIndexes(ctx, request)
	return listing.ToSlice[Index](ctx, iterator)
}

func (a *aiSearchImpl) internalListIndexes(ctx context.Context, request ListIndexesRequest) (*ListIndexesResponse, error) {
	var listIndexesResponse ListIndexesResponse
	path := fmt.Sprintf("/api/2.0/ai-search/%v/indexes", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listIndexesResponse)
	return &listIndexesResponse, err
}

func (a *aiSearchImpl) QueryIndex(ctx context.Context, request QueryIndexRequest) (*QueryIndexResponse, error) {
	var queryIndexResponse QueryIndexResponse
	path := fmt.Sprintf("/api/2.0/ai-search/%v:query", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &queryIndexResponse)
	return &queryIndexResponse, err
}

func (a *aiSearchImpl) RemoveData(ctx context.Context, request RemoveDataRequest) (*RemoveDataResponse, error) {
	var removeDataResponse RemoveDataResponse
	path := fmt.Sprintf("/api/2.0/ai-search/%v:removeData", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &removeDataResponse)
	return &removeDataResponse, err
}

func (a *aiSearchImpl) ScanIndex(ctx context.Context, request ScanIndexRequest) (*ScanIndexResponse, error) {
	var scanIndexResponse ScanIndexResponse
	path := fmt.Sprintf("/api/2.0/ai-search/%v:scan", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &scanIndexResponse)
	return &scanIndexResponse, err
}

func (a *aiSearchImpl) SyncIndex(ctx context.Context, request SyncIndexRequest) (*SyncIndexResponse, error) {
	var syncIndexResponse SyncIndexResponse
	path := fmt.Sprintf("/api/2.0/ai-search/%v:sync", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &syncIndexResponse)
	return &syncIndexResponse, err
}

func (a *aiSearchImpl) UpdateEndpoint(ctx context.Context, request UpdateEndpointRequest) (*Endpoint, error) {
	var endpoint Endpoint
	path := fmt.Sprintf("/api/2.0/ai-search/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Endpoint, &endpoint)
	return &endpoint, err
}

func (a *aiSearchImpl) UpsertData(ctx context.Context, request UpsertDataRequest) (*UpsertDataResponse, error) {
	var upsertDataResponse UpsertDataResponse
	path := fmt.Sprintf("/api/2.0/ai-search/%v:upsertData", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &upsertDataResponse)
	return &upsertDataResponse, err
}
