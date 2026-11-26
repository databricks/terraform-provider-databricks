// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package bundledeployments

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just BundleDeployments API methods
type bundleDeploymentsImpl struct {
	client *client.DatabricksClient
}

func (a *bundleDeploymentsImpl) CompleteVersion(ctx context.Context, request CompleteVersionRequest) (*Version, error) {
	var version Version
	path := fmt.Sprintf("/api/2.0/bundle/%v/complete", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &version)
	return &version, err
}

func (a *bundleDeploymentsImpl) CreateDeployment(ctx context.Context, request CreateDeploymentRequest) (*Deployment, error) {
	var deployment Deployment
	path := "/api/2.0/bundle/deployments"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Deployment, &deployment)
	return &deployment, err
}

func (a *bundleDeploymentsImpl) CreateVersion(ctx context.Context, request CreateVersionRequest) (*Version, error) {
	var version Version
	path := fmt.Sprintf("/api/2.0/bundle/%v/versions", request.Parent)
	queryParams := make(map[string]any)

	if request.VersionId != "" {
		queryParams["version_id"] = request.VersionId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Version, &version)
	return &version, err
}

func (a *bundleDeploymentsImpl) DeleteDeployment(ctx context.Context, request DeleteDeploymentRequest) error {
	path := fmt.Sprintf("/api/2.0/bundle/%v", request.Name)
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

func (a *bundleDeploymentsImpl) GetDeployment(ctx context.Context, request GetDeploymentRequest) (*Deployment, error) {
	var deployment Deployment
	path := fmt.Sprintf("/api/2.0/bundle/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &deployment)
	return &deployment, err
}

func (a *bundleDeploymentsImpl) GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/bundle/%v", request.Name)
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

func (a *bundleDeploymentsImpl) GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error) {
	var resource Resource
	path := fmt.Sprintf("/api/2.0/bundle/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &resource)
	return &resource, err
}

func (a *bundleDeploymentsImpl) GetVersion(ctx context.Context, request GetVersionRequest) (*Version, error) {
	var version Version
	path := fmt.Sprintf("/api/2.0/bundle/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &version)
	return &version, err
}

func (a *bundleDeploymentsImpl) Heartbeat(ctx context.Context, request HeartbeatRequest) (*HeartbeatResponse, error) {
	var heartbeatResponse HeartbeatResponse
	path := fmt.Sprintf("/api/2.0/bundle/%v/heartbeat", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &heartbeatResponse)
	return &heartbeatResponse, err
}

// Lists deployments in the workspace.
func (a *bundleDeploymentsImpl) ListDeployments(ctx context.Context, request ListDeploymentsRequest) listing.Iterator[Deployment] {

	getNextPage := func(ctx context.Context, req ListDeploymentsRequest) (*ListDeploymentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDeployments(ctx, req)
	}
	getItems := func(resp *ListDeploymentsResponse) []Deployment {
		return resp.Deployments
	}
	getNextReq := func(resp *ListDeploymentsResponse) *ListDeploymentsRequest {
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

// Lists deployments in the workspace.
func (a *bundleDeploymentsImpl) ListDeploymentsAll(ctx context.Context, request ListDeploymentsRequest) ([]Deployment, error) {
	iterator := a.ListDeployments(ctx, request)
	return listing.ToSlice[Deployment](ctx, iterator)
}

func (a *bundleDeploymentsImpl) internalListDeployments(ctx context.Context, request ListDeploymentsRequest) (*ListDeploymentsResponse, error) {
	var listDeploymentsResponse ListDeploymentsResponse
	path := "/api/2.0/bundle/deployments"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDeploymentsResponse)
	return &listDeploymentsResponse, err
}

// Lists resource operations under a version.
func (a *bundleDeploymentsImpl) ListOperations(ctx context.Context, request ListOperationsRequest) listing.Iterator[Operation] {

	getNextPage := func(ctx context.Context, req ListOperationsRequest) (*ListOperationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListOperations(ctx, req)
	}
	getItems := func(resp *ListOperationsResponse) []Operation {
		return resp.Operations
	}
	getNextReq := func(resp *ListOperationsResponse) *ListOperationsRequest {
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

// Lists resource operations under a version.
func (a *bundleDeploymentsImpl) ListOperationsAll(ctx context.Context, request ListOperationsRequest) ([]Operation, error) {
	iterator := a.ListOperations(ctx, request)
	return listing.ToSlice[Operation](ctx, iterator)
}

func (a *bundleDeploymentsImpl) internalListOperations(ctx context.Context, request ListOperationsRequest) (*ListOperationsResponse, error) {
	var listOperationsResponse ListOperationsResponse
	path := fmt.Sprintf("/api/2.0/bundle/%v/operations", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listOperationsResponse)
	return &listOperationsResponse, err
}

// Lists resources under a deployment.
func (a *bundleDeploymentsImpl) ListResources(ctx context.Context, request ListResourcesRequest) listing.Iterator[Resource] {

	getNextPage := func(ctx context.Context, req ListResourcesRequest) (*ListResourcesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListResources(ctx, req)
	}
	getItems := func(resp *ListResourcesResponse) []Resource {
		return resp.Resources
	}
	getNextReq := func(resp *ListResourcesResponse) *ListResourcesRequest {
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

// Lists resources under a deployment.
func (a *bundleDeploymentsImpl) ListResourcesAll(ctx context.Context, request ListResourcesRequest) ([]Resource, error) {
	iterator := a.ListResources(ctx, request)
	return listing.ToSlice[Resource](ctx, iterator)
}

func (a *bundleDeploymentsImpl) internalListResources(ctx context.Context, request ListResourcesRequest) (*ListResourcesResponse, error) {
	var listResourcesResponse ListResourcesResponse
	path := fmt.Sprintf("/api/2.0/bundle/%v/resources", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listResourcesResponse)
	return &listResourcesResponse, err
}

// Lists versions under a deployment, ordered numerically by version_id
// descending (most recent first).
func (a *bundleDeploymentsImpl) ListVersions(ctx context.Context, request ListVersionsRequest) listing.Iterator[Version] {

	getNextPage := func(ctx context.Context, req ListVersionsRequest) (*ListVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListVersions(ctx, req)
	}
	getItems := func(resp *ListVersionsResponse) []Version {
		return resp.Versions
	}
	getNextReq := func(resp *ListVersionsResponse) *ListVersionsRequest {
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

// Lists versions under a deployment, ordered numerically by version_id
// descending (most recent first).
func (a *bundleDeploymentsImpl) ListVersionsAll(ctx context.Context, request ListVersionsRequest) ([]Version, error) {
	iterator := a.ListVersions(ctx, request)
	return listing.ToSlice[Version](ctx, iterator)
}

func (a *bundleDeploymentsImpl) internalListVersions(ctx context.Context, request ListVersionsRequest) (*ListVersionsResponse, error) {
	var listVersionsResponse ListVersionsResponse
	path := fmt.Sprintf("/api/2.0/bundle/%v/versions", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listVersionsResponse)
	return &listVersionsResponse, err
}

func (a *bundleDeploymentsImpl) UpdateDeployment(ctx context.Context, request UpdateDeploymentRequest) (*Deployment, error) {
	var deployment Deployment
	path := fmt.Sprintf("/api/2.0/bundle/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Deployment, &deployment)
	return &deployment, err
}

func (a *bundleDeploymentsImpl) UpdateOperation(ctx context.Context, request UpdateOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/bundle/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Operation, &operation)
	return &operation, err
}
