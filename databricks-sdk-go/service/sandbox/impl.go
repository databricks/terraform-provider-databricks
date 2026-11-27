// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sandbox

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

// unexported type that holds implementations of just Sandbox API methods
type sandboxImpl struct {
	client *client.DatabricksClient
}

func (a *sandboxImpl) CreateDirectory(ctx context.Context, request CreateDirectoryRequest) (*CreateDirectoryResponse, error) {
	var createDirectoryResponse CreateDirectoryResponse
	path := fmt.Sprintf("/api/2.0/sandbox-files/%v/directories:make", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createDirectoryResponse)
	return &createDirectoryResponse, err
}

func (a *sandboxImpl) CreateSandbox(ctx context.Context, request CreateSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
	path := "/api/2.0/sandboxes"
	queryParams := make(map[string]any)

	if request.SandboxId != "" {
		queryParams["sandbox_id"] = request.SandboxId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Sandbox, &sandbox)
	return &sandbox, err
}

func (a *sandboxImpl) DeletePath(ctx context.Context, request DeletePathRequest) (*DeletePathResponse, error) {
	var deletePathResponse DeletePathResponse
	path := fmt.Sprintf("/api/2.0/sandbox-files/%v/paths:remove", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deletePathResponse)
	return &deletePathResponse, err
}

func (a *sandboxImpl) DeleteSandbox(ctx context.Context, request DeleteSandboxRequest) error {
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
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

func (a *sandboxImpl) DownloadFile(ctx context.Context, request DownloadFileRequest) (*DownloadFileResponse, error) {
	var downloadFileResponse DownloadFileResponse
	path := fmt.Sprintf("/api/2.0/sandbox-files/%v/content", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/octet-stream"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &downloadFileResponse)
	return &downloadFileResponse, err
}

func (a *sandboxImpl) ExecuteCommandSync(ctx context.Context, request ExecuteCommandSyncRequest) (*ExecuteCommandSyncResponse, error) {
	var executeCommandSyncResponse ExecuteCommandSyncResponse
	path := fmt.Sprintf("/api/2.0/sandbox-exec/%v/exec-sync", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &executeCommandSyncResponse)
	return &executeCommandSyncResponse, err
}

func (a *sandboxImpl) GetMetadata(ctx context.Context, request GetMetadataRequest) (*GetMetadataResponse, error) {
	var getMetadataResponse GetMetadataResponse
	path := fmt.Sprintf("/api/2.0/sandbox-files/%v/paths:describe", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getMetadataResponse)
	return &getMetadataResponse, err
}

func (a *sandboxImpl) GetSandbox(ctx context.Context, request GetSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &sandbox)
	return &sandbox, err
}

func (a *sandboxImpl) ListCommands(ctx context.Context, request ListCommandsRequest) (*ListCommandsResponse, error) {
	var listCommandsResponse ListCommandsResponse
	path := fmt.Sprintf("/api/2.0/sandbox-exec/%v/commands", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCommandsResponse)
	return &listCommandsResponse, err
}

// Lists a directory's immediate children in UTF-8 name order. Later pages may
// reflect concurrent changes.
func (a *sandboxImpl) ListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) listing.Iterator[SandboxFileInfo] {

	getNextPage := func(ctx context.Context, req ListDirectoryContentsRequest) (*ListDirectoryContentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDirectoryContents(ctx, req)
	}
	getItems := func(resp *ListDirectoryContentsResponse) []SandboxFileInfo {
		return resp.Contents
	}
	getNextReq := func(resp *ListDirectoryContentsResponse) *ListDirectoryContentsRequest {
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

// Lists a directory's immediate children in UTF-8 name order. Later pages may
// reflect concurrent changes.
func (a *sandboxImpl) ListDirectoryContentsAll(ctx context.Context, request ListDirectoryContentsRequest) ([]SandboxFileInfo, error) {
	iterator := a.ListDirectoryContents(ctx, request)
	return listing.ToSlice[SandboxFileInfo](ctx, iterator)
}

func (a *sandboxImpl) internalListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) (*ListDirectoryContentsResponse, error) {
	var listDirectoryContentsResponse ListDirectoryContentsResponse
	path := fmt.Sprintf("/api/2.0/sandbox-files/%v/directory-entries", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDirectoryContentsResponse)
	return &listDirectoryContentsResponse, err
}

// Lists the sandboxes in the workspace that are visible to the caller,
// including stopped sandboxes. Sandboxes that are being deleted are not
// returned.
func (a *sandboxImpl) ListSandboxes(ctx context.Context, request ListSandboxesRequest) listing.Iterator[Sandbox] {

	getNextPage := func(ctx context.Context, req ListSandboxesRequest) (*ListSandboxesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSandboxes(ctx, req)
	}
	getItems := func(resp *ListSandboxesResponse) []Sandbox {
		return resp.Sandboxes
	}
	getNextReq := func(resp *ListSandboxesResponse) *ListSandboxesRequest {
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

// Lists the sandboxes in the workspace that are visible to the caller,
// including stopped sandboxes. Sandboxes that are being deleted are not
// returned.
func (a *sandboxImpl) ListSandboxesAll(ctx context.Context, request ListSandboxesRequest) ([]Sandbox, error) {
	iterator := a.ListSandboxes(ctx, request)
	return listing.ToSlice[Sandbox](ctx, iterator)
}

func (a *sandboxImpl) internalListSandboxes(ctx context.Context, request ListSandboxesRequest) (*ListSandboxesResponse, error) {
	var listSandboxesResponse ListSandboxesResponse
	path := "/api/2.0/sandboxes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSandboxesResponse)
	return &listSandboxesResponse, err
}

func (a *sandboxImpl) RenamePath(ctx context.Context, request RenamePathRequest) (*RenamePathResponse, error) {
	var renamePathResponse RenamePathResponse
	path := fmt.Sprintf("/api/2.0/sandbox-files/%v/paths:rename", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &renamePathResponse)
	return &renamePathResponse, err
}

func (a *sandboxImpl) StartSandbox(ctx context.Context, request StartSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
	path := fmt.Sprintf("/api/2.0/%v/start", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &sandbox)
	return &sandbox, err
}

func (a *sandboxImpl) StopSandbox(ctx context.Context, request StopSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
	path := fmt.Sprintf("/api/2.0/%v/stop", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &sandbox)
	return &sandbox, err
}

func (a *sandboxImpl) UpdateSandbox(ctx context.Context, request UpdateSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Sandbox, &sandbox)
	return &sandbox, err
}

func (a *sandboxImpl) UploadFile(ctx context.Context, request UploadFileRequest) error {
	path := fmt.Sprintf("/api/2.0/sandbox-files/%v/content", request.Parent)
	queryParams := make(map[string]any)

	if request.Overwrite != false || slices.Contains(request.ForceSendFields, "Overwrite") {
		queryParams["overwrite"] = request.Overwrite
	}

	if request.Path != "" || slices.Contains(request.ForceSendFields, "Path") {
		queryParams["path"] = request.Path
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request.Contents, nil)
	return err
}
