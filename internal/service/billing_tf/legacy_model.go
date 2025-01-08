// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package billing_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ActionConfiguration_SdkV2 struct {
	// Databricks action configuration ID.
	ActionConfigurationId types.String `tfsdk:"action_configuration_id" tf:"optional"`
	// The type of the action.
	ActionType types.String `tfsdk:"action_type" tf:"optional"`
	// Target for the action. For example, an email address.
	Target types.String `tfsdk:"target" tf:"optional"`
}

func (newState *ActionConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ActionConfiguration_SdkV2) {
}

func (newState *ActionConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState ActionConfiguration_SdkV2) {
}

func (c ActionConfiguration_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ActionConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ActionConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ActionConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o ActionConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_configuration_id": o.ActionConfigurationId,
			"action_type":             o.ActionType,
			"target":                  o.Target,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ActionConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_configuration_id": types.StringType,
			"action_type":             types.StringType,
			"target":                  types.StringType,
		},
	}
}

type AlertConfiguration_SdkV2 struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	ActionConfigurations types.List `tfsdk:"action_configurations" tf:"optional"`
	// Databricks alert configuration ID.
	AlertConfigurationId types.String `tfsdk:"alert_configuration_id" tf:"optional"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	QuantityThreshold types.String `tfsdk:"quantity_threshold" tf:"optional"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	QuantityType types.String `tfsdk:"quantity_type" tf:"optional"`
	// The time window of usage data for the budget.
	TimePeriod types.String `tfsdk:"time_period" tf:"optional"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	TriggerType types.String `tfsdk:"trigger_type" tf:"optional"`
}

func (newState *AlertConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertConfiguration_SdkV2) {
}

func (newState *AlertConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertConfiguration_SdkV2) {
}

func (c AlertConfiguration_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	ActionConfiguration_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "action_configurations")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"action_configurations": reflect.TypeOf(ActionConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_configurations":  o.ActionConfigurations,
			"alert_configuration_id": o.AlertConfigurationId,
			"quantity_threshold":     o.QuantityThreshold,
			"quantity_type":          o.QuantityType,
			"time_period":            o.TimePeriod,
			"trigger_type":           o.TriggerType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_configurations": basetypes.ListType{
				ElemType: ActionConfiguration_SdkV2{}.Type(ctx),
			},
			"alert_configuration_id": types.StringType,
			"quantity_threshold":     types.StringType,
			"quantity_type":          types.StringType,
			"time_period":            types.StringType,
			"trigger_type":           types.StringType,
		},
	}
}

// GetActionConfigurations returns the value of the ActionConfigurations field in AlertConfiguration_SdkV2 as
// a slice of ActionConfiguration_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertConfiguration_SdkV2) GetActionConfigurations(ctx context.Context) ([]ActionConfiguration_SdkV2, bool) {
	if o.ActionConfigurations.IsNull() || o.ActionConfigurations.IsUnknown() {
		return nil, false
	}
	var v []ActionConfiguration_SdkV2
	d := o.ActionConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActionConfigurations sets the value of the ActionConfigurations field in AlertConfiguration_SdkV2.
func (o *AlertConfiguration_SdkV2) SetActionConfigurations(ctx context.Context, v []ActionConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["action_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ActionConfigurations = types.ListValueMust(t, vs)
}

type BudgetConfiguration_SdkV2 struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations types.List `tfsdk:"alert_configurations" tf:"optional"`
	// Databricks budget configuration ID.
	BudgetConfigurationId types.String `tfsdk:"budget_configuration_id" tf:"optional"`
	// Creation time of this budget configuration.
	CreateTime types.Int64 `tfsdk:"create_time" tf:"optional"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter types.List `tfsdk:"filter" tf:"optional,object"`
	// Update time of this budget configuration.
	UpdateTime types.Int64 `tfsdk:"update_time" tf:"optional"`
}

func (newState *BudgetConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfiguration_SdkV2) {
}

func (newState *BudgetConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState BudgetConfiguration_SdkV2) {
}

func (c BudgetConfiguration_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AlertConfiguration_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "alert_configurations")...)
	BudgetConfigurationFilter_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "filter")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_configurations": reflect.TypeOf(AlertConfiguration_SdkV2{}),
		"filter":               reflect.TypeOf(BudgetConfigurationFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o BudgetConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":              o.AccountId,
			"alert_configurations":    o.AlertConfigurations,
			"budget_configuration_id": o.BudgetConfigurationId,
			"create_time":             o.CreateTime,
			"display_name":            o.DisplayName,
			"filter":                  o.Filter,
			"update_time":             o.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"alert_configurations": basetypes.ListType{
				ElemType: AlertConfiguration_SdkV2{}.Type(ctx),
			},
			"budget_configuration_id": types.StringType,
			"create_time":             types.Int64Type,
			"display_name":            types.StringType,
			"filter": basetypes.ListType{
				ElemType: BudgetConfigurationFilter_SdkV2{}.Type(ctx),
			},
			"update_time": types.Int64Type,
		},
	}
}

// GetAlertConfigurations returns the value of the AlertConfigurations field in BudgetConfiguration_SdkV2 as
// a slice of AlertConfiguration_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfiguration_SdkV2) GetAlertConfigurations(ctx context.Context) ([]AlertConfiguration_SdkV2, bool) {
	if o.AlertConfigurations.IsNull() || o.AlertConfigurations.IsUnknown() {
		return nil, false
	}
	var v []AlertConfiguration_SdkV2
	d := o.AlertConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlertConfigurations sets the value of the AlertConfigurations field in BudgetConfiguration_SdkV2.
func (o *BudgetConfiguration_SdkV2) SetAlertConfigurations(ctx context.Context, v []AlertConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AlertConfigurations = types.ListValueMust(t, vs)
}

// GetFilter returns the value of the Filter field in BudgetConfiguration_SdkV2 as
// a BudgetConfigurationFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfiguration_SdkV2) GetFilter(ctx context.Context) (BudgetConfigurationFilter_SdkV2, bool) {
	var e BudgetConfigurationFilter_SdkV2
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilter_SdkV2
	d := o.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in BudgetConfiguration_SdkV2.
func (o *BudgetConfiguration_SdkV2) SetFilter(ctx context.Context, v BudgetConfigurationFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	o.Filter = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilter_SdkV2 struct {
	// A list of tag keys and values that will limit the budget to usage that
	// includes those specific custom tags. Tags are case-sensitive and should
	// be entered exactly as they appear in your usage data.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// If provided, usage must match with the provided Databricks workspace IDs.
	WorkspaceId types.List `tfsdk:"workspace_id" tf:"optional,object"`
}

func (newState *BudgetConfigurationFilter_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfigurationFilter_SdkV2) {
}

func (newState *BudgetConfigurationFilter_SdkV2) SyncEffectiveFieldsDuringRead(existingState BudgetConfigurationFilter_SdkV2) {
}

func (c BudgetConfigurationFilter_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	BudgetConfigurationFilterTagClause_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "tags")...)
	BudgetConfigurationFilterWorkspaceIdClause_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "workspace_id")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfigurationFilter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags":         reflect.TypeOf(BudgetConfigurationFilterTagClause_SdkV2{}),
		"workspace_id": reflect.TypeOf(BudgetConfigurationFilterWorkspaceIdClause_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilter_SdkV2
// only implements ToObjectValue() and Type().
func (o BudgetConfigurationFilter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tags":         o.Tags,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tags": basetypes.ListType{
				ElemType: BudgetConfigurationFilterTagClause_SdkV2{}.Type(ctx),
			},
			"workspace_id": basetypes.ListType{
				ElemType: BudgetConfigurationFilterWorkspaceIdClause_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in BudgetConfigurationFilter_SdkV2 as
// a slice of BudgetConfigurationFilterTagClause_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilter_SdkV2) GetTags(ctx context.Context) ([]BudgetConfigurationFilterTagClause_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []BudgetConfigurationFilterTagClause_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in BudgetConfigurationFilter_SdkV2.
func (o *BudgetConfigurationFilter_SdkV2) SetTags(ctx context.Context, v []BudgetConfigurationFilterTagClause_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// GetWorkspaceId returns the value of the WorkspaceId field in BudgetConfigurationFilter_SdkV2 as
// a BudgetConfigurationFilterWorkspaceIdClause_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilter_SdkV2) GetWorkspaceId(ctx context.Context) (BudgetConfigurationFilterWorkspaceIdClause_SdkV2, bool) {
	var e BudgetConfigurationFilterWorkspaceIdClause_SdkV2
	if o.WorkspaceId.IsNull() || o.WorkspaceId.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilterWorkspaceIdClause_SdkV2
	d := o.WorkspaceId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceId sets the value of the WorkspaceId field in BudgetConfigurationFilter_SdkV2.
func (o *BudgetConfigurationFilter_SdkV2) SetWorkspaceId(ctx context.Context, v BudgetConfigurationFilterWorkspaceIdClause_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_id"]
	o.WorkspaceId = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilterClause_SdkV2 struct {
	Operator types.String `tfsdk:"operator" tf:"optional"`

	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *BudgetConfigurationFilterClause_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfigurationFilterClause_SdkV2) {
}

func (newState *BudgetConfigurationFilterClause_SdkV2) SyncEffectiveFieldsDuringRead(existingState BudgetConfigurationFilterClause_SdkV2) {
}

func (c BudgetConfigurationFilterClause_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilterClause.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfigurationFilterClause_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilterClause_SdkV2
// only implements ToObjectValue() and Type().
func (o BudgetConfigurationFilterClause_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operator": o.Operator,
			"values":   o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterClause_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"operator": types.StringType,
			"values": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetValues returns the value of the Values field in BudgetConfigurationFilterClause_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilterClause_SdkV2) GetValues(ctx context.Context) ([]types.String, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in BudgetConfigurationFilterClause_SdkV2.
func (o *BudgetConfigurationFilterClause_SdkV2) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilterTagClause_SdkV2 struct {
	Key types.String `tfsdk:"key" tf:"optional"`

	Value types.List `tfsdk:"value" tf:"optional,object"`
}

func (newState *BudgetConfigurationFilterTagClause_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfigurationFilterTagClause_SdkV2) {
}

func (newState *BudgetConfigurationFilterTagClause_SdkV2) SyncEffectiveFieldsDuringRead(existingState BudgetConfigurationFilterTagClause_SdkV2) {
}

func (c BudgetConfigurationFilterTagClause_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	BudgetConfigurationFilterClause_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "value")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilterTagClause.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfigurationFilterTagClause_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(BudgetConfigurationFilterClause_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilterTagClause_SdkV2
// only implements ToObjectValue() and Type().
func (o BudgetConfigurationFilterTagClause_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterTagClause_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key": types.StringType,
			"value": basetypes.ListType{
				ElemType: BudgetConfigurationFilterClause_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetValue returns the value of the Value field in BudgetConfigurationFilterTagClause_SdkV2 as
// a BudgetConfigurationFilterClause_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilterTagClause_SdkV2) GetValue(ctx context.Context) (BudgetConfigurationFilterClause_SdkV2, bool) {
	var e BudgetConfigurationFilterClause_SdkV2
	if o.Value.IsNull() || o.Value.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilterClause_SdkV2
	d := o.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in BudgetConfigurationFilterTagClause_SdkV2.
func (o *BudgetConfigurationFilterTagClause_SdkV2) SetValue(ctx context.Context, v BudgetConfigurationFilterClause_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	o.Value = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilterWorkspaceIdClause_SdkV2 struct {
	Operator types.String `tfsdk:"operator" tf:"optional"`

	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *BudgetConfigurationFilterWorkspaceIdClause_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfigurationFilterWorkspaceIdClause_SdkV2) {
}

func (newState *BudgetConfigurationFilterWorkspaceIdClause_SdkV2) SyncEffectiveFieldsDuringRead(existingState BudgetConfigurationFilterWorkspaceIdClause_SdkV2) {
}

func (c BudgetConfigurationFilterWorkspaceIdClause_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilterWorkspaceIdClause.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfigurationFilterWorkspaceIdClause_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilterWorkspaceIdClause_SdkV2
// only implements ToObjectValue() and Type().
func (o BudgetConfigurationFilterWorkspaceIdClause_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operator": o.Operator,
			"values":   o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterWorkspaceIdClause_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"operator": types.StringType,
			"values": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetValues returns the value of the Values field in BudgetConfigurationFilterWorkspaceIdClause_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilterWorkspaceIdClause_SdkV2) GetValues(ctx context.Context) ([]types.Int64, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in BudgetConfigurationFilterWorkspaceIdClause_SdkV2.
func (o *BudgetConfigurationFilterWorkspaceIdClause_SdkV2) SetValues(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type CreateBillingUsageDashboardRequest_SdkV2 struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType types.String `tfsdk:"dashboard_type" tf:"optional"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

func (newState *CreateBillingUsageDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBillingUsageDashboardRequest_SdkV2) {
}

func (newState *CreateBillingUsageDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateBillingUsageDashboardRequest_SdkV2) {
}

func (c CreateBillingUsageDashboardRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBillingUsageDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBillingUsageDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBillingUsageDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateBillingUsageDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_type": o.DashboardType,
			"workspace_id":   o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

type CreateBillingUsageDashboardResponse_SdkV2 struct {
	// The unique id of the usage dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
}

func (newState *CreateBillingUsageDashboardResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBillingUsageDashboardResponse_SdkV2) {
}

func (newState *CreateBillingUsageDashboardResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateBillingUsageDashboardResponse_SdkV2) {
}

func (c CreateBillingUsageDashboardResponse_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBillingUsageDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBillingUsageDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBillingUsageDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateBillingUsageDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type CreateBudgetConfigurationBudget_SdkV2 struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations types.List `tfsdk:"alert_configurations" tf:"optional"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter types.List `tfsdk:"filter" tf:"optional,object"`
}

func (newState *CreateBudgetConfigurationBudget_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationBudget_SdkV2) {
}

func (newState *CreateBudgetConfigurationBudget_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationBudget_SdkV2) {
}

func (c CreateBudgetConfigurationBudget_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "alert_configurations")...)
	BudgetConfigurationFilter_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "filter")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationBudget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationBudget_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_configurations": reflect.TypeOf(CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2{}),
		"filter":               reflect.TypeOf(BudgetConfigurationFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationBudget_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationBudget_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":           o.AccountId,
			"alert_configurations": o.AlertConfigurations,
			"display_name":         o.DisplayName,
			"filter":               o.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudget_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"alert_configurations": basetypes.ListType{
				ElemType: CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2{}.Type(ctx),
			},
			"display_name": types.StringType,
			"filter": basetypes.ListType{
				ElemType: BudgetConfigurationFilter_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAlertConfigurations returns the value of the AlertConfigurations field in CreateBudgetConfigurationBudget_SdkV2 as
// a slice of CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationBudget_SdkV2) GetAlertConfigurations(ctx context.Context) ([]CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2, bool) {
	if o.AlertConfigurations.IsNull() || o.AlertConfigurations.IsUnknown() {
		return nil, false
	}
	var v []CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2
	d := o.AlertConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlertConfigurations sets the value of the AlertConfigurations field in CreateBudgetConfigurationBudget_SdkV2.
func (o *CreateBudgetConfigurationBudget_SdkV2) SetAlertConfigurations(ctx context.Context, v []CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AlertConfigurations = types.ListValueMust(t, vs)
}

// GetFilter returns the value of the Filter field in CreateBudgetConfigurationBudget_SdkV2 as
// a BudgetConfigurationFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationBudget_SdkV2) GetFilter(ctx context.Context) (BudgetConfigurationFilter_SdkV2, bool) {
	var e BudgetConfigurationFilter_SdkV2
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilter_SdkV2
	d := o.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in CreateBudgetConfigurationBudget_SdkV2.
func (o *CreateBudgetConfigurationBudget_SdkV2) SetFilter(ctx context.Context, v BudgetConfigurationFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	o.Filter = types.ListValueMust(t, vs)
}

type CreateBudgetConfigurationBudgetActionConfigurations_SdkV2 struct {
	// The type of the action.
	ActionType types.String `tfsdk:"action_type" tf:"optional"`
	// Target for the action. For example, an email address.
	Target types.String `tfsdk:"target" tf:"optional"`
}

func (newState *CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) {
}

func (newState *CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) {
}

func (c CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationBudgetActionConfigurations.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationBudgetActionConfigurations_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_type": o.ActionType,
			"target":      o.Target,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_type": types.StringType,
			"target":      types.StringType,
		},
	}
}

type CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2 struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	ActionConfigurations types.List `tfsdk:"action_configurations" tf:"optional"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	QuantityThreshold types.String `tfsdk:"quantity_threshold" tf:"optional"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	QuantityType types.String `tfsdk:"quantity_type" tf:"optional"`
	// The time window of usage data for the budget.
	TimePeriod types.String `tfsdk:"time_period" tf:"optional"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	TriggerType types.String `tfsdk:"trigger_type" tf:"optional"`
}

func (newState *CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) {
}

func (newState *CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) {
}

func (c CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	CreateBudgetConfigurationBudgetActionConfigurations_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "action_configurations")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationBudgetAlertConfigurations.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"action_configurations": reflect.TypeOf(CreateBudgetConfigurationBudgetActionConfigurations_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_configurations": o.ActionConfigurations,
			"quantity_threshold":    o.QuantityThreshold,
			"quantity_type":         o.QuantityType,
			"time_period":           o.TimePeriod,
			"trigger_type":          o.TriggerType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_configurations": basetypes.ListType{
				ElemType: CreateBudgetConfigurationBudgetActionConfigurations_SdkV2{}.Type(ctx),
			},
			"quantity_threshold": types.StringType,
			"quantity_type":      types.StringType,
			"time_period":        types.StringType,
			"trigger_type":       types.StringType,
		},
	}
}

// GetActionConfigurations returns the value of the ActionConfigurations field in CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2 as
// a slice of CreateBudgetConfigurationBudgetActionConfigurations_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) GetActionConfigurations(ctx context.Context) ([]CreateBudgetConfigurationBudgetActionConfigurations_SdkV2, bool) {
	if o.ActionConfigurations.IsNull() || o.ActionConfigurations.IsUnknown() {
		return nil, false
	}
	var v []CreateBudgetConfigurationBudgetActionConfigurations_SdkV2
	d := o.ActionConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActionConfigurations sets the value of the ActionConfigurations field in CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2.
func (o *CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) SetActionConfigurations(ctx context.Context, v []CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["action_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ActionConfigurations = types.ListValueMust(t, vs)
}

type CreateBudgetConfigurationRequest_SdkV2 struct {
	// Properties of the new budget configuration.
	Budget types.List `tfsdk:"budget" tf:"object"`
}

func (newState *CreateBudgetConfigurationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationRequest_SdkV2) {
}

func (newState *CreateBudgetConfigurationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationRequest_SdkV2) {
}

func (c CreateBudgetConfigurationRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "budget")...)
	CreateBudgetConfigurationBudget_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "budget")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(CreateBudgetConfigurationBudget_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": o.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: CreateBudgetConfigurationBudget_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetBudget returns the value of the Budget field in CreateBudgetConfigurationRequest_SdkV2 as
// a CreateBudgetConfigurationBudget_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationRequest_SdkV2) GetBudget(ctx context.Context) (CreateBudgetConfigurationBudget_SdkV2, bool) {
	var e CreateBudgetConfigurationBudget_SdkV2
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v []CreateBudgetConfigurationBudget_SdkV2
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in CreateBudgetConfigurationRequest_SdkV2.
func (o *CreateBudgetConfigurationRequest_SdkV2) SetBudget(ctx context.Context, v CreateBudgetConfigurationBudget_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

type CreateBudgetConfigurationResponse_SdkV2 struct {
	// The created budget configuration.
	Budget types.List `tfsdk:"budget" tf:"optional,object"`
}

func (newState *CreateBudgetConfigurationResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationResponse_SdkV2) {
}

func (newState *CreateBudgetConfigurationResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationResponse_SdkV2) {
}

func (c CreateBudgetConfigurationResponse_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	BudgetConfiguration_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "budget")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(BudgetConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": o.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: BudgetConfiguration_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetBudget returns the value of the Budget field in CreateBudgetConfigurationResponse_SdkV2 as
// a BudgetConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationResponse_SdkV2) GetBudget(ctx context.Context) (BudgetConfiguration_SdkV2, bool) {
	var e BudgetConfiguration_SdkV2
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v []BudgetConfiguration_SdkV2
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in CreateBudgetConfigurationResponse_SdkV2.
func (o *CreateBudgetConfigurationResponse_SdkV2) SetBudget(ctx context.Context, v BudgetConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

type CreateLogDeliveryConfigurationParams_SdkV2 struct {
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName types.String `tfsdk:"config_name" tf:"optional"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId types.String `tfsdk:"credentials_id" tf:""`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix types.String `tfsdk:"delivery_path_prefix" tf:"optional"`
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	DeliveryStartTime types.String `tfsdk:"delivery_start_time" tf:"optional"`
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the
	// CSV schema, see the [View billable usage].
	//
	// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema,
	// see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType types.String `tfsdk:"log_type" tf:""`
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be
	// `JSON`. Only the JSON (JavaScript Object Notation) format is supported.
	// For the schema, see the [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat types.String `tfsdk:"output_format" tf:""`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:""`
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	WorkspaceIdsFilter types.List `tfsdk:"workspace_ids_filter" tf:"optional"`
}

func (newState *CreateLogDeliveryConfigurationParams_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateLogDeliveryConfigurationParams_SdkV2) {
}

func (newState *CreateLogDeliveryConfigurationParams_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateLogDeliveryConfigurationParams_SdkV2) {
}

func (c CreateLogDeliveryConfigurationParams_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "credentials_id")...)
	cs.SetRequired(append(path, "log_type")...)
	cs.SetRequired(append(path, "output_format")...)
	cs.SetRequired(append(path, "storage_configuration_id")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateLogDeliveryConfigurationParams.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateLogDeliveryConfigurationParams_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_ids_filter": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLogDeliveryConfigurationParams_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateLogDeliveryConfigurationParams_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config_name":              o.ConfigName,
			"credentials_id":           o.CredentialsId,
			"delivery_path_prefix":     o.DeliveryPathPrefix,
			"delivery_start_time":      o.DeliveryStartTime,
			"log_type":                 o.LogType,
			"output_format":            o.OutputFormat,
			"status":                   o.Status,
			"storage_configuration_id": o.StorageConfigurationId,
			"workspace_ids_filter":     o.WorkspaceIdsFilter,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateLogDeliveryConfigurationParams_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config_name":              types.StringType,
			"credentials_id":           types.StringType,
			"delivery_path_prefix":     types.StringType,
			"delivery_start_time":      types.StringType,
			"log_type":                 types.StringType,
			"output_format":            types.StringType,
			"status":                   types.StringType,
			"storage_configuration_id": types.StringType,
			"workspace_ids_filter": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetWorkspaceIdsFilter returns the value of the WorkspaceIdsFilter field in CreateLogDeliveryConfigurationParams_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateLogDeliveryConfigurationParams_SdkV2) GetWorkspaceIdsFilter(ctx context.Context) ([]types.Int64, bool) {
	if o.WorkspaceIdsFilter.IsNull() || o.WorkspaceIdsFilter.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.WorkspaceIdsFilter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceIdsFilter sets the value of the WorkspaceIdsFilter field in CreateLogDeliveryConfigurationParams_SdkV2.
func (o *CreateLogDeliveryConfigurationParams_SdkV2) SetWorkspaceIdsFilter(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_ids_filter"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WorkspaceIdsFilter = types.ListValueMust(t, vs)
}

// Delete budget
type DeleteBudgetConfigurationRequest_SdkV2 struct {
	// The Databricks budget configuration ID.
	BudgetId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteBudgetConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteBudgetConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBudgetConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteBudgetConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_id": o.BudgetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_id": types.StringType,
		},
	}
}

type DeleteBudgetConfigurationResponse_SdkV2 struct {
}

func (newState *DeleteBudgetConfigurationResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteBudgetConfigurationResponse_SdkV2) {
}

func (newState *DeleteBudgetConfigurationResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteBudgetConfigurationResponse_SdkV2) {
}

func (c DeleteBudgetConfigurationResponse_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteBudgetConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBudgetConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteBudgetConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Return billable usage logs
type DownloadRequest_SdkV2 struct {
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	EndMonth types.String `tfsdk:"-"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	PersonalData types.Bool `tfsdk:"-"`
	// Format: `YYYY-MM`. First month to return billable usage logs for. This
	// field is required.
	StartMonth types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DownloadRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DownloadRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_month":     o.EndMonth,
			"personal_data": o.PersonalData,
			"start_month":   o.StartMonth,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DownloadRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_month":     types.StringType,
			"personal_data": types.BoolType,
			"start_month":   types.StringType,
		},
	}
}

type DownloadResponse_SdkV2 struct {
	Contents types.Object `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DownloadResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DownloadResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DownloadResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": o.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DownloadResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

// Get usage dashboard
type GetBillingUsageDashboardRequest_SdkV2 struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType types.String `tfsdk:"-"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBillingUsageDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBillingUsageDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBillingUsageDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetBillingUsageDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_type": o.DashboardType,
			"workspace_id":   o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

type GetBillingUsageDashboardResponse_SdkV2 struct {
	// The unique id of the usage dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The URL of the usage dashboard.
	DashboardUrl types.String `tfsdk:"dashboard_url" tf:"optional"`
}

func (newState *GetBillingUsageDashboardResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetBillingUsageDashboardResponse_SdkV2) {
}

func (newState *GetBillingUsageDashboardResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetBillingUsageDashboardResponse_SdkV2) {
}

func (c GetBillingUsageDashboardResponse_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBillingUsageDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBillingUsageDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBillingUsageDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetBillingUsageDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":  o.DashboardId,
			"dashboard_url": o.DashboardUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":  types.StringType,
			"dashboard_url": types.StringType,
		},
	}
}

// Get budget
type GetBudgetConfigurationRequest_SdkV2 struct {
	// The budget configuration ID
	BudgetId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBudgetConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBudgetConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBudgetConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetBudgetConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_id": o.BudgetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_id": types.StringType,
		},
	}
}

type GetBudgetConfigurationResponse_SdkV2 struct {
	Budget types.List `tfsdk:"budget" tf:"optional,object"`
}

func (newState *GetBudgetConfigurationResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetBudgetConfigurationResponse_SdkV2) {
}

func (newState *GetBudgetConfigurationResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetBudgetConfigurationResponse_SdkV2) {
}

func (c GetBudgetConfigurationResponse_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	BudgetConfiguration_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "budget")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBudgetConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(BudgetConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBudgetConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetBudgetConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": o.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: BudgetConfiguration_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetBudget returns the value of the Budget field in GetBudgetConfigurationResponse_SdkV2 as
// a BudgetConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetBudgetConfigurationResponse_SdkV2) GetBudget(ctx context.Context) (BudgetConfiguration_SdkV2, bool) {
	var e BudgetConfiguration_SdkV2
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v []BudgetConfiguration_SdkV2
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in GetBudgetConfigurationResponse_SdkV2.
func (o *GetBudgetConfigurationResponse_SdkV2) SetBudget(ctx context.Context, v BudgetConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

// Get log delivery configuration
type GetLogDeliveryRequest_SdkV2 struct {
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLogDeliveryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLogDeliveryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLogDeliveryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetLogDeliveryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration_id": o.LogDeliveryConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLogDeliveryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration_id": types.StringType,
		},
	}
}

// Get all budgets
type ListBudgetConfigurationsRequest_SdkV2 struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBudgetConfigurationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListBudgetConfigurationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetConfigurationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListBudgetConfigurationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

type ListBudgetConfigurationsResponse_SdkV2 struct {
	Budgets types.List `tfsdk:"budgets" tf:"optional"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListBudgetConfigurationsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListBudgetConfigurationsResponse_SdkV2) {
}

func (newState *ListBudgetConfigurationsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListBudgetConfigurationsResponse_SdkV2) {
}

func (c ListBudgetConfigurationsResponse_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	BudgetConfiguration_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "budgets")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBudgetConfigurationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListBudgetConfigurationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budgets": reflect.TypeOf(BudgetConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetConfigurationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListBudgetConfigurationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budgets":         o.Budgets,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budgets": basetypes.ListType{
				ElemType: BudgetConfiguration_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetBudgets returns the value of the Budgets field in ListBudgetConfigurationsResponse_SdkV2 as
// a slice of BudgetConfiguration_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListBudgetConfigurationsResponse_SdkV2) GetBudgets(ctx context.Context) ([]BudgetConfiguration_SdkV2, bool) {
	if o.Budgets.IsNull() || o.Budgets.IsUnknown() {
		return nil, false
	}
	var v []BudgetConfiguration_SdkV2
	d := o.Budgets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBudgets sets the value of the Budgets field in ListBudgetConfigurationsResponse_SdkV2.
func (o *ListBudgetConfigurationsResponse_SdkV2) SetBudgets(ctx context.Context, v []BudgetConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budgets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Budgets = types.ListValueMust(t, vs)
}

// Get all log delivery configurations
type ListLogDeliveryRequest_SdkV2 struct {
	// Filter by credential configuration ID.
	CredentialsId types.String `tfsdk:"-"`
	// Filter by status `ENABLED` or `DISABLED`.
	Status types.String `tfsdk:"-"`
	// Filter by storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListLogDeliveryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListLogDeliveryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListLogDeliveryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListLogDeliveryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials_id":           o.CredentialsId,
			"status":                   o.Status,
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListLogDeliveryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id":           types.StringType,
			"status":                   types.StringType,
			"storage_configuration_id": types.StringType,
		},
	}
}

type LogDeliveryConfiguration_SdkV2 struct {
	// The Databricks account ID that hosts the log delivery configuration.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// Databricks log delivery configuration ID.
	ConfigId types.String `tfsdk:"config_id" tf:"optional"`
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName types.String `tfsdk:"config_name" tf:"optional"`
	// Time in epoch milliseconds when the log delivery configuration was
	// created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"optional"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix types.String `tfsdk:"delivery_path_prefix" tf:"optional"`
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	DeliveryStartTime types.String `tfsdk:"delivery_start_time" tf:"optional"`
	// Databricks log delivery status.
	LogDeliveryStatus types.List `tfsdk:"log_delivery_status" tf:"optional,object"`
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the
	// CSV schema, see the [View billable usage].
	//
	// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema,
	// see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType types.String `tfsdk:"log_type" tf:"optional"`
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be
	// `JSON`. Only the JSON (JavaScript Object Notation) format is supported.
	// For the schema, see the [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat types.String `tfsdk:"output_format" tf:"optional"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// Time in epoch milliseconds when the log delivery configuration was
	// updated.
	UpdateTime types.Int64 `tfsdk:"update_time" tf:"optional"`
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	WorkspaceIdsFilter types.List `tfsdk:"workspace_ids_filter" tf:"optional"`
}

func (newState *LogDeliveryConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogDeliveryConfiguration_SdkV2) {
}

func (newState *LogDeliveryConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState LogDeliveryConfiguration_SdkV2) {
}

func (c LogDeliveryConfiguration_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	LogDeliveryStatus_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "log_delivery_status")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogDeliveryConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogDeliveryConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_status":  reflect.TypeOf(LogDeliveryStatus_SdkV2{}),
		"workspace_ids_filter": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogDeliveryConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o LogDeliveryConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":               o.AccountId,
			"config_id":                o.ConfigId,
			"config_name":              o.ConfigName,
			"creation_time":            o.CreationTime,
			"credentials_id":           o.CredentialsId,
			"delivery_path_prefix":     o.DeliveryPathPrefix,
			"delivery_start_time":      o.DeliveryStartTime,
			"log_delivery_status":      o.LogDeliveryStatus,
			"log_type":                 o.LogType,
			"output_format":            o.OutputFormat,
			"status":                   o.Status,
			"storage_configuration_id": o.StorageConfigurationId,
			"update_time":              o.UpdateTime,
			"workspace_ids_filter":     o.WorkspaceIdsFilter,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogDeliveryConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":           types.StringType,
			"config_id":            types.StringType,
			"config_name":          types.StringType,
			"creation_time":        types.Int64Type,
			"credentials_id":       types.StringType,
			"delivery_path_prefix": types.StringType,
			"delivery_start_time":  types.StringType,
			"log_delivery_status": basetypes.ListType{
				ElemType: LogDeliveryStatus_SdkV2{}.Type(ctx),
			},
			"log_type":                 types.StringType,
			"output_format":            types.StringType,
			"status":                   types.StringType,
			"storage_configuration_id": types.StringType,
			"update_time":              types.Int64Type,
			"workspace_ids_filter": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetLogDeliveryStatus returns the value of the LogDeliveryStatus field in LogDeliveryConfiguration_SdkV2 as
// a LogDeliveryStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LogDeliveryConfiguration_SdkV2) GetLogDeliveryStatus(ctx context.Context) (LogDeliveryStatus_SdkV2, bool) {
	var e LogDeliveryStatus_SdkV2
	if o.LogDeliveryStatus.IsNull() || o.LogDeliveryStatus.IsUnknown() {
		return e, false
	}
	var v []LogDeliveryStatus_SdkV2
	d := o.LogDeliveryStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryStatus sets the value of the LogDeliveryStatus field in LogDeliveryConfiguration_SdkV2.
func (o *LogDeliveryConfiguration_SdkV2) SetLogDeliveryStatus(ctx context.Context, v LogDeliveryStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_status"]
	o.LogDeliveryStatus = types.ListValueMust(t, vs)
}

// GetWorkspaceIdsFilter returns the value of the WorkspaceIdsFilter field in LogDeliveryConfiguration_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *LogDeliveryConfiguration_SdkV2) GetWorkspaceIdsFilter(ctx context.Context) ([]types.Int64, bool) {
	if o.WorkspaceIdsFilter.IsNull() || o.WorkspaceIdsFilter.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.WorkspaceIdsFilter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceIdsFilter sets the value of the WorkspaceIdsFilter field in LogDeliveryConfiguration_SdkV2.
func (o *LogDeliveryConfiguration_SdkV2) SetWorkspaceIdsFilter(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_ids_filter"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WorkspaceIdsFilter = types.ListValueMust(t, vs)
}

// Databricks log delivery status.
type LogDeliveryStatus_SdkV2 struct {
	// The UTC time for the latest log delivery attempt.
	LastAttemptTime types.String `tfsdk:"last_attempt_time" tf:"optional"`
	// The UTC time for the latest successful log delivery.
	LastSuccessfulAttemptTime types.String `tfsdk:"last_successful_attempt_time" tf:"optional"`
	// Informative message about the latest log delivery attempt. If the log
	// delivery fails with USER_FAILURE, error details will be provided for
	// fixing misconfigurations in cloud permissions.
	Message types.String `tfsdk:"message" tf:"optional"`
	// The status string for log delivery. Possible values are: * `CREATED`:
	// There were no log delivery attempts since the config was created. *
	// `SUCCEEDED`: The latest attempt of log delivery has succeeded completely.
	// * `USER_FAILURE`: The latest attempt of log delivery failed because of
	// misconfiguration of customer provided permissions on role or storage. *
	// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
	// Databricks internal error. Contact support if it doesn't go away soon. *
	// `NOT_FOUND`: The log delivery status as the configuration has been
	// disabled since the release of this feature or there are no workspaces in
	// the account.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *LogDeliveryStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogDeliveryStatus_SdkV2) {
}

func (newState *LogDeliveryStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState LogDeliveryStatus_SdkV2) {
}

func (c LogDeliveryStatus_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogDeliveryStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogDeliveryStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogDeliveryStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o LogDeliveryStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_attempt_time":            o.LastAttemptTime,
			"last_successful_attempt_time": o.LastSuccessfulAttemptTime,
			"message":                      o.Message,
			"status":                       o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogDeliveryStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_attempt_time":            types.StringType,
			"last_successful_attempt_time": types.StringType,
			"message":                      types.StringType,
			"status":                       types.StringType,
		},
	}
}

type PatchStatusResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchStatusResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchStatusResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PatchStatusResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PatchStatusResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateBudgetConfigurationBudget_SdkV2 struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations types.List `tfsdk:"alert_configurations" tf:"optional"`
	// Databricks budget configuration ID.
	BudgetConfigurationId types.String `tfsdk:"budget_configuration_id" tf:"optional"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter types.List `tfsdk:"filter" tf:"optional,object"`
}

func (newState *UpdateBudgetConfigurationBudget_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateBudgetConfigurationBudget_SdkV2) {
}

func (newState *UpdateBudgetConfigurationBudget_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateBudgetConfigurationBudget_SdkV2) {
}

func (c UpdateBudgetConfigurationBudget_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AlertConfiguration_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "alert_configurations")...)
	BudgetConfigurationFilter_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "filter")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetConfigurationBudget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateBudgetConfigurationBudget_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_configurations": reflect.TypeOf(AlertConfiguration_SdkV2{}),
		"filter":               reflect.TypeOf(BudgetConfigurationFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetConfigurationBudget_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateBudgetConfigurationBudget_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":              o.AccountId,
			"alert_configurations":    o.AlertConfigurations,
			"budget_configuration_id": o.BudgetConfigurationId,
			"display_name":            o.DisplayName,
			"filter":                  o.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationBudget_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"alert_configurations": basetypes.ListType{
				ElemType: AlertConfiguration_SdkV2{}.Type(ctx),
			},
			"budget_configuration_id": types.StringType,
			"display_name":            types.StringType,
			"filter": basetypes.ListType{
				ElemType: BudgetConfigurationFilter_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAlertConfigurations returns the value of the AlertConfigurations field in UpdateBudgetConfigurationBudget_SdkV2 as
// a slice of AlertConfiguration_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetConfigurationBudget_SdkV2) GetAlertConfigurations(ctx context.Context) ([]AlertConfiguration_SdkV2, bool) {
	if o.AlertConfigurations.IsNull() || o.AlertConfigurations.IsUnknown() {
		return nil, false
	}
	var v []AlertConfiguration_SdkV2
	d := o.AlertConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlertConfigurations sets the value of the AlertConfigurations field in UpdateBudgetConfigurationBudget_SdkV2.
func (o *UpdateBudgetConfigurationBudget_SdkV2) SetAlertConfigurations(ctx context.Context, v []AlertConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AlertConfigurations = types.ListValueMust(t, vs)
}

// GetFilter returns the value of the Filter field in UpdateBudgetConfigurationBudget_SdkV2 as
// a BudgetConfigurationFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetConfigurationBudget_SdkV2) GetFilter(ctx context.Context) (BudgetConfigurationFilter_SdkV2, bool) {
	var e BudgetConfigurationFilter_SdkV2
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilter_SdkV2
	d := o.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in UpdateBudgetConfigurationBudget_SdkV2.
func (o *UpdateBudgetConfigurationBudget_SdkV2) SetFilter(ctx context.Context, v BudgetConfigurationFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	o.Filter = types.ListValueMust(t, vs)
}

type UpdateBudgetConfigurationRequest_SdkV2 struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID.
	Budget types.List `tfsdk:"budget" tf:"object"`
	// The Databricks budget configuration ID.
	BudgetId types.String `tfsdk:"-"`
}

func (newState *UpdateBudgetConfigurationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateBudgetConfigurationRequest_SdkV2) {
}

func (newState *UpdateBudgetConfigurationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateBudgetConfigurationRequest_SdkV2) {
}

func (c UpdateBudgetConfigurationRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "budget")...)
	UpdateBudgetConfigurationBudget_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "budget")...)
	cs.SetRequired(append(path, "budget_id")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateBudgetConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(UpdateBudgetConfigurationBudget_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateBudgetConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget":    o.Budget,
			"budget_id": o.BudgetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: UpdateBudgetConfigurationBudget_SdkV2{}.Type(ctx),
			},
			"budget_id": types.StringType,
		},
	}
}

// GetBudget returns the value of the Budget field in UpdateBudgetConfigurationRequest_SdkV2 as
// a UpdateBudgetConfigurationBudget_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetConfigurationRequest_SdkV2) GetBudget(ctx context.Context) (UpdateBudgetConfigurationBudget_SdkV2, bool) {
	var e UpdateBudgetConfigurationBudget_SdkV2
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v []UpdateBudgetConfigurationBudget_SdkV2
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in UpdateBudgetConfigurationRequest_SdkV2.
func (o *UpdateBudgetConfigurationRequest_SdkV2) SetBudget(ctx context.Context, v UpdateBudgetConfigurationBudget_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

type UpdateBudgetConfigurationResponse_SdkV2 struct {
	// The updated budget.
	Budget types.List `tfsdk:"budget" tf:"optional,object"`
}

func (newState *UpdateBudgetConfigurationResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateBudgetConfigurationResponse_SdkV2) {
}

func (newState *UpdateBudgetConfigurationResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateBudgetConfigurationResponse_SdkV2) {
}

func (c UpdateBudgetConfigurationResponse_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	BudgetConfiguration_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "budget")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateBudgetConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(BudgetConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateBudgetConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": o.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: BudgetConfiguration_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetBudget returns the value of the Budget field in UpdateBudgetConfigurationResponse_SdkV2 as
// a BudgetConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetConfigurationResponse_SdkV2) GetBudget(ctx context.Context) (BudgetConfiguration_SdkV2, bool) {
	var e BudgetConfiguration_SdkV2
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v []BudgetConfiguration_SdkV2
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in UpdateBudgetConfigurationResponse_SdkV2.
func (o *UpdateBudgetConfigurationResponse_SdkV2) SetBudget(ctx context.Context, v BudgetConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

type UpdateLogDeliveryConfigurationStatusRequest_SdkV2 struct {
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId types.String `tfsdk:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status types.String `tfsdk:"status" tf:""`
}

func (newState *UpdateLogDeliveryConfigurationStatusRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateLogDeliveryConfigurationStatusRequest_SdkV2) {
}

func (newState *UpdateLogDeliveryConfigurationStatusRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateLogDeliveryConfigurationStatusRequest_SdkV2) {
}

func (c UpdateLogDeliveryConfigurationStatusRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "log_delivery_configuration_id")...)
	cs.SetRequired(append(path, "status")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLogDeliveryConfigurationStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateLogDeliveryConfigurationStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLogDeliveryConfigurationStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateLogDeliveryConfigurationStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration_id": o.LogDeliveryConfigurationId,
			"status":                        o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateLogDeliveryConfigurationStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration_id": types.StringType,
			"status":                        types.StringType,
		},
	}
}

type WrappedCreateLogDeliveryConfiguration_SdkV2 struct {
	LogDeliveryConfiguration types.List `tfsdk:"log_delivery_configuration" tf:"optional,object"`
}

func (newState *WrappedCreateLogDeliveryConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WrappedCreateLogDeliveryConfiguration_SdkV2) {
}

func (newState *WrappedCreateLogDeliveryConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState WrappedCreateLogDeliveryConfiguration_SdkV2) {
}

func (c WrappedCreateLogDeliveryConfiguration_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	CreateLogDeliveryConfigurationParams_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "log_delivery_configuration")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WrappedCreateLogDeliveryConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WrappedCreateLogDeliveryConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configuration": reflect.TypeOf(CreateLogDeliveryConfigurationParams_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WrappedCreateLogDeliveryConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o WrappedCreateLogDeliveryConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration": o.LogDeliveryConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WrappedCreateLogDeliveryConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration": basetypes.ListType{
				ElemType: CreateLogDeliveryConfigurationParams_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetLogDeliveryConfiguration returns the value of the LogDeliveryConfiguration field in WrappedCreateLogDeliveryConfiguration_SdkV2 as
// a CreateLogDeliveryConfigurationParams_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *WrappedCreateLogDeliveryConfiguration_SdkV2) GetLogDeliveryConfiguration(ctx context.Context) (CreateLogDeliveryConfigurationParams_SdkV2, bool) {
	var e CreateLogDeliveryConfigurationParams_SdkV2
	if o.LogDeliveryConfiguration.IsNull() || o.LogDeliveryConfiguration.IsUnknown() {
		return e, false
	}
	var v []CreateLogDeliveryConfigurationParams_SdkV2
	d := o.LogDeliveryConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in WrappedCreateLogDeliveryConfiguration_SdkV2.
func (o *WrappedCreateLogDeliveryConfiguration_SdkV2) SetLogDeliveryConfiguration(ctx context.Context, v CreateLogDeliveryConfigurationParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configuration"]
	o.LogDeliveryConfiguration = types.ListValueMust(t, vs)
}

type WrappedLogDeliveryConfiguration_SdkV2 struct {
	LogDeliveryConfiguration types.List `tfsdk:"log_delivery_configuration" tf:"optional,object"`
}

func (newState *WrappedLogDeliveryConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WrappedLogDeliveryConfiguration_SdkV2) {
}

func (newState *WrappedLogDeliveryConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState WrappedLogDeliveryConfiguration_SdkV2) {
}

func (c WrappedLogDeliveryConfiguration_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	LogDeliveryConfiguration_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "log_delivery_configuration")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WrappedLogDeliveryConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WrappedLogDeliveryConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configuration": reflect.TypeOf(LogDeliveryConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WrappedLogDeliveryConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o WrappedLogDeliveryConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration": o.LogDeliveryConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration": basetypes.ListType{
				ElemType: LogDeliveryConfiguration_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetLogDeliveryConfiguration returns the value of the LogDeliveryConfiguration field in WrappedLogDeliveryConfiguration_SdkV2 as
// a LogDeliveryConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *WrappedLogDeliveryConfiguration_SdkV2) GetLogDeliveryConfiguration(ctx context.Context) (LogDeliveryConfiguration_SdkV2, bool) {
	var e LogDeliveryConfiguration_SdkV2
	if o.LogDeliveryConfiguration.IsNull() || o.LogDeliveryConfiguration.IsUnknown() {
		return e, false
	}
	var v []LogDeliveryConfiguration_SdkV2
	d := o.LogDeliveryConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in WrappedLogDeliveryConfiguration_SdkV2.
func (o *WrappedLogDeliveryConfiguration_SdkV2) SetLogDeliveryConfiguration(ctx context.Context, v LogDeliveryConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configuration"]
	o.LogDeliveryConfiguration = types.ListValueMust(t, vs)
}

type WrappedLogDeliveryConfigurations_SdkV2 struct {
	LogDeliveryConfigurations types.List `tfsdk:"log_delivery_configurations" tf:"optional"`
}

func (newState *WrappedLogDeliveryConfigurations_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WrappedLogDeliveryConfigurations_SdkV2) {
}

func (newState *WrappedLogDeliveryConfigurations_SdkV2) SyncEffectiveFieldsDuringRead(existingState WrappedLogDeliveryConfigurations_SdkV2) {
}

func (c WrappedLogDeliveryConfigurations_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	LogDeliveryConfiguration_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "log_delivery_configurations")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WrappedLogDeliveryConfigurations.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WrappedLogDeliveryConfigurations_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configurations": reflect.TypeOf(LogDeliveryConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WrappedLogDeliveryConfigurations_SdkV2
// only implements ToObjectValue() and Type().
func (o WrappedLogDeliveryConfigurations_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configurations": o.LogDeliveryConfigurations,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfigurations_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configurations": basetypes.ListType{
				ElemType: LogDeliveryConfiguration_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetLogDeliveryConfigurations returns the value of the LogDeliveryConfigurations field in WrappedLogDeliveryConfigurations_SdkV2 as
// a slice of LogDeliveryConfiguration_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WrappedLogDeliveryConfigurations_SdkV2) GetLogDeliveryConfigurations(ctx context.Context) ([]LogDeliveryConfiguration_SdkV2, bool) {
	if o.LogDeliveryConfigurations.IsNull() || o.LogDeliveryConfigurations.IsUnknown() {
		return nil, false
	}
	var v []LogDeliveryConfiguration_SdkV2
	d := o.LogDeliveryConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLogDeliveryConfigurations sets the value of the LogDeliveryConfigurations field in WrappedLogDeliveryConfigurations_SdkV2.
func (o *WrappedLogDeliveryConfigurations_SdkV2) SetLogDeliveryConfigurations(ctx context.Context, v []LogDeliveryConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LogDeliveryConfigurations = types.ListValueMust(t, vs)
}
