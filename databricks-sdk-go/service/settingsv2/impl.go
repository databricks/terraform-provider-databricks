// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingsv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccountSettingsV2 API methods
type accountSettingsV2Impl struct {
	client *client.DatabricksClient
}

func (a *accountSettingsV2Impl) GetPublicAccountSetting(ctx context.Context, request GetPublicAccountSettingRequest) (*Setting, error) {
	var setting Setting
	path := fmt.Sprintf("/api/2.1/accounts/%v/settings/%v", a.client.ConfiguredAccountID(), request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &setting)
	return &setting, err
}

func (a *accountSettingsV2Impl) GetPublicAccountUserPreference(ctx context.Context, request GetPublicAccountUserPreferenceRequest) (*UserPreference, error) {
	var userPreference UserPreference
	path := fmt.Sprintf("/api/2.1/accounts/%v/users/%v/settings/%v", a.client.ConfiguredAccountID(), request.UserId, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &userPreference)
	return &userPreference, err
}

// List valid setting keys and metadata. These settings are available to be
// referenced via GET :method:settingsv2/getpublicaccountsetting and PATCH
// :method:settingsv2/patchpublicaccountsetting APIs
func (a *accountSettingsV2Impl) ListAccountSettingsMetadata(ctx context.Context, request ListAccountSettingsMetadataRequest) listing.Iterator[SettingsMetadata] {

	getNextPage := func(ctx context.Context, req ListAccountSettingsMetadataRequest) (*ListAccountSettingsMetadataResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListAccountSettingsMetadata(ctx, req)
	}
	getItems := func(resp *ListAccountSettingsMetadataResponse) []SettingsMetadata {
		return resp.SettingsMetadata
	}
	getNextReq := func(resp *ListAccountSettingsMetadataResponse) *ListAccountSettingsMetadataRequest {
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

// List valid setting keys and metadata. These settings are available to be
// referenced via GET :method:settingsv2/getpublicaccountsetting and PATCH
// :method:settingsv2/patchpublicaccountsetting APIs
func (a *accountSettingsV2Impl) ListAccountSettingsMetadataAll(ctx context.Context, request ListAccountSettingsMetadataRequest) ([]SettingsMetadata, error) {
	iterator := a.ListAccountSettingsMetadata(ctx, request)
	return listing.ToSlice[SettingsMetadata](ctx, iterator)
}

func (a *accountSettingsV2Impl) internalListAccountSettingsMetadata(ctx context.Context, request ListAccountSettingsMetadataRequest) (*ListAccountSettingsMetadataResponse, error) {
	var listAccountSettingsMetadataResponse ListAccountSettingsMetadataResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/settings-metadata", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAccountSettingsMetadataResponse)
	return &listAccountSettingsMetadataResponse, err
}

// List valid user preferences and their metadata for a specific user. User
// preferences are personal settings that allow individual customization without
// affecting other users. These settings are available to be referenced via GET
// :method:settingsv2/getpublicaccountuserpreference and PATCH
// :method:settingsv2/patchpublicaccountuserpreference APIs
func (a *accountSettingsV2Impl) ListAccountUserPreferencesMetadata(ctx context.Context, request ListAccountUserPreferencesMetadataRequest) listing.Iterator[SettingsMetadata] {

	getNextPage := func(ctx context.Context, req ListAccountUserPreferencesMetadataRequest) (*ListAccountUserPreferencesMetadataResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListAccountUserPreferencesMetadata(ctx, req)
	}
	getItems := func(resp *ListAccountUserPreferencesMetadataResponse) []SettingsMetadata {
		return resp.SettingsMetadata
	}
	getNextReq := func(resp *ListAccountUserPreferencesMetadataResponse) *ListAccountUserPreferencesMetadataRequest {
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

// List valid user preferences and their metadata for a specific user. User
// preferences are personal settings that allow individual customization without
// affecting other users. These settings are available to be referenced via GET
// :method:settingsv2/getpublicaccountuserpreference and PATCH
// :method:settingsv2/patchpublicaccountuserpreference APIs
func (a *accountSettingsV2Impl) ListAccountUserPreferencesMetadataAll(ctx context.Context, request ListAccountUserPreferencesMetadataRequest) ([]SettingsMetadata, error) {
	iterator := a.ListAccountUserPreferencesMetadata(ctx, request)
	return listing.ToSlice[SettingsMetadata](ctx, iterator)
}

func (a *accountSettingsV2Impl) internalListAccountUserPreferencesMetadata(ctx context.Context, request ListAccountUserPreferencesMetadataRequest) (*ListAccountUserPreferencesMetadataResponse, error) {
	var listAccountUserPreferencesMetadataResponse ListAccountUserPreferencesMetadataResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/users/%v/settings-metadata", a.client.ConfiguredAccountID(), request.UserId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAccountUserPreferencesMetadataResponse)
	return &listAccountUserPreferencesMetadataResponse, err
}

func (a *accountSettingsV2Impl) PatchPublicAccountSetting(ctx context.Context, request PatchPublicAccountSettingRequest) (*Setting, error) {
	var setting Setting
	path := fmt.Sprintf("/api/2.1/accounts/%v/settings/%v", a.client.ConfiguredAccountID(), request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Setting, &setting)
	return &setting, err
}

func (a *accountSettingsV2Impl) PatchPublicAccountUserPreference(ctx context.Context, request PatchPublicAccountUserPreferenceRequest) (*UserPreference, error) {
	var userPreference UserPreference
	path := fmt.Sprintf("/api/2.1/accounts/%v/users/%v/settings/%v", a.client.ConfiguredAccountID(), request.UserId, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Setting, &userPreference)
	return &userPreference, err
}

// unexported type that holds implementations of just WorkspaceSettingsV2 API methods
type workspaceSettingsV2Impl struct {
	client *client.DatabricksClient
}

func (a *workspaceSettingsV2Impl) GetPublicWorkspaceSetting(ctx context.Context, request GetPublicWorkspaceSettingRequest) (*Setting, error) {
	var setting Setting
	path := fmt.Sprintf("/api/2.1/settings/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &setting)
	return &setting, err
}

// List valid setting keys and metadata. These settings are available to be
// referenced via GET :method:settingsv2/getpublicworkspacesetting and PATCH
// :method:settingsv2/patchpublicworkspacesetting APIs
func (a *workspaceSettingsV2Impl) ListWorkspaceSettingsMetadata(ctx context.Context, request ListWorkspaceSettingsMetadataRequest) listing.Iterator[SettingsMetadata] {

	getNextPage := func(ctx context.Context, req ListWorkspaceSettingsMetadataRequest) (*ListWorkspaceSettingsMetadataResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListWorkspaceSettingsMetadata(ctx, req)
	}
	getItems := func(resp *ListWorkspaceSettingsMetadataResponse) []SettingsMetadata {
		return resp.SettingsMetadata
	}
	getNextReq := func(resp *ListWorkspaceSettingsMetadataResponse) *ListWorkspaceSettingsMetadataRequest {
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

// List valid setting keys and metadata. These settings are available to be
// referenced via GET :method:settingsv2/getpublicworkspacesetting and PATCH
// :method:settingsv2/patchpublicworkspacesetting APIs
func (a *workspaceSettingsV2Impl) ListWorkspaceSettingsMetadataAll(ctx context.Context, request ListWorkspaceSettingsMetadataRequest) ([]SettingsMetadata, error) {
	iterator := a.ListWorkspaceSettingsMetadata(ctx, request)
	return listing.ToSlice[SettingsMetadata](ctx, iterator)
}

func (a *workspaceSettingsV2Impl) internalListWorkspaceSettingsMetadata(ctx context.Context, request ListWorkspaceSettingsMetadataRequest) (*ListWorkspaceSettingsMetadataResponse, error) {
	var listWorkspaceSettingsMetadataResponse ListWorkspaceSettingsMetadataResponse
	path := "/api/2.1/settings-metadata"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceSettingsMetadataResponse)
	return &listWorkspaceSettingsMetadataResponse, err
}

func (a *workspaceSettingsV2Impl) PatchPublicWorkspaceSetting(ctx context.Context, request PatchPublicWorkspaceSettingRequest) (*Setting, error) {
	var setting Setting
	path := fmt.Sprintf("/api/2.1/settings/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Setting, &setting)
	return &setting, err
}
