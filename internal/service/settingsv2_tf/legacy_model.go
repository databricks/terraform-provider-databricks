// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package settingsv2_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AibiDashboardEmbeddingAccessPolicy_SdkV2 struct {
	AccessPolicyType types.String `tfsdk:"access_policy_type"`
}

func (toState *AibiDashboardEmbeddingAccessPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AibiDashboardEmbeddingAccessPolicy_SdkV2) {
}

func (toState *AibiDashboardEmbeddingAccessPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AibiDashboardEmbeddingAccessPolicy_SdkV2) {
}

func (c AibiDashboardEmbeddingAccessPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_policy_type"] = attrs["access_policy_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingAccessPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AibiDashboardEmbeddingAccessPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingAccessPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_policy_type": o.AccessPolicyType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingAccessPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_policy_type": types.StringType,
		},
	}
}

type AibiDashboardEmbeddingApprovedDomains_SdkV2 struct {
	ApprovedDomains types.List `tfsdk:"approved_domains"`
}

func (toState *AibiDashboardEmbeddingApprovedDomains_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AibiDashboardEmbeddingApprovedDomains_SdkV2) {
}

func (toState *AibiDashboardEmbeddingApprovedDomains_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AibiDashboardEmbeddingApprovedDomains_SdkV2) {
}

func (c AibiDashboardEmbeddingApprovedDomains_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["approved_domains"] = attrs["approved_domains"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingApprovedDomains.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AibiDashboardEmbeddingApprovedDomains_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"approved_domains": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomains_SdkV2
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingApprovedDomains_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"approved_domains": o.ApprovedDomains,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingApprovedDomains_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"approved_domains": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetApprovedDomains returns the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AibiDashboardEmbeddingApprovedDomains_SdkV2) GetApprovedDomains(ctx context.Context) ([]types.String, bool) {
	if o.ApprovedDomains.IsNull() || o.ApprovedDomains.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ApprovedDomains.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApprovedDomains sets the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains_SdkV2.
func (o *AibiDashboardEmbeddingApprovedDomains_SdkV2) SetApprovedDomains(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["approved_domains"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ApprovedDomains = types.ListValueMust(t, vs)
}

type BooleanMessage_SdkV2 struct {
	Value types.Bool `tfsdk:"value"`
}

func (toState *BooleanMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan BooleanMessage_SdkV2) {
}

func (toState *BooleanMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState BooleanMessage_SdkV2) {
}

func (c BooleanMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BooleanMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BooleanMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BooleanMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o BooleanMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BooleanMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.BoolType,
		},
	}
}

type ClusterAutoRestartMessage_SdkV2 struct {
	CanToggle types.Bool `tfsdk:"can_toggle"`

	Enabled types.Bool `tfsdk:"enabled"`

	EnablementDetails types.List `tfsdk:"enablement_details"`

	MaintenanceWindow types.List `tfsdk:"maintenance_window"`

	RestartEvenIfNoUpdatesAvailable types.Bool `tfsdk:"restart_even_if_no_updates_available"`
}

func (toState *ClusterAutoRestartMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessage_SdkV2) {
	if !fromPlan.EnablementDetails.IsNull() && !fromPlan.EnablementDetails.IsUnknown() {
		if toStateEnablementDetails, ok := toState.GetEnablementDetails(ctx); ok {
			if fromPlanEnablementDetails, ok := fromPlan.GetEnablementDetails(ctx); ok {
				toStateEnablementDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEnablementDetails)
				toState.SetEnablementDetails(ctx, toStateEnablementDetails)
			}
		}
	}
	if !fromPlan.MaintenanceWindow.IsNull() && !fromPlan.MaintenanceWindow.IsUnknown() {
		if toStateMaintenanceWindow, ok := toState.GetMaintenanceWindow(ctx); ok {
			if fromPlanMaintenanceWindow, ok := fromPlan.GetMaintenanceWindow(ctx); ok {
				toStateMaintenanceWindow.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanMaintenanceWindow)
				toState.SetMaintenanceWindow(ctx, toStateMaintenanceWindow)
			}
		}
	}
}

func (toState *ClusterAutoRestartMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessage_SdkV2) {
	if !fromState.EnablementDetails.IsNull() && !fromState.EnablementDetails.IsUnknown() {
		if toStateEnablementDetails, ok := toState.GetEnablementDetails(ctx); ok {
			if fromStateEnablementDetails, ok := fromState.GetEnablementDetails(ctx); ok {
				toStateEnablementDetails.SyncFieldsDuringRead(ctx, fromStateEnablementDetails)
				toState.SetEnablementDetails(ctx, toStateEnablementDetails)
			}
		}
	}
	if !fromState.MaintenanceWindow.IsNull() && !fromState.MaintenanceWindow.IsUnknown() {
		if toStateMaintenanceWindow, ok := toState.GetMaintenanceWindow(ctx); ok {
			if fromStateMaintenanceWindow, ok := fromState.GetMaintenanceWindow(ctx); ok {
				toStateMaintenanceWindow.SyncFieldsDuringRead(ctx, fromStateMaintenanceWindow)
				toState.SetMaintenanceWindow(ctx, toStateMaintenanceWindow)
			}
		}
	}
}

func (c ClusterAutoRestartMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["can_toggle"] = attrs["can_toggle"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["enablement_details"] = attrs["enablement_details"].SetOptional()
	attrs["enablement_details"] = attrs["enablement_details"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["maintenance_window"] = attrs["maintenance_window"].SetOptional()
	attrs["maintenance_window"] = attrs["maintenance_window"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["restart_even_if_no_updates_available"] = attrs["restart_even_if_no_updates_available"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enablement_details": reflect.TypeOf(ClusterAutoRestartMessageEnablementDetails_SdkV2{}),
		"maintenance_window": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindow_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"can_toggle":                           o.CanToggle,
			"enabled":                              o.Enabled,
			"enablement_details":                   o.EnablementDetails,
			"maintenance_window":                   o.MaintenanceWindow,
			"restart_even_if_no_updates_available": o.RestartEvenIfNoUpdatesAvailable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"can_toggle": types.BoolType,
			"enabled":    types.BoolType,
			"enablement_details": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageEnablementDetails_SdkV2{}.Type(ctx),
			},
			"maintenance_window": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageMaintenanceWindow_SdkV2{}.Type(ctx),
			},
			"restart_even_if_no_updates_available": types.BoolType,
		},
	}
}

// GetEnablementDetails returns the value of the EnablementDetails field in ClusterAutoRestartMessage_SdkV2 as
// a ClusterAutoRestartMessageEnablementDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessage_SdkV2) GetEnablementDetails(ctx context.Context) (ClusterAutoRestartMessageEnablementDetails_SdkV2, bool) {
	var e ClusterAutoRestartMessageEnablementDetails_SdkV2
	if o.EnablementDetails.IsNull() || o.EnablementDetails.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageEnablementDetails_SdkV2
	d := o.EnablementDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnablementDetails sets the value of the EnablementDetails field in ClusterAutoRestartMessage_SdkV2.
func (o *ClusterAutoRestartMessage_SdkV2) SetEnablementDetails(ctx context.Context, v ClusterAutoRestartMessageEnablementDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["enablement_details"]
	o.EnablementDetails = types.ListValueMust(t, vs)
}

// GetMaintenanceWindow returns the value of the MaintenanceWindow field in ClusterAutoRestartMessage_SdkV2 as
// a ClusterAutoRestartMessageMaintenanceWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessage_SdkV2) GetMaintenanceWindow(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindow_SdkV2, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindow_SdkV2
	if o.MaintenanceWindow.IsNull() || o.MaintenanceWindow.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindow_SdkV2
	d := o.MaintenanceWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMaintenanceWindow sets the value of the MaintenanceWindow field in ClusterAutoRestartMessage_SdkV2.
func (o *ClusterAutoRestartMessage_SdkV2) SetMaintenanceWindow(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["maintenance_window"]
	o.MaintenanceWindow = types.ListValueMust(t, vs)
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails_SdkV2 struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode types.Bool `tfsdk:"forced_for_compliance_mode"`
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	UnavailableForDisabledEntitlement types.Bool `tfsdk:"unavailable_for_disabled_entitlement"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier types.Bool `tfsdk:"unavailable_for_non_enterprise_tier"`
}

func (toState *ClusterAutoRestartMessageEnablementDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessageEnablementDetails_SdkV2) {
}

func (toState *ClusterAutoRestartMessageEnablementDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessageEnablementDetails_SdkV2) {
}

func (c ClusterAutoRestartMessageEnablementDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["forced_for_compliance_mode"] = attrs["forced_for_compliance_mode"].SetOptional()
	attrs["unavailable_for_disabled_entitlement"] = attrs["unavailable_for_disabled_entitlement"].SetOptional()
	attrs["unavailable_for_non_enterprise_tier"] = attrs["unavailable_for_non_enterprise_tier"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageEnablementDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessageEnablementDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageEnablementDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageEnablementDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"forced_for_compliance_mode":           o.ForcedForComplianceMode,
			"unavailable_for_disabled_entitlement": o.UnavailableForDisabledEntitlement,
			"unavailable_for_non_enterprise_tier":  o.UnavailableForNonEnterpriseTier,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageEnablementDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"forced_for_compliance_mode":           types.BoolType,
			"unavailable_for_disabled_entitlement": types.BoolType,
			"unavailable_for_non_enterprise_tier":  types.BoolType,
		},
	}
}

type ClusterAutoRestartMessageMaintenanceWindow_SdkV2 struct {
	WeekDayBasedSchedule types.List `tfsdk:"week_day_based_schedule"`
}

func (toState *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessageMaintenanceWindow_SdkV2) {
	if !fromPlan.WeekDayBasedSchedule.IsNull() && !fromPlan.WeekDayBasedSchedule.IsUnknown() {
		if toStateWeekDayBasedSchedule, ok := toState.GetWeekDayBasedSchedule(ctx); ok {
			if fromPlanWeekDayBasedSchedule, ok := fromPlan.GetWeekDayBasedSchedule(ctx); ok {
				toStateWeekDayBasedSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanWeekDayBasedSchedule)
				toState.SetWeekDayBasedSchedule(ctx, toStateWeekDayBasedSchedule)
			}
		}
	}
}

func (toState *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessageMaintenanceWindow_SdkV2) {
	if !fromState.WeekDayBasedSchedule.IsNull() && !fromState.WeekDayBasedSchedule.IsUnknown() {
		if toStateWeekDayBasedSchedule, ok := toState.GetWeekDayBasedSchedule(ctx); ok {
			if fromStateWeekDayBasedSchedule, ok := fromState.GetWeekDayBasedSchedule(ctx); ok {
				toStateWeekDayBasedSchedule.SyncFieldsDuringRead(ctx, fromStateWeekDayBasedSchedule)
				toState.SetWeekDayBasedSchedule(ctx, toStateWeekDayBasedSchedule)
			}
		}
	}
}

func (c ClusterAutoRestartMessageMaintenanceWindow_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["week_day_based_schedule"] = attrs["week_day_based_schedule"].SetOptional()
	attrs["week_day_based_schedule"] = attrs["week_day_based_schedule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessageMaintenanceWindow_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"week_day_based_schedule": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindow_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindow_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"week_day_based_schedule": o.WeekDayBasedSchedule,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindow_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"week_day_based_schedule": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWeekDayBasedSchedule returns the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow_SdkV2 as
// a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) GetWeekDayBasedSchedule(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2
	if o.WeekDayBasedSchedule.IsNull() || o.WeekDayBasedSchedule.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2
	d := o.WeekDayBasedSchedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWeekDayBasedSchedule sets the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow_SdkV2.
func (o *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) SetWeekDayBasedSchedule(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["week_day_based_schedule"]
	o.WeekDayBasedSchedule = types.ListValueMust(t, vs)
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2 struct {
	DayOfWeek types.String `tfsdk:"day_of_week"`

	Frequency types.String `tfsdk:"frequency"`

	WindowStartTime types.List `tfsdk:"window_start_time"`
}

func (toState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) {
	if !fromPlan.WindowStartTime.IsNull() && !fromPlan.WindowStartTime.IsUnknown() {
		if toStateWindowStartTime, ok := toState.GetWindowStartTime(ctx); ok {
			if fromPlanWindowStartTime, ok := fromPlan.GetWindowStartTime(ctx); ok {
				toStateWindowStartTime.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanWindowStartTime)
				toState.SetWindowStartTime(ctx, toStateWindowStartTime)
			}
		}
	}
}

func (toState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) {
	if !fromState.WindowStartTime.IsNull() && !fromState.WindowStartTime.IsUnknown() {
		if toStateWindowStartTime, ok := toState.GetWindowStartTime(ctx); ok {
			if fromStateWindowStartTime, ok := fromState.GetWindowStartTime(ctx); ok {
				toStateWindowStartTime.SyncFieldsDuringRead(ctx, fromStateWindowStartTime)
				toState.SetWindowStartTime(ctx, toStateWindowStartTime)
			}
		}
	}
}

func (c ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["day_of_week"] = attrs["day_of_week"].SetOptional()
	attrs["frequency"] = attrs["frequency"].SetOptional()
	attrs["window_start_time"] = attrs["window_start_time"].SetOptional()
	attrs["window_start_time"] = attrs["window_start_time"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"window_start_time": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"day_of_week":       o.DayOfWeek,
			"frequency":         o.Frequency,
			"window_start_time": o.WindowStartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"day_of_week": types.StringType,
			"frequency":   types.StringType,
			"window_start_time": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWindowStartTime returns the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2 as
// a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) GetWindowStartTime(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2
	if o.WindowStartTime.IsNull() || o.WindowStartTime.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2
	d := o.WindowStartTime.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWindowStartTime sets the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2.
func (o *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) SetWindowStartTime(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["window_start_time"]
	o.WindowStartTime = types.ListValueMust(t, vs)
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2 struct {
	Hours types.Int64 `tfsdk:"hours"`

	Minutes types.Int64 `tfsdk:"minutes"`
}

func (toState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) {
}

func (toState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) {
}

func (c ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["hours"] = attrs["hours"].SetOptional()
	attrs["minutes"] = attrs["minutes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindowWindowStartTime.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"hours":   o.Hours,
			"minutes": o.Minutes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"hours":   types.Int64Type,
			"minutes": types.Int64Type,
		},
	}
}

// Changes the behaviour of Jobs service when creating job clusters.
//
// Before this setting is introduced, all workspaces with metastore attached had
// behaviour matching SINGLE_USER setting.
//
// See: - go/defaultdatasecuritymode - go/defaultdatasecuritymode/setting -
// go/datasecuritymode
type DefaultDataSecurityModeMessage_SdkV2 struct {
	Status types.String `tfsdk:"status"`
}

func (toState *DefaultDataSecurityModeMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DefaultDataSecurityModeMessage_SdkV2) {
}

func (toState *DefaultDataSecurityModeMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DefaultDataSecurityModeMessage_SdkV2) {
}

func (c DefaultDataSecurityModeMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["status"] = attrs["status"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DefaultDataSecurityModeMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DefaultDataSecurityModeMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultDataSecurityModeMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o DefaultDataSecurityModeMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DefaultDataSecurityModeMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": types.StringType,
		},
	}
}

type GetPublicAccountSettingRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublicAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublicAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublicAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPublicAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublicAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetPublicWorkspaceSettingRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublicWorkspaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublicWorkspaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublicWorkspaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPublicWorkspaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublicWorkspaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type IntegerMessage_SdkV2 struct {
	Value types.Int64 `tfsdk:"value"`
}

func (toState *IntegerMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan IntegerMessage_SdkV2) {
}

func (toState *IntegerMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState IntegerMessage_SdkV2) {
}

func (c IntegerMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IntegerMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IntegerMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IntegerMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o IntegerMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IntegerMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.Int64Type,
		},
	}
}

type ListAccountSettingsMetadataRequest_SdkV2 struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous
	// `ListAccountSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListAccountSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountSettingsMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountSettingsMetadataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountSettingsMetadataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountSettingsMetadataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountSettingsMetadataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAccountSettingsMetadataResponse_SdkV2 struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata types.List `tfsdk:"settings_metadata"`
}

func (toState *ListAccountSettingsMetadataResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListAccountSettingsMetadataResponse_SdkV2) {
}

func (toState *ListAccountSettingsMetadataResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListAccountSettingsMetadataResponse_SdkV2) {
}

func (c ListAccountSettingsMetadataResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["settings_metadata"] = attrs["settings_metadata"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountSettingsMetadataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountSettingsMetadataResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings_metadata": reflect.TypeOf(SettingsMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountSettingsMetadataResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountSettingsMetadataResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"settings_metadata": o.SettingsMetadata,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountSettingsMetadataResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"settings_metadata": basetypes.ListType{
				ElemType: SettingsMetadata_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSettingsMetadata returns the value of the SettingsMetadata field in ListAccountSettingsMetadataResponse_SdkV2 as
// a slice of SettingsMetadata_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAccountSettingsMetadataResponse_SdkV2) GetSettingsMetadata(ctx context.Context) ([]SettingsMetadata_SdkV2, bool) {
	if o.SettingsMetadata.IsNull() || o.SettingsMetadata.IsUnknown() {
		return nil, false
	}
	var v []SettingsMetadata_SdkV2
	d := o.SettingsMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettingsMetadata sets the value of the SettingsMetadata field in ListAccountSettingsMetadataResponse_SdkV2.
func (o *ListAccountSettingsMetadataResponse_SdkV2) SetSettingsMetadata(ctx context.Context, v []SettingsMetadata_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings_metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SettingsMetadata = types.ListValueMust(t, vs)
}

type ListWorkspaceSettingsMetadataRequest_SdkV2 struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous
	// `ListWorkspaceSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListWorkspaceSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceSettingsMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspaceSettingsMetadataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceSettingsMetadataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWorkspaceSettingsMetadataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceSettingsMetadataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListWorkspaceSettingsMetadataResponse_SdkV2 struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata types.List `tfsdk:"settings_metadata"`
}

func (toState *ListWorkspaceSettingsMetadataResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListWorkspaceSettingsMetadataResponse_SdkV2) {
}

func (toState *ListWorkspaceSettingsMetadataResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListWorkspaceSettingsMetadataResponse_SdkV2) {
}

func (c ListWorkspaceSettingsMetadataResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["settings_metadata"] = attrs["settings_metadata"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceSettingsMetadataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspaceSettingsMetadataResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings_metadata": reflect.TypeOf(SettingsMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceSettingsMetadataResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWorkspaceSettingsMetadataResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"settings_metadata": o.SettingsMetadata,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceSettingsMetadataResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"settings_metadata": basetypes.ListType{
				ElemType: SettingsMetadata_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSettingsMetadata returns the value of the SettingsMetadata field in ListWorkspaceSettingsMetadataResponse_SdkV2 as
// a slice of SettingsMetadata_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListWorkspaceSettingsMetadataResponse_SdkV2) GetSettingsMetadata(ctx context.Context) ([]SettingsMetadata_SdkV2, bool) {
	if o.SettingsMetadata.IsNull() || o.SettingsMetadata.IsUnknown() {
		return nil, false
	}
	var v []SettingsMetadata_SdkV2
	d := o.SettingsMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettingsMetadata sets the value of the SettingsMetadata field in ListWorkspaceSettingsMetadataResponse_SdkV2.
func (o *ListWorkspaceSettingsMetadataResponse_SdkV2) SetSettingsMetadata(ctx context.Context, v []SettingsMetadata_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings_metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SettingsMetadata = types.ListValueMust(t, vs)
}

type PatchPublicAccountSettingRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`

	Setting types.List `tfsdk:"setting"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchPublicAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchPublicAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(Setting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchPublicAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PatchPublicAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"setting": o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchPublicAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"setting": basetypes.ListType{
				ElemType: Setting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in PatchPublicAccountSettingRequest_SdkV2 as
// a Setting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchPublicAccountSettingRequest_SdkV2) GetSetting(ctx context.Context) (Setting_SdkV2, bool) {
	var e Setting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []Setting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in PatchPublicAccountSettingRequest_SdkV2.
func (o *PatchPublicAccountSettingRequest_SdkV2) SetSetting(ctx context.Context, v Setting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
}

type PatchPublicWorkspaceSettingRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`

	Setting types.List `tfsdk:"setting"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchPublicWorkspaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchPublicWorkspaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(Setting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchPublicWorkspaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PatchPublicWorkspaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"setting": o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchPublicWorkspaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"setting": basetypes.ListType{
				ElemType: Setting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in PatchPublicWorkspaceSettingRequest_SdkV2 as
// a Setting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchPublicWorkspaceSettingRequest_SdkV2) GetSetting(ctx context.Context) (Setting_SdkV2, bool) {
	var e Setting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []Setting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in PatchPublicWorkspaceSettingRequest_SdkV2.
func (o *PatchPublicWorkspaceSettingRequest_SdkV2) SetSetting(ctx context.Context, v Setting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
}

type PersonalComputeMessage_SdkV2 struct {
	Value types.String `tfsdk:"value"`
}

func (toState *PersonalComputeMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PersonalComputeMessage_SdkV2) {
}

func (toState *PersonalComputeMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PersonalComputeMessage_SdkV2) {
}

func (c PersonalComputeMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PersonalComputeMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PersonalComputeMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o PersonalComputeMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PersonalComputeMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type RestrictWorkspaceAdminsMessage_SdkV2 struct {
	Status types.String `tfsdk:"status"`
}

func (toState *RestrictWorkspaceAdminsMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RestrictWorkspaceAdminsMessage_SdkV2) {
}

func (toState *RestrictWorkspaceAdminsMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RestrictWorkspaceAdminsMessage_SdkV2) {
}

func (c RestrictWorkspaceAdminsMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["status"] = attrs["status"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestrictWorkspaceAdminsMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestrictWorkspaceAdminsMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o RestrictWorkspaceAdminsMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestrictWorkspaceAdminsMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": types.StringType,
		},
	}
}

type Setting_SdkV2 struct {
	AibiDashboardEmbeddingAccessPolicy types.List `tfsdk:"aibi_dashboard_embedding_access_policy"`

	AibiDashboardEmbeddingApprovedDomains types.List `tfsdk:"aibi_dashboard_embedding_approved_domains"`
	// todo: Mark these Public after onboarded to DSL
	AutomaticClusterUpdateWorkspace types.List `tfsdk:"automatic_cluster_update_workspace"`

	BooleanVal types.List `tfsdk:"boolean_val"`

	DefaultDataSecurityMode types.List `tfsdk:"default_data_security_mode"`

	EffectiveAibiDashboardEmbeddingAccessPolicy types.List `tfsdk:"effective_aibi_dashboard_embedding_access_policy"`

	EffectiveAibiDashboardEmbeddingApprovedDomains types.List `tfsdk:"effective_aibi_dashboard_embedding_approved_domains"`

	EffectiveAutomaticClusterUpdateWorkspace types.List `tfsdk:"effective_automatic_cluster_update_workspace"`

	EffectiveBooleanVal types.List `tfsdk:"effective_boolean_val"`

	EffectiveDefaultDataSecurityMode types.List `tfsdk:"effective_default_data_security_mode"`

	EffectiveIntegerVal types.List `tfsdk:"effective_integer_val"`

	EffectivePersonalCompute types.List `tfsdk:"effective_personal_compute"`

	EffectiveRestrictWorkspaceAdmins types.List `tfsdk:"effective_restrict_workspace_admins"`

	EffectiveStringVal types.List `tfsdk:"effective_string_val"`

	IntegerVal types.List `tfsdk:"integer_val"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`

	PersonalCompute types.List `tfsdk:"personal_compute"`

	RestrictWorkspaceAdmins types.List `tfsdk:"restrict_workspace_admins"`

	StringVal types.List `tfsdk:"string_val"`
}

func (toState *Setting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Setting_SdkV2) {
	if !fromPlan.AibiDashboardEmbeddingAccessPolicy.IsNull() && !fromPlan.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toStateAibiDashboardEmbeddingAccessPolicy, ok := toState.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromPlanAibiDashboardEmbeddingAccessPolicy, ok := fromPlan.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				toStateAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAibiDashboardEmbeddingAccessPolicy)
				toState.SetAibiDashboardEmbeddingAccessPolicy(ctx, toStateAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
	if !fromPlan.AibiDashboardEmbeddingApprovedDomains.IsNull() && !fromPlan.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toStateAibiDashboardEmbeddingApprovedDomains, ok := toState.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromPlanAibiDashboardEmbeddingApprovedDomains, ok := fromPlan.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				toStateAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAibiDashboardEmbeddingApprovedDomains)
				toState.SetAibiDashboardEmbeddingApprovedDomains(ctx, toStateAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
	if !fromPlan.AutomaticClusterUpdateWorkspace.IsNull() && !fromPlan.AutomaticClusterUpdateWorkspace.IsUnknown() {
		if toStateAutomaticClusterUpdateWorkspace, ok := toState.GetAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromPlanAutomaticClusterUpdateWorkspace, ok := fromPlan.GetAutomaticClusterUpdateWorkspace(ctx); ok {
				toStateAutomaticClusterUpdateWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAutomaticClusterUpdateWorkspace)
				toState.SetAutomaticClusterUpdateWorkspace(ctx, toStateAutomaticClusterUpdateWorkspace)
			}
		}
	}
	if !fromPlan.BooleanVal.IsNull() && !fromPlan.BooleanVal.IsUnknown() {
		if toStateBooleanVal, ok := toState.GetBooleanVal(ctx); ok {
			if fromPlanBooleanVal, ok := fromPlan.GetBooleanVal(ctx); ok {
				toStateBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanBooleanVal)
				toState.SetBooleanVal(ctx, toStateBooleanVal)
			}
		}
	}
	if !fromPlan.DefaultDataSecurityMode.IsNull() && !fromPlan.DefaultDataSecurityMode.IsUnknown() {
		if toStateDefaultDataSecurityMode, ok := toState.GetDefaultDataSecurityMode(ctx); ok {
			if fromPlanDefaultDataSecurityMode, ok := fromPlan.GetDefaultDataSecurityMode(ctx); ok {
				toStateDefaultDataSecurityMode.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanDefaultDataSecurityMode)
				toState.SetDefaultDataSecurityMode(ctx, toStateDefaultDataSecurityMode)
			}
		}
	}
	if !fromPlan.EffectiveAibiDashboardEmbeddingAccessPolicy.IsNull() && !fromPlan.EffectiveAibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toStateEffectiveAibiDashboardEmbeddingAccessPolicy, ok := toState.GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromPlanEffectiveAibiDashboardEmbeddingAccessPolicy, ok := fromPlan.GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				toStateEffectiveAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectiveAibiDashboardEmbeddingAccessPolicy)
				toState.SetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx, toStateEffectiveAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
	if !fromPlan.EffectiveAibiDashboardEmbeddingApprovedDomains.IsNull() && !fromPlan.EffectiveAibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toStateEffectiveAibiDashboardEmbeddingApprovedDomains, ok := toState.GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromPlanEffectiveAibiDashboardEmbeddingApprovedDomains, ok := fromPlan.GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				toStateEffectiveAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectiveAibiDashboardEmbeddingApprovedDomains)
				toState.SetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx, toStateEffectiveAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
	if !fromPlan.EffectiveAutomaticClusterUpdateWorkspace.IsNull() && !fromPlan.EffectiveAutomaticClusterUpdateWorkspace.IsUnknown() {
		if toStateEffectiveAutomaticClusterUpdateWorkspace, ok := toState.GetEffectiveAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromPlanEffectiveAutomaticClusterUpdateWorkspace, ok := fromPlan.GetEffectiveAutomaticClusterUpdateWorkspace(ctx); ok {
				toStateEffectiveAutomaticClusterUpdateWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectiveAutomaticClusterUpdateWorkspace)
				toState.SetEffectiveAutomaticClusterUpdateWorkspace(ctx, toStateEffectiveAutomaticClusterUpdateWorkspace)
			}
		}
	}
	if !fromPlan.EffectiveBooleanVal.IsNull() && !fromPlan.EffectiveBooleanVal.IsUnknown() {
		if toStateEffectiveBooleanVal, ok := toState.GetEffectiveBooleanVal(ctx); ok {
			if fromPlanEffectiveBooleanVal, ok := fromPlan.GetEffectiveBooleanVal(ctx); ok {
				toStateEffectiveBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectiveBooleanVal)
				toState.SetEffectiveBooleanVal(ctx, toStateEffectiveBooleanVal)
			}
		}
	}
	if !fromPlan.EffectiveDefaultDataSecurityMode.IsNull() && !fromPlan.EffectiveDefaultDataSecurityMode.IsUnknown() {
		if toStateEffectiveDefaultDataSecurityMode, ok := toState.GetEffectiveDefaultDataSecurityMode(ctx); ok {
			if fromPlanEffectiveDefaultDataSecurityMode, ok := fromPlan.GetEffectiveDefaultDataSecurityMode(ctx); ok {
				toStateEffectiveDefaultDataSecurityMode.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectiveDefaultDataSecurityMode)
				toState.SetEffectiveDefaultDataSecurityMode(ctx, toStateEffectiveDefaultDataSecurityMode)
			}
		}
	}
	if !fromPlan.EffectiveIntegerVal.IsNull() && !fromPlan.EffectiveIntegerVal.IsUnknown() {
		if toStateEffectiveIntegerVal, ok := toState.GetEffectiveIntegerVal(ctx); ok {
			if fromPlanEffectiveIntegerVal, ok := fromPlan.GetEffectiveIntegerVal(ctx); ok {
				toStateEffectiveIntegerVal.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectiveIntegerVal)
				toState.SetEffectiveIntegerVal(ctx, toStateEffectiveIntegerVal)
			}
		}
	}
	if !fromPlan.EffectivePersonalCompute.IsNull() && !fromPlan.EffectivePersonalCompute.IsUnknown() {
		if toStateEffectivePersonalCompute, ok := toState.GetEffectivePersonalCompute(ctx); ok {
			if fromPlanEffectivePersonalCompute, ok := fromPlan.GetEffectivePersonalCompute(ctx); ok {
				toStateEffectivePersonalCompute.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectivePersonalCompute)
				toState.SetEffectivePersonalCompute(ctx, toStateEffectivePersonalCompute)
			}
		}
	}
	if !fromPlan.EffectiveRestrictWorkspaceAdmins.IsNull() && !fromPlan.EffectiveRestrictWorkspaceAdmins.IsUnknown() {
		if toStateEffectiveRestrictWorkspaceAdmins, ok := toState.GetEffectiveRestrictWorkspaceAdmins(ctx); ok {
			if fromPlanEffectiveRestrictWorkspaceAdmins, ok := fromPlan.GetEffectiveRestrictWorkspaceAdmins(ctx); ok {
				toStateEffectiveRestrictWorkspaceAdmins.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectiveRestrictWorkspaceAdmins)
				toState.SetEffectiveRestrictWorkspaceAdmins(ctx, toStateEffectiveRestrictWorkspaceAdmins)
			}
		}
	}
	if !fromPlan.EffectiveStringVal.IsNull() && !fromPlan.EffectiveStringVal.IsUnknown() {
		if toStateEffectiveStringVal, ok := toState.GetEffectiveStringVal(ctx); ok {
			if fromPlanEffectiveStringVal, ok := fromPlan.GetEffectiveStringVal(ctx); ok {
				toStateEffectiveStringVal.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectiveStringVal)
				toState.SetEffectiveStringVal(ctx, toStateEffectiveStringVal)
			}
		}
	}
	if !fromPlan.IntegerVal.IsNull() && !fromPlan.IntegerVal.IsUnknown() {
		if toStateIntegerVal, ok := toState.GetIntegerVal(ctx); ok {
			if fromPlanIntegerVal, ok := fromPlan.GetIntegerVal(ctx); ok {
				toStateIntegerVal.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanIntegerVal)
				toState.SetIntegerVal(ctx, toStateIntegerVal)
			}
		}
	}
	if !fromPlan.PersonalCompute.IsNull() && !fromPlan.PersonalCompute.IsUnknown() {
		if toStatePersonalCompute, ok := toState.GetPersonalCompute(ctx); ok {
			if fromPlanPersonalCompute, ok := fromPlan.GetPersonalCompute(ctx); ok {
				toStatePersonalCompute.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanPersonalCompute)
				toState.SetPersonalCompute(ctx, toStatePersonalCompute)
			}
		}
	}
	if !fromPlan.RestrictWorkspaceAdmins.IsNull() && !fromPlan.RestrictWorkspaceAdmins.IsUnknown() {
		if toStateRestrictWorkspaceAdmins, ok := toState.GetRestrictWorkspaceAdmins(ctx); ok {
			if fromPlanRestrictWorkspaceAdmins, ok := fromPlan.GetRestrictWorkspaceAdmins(ctx); ok {
				toStateRestrictWorkspaceAdmins.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRestrictWorkspaceAdmins)
				toState.SetRestrictWorkspaceAdmins(ctx, toStateRestrictWorkspaceAdmins)
			}
		}
	}
	if !fromPlan.StringVal.IsNull() && !fromPlan.StringVal.IsUnknown() {
		if toStateStringVal, ok := toState.GetStringVal(ctx); ok {
			if fromPlanStringVal, ok := fromPlan.GetStringVal(ctx); ok {
				toStateStringVal.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanStringVal)
				toState.SetStringVal(ctx, toStateStringVal)
			}
		}
	}
}

func (toState *Setting_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Setting_SdkV2) {
	if !fromState.AibiDashboardEmbeddingAccessPolicy.IsNull() && !fromState.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toStateAibiDashboardEmbeddingAccessPolicy, ok := toState.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromStateAibiDashboardEmbeddingAccessPolicy, ok := fromState.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				toStateAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringRead(ctx, fromStateAibiDashboardEmbeddingAccessPolicy)
				toState.SetAibiDashboardEmbeddingAccessPolicy(ctx, toStateAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
	if !fromState.AibiDashboardEmbeddingApprovedDomains.IsNull() && !fromState.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toStateAibiDashboardEmbeddingApprovedDomains, ok := toState.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromStateAibiDashboardEmbeddingApprovedDomains, ok := fromState.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				toStateAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringRead(ctx, fromStateAibiDashboardEmbeddingApprovedDomains)
				toState.SetAibiDashboardEmbeddingApprovedDomains(ctx, toStateAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
	if !fromState.AutomaticClusterUpdateWorkspace.IsNull() && !fromState.AutomaticClusterUpdateWorkspace.IsUnknown() {
		if toStateAutomaticClusterUpdateWorkspace, ok := toState.GetAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromStateAutomaticClusterUpdateWorkspace, ok := fromState.GetAutomaticClusterUpdateWorkspace(ctx); ok {
				toStateAutomaticClusterUpdateWorkspace.SyncFieldsDuringRead(ctx, fromStateAutomaticClusterUpdateWorkspace)
				toState.SetAutomaticClusterUpdateWorkspace(ctx, toStateAutomaticClusterUpdateWorkspace)
			}
		}
	}
	if !fromState.BooleanVal.IsNull() && !fromState.BooleanVal.IsUnknown() {
		if toStateBooleanVal, ok := toState.GetBooleanVal(ctx); ok {
			if fromStateBooleanVal, ok := fromState.GetBooleanVal(ctx); ok {
				toStateBooleanVal.SyncFieldsDuringRead(ctx, fromStateBooleanVal)
				toState.SetBooleanVal(ctx, toStateBooleanVal)
			}
		}
	}
	if !fromState.DefaultDataSecurityMode.IsNull() && !fromState.DefaultDataSecurityMode.IsUnknown() {
		if toStateDefaultDataSecurityMode, ok := toState.GetDefaultDataSecurityMode(ctx); ok {
			if fromStateDefaultDataSecurityMode, ok := fromState.GetDefaultDataSecurityMode(ctx); ok {
				toStateDefaultDataSecurityMode.SyncFieldsDuringRead(ctx, fromStateDefaultDataSecurityMode)
				toState.SetDefaultDataSecurityMode(ctx, toStateDefaultDataSecurityMode)
			}
		}
	}
	if !fromState.EffectiveAibiDashboardEmbeddingAccessPolicy.IsNull() && !fromState.EffectiveAibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toStateEffectiveAibiDashboardEmbeddingAccessPolicy, ok := toState.GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromStateEffectiveAibiDashboardEmbeddingAccessPolicy, ok := fromState.GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				toStateEffectiveAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringRead(ctx, fromStateEffectiveAibiDashboardEmbeddingAccessPolicy)
				toState.SetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx, toStateEffectiveAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
	if !fromState.EffectiveAibiDashboardEmbeddingApprovedDomains.IsNull() && !fromState.EffectiveAibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toStateEffectiveAibiDashboardEmbeddingApprovedDomains, ok := toState.GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromStateEffectiveAibiDashboardEmbeddingApprovedDomains, ok := fromState.GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				toStateEffectiveAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringRead(ctx, fromStateEffectiveAibiDashboardEmbeddingApprovedDomains)
				toState.SetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx, toStateEffectiveAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
	if !fromState.EffectiveAutomaticClusterUpdateWorkspace.IsNull() && !fromState.EffectiveAutomaticClusterUpdateWorkspace.IsUnknown() {
		if toStateEffectiveAutomaticClusterUpdateWorkspace, ok := toState.GetEffectiveAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromStateEffectiveAutomaticClusterUpdateWorkspace, ok := fromState.GetEffectiveAutomaticClusterUpdateWorkspace(ctx); ok {
				toStateEffectiveAutomaticClusterUpdateWorkspace.SyncFieldsDuringRead(ctx, fromStateEffectiveAutomaticClusterUpdateWorkspace)
				toState.SetEffectiveAutomaticClusterUpdateWorkspace(ctx, toStateEffectiveAutomaticClusterUpdateWorkspace)
			}
		}
	}
	if !fromState.EffectiveBooleanVal.IsNull() && !fromState.EffectiveBooleanVal.IsUnknown() {
		if toStateEffectiveBooleanVal, ok := toState.GetEffectiveBooleanVal(ctx); ok {
			if fromStateEffectiveBooleanVal, ok := fromState.GetEffectiveBooleanVal(ctx); ok {
				toStateEffectiveBooleanVal.SyncFieldsDuringRead(ctx, fromStateEffectiveBooleanVal)
				toState.SetEffectiveBooleanVal(ctx, toStateEffectiveBooleanVal)
			}
		}
	}
	if !fromState.EffectiveDefaultDataSecurityMode.IsNull() && !fromState.EffectiveDefaultDataSecurityMode.IsUnknown() {
		if toStateEffectiveDefaultDataSecurityMode, ok := toState.GetEffectiveDefaultDataSecurityMode(ctx); ok {
			if fromStateEffectiveDefaultDataSecurityMode, ok := fromState.GetEffectiveDefaultDataSecurityMode(ctx); ok {
				toStateEffectiveDefaultDataSecurityMode.SyncFieldsDuringRead(ctx, fromStateEffectiveDefaultDataSecurityMode)
				toState.SetEffectiveDefaultDataSecurityMode(ctx, toStateEffectiveDefaultDataSecurityMode)
			}
		}
	}
	if !fromState.EffectiveIntegerVal.IsNull() && !fromState.EffectiveIntegerVal.IsUnknown() {
		if toStateEffectiveIntegerVal, ok := toState.GetEffectiveIntegerVal(ctx); ok {
			if fromStateEffectiveIntegerVal, ok := fromState.GetEffectiveIntegerVal(ctx); ok {
				toStateEffectiveIntegerVal.SyncFieldsDuringRead(ctx, fromStateEffectiveIntegerVal)
				toState.SetEffectiveIntegerVal(ctx, toStateEffectiveIntegerVal)
			}
		}
	}
	if !fromState.EffectivePersonalCompute.IsNull() && !fromState.EffectivePersonalCompute.IsUnknown() {
		if toStateEffectivePersonalCompute, ok := toState.GetEffectivePersonalCompute(ctx); ok {
			if fromStateEffectivePersonalCompute, ok := fromState.GetEffectivePersonalCompute(ctx); ok {
				toStateEffectivePersonalCompute.SyncFieldsDuringRead(ctx, fromStateEffectivePersonalCompute)
				toState.SetEffectivePersonalCompute(ctx, toStateEffectivePersonalCompute)
			}
		}
	}
	if !fromState.EffectiveRestrictWorkspaceAdmins.IsNull() && !fromState.EffectiveRestrictWorkspaceAdmins.IsUnknown() {
		if toStateEffectiveRestrictWorkspaceAdmins, ok := toState.GetEffectiveRestrictWorkspaceAdmins(ctx); ok {
			if fromStateEffectiveRestrictWorkspaceAdmins, ok := fromState.GetEffectiveRestrictWorkspaceAdmins(ctx); ok {
				toStateEffectiveRestrictWorkspaceAdmins.SyncFieldsDuringRead(ctx, fromStateEffectiveRestrictWorkspaceAdmins)
				toState.SetEffectiveRestrictWorkspaceAdmins(ctx, toStateEffectiveRestrictWorkspaceAdmins)
			}
		}
	}
	if !fromState.EffectiveStringVal.IsNull() && !fromState.EffectiveStringVal.IsUnknown() {
		if toStateEffectiveStringVal, ok := toState.GetEffectiveStringVal(ctx); ok {
			if fromStateEffectiveStringVal, ok := fromState.GetEffectiveStringVal(ctx); ok {
				toStateEffectiveStringVal.SyncFieldsDuringRead(ctx, fromStateEffectiveStringVal)
				toState.SetEffectiveStringVal(ctx, toStateEffectiveStringVal)
			}
		}
	}
	if !fromState.IntegerVal.IsNull() && !fromState.IntegerVal.IsUnknown() {
		if toStateIntegerVal, ok := toState.GetIntegerVal(ctx); ok {
			if fromStateIntegerVal, ok := fromState.GetIntegerVal(ctx); ok {
				toStateIntegerVal.SyncFieldsDuringRead(ctx, fromStateIntegerVal)
				toState.SetIntegerVal(ctx, toStateIntegerVal)
			}
		}
	}
	if !fromState.PersonalCompute.IsNull() && !fromState.PersonalCompute.IsUnknown() {
		if toStatePersonalCompute, ok := toState.GetPersonalCompute(ctx); ok {
			if fromStatePersonalCompute, ok := fromState.GetPersonalCompute(ctx); ok {
				toStatePersonalCompute.SyncFieldsDuringRead(ctx, fromStatePersonalCompute)
				toState.SetPersonalCompute(ctx, toStatePersonalCompute)
			}
		}
	}
	if !fromState.RestrictWorkspaceAdmins.IsNull() && !fromState.RestrictWorkspaceAdmins.IsUnknown() {
		if toStateRestrictWorkspaceAdmins, ok := toState.GetRestrictWorkspaceAdmins(ctx); ok {
			if fromStateRestrictWorkspaceAdmins, ok := fromState.GetRestrictWorkspaceAdmins(ctx); ok {
				toStateRestrictWorkspaceAdmins.SyncFieldsDuringRead(ctx, fromStateRestrictWorkspaceAdmins)
				toState.SetRestrictWorkspaceAdmins(ctx, toStateRestrictWorkspaceAdmins)
			}
		}
	}
	if !fromState.StringVal.IsNull() && !fromState.StringVal.IsUnknown() {
		if toStateStringVal, ok := toState.GetStringVal(ctx); ok {
			if fromStateStringVal, ok := fromState.GetStringVal(ctx); ok {
				toStateStringVal.SyncFieldsDuringRead(ctx, fromStateStringVal)
				toState.SetStringVal(ctx, toStateStringVal)
			}
		}
	}
}

func (c Setting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aibi_dashboard_embedding_access_policy"] = attrs["aibi_dashboard_embedding_access_policy"].SetOptional()
	attrs["aibi_dashboard_embedding_access_policy"] = attrs["aibi_dashboard_embedding_access_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["aibi_dashboard_embedding_approved_domains"] = attrs["aibi_dashboard_embedding_approved_domains"].SetOptional()
	attrs["aibi_dashboard_embedding_approved_domains"] = attrs["aibi_dashboard_embedding_approved_domains"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["automatic_cluster_update_workspace"] = attrs["automatic_cluster_update_workspace"].SetOptional()
	attrs["automatic_cluster_update_workspace"] = attrs["automatic_cluster_update_workspace"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["default_data_security_mode"] = attrs["default_data_security_mode"].SetOptional()
	attrs["default_data_security_mode"] = attrs["default_data_security_mode"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_aibi_dashboard_embedding_access_policy"] = attrs["effective_aibi_dashboard_embedding_access_policy"].SetOptional()
	attrs["effective_aibi_dashboard_embedding_access_policy"] = attrs["effective_aibi_dashboard_embedding_access_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_aibi_dashboard_embedding_approved_domains"] = attrs["effective_aibi_dashboard_embedding_approved_domains"].SetOptional()
	attrs["effective_aibi_dashboard_embedding_approved_domains"] = attrs["effective_aibi_dashboard_embedding_approved_domains"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_automatic_cluster_update_workspace"] = attrs["effective_automatic_cluster_update_workspace"].SetOptional()
	attrs["effective_automatic_cluster_update_workspace"] = attrs["effective_automatic_cluster_update_workspace"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].SetComputed()
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_default_data_security_mode"] = attrs["effective_default_data_security_mode"].SetOptional()
	attrs["effective_default_data_security_mode"] = attrs["effective_default_data_security_mode"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_integer_val"] = attrs["effective_integer_val"].SetComputed()
	attrs["effective_integer_val"] = attrs["effective_integer_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_personal_compute"] = attrs["effective_personal_compute"].SetOptional()
	attrs["effective_personal_compute"] = attrs["effective_personal_compute"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_restrict_workspace_admins"] = attrs["effective_restrict_workspace_admins"].SetOptional()
	attrs["effective_restrict_workspace_admins"] = attrs["effective_restrict_workspace_admins"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_string_val"] = attrs["effective_string_val"].SetComputed()
	attrs["effective_string_val"] = attrs["effective_string_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["integer_val"] = attrs["integer_val"].SetOptional()
	attrs["integer_val"] = attrs["integer_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["personal_compute"] = attrs["personal_compute"].SetOptional()
	attrs["personal_compute"] = attrs["personal_compute"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["restrict_workspace_admins"] = attrs["restrict_workspace_admins"].SetOptional()
	attrs["restrict_workspace_admins"] = attrs["restrict_workspace_admins"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["string_val"] = attrs["string_val"].SetOptional()
	attrs["string_val"] = attrs["string_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Setting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Setting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_access_policy":              reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy_SdkV2{}),
		"aibi_dashboard_embedding_approved_domains":           reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains_SdkV2{}),
		"automatic_cluster_update_workspace":                  reflect.TypeOf(ClusterAutoRestartMessage_SdkV2{}),
		"boolean_val":                                         reflect.TypeOf(BooleanMessage_SdkV2{}),
		"default_data_security_mode":                          reflect.TypeOf(DefaultDataSecurityModeMessage_SdkV2{}),
		"effective_aibi_dashboard_embedding_access_policy":    reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy_SdkV2{}),
		"effective_aibi_dashboard_embedding_approved_domains": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains_SdkV2{}),
		"effective_automatic_cluster_update_workspace":        reflect.TypeOf(ClusterAutoRestartMessage_SdkV2{}),
		"effective_boolean_val":                               reflect.TypeOf(BooleanMessage_SdkV2{}),
		"effective_default_data_security_mode":                reflect.TypeOf(DefaultDataSecurityModeMessage_SdkV2{}),
		"effective_integer_val":                               reflect.TypeOf(IntegerMessage_SdkV2{}),
		"effective_personal_compute":                          reflect.TypeOf(PersonalComputeMessage_SdkV2{}),
		"effective_restrict_workspace_admins":                 reflect.TypeOf(RestrictWorkspaceAdminsMessage_SdkV2{}),
		"effective_string_val":                                reflect.TypeOf(StringMessage_SdkV2{}),
		"integer_val":                                         reflect.TypeOf(IntegerMessage_SdkV2{}),
		"personal_compute":                                    reflect.TypeOf(PersonalComputeMessage_SdkV2{}),
		"restrict_workspace_admins":                           reflect.TypeOf(RestrictWorkspaceAdminsMessage_SdkV2{}),
		"string_val":                                          reflect.TypeOf(StringMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Setting_SdkV2
// only implements ToObjectValue() and Type().
func (o Setting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_access_policy":              o.AibiDashboardEmbeddingAccessPolicy,
			"aibi_dashboard_embedding_approved_domains":           o.AibiDashboardEmbeddingApprovedDomains,
			"automatic_cluster_update_workspace":                  o.AutomaticClusterUpdateWorkspace,
			"boolean_val":                                         o.BooleanVal,
			"default_data_security_mode":                          o.DefaultDataSecurityMode,
			"effective_aibi_dashboard_embedding_access_policy":    o.EffectiveAibiDashboardEmbeddingAccessPolicy,
			"effective_aibi_dashboard_embedding_approved_domains": o.EffectiveAibiDashboardEmbeddingApprovedDomains,
			"effective_automatic_cluster_update_workspace":        o.EffectiveAutomaticClusterUpdateWorkspace,
			"effective_boolean_val":                               o.EffectiveBooleanVal,
			"effective_default_data_security_mode":                o.EffectiveDefaultDataSecurityMode,
			"effective_integer_val":                               o.EffectiveIntegerVal,
			"effective_personal_compute":                          o.EffectivePersonalCompute,
			"effective_restrict_workspace_admins":                 o.EffectiveRestrictWorkspaceAdmins,
			"effective_string_val":                                o.EffectiveStringVal,
			"integer_val":                                         o.IntegerVal,
			"name":                                                o.Name,
			"personal_compute":                                    o.PersonalCompute,
			"restrict_workspace_admins":                           o.RestrictWorkspaceAdmins,
			"string_val":                                          o.StringVal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Setting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aibi_dashboard_embedding_access_policy": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingAccessPolicy_SdkV2{}.Type(ctx),
			},
			"aibi_dashboard_embedding_approved_domains": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingApprovedDomains_SdkV2{}.Type(ctx),
			},
			"automatic_cluster_update_workspace": basetypes.ListType{
				ElemType: ClusterAutoRestartMessage_SdkV2{}.Type(ctx),
			},
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"default_data_security_mode": basetypes.ListType{
				ElemType: DefaultDataSecurityModeMessage_SdkV2{}.Type(ctx),
			},
			"effective_aibi_dashboard_embedding_access_policy": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingAccessPolicy_SdkV2{}.Type(ctx),
			},
			"effective_aibi_dashboard_embedding_approved_domains": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingApprovedDomains_SdkV2{}.Type(ctx),
			},
			"effective_automatic_cluster_update_workspace": basetypes.ListType{
				ElemType: ClusterAutoRestartMessage_SdkV2{}.Type(ctx),
			},
			"effective_boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"effective_default_data_security_mode": basetypes.ListType{
				ElemType: DefaultDataSecurityModeMessage_SdkV2{}.Type(ctx),
			},
			"effective_integer_val": basetypes.ListType{
				ElemType: IntegerMessage_SdkV2{}.Type(ctx),
			},
			"effective_personal_compute": basetypes.ListType{
				ElemType: PersonalComputeMessage_SdkV2{}.Type(ctx),
			},
			"effective_restrict_workspace_admins": basetypes.ListType{
				ElemType: RestrictWorkspaceAdminsMessage_SdkV2{}.Type(ctx),
			},
			"effective_string_val": basetypes.ListType{
				ElemType: StringMessage_SdkV2{}.Type(ctx),
			},
			"integer_val": basetypes.ListType{
				ElemType: IntegerMessage_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"personal_compute": basetypes.ListType{
				ElemType: PersonalComputeMessage_SdkV2{}.Type(ctx),
			},
			"restrict_workspace_admins": basetypes.ListType{
				ElemType: RestrictWorkspaceAdminsMessage_SdkV2{}.Type(ctx),
			},
			"string_val": basetypes.ListType{
				ElemType: StringMessage_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAibiDashboardEmbeddingAccessPolicy returns the value of the AibiDashboardEmbeddingAccessPolicy field in Setting_SdkV2 as
// a AibiDashboardEmbeddingAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (AibiDashboardEmbeddingAccessPolicy_SdkV2, bool) {
	var e AibiDashboardEmbeddingAccessPolicy_SdkV2
	if o.AibiDashboardEmbeddingAccessPolicy.IsNull() || o.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingAccessPolicy_SdkV2
	d := o.AibiDashboardEmbeddingAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAibiDashboardEmbeddingAccessPolicy sets the value of the AibiDashboardEmbeddingAccessPolicy field in Setting_SdkV2.
func (o *Setting_SdkV2) SetAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v AibiDashboardEmbeddingAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aibi_dashboard_embedding_access_policy"]
	o.AibiDashboardEmbeddingAccessPolicy = types.ListValueMust(t, vs)
}

// GetAibiDashboardEmbeddingApprovedDomains returns the value of the AibiDashboardEmbeddingApprovedDomains field in Setting_SdkV2 as
// a AibiDashboardEmbeddingApprovedDomains_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (AibiDashboardEmbeddingApprovedDomains_SdkV2, bool) {
	var e AibiDashboardEmbeddingApprovedDomains_SdkV2
	if o.AibiDashboardEmbeddingApprovedDomains.IsNull() || o.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingApprovedDomains_SdkV2
	d := o.AibiDashboardEmbeddingApprovedDomains.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAibiDashboardEmbeddingApprovedDomains sets the value of the AibiDashboardEmbeddingApprovedDomains field in Setting_SdkV2.
func (o *Setting_SdkV2) SetAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v AibiDashboardEmbeddingApprovedDomains_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aibi_dashboard_embedding_approved_domains"]
	o.AibiDashboardEmbeddingApprovedDomains = types.ListValueMust(t, vs)
}

// GetAutomaticClusterUpdateWorkspace returns the value of the AutomaticClusterUpdateWorkspace field in Setting_SdkV2 as
// a ClusterAutoRestartMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetAutomaticClusterUpdateWorkspace(ctx context.Context) (ClusterAutoRestartMessage_SdkV2, bool) {
	var e ClusterAutoRestartMessage_SdkV2
	if o.AutomaticClusterUpdateWorkspace.IsNull() || o.AutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessage_SdkV2
	d := o.AutomaticClusterUpdateWorkspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutomaticClusterUpdateWorkspace sets the value of the AutomaticClusterUpdateWorkspace field in Setting_SdkV2.
func (o *Setting_SdkV2) SetAutomaticClusterUpdateWorkspace(ctx context.Context, v ClusterAutoRestartMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["automatic_cluster_update_workspace"]
	o.AutomaticClusterUpdateWorkspace = types.ListValueMust(t, vs)
}

// GetBooleanVal returns the value of the BooleanVal field in Setting_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := o.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	o.BooleanVal = types.ListValueMust(t, vs)
}

// GetDefaultDataSecurityMode returns the value of the DefaultDataSecurityMode field in Setting_SdkV2 as
// a DefaultDataSecurityModeMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetDefaultDataSecurityMode(ctx context.Context) (DefaultDataSecurityModeMessage_SdkV2, bool) {
	var e DefaultDataSecurityModeMessage_SdkV2
	if o.DefaultDataSecurityMode.IsNull() || o.DefaultDataSecurityMode.IsUnknown() {
		return e, false
	}
	var v []DefaultDataSecurityModeMessage_SdkV2
	d := o.DefaultDataSecurityMode.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultDataSecurityMode sets the value of the DefaultDataSecurityMode field in Setting_SdkV2.
func (o *Setting_SdkV2) SetDefaultDataSecurityMode(ctx context.Context, v DefaultDataSecurityModeMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_data_security_mode"]
	o.DefaultDataSecurityMode = types.ListValueMust(t, vs)
}

// GetEffectiveAibiDashboardEmbeddingAccessPolicy returns the value of the EffectiveAibiDashboardEmbeddingAccessPolicy field in Setting_SdkV2 as
// a AibiDashboardEmbeddingAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (AibiDashboardEmbeddingAccessPolicy_SdkV2, bool) {
	var e AibiDashboardEmbeddingAccessPolicy_SdkV2
	if o.EffectiveAibiDashboardEmbeddingAccessPolicy.IsNull() || o.EffectiveAibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingAccessPolicy_SdkV2
	d := o.EffectiveAibiDashboardEmbeddingAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveAibiDashboardEmbeddingAccessPolicy sets the value of the EffectiveAibiDashboardEmbeddingAccessPolicy field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v AibiDashboardEmbeddingAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_aibi_dashboard_embedding_access_policy"]
	o.EffectiveAibiDashboardEmbeddingAccessPolicy = types.ListValueMust(t, vs)
}

// GetEffectiveAibiDashboardEmbeddingApprovedDomains returns the value of the EffectiveAibiDashboardEmbeddingApprovedDomains field in Setting_SdkV2 as
// a AibiDashboardEmbeddingApprovedDomains_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (AibiDashboardEmbeddingApprovedDomains_SdkV2, bool) {
	var e AibiDashboardEmbeddingApprovedDomains_SdkV2
	if o.EffectiveAibiDashboardEmbeddingApprovedDomains.IsNull() || o.EffectiveAibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingApprovedDomains_SdkV2
	d := o.EffectiveAibiDashboardEmbeddingApprovedDomains.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveAibiDashboardEmbeddingApprovedDomains sets the value of the EffectiveAibiDashboardEmbeddingApprovedDomains field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v AibiDashboardEmbeddingApprovedDomains_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_aibi_dashboard_embedding_approved_domains"]
	o.EffectiveAibiDashboardEmbeddingApprovedDomains = types.ListValueMust(t, vs)
}

// GetEffectiveAutomaticClusterUpdateWorkspace returns the value of the EffectiveAutomaticClusterUpdateWorkspace field in Setting_SdkV2 as
// a ClusterAutoRestartMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveAutomaticClusterUpdateWorkspace(ctx context.Context) (ClusterAutoRestartMessage_SdkV2, bool) {
	var e ClusterAutoRestartMessage_SdkV2
	if o.EffectiveAutomaticClusterUpdateWorkspace.IsNull() || o.EffectiveAutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessage_SdkV2
	d := o.EffectiveAutomaticClusterUpdateWorkspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveAutomaticClusterUpdateWorkspace sets the value of the EffectiveAutomaticClusterUpdateWorkspace field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveAutomaticClusterUpdateWorkspace(ctx context.Context, v ClusterAutoRestartMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_automatic_cluster_update_workspace"]
	o.EffectiveAutomaticClusterUpdateWorkspace = types.ListValueMust(t, vs)
}

// GetEffectiveBooleanVal returns the value of the EffectiveBooleanVal field in Setting_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if o.EffectiveBooleanVal.IsNull() || o.EffectiveBooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := o.EffectiveBooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveBooleanVal sets the value of the EffectiveBooleanVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_boolean_val"]
	o.EffectiveBooleanVal = types.ListValueMust(t, vs)
}

// GetEffectiveDefaultDataSecurityMode returns the value of the EffectiveDefaultDataSecurityMode field in Setting_SdkV2 as
// a DefaultDataSecurityModeMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveDefaultDataSecurityMode(ctx context.Context) (DefaultDataSecurityModeMessage_SdkV2, bool) {
	var e DefaultDataSecurityModeMessage_SdkV2
	if o.EffectiveDefaultDataSecurityMode.IsNull() || o.EffectiveDefaultDataSecurityMode.IsUnknown() {
		return e, false
	}
	var v []DefaultDataSecurityModeMessage_SdkV2
	d := o.EffectiveDefaultDataSecurityMode.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveDefaultDataSecurityMode sets the value of the EffectiveDefaultDataSecurityMode field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveDefaultDataSecurityMode(ctx context.Context, v DefaultDataSecurityModeMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_default_data_security_mode"]
	o.EffectiveDefaultDataSecurityMode = types.ListValueMust(t, vs)
}

// GetEffectiveIntegerVal returns the value of the EffectiveIntegerVal field in Setting_SdkV2 as
// a IntegerMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveIntegerVal(ctx context.Context) (IntegerMessage_SdkV2, bool) {
	var e IntegerMessage_SdkV2
	if o.EffectiveIntegerVal.IsNull() || o.EffectiveIntegerVal.IsUnknown() {
		return e, false
	}
	var v []IntegerMessage_SdkV2
	d := o.EffectiveIntegerVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveIntegerVal sets the value of the EffectiveIntegerVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveIntegerVal(ctx context.Context, v IntegerMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_integer_val"]
	o.EffectiveIntegerVal = types.ListValueMust(t, vs)
}

// GetEffectivePersonalCompute returns the value of the EffectivePersonalCompute field in Setting_SdkV2 as
// a PersonalComputeMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectivePersonalCompute(ctx context.Context) (PersonalComputeMessage_SdkV2, bool) {
	var e PersonalComputeMessage_SdkV2
	if o.EffectivePersonalCompute.IsNull() || o.EffectivePersonalCompute.IsUnknown() {
		return e, false
	}
	var v []PersonalComputeMessage_SdkV2
	d := o.EffectivePersonalCompute.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectivePersonalCompute sets the value of the EffectivePersonalCompute field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectivePersonalCompute(ctx context.Context, v PersonalComputeMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_personal_compute"]
	o.EffectivePersonalCompute = types.ListValueMust(t, vs)
}

// GetEffectiveRestrictWorkspaceAdmins returns the value of the EffectiveRestrictWorkspaceAdmins field in Setting_SdkV2 as
// a RestrictWorkspaceAdminsMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveRestrictWorkspaceAdmins(ctx context.Context) (RestrictWorkspaceAdminsMessage_SdkV2, bool) {
	var e RestrictWorkspaceAdminsMessage_SdkV2
	if o.EffectiveRestrictWorkspaceAdmins.IsNull() || o.EffectiveRestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v []RestrictWorkspaceAdminsMessage_SdkV2
	d := o.EffectiveRestrictWorkspaceAdmins.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveRestrictWorkspaceAdmins sets the value of the EffectiveRestrictWorkspaceAdmins field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveRestrictWorkspaceAdmins(ctx context.Context, v RestrictWorkspaceAdminsMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_restrict_workspace_admins"]
	o.EffectiveRestrictWorkspaceAdmins = types.ListValueMust(t, vs)
}

// GetEffectiveStringVal returns the value of the EffectiveStringVal field in Setting_SdkV2 as
// a StringMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveStringVal(ctx context.Context) (StringMessage_SdkV2, bool) {
	var e StringMessage_SdkV2
	if o.EffectiveStringVal.IsNull() || o.EffectiveStringVal.IsUnknown() {
		return e, false
	}
	var v []StringMessage_SdkV2
	d := o.EffectiveStringVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveStringVal sets the value of the EffectiveStringVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveStringVal(ctx context.Context, v StringMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_string_val"]
	o.EffectiveStringVal = types.ListValueMust(t, vs)
}

// GetIntegerVal returns the value of the IntegerVal field in Setting_SdkV2 as
// a IntegerMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetIntegerVal(ctx context.Context) (IntegerMessage_SdkV2, bool) {
	var e IntegerMessage_SdkV2
	if o.IntegerVal.IsNull() || o.IntegerVal.IsUnknown() {
		return e, false
	}
	var v []IntegerMessage_SdkV2
	d := o.IntegerVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIntegerVal sets the value of the IntegerVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetIntegerVal(ctx context.Context, v IntegerMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["integer_val"]
	o.IntegerVal = types.ListValueMust(t, vs)
}

// GetPersonalCompute returns the value of the PersonalCompute field in Setting_SdkV2 as
// a PersonalComputeMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetPersonalCompute(ctx context.Context) (PersonalComputeMessage_SdkV2, bool) {
	var e PersonalComputeMessage_SdkV2
	if o.PersonalCompute.IsNull() || o.PersonalCompute.IsUnknown() {
		return e, false
	}
	var v []PersonalComputeMessage_SdkV2
	d := o.PersonalCompute.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPersonalCompute sets the value of the PersonalCompute field in Setting_SdkV2.
func (o *Setting_SdkV2) SetPersonalCompute(ctx context.Context, v PersonalComputeMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["personal_compute"]
	o.PersonalCompute = types.ListValueMust(t, vs)
}

// GetRestrictWorkspaceAdmins returns the value of the RestrictWorkspaceAdmins field in Setting_SdkV2 as
// a RestrictWorkspaceAdminsMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetRestrictWorkspaceAdmins(ctx context.Context) (RestrictWorkspaceAdminsMessage_SdkV2, bool) {
	var e RestrictWorkspaceAdminsMessage_SdkV2
	if o.RestrictWorkspaceAdmins.IsNull() || o.RestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v []RestrictWorkspaceAdminsMessage_SdkV2
	d := o.RestrictWorkspaceAdmins.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRestrictWorkspaceAdmins sets the value of the RestrictWorkspaceAdmins field in Setting_SdkV2.
func (o *Setting_SdkV2) SetRestrictWorkspaceAdmins(ctx context.Context, v RestrictWorkspaceAdminsMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["restrict_workspace_admins"]
	o.RestrictWorkspaceAdmins = types.ListValueMust(t, vs)
}

// GetStringVal returns the value of the StringVal field in Setting_SdkV2 as
// a StringMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetStringVal(ctx context.Context) (StringMessage_SdkV2, bool) {
	var e StringMessage_SdkV2
	if o.StringVal.IsNull() || o.StringVal.IsUnknown() {
		return e, false
	}
	var v []StringMessage_SdkV2
	d := o.StringVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStringVal sets the value of the StringVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetStringVal(ctx context.Context, v StringMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["string_val"]
	o.StringVal = types.ListValueMust(t, vs)
}

type SettingsMetadata_SdkV2 struct {
	// Setting description for what this setting controls
	Description types.String `tfsdk:"description"`
	// Link to databricks documentation for the setting
	DocsLink types.String `tfsdk:"docs_link"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`
	// Type of the setting. To set this setting, the value sent must match this
	// type.
	Type_ types.String `tfsdk:"type"`
}

func (toState *SettingsMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SettingsMetadata_SdkV2) {
}

func (toState *SettingsMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SettingsMetadata_SdkV2) {
}

func (c SettingsMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["docs_link"] = attrs["docs_link"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SettingsMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SettingsMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SettingsMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (o SettingsMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"docs_link":   o.DocsLink,
			"name":        o.Name,
			"type":        o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SettingsMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"docs_link":   types.StringType,
			"name":        types.StringType,
			"type":        types.StringType,
		},
	}
}

type StringMessage_SdkV2 struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value"`
}

func (toState *StringMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan StringMessage_SdkV2) {
}

func (toState *StringMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState StringMessage_SdkV2) {
}

func (c StringMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StringMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StringMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StringMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o StringMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StringMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}
