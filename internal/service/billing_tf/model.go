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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ActionConfiguration struct {
	// Databricks action configuration ID.
	ActionConfigurationId types.String `tfsdk:"action_configuration_id" tf:"optional"`
	// The type of the action.
	ActionType types.String `tfsdk:"action_type" tf:"optional"`
	// Target for the action. For example, an email address.
	Target types.String `tfsdk:"target" tf:"optional"`
}

func (newState *ActionConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan ActionConfiguration) {
}

func (newState *ActionConfiguration) SyncEffectiveFieldsDuringRead(existingState ActionConfiguration) {
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

func (newState *AlertConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertConfiguration) {
}

func (newState *AlertConfiguration) SyncEffectiveFieldsDuringRead(existingState AlertConfiguration) {
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

func (newState *BudgetConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfiguration) {
}

func (newState *BudgetConfiguration) SyncEffectiveFieldsDuringRead(existingState BudgetConfiguration) {
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
			"filter": basetypes.ListType{
				ElemType: BudgetConfigurationFilter{}.Type(ctx),
			},
			"update_time": types.Int64Type,
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
	var v []BudgetConfigurationFilter
	d := o.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in BudgetConfiguration.
func (o *BudgetConfiguration) SetFilter(ctx context.Context, v BudgetConfigurationFilter) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	o.Filter = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilter struct {
	// A list of tag keys and values that will limit the budget to usage that
	// includes those specific custom tags. Tags are case-sensitive and should
	// be entered exactly as they appear in your usage data.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// If provided, usage must match with the provided Databricks workspace IDs.
	WorkspaceId types.List `tfsdk:"workspace_id" tf:"optional,object"`
}

func (newState *BudgetConfigurationFilter) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfigurationFilter) {
}

func (newState *BudgetConfigurationFilter) SyncEffectiveFieldsDuringRead(existingState BudgetConfigurationFilter) {
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
			"workspace_id": basetypes.ListType{
				ElemType: BudgetConfigurationFilterWorkspaceIdClause{}.Type(ctx),
			},
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
	var v []BudgetConfigurationFilterWorkspaceIdClause
	d := o.WorkspaceId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceId sets the value of the WorkspaceId field in BudgetConfigurationFilter.
func (o *BudgetConfigurationFilter) SetWorkspaceId(ctx context.Context, v BudgetConfigurationFilterWorkspaceIdClause) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_id"]
	o.WorkspaceId = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilterClause struct {
	Operator types.String `tfsdk:"operator" tf:"optional"`

	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *BudgetConfigurationFilterClause) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfigurationFilterClause) {
}

func (newState *BudgetConfigurationFilterClause) SyncEffectiveFieldsDuringRead(existingState BudgetConfigurationFilterClause) {
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
	Key types.String `tfsdk:"key" tf:"optional"`

	Value types.List `tfsdk:"value" tf:"optional,object"`
}

func (newState *BudgetConfigurationFilterTagClause) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfigurationFilterTagClause) {
}

func (newState *BudgetConfigurationFilterTagClause) SyncEffectiveFieldsDuringRead(existingState BudgetConfigurationFilterTagClause) {
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
			"key": types.StringType,
			"value": basetypes.ListType{
				ElemType: BudgetConfigurationFilterClause{}.Type(ctx),
			},
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
	var v []BudgetConfigurationFilterClause
	d := o.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in BudgetConfigurationFilterTagClause.
func (o *BudgetConfigurationFilterTagClause) SetValue(ctx context.Context, v BudgetConfigurationFilterClause) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	o.Value = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilterWorkspaceIdClause struct {
	Operator types.String `tfsdk:"operator" tf:"optional"`

	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *BudgetConfigurationFilterWorkspaceIdClause) SyncEffectiveFieldsDuringCreateOrUpdate(plan BudgetConfigurationFilterWorkspaceIdClause) {
}

func (newState *BudgetConfigurationFilterWorkspaceIdClause) SyncEffectiveFieldsDuringRead(existingState BudgetConfigurationFilterWorkspaceIdClause) {
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

type CreateBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType types.String `tfsdk:"dashboard_type" tf:"optional"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

func (newState *CreateBillingUsageDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBillingUsageDashboardRequest) {
}

func (newState *CreateBillingUsageDashboardRequest) SyncEffectiveFieldsDuringRead(existingState CreateBillingUsageDashboardRequest) {
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
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
}

func (newState *CreateBillingUsageDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBillingUsageDashboardResponse) {
}

func (newState *CreateBillingUsageDashboardResponse) SyncEffectiveFieldsDuringRead(existingState CreateBillingUsageDashboardResponse) {
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

func (newState *CreateBudgetConfigurationBudget) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationBudget) {
}

func (newState *CreateBudgetConfigurationBudget) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationBudget) {
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
			"filter": basetypes.ListType{
				ElemType: BudgetConfigurationFilter{}.Type(ctx),
			},
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
	var v []BudgetConfigurationFilter
	d := o.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in CreateBudgetConfigurationBudget.
func (o *CreateBudgetConfigurationBudget) SetFilter(ctx context.Context, v BudgetConfigurationFilter) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	o.Filter = types.ListValueMust(t, vs)
}

type CreateBudgetConfigurationBudgetActionConfigurations struct {
	// The type of the action.
	ActionType types.String `tfsdk:"action_type" tf:"optional"`
	// Target for the action. For example, an email address.
	Target types.String `tfsdk:"target" tf:"optional"`
}

func (newState *CreateBudgetConfigurationBudgetActionConfigurations) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationBudgetActionConfigurations) {
}

func (newState *CreateBudgetConfigurationBudgetActionConfigurations) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationBudgetActionConfigurations) {
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

func (newState *CreateBudgetConfigurationBudgetAlertConfigurations) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationBudgetAlertConfigurations) {
}

func (newState *CreateBudgetConfigurationBudgetAlertConfigurations) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationBudgetAlertConfigurations) {
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
	Budget types.List `tfsdk:"budget" tf:"object"`
}

func (newState *CreateBudgetConfigurationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationRequest) {
}

func (newState *CreateBudgetConfigurationRequest) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationRequest) {
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
			"budget": basetypes.ListType{
				ElemType: CreateBudgetConfigurationBudget{}.Type(ctx),
			},
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
	var v []CreateBudgetConfigurationBudget
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in CreateBudgetConfigurationRequest.
func (o *CreateBudgetConfigurationRequest) SetBudget(ctx context.Context, v CreateBudgetConfigurationBudget) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

type CreateBudgetConfigurationResponse struct {
	// The created budget configuration.
	Budget types.List `tfsdk:"budget" tf:"optional,object"`
}

func (newState *CreateBudgetConfigurationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateBudgetConfigurationResponse) {
}

func (newState *CreateBudgetConfigurationResponse) SyncEffectiveFieldsDuringRead(existingState CreateBudgetConfigurationResponse) {
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
			"budget": basetypes.ListType{
				ElemType: BudgetConfiguration{}.Type(ctx),
			},
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
	var v []BudgetConfiguration
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in CreateBudgetConfigurationResponse.
func (o *CreateBudgetConfigurationResponse) SetBudget(ctx context.Context, v BudgetConfiguration) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

type CreateLogDeliveryConfigurationParams struct {
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

func (newState *CreateLogDeliveryConfigurationParams) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateLogDeliveryConfigurationParams) {
}

func (newState *CreateLogDeliveryConfigurationParams) SyncEffectiveFieldsDuringRead(existingState CreateLogDeliveryConfigurationParams) {
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

// Delete budget
type DeleteBudgetConfigurationRequest struct {
	// The Databricks budget configuration ID.
	BudgetId types.String `tfsdk:"-"`
}

func (newState *DeleteBudgetConfigurationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteBudgetConfigurationRequest) {
}

func (newState *DeleteBudgetConfigurationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteBudgetConfigurationRequest) {
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

func (newState *DeleteBudgetConfigurationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteBudgetConfigurationResponse) {
}

func (newState *DeleteBudgetConfigurationResponse) SyncEffectiveFieldsDuringRead(existingState DeleteBudgetConfigurationResponse) {
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

// Return billable usage logs
type DownloadRequest struct {
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

func (newState *DownloadRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DownloadRequest) {
}

func (newState *DownloadRequest) SyncEffectiveFieldsDuringRead(existingState DownloadRequest) {
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

func (newState *DownloadResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DownloadResponse) {
}

func (newState *DownloadResponse) SyncEffectiveFieldsDuringRead(existingState DownloadResponse) {
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

// Get usage dashboard
type GetBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType types.String `tfsdk:"-"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *GetBillingUsageDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetBillingUsageDashboardRequest) {
}

func (newState *GetBillingUsageDashboardRequest) SyncEffectiveFieldsDuringRead(existingState GetBillingUsageDashboardRequest) {
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
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The URL of the usage dashboard.
	DashboardUrl types.String `tfsdk:"dashboard_url" tf:"optional"`
}

func (newState *GetBillingUsageDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetBillingUsageDashboardResponse) {
}

func (newState *GetBillingUsageDashboardResponse) SyncEffectiveFieldsDuringRead(existingState GetBillingUsageDashboardResponse) {
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

// Get budget
type GetBudgetConfigurationRequest struct {
	// The budget configuration ID
	BudgetId types.String `tfsdk:"-"`
}

func (newState *GetBudgetConfigurationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetBudgetConfigurationRequest) {
}

func (newState *GetBudgetConfigurationRequest) SyncEffectiveFieldsDuringRead(existingState GetBudgetConfigurationRequest) {
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
	Budget types.List `tfsdk:"budget" tf:"optional,object"`
}

func (newState *GetBudgetConfigurationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetBudgetConfigurationResponse) {
}

func (newState *GetBudgetConfigurationResponse) SyncEffectiveFieldsDuringRead(existingState GetBudgetConfigurationResponse) {
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
			"budget": basetypes.ListType{
				ElemType: BudgetConfiguration{}.Type(ctx),
			},
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
	var v []BudgetConfiguration
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in GetBudgetConfigurationResponse.
func (o *GetBudgetConfigurationResponse) SetBudget(ctx context.Context, v BudgetConfiguration) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

// Get log delivery configuration
type GetLogDeliveryRequest struct {
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId types.String `tfsdk:"-"`
}

func (newState *GetLogDeliveryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetLogDeliveryRequest) {
}

func (newState *GetLogDeliveryRequest) SyncEffectiveFieldsDuringRead(existingState GetLogDeliveryRequest) {
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

// Get all budgets
type ListBudgetConfigurationsRequest struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListBudgetConfigurationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListBudgetConfigurationsRequest) {
}

func (newState *ListBudgetConfigurationsRequest) SyncEffectiveFieldsDuringRead(existingState ListBudgetConfigurationsRequest) {
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
	Budgets types.List `tfsdk:"budgets" tf:"optional"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListBudgetConfigurationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListBudgetConfigurationsResponse) {
}

func (newState *ListBudgetConfigurationsResponse) SyncEffectiveFieldsDuringRead(existingState ListBudgetConfigurationsResponse) {
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

// Get all log delivery configurations
type ListLogDeliveryRequest struct {
	// Filter by credential configuration ID.
	CredentialsId types.String `tfsdk:"-"`
	// Filter by status `ENABLED` or `DISABLED`.
	Status types.String `tfsdk:"-"`
	// Filter by storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (newState *ListLogDeliveryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListLogDeliveryRequest) {
}

func (newState *ListLogDeliveryRequest) SyncEffectiveFieldsDuringRead(existingState ListLogDeliveryRequest) {
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
			"status":                   o.Status,
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListLogDeliveryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id":           types.StringType,
			"status":                   types.StringType,
			"storage_configuration_id": types.StringType,
		},
	}
}

type LogDeliveryConfiguration struct {
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

func (newState *LogDeliveryConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogDeliveryConfiguration) {
}

func (newState *LogDeliveryConfiguration) SyncEffectiveFieldsDuringRead(existingState LogDeliveryConfiguration) {
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
			"account_id":           types.StringType,
			"config_id":            types.StringType,
			"config_name":          types.StringType,
			"creation_time":        types.Int64Type,
			"credentials_id":       types.StringType,
			"delivery_path_prefix": types.StringType,
			"delivery_start_time":  types.StringType,
			"log_delivery_status": basetypes.ListType{
				ElemType: LogDeliveryStatus{}.Type(ctx),
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

// GetLogDeliveryStatus returns the value of the LogDeliveryStatus field in LogDeliveryConfiguration as
// a LogDeliveryStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *LogDeliveryConfiguration) GetLogDeliveryStatus(ctx context.Context) (LogDeliveryStatus, bool) {
	var e LogDeliveryStatus
	if o.LogDeliveryStatus.IsNull() || o.LogDeliveryStatus.IsUnknown() {
		return e, false
	}
	var v []LogDeliveryStatus
	d := o.LogDeliveryStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryStatus sets the value of the LogDeliveryStatus field in LogDeliveryConfiguration.
func (o *LogDeliveryConfiguration) SetLogDeliveryStatus(ctx context.Context, v LogDeliveryStatus) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_status"]
	o.LogDeliveryStatus = types.ListValueMust(t, vs)
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

// Databricks log delivery status.
type LogDeliveryStatus struct {
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

func (newState *LogDeliveryStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogDeliveryStatus) {
}

func (newState *LogDeliveryStatus) SyncEffectiveFieldsDuringRead(existingState LogDeliveryStatus) {
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

func (newState *PatchStatusResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PatchStatusResponse) {
}

func (newState *PatchStatusResponse) SyncEffectiveFieldsDuringRead(existingState PatchStatusResponse) {
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

type UpdateBudgetConfigurationBudget struct {
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

func (newState *UpdateBudgetConfigurationBudget) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateBudgetConfigurationBudget) {
}

func (newState *UpdateBudgetConfigurationBudget) SyncEffectiveFieldsDuringRead(existingState UpdateBudgetConfigurationBudget) {
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
			"filter": basetypes.ListType{
				ElemType: BudgetConfigurationFilter{}.Type(ctx),
			},
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
	var v []BudgetConfigurationFilter
	d := o.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in UpdateBudgetConfigurationBudget.
func (o *UpdateBudgetConfigurationBudget) SetFilter(ctx context.Context, v BudgetConfigurationFilter) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	o.Filter = types.ListValueMust(t, vs)
}

type UpdateBudgetConfigurationRequest struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID.
	Budget types.List `tfsdk:"budget" tf:"object"`
	// The Databricks budget configuration ID.
	BudgetId types.String `tfsdk:"-"`
}

func (newState *UpdateBudgetConfigurationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateBudgetConfigurationRequest) {
}

func (newState *UpdateBudgetConfigurationRequest) SyncEffectiveFieldsDuringRead(existingState UpdateBudgetConfigurationRequest) {
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
			"budget": basetypes.ListType{
				ElemType: UpdateBudgetConfigurationBudget{}.Type(ctx),
			},
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
	var v []UpdateBudgetConfigurationBudget
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in UpdateBudgetConfigurationRequest.
func (o *UpdateBudgetConfigurationRequest) SetBudget(ctx context.Context, v UpdateBudgetConfigurationBudget) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

type UpdateBudgetConfigurationResponse struct {
	// The updated budget.
	Budget types.List `tfsdk:"budget" tf:"optional,object"`
}

func (newState *UpdateBudgetConfigurationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateBudgetConfigurationResponse) {
}

func (newState *UpdateBudgetConfigurationResponse) SyncEffectiveFieldsDuringRead(existingState UpdateBudgetConfigurationResponse) {
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
			"budget": basetypes.ListType{
				ElemType: BudgetConfiguration{}.Type(ctx),
			},
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
	var v []BudgetConfiguration
	d := o.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in UpdateBudgetConfigurationResponse.
func (o *UpdateBudgetConfigurationResponse) SetBudget(ctx context.Context, v BudgetConfiguration) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	o.Budget = types.ListValueMust(t, vs)
}

type UpdateLogDeliveryConfigurationStatusRequest struct {
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId types.String `tfsdk:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status types.String `tfsdk:"status" tf:""`
}

func (newState *UpdateLogDeliveryConfigurationStatusRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateLogDeliveryConfigurationStatusRequest) {
}

func (newState *UpdateLogDeliveryConfigurationStatusRequest) SyncEffectiveFieldsDuringRead(existingState UpdateLogDeliveryConfigurationStatusRequest) {
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

type WrappedCreateLogDeliveryConfiguration struct {
	LogDeliveryConfiguration types.List `tfsdk:"log_delivery_configuration" tf:"optional,object"`
}

func (newState *WrappedCreateLogDeliveryConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan WrappedCreateLogDeliveryConfiguration) {
}

func (newState *WrappedCreateLogDeliveryConfiguration) SyncEffectiveFieldsDuringRead(existingState WrappedCreateLogDeliveryConfiguration) {
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
			"log_delivery_configuration": basetypes.ListType{
				ElemType: CreateLogDeliveryConfigurationParams{}.Type(ctx),
			},
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
	var v []CreateLogDeliveryConfigurationParams
	d := o.LogDeliveryConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in WrappedCreateLogDeliveryConfiguration.
func (o *WrappedCreateLogDeliveryConfiguration) SetLogDeliveryConfiguration(ctx context.Context, v CreateLogDeliveryConfigurationParams) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configuration"]
	o.LogDeliveryConfiguration = types.ListValueMust(t, vs)
}

type WrappedLogDeliveryConfiguration struct {
	LogDeliveryConfiguration types.List `tfsdk:"log_delivery_configuration" tf:"optional,object"`
}

func (newState *WrappedLogDeliveryConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan WrappedLogDeliveryConfiguration) {
}

func (newState *WrappedLogDeliveryConfiguration) SyncEffectiveFieldsDuringRead(existingState WrappedLogDeliveryConfiguration) {
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
			"log_delivery_configuration": basetypes.ListType{
				ElemType: LogDeliveryConfiguration{}.Type(ctx),
			},
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
	var v []LogDeliveryConfiguration
	d := o.LogDeliveryConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in WrappedLogDeliveryConfiguration.
func (o *WrappedLogDeliveryConfiguration) SetLogDeliveryConfiguration(ctx context.Context, v LogDeliveryConfiguration) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configuration"]
	o.LogDeliveryConfiguration = types.ListValueMust(t, vs)
}

type WrappedLogDeliveryConfigurations struct {
	LogDeliveryConfigurations types.List `tfsdk:"log_delivery_configurations" tf:"optional"`
}

func (newState *WrappedLogDeliveryConfigurations) SyncEffectiveFieldsDuringCreateOrUpdate(plan WrappedLogDeliveryConfigurations) {
}

func (newState *WrappedLogDeliveryConfigurations) SyncEffectiveFieldsDuringRead(existingState WrappedLogDeliveryConfigurations) {
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
		})
}

// Type implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfigurations) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configurations": basetypes.ListType{
				ElemType: LogDeliveryConfiguration{}.Type(ctx),
			},
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
