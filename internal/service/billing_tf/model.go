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
	"io"
	"reflect"

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

// ToObjectType returns the representation of ActionConfiguration in the Terraform plugin framework type
// system.
func (a ActionConfiguration) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of AlertConfiguration in the Terraform plugin framework type
// system.
func (a AlertConfiguration) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_configurations": basetypes.ListType{
				ElemType: ActionConfiguration{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of BudgetConfiguration in the Terraform plugin framework type
// system.
func (a BudgetConfiguration) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"alert_configurations": basetypes.ListType{
				ElemType: AlertConfiguration{}.ToObjectType(ctx),
			},
			"budget_configuration_id": types.StringType,
			"create_time":             types.Int64Type,
			"display_name":            types.StringType,
			"filter": basetypes.ListType{
				ElemType: BudgetConfigurationFilter{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of BudgetConfigurationFilter in the Terraform plugin framework type
// system.
func (a BudgetConfigurationFilter) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tags": basetypes.ListType{
				ElemType: BudgetConfigurationFilterTagClause{}.ToObjectType(ctx),
			},
			"workspace_id": basetypes.ListType{
				ElemType: BudgetConfigurationFilterWorkspaceIdClause{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of BudgetConfigurationFilterClause in the Terraform plugin framework type
// system.
func (a BudgetConfigurationFilterClause) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of BudgetConfigurationFilterTagClause in the Terraform plugin framework type
// system.
func (a BudgetConfigurationFilterTagClause) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key": types.StringType,
			"value": basetypes.ListType{
				ElemType: BudgetConfigurationFilterClause{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of BudgetConfigurationFilterWorkspaceIdClause in the Terraform plugin framework type
// system.
func (a BudgetConfigurationFilterWorkspaceIdClause) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of CreateBillingUsageDashboardRequest in the Terraform plugin framework type
// system.
func (a CreateBillingUsageDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of CreateBillingUsageDashboardResponse in the Terraform plugin framework type
// system.
func (a CreateBillingUsageDashboardResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of CreateBudgetConfigurationBudget in the Terraform plugin framework type
// system.
func (a CreateBudgetConfigurationBudget) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"alert_configurations": basetypes.ListType{
				ElemType: CreateBudgetConfigurationBudgetAlertConfigurations{}.ToObjectType(ctx),
			},
			"display_name": types.StringType,
			"filter": basetypes.ListType{
				ElemType: BudgetConfigurationFilter{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of CreateBudgetConfigurationBudgetActionConfigurations in the Terraform plugin framework type
// system.
func (a CreateBudgetConfigurationBudgetActionConfigurations) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of CreateBudgetConfigurationBudgetAlertConfigurations in the Terraform plugin framework type
// system.
func (a CreateBudgetConfigurationBudgetAlertConfigurations) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_configurations": basetypes.ListType{
				ElemType: CreateBudgetConfigurationBudgetActionConfigurations{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of CreateBudgetConfigurationRequest in the Terraform plugin framework type
// system.
func (a CreateBudgetConfigurationRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: CreateBudgetConfigurationBudget{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of CreateBudgetConfigurationResponse in the Terraform plugin framework type
// system.
func (a CreateBudgetConfigurationResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: BudgetConfiguration{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of CreateLogDeliveryConfigurationParams in the Terraform plugin framework type
// system.
func (a CreateLogDeliveryConfigurationParams) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of DeleteBudgetConfigurationRequest in the Terraform plugin framework type
// system.
func (a DeleteBudgetConfigurationRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of DeleteBudgetConfigurationResponse in the Terraform plugin framework type
// system.
func (a DeleteBudgetConfigurationResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of DownloadRequest in the Terraform plugin framework type
// system.
func (a DownloadRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of DownloadResponse in the Terraform plugin framework type
// system.
func (a DownloadResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of GetBillingUsageDashboardRequest in the Terraform plugin framework type
// system.
func (a GetBillingUsageDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of GetBillingUsageDashboardResponse in the Terraform plugin framework type
// system.
func (a GetBillingUsageDashboardResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of GetBudgetConfigurationRequest in the Terraform plugin framework type
// system.
func (a GetBudgetConfigurationRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of GetBudgetConfigurationResponse in the Terraform plugin framework type
// system.
func (a GetBudgetConfigurationResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: BudgetConfiguration{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of GetLogDeliveryRequest in the Terraform plugin framework type
// system.
func (a GetLogDeliveryRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of ListBudgetConfigurationsRequest in the Terraform plugin framework type
// system.
func (a ListBudgetConfigurationsRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of ListBudgetConfigurationsResponse in the Terraform plugin framework type
// system.
func (a ListBudgetConfigurationsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budgets": basetypes.ListType{
				ElemType: BudgetConfiguration{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of ListLogDeliveryRequest in the Terraform plugin framework type
// system.
func (a ListLogDeliveryRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of LogDeliveryConfiguration in the Terraform plugin framework type
// system.
func (a LogDeliveryConfiguration) ToObjectType(ctx context.Context) types.ObjectType {
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
				ElemType: LogDeliveryStatus{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of LogDeliveryStatus in the Terraform plugin framework type
// system.
func (a LogDeliveryStatus) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of PatchStatusResponse in the Terraform plugin framework type
// system.
func (a PatchStatusResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of UpdateBudgetConfigurationBudget in the Terraform plugin framework type
// system.
func (a UpdateBudgetConfigurationBudget) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"alert_configurations": basetypes.ListType{
				ElemType: AlertConfiguration{}.ToObjectType(ctx),
			},
			"budget_configuration_id": types.StringType,
			"display_name":            types.StringType,
			"filter": basetypes.ListType{
				ElemType: BudgetConfigurationFilter{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of UpdateBudgetConfigurationRequest in the Terraform plugin framework type
// system.
func (a UpdateBudgetConfigurationRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: UpdateBudgetConfigurationBudget{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of UpdateBudgetConfigurationResponse in the Terraform plugin framework type
// system.
func (a UpdateBudgetConfigurationResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget": basetypes.ListType{
				ElemType: BudgetConfiguration{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of UpdateLogDeliveryConfigurationStatusRequest in the Terraform plugin framework type
// system.
func (a UpdateLogDeliveryConfigurationStatusRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of WrappedCreateLogDeliveryConfiguration in the Terraform plugin framework type
// system.
func (a WrappedCreateLogDeliveryConfiguration) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration": basetypes.ListType{
				ElemType: CreateLogDeliveryConfigurationParams{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of WrappedLogDeliveryConfiguration in the Terraform plugin framework type
// system.
func (a WrappedLogDeliveryConfiguration) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configuration": basetypes.ListType{
				ElemType: LogDeliveryConfiguration{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of WrappedLogDeliveryConfigurations in the Terraform plugin framework type
// system.
func (a WrappedLogDeliveryConfigurations) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_delivery_configurations": basetypes.ListType{
				ElemType: LogDeliveryConfiguration{}.ToObjectType(ctx),
			},
		},
	}
}
