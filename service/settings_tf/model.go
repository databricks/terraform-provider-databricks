// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings_tf

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AutomaticClusterUpdateSetting struct {
	AutomaticClusterUpdateWorkspace ClusterAutoRestartMessage `tfsdk:"automatic_cluster_update_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

type ClusterAutoRestartMessage struct {
	CanToggle types.Bool `tfsdk:"can_toggle"`

	Enabled types.Bool `tfsdk:"enabled"`
	// Contains an information about the enablement status judging (e.g. whether
	// the enterprise tier is enabled) This is only additional information that
	// MUST NOT be used to decide whether the setting is enabled or not. This is
	// intended to use only for purposes like showing an error message to the
	// customer with the additional details. For example, using these details we
	// can check why exactly the feature is disabled for this customer.
	EnablementDetails *ClusterAutoRestartMessageEnablementDetails `tfsdk:"enablement_details"`

	MaintenanceWindow *ClusterAutoRestartMessageMaintenanceWindow `tfsdk:"maintenance_window"`

	RestartEvenIfNoUpdatesAvailable types.Bool `tfsdk:"restart_even_if_no_updates_available"`
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode types.Bool `tfsdk:"forced_for_compliance_mode"`
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	UnavailableForDisabledEntitlement types.Bool `tfsdk:"unavailable_for_disabled_entitlement"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier types.Bool `tfsdk:"unavailable_for_non_enterprise_tier"`
}

type ClusterAutoRestartMessageMaintenanceWindow struct {
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule `tfsdk:"week_day_based_schedule"`
}

type ClusterAutoRestartMessageMaintenanceWindowDayOfWeek string

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekDayOfWeekUnspecified ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `DAY_OF_WEEK_UNSPECIFIED`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekFriday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `FRIDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekMonday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `MONDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSaturday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `SATURDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSunday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `SUNDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekThursday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `THURSDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekTuesday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `TUESDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekWednesday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `WEDNESDAY`

// String representation for [fmt.Print]
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Set(v string) error {
	switch v {
	case `DAY_OF_WEEK_UNSPECIFIED`, `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`:
		*f = ClusterAutoRestartMessageMaintenanceWindowDayOfWeek(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAY_OF_WEEK_UNSPECIFIED", "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY", "WEDNESDAY"`, v)
	}
}

// Type always returns ClusterAutoRestartMessageMaintenanceWindowDayOfWeek to satisfy [pflag.Value] interface
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Type() string {
	return "ClusterAutoRestartMessageMaintenanceWindowDayOfWeek"
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	DayOfWeek ClusterAutoRestartMessageMaintenanceWindowDayOfWeek `tfsdk:"day_of_week"`

	Frequency ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency `tfsdk:"frequency"`

	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime `tfsdk:"window_start_time"`
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency string

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyEveryWeek ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `EVERY_WEEK`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstAndThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FIRST_AND_THIRD_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FIRST_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondAndFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `SECOND_AND_FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `SECOND_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `THIRD_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyWeekDayFrequencyUnspecified ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `WEEK_DAY_FREQUENCY_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Set(v string) error {
	switch v {
	case `EVERY_WEEK`, `FIRST_AND_THIRD_OF_MONTH`, `FIRST_OF_MONTH`, `FOURTH_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`, `WEEK_DAY_FREQUENCY_UNSPECIFIED`:
		*f = ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EVERY_WEEK", "FIRST_AND_THIRD_OF_MONTH", "FIRST_OF_MONTH", "FOURTH_OF_MONTH", "SECOND_AND_FOURTH_OF_MONTH", "SECOND_OF_MONTH", "THIRD_OF_MONTH", "WEEK_DAY_FREQUENCY_UNSPECIFIED"`, v)
	}
}

// Type always returns ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency to satisfy [pflag.Value] interface
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Type() string {
	return "ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency"
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	Hours types.Int64 `tfsdk:"hours"`

	Minutes types.Int64 `tfsdk:"minutes"`
}

// SHIELD feature: CSP
type ComplianceSecurityProfile struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	ComplianceStandards types.List `tfsdk:"compliance_standards"`

	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

type ComplianceSecurityProfileSetting struct {
	// SHIELD feature: CSP
	ComplianceSecurityProfileWorkspace ComplianceSecurityProfile `tfsdk:"compliance_security_profile_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

// Compliance stardard for SHIELD customers
type ComplianceStandard string

const ComplianceStandardComplianceStandardUnspecified ComplianceStandard = `COMPLIANCE_STANDARD_UNSPECIFIED`

const ComplianceStandardCyberEssentialPlus ComplianceStandard = `CYBER_ESSENTIAL_PLUS`

const ComplianceStandardFedrampHigh ComplianceStandard = `FEDRAMP_HIGH`

const ComplianceStandardFedrampIl5 ComplianceStandard = `FEDRAMP_IL5`

const ComplianceStandardFedrampModerate ComplianceStandard = `FEDRAMP_MODERATE`

const ComplianceStandardHipaa ComplianceStandard = `HIPAA`

const ComplianceStandardIrapProtected ComplianceStandard = `IRAP_PROTECTED`

const ComplianceStandardItarEar ComplianceStandard = `ITAR_EAR`

const ComplianceStandardNone ComplianceStandard = `NONE`

const ComplianceStandardPciDss ComplianceStandard = `PCI_DSS`

// String representation for [fmt.Print]
func (f *ComplianceStandard) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComplianceStandard) Set(v string) error {
	switch v {
	case `COMPLIANCE_STANDARD_UNSPECIFIED`, `CYBER_ESSENTIAL_PLUS`, `FEDRAMP_HIGH`, `FEDRAMP_IL5`, `FEDRAMP_MODERATE`, `HIPAA`, `IRAP_PROTECTED`, `ITAR_EAR`, `NONE`, `PCI_DSS`:
		*f = ComplianceStandard(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMPLIANCE_STANDARD_UNSPECIFIED", "CYBER_ESSENTIAL_PLUS", "FEDRAMP_HIGH", "FEDRAMP_IL5", "FEDRAMP_MODERATE", "HIPAA", "IRAP_PROTECTED", "ITAR_EAR", "NONE", "PCI_DSS"`, v)
	}
}

// Type always returns ComplianceStandard to satisfy [pflag.Value] interface
func (f *ComplianceStandard) Type() string {
	return "ComplianceStandard"
}

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {
	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `tfsdk:"list_type"`
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `tfsdk:"ip_access_list"`
}

type CreateNetworkConnectivityConfigRequest struct {
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name types.String `tfsdk:"name"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region types.String `tfsdk:"region"`
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	ApplicationId types.String `tfsdk:"application_id"`
	// Comment that describes the purpose of the token.
	Comment types.String `tfsdk:"comment"`
	// The number of seconds before the token expires.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds"`
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo `tfsdk:"token_info"`
	// Value of the token.
	TokenValue types.String `tfsdk:"token_value"`
}

type CreatePrivateEndpointRuleRequest struct {
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId CreatePrivateEndpointRuleRequestGroupId `tfsdk:"group_id"`
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-" url:"-"`
	// The Azure resource ID of the target resource.
	ResourceId types.String `tfsdk:"resource_id"`
}

// The sub-resource type (group ID) of the target resource. Note that to connect
// to workspace root storage (root DBFS), you need two endpoints, one for `blob`
// and one for `dfs`.
type CreatePrivateEndpointRuleRequestGroupId string

const CreatePrivateEndpointRuleRequestGroupIdBlob CreatePrivateEndpointRuleRequestGroupId = `blob`

const CreatePrivateEndpointRuleRequestGroupIdDfs CreatePrivateEndpointRuleRequestGroupId = `dfs`

const CreatePrivateEndpointRuleRequestGroupIdMysqlServer CreatePrivateEndpointRuleRequestGroupId = `mysqlServer`

const CreatePrivateEndpointRuleRequestGroupIdSqlServer CreatePrivateEndpointRuleRequestGroupId = `sqlServer`

// String representation for [fmt.Print]
func (f *CreatePrivateEndpointRuleRequestGroupId) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreatePrivateEndpointRuleRequestGroupId) Set(v string) error {
	switch v {
	case `blob`, `dfs`, `mysqlServer`, `sqlServer`:
		*f = CreatePrivateEndpointRuleRequestGroupId(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "blob", "dfs", "mysqlServer", "sqlServer"`, v)
	}
}

// Type always returns CreatePrivateEndpointRuleRequestGroupId to satisfy [pflag.Value] interface
func (f *CreatePrivateEndpointRuleRequestGroupId) Type() string {
	return "CreatePrivateEndpointRuleRequestGroupId"
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	Comment types.String `tfsdk:"comment"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds"`
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo *PublicTokenInfo `tfsdk:"token_info"`
	// The value of the new token.
	TokenValue types.String `tfsdk:"token_value"`
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards types.List `tfsdk:"compliance_standards"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

type CspEnablementAccountSetting struct {
	// Account level policy for CSP
	CspEnablementAccount CspEnablementAccount `tfsdk:"csp_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
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
	Etag types.String `tfsdk:"etag"`

	Namespace StringMessage `tfsdk:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-" url:"-"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
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
	Etag types.String `tfsdk:"etag"`
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-" url:"-"`
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-" url:"-"`
}

type DeleteNetworkConnectivityConfigurationResponse struct {
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
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
	Etag types.String `tfsdk:"etag"`
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-" url:"-"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
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
	Etag types.String `tfsdk:"etag"`
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId types.String `tfsdk:"-" url:"-"`
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {
	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

type EnhancedSecurityMonitoringSetting struct {
	// SHIELD feature: ESM
	EnhancedSecurityMonitoringWorkspace EnhancedSecurityMonitoring `tfsdk:"enhanced_security_monitoring_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

// Account level policy for ESM
type EsmEnablementAccount struct {
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

type EsmEnablementAccountSetting struct {
	// Account level policy for ESM
	EsmEnablementAccount EsmEnablementAccount `tfsdk:"esm_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken struct {
	// The requested token.
	Credential types.String `tfsdk:"credential"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime types.Int64 `tfsdk:"credentialEolTime"`
	// User ID of the user that owns this token.
	OwnerId types.Int64 `tfsdk:"ownerId"`
	// The scopes of access granted in the token.
	Scopes types.List `tfsdk:"scopes"`
	// The type of this exchange token
	TokenType TokenType `tfsdk:"tokenType"`
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	PartitionId PartitionId `tfsdk:"partitionId"`
	// Array of scopes for the token request.
	Scopes types.List `tfsdk:"scopes"`
	// A list of token types being requested
	TokenType types.List `tfsdk:"tokenType"`
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {
	Values types.List `tfsdk:"values"`
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `tfsdk:"ip_access_list"`
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-" url:"-"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-" url:"-"`
}

type GetIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `tfsdk:"ip_access_list"`
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-" url:"-"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
}

// Get a private endpoint rule
type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-" url:"-"`
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
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
}

// Check configuration status
type GetStatusRequest struct {
	Keys types.String `tfsdk:"-" url:"keys"`
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId types.String `tfsdk:"-" url:"-"`
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {
	TokenInfo *TokenInfo `tfsdk:"token_info"`
}

// Definition of an IP Access list
type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	AddressCount types.Int64 `tfsdk:"address_count"`
	// Creation timestamp in milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// User ID of the user who created this list.
	CreatedBy types.Int64 `tfsdk:"created_by"`
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled"`

	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId types.String `tfsdk:"list_id"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `tfsdk:"list_type"`
	// Update timestamp in milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// User ID of the user who updated this list.
	UpdatedBy types.Int64 `tfsdk:"updated_by"`
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

type ListNccAzurePrivateEndpointRulesResponse struct {
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListNetworkConnectivityConfigurationsResponse struct {
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-" url:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos types.List `tfsdk:"token_infos"`
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById types.Int64 `tfsdk:"-" url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername types.String `tfsdk:"-" url:"created_by_username,omitempty"`
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos types.List `tfsdk:"token_infos"`
}

// Type of IP access list. Valid values are as follows and are case-sensitive:
//
// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block list.
// Exclude this IP or range. IP addresses in the block list are excluded even if
// they are included in an allow list.
type ListType string

// An allow list. Include this IP or range.
const ListTypeAllow ListType = `ALLOW`

// A block list. Exclude this IP or range. IP addresses in the block list are
// excluded even if they are included in an allow list.
const ListTypeBlock ListType = `BLOCK`

// String representation for [fmt.Print]
func (f *ListType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListType) Set(v string) error {
	switch v {
	case `ALLOW`, `BLOCK`:
		*f = ListType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW", "BLOCK"`, v)
	}
}

// Type always returns ListType to satisfy [pflag.Value] interface
func (f *ListType) Type() string {
	return "ListType"
}

// The stable AWS IP CIDR blocks. You can use these to configure the firewall of
// your resources to allow traffic from your Databricks workspace.
type NccAwsStableIpRule struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	CidrBlocks types.List `tfsdk:"cidr_blocks"`
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
	ConnectionState NccAzurePrivateEndpointRuleConnectionState `tfsdk:"connection_state"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Whether this private endpoint is deactivated.
	Deactivated types.Bool `tfsdk:"deactivated"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt types.Int64 `tfsdk:"deactivated_at"`
	// The name of the Azure private endpoint resource.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId NccAzurePrivateEndpointRuleGroupId `tfsdk:"group_id"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// The Azure resource ID of the target resource.
	ResourceId types.String `tfsdk:"resource_id"`
	// The ID of a private endpoint rule.
	RuleId types.String `tfsdk:"rule_id"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time"`
}

// The current status of this private endpoint. The private endpoint rules are
// effective only if the connection state is `ESTABLISHED`. Remember that you
// must approve new endpoints on your resources in the Azure portal before they
// take effect.
//
// The possible values are: - INIT: (deprecated) The endpoint has been created
// and pending approval. - PENDING: The endpoint has been created and pending
// approval. - ESTABLISHED: The endpoint has been approved and is ready to use
// in your serverless compute resources. - REJECTED: Connection was rejected by
// the private link resource owner. - DISCONNECTED: Connection was removed by
// the private link resource owner, the private endpoint becomes informative and
// should be deleted for clean-up.
type NccAzurePrivateEndpointRuleConnectionState string

const NccAzurePrivateEndpointRuleConnectionStateDisconnected NccAzurePrivateEndpointRuleConnectionState = `DISCONNECTED`

const NccAzurePrivateEndpointRuleConnectionStateEstablished NccAzurePrivateEndpointRuleConnectionState = `ESTABLISHED`

const NccAzurePrivateEndpointRuleConnectionStateInit NccAzurePrivateEndpointRuleConnectionState = `INIT`

const NccAzurePrivateEndpointRuleConnectionStatePending NccAzurePrivateEndpointRuleConnectionState = `PENDING`

const NccAzurePrivateEndpointRuleConnectionStateRejected NccAzurePrivateEndpointRuleConnectionState = `REJECTED`

// String representation for [fmt.Print]
func (f *NccAzurePrivateEndpointRuleConnectionState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NccAzurePrivateEndpointRuleConnectionState) Set(v string) error {
	switch v {
	case `DISCONNECTED`, `ESTABLISHED`, `INIT`, `PENDING`, `REJECTED`:
		*f = NccAzurePrivateEndpointRuleConnectionState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISCONNECTED", "ESTABLISHED", "INIT", "PENDING", "REJECTED"`, v)
	}
}

// Type always returns NccAzurePrivateEndpointRuleConnectionState to satisfy [pflag.Value] interface
func (f *NccAzurePrivateEndpointRuleConnectionState) Type() string {
	return "NccAzurePrivateEndpointRuleConnectionState"
}

// The sub-resource type (group ID) of the target resource. Note that to connect
// to workspace root storage (root DBFS), you need two endpoints, one for `blob`
// and one for `dfs`.
type NccAzurePrivateEndpointRuleGroupId string

const NccAzurePrivateEndpointRuleGroupIdBlob NccAzurePrivateEndpointRuleGroupId = `blob`

const NccAzurePrivateEndpointRuleGroupIdDfs NccAzurePrivateEndpointRuleGroupId = `dfs`

const NccAzurePrivateEndpointRuleGroupIdMysqlServer NccAzurePrivateEndpointRuleGroupId = `mysqlServer`

const NccAzurePrivateEndpointRuleGroupIdSqlServer NccAzurePrivateEndpointRuleGroupId = `sqlServer`

// String representation for [fmt.Print]
func (f *NccAzurePrivateEndpointRuleGroupId) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NccAzurePrivateEndpointRuleGroupId) Set(v string) error {
	switch v {
	case `blob`, `dfs`, `mysqlServer`, `sqlServer`:
		*f = NccAzurePrivateEndpointRuleGroupId(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "blob", "dfs", "mysqlServer", "sqlServer"`, v)
	}
}

// Type always returns NccAzurePrivateEndpointRuleGroupId to satisfy [pflag.Value] interface
func (f *NccAzurePrivateEndpointRuleGroupId) Type() string {
	return "NccAzurePrivateEndpointRuleGroupId"
}

// The stable Azure service endpoints. You can configure the firewall of your
// Azure resources to allow traffic from your Databricks serverless compute
// resources.
type NccAzureServiceEndpointRule struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets types.List `tfsdk:"subnets"`
	// The Azure region in which this service endpoint rule applies.
	TargetRegion types.String `tfsdk:"target_region"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices types.List `tfsdk:"target_services"`
}

// The network connectivity rules that apply to network traffic from your
// serverless compute resources.
type NccEgressConfig struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules *NccEgressDefaultRules `tfsdk:"default_rules"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules *NccEgressTargetRules `tfsdk:"target_rules"`
}

// The network connectivity rules that are applied by default without resource
// specific configurations. You can find the stable network information of your
// serverless compute resources here.
type NccEgressDefaultRules struct {
	// The stable AWS IP CIDR blocks. You can use these to configure the
	// firewall of your resources to allow traffic from your Databricks
	// workspace.
	AwsStableIpRule *NccAwsStableIpRule `tfsdk:"aws_stable_ip_rule"`
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	AzureServiceEndpointRule *NccAzureServiceEndpointRule `tfsdk:"azure_service_endpoint_rule"`
}

// The network connectivity rules that configured for each destinations. These
// rules override default rules.
type NccEgressTargetRules struct {
	AzurePrivateEndpointRules types.List `tfsdk:"azure_private_endpoint_rules"`
}

type NetworkConnectivityConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig *NccEgressConfig `tfsdk:"egress_config"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name types.String `tfsdk:"name"`
	// Databricks network connectivity configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region types.String `tfsdk:"region"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time"`
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId types.Int64 `tfsdk:"workspaceId"`
}

type PersonalComputeMessage struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value PersonalComputeMessageEnum `tfsdk:"value"`
}

// ON: Grants all users in all workspaces access to the Personal Compute default
// policy, allowing all users to create single-machine compute resources.
// DELEGATE: Moves access control for the Personal Compute default policy to
// individual workspaces and requires a workspace’s users or groups to be
// added to the ACLs of that workspace’s Personal Compute default policy
// before they will be able to create compute resources through that policy.
type PersonalComputeMessageEnum string

const PersonalComputeMessageEnumDelegate PersonalComputeMessageEnum = `DELEGATE`

const PersonalComputeMessageEnumOn PersonalComputeMessageEnum = `ON`

// String representation for [fmt.Print]
func (f *PersonalComputeMessageEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PersonalComputeMessageEnum) Set(v string) error {
	switch v {
	case `DELEGATE`, `ON`:
		*f = PersonalComputeMessageEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELEGATE", "ON"`, v)
	}
}

// Type always returns PersonalComputeMessageEnum to satisfy [pflag.Value] interface
func (f *PersonalComputeMessageEnum) Type() string {
	return "PersonalComputeMessageEnum"
}

type PersonalComputeSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`

	PersonalCompute PersonalComputeMessage `tfsdk:"personal_compute"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

type PublicTokenInfo struct {
	// Comment the token was created with, if applicable.
	Comment types.String `tfsdk:"comment"`
	// Server time (in epoch milliseconds) when the token was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	ExpiryTime types.Int64 `tfsdk:"expiry_time"`
	// The ID of this token.
	TokenId types.String `tfsdk:"token_id"`
}

// Details required to replace an IP access list.
type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-" url:"-"`

	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `tfsdk:"list_type"`
}

type ReplaceResponse struct {
}

type RestrictWorkspaceAdminsMessage struct {
	Status RestrictWorkspaceAdminsMessageStatus `tfsdk:"status"`
}

type RestrictWorkspaceAdminsMessageStatus string

const RestrictWorkspaceAdminsMessageStatusAllowAll RestrictWorkspaceAdminsMessageStatus = `ALLOW_ALL`

const RestrictWorkspaceAdminsMessageStatusRestrictTokensAndJobRunAs RestrictWorkspaceAdminsMessageStatus = `RESTRICT_TOKENS_AND_JOB_RUN_AS`

const RestrictWorkspaceAdminsMessageStatusStatusUnspecified RestrictWorkspaceAdminsMessageStatus = `STATUS_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *RestrictWorkspaceAdminsMessageStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RestrictWorkspaceAdminsMessageStatus) Set(v string) error {
	switch v {
	case `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`, `STATUS_UNSPECIFIED`:
		*f = RestrictWorkspaceAdminsMessageStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL", "RESTRICT_TOKENS_AND_JOB_RUN_AS", "STATUS_UNSPECIFIED"`, v)
	}
}

// Type always returns RestrictWorkspaceAdminsMessageStatus to satisfy [pflag.Value] interface
func (f *RestrictWorkspaceAdminsMessageStatus) Type() string {
	return "RestrictWorkspaceAdminsMessageStatus"
}

type RestrictWorkspaceAdminsSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`

	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessage `tfsdk:"restrict_workspace_admins"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId types.String `tfsdk:"token_id"`
}

type RevokeTokenResponse struct {
}

type SetStatusResponse struct {
}

type StringMessage struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value"`
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel TokenPermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment types.String `tfsdk:"comment"`
	// User ID of the user that created the token.
	CreatedById types.Int64 `tfsdk:"created_by_id"`
	// Username of the user that created the token.
	CreatedByUsername types.String `tfsdk:"created_by_username"`
	// Timestamp when the token was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Timestamp when the token expires.
	ExpiryTime types.Int64 `tfsdk:"expiry_time"`
	// User ID of the user that owns the token.
	OwnerId types.Int64 `tfsdk:"owner_id"`
	// ID of the token.
	TokenId types.String `tfsdk:"token_id"`
}

type TokenPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel TokenPermissionLevel `tfsdk:"permission_level"`
}

// Permission level
type TokenPermissionLevel string

const TokenPermissionLevelCanUse TokenPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *TokenPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_USE`:
		*f = TokenPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_USE"`, v)
	}
}

// Type always returns TokenPermissionLevel to satisfy [pflag.Value] interface
func (f *TokenPermissionLevel) Type() string {
	return "TokenPermissionLevel"
}

type TokenPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

type TokenPermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel TokenPermissionLevel `tfsdk:"permission_level"`
}

type TokenPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
}

// The type of token request. As of now, only `AZURE_ACTIVE_DIRECTORY_TOKEN` is
// supported.
type TokenType string

const TokenTypeAzureActiveDirectoryToken TokenType = `AZURE_ACTIVE_DIRECTORY_TOKEN`

// String representation for [fmt.Print]
func (f *TokenType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenType) Set(v string) error {
	switch v {
	case `AZURE_ACTIVE_DIRECTORY_TOKEN`:
		*f = TokenType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AZURE_ACTIVE_DIRECTORY_TOKEN"`, v)
	}
}

// Type always returns TokenType to satisfy [pflag.Value] interface
func (f *TokenType) Type() string {
	return "TokenType"
}

// Details required to update a setting.
type UpdateAutomaticClusterUpdateSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask"`

	Setting AutomaticClusterUpdateSetting `tfsdk:"setting"`
}

// Details required to update a setting.
type UpdateComplianceSecurityProfileSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask"`

	Setting ComplianceSecurityProfileSetting `tfsdk:"setting"`
}

// Details required to update a setting.
type UpdateCspEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask"`

	Setting CspEnablementAccountSetting `tfsdk:"setting"`
}

// Details required to update a setting.
type UpdateDefaultNamespaceSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask"`
	// This represents the setting configuration for the default namespace in
	// the Databricks workspace. Setting the default catalog for the workspace
	// determines the catalog that is used when queries do not reference a fully
	// qualified 3 level name. For example, if the default catalog is set to
	// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the
	// object 'retail_prod.default.myTable' (the schema 'default' is always
	// assumed). This setting requires a restart of clusters and SQL warehouses
	// to take effect. Additionally, the default namespace only applies when
	// using Unity Catalog-enabled compute.
	Setting DefaultNamespaceSetting `tfsdk:"setting"`
}

// Details required to update a setting.
type UpdateEnhancedSecurityMonitoringSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask"`

	Setting EnhancedSecurityMonitoringSetting `tfsdk:"setting"`
}

// Details required to update a setting.
type UpdateEsmEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask"`

	Setting EsmEnablementAccountSetting `tfsdk:"setting"`
}

// Details required to update an IP access list.
type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-" url:"-"`

	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `tfsdk:"list_type"`
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask"`

	Setting PersonalComputeSetting `tfsdk:"setting"`
}

type UpdateResponse struct {
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask"`

	Setting RestrictWorkspaceAdminsSetting `tfsdk:"setting"`
}

type WorkspaceConf types.Map

func (tfSdkStruct *WorkspaceConf) ToGoSdk(goSdkStruct settings.WorkspaceConf) error {
	return nil
}

func (tfSdkStruct *WorkspaceConf) FromGoSdk(goSdkStruct settings.WorkspaceConf) error {
	return nil
}
