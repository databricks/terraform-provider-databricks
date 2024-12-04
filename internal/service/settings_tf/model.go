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
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AibiDashboardEmbeddingAccessPolicy struct {
	AccessPolicyType types.String `tfsdk:"access_policy_type" tf:""`
}

func (newState *AibiDashboardEmbeddingAccessPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingAccessPolicy) {
}

func (newState *AibiDashboardEmbeddingAccessPolicy) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingAccessPolicy) {
}

func (a AibiDashboardEmbeddingAccessPolicy) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a AibiDashboardEmbeddingAccessPolicy) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AccessPolicyType": types.StringType,
		},
	}
}

type AibiDashboardEmbeddingAccessPolicySetting struct {
	AibiDashboardEmbeddingAccessPolicy types.List `tfsdk:"aibi_dashboard_embedding_access_policy" tf:"object"`
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

func (newState *AibiDashboardEmbeddingAccessPolicySetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingAccessPolicySetting) {
}

func (newState *AibiDashboardEmbeddingAccessPolicySetting) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingAccessPolicySetting) {
}

func (a AibiDashboardEmbeddingAccessPolicySetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AibiDashboardEmbeddingAccessPolicy": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy{}),
	}
}

func (a AibiDashboardEmbeddingAccessPolicySetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AibiDashboardEmbeddingAccessPolicy": AibiDashboardEmbeddingAccessPolicy{}.ToAttrType(ctx),
			"Etag":                               types.StringType,
			"SettingName":                        types.StringType,
		},
	}
}

type AibiDashboardEmbeddingApprovedDomains struct {
	ApprovedDomains types.List `tfsdk:"approved_domains" tf:"optional"`
}

func (newState *AibiDashboardEmbeddingApprovedDomains) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingApprovedDomains) {
}

func (newState *AibiDashboardEmbeddingApprovedDomains) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingApprovedDomains) {
}

func (a AibiDashboardEmbeddingApprovedDomains) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ApprovedDomains": reflect.TypeOf(types.StringType),
	}
}

func (a AibiDashboardEmbeddingApprovedDomains) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ApprovedDomains": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type AibiDashboardEmbeddingApprovedDomainsSetting struct {
	AibiDashboardEmbeddingApprovedDomains types.List `tfsdk:"aibi_dashboard_embedding_approved_domains" tf:"object"`
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

func (newState *AibiDashboardEmbeddingApprovedDomainsSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingApprovedDomainsSetting) {
}

func (newState *AibiDashboardEmbeddingApprovedDomainsSetting) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingApprovedDomainsSetting) {
}

func (a AibiDashboardEmbeddingApprovedDomainsSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AibiDashboardEmbeddingApprovedDomains": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains{}),
	}
}

func (a AibiDashboardEmbeddingApprovedDomainsSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AibiDashboardEmbeddingApprovedDomains": AibiDashboardEmbeddingApprovedDomains{}.ToAttrType(ctx),
			"Etag":                                  types.StringType,
			"SettingName":                           types.StringType,
		},
	}
}

type AutomaticClusterUpdateSetting struct {
	AutomaticClusterUpdateWorkspace types.List `tfsdk:"automatic_cluster_update_workspace" tf:"object"`
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

func (newState *AutomaticClusterUpdateSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutomaticClusterUpdateSetting) {
}

func (newState *AutomaticClusterUpdateSetting) SyncEffectiveFieldsDuringRead(existingState AutomaticClusterUpdateSetting) {
}

func (a AutomaticClusterUpdateSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AutomaticClusterUpdateWorkspace": reflect.TypeOf(ClusterAutoRestartMessage{}),
	}
}

func (a AutomaticClusterUpdateSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AutomaticClusterUpdateWorkspace": ClusterAutoRestartMessage{}.ToAttrType(ctx),
			"Etag":                            types.StringType,
			"SettingName":                     types.StringType,
		},
	}
}

type BooleanMessage struct {
	Value types.Bool `tfsdk:"value" tf:"optional"`
}

func (newState *BooleanMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan BooleanMessage) {
}

func (newState *BooleanMessage) SyncEffectiveFieldsDuringRead(existingState BooleanMessage) {
}

func (a BooleanMessage) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a BooleanMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Value": types.BoolType,
		},
	}
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
	EnablementDetails types.List `tfsdk:"enablement_details" tf:"optional,object"`

	MaintenanceWindow types.List `tfsdk:"maintenance_window" tf:"optional,object"`

	RestartEvenIfNoUpdatesAvailable types.Bool `tfsdk:"restart_even_if_no_updates_available" tf:"optional"`
}

func (newState *ClusterAutoRestartMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessage) {
}

func (newState *ClusterAutoRestartMessage) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessage) {
}

func (a ClusterAutoRestartMessage) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"EnablementDetails": reflect.TypeOf(ClusterAutoRestartMessageEnablementDetails{}),
		"MaintenanceWindow": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindow{}),
	}
}

func (a ClusterAutoRestartMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CanToggle":                       types.BoolType,
			"Enabled":                         types.BoolType,
			"EnablementDetails":               ClusterAutoRestartMessageEnablementDetails{}.ToAttrType(ctx),
			"MaintenanceWindow":               ClusterAutoRestartMessageMaintenanceWindow{}.ToAttrType(ctx),
			"RestartEvenIfNoUpdatesAvailable": types.BoolType,
		},
	}
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

func (newState *ClusterAutoRestartMessageEnablementDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageEnablementDetails) {
}

func (newState *ClusterAutoRestartMessageEnablementDetails) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageEnablementDetails) {
}

func (a ClusterAutoRestartMessageEnablementDetails) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ClusterAutoRestartMessageEnablementDetails) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ForcedForComplianceMode":           types.BoolType,
			"UnavailableForDisabledEntitlement": types.BoolType,
			"UnavailableForNonEnterpriseTier":   types.BoolType,
		},
	}
}

type ClusterAutoRestartMessageMaintenanceWindow struct {
	WeekDayBasedSchedule types.List `tfsdk:"week_day_based_schedule" tf:"optional,object"`
}

func (newState *ClusterAutoRestartMessageMaintenanceWindow) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageMaintenanceWindow) {
}

func (newState *ClusterAutoRestartMessageMaintenanceWindow) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageMaintenanceWindow) {
}

func (a ClusterAutoRestartMessageMaintenanceWindow) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"WeekDayBasedSchedule": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}),
	}
}

func (a ClusterAutoRestartMessageMaintenanceWindow) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"WeekDayBasedSchedule": ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}.ToAttrType(ctx),
		},
	}
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	DayOfWeek types.String `tfsdk:"day_of_week" tf:"optional"`

	Frequency types.String `tfsdk:"frequency" tf:"optional"`

	WindowStartTime types.List `tfsdk:"window_start_time" tf:"optional,object"`
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
}

func (a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"WindowStartTime": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}),
	}
}

func (a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DayOfWeek":       types.StringType,
			"Frequency":       types.StringType,
			"WindowStartTime": ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}.ToAttrType(ctx),
		},
	}
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	Hours types.Int64 `tfsdk:"hours" tf:"optional"`

	Minutes types.Int64 `tfsdk:"minutes" tf:"optional"`
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
}

func (a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Hours":   types.Int64Type,
			"Minutes": types.Int64Type,
		},
	}
}

// SHIELD feature: CSP
type ComplianceSecurityProfile struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	ComplianceStandards types.List `tfsdk:"compliance_standards" tf:"optional"`

	IsEnabled types.Bool `tfsdk:"is_enabled" tf:"optional"`
}

func (newState *ComplianceSecurityProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplianceSecurityProfile) {
}

func (newState *ComplianceSecurityProfile) SyncEffectiveFieldsDuringRead(existingState ComplianceSecurityProfile) {
}

func (a ComplianceSecurityProfile) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ComplianceStandards": reflect.TypeOf(types.StringType),
	}
}

func (a ComplianceSecurityProfile) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ComplianceStandards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"IsEnabled": types.BoolType,
		},
	}
}

type ComplianceSecurityProfileSetting struct {
	// SHIELD feature: CSP
	ComplianceSecurityProfileWorkspace types.List `tfsdk:"compliance_security_profile_workspace" tf:"object"`
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

func (newState *ComplianceSecurityProfileSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplianceSecurityProfileSetting) {
}

func (newState *ComplianceSecurityProfileSetting) SyncEffectiveFieldsDuringRead(existingState ComplianceSecurityProfileSetting) {
}

func (a ComplianceSecurityProfileSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ComplianceSecurityProfileWorkspace": reflect.TypeOf(ComplianceSecurityProfile{}),
	}
}

func (a ComplianceSecurityProfileSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ComplianceSecurityProfileWorkspace": ComplianceSecurityProfile{}.ToAttrType(ctx),
			"Etag":                               types.StringType,
			"SettingName":                        types.StringType,
		},
	}
}

type Config struct {
	Email types.List `tfsdk:"email" tf:"optional,object"`

	GenericWebhook types.List `tfsdk:"generic_webhook" tf:"optional,object"`

	MicrosoftTeams types.List `tfsdk:"microsoft_teams" tf:"optional,object"`

	Pagerduty types.List `tfsdk:"pagerduty" tf:"optional,object"`

	Slack types.List `tfsdk:"slack" tf:"optional,object"`
}

func (newState *Config) SyncEffectiveFieldsDuringCreateOrUpdate(plan Config) {
}

func (newState *Config) SyncEffectiveFieldsDuringRead(existingState Config) {
}

func (a Config) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Email":          reflect.TypeOf(EmailConfig{}),
		"GenericWebhook": reflect.TypeOf(GenericWebhookConfig{}),
		"MicrosoftTeams": reflect.TypeOf(MicrosoftTeamsConfig{}),
		"Pagerduty":      reflect.TypeOf(PagerdutyConfig{}),
		"Slack":          reflect.TypeOf(SlackConfig{}),
	}
}

func (a Config) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Email":          EmailConfig{}.ToAttrType(ctx),
			"GenericWebhook": GenericWebhookConfig{}.ToAttrType(ctx),
			"MicrosoftTeams": MicrosoftTeamsConfig{}.ToAttrType(ctx),
			"Pagerduty":      PagerdutyConfig{}.ToAttrType(ctx),
			"Slack":          SlackConfig{}.ToAttrType(ctx),
		},
	}
}

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {
	IpAddresses types.List `tfsdk:"ip_addresses" tf:"optional"`
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

func (newState *CreateIpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateIpAccessList) {
}

func (newState *CreateIpAccessList) SyncEffectiveFieldsDuringRead(existingState CreateIpAccessList) {
}

func (a CreateIpAccessList) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"IpAddresses": reflect.TypeOf(types.StringType),
	}
}

func (a CreateIpAccessList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAddresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Label":    types.StringType,
			"ListType": types.StringType,
		},
	}
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList types.List `tfsdk:"ip_access_list" tf:"optional,object"`
}

func (newState *CreateIpAccessListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateIpAccessListResponse) {
}

func (newState *CreateIpAccessListResponse) SyncEffectiveFieldsDuringRead(existingState CreateIpAccessListResponse) {
}

func (a CreateIpAccessListResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"IpAccessList": reflect.TypeOf(IpAccessListInfo{}),
	}
}

func (a CreateIpAccessListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAccessList": IpAccessListInfo{}.ToAttrType(ctx),
		},
	}
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

func (newState *CreateNetworkConnectivityConfigRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateNetworkConnectivityConfigRequest) {
}

func (newState *CreateNetworkConnectivityConfigRequest) SyncEffectiveFieldsDuringRead(existingState CreateNetworkConnectivityConfigRequest) {
}

func (a CreateNetworkConnectivityConfigRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateNetworkConnectivityConfigRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name":   types.StringType,
			"Region": types.StringType,
		},
	}
}

type CreateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config types.List `tfsdk:"config" tf:"optional,object"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
}

func (newState *CreateNotificationDestinationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateNotificationDestinationRequest) {
}

func (newState *CreateNotificationDestinationRequest) SyncEffectiveFieldsDuringRead(existingState CreateNotificationDestinationRequest) {
}

func (a CreateNotificationDestinationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Config": reflect.TypeOf(Config{}),
	}
}

func (a CreateNotificationDestinationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Config":      Config{}.ToAttrType(ctx),
			"DisplayName": types.StringType,
		},
	}
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

func (newState *CreateOboTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateOboTokenRequest) {
}

func (newState *CreateOboTokenRequest) SyncEffectiveFieldsDuringRead(existingState CreateOboTokenRequest) {
}

func (a CreateOboTokenRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateOboTokenRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ApplicationId":   types.StringType,
			"Comment":         types.StringType,
			"LifetimeSeconds": types.Int64Type,
		},
	}
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {
	TokenInfo types.List `tfsdk:"token_info" tf:"optional,object"`
	// Value of the token.
	TokenValue types.String `tfsdk:"token_value" tf:"optional"`
}

func (newState *CreateOboTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateOboTokenResponse) {
}

func (newState *CreateOboTokenResponse) SyncEffectiveFieldsDuringRead(existingState CreateOboTokenResponse) {
}

func (a CreateOboTokenResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"TokenInfo": reflect.TypeOf(TokenInfo{}),
	}
}

func (a CreateOboTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"TokenInfo":  TokenInfo{}.ToAttrType(ctx),
			"TokenValue": types.StringType,
		},
	}
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

func (newState *CreatePrivateEndpointRuleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePrivateEndpointRuleRequest) {
}

func (newState *CreatePrivateEndpointRuleRequest) SyncEffectiveFieldsDuringRead(existingState CreatePrivateEndpointRuleRequest) {
}

func (a CreatePrivateEndpointRuleRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreatePrivateEndpointRuleRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"GroupId":                     types.StringType,
			"NetworkConnectivityConfigId": types.StringType,
			"ResourceId":                  types.StringType,
		},
	}
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds" tf:"optional"`
}

func (newState *CreateTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTokenRequest) {
}

func (newState *CreateTokenRequest) SyncEffectiveFieldsDuringRead(existingState CreateTokenRequest) {
}

func (a CreateTokenRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateTokenRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment":         types.StringType,
			"LifetimeSeconds": types.Int64Type,
		},
	}
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo types.List `tfsdk:"token_info" tf:"optional,object"`
	// The value of the new token.
	TokenValue types.String `tfsdk:"token_value" tf:"optional"`
}

func (newState *CreateTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTokenResponse) {
}

func (newState *CreateTokenResponse) SyncEffectiveFieldsDuringRead(existingState CreateTokenResponse) {
}

func (a CreateTokenResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"TokenInfo": reflect.TypeOf(PublicTokenInfo{}),
	}
}

func (a CreateTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"TokenInfo":  PublicTokenInfo{}.ToAttrType(ctx),
			"TokenValue": types.StringType,
		},
	}
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards types.List `tfsdk:"compliance_standards" tf:"optional"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced types.Bool `tfsdk:"is_enforced" tf:"optional"`
}

func (newState *CspEnablementAccount) SyncEffectiveFieldsDuringCreateOrUpdate(plan CspEnablementAccount) {
}

func (newState *CspEnablementAccount) SyncEffectiveFieldsDuringRead(existingState CspEnablementAccount) {
}

func (a CspEnablementAccount) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ComplianceStandards": reflect.TypeOf(types.StringType),
	}
}

func (a CspEnablementAccount) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ComplianceStandards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"IsEnforced": types.BoolType,
		},
	}
}

type CspEnablementAccountSetting struct {
	// Account level policy for CSP
	CspEnablementAccount types.List `tfsdk:"csp_enablement_account" tf:"object"`
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

func (newState *CspEnablementAccountSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan CspEnablementAccountSetting) {
}

func (newState *CspEnablementAccountSetting) SyncEffectiveFieldsDuringRead(existingState CspEnablementAccountSetting) {
}

func (a CspEnablementAccountSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"CspEnablementAccount": reflect.TypeOf(CspEnablementAccount{}),
	}
}

func (a CspEnablementAccountSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CspEnablementAccount": CspEnablementAccount{}.ToAttrType(ctx),
			"Etag":                 types.StringType,
			"SettingName":          types.StringType,
		},
	}
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

	Namespace types.List `tfsdk:"namespace" tf:"object"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

func (newState *DefaultNamespaceSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan DefaultNamespaceSetting) {
}

func (newState *DefaultNamespaceSetting) SyncEffectiveFieldsDuringRead(existingState DefaultNamespaceSetting) {
}

func (a DefaultNamespaceSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Namespace": reflect.TypeOf(StringMessage{}),
	}
}

func (a DefaultNamespaceSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag":        types.StringType,
			"Namespace":   StringMessage{}.ToAttrType(ctx),
			"SettingName": types.StringType,
		},
	}
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (newState *DeleteAccountIpAccessListRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountIpAccessListRequest) {
}

func (newState *DeleteAccountIpAccessListRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountIpAccessListRequest) {
}

func (a DeleteAccountIpAccessListRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteAccountIpAccessListRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAccessListId": types.StringType,
		},
	}
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

func (newState *DeleteDefaultNamespaceSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDefaultNamespaceSettingRequest) {
}

func (newState *DeleteDefaultNamespaceSettingRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDefaultNamespaceSettingRequest) {
}

func (a DeleteDefaultNamespaceSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteDefaultNamespaceSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *DeleteDefaultNamespaceSettingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDefaultNamespaceSettingResponse) {
}

func (newState *DeleteDefaultNamespaceSettingResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDefaultNamespaceSettingResponse) {
}

func (a DeleteDefaultNamespaceSettingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteDefaultNamespaceSettingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *DeleteDisableLegacyAccessRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyAccessRequest) {
}

func (newState *DeleteDisableLegacyAccessRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyAccessRequest) {
}

func (a DeleteDisableLegacyAccessRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteDisableLegacyAccessRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *DeleteDisableLegacyAccessResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyAccessResponse) {
}

func (newState *DeleteDisableLegacyAccessResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyAccessResponse) {
}

func (a DeleteDisableLegacyAccessResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteDisableLegacyAccessResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
}

// Delete the disable legacy DBFS setting
type DeleteDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (newState *DeleteDisableLegacyDbfsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyDbfsRequest) {
}

func (newState *DeleteDisableLegacyDbfsRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyDbfsRequest) {
}

func (a DeleteDisableLegacyDbfsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteDisableLegacyDbfsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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
	Etag types.String `tfsdk:"etag" tf:""`
}

func (newState *DeleteDisableLegacyDbfsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyDbfsResponse) {
}

func (newState *DeleteDisableLegacyDbfsResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyDbfsResponse) {
}

func (a DeleteDisableLegacyDbfsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteDisableLegacyDbfsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *DeleteDisableLegacyFeaturesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyFeaturesRequest) {
}

func (newState *DeleteDisableLegacyFeaturesRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyFeaturesRequest) {
}

func (a DeleteDisableLegacyFeaturesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteDisableLegacyFeaturesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *DeleteDisableLegacyFeaturesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyFeaturesResponse) {
}

func (newState *DeleteDisableLegacyFeaturesResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyFeaturesResponse) {
}

func (a DeleteDisableLegacyFeaturesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteDisableLegacyFeaturesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (newState *DeleteIpAccessListRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteIpAccessListRequest) {
}

func (newState *DeleteIpAccessListRequest) SyncEffectiveFieldsDuringRead(existingState DeleteIpAccessListRequest) {
}

func (a DeleteIpAccessListRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteIpAccessListRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAccessListId": types.StringType,
		},
	}
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

func (newState *DeleteNetworkConnectivityConfigurationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteNetworkConnectivityConfigurationRequest) {
}

func (newState *DeleteNetworkConnectivityConfigurationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteNetworkConnectivityConfigurationRequest) {
}

func (a DeleteNetworkConnectivityConfigurationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteNetworkConnectivityConfigurationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NetworkConnectivityConfigId": types.StringType,
		},
	}
}

type DeleteNetworkConnectivityConfigurationResponse struct {
}

func (newState *DeleteNetworkConnectivityConfigurationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteNetworkConnectivityConfigurationResponse) {
}

func (newState *DeleteNetworkConnectivityConfigurationResponse) SyncEffectiveFieldsDuringRead(existingState DeleteNetworkConnectivityConfigurationResponse) {
}

func (a DeleteNetworkConnectivityConfigurationResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteNetworkConnectivityConfigurationResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a notification destination
type DeleteNotificationDestinationRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteNotificationDestinationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteNotificationDestinationRequest) {
}

func (newState *DeleteNotificationDestinationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteNotificationDestinationRequest) {
}

func (a DeleteNotificationDestinationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteNotificationDestinationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
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

func (newState *DeletePersonalComputeSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePersonalComputeSettingRequest) {
}

func (newState *DeletePersonalComputeSettingRequest) SyncEffectiveFieldsDuringRead(existingState DeletePersonalComputeSettingRequest) {
}

func (a DeletePersonalComputeSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeletePersonalComputeSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *DeletePersonalComputeSettingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePersonalComputeSettingResponse) {
}

func (newState *DeletePersonalComputeSettingResponse) SyncEffectiveFieldsDuringRead(existingState DeletePersonalComputeSettingResponse) {
}

func (a DeletePersonalComputeSettingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeletePersonalComputeSettingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

func (newState *DeletePrivateEndpointRuleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePrivateEndpointRuleRequest) {
}

func (newState *DeletePrivateEndpointRuleRequest) SyncEffectiveFieldsDuringRead(existingState DeletePrivateEndpointRuleRequest) {
}

func (a DeletePrivateEndpointRuleRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeletePrivateEndpointRuleRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NetworkConnectivityConfigId": types.StringType,
			"PrivateEndpointRuleId":       types.StringType,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

func (a DeleteResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

func (newState *DeleteRestrictWorkspaceAdminsSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRestrictWorkspaceAdminsSettingRequest) {
}

func (newState *DeleteRestrictWorkspaceAdminsSettingRequest) SyncEffectiveFieldsDuringRead(existingState DeleteRestrictWorkspaceAdminsSettingRequest) {
}

func (a DeleteRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteRestrictWorkspaceAdminsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *DeleteRestrictWorkspaceAdminsSettingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRestrictWorkspaceAdminsSettingResponse) {
}

func (newState *DeleteRestrictWorkspaceAdminsSettingResponse) SyncEffectiveFieldsDuringRead(existingState DeleteRestrictWorkspaceAdminsSettingResponse) {
}

func (a DeleteRestrictWorkspaceAdminsSettingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteRestrictWorkspaceAdminsSettingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to revoke.
	TokenId types.String `tfsdk:"-"`
}

func (newState *DeleteTokenManagementRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTokenManagementRequest) {
}

func (newState *DeleteTokenManagementRequest) SyncEffectiveFieldsDuringRead(existingState DeleteTokenManagementRequest) {
}

func (a DeleteTokenManagementRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteTokenManagementRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"TokenId": types.StringType,
		},
	}
}

type DisableLegacyAccess struct {
	DisableLegacyAccess types.List `tfsdk:"disable_legacy_access" tf:"object"`
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

func (newState *DisableLegacyAccess) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableLegacyAccess) {
}

func (newState *DisableLegacyAccess) SyncEffectiveFieldsDuringRead(existingState DisableLegacyAccess) {
}

func (a DisableLegacyAccess) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DisableLegacyAccess": reflect.TypeOf(BooleanMessage{}),
	}
}

func (a DisableLegacyAccess) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DisableLegacyAccess": BooleanMessage{}.ToAttrType(ctx),
			"Etag":                types.StringType,
			"SettingName":         types.StringType,
		},
	}
}

type DisableLegacyDbfs struct {
	DisableLegacyDbfs types.List `tfsdk:"disable_legacy_dbfs" tf:"object"`
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

func (newState *DisableLegacyDbfs) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableLegacyDbfs) {
}

func (newState *DisableLegacyDbfs) SyncEffectiveFieldsDuringRead(existingState DisableLegacyDbfs) {
}

func (a DisableLegacyDbfs) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DisableLegacyDbfs": reflect.TypeOf(BooleanMessage{}),
	}
}

func (a DisableLegacyDbfs) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DisableLegacyDbfs": BooleanMessage{}.ToAttrType(ctx),
			"Etag":              types.StringType,
			"SettingName":       types.StringType,
		},
	}
}

type DisableLegacyFeatures struct {
	DisableLegacyFeatures types.List `tfsdk:"disable_legacy_features" tf:"object"`
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

func (newState *DisableLegacyFeatures) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableLegacyFeatures) {
}

func (newState *DisableLegacyFeatures) SyncEffectiveFieldsDuringRead(existingState DisableLegacyFeatures) {
}

func (a DisableLegacyFeatures) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DisableLegacyFeatures": reflect.TypeOf(BooleanMessage{}),
	}
}

func (a DisableLegacyFeatures) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DisableLegacyFeatures": BooleanMessage{}.ToAttrType(ctx),
			"Etag":                  types.StringType,
			"SettingName":           types.StringType,
		},
	}
}

type EmailConfig struct {
	// Email addresses to notify.
	Addresses types.List `tfsdk:"addresses" tf:"optional"`
}

func (newState *EmailConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmailConfig) {
}

func (newState *EmailConfig) SyncEffectiveFieldsDuringRead(existingState EmailConfig) {
}

func (a EmailConfig) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Addresses": reflect.TypeOf(types.StringType),
	}
}

func (a EmailConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type Empty struct {
}

func (newState *Empty) SyncEffectiveFieldsDuringCreateOrUpdate(plan Empty) {
}

func (newState *Empty) SyncEffectiveFieldsDuringRead(existingState Empty) {
}

func (a Empty) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Empty) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {
	IsEnabled types.Bool `tfsdk:"is_enabled" tf:"optional"`
}

func (newState *EnhancedSecurityMonitoring) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnhancedSecurityMonitoring) {
}

func (newState *EnhancedSecurityMonitoring) SyncEffectiveFieldsDuringRead(existingState EnhancedSecurityMonitoring) {
}

func (a EnhancedSecurityMonitoring) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a EnhancedSecurityMonitoring) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IsEnabled": types.BoolType,
		},
	}
}

type EnhancedSecurityMonitoringSetting struct {
	// SHIELD feature: ESM
	EnhancedSecurityMonitoringWorkspace types.List `tfsdk:"enhanced_security_monitoring_workspace" tf:"object"`
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

func (newState *EnhancedSecurityMonitoringSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnhancedSecurityMonitoringSetting) {
}

func (newState *EnhancedSecurityMonitoringSetting) SyncEffectiveFieldsDuringRead(existingState EnhancedSecurityMonitoringSetting) {
}

func (a EnhancedSecurityMonitoringSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"EnhancedSecurityMonitoringWorkspace": reflect.TypeOf(EnhancedSecurityMonitoring{}),
	}
}

func (a EnhancedSecurityMonitoringSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EnhancedSecurityMonitoringWorkspace": EnhancedSecurityMonitoring{}.ToAttrType(ctx),
			"Etag":                                types.StringType,
			"SettingName":                         types.StringType,
		},
	}
}

// Account level policy for ESM
type EsmEnablementAccount struct {
	IsEnforced types.Bool `tfsdk:"is_enforced" tf:"optional"`
}

func (newState *EsmEnablementAccount) SyncEffectiveFieldsDuringCreateOrUpdate(plan EsmEnablementAccount) {
}

func (newState *EsmEnablementAccount) SyncEffectiveFieldsDuringRead(existingState EsmEnablementAccount) {
}

func (a EsmEnablementAccount) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a EsmEnablementAccount) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IsEnforced": types.BoolType,
		},
	}
}

type EsmEnablementAccountSetting struct {
	// Account level policy for ESM
	EsmEnablementAccount types.List `tfsdk:"esm_enablement_account" tf:"object"`
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

func (newState *EsmEnablementAccountSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan EsmEnablementAccountSetting) {
}

func (newState *EsmEnablementAccountSetting) SyncEffectiveFieldsDuringRead(existingState EsmEnablementAccountSetting) {
}

func (a EsmEnablementAccountSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"EsmEnablementAccount": reflect.TypeOf(EsmEnablementAccount{}),
	}
}

func (a EsmEnablementAccountSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EsmEnablementAccount": EsmEnablementAccount{}.ToAttrType(ctx),
			"Etag":                 types.StringType,
			"SettingName":          types.StringType,
		},
	}
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
	Scopes types.List `tfsdk:"scopes" tf:"optional"`
	// The type of this exchange token
	TokenType types.String `tfsdk:"tokenType" tf:"optional"`
}

func (newState *ExchangeToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeToken) {
}

func (newState *ExchangeToken) SyncEffectiveFieldsDuringRead(existingState ExchangeToken) {
}

func (a ExchangeToken) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Scopes": reflect.TypeOf(types.StringType),
	}
}

func (a ExchangeToken) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Credential":        types.StringType,
			"CredentialEolTime": types.Int64Type,
			"OwnerId":           types.Int64Type,
			"Scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"TokenType": types.StringType,
		},
	}
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	PartitionId types.List `tfsdk:"partitionId" tf:"object"`
	// Array of scopes for the token request.
	Scopes types.List `tfsdk:"scopes" tf:""`
	// A list of token types being requested
	TokenType types.List `tfsdk:"tokenType" tf:""`
}

func (newState *ExchangeTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeTokenRequest) {
}

func (newState *ExchangeTokenRequest) SyncEffectiveFieldsDuringRead(existingState ExchangeTokenRequest) {
}

func (a ExchangeTokenRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PartitionId": reflect.TypeOf(PartitionId{}),
		"Scopes":      reflect.TypeOf(types.StringType),
		"TokenType":   reflect.TypeOf(types.StringType),
	}
}

func (a ExchangeTokenRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PartitionId": PartitionId{}.ToAttrType(ctx),
			"Scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"TokenType": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {
	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *ExchangeTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeTokenResponse) {
}

func (newState *ExchangeTokenResponse) SyncEffectiveFieldsDuringRead(existingState ExchangeTokenResponse) {
}

func (a ExchangeTokenResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Values": reflect.TypeOf(ExchangeToken{}),
	}
}

func (a ExchangeTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Values": basetypes.ListType{
				ElemType: ExchangeToken{}.ToAttrType(ctx),
			},
		},
	}
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList types.List `tfsdk:"ip_access_list" tf:"optional,object"`
}

func (newState *FetchIpAccessListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan FetchIpAccessListResponse) {
}

func (newState *FetchIpAccessListResponse) SyncEffectiveFieldsDuringRead(existingState FetchIpAccessListResponse) {
}

func (a FetchIpAccessListResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"IpAccessList": reflect.TypeOf(IpAccessListInfo{}),
	}
}

func (a FetchIpAccessListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAccessList": IpAccessListInfo{}.ToAttrType(ctx),
		},
	}
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

func (newState *GenericWebhookConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenericWebhookConfig) {
}

func (newState *GenericWebhookConfig) SyncEffectiveFieldsDuringRead(existingState GenericWebhookConfig) {
}

func (a GenericWebhookConfig) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GenericWebhookConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Password":    types.StringType,
			"PasswordSet": types.BoolType,
			"Url":         types.StringType,
			"UrlSet":      types.BoolType,
			"Username":    types.StringType,
			"UsernameSet": types.BoolType,
		},
	}
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (newState *GetAccountIpAccessListRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountIpAccessListRequest) {
}

func (newState *GetAccountIpAccessListRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountIpAccessListRequest) {
}

func (a GetAccountIpAccessListRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetAccountIpAccessListRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAccessListId": types.StringType,
		},
	}
}

// Retrieve the AI/BI dashboard embedding access policy
type GetAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (newState *GetAibiDashboardEmbeddingAccessPolicySettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (newState *GetAibiDashboardEmbeddingAccessPolicySettingRequest) SyncEffectiveFieldsDuringRead(existingState GetAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (a GetAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetAibiDashboardEmbeddingAccessPolicySettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
}

// Retrieve the list of domains approved to host embedded AI/BI dashboards
type GetAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (newState *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (newState *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncEffectiveFieldsDuringRead(existingState GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (a GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *GetAutomaticClusterUpdateSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAutomaticClusterUpdateSettingRequest) {
}

func (newState *GetAutomaticClusterUpdateSettingRequest) SyncEffectiveFieldsDuringRead(existingState GetAutomaticClusterUpdateSettingRequest) {
}

func (a GetAutomaticClusterUpdateSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetAutomaticClusterUpdateSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *GetComplianceSecurityProfileSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetComplianceSecurityProfileSettingRequest) {
}

func (newState *GetComplianceSecurityProfileSettingRequest) SyncEffectiveFieldsDuringRead(existingState GetComplianceSecurityProfileSettingRequest) {
}

func (a GetComplianceSecurityProfileSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetComplianceSecurityProfileSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *GetCspEnablementAccountSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCspEnablementAccountSettingRequest) {
}

func (newState *GetCspEnablementAccountSettingRequest) SyncEffectiveFieldsDuringRead(existingState GetCspEnablementAccountSettingRequest) {
}

func (a GetCspEnablementAccountSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetCspEnablementAccountSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *GetDefaultNamespaceSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDefaultNamespaceSettingRequest) {
}

func (newState *GetDefaultNamespaceSettingRequest) SyncEffectiveFieldsDuringRead(existingState GetDefaultNamespaceSettingRequest) {
}

func (a GetDefaultNamespaceSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetDefaultNamespaceSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *GetDisableLegacyAccessRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDisableLegacyAccessRequest) {
}

func (newState *GetDisableLegacyAccessRequest) SyncEffectiveFieldsDuringRead(existingState GetDisableLegacyAccessRequest) {
}

func (a GetDisableLegacyAccessRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetDisableLegacyAccessRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
}

// Get the disable legacy DBFS setting
type GetDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (newState *GetDisableLegacyDbfsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDisableLegacyDbfsRequest) {
}

func (newState *GetDisableLegacyDbfsRequest) SyncEffectiveFieldsDuringRead(existingState GetDisableLegacyDbfsRequest) {
}

func (a GetDisableLegacyDbfsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetDisableLegacyDbfsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *GetDisableLegacyFeaturesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDisableLegacyFeaturesRequest) {
}

func (newState *GetDisableLegacyFeaturesRequest) SyncEffectiveFieldsDuringRead(existingState GetDisableLegacyFeaturesRequest) {
}

func (a GetDisableLegacyFeaturesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetDisableLegacyFeaturesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *GetEnhancedSecurityMonitoringSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEnhancedSecurityMonitoringSettingRequest) {
}

func (newState *GetEnhancedSecurityMonitoringSettingRequest) SyncEffectiveFieldsDuringRead(existingState GetEnhancedSecurityMonitoringSettingRequest) {
}

func (a GetEnhancedSecurityMonitoringSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetEnhancedSecurityMonitoringSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
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

func (newState *GetEsmEnablementAccountSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEsmEnablementAccountSettingRequest) {
}

func (newState *GetEsmEnablementAccountSettingRequest) SyncEffectiveFieldsDuringRead(existingState GetEsmEnablementAccountSettingRequest) {
}

func (a GetEsmEnablementAccountSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetEsmEnablementAccountSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (newState *GetIpAccessListRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetIpAccessListRequest) {
}

func (newState *GetIpAccessListRequest) SyncEffectiveFieldsDuringRead(existingState GetIpAccessListRequest) {
}

func (a GetIpAccessListRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetIpAccessListRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAccessListId": types.StringType,
		},
	}
}

type GetIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList types.List `tfsdk:"ip_access_list" tf:"optional,object"`
}

func (newState *GetIpAccessListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetIpAccessListResponse) {
}

func (newState *GetIpAccessListResponse) SyncEffectiveFieldsDuringRead(existingState GetIpAccessListResponse) {
}

func (a GetIpAccessListResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"IpAccessList": reflect.TypeOf(IpAccessListInfo{}),
	}
}

func (a GetIpAccessListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAccessList": IpAccessListInfo{}.ToAttrType(ctx),
		},
	}
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists" tf:"optional"`
}

func (newState *GetIpAccessListsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetIpAccessListsResponse) {
}

func (newState *GetIpAccessListsResponse) SyncEffectiveFieldsDuringRead(existingState GetIpAccessListsResponse) {
}

func (a GetIpAccessListsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"IpAccessLists": reflect.TypeOf(IpAccessListInfo{}),
	}
}

func (a GetIpAccessListsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAccessLists": basetypes.ListType{
				ElemType: IpAccessListInfo{}.ToAttrType(ctx),
			},
		},
	}
}

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

func (newState *GetNetworkConnectivityConfigurationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetNetworkConnectivityConfigurationRequest) {
}

func (newState *GetNetworkConnectivityConfigurationRequest) SyncEffectiveFieldsDuringRead(existingState GetNetworkConnectivityConfigurationRequest) {
}

func (a GetNetworkConnectivityConfigurationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetNetworkConnectivityConfigurationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NetworkConnectivityConfigId": types.StringType,
		},
	}
}

// Get a notification destination
type GetNotificationDestinationRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetNotificationDestinationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetNotificationDestinationRequest) {
}

func (newState *GetNotificationDestinationRequest) SyncEffectiveFieldsDuringRead(existingState GetNotificationDestinationRequest) {
}

func (a GetNotificationDestinationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetNotificationDestinationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
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

func (newState *GetPersonalComputeSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPersonalComputeSettingRequest) {
}

func (newState *GetPersonalComputeSettingRequest) SyncEffectiveFieldsDuringRead(existingState GetPersonalComputeSettingRequest) {
}

func (a GetPersonalComputeSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetPersonalComputeSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
}

// Get a private endpoint rule
type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

func (newState *GetPrivateEndpointRuleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPrivateEndpointRuleRequest) {
}

func (newState *GetPrivateEndpointRuleRequest) SyncEffectiveFieldsDuringRead(existingState GetPrivateEndpointRuleRequest) {
}

func (a GetPrivateEndpointRuleRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetPrivateEndpointRuleRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NetworkConnectivityConfigId": types.StringType,
			"PrivateEndpointRuleId":       types.StringType,
		},
	}
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

func (newState *GetRestrictWorkspaceAdminsSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRestrictWorkspaceAdminsSettingRequest) {
}

func (newState *GetRestrictWorkspaceAdminsSettingRequest) SyncEffectiveFieldsDuringRead(existingState GetRestrictWorkspaceAdminsSettingRequest) {
}

func (a GetRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetRestrictWorkspaceAdminsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag": types.StringType,
		},
	}
}

// Check configuration status
type GetStatusRequest struct {
	Keys types.String `tfsdk:"-"`
}

func (newState *GetStatusRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStatusRequest) {
}

func (newState *GetStatusRequest) SyncEffectiveFieldsDuringRead(existingState GetStatusRequest) {
}

func (a GetStatusRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetStatusRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Keys": types.StringType,
		},
	}
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId types.String `tfsdk:"-"`
}

func (newState *GetTokenManagementRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTokenManagementRequest) {
}

func (newState *GetTokenManagementRequest) SyncEffectiveFieldsDuringRead(existingState GetTokenManagementRequest) {
}

func (a GetTokenManagementRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetTokenManagementRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"TokenId": types.StringType,
		},
	}
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetTokenPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTokenPermissionLevelsResponse) {
}

func (newState *GetTokenPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetTokenPermissionLevelsResponse) {
}

func (a GetTokenPermissionLevelsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PermissionLevels": reflect.TypeOf(TokenPermissionsDescription{}),
	}
}

func (a GetTokenPermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PermissionLevels": basetypes.ListType{
				ElemType: TokenPermissionsDescription{}.ToAttrType(ctx),
			},
		},
	}
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {
	TokenInfo types.List `tfsdk:"token_info" tf:"optional,object"`
}

func (newState *GetTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTokenResponse) {
}

func (newState *GetTokenResponse) SyncEffectiveFieldsDuringRead(existingState GetTokenResponse) {
}

func (a GetTokenResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"TokenInfo": reflect.TypeOf(TokenInfo{}),
	}
}

func (a GetTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"TokenInfo": TokenInfo{}.ToAttrType(ctx),
		},
	}
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

	IpAddresses types.List `tfsdk:"ip_addresses" tf:"optional"`
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

func (newState *IpAccessListInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan IpAccessListInfo) {
}

func (newState *IpAccessListInfo) SyncEffectiveFieldsDuringRead(existingState IpAccessListInfo) {
}

func (a IpAccessListInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"IpAddresses": reflect.TypeOf(types.StringType),
	}
}

func (a IpAccessListInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AddressCount": types.Int64Type,
			"CreatedAt":    types.Int64Type,
			"CreatedBy":    types.Int64Type,
			"Enabled":      types.BoolType,
			"IpAddresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Label":     types.StringType,
			"ListId":    types.StringType,
			"ListType":  types.StringType,
			"UpdatedAt": types.Int64Type,
			"UpdatedBy": types.Int64Type,
		},
	}
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists" tf:"optional"`
}

func (newState *ListIpAccessListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListIpAccessListResponse) {
}

func (newState *ListIpAccessListResponse) SyncEffectiveFieldsDuringRead(existingState ListIpAccessListResponse) {
}

func (a ListIpAccessListResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"IpAccessLists": reflect.TypeOf(IpAccessListInfo{}),
	}
}

func (a ListIpAccessListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IpAccessLists": basetypes.ListType{
				ElemType: IpAccessListInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type ListNccAzurePrivateEndpointRulesResponse struct {
	Items types.List `tfsdk:"items" tf:"optional"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListNccAzurePrivateEndpointRulesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNccAzurePrivateEndpointRulesResponse) {
}

func (newState *ListNccAzurePrivateEndpointRulesResponse) SyncEffectiveFieldsDuringRead(existingState ListNccAzurePrivateEndpointRulesResponse) {
}

func (a ListNccAzurePrivateEndpointRulesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Items": reflect.TypeOf(NccAzurePrivateEndpointRule{}),
	}
}

func (a ListNccAzurePrivateEndpointRulesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Items": basetypes.ListType{
				ElemType: NccAzurePrivateEndpointRule{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListNetworkConnectivityConfigurationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNetworkConnectivityConfigurationsRequest) {
}

func (newState *ListNetworkConnectivityConfigurationsRequest) SyncEffectiveFieldsDuringRead(existingState ListNetworkConnectivityConfigurationsRequest) {
}

func (a ListNetworkConnectivityConfigurationsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListNetworkConnectivityConfigurationsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PageToken": types.StringType,
		},
	}
}

type ListNetworkConnectivityConfigurationsResponse struct {
	Items types.List `tfsdk:"items" tf:"optional"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListNetworkConnectivityConfigurationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNetworkConnectivityConfigurationsResponse) {
}

func (newState *ListNetworkConnectivityConfigurationsResponse) SyncEffectiveFieldsDuringRead(existingState ListNetworkConnectivityConfigurationsResponse) {
}

func (a ListNetworkConnectivityConfigurationsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Items": reflect.TypeOf(NetworkConnectivityConfiguration{}),
	}
}

func (a ListNetworkConnectivityConfigurationsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Items": basetypes.ListType{
				ElemType: NetworkConnectivityConfiguration{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List notification destinations
type ListNotificationDestinationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListNotificationDestinationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNotificationDestinationsRequest) {
}

func (newState *ListNotificationDestinationsRequest) SyncEffectiveFieldsDuringRead(existingState ListNotificationDestinationsRequest) {
}

func (a ListNotificationDestinationsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListNotificationDestinationsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PageSize":  types.Int64Type,
			"PageToken": types.StringType,
		},
	}
}

type ListNotificationDestinationsResponse struct {
	// Page token for next of results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ListNotificationDestinationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNotificationDestinationsResponse) {
}

func (newState *ListNotificationDestinationsResponse) SyncEffectiveFieldsDuringRead(existingState ListNotificationDestinationsResponse) {
}

func (a ListNotificationDestinationsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Results": reflect.TypeOf(ListNotificationDestinationsResult{}),
	}
}

func (a ListNotificationDestinationsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NextPageToken": types.StringType,
			"Results": basetypes.ListType{
				ElemType: ListNotificationDestinationsResult{}.ToAttrType(ctx),
			},
		},
	}
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

func (newState *ListNotificationDestinationsResult) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNotificationDestinationsResult) {
}

func (newState *ListNotificationDestinationsResult) SyncEffectiveFieldsDuringRead(existingState ListNotificationDestinationsResult) {
}

func (a ListNotificationDestinationsResult) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListNotificationDestinationsResult) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DestinationType": types.StringType,
			"DisplayName":     types.StringType,
			"Id":              types.StringType,
		},
	}
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListPrivateEndpointRulesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPrivateEndpointRulesRequest) {
}

func (newState *ListPrivateEndpointRulesRequest) SyncEffectiveFieldsDuringRead(existingState ListPrivateEndpointRulesRequest) {
}

func (a ListPrivateEndpointRulesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListPrivateEndpointRulesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NetworkConnectivityConfigId": types.StringType,
			"PageToken":                   types.StringType,
		},
	}
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos types.List `tfsdk:"token_infos" tf:"optional"`
}

func (newState *ListPublicTokensResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPublicTokensResponse) {
}

func (newState *ListPublicTokensResponse) SyncEffectiveFieldsDuringRead(existingState ListPublicTokensResponse) {
}

func (a ListPublicTokensResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"TokenInfos": reflect.TypeOf(PublicTokenInfo{}),
	}
}

func (a ListPublicTokensResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"TokenInfos": basetypes.ListType{
				ElemType: PublicTokenInfo{}.ToAttrType(ctx),
			},
		},
	}
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById types.Int64 `tfsdk:"-"`
	// Username of the user that created the token.
	CreatedByUsername types.String `tfsdk:"-"`
}

func (newState *ListTokenManagementRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTokenManagementRequest) {
}

func (newState *ListTokenManagementRequest) SyncEffectiveFieldsDuringRead(existingState ListTokenManagementRequest) {
}

func (a ListTokenManagementRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListTokenManagementRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreatedById":       types.Int64Type,
			"CreatedByUsername": types.StringType,
		},
	}
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos types.List `tfsdk:"token_infos" tf:"optional"`
}

func (newState *ListTokensResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTokensResponse) {
}

func (newState *ListTokensResponse) SyncEffectiveFieldsDuringRead(existingState ListTokensResponse) {
}

func (a ListTokensResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"TokenInfos": reflect.TypeOf(TokenInfo{}),
	}
}

func (a ListTokensResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"TokenInfos": basetypes.ListType{
				ElemType: TokenInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type MicrosoftTeamsConfig struct {
	// [Input-Only] URL for Microsoft Teams.
	Url types.String `tfsdk:"url" tf:"optional"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set" tf:"optional"`
}

func (newState *MicrosoftTeamsConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan MicrosoftTeamsConfig) {
}

func (newState *MicrosoftTeamsConfig) SyncEffectiveFieldsDuringRead(existingState MicrosoftTeamsConfig) {
}

func (a MicrosoftTeamsConfig) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a MicrosoftTeamsConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Url":    types.StringType,
			"UrlSet": types.BoolType,
		},
	}
}

// The stable AWS IP CIDR blocks. You can use these to configure the firewall of
// your resources to allow traffic from your Databricks workspace.
type NccAwsStableIpRule struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	CidrBlocks types.List `tfsdk:"cidr_blocks" tf:"optional"`
}

func (newState *NccAwsStableIpRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccAwsStableIpRule) {
}

func (newState *NccAwsStableIpRule) SyncEffectiveFieldsDuringRead(existingState NccAwsStableIpRule) {
}

func (a NccAwsStableIpRule) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"CidrBlocks": reflect.TypeOf(types.StringType),
	}
}

func (a NccAwsStableIpRule) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CidrBlocks": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
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
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed,optional"`
	// Whether this private endpoint is deactivated.
	Deactivated types.Bool `tfsdk:"deactivated" tf:"computed,optional"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt types.Int64 `tfsdk:"deactivated_at" tf:"computed,optional"`
	// The name of the Azure private endpoint resource.
	EndpointName types.String `tfsdk:"endpoint_name" tf:"computed,optional"`
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
	RuleId types.String `tfsdk:"rule_id" tf:"computed,optional"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time" tf:"computed,optional"`
}

func (newState *NccAzurePrivateEndpointRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccAzurePrivateEndpointRule) {
}

func (newState *NccAzurePrivateEndpointRule) SyncEffectiveFieldsDuringRead(existingState NccAzurePrivateEndpointRule) {
}

func (a NccAzurePrivateEndpointRule) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a NccAzurePrivateEndpointRule) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ConnectionState":             types.StringType,
			"CreationTime":                types.Int64Type,
			"Deactivated":                 types.BoolType,
			"DeactivatedAt":               types.Int64Type,
			"EndpointName":                types.StringType,
			"GroupId":                     types.StringType,
			"NetworkConnectivityConfigId": types.StringType,
			"ResourceId":                  types.StringType,
			"RuleId":                      types.StringType,
			"UpdatedTime":                 types.Int64Type,
		},
	}
}

// The stable Azure service endpoints. You can configure the firewall of your
// Azure resources to allow traffic from your Databricks serverless compute
// resources.
type NccAzureServiceEndpointRule struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets types.List `tfsdk:"subnets" tf:"optional"`
	// The Azure region in which this service endpoint rule applies.
	TargetRegion types.String `tfsdk:"target_region" tf:"optional"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices types.List `tfsdk:"target_services" tf:"optional"`
}

func (newState *NccAzureServiceEndpointRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccAzureServiceEndpointRule) {
}

func (newState *NccAzureServiceEndpointRule) SyncEffectiveFieldsDuringRead(existingState NccAzureServiceEndpointRule) {
}

func (a NccAzureServiceEndpointRule) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Subnets":        reflect.TypeOf(types.StringType),
		"TargetServices": reflect.TypeOf(types.StringType),
	}
}

func (a NccAzureServiceEndpointRule) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Subnets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"TargetRegion": types.StringType,
			"TargetServices": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// The network connectivity rules that apply to network traffic from your
// serverless compute resources.
type NccEgressConfig struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules types.List `tfsdk:"default_rules" tf:"computed,optional"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules types.List `tfsdk:"target_rules" tf:"optional,object"`
}

func (newState *NccEgressConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccEgressConfig) {
}

func (newState *NccEgressConfig) SyncEffectiveFieldsDuringRead(existingState NccEgressConfig) {
}

func (a NccEgressConfig) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DefaultRules": reflect.TypeOf(NccEgressDefaultRules{}),
		"TargetRules":  reflect.TypeOf(NccEgressTargetRules{}),
	}
}

func (a NccEgressConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DefaultRules": NccEgressDefaultRules{}.ToAttrType(ctx),
			"TargetRules":  NccEgressTargetRules{}.ToAttrType(ctx),
		},
	}
}

// The network connectivity rules that are applied by default without resource
// specific configurations. You can find the stable network information of your
// serverless compute resources here.
type NccEgressDefaultRules struct {
	// The stable AWS IP CIDR blocks. You can use these to configure the
	// firewall of your resources to allow traffic from your Databricks
	// workspace.
	AwsStableIpRule types.List `tfsdk:"aws_stable_ip_rule" tf:"optional,object"`
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	AzureServiceEndpointRule types.List `tfsdk:"azure_service_endpoint_rule" tf:"optional,object"`
}

func (newState *NccEgressDefaultRules) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccEgressDefaultRules) {
}

func (newState *NccEgressDefaultRules) SyncEffectiveFieldsDuringRead(existingState NccEgressDefaultRules) {
}

func (a NccEgressDefaultRules) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AwsStableIpRule":          reflect.TypeOf(NccAwsStableIpRule{}),
		"AzureServiceEndpointRule": reflect.TypeOf(NccAzureServiceEndpointRule{}),
	}
}

func (a NccEgressDefaultRules) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AwsStableIpRule":          NccAwsStableIpRule{}.ToAttrType(ctx),
			"AzureServiceEndpointRule": NccAzureServiceEndpointRule{}.ToAttrType(ctx),
		},
	}
}

// The network connectivity rules that configured for each destinations. These
// rules override default rules.
type NccEgressTargetRules struct {
	AzurePrivateEndpointRules types.List `tfsdk:"azure_private_endpoint_rules" tf:"optional"`
}

func (newState *NccEgressTargetRules) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccEgressTargetRules) {
}

func (newState *NccEgressTargetRules) SyncEffectiveFieldsDuringRead(existingState NccEgressTargetRules) {
}

func (a NccEgressTargetRules) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AzurePrivateEndpointRules": reflect.TypeOf(NccAzurePrivateEndpointRule{}),
	}
}

func (a NccEgressTargetRules) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AzurePrivateEndpointRules": basetypes.ListType{
				ElemType: NccAzurePrivateEndpointRule{}.ToAttrType(ctx),
			},
		},
	}
}

type NetworkConnectivityConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed,optional"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig types.List `tfsdk:"egress_config" tf:"optional,object"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Databricks network connectivity configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id" tf:"computed,optional"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region types.String `tfsdk:"region" tf:"optional"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time" tf:"computed,optional"`
}

func (newState *NetworkConnectivityConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkConnectivityConfiguration) {
}

func (newState *NetworkConnectivityConfiguration) SyncEffectiveFieldsDuringRead(existingState NetworkConnectivityConfiguration) {
}

func (a NetworkConnectivityConfiguration) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"EgressConfig": reflect.TypeOf(NccEgressConfig{}),
	}
}

func (a NetworkConnectivityConfiguration) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AccountId":                   types.StringType,
			"CreationTime":                types.Int64Type,
			"EgressConfig":                NccEgressConfig{}.ToAttrType(ctx),
			"Name":                        types.StringType,
			"NetworkConnectivityConfigId": types.StringType,
			"Region":                      types.StringType,
			"UpdatedTime":                 types.Int64Type,
		},
	}
}

type NotificationDestination struct {
	// The configuration for the notification destination. Will be exactly one
	// of the nested configs. Only returns for users with workspace admin
	// permissions.
	Config types.List `tfsdk:"config" tf:"optional,object"`
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType types.String `tfsdk:"destination_type" tf:"optional"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *NotificationDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotificationDestination) {
}

func (newState *NotificationDestination) SyncEffectiveFieldsDuringRead(existingState NotificationDestination) {
}

func (a NotificationDestination) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Config": reflect.TypeOf(Config{}),
	}
}

func (a NotificationDestination) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Config":          Config{}.ToAttrType(ctx),
			"DestinationType": types.StringType,
			"DisplayName":     types.StringType,
			"Id":              types.StringType,
		},
	}
}

type PagerdutyConfig struct {
	// [Input-Only] Integration key for PagerDuty.
	IntegrationKey types.String `tfsdk:"integration_key" tf:"optional"`
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet types.Bool `tfsdk:"integration_key_set" tf:"optional"`
}

func (newState *PagerdutyConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan PagerdutyConfig) {
}

func (newState *PagerdutyConfig) SyncEffectiveFieldsDuringRead(existingState PagerdutyConfig) {
}

func (a PagerdutyConfig) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PagerdutyConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IntegrationKey":    types.StringType,
			"IntegrationKeySet": types.BoolType,
		},
	}
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId types.Int64 `tfsdk:"workspaceId" tf:"optional"`
}

func (newState *PartitionId) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionId) {
}

func (newState *PartitionId) SyncEffectiveFieldsDuringRead(existingState PartitionId) {
}

func (a PartitionId) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PartitionId) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"WorkspaceId": types.Int64Type,
		},
	}
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

func (newState *PersonalComputeMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan PersonalComputeMessage) {
}

func (newState *PersonalComputeMessage) SyncEffectiveFieldsDuringRead(existingState PersonalComputeMessage) {
}

func (a PersonalComputeMessage) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PersonalComputeMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Value": types.StringType,
		},
	}
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

	PersonalCompute types.List `tfsdk:"personal_compute" tf:"object"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

func (newState *PersonalComputeSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan PersonalComputeSetting) {
}

func (newState *PersonalComputeSetting) SyncEffectiveFieldsDuringRead(existingState PersonalComputeSetting) {
}

func (a PersonalComputeSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PersonalCompute": reflect.TypeOf(PersonalComputeMessage{}),
	}
}

func (a PersonalComputeSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag":            types.StringType,
			"PersonalCompute": PersonalComputeMessage{}.ToAttrType(ctx),
			"SettingName":     types.StringType,
		},
	}
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

func (newState *PublicTokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublicTokenInfo) {
}

func (newState *PublicTokenInfo) SyncEffectiveFieldsDuringRead(existingState PublicTokenInfo) {
}

func (a PublicTokenInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PublicTokenInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment":      types.StringType,
			"CreationTime": types.Int64Type,
			"ExpiryTime":   types.Int64Type,
			"TokenId":      types.StringType,
		},
	}
}

// Details required to replace an IP access list.
type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:""`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`

	IpAddresses types.List `tfsdk:"ip_addresses" tf:"optional"`
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

func (newState *ReplaceIpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReplaceIpAccessList) {
}

func (newState *ReplaceIpAccessList) SyncEffectiveFieldsDuringRead(existingState ReplaceIpAccessList) {
}

func (a ReplaceIpAccessList) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"IpAddresses": reflect.TypeOf(types.StringType),
	}
}

func (a ReplaceIpAccessList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Enabled":        types.BoolType,
			"IpAccessListId": types.StringType,
			"IpAddresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Label":    types.StringType,
			"ListType": types.StringType,
		},
	}
}

type ReplaceResponse struct {
}

func (newState *ReplaceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReplaceResponse) {
}

func (newState *ReplaceResponse) SyncEffectiveFieldsDuringRead(existingState ReplaceResponse) {
}

func (a ReplaceResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ReplaceResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestrictWorkspaceAdminsMessage struct {
	Status types.String `tfsdk:"status" tf:""`
}

func (newState *RestrictWorkspaceAdminsMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestrictWorkspaceAdminsMessage) {
}

func (newState *RestrictWorkspaceAdminsMessage) SyncEffectiveFieldsDuringRead(existingState RestrictWorkspaceAdminsMessage) {
}

func (a RestrictWorkspaceAdminsMessage) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RestrictWorkspaceAdminsMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Status": types.StringType,
		},
	}
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

	RestrictWorkspaceAdmins types.List `tfsdk:"restrict_workspace_admins" tf:"object"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name" tf:"optional"`
}

func (newState *RestrictWorkspaceAdminsSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestrictWorkspaceAdminsSetting) {
}

func (newState *RestrictWorkspaceAdminsSetting) SyncEffectiveFieldsDuringRead(existingState RestrictWorkspaceAdminsSetting) {
}

func (a RestrictWorkspaceAdminsSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"RestrictWorkspaceAdmins": reflect.TypeOf(RestrictWorkspaceAdminsMessage{}),
	}
}

func (a RestrictWorkspaceAdminsSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Etag":                    types.StringType,
			"RestrictWorkspaceAdmins": RestrictWorkspaceAdminsMessage{}.ToAttrType(ctx),
			"SettingName":             types.StringType,
		},
	}
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId types.String `tfsdk:"token_id" tf:""`
}

func (newState *RevokeTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RevokeTokenRequest) {
}

func (newState *RevokeTokenRequest) SyncEffectiveFieldsDuringRead(existingState RevokeTokenRequest) {
}

func (a RevokeTokenRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RevokeTokenRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"TokenId": types.StringType,
		},
	}
}

type RevokeTokenResponse struct {
}

func (newState *RevokeTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RevokeTokenResponse) {
}

func (newState *RevokeTokenResponse) SyncEffectiveFieldsDuringRead(existingState RevokeTokenResponse) {
}

func (a RevokeTokenResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RevokeTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetStatusResponse struct {
}

func (newState *SetStatusResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetStatusResponse) {
}

func (newState *SetStatusResponse) SyncEffectiveFieldsDuringRead(existingState SetStatusResponse) {
}

func (a SetStatusResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SetStatusResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SlackConfig struct {
	// [Input-Only] URL for Slack destination.
	Url types.String `tfsdk:"url" tf:"optional"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set" tf:"optional"`
}

func (newState *SlackConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan SlackConfig) {
}

func (newState *SlackConfig) SyncEffectiveFieldsDuringRead(existingState SlackConfig) {
}

func (a SlackConfig) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SlackConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Url":    types.StringType,
			"UrlSet": types.BoolType,
		},
	}
}

type StringMessage struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *StringMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan StringMessage) {
}

func (newState *StringMessage) SyncEffectiveFieldsDuringRead(existingState StringMessage) {
}

func (a StringMessage) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a StringMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Value": types.StringType,
		},
	}
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

func (newState *TokenAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenAccessControlRequest) {
}

func (newState *TokenAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState TokenAccessControlRequest) {
}

func (a TokenAccessControlRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a TokenAccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"GroupName":            types.StringType,
			"PermissionLevel":      types.StringType,
			"ServicePrincipalName": types.StringType,
			"UserName":             types.StringType,
		},
	}
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *TokenAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenAccessControlResponse) {
}

func (newState *TokenAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState TokenAccessControlResponse) {
}

func (a TokenAccessControlResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AllPermissions": reflect.TypeOf(TokenPermission{}),
	}
}

func (a TokenAccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllPermissions": basetypes.ListType{
				ElemType: TokenPermission{}.ToAttrType(ctx),
			},
			"DisplayName":          types.StringType,
			"GroupName":            types.StringType,
			"ServicePrincipalName": types.StringType,
			"UserName":             types.StringType,
		},
	}
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
	// Approximate timestamp for the day the token was last used. Accurate up to
	// 1 day.
	LastUsedDay types.Int64 `tfsdk:"last_used_day" tf:"optional"`
	// User ID of the user that owns the token.
	OwnerId types.Int64 `tfsdk:"owner_id" tf:"optional"`
	// ID of the token.
	TokenId types.String `tfsdk:"token_id" tf:"optional"`
	// If applicable, the ID of the workspace that the token was created in.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

func (newState *TokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenInfo) {
}

func (newState *TokenInfo) SyncEffectiveFieldsDuringRead(existingState TokenInfo) {
}

func (a TokenInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a TokenInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment":           types.StringType,
			"CreatedById":       types.Int64Type,
			"CreatedByUsername": types.StringType,
			"CreationTime":      types.Int64Type,
			"ExpiryTime":        types.Int64Type,
			"LastUsedDay":       types.Int64Type,
			"OwnerId":           types.Int64Type,
			"TokenId":           types.StringType,
			"WorkspaceId":       types.Int64Type,
		},
	}
}

type TokenPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *TokenPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermission) {
}

func (newState *TokenPermission) SyncEffectiveFieldsDuringRead(existingState TokenPermission) {
}

func (a TokenPermission) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"InheritedFromObject": reflect.TypeOf(types.StringType),
	}
}

func (a TokenPermission) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Inherited": types.BoolType,
			"InheritedFromObject": basetypes.ListType{
				ElemType: types.StringType,
			},
			"PermissionLevel": types.StringType,
		},
	}
}

type TokenPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *TokenPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermissions) {
}

func (newState *TokenPermissions) SyncEffectiveFieldsDuringRead(existingState TokenPermissions) {
}

func (a TokenPermissions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(TokenAccessControlResponse{}),
	}
}

func (a TokenPermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AccessControlList": basetypes.ListType{
				ElemType: TokenAccessControlResponse{}.ToAttrType(ctx),
			},
			"ObjectId":   types.StringType,
			"ObjectType": types.StringType,
		},
	}
}

type TokenPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *TokenPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermissionsDescription) {
}

func (newState *TokenPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState TokenPermissionsDescription) {
}

func (a TokenPermissionsDescription) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a TokenPermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Description":     types.StringType,
			"PermissionLevel": types.StringType,
		},
	}
}

type TokenPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
}

func (newState *TokenPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermissionsRequest) {
}

func (newState *TokenPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState TokenPermissionsRequest) {
}

func (a TokenPermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(TokenAccessControlRequest{}),
	}
}

func (a TokenPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AccessControlList": basetypes.ListType{
				ElemType: TokenAccessControlRequest{}.ToAttrType(ctx),
			},
		},
	}
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (newState *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (a UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicySetting{}),
	}
}

func (a UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      AibiDashboardEmbeddingAccessPolicySetting{}.ToAttrType(ctx),
		},
	}
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (newState *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (a UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomainsSetting{}),
	}
}

func (a UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      AibiDashboardEmbeddingApprovedDomainsSetting{}.ToAttrType(ctx),
		},
	}
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

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateAutomaticClusterUpdateSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAutomaticClusterUpdateSettingRequest) {
}

func (newState *UpdateAutomaticClusterUpdateSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAutomaticClusterUpdateSettingRequest) {
}

func (a UpdateAutomaticClusterUpdateSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(AutomaticClusterUpdateSetting{}),
	}
}

func (a UpdateAutomaticClusterUpdateSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      AutomaticClusterUpdateSetting{}.ToAttrType(ctx),
		},
	}
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

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateComplianceSecurityProfileSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateComplianceSecurityProfileSettingRequest) {
}

func (newState *UpdateComplianceSecurityProfileSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateComplianceSecurityProfileSettingRequest) {
}

func (a UpdateComplianceSecurityProfileSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(ComplianceSecurityProfileSetting{}),
	}
}

func (a UpdateComplianceSecurityProfileSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      ComplianceSecurityProfileSetting{}.ToAttrType(ctx),
		},
	}
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

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateCspEnablementAccountSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCspEnablementAccountSettingRequest) {
}

func (newState *UpdateCspEnablementAccountSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateCspEnablementAccountSettingRequest) {
}

func (a UpdateCspEnablementAccountSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(CspEnablementAccountSetting{}),
	}
}

func (a UpdateCspEnablementAccountSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      CspEnablementAccountSetting{}.ToAttrType(ctx),
		},
	}
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
	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateDefaultNamespaceSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDefaultNamespaceSettingRequest) {
}

func (newState *UpdateDefaultNamespaceSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDefaultNamespaceSettingRequest) {
}

func (a UpdateDefaultNamespaceSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(DefaultNamespaceSetting{}),
	}
}

func (a UpdateDefaultNamespaceSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      DefaultNamespaceSetting{}.ToAttrType(ctx),
		},
	}
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

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateDisableLegacyAccessRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDisableLegacyAccessRequest) {
}

func (newState *UpdateDisableLegacyAccessRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDisableLegacyAccessRequest) {
}

func (a UpdateDisableLegacyAccessRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(DisableLegacyAccess{}),
	}
}

func (a UpdateDisableLegacyAccessRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      DisableLegacyAccess{}.ToAttrType(ctx),
		},
	}
}

// Details required to update a setting.
type UpdateDisableLegacyDbfsRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing" tf:""`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask types.String `tfsdk:"field_mask" tf:""`

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateDisableLegacyDbfsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDisableLegacyDbfsRequest) {
}

func (newState *UpdateDisableLegacyDbfsRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDisableLegacyDbfsRequest) {
}

func (a UpdateDisableLegacyDbfsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(DisableLegacyDbfs{}),
	}
}

func (a UpdateDisableLegacyDbfsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      DisableLegacyDbfs{}.ToAttrType(ctx),
		},
	}
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

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateDisableLegacyFeaturesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDisableLegacyFeaturesRequest) {
}

func (newState *UpdateDisableLegacyFeaturesRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDisableLegacyFeaturesRequest) {
}

func (a UpdateDisableLegacyFeaturesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(DisableLegacyFeatures{}),
	}
}

func (a UpdateDisableLegacyFeaturesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      DisableLegacyFeatures{}.ToAttrType(ctx),
		},
	}
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

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateEnhancedSecurityMonitoringSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateEnhancedSecurityMonitoringSettingRequest) {
}

func (newState *UpdateEnhancedSecurityMonitoringSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateEnhancedSecurityMonitoringSettingRequest) {
}

func (a UpdateEnhancedSecurityMonitoringSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(EnhancedSecurityMonitoringSetting{}),
	}
}

func (a UpdateEnhancedSecurityMonitoringSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      EnhancedSecurityMonitoringSetting{}.ToAttrType(ctx),
		},
	}
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

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateEsmEnablementAccountSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateEsmEnablementAccountSettingRequest) {
}

func (newState *UpdateEsmEnablementAccountSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateEsmEnablementAccountSettingRequest) {
}

func (a UpdateEsmEnablementAccountSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(EsmEnablementAccountSetting{}),
	}
}

func (a UpdateEsmEnablementAccountSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      EsmEnablementAccountSetting{}.ToAttrType(ctx),
		},
	}
}

// Details required to update an IP access list.
type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`

	IpAddresses types.List `tfsdk:"ip_addresses" tf:"optional"`
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

func (newState *UpdateIpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateIpAccessList) {
}

func (newState *UpdateIpAccessList) SyncEffectiveFieldsDuringRead(existingState UpdateIpAccessList) {
}

func (a UpdateIpAccessList) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"IpAddresses": reflect.TypeOf(types.StringType),
	}
}

func (a UpdateIpAccessList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Enabled":        types.BoolType,
			"IpAccessListId": types.StringType,
			"IpAddresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Label":    types.StringType,
			"ListType": types.StringType,
		},
	}
}

type UpdateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config types.List `tfsdk:"config" tf:"optional,object"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"-"`
}

func (newState *UpdateNotificationDestinationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateNotificationDestinationRequest) {
}

func (newState *UpdateNotificationDestinationRequest) SyncEffectiveFieldsDuringRead(existingState UpdateNotificationDestinationRequest) {
}

func (a UpdateNotificationDestinationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Config": reflect.TypeOf(Config{}),
	}
}

func (a UpdateNotificationDestinationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Config":      Config{}.ToAttrType(ctx),
			"DisplayName": types.StringType,
			"Id":          types.StringType,
		},
	}
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

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdatePersonalComputeSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePersonalComputeSettingRequest) {
}

func (newState *UpdatePersonalComputeSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdatePersonalComputeSettingRequest) {
}

func (a UpdatePersonalComputeSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(PersonalComputeSetting{}),
	}
}

func (a UpdatePersonalComputeSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      PersonalComputeSetting{}.ToAttrType(ctx),
		},
	}
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

func (a UpdateResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

	Setting types.List `tfsdk:"setting" tf:"object"`
}

func (newState *UpdateRestrictWorkspaceAdminsSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRestrictWorkspaceAdminsSettingRequest) {
}

func (newState *UpdateRestrictWorkspaceAdminsSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateRestrictWorkspaceAdminsSettingRequest) {
}

func (a UpdateRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Setting": reflect.TypeOf(RestrictWorkspaceAdminsSetting{}),
	}
}

func (a UpdateRestrictWorkspaceAdminsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllowMissing": types.BoolType,
			"FieldMask":    types.StringType,
			"Setting":      RestrictWorkspaceAdminsSetting{}.ToAttrType(ctx),
		},
	}
}
