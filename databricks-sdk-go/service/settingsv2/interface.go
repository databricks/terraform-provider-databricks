// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingsv2

import (
	"context"
)

// APIs to manage account level settings
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountSettingsV2Service interface {

	// Get a setting value at account level. See
	// :method:settingsv2/listaccountsettingsmetadata for list of setting
	// available via public APIs at account level.
	GetPublicAccountSetting(ctx context.Context, request GetPublicAccountSettingRequest) (*Setting, error)

	// Get a user preference for a specific user. User preferences are personal
	// settings that allow individual customization without affecting other
	// users. See :method:settingsv2/listaccountuserpreferencesmetadata for list
	// of user preferences available via public APIs.
	GetPublicAccountUserPreference(ctx context.Context, request GetPublicAccountUserPreferenceRequest) (*UserPreference, error)

	// List valid setting keys and metadata. These settings are available to be
	// referenced via GET :method:settingsv2/getpublicaccountsetting and PATCH
	// :method:settingsv2/patchpublicaccountsetting APIs
	ListAccountSettingsMetadata(ctx context.Context, request ListAccountSettingsMetadataRequest) (*ListAccountSettingsMetadataResponse, error)

	// List valid user preferences and their metadata for a specific user. User
	// preferences are personal settings that allow individual customization
	// without affecting other users. These settings are available to be
	// referenced via GET :method:settingsv2/getpublicaccountuserpreference and
	// PATCH :method:settingsv2/patchpublicaccountuserpreference APIs
	ListAccountUserPreferencesMetadata(ctx context.Context, request ListAccountUserPreferencesMetadataRequest) (*ListAccountUserPreferencesMetadataResponse, error)

	// Patch a setting value at account level. See
	// :method:settingsv2/listaccountsettingsmetadata for list of setting
	// available via public APIs at account level. To determine the correct
	// field to include in a patch request, refer to the type field of the
	// setting returned in the :method:settingsv2/listaccountsettingsmetadata
	// response.
	//
	// Note: Page refresh is required for changes to take effect in UI.
	PatchPublicAccountSetting(ctx context.Context, request PatchPublicAccountSettingRequest) (*Setting, error)

	// Update a user preference for a specific user. User preferences are
	// personal settings that allow individual customization without affecting
	// other users. See :method:settingsv2/listaccountuserpreferencesmetadata
	// for list of user preferences available via public APIs.
	//
	// Note: Page refresh is required for changes to take effect in UI.
	PatchPublicAccountUserPreference(ctx context.Context, request PatchPublicAccountUserPreferenceRequest) (*UserPreference, error)
}

// APIs to manage workspace level settings
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceSettingsV2Service interface {

	// Get a setting value at workspace level. See
	// :method:settingsv2/listworkspacesettingsmetadata for the list of settings
	// available via public APIs.
	GetPublicWorkspaceSetting(ctx context.Context, request GetPublicWorkspaceSettingRequest) (*Setting, error)

	// List valid setting keys and metadata. These settings are available to be
	// referenced via GET :method:settingsv2/getpublicworkspacesetting and PATCH
	// :method:settingsv2/patchpublicworkspacesetting APIs
	ListWorkspaceSettingsMetadata(ctx context.Context, request ListWorkspaceSettingsMetadataRequest) (*ListWorkspaceSettingsMetadataResponse, error)

	// Patch a setting value at workspace level. See
	// :method:settingsv2/listworkspacesettingsmetadata for the list of settings
	// available via public APIs at workspace level. To determine the correct
	// field to include in a patch request, refer to the type field of the
	// setting returned in the :method:settingsv2/listworkspacesettingsmetadata
	// response.
	//
	// Note: Page refresh is required for changes to take effect in UI.
	PatchPublicWorkspaceSetting(ctx context.Context, request PatchPublicWorkspaceSettingRequest) (*Setting, error)
}
