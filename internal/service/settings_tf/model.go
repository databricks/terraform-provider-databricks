// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package settings_tf

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AutomaticClusterUpdateSetting struct {
	AutomaticClusterUpdateWorkspace []ClusterAutoRestartMessage `tfsdk:"automatic_cluster_update_workspace" tf:"object"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

type BooleanMessage struct {
	Value types.Bool `tfsdk:"value" tf:"optional"`
}

type ClusterAutoRestartMessage struct {
	CanToggle types.Bool `tfsdk:"can_toggle" tf:"optional"`

	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// Contains an information about the enablement status judging (e.g. whether
	// the enterprise tier is enabled) This is only additional information that
	// MUST NOT be used to decide whether the setting is enabled or not. This is
	// intended to use only for purposes like showing an error message to the
	// customer with the additional details. For example, using these details we
	// can check why exactly the feature is disabled for this customer.
	EnablementDetails []ClusterAutoRestartMessageEnablementDetails `tfsdk:"enablement_details" tf:"optional,object"`

	MaintenanceWindow []ClusterAutoRestartMessageMaintenanceWindow `tfsdk:"maintenance_window" tf:"optional,object"`

	RestartEvenIfNoUpdatesAvailable types.Bool `tfsdk:"restart_even_if_no_updates_available" tf:"optional"`
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode types.Bool `tfsdk:"forced_for_compliance_mode" tf:"optional"`
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	UnavailableForDisabledEntitlement types.Bool `tfsdk:"unavailable_for_disabled_entitlement" tf:"optional"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier types.Bool `tfsdk:"unavailable_for_non_enterprise_tier" tf:"optional"`
}

type ClusterAutoRestartMessageMaintenanceWindow struct {
	WeekDayBasedSchedule []ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule `tfsdk:"week_day_based_schedule" tf:"optional,object"`
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	DayOfWeek types.String `tfsdk:"day_of_week" tf:"optional"`

	Frequency types.String `tfsdk:"frequency" tf:"optional"`

	WindowStartTime []ClusterAutoRestartMessageMaintenanceWindowWindowStartTime `tfsdk:"window_start_time" tf:"optional,object"`
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	Hours types.Int64 `tfsdk:"hours" tf:"optional"`

	Minutes types.Int64 `tfsdk:"minutes" tf:"optional"`
}

// SHIELD feature: CSP
type ComplianceSecurityProfile struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	ComplianceStandards []types.String `tfsdk:"compliance_standards" tf:"optional"`

	IsEnabled types.Bool `tfsdk:"is_enabled" tf:"optional"`
}

type ComplianceSecurityProfileSetting struct {
	// SHIELD feature: CSP
	ComplianceSecurityProfileWorkspace []ComplianceSecurityProfile `tfsdk:"compliance_security_profile_workspace" tf:"object"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

type Config struct {
	Email []EmailConfig `tfsdk:"email" tf:"optional,object"`

	GenericWebhook []GenericWebhookConfig `tfsdk:"generic_webhook" tf:"optional,object"`

	MicrosoftTeams []MicrosoftTeamsConfig `tfsdk:"microsoft_teams" tf:"optional,object"`

	Pagerduty []PagerdutyConfig `tfsdk:"pagerduty" tf:"optional,object"`

	Slack []SlackConfig `tfsdk:"slack" tf:"optional,object"`
}

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {
	IpAddresses []types.String `tfsdk:"ip_addresses" tf:"optional"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label" tf:""`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type" tf:""`
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList []IpAccessListInfo `tfsdk:"ip_access_list" tf:"optional,object"`
}

type CreateNetworkConnectivityConfigRequest struct {
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name types.String `tfsdk:"name" tf:""`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region types.String `tfsdk:"region" tf:""`
}

type CreateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config []Config `tfsdk:"config" tf:"optional,object"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	ApplicationId types.String `tfsdk:"application_id" tf:""`
	// Comment that describes the purpose of the token.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The number of seconds before the token expires.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds" tf:"optional"`
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {
	TokenInfo []TokenInfo `tfsdk:"token_info" tf:"optional,object"`
	// Value of the token.
	TokenValue types.String `tfsdk:"token_value" tf:"optional"`
}

type CreatePrivateEndpointRuleRequest struct {
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId types.String `tfsdk:"group_id" tf:""`
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// The Azure resource ID of the target resource.
	ResourceId types.String `tfsdk:"resource_id" tf:""`
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds" tf:"optional"`
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo []PublicTokenInfo `tfsdk:"token_info" tf:"optional,object"`
	// The value of the new token.
	TokenValue types.String `tfsdk:"token_value" tf:"optional"`
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards []types.String `tfsdk:"compliance_standards" tf:"optional"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced types.Bool `tfsdk:"is_enforced" tf:"optional"`
}

type CspEnablementAccountSetting struct {
	// Account level policy for CSP
	CspEnablementAccount []CspEnablementAccount `tfsdk:"csp_enablement_account" tf:"object"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

// This represents the setting configuration for the default namespace in the
// Databricks workspace. Setting the default catalog for the workspace
// determines the catalog that is used when queries do not reference a fully
// qualified 3 level name. For example, if the default catalog is set to
// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the object
// 'retail_prod.default.myTable' (the schema 'default' is always assumed). This
// setting requires a restart of clusters and SQL warehouses to take effect.
// Additionally, the default namespace only applies when using Unity
// Catalog-enabled compute.
type DefaultNamespaceSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`

	Namespace []StringMessage `tfsdk:"namespace" tf:"object"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

// Delete the default namespace setting
type DeleteDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// The etag is returned.
type DeleteDefaultNamespaceSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag" tf:""`
}

// Delete Legacy Access Disablement Status
type DeleteDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// The etag is returned.
type DeleteDisableLegacyAccessResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag" tf:""`
}

// Delete the disable legacy features setting
type DeleteDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// The etag is returned.
type DeleteDisableLegacyFeaturesResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag" tf:""`
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

type DeleteNetworkConnectivityConfigurationResponse struct {
}

// Delete a notification destination
type DeleteNotificationDestinationRequest struct {
	Id types.String `tfsdk:"-"`
}

// Delete Personal Compute setting
type DeletePersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// The etag is returned.
type DeletePersonalComputeSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag" tf:""`
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

type DeleteResponse struct {
}

// Delete the restrict workspace admins setting
type DeleteRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// The etag is returned.
type DeleteRestrictWorkspaceAdminsSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag" tf:""`
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId types.String `tfsdk:"-"`
}

type DisableLegacyAccess struct {
	DisableLegacyAccess []BooleanMessage `tfsdk:"disable_legacy_access" tf:"object"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

type DisableLegacyFeatures struct {
	DisableLegacyFeatures []BooleanMessage `tfsdk:"disable_legacy_features" tf:"object"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

type EmailConfig struct {
	// Email addresses to notify.
	Addresses []types.String `tfsdk:"addresses" tf:"optional"`
}

type Empty struct {
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {
	IsEnabled types.Bool `tfsdk:"is_enabled" tf:"optional"`
}

type EnhancedSecurityMonitoringSetting struct {
	// SHIELD feature: ESM
	EnhancedSecurityMonitoringWorkspace []EnhancedSecurityMonitoring `tfsdk:"enhanced_security_monitoring_workspace" tf:"object"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

// Account level policy for ESM
type EsmEnablementAccount struct {
	IsEnforced types.Bool `tfsdk:"is_enforced" tf:"optional"`
}

type EsmEnablementAccountSetting struct {
	// Account level policy for ESM
	EsmEnablementAccount []EsmEnablementAccount `tfsdk:"esm_enablement_account" tf:"object"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken struct {
	// The requested token.
	Credential types.String `tfsdk:"credential" tf:"optional"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime types.Int64 `tfsdk:"credentialEolTime" tf:"optional"`
	// User ID of the user that owns this token.
	OwnerId types.Int64 `tfsdk:"ownerId" tf:"optional"`
	// The scopes of access granted in the token.
	Scopes []types.String `tfsdk:"scopes" tf:"optional"`
	// The type of this exchange token
	TokenType types.String `tfsdk:"tokenType" tf:"optional"`
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	PartitionId []PartitionId `tfsdk:"partitionId" tf:"object"`
	// Array of scopes for the token request.
	Scopes []types.String `tfsdk:"scopes" tf:""`
	// A list of token types being requested
	TokenType []types.String `tfsdk:"tokenType" tf:""`
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {
	Values []ExchangeToken `tfsdk:"values" tf:"optional"`
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList []IpAccessListInfo `tfsdk:"ip_access_list" tf:"optional,object"`
}

type GenericWebhookConfig struct {
	// [Input-Only][Optional] Password for webhook.
	Password types.String `tfsdk:"password" tf:"optional"`
	// [Output-Only] Whether password is set.
	PasswordSet types.Bool `tfsdk:"password_set" tf:"optional"`
	// [Input-Only] URL for webhook.
	Url types.String `tfsdk:"url" tf:"optional"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set" tf:"optional"`
	// [Input-Only][Optional] Username for webhook.
	Username types.String `tfsdk:"username" tf:"optional"`
	// [Output-Only] Whether username is set.
	UsernameSet types.Bool `tfsdk:"username_set" tf:"optional"`
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

// Get the automatic cluster update setting
type GetAutomaticClusterUpdateSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Get the compliance security profile setting
type GetComplianceSecurityProfileSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Get the compliance security profile setting for new workspaces
type GetCspEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Get the default namespace setting
type GetDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Retrieve Legacy Access Disablement Status
type GetDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Get the disable legacy features setting
type GetDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Get the enhanced security monitoring setting
type GetEnhancedSecurityMonitoringSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Get the enhanced security monitoring setting for new workspaces
type GetEsmEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

type GetIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList []IpAccessListInfo `tfsdk:"ip_access_list" tf:"optional,object"`
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {
	IpAccessLists []IpAccessListInfo `tfsdk:"ip_access_lists" tf:"optional"`
}

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

// Get a notification destination
type GetNotificationDestinationRequest struct {
	Id types.String `tfsdk:"-"`
}

// Get Personal Compute setting
type GetPersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Get a private endpoint rule
type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

// Get the restrict workspace admins setting
type GetRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

// Check configuration status
type GetStatusRequest struct {
	Keys types.String `tfsdk:"-"`
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId types.String `tfsdk:"-"`
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []TokenPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {
	TokenInfo []TokenInfo `tfsdk:"token_info" tf:"optional,object"`
}

// Definition of an IP Access list
type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	AddressCount types.Int64 `tfsdk:"address_count" tf:"optional"`
	// Creation timestamp in milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// User ID of the user who created this list.
	CreatedBy types.Int64 `tfsdk:"created_by" tf:"optional"`
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`

	IpAddresses []types.String `tfsdk:"ip_addresses" tf:"optional"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label" tf:"optional"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId types.String `tfsdk:"list_id" tf:"optional"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type" tf:"optional"`
	// Update timestamp in milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// User ID of the user who updated this list.
	UpdatedBy types.Int64 `tfsdk:"updated_by" tf:"optional"`
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists []IpAccessListInfo `tfsdk:"ip_access_lists" tf:"optional"`
}

type ListNccAzurePrivateEndpointRulesResponse struct {
	Items []NccAzurePrivateEndpointRule `tfsdk:"items" tf:"optional"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

type ListNetworkConnectivityConfigurationsResponse struct {
	Items []NetworkConnectivityConfiguration `tfsdk:"items" tf:"optional"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List notification destinations
type ListNotificationDestinationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListNotificationDestinationsResponse struct {
	// Page token for next of results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Results []ListNotificationDestinationsResult `tfsdk:"results" tf:"optional"`
}

type ListNotificationDestinationsResult struct {
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType types.String `tfsdk:"destination_type" tf:"optional"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"id" tf:"optional"`
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos []PublicTokenInfo `tfsdk:"token_infos" tf:"optional"`
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById types.Int64 `tfsdk:"-"`
	// Username of the user that created the token.
	CreatedByUsername types.String `tfsdk:"-"`
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos []TokenInfo `tfsdk:"token_infos" tf:"optional"`
}

type MicrosoftTeamsConfig struct {
	// [Input-Only] URL for Microsoft Teams.
	Url types.String `tfsdk:"url" tf:"optional"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set" tf:"optional"`
}

// The stable AWS IP CIDR blocks. You can use these to configure the firewall of
// your resources to allow traffic from your Databricks workspace.
type NccAwsStableIpRule struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	CidrBlocks []types.String `tfsdk:"cidr_blocks" tf:"optional"`
}

type NccAzurePrivateEndpointRule struct {
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is `ESTABLISHED`. Remember
	// that you must approve new endpoints on your resources in the Azure portal
	// before they take effect.
	//
	// The possible values are: - INIT: (deprecated) The endpoint has been
	// created and pending approval. - PENDING: The endpoint has been created
	// and pending approval. - ESTABLISHED: The endpoint has been approved and
	// is ready to use in your serverless compute resources. - REJECTED:
	// Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up.
	ConnectionState types.String `tfsdk:"connection_state" tf:"optional"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"optional"`
	// Whether this private endpoint is deactivated.
	Deactivated types.Bool `tfsdk:"deactivated" tf:"optional"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt types.Int64 `tfsdk:"deactivated_at" tf:"optional"`
	// The name of the Azure private endpoint resource.
	EndpointName types.String `tfsdk:"endpoint_name" tf:"optional"`
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId types.String `tfsdk:"group_id" tf:"optional"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id" tf:"optional"`
	// The Azure resource ID of the target resource.
	ResourceId types.String `tfsdk:"resource_id" tf:"optional"`
	// The ID of a private endpoint rule.
	RuleId types.String `tfsdk:"rule_id" tf:"optional"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time" tf:"optional"`
}

// The stable Azure service endpoints. You can configure the firewall of your
// Azure resources to allow traffic from your Databricks serverless compute
// resources.
type NccAzureServiceEndpointRule struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets []types.String `tfsdk:"subnets" tf:"optional"`
	// The Azure region in which this service endpoint rule applies.
	TargetRegion types.String `tfsdk:"target_region" tf:"optional"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices []types.String `tfsdk:"target_services" tf:"optional"`
}

// The network connectivity rules that apply to network traffic from your
// serverless compute resources.
type NccEgressConfig struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules []NccEgressDefaultRules `tfsdk:"default_rules" tf:"optional,object"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules []NccEgressTargetRules `tfsdk:"target_rules" tf:"optional,object"`
}

// The network connectivity rules that are applied by default without resource
// specific configurations. You can find the stable network information of your
// serverless compute resources here.
type NccEgressDefaultRules struct {
	// The stable AWS IP CIDR blocks. You can use these to configure the
	// firewall of your resources to allow traffic from your Databricks
	// workspace.
	AwsStableIpRule []NccAwsStableIpRule `tfsdk:"aws_stable_ip_rule" tf:"optional,object"`
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	AzureServiceEndpointRule []NccAzureServiceEndpointRule `tfsdk:"azure_service_endpoint_rule" tf:"optional,object"`
}

// The network connectivity rules that configured for each destinations. These
// rules override default rules.
type NccEgressTargetRules struct {
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRule `tfsdk:"azure_private_endpoint_rules" tf:"optional"`
}

type NetworkConnectivityConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"optional"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig []NccEgressConfig `tfsdk:"egress_config" tf:"optional,object"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Databricks network connectivity configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id" tf:"optional"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region types.String `tfsdk:"region" tf:"optional"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time" tf:"optional"`
}

type NotificationDestination struct {
	// The configuration for the notification destination. Will be exactly one
	// of the nested configs. Only returns for users with workspace admin
	// permissions.
	Config []Config `tfsdk:"config" tf:"optional,object"`
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType types.String `tfsdk:"destination_type" tf:"optional"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"id" tf:"optional"`
}

type PagerdutyConfig struct {
	// [Input-Only] Integration key for PagerDuty.
	IntegrationKey types.String `tfsdk:"integration_key" tf:"optional"`
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet types.Bool `tfsdk:"integration_key_set" tf:"optional"`
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId types.Int64 `tfsdk:"workspaceId" tf:"optional"`
}

type PersonalComputeMessage struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value types.String `tfsdk:"value" tf:""`
}

type PersonalComputeSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`

	PersonalCompute []PersonalComputeMessage `tfsdk:"personal_compute" tf:"object"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

type PublicTokenInfo struct {
	// Comment the token was created with, if applicable.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Server time (in epoch milliseconds) when the token was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"optional"`
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	ExpiryTime types.Int64 `tfsdk:"expiry_time" tf:"optional"`
	// The ID of this token.
	TokenId types.String `tfsdk:"token_id" tf:"optional"`
}

// Details required to replace an IP access list.
type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:""`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`

	IpAddresses []types.String `tfsdk:"ip_addresses" tf:"optional"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label" tf:""`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type" tf:""`
}

type ReplaceResponse struct {
}

type RestrictWorkspaceAdminsMessage struct {
	Status types.String `tfsdk:"status" tf:""`
}

type RestrictWorkspaceAdminsSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag" tf:"optional"`

	RestrictWorkspaceAdmins []RestrictWorkspaceAdminsMessage `tfsdk:"restrict_workspace_admins" tf:"object"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId types.String `tfsdk:"token_id" tf:""`
}

type RevokeTokenResponse struct {
}

type SetStatusResponse struct {
}

type SlackConfig struct {
	// [Input-Only] URL for Slack destination.
	Url types.String `tfsdk:"url" tf:"optional"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set" tf:"optional"`
}

type StringMessage struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions []TokenPermission `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// User ID of the user that created the token.
	CreatedById types.Int64 `tfsdk:"created_by_id" tf:"optional"`
	// Username of the user that created the token.
	CreatedByUsername types.String `tfsdk:"created_by_username" tf:"optional"`
	// Timestamp when the token was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"optional"`
	// Timestamp when the token expires.
	ExpiryTime types.Int64 `tfsdk:"expiry_time" tf:"optional"`
	// User ID of the user that owns the token.
	OwnerId types.Int64 `tfsdk:"owner_id" tf:"optional"`
	// ID of the token.
	TokenId types.String `tfsdk:"token_id" tf:"optional"`
	// If applicable, the ID of the workspace that the token was created in.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

type TokenPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

type TokenPermissions struct {
	AccessControlList []TokenAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type TokenPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

type TokenPermissionsRequest struct {
	AccessControlList []TokenAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
}

// Details required to update a setting.
type UpdateAutomaticClusterUpdateSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting []AutomaticClusterUpdateSetting `tfsdk:"setting" tf:"object"`
}

// Details required to update a setting.
type UpdateComplianceSecurityProfileSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting []ComplianceSecurityProfileSetting `tfsdk:"setting" tf:"object"`
}

// Details required to update a setting.
type UpdateCspEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting []CspEnablementAccountSetting `tfsdk:"setting" tf:"object"`
}

// Details required to update a setting.
type UpdateDefaultNamespaceSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`
	// This represents the setting configuration for the default namespace in
	// the Databricks workspace. Setting the default catalog for the workspace
	// determines the catalog that is used when queries do not reference a fully
	// qualified 3 level name. For example, if the default catalog is set to
	// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the
	// object 'retail_prod.default.myTable' (the schema 'default' is always
	// assumed). This setting requires a restart of clusters and SQL warehouses
	// to take effect. Additionally, the default namespace only applies when
	// using Unity Catalog-enabled compute.
	Setting []DefaultNamespaceSetting `tfsdk:"setting" tf:"object"`
}

// Details required to update a setting.
type UpdateDisableLegacyAccessRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting []DisableLegacyAccess `tfsdk:"setting" tf:"object"`
}

// Details required to update a setting.
type UpdateDisableLegacyFeaturesRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting []DisableLegacyFeatures `tfsdk:"setting" tf:"object"`
}

// Details required to update a setting.
type UpdateEnhancedSecurityMonitoringSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting []EnhancedSecurityMonitoringSetting `tfsdk:"setting" tf:"object"`
}

// Details required to update a setting.
type UpdateEsmEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting []EsmEnablementAccountSetting `tfsdk:"setting" tf:"object"`
}

// Details required to update an IP access list.
type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`

	IpAddresses []types.String `tfsdk:"ip_addresses" tf:"optional"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label" tf:"optional"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type" tf:"optional"`
}

type UpdateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config []Config `tfsdk:"config" tf:"optional,object"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`

	Id types.String `tfsdk:"-"`
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting []PersonalComputeSetting `tfsdk:"setting" tf:"object"`
}

type UpdateResponse struct {
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting []RestrictWorkspaceAdminsSetting `tfsdk:"setting" tf:"object"`
}
