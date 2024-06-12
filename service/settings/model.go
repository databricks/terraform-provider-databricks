// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
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
	Etag string `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `tfsdk:"setting_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AutomaticClusterUpdateSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutomaticClusterUpdateSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterAutoRestartMessage struct {
	CanToggle bool `tfsdk:"can_toggle"`

	Enabled bool `tfsdk:"enabled"`
	// Contains an information about the enablement status judging (e.g. whether
	// the enterprise tier is enabled) This is only additional information that
	// MUST NOT be used to decide whether the setting is enabled or not. This is
	// intended to use only for purposes like showing an error message to the
	// customer with the additional details. For example, using these details we
	// can check why exactly the feature is disabled for this customer.
	EnablementDetails *ClusterAutoRestartMessageEnablementDetails `tfsdk:"enablement_details"`

	MaintenanceWindow *ClusterAutoRestartMessageMaintenanceWindow `tfsdk:"maintenance_window"`

	RestartEvenIfNoUpdatesAvailable bool `tfsdk:"restart_even_if_no_updates_available"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ClusterAutoRestartMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode bool `tfsdk:"forced_for_compliance_mode"`
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	UnavailableForDisabledEntitlement bool `tfsdk:"unavailable_for_disabled_entitlement"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier bool `tfsdk:"unavailable_for_non_enterprise_tier"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ClusterAutoRestartMessageEnablementDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessageEnablementDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Hours int `tfsdk:"hours"`

	Minutes int `tfsdk:"minutes"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// SHIELD feature: CSP
type ComplianceSecurityProfile struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	ComplianceStandards []ComplianceStandard `tfsdk:"compliance_standards"`

	IsEnabled bool `tfsdk:"is_enabled"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ComplianceSecurityProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplianceSecurityProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `tfsdk:"setting_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ComplianceSecurityProfileSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplianceSecurityProfileSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Compliance stardard for SHIELD customers
type ComplianceStandard string

const ComplianceStandardComplianceStandardUnspecified ComplianceStandard = `COMPLIANCE_STANDARD_UNSPECIFIED`

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
	case `COMPLIANCE_STANDARD_UNSPECIFIED`, `FEDRAMP_HIGH`, `FEDRAMP_IL5`, `FEDRAMP_MODERATE`, `HIPAA`, `IRAP_PROTECTED`, `ITAR_EAR`, `NONE`, `PCI_DSS`:
		*f = ComplianceStandard(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMPLIANCE_STANDARD_UNSPECIFIED", "FEDRAMP_HIGH", "FEDRAMP_IL5", "FEDRAMP_MODERATE", "HIPAA", "IRAP_PROTECTED", "ITAR_EAR", "NONE", "PCI_DSS"`, v)
	}
}

// Type always returns ComplianceStandard to satisfy [pflag.Value] interface
func (f *ComplianceStandard) Type() string {
	return "ComplianceStandard"
}

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {
	IpAddresses []string `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `tfsdk:"label"`
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
	Name string `tfsdk:"name"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region string `tfsdk:"region"`
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	ApplicationId string `tfsdk:"application_id"`
	// Comment that describes the purpose of the token.
	Comment string `tfsdk:"comment"`
	// The number of seconds before the token expires.
	LifetimeSeconds int64 `tfsdk:"lifetime_seconds"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateOboTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo `tfsdk:"token_info"`
	// Value of the token.
	TokenValue string `tfsdk:"token_value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateOboTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePrivateEndpointRuleRequest struct {
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId CreatePrivateEndpointRuleRequestGroupId `tfsdk:"group_id"`
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `tfsdk:"-" url:"-"`
	// The Azure resource ID of the target resource.
	ResourceId string `tfsdk:"resource_id"`
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
	Comment string `tfsdk:"comment"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds int64 `tfsdk:"lifetime_seconds"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo *PublicTokenInfo `tfsdk:"token_info"`
	// The value of the new token.
	TokenValue string `tfsdk:"token_value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards []ComplianceStandard `tfsdk:"compliance_standards"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced bool `tfsdk:"is_enforced"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CspEnablementAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CspEnablementAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `tfsdk:"setting_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CspEnablementAccountSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CspEnablementAccountSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"etag"`

	Namespace StringMessage `tfsdk:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `tfsdk:"setting_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DefaultNamespaceSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DefaultNamespaceSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `tfsdk:"-" url:"-"`
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"etag"`
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `tfsdk:"-" url:"-"`
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `tfsdk:"-" url:"-"`
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeletePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeletePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"etag"`
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `tfsdk:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `tfsdk:"-" url:"-"`
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"etag"`
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `tfsdk:"-" url:"-"`
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {
	IsEnabled bool `tfsdk:"is_enabled"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EnhancedSecurityMonitoring) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnhancedSecurityMonitoring) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `tfsdk:"setting_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EnhancedSecurityMonitoringSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnhancedSecurityMonitoringSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Account level policy for ESM
type EsmEnablementAccount struct {
	IsEnforced bool `tfsdk:"is_enforced"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EsmEnablementAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EsmEnablementAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `tfsdk:"setting_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EsmEnablementAccountSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EsmEnablementAccountSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken struct {
	// The requested token.
	Credential string `tfsdk:"credential"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime int64 `tfsdk:"credentialEolTime"`
	// User ID of the user that owns this token.
	OwnerId int64 `tfsdk:"ownerId"`
	// The scopes of access granted in the token.
	Scopes []string `tfsdk:"scopes"`
	// The type of this exchange token
	TokenType TokenType `tfsdk:"tokenType"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ExchangeToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	PartitionId PartitionId `tfsdk:"partitionId"`
	// Array of scopes for the token request.
	Scopes []string `tfsdk:"scopes"`
	// A list of token types being requested
	TokenType []TokenType `tfsdk:"tokenType"`
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {
	Values []ExchangeToken `tfsdk:"values"`
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `tfsdk:"ip_access_list"`
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `tfsdk:"-" url:"-"`
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetAutomaticClusterUpdateSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAutomaticClusterUpdateSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetComplianceSecurityProfileSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetComplianceSecurityProfileSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetCspEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCspEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetEnhancedSecurityMonitoringSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEnhancedSecurityMonitoringSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetEsmEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEsmEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `tfsdk:"-" url:"-"`
}

type GetIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `tfsdk:"ip_access_list"`
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {
	IpAccessLists []IpAccessListInfo `tfsdk:"ip_access_lists"`
}

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `tfsdk:"-" url:"-"`
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetPersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a private endpoint rule
type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `tfsdk:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `tfsdk:"-" url:"-"`
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
	Etag string `tfsdk:"-" url:"etag,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Check configuration status
type GetStatusRequest struct {
	Keys string `tfsdk:"-" url:"keys"`
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `tfsdk:"-" url:"-"`
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []TokenPermissionsDescription `tfsdk:"permission_levels"`
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {
	TokenInfo *TokenInfo `tfsdk:"token_info"`
}

// Definition of an IP Access list
type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	AddressCount int `tfsdk:"address_count"`
	// Creation timestamp in milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// User ID of the user who created this list.
	CreatedBy int64 `tfsdk:"created_by"`
	// Specifies whether this IP access list is enabled.
	Enabled bool `tfsdk:"enabled"`

	IpAddresses []string `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `tfsdk:"label"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId string `tfsdk:"list_id"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `tfsdk:"list_type"`
	// Update timestamp in milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// User ID of the user who updated this list.
	UpdatedBy int64 `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *IpAccessListInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IpAccessListInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists []IpAccessListInfo `tfsdk:"ip_access_lists"`
}

type ListNccAzurePrivateEndpointRulesResponse struct {
	Items []NccAzurePrivateEndpointRule `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListNccAzurePrivateEndpointRulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNccAzurePrivateEndpointRulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListNetworkConnectivityConfigurationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkConnectivityConfigurationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListNetworkConnectivityConfigurationsResponse struct {
	Items []NetworkConnectivityConfiguration `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListNetworkConnectivityConfigurationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkConnectivityConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `tfsdk:"-" url:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListPrivateEndpointRulesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPrivateEndpointRulesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos []PublicTokenInfo `tfsdk:"token_infos"`
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById int64 `tfsdk:"-" url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `tfsdk:"-" url:"created_by_username,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListTokenManagementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTokenManagementRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos []TokenInfo `tfsdk:"token_infos"`
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
	CidrBlocks []string `tfsdk:"cidr_blocks"`
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
	CreationTime int64 `tfsdk:"creation_time"`
	// Whether this private endpoint is deactivated.
	Deactivated bool `tfsdk:"deactivated"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt int64 `tfsdk:"deactivated_at"`
	// The name of the Azure private endpoint resource.
	EndpointName string `tfsdk:"endpoint_name"`
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId NccAzurePrivateEndpointRuleGroupId `tfsdk:"group_id"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId string `tfsdk:"network_connectivity_config_id"`
	// The Azure resource ID of the target resource.
	ResourceId string `tfsdk:"resource_id"`
	// The ID of a private endpoint rule.
	RuleId string `tfsdk:"rule_id"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `tfsdk:"updated_time"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *NccAzurePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NccAzurePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Subnets []string `tfsdk:"subnets"`
	// The Azure region in which this service endpoint rule applies.
	TargetRegion string `tfsdk:"target_region"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices []string `tfsdk:"target_services"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *NccAzureServiceEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NccAzureServiceEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRule `tfsdk:"azure_private_endpoint_rules"`
}

type NetworkConnectivityConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `tfsdk:"account_id"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `tfsdk:"creation_time"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig *NccEgressConfig `tfsdk:"egress_config"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name string `tfsdk:"name"`
	// Databricks network connectivity configuration ID.
	NetworkConnectivityConfigId string `tfsdk:"network_connectivity_config_id"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region string `tfsdk:"region"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `tfsdk:"updated_time"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *NetworkConnectivityConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkConnectivityConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId int64 `tfsdk:"workspaceId"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PartitionId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PartitionId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `tfsdk:"etag"`

	PersonalCompute PersonalComputeMessage `tfsdk:"personal_compute"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `tfsdk:"setting_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PersonalComputeSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PersonalComputeSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublicTokenInfo struct {
	// Comment the token was created with, if applicable.
	Comment string `tfsdk:"comment"`
	// Server time (in epoch milliseconds) when the token was created.
	CreationTime int64 `tfsdk:"creation_time"`
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	ExpiryTime int64 `tfsdk:"expiry_time"`
	// The ID of this token.
	TokenId string `tfsdk:"token_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PublicTokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublicTokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to replace an IP access list.
type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `tfsdk:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId string `tfsdk:"-" url:"-"`

	IpAddresses []string `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `tfsdk:"label"`
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
	Etag string `tfsdk:"etag"`

	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessage `tfsdk:"restrict_workspace_admins"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `tfsdk:"setting_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RestrictWorkspaceAdminsSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestrictWorkspaceAdminsSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId string `tfsdk:"token_id"`
}

type RevokeTokenResponse struct {
}

type SetStatusResponse struct {
}

type StringMessage struct {
	// Represents a generic string value.
	Value string `tfsdk:"value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *StringMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StringMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Permission level
	PermissionLevel TokenPermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TokenAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions []TokenPermission `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName string `tfsdk:"display_name"`
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TokenAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment string `tfsdk:"comment"`
	// User ID of the user that created the token.
	CreatedById int64 `tfsdk:"created_by_id"`
	// Username of the user that created the token.
	CreatedByUsername string `tfsdk:"created_by_username"`
	// Timestamp when the token was created.
	CreationTime int64 `tfsdk:"creation_time"`
	// Timestamp when the token expires.
	ExpiryTime int64 `tfsdk:"expiry_time"`
	// User ID of the user that owns the token.
	OwnerId int64 `tfsdk:"owner_id"`
	// ID of the token.
	TokenId string `tfsdk:"token_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermission struct {
	Inherited bool `tfsdk:"inherited"`

	InheritedFromObject []string `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel TokenPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TokenPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	AccessControlList []TokenAccessControlResponse `tfsdk:"access_control_list"`

	ObjectId string `tfsdk:"object_id"`

	ObjectType string `tfsdk:"object_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TokenPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsDescription struct {
	Description string `tfsdk:"description"`
	// Permission level
	PermissionLevel TokenPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TokenPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsRequest struct {
	AccessControlList []TokenAccessControlRequest `tfsdk:"access_control_list"`
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
	AllowMissing bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `tfsdk:"field_mask"`

	Setting AutomaticClusterUpdateSetting `tfsdk:"setting"`
}

// Details required to update a setting.
type UpdateComplianceSecurityProfileSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `tfsdk:"field_mask"`

	Setting ComplianceSecurityProfileSetting `tfsdk:"setting"`
}

// Details required to update a setting.
type UpdateCspEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `tfsdk:"field_mask"`

	Setting CspEnablementAccountSetting `tfsdk:"setting"`
}

// Details required to update a setting.
type UpdateDefaultNamespaceSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `tfsdk:"field_mask"`
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
	AllowMissing bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `tfsdk:"field_mask"`

	Setting EnhancedSecurityMonitoringSetting `tfsdk:"setting"`
}

// Details required to update a setting.
type UpdateEsmEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `tfsdk:"field_mask"`

	Setting EsmEnablementAccountSetting `tfsdk:"setting"`
}

// Details required to update an IP access list.
type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `tfsdk:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId string `tfsdk:"-" url:"-"`

	IpAddresses []string `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `tfsdk:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `tfsdk:"list_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateIpAccessList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateIpAccessList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `tfsdk:"field_mask"`

	Setting PersonalComputeSetting `tfsdk:"setting"`
}

type UpdateResponse struct {
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `tfsdk:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `tfsdk:"field_mask"`

	Setting RestrictWorkspaceAdminsSetting `tfsdk:"setting"`
}

type WorkspaceConf map[string]string
