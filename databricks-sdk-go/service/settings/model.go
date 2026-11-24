// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/iam"
)

type AccountIpAccessEnable struct {
	AcctIpAclEnable BooleanMessage `json:"acct_ip_acl_enable"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AccountIpAccessEnable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccountIpAccessEnable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AccountNetworkPolicy struct {
	// The associated account ID for this Network Policy object.
	AccountId string `json:"account_id,omitempty"`
	// The network policies applying for egress traffic.
	Egress *NetworkPolicyEgress `json:"egress,omitempty"`
	// The network policies applying for ingress traffic.
	Ingress *CustomerFacingIngressNetworkPolicy `json:"ingress,omitempty"`
	// The ingress policy for dry run mode. Dry run will always run even if the
	// request is allowed by the ingress policy. When this field is set, the
	// policy will be evaluated and emit logs only without blocking requests.
	IngressDryRun *CustomerFacingIngressNetworkPolicy `json:"ingress_dry_run,omitempty"`
	// The unique identifier for the network policy.
	NetworkPolicyId string `json:"network_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AccountNetworkPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccountNetworkPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AibiDashboardEmbeddingAccessPolicy struct {
	// Where AI/BI dashboards can be embedded: in any domain
	// (ALLOW_ALL_DOMAINS), only in the approved domains configured for the
	// workspace (ALLOW_APPROVED_DOMAINS), or nowhere (DENY_ALL_DOMAINS).
	// Defaults to ALLOW_APPROVED_DOMAINS.
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyType `json:"access_policy_type"`
}

func (s *AibiDashboardEmbeddingAccessPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AibiDashboardEmbeddingAccessPolicyAccessPolicyType string

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowAllDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyType = `ALLOW_ALL_DOMAINS`

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowApprovedDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyType = `ALLOW_APPROVED_DOMAINS`

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeDenyAllDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyType = `DENY_ALL_DOMAINS`

// String representation for [fmt.Print]
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Set(v string) error {
	switch v {
	case `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`:
		*f = AibiDashboardEmbeddingAccessPolicyAccessPolicyType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL_DOMAINS", "ALLOW_APPROVED_DOMAINS", "DENY_ALL_DOMAINS"`, v)
	}
}

// Values returns all possible values for AibiDashboardEmbeddingAccessPolicyAccessPolicyType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Values() []AibiDashboardEmbeddingAccessPolicyAccessPolicyType {
	return []AibiDashboardEmbeddingAccessPolicyAccessPolicyType{
		AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowAllDomains,
		AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowApprovedDomains,
		AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeDenyAllDomains,
	}
}

// Type always returns AibiDashboardEmbeddingAccessPolicyAccessPolicyType to satisfy [pflag.Value] interface
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Type() string {
	return "AibiDashboardEmbeddingAccessPolicyAccessPolicyType"
}

type AibiDashboardEmbeddingAccessPolicySetting struct {
	AibiDashboardEmbeddingAccessPolicy AibiDashboardEmbeddingAccessPolicy `json:"aibi_dashboard_embedding_access_policy"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AibiDashboardEmbeddingAccessPolicySetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AibiDashboardEmbeddingAccessPolicySetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AibiDashboardEmbeddingApprovedDomains struct {
	// Domains allowed to host embedded AI/BI dashboards. A leading wildcard
	// allows every subdomain of a domain. The list applies, and can only be
	// changed, while the embedding access policy is ALLOW_APPROVED_DOMAINS.
	ApprovedDomains []string `json:"approved_domains,omitempty"`
}

func (s *AibiDashboardEmbeddingApprovedDomains) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AibiDashboardEmbeddingApprovedDomainsSetting struct {
	AibiDashboardEmbeddingApprovedDomains AibiDashboardEmbeddingApprovedDomains `json:"aibi_dashboard_embedding_approved_domains"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AibiDashboardEmbeddingApprovedDomainsSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AibiDashboardEmbeddingApprovedDomainsSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AutomaticClusterUpdateSetting struct {
	AutomaticClusterUpdateWorkspace ClusterAutoRestartMessage `json:"automatic_cluster_update_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AutomaticClusterUpdateSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutomaticClusterUpdateSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BooleanMessage struct {
	Value bool `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BooleanMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BooleanMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterAutoRestartMessage struct {
	// Whether the workspace can change the enabled state. Databricks computes
	// this value: it is false when the compliance security profile forces the
	// feature on or when the workspace lacks the Enhanced Security and
	// Compliance entitlement. A value sent in an update request has no effect.
	CanToggle bool `json:"can_toggle,omitempty"`
	// Whether automatic cluster update is enabled for the workspace. The
	// effective value is always true while the compliance security profile is
	// enabled on the workspace.
	Enabled bool `json:"enabled,omitempty"`
	// Why the feature is forced on or unavailable. Databricks sets these
	// details for information only; they do not determine whether the feature
	// is enabled.
	EnablementDetails *ClusterAutoRestartMessageEnablementDetails `json:"enablement_details,omitempty"`
	// When the maintenance window runs. When unset, it defaults to the first
	// Sunday of every month at 01:00 UTC.
	MaintenanceWindow *ClusterAutoRestartMessageMaintenanceWindow `json:"maintenance_window,omitempty"`
	// Whether compute resources restart during every maintenance window even
	// when no update is available. When false, they restart only when an update
	// is available.
	RestartEvenIfNoUpdatesAvailable bool `json:"restart_even_if_no_updates_available,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	ForcedForComplianceMode bool `json:"forced_for_compliance_mode,omitempty"`
	// The feature is unavailable because the workspace does not have the
	// Enhanced Security and Compliance entitlement.
	UnavailableForDisabledEntitlement bool `json:"unavailable_for_disabled_entitlement,omitempty"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier bool `json:"unavailable_for_non_enterprise_tier,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAutoRestartMessageEnablementDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessageEnablementDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterAutoRestartMessageMaintenanceWindow struct {
	// Schedule anchored to a day of the week, either every week or on given
	// weeks of the month.
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule `json:"week_day_based_schedule,omitempty"`
}

func (s *ClusterAutoRestartMessageMaintenanceWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ClusterAutoRestartMessageMaintenanceWindowDayOfWeek string

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
	case `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`:
		*f = ClusterAutoRestartMessageMaintenanceWindowDayOfWeek(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY", "WEDNESDAY"`, v)
	}
}

// Values returns all possible values for ClusterAutoRestartMessageMaintenanceWindowDayOfWeek.
//
// There is no guarantee on the order of the values in the slice.
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Values() []ClusterAutoRestartMessageMaintenanceWindowDayOfWeek {
	return []ClusterAutoRestartMessageMaintenanceWindowDayOfWeek{
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekFriday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekMonday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSaturday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSunday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekThursday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekTuesday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekWednesday,
	}
}

// Type always returns ClusterAutoRestartMessageMaintenanceWindowDayOfWeek to satisfy [pflag.Value] interface
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Type() string {
	return "ClusterAutoRestartMessageMaintenanceWindowDayOfWeek"
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	// Day of the week on which the maintenance window starts.
	DayOfWeek ClusterAutoRestartMessageMaintenanceWindowDayOfWeek `json:"day_of_week,omitempty"`
	// How often the maintenance window occurs: every week, or only on the
	// listed occurrences of the day of the week within each month.
	Frequency ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency `json:"frequency,omitempty"`
	// Time of day, in UTC, at which the maintenance window starts.
	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime `json:"window_start_time,omitempty"`
}

func (s *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency string

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyEveryWeek ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `EVERY_WEEK`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstAndThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FIRST_AND_THIRD_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FIRST_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondAndFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `SECOND_AND_FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `SECOND_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `THIRD_OF_MONTH`

// String representation for [fmt.Print]
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Set(v string) error {
	switch v {
	case `EVERY_WEEK`, `FIRST_AND_THIRD_OF_MONTH`, `FIRST_OF_MONTH`, `FOURTH_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`:
		*f = ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EVERY_WEEK", "FIRST_AND_THIRD_OF_MONTH", "FIRST_OF_MONTH", "FOURTH_OF_MONTH", "SECOND_AND_FOURTH_OF_MONTH", "SECOND_OF_MONTH", "THIRD_OF_MONTH"`, v)
	}
}

// Values returns all possible values for ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency.
//
// There is no guarantee on the order of the values in the slice.
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Values() []ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency {
	return []ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency{
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyEveryWeek,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstAndThirdOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFourthOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondAndFourthOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyThirdOfMonth,
	}
}

// Type always returns ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency to satisfy [pflag.Value] interface
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Type() string {
	return "ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency"
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	// Hour of the day on a 24-hour clock.
	Hours int `json:"hours,omitempty"`
	// Minute of the hour.
	Minutes int `json:"minutes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// SHIELD feature: CSP Compliance Security Profile (CSP) enables enhanced
// compliance controls on the workspace.
type ComplianceSecurityProfile struct {
	// Compliance standards selected by the customer for this Compliance
	// Security Profile.
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`
	// Whether Compliance Security Profile (CSP) is enabled on the workspace.
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComplianceSecurityProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplianceSecurityProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ComplianceSecurityProfileSetting struct {
	ComplianceSecurityProfileWorkspace ComplianceSecurityProfile `json:"compliance_security_profile_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComplianceSecurityProfileSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplianceSecurityProfileSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Compliance standard for SHIELD customers. See README.md for how instructions
// of how to add new standards.
type ComplianceStandard string

const ComplianceStandardCanadaProtectedB ComplianceStandard = `CANADA_PROTECTED_B`

const ComplianceStandardCyberEssentialPlus ComplianceStandard = `CYBER_ESSENTIAL_PLUS`

const ComplianceStandardFedrampHigh ComplianceStandard = `FEDRAMP_HIGH`

const ComplianceStandardFedrampIl5 ComplianceStandard = `FEDRAMP_IL5`

const ComplianceStandardFedrampModerate ComplianceStandard = `FEDRAMP_MODERATE`

const ComplianceStandardGermanyC5 ComplianceStandard = `GERMANY_C5`

const ComplianceStandardGermanyTisax ComplianceStandard = `GERMANY_TISAX`

const ComplianceStandardHipaa ComplianceStandard = `HIPAA`

const ComplianceStandardHitrust ComplianceStandard = `HITRUST`

const ComplianceStandardIrapProtected ComplianceStandard = `IRAP_PROTECTED`

const ComplianceStandardIsmap ComplianceStandard = `ISMAP`

const ComplianceStandardItarEar ComplianceStandard = `ITAR_EAR`

const ComplianceStandardKsaEccCccDcc ComplianceStandard = `KSA_ECC_CCC_DCC`

const ComplianceStandardKFsi ComplianceStandard = `K_FSI`

const ComplianceStandardNone ComplianceStandard = `NONE`

const ComplianceStandardPciDss ComplianceStandard = `PCI_DSS`

// String representation for [fmt.Print]
func (f *ComplianceStandard) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComplianceStandard) Set(v string) error {
	switch v {
	case `CANADA_PROTECTED_B`, `CYBER_ESSENTIAL_PLUS`, `FEDRAMP_HIGH`, `FEDRAMP_IL5`, `FEDRAMP_MODERATE`, `GERMANY_C5`, `GERMANY_TISAX`, `HIPAA`, `HITRUST`, `IRAP_PROTECTED`, `ISMAP`, `ITAR_EAR`, `KSA_ECC_CCC_DCC`, `K_FSI`, `NONE`, `PCI_DSS`:
		*f = ComplianceStandard(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANADA_PROTECTED_B", "CYBER_ESSENTIAL_PLUS", "FEDRAMP_HIGH", "FEDRAMP_IL5", "FEDRAMP_MODERATE", "GERMANY_C5", "GERMANY_TISAX", "HIPAA", "HITRUST", "IRAP_PROTECTED", "ISMAP", "ITAR_EAR", "KSA_ECC_CCC_DCC", "K_FSI", "NONE", "PCI_DSS"`, v)
	}
}

// Values returns all possible values for ComplianceStandard.
//
// There is no guarantee on the order of the values in the slice.
func (f *ComplianceStandard) Values() []ComplianceStandard {
	return []ComplianceStandard{
		ComplianceStandardCanadaProtectedB,
		ComplianceStandardCyberEssentialPlus,
		ComplianceStandardFedrampHigh,
		ComplianceStandardFedrampIl5,
		ComplianceStandardFedrampModerate,
		ComplianceStandardGermanyC5,
		ComplianceStandardGermanyTisax,
		ComplianceStandardHipaa,
		ComplianceStandardHitrust,
		ComplianceStandardIrapProtected,
		ComplianceStandardIsmap,
		ComplianceStandardItarEar,
		ComplianceStandardKsaEccCccDcc,
		ComplianceStandardKFsi,
		ComplianceStandardNone,
		ComplianceStandardPciDss,
	}
}

// Type always returns ComplianceStandard to satisfy [pflag.Value] interface
func (f *ComplianceStandard) Type() string {
	return "ComplianceStandard"
}

// Holds the destination-specific configuration. Exactly one of the nested
// configs is set.
type Config struct {
	// Configures a Slack channel reached through the Databricks Genie app for
	// Slack.
	DatabricksAppSlack *DatabricksAppSlackConfig `json:"databricks_app_slack,omitempty"`
	// Configures a Microsoft Teams channel reached through the Databricks Genie
	// app for Microsoft Teams.
	DatabricksAppTeams *DatabricksAppTeamsConfig `json:"databricks_app_teams,omitempty"`
	// Configures an email destination.
	Email *EmailConfig `json:"email,omitempty"`
	// Configures a generic HTTPS webhook destination.
	GenericWebhook *GenericWebhookConfig `json:"generic_webhook,omitempty"`
	// Configures a Microsoft Teams incoming webhook destination.
	MicrosoftTeams *MicrosoftTeamsConfig `json:"microsoft_teams,omitempty"`
	// Configures a PagerDuty destination.
	Pagerduty *PagerdutyConfig `json:"pagerduty,omitempty"`
	// Configures a Slack incoming webhook destination.
	Slack *SlackConfig `json:"slack,omitempty"`
}

func (s *Config) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`

	ListType ListType `json:"list_type"`
}

func (s *CreateIpAccessList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func (s *CreateIpAccessListResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateNetworkConnectivityConfigRequest struct {
	NetworkConnectivityConfig CreateNetworkConnectivityConfiguration `json:"network_connectivity_config"`
}

func (s *CreateNetworkConnectivityConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Properties of the new network connectivity configuration.
type CreateNetworkConnectivityConfiguration struct {
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
	Name string `json:"name"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region string `json:"region"`
}

func (s *CreateNetworkConnectivityConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateNetworkPolicyRequest struct {
	// Network policy configuration details.
	NetworkPolicy AccountNetworkPolicy `json:"network_policy"`
}

func (s *CreateNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config *Config `json:"config,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	ApplicationId string `json:"application_id"`
	// Whether to enable autoscoping for this token. `disable_autoscoping`
	// supersedes this field and is the one read when both are set.
	AutoscopeEnabled bool `json:"autoscope_enabled,omitempty"`
	// Comment that describes the purpose of the token.
	Comment string `json:"comment,omitempty"`
	// The number of seconds before the token expires.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	Scopes []string `json:"scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateOboTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
	// Value of the token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateOboTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type CreatePrivateEndpointRule struct {
	// Only used by private endpoints to customer-managed private endpoint
	// services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string `json:"domain_names,omitempty"`
	// The full target AWS endpoint service name that connects to the
	// destination resources of the private endpoint.
	EndpointService string `json:"endpoint_service,omitempty"`

	ErrorMessage string `json:"error_message,omitempty"`

	GcpEndpoint *GcpEndpoint `json:"gcp_endpoint,omitempty"`
	// Not used by customer-managed private endpoint services.
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId string `json:"group_id,omitempty"`
	// The Azure resource ID of the target resource.
	ResourceId string `json:"resource_id,omitempty"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	ResourceNames []string `json:"resource_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePrivateEndpointRuleRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`

	PrivateEndpointRule CreatePrivateEndpointRule `json:"private_endpoint_rule"`
}

func (s *CreatePrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateTokenRequest struct {
	// Whether to enable autoscoping for this token. When true, the token will
	// automatically collect inferred API path scopes as it is used.
	// `disable_autoscoping` supersedes this field and is the one read when both
	// are set.
	AutoscopeEnabled bool `json:"autoscope_enabled,omitempty"`
	// Optional description to attach to the token.
	Comment string `json:"comment,omitempty"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid for 2 years.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`
	// Optional scopes of the token.
	Scopes []string `json:"scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo *PublicTokenInfo `json:"token_info,omitempty"`
	// The value of the new token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CspEnablementAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CspEnablementAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CspEnablementAccountSetting struct {
	CspEnablementAccount CspEnablementAccount `json:"csp_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CspEnablementAccountSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CspEnablementAccountSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network policies applying for ingress traffic.
type CustomerFacingIngressNetworkPolicy struct {
	// The network policy restrictions for cross-workspace access: which other
	// workspaces may access this workspace.
	CrossWorkspaceAccess *CustomerFacingIngressNetworkPolicyCrossWorkspaceAccess `json:"cross_workspace_access,omitempty"`
	// The network policy restrictions for private access. Configures how
	// requests arriving over private connectivity are governed.
	PrivateAccess *CustomerFacingIngressNetworkPolicyPrivateAccess `json:"private_access,omitempty"`
	// The network policy restrictions for public access to the workspace.
	// Configures how public internet traffic is allowed or denied access.
	PublicAccess *CustomerFacingIngressNetworkPolicyPublicAccess `json:"public_access,omitempty"`
}

func (s *CustomerFacingIngressNetworkPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Matches account-level Databricks API endpoints for an ingress network policy
// rule.
type CustomerFacingIngressNetworkPolicyAccountApiDestination struct {
	// Qualifies the breadth of API access for the listed scopes. See
	// ApiScopeQualifier.
	ScopeQualifier CustomerFacingIngressNetworkPolicyApiScopeQualifier `json:"scope_qualifier,omitempty"`
	// The API scopes to match. Use "all-apis" to match any account-level API.
	Scopes []string `json:"scopes,omitempty"`
}

func (s *CustomerFacingIngressNetworkPolicyAccountApiDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CustomerFacingIngressNetworkPolicyAccountDatabricksOneDestination struct {
	// Must be set to true.
	AllDestinations bool `json:"all_destinations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyAccountDatabricksOneDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyAccountDatabricksOneDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The account console UI destination.
type CustomerFacingIngressNetworkPolicyAccountUiDestination struct {
	// Must be set to true.
	AllDestinations bool `json:"all_destinations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyAccountUiDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyAccountUiDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Qualifies the breadth of API access permitted by an ingress network policy
// rule. API_SCOPE_QUALIFIER_READ narrows matching to read-only variants of the
// listed scopes; API_SCOPE_QUALIFIER_ALL matches any scope. When unset, scopes
// match exactly as listed.
type CustomerFacingIngressNetworkPolicyApiScopeQualifier string

const CustomerFacingIngressNetworkPolicyApiScopeQualifierApiScopeQualifierAll CustomerFacingIngressNetworkPolicyApiScopeQualifier = `API_SCOPE_QUALIFIER_ALL`

const CustomerFacingIngressNetworkPolicyApiScopeQualifierApiScopeQualifierRead CustomerFacingIngressNetworkPolicyApiScopeQualifier = `API_SCOPE_QUALIFIER_READ`

// String representation for [fmt.Print]
func (f *CustomerFacingIngressNetworkPolicyApiScopeQualifier) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingIngressNetworkPolicyApiScopeQualifier) Set(v string) error {
	switch v {
	case `API_SCOPE_QUALIFIER_ALL`, `API_SCOPE_QUALIFIER_READ`:
		*f = CustomerFacingIngressNetworkPolicyApiScopeQualifier(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API_SCOPE_QUALIFIER_ALL", "API_SCOPE_QUALIFIER_READ"`, v)
	}
}

// Values returns all possible values for CustomerFacingIngressNetworkPolicyApiScopeQualifier.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingIngressNetworkPolicyApiScopeQualifier) Values() []CustomerFacingIngressNetworkPolicyApiScopeQualifier {
	return []CustomerFacingIngressNetworkPolicyApiScopeQualifier{
		CustomerFacingIngressNetworkPolicyApiScopeQualifierApiScopeQualifierAll,
		CustomerFacingIngressNetworkPolicyApiScopeQualifierApiScopeQualifierRead,
	}
}

// Type always returns CustomerFacingIngressNetworkPolicyApiScopeQualifier to satisfy [pflag.Value] interface
func (f *CustomerFacingIngressNetworkPolicyApiScopeQualifier) Type() string {
	return "CustomerFacingIngressNetworkPolicyApiScopeQualifier"
}

type CustomerFacingIngressNetworkPolicyAppsRuntimeDestination struct {
	// Must be set to true.
	AllDestinations bool `json:"all_destinations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyAppsRuntimeDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyAppsRuntimeDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerFacingIngressNetworkPolicyAuthentication struct {
	// Valid only when IdentityType is IDENTITY_TYPE_SELECTED_IDENTITIES.
	Identities []CustomerFacingIngressNetworkPolicyAuthenticationIdentity `json:"identities,omitempty"`

	IdentityType CustomerFacingIngressNetworkPolicyAuthenticationIdentityType `json:"identity_type,omitempty"`
}

func (s *CustomerFacingIngressNetworkPolicyAuthentication) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CustomerFacingIngressNetworkPolicyAuthenticationIdentity struct {
	PrincipalId int64 `json:"principal_id,omitempty"`

	PrincipalType CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType `json:"principal_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyAuthenticationIdentity) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyAuthenticationIdentity) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType string

const CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalTypePrincipalTypeServicePrincipal CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType = `PRINCIPAL_TYPE_SERVICE_PRINCIPAL`

const CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalTypePrincipalTypeUser CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType = `PRINCIPAL_TYPE_USER`

// String representation for [fmt.Print]
func (f *CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType) Set(v string) error {
	switch v {
	case `PRINCIPAL_TYPE_SERVICE_PRINCIPAL`, `PRINCIPAL_TYPE_USER`:
		*f = CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PRINCIPAL_TYPE_SERVICE_PRINCIPAL", "PRINCIPAL_TYPE_USER"`, v)
	}
}

// Values returns all possible values for CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType) Values() []CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType {
	return []CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType{
		CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalTypePrincipalTypeServicePrincipal,
		CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalTypePrincipalTypeUser,
	}
}

// Type always returns CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType to satisfy [pflag.Value] interface
func (f *CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType) Type() string {
	return "CustomerFacingIngressNetworkPolicyAuthenticationIdentityPrincipalType"
}

type CustomerFacingIngressNetworkPolicyAuthenticationIdentityType string

const CustomerFacingIngressNetworkPolicyAuthenticationIdentityTypeIdentityTypeAllServicePrincipals CustomerFacingIngressNetworkPolicyAuthenticationIdentityType = `IDENTITY_TYPE_ALL_SERVICE_PRINCIPALS`

const CustomerFacingIngressNetworkPolicyAuthenticationIdentityTypeIdentityTypeAllUsers CustomerFacingIngressNetworkPolicyAuthenticationIdentityType = `IDENTITY_TYPE_ALL_USERS`

const CustomerFacingIngressNetworkPolicyAuthenticationIdentityTypeIdentityTypeSelectedIdentities CustomerFacingIngressNetworkPolicyAuthenticationIdentityType = `IDENTITY_TYPE_SELECTED_IDENTITIES`

// String representation for [fmt.Print]
func (f *CustomerFacingIngressNetworkPolicyAuthenticationIdentityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingIngressNetworkPolicyAuthenticationIdentityType) Set(v string) error {
	switch v {
	case `IDENTITY_TYPE_ALL_SERVICE_PRINCIPALS`, `IDENTITY_TYPE_ALL_USERS`, `IDENTITY_TYPE_SELECTED_IDENTITIES`:
		*f = CustomerFacingIngressNetworkPolicyAuthenticationIdentityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IDENTITY_TYPE_ALL_SERVICE_PRINCIPALS", "IDENTITY_TYPE_ALL_USERS", "IDENTITY_TYPE_SELECTED_IDENTITIES"`, v)
	}
}

// Values returns all possible values for CustomerFacingIngressNetworkPolicyAuthenticationIdentityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingIngressNetworkPolicyAuthenticationIdentityType) Values() []CustomerFacingIngressNetworkPolicyAuthenticationIdentityType {
	return []CustomerFacingIngressNetworkPolicyAuthenticationIdentityType{
		CustomerFacingIngressNetworkPolicyAuthenticationIdentityTypeIdentityTypeAllServicePrincipals,
		CustomerFacingIngressNetworkPolicyAuthenticationIdentityTypeIdentityTypeAllUsers,
		CustomerFacingIngressNetworkPolicyAuthenticationIdentityTypeIdentityTypeSelectedIdentities,
	}
}

// Type always returns CustomerFacingIngressNetworkPolicyAuthenticationIdentityType to satisfy [pflag.Value] interface
func (f *CustomerFacingIngressNetworkPolicyAuthenticationIdentityType) Type() string {
	return "CustomerFacingIngressNetworkPolicyAuthenticationIdentityType"
}

type CustomerFacingIngressNetworkPolicyCrossWorkspaceAccess struct {
	// Allow rules are evaluated after deny rules. A request matching any allow
	// rule is allowed; a request matching no rule is denied by default. Only
	// applies when restriction_mode is RESTRICTED_ACCESS.
	AllowRules []CustomerFacingIngressNetworkPolicyCrossWorkspaceIngressRule `json:"allow_rules,omitempty"`
	// Deny rules are evaluated first. A request matching any deny rule is
	// denied, regardless of allow rules. Only applies when restriction_mode is
	// RESTRICTED_ACCESS.
	DenyRules []CustomerFacingIngressNetworkPolicyCrossWorkspaceIngressRule `json:"deny_rules,omitempty"`
	// The restriction mode for cross-workspace access.
	RestrictionMode CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode `json:"restriction_mode"`
}

func (s *CustomerFacingIngressNetworkPolicyCrossWorkspaceAccess) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The restriction mode for cross-workspace access. In FULL_ACCESS mode,
// requests from any source workspace (in any account) are allowed, and deny
// rules and allow rules cannot be set. In RESTRICTED_ACCESS mode, access is
// restricted based on deny rules and allow rules; requests that do not match
// any allow rule are denied. In LEGACY_MODE, cross-workspace ingress is not
// governed by this policy.
type CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode string

const CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionModeFullAccess CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode = `FULL_ACCESS`

const CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionModeLegacyMode CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode = `LEGACY_MODE`

const CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionModeRestrictedAccess CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode = `RESTRICTED_ACCESS`

// String representation for [fmt.Print]
func (f *CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode) Set(v string) error {
	switch v {
	case `FULL_ACCESS`, `LEGACY_MODE`, `RESTRICTED_ACCESS`:
		*f = CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL_ACCESS", "LEGACY_MODE", "RESTRICTED_ACCESS"`, v)
	}
}

// Values returns all possible values for CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode) Values() []CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode {
	return []CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode{
		CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionModeFullAccess,
		CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionModeLegacyMode,
		CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionModeRestrictedAccess,
	}
}

// Type always returns CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode to satisfy [pflag.Value] interface
func (f *CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode) Type() string {
	return "CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode"
}

// An ingress rule is enforced when a request satisfies all specified attributes
// — including request origin, destination, and authentication.
type CustomerFacingIngressNetworkPolicyCrossWorkspaceIngressRule struct {
	// The authenticated identity the request must match. When unset, the rule
	// matches all users and service principals.
	Authentication *CustomerFacingIngressNetworkPolicyAuthentication `json:"authentication,omitempty"`
	// The destination the request must match — the resource being accessed,
	// for example the workspace UI or workspace APIs. See RequestDestination.
	Destination *CustomerFacingIngressNetworkPolicyRequestDestination `json:"destination,omitempty"`
	// The label for this ingress rule.
	Label string `json:"label,omitempty"`
	// The origin the request must match — the source workspace the request
	// comes from, either specific source workspaces or any source workspace in
	// any account. See CrossWorkspaceRequestOrigin.
	Origin *CustomerFacingIngressNetworkPolicyCrossWorkspaceRequestOrigin `json:"origin,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyCrossWorkspaceIngressRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyCrossWorkspaceIngressRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerFacingIngressNetworkPolicyCrossWorkspaceRequestOrigin struct {
	// Matches all source workspaces.
	AllSourceWorkspaces bool `json:"all_source_workspaces,omitempty"`
	// Specific source workspace IDs to match.
	SelectedWorkspaces *CustomerFacingIngressNetworkPolicyWorkspaceIdList `json:"selected_workspaces,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyCrossWorkspaceRequestOrigin) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyCrossWorkspaceRequestOrigin) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A set of registered endpoints, identified by their endpoint IDs.
type CustomerFacingIngressNetworkPolicyEndpoints struct {
	// The IDs of the registered endpoints. Must contain at least one endpoint
	// ID.
	EndpointIds []string `json:"endpoint_ids,omitempty"`
}

func (s *CustomerFacingIngressNetworkPolicyEndpoints) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CustomerFacingIngressNetworkPolicyIpRanges struct {
	// We only support IPv4 and IPv4 CIDR notation for now.
	IpRanges []string `json:"ip_ranges,omitempty"`
}

func (s *CustomerFacingIngressNetworkPolicyIpRanges) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CustomerFacingIngressNetworkPolicyLakebaseRuntimeDestination struct {
	// Must be set to true.
	AllDestinations bool `json:"all_destinations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyLakebaseRuntimeDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyLakebaseRuntimeDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configures how requests arriving over private connectivity, such as
// registered endpoints, are allowed or denied access.
type CustomerFacingIngressNetworkPolicyPrivateAccess struct {
	// Allow rules are evaluated after deny rules. A request matching any allow
	// rule is allowed; a request matching no rule is denied by default. Only
	// applies when restriction_mode is RESTRICTED_ACCESS.
	AllowRules []CustomerFacingIngressNetworkPolicyPrivateIngressRule `json:"allow_rules,omitempty"`
	// Deny rules are evaluated first. A request matching any deny rule is
	// denied, regardless of allow rules. Only applies when restriction_mode is
	// RESTRICTED_ACCESS.
	DenyRules []CustomerFacingIngressNetworkPolicyPrivateIngressRule `json:"deny_rules,omitempty"`
	// The restriction mode for private access.
	RestrictionMode CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode `json:"restriction_mode"`
}

func (s *CustomerFacingIngressNetworkPolicyPrivateAccess) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The restriction mode for private access. In ALLOW_ALL_REGISTERED_ENDPOINTS
// mode, requests arriving through any endpoint registered to the account are
// allowed, and deny rules and allow rules cannot be set. In RESTRICTED_ACCESS
// mode, access is restricted based on deny rules and allow rules; requests that
// do not match any allow rule are denied.
type CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode string

const CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionModeAllowAllRegisteredEndpoints CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode = `ALLOW_ALL_REGISTERED_ENDPOINTS`

const CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionModeRestrictedAccess CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode = `RESTRICTED_ACCESS`

// String representation for [fmt.Print]
func (f *CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode) Set(v string) error {
	switch v {
	case `ALLOW_ALL_REGISTERED_ENDPOINTS`, `RESTRICTED_ACCESS`:
		*f = CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL_REGISTERED_ENDPOINTS", "RESTRICTED_ACCESS"`, v)
	}
}

// Values returns all possible values for CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode) Values() []CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode {
	return []CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode{
		CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionModeAllowAllRegisteredEndpoints,
		CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionModeRestrictedAccess,
	}
}

// Type always returns CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode to satisfy [pflag.Value] interface
func (f *CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode) Type() string {
	return "CustomerFacingIngressNetworkPolicyPrivateAccessRestrictionMode"
}

// An ingress rule is enforced when a request satisfies all specified attributes
// — including request origin, destination, and authentication.
type CustomerFacingIngressNetworkPolicyPrivateIngressRule struct {
	// The authenticated identity the request must match. When unset, the rule
	// matches all users and service principals.
	Authentication *CustomerFacingIngressNetworkPolicyAuthentication `json:"authentication,omitempty"`
	// The destination the request must match — the resource being accessed,
	// for example the workspace UI, workspace APIs, or account-level APIs. See
	// RequestDestination.
	Destination *CustomerFacingIngressNetworkPolicyRequestDestination `json:"destination,omitempty"`
	// The label for this ingress rule.
	Label string `json:"label,omitempty"`
	// The origin the request must match — the private connectivity the
	// request arrives through, for example a specific set of registered
	// endpoints or any endpoint registered to the account. See
	// PrivateRequestOrigin.
	Origin *CustomerFacingIngressNetworkPolicyPrivateRequestOrigin `json:"origin,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyPrivateIngressRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyPrivateIngressRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The origin of a private access request, identified by the endpoint through
// which the request arrives.
type CustomerFacingIngressNetworkPolicyPrivateRequestOrigin struct {
	// Matches requests arriving over any private connectivity, including
	// registered endpoints and the workspace's Azure Private Link (ui-api)
	// endpoints. Can only be used in deny rules of workspace-level network
	// policies. Must be set to true when specified.
	AllPrivateAccess bool `json:"all_private_access,omitempty"`
	// Matches requests arriving through any endpoint registered to the account.
	// Must be set to true when specified.
	AllRegisteredEndpoints bool `json:"all_registered_endpoints,omitempty"`
	// Matches requests arriving through the workspace's Azure Private Link
	// (ui-api) endpoints. Can only be used in deny rules of workspace-level
	// network policies. Must be set to true when specified.
	AzureWorkspacePrivateLink bool `json:"azure_workspace_private_link,omitempty"`
	// Matches requests arriving through any of the specified registered
	// endpoints.
	Endpoints *CustomerFacingIngressNetworkPolicyEndpoints `json:"endpoints,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyPrivateRequestOrigin) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyPrivateRequestOrigin) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerFacingIngressNetworkPolicyPublicAccess struct {
	AllowRules []CustomerFacingIngressNetworkPolicyPublicIngressRule `json:"allow_rules,omitempty"`

	DenyRules []CustomerFacingIngressNetworkPolicyPublicIngressRule `json:"deny_rules,omitempty"`

	RestrictionMode CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode `json:"restriction_mode"`
}

func (s *CustomerFacingIngressNetworkPolicyPublicAccess) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode string

const CustomerFacingIngressNetworkPolicyPublicAccessRestrictionModeFullAccess CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode = `FULL_ACCESS`

const CustomerFacingIngressNetworkPolicyPublicAccessRestrictionModeRestrictedAccess CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode = `RESTRICTED_ACCESS`

// String representation for [fmt.Print]
func (f *CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode) Set(v string) error {
	switch v {
	case `FULL_ACCESS`, `RESTRICTED_ACCESS`:
		*f = CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL_ACCESS", "RESTRICTED_ACCESS"`, v)
	}
}

// Values returns all possible values for CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode) Values() []CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode {
	return []CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode{
		CustomerFacingIngressNetworkPolicyPublicAccessRestrictionModeFullAccess,
		CustomerFacingIngressNetworkPolicyPublicAccessRestrictionModeRestrictedAccess,
	}
}

// Type always returns CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode to satisfy [pflag.Value] interface
func (f *CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode) Type() string {
	return "CustomerFacingIngressNetworkPolicyPublicAccessRestrictionMode"
}

// An ingress rule is enforced when a request satisfies all specified attributes
// — including request origin, destination, and authentication.
type CustomerFacingIngressNetworkPolicyPublicIngressRule struct {
	Authentication *CustomerFacingIngressNetworkPolicyAuthentication `json:"authentication,omitempty"`

	Destination *CustomerFacingIngressNetworkPolicyRequestDestination `json:"destination,omitempty"`
	// The label for this ingress rule.
	Label string `json:"label,omitempty"`

	Origin *CustomerFacingIngressNetworkPolicyPublicRequestOrigin `json:"origin,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyPublicIngressRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyPublicIngressRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerFacingIngressNetworkPolicyPublicRequestOrigin struct {
	// Matches all IPv4 and IPv6 ranges (both public and private).
	AllIpRanges bool `json:"all_ip_ranges,omitempty"`
	// Excluded means: all public IP ranges except this one.
	ExcludedIpRanges *CustomerFacingIngressNetworkPolicyIpRanges `json:"excluded_ip_ranges,omitempty"`
	// Will not allow IP ranges with private IPs.
	IncludedIpRanges *CustomerFacingIngressNetworkPolicyIpRanges `json:"included_ip_ranges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyPublicRequestOrigin) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyPublicRequestOrigin) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerFacingIngressNetworkPolicyRequestDestination struct {
	// Matches requests to account-level APIs. Can only be used in the
	// account-level network policy.
	AccountApi *CustomerFacingIngressNetworkPolicyAccountApiDestination `json:"account_api,omitempty"`
	// Account DatabricksOne destination is not supported.
	AccountDatabricksOne *CustomerFacingIngressNetworkPolicyAccountDatabricksOneDestination `json:"account_databricks_one,omitempty"`
	// Matches requests to the account console UI. Can only be used in the
	// account-level network policy.
	AccountUi *CustomerFacingIngressNetworkPolicyAccountUiDestination `json:"account_ui,omitempty"`
	// When true, match all destinations, no other destination fields can be
	// set. When not set or false, at least one specific destination must be
	// provided.
	AllDestinations bool `json:"all_destinations,omitempty"`

	AppsRuntime *CustomerFacingIngressNetworkPolicyAppsRuntimeDestination `json:"apps_runtime,omitempty"`

	LakebaseRuntime *CustomerFacingIngressNetworkPolicyLakebaseRuntimeDestination `json:"lakebase_runtime,omitempty"`

	WorkspaceApi *CustomerFacingIngressNetworkPolicyWorkspaceApiDestination `json:"workspace_api,omitempty"`

	WorkspaceUi *CustomerFacingIngressNetworkPolicyWorkspaceUiDestination `json:"workspace_ui,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyRequestDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyRequestDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Matches workspace-level Databricks API endpoints for an ingress network
// policy rule.
type CustomerFacingIngressNetworkPolicyWorkspaceApiDestination struct {
	// Inverse of `scopes`: matches every API scope EXCEPT those listed here
	// ("allow all except"). Mutually exclusive with `scopes` — a single
	// destination may set at most one of the two.
	ExcludedScopes []string `json:"excluded_scopes,omitempty"`
	// Qualifies the breadth of API access for the listed scopes. See
	// ApiScopeQualifier.
	ScopeQualifier CustomerFacingIngressNetworkPolicyApiScopeQualifier `json:"scope_qualifier,omitempty"`

	Scopes []string `json:"scopes,omitempty"`
}

func (s *CustomerFacingIngressNetworkPolicyWorkspaceApiDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CustomerFacingIngressNetworkPolicyWorkspaceIdList struct {
	WorkspaceIds []int64 `json:"workspace_ids,omitempty"`
}

func (s *CustomerFacingIngressNetworkPolicyWorkspaceIdList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CustomerFacingIngressNetworkPolicyWorkspaceUiDestination struct {
	// Must be set to true.
	AllDestinations bool `json:"all_destinations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingIngressNetworkPolicyWorkspaceUiDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingIngressNetworkPolicyWorkspaceUiDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Properties of the new private endpoint rule. Note that for private endpoints
// towards a VPC endpoint service behind a customer-managed NLB, you must
// approve the endpoint in AWS console after initialization.
type CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule struct {
	// Databricks account ID. You can find your account ID from the Accounts
	// Console.
	AccountId string `json:"account_id,omitempty"`
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is ESTABLISHED. Remember that
	// you must approve new endpoints on your resources in the AWS console
	// before they take effect. The possible values are: - PENDING: The endpoint
	// has been created and pending approval. - ESTABLISHED: The endpoint has
	// been approved and is ready to use in your serverless compute resources. -
	// REJECTED: Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up. - EXPIRED: If the endpoint is created but not approved in 14
	// days, it is EXPIRED.
	ConnectionState CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState `json:"connection_state,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Whether this private endpoint is deactivated.
	Deactivated bool `json:"deactivated,omitempty"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt int64 `json:"deactivated_at,omitempty"`
	// Only used by private endpoints towards a VPC endpoint service for
	// customer-managed VPC endpoint service.
	//
	// The target AWS resource FQDNs accessible via the VPC endpoint service.
	// When updating this field, we perform full update on this field. Please
	// ensure a full list of desired domain_names is provided.
	DomainNames []string `json:"domain_names,omitempty"`
	// Only used by private endpoints towards an AWS S3 service.
	//
	// Update this field to activate/deactivate this private endpoint to allow
	// egress access from serverless compute resources.
	Enabled bool `json:"enabled,omitempty"`
	// The full target AWS endpoint service name that connects to the
	// destination resources of the private endpoint.
	EndpointService string `json:"endpoint_service,omitempty"`

	ErrorMessage string `json:"error_message,omitempty"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	ResourceNames []string `json:"resource_names,omitempty"`
	// The ID of a private endpoint rule.
	RuleId string `json:"rule_id,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`
	// The AWS VPC endpoint ID. You can use this ID to identify VPC endpoint
	// created by Databricks.
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState string

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateCreateFailed CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `CREATE_FAILED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateCreating CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `CREATING`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateDisconnected CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `DISCONNECTED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateEstablished CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `ESTABLISHED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateExpired CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `EXPIRED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePending CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `PENDING`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateRejected CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `REJECTED`

// String representation for [fmt.Print]
func (f *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState) Set(v string) error {
	switch v {
	case `CREATE_FAILED`, `CREATING`, `DISCONNECTED`, `ESTABLISHED`, `EXPIRED`, `PENDING`, `REJECTED`:
		*f = CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATE_FAILED", "CREATING", "DISCONNECTED", "ESTABLISHED", "EXPIRED", "PENDING", "REJECTED"`, v)
	}
}

// Values returns all possible values for CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState) Values() []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState {
	return []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState{
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateCreateFailed,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateCreating,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateDisconnected,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateEstablished,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateExpired,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePending,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateRejected,
	}
}

// Type always returns CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState to satisfy [pflag.Value] interface
func (f *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState) Type() string {
	return "CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState"
}

type DashboardEmailSubscriptions struct {
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardEmailSubscriptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardEmailSubscriptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Slack channel reached through the Databricks app. The app must already be
// installed in the Slack workspace and connected to this Databricks workspace.
type DatabricksAppSlackConfig struct {
	// [Input-Only] Slack channel ID notifications are posted to.
	ChannelId string `json:"channel_id,omitempty"`
	// [Output-Only] Whether channel ID is set.
	ChannelIdSet bool `json:"channel_id_set,omitempty"`
	// [Input-Only] Slack workspace (team) ID the Databricks app is installed
	// in.
	TeamId string `json:"team_id,omitempty"`
	// [Output-Only] Whether team ID is set.
	TeamIdSet bool `json:"team_id_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksAppSlackConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksAppSlackConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Microsoft Teams channel reached through the Databricks app. The app must
// already be installed in the Microsoft 365 tenant and connected to this
// Databricks workspace.
type DatabricksAppTeamsConfig struct {
	// [Input-Only] Teams channel ID notifications are posted to.
	ChannelId string `json:"channel_id,omitempty"`
	// [Output-Only] Whether channel ID is set.
	ChannelIdSet bool `json:"channel_id_set,omitempty"`
	// [Input-Only] Microsoft Entra group ID of the team that owns the channel.
	TeamId string `json:"team_id,omitempty"`
	// [Output-Only] Whether team ID is set.
	TeamIdSet bool `json:"team_id_set,omitempty"`
	// [Input-Only] Microsoft Entra tenant ID of the Microsoft 365 tenant the
	// app is installed in.
	TenantId string `json:"tenant_id,omitempty"`
	// [Output-Only] Whether tenant ID is set.
	TenantIdSet bool `json:"tenant_id_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksAppTeamsConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksAppTeamsConfig) MarshalJSON() ([]byte, error) {
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
	Etag string `json:"etag,omitempty"`

	Namespace StringMessage `json:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DefaultNamespaceSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DefaultNamespaceSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DefaultWarehouseId struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	StringVal StringMessage `json:"string_val"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DefaultWarehouseId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DefaultWarehouseId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteAccountIpAccessEnableResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func (s *DeleteAccountIpAccessEnableResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

func (s *DeleteAccountIpAccessListRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteAibiDashboardEmbeddingAccessPolicySettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func (s *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func (s *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDashboardEmailSubscriptionsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDashboardEmailSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDashboardEmailSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteDashboardEmailSubscriptionsResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func (s *DeleteDashboardEmailSubscriptionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	Etag string `json:"etag"`
}

func (s *DeleteDefaultNamespaceSettingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDefaultWarehouseIdRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDefaultWarehouseIdRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDefaultWarehouseIdRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteDefaultWarehouseIdResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func (s *DeleteDefaultWarehouseIdResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `json:"etag"`
}

func (s *DeleteDisableLegacyAccessResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteDisableLegacyDbfsResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func (s *DeleteDisableLegacyDbfsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Etag string `json:"etag"`
}

func (s *DeleteDisableLegacyFeaturesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

func (s *DeleteIpAccessListRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteLlmProxyPartnerPoweredWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteLlmProxyPartnerPoweredWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteLlmProxyPartnerPoweredWorkspaceResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func (s *DeleteLlmProxyPartnerPoweredWorkspaceResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

func (s *DeleteNetworkConnectivityConfigurationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteNetworkPolicyRequest struct {
	// The unique identifier of the network policy to delete.
	NetworkPolicyId string `json:"-" url:"-"`
}

func (s *DeleteNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteNotificationDestinationRequest struct {
	// A UUID identifying the notification destination.
	Id string `json:"-" url:"-"`
}

func (s *DeleteNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeletePersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	Etag string `json:"etag"`
}

func (s *DeletePersonalComputeSettingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
}

func (s *DeletePrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	Etag string `json:"etag"`
}

func (s *DeleteRestrictWorkspaceAdminsSettingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteSqlResultsDownloadRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteSqlResultsDownloadRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSqlResultsDownloadRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteSqlResultsDownloadResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func (s *DeleteSqlResultsDownloadResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteTokenManagementRequest struct {
	// The ID of the token to revoke.
	TokenId string `json:"-" url:"-"`
}

func (s *DeleteTokenManagementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The kind of destination that notifications are sent to.
type DestinationType string

const DestinationTypeDatabricksAppSlack DestinationType = `DATABRICKS_APP_SLACK`

const DestinationTypeDatabricksAppTeams DestinationType = `DATABRICKS_APP_TEAMS`

const DestinationTypeEmail DestinationType = `EMAIL`

const DestinationTypeMicrosoftTeams DestinationType = `MICROSOFT_TEAMS`

const DestinationTypePagerduty DestinationType = `PAGERDUTY`

const DestinationTypeSlack DestinationType = `SLACK`

const DestinationTypeWebhook DestinationType = `WEBHOOK`

// String representation for [fmt.Print]
func (f *DestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DestinationType) Set(v string) error {
	switch v {
	case `DATABRICKS_APP_SLACK`, `DATABRICKS_APP_TEAMS`, `EMAIL`, `MICROSOFT_TEAMS`, `PAGERDUTY`, `SLACK`, `WEBHOOK`:
		*f = DestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS_APP_SLACK", "DATABRICKS_APP_TEAMS", "EMAIL", "MICROSOFT_TEAMS", "PAGERDUTY", "SLACK", "WEBHOOK"`, v)
	}
}

// Values returns all possible values for DestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DestinationType) Values() []DestinationType {
	return []DestinationType{
		DestinationTypeDatabricksAppSlack,
		DestinationTypeDatabricksAppTeams,
		DestinationTypeEmail,
		DestinationTypeMicrosoftTeams,
		DestinationTypePagerduty,
		DestinationTypeSlack,
		DestinationTypeWebhook,
	}
}

// Type always returns DestinationType to satisfy [pflag.Value] interface
func (f *DestinationType) Type() string {
	return "DestinationType"
}

type DisableLegacyAccess struct {
	DisableLegacyAccess BooleanMessage `json:"disable_legacy_access"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DisableLegacyAccess) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DisableLegacyAccess) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DisableLegacyDbfs struct {
	DisableLegacyDbfs BooleanMessage `json:"disable_legacy_dbfs"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DisableLegacyDbfs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DisableLegacyDbfs) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DisableLegacyFeatures struct {
	DisableLegacyFeatures BooleanMessage `json:"disable_legacy_features"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DisableLegacyFeatures) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DisableLegacyFeatures) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto).
type EgressNetworkPolicy struct {
	// The access policy enforced for egress traffic to the internet.
	InternetAccess *EgressNetworkPolicyInternetAccessPolicy `json:"internet_access,omitempty"`
}

func (s *EgressNetworkPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EgressNetworkPolicyInternetAccessPolicy struct {
	AllowedInternetDestinations []EgressNetworkPolicyInternetAccessPolicyInternetDestination `json:"allowed_internet_destinations,omitempty"`

	AllowedStorageDestinations []EgressNetworkPolicyInternetAccessPolicyStorageDestination `json:"allowed_storage_destinations,omitempty"`
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	LogOnlyMode *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode `json:"log_only_mode,omitempty"`

	RestrictionMode EgressNetworkPolicyInternetAccessPolicyRestrictionMode `json:"restriction_mode,omitempty"`
}

func (s *EgressNetworkPolicyInternetAccessPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support domain name (FQDN) destinations for the time
// being, though going forwards we want to support host names and IP addresses.
type EgressNetworkPolicyInternetAccessPolicyInternetDestination struct {
	Destination string `json:"destination,omitempty"`

	Protocol EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol `json:"protocol,omitempty"`

	Type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EgressNetworkPolicyInternetAccessPolicyInternetDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EgressNetworkPolicyInternetAccessPolicyInternetDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The filtering protocol used by the DP. For private and public preview, SEG
// will only support TCP filtering (i.e. DNS based filtering, filtering by
// destination IP address), so protocol will be set to TCP by default and hidden
// from the user. In the future, users may be able to select HTTP filtering
// (i.e. SNI based filtering, filtering by FQDN).
type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol string

const EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolTcp EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol = `TCP`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) Set(v string) error {
	switch v {
	case `TCP`:
		*f = EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TCP"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) Values() []EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol {
	return []EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol{
		EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolTcp,
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol"
}

type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType string

const EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeFqdn EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType = `FQDN`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) Set(v string) error {
	switch v {
	case `FQDN`:
		*f = EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FQDN"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) Values() []EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType {
	return []EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType{
		EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeFqdn,
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType"
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyMode struct {
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType `json:"log_only_mode_type,omitempty"`

	Workloads []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType `json:"workloads,omitempty"`
}

func (s *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType string

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeAllServices EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType = `ALL_SERVICES`

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeSelectedServices EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType = `SELECTED_SERVICES`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) Set(v string) error {
	switch v {
	case `ALL_SERVICES`, `SELECTED_SERVICES`:
		*f = EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_SERVICES", "SELECTED_SERVICES"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) Values() []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType {
	return []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType{
		EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeAllServices,
		EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeSelectedServices,
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType"
}

// The values should match the list of workloads used in networkconfig.proto
type EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType string

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeDbsql EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType = `DBSQL`

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeMlServing EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType = `ML_SERVING`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) Set(v string) error {
	switch v {
	case `DBSQL`, `ML_SERVING`:
		*f = EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DBSQL", "ML_SERVING"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) Values() []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType {
	return []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType{
		EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeDbsql,
		EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeMlServing,
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType"
}

// At which level can Databricks and Databricks managed compute access Internet.
// FULL_ACCESS: Databricks can access Internet. No blocking rules will apply.
// RESTRICTED_ACCESS: Databricks can only access explicitly allowed internet and
// storage destinations, as well as UC connections and external locations.
// PRIVATE_ACCESS_ONLY (not used): Databricks can only access destinations via
// private link.
type EgressNetworkPolicyInternetAccessPolicyRestrictionMode string

const EgressNetworkPolicyInternetAccessPolicyRestrictionModeFullAccess EgressNetworkPolicyInternetAccessPolicyRestrictionMode = `FULL_ACCESS`

const EgressNetworkPolicyInternetAccessPolicyRestrictionModePrivateAccessOnly EgressNetworkPolicyInternetAccessPolicyRestrictionMode = `PRIVATE_ACCESS_ONLY`

const EgressNetworkPolicyInternetAccessPolicyRestrictionModeRestrictedAccess EgressNetworkPolicyInternetAccessPolicyRestrictionMode = `RESTRICTED_ACCESS`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) Set(v string) error {
	switch v {
	case `FULL_ACCESS`, `PRIVATE_ACCESS_ONLY`, `RESTRICTED_ACCESS`:
		*f = EgressNetworkPolicyInternetAccessPolicyRestrictionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL_ACCESS", "PRIVATE_ACCESS_ONLY", "RESTRICTED_ACCESS"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyRestrictionMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) Values() []EgressNetworkPolicyInternetAccessPolicyRestrictionMode {
	return []EgressNetworkPolicyInternetAccessPolicyRestrictionMode{
		EgressNetworkPolicyInternetAccessPolicyRestrictionModeFullAccess,
		EgressNetworkPolicyInternetAccessPolicyRestrictionModePrivateAccessOnly,
		EgressNetworkPolicyInternetAccessPolicyRestrictionModeRestrictedAccess,
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyRestrictionMode to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyRestrictionMode"
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyInternetAccessPolicyStorageDestination struct {
	AllowedPaths []string `json:"allowed_paths,omitempty"`

	AzureContainer string `json:"azure_container,omitempty"`

	AzureDnsZone string `json:"azure_dns_zone,omitempty"`

	AzureStorageAccount string `json:"azure_storage_account,omitempty"`

	AzureStorageService string `json:"azure_storage_service,omitempty"`

	BucketName string `json:"bucket_name,omitempty"`

	Region string `json:"region,omitempty"`

	Type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EgressNetworkPolicyInternetAccessPolicyStorageDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EgressNetworkPolicyInternetAccessPolicyStorageDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType string

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAwsS3 EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `AWS_S3`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `AZURE_STORAGE`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeCloudflareR2 EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `CLOUDFLARE_R2`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `GOOGLE_CLOUD_STORAGE`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) Set(v string) error {
	switch v {
	case `AWS_S3`, `AZURE_STORAGE`, `CLOUDFLARE_R2`, `GOOGLE_CLOUD_STORAGE`:
		*f = EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_S3", "AZURE_STORAGE", "CLOUDFLARE_R2", "GOOGLE_CLOUD_STORAGE"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) Values() []EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType {
	return []EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType{
		EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAwsS3,
		EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage,
		EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeCloudflareR2,
		EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage,
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType"
}

type EgressNetworkPolicyNetworkAccessPolicy struct {
	// List of Databricks workspace destinations that serverless workloads are
	// allowed to access when in RESTRICTED_ACCESS mode.
	AllowedDatabricksDestinations []EgressNetworkPolicyNetworkAccessPolicyDatabricksDestination `json:"allowed_databricks_destinations,omitempty"`
	// List of internet destinations that serverless workloads are allowed to
	// access when in RESTRICTED_ACCESS mode.
	AllowedInternetDestinations []EgressNetworkPolicyNetworkAccessPolicyInternetDestination `json:"allowed_internet_destinations,omitempty"`
	// List of storage destinations that serverless workloads are allowed to
	// access when in RESTRICTED_ACCESS mode.
	AllowedStorageDestinations []EgressNetworkPolicyNetworkAccessPolicyStorageDestination `json:"allowed_storage_destinations,omitempty"`
	// List of internet destinations that serverless workloads are blocked from
	// accessing. These destinations are enforced when restriction mode is
	// RESTRICTED_ACCESS or DRY_RUN. Currently supports DNS_NAME type only;
	// IP_RANGE support is planned.
	BlockedInternetDestinations []EgressNetworkPolicyNetworkAccessPolicyInternetDestination `json:"blocked_internet_destinations,omitempty"`
	// Optional. When policy_enforcement is not provided, we default to
	// ENFORCE_MODE_ALL_SERVICES
	PolicyEnforcement *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement `json:"policy_enforcement,omitempty"`
	// The restriction mode that controls how serverless workloads can access
	// the internet.
	RestrictionMode EgressNetworkPolicyNetworkAccessPolicyRestrictionMode `json:"restriction_mode"`
}

func (s *EgressNetworkPolicyNetworkAccessPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EgressNetworkPolicyNetworkAccessPolicyDatabricksDestination struct {
	// The workspace IDs to allow egress traffic to.
	WorkspaceIds []int64 `json:"workspace_ids,omitempty"`
}

func (s *EgressNetworkPolicyNetworkAccessPolicyDatabricksDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support DNS_NAME (FQDN format) destinations for the time
// being. Going forward we may extend support to host names and IP addresses.
type EgressNetworkPolicyNetworkAccessPolicyInternetDestination struct {
	// The internet destination to which access will be allowed. Format
	// dependent on the destination type.
	Destination string `json:"destination,omitempty"`
	// The type of internet destination. Currently only DNS_NAME is supported.
	InternetDestinationType EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType `json:"internet_destination_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EgressNetworkPolicyNetworkAccessPolicyInternetDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType string

const EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypeDnsName EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType = `DNS_NAME`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType) Set(v string) error {
	switch v {
	case `DNS_NAME`:
		*f = EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DNS_NAME"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType) Values() []EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType {
	return []EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType{
		EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypeDnsName,
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType"
}

type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement struct {
	// When empty, it means dry run for all products. When non-empty, it means
	// dry run for specific products and for the other products, they will run
	// in enforced mode.
	DryRunModeProductFilter []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter `json:"dry_run_mode_product_filter,omitempty"`
	// The mode of policy enforcement. ENFORCED blocks traffic that violates
	// policy, while DRY_RUN only logs violations without blocking. When not
	// specified, defaults to ENFORCED.
	EnforcementMode EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode `json:"enforcement_mode,omitempty"`
}

func (s *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The values should match the list of workloads used in networkconfig.proto
type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter string

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterDbsql EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter = `DBSQL`

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterMlServing EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter = `ML_SERVING`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter) Set(v string) error {
	switch v {
	case `DBSQL`, `ML_SERVING`:
		*f = EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DBSQL", "ML_SERVING"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter) Values() []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter {
	return []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter{
		EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterDbsql,
		EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterMlServing,
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter"
}

type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode string

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeDryRun EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode = `DRY_RUN`

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeEnforced EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode = `ENFORCED`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode) Set(v string) error {
	switch v {
	case `DRY_RUN`, `ENFORCED`:
		*f = EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DRY_RUN", "ENFORCED"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode) Values() []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode {
	return []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode{
		EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeDryRun,
		EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeEnforced,
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode"
}

// At which level can Databricks and Databricks managed compute access Internet.
// FULL_ACCESS: Databricks can access Internet. No blocking rules will apply.
// RESTRICTED_ACCESS: Databricks can only access explicitly allowed internet and
// storage destinations, as well as UC connections and external locations.
type EgressNetworkPolicyNetworkAccessPolicyRestrictionMode string

const EgressNetworkPolicyNetworkAccessPolicyRestrictionModeFullAccess EgressNetworkPolicyNetworkAccessPolicyRestrictionMode = `FULL_ACCESS`

const EgressNetworkPolicyNetworkAccessPolicyRestrictionModeRestrictedAccess EgressNetworkPolicyNetworkAccessPolicyRestrictionMode = `RESTRICTED_ACCESS`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyRestrictionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyRestrictionMode) Set(v string) error {
	switch v {
	case `FULL_ACCESS`, `RESTRICTED_ACCESS`:
		*f = EgressNetworkPolicyNetworkAccessPolicyRestrictionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL_ACCESS", "RESTRICTED_ACCESS"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyRestrictionMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyRestrictionMode) Values() []EgressNetworkPolicyNetworkAccessPolicyRestrictionMode {
	return []EgressNetworkPolicyNetworkAccessPolicyRestrictionMode{
		EgressNetworkPolicyNetworkAccessPolicyRestrictionModeFullAccess,
		EgressNetworkPolicyNetworkAccessPolicyRestrictionModeRestrictedAccess,
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyRestrictionMode to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyRestrictionMode) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyRestrictionMode"
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyNetworkAccessPolicyStorageDestination struct {
	// The Azure storage account name.
	AzureStorageAccount string `json:"azure_storage_account,omitempty"`
	// The Azure storage service type (blob, dfs, etc.).
	AzureStorageService string `json:"azure_storage_service,omitempty"`

	BucketName string `json:"bucket_name,omitempty"`

	Region string `json:"region,omitempty"`
	// The type of storage destination.
	StorageDestinationType EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType `json:"storage_destination_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EgressNetworkPolicyNetworkAccessPolicyStorageDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType string

const EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAwsS3 EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType = `AWS_S3`

const EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType = `AZURE_STORAGE`

const EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType = `GOOGLE_CLOUD_STORAGE`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType) Set(v string) error {
	switch v {
	case `AWS_S3`, `AZURE_STORAGE`, `GOOGLE_CLOUD_STORAGE`:
		*f = EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_S3", "AZURE_STORAGE", "GOOGLE_CLOUD_STORAGE"`, v)
	}
}

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType) Values() []EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType {
	return []EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType{
		EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAwsS3,
		EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage,
		EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage,
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType"
}

// The target resources that are supported by Network Connectivity Config. Note:
// some egress types can support general types that are not defined in
// EgressResourceType. E.g.: Azure private endpoint supports private link
// enabled Azure services.
type EgressResourceType string

const EgressResourceTypeAzureBlobStorage EgressResourceType = `AZURE_BLOB_STORAGE`

// String representation for [fmt.Print]
func (f *EgressResourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressResourceType) Set(v string) error {
	switch v {
	case `AZURE_BLOB_STORAGE`:
		*f = EgressResourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AZURE_BLOB_STORAGE"`, v)
	}
}

// Values returns all possible values for EgressResourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressResourceType) Values() []EgressResourceType {
	return []EgressResourceType{
		EgressResourceTypeAzureBlobStorage,
	}
}

// Type always returns EgressResourceType to satisfy [pflag.Value] interface
func (f *EgressResourceType) Type() string {
	return "EgressResourceType"
}

// Configures an email destination.
type EmailConfig struct {
	// Email addresses to notify. At least one address is required.
	Addresses []string `json:"addresses,omitempty"`
}

func (s *EmailConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EnableExportNotebook struct {
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnableExportNotebook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnableExportNotebook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnableNotebookTableClipboard struct {
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnableNotebookTableClipboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnableNotebookTableClipboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnableResultsDownloading struct {
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnableResultsDownloading) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnableResultsDownloading) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// SHIELD feature: ESM Enhanced Security Monitoring (ESM) enables additional
// security monitoring on the workspace.
type EnhancedSecurityMonitoring struct {
	// Whether Enhanced Security Monitoring (ESM) is enabled on the workspace.
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnhancedSecurityMonitoring) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnhancedSecurityMonitoring) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnhancedSecurityMonitoringSetting struct {
	EnhancedSecurityMonitoringWorkspace EnhancedSecurityMonitoring `json:"enhanced_security_monitoring_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnhancedSecurityMonitoringSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnhancedSecurityMonitoringSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Account level policy for ESM
type EsmEnablementAccount struct {
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EsmEnablementAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EsmEnablementAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EsmEnablementAccountSetting struct {
	EsmEnablementAccount EsmEnablementAccount `json:"esm_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	Credential string `json:"credential,omitempty"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime int64 `json:"credentialEolTime,omitempty"`
	// User ID of the user that owns this token.
	OwnerId int64 `json:"ownerId,omitempty"`
	// The scopes of access granted in the token.
	Scopes []string `json:"scopes,omitempty"`
	// The type of this exchange token
	TokenType TokenType `json:"tokenType,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	PartitionId PartitionId `json:"partitionId"`
	// Array of scopes for the token request.
	Scopes []string `json:"scopes"`
	// A list of token types being requested
	TokenType []TokenType `json:"tokenType"`
}

func (s *ExchangeTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {
	Values []ExchangeToken `json:"values,omitempty"`
}

func (s *ExchangeTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func (s *FetchIpAccessListResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GcpEndpoint struct {
	// All Google APIs that support VPC Service Controls (a subset of all Google
	// APIs).
	AllVpcScServices bool `json:"all_vpc_sc_services,omitempty"`
	// Selected Google API hostnames, e.g. "storage.googleapis.com",
	// "bigquery.googleapis.com".
	GoogleApiEndpoints *GoogleApiEndpoints `json:"google_api_endpoints,omitempty"`
	// Output only. The URI of the created PSC endpoint.
	PscEndpointUri string `json:"psc_endpoint_uri,omitempty"`
	// The full url of the target service attachment. Example:
	// projects/my-gcp-project/regions/us-east4/serviceAttachments/my-service-attachment
	ServiceAttachment string `json:"service_attachment,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configures a generic webhook destination. Notifications are delivered as
// HTTPS POST requests to url, optionally authenticated with HTTP Basic
// authentication.
type GenericWebhookConfig struct {
	// [Input-Only][Optional] Password for HTTP Basic authentication. Set it
	// together with username.
	Password string `json:"password,omitempty"`
	// [Output-Only] Whether password is set.
	PasswordSet bool `json:"password_set,omitempty"`
	// [Input-Only] HTTPS URL that notifications are posted to.
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	UrlSet bool `json:"url_set,omitempty"`
	// [Input-Only][Optional] Username for HTTP Basic authentication. Set it
	// together with password.
	Username string `json:"username,omitempty"`
	// [Output-Only] Whether username is set.
	UsernameSet bool `json:"username_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenericWebhookConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenericWebhookConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

func (s *GetAccountIpAccessListRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetAutomaticClusterUpdateSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetAutomaticClusterUpdateSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAutomaticClusterUpdateSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetComplianceSecurityProfileSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetComplianceSecurityProfileSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetComplianceSecurityProfileSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetCspEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetCspEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCspEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetDashboardEmailSubscriptionsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDashboardEmailSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDashboardEmailSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetDefaultWarehouseIdRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDefaultWarehouseIdRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDefaultWarehouseIdRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetEnhancedSecurityMonitoringSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetEnhancedSecurityMonitoringSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEnhancedSecurityMonitoringSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetEsmEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetEsmEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEsmEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

func (s *GetIpAccessListRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetIpAccessListResponse struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func (s *GetIpAccessListResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

func (s *GetIpAccessListsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetLlmProxyPartnerPoweredAccountRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetLlmProxyPartnerPoweredAccountRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetLlmProxyPartnerPoweredAccountRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetLlmProxyPartnerPoweredEnforceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetLlmProxyPartnerPoweredEnforceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetLlmProxyPartnerPoweredEnforceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetLlmProxyPartnerPoweredWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetLlmProxyPartnerPoweredWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

func (s *GetNetworkConnectivityConfigurationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetNetworkPolicyRequest struct {
	// The unique identifier of the network policy to retrieve.
	NetworkPolicyId string `json:"-" url:"-"`
}

func (s *GetNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetNotificationDestinationRequest struct {
	// A UUID identifying the notification destination.
	Id string `json:"-" url:"-"`
}

func (s *GetNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
}

func (s *GetPrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetSqlResultsDownloadRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetSqlResultsDownloadRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSqlResultsDownloadRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetStatusRequest struct {
	Keys string `json:"-" url:"keys"`
}

func (s *GetStatusRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`
}

func (s *GetTokenManagementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []TokenPermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetTokenPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
}

func (s *GetTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceNetworkOptionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Wrapper for a list of Google API hostnames. Wrapped in a message because
// proto3 oneof does not support repeated fields directly.
type GoogleApiEndpoints struct {
	// Google API hostnames, e.g. "storage.googleapis.com",
	// "bigquery.googleapis.com". Use "googleapis.com" to cover all Google APIs.
	Endpoints []string `json:"endpoints,omitempty"`
}

func (s *GoogleApiEndpoints) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Definition of an IP Access list
type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	AddressCount int `json:"address_count,omitempty"`
	// Creation timestamp in milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// User ID of the user who created this list.
	CreatedBy int64 `json:"created_by,omitempty"`
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId string `json:"list_id,omitempty"`

	ListType ListType `json:"list_type,omitempty"`
	// Update timestamp in milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// User ID of the user who updated this list.
	UpdatedBy int64 `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IpAccessListInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IpAccessListInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

func (s *ListIpAccessListResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNetworkConnectivityConfigurationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkConnectivityConfigurationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network connectivity configuration list was successfully retrieved.
type ListNetworkConnectivityConfigurationsResponse struct {
	Items []NetworkConnectivityConfiguration `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNetworkConnectivityConfigurationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkConnectivityConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListNetworkPoliciesRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNetworkPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListNetworkPoliciesResponse struct {
	// List of network policies.
	Items []AccountNetworkPolicy `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNetworkPoliciesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkPoliciesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListNotificationDestinationsRequest struct {
	// Maximum number of notification destinations to return per page. Defaults
	// to 20.
	PageSize int64 `json:"-" url:"page_size,omitempty"`
	// Opaque token returned as next_page_token by a previous call. Omit it to
	// request the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNotificationDestinationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNotificationDestinationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListNotificationDestinationsResponse struct {
	// Token to pass as page_token to fetch the next page of results. Absent
	// when there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Notification destinations in the current page.
	Results []ListNotificationDestinationsResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNotificationDestinationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNotificationDestinationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A reusable definition of where a workspace sends notifications, such as a
// Slack channel, a set of email addresses, a PagerDuty service, a Microsoft
// Teams channel or a generic webhook. Job notifications, SQL alerts and other
// workspace features reference a destination by its id. Only workspace admins
// can create, update and delete destinations; once created, a destination is
// available to all users in the workspace. Secret values such as webhook URLs
// and tokens are write-only: responses report whether each one is set, never
// its value.
type ListNotificationDestinationsResult struct {
	// [Output-only] The type of the notification destination, determined by
	// which nested config was provided on creation. The type can not be changed
	// once set.
	DestinationType DestinationType `json:"destination_type,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`
	// A UUID identifying the notification destination.
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNotificationDestinationsResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNotificationDestinationsResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPrivateEndpointRulesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPrivateEndpointRulesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The private endpoint rule list was successfully retrieved.
type ListPrivateEndpointRulesResponse struct {
	Items []NccPrivateEndpointRule `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPrivateEndpointRulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPrivateEndpointRulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos []PublicTokenInfo `json:"token_infos,omitempty"`
}

func (s *ListPublicTokensResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById int64 `json:"-" url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"-" url:"created_by_username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

func (s *ListTokensResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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

// Values returns all possible values for ListType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListType) Values() []ListType {
	return []ListType{
		ListTypeAllow,
		ListTypeBlock,
	}
}

// Type always returns ListType to satisfy [pflag.Value] interface
func (f *ListType) Type() string {
	return "ListType"
}

type LlmProxyPartnerPoweredAccount struct {
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LlmProxyPartnerPoweredAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LlmProxyPartnerPoweredAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LlmProxyPartnerPoweredEnforce struct {
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LlmProxyPartnerPoweredEnforce) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LlmProxyPartnerPoweredEnforce) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LlmProxyPartnerPoweredWorkspace struct {
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LlmProxyPartnerPoweredWorkspace) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LlmProxyPartnerPoweredWorkspace) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configures a Microsoft Teams destination. Notifications are posted to the
// incoming webhook url, which is required. The app_id, auth_secret, channel_url
// and tenant_id fields can optionally be set together to post through a
// Microsoft Teams app.
type MicrosoftTeamsConfig struct {
	// [Input-Only] Application (client) ID of the Microsoft Teams app, as a
	// UUID. Optional; set it together with auth_secret, channel_url and
	// tenant_id.
	AppId string `json:"app_id,omitempty"`
	// [Output-Only] Whether App ID is set.
	AppIdSet bool `json:"app_id_set,omitempty"`
	// [Input-Only] Client secret of the Microsoft Teams app. Optional; set it
	// together with app_id, channel_url and tenant_id.
	AuthSecret string `json:"auth_secret,omitempty"`
	// [Output-Only] Whether secret is set.
	AuthSecretSet bool `json:"auth_secret_set,omitempty"`
	// [Input-Only] URL of the Microsoft Teams channel that the app posts to.
	// Optional; set it together with app_id, auth_secret and tenant_id.
	ChannelUrl string `json:"channel_url,omitempty"`
	// [Output-Only] Whether Channel URL is set.
	ChannelUrlSet bool `json:"channel_url_set,omitempty"`
	// [Input-Only] Microsoft Entra tenant ID of the Microsoft Teams app, as a
	// UUID. Optional; set it together with app_id, auth_secret and channel_url.
	TenantId string `json:"tenant_id,omitempty"`
	// [Output-Only] Whether Tenant ID is set.
	TenantIdSet bool `json:"tenant_id_set,omitempty"`
	// [Input-Only] Microsoft Teams incoming webhook URL that notifications are
	// posted to. Must be an HTTPS URL. The URL is required even when the app
	// fields are set.
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	UrlSet bool `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MicrosoftTeamsConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MicrosoftTeamsConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The stable AWS IP CIDR blocks. You can use these to configure the firewall of
// your resources to allow traffic from your Databricks workspace.
type NccAwsStableIpRule struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
}

func (s *NccAwsStableIpRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type NccAzurePrivateEndpointRule struct {
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is ESTABLISHED. Remember that
	// you must approve new endpoints on your resources in the Azure portal
	// before they take effect. The possible values are: - INIT: (deprecated)
	// The endpoint has been created and pending approval. - PENDING: The
	// endpoint has been created and pending approval. - ESTABLISHED: The
	// endpoint has been approved and is ready to use in your serverless compute
	// resources. - REJECTED: Connection was rejected by the private link
	// resource owner. - DISCONNECTED: Connection was removed by the private
	// link resource owner, the private endpoint becomes informative and should
	// be deleted for clean-up. - EXPIRED: If the endpoint was created but not
	// approved in 14 days, it will be EXPIRED.
	ConnectionState NccAzurePrivateEndpointRuleConnectionState `json:"connection_state,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Whether this private endpoint is deactivated.
	Deactivated bool `json:"deactivated,omitempty"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt int64 `json:"deactivated_at,omitempty"`
	// Not used by customer-managed private endpoint services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string `json:"domain_names,omitempty"`
	// The name of the Azure private endpoint resource.
	EndpointName string `json:"endpoint_name,omitempty"`

	ErrorMessage string `json:"error_message,omitempty"`
	// Only used by private endpoints to Azure first-party services.
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId string `json:"group_id,omitempty"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The Azure resource ID of the target resource.
	ResourceId string `json:"resource_id,omitempty"`
	// The ID of a private endpoint rule.
	RuleId string `json:"rule_id,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NccAzurePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NccAzurePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NccAzurePrivateEndpointRuleConnectionState string

const NccAzurePrivateEndpointRuleConnectionStateCreateFailed NccAzurePrivateEndpointRuleConnectionState = `CREATE_FAILED`

const NccAzurePrivateEndpointRuleConnectionStateCreating NccAzurePrivateEndpointRuleConnectionState = `CREATING`

const NccAzurePrivateEndpointRuleConnectionStateDisconnected NccAzurePrivateEndpointRuleConnectionState = `DISCONNECTED`

const NccAzurePrivateEndpointRuleConnectionStateEstablished NccAzurePrivateEndpointRuleConnectionState = `ESTABLISHED`

const NccAzurePrivateEndpointRuleConnectionStateExpired NccAzurePrivateEndpointRuleConnectionState = `EXPIRED`

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
	case `CREATE_FAILED`, `CREATING`, `DISCONNECTED`, `ESTABLISHED`, `EXPIRED`, `INIT`, `PENDING`, `REJECTED`:
		*f = NccAzurePrivateEndpointRuleConnectionState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATE_FAILED", "CREATING", "DISCONNECTED", "ESTABLISHED", "EXPIRED", "INIT", "PENDING", "REJECTED"`, v)
	}
}

// Values returns all possible values for NccAzurePrivateEndpointRuleConnectionState.
//
// There is no guarantee on the order of the values in the slice.
func (f *NccAzurePrivateEndpointRuleConnectionState) Values() []NccAzurePrivateEndpointRuleConnectionState {
	return []NccAzurePrivateEndpointRuleConnectionState{
		NccAzurePrivateEndpointRuleConnectionStateCreateFailed,
		NccAzurePrivateEndpointRuleConnectionStateCreating,
		NccAzurePrivateEndpointRuleConnectionStateDisconnected,
		NccAzurePrivateEndpointRuleConnectionStateEstablished,
		NccAzurePrivateEndpointRuleConnectionStateExpired,
		NccAzurePrivateEndpointRuleConnectionStateInit,
		NccAzurePrivateEndpointRuleConnectionStatePending,
		NccAzurePrivateEndpointRuleConnectionStateRejected,
	}
}

// Type always returns NccAzurePrivateEndpointRuleConnectionState to satisfy [pflag.Value] interface
func (f *NccAzurePrivateEndpointRuleConnectionState) Type() string {
	return "NccAzurePrivateEndpointRuleConnectionState"
}

// The stable Azure service endpoints. You can configure the firewall of your
// Azure resources to allow traffic from your Databricks serverless compute
// resources.
type NccAzureServiceEndpointRule struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets []string `json:"subnets,omitempty"`
	// The Azure region in which this service endpoint rule applies..
	TargetRegion string `json:"target_region,omitempty"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices []EgressResourceType `json:"target_services,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NccAzureServiceEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NccAzureServiceEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NccEgressConfig struct {
	// Deprecated. This field is retained for compatibility with existing
	// clients. Avoid using it in new integrations.
	DefaultRules *NccEgressDefaultRules `json:"default_rules,omitempty"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules *NccEgressTargetRules `json:"target_rules,omitempty"`
}

func (s *NccEgressConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Default rules don't have specific targets.
type NccEgressDefaultRules struct {
	AwsStableIpRule *NccAwsStableIpRule `json:"aws_stable_ip_rule,omitempty"`

	AzureServiceEndpointRule *NccAzureServiceEndpointRule `json:"azure_service_endpoint_rule,omitempty"`
}

func (s *NccEgressDefaultRules) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Target rule controls the egress rules that are dedicated to specific
// resources.
type NccEgressTargetRules struct {
	// Deprecated. This field is not supported anymore and kept purely for
	// interface stability.
	AwsPrivateEndpointRules []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule `json:"aws_private_endpoint_rules,omitempty"`
	// Deprecated. This field is not supported anymore and kept purely for
	// interface stability.
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRule `json:"azure_private_endpoint_rules,omitempty"`
}

func (s *NccEgressTargetRules) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type NccPrivateEndpointRule struct {
	// Databricks account ID. You can find your account ID from the Accounts
	// Console.
	AccountId string `json:"account_id,omitempty"`
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is ESTABLISHED. Remember that
	// you must approve new endpoints on your resources in the Cloud console
	// before they take effect. The possible values are: - PENDING: The endpoint
	// has been created and pending approval. - ESTABLISHED: The endpoint has
	// been approved and is ready to use in your serverless compute resources. -
	// REJECTED: Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up. - EXPIRED: If the endpoint was created but not approved in 14
	// days, it will be EXPIRED. - CREATING: The endpoint creation is in
	// progress. Once successfully created, the state will transition to
	// PENDING. - CREATE_FAILED: The endpoint creation failed. You can check the
	// error_message field for more details.
	ConnectionState NccPrivateEndpointRulePrivateLinkConnectionState `json:"connection_state,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Whether this private endpoint is deactivated.
	Deactivated bool `json:"deactivated,omitempty"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt int64 `json:"deactivated_at,omitempty"`
	// Only used by private endpoints to customer-managed private endpoint
	// services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string `json:"domain_names,omitempty"`
	// Update this field to activate/deactivate this private endpoint to allow
	// egress access from serverless compute resources. Only honored for
	// first-party services on each cloud (e.g. AWS S3).
	Enabled bool `json:"enabled,omitempty"`
	// The name of the Azure private endpoint resource.
	EndpointName string `json:"endpoint_name,omitempty"`
	// The full target AWS endpoint service name that connects to the
	// destination resources of the private endpoint.
	EndpointService string `json:"endpoint_service,omitempty"`

	ErrorMessage string `json:"error_message,omitempty"`

	GcpEndpoint *GcpEndpoint `json:"gcp_endpoint,omitempty"`
	// Not used by customer-managed private endpoint services.
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId string `json:"group_id,omitempty"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The Azure resource ID of the target resource.
	ResourceId string `json:"resource_id,omitempty"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	ResourceNames []string `json:"resource_names,omitempty"`
	// The ID of a private endpoint rule.
	RuleId string `json:"rule_id,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`
	// The AWS VPC endpoint ID. You can use this ID to identify the VPC endpoint
	// created by Databricks.
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NccPrivateEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NccPrivateEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NccPrivateEndpointRulePrivateLinkConnectionState string

const NccPrivateEndpointRulePrivateLinkConnectionStateCreateFailed NccPrivateEndpointRulePrivateLinkConnectionState = `CREATE_FAILED`

const NccPrivateEndpointRulePrivateLinkConnectionStateCreating NccPrivateEndpointRulePrivateLinkConnectionState = `CREATING`

const NccPrivateEndpointRulePrivateLinkConnectionStateDisconnected NccPrivateEndpointRulePrivateLinkConnectionState = `DISCONNECTED`

const NccPrivateEndpointRulePrivateLinkConnectionStateEstablished NccPrivateEndpointRulePrivateLinkConnectionState = `ESTABLISHED`

const NccPrivateEndpointRulePrivateLinkConnectionStateExpired NccPrivateEndpointRulePrivateLinkConnectionState = `EXPIRED`

const NccPrivateEndpointRulePrivateLinkConnectionStatePending NccPrivateEndpointRulePrivateLinkConnectionState = `PENDING`

const NccPrivateEndpointRulePrivateLinkConnectionStateRejected NccPrivateEndpointRulePrivateLinkConnectionState = `REJECTED`

// String representation for [fmt.Print]
func (f *NccPrivateEndpointRulePrivateLinkConnectionState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NccPrivateEndpointRulePrivateLinkConnectionState) Set(v string) error {
	switch v {
	case `CREATE_FAILED`, `CREATING`, `DISCONNECTED`, `ESTABLISHED`, `EXPIRED`, `PENDING`, `REJECTED`:
		*f = NccPrivateEndpointRulePrivateLinkConnectionState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATE_FAILED", "CREATING", "DISCONNECTED", "ESTABLISHED", "EXPIRED", "PENDING", "REJECTED"`, v)
	}
}

// Values returns all possible values for NccPrivateEndpointRulePrivateLinkConnectionState.
//
// There is no guarantee on the order of the values in the slice.
func (f *NccPrivateEndpointRulePrivateLinkConnectionState) Values() []NccPrivateEndpointRulePrivateLinkConnectionState {
	return []NccPrivateEndpointRulePrivateLinkConnectionState{
		NccPrivateEndpointRulePrivateLinkConnectionStateCreateFailed,
		NccPrivateEndpointRulePrivateLinkConnectionStateCreating,
		NccPrivateEndpointRulePrivateLinkConnectionStateDisconnected,
		NccPrivateEndpointRulePrivateLinkConnectionStateEstablished,
		NccPrivateEndpointRulePrivateLinkConnectionStateExpired,
		NccPrivateEndpointRulePrivateLinkConnectionStatePending,
		NccPrivateEndpointRulePrivateLinkConnectionStateRejected,
	}
}

// Type always returns NccPrivateEndpointRulePrivateLinkConnectionState to satisfy [pflag.Value] interface
func (f *NccPrivateEndpointRulePrivateLinkConnectionState) Type() string {
	return "NccPrivateEndpointRulePrivateLinkConnectionState"
}

// Properties of the new network connectivity configuration.
type NetworkConnectivityConfiguration struct {
	// Your Databricks account ID. You can find your account ID in your
	// Databricks accounts console.
	AccountId string `json:"account_id,omitempty"`
	// The workspace to which charges for this network connectivity
	// configuration are attributed. The workspace must be active and attached
	// to this network connectivity configuration when set.
	BillingWorkspaceId int64 `json:"billing_workspace_id,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig *NccEgressConfig `json:"egress_config,omitempty"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
	Name string `json:"name,omitempty"`
	// Databricks network connectivity configuration ID.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region string `json:"region,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NetworkConnectivityConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkConnectivityConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network policies applying for egress traffic.
type NetworkPolicyEgress struct {
	// The access policy enforced for egress traffic to the internet.
	NetworkAccess *EgressNetworkPolicyNetworkAccessPolicy `json:"network_access,omitempty"`
}

func (s *NetworkPolicyEgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A reusable definition of where a workspace sends notifications, such as a
// Slack channel, a set of email addresses, a PagerDuty service, a Microsoft
// Teams channel or a generic webhook. Job notifications, SQL alerts and other
// workspace features reference a destination by its id. Only workspace admins
// can create, update and delete destinations; once created, a destination is
// available to all users in the workspace. Secret values such as webhook URLs
// and tokens are write-only: responses report whether each one is set, never
// its value.
type NotificationDestination struct {
	// The configuration for the notification destination. Will be exactly one
	// of the nested configs. Only returned to users with workspace admin
	// permissions.
	Config *Config `json:"config,omitempty"`
	// [Output-only] The type of the notification destination, determined by
	// which nested config was provided on creation. The type can not be changed
	// once set.
	DestinationType DestinationType `json:"destination_type,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`
	// A UUID identifying the notification destination.
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NotificationDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotificationDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configures a PagerDuty destination. Notifications are sent as events to the
// PagerDuty Events API v2.
type PagerdutyConfig struct {
	// [Input-Only] Integration key from an Events API v2 integration on the
	// PagerDuty service that receives the notifications.
	IntegrationKey string `json:"integration_key,omitempty"`
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet bool `json:"integration_key_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PagerdutyConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PagerdutyConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId int64 `json:"workspaceId,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PartitionId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PartitionId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PersonalComputeMessage struct {
	Value PersonalComputeMessageEnum `json:"value"`
}

func (s *PersonalComputeMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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

// Values returns all possible values for PersonalComputeMessageEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *PersonalComputeMessageEnum) Values() []PersonalComputeMessageEnum {
	return []PersonalComputeMessageEnum{
		PersonalComputeMessageEnumDelegate,
		PersonalComputeMessageEnumOn,
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
	Etag string `json:"etag,omitempty"`

	PersonalCompute PersonalComputeMessage `json:"personal_compute"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PersonalComputeSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PersonalComputeSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublicTokenInfo struct {
	// Output only. The autoscope state of this token.
	AutoscopeState iam.AutoscopeState `json:"autoscope_state,omitempty"`
	// Output only. Scopes inferred from offline backfill processing.
	BackfillScopes []string `json:"backfill_scopes,omitempty"`
	// Comment the token was created with, if applicable.
	Comment string `json:"comment,omitempty"`
	// Server time (in epoch milliseconds) when the token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// Output only. Inferred API path scopes collected for this token when
	// autoscope is enabled.
	InferredScopes []string `json:"inferred_scopes,omitempty"`
	// Scope of the token was created with, if applicable.
	Scopes []string `json:"scopes,omitempty"`
	// The ID of this token.
	TokenId string `json:"token_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	Enabled bool `json:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`

	ListType ListType `json:"list_type"`
}

func (s *ReplaceIpAccessList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RestrictWorkspaceAdminsMessage struct {
	// When true, workspace admins cannot create governance tags. ALLOW_ALL
	// status does not override this; they are independent.
	DisableGovTagCreation bool `json:"disable_gov_tag_creation,omitempty"`
	// Controls what workspace admins can do with service principal tokens and
	// job identities. With ALLOW_ALL, workspace admins can create personal
	// access tokens on behalf of any service principal in the workspace, change
	// a job owner to any workspace user, and set a job's run_as identity to any
	// workspace user or to a service principal they have the Service Principal
	// User role on. With RESTRICT_TOKENS_AND_JOB_RUN_AS, they can only create
	// tokens for service principals they have the Service Principal User role
	// on, change a job owner to themselves, and set run_as to themselves or to
	// a service principal they have that role on.
	Status RestrictWorkspaceAdminsMessageStatus `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestrictWorkspaceAdminsMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestrictWorkspaceAdminsMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestrictWorkspaceAdminsMessageStatus string

const RestrictWorkspaceAdminsMessageStatusAllowAll RestrictWorkspaceAdminsMessageStatus = `ALLOW_ALL`

const RestrictWorkspaceAdminsMessageStatusRestrictTokensAndJobRunAs RestrictWorkspaceAdminsMessageStatus = `RESTRICT_TOKENS_AND_JOB_RUN_AS`

// String representation for [fmt.Print]
func (f *RestrictWorkspaceAdminsMessageStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RestrictWorkspaceAdminsMessageStatus) Set(v string) error {
	switch v {
	case `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`:
		*f = RestrictWorkspaceAdminsMessageStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL", "RESTRICT_TOKENS_AND_JOB_RUN_AS"`, v)
	}
}

// Values returns all possible values for RestrictWorkspaceAdminsMessageStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *RestrictWorkspaceAdminsMessageStatus) Values() []RestrictWorkspaceAdminsMessageStatus {
	return []RestrictWorkspaceAdminsMessageStatus{
		RestrictWorkspaceAdminsMessageStatusAllowAll,
		RestrictWorkspaceAdminsMessageStatusRestrictTokensAndJobRunAs,
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
	Etag string `json:"etag,omitempty"`

	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessage `json:"restrict_workspace_admins"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestrictWorkspaceAdminsSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestrictWorkspaceAdminsSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId string `json:"token_id"`
}

func (s *RevokeTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Configures a Slack destination. Notifications are posted to the incoming
// webhook url, which is required. A bot oauth_token and channel_id can
// optionally be set together to post to the channel through the Slack API.
type SlackConfig struct {
	// [Input-Only] ID of the Slack channel that notifications are posted to
	// through the Slack API. Optional; set it together with oauth_token.
	ChannelId string `json:"channel_id,omitempty"`
	// [Output-Only] Whether channel ID is set.
	ChannelIdSet bool `json:"channel_id_set,omitempty"`
	// [Input-Only] Slack bot OAuth token (xoxb-...). Optional; when set
	// together with channel_id, notifications are posted to that channel
	// through the Slack API.
	OauthToken string `json:"oauth_token,omitempty"`
	// [Output-Only] Whether OAuth token is set.
	OauthTokenSet bool `json:"oauth_token_set,omitempty"`
	// [Input-Only] Slack incoming webhook URL that notifications are posted to.
	// Must be an HTTPS URL. The URL is required even when oauth_token and
	// channel_id are set.
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	UrlSet bool `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SlackConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SlackConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlResultsDownload struct {
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlResultsDownload) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlResultsDownload) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StringMessage struct {
	// Represents a generic string value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StringMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StringMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions []TokenPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenInfo struct {
	// Output only. The autoscope state of this token.
	AutoscopeState iam.AutoscopeState `json:"autoscope_state,omitempty"`
	// Output only. Scopes inferred from offline backfill processing.
	BackfillScopes []string `json:"backfill_scopes,omitempty"`
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment string `json:"comment,omitempty"`
	// User ID of the user that created the token.
	CreatedById int64 `json:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"created_by_username,omitempty"`
	// Timestamp when the token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Timestamp when the token expires.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// Output only. Inferred API path scopes collected for this token when
	// autoscope is enabled.
	InferredScopes []string `json:"inferred_scopes,omitempty"`
	// Approximate timestamp for the day the token was last used. Accurate up to
	// 1 day.
	LastUsedDay int64 `json:"last_used_day,omitempty"`
	// User ID of the user that owns the token.
	OwnerId int64 `json:"owner_id,omitempty"`
	// Scope of the token was created with, if applicable.
	Scopes []string `json:"scopes,omitempty"`
	// ID of the token.
	TokenId string `json:"token_id,omitempty"`
	// If applicable, the ID of the workspace that the token was created in.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

// Values returns all possible values for TokenPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *TokenPermissionLevel) Values() []TokenPermissionLevel {
	return []TokenPermissionLevel{
		TokenPermissionLevelCanUse,
	}
}

// Type always returns TokenPermissionLevel to satisfy [pflag.Value] interface
func (f *TokenPermissionLevel) Type() string {
	return "TokenPermissionLevel"
}

type TokenPermissions struct {
	AccessControlList []TokenAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsRequest struct {
	AccessControlList []TokenAccessControlRequest `json:"access_control_list,omitempty"`
}

func (s *TokenPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The type of token request. As of now, only `AZURE_ACTIVE_DIRECTORY_TOKEN` is
// supported.
type TokenType string

const TokenTypeArclightAzureExchangeToken TokenType = `ARCLIGHT_AZURE_EXCHANGE_TOKEN`

const TokenTypeArclightAzureExchangeTokenWithUserDelegationKey TokenType = `ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`

const TokenTypeArclightMultiTenantAzureExchangeToken TokenType = `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN`

const TokenTypeArclightMultiTenantAzureExchangeTokenWithUserDelegationKey TokenType = `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`

const TokenTypeAzureActiveDirectoryToken TokenType = `AZURE_ACTIVE_DIRECTORY_TOKEN`

// String representation for [fmt.Print]
func (f *TokenType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenType) Set(v string) error {
	switch v {
	case `ARCLIGHT_AZURE_EXCHANGE_TOKEN`, `ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`, `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN`, `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`, `AZURE_ACTIVE_DIRECTORY_TOKEN`:
		*f = TokenType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARCLIGHT_AZURE_EXCHANGE_TOKEN", "ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY", "ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN", "ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY", "AZURE_ACTIVE_DIRECTORY_TOKEN"`, v)
	}
}

// Values returns all possible values for TokenType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TokenType) Values() []TokenType {
	return []TokenType{
		TokenTypeArclightAzureExchangeToken,
		TokenTypeArclightAzureExchangeTokenWithUserDelegationKey,
		TokenTypeArclightMultiTenantAzureExchangeToken,
		TokenTypeArclightMultiTenantAzureExchangeTokenWithUserDelegationKey,
		TokenTypeAzureActiveDirectoryToken,
	}
}

// Type always returns TokenType to satisfy [pflag.Value] interface
func (f *TokenType) Type() string {
	return "TokenType"
}

// Details required to update a setting.
type UpdateAccountIpAccessEnableRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting AccountIpAccessEnable `json:"setting"`
}

func (s *UpdateAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting AibiDashboardEmbeddingAccessPolicySetting `json:"setting"`
}

func (s *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting AibiDashboardEmbeddingApprovedDomainsSetting `json:"setting"`
}

func (s *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateAutomaticClusterUpdateSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting AutomaticClusterUpdateSetting `json:"setting"`
}

func (s *UpdateAutomaticClusterUpdateSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateComplianceSecurityProfileSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting ComplianceSecurityProfileSetting `json:"setting"`
}

func (s *UpdateComplianceSecurityProfileSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateCspEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting CspEnablementAccountSetting `json:"setting"`
}

func (s *UpdateCspEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateDashboardEmailSubscriptionsRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting DashboardEmailSubscriptions `json:"setting"`
}

func (s *UpdateDashboardEmailSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateDefaultNamespaceSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting DefaultNamespaceSetting `json:"setting"`
}

func (s *UpdateDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateDefaultWarehouseIdRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting DefaultWarehouseId `json:"setting"`
}

func (s *UpdateDefaultWarehouseIdRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateDisableLegacyAccessRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting DisableLegacyAccess `json:"setting"`
}

func (s *UpdateDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateDisableLegacyDbfsRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting DisableLegacyDbfs `json:"setting"`
}

func (s *UpdateDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateDisableLegacyFeaturesRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting DisableLegacyFeatures `json:"setting"`
}

func (s *UpdateDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateEnableExportNotebookRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting EnableExportNotebook `json:"setting"`
}

func (s *UpdateEnableExportNotebookRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateEnableNotebookTableClipboardRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting EnableNotebookTableClipboard `json:"setting"`
}

func (s *UpdateEnableNotebookTableClipboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateEnableResultsDownloadingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting EnableResultsDownloading `json:"setting"`
}

func (s *UpdateEnableResultsDownloadingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateEnhancedSecurityMonitoringSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting EnhancedSecurityMonitoringSetting `json:"setting"`
}

func (s *UpdateEnhancedSecurityMonitoringSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateEsmEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting EsmEnablementAccountSetting `json:"setting"`
}

func (s *UpdateEsmEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update an IP access list.
type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`

	ListType ListType `json:"list_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateIpAccessList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateIpAccessList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredAccountRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting LlmProxyPartnerPoweredAccount `json:"setting"`
}

func (s *UpdateLlmProxyPartnerPoweredAccountRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredEnforceRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting LlmProxyPartnerPoweredEnforce `json:"setting"`
}

func (s *UpdateLlmProxyPartnerPoweredEnforceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredWorkspaceRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting LlmProxyPartnerPoweredWorkspace `json:"setting"`
}

func (s *UpdateLlmProxyPartnerPoweredWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateNccPrivateEndpointRuleRequest struct {
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId string `json:"-" url:"-"`

	PrivateEndpointRule UpdatePrivateEndpointRule `json:"private_endpoint_rule"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateNccPrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateNetworkConnectivityConfigRequest struct {
	// The network connectivity configuration fields to update.
	NetworkConnectivityConfig NetworkConnectivityConfiguration `json:"network_connectivity_config"`
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Must contain at least one of `name` or `billing_workspace_id`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateNetworkConnectivityConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateNetworkPolicyRequest struct {
	// Updated network policy configuration details.
	NetworkPolicy AccountNetworkPolicy `json:"network_policy"`
	// The unique identifier for the network policy.
	NetworkPolicyId string `json:"-" url:"-"`
}

func (s *UpdateNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config *Config `json:"config,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`
	// A UUID identifying the notification destination.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting PersonalComputeSetting `json:"setting"`
}

func (s *UpdatePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type UpdatePrivateEndpointRule struct {
	// Only used by private endpoints to customer-managed private endpoint
	// services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string `json:"domain_names,omitempty"`
	// Update this field to activate/deactivate this private endpoint to allow
	// egress access from serverless compute resources. Only honored for
	// first-party services on each cloud (e.g. AWS S3).
	Enabled bool `json:"enabled,omitempty"`

	ErrorMessage string `json:"error_message,omitempty"`

	GcpEndpoint *GcpEndpoint `json:"gcp_endpoint,omitempty"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	ResourceNames []string `json:"resource_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdatePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdatePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting RestrictWorkspaceAdminsSetting `json:"setting"`
}

func (s *UpdateRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a setting.
type UpdateSqlResultsDownloadRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask string `json:"field_mask"`

	Setting SqlResultsDownload `json:"setting"`
}

func (s *UpdateSqlResultsDownloadRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// For the list of supported token scopes, see
// https://docs.databricks.com/api/workspace/api/scopes.
type UpdateTokenManagementRequest struct {
	Token TokenInfo `json:"token"`
	// ID of the token.
	TokenId string `json:"-" url:"-"`
	// A list of field name under token, For example, {"update_mask":
	// "comment,scopes"}
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask fieldmask.FieldMask `json:"update_mask"`
}

func (s *UpdateTokenManagementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// For the list of supported token scopes, see
// https://docs.databricks.com/api/workspace/api/scopes.
type UpdateTokenRequest struct {
	Token PublicTokenInfo `json:"token"`
	// The SHA-256 hash of the token to be updated.
	TokenId string `json:"-" url:"-"`
	// A list of field name under token, For example, {"update_mask":
	// "comment,scopes"}
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask fieldmask.FieldMask `json:"update_mask"`
}

func (s *UpdateTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateTokenResponse struct {
}

func (s *UpdateTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
	// The network option details for the workspace.
	WorkspaceNetworkOption WorkspaceNetworkOption `json:"workspace_network_option"`
}

func (s *UpdateWorkspaceNetworkOptionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type WorkspaceConf map[string]string

type WorkspaceNetworkOption struct {
	// The network policy ID to apply to the workspace. This controls the
	// network access rules for all serverless compute resources in the
	// workspace. Each workspace can only be linked to one policy at a time. If
	// no policy is explicitly assigned, the workspace will use
	// 'default-policy'.
	NetworkPolicyId string `json:"network_policy_id,omitempty"`
	// The workspace ID.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceNetworkOption) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceNetworkOption) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
