// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just QualityMonitorV2 API methods
type qualityMonitorV2Impl struct {
	client *client.DatabricksClient
}

func (a *qualityMonitorV2Impl) CreateQualityMonitor(ctx context.Context, request CreateQualityMonitorRequest) (*QualityMonitor, error) {
	var qualityMonitor QualityMonitor
	path := "/api/2.0/quality-monitors"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.QualityMonitor, &qualityMonitor)
	return &qualityMonitor, err
}

func (a *qualityMonitorV2Impl) DeleteQualityMonitor(ctx context.Context, request DeleteQualityMonitorRequest) error {
	path := fmt.Sprintf("/api/2.0/quality-monitors/%v/%v", request.ObjectType, request.ObjectId)
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

func (a *qualityMonitorV2Impl) GetQualityMonitor(ctx context.Context, request GetQualityMonitorRequest) (*QualityMonitor, error) {
	var qualityMonitor QualityMonitor
	path := fmt.Sprintf("/api/2.0/quality-monitors/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &qualityMonitor)
	return &qualityMonitor, err
}

// Deprecated: Use Data Quality Monitoring API instead
// (/api/data-quality/v1/monitors). (Unimplemented) List quality monitors.
func (a *qualityMonitorV2Impl) ListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) listing.Iterator[QualityMonitor] {

	getNextPage := func(ctx context.Context, req ListQualityMonitorRequest) (*ListQualityMonitorResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListQualityMonitor(ctx, req)
	}
	getItems := func(resp *ListQualityMonitorResponse) []QualityMonitor {
		return resp.QualityMonitors
	}
	getNextReq := func(resp *ListQualityMonitorResponse) *ListQualityMonitorRequest {
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

// Deprecated: Use Data Quality Monitoring API instead
// (/api/data-quality/v1/monitors). (Unimplemented) List quality monitors.
func (a *qualityMonitorV2Impl) ListQualityMonitorAll(ctx context.Context, request ListQualityMonitorRequest) ([]QualityMonitor, error) {
	iterator := a.ListQualityMonitor(ctx, request)
	return listing.ToSlice[QualityMonitor](ctx, iterator)
}

func (a *qualityMonitorV2Impl) internalListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) (*ListQualityMonitorResponse, error) {
	var listQualityMonitorResponse ListQualityMonitorResponse
	path := "/api/2.0/quality-monitors"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listQualityMonitorResponse)
	return &listQualityMonitorResponse, err
}

func (a *qualityMonitorV2Impl) UpdateQualityMonitor(ctx context.Context, request UpdateQualityMonitorRequest) (*QualityMonitor, error) {
	var qualityMonitor QualityMonitor
	path := fmt.Sprintf("/api/2.0/quality-monitors/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request.QualityMonitor, &qualityMonitor)
	return &qualityMonitor, err
}
