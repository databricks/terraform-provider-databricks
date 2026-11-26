// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just GitCredentials API methods
type gitCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *gitCredentialsImpl) Create(ctx context.Context, request CreateCredentialsRequest) (*CreateCredentialsResponse, error) {
	var createCredentialsResponse CreateCredentialsResponse
	path := "/api/2.0/git-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createCredentialsResponse)
	return &createCredentialsResponse, err
}

func (a *gitCredentialsImpl) Delete(ctx context.Context, request DeleteCredentialsRequest) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
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

func (a *gitCredentialsImpl) Get(ctx context.Context, request GetCredentialsRequest) (*GetCredentialsResponse, error) {
	var getCredentialsResponse GetCredentialsResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getCredentialsResponse)
	return &getCredentialsResponse, err
}

// Lists the calling user's Git credentials.
func (a *gitCredentialsImpl) List(ctx context.Context, request ListCredentialsRequest) listing.Iterator[CredentialInfo] {

	getNextPage := func(ctx context.Context, req ListCredentialsRequest) (*ListCredentialsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCredentialsResponse) []CredentialInfo {
		return resp.Credentials
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Lists the calling user's Git credentials.
func (a *gitCredentialsImpl) ListAll(ctx context.Context, request ListCredentialsRequest) ([]CredentialInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CredentialInfo](ctx, iterator)
}

func (a *gitCredentialsImpl) internalList(ctx context.Context, request ListCredentialsRequest) (*ListCredentialsResponse, error) {
	var listCredentialsResponse ListCredentialsResponse
	path := "/api/2.0/git-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCredentialsResponse)
	return &listCredentialsResponse, err
}

func (a *gitCredentialsImpl) Update(ctx context.Context, request UpdateCredentialsRequest) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just Repos API methods
type reposImpl struct {
	client *client.DatabricksClient
}

func (a *reposImpl) Create(ctx context.Context, request CreateRepoRequest) (*CreateRepoResponse, error) {
	var createRepoResponse CreateRepoResponse
	path := "/api/2.0/repos"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createRepoResponse)
	return &createRepoResponse, err
}

func (a *reposImpl) Delete(ctx context.Context, request DeleteRepoRequest) error {
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
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

func (a *reposImpl) Get(ctx context.Context, request GetRepoRequest) (*GetRepoResponse, error) {
	var getRepoResponse GetRepoResponse
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRepoResponse)
	return &getRepoResponse, err
}

func (a *reposImpl) GetPermissionLevels(ctx context.Context, request GetRepoPermissionLevelsRequest) (*GetRepoPermissionLevelsResponse, error) {
	var getRepoPermissionLevelsResponse GetRepoPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v/permissionLevels", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRepoPermissionLevelsResponse)
	return &getRepoPermissionLevelsResponse, err
}

func (a *reposImpl) GetPermissions(ctx context.Context, request GetRepoPermissionsRequest) (*RepoPermissions, error) {
	var repoPermissions RepoPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &repoPermissions)
	return &repoPermissions, err
}

// Returns repos that the calling user has Manage permissions on. Use
// `next_page_token` to iterate through additional pages.
//
// Deprecated: This operation does not return a complete list of the repos in
// the workspace, because repos with the Git CLI enabled are not included in its
// results. Instead, find Git folders with `GET /api/2.0/workspace/get-status`
// (for one path) or `GET /api/2.0/workspace/list` (keep objects with
// `object_type == "REPO"` or `directory_info.is_git_folder == true`), then call
// `GET /api/2.0/repos/{object_id}` for Git metadata.
func (a *reposImpl) List(ctx context.Context, request ListReposRequest) listing.Iterator[RepoInfo] {

	getNextPage := func(ctx context.Context, req ListReposRequest) (*ListReposResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListReposResponse) []RepoInfo {
		return resp.Repos
	}
	getNextReq := func(resp *ListReposResponse) *ListReposRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.NextPageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Returns repos that the calling user has Manage permissions on. Use
// `next_page_token` to iterate through additional pages.
//
// Deprecated: This operation does not return a complete list of the repos in
// the workspace, because repos with the Git CLI enabled are not included in its
// results. Instead, find Git folders with `GET /api/2.0/workspace/get-status`
// (for one path) or `GET /api/2.0/workspace/list` (keep objects with
// `object_type == "REPO"` or `directory_info.is_git_folder == true`), then call
// `GET /api/2.0/repos/{object_id}` for Git metadata.
func (a *reposImpl) ListAll(ctx context.Context, request ListReposRequest) ([]RepoInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[RepoInfo](ctx, iterator)
}

func (a *reposImpl) internalList(ctx context.Context, request ListReposRequest) (*ListReposResponse, error) {
	var listReposResponse ListReposResponse
	path := "/api/2.0/repos"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listReposResponse)
	return &listReposResponse, err
}

func (a *reposImpl) SetPermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {
	var repoPermissions RepoPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &repoPermissions)
	return &repoPermissions, err
}

func (a *reposImpl) Update(ctx context.Context, request UpdateRepoRequest) error {
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *reposImpl) UpdatePermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {
	var repoPermissions RepoPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &repoPermissions)
	return &repoPermissions, err
}

// unexported type that holds implementations of just Secrets API methods
type secretsImpl struct {
	client *client.DatabricksClient
}

func (a *secretsImpl) CreateScope(ctx context.Context, request CreateScope) error {
	path := "/api/2.0/secrets/scopes/create"
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

func (a *secretsImpl) DeleteAcl(ctx context.Context, request DeleteAcl) error {
	path := "/api/2.0/secrets/acls/delete"
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

func (a *secretsImpl) DeleteScope(ctx context.Context, request DeleteScope) error {
	path := "/api/2.0/secrets/scopes/delete"
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

func (a *secretsImpl) DeleteSecret(ctx context.Context, request DeleteSecret) error {
	path := "/api/2.0/secrets/delete"
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

func (a *secretsImpl) GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error) {
	var aclItem AclItem
	path := "/api/2.0/secrets/acls/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &aclItem)
	return &aclItem, err
}

func (a *secretsImpl) GetSecret(ctx context.Context, request GetSecretRequest) (*GetSecretResponse, error) {
	var getSecretResponse GetSecretResponse
	path := "/api/2.0/secrets/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getSecretResponse)
	return &getSecretResponse, err
}

// Lists the ACLs set on the given scope.
//
// Users must have the “MANAGE“ permission to invoke this API.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListAcls(ctx context.Context, request ListAclsRequest) listing.Iterator[AclItem] {

	getNextPage := func(ctx context.Context, req ListAclsRequest) (*ListAclsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListAcls(ctx, req)
	}
	getItems := func(resp *ListAclsResponse) []AclItem {
		return resp.Items
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Lists the ACLs set on the given scope.
//
// Users must have the “MANAGE“ permission to invoke this API.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListAclsAll(ctx context.Context, request ListAclsRequest) ([]AclItem, error) {
	iterator := a.ListAcls(ctx, request)
	return listing.ToSlice[AclItem](ctx, iterator)
}

func (a *secretsImpl) internalListAcls(ctx context.Context, request ListAclsRequest) (*ListAclsResponse, error) {
	var listAclsResponse ListAclsResponse
	path := "/api/2.0/secrets/acls/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAclsResponse)
	return &listAclsResponse, err
}

// Lists all secret scopes available in the workspace.
//
// Throws “PERMISSION_DENIED“ if the user does not have permission to make
// this API call.
func (a *secretsImpl) ListScopes(ctx context.Context) listing.Iterator[SecretScope] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListScopesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListScopes(ctx)
	}
	getItems := func(resp *ListScopesResponse) []SecretScope {
		return resp.Scopes
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Lists all secret scopes available in the workspace.
//
// Throws “PERMISSION_DENIED“ if the user does not have permission to make
// this API call.
func (a *secretsImpl) ListScopesAll(ctx context.Context) ([]SecretScope, error) {
	iterator := a.ListScopes(ctx)
	return listing.ToSlice[SecretScope](ctx, iterator)
}

func (a *secretsImpl) internalListScopes(ctx context.Context) (*ListScopesResponse, error) {
	var listScopesResponse ListScopesResponse
	path := "/api/2.0/secrets/scopes/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listScopesResponse)
	return &listScopesResponse, err
}

// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListSecrets(ctx context.Context, request ListSecretsRequest) listing.Iterator[SecretMetadata] {

	getNextPage := func(ctx context.Context, req ListSecretsRequest) (*ListSecretsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSecrets(ctx, req)
	}
	getItems := func(resp *ListSecretsResponse) []SecretMetadata {
		return resp.Secrets
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListSecretsAll(ctx context.Context, request ListSecretsRequest) ([]SecretMetadata, error) {
	iterator := a.ListSecrets(ctx, request)
	return listing.ToSlice[SecretMetadata](ctx, iterator)
}

func (a *secretsImpl) internalListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error) {
	var listSecretsResponse ListSecretsResponse
	path := "/api/2.0/secrets/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSecretsResponse)
	return &listSecretsResponse, err
}

func (a *secretsImpl) PutAcl(ctx context.Context, request PutAcl) error {
	path := "/api/2.0/secrets/acls/put"
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

func (a *secretsImpl) PutSecret(ctx context.Context, request PutSecret) error {
	path := "/api/2.0/secrets/put"
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

// unexported type that holds implementations of just Workspace API methods
type workspaceImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceImpl) Delete(ctx context.Context, request Delete) error {
	path := "/api/2.0/workspace/delete"
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

func (a *workspaceImpl) Export(ctx context.Context, request ExportRequest) (*ExportResponse, error) {
	var exportResponse ExportResponse
	path := "/api/2.0/workspace/export"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &exportResponse)
	return &exportResponse, err
}

func (a *workspaceImpl) GetPermissionLevels(ctx context.Context, request GetWorkspaceObjectPermissionLevelsRequest) (*GetWorkspaceObjectPermissionLevelsResponse, error) {
	var getWorkspaceObjectPermissionLevelsResponse GetWorkspaceObjectPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getWorkspaceObjectPermissionLevelsResponse)
	return &getWorkspaceObjectPermissionLevelsResponse, err
}

func (a *workspaceImpl) GetPermissions(ctx context.Context, request GetWorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	var workspaceObjectPermissions WorkspaceObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceObjectPermissions)
	return &workspaceObjectPermissions, err
}

func (a *workspaceImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error) {
	var objectInfo ObjectInfo
	path := "/api/2.0/workspace/get-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &objectInfo)
	return &objectInfo, err
}

func (a *workspaceImpl) Import(ctx context.Context, request Import) error {
	path := "/api/2.0/workspace/import"
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

// Lists the contents of a directory, or the object if it is not a directory. If
// the input path does not exist, this call returns an error
// `RESOURCE_DOES_NOT_EXIST`.
func (a *workspaceImpl) List(ctx context.Context, request ListWorkspaceRequest) listing.Iterator[ObjectInfo] {

	getNextPage := func(ctx context.Context, req ListWorkspaceRequest) (*ListResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListResponse) []ObjectInfo {
		return resp.Objects
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Lists the contents of a directory, or the object if it is not a directory. If
// the input path does not exist, this call returns an error
// `RESOURCE_DOES_NOT_EXIST`.
func (a *workspaceImpl) ListAll(ctx context.Context, request ListWorkspaceRequest) ([]ObjectInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ObjectInfo](ctx, iterator)
}

func (a *workspaceImpl) internalList(ctx context.Context, request ListWorkspaceRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/workspace/list"
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

func (a *workspaceImpl) Mkdirs(ctx context.Context, request Mkdirs) error {
	path := "/api/2.0/workspace/mkdirs"
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

func (a *workspaceImpl) SetPermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	var workspaceObjectPermissions WorkspaceObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &workspaceObjectPermissions)
	return &workspaceObjectPermissions, err
}

func (a *workspaceImpl) UpdatePermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	var workspaceObjectPermissions WorkspaceObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &workspaceObjectPermissions)
	return &workspaceObjectPermissions, err
}
