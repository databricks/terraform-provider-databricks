// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"context"
)

// The Accounts IP Access List API enables account admins to configure IP access
// lists for access to the account console.
//
// Account IP Access Lists affect web application access and REST API access to
// the account console and account APIs. If the feature is disabled for the
// account, all access is allowed for this account. There is support for allow
// lists (inclusion) and block lists (exclusion).
//
// When a connection is attempted: 1. **First, all block lists are checked.** If
// the connection IP address matches any block list, the connection is rejected.
// 2. **If the connection was not rejected by block lists**, the IP address is
// compared with the allow lists.
//
// If there is at least one allow list for the account, the connection is
// allowed only if the IP address matches an allow list. If there are no allow
// lists for the account, all IP addresses are allowed.
//
// For all allow lists and block lists combined, the account supports a maximum
// of 1000 IP/CIDR values, where one CIDR counts as a single value.
//
// After changes to the account-level IP access lists, it can take a few minutes
// for changes to take effect.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountIpAccessListsService interface {

	// Creates an IP access list for the account.
	//
	// A list can be an allow list or a block list. See the top of this file for
	// a description of how the server treats allow lists and block lists at
	// runtime.
	//
	// When creating or updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the new list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect.
	Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error)

	// Deletes an IP access list, specified by its list ID.
	Delete(ctx context.Context, request DeleteAccountIpAccessListRequest) error

	// Gets an IP access list, specified by its list ID.
	Get(ctx context.Context, request GetAccountIpAccessListRequest) (*GetIpAccessListResponse, error)

	// Gets all IP access lists for the specified account.
	List(ctx context.Context) (*GetIpAccessListsResponse, error)

	// Replaces an IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time. When replacing an IP access list: * For all allow lists and
	// block lists combined, the API supports a maximum of 1000 IP/CIDR values,
	// where one CIDR counts as a single value. Attempts to exceed that number
	// return error 400 with `error_code` value `QUOTA_EXCEEDED`. * If the
	// resulting list would block the calling user's current IP, error 400 is
	// returned with `error_code` value `INVALID_STATE`. It can take a few
	// minutes for the changes to take effect.
	Replace(ctx context.Context, request ReplaceIpAccessList) error

	// Updates an existing IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time.
	//
	// When updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the updated list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect.
	Update(ctx context.Context, request UpdateIpAccessList) error
}

// Accounts Settings API allows users to manage settings at the account level.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountSettingsService interface {
}

// Controls whether AI/BI published dashboard embedding is enabled,
// conditionally enabled, or disabled at the workspace level. By default, this
// setting is conditionally enabled (ALLOW_APPROVED_DOMAINS).
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AibiDashboardEmbeddingAccessPolicyService interface {

	// Delete the AI/BI dashboard embedding access policy, reverting back to the
	// default.
	Delete(ctx context.Context, request DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error)

	// Retrieves the AI/BI dashboard embedding access policy. The default
	// setting is ALLOW_APPROVED_DOMAINS, permitting AI/BI dashboards to be
	// embedded on approved domains.
	Get(ctx context.Context, request GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error)

	// Updates the AI/BI dashboard embedding access policy at the workspace
	// level.
	Update(ctx context.Context, request UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error)
}

// Controls the list of domains approved to host the embedded AI/BI dashboards.
// The approved domains list can't be mutated when the current access policy is
// not set to ALLOW_APPROVED_DOMAINS.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AibiDashboardEmbeddingApprovedDomainsService interface {

	// Delete the list of domains approved to host embedded AI/BI dashboards,
	// reverting back to the default empty list.
	Delete(ctx context.Context, request DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error)

	// Retrieves the list of domains approved to host embedded AI/BI dashboards.
	Get(ctx context.Context, request GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error)

	// Updates the list of domains approved to host embedded AI/BI dashboards.
	// This update will fail if the current workspace access policy is not
	// ALLOW_APPROVED_DOMAINS.
	Update(ctx context.Context, request UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error)
}

// Controls whether automatic cluster update is enabled for the current
// workspace. By default, it is turned off.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AutomaticClusterUpdateService interface {

	// Gets the automatic cluster update setting.
	Get(ctx context.Context, request GetAutomaticClusterUpdateSettingRequest) (*AutomaticClusterUpdateSetting, error)

	// Updates the automatic cluster update setting for the workspace. A fresh
	// etag needs to be provided in `PATCH` requests (as part of the setting
	// field). The etag can be retrieved by making a `GET` request before the
	// `PATCH` request. If the setting is updated concurrently, `PATCH` fails
	// with 409 and the request must be retried by using the fresh etag in the
	// 409 response.
	Update(ctx context.Context, request UpdateAutomaticClusterUpdateSettingRequest) (*AutomaticClusterUpdateSetting, error)
}

// Controls whether to enable the compliance security profile for the current
// workspace. Enabling it on a workspace is permanent. By default, it is turned
// off.
//
// This settings can NOT be disabled once it is enabled.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ComplianceSecurityProfileService interface {

	// Gets the compliance security profile setting.
	Get(ctx context.Context, request GetComplianceSecurityProfileSettingRequest) (*ComplianceSecurityProfileSetting, error)

	// Updates the compliance security profile setting for the workspace. A
	// fresh etag needs to be provided in `PATCH` requests (as part of the
	// setting field). The etag can be retrieved by making a `GET` request
	// before the `PATCH` request. If the setting is updated concurrently,
	// `PATCH` fails with 409 and the request must be retried by using the fresh
	// etag in the 409 response.
	Update(ctx context.Context, request UpdateComplianceSecurityProfileSettingRequest) (*ComplianceSecurityProfileSetting, error)
}

// Credentials manager interacts with with Identity Providers to to perform
// token exchanges using stored credentials and refresh tokens.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CredentialsManagerService interface {

	// Exchange tokens with an Identity Provider to get a new access token. It
	// allows specifying scopes to determine token permissions.
	//
	// POST /exchange-tokens/token is the documented public form, expressed via
	// `google.api.http` below. GET /exchange-tokens/$exchange is a legacy alias
	// used by the Spark driver's OAuth refresh path (DBHttpClient#get sends a
	// body via HttpGetWithEntity) and stays on the legacy `option
	// (rpc).endpoints` annotation: its path contains a literal `$`, which
	// `google.api.http`'s LITERAL grammar does not allow, and `HttpPathParser`
	// does not percent-decode template segments (so encoding as `%24exchange`
	// would not match the literal `$exchange` path the Spark driver sends).
	// Per-endpoint `visibility: PUBLIC_UNDOCUMENTED` preserves the DECO-7732
	// intent of suppressing the GET alias from the public API spec.
	ExchangeToken(ctx context.Context, request ExchangeTokenRequest) (*ExchangeTokenResponse, error)
}

// The compliance security profile settings at the account level control whether
// to enable it for new workspaces. By default, this account-level setting is
// disabled for new workspaces. After workspace creation, account admins can
// enable the compliance security profile individually for each workspace.
//
// This settings can be disabled so that new workspaces do not have compliance
// security profile enabled by default.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CspEnablementAccountService interface {

	// Gets the compliance security profile setting for new workspaces.
	Get(ctx context.Context, request GetCspEnablementAccountSettingRequest) (*CspEnablementAccountSetting, error)

	// Updates the value of the compliance security profile setting for new
	// workspaces.
	Update(ctx context.Context, request UpdateCspEnablementAccountSettingRequest) (*CspEnablementAccountSetting, error)
}

// Controls whether schedules or workload tasks for refreshing AI/BI Dashboards
// in the workspace can send subscription emails containing PDFs and/or images
// of the dashboard. By default, this setting is enabled (set to `true`)
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DashboardEmailSubscriptionsService interface {

	// Reverts the Dashboard Email Subscriptions setting to its default value.
	Delete(ctx context.Context, request DeleteDashboardEmailSubscriptionsRequest) (*DeleteDashboardEmailSubscriptionsResponse, error)

	// Gets the Dashboard Email Subscriptions setting.
	Get(ctx context.Context, request GetDashboardEmailSubscriptionsRequest) (*DashboardEmailSubscriptions, error)

	// Updates the Dashboard Email Subscriptions setting.
	Update(ctx context.Context, request UpdateDashboardEmailSubscriptionsRequest) (*DashboardEmailSubscriptions, error)
}

// The default namespace setting API allows users to configure the default
// namespace for a Databricks workspace.
//
// Through this API, users can retrieve, set, or modify the default namespace
// used when queries do not reference a fully qualified three-level name. For
// example, if you use the API to set 'retail_prod' as the default catalog, then
// a query 'SELECT * FROM myTable' would reference the object
// 'retail_prod.default.myTable' (the schema 'default' is always assumed).
//
// This setting requires a restart of clusters and SQL warehouses to take
// effect. Additionally, the default namespace only applies when using Unity
// Catalog-enabled compute.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DefaultNamespaceService interface {

	// Deletes the default namespace setting for the workspace. A fresh etag
	// needs to be provided in `DELETE` requests (as a query parameter). The
	// etag can be retrieved by making a `GET` request before the `DELETE`
	// request. If the setting is updated/deleted concurrently, `DELETE` fails
	// with 409 and the request must be retried by using the fresh etag in the
	// 409 response.
	Delete(ctx context.Context, request DeleteDefaultNamespaceSettingRequest) (*DeleteDefaultNamespaceSettingResponse, error)

	// Gets the default namespace setting.
	Get(ctx context.Context, request GetDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error)

	// Updates the default namespace setting for the workspace. A fresh etag
	// needs to be provided in `PATCH` requests (as part of the setting field).
	// The etag can be retrieved by making a `GET` request before the `PATCH`
	// request. Note that if the setting does not exist, `GET` returns a
	// NOT_FOUND error and the etag is present in the error response, which
	// should be set in the `PATCH` request. If the setting is updated
	// concurrently, `PATCH` fails with 409 and the request must be retried by
	// using the fresh etag in the 409 response.
	Update(ctx context.Context, request UpdateDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error)
}

// Warehouse to be selected by default for users in this workspace. Covers SQL
// workloads only and can be overridden by users.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DefaultWarehouseIdService interface {

	// Reverts the Default Warehouse Id setting to its default value.
	Delete(ctx context.Context, request DeleteDefaultWarehouseIdRequest) (*DeleteDefaultWarehouseIdResponse, error)

	// Gets the Default Warehouse Id setting.
	Get(ctx context.Context, request GetDefaultWarehouseIdRequest) (*DefaultWarehouseId, error)

	// Updates the Default Warehouse Id setting.
	Update(ctx context.Context, request UpdateDefaultWarehouseIdRequest) (*DefaultWarehouseId, error)
}

// 'Disabling legacy access' has the following impacts:
//
// 1. Disables direct access to Hive Metastores from the workspace. However, you
// can still access a Hive Metastore through Hive Metastore federation. 2.
// Disables fallback mode on external location access from the workspace. 3.
// Disables Databricks Runtime versions prior to 13.3LTS.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DisableLegacyAccessService interface {

	// Deletes legacy access disablement status.
	Delete(ctx context.Context, request DeleteDisableLegacyAccessRequest) (*DeleteDisableLegacyAccessResponse, error)

	// Retrieves legacy access disablement Status.
	Get(ctx context.Context, request GetDisableLegacyAccessRequest) (*DisableLegacyAccess, error)

	// Updates legacy access disablement status.
	Update(ctx context.Context, request UpdateDisableLegacyAccessRequest) (*DisableLegacyAccess, error)
}

// Disabling legacy DBFS has the following implications:
//
// 1. Access to DBFS root and DBFS mounts is disallowed (as well as the creation
// of new mounts). 2. Disables Databricks Runtime versions prior to 13.3LTS.
//
// When the setting is off, all DBFS functionality is enabled and no
// restrictions are imposed on Databricks Runtime versions. This setting can
// take up to 20 minutes to take effect and requires a manual restart of
// all-purpose compute clusters and SQL warehouses.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DisableLegacyDbfsService interface {

	// Deletes the disable legacy DBFS setting for a workspace, reverting back
	// to the default.
	Delete(ctx context.Context, request DeleteDisableLegacyDbfsRequest) (*DeleteDisableLegacyDbfsResponse, error)

	// Gets the disable legacy DBFS setting.
	Get(ctx context.Context, request GetDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error)

	// Updates the disable legacy DBFS setting for the workspace.
	Update(ctx context.Context, request UpdateDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error)
}

// Disable legacy features for new Databricks workspaces.
//
// For newly created workspaces: 1. Disables the use of DBFS root and mounts. 2.
// Hive Metastore will not be provisioned. 3. Disables the use of
// ‘No-isolation clusters’. 4. Disables Databricks Runtime versions prior to
// 13.3LTS.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DisableLegacyFeaturesService interface {

	// Deletes the disable legacy features setting.
	Delete(ctx context.Context, request DeleteDisableLegacyFeaturesRequest) (*DeleteDisableLegacyFeaturesResponse, error)

	// Gets the value of the disable legacy features setting.
	Get(ctx context.Context, request GetDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error)

	// Updates the value of the disable legacy features setting.
	Update(ctx context.Context, request UpdateDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error)
}

// Controls whether users can export notebooks and files from the Workspace UI.
// By default, this setting is enabled.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type EnableExportNotebookService interface {

	// Gets the Notebook and File exporting setting.
	GetEnableExportNotebook(ctx context.Context) (*EnableExportNotebook, error)

	// Updates the Notebook and File exporting setting. The model follows
	// eventual consistency, which means the get after the update operation
	// might receive stale values for some time.
	PatchEnableExportNotebook(ctx context.Context, request UpdateEnableExportNotebookRequest) (*EnableExportNotebook, error)
}

// Controls the enforcement of IP access lists for accessing the account
// console. Allowing you to enable or disable restricted access based on IP
// addresses.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type EnableIpAccessListsService interface {

	// Reverts the value of the account IP access toggle setting to default (ON)
	Delete(ctx context.Context, request DeleteAccountIpAccessEnableRequest) (*DeleteAccountIpAccessEnableResponse, error)

	// Gets the value of the account IP access toggle setting.
	Get(ctx context.Context, request GetAccountIpAccessEnableRequest) (*AccountIpAccessEnable, error)

	// Updates the value of the account IP access toggle setting.
	Update(ctx context.Context, request UpdateAccountIpAccessEnableRequest) (*AccountIpAccessEnable, error)
}

// Controls whether users can copy tabular data to the clipboard via the UI. By
// default, this setting is enabled.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type EnableNotebookTableClipboardService interface {

	// Gets the Results Table Clipboard features setting.
	GetEnableNotebookTableClipboard(ctx context.Context) (*EnableNotebookTableClipboard, error)

	// Updates the Results Table Clipboard features setting. The model follows
	// eventual consistency, which means the get after the update operation
	// might receive stale values for some time.
	PatchEnableNotebookTableClipboard(ctx context.Context, request UpdateEnableNotebookTableClipboardRequest) (*EnableNotebookTableClipboard, error)
}

// Controls whether users can download notebook results. By default, this
// setting is enabled.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type EnableResultsDownloadingService interface {

	// Gets the Notebook results download setting.
	GetEnableResultsDownloading(ctx context.Context) (*EnableResultsDownloading, error)

	// Updates the Notebook results download setting. The model follows eventual
	// consistency, which means the get after the update operation might receive
	// stale values for some time.
	PatchEnableResultsDownloading(ctx context.Context, request UpdateEnableResultsDownloadingRequest) (*EnableResultsDownloading, error)
}

// Controls whether enhanced security monitoring is enabled for the current
// workspace. If the compliance security profile is enabled, this is
// automatically enabled. By default, it is disabled. However, if the compliance
// security profile is enabled, this is automatically enabled.
//
// If the compliance security profile is disabled, you can enable or disable
// this setting and it is not permanent.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type EnhancedSecurityMonitoringService interface {

	// Gets the enhanced security monitoring setting.
	Get(ctx context.Context, request GetEnhancedSecurityMonitoringSettingRequest) (*EnhancedSecurityMonitoringSetting, error)

	// Updates the enhanced security monitoring setting for the workspace. A
	// fresh etag needs to be provided in `PATCH` requests (as part of the
	// setting field). The etag can be retrieved by making a `GET` request
	// before the `PATCH` request. If the setting is updated concurrently,
	// `PATCH` fails with 409 and the request must be retried by using the fresh
	// etag in the 409 response.
	Update(ctx context.Context, request UpdateEnhancedSecurityMonitoringSettingRequest) (*EnhancedSecurityMonitoringSetting, error)
}

// The enhanced security monitoring setting at the account level controls
// whether to enable the feature on new workspaces. By default, this
// account-level setting is disabled for new workspaces. After workspace
// creation, account admins can enable enhanced security monitoring individually
// for each workspace.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type EsmEnablementAccountService interface {

	// Gets the enhanced security monitoring setting for new workspaces.
	Get(ctx context.Context, request GetEsmEnablementAccountSettingRequest) (*EsmEnablementAccountSetting, error)

	// Updates the value of the enhanced security monitoring setting for new
	// workspaces.
	Update(ctx context.Context, request UpdateEsmEnablementAccountSettingRequest) (*EsmEnablementAccountSetting, error)
}

// IP Access List enables admins to configure IP access lists.
//
// IP access lists affect web application access and REST API access to this
// workspace only. If the feature is disabled for a workspace, all access is
// allowed for this workspace. There is support for allow lists (inclusion) and
// block lists (exclusion).
//
// When a connection is attempted: 1. **First, all block lists are checked.** If
// the connection IP address matches any block list, the connection is rejected.
// 2. **If the connection was not rejected by block lists**, the IP address is
// compared with the allow lists.
//
// If there is at least one allow list for the workspace, the connection is
// allowed only if the IP address matches an allow list. If there are no allow
// lists for the workspace, all IP addresses are allowed.
//
// For all allow lists and block lists combined, the workspace supports a
// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
//
// After changes to the IP access list feature, it can take a few minutes for
// changes to take effect.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type IpAccessListsService interface {

	// Creates an IP access list for this workspace.
	//
	// A list can be an allow list or a block list. See the top of this file for
	// a description of how the server treats allow lists and block lists at
	// runtime.
	//
	// When creating or updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the new list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect. **Note**: Your
	// new IP access list has no effect until you enable the feature. See
	// [workspaceconf/setStatus]
	//
	// [workspaceconf/setStatus]: https://docs.databricks.com/api/workspace/workspaceconf/setstatus
	Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error)

	// Deletes an IP access list, specified by its list ID.
	Delete(ctx context.Context, request DeleteIpAccessListRequest) error

	// Gets an IP access list, specified by its list ID.
	Get(ctx context.Context, request GetIpAccessListRequest) (*FetchIpAccessListResponse, error)

	// Gets all IP access lists for the specified workspace.
	List(ctx context.Context) (*ListIpAccessListResponse, error)

	// Replaces an IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time. When replacing an IP access list: * For all allow lists and
	// block lists combined, the API supports a maximum of 1000 IP/CIDR values,
	// where one CIDR counts as a single value. Attempts to exceed that number
	// return error 400 with `error_code` value `QUOTA_EXCEEDED`. * If the
	// resulting list would block the calling user's current IP, error 400 is
	// returned with `error_code` value `INVALID_STATE`. It can take a few
	// minutes for the changes to take effect. Note that your resulting IP
	// access list has no effect until you enable the feature. See
	// [workspaceconf/setStatus].
	//
	// [workspaceconf/setStatus]: https://docs.databricks.com/api/workspace/workspaceconf/setstatus
	Replace(ctx context.Context, request ReplaceIpAccessList) error

	// Updates an existing IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time.
	//
	// When updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the updated list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect. Note that your
	// resulting IP access list has no effect until you enable the feature. See
	// [workspaceconf/setStatus].
	//
	// [workspaceconf/setStatus]: https://docs.databricks.com/api/workspace/workspaceconf/setstatus
	Update(ctx context.Context, request UpdateIpAccessList) error
}

// Determines if partner powered models are enabled or not for a specific
// account
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LlmProxyPartnerPoweredAccountService interface {

	// Gets the enable partner powered AI features account setting.
	Get(ctx context.Context, request GetLlmProxyPartnerPoweredAccountRequest) (*LlmProxyPartnerPoweredAccount, error)

	// Updates the enable partner powered AI features account setting.
	Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredAccountRequest) (*LlmProxyPartnerPoweredAccount, error)
}

// Determines if the account-level partner-powered setting value is enforced
// upon the workspace-level partner-powered setting
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LlmProxyPartnerPoweredEnforceService interface {

	// Gets the enforcement status of partner powered AI features account
	// setting.
	Get(ctx context.Context, request GetLlmProxyPartnerPoweredEnforceRequest) (*LlmProxyPartnerPoweredEnforce, error)

	// Updates the enable enforcement status of partner powered AI features
	// account setting.
	Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredEnforceRequest) (*LlmProxyPartnerPoweredEnforce, error)
}

// Determines if partner powered models are enabled or not for a specific
// workspace
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LlmProxyPartnerPoweredWorkspaceService interface {

	// Reverts the enable partner powered AI features workspace setting to its
	// default value.
	Delete(ctx context.Context, request DeleteLlmProxyPartnerPoweredWorkspaceRequest) (*DeleteLlmProxyPartnerPoweredWorkspaceResponse, error)

	// Gets the enable partner powered AI features workspace setting.
	Get(ctx context.Context, request GetLlmProxyPartnerPoweredWorkspaceRequest) (*LlmProxyPartnerPoweredWorkspace, error)

	// Updates the enable partner powered AI features workspace setting.
	Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredWorkspaceRequest) (*LlmProxyPartnerPoweredWorkspace, error)
}

// These APIs provide configurations for the network connectivity of your
// workspaces for serverless compute resources. This API provides stable subnets
// for your workspace so that you can configure your firewalls on your Azure
// Storage accounts to allow access from Databricks. You can also use the API to
// provision private endpoints for Databricks to privately connect serverless
// compute resources to your Azure resources using Azure Private Link. See
// [configure serverless secure connectivity].
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [configure serverless secure connectivity]: https://learn.microsoft.com/azure/databricks/security/network/serverless-network-security
type NetworkConnectivityService interface {

	// Creates a network connectivity configuration (NCC), which provides stable
	// Azure service subnets when accessing your Azure Storage accounts. You can
	// also use a network connectivity configuration to create Databricks
	// managed private endpoints so that Databricks serverless compute resources
	// privately access your resources.
	//
	// **IMPORTANT**: After you create the network connectivity configuration,
	// you must assign one or more workspaces to the new network connectivity
	// configuration. You can share one network connectivity configuration with
	// multiple workspaces from the same Azure region within the same Databricks
	// account. See [configure serverless secure connectivity].
	//
	// [configure serverless secure connectivity]: https://learn.microsoft.com/azure/databricks/security/network/serverless-network-security
	CreateNetworkConnectivityConfiguration(ctx context.Context, request CreateNetworkConnectivityConfigRequest) (*NetworkConnectivityConfiguration, error)

	// Create a private endpoint rule for the specified network connectivity
	// config object. Once the object is created, Databricks asynchronously
	// provisions a new Azure private endpoint to your specified Azure resource.
	//
	// **IMPORTANT**: You must use Azure portal or other Azure tools to approve
	// the private endpoint to complete the connection. To get the information
	// of the private endpoint created, make a `GET` request on the new private
	// endpoint rule. See [serverless private link].
	//
	// [serverless private link]: https://learn.microsoft.com/azure/databricks/security/network/serverless-network-security/serverless-private-link
	CreatePrivateEndpointRule(ctx context.Context, request CreatePrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error)

	// Deletes a network connectivity configuration.
	DeleteNetworkConnectivityConfiguration(ctx context.Context, request DeleteNetworkConnectivityConfigurationRequest) error

	// Initiates deleting a private endpoint rule. If the connection state is
	// PENDING or EXPIRED, the private endpoint is immediately deleted.
	// Otherwise, the private endpoint is deactivated and will be deleted after
	// one day of deactivation. When a private endpoint is deactivated, the
	// `deactivated` field is set to `true` and the private endpoint is not
	// available to your serverless compute resources.
	DeletePrivateEndpointRule(ctx context.Context, request DeletePrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error)

	// Gets a network connectivity configuration.
	GetNetworkConnectivityConfiguration(ctx context.Context, request GetNetworkConnectivityConfigurationRequest) (*NetworkConnectivityConfiguration, error)

	// Gets the private endpoint rule.
	GetPrivateEndpointRule(ctx context.Context, request GetPrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error)

	// Gets an array of network connectivity configurations.
	ListNetworkConnectivityConfigurations(ctx context.Context, request ListNetworkConnectivityConfigurationsRequest) (*ListNetworkConnectivityConfigurationsResponse, error)

	// Gets an array of private endpoint rules.
	ListPrivateEndpointRules(ctx context.Context, request ListPrivateEndpointRulesRequest) (*ListPrivateEndpointRulesResponse, error)

	// Updates the name or billing workspace for a network connectivity
	// configuration.
	UpdateNetworkConnectivityConfig(ctx context.Context, request UpdateNetworkConnectivityConfigRequest) (*NetworkConnectivityConfiguration, error)

	// Updates a private endpoint rule. Currently only a private endpoint rule
	// to customer-managed resources is allowed to be updated.
	UpdatePrivateEndpointRule(ctx context.Context, request UpdateNccPrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error)
}

// These APIs manage network policies for this account. Network policies control
// which network destinations can be accessed from the Databricks environment.
// Each Databricks account includes a default policy named 'default-policy'.
// 'default-policy' is associated with any workspace lacking an explicit network
// policy assignment, and is automatically associated with each newly created
// workspace. 'default-policy' is reserved and cannot be deleted, but it can be
// updated to customize the default network access rules for your account.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type NetworkPoliciesService interface {

	// Creates a new network policy to manage which network destinations can be
	// accessed from the Databricks environment.
	CreateNetworkPolicyRpc(ctx context.Context, request CreateNetworkPolicyRequest) (*AccountNetworkPolicy, error)

	// Deletes a network policy. Cannot be called on 'default-policy'.
	DeleteNetworkPolicyRpc(ctx context.Context, request DeleteNetworkPolicyRequest) error

	// Gets a network policy.
	GetNetworkPolicyRpc(ctx context.Context, request GetNetworkPolicyRequest) (*AccountNetworkPolicy, error)

	// Gets an array of network policies.
	ListNetworkPoliciesRpc(ctx context.Context, request ListNetworkPoliciesRequest) (*ListNetworkPoliciesResponse, error)

	// Updates a network policy. This allows you to modify the configuration of
	// a network policy.
	UpdateNetworkPolicyRpc(ctx context.Context, request UpdateNetworkPolicyRequest) (*AccountNetworkPolicy, error)
}

// The notification destinations API lets you programmatically manage a
// workspace's notification destinations. Notification destinations are used to
// send notifications for query alerts and jobs to destinations outside of
// Databricks. Only workspace admins can create, update, and delete notification
// destinations.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type NotificationDestinationsService interface {

	// Creates a notification destination. Requires workspace admin permissions.
	Create(ctx context.Context, request CreateNotificationDestinationRequest) (*NotificationDestination, error)

	// Deletes a notification destination. Requires workspace admin permissions.
	Delete(ctx context.Context, request DeleteNotificationDestinationRequest) error

	// Gets a notification destination by ID. Workspace admins also receive its
	// config.
	Get(ctx context.Context, request GetNotificationDestinationRequest) (*NotificationDestination, error)

	// Lists the notification destinations in the workspace. Results never
	// include a destination's config; workspace admins can retrieve it with
	// Get.
	List(ctx context.Context, request ListNotificationDestinationsRequest) (*ListNotificationDestinationsResponse, error)

	// Updates a notification destination. Requires workspace admin permissions.
	// At least one field is required in the request body. A config provided in
	// the request replaces the existing one.
	Update(ctx context.Context, request UpdateNotificationDestinationRequest) (*NotificationDestination, error)
}

// The Personal Compute enablement setting lets you control which users can use
// the Personal Compute default policy to create compute resources. By default
// all users in all workspaces have access (ON), but you can change the setting
// to instead let individual workspaces configure access control (DELEGATE).
//
// There is only one instance of this setting per account. Since this setting
// has a default value, this setting is present on all accounts even though it's
// never set on a given account. Deletion reverts the value of the setting back
// to the default value.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PersonalComputeService interface {

	// Reverts back the Personal Compute setting value to default (ON)
	Delete(ctx context.Context, request DeletePersonalComputeSettingRequest) (*DeletePersonalComputeSettingResponse, error)

	// Gets the value of the Personal Compute setting.
	Get(ctx context.Context, request GetPersonalComputeSettingRequest) (*PersonalComputeSetting, error)

	// Updates the value of the Personal Compute setting.
	Update(ctx context.Context, request UpdatePersonalComputeSettingRequest) (*PersonalComputeSetting, error)
}

// The Restrict Workspace Admins setting lets you control the capabilities of
// workspace admins. With the setting status set to ALLOW_ALL, workspace admins
// can create service principal personal access tokens on behalf of any service
// principal in their workspace. Workspace admins can also change a job owner to
// any user in their workspace. And they can change the job run_as setting to
// any user in their workspace or to a service principal on which they have the
// Service Principal User role. With the setting status set to
// RESTRICT_TOKENS_AND_JOB_RUN_AS, workspace admins can only create personal
// access tokens on behalf of service principals they have the Service Principal
// User role on. They can also only change a job owner to themselves. And they
// can change the job run_as setting to themselves or to a service principal on
// which they have the Service Principal User role.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type RestrictWorkspaceAdminsService interface {

	// Reverts the restrict workspace admins setting status for the workspace. A
	// fresh etag needs to be provided in `DELETE` requests (as a query
	// parameter). The etag can be retrieved by making a `GET` request before
	// the DELETE request. If the setting is updated/deleted concurrently,
	// `DELETE` fails with 409 and the request must be retried by using the
	// fresh etag in the 409 response.
	Delete(ctx context.Context, request DeleteRestrictWorkspaceAdminsSettingRequest) (*DeleteRestrictWorkspaceAdminsSettingResponse, error)

	// Gets the restrict workspace admins setting.
	Get(ctx context.Context, request GetRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error)

	// Updates the restrict workspace admins setting for the workspace. A fresh
	// etag needs to be provided in `PATCH` requests (as part of the setting
	// field). The etag can be retrieved by making a GET request before the
	// `PATCH` request. If the setting is updated concurrently, `PATCH` fails
	// with 409 and the request must be retried by using the fresh etag in the
	// 409 response.
	Update(ctx context.Context, request UpdateRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error)
}

// Workspace Settings API allows users to manage settings at the workspace
// level.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type SettingsService interface {
}

// Controls whether users within the workspace are allowed to download results
// from the SQL Editor and AI/BI Dashboards UIs. By default, this setting is
// enabled (set to `true`)
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type SqlResultsDownloadService interface {

	// Reverts the SQL Results Download setting to its default value.
	Delete(ctx context.Context, request DeleteSqlResultsDownloadRequest) (*DeleteSqlResultsDownloadResponse, error)

	// Gets the SQL Results Download setting.
	Get(ctx context.Context, request GetSqlResultsDownloadRequest) (*SqlResultsDownload, error)

	// Updates the SQL Results Download setting.
	Update(ctx context.Context, request UpdateSqlResultsDownloadRequest) (*SqlResultsDownload, error)
}

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type TokenManagementService interface {

	// Creates a token on behalf of a service principal.
	CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error)

	// Deletes a token, specified by its ID.
	Delete(ctx context.Context, request DeleteTokenManagementRequest) error

	// Gets information about a token, specified by its ID.
	Get(ctx context.Context, request GetTokenManagementRequest) (*GetTokenResponse, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context) (*GetTokenPermissionLevelsResponse, error)

	// Gets the permissions of all tokens. Tokens can inherit permissions from
	// their root object.
	GetPermissions(ctx context.Context) (*TokenPermissions, error)

	// Lists all tokens associated with the specified workspace or user.
	List(ctx context.Context, request ListTokenManagementRequest) (*ListTokensResponse, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error)

	// Updates the permissions on all tokens. Tokens can inherit permissions
	// from their root object.
	UpdatePermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error)

	// Updates a token, specified by its ID.
	UpdateTokenManagement(ctx context.Context, request UpdateTokenManagementRequest) (*TokenInfo, error)
}

// The Token API allows you to create, list, and revoke tokens that can be used
// to authenticate and access Databricks REST APIs.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type TokensService interface {

	// Creates and returns a token for a user. If this call is made through
	// token authentication, it creates a token with the same client ID as the
	// authenticated token. If the user's token quota is exceeded, this call
	// returns an error **QUOTA_EXCEEDED**.
	Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error)

	// Revokes an access token.
	//
	// If a token with the specified ID is not valid, this call returns an error
	// **RESOURCE_DOES_NOT_EXIST**.
	Delete(ctx context.Context, request RevokeTokenRequest) error

	// Lists all the valid tokens for a user-workspace pair.
	List(ctx context.Context) (*ListPublicTokensResponse, error)

	// Updates the comment or scopes of a token.
	//
	// If a token with the specified ID is not valid, this call returns an error
	// **NOT_FOUND**.
	Update(ctx context.Context, request UpdateTokenRequest) (*UpdateTokenResponse, error)
}

// This API allows updating known workspace settings for advanced users.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceConfService interface {

	// Gets the configuration status for a workspace.
	GetStatus(ctx context.Context, request GetStatusRequest) (*map[string]string, error)

	// Sets the configuration status for a workspace, including enabling or
	// disabling it.
	SetStatus(ctx context.Context, request WorkspaceConf) error
}

// These APIs allow configuration of network settings for Databricks workspaces
// by selecting which network policy to associate with the workspace. Each
// workspace is always associated with exactly one network policy that controls
// which network destinations can be accessed from the Databricks environment.
// By default, workspaces are associated with the 'default-policy' network
// policy. You cannot create or delete a workspace's network option, only update
// it to associate the workspace with a different policy
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceNetworkConfigurationService interface {

	// Gets the network option for a workspace. Every workspace has exactly one
	// network policy binding, with 'default-policy' used if no explicit
	// assignment exists.
	GetWorkspaceNetworkOptionRpc(ctx context.Context, request GetWorkspaceNetworkOptionRequest) (*WorkspaceNetworkOption, error)

	// Updates the network option for a workspace. This operation associates the
	// workspace with the specified network policy. To revert to the default
	// policy, specify 'default-policy' as the network_policy_id.
	UpdateWorkspaceNetworkOptionRpc(ctx context.Context, request UpdateWorkspaceNetworkOptionRequest) (*WorkspaceNetworkOption, error)
}
