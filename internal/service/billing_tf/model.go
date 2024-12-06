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
	"fmt"
	"io"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ActionConfiguration{}

// Equal implements basetypes.ObjectValuable.
func (o ActionConfiguration) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ActionConfiguration) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ActionConfiguration) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ActionConfiguration) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ActionConfiguration) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ActionConfiguration) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = AlertConfiguration{}

// Equal implements basetypes.ObjectValuable.
func (o AlertConfiguration) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o AlertConfiguration) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o AlertConfiguration) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o AlertConfiguration) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o AlertConfiguration) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o AlertConfiguration) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = BudgetConfiguration{}

// Equal implements basetypes.ObjectValuable.
func (o BudgetConfiguration) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o BudgetConfiguration) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o BudgetConfiguration) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o BudgetConfiguration) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o BudgetConfiguration) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o BudgetConfiguration) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = BudgetConfigurationFilter{}

// Equal implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilter) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilter) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilter) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilter) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilter) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilter) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = BudgetConfigurationFilterClause{}

// Equal implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterClause) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterClause) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterClause) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterClause) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterClause) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterClause) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = BudgetConfigurationFilterTagClause{}

// Equal implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterTagClause) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterTagClause) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterTagClause) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterTagClause) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterTagClause) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterTagClause) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = BudgetConfigurationFilterWorkspaceIdClause{}

// Equal implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterWorkspaceIdClause) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterWorkspaceIdClause) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterWorkspaceIdClause) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterWorkspaceIdClause) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterWorkspaceIdClause) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o BudgetConfigurationFilterWorkspaceIdClause) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateBillingUsageDashboardRequest{}

// Equal implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateBillingUsageDashboardResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateBillingUsageDashboardResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateBudgetConfigurationBudget{}

// Equal implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudget) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudget) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudget) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudget) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudget) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudget) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateBudgetConfigurationBudgetActionConfigurations{}

// Equal implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetActionConfigurations) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetActionConfigurations) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetActionConfigurations) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetActionConfigurations) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetActionConfigurations) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetActionConfigurations) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateBudgetConfigurationBudgetAlertConfigurations{}

// Equal implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetAlertConfigurations) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetAlertConfigurations) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetAlertConfigurations) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetAlertConfigurations) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetAlertConfigurations) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationBudgetAlertConfigurations) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateBudgetConfigurationRequest{}

// Equal implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateBudgetConfigurationResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateBudgetConfigurationResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateLogDeliveryConfigurationParams{}

// Equal implements basetypes.ObjectValuable.
func (o CreateLogDeliveryConfigurationParams) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateLogDeliveryConfigurationParams) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateLogDeliveryConfigurationParams) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateLogDeliveryConfigurationParams) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateLogDeliveryConfigurationParams) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateLogDeliveryConfigurationParams) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteBudgetConfigurationRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteBudgetConfigurationResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteBudgetConfigurationResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DownloadRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DownloadRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DownloadRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DownloadRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DownloadRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DownloadRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DownloadRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Contents io.ReadCloser `tfsdk:"-"`
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DownloadResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DownloadResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DownloadResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DownloadResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DownloadResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DownloadResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DownloadResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetBillingUsageDashboardRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetBillingUsageDashboardResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetBillingUsageDashboardResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetBudgetConfigurationRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetBudgetConfigurationResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetBudgetConfigurationResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetLogDeliveryRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetLogDeliveryRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetLogDeliveryRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetLogDeliveryRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetLogDeliveryRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetLogDeliveryRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetLogDeliveryRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListBudgetConfigurationsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListBudgetConfigurationsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListBudgetConfigurationsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListLogDeliveryRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListLogDeliveryRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListLogDeliveryRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListLogDeliveryRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListLogDeliveryRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListLogDeliveryRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListLogDeliveryRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogDeliveryConfiguration{}

// Equal implements basetypes.ObjectValuable.
func (o LogDeliveryConfiguration) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogDeliveryConfiguration) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogDeliveryConfiguration) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogDeliveryConfiguration) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogDeliveryConfiguration) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogDeliveryConfiguration) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogDeliveryStatus{}

// Equal implements basetypes.ObjectValuable.
func (o LogDeliveryStatus) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogDeliveryStatus) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogDeliveryStatus) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogDeliveryStatus) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogDeliveryStatus) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogDeliveryStatus) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PatchStatusResponse{}

// Equal implements basetypes.ObjectValuable.
func (o PatchStatusResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PatchStatusResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PatchStatusResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PatchStatusResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PatchStatusResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PatchStatusResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateBudgetConfigurationBudget{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationBudget) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationBudget) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationBudget) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationBudget) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationBudget) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationBudget) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateBudgetConfigurationRequest{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateBudgetConfigurationResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateBudgetConfigurationResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateLogDeliveryConfigurationStatusRequest{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateLogDeliveryConfigurationStatusRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateLogDeliveryConfigurationStatusRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateLogDeliveryConfigurationStatusRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateLogDeliveryConfigurationStatusRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateLogDeliveryConfigurationStatusRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateLogDeliveryConfigurationStatusRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = WrappedCreateLogDeliveryConfiguration{}

// Equal implements basetypes.ObjectValuable.
func (o WrappedCreateLogDeliveryConfiguration) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o WrappedCreateLogDeliveryConfiguration) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o WrappedCreateLogDeliveryConfiguration) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o WrappedCreateLogDeliveryConfiguration) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o WrappedCreateLogDeliveryConfiguration) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o WrappedCreateLogDeliveryConfiguration) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = WrappedLogDeliveryConfiguration{}

// Equal implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfiguration) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfiguration) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfiguration) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfiguration) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfiguration) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfiguration) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = WrappedLogDeliveryConfigurations{}

// Equal implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfigurations) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfigurations) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfigurations) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfigurations) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfigurations) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o WrappedLogDeliveryConfigurations) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// The status string for log delivery. Possible values are: * `CREATED`: There
// were no log delivery attempts since the config was created. * `SUCCEEDED`:
// The latest attempt of log delivery has succeeded completely. *
// `USER_FAILURE`: The latest attempt of log delivery failed because of
// misconfiguration of customer provided permissions on role or storage. *
// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
// Databricks internal error. Contact support if it doesn't go away soon. *
// `NOT_FOUND`: The log delivery status as the configuration has been disabled
// since the release of this feature or there are no workspaces in the account.

// Status of log delivery configuration. Set to `ENABLED` (enabled) or
// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable the
// configuration](#operation/patch-log-delivery-config-status) later. Deletion
// of a configuration is not supported, so disable a log delivery configuration
// that is no longer needed.

// Log delivery type. Supported values are:
//
// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the CSV
// schema, see the [View billable usage].
//
// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema, see
// [Configure audit logging]
//
// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html

// The file type of log delivery.
//
// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the CSV
// (comma-separated values) format is supported. For the schema, see the [View
// billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be `JSON`.
// Only the JSON (JavaScript Object Notation) format is supported. For the
// schema, see the [Configuring audit logs].
//
// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
