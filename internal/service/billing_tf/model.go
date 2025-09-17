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
	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ActionConfiguration struct {
	// Databricks action configuration ID.
	ActionConfigurationId types.String `tfsdk:"action_configuration_id"`
	// The type of the action.
	ActionType types.String `tfsdk:"action_type"`
	// Target for the action. For example, an email address.
	Target types.String `tfsdk:"target"`
}

func (toState *ActionConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ActionConfiguration) {
}

func (toState *ActionConfiguration) SyncFieldsDuringRead(ctx context.Context, fromState ActionConfiguration) {
}

func (c ActionConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["action_configuration_id"] = attrs["action_configuration_id"].SetOptional()
	attrs["action_type"] = attrs["action_type"].SetOptional()
	attrs["target"] = attrs["target"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ActionConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ActionConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ActionConfiguration
// only implements ToObjectValue() and Type().
func (o ActionConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_configuration_id": o.ActionConfigurationId,
			"action_type":             o.ActionType,
			"target":                  o.Target,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ActionConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_configuration_id": types.StringType,
			"action_type":             types.StringType,
			"target":                  types.StringType,
		},
	}
}

type AlertConfiguration struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	ActionConfigurations types.List `tfsdk:"action_configurations"`
	// Databricks alert configuration ID.
	AlertConfigurationId types.String `tfsdk:"alert_configuration_id"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	QuantityThreshold types.String `tfsdk:"quantity_threshold"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	QuantityType types.String `tfsdk:"quantity_type"`
	// The time window of usage data for the budget.
	TimePeriod types.String `tfsdk:"time_period"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	TriggerType types.String `tfsdk:"trigger_type"`
}

func (toState *AlertConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AlertConfiguration) {
}

func (toState *AlertConfiguration) SyncFieldsDuringRead(ctx context.Context, fromState AlertConfiguration) {
}

func (c AlertConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["action_configurations"] = attrs["action_configurations"].SetOptional()
	attrs["alert_configuration_id"] = attrs["alert_configuration_id"].SetOptional()
	attrs["quantity_threshold"] = attrs["quantity_threshold"].SetOptional()
	attrs["quantity_type"] = attrs["quantity_type"].SetOptional()
	attrs["time_period"] = attrs["time_period"].SetOptional()
	attrs["trigger_type"] = attrs["trigger_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"action_configurations": reflect.TypeOf(ActionConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConfiguration
// only implements ToObjectValue() and Type().
func (o AlertConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AlertConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_configurations": basetypes.ListType{
				ElemType: ActionConfiguration{}.Type(ctx),
			},
			"alert_configuration_id": types.StringType,
			"quantity_threshold":     types.StringType,
			"quantity_type":          types.StringType,
			"time_period":            types.StringType,
			"trigger_type":           types.StringType,
		},
	}
}

// GetActionConfigurations returns the value of the ActionConfigurations field in AlertConfiguration as
// a slice of ActionConfiguration values.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertConfiguration) GetActionConfigurations(ctx context.Context) ([]ActionConfiguration, bool) {
	if o.ActionConfigurations.IsNull() || o.ActionConfigurations.IsUnknown() {
		return nil, false
	}
	var v []ActionConfiguration
	d := o.ActionConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActionConfigurations sets the value of the ActionConfigurations field in AlertConfiguration.
func (o *AlertConfiguration) SetActionConfigurations(ctx context.Context, v []ActionConfiguration) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["action_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ActionConfigurations = types.ListValueMust(t, vs)
}

type BudgetConfiguration struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations types.List `tfsdk:"alert_configurations"`
	// Databricks budget configuration ID.
	BudgetConfigurationId types.String `tfsdk:"budget_configuration_id"`
	// Creation time of this budget configuration.
	CreateTime types.Int64 `tfsdk:"create_time"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName types.String `tfsdk:"display_name"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter types.Object `tfsdk:"filter"`
	// Update time of this budget configuration.
	UpdateTime types.Int64 `tfsdk:"update_time"`
}

func (toState *BudgetConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan BudgetConfiguration) {
	if !fromPlan.Filter.IsNull() && !fromPlan.Filter.IsUnknown() {
		if toStateFilter, ok := toState.GetFilter(ctx); ok {
			if fromPlanFilter, ok := fromPlan.GetFilter(ctx); ok {
				toStateFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFilter)
				toState.SetFilter(ctx, toStateFilter)
			}
		}
	}
}

func (toState *BudgetConfiguration) SyncFieldsDuringRead(ctx context.Context, fromState BudgetConfiguration) {
	if !fromState.Filter.IsNull() && !fromState.Filter.IsUnknown() {
		if toStateFilter, ok := toState.GetFilter(ctx); ok {
			if fromStateFilter, ok := fromState.GetFilter(ctx); ok {
				toStateFilter.SyncFieldsDuringRead(ctx, fromStateFilter)
				toState.SetFilter(ctx, toStateFilter)
			}
		}
	}
}

func (c BudgetConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["alert_configurations"] = attrs["alert_configurations"].SetOptional()
	attrs["budget_configuration_id"] = attrs["budget_configuration_id"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_configurations": reflect.TypeOf(AlertConfiguration{}),
		"filter":               reflect.TypeOf(BudgetConfigurationFilter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfiguration
// only implements ToObjectValue() and Type().
func (o BudgetConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o BudgetConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"alert_configurations": basetypes.ListType{
				ElemType: AlertConfiguration{}.Type(ctx),
			},
			"budget_configuration_id": types.StringType,
			"create_time":             types.Int64Type,
			"display_name":            types.StringType,
			"filter":                  BudgetConfigurationFilter{}.Type(ctx),
			"update_time":             types.Int64Type,
		},
	}
}

// GetAlertConfigurations returns the value of the AlertConfigurations field in BudgetConfiguration as
// a slice of AlertConfiguration values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfiguration) GetAlertConfigurations(ctx context.Context) ([]AlertConfiguration, bool) {
	if o.AlertConfigurations.IsNull() || o.AlertConfigurations.IsUnknown() {
		return nil, false
	}
	var v []AlertConfiguration
	d := o.AlertConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlertConfigurations sets the value of the AlertConfigurations field in BudgetConfiguration.
func (o *BudgetConfiguration) SetAlertConfigurations(ctx context.Context, v []AlertConfiguration) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AlertConfigurations = types.ListValueMust(t, vs)
}

// GetFilter returns the value of the Filter field in BudgetConfiguration as
// a BudgetConfigurationFilter value.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfiguration) GetFilter(ctx context.Context) (BudgetConfigurationFilter, bool) {
	var e BudgetConfigurationFilter
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v BudgetConfigurationFilter
	d := o.Filter.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilter sets the value of the Filter field in BudgetConfiguration.
func (o *BudgetConfiguration) SetFilter(ctx context.Context, v BudgetConfigurationFilter) {
	vs := v.ToObjectValue(ctx)
	o.Filter = vs
}

type BudgetConfigurationFilter struct {
	// A list of tag keys and values that will limit the budget to usage that
	// includes those specific custom tags. Tags are case-sensitive and should
	// be entered exactly as they appear in your usage data.
	Tags types.List `tfsdk:"tags"`
	// If provided, usage must match with the provided Databricks workspace IDs.
	WorkspaceId types.Object `tfsdk:"workspace_id"`
}

func (toState *BudgetConfigurationFilter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan BudgetConfigurationFilter) {
	if !fromPlan.WorkspaceId.IsNull() && !fromPlan.WorkspaceId.IsUnknown() {
		if toStateWorkspaceId, ok := toState.GetWorkspaceId(ctx); ok {
			if fromPlanWorkspaceId, ok := fromPlan.GetWorkspaceId(ctx); ok {
				toStateWorkspaceId.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanWorkspaceId)
				toState.SetWorkspaceId(ctx, toStateWorkspaceId)
			}
		}
	}
}

func (toState *BudgetConfigurationFilter) SyncFieldsDuringRead(ctx context.Context, fromState BudgetConfigurationFilter) {
	if !fromState.WorkspaceId.IsNull() && !fromState.WorkspaceId.IsUnknown() {
		if toStateWorkspaceId, ok := toState.GetWorkspaceId(ctx); ok {
			if fromStateWorkspaceId, ok := fromState.GetWorkspaceId(ctx); ok {
				toStateWorkspaceId.SyncFieldsDuringRead(ctx, fromStateWorkspaceId)
				toState.SetWorkspaceId(ctx, toStateWorkspaceId)
			}
		}
	}
}

func (c BudgetConfigurationFilter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfigurationFilter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags":         reflect.TypeOf(BudgetConfigurationFilterTagClause{}),
		"workspace_id": reflect.TypeOf(BudgetConfigurationFilterWorkspaceIdClause{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilter
// only implements ToObjectValue() and Type().
func (o BudgetConfigurationFilter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tags":         o.Tags,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tags": basetypes.ListType{
				ElemType: BudgetConfigurationFilterTagClause{}.Type(ctx),
			},
			"workspace_id": BudgetConfigurationFilterWorkspaceIdClause{}.Type(ctx),
		},
	}
}

// GetTags returns the value of the Tags field in BudgetConfigurationFilter as
// a slice of BudgetConfigurationFilterTagClause values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilter) GetTags(ctx context.Context) ([]BudgetConfigurationFilterTagClause, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []BudgetConfigurationFilterTagClause
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in BudgetConfigurationFilter.
func (o *BudgetConfigurationFilter) SetTags(ctx context.Context, v []BudgetConfigurationFilterTagClause) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// GetWorkspaceId returns the value of the WorkspaceId field in BudgetConfigurationFilter as
// a BudgetConfigurationFilterWorkspaceIdClause value.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilter) GetWorkspaceId(ctx context.Context) (BudgetConfigurationFilterWorkspaceIdClause, bool) {
	var e BudgetConfigurationFilterWorkspaceIdClause
	if o.WorkspaceId.IsNull() || o.WorkspaceId.IsUnknown() {
		return e, false
	}
	var v BudgetConfigurationFilterWorkspaceIdClause
	d := o.WorkspaceId.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceId sets the value of the WorkspaceId field in BudgetConfigurationFilter.
func (o *BudgetConfigurationFilter) SetWorkspaceId(ctx context.Context, v BudgetConfigurationFilterWorkspaceIdClause) {
	vs := v.ToObjectValue(ctx)
	o.WorkspaceId = vs
}

type BudgetConfigurationFilterClause struct {
	Operator types.String `tfsdk:"operator"`

	Values types.List `tfsdk:"values"`
}

func (toState *BudgetConfigurationFilterClause) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan BudgetConfigurationFilterClause) {
}

func (toState *BudgetConfigurationFilterClause) SyncFieldsDuringRead(ctx context.Context, fromState BudgetConfigurationFilterClause) {
}

func (c BudgetConfigurationFilterClause) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["operator"] = attrs["operator"].SetOptional()
	attrs["values"] = attrs["values"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilterClause.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfigurationFilterClause) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilterClause
// only implements ToObjectValue() and Type().
func (o BudgetConfigurationFilterClause) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operator": o.Operator,
			"values":   o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterClause) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"operator": types.StringType,
			"values": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetValues returns the value of the Values field in BudgetConfigurationFilterClause as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilterClause) GetValues(ctx context.Context) ([]types.String, bool) {
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

// SetValues sets the value of the Values field in BudgetConfigurationFilterClause.
func (o *BudgetConfigurationFilterClause) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilterTagClause struct {
	Key types.String `tfsdk:"key"`

	Value types.Object `tfsdk:"value"`
}

func (toState *BudgetConfigurationFilterTagClause) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan BudgetConfigurationFilterTagClause) {
	if !fromPlan.Value.IsNull() && !fromPlan.Value.IsUnknown() {
		if toStateValue, ok := toState.GetValue(ctx); ok {
			if fromPlanValue, ok := fromPlan.GetValue(ctx); ok {
				toStateValue.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanValue)
				toState.SetValue(ctx, toStateValue)
			}
		}
	}
}

func (toState *BudgetConfigurationFilterTagClause) SyncFieldsDuringRead(ctx context.Context, fromState BudgetConfigurationFilterTagClause) {
	if !fromState.Value.IsNull() && !fromState.Value.IsUnknown() {
		if toStateValue, ok := toState.GetValue(ctx); ok {
			if fromStateValue, ok := fromState.GetValue(ctx); ok {
				toStateValue.SyncFieldsDuringRead(ctx, fromStateValue)
				toState.SetValue(ctx, toStateValue)
			}
		}
	}
}

func (c BudgetConfigurationFilterTagClause) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilterTagClause.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfigurationFilterTagClause) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(BudgetConfigurationFilterClause{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilterTagClause
// only implements ToObjectValue() and Type().
func (o BudgetConfigurationFilterTagClause) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterTagClause) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": BudgetConfigurationFilterClause{}.Type(ctx),
		},
	}
}

// GetValue returns the value of the Value field in BudgetConfigurationFilterTagClause as
// a BudgetConfigurationFilterClause value.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilterTagClause) GetValue(ctx context.Context) (BudgetConfigurationFilterClause, bool) {
	var e BudgetConfigurationFilterClause
	if o.Value.IsNull() || o.Value.IsUnknown() {
		return e, false
	}
	var v BudgetConfigurationFilterClause
	d := o.Value.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValue sets the value of the Value field in BudgetConfigurationFilterTagClause.
func (o *BudgetConfigurationFilterTagClause) SetValue(ctx context.Context, v BudgetConfigurationFilterClause) {
	vs := v.ToObjectValue(ctx)
	o.Value = vs
}

type BudgetConfigurationFilterWorkspaceIdClause struct {
	Operator types.String `tfsdk:"operator"`

	Values types.List `tfsdk:"values"`
}

func (toState *BudgetConfigurationFilterWorkspaceIdClause) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan BudgetConfigurationFilterWorkspaceIdClause) {
}

func (toState *BudgetConfigurationFilterWorkspaceIdClause) SyncFieldsDuringRead(ctx context.Context, fromState BudgetConfigurationFilterWorkspaceIdClause) {
}

func (c BudgetConfigurationFilterWorkspaceIdClause) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["operator"] = attrs["operator"].SetOptional()
	attrs["values"] = attrs["values"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilterWorkspaceIdClause.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetConfigurationFilterWorkspaceIdClause) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilterWorkspaceIdClause
// only implements ToObjectValue() and Type().
func (o BudgetConfigurationFilterWorkspaceIdClause) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operator": o.Operator,
			"values":   o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterWorkspaceIdClause) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"operator": types.StringType,
			"values": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetValues returns the value of the Values field in BudgetConfigurationFilterWorkspaceIdClause as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetConfigurationFilterWorkspaceIdClause) GetValues(ctx context.Context) ([]types.Int64, bool) {
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

// SetValues sets the value of the Values field in BudgetConfigurationFilterWorkspaceIdClause.
func (o *BudgetConfigurationFilterWorkspaceIdClause) SetValues(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

// Contains the BudgetPolicy details.
type BudgetPolicy struct {
	// List of workspaces that this budget policy will be exclusively bound to.
	// An empty binding implies that this budget policy is open to any workspace
	// in the account.
	BindingWorkspaceIds types.List `tfsdk:"binding_workspace_ids"`
	// A list of tags defined by the customer. At most 20 entries are allowed
	// per policy.
	CustomTags types.List `tfsdk:"custom_tags"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId types.String `tfsdk:"policy_id"`
	// The name of the policy. - Must be unique among active policies. - Can
	// contain only characters from the ISO 8859-1 (latin1) set. - Can't start
	// with reserved keywords such as `databricks:default-policy`.
	PolicyName types.String `tfsdk:"policy_name"`
}

func (toState *BudgetPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan BudgetPolicy) {
}

func (toState *BudgetPolicy) SyncFieldsDuringRead(ctx context.Context, fromState BudgetPolicy) {
}

func (c BudgetPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["binding_workspace_ids"] = attrs["binding_workspace_ids"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetComputed()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()
	attrs["policy_name"] = attrs["policy_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BudgetPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"binding_workspace_ids": reflect.TypeOf(types.Int64{}),
		"custom_tags":           reflect.TypeOf(compute_tf.CustomPolicyTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetPolicy
// only implements ToObjectValue() and Type().
func (o BudgetPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"binding_workspace_ids": o.BindingWorkspaceIds,
			"custom_tags":           o.CustomTags,
			"policy_id":             o.PolicyId,
			"policy_name":           o.PolicyName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BudgetPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"binding_workspace_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"custom_tags": basetypes.ListType{
				ElemType: compute_tf.CustomPolicyTag{}.Type(ctx),
			},
			"policy_id":   types.StringType,
			"policy_name": types.StringType,
		},
	}
}

// GetBindingWorkspaceIds returns the value of the BindingWorkspaceIds field in BudgetPolicy as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetPolicy) GetBindingWorkspaceIds(ctx context.Context) ([]types.Int64, bool) {
	if o.BindingWorkspaceIds.IsNull() || o.BindingWorkspaceIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.BindingWorkspaceIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBindingWorkspaceIds sets the value of the BindingWorkspaceIds field in BudgetPolicy.
func (o *BudgetPolicy) SetBindingWorkspaceIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["binding_workspace_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.BindingWorkspaceIds = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in BudgetPolicy as
// a slice of compute_tf.CustomPolicyTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *BudgetPolicy) GetCustomTags(ctx context.Context) ([]compute_tf.CustomPolicyTag, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.CustomPolicyTag
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in BudgetPolicy.
func (o *BudgetPolicy) SetCustomTags(ctx context.Context, v []compute_tf.CustomPolicyTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.ListValueMust(t, vs)
}

type CreateBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType types.String `tfsdk:"dashboard_type"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (toState *CreateBillingUsageDashboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateBillingUsageDashboardRequest) {
}

func (toState *CreateBillingUsageDashboardRequest) SyncFieldsDuringRead(ctx context.Context, fromState CreateBillingUsageDashboardRequest) {
}

func (c CreateBillingUsageDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_type"] = attrs["dashboard_type"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBillingUsageDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBillingUsageDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBillingUsageDashboardRequest
// only implements ToObjectValue() and Type().
func (o CreateBillingUsageDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_type": o.DashboardType,
			"workspace_id":   o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

type CreateBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`
}

func (toState *CreateBillingUsageDashboardResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateBillingUsageDashboardResponse) {
}

func (toState *CreateBillingUsageDashboardResponse) SyncFieldsDuringRead(ctx context.Context, fromState CreateBillingUsageDashboardResponse) {
}

func (c CreateBillingUsageDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBillingUsageDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBillingUsageDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBillingUsageDashboardResponse
// only implements ToObjectValue() and Type().
func (o CreateBillingUsageDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type CreateBudgetConfigurationBudget struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations types.List `tfsdk:"alert_configurations"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName types.String `tfsdk:"display_name"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter types.Object `tfsdk:"filter"`
}

func (toState *CreateBudgetConfigurationBudget) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateBudgetConfigurationBudget) {
	if !fromPlan.Filter.IsNull() && !fromPlan.Filter.IsUnknown() {
		if toStateFilter, ok := toState.GetFilter(ctx); ok {
			if fromPlanFilter, ok := fromPlan.GetFilter(ctx); ok {
				toStateFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFilter)
				toState.SetFilter(ctx, toStateFilter)
			}
		}
	}
}

func (toState *CreateBudgetConfigurationBudget) SyncFieldsDuringRead(ctx context.Context, fromState CreateBudgetConfigurationBudget) {
	if !fromState.Filter.IsNull() && !fromState.Filter.IsUnknown() {
		if toStateFilter, ok := toState.GetFilter(ctx); ok {
			if fromStateFilter, ok := fromState.GetFilter(ctx); ok {
				toStateFilter.SyncFieldsDuringRead(ctx, fromStateFilter)
				toState.SetFilter(ctx, toStateFilter)
			}
		}
	}
}

func (c CreateBudgetConfigurationBudget) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["alert_configurations"] = attrs["alert_configurations"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationBudget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationBudget) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_configurations": reflect.TypeOf(CreateBudgetConfigurationBudgetAlertConfigurations{}),
		"filter":               reflect.TypeOf(BudgetConfigurationFilter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationBudget
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationBudget) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateBudgetConfigurationBudget) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"alert_configurations": basetypes.ListType{
				ElemType: CreateBudgetConfigurationBudgetAlertConfigurations{}.Type(ctx),
			},
			"display_name": types.StringType,
			"filter":       BudgetConfigurationFilter{}.Type(ctx),
		},
	}
}

// GetAlertConfigurations returns the value of the AlertConfigurations field in CreateBudgetConfigurationBudget as
// a slice of CreateBudgetConfigurationBudgetAlertConfigurations values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationBudget) GetAlertConfigurations(ctx context.Context) ([]CreateBudgetConfigurationBudgetAlertConfigurations, bool) {
	if o.AlertConfigurations.IsNull() || o.AlertConfigurations.IsUnknown() {
		return nil, false
	}
	var v []CreateBudgetConfigurationBudgetAlertConfigurations
	d := o.AlertConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlertConfigurations sets the value of the AlertConfigurations field in CreateBudgetConfigurationBudget.
func (o *CreateBudgetConfigurationBudget) SetAlertConfigurations(ctx context.Context, v []CreateBudgetConfigurationBudgetAlertConfigurations) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AlertConfigurations = types.ListValueMust(t, vs)
}

// GetFilter returns the value of the Filter field in CreateBudgetConfigurationBudget as
// a BudgetConfigurationFilter value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationBudget) GetFilter(ctx context.Context) (BudgetConfigurationFilter, bool) {
	var e BudgetConfigurationFilter
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v BudgetConfigurationFilter
	d := o.Filter.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilter sets the value of the Filter field in CreateBudgetConfigurationBudget.
func (o *CreateBudgetConfigurationBudget) SetFilter(ctx context.Context, v BudgetConfigurationFilter) {
	vs := v.ToObjectValue(ctx)
	o.Filter = vs
}

type CreateBudgetConfigurationBudgetActionConfigurations struct {
	// The type of the action.
	ActionType types.String `tfsdk:"action_type"`
	// Target for the action. For example, an email address.
	Target types.String `tfsdk:"target"`
}

func (toState *CreateBudgetConfigurationBudgetActionConfigurations) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateBudgetConfigurationBudgetActionConfigurations) {
}

func (toState *CreateBudgetConfigurationBudgetActionConfigurations) SyncFieldsDuringRead(ctx context.Context, fromState CreateBudgetConfigurationBudgetActionConfigurations) {
}

func (c CreateBudgetConfigurationBudgetActionConfigurations) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["action_type"] = attrs["action_type"].SetOptional()
	attrs["target"] = attrs["target"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationBudgetActionConfigurations.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationBudgetActionConfigurations) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationBudgetActionConfigurations
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationBudgetActionConfigurations) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_type": o.ActionType,
			"target":      o.Target,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetActionConfigurations) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_type": types.StringType,
			"target":      types.StringType,
		},
	}
}

type CreateBudgetConfigurationBudgetAlertConfigurations struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	ActionConfigurations types.List `tfsdk:"action_configurations"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	QuantityThreshold types.String `tfsdk:"quantity_threshold"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	QuantityType types.String `tfsdk:"quantity_type"`
	// The time window of usage data for the budget.
	TimePeriod types.String `tfsdk:"time_period"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	TriggerType types.String `tfsdk:"trigger_type"`
}

func (toState *CreateBudgetConfigurationBudgetAlertConfigurations) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateBudgetConfigurationBudgetAlertConfigurations) {
}

func (toState *CreateBudgetConfigurationBudgetAlertConfigurations) SyncFieldsDuringRead(ctx context.Context, fromState CreateBudgetConfigurationBudgetAlertConfigurations) {
}

func (c CreateBudgetConfigurationBudgetAlertConfigurations) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["action_configurations"] = attrs["action_configurations"].SetOptional()
	attrs["quantity_threshold"] = attrs["quantity_threshold"].SetOptional()
	attrs["quantity_type"] = attrs["quantity_type"].SetOptional()
	attrs["time_period"] = attrs["time_period"].SetOptional()
	attrs["trigger_type"] = attrs["trigger_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationBudgetAlertConfigurations.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationBudgetAlertConfigurations) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"action_configurations": reflect.TypeOf(CreateBudgetConfigurationBudgetActionConfigurations{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationBudgetAlertConfigurations
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationBudgetAlertConfigurations) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateBudgetConfigurationBudgetAlertConfigurations) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_configurations": basetypes.ListType{
				ElemType: CreateBudgetConfigurationBudgetActionConfigurations{}.Type(ctx),
			},
			"quantity_threshold": types.StringType,
			"quantity_type":      types.StringType,
			"time_period":        types.StringType,
			"trigger_type":       types.StringType,
		},
	}
}

// GetActionConfigurations returns the value of the ActionConfigurations field in CreateBudgetConfigurationBudgetAlertConfigurations as
// a slice of CreateBudgetConfigurationBudgetActionConfigurations values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationBudgetAlertConfigurations) GetActionConfigurations(ctx context.Context) ([]CreateBudgetConfigurationBudgetActionConfigurations, bool) {
	if o.ActionConfigurations.IsNull() || o.ActionConfigurations.IsUnknown() {
		return nil, false
	}
	var v []CreateBudgetConfigurationBudgetActionConfigurations
	d := o.ActionConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActionConfigurations sets the value of the ActionConfigurations field in CreateBudgetConfigurationBudgetAlertConfigurations.
func (o *CreateBudgetConfigurationBudgetAlertConfigurations) SetActionConfigurations(ctx context.Context, v []CreateBudgetConfigurationBudgetActionConfigurations) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["action_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ActionConfigurations = types.ListValueMust(t, vs)
}

type CreateBudgetConfigurationRequest struct {
	// Properties of the new budget configuration.
	Budget types.Object `tfsdk:"budget"`
}

func (toState *CreateBudgetConfigurationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateBudgetConfigurationRequest) {
	if !fromPlan.Budget.IsNull() && !fromPlan.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromPlanBudget, ok := fromPlan.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (toState *CreateBudgetConfigurationRequest) SyncFieldsDuringRead(ctx context.Context, fromState CreateBudgetConfigurationRequest) {
	if !fromState.Budget.IsNull() && !fromState.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromStateBudget, ok := fromState.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringRead(ctx, fromStateBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (c CreateBudgetConfigurationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(CreateBudgetConfigurationBudget{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationRequest
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": o.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": CreateBudgetConfigurationBudget{}.Type(ctx),
		},
	}
}

// GetBudget returns the value of the Budget field in CreateBudgetConfigurationRequest as
// a CreateBudgetConfigurationBudget value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationRequest) GetBudget(ctx context.Context) (CreateBudgetConfigurationBudget, bool) {
	var e CreateBudgetConfigurationBudget
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v CreateBudgetConfigurationBudget
	d := o.Budget.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBudget sets the value of the Budget field in CreateBudgetConfigurationRequest.
func (o *CreateBudgetConfigurationRequest) SetBudget(ctx context.Context, v CreateBudgetConfigurationBudget) {
	vs := v.ToObjectValue(ctx)
	o.Budget = vs
}

type CreateBudgetConfigurationResponse struct {
	// The created budget configuration.
	Budget types.Object `tfsdk:"budget"`
}

func (toState *CreateBudgetConfigurationResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateBudgetConfigurationResponse) {
	if !fromPlan.Budget.IsNull() && !fromPlan.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromPlanBudget, ok := fromPlan.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (toState *CreateBudgetConfigurationResponse) SyncFieldsDuringRead(ctx context.Context, fromState CreateBudgetConfigurationResponse) {
	if !fromState.Budget.IsNull() && !fromState.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromStateBudget, ok := fromState.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringRead(ctx, fromStateBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (c CreateBudgetConfigurationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetConfigurationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(BudgetConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationResponse
// only implements ToObjectValue() and Type().
func (o CreateBudgetConfigurationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": o.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": BudgetConfiguration{}.Type(ctx),
		},
	}
}

// GetBudget returns the value of the Budget field in CreateBudgetConfigurationResponse as
// a BudgetConfiguration value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetConfigurationResponse) GetBudget(ctx context.Context) (BudgetConfiguration, bool) {
	var e BudgetConfiguration
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v BudgetConfiguration
	d := o.Budget.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBudget sets the value of the Budget field in CreateBudgetConfigurationResponse.
func (o *CreateBudgetConfigurationResponse) SetBudget(ctx context.Context, v BudgetConfiguration) {
	vs := v.ToObjectValue(ctx)
	o.Budget = vs
}

// A request to create a BudgetPolicy.
type CreateBudgetPolicyRequest struct {
	// The policy to create. `policy_id` needs to be empty as it will be
	// generated `policy_name` must be provided, custom_tags may need to be
	// provided depending on the cloud provider. All other fields are optional.
	Policy types.Object `tfsdk:"policy"`
	// A unique identifier for this request. Restricted to 36 ASCII characters.
	// A random UUID is recommended. This request is only idempotent if a
	// `request_id` is provided.
	RequestId types.String `tfsdk:"request_id"`
}

func (toState *CreateBudgetPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateBudgetPolicyRequest) {
	if !fromPlan.Policy.IsNull() && !fromPlan.Policy.IsUnknown() {
		if toStatePolicy, ok := toState.GetPolicy(ctx); ok {
			if fromPlanPolicy, ok := fromPlan.GetPolicy(ctx); ok {
				toStatePolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanPolicy)
				toState.SetPolicy(ctx, toStatePolicy)
			}
		}
	}
}

func (toState *CreateBudgetPolicyRequest) SyncFieldsDuringRead(ctx context.Context, fromState CreateBudgetPolicyRequest) {
	if !fromState.Policy.IsNull() && !fromState.Policy.IsUnknown() {
		if toStatePolicy, ok := toState.GetPolicy(ctx); ok {
			if fromStatePolicy, ok := fromState.GetPolicy(ctx); ok {
				toStatePolicy.SyncFieldsDuringRead(ctx, fromStatePolicy)
				toState.SetPolicy(ctx, toStatePolicy)
			}
		}
	}
}

func (c CreateBudgetPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetOptional()
	attrs["request_id"] = attrs["request_id"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateBudgetPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(BudgetPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetPolicyRequest
// only implements ToObjectValue() and Type().
func (o CreateBudgetPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":     o.Policy,
			"request_id": o.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateBudgetPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy":     BudgetPolicy{}.Type(ctx),
			"request_id": types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in CreateBudgetPolicyRequest as
// a BudgetPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateBudgetPolicyRequest) GetPolicy(ctx context.Context) (BudgetPolicy, bool) {
	var e BudgetPolicy
	if o.Policy.IsNull() || o.Policy.IsUnknown() {
		return e, false
	}
	var v BudgetPolicy
	d := o.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicy sets the value of the Policy field in CreateBudgetPolicyRequest.
func (o *CreateBudgetPolicyRequest) SetPolicy(ctx context.Context, v BudgetPolicy) {
	vs := v.ToObjectValue(ctx)
	o.Policy = vs
}

// * Log Delivery Configuration
type CreateLogDeliveryConfigurationParams struct {
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName types.String `tfsdk:"config_name"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix types.String `tfsdk:"delivery_path_prefix"`
	// This field applies only if log_type is BILLABLE_USAGE. This is the
	// optional start month and year for delivery, specified in YYYY-MM format.
	// Defaults to current year and month. BILLABLE_USAGE logs are not available
	// for usage before March 2019 (2019-03).
	DeliveryStartTime types.String `tfsdk:"delivery_start_time"`
	// Log delivery type. Supported values are: * `BILLABLE_USAGE` — Configure
	// [billable usage log delivery]. For the CSV schema, see the [View billable
	// usage]. * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON
	// schema, see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType types.String `tfsdk:"log_type"`
	// The file type of log delivery. * If `log_type` is `BILLABLE_USAGE`, this
	// value must be `CSV`. Only the CSV (comma-separated values) format is
	// supported. For the schema, see the [View billable usage] * If `log_type`
	// is `AUDIT_LOGS`, this value must be `JSON`. Only the JSON (JavaScript
	// Object Notation) format is supported. For the schema, see the
	// [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat types.String `tfsdk:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status types.String `tfsdk:"status"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
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
	WorkspaceIdsFilter types.List `tfsdk:"workspace_ids_filter"`
}

func (toState *CreateLogDeliveryConfigurationParams) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateLogDeliveryConfigurationParams) {
}

func (toState *CreateLogDeliveryConfigurationParams) SyncFieldsDuringRead(ctx context.Context, fromState CreateLogDeliveryConfigurationParams) {
}

func (c CreateLogDeliveryConfigurationParams) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config_name"] = attrs["config_name"].SetOptional()
	attrs["credentials_id"] = attrs["credentials_id"].SetRequired()
	attrs["delivery_path_prefix"] = attrs["delivery_path_prefix"].SetOptional()
	attrs["delivery_start_time"] = attrs["delivery_start_time"].SetOptional()
	attrs["log_type"] = attrs["log_type"].SetRequired()
	attrs["output_format"] = attrs["output_format"].SetRequired()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetRequired()
	attrs["workspace_ids_filter"] = attrs["workspace_ids_filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateLogDeliveryConfigurationParams.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateLogDeliveryConfigurationParams) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_ids_filter": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLogDeliveryConfigurationParams
// only implements ToObjectValue() and Type().
func (o CreateLogDeliveryConfigurationParams) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateLogDeliveryConfigurationParams) Type(ctx context.Context) attr.Type {
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

// GetWorkspaceIdsFilter returns the value of the WorkspaceIdsFilter field in CreateLogDeliveryConfigurationParams as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateLogDeliveryConfigurationParams) GetWorkspaceIdsFilter(ctx context.Context) ([]types.Int64, bool) {
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

// SetWorkspaceIdsFilter sets the value of the WorkspaceIdsFilter field in CreateLogDeliveryConfigurationParams.
func (o *CreateLogDeliveryConfigurationParams) SetWorkspaceIdsFilter(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_ids_filter"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WorkspaceIdsFilter = types.ListValueMust(t, vs)
}

type DeleteBudgetConfigurationRequest struct {
	// The Databricks budget configuration ID.
	BudgetId types.String `tfsdk:"-"`
}

func (toState *DeleteBudgetConfigurationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteBudgetConfigurationRequest) {
}

func (toState *DeleteBudgetConfigurationRequest) SyncFieldsDuringRead(ctx context.Context, fromState DeleteBudgetConfigurationRequest) {
}

func (c DeleteBudgetConfigurationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["budget_id"] = attrs["budget_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteBudgetConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteBudgetConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBudgetConfigurationRequest
// only implements ToObjectValue() and Type().
func (o DeleteBudgetConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_id": o.BudgetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_id": types.StringType,
		},
	}
}

type DeleteBudgetConfigurationResponse struct {
}

func (toState *DeleteBudgetConfigurationResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteBudgetConfigurationResponse) {
}

func (toState *DeleteBudgetConfigurationResponse) SyncFieldsDuringRead(ctx context.Context, fromState DeleteBudgetConfigurationResponse) {
}

func (c DeleteBudgetConfigurationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteBudgetConfigurationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBudgetConfigurationResponse
// only implements ToObjectValue() and Type().
func (o DeleteBudgetConfigurationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId types.String `tfsdk:"-"`
}

func (toState *DeleteBudgetPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteBudgetPolicyRequest) {
}

func (toState *DeleteBudgetPolicyRequest) SyncFieldsDuringRead(ctx context.Context, fromState DeleteBudgetPolicyRequest) {
}

func (c DeleteBudgetPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteBudgetPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteBudgetPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBudgetPolicyRequest
// only implements ToObjectValue() and Type().
func (o DeleteBudgetPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteBudgetPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type DownloadRequest struct {
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	EndMonth types.String `tfsdk:"-"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	PersonalData types.Bool `tfsdk:"-"`
	// Format specification for month in the format `YYYY-MM`. This is used to
	// specify billable usage `start_month` and `end_month` properties.
	// **Note**: Billable usage logs are unavailable before March 2019
	// (`2019-03`).
	StartMonth types.String `tfsdk:"-"`
}

func (toState *DownloadRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DownloadRequest) {
}

func (toState *DownloadRequest) SyncFieldsDuringRead(ctx context.Context, fromState DownloadRequest) {
}

func (c DownloadRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["start_month"] = attrs["start_month"].SetRequired()
	attrs["end_month"] = attrs["end_month"].SetRequired()
	attrs["personal_data"] = attrs["personal_data"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DownloadRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadRequest
// only implements ToObjectValue() and Type().
func (o DownloadRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_month":     o.EndMonth,
			"personal_data": o.PersonalData,
			"start_month":   o.StartMonth,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DownloadRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_month":     types.StringType,
			"personal_data": types.BoolType,
			"start_month":   types.StringType,
		},
	}
}

type DownloadResponse struct {
	Contents types.Object `tfsdk:"-"`
}

func (toState *DownloadResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DownloadResponse) {
}

func (toState *DownloadResponse) SyncFieldsDuringRead(ctx context.Context, fromState DownloadResponse) {
}

func (c DownloadResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["contents"] = attrs["contents"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DownloadResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DownloadResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadResponse
// only implements ToObjectValue() and Type().
func (o DownloadResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": o.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DownloadResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

// Structured representation of a filter to be applied to a list of policies.
// All specified filters will be applied in conjunction.
type Filter struct {
	// The policy creator user id to be filtered on. If unspecified, all
	// policies will be returned.
	CreatorUserId types.Int64 `tfsdk:"creator_user_id"`
	// The policy creator user name to be filtered on. If unspecified, all
	// policies will be returned.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// The partial name of policies to be filtered on. If unspecified, all
	// policies will be returned.
	PolicyName types.String `tfsdk:"policy_name"`
}

func (toState *Filter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Filter) {
}

func (toState *Filter) SyncFieldsDuringRead(ctx context.Context, fromState Filter) {
}

func (c Filter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creator_user_id"] = attrs["creator_user_id"].SetOptional()
	attrs["creator_user_name"] = attrs["creator_user_name"].SetOptional()
	attrs["policy_name"] = attrs["policy_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Filter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Filter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Filter
// only implements ToObjectValue() and Type().
func (o Filter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator_user_id":   o.CreatorUserId,
			"creator_user_name": o.CreatorUserName,
			"policy_name":       o.PolicyName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Filter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator_user_id":   types.Int64Type,
			"creator_user_name": types.StringType,
			"policy_name":       types.StringType,
		},
	}
}

type GetBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType types.String `tfsdk:"-"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (toState *GetBillingUsageDashboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetBillingUsageDashboardRequest) {
}

func (toState *GetBillingUsageDashboardRequest) SyncFieldsDuringRead(ctx context.Context, fromState GetBillingUsageDashboardRequest) {
}

func (c GetBillingUsageDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["dashboard_type"] = attrs["dashboard_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBillingUsageDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBillingUsageDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBillingUsageDashboardRequest
// only implements ToObjectValue() and Type().
func (o GetBillingUsageDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_type": o.DashboardType,
			"workspace_id":   o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

type GetBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The URL of the usage dashboard.
	DashboardUrl types.String `tfsdk:"dashboard_url"`
}

func (toState *GetBillingUsageDashboardResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetBillingUsageDashboardResponse) {
}

func (toState *GetBillingUsageDashboardResponse) SyncFieldsDuringRead(ctx context.Context, fromState GetBillingUsageDashboardResponse) {
}

func (c GetBillingUsageDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetOptional()
	attrs["dashboard_url"] = attrs["dashboard_url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBillingUsageDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBillingUsageDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBillingUsageDashboardResponse
// only implements ToObjectValue() and Type().
func (o GetBillingUsageDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":  o.DashboardId,
			"dashboard_url": o.DashboardUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":  types.StringType,
			"dashboard_url": types.StringType,
		},
	}
}

type GetBudgetConfigurationRequest struct {
	// The budget configuration ID
	BudgetId types.String `tfsdk:"-"`
}

func (toState *GetBudgetConfigurationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetBudgetConfigurationRequest) {
}

func (toState *GetBudgetConfigurationRequest) SyncFieldsDuringRead(ctx context.Context, fromState GetBudgetConfigurationRequest) {
}

func (c GetBudgetConfigurationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["budget_id"] = attrs["budget_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBudgetConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBudgetConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBudgetConfigurationRequest
// only implements ToObjectValue() and Type().
func (o GetBudgetConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_id": o.BudgetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_id": types.StringType,
		},
	}
}

type GetBudgetConfigurationResponse struct {
	Budget types.Object `tfsdk:"budget"`
}

func (toState *GetBudgetConfigurationResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetBudgetConfigurationResponse) {
	if !fromPlan.Budget.IsNull() && !fromPlan.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromPlanBudget, ok := fromPlan.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (toState *GetBudgetConfigurationResponse) SyncFieldsDuringRead(ctx context.Context, fromState GetBudgetConfigurationResponse) {
	if !fromState.Budget.IsNull() && !fromState.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromStateBudget, ok := fromState.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringRead(ctx, fromStateBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (c GetBudgetConfigurationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBudgetConfigurationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(BudgetConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBudgetConfigurationResponse
// only implements ToObjectValue() and Type().
func (o GetBudgetConfigurationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": o.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": BudgetConfiguration{}.Type(ctx),
		},
	}
}

// GetBudget returns the value of the Budget field in GetBudgetConfigurationResponse as
// a BudgetConfiguration value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetBudgetConfigurationResponse) GetBudget(ctx context.Context) (BudgetConfiguration, bool) {
	var e BudgetConfiguration
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v BudgetConfiguration
	d := o.Budget.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBudget sets the value of the Budget field in GetBudgetConfigurationResponse.
func (o *GetBudgetConfigurationResponse) SetBudget(ctx context.Context, v BudgetConfiguration) {
	vs := v.ToObjectValue(ctx)
	o.Budget = vs
}

type GetBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId types.String `tfsdk:"-"`
}

func (toState *GetBudgetPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetBudgetPolicyRequest) {
}

func (toState *GetBudgetPolicyRequest) SyncFieldsDuringRead(ctx context.Context, fromState GetBudgetPolicyRequest) {
}

func (c GetBudgetPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBudgetPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBudgetPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBudgetPolicyRequest
// only implements ToObjectValue() and Type().
func (o GetBudgetPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBudgetPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type GetLogDeliveryConfigurationResponse struct {
	// The fetched log delivery configuration
	LogDeliveryConfiguration types.Object `tfsdk:"log_delivery_configuration"`
}

func (toState *GetLogDeliveryConfigurationResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetLogDeliveryConfigurationResponse) {
	if !fromPlan.LogDeliveryConfiguration.IsNull() && !fromPlan.LogDeliveryConfiguration.IsUnknown() {
		if toStateLogDeliveryConfiguration, ok := toState.GetLogDeliveryConfiguration(ctx); ok {
			if fromPlanLogDeliveryConfiguration, ok := fromPlan.GetLogDeliveryConfiguration(ctx); ok {
				toStateLogDeliveryConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanLogDeliveryConfiguration)
				toState.SetLogDeliveryConfiguration(ctx, toStateLogDeliveryConfiguration)
			}
		}
	}
}

func (toState *GetLogDeliveryConfigurationResponse) SyncFieldsDuringRead(ctx context.Context, fromState GetLogDeliveryConfigurationResponse) {
	if !fromState.LogDeliveryConfiguration.IsNull() && !fromState.LogDeliveryConfiguration.IsUnknown() {
		if toStateLogDeliveryConfiguration, ok := toState.GetLogDeliveryConfiguration(ctx); ok {
			if fromStateLogDeliveryConfiguration, ok := fromState.GetLogDeliveryConfiguration(ctx); ok {
				toStateLogDeliveryConfiguration.SyncFieldsDuringRead(ctx, fromStateLogDeliveryConfiguration)
				toState.SetLogDeliveryConfiguration(ctx, toStateLogDeliveryConfiguration)
			}
		}
	}
}

func (c GetLogDeliveryConfigurationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["log_delivery_configuration"] = attrs["log_delivery_configuration"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLogDeliveryConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLogDeliveryConfigurationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configuration": reflect.TypeOf(LogDeliveryConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLogDeliveryConfigurationResponse
// only implements ToObjectValue() and Type().
func (o GetLogDeliveryConfigurationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration": o.LogDeliveryConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLogDeliveryConfigurationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration": LogDeliveryConfiguration{}.Type(ctx),
		},
	}
}

// GetLogDeliveryConfiguration returns the value of the LogDeliveryConfiguration field in GetLogDeliveryConfigurationResponse as
// a LogDeliveryConfiguration value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetLogDeliveryConfigurationResponse) GetLogDeliveryConfiguration(ctx context.Context) (LogDeliveryConfiguration, bool) {
	var e LogDeliveryConfiguration
	if o.LogDeliveryConfiguration.IsNull() || o.LogDeliveryConfiguration.IsUnknown() {
		return e, false
	}
	var v LogDeliveryConfiguration
	d := o.LogDeliveryConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in GetLogDeliveryConfigurationResponse.
func (o *GetLogDeliveryConfigurationResponse) SetLogDeliveryConfiguration(ctx context.Context, v LogDeliveryConfiguration) {
	vs := v.ToObjectValue(ctx)
	o.LogDeliveryConfiguration = vs
}

type GetLogDeliveryRequest struct {
	// The log delivery configuration id of customer
	LogDeliveryConfigurationId types.String `tfsdk:"-"`
}

func (toState *GetLogDeliveryRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetLogDeliveryRequest) {
}

func (toState *GetLogDeliveryRequest) SyncFieldsDuringRead(ctx context.Context, fromState GetLogDeliveryRequest) {
}

func (c GetLogDeliveryRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["log_delivery_configuration_id"] = attrs["log_delivery_configuration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLogDeliveryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLogDeliveryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLogDeliveryRequest
// only implements ToObjectValue() and Type().
func (o GetLogDeliveryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration_id": o.LogDeliveryConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLogDeliveryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration_id": types.StringType,
		},
	}
}

// The limit configuration of the policy. Limit configuration provide a budget
// policy level cost control by enforcing the limit.
type LimitConfig struct {
}

func (toState *LimitConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LimitConfig) {
}

func (toState *LimitConfig) SyncFieldsDuringRead(ctx context.Context, fromState LimitConfig) {
}

func (c LimitConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LimitConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LimitConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LimitConfig
// only implements ToObjectValue() and Type().
func (o LimitConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LimitConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListBudgetConfigurationsRequest struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (toState *ListBudgetConfigurationsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListBudgetConfigurationsRequest) {
}

func (toState *ListBudgetConfigurationsRequest) SyncFieldsDuringRead(ctx context.Context, fromState ListBudgetConfigurationsRequest) {
}

func (c ListBudgetConfigurationsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBudgetConfigurationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListBudgetConfigurationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetConfigurationsRequest
// only implements ToObjectValue() and Type().
func (o ListBudgetConfigurationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

type ListBudgetConfigurationsResponse struct {
	Budgets types.List `tfsdk:"budgets"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *ListBudgetConfigurationsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListBudgetConfigurationsResponse) {
}

func (toState *ListBudgetConfigurationsResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListBudgetConfigurationsResponse) {
}

func (c ListBudgetConfigurationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budgets"] = attrs["budgets"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBudgetConfigurationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListBudgetConfigurationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budgets": reflect.TypeOf(BudgetConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetConfigurationsResponse
// only implements ToObjectValue() and Type().
func (o ListBudgetConfigurationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budgets":         o.Budgets,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budgets": basetypes.ListType{
				ElemType: BudgetConfiguration{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetBudgets returns the value of the Budgets field in ListBudgetConfigurationsResponse as
// a slice of BudgetConfiguration values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListBudgetConfigurationsResponse) GetBudgets(ctx context.Context) ([]BudgetConfiguration, bool) {
	if o.Budgets.IsNull() || o.Budgets.IsUnknown() {
		return nil, false
	}
	var v []BudgetConfiguration
	d := o.Budgets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBudgets sets the value of the Budgets field in ListBudgetConfigurationsResponse.
func (o *ListBudgetConfigurationsResponse) SetBudgets(ctx context.Context, v []BudgetConfiguration) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budgets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Budgets = types.ListValueMust(t, vs)
}

type ListBudgetPoliciesRequest struct {
	// A filter to apply to the list of policies.
	FilterBy types.Object `tfsdk:"-"`
	// The maximum number of budget policies to return. If unspecified, at most
	// 100 budget policies will be returned. The maximum value is 1000; values
	// above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListServerlessPolicies` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	//
	// When paginating, all other parameters provided to
	// `ListServerlessPoliciesRequest` must match the call that provided the
	// page token.
	PageToken types.String `tfsdk:"-"`
	// The sort specification.
	SortSpec types.Object `tfsdk:"-"`
}

func (toState *ListBudgetPoliciesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListBudgetPoliciesRequest) {
	if !fromPlan.FilterBy.IsNull() && !fromPlan.FilterBy.IsUnknown() {
		if toStateFilterBy, ok := toState.GetFilterBy(ctx); ok {
			if fromPlanFilterBy, ok := fromPlan.GetFilterBy(ctx); ok {
				toStateFilterBy.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFilterBy)
				toState.SetFilterBy(ctx, toStateFilterBy)
			}
		}
	}
	if !fromPlan.SortSpec.IsNull() && !fromPlan.SortSpec.IsUnknown() {
		if toStateSortSpec, ok := toState.GetSortSpec(ctx); ok {
			if fromPlanSortSpec, ok := fromPlan.GetSortSpec(ctx); ok {
				toStateSortSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSortSpec)
				toState.SetSortSpec(ctx, toStateSortSpec)
			}
		}
	}
}

func (toState *ListBudgetPoliciesRequest) SyncFieldsDuringRead(ctx context.Context, fromState ListBudgetPoliciesRequest) {
	if !fromState.FilterBy.IsNull() && !fromState.FilterBy.IsUnknown() {
		if toStateFilterBy, ok := toState.GetFilterBy(ctx); ok {
			if fromStateFilterBy, ok := fromState.GetFilterBy(ctx); ok {
				toStateFilterBy.SyncFieldsDuringRead(ctx, fromStateFilterBy)
				toState.SetFilterBy(ctx, toStateFilterBy)
			}
		}
	}
	if !fromState.SortSpec.IsNull() && !fromState.SortSpec.IsUnknown() {
		if toStateSortSpec, ok := toState.GetSortSpec(ctx); ok {
			if fromStateSortSpec, ok := fromState.GetSortSpec(ctx); ok {
				toStateSortSpec.SyncFieldsDuringRead(ctx, fromStateSortSpec)
				toState.SetSortSpec(ctx, toStateSortSpec)
			}
		}
	}
}

func (c ListBudgetPoliciesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter_by"] = attrs["filter_by"].SetOptional()
	attrs["sort_spec"] = attrs["sort_spec"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBudgetPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListBudgetPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter_by": reflect.TypeOf(Filter{}),
		"sort_spec": reflect.TypeOf(SortSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetPoliciesRequest
// only implements ToObjectValue() and Type().
func (o ListBudgetPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter_by":  o.FilterBy,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
			"sort_spec":  o.SortSpec,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListBudgetPoliciesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_by":  Filter{}.Type(ctx),
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"sort_spec":  SortSpec{}.Type(ctx),
		},
	}
}

// GetFilterBy returns the value of the FilterBy field in ListBudgetPoliciesRequest as
// a Filter value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListBudgetPoliciesRequest) GetFilterBy(ctx context.Context) (Filter, bool) {
	var e Filter
	if o.FilterBy.IsNull() || o.FilterBy.IsUnknown() {
		return e, false
	}
	var v Filter
	d := o.FilterBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilterBy sets the value of the FilterBy field in ListBudgetPoliciesRequest.
func (o *ListBudgetPoliciesRequest) SetFilterBy(ctx context.Context, v Filter) {
	vs := v.ToObjectValue(ctx)
	o.FilterBy = vs
}

// GetSortSpec returns the value of the SortSpec field in ListBudgetPoliciesRequest as
// a SortSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListBudgetPoliciesRequest) GetSortSpec(ctx context.Context) (SortSpec, bool) {
	var e SortSpec
	if o.SortSpec.IsNull() || o.SortSpec.IsUnknown() {
		return e, false
	}
	var v SortSpec
	d := o.SortSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSortSpec sets the value of the SortSpec field in ListBudgetPoliciesRequest.
func (o *ListBudgetPoliciesRequest) SetSortSpec(ctx context.Context, v SortSpec) {
	vs := v.ToObjectValue(ctx)
	o.SortSpec = vs
}

// A list of policies.
type ListBudgetPoliciesResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Policies types.List `tfsdk:"policies"`
	// A token that can be sent as `page_token` to retrieve the previous page.
	// In this field is omitted, there are no previous pages.
	PreviousPageToken types.String `tfsdk:"previous_page_token"`
}

func (toState *ListBudgetPoliciesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListBudgetPoliciesResponse) {
}

func (toState *ListBudgetPoliciesResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListBudgetPoliciesResponse) {
}

func (c ListBudgetPoliciesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["policies"] = attrs["policies"].SetOptional()
	attrs["previous_page_token"] = attrs["previous_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBudgetPoliciesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListBudgetPoliciesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(BudgetPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetPoliciesResponse
// only implements ToObjectValue() and Type().
func (o ListBudgetPoliciesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":     o.NextPageToken,
			"policies":            o.Policies,
			"previous_page_token": o.PreviousPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListBudgetPoliciesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"policies": basetypes.ListType{
				ElemType: BudgetPolicy{}.Type(ctx),
			},
			"previous_page_token": types.StringType,
		},
	}
}

// GetPolicies returns the value of the Policies field in ListBudgetPoliciesResponse as
// a slice of BudgetPolicy values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListBudgetPoliciesResponse) GetPolicies(ctx context.Context) ([]BudgetPolicy, bool) {
	if o.Policies.IsNull() || o.Policies.IsUnknown() {
		return nil, false
	}
	var v []BudgetPolicy
	d := o.Policies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicies sets the value of the Policies field in ListBudgetPoliciesResponse.
func (o *ListBudgetPoliciesResponse) SetPolicies(ctx context.Context, v []BudgetPolicy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Policies = types.ListValueMust(t, vs)
}

type ListLogDeliveryRequest struct {
	// The Credentials id to filter the search results with
	CredentialsId types.String `tfsdk:"-"`
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken types.String `tfsdk:"-"`
	// The log delivery status to filter the search results with
	Status types.String `tfsdk:"-"`
	// The Storage Configuration id to filter the search results with
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (toState *ListLogDeliveryRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListLogDeliveryRequest) {
}

func (toState *ListLogDeliveryRequest) SyncFieldsDuringRead(ctx context.Context, fromState ListLogDeliveryRequest) {
}

func (c ListLogDeliveryRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListLogDeliveryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListLogDeliveryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListLogDeliveryRequest
// only implements ToObjectValue() and Type().
func (o ListLogDeliveryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials_id":           o.CredentialsId,
			"page_token":               o.PageToken,
			"status":                   o.Status,
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListLogDeliveryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id":           types.StringType,
			"page_token":               types.StringType,
			"status":                   types.StringType,
			"storage_configuration_id": types.StringType,
		},
	}
}

// * Log Delivery Configuration
type LogDeliveryConfiguration struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id"`
	// The unique UUID of log delivery configuration
	ConfigId types.String `tfsdk:"config_id"`
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName types.String `tfsdk:"config_name"`
	// Time in epoch milliseconds when the log delivery configuration was
	// created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix types.String `tfsdk:"delivery_path_prefix"`
	// This field applies only if log_type is BILLABLE_USAGE. This is the
	// optional start month and year for delivery, specified in YYYY-MM format.
	// Defaults to current year and month. BILLABLE_USAGE logs are not available
	// for usage before March 2019 (2019-03).
	DeliveryStartTime types.String `tfsdk:"delivery_start_time"`
	// The LogDeliveryStatus of this log delivery configuration
	LogDeliveryStatus types.Object `tfsdk:"log_delivery_status"`
	// Log delivery type. Supported values are: * `BILLABLE_USAGE` — Configure
	// [billable usage log delivery]. For the CSV schema, see the [View billable
	// usage]. * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON
	// schema, see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType types.String `tfsdk:"log_type"`
	// The file type of log delivery. * If `log_type` is `BILLABLE_USAGE`, this
	// value must be `CSV`. Only the CSV (comma-separated values) format is
	// supported. For the schema, see the [View billable usage] * If `log_type`
	// is `AUDIT_LOGS`, this value must be `JSON`. Only the JSON (JavaScript
	// Object Notation) format is supported. For the schema, see the
	// [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat types.String `tfsdk:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status types.String `tfsdk:"status"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
	// Time in epoch milliseconds when the log delivery configuration was
	// updated.
	UpdateTime types.Int64 `tfsdk:"update_time"`
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
	WorkspaceIdsFilter types.List `tfsdk:"workspace_ids_filter"`
}

func (toState *LogDeliveryConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogDeliveryConfiguration) {
	if !fromPlan.LogDeliveryStatus.IsNull() && !fromPlan.LogDeliveryStatus.IsUnknown() {
		if toStateLogDeliveryStatus, ok := toState.GetLogDeliveryStatus(ctx); ok {
			if fromPlanLogDeliveryStatus, ok := fromPlan.GetLogDeliveryStatus(ctx); ok {
				toStateLogDeliveryStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanLogDeliveryStatus)
				toState.SetLogDeliveryStatus(ctx, toStateLogDeliveryStatus)
			}
		}
	}
}

func (toState *LogDeliveryConfiguration) SyncFieldsDuringRead(ctx context.Context, fromState LogDeliveryConfiguration) {
	if !fromState.LogDeliveryStatus.IsNull() && !fromState.LogDeliveryStatus.IsUnknown() {
		if toStateLogDeliveryStatus, ok := toState.GetLogDeliveryStatus(ctx); ok {
			if fromStateLogDeliveryStatus, ok := fromState.GetLogDeliveryStatus(ctx); ok {
				toStateLogDeliveryStatus.SyncFieldsDuringRead(ctx, fromStateLogDeliveryStatus)
				toState.SetLogDeliveryStatus(ctx, toStateLogDeliveryStatus)
			}
		}
	}
}

func (c LogDeliveryConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["config_id"] = attrs["config_id"].SetOptional()
	attrs["config_name"] = attrs["config_name"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["credentials_id"] = attrs["credentials_id"].SetRequired()
	attrs["delivery_path_prefix"] = attrs["delivery_path_prefix"].SetOptional()
	attrs["delivery_start_time"] = attrs["delivery_start_time"].SetOptional()
	attrs["log_delivery_status"] = attrs["log_delivery_status"].SetOptional()
	attrs["log_type"] = attrs["log_type"].SetRequired()
	attrs["output_format"] = attrs["output_format"].SetRequired()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetRequired()
	attrs["update_time"] = attrs["update_time"].SetOptional()
	attrs["workspace_ids_filter"] = attrs["workspace_ids_filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogDeliveryConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogDeliveryConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_status":  reflect.TypeOf(LogDeliveryStatus{}),
		"workspace_ids_filter": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogDeliveryConfiguration
// only implements ToObjectValue() and Type().
func (o LogDeliveryConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LogDeliveryConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":               types.StringType,
			"config_id":                types.StringType,
			"config_name":              types.StringType,
			"creation_time":            types.Int64Type,
			"credentials_id":           types.StringType,
			"delivery_path_prefix":     types.StringType,
			"delivery_start_time":      types.StringType,
			"log_delivery_status":      LogDeliveryStatus{}.Type(ctx),
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

// GetLogDeliveryStatus returns the value of the LogDeliveryStatus field in LogDeliveryConfiguration as
// a LogDeliveryStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *LogDeliveryConfiguration) GetLogDeliveryStatus(ctx context.Context) (LogDeliveryStatus, bool) {
	var e LogDeliveryStatus
	if o.LogDeliveryStatus.IsNull() || o.LogDeliveryStatus.IsUnknown() {
		return e, false
	}
	var v LogDeliveryStatus
	d := o.LogDeliveryStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLogDeliveryStatus sets the value of the LogDeliveryStatus field in LogDeliveryConfiguration.
func (o *LogDeliveryConfiguration) SetLogDeliveryStatus(ctx context.Context, v LogDeliveryStatus) {
	vs := v.ToObjectValue(ctx)
	o.LogDeliveryStatus = vs
}

// GetWorkspaceIdsFilter returns the value of the WorkspaceIdsFilter field in LogDeliveryConfiguration as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *LogDeliveryConfiguration) GetWorkspaceIdsFilter(ctx context.Context) ([]types.Int64, bool) {
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

// SetWorkspaceIdsFilter sets the value of the WorkspaceIdsFilter field in LogDeliveryConfiguration.
func (o *LogDeliveryConfiguration) SetWorkspaceIdsFilter(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_ids_filter"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WorkspaceIdsFilter = types.ListValueMust(t, vs)
}

type LogDeliveryStatus struct {
	// The UTC time for the latest log delivery attempt.
	LastAttemptTime types.String `tfsdk:"last_attempt_time"`
	// The UTC time for the latest successful log delivery.
	LastSuccessfulAttemptTime types.String `tfsdk:"last_successful_attempt_time"`
	// Informative message about the latest log delivery attempt. If the log
	// delivery fails with USER_FAILURE, error details will be provided for
	// fixing misconfigurations in cloud permissions.
	Message types.String `tfsdk:"message"`
	// Enum that describes the status. Possible values are: * `CREATED`: There
	// were no log delivery attempts since the config was created. *
	// `SUCCEEDED`: The latest attempt of log delivery has succeeded completely.
	// * `USER_FAILURE`: The latest attempt of log delivery failed because of
	// misconfiguration of customer provided permissions on role or storage. *
	// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
	// Databricks internal error. Contact support if it doesn't go away soon. *
	// `NOT_FOUND`: The log delivery status as the configuration has been
	// disabled since the release of this feature or there are no workspaces in
	// the account.
	Status types.String `tfsdk:"status"`
}

func (toState *LogDeliveryStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogDeliveryStatus) {
}

func (toState *LogDeliveryStatus) SyncFieldsDuringRead(ctx context.Context, fromState LogDeliveryStatus) {
}

func (c LogDeliveryStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_attempt_time"] = attrs["last_attempt_time"].SetOptional()
	attrs["last_successful_attempt_time"] = attrs["last_successful_attempt_time"].SetOptional()
	attrs["message"] = attrs["message"].SetRequired()
	attrs["status"] = attrs["status"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogDeliveryStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogDeliveryStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogDeliveryStatus
// only implements ToObjectValue() and Type().
func (o LogDeliveryStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LogDeliveryStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_attempt_time":            types.StringType,
			"last_successful_attempt_time": types.StringType,
			"message":                      types.StringType,
			"status":                       types.StringType,
		},
	}
}

type PatchStatusResponse struct {
}

func (toState *PatchStatusResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PatchStatusResponse) {
}

func (toState *PatchStatusResponse) SyncFieldsDuringRead(ctx context.Context, fromState PatchStatusResponse) {
}

func (c PatchStatusResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchStatusResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchStatusResponse
// only implements ToObjectValue() and Type().
func (o PatchStatusResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PatchStatusResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SortSpec struct {
	// Whether to sort in descending order.
	Descending types.Bool `tfsdk:"descending"`
	// The filed to sort by
	Field types.String `tfsdk:"field"`
}

func (toState *SortSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SortSpec) {
}

func (toState *SortSpec) SyncFieldsDuringRead(ctx context.Context, fromState SortSpec) {
}

func (c SortSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["descending"] = attrs["descending"].SetOptional()
	attrs["field"] = attrs["field"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SortSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SortSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SortSpec
// only implements ToObjectValue() and Type().
func (o SortSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"descending": o.Descending,
			"field":      o.Field,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SortSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"descending": types.BoolType,
			"field":      types.StringType,
		},
	}
}

type UpdateBudgetConfigurationBudget struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations types.List `tfsdk:"alert_configurations"`
	// Databricks budget configuration ID.
	BudgetConfigurationId types.String `tfsdk:"budget_configuration_id"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName types.String `tfsdk:"display_name"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter types.Object `tfsdk:"filter"`
}

func (toState *UpdateBudgetConfigurationBudget) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateBudgetConfigurationBudget) {
	if !fromPlan.Filter.IsNull() && !fromPlan.Filter.IsUnknown() {
		if toStateFilter, ok := toState.GetFilter(ctx); ok {
			if fromPlanFilter, ok := fromPlan.GetFilter(ctx); ok {
				toStateFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFilter)
				toState.SetFilter(ctx, toStateFilter)
			}
		}
	}
}

func (toState *UpdateBudgetConfigurationBudget) SyncFieldsDuringRead(ctx context.Context, fromState UpdateBudgetConfigurationBudget) {
	if !fromState.Filter.IsNull() && !fromState.Filter.IsUnknown() {
		if toStateFilter, ok := toState.GetFilter(ctx); ok {
			if fromStateFilter, ok := fromState.GetFilter(ctx); ok {
				toStateFilter.SyncFieldsDuringRead(ctx, fromStateFilter)
				toState.SetFilter(ctx, toStateFilter)
			}
		}
	}
}

func (c UpdateBudgetConfigurationBudget) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["alert_configurations"] = attrs["alert_configurations"].SetOptional()
	attrs["budget_configuration_id"] = attrs["budget_configuration_id"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetConfigurationBudget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateBudgetConfigurationBudget) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_configurations": reflect.TypeOf(AlertConfiguration{}),
		"filter":               reflect.TypeOf(BudgetConfigurationFilter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetConfigurationBudget
// only implements ToObjectValue() and Type().
func (o UpdateBudgetConfigurationBudget) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateBudgetConfigurationBudget) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"alert_configurations": basetypes.ListType{
				ElemType: AlertConfiguration{}.Type(ctx),
			},
			"budget_configuration_id": types.StringType,
			"display_name":            types.StringType,
			"filter":                  BudgetConfigurationFilter{}.Type(ctx),
		},
	}
}

// GetAlertConfigurations returns the value of the AlertConfigurations field in UpdateBudgetConfigurationBudget as
// a slice of AlertConfiguration values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetConfigurationBudget) GetAlertConfigurations(ctx context.Context) ([]AlertConfiguration, bool) {
	if o.AlertConfigurations.IsNull() || o.AlertConfigurations.IsUnknown() {
		return nil, false
	}
	var v []AlertConfiguration
	d := o.AlertConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlertConfigurations sets the value of the AlertConfigurations field in UpdateBudgetConfigurationBudget.
func (o *UpdateBudgetConfigurationBudget) SetAlertConfigurations(ctx context.Context, v []AlertConfiguration) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AlertConfigurations = types.ListValueMust(t, vs)
}

// GetFilter returns the value of the Filter field in UpdateBudgetConfigurationBudget as
// a BudgetConfigurationFilter value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetConfigurationBudget) GetFilter(ctx context.Context) (BudgetConfigurationFilter, bool) {
	var e BudgetConfigurationFilter
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v BudgetConfigurationFilter
	d := o.Filter.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilter sets the value of the Filter field in UpdateBudgetConfigurationBudget.
func (o *UpdateBudgetConfigurationBudget) SetFilter(ctx context.Context, v BudgetConfigurationFilter) {
	vs := v.ToObjectValue(ctx)
	o.Filter = vs
}

type UpdateBudgetConfigurationRequest struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID.
	Budget types.Object `tfsdk:"budget"`
	// The Databricks budget configuration ID.
	BudgetId types.String `tfsdk:"-"`
}

func (toState *UpdateBudgetConfigurationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateBudgetConfigurationRequest) {
	if !fromPlan.Budget.IsNull() && !fromPlan.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromPlanBudget, ok := fromPlan.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (toState *UpdateBudgetConfigurationRequest) SyncFieldsDuringRead(ctx context.Context, fromState UpdateBudgetConfigurationRequest) {
	if !fromState.Budget.IsNull() && !fromState.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromStateBudget, ok := fromState.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringRead(ctx, fromStateBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (c UpdateBudgetConfigurationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["budget_id"] = attrs["budget_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateBudgetConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(UpdateBudgetConfigurationBudget{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetConfigurationRequest
// only implements ToObjectValue() and Type().
func (o UpdateBudgetConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget":    o.Budget,
			"budget_id": o.BudgetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget":    UpdateBudgetConfigurationBudget{}.Type(ctx),
			"budget_id": types.StringType,
		},
	}
}

// GetBudget returns the value of the Budget field in UpdateBudgetConfigurationRequest as
// a UpdateBudgetConfigurationBudget value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetConfigurationRequest) GetBudget(ctx context.Context) (UpdateBudgetConfigurationBudget, bool) {
	var e UpdateBudgetConfigurationBudget
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v UpdateBudgetConfigurationBudget
	d := o.Budget.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBudget sets the value of the Budget field in UpdateBudgetConfigurationRequest.
func (o *UpdateBudgetConfigurationRequest) SetBudget(ctx context.Context, v UpdateBudgetConfigurationBudget) {
	vs := v.ToObjectValue(ctx)
	o.Budget = vs
}

type UpdateBudgetConfigurationResponse struct {
	// The updated budget.
	Budget types.Object `tfsdk:"budget"`
}

func (toState *UpdateBudgetConfigurationResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateBudgetConfigurationResponse) {
	if !fromPlan.Budget.IsNull() && !fromPlan.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromPlanBudget, ok := fromPlan.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (toState *UpdateBudgetConfigurationResponse) SyncFieldsDuringRead(ctx context.Context, fromState UpdateBudgetConfigurationResponse) {
	if !fromState.Budget.IsNull() && !fromState.Budget.IsUnknown() {
		if toStateBudget, ok := toState.GetBudget(ctx); ok {
			if fromStateBudget, ok := fromState.GetBudget(ctx); ok {
				toStateBudget.SyncFieldsDuringRead(ctx, fromStateBudget)
				toState.SetBudget(ctx, toStateBudget)
			}
		}
	}
}

func (c UpdateBudgetConfigurationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateBudgetConfigurationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(BudgetConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetConfigurationResponse
// only implements ToObjectValue() and Type().
func (o UpdateBudgetConfigurationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": o.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": BudgetConfiguration{}.Type(ctx),
		},
	}
}

// GetBudget returns the value of the Budget field in UpdateBudgetConfigurationResponse as
// a BudgetConfiguration value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetConfigurationResponse) GetBudget(ctx context.Context) (BudgetConfiguration, bool) {
	var e BudgetConfiguration
	if o.Budget.IsNull() || o.Budget.IsUnknown() {
		return e, false
	}
	var v BudgetConfiguration
	d := o.Budget.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBudget sets the value of the Budget field in UpdateBudgetConfigurationResponse.
func (o *UpdateBudgetConfigurationResponse) SetBudget(ctx context.Context, v BudgetConfiguration) {
	vs := v.ToObjectValue(ctx)
	o.Budget = vs
}

type UpdateBudgetPolicyRequest struct {
	// DEPRECATED. This is redundant field as LimitConfig is part of the
	// BudgetPolicy
	LimitConfig types.Object `tfsdk:"-"`
	// The policy to update. `creator_user_id` cannot be specified in the
	// request. All other fields must be specified even if not changed. The
	// `policy_id` is used to identify the policy to update.
	Policy types.Object `tfsdk:"policy"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId types.String `tfsdk:"-"`
}

func (toState *UpdateBudgetPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateBudgetPolicyRequest) {
	if !fromPlan.LimitConfig.IsNull() && !fromPlan.LimitConfig.IsUnknown() {
		if toStateLimitConfig, ok := toState.GetLimitConfig(ctx); ok {
			if fromPlanLimitConfig, ok := fromPlan.GetLimitConfig(ctx); ok {
				toStateLimitConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanLimitConfig)
				toState.SetLimitConfig(ctx, toStateLimitConfig)
			}
		}
	}
	if !fromPlan.Policy.IsNull() && !fromPlan.Policy.IsUnknown() {
		if toStatePolicy, ok := toState.GetPolicy(ctx); ok {
			if fromPlanPolicy, ok := fromPlan.GetPolicy(ctx); ok {
				toStatePolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanPolicy)
				toState.SetPolicy(ctx, toStatePolicy)
			}
		}
	}
}

func (toState *UpdateBudgetPolicyRequest) SyncFieldsDuringRead(ctx context.Context, fromState UpdateBudgetPolicyRequest) {
	if !fromState.LimitConfig.IsNull() && !fromState.LimitConfig.IsUnknown() {
		if toStateLimitConfig, ok := toState.GetLimitConfig(ctx); ok {
			if fromStateLimitConfig, ok := fromState.GetLimitConfig(ctx); ok {
				toStateLimitConfig.SyncFieldsDuringRead(ctx, fromStateLimitConfig)
				toState.SetLimitConfig(ctx, toStateLimitConfig)
			}
		}
	}
	if !fromState.Policy.IsNull() && !fromState.Policy.IsUnknown() {
		if toStatePolicy, ok := toState.GetPolicy(ctx); ok {
			if fromStatePolicy, ok := fromState.GetPolicy(ctx); ok {
				toStatePolicy.SyncFieldsDuringRead(ctx, fromStatePolicy)
				toState.SetPolicy(ctx, toStatePolicy)
			}
		}
	}
}

func (c UpdateBudgetPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetRequired()
	attrs["limit_config"] = attrs["limit_config"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateBudgetPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"limit_config": reflect.TypeOf(LimitConfig{}),
		"policy":       reflect.TypeOf(BudgetPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetPolicyRequest
// only implements ToObjectValue() and Type().
func (o UpdateBudgetPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"limit_config": o.LimitConfig,
			"policy":       o.Policy,
			"policy_id":    o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateBudgetPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"limit_config": LimitConfig{}.Type(ctx),
			"policy":       BudgetPolicy{}.Type(ctx),
			"policy_id":    types.StringType,
		},
	}
}

// GetLimitConfig returns the value of the LimitConfig field in UpdateBudgetPolicyRequest as
// a LimitConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetPolicyRequest) GetLimitConfig(ctx context.Context) (LimitConfig, bool) {
	var e LimitConfig
	if o.LimitConfig.IsNull() || o.LimitConfig.IsUnknown() {
		return e, false
	}
	var v LimitConfig
	d := o.LimitConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLimitConfig sets the value of the LimitConfig field in UpdateBudgetPolicyRequest.
func (o *UpdateBudgetPolicyRequest) SetLimitConfig(ctx context.Context, v LimitConfig) {
	vs := v.ToObjectValue(ctx)
	o.LimitConfig = vs
}

// GetPolicy returns the value of the Policy field in UpdateBudgetPolicyRequest as
// a BudgetPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateBudgetPolicyRequest) GetPolicy(ctx context.Context) (BudgetPolicy, bool) {
	var e BudgetPolicy
	if o.Policy.IsNull() || o.Policy.IsUnknown() {
		return e, false
	}
	var v BudgetPolicy
	d := o.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicy sets the value of the Policy field in UpdateBudgetPolicyRequest.
func (o *UpdateBudgetPolicyRequest) SetPolicy(ctx context.Context, v BudgetPolicy) {
	vs := v.ToObjectValue(ctx)
	o.Policy = vs
}

// * Update Log Delivery Configuration
type UpdateLogDeliveryConfigurationStatusRequest struct {
	// The log delivery configuration id of customer
	LogDeliveryConfigurationId types.String `tfsdk:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status types.String `tfsdk:"status"`
}

func (toState *UpdateLogDeliveryConfigurationStatusRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateLogDeliveryConfigurationStatusRequest) {
}

func (toState *UpdateLogDeliveryConfigurationStatusRequest) SyncFieldsDuringRead(ctx context.Context, fromState UpdateLogDeliveryConfigurationStatusRequest) {
}

func (c UpdateLogDeliveryConfigurationStatusRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["status"] = attrs["status"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["log_delivery_configuration_id"] = attrs["log_delivery_configuration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLogDeliveryConfigurationStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateLogDeliveryConfigurationStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLogDeliveryConfigurationStatusRequest
// only implements ToObjectValue() and Type().
func (o UpdateLogDeliveryConfigurationStatusRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration_id": o.LogDeliveryConfigurationId,
			"status":                        o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateLogDeliveryConfigurationStatusRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration_id": types.StringType,
			"status":                        types.StringType,
		},
	}
}

// * Properties of the new log delivery configuration.
type WrappedCreateLogDeliveryConfiguration struct {
	LogDeliveryConfiguration types.Object `tfsdk:"log_delivery_configuration"`
}

func (toState *WrappedCreateLogDeliveryConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan WrappedCreateLogDeliveryConfiguration) {
	if !fromPlan.LogDeliveryConfiguration.IsNull() && !fromPlan.LogDeliveryConfiguration.IsUnknown() {
		if toStateLogDeliveryConfiguration, ok := toState.GetLogDeliveryConfiguration(ctx); ok {
			if fromPlanLogDeliveryConfiguration, ok := fromPlan.GetLogDeliveryConfiguration(ctx); ok {
				toStateLogDeliveryConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanLogDeliveryConfiguration)
				toState.SetLogDeliveryConfiguration(ctx, toStateLogDeliveryConfiguration)
			}
		}
	}
}

func (toState *WrappedCreateLogDeliveryConfiguration) SyncFieldsDuringRead(ctx context.Context, fromState WrappedCreateLogDeliveryConfiguration) {
	if !fromState.LogDeliveryConfiguration.IsNull() && !fromState.LogDeliveryConfiguration.IsUnknown() {
		if toStateLogDeliveryConfiguration, ok := toState.GetLogDeliveryConfiguration(ctx); ok {
			if fromStateLogDeliveryConfiguration, ok := fromState.GetLogDeliveryConfiguration(ctx); ok {
				toStateLogDeliveryConfiguration.SyncFieldsDuringRead(ctx, fromStateLogDeliveryConfiguration)
				toState.SetLogDeliveryConfiguration(ctx, toStateLogDeliveryConfiguration)
			}
		}
	}
}

func (c WrappedCreateLogDeliveryConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["log_delivery_configuration"] = attrs["log_delivery_configuration"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WrappedCreateLogDeliveryConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WrappedCreateLogDeliveryConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configuration": reflect.TypeOf(CreateLogDeliveryConfigurationParams{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WrappedCreateLogDeliveryConfiguration
// only implements ToObjectValue() and Type().
func (o WrappedCreateLogDeliveryConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration": o.LogDeliveryConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WrappedCreateLogDeliveryConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration": CreateLogDeliveryConfigurationParams{}.Type(ctx),
		},
	}
}

// GetLogDeliveryConfiguration returns the value of the LogDeliveryConfiguration field in WrappedCreateLogDeliveryConfiguration as
// a CreateLogDeliveryConfigurationParams value.
// If the field is unknown or null, the boolean return value is false.
func (o *WrappedCreateLogDeliveryConfiguration) GetLogDeliveryConfiguration(ctx context.Context) (CreateLogDeliveryConfigurationParams, bool) {
	var e CreateLogDeliveryConfigurationParams
	if o.LogDeliveryConfiguration.IsNull() || o.LogDeliveryConfiguration.IsUnknown() {
		return e, false
	}
	var v CreateLogDeliveryConfigurationParams
	d := o.LogDeliveryConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in WrappedCreateLogDeliveryConfiguration.
func (o *WrappedCreateLogDeliveryConfiguration) SetLogDeliveryConfiguration(ctx context.Context, v CreateLogDeliveryConfigurationParams) {
	vs := v.ToObjectValue(ctx)
	o.LogDeliveryConfiguration = vs
}

type WrappedLogDeliveryConfiguration struct {
	// The created log delivery configuration
	LogDeliveryConfiguration types.Object `tfsdk:"log_delivery_configuration"`
}

func (toState *WrappedLogDeliveryConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan WrappedLogDeliveryConfiguration) {
	if !fromPlan.LogDeliveryConfiguration.IsNull() && !fromPlan.LogDeliveryConfiguration.IsUnknown() {
		if toStateLogDeliveryConfiguration, ok := toState.GetLogDeliveryConfiguration(ctx); ok {
			if fromPlanLogDeliveryConfiguration, ok := fromPlan.GetLogDeliveryConfiguration(ctx); ok {
				toStateLogDeliveryConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanLogDeliveryConfiguration)
				toState.SetLogDeliveryConfiguration(ctx, toStateLogDeliveryConfiguration)
			}
		}
	}
}

func (toState *WrappedLogDeliveryConfiguration) SyncFieldsDuringRead(ctx context.Context, fromState WrappedLogDeliveryConfiguration) {
	if !fromState.LogDeliveryConfiguration.IsNull() && !fromState.LogDeliveryConfiguration.IsUnknown() {
		if toStateLogDeliveryConfiguration, ok := toState.GetLogDeliveryConfiguration(ctx); ok {
			if fromStateLogDeliveryConfiguration, ok := fromState.GetLogDeliveryConfiguration(ctx); ok {
				toStateLogDeliveryConfiguration.SyncFieldsDuringRead(ctx, fromStateLogDeliveryConfiguration)
				toState.SetLogDeliveryConfiguration(ctx, toStateLogDeliveryConfiguration)
			}
		}
	}
}

func (c WrappedLogDeliveryConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["log_delivery_configuration"] = attrs["log_delivery_configuration"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WrappedLogDeliveryConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WrappedLogDeliveryConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configuration": reflect.TypeOf(LogDeliveryConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WrappedLogDeliveryConfiguration
// only implements ToObjectValue() and Type().
func (o WrappedLogDeliveryConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration": o.LogDeliveryConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration": LogDeliveryConfiguration{}.Type(ctx),
		},
	}
}

// GetLogDeliveryConfiguration returns the value of the LogDeliveryConfiguration field in WrappedLogDeliveryConfiguration as
// a LogDeliveryConfiguration value.
// If the field is unknown or null, the boolean return value is false.
func (o *WrappedLogDeliveryConfiguration) GetLogDeliveryConfiguration(ctx context.Context) (LogDeliveryConfiguration, bool) {
	var e LogDeliveryConfiguration
	if o.LogDeliveryConfiguration.IsNull() || o.LogDeliveryConfiguration.IsUnknown() {
		return e, false
	}
	var v LogDeliveryConfiguration
	d := o.LogDeliveryConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in WrappedLogDeliveryConfiguration.
func (o *WrappedLogDeliveryConfiguration) SetLogDeliveryConfiguration(ctx context.Context, v LogDeliveryConfiguration) {
	vs := v.ToObjectValue(ctx)
	o.LogDeliveryConfiguration = vs
}

type WrappedLogDeliveryConfigurations struct {
	// Log delivery configurations were returned successfully.
	LogDeliveryConfigurations types.List `tfsdk:"log_delivery_configurations"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *WrappedLogDeliveryConfigurations) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan WrappedLogDeliveryConfigurations) {
}

func (toState *WrappedLogDeliveryConfigurations) SyncFieldsDuringRead(ctx context.Context, fromState WrappedLogDeliveryConfigurations) {
}

func (c WrappedLogDeliveryConfigurations) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["log_delivery_configurations"] = attrs["log_delivery_configurations"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WrappedLogDeliveryConfigurations.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WrappedLogDeliveryConfigurations) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configurations": reflect.TypeOf(LogDeliveryConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WrappedLogDeliveryConfigurations
// only implements ToObjectValue() and Type().
func (o WrappedLogDeliveryConfigurations) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configurations": o.LogDeliveryConfigurations,
			"next_page_token":             o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfigurations) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configurations": basetypes.ListType{
				ElemType: LogDeliveryConfiguration{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetLogDeliveryConfigurations returns the value of the LogDeliveryConfigurations field in WrappedLogDeliveryConfigurations as
// a slice of LogDeliveryConfiguration values.
// If the field is unknown or null, the boolean return value is false.
func (o *WrappedLogDeliveryConfigurations) GetLogDeliveryConfigurations(ctx context.Context) ([]LogDeliveryConfiguration, bool) {
	if o.LogDeliveryConfigurations.IsNull() || o.LogDeliveryConfigurations.IsUnknown() {
		return nil, false
	}
	var v []LogDeliveryConfiguration
	d := o.LogDeliveryConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLogDeliveryConfigurations sets the value of the LogDeliveryConfigurations field in WrappedLogDeliveryConfigurations.
func (o *WrappedLogDeliveryConfigurations) SetLogDeliveryConfigurations(ctx context.Context, v []LogDeliveryConfiguration) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LogDeliveryConfigurations = types.ListValueMust(t, vs)
}
