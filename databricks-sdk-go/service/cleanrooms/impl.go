// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just CleanRoomAssetRevisions API methods
type cleanRoomAssetRevisionsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomAssetRevisionsImpl) Get(ctx context.Context, request GetCleanRoomAssetRevisionRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v/revisions/%v", request.CleanRoomName, request.AssetType, request.Name, request.Etag)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

// List revisions for an asset
func (a *cleanRoomAssetRevisionsImpl) List(ctx context.Context, request ListCleanRoomAssetRevisionsRequest) listing.Iterator[CleanRoomAsset] {

	getNextPage := func(ctx context.Context, req ListCleanRoomAssetRevisionsRequest) (*ListCleanRoomAssetRevisionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomAssetRevisionsResponse) []CleanRoomAsset {
		return resp.Revisions
	}
	getNextReq := func(resp *ListCleanRoomAssetRevisionsResponse) *ListCleanRoomAssetRevisionsRequest {
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

// List revisions for an asset
func (a *cleanRoomAssetRevisionsImpl) ListAll(ctx context.Context, request ListCleanRoomAssetRevisionsRequest) ([]CleanRoomAsset, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomAsset](ctx, iterator)
}

func (a *cleanRoomAssetRevisionsImpl) internalList(ctx context.Context, request ListCleanRoomAssetRevisionsRequest) (*ListCleanRoomAssetRevisionsResponse, error) {
	var listCleanRoomAssetRevisionsResponse ListCleanRoomAssetRevisionsResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v/revisions", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCleanRoomAssetRevisionsResponse)
	return &listCleanRoomAssetRevisionsResponse, err
}

// unexported type that holds implementations of just CleanRoomAssets API methods
type cleanRoomAssetsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomAssetsImpl) Create(ctx context.Context, request CreateCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Asset, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

func (a *cleanRoomAssetsImpl) CreateCleanRoomAssetReview(ctx context.Context, request CreateCleanRoomAssetReviewRequest) (*CreateCleanRoomAssetReviewResponse, error) {
	var createCleanRoomAssetReviewResponse CreateCleanRoomAssetReviewResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v/reviews", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createCleanRoomAssetReviewResponse)
	return &createCleanRoomAssetReviewResponse, err
}

func (a *cleanRoomAssetsImpl) Delete(ctx context.Context, request DeleteCleanRoomAssetRequest) error {
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.Name)
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

func (a *cleanRoomAssetsImpl) Get(ctx context.Context, request GetCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

// List the assets in a clean room. Listed assets do not include their
// asset-type specific details, such as table columns or notebook content; use
// the get method to retrieve them.
func (a *cleanRoomAssetsImpl) List(ctx context.Context, request ListCleanRoomAssetsRequest) listing.Iterator[CleanRoomAsset] {

	getNextPage := func(ctx context.Context, req ListCleanRoomAssetsRequest) (*ListCleanRoomAssetsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomAssetsResponse) []CleanRoomAsset {
		return resp.Assets
	}
	getNextReq := func(resp *ListCleanRoomAssetsResponse) *ListCleanRoomAssetsRequest {
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

// List the assets in a clean room. Listed assets do not include their
// asset-type specific details, such as table columns or notebook content; use
// the get method to retrieve them.
func (a *cleanRoomAssetsImpl) ListAll(ctx context.Context, request ListCleanRoomAssetsRequest) ([]CleanRoomAsset, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomAsset](ctx, iterator)
}

func (a *cleanRoomAssetsImpl) internalList(ctx context.Context, request ListCleanRoomAssetsRequest) (*ListCleanRoomAssetsResponse, error) {
	var listCleanRoomAssetsResponse ListCleanRoomAssetsResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCleanRoomAssetsResponse)
	return &listCleanRoomAssetsResponse, err
}

func (a *cleanRoomAssetsImpl) Update(ctx context.Context, request UpdateCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Asset, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

// unexported type that holds implementations of just CleanRoomAutoApprovalRules API methods
type cleanRoomAutoApprovalRulesImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomAutoApprovalRulesImpl) Create(ctx context.Context, request CreateCleanRoomAutoApprovalRuleRequest) (*CleanRoomAutoApprovalRule, error) {
	var cleanRoomAutoApprovalRule CleanRoomAutoApprovalRule
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cleanRoomAutoApprovalRule)
	return &cleanRoomAutoApprovalRule, err
}

func (a *cleanRoomAutoApprovalRulesImpl) Delete(ctx context.Context, request DeleteCleanRoomAutoApprovalRuleRequest) error {
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules/%v", request.CleanRoomName, request.RuleId)
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

func (a *cleanRoomAutoApprovalRulesImpl) Get(ctx context.Context, request GetCleanRoomAutoApprovalRuleRequest) (*CleanRoomAutoApprovalRule, error) {
	var cleanRoomAutoApprovalRule CleanRoomAutoApprovalRule
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules/%v", request.CleanRoomName, request.RuleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &cleanRoomAutoApprovalRule)
	return &cleanRoomAutoApprovalRule, err
}

// List all auto-approval rules for the caller
func (a *cleanRoomAutoApprovalRulesImpl) List(ctx context.Context, request ListCleanRoomAutoApprovalRulesRequest) listing.Iterator[CleanRoomAutoApprovalRule] {

	getNextPage := func(ctx context.Context, req ListCleanRoomAutoApprovalRulesRequest) (*ListCleanRoomAutoApprovalRulesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomAutoApprovalRulesResponse) []CleanRoomAutoApprovalRule {
		return resp.Rules
	}
	getNextReq := func(resp *ListCleanRoomAutoApprovalRulesResponse) *ListCleanRoomAutoApprovalRulesRequest {
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

// List all auto-approval rules for the caller
func (a *cleanRoomAutoApprovalRulesImpl) ListAll(ctx context.Context, request ListCleanRoomAutoApprovalRulesRequest) ([]CleanRoomAutoApprovalRule, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomAutoApprovalRule](ctx, iterator)
}

func (a *cleanRoomAutoApprovalRulesImpl) internalList(ctx context.Context, request ListCleanRoomAutoApprovalRulesRequest) (*ListCleanRoomAutoApprovalRulesResponse, error) {
	var listCleanRoomAutoApprovalRulesResponse ListCleanRoomAutoApprovalRulesResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCleanRoomAutoApprovalRulesResponse)
	return &listCleanRoomAutoApprovalRulesResponse, err
}

func (a *cleanRoomAutoApprovalRulesImpl) Update(ctx context.Context, request UpdateCleanRoomAutoApprovalRuleRequest) (*CleanRoomAutoApprovalRule, error) {
	var cleanRoomAutoApprovalRule CleanRoomAutoApprovalRule
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules/%v", request.CleanRoomName, request.RuleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.AutoApprovalRule, &cleanRoomAutoApprovalRule)
	return &cleanRoomAutoApprovalRule, err
}

// unexported type that holds implementations of just CleanRoomTaskRuns API methods
type cleanRoomTaskRunsImpl struct {
	client *client.DatabricksClient
}

// List all the historical notebook task runs in a clean room.
func (a *cleanRoomTaskRunsImpl) List(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) listing.Iterator[CleanRoomNotebookTaskRun] {

	getNextPage := func(ctx context.Context, req ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomNotebookTaskRunsResponse) []CleanRoomNotebookTaskRun {
		return resp.Runs
	}
	getNextReq := func(resp *ListCleanRoomNotebookTaskRunsResponse) *ListCleanRoomNotebookTaskRunsRequest {
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

// List all the historical notebook task runs in a clean room.
func (a *cleanRoomTaskRunsImpl) ListAll(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) ([]CleanRoomNotebookTaskRun, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomNotebookTaskRun](ctx, iterator)
}

func (a *cleanRoomTaskRunsImpl) internalList(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error) {
	var listCleanRoomNotebookTaskRunsResponse ListCleanRoomNotebookTaskRunsResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/runs", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCleanRoomNotebookTaskRunsResponse)
	return &listCleanRoomNotebookTaskRunsResponse, err
}

// List all the historical task runs in a clean room.
func (a *cleanRoomTaskRunsImpl) ListCleanRoomTaskRunsHandler(ctx context.Context, request ListCleanRoomTaskRunsRequest) listing.Iterator[CleanRoomTaskRun] {

	getNextPage := func(ctx context.Context, req ListCleanRoomTaskRunsRequest) (*ListCleanRoomTaskRunsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCleanRoomTaskRunsHandler(ctx, req)
	}
	getItems := func(resp *ListCleanRoomTaskRunsResponse) []CleanRoomTaskRun {
		return resp.Runs
	}
	getNextReq := func(resp *ListCleanRoomTaskRunsResponse) *ListCleanRoomTaskRunsRequest {
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

// List all the historical task runs in a clean room.
func (a *cleanRoomTaskRunsImpl) ListCleanRoomTaskRunsHandlerAll(ctx context.Context, request ListCleanRoomTaskRunsRequest) ([]CleanRoomTaskRun, error) {
	iterator := a.ListCleanRoomTaskRunsHandler(ctx, request)
	return listing.ToSlice[CleanRoomTaskRun](ctx, iterator)
}

func (a *cleanRoomTaskRunsImpl) internalListCleanRoomTaskRunsHandler(ctx context.Context, request ListCleanRoomTaskRunsRequest) (*ListCleanRoomTaskRunsResponse, error) {
	var listCleanRoomTaskRunsResponse ListCleanRoomTaskRunsResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/task-runs", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCleanRoomTaskRunsResponse)
	return &listCleanRoomTaskRunsResponse, err
}

// unexported type that holds implementations of just CleanRooms API methods
type cleanRoomsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomsImpl) Create(ctx context.Context, request CreateCleanRoomRequest) (*CleanRoom, error) {
	var cleanRoom CleanRoom
	path := "/api/2.0/clean-rooms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.CleanRoom, &cleanRoom)
	return &cleanRoom, err
}

func (a *cleanRoomsImpl) CreateOutputCatalog(ctx context.Context, request CreateCleanRoomOutputCatalogRequest) (*CreateCleanRoomOutputCatalogResponse, error) {
	var createCleanRoomOutputCatalogResponse CreateCleanRoomOutputCatalogResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/output-catalogs", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.OutputCatalog, &createCleanRoomOutputCatalogResponse)
	return &createCleanRoomOutputCatalogResponse, err
}

func (a *cleanRoomsImpl) Delete(ctx context.Context, request DeleteCleanRoomRequest) error {
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
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

func (a *cleanRoomsImpl) Get(ctx context.Context, request GetCleanRoomRequest) (*CleanRoom, error) {
	var cleanRoom CleanRoom
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &cleanRoom)
	return &cleanRoom, err
}

// Get a list of all clean rooms of the metastore. Only clean rooms the caller
// has access to are returned.
func (a *cleanRoomsImpl) List(ctx context.Context, request ListCleanRoomsRequest) listing.Iterator[CleanRoom] {

	getNextPage := func(ctx context.Context, req ListCleanRoomsRequest) (*ListCleanRoomsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomsResponse) []CleanRoom {
		return resp.CleanRooms
	}
	getNextReq := func(resp *ListCleanRoomsResponse) *ListCleanRoomsRequest {
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

// Get a list of all clean rooms of the metastore. Only clean rooms the caller
// has access to are returned.
func (a *cleanRoomsImpl) ListAll(ctx context.Context, request ListCleanRoomsRequest) ([]CleanRoom, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoom](ctx, iterator)
}

func (a *cleanRoomsImpl) internalList(ctx context.Context, request ListCleanRoomsRequest) (*ListCleanRoomsResponse, error) {
	var listCleanRoomsResponse ListCleanRoomsResponse
	path := "/api/2.0/clean-rooms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCleanRoomsResponse)
	return &listCleanRoomsResponse, err
}

func (a *cleanRoomsImpl) Update(ctx context.Context, request UpdateCleanRoomRequest) (*CleanRoom, error) {
	var cleanRoom CleanRoom
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &cleanRoom)
	return &cleanRoom, err
}
