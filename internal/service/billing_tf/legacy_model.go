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

	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf" // .tmpl
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ActionConfiguration_SdkV2 struct {
	// Databricks action configuration ID.
	ActionConfigurationId types.String `tfsdk:"action_configuration_id"`
	// The type of the action.
	ActionType types.String `tfsdk:"action_type"`
	// Target for the action. For example, an email address.
	Target types.String `tfsdk:"target"`
}

func (to *ActionConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ActionConfiguration_SdkV2) {
}

func (to *ActionConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ActionConfiguration_SdkV2) {
}

func (m ActionConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ActionConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ActionConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (m ActionConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_configuration_id": m.ActionConfigurationId,
			"action_type":             m.ActionType,
			"target":                  m.Target,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ActionConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *AlertConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertConfiguration_SdkV2) {
	if !from.ActionConfigurations.IsNull() && !from.ActionConfigurations.IsUnknown() && to.ActionConfigurations.IsNull() && len(from.ActionConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ActionConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ActionConfigurations = from.ActionConfigurations
	}
}

func (to *AlertConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertConfiguration_SdkV2) {
	if !from.ActionConfigurations.IsNull() && !from.ActionConfigurations.IsUnknown() && to.ActionConfigurations.IsNull() && len(from.ActionConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ActionConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ActionConfigurations = from.ActionConfigurations
	}
}

func (m AlertConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"action_configurations": reflect.TypeOf(ActionConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_configurations":  m.ActionConfigurations,
			"alert_configuration_id": m.AlertConfigurationId,
			"quantity_threshold":     m.QuantityThreshold,
			"quantity_type":          m.QuantityType,
			"time_period":            m.TimePeriod,
			"trigger_type":           m.TriggerType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *AlertConfiguration_SdkV2) GetActionConfigurations(ctx context.Context) ([]ActionConfiguration_SdkV2, bool) {
	if m.ActionConfigurations.IsNull() || m.ActionConfigurations.IsUnknown() {
		return nil, false
	}
	var v []ActionConfiguration_SdkV2
	d := m.ActionConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActionConfigurations sets the value of the ActionConfigurations field in AlertConfiguration_SdkV2.
func (m *AlertConfiguration_SdkV2) SetActionConfigurations(ctx context.Context, v []ActionConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["action_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ActionConfigurations = types.ListValueMust(t, vs)
}

type BudgetConfiguration_SdkV2 struct {
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
	Filter types.List `tfsdk:"filter"`
	// Update time of this budget configuration.
	UpdateTime types.Int64 `tfsdk:"update_time"`
}

func (to *BudgetConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BudgetConfiguration_SdkV2) {
	if !from.AlertConfigurations.IsNull() && !from.AlertConfigurations.IsUnknown() && to.AlertConfigurations.IsNull() && len(from.AlertConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AlertConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AlertConfigurations = from.AlertConfigurations
	}
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				// Recursively sync the fields of Filter
				toFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (to *BudgetConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BudgetConfiguration_SdkV2) {
	if !from.AlertConfigurations.IsNull() && !from.AlertConfigurations.IsUnknown() && to.AlertConfigurations.IsNull() && len(from.AlertConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AlertConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AlertConfigurations = from.AlertConfigurations
	}
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				toFilter.SyncFieldsDuringRead(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (m BudgetConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["alert_configurations"] = attrs["alert_configurations"].SetOptional()
	attrs["budget_configuration_id"] = attrs["budget_configuration_id"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["filter"] = attrs["filter"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m BudgetConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_configurations": reflect.TypeOf(AlertConfiguration_SdkV2{}),
		"filter":               reflect.TypeOf(BudgetConfigurationFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (m BudgetConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":              m.AccountId,
			"alert_configurations":    m.AlertConfigurations,
			"budget_configuration_id": m.BudgetConfigurationId,
			"create_time":             m.CreateTime,
			"display_name":            m.DisplayName,
			"filter":                  m.Filter,
			"update_time":             m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BudgetConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *BudgetConfiguration_SdkV2) GetAlertConfigurations(ctx context.Context) ([]AlertConfiguration_SdkV2, bool) {
	if m.AlertConfigurations.IsNull() || m.AlertConfigurations.IsUnknown() {
		return nil, false
	}
	var v []AlertConfiguration_SdkV2
	d := m.AlertConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlertConfigurations sets the value of the AlertConfigurations field in BudgetConfiguration_SdkV2.
func (m *BudgetConfiguration_SdkV2) SetAlertConfigurations(ctx context.Context, v []AlertConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AlertConfigurations = types.ListValueMust(t, vs)
}

// GetFilter returns the value of the Filter field in BudgetConfiguration_SdkV2 as
// a BudgetConfigurationFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *BudgetConfiguration_SdkV2) GetFilter(ctx context.Context) (BudgetConfigurationFilter_SdkV2, bool) {
	var e BudgetConfigurationFilter_SdkV2
	if m.Filter.IsNull() || m.Filter.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilter_SdkV2
	d := m.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in BudgetConfiguration_SdkV2.
func (m *BudgetConfiguration_SdkV2) SetFilter(ctx context.Context, v BudgetConfigurationFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	m.Filter = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilter_SdkV2 struct {
	// A list of tag keys and values that will limit the budget to usage that
	// includes those specific custom tags. Tags are case-sensitive and should
	// be entered exactly as they appear in your usage data.
	Tags types.List `tfsdk:"tags"`
	// If provided, usage must match with the provided Databricks workspace IDs.
	WorkspaceId types.List `tfsdk:"workspace_id"`
}

func (to *BudgetConfigurationFilter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BudgetConfigurationFilter_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
	if !from.WorkspaceId.IsNull() && !from.WorkspaceId.IsUnknown() {
		if toWorkspaceId, ok := to.GetWorkspaceId(ctx); ok {
			if fromWorkspaceId, ok := from.GetWorkspaceId(ctx); ok {
				// Recursively sync the fields of WorkspaceId
				toWorkspaceId.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceId)
				to.SetWorkspaceId(ctx, toWorkspaceId)
			}
		}
	}
}

func (to *BudgetConfigurationFilter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BudgetConfigurationFilter_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
	if !from.WorkspaceId.IsNull() && !from.WorkspaceId.IsUnknown() {
		if toWorkspaceId, ok := to.GetWorkspaceId(ctx); ok {
			if fromWorkspaceId, ok := from.GetWorkspaceId(ctx); ok {
				toWorkspaceId.SyncFieldsDuringRead(ctx, fromWorkspaceId)
				to.SetWorkspaceId(ctx, toWorkspaceId)
			}
		}
	}
}

func (m BudgetConfigurationFilter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BudgetConfigurationFilter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags":         reflect.TypeOf(BudgetConfigurationFilterTagClause_SdkV2{}),
		"workspace_id": reflect.TypeOf(BudgetConfigurationFilterWorkspaceIdClause_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilter_SdkV2
// only implements ToObjectValue() and Type().
func (m BudgetConfigurationFilter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tags":         m.Tags,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BudgetConfigurationFilter_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *BudgetConfigurationFilter_SdkV2) GetTags(ctx context.Context) ([]BudgetConfigurationFilterTagClause_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []BudgetConfigurationFilterTagClause_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in BudgetConfigurationFilter_SdkV2.
func (m *BudgetConfigurationFilter_SdkV2) SetTags(ctx context.Context, v []BudgetConfigurationFilterTagClause_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// GetWorkspaceId returns the value of the WorkspaceId field in BudgetConfigurationFilter_SdkV2 as
// a BudgetConfigurationFilterWorkspaceIdClause_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *BudgetConfigurationFilter_SdkV2) GetWorkspaceId(ctx context.Context) (BudgetConfigurationFilterWorkspaceIdClause_SdkV2, bool) {
	var e BudgetConfigurationFilterWorkspaceIdClause_SdkV2
	if m.WorkspaceId.IsNull() || m.WorkspaceId.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilterWorkspaceIdClause_SdkV2
	d := m.WorkspaceId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceId sets the value of the WorkspaceId field in BudgetConfigurationFilter_SdkV2.
func (m *BudgetConfigurationFilter_SdkV2) SetWorkspaceId(ctx context.Context, v BudgetConfigurationFilterWorkspaceIdClause_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_id"]
	m.WorkspaceId = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilterClause_SdkV2 struct {
	Operator types.String `tfsdk:"operator"`

	Values types.List `tfsdk:"values"`
}

func (to *BudgetConfigurationFilterClause_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BudgetConfigurationFilterClause_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *BudgetConfigurationFilterClause_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BudgetConfigurationFilterClause_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m BudgetConfigurationFilterClause_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BudgetConfigurationFilterClause_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilterClause_SdkV2
// only implements ToObjectValue() and Type().
func (m BudgetConfigurationFilterClause_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operator": m.Operator,
			"values":   m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BudgetConfigurationFilterClause_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *BudgetConfigurationFilterClause_SdkV2) GetValues(ctx context.Context) ([]types.String, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in BudgetConfigurationFilterClause_SdkV2.
func (m *BudgetConfigurationFilterClause_SdkV2) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilterTagClause_SdkV2 struct {
	Key types.String `tfsdk:"key"`

	Value types.List `tfsdk:"value"`
}

func (to *BudgetConfigurationFilterTagClause_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BudgetConfigurationFilterTagClause_SdkV2) {
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				// Recursively sync the fields of Value
				toValue.SyncFieldsDuringCreateOrUpdate(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (to *BudgetConfigurationFilterTagClause_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BudgetConfigurationFilterTagClause_SdkV2) {
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				toValue.SyncFieldsDuringRead(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (m BudgetConfigurationFilterTagClause_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()
	attrs["value"] = attrs["value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BudgetConfigurationFilterTagClause.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BudgetConfigurationFilterTagClause_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(BudgetConfigurationFilterClause_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilterTagClause_SdkV2
// only implements ToObjectValue() and Type().
func (m BudgetConfigurationFilterTagClause_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BudgetConfigurationFilterTagClause_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *BudgetConfigurationFilterTagClause_SdkV2) GetValue(ctx context.Context) (BudgetConfigurationFilterClause_SdkV2, bool) {
	var e BudgetConfigurationFilterClause_SdkV2
	if m.Value.IsNull() || m.Value.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilterClause_SdkV2
	d := m.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in BudgetConfigurationFilterTagClause_SdkV2.
func (m *BudgetConfigurationFilterTagClause_SdkV2) SetValue(ctx context.Context, v BudgetConfigurationFilterClause_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	m.Value = types.ListValueMust(t, vs)
}

type BudgetConfigurationFilterWorkspaceIdClause_SdkV2 struct {
	Operator types.String `tfsdk:"operator"`

	Values types.List `tfsdk:"values"`
}

func (to *BudgetConfigurationFilterWorkspaceIdClause_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BudgetConfigurationFilterWorkspaceIdClause_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *BudgetConfigurationFilterWorkspaceIdClause_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BudgetConfigurationFilterWorkspaceIdClause_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m BudgetConfigurationFilterWorkspaceIdClause_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BudgetConfigurationFilterWorkspaceIdClause_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetConfigurationFilterWorkspaceIdClause_SdkV2
// only implements ToObjectValue() and Type().
func (m BudgetConfigurationFilterWorkspaceIdClause_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operator": m.Operator,
			"values":   m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BudgetConfigurationFilterWorkspaceIdClause_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *BudgetConfigurationFilterWorkspaceIdClause_SdkV2) GetValues(ctx context.Context) ([]types.Int64, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in BudgetConfigurationFilterWorkspaceIdClause_SdkV2.
func (m *BudgetConfigurationFilterWorkspaceIdClause_SdkV2) SetValues(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

// Contains the BudgetPolicy details.
type BudgetPolicy_SdkV2 struct {
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

func (to *BudgetPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BudgetPolicy_SdkV2) {
	if !from.BindingWorkspaceIds.IsNull() && !from.BindingWorkspaceIds.IsUnknown() && to.BindingWorkspaceIds.IsNull() && len(from.BindingWorkspaceIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for BindingWorkspaceIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.BindingWorkspaceIds = from.BindingWorkspaceIds
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (to *BudgetPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BudgetPolicy_SdkV2) {
	if !from.BindingWorkspaceIds.IsNull() && !from.BindingWorkspaceIds.IsUnknown() && to.BindingWorkspaceIds.IsNull() && len(from.BindingWorkspaceIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for BindingWorkspaceIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.BindingWorkspaceIds = from.BindingWorkspaceIds
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (m BudgetPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["binding_workspace_ids"] = attrs["binding_workspace_ids"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetComputed()
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
func (m BudgetPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"binding_workspace_ids": reflect.TypeOf(types.Int64{}),
		"custom_tags":           reflect.TypeOf(compute_tf.CustomPolicyTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BudgetPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m BudgetPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"binding_workspace_ids": m.BindingWorkspaceIds,
			"custom_tags":           m.CustomTags,
			"policy_id":             m.PolicyId,
			"policy_name":           m.PolicyName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BudgetPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"binding_workspace_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"custom_tags": basetypes.ListType{
				ElemType: compute_tf.CustomPolicyTag_SdkV2{}.Type(ctx),
			},
			"policy_id":   types.StringType,
			"policy_name": types.StringType,
		},
	}
}

// GetBindingWorkspaceIds returns the value of the BindingWorkspaceIds field in BudgetPolicy_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *BudgetPolicy_SdkV2) GetBindingWorkspaceIds(ctx context.Context) ([]types.Int64, bool) {
	if m.BindingWorkspaceIds.IsNull() || m.BindingWorkspaceIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.BindingWorkspaceIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBindingWorkspaceIds sets the value of the BindingWorkspaceIds field in BudgetPolicy_SdkV2.
func (m *BudgetPolicy_SdkV2) SetBindingWorkspaceIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["binding_workspace_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.BindingWorkspaceIds = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in BudgetPolicy_SdkV2 as
// a slice of compute_tf.CustomPolicyTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *BudgetPolicy_SdkV2) GetCustomTags(ctx context.Context) ([]compute_tf.CustomPolicyTag_SdkV2, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.CustomPolicyTag_SdkV2
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in BudgetPolicy_SdkV2.
func (m *BudgetPolicy_SdkV2) SetCustomTags(ctx context.Context, v []compute_tf.CustomPolicyTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

type CreateBillingUsageDashboardRequest_SdkV2 struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType types.String `tfsdk:"dashboard_type"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *CreateBillingUsageDashboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBillingUsageDashboardRequest_SdkV2) {
}

func (to *CreateBillingUsageDashboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBillingUsageDashboardRequest_SdkV2) {
}

func (m CreateBillingUsageDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateBillingUsageDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBillingUsageDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBillingUsageDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_type": m.DashboardType,
			"workspace_id":   m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBillingUsageDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

type CreateBillingUsageDashboardResponse_SdkV2 struct {
	// The unique id of the usage dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`
}

func (to *CreateBillingUsageDashboardResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBillingUsageDashboardResponse_SdkV2) {
}

func (to *CreateBillingUsageDashboardResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBillingUsageDashboardResponse_SdkV2) {
}

func (m CreateBillingUsageDashboardResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateBillingUsageDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBillingUsageDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBillingUsageDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": m.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBillingUsageDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type CreateBudgetConfigurationBudget_SdkV2 struct {
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
	Filter types.List `tfsdk:"filter"`
}

func (to *CreateBudgetConfigurationBudget_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBudgetConfigurationBudget_SdkV2) {
	if !from.AlertConfigurations.IsNull() && !from.AlertConfigurations.IsUnknown() && to.AlertConfigurations.IsNull() && len(from.AlertConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AlertConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AlertConfigurations = from.AlertConfigurations
	}
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				// Recursively sync the fields of Filter
				toFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (to *CreateBudgetConfigurationBudget_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBudgetConfigurationBudget_SdkV2) {
	if !from.AlertConfigurations.IsNull() && !from.AlertConfigurations.IsUnknown() && to.AlertConfigurations.IsNull() && len(from.AlertConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AlertConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AlertConfigurations = from.AlertConfigurations
	}
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				toFilter.SyncFieldsDuringRead(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (m CreateBudgetConfigurationBudget_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["alert_configurations"] = attrs["alert_configurations"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["filter"] = attrs["filter"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationBudget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateBudgetConfigurationBudget_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_configurations": reflect.TypeOf(CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2{}),
		"filter":               reflect.TypeOf(BudgetConfigurationFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationBudget_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBudgetConfigurationBudget_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":           m.AccountId,
			"alert_configurations": m.AlertConfigurations,
			"display_name":         m.DisplayName,
			"filter":               m.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBudgetConfigurationBudget_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateBudgetConfigurationBudget_SdkV2) GetAlertConfigurations(ctx context.Context) ([]CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2, bool) {
	if m.AlertConfigurations.IsNull() || m.AlertConfigurations.IsUnknown() {
		return nil, false
	}
	var v []CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2
	d := m.AlertConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlertConfigurations sets the value of the AlertConfigurations field in CreateBudgetConfigurationBudget_SdkV2.
func (m *CreateBudgetConfigurationBudget_SdkV2) SetAlertConfigurations(ctx context.Context, v []CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AlertConfigurations = types.ListValueMust(t, vs)
}

// GetFilter returns the value of the Filter field in CreateBudgetConfigurationBudget_SdkV2 as
// a BudgetConfigurationFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateBudgetConfigurationBudget_SdkV2) GetFilter(ctx context.Context) (BudgetConfigurationFilter_SdkV2, bool) {
	var e BudgetConfigurationFilter_SdkV2
	if m.Filter.IsNull() || m.Filter.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilter_SdkV2
	d := m.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in CreateBudgetConfigurationBudget_SdkV2.
func (m *CreateBudgetConfigurationBudget_SdkV2) SetFilter(ctx context.Context, v BudgetConfigurationFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	m.Filter = types.ListValueMust(t, vs)
}

type CreateBudgetConfigurationBudgetActionConfigurations_SdkV2 struct {
	// The type of the action.
	ActionType types.String `tfsdk:"action_type"`
	// Target for the action. For example, an email address.
	Target types.String `tfsdk:"target"`
}

func (to *CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) {
}

func (to *CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) {
}

func (m CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationBudgetActionConfigurations_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_type": m.ActionType,
			"target":      m.Target,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) {
	if !from.ActionConfigurations.IsNull() && !from.ActionConfigurations.IsUnknown() && to.ActionConfigurations.IsNull() && len(from.ActionConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ActionConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ActionConfigurations = from.ActionConfigurations
	}
}

func (to *CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) {
	if !from.ActionConfigurations.IsNull() && !from.ActionConfigurations.IsUnknown() && to.ActionConfigurations.IsNull() && len(from.ActionConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ActionConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ActionConfigurations = from.ActionConfigurations
	}
}

func (m CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"action_configurations": reflect.TypeOf(CreateBudgetConfigurationBudgetActionConfigurations_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_configurations": m.ActionConfigurations,
			"quantity_threshold":    m.QuantityThreshold,
			"quantity_type":         m.QuantityType,
			"time_period":           m.TimePeriod,
			"trigger_type":          m.TriggerType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) GetActionConfigurations(ctx context.Context) ([]CreateBudgetConfigurationBudgetActionConfigurations_SdkV2, bool) {
	if m.ActionConfigurations.IsNull() || m.ActionConfigurations.IsUnknown() {
		return nil, false
	}
	var v []CreateBudgetConfigurationBudgetActionConfigurations_SdkV2
	d := m.ActionConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActionConfigurations sets the value of the ActionConfigurations field in CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2.
func (m *CreateBudgetConfigurationBudgetAlertConfigurations_SdkV2) SetActionConfigurations(ctx context.Context, v []CreateBudgetConfigurationBudgetActionConfigurations_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["action_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ActionConfigurations = types.ListValueMust(t, vs)
}

type CreateBudgetConfigurationRequest_SdkV2 struct {
	// Properties of the new budget configuration.
	Budget types.List `tfsdk:"budget"`
}

func (to *CreateBudgetConfigurationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBudgetConfigurationRequest_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				// Recursively sync the fields of Budget
				toBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (to *CreateBudgetConfigurationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBudgetConfigurationRequest_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				toBudget.SyncFieldsDuringRead(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (m CreateBudgetConfigurationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetRequired()
	attrs["budget"] = attrs["budget"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateBudgetConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(CreateBudgetConfigurationBudget_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBudgetConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": m.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBudgetConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateBudgetConfigurationRequest_SdkV2) GetBudget(ctx context.Context) (CreateBudgetConfigurationBudget_SdkV2, bool) {
	var e CreateBudgetConfigurationBudget_SdkV2
	if m.Budget.IsNull() || m.Budget.IsUnknown() {
		return e, false
	}
	var v []CreateBudgetConfigurationBudget_SdkV2
	d := m.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in CreateBudgetConfigurationRequest_SdkV2.
func (m *CreateBudgetConfigurationRequest_SdkV2) SetBudget(ctx context.Context, v CreateBudgetConfigurationBudget_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	m.Budget = types.ListValueMust(t, vs)
}

type CreateBudgetConfigurationResponse_SdkV2 struct {
	// The created budget configuration.
	Budget types.List `tfsdk:"budget"`
}

func (to *CreateBudgetConfigurationResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBudgetConfigurationResponse_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				// Recursively sync the fields of Budget
				toBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (to *CreateBudgetConfigurationResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBudgetConfigurationResponse_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				toBudget.SyncFieldsDuringRead(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (m CreateBudgetConfigurationResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetOptional()
	attrs["budget"] = attrs["budget"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateBudgetConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(BudgetConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBudgetConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": m.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBudgetConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateBudgetConfigurationResponse_SdkV2) GetBudget(ctx context.Context) (BudgetConfiguration_SdkV2, bool) {
	var e BudgetConfiguration_SdkV2
	if m.Budget.IsNull() || m.Budget.IsUnknown() {
		return e, false
	}
	var v []BudgetConfiguration_SdkV2
	d := m.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in CreateBudgetConfigurationResponse_SdkV2.
func (m *CreateBudgetConfigurationResponse_SdkV2) SetBudget(ctx context.Context, v BudgetConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	m.Budget = types.ListValueMust(t, vs)
}

// A request to create a BudgetPolicy.
type CreateBudgetPolicyRequest_SdkV2 struct {
	// The policy to create. `policy_id` needs to be empty as it will be
	// generated `policy_name` must be provided, custom_tags may need to be
	// provided depending on the cloud provider. All other fields are optional.
	Policy types.List `tfsdk:"policy"`
	// A unique identifier for this request. Restricted to 36 ASCII characters.
	// A random UUID is recommended. This request is only idempotent if a
	// `request_id` is provided.
	RequestId types.String `tfsdk:"request_id"`
}

func (to *CreateBudgetPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBudgetPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				// Recursively sync the fields of Policy
				toPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (to *CreateBudgetPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBudgetPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m CreateBudgetPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetOptional()
	attrs["policy"] = attrs["policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateBudgetPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(BudgetPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBudgetPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBudgetPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":     m.Policy,
			"request_id": m.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBudgetPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy": basetypes.ListType{
				ElemType: BudgetPolicy_SdkV2{}.Type(ctx),
			},
			"request_id": types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in CreateBudgetPolicyRequest_SdkV2 as
// a BudgetPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateBudgetPolicyRequest_SdkV2) GetPolicy(ctx context.Context) (BudgetPolicy_SdkV2, bool) {
	var e BudgetPolicy_SdkV2
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v []BudgetPolicy_SdkV2
	d := m.Policy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPolicy sets the value of the Policy field in CreateBudgetPolicyRequest_SdkV2.
func (m *CreateBudgetPolicyRequest_SdkV2) SetPolicy(ctx context.Context, v BudgetPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy"]
	m.Policy = types.ListValueMust(t, vs)
}

// * Log Delivery Configuration
type CreateLogDeliveryConfigurationParams_SdkV2 struct {
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

func (to *CreateLogDeliveryConfigurationParams_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateLogDeliveryConfigurationParams_SdkV2) {
	if !from.WorkspaceIdsFilter.IsNull() && !from.WorkspaceIdsFilter.IsUnknown() && to.WorkspaceIdsFilter.IsNull() && len(from.WorkspaceIdsFilter.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceIdsFilter, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceIdsFilter = from.WorkspaceIdsFilter
	}
}

func (to *CreateLogDeliveryConfigurationParams_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateLogDeliveryConfigurationParams_SdkV2) {
	if !from.WorkspaceIdsFilter.IsNull() && !from.WorkspaceIdsFilter.IsUnknown() && to.WorkspaceIdsFilter.IsNull() && len(from.WorkspaceIdsFilter.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceIdsFilter, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceIdsFilter = from.WorkspaceIdsFilter
	}
}

func (m CreateLogDeliveryConfigurationParams_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateLogDeliveryConfigurationParams_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_ids_filter": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLogDeliveryConfigurationParams_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateLogDeliveryConfigurationParams_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config_name":              m.ConfigName,
			"credentials_id":           m.CredentialsId,
			"delivery_path_prefix":     m.DeliveryPathPrefix,
			"delivery_start_time":      m.DeliveryStartTime,
			"log_type":                 m.LogType,
			"output_format":            m.OutputFormat,
			"status":                   m.Status,
			"storage_configuration_id": m.StorageConfigurationId,
			"workspace_ids_filter":     m.WorkspaceIdsFilter,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateLogDeliveryConfigurationParams_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateLogDeliveryConfigurationParams_SdkV2) GetWorkspaceIdsFilter(ctx context.Context) ([]types.Int64, bool) {
	if m.WorkspaceIdsFilter.IsNull() || m.WorkspaceIdsFilter.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.WorkspaceIdsFilter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceIdsFilter sets the value of the WorkspaceIdsFilter field in CreateLogDeliveryConfigurationParams_SdkV2.
func (m *CreateLogDeliveryConfigurationParams_SdkV2) SetWorkspaceIdsFilter(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_ids_filter"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceIdsFilter = types.ListValueMust(t, vs)
}

type DeleteBudgetConfigurationRequest_SdkV2 struct {
	// The Databricks budget configuration ID.
	BudgetId types.String `tfsdk:"-"`
}

func (to *DeleteBudgetConfigurationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteBudgetConfigurationRequest_SdkV2) {
}

func (to *DeleteBudgetConfigurationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteBudgetConfigurationRequest_SdkV2) {
}

func (m DeleteBudgetConfigurationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteBudgetConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBudgetConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteBudgetConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_id": m.BudgetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteBudgetConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_id": types.StringType,
		},
	}
}

type DeleteBudgetConfigurationResponse_SdkV2 struct {
}

func (to *DeleteBudgetConfigurationResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteBudgetConfigurationResponse_SdkV2) {
}

func (to *DeleteBudgetConfigurationResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteBudgetConfigurationResponse_SdkV2) {
}

func (m DeleteBudgetConfigurationResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteBudgetConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBudgetConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteBudgetConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteBudgetConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteBudgetPolicyRequest_SdkV2 struct {
	// The Id of the policy.
	PolicyId types.String `tfsdk:"-"`
}

func (to *DeleteBudgetPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteBudgetPolicyRequest_SdkV2) {
}

func (to *DeleteBudgetPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteBudgetPolicyRequest_SdkV2) {
}

func (m DeleteBudgetPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteBudgetPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBudgetPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteBudgetPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteBudgetPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type DownloadRequest_SdkV2 struct {
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

func (to *DownloadRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DownloadRequest_SdkV2) {
}

func (to *DownloadRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DownloadRequest_SdkV2) {
}

func (m DownloadRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DownloadRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DownloadRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_month":     m.EndMonth,
			"personal_data": m.PersonalData,
			"start_month":   m.StartMonth,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DownloadRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DownloadResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DownloadResponse_SdkV2) {
}

func (to *DownloadResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DownloadResponse_SdkV2) {
}

func (m DownloadResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DownloadResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DownloadResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": m.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DownloadResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

// Structured representation of a filter to be applied to a list of policies.
// All specified filters will be applied in conjunction.
type Filter_SdkV2 struct {
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

func (to *Filter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Filter_SdkV2) {
}

func (to *Filter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Filter_SdkV2) {
}

func (m Filter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Filter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Filter_SdkV2
// only implements ToObjectValue() and Type().
func (m Filter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator_user_id":   m.CreatorUserId,
			"creator_user_name": m.CreatorUserName,
			"policy_name":       m.PolicyName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Filter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator_user_id":   types.Int64Type,
			"creator_user_name": types.StringType,
			"policy_name":       types.StringType,
		},
	}
}

type GetBillingUsageDashboardRequest_SdkV2 struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType types.String `tfsdk:"-"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *GetBillingUsageDashboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetBillingUsageDashboardRequest_SdkV2) {
}

func (to *GetBillingUsageDashboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetBillingUsageDashboardRequest_SdkV2) {
}

func (m GetBillingUsageDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetBillingUsageDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBillingUsageDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetBillingUsageDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_type": m.DashboardType,
			"workspace_id":   m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetBillingUsageDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

type GetBillingUsageDashboardResponse_SdkV2 struct {
	// The unique id of the usage dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The URL of the usage dashboard.
	DashboardUrl types.String `tfsdk:"dashboard_url"`
}

func (to *GetBillingUsageDashboardResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetBillingUsageDashboardResponse_SdkV2) {
}

func (to *GetBillingUsageDashboardResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetBillingUsageDashboardResponse_SdkV2) {
}

func (m GetBillingUsageDashboardResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetBillingUsageDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBillingUsageDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetBillingUsageDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":  m.DashboardId,
			"dashboard_url": m.DashboardUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetBillingUsageDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":  types.StringType,
			"dashboard_url": types.StringType,
		},
	}
}

type GetBudgetConfigurationRequest_SdkV2 struct {
	// The budget configuration ID
	BudgetId types.String `tfsdk:"-"`
}

func (to *GetBudgetConfigurationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetBudgetConfigurationRequest_SdkV2) {
}

func (to *GetBudgetConfigurationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetBudgetConfigurationRequest_SdkV2) {
}

func (m GetBudgetConfigurationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetBudgetConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBudgetConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetBudgetConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_id": m.BudgetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetBudgetConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_id": types.StringType,
		},
	}
}

type GetBudgetConfigurationResponse_SdkV2 struct {
	Budget types.List `tfsdk:"budget"`
}

func (to *GetBudgetConfigurationResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetBudgetConfigurationResponse_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				// Recursively sync the fields of Budget
				toBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (to *GetBudgetConfigurationResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetBudgetConfigurationResponse_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				toBudget.SyncFieldsDuringRead(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (m GetBudgetConfigurationResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetOptional()
	attrs["budget"] = attrs["budget"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetBudgetConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(BudgetConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBudgetConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetBudgetConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": m.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetBudgetConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *GetBudgetConfigurationResponse_SdkV2) GetBudget(ctx context.Context) (BudgetConfiguration_SdkV2, bool) {
	var e BudgetConfiguration_SdkV2
	if m.Budget.IsNull() || m.Budget.IsUnknown() {
		return e, false
	}
	var v []BudgetConfiguration_SdkV2
	d := m.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in GetBudgetConfigurationResponse_SdkV2.
func (m *GetBudgetConfigurationResponse_SdkV2) SetBudget(ctx context.Context, v BudgetConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	m.Budget = types.ListValueMust(t, vs)
}

type GetBudgetPolicyRequest_SdkV2 struct {
	// The Id of the policy.
	PolicyId types.String `tfsdk:"-"`
}

func (to *GetBudgetPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetBudgetPolicyRequest_SdkV2) {
}

func (to *GetBudgetPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetBudgetPolicyRequest_SdkV2) {
}

func (m GetBudgetPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetBudgetPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBudgetPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetBudgetPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetBudgetPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type GetLogDeliveryConfigurationResponse_SdkV2 struct {
	// The fetched log delivery configuration
	LogDeliveryConfiguration types.List `tfsdk:"log_delivery_configuration"`
}

func (to *GetLogDeliveryConfigurationResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLogDeliveryConfigurationResponse_SdkV2) {
	if !from.LogDeliveryConfiguration.IsNull() && !from.LogDeliveryConfiguration.IsUnknown() {
		if toLogDeliveryConfiguration, ok := to.GetLogDeliveryConfiguration(ctx); ok {
			if fromLogDeliveryConfiguration, ok := from.GetLogDeliveryConfiguration(ctx); ok {
				// Recursively sync the fields of LogDeliveryConfiguration
				toLogDeliveryConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromLogDeliveryConfiguration)
				to.SetLogDeliveryConfiguration(ctx, toLogDeliveryConfiguration)
			}
		}
	}
}

func (to *GetLogDeliveryConfigurationResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLogDeliveryConfigurationResponse_SdkV2) {
	if !from.LogDeliveryConfiguration.IsNull() && !from.LogDeliveryConfiguration.IsUnknown() {
		if toLogDeliveryConfiguration, ok := to.GetLogDeliveryConfiguration(ctx); ok {
			if fromLogDeliveryConfiguration, ok := from.GetLogDeliveryConfiguration(ctx); ok {
				toLogDeliveryConfiguration.SyncFieldsDuringRead(ctx, fromLogDeliveryConfiguration)
				to.SetLogDeliveryConfiguration(ctx, toLogDeliveryConfiguration)
			}
		}
	}
}

func (m GetLogDeliveryConfigurationResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["log_delivery_configuration"] = attrs["log_delivery_configuration"].SetOptional()
	attrs["log_delivery_configuration"] = attrs["log_delivery_configuration"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLogDeliveryConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLogDeliveryConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configuration": reflect.TypeOf(LogDeliveryConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLogDeliveryConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetLogDeliveryConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration": m.LogDeliveryConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLogDeliveryConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration": basetypes.ListType{
				ElemType: LogDeliveryConfiguration_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetLogDeliveryConfiguration returns the value of the LogDeliveryConfiguration field in GetLogDeliveryConfigurationResponse_SdkV2 as
// a LogDeliveryConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetLogDeliveryConfigurationResponse_SdkV2) GetLogDeliveryConfiguration(ctx context.Context) (LogDeliveryConfiguration_SdkV2, bool) {
	var e LogDeliveryConfiguration_SdkV2
	if m.LogDeliveryConfiguration.IsNull() || m.LogDeliveryConfiguration.IsUnknown() {
		return e, false
	}
	var v []LogDeliveryConfiguration_SdkV2
	d := m.LogDeliveryConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in GetLogDeliveryConfigurationResponse_SdkV2.
func (m *GetLogDeliveryConfigurationResponse_SdkV2) SetLogDeliveryConfiguration(ctx context.Context, v LogDeliveryConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configuration"]
	m.LogDeliveryConfiguration = types.ListValueMust(t, vs)
}

type GetLogDeliveryRequest_SdkV2 struct {
	// The log delivery configuration id of customer
	LogDeliveryConfigurationId types.String `tfsdk:"-"`
}

func (to *GetLogDeliveryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLogDeliveryRequest_SdkV2) {
}

func (to *GetLogDeliveryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLogDeliveryRequest_SdkV2) {
}

func (m GetLogDeliveryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetLogDeliveryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLogDeliveryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetLogDeliveryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration_id": m.LogDeliveryConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLogDeliveryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration_id": types.StringType,
		},
	}
}

// The limit configuration of the policy. Limit configuration provide a budget
// policy level cost control by enforcing the limit.
type LimitConfig_SdkV2 struct {
}

func (to *LimitConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LimitConfig_SdkV2) {
}

func (to *LimitConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LimitConfig_SdkV2) {
}

func (m LimitConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LimitConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LimitConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LimitConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m LimitConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LimitConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListBudgetConfigurationsRequest_SdkV2 struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListBudgetConfigurationsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBudgetConfigurationsRequest_SdkV2) {
}

func (to *ListBudgetConfigurationsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListBudgetConfigurationsRequest_SdkV2) {
}

func (m ListBudgetConfigurationsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListBudgetConfigurationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetConfigurationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListBudgetConfigurationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBudgetConfigurationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

type ListBudgetConfigurationsResponse_SdkV2 struct {
	Budgets types.List `tfsdk:"budgets"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListBudgetConfigurationsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBudgetConfigurationsResponse_SdkV2) {
	if !from.Budgets.IsNull() && !from.Budgets.IsUnknown() && to.Budgets.IsNull() && len(from.Budgets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Budgets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Budgets = from.Budgets
	}
}

func (to *ListBudgetConfigurationsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListBudgetConfigurationsResponse_SdkV2) {
	if !from.Budgets.IsNull() && !from.Budgets.IsUnknown() && to.Budgets.IsNull() && len(from.Budgets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Budgets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Budgets = from.Budgets
	}
}

func (m ListBudgetConfigurationsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListBudgetConfigurationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budgets": reflect.TypeOf(BudgetConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetConfigurationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListBudgetConfigurationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budgets":         m.Budgets,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBudgetConfigurationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListBudgetConfigurationsResponse_SdkV2) GetBudgets(ctx context.Context) ([]BudgetConfiguration_SdkV2, bool) {
	if m.Budgets.IsNull() || m.Budgets.IsUnknown() {
		return nil, false
	}
	var v []BudgetConfiguration_SdkV2
	d := m.Budgets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBudgets sets the value of the Budgets field in ListBudgetConfigurationsResponse_SdkV2.
func (m *ListBudgetConfigurationsResponse_SdkV2) SetBudgets(ctx context.Context, v []BudgetConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["budgets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Budgets = types.ListValueMust(t, vs)
}

type ListBudgetPoliciesRequest_SdkV2 struct {
	// A filter to apply to the list of policies.
	FilterBy types.List `tfsdk:"-"`
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
	SortSpec types.List `tfsdk:"-"`
}

func (to *ListBudgetPoliciesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBudgetPoliciesRequest_SdkV2) {
	if !from.FilterBy.IsNull() && !from.FilterBy.IsUnknown() {
		if toFilterBy, ok := to.GetFilterBy(ctx); ok {
			if fromFilterBy, ok := from.GetFilterBy(ctx); ok {
				// Recursively sync the fields of FilterBy
				toFilterBy.SyncFieldsDuringCreateOrUpdate(ctx, fromFilterBy)
				to.SetFilterBy(ctx, toFilterBy)
			}
		}
	}
	if !from.SortSpec.IsNull() && !from.SortSpec.IsUnknown() {
		if toSortSpec, ok := to.GetSortSpec(ctx); ok {
			if fromSortSpec, ok := from.GetSortSpec(ctx); ok {
				// Recursively sync the fields of SortSpec
				toSortSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSortSpec)
				to.SetSortSpec(ctx, toSortSpec)
			}
		}
	}
}

func (to *ListBudgetPoliciesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListBudgetPoliciesRequest_SdkV2) {
	if !from.FilterBy.IsNull() && !from.FilterBy.IsUnknown() {
		if toFilterBy, ok := to.GetFilterBy(ctx); ok {
			if fromFilterBy, ok := from.GetFilterBy(ctx); ok {
				toFilterBy.SyncFieldsDuringRead(ctx, fromFilterBy)
				to.SetFilterBy(ctx, toFilterBy)
			}
		}
	}
	if !from.SortSpec.IsNull() && !from.SortSpec.IsUnknown() {
		if toSortSpec, ok := to.GetSortSpec(ctx); ok {
			if fromSortSpec, ok := from.GetSortSpec(ctx); ok {
				toSortSpec.SyncFieldsDuringRead(ctx, fromSortSpec)
				to.SetSortSpec(ctx, toSortSpec)
			}
		}
	}
}

func (m ListBudgetPoliciesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter_by"] = attrs["filter_by"].SetOptional()
	attrs["filter_by"] = attrs["filter_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["sort_spec"] = attrs["sort_spec"].SetOptional()
	attrs["sort_spec"] = attrs["sort_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBudgetPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListBudgetPoliciesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter_by": reflect.TypeOf(Filter_SdkV2{}),
		"sort_spec": reflect.TypeOf(SortSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetPoliciesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListBudgetPoliciesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter_by":  m.FilterBy,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"sort_spec":  m.SortSpec,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBudgetPoliciesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_by": basetypes.ListType{
				ElemType: Filter_SdkV2{}.Type(ctx),
			},
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"sort_spec": basetypes.ListType{
				ElemType: SortSpec_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFilterBy returns the value of the FilterBy field in ListBudgetPoliciesRequest_SdkV2 as
// a Filter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListBudgetPoliciesRequest_SdkV2) GetFilterBy(ctx context.Context) (Filter_SdkV2, bool) {
	var e Filter_SdkV2
	if m.FilterBy.IsNull() || m.FilterBy.IsUnknown() {
		return e, false
	}
	var v []Filter_SdkV2
	d := m.FilterBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilterBy sets the value of the FilterBy field in ListBudgetPoliciesRequest_SdkV2.
func (m *ListBudgetPoliciesRequest_SdkV2) SetFilterBy(ctx context.Context, v Filter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filter_by"]
	m.FilterBy = types.ListValueMust(t, vs)
}

// GetSortSpec returns the value of the SortSpec field in ListBudgetPoliciesRequest_SdkV2 as
// a SortSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListBudgetPoliciesRequest_SdkV2) GetSortSpec(ctx context.Context) (SortSpec_SdkV2, bool) {
	var e SortSpec_SdkV2
	if m.SortSpec.IsNull() || m.SortSpec.IsUnknown() {
		return e, false
	}
	var v []SortSpec_SdkV2
	d := m.SortSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSortSpec sets the value of the SortSpec field in ListBudgetPoliciesRequest_SdkV2.
func (m *ListBudgetPoliciesRequest_SdkV2) SetSortSpec(ctx context.Context, v SortSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sort_spec"]
	m.SortSpec = types.ListValueMust(t, vs)
}

// A list of policies.
type ListBudgetPoliciesResponse_SdkV2 struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Policies types.List `tfsdk:"policies"`
	// A token that can be sent as `page_token` to retrieve the previous page.
	// In this field is omitted, there are no previous pages.
	PreviousPageToken types.String `tfsdk:"previous_page_token"`
}

func (to *ListBudgetPoliciesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBudgetPoliciesResponse_SdkV2) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (to *ListBudgetPoliciesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListBudgetPoliciesResponse_SdkV2) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (m ListBudgetPoliciesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListBudgetPoliciesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(BudgetPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBudgetPoliciesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListBudgetPoliciesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":     m.NextPageToken,
			"policies":            m.Policies,
			"previous_page_token": m.PreviousPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBudgetPoliciesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"policies": basetypes.ListType{
				ElemType: BudgetPolicy_SdkV2{}.Type(ctx),
			},
			"previous_page_token": types.StringType,
		},
	}
}

// GetPolicies returns the value of the Policies field in ListBudgetPoliciesResponse_SdkV2 as
// a slice of BudgetPolicy_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListBudgetPoliciesResponse_SdkV2) GetPolicies(ctx context.Context) ([]BudgetPolicy_SdkV2, bool) {
	if m.Policies.IsNull() || m.Policies.IsUnknown() {
		return nil, false
	}
	var v []BudgetPolicy_SdkV2
	d := m.Policies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicies sets the value of the Policies field in ListBudgetPoliciesResponse_SdkV2.
func (m *ListBudgetPoliciesResponse_SdkV2) SetPolicies(ctx context.Context, v []BudgetPolicy_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Policies = types.ListValueMust(t, vs)
}

type ListLogDeliveryRequest_SdkV2 struct {
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

func (to *ListLogDeliveryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListLogDeliveryRequest_SdkV2) {
}

func (to *ListLogDeliveryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListLogDeliveryRequest_SdkV2) {
}

func (m ListLogDeliveryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListLogDeliveryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListLogDeliveryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListLogDeliveryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials_id":           m.CredentialsId,
			"page_token":               m.PageToken,
			"status":                   m.Status,
			"storage_configuration_id": m.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListLogDeliveryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type LogDeliveryConfiguration_SdkV2 struct {
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
	LogDeliveryStatus types.List `tfsdk:"log_delivery_status"`
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

func (to *LogDeliveryConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogDeliveryConfiguration_SdkV2) {
	if !from.LogDeliveryStatus.IsNull() && !from.LogDeliveryStatus.IsUnknown() {
		if toLogDeliveryStatus, ok := to.GetLogDeliveryStatus(ctx); ok {
			if fromLogDeliveryStatus, ok := from.GetLogDeliveryStatus(ctx); ok {
				// Recursively sync the fields of LogDeliveryStatus
				toLogDeliveryStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromLogDeliveryStatus)
				to.SetLogDeliveryStatus(ctx, toLogDeliveryStatus)
			}
		}
	}
	if !from.WorkspaceIdsFilter.IsNull() && !from.WorkspaceIdsFilter.IsUnknown() && to.WorkspaceIdsFilter.IsNull() && len(from.WorkspaceIdsFilter.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceIdsFilter, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceIdsFilter = from.WorkspaceIdsFilter
	}
}

func (to *LogDeliveryConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogDeliveryConfiguration_SdkV2) {
	if !from.LogDeliveryStatus.IsNull() && !from.LogDeliveryStatus.IsUnknown() {
		if toLogDeliveryStatus, ok := to.GetLogDeliveryStatus(ctx); ok {
			if fromLogDeliveryStatus, ok := from.GetLogDeliveryStatus(ctx); ok {
				toLogDeliveryStatus.SyncFieldsDuringRead(ctx, fromLogDeliveryStatus)
				to.SetLogDeliveryStatus(ctx, toLogDeliveryStatus)
			}
		}
	}
	if !from.WorkspaceIdsFilter.IsNull() && !from.WorkspaceIdsFilter.IsUnknown() && to.WorkspaceIdsFilter.IsNull() && len(from.WorkspaceIdsFilter.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceIdsFilter, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceIdsFilter = from.WorkspaceIdsFilter
	}
}

func (m LogDeliveryConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["config_id"] = attrs["config_id"].SetOptional()
	attrs["config_name"] = attrs["config_name"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["credentials_id"] = attrs["credentials_id"].SetRequired()
	attrs["delivery_path_prefix"] = attrs["delivery_path_prefix"].SetOptional()
	attrs["delivery_start_time"] = attrs["delivery_start_time"].SetOptional()
	attrs["log_delivery_status"] = attrs["log_delivery_status"].SetOptional()
	attrs["log_delivery_status"] = attrs["log_delivery_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m LogDeliveryConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_status":  reflect.TypeOf(LogDeliveryStatus_SdkV2{}),
		"workspace_ids_filter": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogDeliveryConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (m LogDeliveryConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":               m.AccountId,
			"config_id":                m.ConfigId,
			"config_name":              m.ConfigName,
			"creation_time":            m.CreationTime,
			"credentials_id":           m.CredentialsId,
			"delivery_path_prefix":     m.DeliveryPathPrefix,
			"delivery_start_time":      m.DeliveryStartTime,
			"log_delivery_status":      m.LogDeliveryStatus,
			"log_type":                 m.LogType,
			"output_format":            m.OutputFormat,
			"status":                   m.Status,
			"storage_configuration_id": m.StorageConfigurationId,
			"update_time":              m.UpdateTime,
			"workspace_ids_filter":     m.WorkspaceIdsFilter,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogDeliveryConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *LogDeliveryConfiguration_SdkV2) GetLogDeliveryStatus(ctx context.Context) (LogDeliveryStatus_SdkV2, bool) {
	var e LogDeliveryStatus_SdkV2
	if m.LogDeliveryStatus.IsNull() || m.LogDeliveryStatus.IsUnknown() {
		return e, false
	}
	var v []LogDeliveryStatus_SdkV2
	d := m.LogDeliveryStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryStatus sets the value of the LogDeliveryStatus field in LogDeliveryConfiguration_SdkV2.
func (m *LogDeliveryConfiguration_SdkV2) SetLogDeliveryStatus(ctx context.Context, v LogDeliveryStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_status"]
	m.LogDeliveryStatus = types.ListValueMust(t, vs)
}

// GetWorkspaceIdsFilter returns the value of the WorkspaceIdsFilter field in LogDeliveryConfiguration_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogDeliveryConfiguration_SdkV2) GetWorkspaceIdsFilter(ctx context.Context) ([]types.Int64, bool) {
	if m.WorkspaceIdsFilter.IsNull() || m.WorkspaceIdsFilter.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.WorkspaceIdsFilter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceIdsFilter sets the value of the WorkspaceIdsFilter field in LogDeliveryConfiguration_SdkV2.
func (m *LogDeliveryConfiguration_SdkV2) SetWorkspaceIdsFilter(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_ids_filter"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceIdsFilter = types.ListValueMust(t, vs)
}

type LogDeliveryStatus_SdkV2 struct {
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

func (to *LogDeliveryStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogDeliveryStatus_SdkV2) {
}

func (to *LogDeliveryStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogDeliveryStatus_SdkV2) {
}

func (m LogDeliveryStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LogDeliveryStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogDeliveryStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m LogDeliveryStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_attempt_time":            m.LastAttemptTime,
			"last_successful_attempt_time": m.LastSuccessfulAttemptTime,
			"message":                      m.Message,
			"status":                       m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogDeliveryStatus_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *PatchStatusResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PatchStatusResponse_SdkV2) {
}

func (to *PatchStatusResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PatchStatusResponse_SdkV2) {
}

func (m PatchStatusResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PatchStatusResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchStatusResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m PatchStatusResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m PatchStatusResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SortSpec_SdkV2 struct {
	// Whether to sort in descending order.
	Descending types.Bool `tfsdk:"descending"`
	// The filed to sort by
	Field types.String `tfsdk:"field"`
}

func (to *SortSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SortSpec_SdkV2) {
}

func (to *SortSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SortSpec_SdkV2) {
}

func (m SortSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SortSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SortSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m SortSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"descending": m.Descending,
			"field":      m.Field,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SortSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"descending": types.BoolType,
			"field":      types.StringType,
		},
	}
}

type UpdateBudgetConfigurationBudget_SdkV2 struct {
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
	Filter types.List `tfsdk:"filter"`
}

func (to *UpdateBudgetConfigurationBudget_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateBudgetConfigurationBudget_SdkV2) {
	if !from.AlertConfigurations.IsNull() && !from.AlertConfigurations.IsUnknown() && to.AlertConfigurations.IsNull() && len(from.AlertConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AlertConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AlertConfigurations = from.AlertConfigurations
	}
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				// Recursively sync the fields of Filter
				toFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (to *UpdateBudgetConfigurationBudget_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateBudgetConfigurationBudget_SdkV2) {
	if !from.AlertConfigurations.IsNull() && !from.AlertConfigurations.IsUnknown() && to.AlertConfigurations.IsNull() && len(from.AlertConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AlertConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AlertConfigurations = from.AlertConfigurations
	}
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				toFilter.SyncFieldsDuringRead(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (m UpdateBudgetConfigurationBudget_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["alert_configurations"] = attrs["alert_configurations"].SetOptional()
	attrs["budget_configuration_id"] = attrs["budget_configuration_id"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["filter"] = attrs["filter"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetConfigurationBudget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateBudgetConfigurationBudget_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_configurations": reflect.TypeOf(AlertConfiguration_SdkV2{}),
		"filter":               reflect.TypeOf(BudgetConfigurationFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetConfigurationBudget_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateBudgetConfigurationBudget_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":              m.AccountId,
			"alert_configurations":    m.AlertConfigurations,
			"budget_configuration_id": m.BudgetConfigurationId,
			"display_name":            m.DisplayName,
			"filter":                  m.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateBudgetConfigurationBudget_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateBudgetConfigurationBudget_SdkV2) GetAlertConfigurations(ctx context.Context) ([]AlertConfiguration_SdkV2, bool) {
	if m.AlertConfigurations.IsNull() || m.AlertConfigurations.IsUnknown() {
		return nil, false
	}
	var v []AlertConfiguration_SdkV2
	d := m.AlertConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlertConfigurations sets the value of the AlertConfigurations field in UpdateBudgetConfigurationBudget_SdkV2.
func (m *UpdateBudgetConfigurationBudget_SdkV2) SetAlertConfigurations(ctx context.Context, v []AlertConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AlertConfigurations = types.ListValueMust(t, vs)
}

// GetFilter returns the value of the Filter field in UpdateBudgetConfigurationBudget_SdkV2 as
// a BudgetConfigurationFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateBudgetConfigurationBudget_SdkV2) GetFilter(ctx context.Context) (BudgetConfigurationFilter_SdkV2, bool) {
	var e BudgetConfigurationFilter_SdkV2
	if m.Filter.IsNull() || m.Filter.IsUnknown() {
		return e, false
	}
	var v []BudgetConfigurationFilter_SdkV2
	d := m.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in UpdateBudgetConfigurationBudget_SdkV2.
func (m *UpdateBudgetConfigurationBudget_SdkV2) SetFilter(ctx context.Context, v BudgetConfigurationFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	m.Filter = types.ListValueMust(t, vs)
}

type UpdateBudgetConfigurationRequest_SdkV2 struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID.
	Budget types.List `tfsdk:"budget"`
	// The Databricks budget configuration ID.
	BudgetId types.String `tfsdk:"-"`
}

func (to *UpdateBudgetConfigurationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateBudgetConfigurationRequest_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				// Recursively sync the fields of Budget
				toBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (to *UpdateBudgetConfigurationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateBudgetConfigurationRequest_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				toBudget.SyncFieldsDuringRead(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (m UpdateBudgetConfigurationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetRequired()
	attrs["budget"] = attrs["budget"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateBudgetConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(UpdateBudgetConfigurationBudget_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateBudgetConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget":    m.Budget,
			"budget_id": m.BudgetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateBudgetConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateBudgetConfigurationRequest_SdkV2) GetBudget(ctx context.Context) (UpdateBudgetConfigurationBudget_SdkV2, bool) {
	var e UpdateBudgetConfigurationBudget_SdkV2
	if m.Budget.IsNull() || m.Budget.IsUnknown() {
		return e, false
	}
	var v []UpdateBudgetConfigurationBudget_SdkV2
	d := m.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in UpdateBudgetConfigurationRequest_SdkV2.
func (m *UpdateBudgetConfigurationRequest_SdkV2) SetBudget(ctx context.Context, v UpdateBudgetConfigurationBudget_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	m.Budget = types.ListValueMust(t, vs)
}

type UpdateBudgetConfigurationResponse_SdkV2 struct {
	// The updated budget.
	Budget types.List `tfsdk:"budget"`
}

func (to *UpdateBudgetConfigurationResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateBudgetConfigurationResponse_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				// Recursively sync the fields of Budget
				toBudget.SyncFieldsDuringCreateOrUpdate(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (to *UpdateBudgetConfigurationResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateBudgetConfigurationResponse_SdkV2) {
	if !from.Budget.IsNull() && !from.Budget.IsUnknown() {
		if toBudget, ok := to.GetBudget(ctx); ok {
			if fromBudget, ok := from.GetBudget(ctx); ok {
				toBudget.SyncFieldsDuringRead(ctx, fromBudget)
				to.SetBudget(ctx, toBudget)
			}
		}
	}
}

func (m UpdateBudgetConfigurationResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget"] = attrs["budget"].SetOptional()
	attrs["budget"] = attrs["budget"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateBudgetConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget": reflect.TypeOf(BudgetConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateBudgetConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget": m.Budget,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateBudgetConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateBudgetConfigurationResponse_SdkV2) GetBudget(ctx context.Context) (BudgetConfiguration_SdkV2, bool) {
	var e BudgetConfiguration_SdkV2
	if m.Budget.IsNull() || m.Budget.IsUnknown() {
		return e, false
	}
	var v []BudgetConfiguration_SdkV2
	d := m.Budget.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBudget sets the value of the Budget field in UpdateBudgetConfigurationResponse_SdkV2.
func (m *UpdateBudgetConfigurationResponse_SdkV2) SetBudget(ctx context.Context, v BudgetConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["budget"]
	m.Budget = types.ListValueMust(t, vs)
}

type UpdateBudgetPolicyRequest_SdkV2 struct {
	// DEPRECATED. This is redundant field as LimitConfig is part of the
	// BudgetPolicy
	LimitConfig types.List `tfsdk:"-"`
	// The policy to update. `creator_user_id` cannot be specified in the
	// request. All other fields must be specified even if not changed. The
	// `policy_id` is used to identify the policy to update.
	Policy types.List `tfsdk:"policy"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId types.String `tfsdk:"-"`
}

func (to *UpdateBudgetPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateBudgetPolicyRequest_SdkV2) {
	if !from.LimitConfig.IsNull() && !from.LimitConfig.IsUnknown() {
		if toLimitConfig, ok := to.GetLimitConfig(ctx); ok {
			if fromLimitConfig, ok := from.GetLimitConfig(ctx); ok {
				// Recursively sync the fields of LimitConfig
				toLimitConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromLimitConfig)
				to.SetLimitConfig(ctx, toLimitConfig)
			}
		}
	}
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				// Recursively sync the fields of Policy
				toPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (to *UpdateBudgetPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateBudgetPolicyRequest_SdkV2) {
	if !from.LimitConfig.IsNull() && !from.LimitConfig.IsUnknown() {
		if toLimitConfig, ok := to.GetLimitConfig(ctx); ok {
			if fromLimitConfig, ok := from.GetLimitConfig(ctx); ok {
				toLimitConfig.SyncFieldsDuringRead(ctx, fromLimitConfig)
				to.SetLimitConfig(ctx, toLimitConfig)
			}
		}
	}
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m UpdateBudgetPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["policy"] = attrs["policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetRequired()
	attrs["limit_config"] = attrs["limit_config"].SetOptional()
	attrs["limit_config"] = attrs["limit_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBudgetPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateBudgetPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"limit_config": reflect.TypeOf(LimitConfig_SdkV2{}),
		"policy":       reflect.TypeOf(BudgetPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBudgetPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateBudgetPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"limit_config": m.LimitConfig,
			"policy":       m.Policy,
			"policy_id":    m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateBudgetPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"limit_config": basetypes.ListType{
				ElemType: LimitConfig_SdkV2{}.Type(ctx),
			},
			"policy": basetypes.ListType{
				ElemType: BudgetPolicy_SdkV2{}.Type(ctx),
			},
			"policy_id": types.StringType,
		},
	}
}

// GetLimitConfig returns the value of the LimitConfig field in UpdateBudgetPolicyRequest_SdkV2 as
// a LimitConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateBudgetPolicyRequest_SdkV2) GetLimitConfig(ctx context.Context) (LimitConfig_SdkV2, bool) {
	var e LimitConfig_SdkV2
	if m.LimitConfig.IsNull() || m.LimitConfig.IsUnknown() {
		return e, false
	}
	var v []LimitConfig_SdkV2
	d := m.LimitConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLimitConfig sets the value of the LimitConfig field in UpdateBudgetPolicyRequest_SdkV2.
func (m *UpdateBudgetPolicyRequest_SdkV2) SetLimitConfig(ctx context.Context, v LimitConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["limit_config"]
	m.LimitConfig = types.ListValueMust(t, vs)
}

// GetPolicy returns the value of the Policy field in UpdateBudgetPolicyRequest_SdkV2 as
// a BudgetPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateBudgetPolicyRequest_SdkV2) GetPolicy(ctx context.Context) (BudgetPolicy_SdkV2, bool) {
	var e BudgetPolicy_SdkV2
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v []BudgetPolicy_SdkV2
	d := m.Policy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPolicy sets the value of the Policy field in UpdateBudgetPolicyRequest_SdkV2.
func (m *UpdateBudgetPolicyRequest_SdkV2) SetPolicy(ctx context.Context, v BudgetPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy"]
	m.Policy = types.ListValueMust(t, vs)
}

// * Update Log Delivery Configuration
type UpdateLogDeliveryConfigurationStatusRequest_SdkV2 struct {
	// The log delivery configuration id of customer
	LogDeliveryConfigurationId types.String `tfsdk:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status types.String `tfsdk:"status"`
}

func (to *UpdateLogDeliveryConfigurationStatusRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateLogDeliveryConfigurationStatusRequest_SdkV2) {
}

func (to *UpdateLogDeliveryConfigurationStatusRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateLogDeliveryConfigurationStatusRequest_SdkV2) {
}

func (m UpdateLogDeliveryConfigurationStatusRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateLogDeliveryConfigurationStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLogDeliveryConfigurationStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateLogDeliveryConfigurationStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration_id": m.LogDeliveryConfigurationId,
			"status":                        m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateLogDeliveryConfigurationStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration_id": types.StringType,
			"status":                        types.StringType,
		},
	}
}

// * Properties of the new log delivery configuration.
type WrappedCreateLogDeliveryConfiguration_SdkV2 struct {
	LogDeliveryConfiguration types.List `tfsdk:"log_delivery_configuration"`
}

func (to *WrappedCreateLogDeliveryConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WrappedCreateLogDeliveryConfiguration_SdkV2) {
	if !from.LogDeliveryConfiguration.IsNull() && !from.LogDeliveryConfiguration.IsUnknown() {
		if toLogDeliveryConfiguration, ok := to.GetLogDeliveryConfiguration(ctx); ok {
			if fromLogDeliveryConfiguration, ok := from.GetLogDeliveryConfiguration(ctx); ok {
				// Recursively sync the fields of LogDeliveryConfiguration
				toLogDeliveryConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromLogDeliveryConfiguration)
				to.SetLogDeliveryConfiguration(ctx, toLogDeliveryConfiguration)
			}
		}
	}
}

func (to *WrappedCreateLogDeliveryConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WrappedCreateLogDeliveryConfiguration_SdkV2) {
	if !from.LogDeliveryConfiguration.IsNull() && !from.LogDeliveryConfiguration.IsUnknown() {
		if toLogDeliveryConfiguration, ok := to.GetLogDeliveryConfiguration(ctx); ok {
			if fromLogDeliveryConfiguration, ok := from.GetLogDeliveryConfiguration(ctx); ok {
				toLogDeliveryConfiguration.SyncFieldsDuringRead(ctx, fromLogDeliveryConfiguration)
				to.SetLogDeliveryConfiguration(ctx, toLogDeliveryConfiguration)
			}
		}
	}
}

func (m WrappedCreateLogDeliveryConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["log_delivery_configuration"] = attrs["log_delivery_configuration"].SetRequired()
	attrs["log_delivery_configuration"] = attrs["log_delivery_configuration"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m WrappedCreateLogDeliveryConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configuration": reflect.TypeOf(CreateLogDeliveryConfigurationParams_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WrappedCreateLogDeliveryConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (m WrappedCreateLogDeliveryConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration": m.LogDeliveryConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WrappedCreateLogDeliveryConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *WrappedCreateLogDeliveryConfiguration_SdkV2) GetLogDeliveryConfiguration(ctx context.Context) (CreateLogDeliveryConfigurationParams_SdkV2, bool) {
	var e CreateLogDeliveryConfigurationParams_SdkV2
	if m.LogDeliveryConfiguration.IsNull() || m.LogDeliveryConfiguration.IsUnknown() {
		return e, false
	}
	var v []CreateLogDeliveryConfigurationParams_SdkV2
	d := m.LogDeliveryConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in WrappedCreateLogDeliveryConfiguration_SdkV2.
func (m *WrappedCreateLogDeliveryConfiguration_SdkV2) SetLogDeliveryConfiguration(ctx context.Context, v CreateLogDeliveryConfigurationParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configuration"]
	m.LogDeliveryConfiguration = types.ListValueMust(t, vs)
}

type WrappedLogDeliveryConfiguration_SdkV2 struct {
	// The created log delivery configuration
	LogDeliveryConfiguration types.List `tfsdk:"log_delivery_configuration"`
}

func (to *WrappedLogDeliveryConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WrappedLogDeliveryConfiguration_SdkV2) {
	if !from.LogDeliveryConfiguration.IsNull() && !from.LogDeliveryConfiguration.IsUnknown() {
		if toLogDeliveryConfiguration, ok := to.GetLogDeliveryConfiguration(ctx); ok {
			if fromLogDeliveryConfiguration, ok := from.GetLogDeliveryConfiguration(ctx); ok {
				// Recursively sync the fields of LogDeliveryConfiguration
				toLogDeliveryConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromLogDeliveryConfiguration)
				to.SetLogDeliveryConfiguration(ctx, toLogDeliveryConfiguration)
			}
		}
	}
}

func (to *WrappedLogDeliveryConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WrappedLogDeliveryConfiguration_SdkV2) {
	if !from.LogDeliveryConfiguration.IsNull() && !from.LogDeliveryConfiguration.IsUnknown() {
		if toLogDeliveryConfiguration, ok := to.GetLogDeliveryConfiguration(ctx); ok {
			if fromLogDeliveryConfiguration, ok := from.GetLogDeliveryConfiguration(ctx); ok {
				toLogDeliveryConfiguration.SyncFieldsDuringRead(ctx, fromLogDeliveryConfiguration)
				to.SetLogDeliveryConfiguration(ctx, toLogDeliveryConfiguration)
			}
		}
	}
}

func (m WrappedLogDeliveryConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["log_delivery_configuration"] = attrs["log_delivery_configuration"].SetOptional()
	attrs["log_delivery_configuration"] = attrs["log_delivery_configuration"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WrappedLogDeliveryConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WrappedLogDeliveryConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configuration": reflect.TypeOf(LogDeliveryConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WrappedLogDeliveryConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (m WrappedLogDeliveryConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configuration": m.LogDeliveryConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WrappedLogDeliveryConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *WrappedLogDeliveryConfiguration_SdkV2) GetLogDeliveryConfiguration(ctx context.Context) (LogDeliveryConfiguration_SdkV2, bool) {
	var e LogDeliveryConfiguration_SdkV2
	if m.LogDeliveryConfiguration.IsNull() || m.LogDeliveryConfiguration.IsUnknown() {
		return e, false
	}
	var v []LogDeliveryConfiguration_SdkV2
	d := m.LogDeliveryConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogDeliveryConfiguration sets the value of the LogDeliveryConfiguration field in WrappedLogDeliveryConfiguration_SdkV2.
func (m *WrappedLogDeliveryConfiguration_SdkV2) SetLogDeliveryConfiguration(ctx context.Context, v LogDeliveryConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configuration"]
	m.LogDeliveryConfiguration = types.ListValueMust(t, vs)
}

type WrappedLogDeliveryConfigurations_SdkV2 struct {
	// Log delivery configurations were returned successfully.
	LogDeliveryConfigurations types.List `tfsdk:"log_delivery_configurations"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *WrappedLogDeliveryConfigurations_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WrappedLogDeliveryConfigurations_SdkV2) {
	if !from.LogDeliveryConfigurations.IsNull() && !from.LogDeliveryConfigurations.IsUnknown() && to.LogDeliveryConfigurations.IsNull() && len(from.LogDeliveryConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LogDeliveryConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LogDeliveryConfigurations = from.LogDeliveryConfigurations
	}
}

func (to *WrappedLogDeliveryConfigurations_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WrappedLogDeliveryConfigurations_SdkV2) {
	if !from.LogDeliveryConfigurations.IsNull() && !from.LogDeliveryConfigurations.IsUnknown() && to.LogDeliveryConfigurations.IsNull() && len(from.LogDeliveryConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LogDeliveryConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LogDeliveryConfigurations = from.LogDeliveryConfigurations
	}
}

func (m WrappedLogDeliveryConfigurations_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WrappedLogDeliveryConfigurations_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_delivery_configurations": reflect.TypeOf(LogDeliveryConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WrappedLogDeliveryConfigurations_SdkV2
// only implements ToObjectValue() and Type().
func (m WrappedLogDeliveryConfigurations_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_delivery_configurations": m.LogDeliveryConfigurations,
			"next_page_token":             m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WrappedLogDeliveryConfigurations_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configurations": basetypes.ListType{
				ElemType: LogDeliveryConfiguration_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetLogDeliveryConfigurations returns the value of the LogDeliveryConfigurations field in WrappedLogDeliveryConfigurations_SdkV2 as
// a slice of LogDeliveryConfiguration_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *WrappedLogDeliveryConfigurations_SdkV2) GetLogDeliveryConfigurations(ctx context.Context) ([]LogDeliveryConfiguration_SdkV2, bool) {
	if m.LogDeliveryConfigurations.IsNull() || m.LogDeliveryConfigurations.IsUnknown() {
		return nil, false
	}
	var v []LogDeliveryConfiguration_SdkV2
	d := m.LogDeliveryConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLogDeliveryConfigurations sets the value of the LogDeliveryConfigurations field in WrappedLogDeliveryConfigurations_SdkV2.
func (m *WrappedLogDeliveryConfigurations_SdkV2) SetLogDeliveryConfigurations(ctx context.Context, v []LogDeliveryConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["log_delivery_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LogDeliveryConfigurations = types.ListValueMust(t, vs)
}
