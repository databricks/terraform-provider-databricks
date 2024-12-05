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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingAccessPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AibiDashboardEmbeddingAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of AibiDashboardEmbeddingAccessPolicy in the Terraform plugin framework type
// system.
func (a AibiDashboardEmbeddingAccessPolicy) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_policy_type": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingAccessPolicySetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AibiDashboardEmbeddingAccessPolicySetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_access_policy": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy{}),
	}
}

// ToAttrType returns the representation of AibiDashboardEmbeddingAccessPolicySetting in the Terraform plugin framework type
// system.
func (a AibiDashboardEmbeddingAccessPolicySetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aibi_dashboard_embedding_access_policy": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingAccessPolicy{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingApprovedDomains.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AibiDashboardEmbeddingApprovedDomains) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"approved_domains": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of AibiDashboardEmbeddingApprovedDomains in the Terraform plugin framework type
// system.
func (a AibiDashboardEmbeddingApprovedDomains) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"approved_domains": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingApprovedDomainsSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AibiDashboardEmbeddingApprovedDomainsSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_approved_domains": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains{}),
	}
}

// ToAttrType returns the representation of AibiDashboardEmbeddingApprovedDomainsSetting in the Terraform plugin framework type
// system.
func (a AibiDashboardEmbeddingApprovedDomainsSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aibi_dashboard_embedding_approved_domains": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingApprovedDomains{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutomaticClusterUpdateSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AutomaticClusterUpdateSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"automatic_cluster_update_workspace": reflect.TypeOf(ClusterAutoRestartMessage{}),
	}
}

// ToAttrType returns the representation of AutomaticClusterUpdateSetting in the Terraform plugin framework type
// system.
func (a AutomaticClusterUpdateSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"automatic_cluster_update_workspace": basetypes.ListType{
				ElemType: ClusterAutoRestartMessage{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BooleanMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BooleanMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of BooleanMessage in the Terraform plugin framework type
// system.
func (a BooleanMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enablement_details": reflect.TypeOf(ClusterAutoRestartMessageEnablementDetails{}),
		"maintenance_window": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindow{}),
	}
}

// ToAttrType returns the representation of ClusterAutoRestartMessage in the Terraform plugin framework type
// system.
func (a ClusterAutoRestartMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"can_toggle": types.BoolType,
			"enabled":    types.BoolType,
			"enablement_details": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageEnablementDetails{}.ToAttrType(ctx),
			},
			"maintenance_window": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageMaintenanceWindow{}.ToAttrType(ctx),
			},
			"restart_even_if_no_updates_available": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageEnablementDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessageEnablementDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ClusterAutoRestartMessageEnablementDetails in the Terraform plugin framework type
// system.
func (a ClusterAutoRestartMessageEnablementDetails) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"forced_for_compliance_mode":           types.BoolType,
			"unavailable_for_disabled_entitlement": types.BoolType,
			"unavailable_for_non_enterprise_tier":  types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessageMaintenanceWindow) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"week_day_based_schedule": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}),
	}
}

// ToAttrType returns the representation of ClusterAutoRestartMessageMaintenanceWindow in the Terraform plugin framework type
// system.
func (a ClusterAutoRestartMessageMaintenanceWindow) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"week_day_based_schedule": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"window_start_time": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}),
	}
}

// ToAttrType returns the representation of ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule in the Terraform plugin framework type
// system.
func (a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"day_of_week": types.StringType,
			"frequency":   types.StringType,
			"window_start_time": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindowWindowStartTime.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ClusterAutoRestartMessageMaintenanceWindowWindowStartTime in the Terraform plugin framework type
// system.
func (a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"hours":   types.Int64Type,
			"minutes": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComplianceSecurityProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ComplianceSecurityProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of ComplianceSecurityProfile in the Terraform plugin framework type
// system.
func (a ComplianceSecurityProfile) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_standards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_enabled": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComplianceSecurityProfileSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ComplianceSecurityProfileSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_security_profile_workspace": reflect.TypeOf(ComplianceSecurityProfile{}),
	}
}

// ToAttrType returns the representation of ComplianceSecurityProfileSetting in the Terraform plugin framework type
// system.
func (a ComplianceSecurityProfileSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_security_profile_workspace": basetypes.ListType{
				ElemType: ComplianceSecurityProfile{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Config.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Config) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email":           reflect.TypeOf(EmailConfig{}),
		"generic_webhook": reflect.TypeOf(GenericWebhookConfig{}),
		"microsoft_teams": reflect.TypeOf(MicrosoftTeamsConfig{}),
		"pagerduty":       reflect.TypeOf(PagerdutyConfig{}),
		"slack":           reflect.TypeOf(SlackConfig{}),
	}
}

// ToAttrType returns the representation of Config in the Terraform plugin framework type
// system.
func (a Config) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email": basetypes.ListType{
				ElemType: EmailConfig{}.ToAttrType(ctx),
			},
			"generic_webhook": basetypes.ListType{
				ElemType: GenericWebhookConfig{}.ToAttrType(ctx),
			},
			"microsoft_teams": basetypes.ListType{
				ElemType: MicrosoftTeamsConfig{}.ToAttrType(ctx),
			},
			"pagerduty": basetypes.ListType{
				ElemType: PagerdutyConfig{}.ToAttrType(ctx),
			},
			"slack": basetypes.ListType{
				ElemType: SlackConfig{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateIpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateIpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of CreateIpAccessList in the Terraform plugin framework type
// system.
func (a CreateIpAccessList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label":     types.StringType,
			"list_type": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// ToAttrType returns the representation of CreateIpAccessListResponse in the Terraform plugin framework type
// system.
func (a CreateIpAccessListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessListInfo{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkConnectivityConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateNetworkConnectivityConfigRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of CreateNetworkConnectivityConfigRequest in the Terraform plugin framework type
// system.
func (a CreateNetworkConnectivityConfigRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"region": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config{}),
	}
}

// ToAttrType returns the representation of CreateNotificationDestinationRequest in the Terraform plugin framework type
// system.
func (a CreateNotificationDestinationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": basetypes.ListType{
				ElemType: Config{}.ToAttrType(ctx),
			},
			"display_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOboTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateOboTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of CreateOboTokenRequest in the Terraform plugin framework type
// system.
func (a CreateOboTokenRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"application_id":   types.StringType,
			"comment":          types.StringType,
			"lifetime_seconds": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOboTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateOboTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo{}),
	}
}

// ToAttrType returns the representation of CreateOboTokenResponse in the Terraform plugin framework type
// system.
func (a CreateOboTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info": basetypes.ListType{
				ElemType: TokenInfo{}.ToAttrType(ctx),
			},
			"token_value": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of CreatePrivateEndpointRuleRequest in the Terraform plugin framework type
// system.
func (a CreatePrivateEndpointRuleRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":                       types.StringType,
			"network_connectivity_config_id": types.StringType,
			"resource_id":                    types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of CreateTokenRequest in the Terraform plugin framework type
// system.
func (a CreateTokenRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":          types.StringType,
			"lifetime_seconds": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(PublicTokenInfo{}),
	}
}

// ToAttrType returns the representation of CreateTokenResponse in the Terraform plugin framework type
// system.
func (a CreateTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info": basetypes.ListType{
				ElemType: PublicTokenInfo{}.ToAttrType(ctx),
			},
			"token_value": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CspEnablementAccount.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CspEnablementAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of CspEnablementAccount in the Terraform plugin framework type
// system.
func (a CspEnablementAccount) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_standards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_enforced": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CspEnablementAccountSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CspEnablementAccountSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"csp_enablement_account": reflect.TypeOf(CspEnablementAccount{}),
	}
}

// ToAttrType returns the representation of CspEnablementAccountSetting in the Terraform plugin framework type
// system.
func (a CspEnablementAccountSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"csp_enablement_account": basetypes.ListType{
				ElemType: CspEnablementAccount{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DefaultNamespaceSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DefaultNamespaceSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"namespace": reflect.TypeOf(StringMessage{}),
	}
}

// ToAttrType returns the representation of DefaultNamespaceSetting in the Terraform plugin framework type
// system.
func (a DefaultNamespaceSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"namespace": basetypes.ListType{
				ElemType: StringMessage{}.ToAttrType(ctx),
			},
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteAccountIpAccessListRequest in the Terraform plugin framework type
// system.
func (a DeleteAccountIpAccessListRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDefaultNamespaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDefaultNamespaceSettingRequest in the Terraform plugin framework type
// system.
func (a DeleteDefaultNamespaceSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultNamespaceSettingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDefaultNamespaceSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDefaultNamespaceSettingResponse in the Terraform plugin framework type
// system.
func (a DeleteDefaultNamespaceSettingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDisableLegacyAccessRequest in the Terraform plugin framework type
// system.
func (a DeleteDisableLegacyAccessRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyAccessResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyAccessResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDisableLegacyAccessResponse in the Terraform plugin framework type
// system.
func (a DeleteDisableLegacyAccessResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDisableLegacyDbfsRequest in the Terraform plugin framework type
// system.
func (a DeleteDisableLegacyDbfsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyDbfsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyDbfsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDisableLegacyDbfsResponse in the Terraform plugin framework type
// system.
func (a DeleteDisableLegacyDbfsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDisableLegacyFeaturesRequest in the Terraform plugin framework type
// system.
func (a DeleteDisableLegacyFeaturesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyFeaturesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyFeaturesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDisableLegacyFeaturesResponse in the Terraform plugin framework type
// system.
func (a DeleteDisableLegacyFeaturesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteIpAccessListRequest in the Terraform plugin framework type
// system.
func (a DeleteIpAccessListRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkConnectivityConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkConnectivityConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteNetworkConnectivityConfigurationRequest in the Terraform plugin framework type
// system.
func (a DeleteNetworkConnectivityConfigurationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
		},
	}
}

type DeleteNetworkConnectivityConfigurationResponse struct {
}

func (newState *DeleteNetworkConnectivityConfigurationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteNetworkConnectivityConfigurationResponse) {
}

func (newState *DeleteNetworkConnectivityConfigurationResponse) SyncEffectiveFieldsDuringRead(existingState DeleteNetworkConnectivityConfigurationResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkConnectivityConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkConnectivityConfigurationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteNetworkConnectivityConfigurationResponse in the Terraform plugin framework type
// system.
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteNotificationDestinationRequest in the Terraform plugin framework type
// system.
func (a DeleteNotificationDestinationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePersonalComputeSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeletePersonalComputeSettingRequest in the Terraform plugin framework type
// system.
func (a DeletePersonalComputeSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePersonalComputeSettingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePersonalComputeSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeletePersonalComputeSettingResponse in the Terraform plugin framework type
// system.
func (a DeletePersonalComputeSettingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeletePrivateEndpointRuleRequest in the Terraform plugin framework type
// system.
func (a DeletePrivateEndpointRuleRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteResponse in the Terraform plugin framework type
// system.
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteRestrictWorkspaceAdminsSettingRequest in the Terraform plugin framework type
// system.
func (a DeleteRestrictWorkspaceAdminsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRestrictWorkspaceAdminsSettingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRestrictWorkspaceAdminsSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteRestrictWorkspaceAdminsSettingResponse in the Terraform plugin framework type
// system.
func (a DeleteRestrictWorkspaceAdminsSettingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTokenManagementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTokenManagementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteTokenManagementRequest in the Terraform plugin framework type
// system.
func (a DeleteTokenManagementRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableLegacyAccess.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DisableLegacyAccess) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_access": reflect.TypeOf(BooleanMessage{}),
	}
}

// ToAttrType returns the representation of DisableLegacyAccess in the Terraform plugin framework type
// system.
func (a DisableLegacyAccess) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disable_legacy_access": basetypes.ListType{
				ElemType: BooleanMessage{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableLegacyDbfs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DisableLegacyDbfs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_dbfs": reflect.TypeOf(BooleanMessage{}),
	}
}

// ToAttrType returns the representation of DisableLegacyDbfs in the Terraform plugin framework type
// system.
func (a DisableLegacyDbfs) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disable_legacy_dbfs": basetypes.ListType{
				ElemType: BooleanMessage{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableLegacyFeatures.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DisableLegacyFeatures) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_features": reflect.TypeOf(BooleanMessage{}),
	}
}

// ToAttrType returns the representation of DisableLegacyFeatures in the Terraform plugin framework type
// system.
func (a DisableLegacyFeatures) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disable_legacy_features": basetypes.ListType{
				ElemType: BooleanMessage{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmailConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EmailConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"addresses": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of EmailConfig in the Terraform plugin framework type
// system.
func (a EmailConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"addresses": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Empty.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Empty) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of Empty in the Terraform plugin framework type
// system.
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnhancedSecurityMonitoring.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnhancedSecurityMonitoring) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of EnhancedSecurityMonitoring in the Terraform plugin framework type
// system.
func (a EnhancedSecurityMonitoring) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enabled": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnhancedSecurityMonitoringSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnhancedSecurityMonitoringSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enhanced_security_monitoring_workspace": reflect.TypeOf(EnhancedSecurityMonitoring{}),
	}
}

// ToAttrType returns the representation of EnhancedSecurityMonitoringSetting in the Terraform plugin framework type
// system.
func (a EnhancedSecurityMonitoringSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enhanced_security_monitoring_workspace": basetypes.ListType{
				ElemType: EnhancedSecurityMonitoring{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EsmEnablementAccount.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EsmEnablementAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of EsmEnablementAccount in the Terraform plugin framework type
// system.
func (a EsmEnablementAccount) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enforced": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EsmEnablementAccountSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EsmEnablementAccountSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"esm_enablement_account": reflect.TypeOf(EsmEnablementAccount{}),
	}
}

// ToAttrType returns the representation of EsmEnablementAccountSetting in the Terraform plugin framework type
// system.
func (a EsmEnablementAccountSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"esm_enablement_account": basetypes.ListType{
				ElemType: EsmEnablementAccount{}.ToAttrType(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExchangeToken) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of ExchangeToken in the Terraform plugin framework type
// system.
func (a ExchangeToken) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential":        types.StringType,
			"credentialEolTime": types.Int64Type,
			"ownerId":           types.Int64Type,
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"tokenType": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExchangeTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partitionId": reflect.TypeOf(PartitionId{}),
		"scopes":      reflect.TypeOf(types.String{}),
		"tokenType":   reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of ExchangeTokenRequest in the Terraform plugin framework type
// system.
func (a ExchangeTokenRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"partitionId": basetypes.ListType{
				ElemType: PartitionId{}.ToAttrType(ctx),
			},
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"tokenType": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExchangeTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(ExchangeToken{}),
	}
}

// ToAttrType returns the representation of ExchangeTokenResponse in the Terraform plugin framework type
// system.
func (a ExchangeTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"values": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FetchIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FetchIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// ToAttrType returns the representation of FetchIpAccessListResponse in the Terraform plugin framework type
// system.
func (a FetchIpAccessListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessListInfo{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenericWebhookConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenericWebhookConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GenericWebhookConfig in the Terraform plugin framework type
// system.
func (a GenericWebhookConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"password":     types.StringType,
			"password_set": types.BoolType,
			"url":          types.StringType,
			"url_set":      types.BoolType,
			"username":     types.StringType,
			"username_set": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetAccountIpAccessListRequest in the Terraform plugin framework type
// system.
func (a GetAccountIpAccessListRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetAibiDashboardEmbeddingAccessPolicySettingRequest in the Terraform plugin framework type
// system.
func (a GetAibiDashboardEmbeddingAccessPolicySettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetAibiDashboardEmbeddingApprovedDomainsSettingRequest in the Terraform plugin framework type
// system.
func (a GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAutomaticClusterUpdateSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAutomaticClusterUpdateSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetAutomaticClusterUpdateSettingRequest in the Terraform plugin framework type
// system.
func (a GetAutomaticClusterUpdateSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetComplianceSecurityProfileSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetComplianceSecurityProfileSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetComplianceSecurityProfileSettingRequest in the Terraform plugin framework type
// system.
func (a GetComplianceSecurityProfileSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCspEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCspEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetCspEnablementAccountSettingRequest in the Terraform plugin framework type
// system.
func (a GetCspEnablementAccountSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDefaultNamespaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetDefaultNamespaceSettingRequest in the Terraform plugin framework type
// system.
func (a GetDefaultNamespaceSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDisableLegacyAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetDisableLegacyAccessRequest in the Terraform plugin framework type
// system.
func (a GetDisableLegacyAccessRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDisableLegacyDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetDisableLegacyDbfsRequest in the Terraform plugin framework type
// system.
func (a GetDisableLegacyDbfsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDisableLegacyFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetDisableLegacyFeaturesRequest in the Terraform plugin framework type
// system.
func (a GetDisableLegacyFeaturesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnhancedSecurityMonitoringSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEnhancedSecurityMonitoringSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetEnhancedSecurityMonitoringSettingRequest in the Terraform plugin framework type
// system.
func (a GetEnhancedSecurityMonitoringSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEsmEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEsmEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetEsmEnablementAccountSettingRequest in the Terraform plugin framework type
// system.
func (a GetEsmEnablementAccountSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetIpAccessListRequest in the Terraform plugin framework type
// system.
func (a GetIpAccessListRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// ToAttrType returns the representation of GetIpAccessListResponse in the Terraform plugin framework type
// system.
func (a GetIpAccessListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessListInfo{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIpAccessListsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetIpAccessListsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// ToAttrType returns the representation of GetIpAccessListsResponse in the Terraform plugin framework type
// system.
func (a GetIpAccessListsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_lists": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkConnectivityConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNetworkConnectivityConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetNetworkConnectivityConfigurationRequest in the Terraform plugin framework type
// system.
func (a GetNetworkConnectivityConfigurationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetNotificationDestinationRequest in the Terraform plugin framework type
// system.
func (a GetNotificationDestinationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPersonalComputeSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetPersonalComputeSettingRequest in the Terraform plugin framework type
// system.
func (a GetPersonalComputeSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetPrivateEndpointRuleRequest in the Terraform plugin framework type
// system.
func (a GetPrivateEndpointRuleRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetRestrictWorkspaceAdminsSettingRequest in the Terraform plugin framework type
// system.
func (a GetRestrictWorkspaceAdminsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetStatusRequest in the Terraform plugin framework type
// system.
func (a GetStatusRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"keys": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenManagementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetTokenManagementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetTokenManagementRequest in the Terraform plugin framework type
// system.
func (a GetTokenManagementRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetTokenPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(TokenPermissionsDescription{}),
	}
}

// ToAttrType returns the representation of GetTokenPermissionLevelsResponse in the Terraform plugin framework type
// system.
func (a GetTokenPermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo{}),
	}
}

// ToAttrType returns the representation of GetTokenResponse in the Terraform plugin framework type
// system.
func (a GetTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info": basetypes.ListType{
				ElemType: TokenInfo{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IpAccessListInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IpAccessListInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of IpAccessListInfo in the Terraform plugin framework type
// system.
func (a IpAccessListInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"address_count": types.Int64Type,
			"created_at":    types.Int64Type,
			"created_by":    types.Int64Type,
			"enabled":       types.BoolType,
			"ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label":      types.StringType,
			"list_id":    types.StringType,
			"list_type":  types.StringType,
			"updated_at": types.Int64Type,
			"updated_by": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// ToAttrType returns the representation of ListIpAccessListResponse in the Terraform plugin framework type
// system.
func (a ListIpAccessListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_lists": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNccAzurePrivateEndpointRulesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNccAzurePrivateEndpointRulesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NccAzurePrivateEndpointRule{}),
	}
}

// ToAttrType returns the representation of ListNccAzurePrivateEndpointRulesResponse in the Terraform plugin framework type
// system.
func (a ListNccAzurePrivateEndpointRulesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: NccAzurePrivateEndpointRule{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworkConnectivityConfigurationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNetworkConnectivityConfigurationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ListNetworkConnectivityConfigurationsRequest in the Terraform plugin framework type
// system.
func (a ListNetworkConnectivityConfigurationsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworkConnectivityConfigurationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNetworkConnectivityConfigurationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NetworkConnectivityConfiguration{}),
	}
}

// ToAttrType returns the representation of ListNetworkConnectivityConfigurationsResponse in the Terraform plugin framework type
// system.
func (a ListNetworkConnectivityConfigurationsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: NetworkConnectivityConfiguration{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNotificationDestinationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNotificationDestinationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ListNotificationDestinationsRequest in the Terraform plugin framework type
// system.
func (a ListNotificationDestinationsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNotificationDestinationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNotificationDestinationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListNotificationDestinationsResult{}),
	}
}

// ToAttrType returns the representation of ListNotificationDestinationsResponse in the Terraform plugin framework type
// system.
func (a ListNotificationDestinationsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNotificationDestinationsResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNotificationDestinationsResult) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ListNotificationDestinationsResult in the Terraform plugin framework type
// system.
func (a ListNotificationDestinationsResult) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_type": types.StringType,
			"display_name":     types.StringType,
			"id":               types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPrivateEndpointRulesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPrivateEndpointRulesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ListPrivateEndpointRulesRequest in the Terraform plugin framework type
// system.
func (a ListPrivateEndpointRulesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"page_token":                     types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPublicTokensResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPublicTokensResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(PublicTokenInfo{}),
	}
}

// ToAttrType returns the representation of ListPublicTokensResponse in the Terraform plugin framework type
// system.
func (a ListPublicTokensResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_infos": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTokenManagementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTokenManagementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ListTokenManagementRequest in the Terraform plugin framework type
// system.
func (a ListTokenManagementRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_by_id":       types.Int64Type,
			"created_by_username": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTokensResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTokensResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(TokenInfo{}),
	}
}

// ToAttrType returns the representation of ListTokensResponse in the Terraform plugin framework type
// system.
func (a ListTokensResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_infos": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MicrosoftTeamsConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MicrosoftTeamsConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of MicrosoftTeamsConfig in the Terraform plugin framework type
// system.
func (a MicrosoftTeamsConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"url":     types.StringType,
			"url_set": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccAwsStableIpRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NccAwsStableIpRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cidr_blocks": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of NccAwsStableIpRule in the Terraform plugin framework type
// system.
func (a NccAwsStableIpRule) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cidr_blocks": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccAzurePrivateEndpointRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NccAzurePrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of NccAzurePrivateEndpointRule in the Terraform plugin framework type
// system.
func (a NccAzurePrivateEndpointRule) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_state":               types.StringType,
			"creation_time":                  types.Int64Type,
			"deactivated":                    types.BoolType,
			"deactivated_at":                 types.Int64Type,
			"endpoint_name":                  types.StringType,
			"group_id":                       types.StringType,
			"network_connectivity_config_id": types.StringType,
			"resource_id":                    types.StringType,
			"rule_id":                        types.StringType,
			"updated_time":                   types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccAzureServiceEndpointRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NccAzureServiceEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subnets":         reflect.TypeOf(types.String{}),
		"target_services": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of NccAzureServiceEndpointRule in the Terraform plugin framework type
// system.
func (a NccAzureServiceEndpointRule) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"subnets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"target_region": types.StringType,
			"target_services": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccEgressConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NccEgressConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_rules": reflect.TypeOf(NccEgressDefaultRules{}),
		"target_rules":  reflect.TypeOf(NccEgressTargetRules{}),
	}
}

// ToAttrType returns the representation of NccEgressConfig in the Terraform plugin framework type
// system.
func (a NccEgressConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_rules": basetypes.ListType{
				ElemType: NccEgressDefaultRules{}.ToAttrType(ctx),
			},
			"target_rules": basetypes.ListType{
				ElemType: NccEgressTargetRules{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccEgressDefaultRules.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NccEgressDefaultRules) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_stable_ip_rule":          reflect.TypeOf(NccAwsStableIpRule{}),
		"azure_service_endpoint_rule": reflect.TypeOf(NccAzureServiceEndpointRule{}),
	}
}

// ToAttrType returns the representation of NccEgressDefaultRules in the Terraform plugin framework type
// system.
func (a NccEgressDefaultRules) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_stable_ip_rule": basetypes.ListType{
				ElemType: NccAwsStableIpRule{}.ToAttrType(ctx),
			},
			"azure_service_endpoint_rule": basetypes.ListType{
				ElemType: NccAzureServiceEndpointRule{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccEgressTargetRules.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NccEgressTargetRules) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"azure_private_endpoint_rules": reflect.TypeOf(NccAzurePrivateEndpointRule{}),
	}
}

// ToAttrType returns the representation of NccEgressTargetRules in the Terraform plugin framework type
// system.
func (a NccEgressTargetRules) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"azure_private_endpoint_rules": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkConnectivityConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkConnectivityConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"egress_config": reflect.TypeOf(NccEgressConfig{}),
	}
}

// ToAttrType returns the representation of NetworkConnectivityConfiguration in the Terraform plugin framework type
// system.
func (a NetworkConnectivityConfiguration) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":    types.StringType,
			"creation_time": types.Int64Type,
			"egress_config": basetypes.ListType{
				ElemType: NccEgressConfig{}.ToAttrType(ctx),
			},
			"name":                           types.StringType,
			"network_connectivity_config_id": types.StringType,
			"region":                         types.StringType,
			"updated_time":                   types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotificationDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NotificationDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config{}),
	}
}

// ToAttrType returns the representation of NotificationDestination in the Terraform plugin framework type
// system.
func (a NotificationDestination) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": basetypes.ListType{
				ElemType: Config{}.ToAttrType(ctx),
			},
			"destination_type": types.StringType,
			"display_name":     types.StringType,
			"id":               types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PagerdutyConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PagerdutyConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of PagerdutyConfig in the Terraform plugin framework type
// system.
func (a PagerdutyConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_key":     types.StringType,
			"integration_key_set": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartitionId.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PartitionId) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of PartitionId in the Terraform plugin framework type
// system.
func (a PartitionId) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspaceId": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PersonalComputeMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PersonalComputeMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of PersonalComputeMessage in the Terraform plugin framework type
// system.
func (a PersonalComputeMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PersonalComputeSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PersonalComputeSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personal_compute": reflect.TypeOf(PersonalComputeMessage{}),
	}
}

// ToAttrType returns the representation of PersonalComputeSetting in the Terraform plugin framework type
// system.
func (a PersonalComputeSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"personal_compute": basetypes.ListType{
				ElemType: PersonalComputeMessage{}.ToAttrType(ctx),
			},
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublicTokenInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublicTokenInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of PublicTokenInfo in the Terraform plugin framework type
// system.
func (a PublicTokenInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":       types.StringType,
			"creation_time": types.Int64Type,
			"expiry_time":   types.Int64Type,
			"token_id":      types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplaceIpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReplaceIpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of ReplaceIpAccessList in the Terraform plugin framework type
// system.
func (a ReplaceIpAccessList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled":           types.BoolType,
			"ip_access_list_id": types.StringType,
			"ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label":     types.StringType,
			"list_type": types.StringType,
		},
	}
}

type ReplaceResponse struct {
}

func (newState *ReplaceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReplaceResponse) {
}

func (newState *ReplaceResponse) SyncEffectiveFieldsDuringRead(existingState ReplaceResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplaceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReplaceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ReplaceResponse in the Terraform plugin framework type
// system.
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestrictWorkspaceAdminsMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestrictWorkspaceAdminsMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of RestrictWorkspaceAdminsMessage in the Terraform plugin framework type
// system.
func (a RestrictWorkspaceAdminsMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestrictWorkspaceAdminsSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestrictWorkspaceAdminsSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"restrict_workspace_admins": reflect.TypeOf(RestrictWorkspaceAdminsMessage{}),
	}
}

// ToAttrType returns the representation of RestrictWorkspaceAdminsSetting in the Terraform plugin framework type
// system.
func (a RestrictWorkspaceAdminsSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"restrict_workspace_admins": basetypes.ListType{
				ElemType: RestrictWorkspaceAdminsMessage{}.ToAttrType(ctx),
			},
			"setting_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RevokeTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RevokeTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of RevokeTokenRequest in the Terraform plugin framework type
// system.
func (a RevokeTokenRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type RevokeTokenResponse struct {
}

func (newState *RevokeTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RevokeTokenResponse) {
}

func (newState *RevokeTokenResponse) SyncEffectiveFieldsDuringRead(existingState RevokeTokenResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RevokeTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RevokeTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of RevokeTokenResponse in the Terraform plugin framework type
// system.
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetStatusResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of SetStatusResponse in the Terraform plugin framework type
// system.
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SlackConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SlackConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of SlackConfig in the Terraform plugin framework type
// system.
func (a SlackConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"url":     types.StringType,
			"url_set": types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StringMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StringMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of StringMessage in the Terraform plugin framework type
// system.
func (a StringMessage) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of TokenAccessControlRequest in the Terraform plugin framework type
// system.
func (a TokenAccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(TokenPermission{}),
	}
}

// ToAttrType returns the representation of TokenAccessControlResponse in the Terraform plugin framework type
// system.
func (a TokenAccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: TokenPermission{}.ToAttrType(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of TokenInfo in the Terraform plugin framework type
// system.
func (a TokenInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":             types.StringType,
			"created_by_id":       types.Int64Type,
			"created_by_username": types.StringType,
			"creation_time":       types.Int64Type,
			"expiry_time":         types.Int64Type,
			"last_used_day":       types.Int64Type,
			"owner_id":            types.Int64Type,
			"token_id":            types.StringType,
			"workspace_id":        types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of TokenPermission in the Terraform plugin framework type
// system.
func (a TokenPermission) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlResponse{}),
	}
}

// ToAttrType returns the representation of TokenPermissions in the Terraform plugin framework type
// system.
func (a TokenPermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: TokenAccessControlResponse{}.ToAttrType(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of TokenPermissionsDescription in the Terraform plugin framework type
// system.
func (a TokenPermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlRequest{}),
	}
}

// ToAttrType returns the representation of TokenPermissionsRequest in the Terraform plugin framework type
// system.
func (a TokenPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicySetting{}),
	}
}

// ToAttrType returns the representation of UpdateAibiDashboardEmbeddingAccessPolicySettingRequest in the Terraform plugin framework type
// system.
func (a UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingAccessPolicySetting{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomainsSetting{}),
	}
}

// ToAttrType returns the representation of UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest in the Terraform plugin framework type
// system.
func (a UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingApprovedDomainsSetting{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAutomaticClusterUpdateSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAutomaticClusterUpdateSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AutomaticClusterUpdateSetting{}),
	}
}

// ToAttrType returns the representation of UpdateAutomaticClusterUpdateSettingRequest in the Terraform plugin framework type
// system.
func (a UpdateAutomaticClusterUpdateSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: AutomaticClusterUpdateSetting{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateComplianceSecurityProfileSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateComplianceSecurityProfileSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(ComplianceSecurityProfileSetting{}),
	}
}

// ToAttrType returns the representation of UpdateComplianceSecurityProfileSettingRequest in the Terraform plugin framework type
// system.
func (a UpdateComplianceSecurityProfileSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: ComplianceSecurityProfileSetting{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCspEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCspEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(CspEnablementAccountSetting{}),
	}
}

// ToAttrType returns the representation of UpdateCspEnablementAccountSettingRequest in the Terraform plugin framework type
// system.
func (a UpdateCspEnablementAccountSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: CspEnablementAccountSetting{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDefaultNamespaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DefaultNamespaceSetting{}),
	}
}

// ToAttrType returns the representation of UpdateDefaultNamespaceSettingRequest in the Terraform plugin framework type
// system.
func (a UpdateDefaultNamespaceSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DefaultNamespaceSetting{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDisableLegacyAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyAccess{}),
	}
}

// ToAttrType returns the representation of UpdateDisableLegacyAccessRequest in the Terraform plugin framework type
// system.
func (a UpdateDisableLegacyAccessRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DisableLegacyAccess{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDisableLegacyDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyDbfs{}),
	}
}

// ToAttrType returns the representation of UpdateDisableLegacyDbfsRequest in the Terraform plugin framework type
// system.
func (a UpdateDisableLegacyDbfsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DisableLegacyDbfs{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDisableLegacyFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyFeatures{}),
	}
}

// ToAttrType returns the representation of UpdateDisableLegacyFeaturesRequest in the Terraform plugin framework type
// system.
func (a UpdateDisableLegacyFeaturesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DisableLegacyFeatures{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEnhancedSecurityMonitoringSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateEnhancedSecurityMonitoringSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnhancedSecurityMonitoringSetting{}),
	}
}

// ToAttrType returns the representation of UpdateEnhancedSecurityMonitoringSettingRequest in the Terraform plugin framework type
// system.
func (a UpdateEnhancedSecurityMonitoringSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: EnhancedSecurityMonitoringSetting{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEsmEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateEsmEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EsmEnablementAccountSetting{}),
	}
}

// ToAttrType returns the representation of UpdateEsmEnablementAccountSettingRequest in the Terraform plugin framework type
// system.
func (a UpdateEsmEnablementAccountSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: EsmEnablementAccountSetting{}.ToAttrType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateIpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateIpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of UpdateIpAccessList in the Terraform plugin framework type
// system.
func (a UpdateIpAccessList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled":           types.BoolType,
			"ip_access_list_id": types.StringType,
			"ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label":     types.StringType,
			"list_type": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config{}),
	}
}

// ToAttrType returns the representation of UpdateNotificationDestinationRequest in the Terraform plugin framework type
// system.
func (a UpdateNotificationDestinationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": basetypes.ListType{
				ElemType: Config{}.ToAttrType(ctx),
			},
			"display_name": types.StringType,
			"id":           types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePersonalComputeSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(PersonalComputeSetting{}),
	}
}

// ToAttrType returns the representation of UpdatePersonalComputeSettingRequest in the Terraform plugin framework type
// system.
func (a UpdatePersonalComputeSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: PersonalComputeSetting{}.ToAttrType(ctx),
			},
		},
	}
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of UpdateResponse in the Terraform plugin framework type
// system.
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(RestrictWorkspaceAdminsSetting{}),
	}
}

// ToAttrType returns the representation of UpdateRestrictWorkspaceAdminsSettingRequest in the Terraform plugin framework type
// system.
func (a UpdateRestrictWorkspaceAdminsSettingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: RestrictWorkspaceAdminsSetting{}.ToAttrType(ctx),
			},
		},
	}
}
