// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package geniemigrate

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

// unexported type that holds implementations of just GenieMigrate API methods
type genieMigrateImpl struct {
	client *client.DatabricksClient
}

func (a *genieMigrateImpl) AddFiles(ctx context.Context, request AddProjectFilesRequest) (*AddProjectFilesResponse, error) {
	var addProjectFilesResponse AddProjectFilesResponse
	path := fmt.Sprintf("/api/2.0/migrations/%v/files:add", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &addProjectFilesResponse)
	return &addProjectFilesResponse, err
}

func (a *genieMigrateImpl) CancelOperation(ctx context.Context, request CancelOperationRequest) error {
	path := fmt.Sprintf("/api/2.0/migrations/%v/cancel", request.Name)
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

func (a *genieMigrateImpl) ConvertProjectFiles(ctx context.Context, request ConvertProjectFilesRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/migrations/%v/files:convert", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *genieMigrateImpl) Create(ctx context.Context, request CreateProjectRequest) (*Project, error) {
	var project Project
	path := "/api/2.0/migrations/projects"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Project, &project)
	return &project, err
}

func (a *genieMigrateImpl) Delete(ctx context.Context, request DeleteProjectRequest) error {
	path := fmt.Sprintf("/api/2.0/migrations/%v", request.Name)
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

func (a *genieMigrateImpl) Get(ctx context.Context, request GetProjectRequest) (*Project, error) {
	var project Project
	path := fmt.Sprintf("/api/2.0/migrations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &project)
	return &project, err
}

func (a *genieMigrateImpl) GetFile(ctx context.Context, request GetProjectFileRequest) (*ProjectFile, error) {
	var projectFile ProjectFile
	path := fmt.Sprintf("/api/2.0/migrations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &projectFile)
	return &projectFile, err
}

func (a *genieMigrateImpl) GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/migrations/%v", request.Name)
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

// Lists the migration projects in a workspace folder. Select the folder with
// `filter`; without a filter, lists the caller's home folder. Listing every
// project in the workspace isn't supported.
//
// Returns project summaries without conversion settings; call `GetProject` for
// the full project. A page can be empty even when more pages follow, so keep
// requesting pages until `next_page_token` is empty.
func (a *genieMigrateImpl) List(ctx context.Context, request ListProjectsRequest) listing.Iterator[Project] {

	getNextPage := func(ctx context.Context, req ListProjectsRequest) (*ListProjectsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListProjectsResponse) []Project {
		return resp.Projects
	}
	getNextReq := func(resp *ListProjectsResponse) *ListProjectsRequest {
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

// Lists the migration projects in a workspace folder. Select the folder with
// `filter`; without a filter, lists the caller's home folder. Listing every
// project in the workspace isn't supported.
//
// Returns project summaries without conversion settings; call `GetProject` for
// the full project. A page can be empty even when more pages follow, so keep
// requesting pages until `next_page_token` is empty.
func (a *genieMigrateImpl) ListAll(ctx context.Context, request ListProjectsRequest) ([]Project, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Project](ctx, iterator)
}

func (a *genieMigrateImpl) internalList(ctx context.Context, request ListProjectsRequest) (*ListProjectsResponse, error) {
	var listProjectsResponse ListProjectsResponse
	path := "/api/2.0/migrations/projects"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProjectsResponse)
	return &listProjectsResponse, err
}

// Lists the files of a migration project with their latest conversion statuses.
// To read a file's changes, to-do items, blockers, and notes, call
// `GetProjectFile`. A page can have fewer files than requested, or none, even
// when more pages follow; keep requesting pages until `next_page_token` is
// empty.
func (a *genieMigrateImpl) ListFiles(ctx context.Context, request ListProjectFilesRequest) listing.Iterator[ProjectFile] {

	getNextPage := func(ctx context.Context, req ListProjectFilesRequest) (*ListProjectFilesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListFiles(ctx, req)
	}
	getItems := func(resp *ListProjectFilesResponse) []ProjectFile {
		return resp.ProjectFiles
	}
	getNextReq := func(resp *ListProjectFilesResponse) *ListProjectFilesRequest {
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

// Lists the files of a migration project with their latest conversion statuses.
// To read a file's changes, to-do items, blockers, and notes, call
// `GetProjectFile`. A page can have fewer files than requested, or none, even
// when more pages follow; keep requesting pages until `next_page_token` is
// empty.
func (a *genieMigrateImpl) ListFilesAll(ctx context.Context, request ListProjectFilesRequest) ([]ProjectFile, error) {
	iterator := a.ListFiles(ctx, request)
	return listing.ToSlice[ProjectFile](ctx, iterator)
}

func (a *genieMigrateImpl) internalListFiles(ctx context.Context, request ListProjectFilesRequest) (*ListProjectFilesResponse, error) {
	var listProjectFilesResponse ListProjectFilesResponse
	path := fmt.Sprintf("/api/2.0/migrations/%v/files", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProjectFilesResponse)
	return &listProjectFilesResponse, err
}

func (a *genieMigrateImpl) RemoveFiles(ctx context.Context, request RemoveProjectFilesRequest) (*RemoveProjectFilesResponse, error) {
	var removeProjectFilesResponse RemoveProjectFilesResponse
	path := fmt.Sprintf("/api/2.0/migrations/%v/files:remove", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &removeProjectFilesResponse)
	return &removeProjectFilesResponse, err
}

func (a *genieMigrateImpl) ReviewProjectFiles(ctx context.Context, request ReviewProjectFilesRequest) (*ReviewProjectFilesResponse, error) {
	var reviewProjectFilesResponse ReviewProjectFilesResponse
	path := fmt.Sprintf("/api/2.0/migrations/%v/files:review", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &reviewProjectFilesResponse)
	return &reviewProjectFilesResponse, err
}

func (a *genieMigrateImpl) UnreviewProjectFiles(ctx context.Context, request UnreviewProjectFilesRequest) (*UnreviewProjectFilesResponse, error) {
	var unreviewProjectFilesResponse UnreviewProjectFilesResponse
	path := fmt.Sprintf("/api/2.0/migrations/%v/files:unreview", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &unreviewProjectFilesResponse)
	return &unreviewProjectFilesResponse, err
}

func (a *genieMigrateImpl) Update(ctx context.Context, request UpdateProjectRequest) (*Project, error) {
	var project Project
	path := fmt.Sprintf("/api/2.0/migrations/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Project, &project)
	return &project, err
}
