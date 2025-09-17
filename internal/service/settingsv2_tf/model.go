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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AibiDashboardEmbeddingAccessPolicy struct {
	AccessPolicyType types.String `tfsdk:"access_policy_type"`
}

func (toState *AibiDashboardEmbeddingAccessPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AibiDashboardEmbeddingAccessPolicy) {
}

func (toState *AibiDashboardEmbeddingAccessPolicy) SyncFieldsDuringRead(ctx context.Context, fromState AibiDashboardEmbeddingAccessPolicy) {
}

func (c AibiDashboardEmbeddingAccessPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicy
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingAccessPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_policy_type": o.AccessPolicyType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingAccessPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_policy_type": types.StringType,
		},
	}
}

type AibiDashboardEmbeddingApprovedDomains struct {
	ApprovedDomains types.List `tfsdk:"approved_domains"`
}

func (toState *AibiDashboardEmbeddingApprovedDomains) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AibiDashboardEmbeddingApprovedDomains) {
}

func (toState *AibiDashboardEmbeddingApprovedDomains) SyncFieldsDuringRead(ctx context.Context, fromState AibiDashboardEmbeddingApprovedDomains) {
}

func (c AibiDashboardEmbeddingApprovedDomains) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingApprovedDomains) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"approved_domains": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomains
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingApprovedDomains) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"approved_domains": o.ApprovedDomains,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingApprovedDomains) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"approved_domains": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetApprovedDomains returns the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AibiDashboardEmbeddingApprovedDomains) GetApprovedDomains(ctx context.Context) ([]types.String, bool) {
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

// SetApprovedDomains sets the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains.
func (o *AibiDashboardEmbeddingApprovedDomains) SetApprovedDomains(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["approved_domains"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ApprovedDomains = types.ListValueMust(t, vs)
}

type BooleanMessage struct {
	Value types.Bool `tfsdk:"value"`
}

func (toState *BooleanMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan BooleanMessage) {
}

func (toState *BooleanMessage) SyncFieldsDuringRead(ctx context.Context, fromState BooleanMessage) {
}

func (c BooleanMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a BooleanMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BooleanMessage
// only implements ToObjectValue() and Type().
func (o BooleanMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BooleanMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.BoolType,
		},
	}
}

type ClusterAutoRestartMessage struct {
	CanToggle types.Bool `tfsdk:"can_toggle"`

	Enabled types.Bool `tfsdk:"enabled"`

	EnablementDetails types.Object `tfsdk:"enablement_details"`

	MaintenanceWindow types.Object `tfsdk:"maintenance_window"`

	RestartEvenIfNoUpdatesAvailable types.Bool `tfsdk:"restart_even_if_no_updates_available"`
}

func (toState *ClusterAutoRestartMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessage) {
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

func (toState *ClusterAutoRestartMessage) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessage) {
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

func (c ClusterAutoRestartMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["can_toggle"] = attrs["can_toggle"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["enablement_details"] = attrs["enablement_details"].SetOptional()
	attrs["maintenance_window"] = attrs["maintenance_window"].SetOptional()
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
func (a ClusterAutoRestartMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enablement_details": reflect.TypeOf(ClusterAutoRestartMessageEnablementDetails{}),
		"maintenance_window": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindow{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessage
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterAutoRestartMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"can_toggle":                           types.BoolType,
			"enabled":                              types.BoolType,
			"enablement_details":                   ClusterAutoRestartMessageEnablementDetails{}.Type(ctx),
			"maintenance_window":                   ClusterAutoRestartMessageMaintenanceWindow{}.Type(ctx),
			"restart_even_if_no_updates_available": types.BoolType,
		},
	}
}

// GetEnablementDetails returns the value of the EnablementDetails field in ClusterAutoRestartMessage as
// a ClusterAutoRestartMessageEnablementDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessage) GetEnablementDetails(ctx context.Context) (ClusterAutoRestartMessageEnablementDetails, bool) {
	var e ClusterAutoRestartMessageEnablementDetails
	if o.EnablementDetails.IsNull() || o.EnablementDetails.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessageEnablementDetails
	d := o.EnablementDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnablementDetails sets the value of the EnablementDetails field in ClusterAutoRestartMessage.
func (o *ClusterAutoRestartMessage) SetEnablementDetails(ctx context.Context, v ClusterAutoRestartMessageEnablementDetails) {
	vs := v.ToObjectValue(ctx)
	o.EnablementDetails = vs
}

// GetMaintenanceWindow returns the value of the MaintenanceWindow field in ClusterAutoRestartMessage as
// a ClusterAutoRestartMessageMaintenanceWindow value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessage) GetMaintenanceWindow(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindow, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindow
	if o.MaintenanceWindow.IsNull() || o.MaintenanceWindow.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessageMaintenanceWindow
	d := o.MaintenanceWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMaintenanceWindow sets the value of the MaintenanceWindow field in ClusterAutoRestartMessage.
func (o *ClusterAutoRestartMessage) SetMaintenanceWindow(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindow) {
	vs := v.ToObjectValue(ctx)
	o.MaintenanceWindow = vs
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

func (toState *ClusterAutoRestartMessageEnablementDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessageEnablementDetails) {
}

func (toState *ClusterAutoRestartMessageEnablementDetails) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessageEnablementDetails) {
}

func (c ClusterAutoRestartMessageEnablementDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageEnablementDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageEnablementDetails
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageEnablementDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"forced_for_compliance_mode":           o.ForcedForComplianceMode,
			"unavailable_for_disabled_entitlement": o.UnavailableForDisabledEntitlement,
			"unavailable_for_non_enterprise_tier":  o.UnavailableForNonEnterpriseTier,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageEnablementDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"forced_for_compliance_mode":           types.BoolType,
			"unavailable_for_disabled_entitlement": types.BoolType,
			"unavailable_for_non_enterprise_tier":  types.BoolType,
		},
	}
}

type ClusterAutoRestartMessageMaintenanceWindow struct {
	WeekDayBasedSchedule types.Object `tfsdk:"week_day_based_schedule"`
}

func (toState *ClusterAutoRestartMessageMaintenanceWindow) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessageMaintenanceWindow) {
	if !fromPlan.WeekDayBasedSchedule.IsNull() && !fromPlan.WeekDayBasedSchedule.IsUnknown() {
		if toStateWeekDayBasedSchedule, ok := toState.GetWeekDayBasedSchedule(ctx); ok {
			if fromPlanWeekDayBasedSchedule, ok := fromPlan.GetWeekDayBasedSchedule(ctx); ok {
				toStateWeekDayBasedSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanWeekDayBasedSchedule)
				toState.SetWeekDayBasedSchedule(ctx, toStateWeekDayBasedSchedule)
			}
		}
	}
}

func (toState *ClusterAutoRestartMessageMaintenanceWindow) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessageMaintenanceWindow) {
	if !fromState.WeekDayBasedSchedule.IsNull() && !fromState.WeekDayBasedSchedule.IsUnknown() {
		if toStateWeekDayBasedSchedule, ok := toState.GetWeekDayBasedSchedule(ctx); ok {
			if fromStateWeekDayBasedSchedule, ok := fromState.GetWeekDayBasedSchedule(ctx); ok {
				toStateWeekDayBasedSchedule.SyncFieldsDuringRead(ctx, fromStateWeekDayBasedSchedule)
				toState.SetWeekDayBasedSchedule(ctx, toStateWeekDayBasedSchedule)
			}
		}
	}
}

func (c ClusterAutoRestartMessageMaintenanceWindow) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["week_day_based_schedule"] = attrs["week_day_based_schedule"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindow
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindow) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"week_day_based_schedule": o.WeekDayBasedSchedule,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindow) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"week_day_based_schedule": ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}.Type(ctx),
		},
	}
}

// GetWeekDayBasedSchedule returns the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow as
// a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessageMaintenanceWindow) GetWeekDayBasedSchedule(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
	if o.WeekDayBasedSchedule.IsNull() || o.WeekDayBasedSchedule.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
	d := o.WeekDayBasedSchedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWeekDayBasedSchedule sets the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow.
func (o *ClusterAutoRestartMessageMaintenanceWindow) SetWeekDayBasedSchedule(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
	vs := v.ToObjectValue(ctx)
	o.WeekDayBasedSchedule = vs
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	DayOfWeek types.String `tfsdk:"day_of_week"`

	Frequency types.String `tfsdk:"frequency"`

	WindowStartTime types.Object `tfsdk:"window_start_time"`
}

func (toState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
	if !fromPlan.WindowStartTime.IsNull() && !fromPlan.WindowStartTime.IsUnknown() {
		if toStateWindowStartTime, ok := toState.GetWindowStartTime(ctx); ok {
			if fromPlanWindowStartTime, ok := fromPlan.GetWindowStartTime(ctx); ok {
				toStateWindowStartTime.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanWindowStartTime)
				toState.SetWindowStartTime(ctx, toStateWindowStartTime)
			}
		}
	}
}

func (toState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
	if !fromState.WindowStartTime.IsNull() && !fromState.WindowStartTime.IsUnknown() {
		if toStateWindowStartTime, ok := toState.GetWindowStartTime(ctx); ok {
			if fromStateWindowStartTime, ok := fromState.GetWindowStartTime(ctx); ok {
				toStateWindowStartTime.SyncFieldsDuringRead(ctx, fromStateWindowStartTime)
				toState.SetWindowStartTime(ctx, toStateWindowStartTime)
			}
		}
	}
}

func (c ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["day_of_week"] = attrs["day_of_week"].SetOptional()
	attrs["frequency"] = attrs["frequency"].SetOptional()
	attrs["window_start_time"] = attrs["window_start_time"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"day_of_week":       o.DayOfWeek,
			"frequency":         o.Frequency,
			"window_start_time": o.WindowStartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"day_of_week":       types.StringType,
			"frequency":         types.StringType,
			"window_start_time": ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}.Type(ctx),
		},
	}
}

// GetWindowStartTime returns the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule as
// a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) GetWindowStartTime(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
	if o.WindowStartTime.IsNull() || o.WindowStartTime.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
	d := o.WindowStartTime.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWindowStartTime sets the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule.
func (o *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SetWindowStartTime(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
	vs := v.ToObjectValue(ctx)
	o.WindowStartTime = vs
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	Hours types.Int64 `tfsdk:"hours"`

	Minutes types.Int64 `tfsdk:"minutes"`
}

func (toState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
}

func (toState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) SyncFieldsDuringRead(ctx context.Context, fromState ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
}

func (c ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"hours":   o.Hours,
			"minutes": o.Minutes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) Type(ctx context.Context) attr.Type {
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
type DefaultDataSecurityModeMessage struct {
	Status types.String `tfsdk:"status"`
}

func (toState *DefaultDataSecurityModeMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DefaultDataSecurityModeMessage) {
}

func (toState *DefaultDataSecurityModeMessage) SyncFieldsDuringRead(ctx context.Context, fromState DefaultDataSecurityModeMessage) {
}

func (c DefaultDataSecurityModeMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DefaultDataSecurityModeMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultDataSecurityModeMessage
// only implements ToObjectValue() and Type().
func (o DefaultDataSecurityModeMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DefaultDataSecurityModeMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": types.StringType,
		},
	}
}

type GetPublicAccountSettingRequest struct {
	Name types.String `tfsdk:"-"`
}

func (toState *GetPublicAccountSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetPublicAccountSettingRequest) {
}

func (toState *GetPublicAccountSettingRequest) SyncFieldsDuringRead(ctx context.Context, fromState GetPublicAccountSettingRequest) {
}

func (c GetPublicAccountSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublicAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublicAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublicAccountSettingRequest
// only implements ToObjectValue() and Type().
func (o GetPublicAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublicAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetPublicWorkspaceSettingRequest struct {
	Name types.String `tfsdk:"-"`
}

func (toState *GetPublicWorkspaceSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetPublicWorkspaceSettingRequest) {
}

func (toState *GetPublicWorkspaceSettingRequest) SyncFieldsDuringRead(ctx context.Context, fromState GetPublicWorkspaceSettingRequest) {
}

func (c GetPublicWorkspaceSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublicWorkspaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublicWorkspaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublicWorkspaceSettingRequest
// only implements ToObjectValue() and Type().
func (o GetPublicWorkspaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublicWorkspaceSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type IntegerMessage struct {
	Value types.Int64 `tfsdk:"value"`
}

func (toState *IntegerMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan IntegerMessage) {
}

func (toState *IntegerMessage) SyncFieldsDuringRead(ctx context.Context, fromState IntegerMessage) {
}

func (c IntegerMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a IntegerMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IntegerMessage
// only implements ToObjectValue() and Type().
func (o IntegerMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IntegerMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.Int64Type,
		},
	}
}

type ListAccountSettingsMetadataRequest struct {
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

func (toState *ListAccountSettingsMetadataRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListAccountSettingsMetadataRequest) {
}

func (toState *ListAccountSettingsMetadataRequest) SyncFieldsDuringRead(ctx context.Context, fromState ListAccountSettingsMetadataRequest) {
}

func (c ListAccountSettingsMetadataRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountSettingsMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountSettingsMetadataRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountSettingsMetadataRequest
// only implements ToObjectValue() and Type().
func (o ListAccountSettingsMetadataRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountSettingsMetadataRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAccountSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata types.List `tfsdk:"settings_metadata"`
}

func (toState *ListAccountSettingsMetadataResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListAccountSettingsMetadataResponse) {
}

func (toState *ListAccountSettingsMetadataResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListAccountSettingsMetadataResponse) {
}

func (c ListAccountSettingsMetadataResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAccountSettingsMetadataResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings_metadata": reflect.TypeOf(SettingsMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountSettingsMetadataResponse
// only implements ToObjectValue() and Type().
func (o ListAccountSettingsMetadataResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"settings_metadata": o.SettingsMetadata,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountSettingsMetadataResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"settings_metadata": basetypes.ListType{
				ElemType: SettingsMetadata{}.Type(ctx),
			},
		},
	}
}

// GetSettingsMetadata returns the value of the SettingsMetadata field in ListAccountSettingsMetadataResponse as
// a slice of SettingsMetadata values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAccountSettingsMetadataResponse) GetSettingsMetadata(ctx context.Context) ([]SettingsMetadata, bool) {
	if o.SettingsMetadata.IsNull() || o.SettingsMetadata.IsUnknown() {
		return nil, false
	}
	var v []SettingsMetadata
	d := o.SettingsMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettingsMetadata sets the value of the SettingsMetadata field in ListAccountSettingsMetadataResponse.
func (o *ListAccountSettingsMetadataResponse) SetSettingsMetadata(ctx context.Context, v []SettingsMetadata) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings_metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SettingsMetadata = types.ListValueMust(t, vs)
}

type ListWorkspaceSettingsMetadataRequest struct {
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

func (toState *ListWorkspaceSettingsMetadataRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListWorkspaceSettingsMetadataRequest) {
}

func (toState *ListWorkspaceSettingsMetadataRequest) SyncFieldsDuringRead(ctx context.Context, fromState ListWorkspaceSettingsMetadataRequest) {
}

func (c ListWorkspaceSettingsMetadataRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceSettingsMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspaceSettingsMetadataRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceSettingsMetadataRequest
// only implements ToObjectValue() and Type().
func (o ListWorkspaceSettingsMetadataRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceSettingsMetadataRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListWorkspaceSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata types.List `tfsdk:"settings_metadata"`
}

func (toState *ListWorkspaceSettingsMetadataResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListWorkspaceSettingsMetadataResponse) {
}

func (toState *ListWorkspaceSettingsMetadataResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListWorkspaceSettingsMetadataResponse) {
}

func (c ListWorkspaceSettingsMetadataResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListWorkspaceSettingsMetadataResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings_metadata": reflect.TypeOf(SettingsMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceSettingsMetadataResponse
// only implements ToObjectValue() and Type().
func (o ListWorkspaceSettingsMetadataResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"settings_metadata": o.SettingsMetadata,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceSettingsMetadataResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"settings_metadata": basetypes.ListType{
				ElemType: SettingsMetadata{}.Type(ctx),
			},
		},
	}
}

// GetSettingsMetadata returns the value of the SettingsMetadata field in ListWorkspaceSettingsMetadataResponse as
// a slice of SettingsMetadata values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListWorkspaceSettingsMetadataResponse) GetSettingsMetadata(ctx context.Context) ([]SettingsMetadata, bool) {
	if o.SettingsMetadata.IsNull() || o.SettingsMetadata.IsUnknown() {
		return nil, false
	}
	var v []SettingsMetadata
	d := o.SettingsMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettingsMetadata sets the value of the SettingsMetadata field in ListWorkspaceSettingsMetadataResponse.
func (o *ListWorkspaceSettingsMetadataResponse) SetSettingsMetadata(ctx context.Context, v []SettingsMetadata) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings_metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SettingsMetadata = types.ListValueMust(t, vs)
}

type PatchPublicAccountSettingRequest struct {
	Name types.String `tfsdk:"-"`

	Setting types.Object `tfsdk:"setting"`
}

func (toState *PatchPublicAccountSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PatchPublicAccountSettingRequest) {
	if !fromPlan.Setting.IsNull() && !fromPlan.Setting.IsUnknown() {
		if toStateSetting, ok := toState.GetSetting(ctx); ok {
			if fromPlanSetting, ok := fromPlan.GetSetting(ctx); ok {
				toStateSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSetting)
				toState.SetSetting(ctx, toStateSetting)
			}
		}
	}
}

func (toState *PatchPublicAccountSettingRequest) SyncFieldsDuringRead(ctx context.Context, fromState PatchPublicAccountSettingRequest) {
	if !fromState.Setting.IsNull() && !fromState.Setting.IsUnknown() {
		if toStateSetting, ok := toState.GetSetting(ctx); ok {
			if fromStateSetting, ok := fromState.GetSetting(ctx); ok {
				toStateSetting.SyncFieldsDuringRead(ctx, fromStateSetting)
				toState.SetSetting(ctx, toStateSetting)
			}
		}
	}
}

func (c PatchPublicAccountSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchPublicAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchPublicAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(Setting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchPublicAccountSettingRequest
// only implements ToObjectValue() and Type().
func (o PatchPublicAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"setting": o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchPublicAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"setting": Setting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in PatchPublicAccountSettingRequest as
// a Setting value.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchPublicAccountSettingRequest) GetSetting(ctx context.Context) (Setting, bool) {
	var e Setting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v Setting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in PatchPublicAccountSettingRequest.
func (o *PatchPublicAccountSettingRequest) SetSetting(ctx context.Context, v Setting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
}

type PatchPublicWorkspaceSettingRequest struct {
	Name types.String `tfsdk:"-"`

	Setting types.Object `tfsdk:"setting"`
}

func (toState *PatchPublicWorkspaceSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PatchPublicWorkspaceSettingRequest) {
	if !fromPlan.Setting.IsNull() && !fromPlan.Setting.IsUnknown() {
		if toStateSetting, ok := toState.GetSetting(ctx); ok {
			if fromPlanSetting, ok := fromPlan.GetSetting(ctx); ok {
				toStateSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSetting)
				toState.SetSetting(ctx, toStateSetting)
			}
		}
	}
}

func (toState *PatchPublicWorkspaceSettingRequest) SyncFieldsDuringRead(ctx context.Context, fromState PatchPublicWorkspaceSettingRequest) {
	if !fromState.Setting.IsNull() && !fromState.Setting.IsUnknown() {
		if toStateSetting, ok := toState.GetSetting(ctx); ok {
			if fromStateSetting, ok := fromState.GetSetting(ctx); ok {
				toStateSetting.SyncFieldsDuringRead(ctx, fromStateSetting)
				toState.SetSetting(ctx, toStateSetting)
			}
		}
	}
}

func (c PatchPublicWorkspaceSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchPublicWorkspaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchPublicWorkspaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(Setting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchPublicWorkspaceSettingRequest
// only implements ToObjectValue() and Type().
func (o PatchPublicWorkspaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"setting": o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchPublicWorkspaceSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"setting": Setting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in PatchPublicWorkspaceSettingRequest as
// a Setting value.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchPublicWorkspaceSettingRequest) GetSetting(ctx context.Context) (Setting, bool) {
	var e Setting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v Setting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in PatchPublicWorkspaceSettingRequest.
func (o *PatchPublicWorkspaceSettingRequest) SetSetting(ctx context.Context, v Setting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
}

type PersonalComputeMessage struct {
	Value types.String `tfsdk:"value"`
}

func (toState *PersonalComputeMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PersonalComputeMessage) {
}

func (toState *PersonalComputeMessage) SyncFieldsDuringRead(ctx context.Context, fromState PersonalComputeMessage) {
}

func (c PersonalComputeMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PersonalComputeMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeMessage
// only implements ToObjectValue() and Type().
func (o PersonalComputeMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PersonalComputeMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type RestrictWorkspaceAdminsMessage struct {
	Status types.String `tfsdk:"status"`
}

func (toState *RestrictWorkspaceAdminsMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RestrictWorkspaceAdminsMessage) {
}

func (toState *RestrictWorkspaceAdminsMessage) SyncFieldsDuringRead(ctx context.Context, fromState RestrictWorkspaceAdminsMessage) {
}

func (c RestrictWorkspaceAdminsMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestrictWorkspaceAdminsMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsMessage
// only implements ToObjectValue() and Type().
func (o RestrictWorkspaceAdminsMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestrictWorkspaceAdminsMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": types.StringType,
		},
	}
}

type Setting struct {
	AibiDashboardEmbeddingAccessPolicy types.Object `tfsdk:"aibi_dashboard_embedding_access_policy"`

	AibiDashboardEmbeddingApprovedDomains types.Object `tfsdk:"aibi_dashboard_embedding_approved_domains"`
	// todo: Mark these Public after onboarded to DSL
	AutomaticClusterUpdateWorkspace types.Object `tfsdk:"automatic_cluster_update_workspace"`

	BooleanVal types.Object `tfsdk:"boolean_val"`

	DefaultDataSecurityMode types.Object `tfsdk:"default_data_security_mode"`

	EffectiveAibiDashboardEmbeddingAccessPolicy types.Object `tfsdk:"effective_aibi_dashboard_embedding_access_policy"`

	EffectiveAibiDashboardEmbeddingApprovedDomains types.Object `tfsdk:"effective_aibi_dashboard_embedding_approved_domains"`

	EffectiveAutomaticClusterUpdateWorkspace types.Object `tfsdk:"effective_automatic_cluster_update_workspace"`

	EffectiveBooleanVal types.Object `tfsdk:"effective_boolean_val"`

	EffectiveDefaultDataSecurityMode types.Object `tfsdk:"effective_default_data_security_mode"`

	EffectiveIntegerVal types.Object `tfsdk:"effective_integer_val"`

	EffectivePersonalCompute types.Object `tfsdk:"effective_personal_compute"`

	EffectiveRestrictWorkspaceAdmins types.Object `tfsdk:"effective_restrict_workspace_admins"`

	EffectiveStringVal types.Object `tfsdk:"effective_string_val"`

	IntegerVal types.Object `tfsdk:"integer_val"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`

	PersonalCompute types.Object `tfsdk:"personal_compute"`

	RestrictWorkspaceAdmins types.Object `tfsdk:"restrict_workspace_admins"`

	StringVal types.Object `tfsdk:"string_val"`
}

func (toState *Setting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Setting) {
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

func (toState *Setting) SyncFieldsDuringRead(ctx context.Context, fromState Setting) {
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

func (c Setting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aibi_dashboard_embedding_access_policy"] = attrs["aibi_dashboard_embedding_access_policy"].SetOptional()
	attrs["aibi_dashboard_embedding_approved_domains"] = attrs["aibi_dashboard_embedding_approved_domains"].SetOptional()
	attrs["automatic_cluster_update_workspace"] = attrs["automatic_cluster_update_workspace"].SetOptional()
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["default_data_security_mode"] = attrs["default_data_security_mode"].SetOptional()
	attrs["effective_aibi_dashboard_embedding_access_policy"] = attrs["effective_aibi_dashboard_embedding_access_policy"].SetOptional()
	attrs["effective_aibi_dashboard_embedding_approved_domains"] = attrs["effective_aibi_dashboard_embedding_approved_domains"].SetOptional()
	attrs["effective_automatic_cluster_update_workspace"] = attrs["effective_automatic_cluster_update_workspace"].SetOptional()
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].SetComputed()
	attrs["effective_default_data_security_mode"] = attrs["effective_default_data_security_mode"].SetOptional()
	attrs["effective_integer_val"] = attrs["effective_integer_val"].SetComputed()
	attrs["effective_personal_compute"] = attrs["effective_personal_compute"].SetOptional()
	attrs["effective_restrict_workspace_admins"] = attrs["effective_restrict_workspace_admins"].SetOptional()
	attrs["effective_string_val"] = attrs["effective_string_val"].SetComputed()
	attrs["integer_val"] = attrs["integer_val"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["personal_compute"] = attrs["personal_compute"].SetOptional()
	attrs["restrict_workspace_admins"] = attrs["restrict_workspace_admins"].SetOptional()
	attrs["string_val"] = attrs["string_val"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Setting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Setting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_access_policy":              reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy{}),
		"aibi_dashboard_embedding_approved_domains":           reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains{}),
		"automatic_cluster_update_workspace":                  reflect.TypeOf(ClusterAutoRestartMessage{}),
		"boolean_val":                                         reflect.TypeOf(BooleanMessage{}),
		"default_data_security_mode":                          reflect.TypeOf(DefaultDataSecurityModeMessage{}),
		"effective_aibi_dashboard_embedding_access_policy":    reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy{}),
		"effective_aibi_dashboard_embedding_approved_domains": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains{}),
		"effective_automatic_cluster_update_workspace":        reflect.TypeOf(ClusterAutoRestartMessage{}),
		"effective_boolean_val":                               reflect.TypeOf(BooleanMessage{}),
		"effective_default_data_security_mode":                reflect.TypeOf(DefaultDataSecurityModeMessage{}),
		"effective_integer_val":                               reflect.TypeOf(IntegerMessage{}),
		"effective_personal_compute":                          reflect.TypeOf(PersonalComputeMessage{}),
		"effective_restrict_workspace_admins":                 reflect.TypeOf(RestrictWorkspaceAdminsMessage{}),
		"effective_string_val":                                reflect.TypeOf(StringMessage{}),
		"integer_val":                                         reflect.TypeOf(IntegerMessage{}),
		"personal_compute":                                    reflect.TypeOf(PersonalComputeMessage{}),
		"restrict_workspace_admins":                           reflect.TypeOf(RestrictWorkspaceAdminsMessage{}),
		"string_val":                                          reflect.TypeOf(StringMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Setting
// only implements ToObjectValue() and Type().
func (o Setting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Setting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aibi_dashboard_embedding_access_policy":              AibiDashboardEmbeddingAccessPolicy{}.Type(ctx),
			"aibi_dashboard_embedding_approved_domains":           AibiDashboardEmbeddingApprovedDomains{}.Type(ctx),
			"automatic_cluster_update_workspace":                  ClusterAutoRestartMessage{}.Type(ctx),
			"boolean_val":                                         BooleanMessage{}.Type(ctx),
			"default_data_security_mode":                          DefaultDataSecurityModeMessage{}.Type(ctx),
			"effective_aibi_dashboard_embedding_access_policy":    AibiDashboardEmbeddingAccessPolicy{}.Type(ctx),
			"effective_aibi_dashboard_embedding_approved_domains": AibiDashboardEmbeddingApprovedDomains{}.Type(ctx),
			"effective_automatic_cluster_update_workspace":        ClusterAutoRestartMessage{}.Type(ctx),
			"effective_boolean_val":                               BooleanMessage{}.Type(ctx),
			"effective_default_data_security_mode":                DefaultDataSecurityModeMessage{}.Type(ctx),
			"effective_integer_val":                               IntegerMessage{}.Type(ctx),
			"effective_personal_compute":                          PersonalComputeMessage{}.Type(ctx),
			"effective_restrict_workspace_admins":                 RestrictWorkspaceAdminsMessage{}.Type(ctx),
			"effective_string_val":                                StringMessage{}.Type(ctx),
			"integer_val":                                         IntegerMessage{}.Type(ctx),
			"name":                                                types.StringType,
			"personal_compute":                                    PersonalComputeMessage{}.Type(ctx),
			"restrict_workspace_admins":                           RestrictWorkspaceAdminsMessage{}.Type(ctx),
			"string_val":                                          StringMessage{}.Type(ctx),
		},
	}
}

// GetAibiDashboardEmbeddingAccessPolicy returns the value of the AibiDashboardEmbeddingAccessPolicy field in Setting as
// a AibiDashboardEmbeddingAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (AibiDashboardEmbeddingAccessPolicy, bool) {
	var e AibiDashboardEmbeddingAccessPolicy
	if o.AibiDashboardEmbeddingAccessPolicy.IsNull() || o.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v AibiDashboardEmbeddingAccessPolicy
	d := o.AibiDashboardEmbeddingAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAibiDashboardEmbeddingAccessPolicy sets the value of the AibiDashboardEmbeddingAccessPolicy field in Setting.
func (o *Setting) SetAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v AibiDashboardEmbeddingAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.AibiDashboardEmbeddingAccessPolicy = vs
}

// GetAibiDashboardEmbeddingApprovedDomains returns the value of the AibiDashboardEmbeddingApprovedDomains field in Setting as
// a AibiDashboardEmbeddingApprovedDomains value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (AibiDashboardEmbeddingApprovedDomains, bool) {
	var e AibiDashboardEmbeddingApprovedDomains
	if o.AibiDashboardEmbeddingApprovedDomains.IsNull() || o.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v AibiDashboardEmbeddingApprovedDomains
	d := o.AibiDashboardEmbeddingApprovedDomains.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAibiDashboardEmbeddingApprovedDomains sets the value of the AibiDashboardEmbeddingApprovedDomains field in Setting.
func (o *Setting) SetAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v AibiDashboardEmbeddingApprovedDomains) {
	vs := v.ToObjectValue(ctx)
	o.AibiDashboardEmbeddingApprovedDomains = vs
}

// GetAutomaticClusterUpdateWorkspace returns the value of the AutomaticClusterUpdateWorkspace field in Setting as
// a ClusterAutoRestartMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetAutomaticClusterUpdateWorkspace(ctx context.Context) (ClusterAutoRestartMessage, bool) {
	var e ClusterAutoRestartMessage
	if o.AutomaticClusterUpdateWorkspace.IsNull() || o.AutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessage
	d := o.AutomaticClusterUpdateWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutomaticClusterUpdateWorkspace sets the value of the AutomaticClusterUpdateWorkspace field in Setting.
func (o *Setting) SetAutomaticClusterUpdateWorkspace(ctx context.Context, v ClusterAutoRestartMessage) {
	vs := v.ToObjectValue(ctx)
	o.AutomaticClusterUpdateWorkspace = vs
}

// GetBooleanVal returns the value of the BooleanVal field in Setting as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in Setting.
func (o *Setting) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
}

// GetDefaultDataSecurityMode returns the value of the DefaultDataSecurityMode field in Setting as
// a DefaultDataSecurityModeMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetDefaultDataSecurityMode(ctx context.Context) (DefaultDataSecurityModeMessage, bool) {
	var e DefaultDataSecurityModeMessage
	if o.DefaultDataSecurityMode.IsNull() || o.DefaultDataSecurityMode.IsUnknown() {
		return e, false
	}
	var v DefaultDataSecurityModeMessage
	d := o.DefaultDataSecurityMode.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultDataSecurityMode sets the value of the DefaultDataSecurityMode field in Setting.
func (o *Setting) SetDefaultDataSecurityMode(ctx context.Context, v DefaultDataSecurityModeMessage) {
	vs := v.ToObjectValue(ctx)
	o.DefaultDataSecurityMode = vs
}

// GetEffectiveAibiDashboardEmbeddingAccessPolicy returns the value of the EffectiveAibiDashboardEmbeddingAccessPolicy field in Setting as
// a AibiDashboardEmbeddingAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (AibiDashboardEmbeddingAccessPolicy, bool) {
	var e AibiDashboardEmbeddingAccessPolicy
	if o.EffectiveAibiDashboardEmbeddingAccessPolicy.IsNull() || o.EffectiveAibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v AibiDashboardEmbeddingAccessPolicy
	d := o.EffectiveAibiDashboardEmbeddingAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveAibiDashboardEmbeddingAccessPolicy sets the value of the EffectiveAibiDashboardEmbeddingAccessPolicy field in Setting.
func (o *Setting) SetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v AibiDashboardEmbeddingAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveAibiDashboardEmbeddingAccessPolicy = vs
}

// GetEffectiveAibiDashboardEmbeddingApprovedDomains returns the value of the EffectiveAibiDashboardEmbeddingApprovedDomains field in Setting as
// a AibiDashboardEmbeddingApprovedDomains value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (AibiDashboardEmbeddingApprovedDomains, bool) {
	var e AibiDashboardEmbeddingApprovedDomains
	if o.EffectiveAibiDashboardEmbeddingApprovedDomains.IsNull() || o.EffectiveAibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v AibiDashboardEmbeddingApprovedDomains
	d := o.EffectiveAibiDashboardEmbeddingApprovedDomains.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveAibiDashboardEmbeddingApprovedDomains sets the value of the EffectiveAibiDashboardEmbeddingApprovedDomains field in Setting.
func (o *Setting) SetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v AibiDashboardEmbeddingApprovedDomains) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveAibiDashboardEmbeddingApprovedDomains = vs
}

// GetEffectiveAutomaticClusterUpdateWorkspace returns the value of the EffectiveAutomaticClusterUpdateWorkspace field in Setting as
// a ClusterAutoRestartMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveAutomaticClusterUpdateWorkspace(ctx context.Context) (ClusterAutoRestartMessage, bool) {
	var e ClusterAutoRestartMessage
	if o.EffectiveAutomaticClusterUpdateWorkspace.IsNull() || o.EffectiveAutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessage
	d := o.EffectiveAutomaticClusterUpdateWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveAutomaticClusterUpdateWorkspace sets the value of the EffectiveAutomaticClusterUpdateWorkspace field in Setting.
func (o *Setting) SetEffectiveAutomaticClusterUpdateWorkspace(ctx context.Context, v ClusterAutoRestartMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveAutomaticClusterUpdateWorkspace = vs
}

// GetEffectiveBooleanVal returns the value of the EffectiveBooleanVal field in Setting as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.EffectiveBooleanVal.IsNull() || o.EffectiveBooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := o.EffectiveBooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveBooleanVal sets the value of the EffectiveBooleanVal field in Setting.
func (o *Setting) SetEffectiveBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveBooleanVal = vs
}

// GetEffectiveDefaultDataSecurityMode returns the value of the EffectiveDefaultDataSecurityMode field in Setting as
// a DefaultDataSecurityModeMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveDefaultDataSecurityMode(ctx context.Context) (DefaultDataSecurityModeMessage, bool) {
	var e DefaultDataSecurityModeMessage
	if o.EffectiveDefaultDataSecurityMode.IsNull() || o.EffectiveDefaultDataSecurityMode.IsUnknown() {
		return e, false
	}
	var v DefaultDataSecurityModeMessage
	d := o.EffectiveDefaultDataSecurityMode.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveDefaultDataSecurityMode sets the value of the EffectiveDefaultDataSecurityMode field in Setting.
func (o *Setting) SetEffectiveDefaultDataSecurityMode(ctx context.Context, v DefaultDataSecurityModeMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveDefaultDataSecurityMode = vs
}

// GetEffectiveIntegerVal returns the value of the EffectiveIntegerVal field in Setting as
// a IntegerMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveIntegerVal(ctx context.Context) (IntegerMessage, bool) {
	var e IntegerMessage
	if o.EffectiveIntegerVal.IsNull() || o.EffectiveIntegerVal.IsUnknown() {
		return e, false
	}
	var v IntegerMessage
	d := o.EffectiveIntegerVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveIntegerVal sets the value of the EffectiveIntegerVal field in Setting.
func (o *Setting) SetEffectiveIntegerVal(ctx context.Context, v IntegerMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveIntegerVal = vs
}

// GetEffectivePersonalCompute returns the value of the EffectivePersonalCompute field in Setting as
// a PersonalComputeMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectivePersonalCompute(ctx context.Context) (PersonalComputeMessage, bool) {
	var e PersonalComputeMessage
	if o.EffectivePersonalCompute.IsNull() || o.EffectivePersonalCompute.IsUnknown() {
		return e, false
	}
	var v PersonalComputeMessage
	d := o.EffectivePersonalCompute.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectivePersonalCompute sets the value of the EffectivePersonalCompute field in Setting.
func (o *Setting) SetEffectivePersonalCompute(ctx context.Context, v PersonalComputeMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectivePersonalCompute = vs
}

// GetEffectiveRestrictWorkspaceAdmins returns the value of the EffectiveRestrictWorkspaceAdmins field in Setting as
// a RestrictWorkspaceAdminsMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveRestrictWorkspaceAdmins(ctx context.Context) (RestrictWorkspaceAdminsMessage, bool) {
	var e RestrictWorkspaceAdminsMessage
	if o.EffectiveRestrictWorkspaceAdmins.IsNull() || o.EffectiveRestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v RestrictWorkspaceAdminsMessage
	d := o.EffectiveRestrictWorkspaceAdmins.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveRestrictWorkspaceAdmins sets the value of the EffectiveRestrictWorkspaceAdmins field in Setting.
func (o *Setting) SetEffectiveRestrictWorkspaceAdmins(ctx context.Context, v RestrictWorkspaceAdminsMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveRestrictWorkspaceAdmins = vs
}

// GetEffectiveStringVal returns the value of the EffectiveStringVal field in Setting as
// a StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveStringVal(ctx context.Context) (StringMessage, bool) {
	var e StringMessage
	if o.EffectiveStringVal.IsNull() || o.EffectiveStringVal.IsUnknown() {
		return e, false
	}
	var v StringMessage
	d := o.EffectiveStringVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveStringVal sets the value of the EffectiveStringVal field in Setting.
func (o *Setting) SetEffectiveStringVal(ctx context.Context, v StringMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveStringVal = vs
}

// GetIntegerVal returns the value of the IntegerVal field in Setting as
// a IntegerMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetIntegerVal(ctx context.Context) (IntegerMessage, bool) {
	var e IntegerMessage
	if o.IntegerVal.IsNull() || o.IntegerVal.IsUnknown() {
		return e, false
	}
	var v IntegerMessage
	d := o.IntegerVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIntegerVal sets the value of the IntegerVal field in Setting.
func (o *Setting) SetIntegerVal(ctx context.Context, v IntegerMessage) {
	vs := v.ToObjectValue(ctx)
	o.IntegerVal = vs
}

// GetPersonalCompute returns the value of the PersonalCompute field in Setting as
// a PersonalComputeMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetPersonalCompute(ctx context.Context) (PersonalComputeMessage, bool) {
	var e PersonalComputeMessage
	if o.PersonalCompute.IsNull() || o.PersonalCompute.IsUnknown() {
		return e, false
	}
	var v PersonalComputeMessage
	d := o.PersonalCompute.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPersonalCompute sets the value of the PersonalCompute field in Setting.
func (o *Setting) SetPersonalCompute(ctx context.Context, v PersonalComputeMessage) {
	vs := v.ToObjectValue(ctx)
	o.PersonalCompute = vs
}

// GetRestrictWorkspaceAdmins returns the value of the RestrictWorkspaceAdmins field in Setting as
// a RestrictWorkspaceAdminsMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetRestrictWorkspaceAdmins(ctx context.Context) (RestrictWorkspaceAdminsMessage, bool) {
	var e RestrictWorkspaceAdminsMessage
	if o.RestrictWorkspaceAdmins.IsNull() || o.RestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v RestrictWorkspaceAdminsMessage
	d := o.RestrictWorkspaceAdmins.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestrictWorkspaceAdmins sets the value of the RestrictWorkspaceAdmins field in Setting.
func (o *Setting) SetRestrictWorkspaceAdmins(ctx context.Context, v RestrictWorkspaceAdminsMessage) {
	vs := v.ToObjectValue(ctx)
	o.RestrictWorkspaceAdmins = vs
}

// GetStringVal returns the value of the StringVal field in Setting as
// a StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetStringVal(ctx context.Context) (StringMessage, bool) {
	var e StringMessage
	if o.StringVal.IsNull() || o.StringVal.IsUnknown() {
		return e, false
	}
	var v StringMessage
	d := o.StringVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStringVal sets the value of the StringVal field in Setting.
func (o *Setting) SetStringVal(ctx context.Context, v StringMessage) {
	vs := v.ToObjectValue(ctx)
	o.StringVal = vs
}

type SettingsMetadata struct {
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

func (toState *SettingsMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SettingsMetadata) {
}

func (toState *SettingsMetadata) SyncFieldsDuringRead(ctx context.Context, fromState SettingsMetadata) {
}

func (c SettingsMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SettingsMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SettingsMetadata
// only implements ToObjectValue() and Type().
func (o SettingsMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SettingsMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"docs_link":   types.StringType,
			"name":        types.StringType,
			"type":        types.StringType,
		},
	}
}

type StringMessage struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value"`
}

func (toState *StringMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan StringMessage) {
}

func (toState *StringMessage) SyncFieldsDuringRead(ctx context.Context, fromState StringMessage) {
}

func (c StringMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a StringMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StringMessage
// only implements ToObjectValue() and Type().
func (o StringMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StringMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}
