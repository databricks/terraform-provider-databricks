// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dataquality

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just DataQuality API methods
type dataQualityImpl struct {
	client *client.DatabricksClient
}

func (a *dataQualityImpl) CancelRefresh(ctx context.Context, request CancelRefreshRequest) (*CancelRefreshResponse, error) {
	var cancelRefreshResponse CancelRefreshResponse
	path := fmt.Sprintf("/api/data-quality/v1/monitors/%v/%v/refreshes/%v/cancel", request.ObjectType, request.ObjectId, request.RefreshId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelRefreshResponse)
	return &cancelRefreshResponse, err
}

func (a *dataQualityImpl) CreateMonitor(ctx context.Context, request CreateMonitorRequest) (*Monitor, error) {
	var monitor Monitor
	path := "/api/data-quality/v1/monitors"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Monitor, &monitor)
	return &monitor, err
}

func (a *dataQualityImpl) CreateRefresh(ctx context.Context, request CreateRefreshRequest) (*Refresh, error) {
	var refresh Refresh
	path := fmt.Sprintf("/api/data-quality/v1/monitors/%v/%v/refreshes", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Refresh, &refresh)
	return &refresh, err
}

func (a *dataQualityImpl) DeleteMonitor(ctx context.Context, request DeleteMonitorRequest) error {
	path := fmt.Sprintf("/api/data-quality/v1/monitors/%v/%v", request.ObjectType, request.ObjectId)
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

func (a *dataQualityImpl) DeleteRefresh(ctx context.Context, request DeleteRefreshRequest) error {
	path := fmt.Sprintf("/api/data-quality/v1/monitors/%v/%v/refreshes/%v", request.ObjectType, request.ObjectId, request.RefreshId)
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

func (a *dataQualityImpl) GetMonitor(ctx context.Context, request GetMonitorRequest) (*Monitor, error) {
	var monitor Monitor
	path := fmt.Sprintf("/api/data-quality/v1/monitors/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &monitor)
	return &monitor, err
}

func (a *dataQualityImpl) GetRefresh(ctx context.Context, request GetRefreshRequest) (*Refresh, error) {
	var refresh Refresh
	path := fmt.Sprintf("/api/data-quality/v1/monitors/%v/%v/refreshes/%v", request.ObjectType, request.ObjectId, request.RefreshId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &refresh)
	return &refresh, err
}

// (Unimplemented) List data quality monitors.
func (a *dataQualityImpl) ListMonitor(ctx context.Context, request ListMonitorRequest) listing.Iterator[Monitor] {

	getNextPage := func(ctx context.Context, req ListMonitorRequest) (*ListMonitorResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListMonitor(ctx, req)
	}
	getItems := func(resp *ListMonitorResponse) []Monitor {
		return resp.Monitors
	}
	getNextReq := func(resp *ListMonitorResponse) *ListMonitorRequest {
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

// (Unimplemented) List data quality monitors.
func (a *dataQualityImpl) ListMonitorAll(ctx context.Context, request ListMonitorRequest) ([]Monitor, error) {
	iterator := a.ListMonitor(ctx, request)
	return listing.ToSlice[Monitor](ctx, iterator)
}

func (a *dataQualityImpl) internalListMonitor(ctx context.Context, request ListMonitorRequest) (*ListMonitorResponse, error) {
	var listMonitorResponse ListMonitorResponse
	path := "/api/data-quality/v1/monitors"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listMonitorResponse)
	return &listMonitorResponse, err
}

// List data quality monitor refreshes. The call must be made in the same
// workspace as where the monitor was created.
//
// For the `table` `object_type`, the caller must have either of the following
// sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the table's parent
// catalog. 2. **USE_CATALOG** on the table's parent catalog, and **MANAGE** and
// **USE_SCHEMA** on the table's parent schema. 3. **USE_CATALOG** on the
// table's parent catalog, **USE_SCHEMA** on the table's parent schema, and
// **SELECT** on the table.
//
// For the `schema` `object_type`, the caller must have either of the following
// sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the schema's parent
// catalog. 2. **USE_CATALOG** on the schema's parent catalog, and
// **USE_SCHEMA** on the schema.
func (a *dataQualityImpl) ListRefresh(ctx context.Context, request ListRefreshRequest) listing.Iterator[Refresh] {

	getNextPage := func(ctx context.Context, req ListRefreshRequest) (*ListRefreshResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListRefresh(ctx, req)
	}
	getItems := func(resp *ListRefreshResponse) []Refresh {
		return resp.Refreshes
	}
	getNextReq := func(resp *ListRefreshResponse) *ListRefreshRequest {
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

// List data quality monitor refreshes. The call must be made in the same
// workspace as where the monitor was created.
//
// For the `table` `object_type`, the caller must have either of the following
// sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the table's parent
// catalog. 2. **USE_CATALOG** on the table's parent catalog, and **MANAGE** and
// **USE_SCHEMA** on the table's parent schema. 3. **USE_CATALOG** on the
// table's parent catalog, **USE_SCHEMA** on the table's parent schema, and
// **SELECT** on the table.
//
// For the `schema` `object_type`, the caller must have either of the following
// sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the schema's parent
// catalog. 2. **USE_CATALOG** on the schema's parent catalog, and
// **USE_SCHEMA** on the schema.
func (a *dataQualityImpl) ListRefreshAll(ctx context.Context, request ListRefreshRequest) ([]Refresh, error) {
	iterator := a.ListRefresh(ctx, request)
	return listing.ToSlice[Refresh](ctx, iterator)
}

func (a *dataQualityImpl) internalListRefresh(ctx context.Context, request ListRefreshRequest) (*ListRefreshResponse, error) {
	var listRefreshResponse ListRefreshResponse
	path := fmt.Sprintf("/api/data-quality/v1/monitors/%v/%v/refreshes", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRefreshResponse)
	return &listRefreshResponse, err
}

func (a *dataQualityImpl) UpdateMonitor(ctx context.Context, request UpdateMonitorRequest) (*Monitor, error) {
	var monitor Monitor
	path := fmt.Sprintf("/api/data-quality/v1/monitors/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Monitor, &monitor)
	return &monitor, err
}

func (a *dataQualityImpl) UpdateRefresh(ctx context.Context, request UpdateRefreshRequest) (*Refresh, error) {
	var refresh Refresh
	path := fmt.Sprintf("/api/data-quality/v1/monitors/%v/%v/refreshes/%v", request.ObjectType, request.ObjectId, request.RefreshId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Refresh, &refresh)
	return &refresh, err
}
