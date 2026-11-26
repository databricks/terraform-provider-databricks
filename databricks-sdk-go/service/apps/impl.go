// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

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

// unexported type that holds implementations of just Apps API methods
type appsImpl struct {
	client *client.DatabricksClient
}

func (a *appsImpl) Create(ctx context.Context, request CreateAppRequest) (*App, error) {
	var app App
	path := "/api/2.0/apps"
	queryParams := make(map[string]any)

	if request.NoCompute != false || slices.Contains(request.ForceSendFields, "NoCompute") {
		queryParams["no_compute"] = request.NoCompute
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.App, &app)
	return &app, err
}

func (a *appsImpl) CreateSpace(ctx context.Context, request CreateSpaceRequest) (*Operation, error) {
	var operation Operation
	path := "/api/2.0/app-spaces"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Space, &operation)
	return &operation, err
}

func (a *appsImpl) CreateUpdate(ctx context.Context, request AsyncUpdateAppRequest) (*AppUpdate, error) {
	var appUpdate AppUpdate
	path := fmt.Sprintf("/api/2.0/apps/%v/update", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &appUpdate)
	return &appUpdate, err
}

func (a *appsImpl) Delete(ctx context.Context, request DeleteAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsImpl) DeleteAppThumbnail(ctx context.Context, request DeleteAppThumbnailRequest) error {
	path := fmt.Sprintf("/api/2.0/apps/%v/thumbnail", request.Name)
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

func (a *appsImpl) DeleteSpace(ctx context.Context, request DeleteSpaceRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/app-spaces/%v", request.Name)
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

func (a *appsImpl) Deploy(ctx context.Context, request CreateAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.AppDeployment, &appDeployment)
	return &appDeployment, err
}

func (a *appsImpl) Get(ctx context.Context, request GetAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsImpl) GetAppDeploymentScanFinding(ctx context.Context, request GetAppDeploymentScanFindingRequest) (*AppDeploymentScanFinding, error) {
	var appDeploymentScanFinding AppDeploymentScanFinding
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &appDeploymentScanFinding)
	return &appDeploymentScanFinding, err
}

func (a *appsImpl) GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments/%v", request.AppName, request.DeploymentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &appDeployment)
	return &appDeployment, err
}

func (a *appsImpl) GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error) {
	var getAppPermissionLevelsResponse GetAppPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v/permissionLevels", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getAppPermissionLevelsResponse)
	return &getAppPermissionLevelsResponse, err
}

func (a *appsImpl) GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &appPermissions)
	return &appPermissions, err
}

func (a *appsImpl) GetSpace(ctx context.Context, request GetSpaceRequest) (*Space, error) {
	var space Space
	path := fmt.Sprintf("/api/2.0/app-spaces/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &space)
	return &space, err
}

func (a *appsImpl) GetSpaceOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/app-spaces/%v/operation", request.Name)
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

func (a *appsImpl) GetUpdate(ctx context.Context, request GetAppUpdateRequest) (*AppUpdate, error) {
	var appUpdate AppUpdate
	path := fmt.Sprintf("/api/2.0/apps/%v/update", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &appUpdate)
	return &appUpdate, err
}

// Lists all apps in the workspace.
func (a *appsImpl) List(ctx context.Context, request ListAppsRequest) listing.Iterator[App] {

	getNextPage := func(ctx context.Context, req ListAppsRequest) (*ListAppsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAppsResponse) []App {
		return resp.Apps
	}
	getNextReq := func(resp *ListAppsResponse) *ListAppsRequest {
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

// Lists all apps in the workspace.
func (a *appsImpl) ListAll(ctx context.Context, request ListAppsRequest) ([]App, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[App](ctx, iterator)
}

func (a *appsImpl) internalList(ctx context.Context, request ListAppsRequest) (*ListAppsResponse, error) {
	var listAppsResponse ListAppsResponse
	path := "/api/2.0/apps"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAppsResponse)
	return &listAppsResponse, err
}

// Lists the security scan findings for an app deployment.
func (a *appsImpl) ListAppDeploymentScanFindings(ctx context.Context, request ListAppDeploymentScanFindingsRequest) listing.Iterator[AppDeploymentScanFinding] {

	getNextPage := func(ctx context.Context, req ListAppDeploymentScanFindingsRequest) (*ListAppDeploymentScanFindingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListAppDeploymentScanFindings(ctx, req)
	}
	getItems := func(resp *ListAppDeploymentScanFindingsResponse) []AppDeploymentScanFinding {
		return resp.AppDeploymentScanFindings
	}
	getNextReq := func(resp *ListAppDeploymentScanFindingsResponse) *ListAppDeploymentScanFindingsRequest {
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

// Lists the security scan findings for an app deployment.
func (a *appsImpl) ListAppDeploymentScanFindingsAll(ctx context.Context, request ListAppDeploymentScanFindingsRequest) ([]AppDeploymentScanFinding, error) {
	iterator := a.ListAppDeploymentScanFindings(ctx, request)
	return listing.ToSlice[AppDeploymentScanFinding](ctx, iterator)
}

func (a *appsImpl) internalListAppDeploymentScanFindings(ctx context.Context, request ListAppDeploymentScanFindingsRequest) (*ListAppDeploymentScanFindingsResponse, error) {
	var listAppDeploymentScanFindingsResponse ListAppDeploymentScanFindingsResponse
	path := fmt.Sprintf("/api/2.0/%v/scanFindings", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAppDeploymentScanFindingsResponse)
	return &listAppDeploymentScanFindingsResponse, err
}

// Lists all app deployments for the app with the supplied name. The history
// holds only the app's most recent deployments.
func (a *appsImpl) ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) listing.Iterator[AppDeployment] {

	getNextPage := func(ctx context.Context, req ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDeployments(ctx, req)
	}
	getItems := func(resp *ListAppDeploymentsResponse) []AppDeployment {
		return resp.AppDeployments
	}
	getNextReq := func(resp *ListAppDeploymentsResponse) *ListAppDeploymentsRequest {
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

// Lists all app deployments for the app with the supplied name. The history
// holds only the app's most recent deployments.
func (a *appsImpl) ListDeploymentsAll(ctx context.Context, request ListAppDeploymentsRequest) ([]AppDeployment, error) {
	iterator := a.ListDeployments(ctx, request)
	return listing.ToSlice[AppDeployment](ctx, iterator)
}

func (a *appsImpl) internalListDeployments(ctx context.Context, request ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error) {
	var listAppDeploymentsResponse ListAppDeploymentsResponse
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAppDeploymentsResponse)
	return &listAppDeploymentsResponse, err
}

// Lists all app spaces in the workspace.
func (a *appsImpl) ListSpaces(ctx context.Context, request ListSpacesRequest) listing.Iterator[Space] {

	getNextPage := func(ctx context.Context, req ListSpacesRequest) (*ListSpacesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSpaces(ctx, req)
	}
	getItems := func(resp *ListSpacesResponse) []Space {
		return resp.Spaces
	}
	getNextReq := func(resp *ListSpacesResponse) *ListSpacesRequest {
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

// Lists all app spaces in the workspace.
func (a *appsImpl) ListSpacesAll(ctx context.Context, request ListSpacesRequest) ([]Space, error) {
	iterator := a.ListSpaces(ctx, request)
	return listing.ToSlice[Space](ctx, iterator)
}

func (a *appsImpl) internalListSpaces(ctx context.Context, request ListSpacesRequest) (*ListSpacesResponse, error) {
	var listSpacesResponse ListSpacesResponse
	path := "/api/2.0/app-spaces"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSpacesResponse)
	return &listSpacesResponse, err
}

func (a *appsImpl) SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &appPermissions)
	return &appPermissions, err
}

func (a *appsImpl) Start(ctx context.Context, request StartAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v/start", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsImpl) Stop(ctx context.Context, request StopAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v/stop", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsImpl) Update(ctx context.Context, request UpdateAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.App, &app)
	return &app, err
}

func (a *appsImpl) UpdateAppDeploymentScanFinding(ctx context.Context, request UpdateAppDeploymentScanFindingRequest) (*AppDeploymentScanFinding, error) {
	var appDeploymentScanFinding AppDeploymentScanFinding
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.AppDeploymentScanFinding, &appDeploymentScanFinding)
	return &appDeploymentScanFinding, err
}

func (a *appsImpl) UpdateAppThumbnail(ctx context.Context, request UpdateAppThumbnailRequest) (*AppThumbnail, error) {
	var appThumbnail AppThumbnail
	path := fmt.Sprintf("/api/2.0/apps/%v/thumbnail", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &appThumbnail)
	return &appThumbnail, err
}

func (a *appsImpl) UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &appPermissions)
	return &appPermissions, err
}

func (a *appsImpl) UpdateSpace(ctx context.Context, request UpdateSpaceRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/app-spaces/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Space, &operation)
	return &operation, err
}

// unexported type that holds implementations of just AppsSettings API methods
type appsSettingsImpl struct {
	client *client.DatabricksClient
}

func (a *appsSettingsImpl) CreateCustomTemplate(ctx context.Context, request CreateCustomTemplateRequest) (*CustomTemplate, error) {
	var customTemplate CustomTemplate
	path := "/api/2.0/apps-settings/templates"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Template, &customTemplate)
	return &customTemplate, err
}

func (a *appsSettingsImpl) DeleteCustomTemplate(ctx context.Context, request DeleteCustomTemplateRequest) (*CustomTemplate, error) {
	var customTemplate CustomTemplate
	path := fmt.Sprintf("/api/2.0/apps-settings/templates/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &customTemplate)
	return &customTemplate, err
}

func (a *appsSettingsImpl) GetCustomTemplate(ctx context.Context, request GetCustomTemplateRequest) (*CustomTemplate, error) {
	var customTemplate CustomTemplate
	path := fmt.Sprintf("/api/2.0/apps-settings/templates/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &customTemplate)
	return &customTemplate, err
}

// Lists all custom templates in the workspace.
func (a *appsSettingsImpl) ListCustomTemplates(ctx context.Context, request ListCustomTemplatesRequest) listing.Iterator[CustomTemplate] {

	getNextPage := func(ctx context.Context, req ListCustomTemplatesRequest) (*ListCustomTemplatesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCustomTemplates(ctx, req)
	}
	getItems := func(resp *ListCustomTemplatesResponse) []CustomTemplate {
		return resp.Templates
	}
	getNextReq := func(resp *ListCustomTemplatesResponse) *ListCustomTemplatesRequest {
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

// Lists all custom templates in the workspace.
func (a *appsSettingsImpl) ListCustomTemplatesAll(ctx context.Context, request ListCustomTemplatesRequest) ([]CustomTemplate, error) {
	iterator := a.ListCustomTemplates(ctx, request)
	return listing.ToSlice[CustomTemplate](ctx, iterator)
}

func (a *appsSettingsImpl) internalListCustomTemplates(ctx context.Context, request ListCustomTemplatesRequest) (*ListCustomTemplatesResponse, error) {
	var listCustomTemplatesResponse ListCustomTemplatesResponse
	path := "/api/2.0/apps-settings/templates"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCustomTemplatesResponse)
	return &listCustomTemplatesResponse, err
}

func (a *appsSettingsImpl) UpdateCustomTemplate(ctx context.Context, request UpdateCustomTemplateRequest) (*CustomTemplate, error) {
	var customTemplate CustomTemplate
	path := fmt.Sprintf("/api/2.0/apps-settings/templates/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request.Template, &customTemplate)
	return &customTemplate, err
}
