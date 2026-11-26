// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

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

// unexported type that holds implementations of just Alerts API methods
type alertsImpl struct {
	client *client.DatabricksClient
}

func (a *alertsImpl) Create(ctx context.Context, request CreateAlertRequest) (*Alert, error) {
	var alert Alert
	path := "/api/2.0/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &alert)
	return &alert, err
}

func (a *alertsImpl) Delete(ctx context.Context, request TrashAlertRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
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

func (a *alertsImpl) Get(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &alert)
	return &alert, err
}

// Gets a list of alerts accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *alertsImpl) List(ctx context.Context, request ListAlertsRequest) listing.Iterator[ListAlertsResponseAlert] {

	getNextPage := func(ctx context.Context, req ListAlertsRequest) (*ListAlertsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAlertsResponse) []ListAlertsResponseAlert {
		return resp.Results
	}
	getNextReq := func(resp *ListAlertsResponse) *ListAlertsRequest {
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

// Gets a list of alerts accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *alertsImpl) ListAll(ctx context.Context, request ListAlertsRequest) ([]ListAlertsResponseAlert, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListAlertsResponseAlert](ctx, iterator)
}

func (a *alertsImpl) internalList(ctx context.Context, request ListAlertsRequest) (*ListAlertsResponse, error) {
	var listAlertsResponse ListAlertsResponse
	path := "/api/2.0/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAlertsResponse)
	return &listAlertsResponse, err
}

func (a *alertsImpl) Update(ctx context.Context, request UpdateAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &alert)
	return &alert, err
}

// unexported type that holds implementations of just AlertsLegacy API methods
type alertsLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *alertsLegacyImpl) Create(ctx context.Context, request CreateAlert) (*LegacyAlert, error) {
	var legacyAlert LegacyAlert
	path := "/api/2.0/preview/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyAlert)
	return &legacyAlert, err
}

func (a *alertsLegacyImpl) Delete(ctx context.Context, request DeleteAlertsLegacyRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
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

func (a *alertsLegacyImpl) Get(ctx context.Context, request GetAlertsLegacyRequest) (*LegacyAlert, error) {
	var legacyAlert LegacyAlert
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &legacyAlert)
	return &legacyAlert, err
}

func (a *alertsLegacyImpl) List(ctx context.Context) ([]LegacyAlert, error) {
	var legacyAlertList []LegacyAlert
	path := "/api/2.0/preview/sql/alerts"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &legacyAlertList)
	return legacyAlertList, err
}

func (a *alertsLegacyImpl) Update(ctx context.Context, request EditAlert) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just AlertsV2 API methods
type alertsV2Impl struct {
	client *client.DatabricksClient
}

func (a *alertsV2Impl) CreateAlert(ctx context.Context, request CreateAlertV2Request) (*AlertV2, error) {
	var alertV2 AlertV2
	path := "/api/2.0/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Alert, &alertV2)
	return &alertV2, err
}

func (a *alertsV2Impl) GetAlert(ctx context.Context, request GetAlertV2Request) (*AlertV2, error) {
	var alertV2 AlertV2
	path := fmt.Sprintf("/api/2.0/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &alertV2)
	return &alertV2, err
}

// Gets a list of alerts accessible to the user, ordered by creation time.
func (a *alertsV2Impl) ListAlerts(ctx context.Context, request ListAlertsV2Request) listing.Iterator[AlertV2] {

	getNextPage := func(ctx context.Context, req ListAlertsV2Request) (*ListAlertsV2Response, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListAlerts(ctx, req)
	}
	getItems := func(resp *ListAlertsV2Response) []AlertV2 {
		return resp.Alerts
	}
	getNextReq := func(resp *ListAlertsV2Response) *ListAlertsV2Request {
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

// Gets a list of alerts accessible to the user, ordered by creation time.
func (a *alertsV2Impl) ListAlertsAll(ctx context.Context, request ListAlertsV2Request) ([]AlertV2, error) {
	iterator := a.ListAlerts(ctx, request)
	return listing.ToSlice[AlertV2](ctx, iterator)
}

func (a *alertsV2Impl) internalListAlerts(ctx context.Context, request ListAlertsV2Request) (*ListAlertsV2Response, error) {
	var listAlertsV2Response ListAlertsV2Response
	path := "/api/2.0/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAlertsV2Response)
	return &listAlertsV2Response, err
}

func (a *alertsV2Impl) TrashAlert(ctx context.Context, request TrashAlertV2Request) error {
	path := fmt.Sprintf("/api/2.0/alerts/%v", request.Id)
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

func (a *alertsV2Impl) UpdateAlert(ctx context.Context, request UpdateAlertV2Request) (*AlertV2, error) {
	var alertV2 AlertV2
	path := fmt.Sprintf("/api/2.0/alerts/%v", request.Id)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Alert, &alertV2)
	return &alertV2, err
}

// unexported type that holds implementations of just DashboardWidgets API methods
type dashboardWidgetsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardWidgetsImpl) Create(ctx context.Context, request CreateWidget) (*Widget, error) {
	var widget Widget
	path := "/api/2.0/preview/sql/widgets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &widget)
	return &widget, err
}

func (a *dashboardWidgetsImpl) Delete(ctx context.Context, request DeleteDashboardWidgetRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", request.Id)
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

func (a *dashboardWidgetsImpl) Update(ctx context.Context, request UpdateWidgetRequest) (*Widget, error) {
	var widget Widget
	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &widget)
	return &widget, err
}

// unexported type that holds implementations of just Dashboards API methods
type dashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardsImpl) Delete(ctx context.Context, request DeleteDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
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

func (a *dashboardsImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

// Fetch a paginated list of dashboard objects.
//
// **Warning**: Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
//
// **Warning**: This API is deprecated. Please use the AI/BI Dashboards API
// instead. [Learn more]
//
// [Learn more]: https://docs.databricks.com/en/dashboards/
func (a *dashboardsImpl) List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard] {

	request.Page = 1 // start iterating from the first page

	getNextPage := func(ctx context.Context, req ListDashboardsRequest) (*ListResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListResponse) []Dashboard {
		return resp.Results
	}
	getNextReq := func(resp *ListResponse) *ListDashboardsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.Page = resp.Page + 1
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Fetch a paginated list of dashboard objects.
//
// **Warning**: Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
//
// **Warning**: This API is deprecated. Please use the AI/BI Dashboards API
// instead. [Learn more]
//
// [Learn more]: https://docs.databricks.com/en/dashboards/
func (a *dashboardsImpl) ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Dashboard, int](ctx, iterator, request.PageSize)

}

func (a *dashboardsImpl) internalList(ctx context.Context, request ListDashboardsRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/preview/sql/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listResponse)
	return &listResponse, err
}

func (a *dashboardsImpl) Restore(ctx context.Context, request RestoreDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/trash/%v", request.DashboardId)
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

func (a *dashboardsImpl) Update(ctx context.Context, request DashboardEditContent) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

// unexported type that holds implementations of just DataSources API methods
type dataSourcesImpl struct {
	client *client.DatabricksClient
}

func (a *dataSourcesImpl) List(ctx context.Context) ([]DataSource, error) {
	var dataSourceList []DataSource
	path := "/api/2.0/preview/sql/data_sources"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &dataSourceList)
	return dataSourceList, err
}

// unexported type that holds implementations of just DbsqlPermissions API methods
type dbsqlPermissionsImpl struct {
	client *client.DatabricksClient
}

func (a *dbsqlPermissionsImpl) Get(ctx context.Context, request GetDbsqlPermissionRequest) (*GetResponse, error) {
	var getResponse GetResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getResponse)
	return &getResponse, err
}

func (a *dbsqlPermissionsImpl) Set(ctx context.Context, request SetRequest) (*SetResponse, error) {
	var setResponse SetResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &setResponse)
	return &setResponse, err
}

func (a *dbsqlPermissionsImpl) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	var success Success
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v/transfer", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &success)
	return &success, err
}

// unexported type that holds implementations of just Queries API methods
type queriesImpl struct {
	client *client.DatabricksClient
}

func (a *queriesImpl) Create(ctx context.Context, request CreateQueryRequest) (*Query, error) {
	var query Query
	path := "/api/2.0/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &query)
	return &query, err
}

func (a *queriesImpl) Delete(ctx context.Context, request TrashQueryRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
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

func (a *queriesImpl) Get(ctx context.Context, request GetQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &query)
	return &query, err
}

// Gets a list of queries accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *queriesImpl) List(ctx context.Context, request ListQueriesRequest) listing.Iterator[ListQueryObjectsResponseQuery] {

	getNextPage := func(ctx context.Context, req ListQueriesRequest) (*ListQueryObjectsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListQueryObjectsResponse) []ListQueryObjectsResponseQuery {
		return resp.Results
	}
	getNextReq := func(resp *ListQueryObjectsResponse) *ListQueriesRequest {
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

// Gets a list of queries accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *queriesImpl) ListAll(ctx context.Context, request ListQueriesRequest) ([]ListQueryObjectsResponseQuery, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListQueryObjectsResponseQuery](ctx, iterator)
}

func (a *queriesImpl) internalList(ctx context.Context, request ListQueriesRequest) (*ListQueryObjectsResponse, error) {
	var listQueryObjectsResponse ListQueryObjectsResponse
	path := "/api/2.0/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listQueryObjectsResponse)
	return &listQueryObjectsResponse, err
}

// Gets a list of visualizations on a query.
func (a *queriesImpl) ListVisualizations(ctx context.Context, request ListVisualizationsForQueryRequest) listing.Iterator[Visualization] {

	getNextPage := func(ctx context.Context, req ListVisualizationsForQueryRequest) (*ListVisualizationsForQueryResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListVisualizations(ctx, req)
	}
	getItems := func(resp *ListVisualizationsForQueryResponse) []Visualization {
		return resp.Results
	}
	getNextReq := func(resp *ListVisualizationsForQueryResponse) *ListVisualizationsForQueryRequest {
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

// Gets a list of visualizations on a query.
func (a *queriesImpl) ListVisualizationsAll(ctx context.Context, request ListVisualizationsForQueryRequest) ([]Visualization, error) {
	iterator := a.ListVisualizations(ctx, request)
	return listing.ToSlice[Visualization](ctx, iterator)
}

func (a *queriesImpl) internalListVisualizations(ctx context.Context, request ListVisualizationsForQueryRequest) (*ListVisualizationsForQueryResponse, error) {
	var listVisualizationsForQueryResponse ListVisualizationsForQueryResponse
	path := fmt.Sprintf("/api/2.0/sql/queries/%v/visualizations", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listVisualizationsForQueryResponse)
	return &listVisualizationsForQueryResponse, err
}

func (a *queriesImpl) Update(ctx context.Context, request UpdateQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &query)
	return &query, err
}

// unexported type that holds implementations of just QueriesLegacy API methods
type queriesLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *queriesLegacyImpl) Create(ctx context.Context, request QueryPostContent) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := "/api/2.0/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyQuery)
	return &legacyQuery, err
}

func (a *queriesLegacyImpl) Delete(ctx context.Context, request DeleteQueriesLegacyRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
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

func (a *queriesLegacyImpl) Get(ctx context.Context, request GetQueriesLegacyRequest) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &legacyQuery)
	return &legacyQuery, err
}

// Gets a list of queries. Optionally, this list can be filtered by a search
// term.
//
// **Warning**: Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
//
// **Warning**: This API is deprecated. Please use :method:queries/list instead.
// [Learn more]
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
func (a *queriesLegacyImpl) List(ctx context.Context, request ListQueriesLegacyRequest) listing.Iterator[LegacyQuery] {

	request.Page = 1 // start iterating from the first page

	getNextPage := func(ctx context.Context, req ListQueriesLegacyRequest) (*QueryList, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *QueryList) []LegacyQuery {
		return resp.Results
	}
	getNextReq := func(resp *QueryList) *ListQueriesLegacyRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.Page = resp.Page + 1
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Gets a list of queries. Optionally, this list can be filtered by a search
// term.
//
// **Warning**: Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
//
// **Warning**: This API is deprecated. Please use :method:queries/list instead.
// [Learn more]
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
func (a *queriesLegacyImpl) ListAll(ctx context.Context, request ListQueriesLegacyRequest) ([]LegacyQuery, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[LegacyQuery, int](ctx, iterator, request.PageSize)

}

func (a *queriesLegacyImpl) internalList(ctx context.Context, request ListQueriesLegacyRequest) (*QueryList, error) {
	var queryList QueryList
	path := "/api/2.0/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &queryList)
	return &queryList, err
}

func (a *queriesLegacyImpl) Restore(ctx context.Context, request RestoreQueriesLegacyRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/trash/%v", request.QueryId)
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

func (a *queriesLegacyImpl) Update(ctx context.Context, request QueryEditContent) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyQuery)
	return &legacyQuery, err
}

// unexported type that holds implementations of just QueryHistory API methods
type queryHistoryImpl struct {
	client *client.DatabricksClient
}

func (a *queryHistoryImpl) List(ctx context.Context, request ListQueryHistoryRequest) (*ListQueriesResponse, error) {
	var listQueriesResponse ListQueriesResponse
	path := "/api/2.0/sql/history/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listQueriesResponse)
	return &listQueriesResponse, err
}

// unexported type that holds implementations of just QueryVisualizations API methods
type queryVisualizationsImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsImpl) Create(ctx context.Context, request CreateVisualizationRequest) (*Visualization, error) {
	var visualization Visualization
	path := "/api/2.0/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &visualization)
	return &visualization, err
}

func (a *queryVisualizationsImpl) Delete(ctx context.Context, request DeleteVisualizationRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", request.Id)
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

func (a *queryVisualizationsImpl) Update(ctx context.Context, request UpdateVisualizationRequest) (*Visualization, error) {
	var visualization Visualization
	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &visualization)
	return &visualization, err
}

// unexported type that holds implementations of just QueryVisualizationsLegacy API methods
type queryVisualizationsLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsLegacyImpl) Create(ctx context.Context, request CreateQueryVisualizationsLegacyRequest) (*LegacyVisualization, error) {
	var legacyVisualization LegacyVisualization
	path := "/api/2.0/preview/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyVisualization)
	return &legacyVisualization, err
}

func (a *queryVisualizationsLegacyImpl) Delete(ctx context.Context, request DeleteQueryVisualizationsLegacyRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", request.Id)
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

func (a *queryVisualizationsLegacyImpl) Update(ctx context.Context, request LegacyVisualization) (*LegacyVisualization, error) {
	var legacyVisualization LegacyVisualization
	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyVisualization)
	return &legacyVisualization, err
}

// unexported type that holds implementations of just RedashConfig API methods
type redashConfigImpl struct {
	client *client.DatabricksClient
}

func (a *redashConfigImpl) GetConfig(ctx context.Context) (*ClientConfig, error) {
	var clientConfig ClientConfig
	path := "/api/2.0/redash-v2/config"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &clientConfig)
	return &clientConfig, err
}

// unexported type that holds implementations of just StatementExecution API methods
type statementExecutionImpl struct {
	client *client.DatabricksClient
}

func (a *statementExecutionImpl) CancelExecution(ctx context.Context, request CancelExecutionRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/cancel", request.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, nil)
	return err
}

func (a *statementExecutionImpl) ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := "/api/2.0/sql/statements"
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

func (a *statementExecutionImpl) GetStatement(ctx context.Context, request GetStatementRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := fmt.Sprintf("/api/2.0/sql/statements/%v", request.StatementId)
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

func (a *statementExecutionImpl) GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error) {
	var resultData ResultData
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/result/chunks/%v", request.StatementId, request.ChunkIndex)
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

// unexported type that holds implementations of just Warehouses API methods
type warehousesImpl struct {
	client *client.DatabricksClient
}

func (a *warehousesImpl) Create(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	var createWarehouseResponse CreateWarehouseResponse
	path := "/api/2.0/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createWarehouseResponse)
	return &createWarehouseResponse, err
}

func (a *warehousesImpl) CreateDefaultWarehouseOverride(ctx context.Context, request CreateDefaultWarehouseOverrideRequest) (*DefaultWarehouseOverride, error) {
	var defaultWarehouseOverride DefaultWarehouseOverride
	path := "/api/warehouses/v1/default-warehouse-overrides"
	queryParams := make(map[string]any)

	if request.DefaultWarehouseOverrideId != "" {
		queryParams["default_warehouse_override_id"] = request.DefaultWarehouseOverrideId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DefaultWarehouseOverride, &defaultWarehouseOverride)
	return &defaultWarehouseOverride, err
}

func (a *warehousesImpl) Delete(ctx context.Context, request DeleteWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
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

func (a *warehousesImpl) DeleteDefaultWarehouseOverride(ctx context.Context, request DeleteDefaultWarehouseOverrideRequest) error {
	path := fmt.Sprintf("/api/warehouses/v1/%v", request.Name)
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

func (a *warehousesImpl) Edit(ctx context.Context, request EditWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", request.Id)
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

func (a *warehousesImpl) Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	var getWarehouseResponse GetWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getWarehouseResponse)
	return &getWarehouseResponse, err
}

func (a *warehousesImpl) GetDefaultWarehouseOverride(ctx context.Context, request GetDefaultWarehouseOverrideRequest) (*DefaultWarehouseOverride, error) {
	var defaultWarehouseOverride DefaultWarehouseOverride
	path := fmt.Sprintf("/api/warehouses/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &defaultWarehouseOverride)
	return &defaultWarehouseOverride, err
}

func (a *warehousesImpl) GetPermissionLevels(ctx context.Context, request GetWarehousePermissionLevelsRequest) (*GetWarehousePermissionLevelsResponse, error) {
	var getWarehousePermissionLevelsResponse GetWarehousePermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v/permissionLevels", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getWarehousePermissionLevelsResponse)
	return &getWarehousePermissionLevelsResponse, err
}

func (a *warehousesImpl) GetPermissions(ctx context.Context, request GetWarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &warehousePermissions)
	return &warehousePermissions, err
}

func (a *warehousesImpl) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {
	var getWorkspaceWarehouseConfigResponse GetWorkspaceWarehouseConfigResponse
	path := "/api/2.0/sql/config/warehouses"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getWorkspaceWarehouseConfigResponse)
	return &getWorkspaceWarehouseConfigResponse, err
}

// Lists all SQL warehouses that a user has access to.
func (a *warehousesImpl) List(ctx context.Context, request ListWarehousesRequest) listing.Iterator[EndpointInfo] {

	getNextPage := func(ctx context.Context, req ListWarehousesRequest) (*ListWarehousesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListWarehousesResponse) []EndpointInfo {
		return resp.Warehouses
	}
	getNextReq := func(resp *ListWarehousesResponse) *ListWarehousesRequest {
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

// Lists all SQL warehouses that a user has access to.
func (a *warehousesImpl) ListAll(ctx context.Context, request ListWarehousesRequest) ([]EndpointInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[EndpointInfo](ctx, iterator)
}

func (a *warehousesImpl) internalList(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {
	var listWarehousesResponse ListWarehousesResponse
	path := "/api/2.0/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWarehousesResponse)
	return &listWarehousesResponse, err
}

// Lists all default warehouse overrides in the workspace. Only workspace
// administrators can list all overrides.
func (a *warehousesImpl) ListDefaultWarehouseOverrides(ctx context.Context, request ListDefaultWarehouseOverridesRequest) listing.Iterator[DefaultWarehouseOverride] {

	getNextPage := func(ctx context.Context, req ListDefaultWarehouseOverridesRequest) (*ListDefaultWarehouseOverridesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDefaultWarehouseOverrides(ctx, req)
	}
	getItems := func(resp *ListDefaultWarehouseOverridesResponse) []DefaultWarehouseOverride {
		return resp.DefaultWarehouseOverrides
	}
	getNextReq := func(resp *ListDefaultWarehouseOverridesResponse) *ListDefaultWarehouseOverridesRequest {
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

// Lists all default warehouse overrides in the workspace. Only workspace
// administrators can list all overrides.
func (a *warehousesImpl) ListDefaultWarehouseOverridesAll(ctx context.Context, request ListDefaultWarehouseOverridesRequest) ([]DefaultWarehouseOverride, error) {
	iterator := a.ListDefaultWarehouseOverrides(ctx, request)
	return listing.ToSlice[DefaultWarehouseOverride](ctx, iterator)
}

func (a *warehousesImpl) internalListDefaultWarehouseOverrides(ctx context.Context, request ListDefaultWarehouseOverridesRequest) (*ListDefaultWarehouseOverridesResponse, error) {
	var listDefaultWarehouseOverridesResponse ListDefaultWarehouseOverridesResponse
	path := "/api/warehouses/v1/default-warehouse-overrides"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDefaultWarehouseOverridesResponse)
	return &listDefaultWarehouseOverridesResponse, err
}

func (a *warehousesImpl) SetPermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &warehousePermissions)
	return &warehousePermissions, err
}

func (a *warehousesImpl) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	path := "/api/2.0/sql/config/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

func (a *warehousesImpl) Start(ctx context.Context, request StartRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/start", request.Id)
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

func (a *warehousesImpl) Stop(ctx context.Context, request StopRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/stop", request.Id)
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

func (a *warehousesImpl) UpdateDefaultWarehouseOverride(ctx context.Context, request UpdateDefaultWarehouseOverrideRequest) (*DefaultWarehouseOverride, error) {
	var defaultWarehouseOverride DefaultWarehouseOverride
	path := fmt.Sprintf("/api/warehouses/v1/%v", request.Name)
	queryParams := make(map[string]any)

	if request.AllowMissing != false || slices.Contains(request.ForceSendFields, "AllowMissing") {
		queryParams["allow_missing"] = request.AllowMissing
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.DefaultWarehouseOverride, &defaultWarehouseOverride)
	return &defaultWarehouseOverride, err
}

func (a *warehousesImpl) UpdatePermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &warehousePermissions)
	return &warehousePermissions, err
}
