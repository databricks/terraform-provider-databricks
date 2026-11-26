// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package softwareartifacts

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

// unexported type that holds implementations of just SoftwareArtifacts API methods
type softwareArtifactsImpl struct {
	client *client.DatabricksClient
}

func (a *softwareArtifactsImpl) DeleteSoftwareArtifact(ctx context.Context, request DeleteSoftwareArtifactRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
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

func (a *softwareArtifactsImpl) DeleteSoftwareArtifactVersion(ctx context.Context, request DeleteSoftwareArtifactVersionRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
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

func (a *softwareArtifactsImpl) GetSoftwareArtifact(ctx context.Context, request GetSoftwareArtifactRequest) (*SoftwareArtifact, error) {
	var softwareArtifact SoftwareArtifact
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &softwareArtifact)
	return &softwareArtifact, err
}

func (a *softwareArtifactsImpl) GetSoftwareArtifactVersion(ctx context.Context, request GetSoftwareArtifactVersionRequest) (*SoftwareArtifactVersion, error) {
	var softwareArtifactVersion SoftwareArtifactVersion
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &softwareArtifactVersion)
	return &softwareArtifactVersion, err
}

// Lists the versions of a software artifact.
func (a *softwareArtifactsImpl) ListSoftwareArtifactVersions(ctx context.Context, request ListSoftwareArtifactVersionsRequest) listing.Iterator[SoftwareArtifactVersion] {

	getNextPage := func(ctx context.Context, req ListSoftwareArtifactVersionsRequest) (*ListSoftwareArtifactVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSoftwareArtifactVersions(ctx, req)
	}
	getItems := func(resp *ListSoftwareArtifactVersionsResponse) []SoftwareArtifactVersion {
		return resp.SoftwareArtifactVersions
	}
	getNextReq := func(resp *ListSoftwareArtifactVersionsResponse) *ListSoftwareArtifactVersionsRequest {
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

// Lists the versions of a software artifact.
func (a *softwareArtifactsImpl) ListSoftwareArtifactVersionsAll(ctx context.Context, request ListSoftwareArtifactVersionsRequest) ([]SoftwareArtifactVersion, error) {
	iterator := a.ListSoftwareArtifactVersions(ctx, request)
	return listing.ToSlice[SoftwareArtifactVersion](ctx, iterator)
}

func (a *softwareArtifactsImpl) internalListSoftwareArtifactVersions(ctx context.Context, request ListSoftwareArtifactVersionsRequest) (*ListSoftwareArtifactVersionsResponse, error) {
	var listSoftwareArtifactVersionsResponse ListSoftwareArtifactVersionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v/versions", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSoftwareArtifactVersionsResponse)
	return &listSoftwareArtifactVersionsResponse, err
}

// Lists software artifacts in a schema.
func (a *softwareArtifactsImpl) ListSoftwareArtifacts(ctx context.Context, request ListSoftwareArtifactsRequest) listing.Iterator[SoftwareArtifact] {

	getNextPage := func(ctx context.Context, req ListSoftwareArtifactsRequest) (*ListSoftwareArtifactsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSoftwareArtifacts(ctx, req)
	}
	getItems := func(resp *ListSoftwareArtifactsResponse) []SoftwareArtifact {
		return resp.SoftwareArtifacts
	}
	getNextReq := func(resp *ListSoftwareArtifactsResponse) *ListSoftwareArtifactsRequest {
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

// Lists software artifacts in a schema.
func (a *softwareArtifactsImpl) ListSoftwareArtifactsAll(ctx context.Context, request ListSoftwareArtifactsRequest) ([]SoftwareArtifact, error) {
	iterator := a.ListSoftwareArtifacts(ctx, request)
	return listing.ToSlice[SoftwareArtifact](ctx, iterator)
}

func (a *softwareArtifactsImpl) internalListSoftwareArtifacts(ctx context.Context, request ListSoftwareArtifactsRequest) (*ListSoftwareArtifactsResponse, error) {
	var listSoftwareArtifactsResponse ListSoftwareArtifactsResponse
	path := "/api/2.1/unity-catalog/software-artifacts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSoftwareArtifactsResponse)
	return &listSoftwareArtifactsResponse, err
}

func (a *softwareArtifactsImpl) ResolveSoftwareArtifactVersion(ctx context.Context, request ResolveSoftwareArtifactVersionRequest) (*ResolveSoftwareArtifactVersionResponse, error) {
	var resolveSoftwareArtifactVersionResponse ResolveSoftwareArtifactVersionResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v/versions:resolve", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveSoftwareArtifactVersionResponse)
	return &resolveSoftwareArtifactVersionResponse, err
}

func (a *softwareArtifactsImpl) UpdateSoftwareArtifact(ctx context.Context, request UpdateSoftwareArtifactRequest) (*SoftwareArtifact, error) {
	var softwareArtifact SoftwareArtifact
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.SoftwareArtifact, &softwareArtifact)
	return &softwareArtifact, err
}
