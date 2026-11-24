// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package bundle

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Bundle API methods
type bundleImpl struct {
	client *client.DatabricksClient
}

func (a *bundleImpl) CompleteVersion(ctx context.Context, request CompleteVersionRequest) (*Version, error) {
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

func (a *bundleImpl) CreateDeployment(ctx context.Context, request CreateDeploymentRequest) (*Deployment, error) {
	var deployment Deployment
	path := "/api/2.0/bundle/deployments"
	queryParams := make(map[string]any)

	if request.DeploymentId != "" {
		queryParams["deployment_id"] = request.DeploymentId
	}
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

func (a *bundleImpl) CreateOperation(ctx context.Context, request CreateOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/bundle/%v/operations", request.Parent)
	queryParams := make(map[string]any)

	if request.ResourceKey != "" {
		queryParams["resource_key"] = request.ResourceKey
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Operation, &operation)
	return &operation, err
}

func (a *bundleImpl) CreateVersion(ctx context.Context, request CreateVersionRequest) (*Version, error) {
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

func (a *bundleImpl) DeleteDeployment(ctx context.Context, request DeleteDeploymentRequest) error {
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

func (a *bundleImpl) GetDeployment(ctx context.Context, request GetDeploymentRequest) (*Deployment, error) {
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

func (a *bundleImpl) GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
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

func (a *bundleImpl) GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error) {
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

func (a *bundleImpl) GetVersion(ctx context.Context, request GetVersionRequest) (*Version, error) {
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

func (a *bundleImpl) Heartbeat(ctx context.Context, request HeartbeatRequest) (*HeartbeatResponse, error) {
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
func (a *bundleImpl) ListDeployments(ctx context.Context, request ListDeploymentsRequest) listing.Iterator[Deployment] {

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
func (a *bundleImpl) ListDeploymentsAll(ctx context.Context, request ListDeploymentsRequest) ([]Deployment, error) {
	iterator := a.ListDeployments(ctx, request)
	return listing.ToSlice[Deployment](ctx, iterator)
}

func (a *bundleImpl) internalListDeployments(ctx context.Context, request ListDeploymentsRequest) (*ListDeploymentsResponse, error) {
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
func (a *bundleImpl) ListOperations(ctx context.Context, request ListOperationsRequest) listing.Iterator[Operation] {

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
func (a *bundleImpl) ListOperationsAll(ctx context.Context, request ListOperationsRequest) ([]Operation, error) {
	iterator := a.ListOperations(ctx, request)
	return listing.ToSlice[Operation](ctx, iterator)
}

func (a *bundleImpl) internalListOperations(ctx context.Context, request ListOperationsRequest) (*ListOperationsResponse, error) {
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
func (a *bundleImpl) ListResources(ctx context.Context, request ListResourcesRequest) listing.Iterator[Resource] {

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
func (a *bundleImpl) ListResourcesAll(ctx context.Context, request ListResourcesRequest) ([]Resource, error) {
	iterator := a.ListResources(ctx, request)
	return listing.ToSlice[Resource](ctx, iterator)
}

func (a *bundleImpl) internalListResources(ctx context.Context, request ListResourcesRequest) (*ListResourcesResponse, error) {
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

// Lists versions under a deployment, ordered by version_id descending (most
// recent first).
func (a *bundleImpl) ListVersions(ctx context.Context, request ListVersionsRequest) listing.Iterator[Version] {

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

// Lists versions under a deployment, ordered by version_id descending (most
// recent first).
func (a *bundleImpl) ListVersionsAll(ctx context.Context, request ListVersionsRequest) ([]Version, error) {
	iterator := a.ListVersions(ctx, request)
	return listing.ToSlice[Version](ctx, iterator)
}

func (a *bundleImpl) internalListVersions(ctx context.Context, request ListVersionsRequest) (*ListVersionsResponse, error) {
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
