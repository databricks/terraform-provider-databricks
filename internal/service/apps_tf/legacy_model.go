// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package apps_tf

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

type App_SdkV2 struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment types.List `tfsdk:"active_deployment"`

	AppStatus types.List `tfsdk:"app_status"`

	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`

	ComputeStatus types.List `tfsdk:"compute_status"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time"`
	// The email of the user that created the app.
	Creator types.String `tfsdk:"creator"`
	// The default workspace file system path of the source code from which app
	// deployment are created. This field tracks the workspace source code path
	// of the last active deployment.
	DefaultSourceCodePath types.String `tfsdk:"default_source_code_path"`
	// The description of the app.
	Description types.String `tfsdk:"description"`

	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
	// The effective api scopes granted to the user access token.
	EffectiveUserApiScopes types.List `tfsdk:"effective_user_api_scopes"`
	// The unique identifier of the app.
	Id types.String `tfsdk:"id"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name"`

	Oauth2AppClientId types.String `tfsdk:"oauth2_app_client_id"`

	Oauth2AppIntegrationId types.String `tfsdk:"oauth2_app_integration_id"`
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	PendingDeployment types.List `tfsdk:"pending_deployment"`
	// Resources for the app.
	Resources types.List `tfsdk:"resources"`

	ServicePrincipalClientId types.String `tfsdk:"service_principal_client_id"`

	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id"`

	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time"`
	// The email of the user that last updated the app.
	Updater types.String `tfsdk:"updater"`
	// The URL of the app once it is deployed.
	Url types.String `tfsdk:"url"`

	UserApiScopes types.List `tfsdk:"user_api_scopes"`
}

func (toState *App_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan App_SdkV2) {
	if !fromPlan.ActiveDeployment.IsNull() && !fromPlan.ActiveDeployment.IsUnknown() {
		if toStateActiveDeployment, ok := toState.GetActiveDeployment(ctx); ok {
			if fromPlanActiveDeployment, ok := fromPlan.GetActiveDeployment(ctx); ok {
				toStateActiveDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanActiveDeployment)
				toState.SetActiveDeployment(ctx, toStateActiveDeployment)
			}
		}
	}
	if !fromPlan.AppStatus.IsNull() && !fromPlan.AppStatus.IsUnknown() {
		if toStateAppStatus, ok := toState.GetAppStatus(ctx); ok {
			if fromPlanAppStatus, ok := fromPlan.GetAppStatus(ctx); ok {
				toStateAppStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAppStatus)
				toState.SetAppStatus(ctx, toStateAppStatus)
			}
		}
	}
	if !fromPlan.ComputeStatus.IsNull() && !fromPlan.ComputeStatus.IsUnknown() {
		if toStateComputeStatus, ok := toState.GetComputeStatus(ctx); ok {
			if fromPlanComputeStatus, ok := fromPlan.GetComputeStatus(ctx); ok {
				toStateComputeStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanComputeStatus)
				toState.SetComputeStatus(ctx, toStateComputeStatus)
			}
		}
	}
	if !fromPlan.PendingDeployment.IsNull() && !fromPlan.PendingDeployment.IsUnknown() {
		if toStatePendingDeployment, ok := toState.GetPendingDeployment(ctx); ok {
			if fromPlanPendingDeployment, ok := fromPlan.GetPendingDeployment(ctx); ok {
				toStatePendingDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanPendingDeployment)
				toState.SetPendingDeployment(ctx, toStatePendingDeployment)
			}
		}
	}
}

func (toState *App_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState App_SdkV2) {
	if !fromState.ActiveDeployment.IsNull() && !fromState.ActiveDeployment.IsUnknown() {
		if toStateActiveDeployment, ok := toState.GetActiveDeployment(ctx); ok {
			if fromStateActiveDeployment, ok := fromState.GetActiveDeployment(ctx); ok {
				toStateActiveDeployment.SyncFieldsDuringRead(ctx, fromStateActiveDeployment)
				toState.SetActiveDeployment(ctx, toStateActiveDeployment)
			}
		}
	}
	if !fromState.AppStatus.IsNull() && !fromState.AppStatus.IsUnknown() {
		if toStateAppStatus, ok := toState.GetAppStatus(ctx); ok {
			if fromStateAppStatus, ok := fromState.GetAppStatus(ctx); ok {
				toStateAppStatus.SyncFieldsDuringRead(ctx, fromStateAppStatus)
				toState.SetAppStatus(ctx, toStateAppStatus)
			}
		}
	}
	if !fromState.ComputeStatus.IsNull() && !fromState.ComputeStatus.IsUnknown() {
		if toStateComputeStatus, ok := toState.GetComputeStatus(ctx); ok {
			if fromStateComputeStatus, ok := fromState.GetComputeStatus(ctx); ok {
				toStateComputeStatus.SyncFieldsDuringRead(ctx, fromStateComputeStatus)
				toState.SetComputeStatus(ctx, toStateComputeStatus)
			}
		}
	}
	if !fromState.PendingDeployment.IsNull() && !fromState.PendingDeployment.IsUnknown() {
		if toStatePendingDeployment, ok := toState.GetPendingDeployment(ctx); ok {
			if fromStatePendingDeployment, ok := fromState.GetPendingDeployment(ctx); ok {
				toStatePendingDeployment.SyncFieldsDuringRead(ctx, fromStatePendingDeployment)
				toState.SetPendingDeployment(ctx, toStatePendingDeployment)
			}
		}
	}
}

func (c App_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["active_deployment"] = attrs["active_deployment"].SetComputed()
	attrs["active_deployment"] = attrs["active_deployment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["app_status"] = attrs["app_status"].SetComputed()
	attrs["app_status"] = attrs["app_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["compute_status"] = attrs["compute_status"].SetComputed()
	attrs["compute_status"] = attrs["compute_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["default_source_code_path"] = attrs["default_source_code_path"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetComputed()
	attrs["effective_user_api_scopes"] = attrs["effective_user_api_scopes"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["oauth2_app_client_id"] = attrs["oauth2_app_client_id"].SetComputed()
	attrs["oauth2_app_integration_id"] = attrs["oauth2_app_integration_id"].SetComputed()
	attrs["pending_deployment"] = attrs["pending_deployment"].SetComputed()
	attrs["pending_deployment"] = attrs["pending_deployment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["resources"] = attrs["resources"].SetOptional()
	attrs["service_principal_client_id"] = attrs["service_principal_client_id"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updater"] = attrs["updater"].SetComputed()
	attrs["url"] = attrs["url"].SetComputed()
	attrs["user_api_scopes"] = attrs["user_api_scopes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in App.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a App_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"active_deployment":         reflect.TypeOf(AppDeployment_SdkV2{}),
		"app_status":                reflect.TypeOf(ApplicationStatus_SdkV2{}),
		"compute_status":            reflect.TypeOf(ComputeStatus_SdkV2{}),
		"effective_user_api_scopes": reflect.TypeOf(types.String{}),
		"pending_deployment":        reflect.TypeOf(AppDeployment_SdkV2{}),
		"resources":                 reflect.TypeOf(AppResource_SdkV2{}),
		"user_api_scopes":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, App_SdkV2
// only implements ToObjectValue() and Type().
func (o App_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"active_deployment":           o.ActiveDeployment,
			"app_status":                  o.AppStatus,
			"budget_policy_id":            o.BudgetPolicyId,
			"compute_status":              o.ComputeStatus,
			"create_time":                 o.CreateTime,
			"creator":                     o.Creator,
			"default_source_code_path":    o.DefaultSourceCodePath,
			"description":                 o.Description,
			"effective_budget_policy_id":  o.EffectiveBudgetPolicyId,
			"effective_user_api_scopes":   o.EffectiveUserApiScopes,
			"id":                          o.Id,
			"name":                        o.Name,
			"oauth2_app_client_id":        o.Oauth2AppClientId,
			"oauth2_app_integration_id":   o.Oauth2AppIntegrationId,
			"pending_deployment":          o.PendingDeployment,
			"resources":                   o.Resources,
			"service_principal_client_id": o.ServicePrincipalClientId,
			"service_principal_id":        o.ServicePrincipalId,
			"service_principal_name":      o.ServicePrincipalName,
			"update_time":                 o.UpdateTime,
			"updater":                     o.Updater,
			"url":                         o.Url,
			"user_api_scopes":             o.UserApiScopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o App_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active_deployment": basetypes.ListType{
				ElemType: AppDeployment_SdkV2{}.Type(ctx),
			},
			"app_status": basetypes.ListType{
				ElemType: ApplicationStatus_SdkV2{}.Type(ctx),
			},
			"budget_policy_id": types.StringType,
			"compute_status": basetypes.ListType{
				ElemType: ComputeStatus_SdkV2{}.Type(ctx),
			},
			"create_time":                types.StringType,
			"creator":                    types.StringType,
			"default_source_code_path":   types.StringType,
			"description":                types.StringType,
			"effective_budget_policy_id": types.StringType,
			"effective_user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"id":                        types.StringType,
			"name":                      types.StringType,
			"oauth2_app_client_id":      types.StringType,
			"oauth2_app_integration_id": types.StringType,
			"pending_deployment": basetypes.ListType{
				ElemType: AppDeployment_SdkV2{}.Type(ctx),
			},
			"resources": basetypes.ListType{
				ElemType: AppResource_SdkV2{}.Type(ctx),
			},
			"service_principal_client_id": types.StringType,
			"service_principal_id":        types.Int64Type,
			"service_principal_name":      types.StringType,
			"update_time":                 types.StringType,
			"updater":                     types.StringType,
			"url":                         types.StringType,
			"user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetActiveDeployment returns the value of the ActiveDeployment field in App_SdkV2 as
// a AppDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *App_SdkV2) GetActiveDeployment(ctx context.Context) (AppDeployment_SdkV2, bool) {
	var e AppDeployment_SdkV2
	if o.ActiveDeployment.IsNull() || o.ActiveDeployment.IsUnknown() {
		return e, false
	}
	var v []AppDeployment_SdkV2
	d := o.ActiveDeployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActiveDeployment sets the value of the ActiveDeployment field in App_SdkV2.
func (o *App_SdkV2) SetActiveDeployment(ctx context.Context, v AppDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["active_deployment"]
	o.ActiveDeployment = types.ListValueMust(t, vs)
}

// GetAppStatus returns the value of the AppStatus field in App_SdkV2 as
// a ApplicationStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *App_SdkV2) GetAppStatus(ctx context.Context) (ApplicationStatus_SdkV2, bool) {
	var e ApplicationStatus_SdkV2
	if o.AppStatus.IsNull() || o.AppStatus.IsUnknown() {
		return e, false
	}
	var v []ApplicationStatus_SdkV2
	d := o.AppStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAppStatus sets the value of the AppStatus field in App_SdkV2.
func (o *App_SdkV2) SetAppStatus(ctx context.Context, v ApplicationStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["app_status"]
	o.AppStatus = types.ListValueMust(t, vs)
}

// GetComputeStatus returns the value of the ComputeStatus field in App_SdkV2 as
// a ComputeStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *App_SdkV2) GetComputeStatus(ctx context.Context) (ComputeStatus_SdkV2, bool) {
	var e ComputeStatus_SdkV2
	if o.ComputeStatus.IsNull() || o.ComputeStatus.IsUnknown() {
		return e, false
	}
	var v []ComputeStatus_SdkV2
	d := o.ComputeStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComputeStatus sets the value of the ComputeStatus field in App_SdkV2.
func (o *App_SdkV2) SetComputeStatus(ctx context.Context, v ComputeStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compute_status"]
	o.ComputeStatus = types.ListValueMust(t, vs)
}

// GetEffectiveUserApiScopes returns the value of the EffectiveUserApiScopes field in App_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *App_SdkV2) GetEffectiveUserApiScopes(ctx context.Context) ([]types.String, bool) {
	if o.EffectiveUserApiScopes.IsNull() || o.EffectiveUserApiScopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.EffectiveUserApiScopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveUserApiScopes sets the value of the EffectiveUserApiScopes field in App_SdkV2.
func (o *App_SdkV2) SetEffectiveUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EffectiveUserApiScopes = types.ListValueMust(t, vs)
}

// GetPendingDeployment returns the value of the PendingDeployment field in App_SdkV2 as
// a AppDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *App_SdkV2) GetPendingDeployment(ctx context.Context) (AppDeployment_SdkV2, bool) {
	var e AppDeployment_SdkV2
	if o.PendingDeployment.IsNull() || o.PendingDeployment.IsUnknown() {
		return e, false
	}
	var v []AppDeployment_SdkV2
	d := o.PendingDeployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPendingDeployment sets the value of the PendingDeployment field in App_SdkV2.
func (o *App_SdkV2) SetPendingDeployment(ctx context.Context, v AppDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pending_deployment"]
	o.PendingDeployment = types.ListValueMust(t, vs)
}

// GetResources returns the value of the Resources field in App_SdkV2 as
// a slice of AppResource_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *App_SdkV2) GetResources(ctx context.Context) ([]AppResource_SdkV2, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []AppResource_SdkV2
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in App_SdkV2.
func (o *App_SdkV2) SetResources(ctx context.Context, v []AppResource_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// GetUserApiScopes returns the value of the UserApiScopes field in App_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *App_SdkV2) GetUserApiScopes(ctx context.Context) ([]types.String, bool) {
	if o.UserApiScopes.IsNull() || o.UserApiScopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.UserApiScopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUserApiScopes sets the value of the UserApiScopes field in App_SdkV2.
func (o *App_SdkV2) SetUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UserApiScopes = types.ListValueMust(t, vs)
}

type AppAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (toState *AppAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppAccessControlRequest_SdkV2) {
}

func (toState *AppAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppAccessControlRequest_SdkV2) {
}

func (c AppAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o AppAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AppAccessControlResponse_SdkV2 struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (toState *AppAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppAccessControlResponse_SdkV2) {
}

func (toState *AppAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppAccessControlResponse_SdkV2) {
}

func (c AppAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(AppPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o AppAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: AppPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in AppAccessControlResponse_SdkV2 as
// a slice of AppPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AppAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]AppPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []AppPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in AppAccessControlResponse_SdkV2.
func (o *AppAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []AppPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type AppDeployment_SdkV2 struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time"`
	// The email of the user creates the deployment.
	Creator types.String `tfsdk:"creator"`
	// The deployment artifacts for an app.
	DeploymentArtifacts types.List `tfsdk:"deployment_artifacts"`
	// The unique id of the deployment.
	DeploymentId types.String `tfsdk:"deployment_id"`
	// The mode of which the deployment will manage the source code.
	Mode types.String `tfsdk:"mode"`
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	SourceCodePath types.String `tfsdk:"source_code_path"`
	// Status and status message of the deployment
	Status types.List `tfsdk:"status"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (toState *AppDeployment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppDeployment_SdkV2) {
	if !fromPlan.DeploymentArtifacts.IsNull() && !fromPlan.DeploymentArtifacts.IsUnknown() {
		if toStateDeploymentArtifacts, ok := toState.GetDeploymentArtifacts(ctx); ok {
			if fromPlanDeploymentArtifacts, ok := fromPlan.GetDeploymentArtifacts(ctx); ok {
				toStateDeploymentArtifacts.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanDeploymentArtifacts)
				toState.SetDeploymentArtifacts(ctx, toStateDeploymentArtifacts)
			}
		}
	}
	if !fromPlan.Status.IsNull() && !fromPlan.Status.IsUnknown() {
		if toStateStatus, ok := toState.GetStatus(ctx); ok {
			if fromPlanStatus, ok := fromPlan.GetStatus(ctx); ok {
				toStateStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanStatus)
				toState.SetStatus(ctx, toStateStatus)
			}
		}
	}
}

func (toState *AppDeployment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppDeployment_SdkV2) {
	if !fromState.DeploymentArtifacts.IsNull() && !fromState.DeploymentArtifacts.IsUnknown() {
		if toStateDeploymentArtifacts, ok := toState.GetDeploymentArtifacts(ctx); ok {
			if fromStateDeploymentArtifacts, ok := fromState.GetDeploymentArtifacts(ctx); ok {
				toStateDeploymentArtifacts.SyncFieldsDuringRead(ctx, fromStateDeploymentArtifacts)
				toState.SetDeploymentArtifacts(ctx, toStateDeploymentArtifacts)
			}
		}
	}
	if !fromState.Status.IsNull() && !fromState.Status.IsUnknown() {
		if toStateStatus, ok := toState.GetStatus(ctx); ok {
			if fromStateStatus, ok := fromState.GetStatus(ctx); ok {
				toStateStatus.SyncFieldsDuringRead(ctx, fromStateStatus)
				toState.SetStatus(ctx, toStateStatus)
			}
		}
	}
}

func (c AppDeployment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["deployment_artifacts"] = attrs["deployment_artifacts"].SetComputed()
	attrs["deployment_artifacts"] = attrs["deployment_artifacts"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["deployment_id"] = attrs["deployment_id"].SetOptional()
	attrs["mode"] = attrs["mode"].SetOptional()
	attrs["source_code_path"] = attrs["source_code_path"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppDeployment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppDeployment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"deployment_artifacts": reflect.TypeOf(AppDeploymentArtifacts_SdkV2{}),
		"status":               reflect.TypeOf(AppDeploymentStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppDeployment_SdkV2
// only implements ToObjectValue() and Type().
func (o AppDeployment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":          o.CreateTime,
			"creator":              o.Creator,
			"deployment_artifacts": o.DeploymentArtifacts,
			"deployment_id":        o.DeploymentId,
			"mode":                 o.Mode,
			"source_code_path":     o.SourceCodePath,
			"status":               o.Status,
			"update_time":          o.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppDeployment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"creator":     types.StringType,
			"deployment_artifacts": basetypes.ListType{
				ElemType: AppDeploymentArtifacts_SdkV2{}.Type(ctx),
			},
			"deployment_id":    types.StringType,
			"mode":             types.StringType,
			"source_code_path": types.StringType,
			"status": basetypes.ListType{
				ElemType: AppDeploymentStatus_SdkV2{}.Type(ctx),
			},
			"update_time": types.StringType,
		},
	}
}

// GetDeploymentArtifacts returns the value of the DeploymentArtifacts field in AppDeployment_SdkV2 as
// a AppDeploymentArtifacts_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppDeployment_SdkV2) GetDeploymentArtifacts(ctx context.Context) (AppDeploymentArtifacts_SdkV2, bool) {
	var e AppDeploymentArtifacts_SdkV2
	if o.DeploymentArtifacts.IsNull() || o.DeploymentArtifacts.IsUnknown() {
		return e, false
	}
	var v []AppDeploymentArtifacts_SdkV2
	d := o.DeploymentArtifacts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeploymentArtifacts sets the value of the DeploymentArtifacts field in AppDeployment_SdkV2.
func (o *AppDeployment_SdkV2) SetDeploymentArtifacts(ctx context.Context, v AppDeploymentArtifacts_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment_artifacts"]
	o.DeploymentArtifacts = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in AppDeployment_SdkV2 as
// a AppDeploymentStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppDeployment_SdkV2) GetStatus(ctx context.Context) (AppDeploymentStatus_SdkV2, bool) {
	var e AppDeploymentStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []AppDeploymentStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in AppDeployment_SdkV2.
func (o *AppDeployment_SdkV2) SetStatus(ctx context.Context, v AppDeploymentStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

type AppDeploymentArtifacts_SdkV2 struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	SourceCodePath types.String `tfsdk:"source_code_path"`
}

func (toState *AppDeploymentArtifacts_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppDeploymentArtifacts_SdkV2) {
}

func (toState *AppDeploymentArtifacts_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppDeploymentArtifacts_SdkV2) {
}

func (c AppDeploymentArtifacts_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["source_code_path"] = attrs["source_code_path"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppDeploymentArtifacts.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppDeploymentArtifacts_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppDeploymentArtifacts_SdkV2
// only implements ToObjectValue() and Type().
func (o AppDeploymentArtifacts_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"source_code_path": o.SourceCodePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppDeploymentArtifacts_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"source_code_path": types.StringType,
		},
	}
}

type AppDeploymentStatus_SdkV2 struct {
	// Message corresponding with the deployment state.
	Message types.String `tfsdk:"message"`
	// State of the deployment.
	State types.String `tfsdk:"state"`
}

func (toState *AppDeploymentStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppDeploymentStatus_SdkV2) {
}

func (toState *AppDeploymentStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppDeploymentStatus_SdkV2) {
}

func (c AppDeploymentStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppDeploymentStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppDeploymentStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppDeploymentStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o AppDeploymentStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"state":   o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppDeploymentStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

// App manifest definition
type AppManifest_SdkV2 struct {
	// Description of the app defined by manifest author / publisher
	Description types.String `tfsdk:"description"`
	// Name of the app defined by manifest author / publisher
	Name types.String `tfsdk:"name"`

	ResourceSpecs types.List `tfsdk:"resource_specs"`
	// The manifest schema version, for now only 1 is allowed
	Version types.Int64 `tfsdk:"version"`
}

func (toState *AppManifest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppManifest_SdkV2) {
}

func (toState *AppManifest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppManifest_SdkV2) {
}

func (c AppManifest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["resource_specs"] = attrs["resource_specs"].SetOptional()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppManifest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resource_specs": reflect.TypeOf(AppManifestAppResourceSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifest_SdkV2
// only implements ToObjectValue() and Type().
func (o AppManifest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":    o.Description,
			"name":           o.Name,
			"resource_specs": o.ResourceSpecs,
			"version":        o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppManifest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"resource_specs": basetypes.ListType{
				ElemType: AppManifestAppResourceSpec_SdkV2{}.Type(ctx),
			},
			"version": types.Int64Type,
		},
	}
}

// GetResourceSpecs returns the value of the ResourceSpecs field in AppManifest_SdkV2 as
// a slice of AppManifestAppResourceSpec_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AppManifest_SdkV2) GetResourceSpecs(ctx context.Context) ([]AppManifestAppResourceSpec_SdkV2, bool) {
	if o.ResourceSpecs.IsNull() || o.ResourceSpecs.IsUnknown() {
		return nil, false
	}
	var v []AppManifestAppResourceSpec_SdkV2
	d := o.ResourceSpecs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceSpecs sets the value of the ResourceSpecs field in AppManifest_SdkV2.
func (o *AppManifest_SdkV2) SetResourceSpecs(ctx context.Context, v []AppManifestAppResourceSpec_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_specs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ResourceSpecs = types.ListValueMust(t, vs)
}

type AppManifestAppResourceJobSpec_SdkV2 struct {
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (toState *AppManifestAppResourceJobSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppManifestAppResourceJobSpec_SdkV2) {
}

func (toState *AppManifestAppResourceJobSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppManifestAppResourceJobSpec_SdkV2) {
}

func (c AppManifestAppResourceJobSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceJobSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppManifestAppResourceJobSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceJobSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o AppManifestAppResourceJobSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppManifestAppResourceJobSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
		},
	}
}

type AppManifestAppResourceSecretSpec_SdkV2 struct {
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission types.String `tfsdk:"permission"`
}

func (toState *AppManifestAppResourceSecretSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppManifestAppResourceSecretSpec_SdkV2) {
}

func (toState *AppManifestAppResourceSecretSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppManifestAppResourceSecretSpec_SdkV2) {
}

func (c AppManifestAppResourceSecretSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceSecretSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppManifestAppResourceSecretSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceSecretSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o AppManifestAppResourceSecretSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppManifestAppResourceSecretSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
		},
	}
}

type AppManifestAppResourceServingEndpointSpec_SdkV2 struct {
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (toState *AppManifestAppResourceServingEndpointSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppManifestAppResourceServingEndpointSpec_SdkV2) {
}

func (toState *AppManifestAppResourceServingEndpointSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppManifestAppResourceServingEndpointSpec_SdkV2) {
}

func (c AppManifestAppResourceServingEndpointSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceServingEndpointSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppManifestAppResourceServingEndpointSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceServingEndpointSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o AppManifestAppResourceServingEndpointSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppManifestAppResourceServingEndpointSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
		},
	}
}

// AppResource related fields are copied from app.proto but excludes resource
// identifiers (e.g. name, id, key, scope, etc.)
type AppManifestAppResourceSpec_SdkV2 struct {
	// Description of the App Resource.
	Description types.String `tfsdk:"description"`

	JobSpec types.List `tfsdk:"job_spec"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name"`

	SecretSpec types.List `tfsdk:"secret_spec"`

	ServingEndpointSpec types.List `tfsdk:"serving_endpoint_spec"`

	SqlWarehouseSpec types.List `tfsdk:"sql_warehouse_spec"`

	UcSecurableSpec types.List `tfsdk:"uc_securable_spec"`
}

func (toState *AppManifestAppResourceSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppManifestAppResourceSpec_SdkV2) {
	if !fromPlan.JobSpec.IsNull() && !fromPlan.JobSpec.IsUnknown() {
		if toStateJobSpec, ok := toState.GetJobSpec(ctx); ok {
			if fromPlanJobSpec, ok := fromPlan.GetJobSpec(ctx); ok {
				toStateJobSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanJobSpec)
				toState.SetJobSpec(ctx, toStateJobSpec)
			}
		}
	}
	if !fromPlan.SecretSpec.IsNull() && !fromPlan.SecretSpec.IsUnknown() {
		if toStateSecretSpec, ok := toState.GetSecretSpec(ctx); ok {
			if fromPlanSecretSpec, ok := fromPlan.GetSecretSpec(ctx); ok {
				toStateSecretSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSecretSpec)
				toState.SetSecretSpec(ctx, toStateSecretSpec)
			}
		}
	}
	if !fromPlan.ServingEndpointSpec.IsNull() && !fromPlan.ServingEndpointSpec.IsUnknown() {
		if toStateServingEndpointSpec, ok := toState.GetServingEndpointSpec(ctx); ok {
			if fromPlanServingEndpointSpec, ok := fromPlan.GetServingEndpointSpec(ctx); ok {
				toStateServingEndpointSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanServingEndpointSpec)
				toState.SetServingEndpointSpec(ctx, toStateServingEndpointSpec)
			}
		}
	}
	if !fromPlan.SqlWarehouseSpec.IsNull() && !fromPlan.SqlWarehouseSpec.IsUnknown() {
		if toStateSqlWarehouseSpec, ok := toState.GetSqlWarehouseSpec(ctx); ok {
			if fromPlanSqlWarehouseSpec, ok := fromPlan.GetSqlWarehouseSpec(ctx); ok {
				toStateSqlWarehouseSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSqlWarehouseSpec)
				toState.SetSqlWarehouseSpec(ctx, toStateSqlWarehouseSpec)
			}
		}
	}
	if !fromPlan.UcSecurableSpec.IsNull() && !fromPlan.UcSecurableSpec.IsUnknown() {
		if toStateUcSecurableSpec, ok := toState.GetUcSecurableSpec(ctx); ok {
			if fromPlanUcSecurableSpec, ok := fromPlan.GetUcSecurableSpec(ctx); ok {
				toStateUcSecurableSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanUcSecurableSpec)
				toState.SetUcSecurableSpec(ctx, toStateUcSecurableSpec)
			}
		}
	}
}

func (toState *AppManifestAppResourceSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppManifestAppResourceSpec_SdkV2) {
	if !fromState.JobSpec.IsNull() && !fromState.JobSpec.IsUnknown() {
		if toStateJobSpec, ok := toState.GetJobSpec(ctx); ok {
			if fromStateJobSpec, ok := fromState.GetJobSpec(ctx); ok {
				toStateJobSpec.SyncFieldsDuringRead(ctx, fromStateJobSpec)
				toState.SetJobSpec(ctx, toStateJobSpec)
			}
		}
	}
	if !fromState.SecretSpec.IsNull() && !fromState.SecretSpec.IsUnknown() {
		if toStateSecretSpec, ok := toState.GetSecretSpec(ctx); ok {
			if fromStateSecretSpec, ok := fromState.GetSecretSpec(ctx); ok {
				toStateSecretSpec.SyncFieldsDuringRead(ctx, fromStateSecretSpec)
				toState.SetSecretSpec(ctx, toStateSecretSpec)
			}
		}
	}
	if !fromState.ServingEndpointSpec.IsNull() && !fromState.ServingEndpointSpec.IsUnknown() {
		if toStateServingEndpointSpec, ok := toState.GetServingEndpointSpec(ctx); ok {
			if fromStateServingEndpointSpec, ok := fromState.GetServingEndpointSpec(ctx); ok {
				toStateServingEndpointSpec.SyncFieldsDuringRead(ctx, fromStateServingEndpointSpec)
				toState.SetServingEndpointSpec(ctx, toStateServingEndpointSpec)
			}
		}
	}
	if !fromState.SqlWarehouseSpec.IsNull() && !fromState.SqlWarehouseSpec.IsUnknown() {
		if toStateSqlWarehouseSpec, ok := toState.GetSqlWarehouseSpec(ctx); ok {
			if fromStateSqlWarehouseSpec, ok := fromState.GetSqlWarehouseSpec(ctx); ok {
				toStateSqlWarehouseSpec.SyncFieldsDuringRead(ctx, fromStateSqlWarehouseSpec)
				toState.SetSqlWarehouseSpec(ctx, toStateSqlWarehouseSpec)
			}
		}
	}
	if !fromState.UcSecurableSpec.IsNull() && !fromState.UcSecurableSpec.IsUnknown() {
		if toStateUcSecurableSpec, ok := toState.GetUcSecurableSpec(ctx); ok {
			if fromStateUcSecurableSpec, ok := fromState.GetUcSecurableSpec(ctx); ok {
				toStateUcSecurableSpec.SyncFieldsDuringRead(ctx, fromStateUcSecurableSpec)
				toState.SetUcSecurableSpec(ctx, toStateUcSecurableSpec)
			}
		}
	}
}

func (c AppManifestAppResourceSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["job_spec"] = attrs["job_spec"].SetOptional()
	attrs["job_spec"] = attrs["job_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["secret_spec"] = attrs["secret_spec"].SetOptional()
	attrs["secret_spec"] = attrs["secret_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["serving_endpoint_spec"] = attrs["serving_endpoint_spec"].SetOptional()
	attrs["serving_endpoint_spec"] = attrs["serving_endpoint_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["sql_warehouse_spec"] = attrs["sql_warehouse_spec"].SetOptional()
	attrs["sql_warehouse_spec"] = attrs["sql_warehouse_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["uc_securable_spec"] = attrs["uc_securable_spec"].SetOptional()
	attrs["uc_securable_spec"] = attrs["uc_securable_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppManifestAppResourceSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"job_spec":              reflect.TypeOf(AppManifestAppResourceJobSpec_SdkV2{}),
		"secret_spec":           reflect.TypeOf(AppManifestAppResourceSecretSpec_SdkV2{}),
		"serving_endpoint_spec": reflect.TypeOf(AppManifestAppResourceServingEndpointSpec_SdkV2{}),
		"sql_warehouse_spec":    reflect.TypeOf(AppManifestAppResourceSqlWarehouseSpec_SdkV2{}),
		"uc_securable_spec":     reflect.TypeOf(AppManifestAppResourceUcSecurableSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o AppManifestAppResourceSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":           o.Description,
			"job_spec":              o.JobSpec,
			"name":                  o.Name,
			"secret_spec":           o.SecretSpec,
			"serving_endpoint_spec": o.ServingEndpointSpec,
			"sql_warehouse_spec":    o.SqlWarehouseSpec,
			"uc_securable_spec":     o.UcSecurableSpec,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppManifestAppResourceSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"job_spec": basetypes.ListType{
				ElemType: AppManifestAppResourceJobSpec_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"secret_spec": basetypes.ListType{
				ElemType: AppManifestAppResourceSecretSpec_SdkV2{}.Type(ctx),
			},
			"serving_endpoint_spec": basetypes.ListType{
				ElemType: AppManifestAppResourceServingEndpointSpec_SdkV2{}.Type(ctx),
			},
			"sql_warehouse_spec": basetypes.ListType{
				ElemType: AppManifestAppResourceSqlWarehouseSpec_SdkV2{}.Type(ctx),
			},
			"uc_securable_spec": basetypes.ListType{
				ElemType: AppManifestAppResourceUcSecurableSpec_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetJobSpec returns the value of the JobSpec field in AppManifestAppResourceSpec_SdkV2 as
// a AppManifestAppResourceJobSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppManifestAppResourceSpec_SdkV2) GetJobSpec(ctx context.Context) (AppManifestAppResourceJobSpec_SdkV2, bool) {
	var e AppManifestAppResourceJobSpec_SdkV2
	if o.JobSpec.IsNull() || o.JobSpec.IsUnknown() {
		return e, false
	}
	var v []AppManifestAppResourceJobSpec_SdkV2
	d := o.JobSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in AppManifestAppResourceSpec_SdkV2.
func (o *AppManifestAppResourceSpec_SdkV2) SetJobSpec(ctx context.Context, v AppManifestAppResourceJobSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_spec"]
	o.JobSpec = types.ListValueMust(t, vs)
}

// GetSecretSpec returns the value of the SecretSpec field in AppManifestAppResourceSpec_SdkV2 as
// a AppManifestAppResourceSecretSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppManifestAppResourceSpec_SdkV2) GetSecretSpec(ctx context.Context) (AppManifestAppResourceSecretSpec_SdkV2, bool) {
	var e AppManifestAppResourceSecretSpec_SdkV2
	if o.SecretSpec.IsNull() || o.SecretSpec.IsUnknown() {
		return e, false
	}
	var v []AppManifestAppResourceSecretSpec_SdkV2
	d := o.SecretSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSecretSpec sets the value of the SecretSpec field in AppManifestAppResourceSpec_SdkV2.
func (o *AppManifestAppResourceSpec_SdkV2) SetSecretSpec(ctx context.Context, v AppManifestAppResourceSecretSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["secret_spec"]
	o.SecretSpec = types.ListValueMust(t, vs)
}

// GetServingEndpointSpec returns the value of the ServingEndpointSpec field in AppManifestAppResourceSpec_SdkV2 as
// a AppManifestAppResourceServingEndpointSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppManifestAppResourceSpec_SdkV2) GetServingEndpointSpec(ctx context.Context) (AppManifestAppResourceServingEndpointSpec_SdkV2, bool) {
	var e AppManifestAppResourceServingEndpointSpec_SdkV2
	if o.ServingEndpointSpec.IsNull() || o.ServingEndpointSpec.IsUnknown() {
		return e, false
	}
	var v []AppManifestAppResourceServingEndpointSpec_SdkV2
	d := o.ServingEndpointSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServingEndpointSpec sets the value of the ServingEndpointSpec field in AppManifestAppResourceSpec_SdkV2.
func (o *AppManifestAppResourceSpec_SdkV2) SetServingEndpointSpec(ctx context.Context, v AppManifestAppResourceServingEndpointSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["serving_endpoint_spec"]
	o.ServingEndpointSpec = types.ListValueMust(t, vs)
}

// GetSqlWarehouseSpec returns the value of the SqlWarehouseSpec field in AppManifestAppResourceSpec_SdkV2 as
// a AppManifestAppResourceSqlWarehouseSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppManifestAppResourceSpec_SdkV2) GetSqlWarehouseSpec(ctx context.Context) (AppManifestAppResourceSqlWarehouseSpec_SdkV2, bool) {
	var e AppManifestAppResourceSqlWarehouseSpec_SdkV2
	if o.SqlWarehouseSpec.IsNull() || o.SqlWarehouseSpec.IsUnknown() {
		return e, false
	}
	var v []AppManifestAppResourceSqlWarehouseSpec_SdkV2
	d := o.SqlWarehouseSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlWarehouseSpec sets the value of the SqlWarehouseSpec field in AppManifestAppResourceSpec_SdkV2.
func (o *AppManifestAppResourceSpec_SdkV2) SetSqlWarehouseSpec(ctx context.Context, v AppManifestAppResourceSqlWarehouseSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_warehouse_spec"]
	o.SqlWarehouseSpec = types.ListValueMust(t, vs)
}

// GetUcSecurableSpec returns the value of the UcSecurableSpec field in AppManifestAppResourceSpec_SdkV2 as
// a AppManifestAppResourceUcSecurableSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppManifestAppResourceSpec_SdkV2) GetUcSecurableSpec(ctx context.Context) (AppManifestAppResourceUcSecurableSpec_SdkV2, bool) {
	var e AppManifestAppResourceUcSecurableSpec_SdkV2
	if o.UcSecurableSpec.IsNull() || o.UcSecurableSpec.IsUnknown() {
		return e, false
	}
	var v []AppManifestAppResourceUcSecurableSpec_SdkV2
	d := o.UcSecurableSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUcSecurableSpec sets the value of the UcSecurableSpec field in AppManifestAppResourceSpec_SdkV2.
func (o *AppManifestAppResourceSpec_SdkV2) SetUcSecurableSpec(ctx context.Context, v AppManifestAppResourceUcSecurableSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["uc_securable_spec"]
	o.UcSecurableSpec = types.ListValueMust(t, vs)
}

type AppManifestAppResourceSqlWarehouseSpec_SdkV2 struct {
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission types.String `tfsdk:"permission"`
}

func (toState *AppManifestAppResourceSqlWarehouseSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppManifestAppResourceSqlWarehouseSpec_SdkV2) {
}

func (toState *AppManifestAppResourceSqlWarehouseSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppManifestAppResourceSqlWarehouseSpec_SdkV2) {
}

func (c AppManifestAppResourceSqlWarehouseSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceSqlWarehouseSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppManifestAppResourceSqlWarehouseSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceSqlWarehouseSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o AppManifestAppResourceSqlWarehouseSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppManifestAppResourceSqlWarehouseSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
		},
	}
}

type AppManifestAppResourceUcSecurableSpec_SdkV2 struct {
	Permission types.String `tfsdk:"permission"`

	SecurableType types.String `tfsdk:"securable_type"`
}

func (toState *AppManifestAppResourceUcSecurableSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppManifestAppResourceUcSecurableSpec_SdkV2) {
}

func (toState *AppManifestAppResourceUcSecurableSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppManifestAppResourceUcSecurableSpec_SdkV2) {
}

func (c AppManifestAppResourceUcSecurableSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()
	attrs["securable_type"] = attrs["securable_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceUcSecurableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppManifestAppResourceUcSecurableSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceUcSecurableSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o AppManifestAppResourceUcSecurableSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission":     o.Permission,
			"securable_type": o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppManifestAppResourceUcSecurableSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission":     types.StringType,
			"securable_type": types.StringType,
		},
	}
}

type AppPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *AppPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppPermission_SdkV2) {
}

func (toState *AppPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppPermission_SdkV2) {
}

func (c AppPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o AppPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in AppPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AppPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in AppPermission_SdkV2.
func (o *AppPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type AppPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (toState *AppPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppPermissions_SdkV2) {
}

func (toState *AppPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppPermissions_SdkV2) {
}

func (c AppPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AppAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o AppPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AppAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in AppPermissions_SdkV2 as
// a slice of AppAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AppPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]AppAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AppAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in AppPermissions_SdkV2.
func (o *AppPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []AppAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type AppPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *AppPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppPermissionsDescription_SdkV2) {
}

func (toState *AppPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppPermissionsDescription_SdkV2) {
}

func (c AppPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o AppPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type AppPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AppAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o AppPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"app_name":            o.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AppAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"app_name": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in AppPermissionsRequest_SdkV2 as
// a slice of AppAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AppPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]AppAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AppAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in AppPermissionsRequest_SdkV2.
func (o *AppPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []AppAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type AppResource_SdkV2 struct {
	Database types.List `tfsdk:"database"`
	// Description of the App Resource.
	Description types.String `tfsdk:"description"`

	Job types.List `tfsdk:"job"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name"`

	Secret types.List `tfsdk:"secret"`

	ServingEndpoint types.List `tfsdk:"serving_endpoint"`

	SqlWarehouse types.List `tfsdk:"sql_warehouse"`

	UcSecurable types.List `tfsdk:"uc_securable"`
}

func (toState *AppResource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppResource_SdkV2) {
	if !fromPlan.Database.IsNull() && !fromPlan.Database.IsUnknown() {
		if toStateDatabase, ok := toState.GetDatabase(ctx); ok {
			if fromPlanDatabase, ok := fromPlan.GetDatabase(ctx); ok {
				toStateDatabase.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanDatabase)
				toState.SetDatabase(ctx, toStateDatabase)
			}
		}
	}
	if !fromPlan.Job.IsNull() && !fromPlan.Job.IsUnknown() {
		if toStateJob, ok := toState.GetJob(ctx); ok {
			if fromPlanJob, ok := fromPlan.GetJob(ctx); ok {
				toStateJob.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanJob)
				toState.SetJob(ctx, toStateJob)
			}
		}
	}
	if !fromPlan.Secret.IsNull() && !fromPlan.Secret.IsUnknown() {
		if toStateSecret, ok := toState.GetSecret(ctx); ok {
			if fromPlanSecret, ok := fromPlan.GetSecret(ctx); ok {
				toStateSecret.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSecret)
				toState.SetSecret(ctx, toStateSecret)
			}
		}
	}
	if !fromPlan.ServingEndpoint.IsNull() && !fromPlan.ServingEndpoint.IsUnknown() {
		if toStateServingEndpoint, ok := toState.GetServingEndpoint(ctx); ok {
			if fromPlanServingEndpoint, ok := fromPlan.GetServingEndpoint(ctx); ok {
				toStateServingEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanServingEndpoint)
				toState.SetServingEndpoint(ctx, toStateServingEndpoint)
			}
		}
	}
	if !fromPlan.SqlWarehouse.IsNull() && !fromPlan.SqlWarehouse.IsUnknown() {
		if toStateSqlWarehouse, ok := toState.GetSqlWarehouse(ctx); ok {
			if fromPlanSqlWarehouse, ok := fromPlan.GetSqlWarehouse(ctx); ok {
				toStateSqlWarehouse.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSqlWarehouse)
				toState.SetSqlWarehouse(ctx, toStateSqlWarehouse)
			}
		}
	}
	if !fromPlan.UcSecurable.IsNull() && !fromPlan.UcSecurable.IsUnknown() {
		if toStateUcSecurable, ok := toState.GetUcSecurable(ctx); ok {
			if fromPlanUcSecurable, ok := fromPlan.GetUcSecurable(ctx); ok {
				toStateUcSecurable.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanUcSecurable)
				toState.SetUcSecurable(ctx, toStateUcSecurable)
			}
		}
	}
}

func (toState *AppResource_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppResource_SdkV2) {
	if !fromState.Database.IsNull() && !fromState.Database.IsUnknown() {
		if toStateDatabase, ok := toState.GetDatabase(ctx); ok {
			if fromStateDatabase, ok := fromState.GetDatabase(ctx); ok {
				toStateDatabase.SyncFieldsDuringRead(ctx, fromStateDatabase)
				toState.SetDatabase(ctx, toStateDatabase)
			}
		}
	}
	if !fromState.Job.IsNull() && !fromState.Job.IsUnknown() {
		if toStateJob, ok := toState.GetJob(ctx); ok {
			if fromStateJob, ok := fromState.GetJob(ctx); ok {
				toStateJob.SyncFieldsDuringRead(ctx, fromStateJob)
				toState.SetJob(ctx, toStateJob)
			}
		}
	}
	if !fromState.Secret.IsNull() && !fromState.Secret.IsUnknown() {
		if toStateSecret, ok := toState.GetSecret(ctx); ok {
			if fromStateSecret, ok := fromState.GetSecret(ctx); ok {
				toStateSecret.SyncFieldsDuringRead(ctx, fromStateSecret)
				toState.SetSecret(ctx, toStateSecret)
			}
		}
	}
	if !fromState.ServingEndpoint.IsNull() && !fromState.ServingEndpoint.IsUnknown() {
		if toStateServingEndpoint, ok := toState.GetServingEndpoint(ctx); ok {
			if fromStateServingEndpoint, ok := fromState.GetServingEndpoint(ctx); ok {
				toStateServingEndpoint.SyncFieldsDuringRead(ctx, fromStateServingEndpoint)
				toState.SetServingEndpoint(ctx, toStateServingEndpoint)
			}
		}
	}
	if !fromState.SqlWarehouse.IsNull() && !fromState.SqlWarehouse.IsUnknown() {
		if toStateSqlWarehouse, ok := toState.GetSqlWarehouse(ctx); ok {
			if fromStateSqlWarehouse, ok := fromState.GetSqlWarehouse(ctx); ok {
				toStateSqlWarehouse.SyncFieldsDuringRead(ctx, fromStateSqlWarehouse)
				toState.SetSqlWarehouse(ctx, toStateSqlWarehouse)
			}
		}
	}
	if !fromState.UcSecurable.IsNull() && !fromState.UcSecurable.IsUnknown() {
		if toStateUcSecurable, ok := toState.GetUcSecurable(ctx); ok {
			if fromStateUcSecurable, ok := fromState.GetUcSecurable(ctx); ok {
				toStateUcSecurable.SyncFieldsDuringRead(ctx, fromStateUcSecurable)
				toState.SetUcSecurable(ctx, toStateUcSecurable)
			}
		}
	}
}

func (c AppResource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database"] = attrs["database"].SetOptional()
	attrs["database"] = attrs["database"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetOptional()
	attrs["job"] = attrs["job"].SetOptional()
	attrs["job"] = attrs["job"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["secret"] = attrs["secret"].SetOptional()
	attrs["secret"] = attrs["secret"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["serving_endpoint"] = attrs["serving_endpoint"].SetOptional()
	attrs["serving_endpoint"] = attrs["serving_endpoint"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["sql_warehouse"] = attrs["sql_warehouse"].SetOptional()
	attrs["sql_warehouse"] = attrs["sql_warehouse"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["uc_securable"] = attrs["uc_securable"].SetOptional()
	attrs["uc_securable"] = attrs["uc_securable"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppResource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database":         reflect.TypeOf(AppResourceDatabase_SdkV2{}),
		"job":              reflect.TypeOf(AppResourceJob_SdkV2{}),
		"secret":           reflect.TypeOf(AppResourceSecret_SdkV2{}),
		"serving_endpoint": reflect.TypeOf(AppResourceServingEndpoint_SdkV2{}),
		"sql_warehouse":    reflect.TypeOf(AppResourceSqlWarehouse_SdkV2{}),
		"uc_securable":     reflect.TypeOf(AppResourceUcSecurable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResource_SdkV2
// only implements ToObjectValue() and Type().
func (o AppResource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database":         o.Database,
			"description":      o.Description,
			"job":              o.Job,
			"name":             o.Name,
			"secret":           o.Secret,
			"serving_endpoint": o.ServingEndpoint,
			"sql_warehouse":    o.SqlWarehouse,
			"uc_securable":     o.UcSecurable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database": basetypes.ListType{
				ElemType: AppResourceDatabase_SdkV2{}.Type(ctx),
			},
			"description": types.StringType,
			"job": basetypes.ListType{
				ElemType: AppResourceJob_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"secret": basetypes.ListType{
				ElemType: AppResourceSecret_SdkV2{}.Type(ctx),
			},
			"serving_endpoint": basetypes.ListType{
				ElemType: AppResourceServingEndpoint_SdkV2{}.Type(ctx),
			},
			"sql_warehouse": basetypes.ListType{
				ElemType: AppResourceSqlWarehouse_SdkV2{}.Type(ctx),
			},
			"uc_securable": basetypes.ListType{
				ElemType: AppResourceUcSecurable_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDatabase returns the value of the Database field in AppResource_SdkV2 as
// a AppResourceDatabase_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource_SdkV2) GetDatabase(ctx context.Context) (AppResourceDatabase_SdkV2, bool) {
	var e AppResourceDatabase_SdkV2
	if o.Database.IsNull() || o.Database.IsUnknown() {
		return e, false
	}
	var v []AppResourceDatabase_SdkV2
	d := o.Database.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabase sets the value of the Database field in AppResource_SdkV2.
func (o *AppResource_SdkV2) SetDatabase(ctx context.Context, v AppResourceDatabase_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["database"]
	o.Database = types.ListValueMust(t, vs)
}

// GetJob returns the value of the Job field in AppResource_SdkV2 as
// a AppResourceJob_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource_SdkV2) GetJob(ctx context.Context) (AppResourceJob_SdkV2, bool) {
	var e AppResourceJob_SdkV2
	if o.Job.IsNull() || o.Job.IsUnknown() {
		return e, false
	}
	var v []AppResourceJob_SdkV2
	d := o.Job.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJob sets the value of the Job field in AppResource_SdkV2.
func (o *AppResource_SdkV2) SetJob(ctx context.Context, v AppResourceJob_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job"]
	o.Job = types.ListValueMust(t, vs)
}

// GetSecret returns the value of the Secret field in AppResource_SdkV2 as
// a AppResourceSecret_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource_SdkV2) GetSecret(ctx context.Context) (AppResourceSecret_SdkV2, bool) {
	var e AppResourceSecret_SdkV2
	if o.Secret.IsNull() || o.Secret.IsUnknown() {
		return e, false
	}
	var v []AppResourceSecret_SdkV2
	d := o.Secret.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSecret sets the value of the Secret field in AppResource_SdkV2.
func (o *AppResource_SdkV2) SetSecret(ctx context.Context, v AppResourceSecret_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["secret"]
	o.Secret = types.ListValueMust(t, vs)
}

// GetServingEndpoint returns the value of the ServingEndpoint field in AppResource_SdkV2 as
// a AppResourceServingEndpoint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource_SdkV2) GetServingEndpoint(ctx context.Context) (AppResourceServingEndpoint_SdkV2, bool) {
	var e AppResourceServingEndpoint_SdkV2
	if o.ServingEndpoint.IsNull() || o.ServingEndpoint.IsUnknown() {
		return e, false
	}
	var v []AppResourceServingEndpoint_SdkV2
	d := o.ServingEndpoint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServingEndpoint sets the value of the ServingEndpoint field in AppResource_SdkV2.
func (o *AppResource_SdkV2) SetServingEndpoint(ctx context.Context, v AppResourceServingEndpoint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["serving_endpoint"]
	o.ServingEndpoint = types.ListValueMust(t, vs)
}

// GetSqlWarehouse returns the value of the SqlWarehouse field in AppResource_SdkV2 as
// a AppResourceSqlWarehouse_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource_SdkV2) GetSqlWarehouse(ctx context.Context) (AppResourceSqlWarehouse_SdkV2, bool) {
	var e AppResourceSqlWarehouse_SdkV2
	if o.SqlWarehouse.IsNull() || o.SqlWarehouse.IsUnknown() {
		return e, false
	}
	var v []AppResourceSqlWarehouse_SdkV2
	d := o.SqlWarehouse.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlWarehouse sets the value of the SqlWarehouse field in AppResource_SdkV2.
func (o *AppResource_SdkV2) SetSqlWarehouse(ctx context.Context, v AppResourceSqlWarehouse_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_warehouse"]
	o.SqlWarehouse = types.ListValueMust(t, vs)
}

// GetUcSecurable returns the value of the UcSecurable field in AppResource_SdkV2 as
// a AppResourceUcSecurable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource_SdkV2) GetUcSecurable(ctx context.Context) (AppResourceUcSecurable_SdkV2, bool) {
	var e AppResourceUcSecurable_SdkV2
	if o.UcSecurable.IsNull() || o.UcSecurable.IsUnknown() {
		return e, false
	}
	var v []AppResourceUcSecurable_SdkV2
	d := o.UcSecurable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUcSecurable sets the value of the UcSecurable field in AppResource_SdkV2.
func (o *AppResource_SdkV2) SetUcSecurable(ctx context.Context, v AppResourceUcSecurable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["uc_securable"]
	o.UcSecurable = types.ListValueMust(t, vs)
}

type AppResourceDatabase_SdkV2 struct {
	DatabaseName types.String `tfsdk:"database_name"`

	InstanceName types.String `tfsdk:"instance_name"`

	Permission types.String `tfsdk:"permission"`
}

func (toState *AppResourceDatabase_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppResourceDatabase_SdkV2) {
}

func (toState *AppResourceDatabase_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppResourceDatabase_SdkV2) {
}

func (c AppResourceDatabase_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_name"] = attrs["database_name"].SetRequired()
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceDatabase.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppResourceDatabase_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceDatabase_SdkV2
// only implements ToObjectValue() and Type().
func (o AppResourceDatabase_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_name": o.DatabaseName,
			"instance_name": o.InstanceName,
			"permission":    o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceDatabase_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_name": types.StringType,
			"instance_name": types.StringType,
			"permission":    types.StringType,
		},
	}
}

type AppResourceJob_SdkV2 struct {
	// Id of the job to grant permission on.
	Id types.String `tfsdk:"id"`
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (toState *AppResourceJob_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppResourceJob_SdkV2) {
}

func (toState *AppResourceJob_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppResourceJob_SdkV2) {
}

func (c AppResourceJob_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppResourceJob_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceJob_SdkV2
// only implements ToObjectValue() and Type().
func (o AppResourceJob_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         o.Id,
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceJob_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"permission": types.StringType,
		},
	}
}

type AppResourceSecret_SdkV2 struct {
	// Key of the secret to grant permission on.
	Key types.String `tfsdk:"key"`
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission types.String `tfsdk:"permission"`
	// Scope of the secret to grant permission on.
	Scope types.String `tfsdk:"scope"`
}

func (toState *AppResourceSecret_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppResourceSecret_SdkV2) {
}

func (toState *AppResourceSecret_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppResourceSecret_SdkV2) {
}

func (c AppResourceSecret_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()
	attrs["scope"] = attrs["scope"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceSecret.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppResourceSecret_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceSecret_SdkV2
// only implements ToObjectValue() and Type().
func (o AppResourceSecret_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":        o.Key,
			"permission": o.Permission,
			"scope":      o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceSecret_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":        types.StringType,
			"permission": types.StringType,
			"scope":      types.StringType,
		},
	}
}

type AppResourceServingEndpoint_SdkV2 struct {
	// Name of the serving endpoint to grant permission on.
	Name types.String `tfsdk:"name"`
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (toState *AppResourceServingEndpoint_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppResourceServingEndpoint_SdkV2) {
}

func (toState *AppResourceServingEndpoint_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppResourceServingEndpoint_SdkV2) {
}

func (c AppResourceServingEndpoint_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceServingEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppResourceServingEndpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceServingEndpoint_SdkV2
// only implements ToObjectValue() and Type().
func (o AppResourceServingEndpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":       o.Name,
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceServingEndpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":       types.StringType,
			"permission": types.StringType,
		},
	}
}

type AppResourceSqlWarehouse_SdkV2 struct {
	// Id of the SQL warehouse to grant permission on.
	Id types.String `tfsdk:"id"`
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission types.String `tfsdk:"permission"`
}

func (toState *AppResourceSqlWarehouse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppResourceSqlWarehouse_SdkV2) {
}

func (toState *AppResourceSqlWarehouse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppResourceSqlWarehouse_SdkV2) {
}

func (c AppResourceSqlWarehouse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceSqlWarehouse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppResourceSqlWarehouse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceSqlWarehouse_SdkV2
// only implements ToObjectValue() and Type().
func (o AppResourceSqlWarehouse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         o.Id,
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceSqlWarehouse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"permission": types.StringType,
		},
	}
}

type AppResourceUcSecurable_SdkV2 struct {
	Permission types.String `tfsdk:"permission"`

	SecurableFullName types.String `tfsdk:"securable_full_name"`

	SecurableType types.String `tfsdk:"securable_type"`
}

func (toState *AppResourceUcSecurable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AppResourceUcSecurable_SdkV2) {
}

func (toState *AppResourceUcSecurable_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AppResourceUcSecurable_SdkV2) {
}

func (c AppResourceUcSecurable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()
	attrs["securable_full_name"] = attrs["securable_full_name"].SetRequired()
	attrs["securable_type"] = attrs["securable_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceUcSecurable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppResourceUcSecurable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceUcSecurable_SdkV2
// only implements ToObjectValue() and Type().
func (o AppResourceUcSecurable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission":          o.Permission,
			"securable_full_name": o.SecurableFullName,
			"securable_type":      o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceUcSecurable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission":          types.StringType,
			"securable_full_name": types.StringType,
			"securable_type":      types.StringType,
		},
	}
}

type ApplicationStatus_SdkV2 struct {
	// Application status message
	Message types.String `tfsdk:"message"`
	// State of the application.
	State types.String `tfsdk:"state"`
}

func (toState *ApplicationStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ApplicationStatus_SdkV2) {
}

func (toState *ApplicationStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ApplicationStatus_SdkV2) {
}

func (c ApplicationStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApplicationStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ApplicationStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApplicationStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o ApplicationStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"state":   o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ApplicationStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type ComputeStatus_SdkV2 struct {
	// Compute status message
	Message types.String `tfsdk:"message"`
	// State of the app compute.
	State types.String `tfsdk:"state"`
}

func (toState *ComputeStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ComputeStatus_SdkV2) {
}

func (toState *ComputeStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ComputeStatus_SdkV2) {
}

func (c ComputeStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComputeStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ComputeStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComputeStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o ComputeStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"state":   o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ComputeStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type CreateAppDeploymentRequest_SdkV2 struct {
	// The app deployment configuration.
	AppDeployment types.List `tfsdk:"app_deployment"`
	// The name of the app.
	AppName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAppDeploymentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAppDeploymentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app_deployment": reflect.TypeOf(AppDeployment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAppDeploymentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateAppDeploymentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_deployment": o.AppDeployment,
			"app_name":       o.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAppDeploymentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_deployment": basetypes.ListType{
				ElemType: AppDeployment_SdkV2{}.Type(ctx),
			},
			"app_name": types.StringType,
		},
	}
}

// GetAppDeployment returns the value of the AppDeployment field in CreateAppDeploymentRequest_SdkV2 as
// a AppDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAppDeploymentRequest_SdkV2) GetAppDeployment(ctx context.Context) (AppDeployment_SdkV2, bool) {
	var e AppDeployment_SdkV2
	if o.AppDeployment.IsNull() || o.AppDeployment.IsUnknown() {
		return e, false
	}
	var v []AppDeployment_SdkV2
	d := o.AppDeployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAppDeployment sets the value of the AppDeployment field in CreateAppDeploymentRequest_SdkV2.
func (o *CreateAppDeploymentRequest_SdkV2) SetAppDeployment(ctx context.Context, v AppDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["app_deployment"]
	o.AppDeployment = types.ListValueMust(t, vs)
}

type CreateAppRequest_SdkV2 struct {
	App types.List `tfsdk:"app"`
	// If true, the app will not be started after creation.
	NoCompute types.Bool `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAppRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app": reflect.TypeOf(App_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAppRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateAppRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":        o.App,
			"no_compute": o.NoCompute,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAppRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app": basetypes.ListType{
				ElemType: App_SdkV2{}.Type(ctx),
			},
			"no_compute": types.BoolType,
		},
	}
}

// GetApp returns the value of the App field in CreateAppRequest_SdkV2 as
// a App_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAppRequest_SdkV2) GetApp(ctx context.Context) (App_SdkV2, bool) {
	var e App_SdkV2
	if o.App.IsNull() || o.App.IsUnknown() {
		return e, false
	}
	var v []App_SdkV2
	d := o.App.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetApp sets the value of the App field in CreateAppRequest_SdkV2.
func (o *CreateAppRequest_SdkV2) SetApp(ctx context.Context, v App_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["app"]
	o.App = types.ListValueMust(t, vs)
}

type CreateCustomTemplateRequest_SdkV2 struct {
	Template types.List `tfsdk:"template"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomTemplateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCustomTemplateRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"template": reflect.TypeOf(CustomTemplate_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomTemplateRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCustomTemplateRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"template": o.Template,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCustomTemplateRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"template": basetypes.ListType{
				ElemType: CustomTemplate_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTemplate returns the value of the Template field in CreateCustomTemplateRequest_SdkV2 as
// a CustomTemplate_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomTemplateRequest_SdkV2) GetTemplate(ctx context.Context) (CustomTemplate_SdkV2, bool) {
	var e CustomTemplate_SdkV2
	if o.Template.IsNull() || o.Template.IsUnknown() {
		return e, false
	}
	var v []CustomTemplate_SdkV2
	d := o.Template.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTemplate sets the value of the Template field in CreateCustomTemplateRequest_SdkV2.
func (o *CreateCustomTemplateRequest_SdkV2) SetTemplate(ctx context.Context, v CustomTemplate_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["template"]
	o.Template = types.ListValueMust(t, vs)
}

type CustomTemplate_SdkV2 struct {
	Creator types.String `tfsdk:"creator"`
	// The description of the template.
	Description types.String `tfsdk:"description"`
	// The Git provider of the template.
	GitProvider types.String `tfsdk:"git_provider"`
	// The Git repository URL that the template resides in.
	GitRepo types.String `tfsdk:"git_repo"`
	// The manifest of the template. It defines fields and default values when
	// installing the template.
	Manifest types.List `tfsdk:"manifest"`
	// The name of the template. It must contain only alphanumeric characters,
	// hyphens, underscores, and whitespaces. It must be unique within the
	// workspace.
	Name types.String `tfsdk:"name"`
	// The path to the template within the Git repository.
	Path types.String `tfsdk:"path"`
}

func (toState *CustomTemplate_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CustomTemplate_SdkV2) {
	if !fromPlan.Manifest.IsNull() && !fromPlan.Manifest.IsUnknown() {
		if toStateManifest, ok := toState.GetManifest(ctx); ok {
			if fromPlanManifest, ok := fromPlan.GetManifest(ctx); ok {
				toStateManifest.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanManifest)
				toState.SetManifest(ctx, toStateManifest)
			}
		}
	}
}

func (toState *CustomTemplate_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CustomTemplate_SdkV2) {
	if !fromState.Manifest.IsNull() && !fromState.Manifest.IsUnknown() {
		if toStateManifest, ok := toState.GetManifest(ctx); ok {
			if fromStateManifest, ok := fromState.GetManifest(ctx); ok {
				toStateManifest.SyncFieldsDuringRead(ctx, fromStateManifest)
				toState.SetManifest(ctx, toStateManifest)
			}
		}
	}
}

func (c CustomTemplate_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["git_provider"] = attrs["git_provider"].SetRequired()
	attrs["git_repo"] = attrs["git_repo"].SetRequired()
	attrs["manifest"] = attrs["manifest"].SetRequired()
	attrs["manifest"] = attrs["manifest"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomTemplate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CustomTemplate_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(AppManifest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomTemplate_SdkV2
// only implements ToObjectValue() and Type().
func (o CustomTemplate_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator":      o.Creator,
			"description":  o.Description,
			"git_provider": o.GitProvider,
			"git_repo":     o.GitRepo,
			"manifest":     o.Manifest,
			"name":         o.Name,
			"path":         o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomTemplate_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator":      types.StringType,
			"description":  types.StringType,
			"git_provider": types.StringType,
			"git_repo":     types.StringType,
			"manifest": basetypes.ListType{
				ElemType: AppManifest_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"path": types.StringType,
		},
	}
}

// GetManifest returns the value of the Manifest field in CustomTemplate_SdkV2 as
// a AppManifest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomTemplate_SdkV2) GetManifest(ctx context.Context) (AppManifest_SdkV2, bool) {
	var e AppManifest_SdkV2
	if o.Manifest.IsNull() || o.Manifest.IsUnknown() {
		return e, false
	}
	var v []AppManifest_SdkV2
	d := o.Manifest.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManifest sets the value of the Manifest field in CustomTemplate_SdkV2.
func (o *CustomTemplate_SdkV2) SetManifest(ctx context.Context, v AppManifest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["manifest"]
	o.Manifest = types.ListValueMust(t, vs)
}

type DeleteAppRequest_SdkV2 struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAppRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAppRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAppRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAppRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteCustomTemplateRequest_SdkV2 struct {
	// The name of the custom template.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCustomTemplateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCustomTemplateRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomTemplateRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCustomTemplateRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCustomTemplateRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetAppDeploymentRequest_SdkV2 struct {
	// The name of the app.
	AppName types.String `tfsdk:"-"`
	// The unique id of the deployment.
	DeploymentId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppDeploymentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAppDeploymentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppDeploymentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAppDeploymentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name":      o.AppName,
			"deployment_id": o.DeploymentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppDeploymentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name":      types.StringType,
			"deployment_id": types.StringType,
		},
	}
}

type GetAppPermissionLevelsRequest_SdkV2 struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAppPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAppPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name": o.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name": types.StringType,
		},
	}
}

type GetAppPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (toState *GetAppPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetAppPermissionLevelsResponse_SdkV2) {
}

func (toState *GetAppPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetAppPermissionLevelsResponse_SdkV2) {
}

func (c GetAppPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAppPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(AppPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAppPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: AppPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetAppPermissionLevelsResponse_SdkV2 as
// a slice of AppPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetAppPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]AppPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []AppPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetAppPermissionLevelsResponse_SdkV2.
func (o *GetAppPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []AppPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetAppPermissionsRequest_SdkV2 struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAppPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAppPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name": o.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name": types.StringType,
		},
	}
}

type GetAppRequest_SdkV2 struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAppRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAppRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetCustomTemplateRequest_SdkV2 struct {
	// The name of the custom template.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCustomTemplateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCustomTemplateRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomTemplateRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCustomTemplateRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCustomTemplateRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListAppDeploymentsRequest_SdkV2 struct {
	// The name of the app.
	AppName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAppDeploymentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAppDeploymentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppDeploymentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAppDeploymentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name":   o.AppName,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAppDeploymentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name":   types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAppDeploymentsResponse_SdkV2 struct {
	// Deployment history of the app.
	AppDeployments types.List `tfsdk:"app_deployments"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *ListAppDeploymentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListAppDeploymentsResponse_SdkV2) {
}

func (toState *ListAppDeploymentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListAppDeploymentsResponse_SdkV2) {
}

func (c ListAppDeploymentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_deployments"] = attrs["app_deployments"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAppDeploymentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAppDeploymentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app_deployments": reflect.TypeOf(AppDeployment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppDeploymentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAppDeploymentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_deployments": o.AppDeployments,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAppDeploymentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_deployments": basetypes.ListType{
				ElemType: AppDeployment_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetAppDeployments returns the value of the AppDeployments field in ListAppDeploymentsResponse_SdkV2 as
// a slice of AppDeployment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAppDeploymentsResponse_SdkV2) GetAppDeployments(ctx context.Context) ([]AppDeployment_SdkV2, bool) {
	if o.AppDeployments.IsNull() || o.AppDeployments.IsUnknown() {
		return nil, false
	}
	var v []AppDeployment_SdkV2
	d := o.AppDeployments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAppDeployments sets the value of the AppDeployments field in ListAppDeploymentsResponse_SdkV2.
func (o *ListAppDeploymentsResponse_SdkV2) SetAppDeployments(ctx context.Context, v []AppDeployment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["app_deployments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AppDeployments = types.ListValueMust(t, vs)
}

type ListAppsRequest_SdkV2 struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAppsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAppsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAppsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAppsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAppsResponse_SdkV2 struct {
	Apps types.List `tfsdk:"apps"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *ListAppsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListAppsResponse_SdkV2) {
}

func (toState *ListAppsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListAppsResponse_SdkV2) {
}

func (c ListAppsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apps"] = attrs["apps"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAppsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAppsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(App_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAppsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            o.Apps,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAppsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: App_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetApps returns the value of the Apps field in ListAppsResponse_SdkV2 as
// a slice of App_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAppsResponse_SdkV2) GetApps(ctx context.Context) ([]App_SdkV2, bool) {
	if o.Apps.IsNull() || o.Apps.IsUnknown() {
		return nil, false
	}
	var v []App_SdkV2
	d := o.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in ListAppsResponse_SdkV2.
func (o *ListAppsResponse_SdkV2) SetApps(ctx context.Context, v []App_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Apps = types.ListValueMust(t, vs)
}

type ListCustomTemplatesRequest_SdkV2 struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of custom templates. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCustomTemplatesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCustomTemplatesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCustomTemplatesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCustomTemplatesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCustomTemplatesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListCustomTemplatesResponse_SdkV2 struct {
	// Pagination token to request the next page of custom templates.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Templates types.List `tfsdk:"templates"`
}

func (toState *ListCustomTemplatesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCustomTemplatesResponse_SdkV2) {
}

func (toState *ListCustomTemplatesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCustomTemplatesResponse_SdkV2) {
}

func (c ListCustomTemplatesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["templates"] = attrs["templates"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCustomTemplatesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCustomTemplatesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"templates": reflect.TypeOf(CustomTemplate_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCustomTemplatesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCustomTemplatesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"templates":       o.Templates,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCustomTemplatesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"templates": basetypes.ListType{
				ElemType: CustomTemplate_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTemplates returns the value of the Templates field in ListCustomTemplatesResponse_SdkV2 as
// a slice of CustomTemplate_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCustomTemplatesResponse_SdkV2) GetTemplates(ctx context.Context) ([]CustomTemplate_SdkV2, bool) {
	if o.Templates.IsNull() || o.Templates.IsUnknown() {
		return nil, false
	}
	var v []CustomTemplate_SdkV2
	d := o.Templates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTemplates sets the value of the Templates field in ListCustomTemplatesResponse_SdkV2.
func (o *ListCustomTemplatesResponse_SdkV2) SetTemplates(ctx context.Context, v []CustomTemplate_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["templates"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Templates = types.ListValueMust(t, vs)
}

type StartAppRequest_SdkV2 struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartAppRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartAppRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o StartAppRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartAppRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type StopAppRequest_SdkV2 struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopAppRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopAppRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o StopAppRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StopAppRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UpdateAppRequest_SdkV2 struct {
	App types.List `tfsdk:"app"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAppRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app": reflect.TypeOf(App_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAppRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateAppRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":  o.App,
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAppRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app": basetypes.ListType{
				ElemType: App_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetApp returns the value of the App field in UpdateAppRequest_SdkV2 as
// a App_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateAppRequest_SdkV2) GetApp(ctx context.Context) (App_SdkV2, bool) {
	var e App_SdkV2
	if o.App.IsNull() || o.App.IsUnknown() {
		return e, false
	}
	var v []App_SdkV2
	d := o.App.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetApp sets the value of the App field in UpdateAppRequest_SdkV2.
func (o *UpdateAppRequest_SdkV2) SetApp(ctx context.Context, v App_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["app"]
	o.App = types.ListValueMust(t, vs)
}

type UpdateCustomTemplateRequest_SdkV2 struct {
	// The name of the template. It must contain only alphanumeric characters,
	// hyphens, underscores, and whitespaces. It must be unique within the
	// workspace.
	Name types.String `tfsdk:"-"`

	Template types.List `tfsdk:"template"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCustomTemplateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCustomTemplateRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"template": reflect.TypeOf(CustomTemplate_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomTemplateRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCustomTemplateRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     o.Name,
			"template": o.Template,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCustomTemplateRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"template": basetypes.ListType{
				ElemType: CustomTemplate_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTemplate returns the value of the Template field in UpdateCustomTemplateRequest_SdkV2 as
// a CustomTemplate_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCustomTemplateRequest_SdkV2) GetTemplate(ctx context.Context) (CustomTemplate_SdkV2, bool) {
	var e CustomTemplate_SdkV2
	if o.Template.IsNull() || o.Template.IsUnknown() {
		return e, false
	}
	var v []CustomTemplate_SdkV2
	d := o.Template.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTemplate sets the value of the Template field in UpdateCustomTemplateRequest_SdkV2.
func (o *UpdateCustomTemplateRequest_SdkV2) SetTemplate(ctx context.Context, v CustomTemplate_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["template"]
	o.Template = types.ListValueMust(t, vs)
}
