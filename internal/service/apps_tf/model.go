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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment types.Object `tfsdk:"active_deployment"`

	AppStatus types.Object `tfsdk:"app_status"`

	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`

	ComputeStatus types.Object `tfsdk:"compute_status"`
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
	PendingDeployment types.Object `tfsdk:"pending_deployment"`
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

func (newState *App) SyncEffectiveFieldsDuringCreateOrUpdate(plan App) {
}

func (newState *App) SyncEffectiveFieldsDuringRead(existingState App) {
}

func (c App) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["active_deployment"] = attrs["active_deployment"].SetComputed()
	attrs["app_status"] = attrs["app_status"].SetComputed()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["compute_status"] = attrs["compute_status"].SetComputed()
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
func (a App) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"active_deployment":         reflect.TypeOf(AppDeployment{}),
		"app_status":                reflect.TypeOf(ApplicationStatus{}),
		"compute_status":            reflect.TypeOf(ComputeStatus{}),
		"effective_user_api_scopes": reflect.TypeOf(types.String{}),
		"pending_deployment":        reflect.TypeOf(AppDeployment{}),
		"resources":                 reflect.TypeOf(AppResource{}),
		"user_api_scopes":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, App
// only implements ToObjectValue() and Type().
func (o App) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o App) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active_deployment":          AppDeployment{}.Type(ctx),
			"app_status":                 ApplicationStatus{}.Type(ctx),
			"budget_policy_id":           types.StringType,
			"compute_status":             ComputeStatus{}.Type(ctx),
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
			"pending_deployment":        AppDeployment{}.Type(ctx),
			"resources": basetypes.ListType{
				ElemType: AppResource{}.Type(ctx),
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

// GetActiveDeployment returns the value of the ActiveDeployment field in App as
// a AppDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (o *App) GetActiveDeployment(ctx context.Context) (AppDeployment, bool) {
	var e AppDeployment
	if o.ActiveDeployment.IsNull() || o.ActiveDeployment.IsUnknown() {
		return e, false
	}
	var v []AppDeployment
	d := o.ActiveDeployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActiveDeployment sets the value of the ActiveDeployment field in App.
func (o *App) SetActiveDeployment(ctx context.Context, v AppDeployment) {
	vs := v.ToObjectValue(ctx)
	o.ActiveDeployment = vs
}

// GetAppStatus returns the value of the AppStatus field in App as
// a ApplicationStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *App) GetAppStatus(ctx context.Context) (ApplicationStatus, bool) {
	var e ApplicationStatus
	if o.AppStatus.IsNull() || o.AppStatus.IsUnknown() {
		return e, false
	}
	var v []ApplicationStatus
	d := o.AppStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAppStatus sets the value of the AppStatus field in App.
func (o *App) SetAppStatus(ctx context.Context, v ApplicationStatus) {
	vs := v.ToObjectValue(ctx)
	o.AppStatus = vs
}

// GetComputeStatus returns the value of the ComputeStatus field in App as
// a ComputeStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *App) GetComputeStatus(ctx context.Context) (ComputeStatus, bool) {
	var e ComputeStatus
	if o.ComputeStatus.IsNull() || o.ComputeStatus.IsUnknown() {
		return e, false
	}
	var v []ComputeStatus
	d := o.ComputeStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComputeStatus sets the value of the ComputeStatus field in App.
func (o *App) SetComputeStatus(ctx context.Context, v ComputeStatus) {
	vs := v.ToObjectValue(ctx)
	o.ComputeStatus = vs
}

// GetEffectiveUserApiScopes returns the value of the EffectiveUserApiScopes field in App as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *App) GetEffectiveUserApiScopes(ctx context.Context) ([]types.String, bool) {
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

// SetEffectiveUserApiScopes sets the value of the EffectiveUserApiScopes field in App.
func (o *App) SetEffectiveUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EffectiveUserApiScopes = types.ListValueMust(t, vs)
}

// GetPendingDeployment returns the value of the PendingDeployment field in App as
// a AppDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (o *App) GetPendingDeployment(ctx context.Context) (AppDeployment, bool) {
	var e AppDeployment
	if o.PendingDeployment.IsNull() || o.PendingDeployment.IsUnknown() {
		return e, false
	}
	var v []AppDeployment
	d := o.PendingDeployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPendingDeployment sets the value of the PendingDeployment field in App.
func (o *App) SetPendingDeployment(ctx context.Context, v AppDeployment) {
	vs := v.ToObjectValue(ctx)
	o.PendingDeployment = vs
}

// GetResources returns the value of the Resources field in App as
// a slice of AppResource values.
// If the field is unknown or null, the boolean return value is false.
func (o *App) GetResources(ctx context.Context) ([]AppResource, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []AppResource
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in App.
func (o *App) SetResources(ctx context.Context, v []AppResource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// GetUserApiScopes returns the value of the UserApiScopes field in App as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *App) GetUserApiScopes(ctx context.Context) ([]types.String, bool) {
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

// SetUserApiScopes sets the value of the UserApiScopes field in App.
func (o *App) SetUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UserApiScopes = types.ListValueMust(t, vs)
}

type AppAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *AppAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppAccessControlRequest) {
}

func (newState *AppAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState AppAccessControlRequest) {
}

func (c AppAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppAccessControlRequest
// only implements ToObjectValue() and Type().
func (o AppAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AppAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AppAccessControlResponse struct {
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

func (newState *AppAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppAccessControlResponse) {
}

func (newState *AppAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState AppAccessControlResponse) {
}

func (c AppAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(AppPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppAccessControlResponse
// only implements ToObjectValue() and Type().
func (o AppAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AppAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: AppPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in AppAccessControlResponse as
// a slice of AppPermission values.
// If the field is unknown or null, the boolean return value is false.
func (o *AppAccessControlResponse) GetAllPermissions(ctx context.Context) ([]AppPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []AppPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in AppAccessControlResponse.
func (o *AppAccessControlResponse) SetAllPermissions(ctx context.Context, v []AppPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time"`
	// The email of the user creates the deployment.
	Creator types.String `tfsdk:"creator"`
	// The deployment artifacts for an app.
	DeploymentArtifacts types.Object `tfsdk:"deployment_artifacts"`
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
	Status types.Object `tfsdk:"status"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeployment) {
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringRead(existingState AppDeployment) {
}

func (c AppDeployment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["deployment_artifacts"] = attrs["deployment_artifacts"].SetComputed()
	attrs["deployment_id"] = attrs["deployment_id"].SetOptional()
	attrs["mode"] = attrs["mode"].SetOptional()
	attrs["source_code_path"] = attrs["source_code_path"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
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
func (a AppDeployment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"deployment_artifacts": reflect.TypeOf(AppDeploymentArtifacts{}),
		"status":               reflect.TypeOf(AppDeploymentStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppDeployment
// only implements ToObjectValue() and Type().
func (o AppDeployment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AppDeployment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":          types.StringType,
			"creator":              types.StringType,
			"deployment_artifacts": AppDeploymentArtifacts{}.Type(ctx),
			"deployment_id":        types.StringType,
			"mode":                 types.StringType,
			"source_code_path":     types.StringType,
			"status":               AppDeploymentStatus{}.Type(ctx),
			"update_time":          types.StringType,
		},
	}
}

// GetDeploymentArtifacts returns the value of the DeploymentArtifacts field in AppDeployment as
// a AppDeploymentArtifacts value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppDeployment) GetDeploymentArtifacts(ctx context.Context) (AppDeploymentArtifacts, bool) {
	var e AppDeploymentArtifacts
	if o.DeploymentArtifacts.IsNull() || o.DeploymentArtifacts.IsUnknown() {
		return e, false
	}
	var v []AppDeploymentArtifacts
	d := o.DeploymentArtifacts.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeploymentArtifacts sets the value of the DeploymentArtifacts field in AppDeployment.
func (o *AppDeployment) SetDeploymentArtifacts(ctx context.Context, v AppDeploymentArtifacts) {
	vs := v.ToObjectValue(ctx)
	o.DeploymentArtifacts = vs
}

// GetStatus returns the value of the Status field in AppDeployment as
// a AppDeploymentStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppDeployment) GetStatus(ctx context.Context) (AppDeploymentStatus, bool) {
	var e AppDeploymentStatus
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []AppDeploymentStatus
	d := o.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in AppDeployment.
func (o *AppDeployment) SetStatus(ctx context.Context, v AppDeploymentStatus) {
	vs := v.ToObjectValue(ctx)
	o.Status = vs
}

type AppDeploymentArtifacts struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	SourceCodePath types.String `tfsdk:"source_code_path"`
}

func (newState *AppDeploymentArtifacts) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeploymentArtifacts) {
}

func (newState *AppDeploymentArtifacts) SyncEffectiveFieldsDuringRead(existingState AppDeploymentArtifacts) {
}

func (c AppDeploymentArtifacts) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppDeploymentArtifacts) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppDeploymentArtifacts
// only implements ToObjectValue() and Type().
func (o AppDeploymentArtifacts) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"source_code_path": o.SourceCodePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppDeploymentArtifacts) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"source_code_path": types.StringType,
		},
	}
}

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	Message types.String `tfsdk:"message"`
	// State of the deployment.
	State types.String `tfsdk:"state"`
}

func (newState *AppDeploymentStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeploymentStatus) {
}

func (newState *AppDeploymentStatus) SyncEffectiveFieldsDuringRead(existingState AppDeploymentStatus) {
}

func (c AppDeploymentStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppDeploymentStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppDeploymentStatus
// only implements ToObjectValue() and Type().
func (o AppDeploymentStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"state":   o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppDeploymentStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type AppPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *AppPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermission) {
}

func (newState *AppPermission) SyncEffectiveFieldsDuringRead(existingState AppPermission) {
}

func (c AppPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermission
// only implements ToObjectValue() and Type().
func (o AppPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppPermission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in AppPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AppPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in AppPermission.
func (o *AppPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type AppPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *AppPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissions) {
}

func (newState *AppPermissions) SyncEffectiveFieldsDuringRead(existingState AppPermissions) {
}

func (c AppPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AppAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermissions
// only implements ToObjectValue() and Type().
func (o AppPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AppAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in AppPermissions as
// a slice of AppAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (o *AppPermissions) GetAccessControlList(ctx context.Context) ([]AppAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AppAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in AppPermissions.
func (o *AppPermissions) SetAccessControlList(ctx context.Context, v []AppAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type AppPermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *AppPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissionsDescription) {
}

func (newState *AppPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState AppPermissionsDescription) {
}

func (c AppPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermissionsDescription
// only implements ToObjectValue() and Type().
func (o AppPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type AppPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (newState *AppPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissionsRequest) {
}

func (newState *AppPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState AppPermissionsRequest) {
}

func (c AppPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["app_name"] = attrs["app_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AppAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermissionsRequest
// only implements ToObjectValue() and Type().
func (o AppPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"app_name":            o.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AppAccessControlRequest{}.Type(ctx),
			},
			"app_name": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in AppPermissionsRequest as
// a slice of AppAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *AppPermissionsRequest) GetAccessControlList(ctx context.Context) ([]AppAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AppAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in AppPermissionsRequest.
func (o *AppPermissionsRequest) SetAccessControlList(ctx context.Context, v []AppAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type AppResource struct {
	// Description of the App Resource.
	Description types.String `tfsdk:"description"`

	Job types.Object `tfsdk:"job"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name"`

	Secret types.Object `tfsdk:"secret"`

	ServingEndpoint types.Object `tfsdk:"serving_endpoint"`

	SqlWarehouse types.Object `tfsdk:"sql_warehouse"`
}

func (newState *AppResource) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResource) {
}

func (newState *AppResource) SyncEffectiveFieldsDuringRead(existingState AppResource) {
}

func (c AppResource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["job"] = attrs["job"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["secret"] = attrs["secret"].SetOptional()
	attrs["serving_endpoint"] = attrs["serving_endpoint"].SetOptional()
	attrs["sql_warehouse"] = attrs["sql_warehouse"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AppResource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"job":              reflect.TypeOf(AppResourceJob{}),
		"secret":           reflect.TypeOf(AppResourceSecret{}),
		"serving_endpoint": reflect.TypeOf(AppResourceServingEndpoint{}),
		"sql_warehouse":    reflect.TypeOf(AppResourceSqlWarehouse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResource
// only implements ToObjectValue() and Type().
func (o AppResource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"job":              o.Job,
			"name":             o.Name,
			"secret":           o.Secret,
			"serving_endpoint": o.ServingEndpoint,
			"sql_warehouse":    o.SqlWarehouse,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"job":              AppResourceJob{}.Type(ctx),
			"name":             types.StringType,
			"secret":           AppResourceSecret{}.Type(ctx),
			"serving_endpoint": AppResourceServingEndpoint{}.Type(ctx),
			"sql_warehouse":    AppResourceSqlWarehouse{}.Type(ctx),
		},
	}
}

// GetJob returns the value of the Job field in AppResource as
// a AppResourceJob value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource) GetJob(ctx context.Context) (AppResourceJob, bool) {
	var e AppResourceJob
	if o.Job.IsNull() || o.Job.IsUnknown() {
		return e, false
	}
	var v []AppResourceJob
	d := o.Job.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJob sets the value of the Job field in AppResource.
func (o *AppResource) SetJob(ctx context.Context, v AppResourceJob) {
	vs := v.ToObjectValue(ctx)
	o.Job = vs
}

// GetSecret returns the value of the Secret field in AppResource as
// a AppResourceSecret value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource) GetSecret(ctx context.Context) (AppResourceSecret, bool) {
	var e AppResourceSecret
	if o.Secret.IsNull() || o.Secret.IsUnknown() {
		return e, false
	}
	var v []AppResourceSecret
	d := o.Secret.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSecret sets the value of the Secret field in AppResource.
func (o *AppResource) SetSecret(ctx context.Context, v AppResourceSecret) {
	vs := v.ToObjectValue(ctx)
	o.Secret = vs
}

// GetServingEndpoint returns the value of the ServingEndpoint field in AppResource as
// a AppResourceServingEndpoint value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource) GetServingEndpoint(ctx context.Context) (AppResourceServingEndpoint, bool) {
	var e AppResourceServingEndpoint
	if o.ServingEndpoint.IsNull() || o.ServingEndpoint.IsUnknown() {
		return e, false
	}
	var v []AppResourceServingEndpoint
	d := o.ServingEndpoint.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServingEndpoint sets the value of the ServingEndpoint field in AppResource.
func (o *AppResource) SetServingEndpoint(ctx context.Context, v AppResourceServingEndpoint) {
	vs := v.ToObjectValue(ctx)
	o.ServingEndpoint = vs
}

// GetSqlWarehouse returns the value of the SqlWarehouse field in AppResource as
// a AppResourceSqlWarehouse value.
// If the field is unknown or null, the boolean return value is false.
func (o *AppResource) GetSqlWarehouse(ctx context.Context) (AppResourceSqlWarehouse, bool) {
	var e AppResourceSqlWarehouse
	if o.SqlWarehouse.IsNull() || o.SqlWarehouse.IsUnknown() {
		return e, false
	}
	var v []AppResourceSqlWarehouse
	d := o.SqlWarehouse.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlWarehouse sets the value of the SqlWarehouse field in AppResource.
func (o *AppResource) SetSqlWarehouse(ctx context.Context, v AppResourceSqlWarehouse) {
	vs := v.ToObjectValue(ctx)
	o.SqlWarehouse = vs
}

type AppResourceJob struct {
	// Id of the job to grant permission on.
	Id types.String `tfsdk:"id"`
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (newState *AppResourceJob) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceJob) {
}

func (newState *AppResourceJob) SyncEffectiveFieldsDuringRead(existingState AppResourceJob) {
}

func (c AppResourceJob) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppResourceJob) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceJob
// only implements ToObjectValue() and Type().
func (o AppResourceJob) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         o.Id,
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceJob) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"permission": types.StringType,
		},
	}
}

type AppResourceSecret struct {
	// Key of the secret to grant permission on.
	Key types.String `tfsdk:"key"`
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission types.String `tfsdk:"permission"`
	// Scope of the secret to grant permission on.
	Scope types.String `tfsdk:"scope"`
}

func (newState *AppResourceSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceSecret) {
}

func (newState *AppResourceSecret) SyncEffectiveFieldsDuringRead(existingState AppResourceSecret) {
}

func (c AppResourceSecret) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppResourceSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceSecret
// only implements ToObjectValue() and Type().
func (o AppResourceSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":        o.Key,
			"permission": o.Permission,
			"scope":      o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceSecret) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":        types.StringType,
			"permission": types.StringType,
			"scope":      types.StringType,
		},
	}
}

type AppResourceServingEndpoint struct {
	// Name of the serving endpoint to grant permission on.
	Name types.String `tfsdk:"name"`
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (newState *AppResourceServingEndpoint) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceServingEndpoint) {
}

func (newState *AppResourceServingEndpoint) SyncEffectiveFieldsDuringRead(existingState AppResourceServingEndpoint) {
}

func (c AppResourceServingEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppResourceServingEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceServingEndpoint
// only implements ToObjectValue() and Type().
func (o AppResourceServingEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":       o.Name,
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceServingEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":       types.StringType,
			"permission": types.StringType,
		},
	}
}

type AppResourceSqlWarehouse struct {
	// Id of the SQL warehouse to grant permission on.
	Id types.String `tfsdk:"id"`
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission types.String `tfsdk:"permission"`
}

func (newState *AppResourceSqlWarehouse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceSqlWarehouse) {
}

func (newState *AppResourceSqlWarehouse) SyncEffectiveFieldsDuringRead(existingState AppResourceSqlWarehouse) {
}

func (c AppResourceSqlWarehouse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AppResourceSqlWarehouse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceSqlWarehouse
// only implements ToObjectValue() and Type().
func (o AppResourceSqlWarehouse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         o.Id,
			"permission": o.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AppResourceSqlWarehouse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"permission": types.StringType,
		},
	}
}

type ApplicationStatus struct {
	// Application status message
	Message types.String `tfsdk:"message"`
	// State of the application.
	State types.String `tfsdk:"state"`
}

func (newState *ApplicationStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApplicationStatus) {
}

func (newState *ApplicationStatus) SyncEffectiveFieldsDuringRead(existingState ApplicationStatus) {
}

func (c ApplicationStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ApplicationStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApplicationStatus
// only implements ToObjectValue() and Type().
func (o ApplicationStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"state":   o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ApplicationStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type ComputeStatus struct {
	// Compute status message
	Message types.String `tfsdk:"message"`
	// State of the app compute.
	State types.String `tfsdk:"state"`
}

func (newState *ComputeStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComputeStatus) {
}

func (newState *ComputeStatus) SyncEffectiveFieldsDuringRead(existingState ComputeStatus) {
}

func (c ComputeStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ComputeStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComputeStatus
// only implements ToObjectValue() and Type().
func (o ComputeStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"state":   o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ComputeStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

// Create an app deployment
type CreateAppDeploymentRequest struct {
	AppDeployment types.Object `tfsdk:"app_deployment"`
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
func (a CreateAppDeploymentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app_deployment": reflect.TypeOf(AppDeployment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAppDeploymentRequest
// only implements ToObjectValue() and Type().
func (o CreateAppDeploymentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_deployment": o.AppDeployment,
			"app_name":       o.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAppDeploymentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_deployment": AppDeployment{}.Type(ctx),
			"app_name":       types.StringType,
		},
	}
}

// GetAppDeployment returns the value of the AppDeployment field in CreateAppDeploymentRequest as
// a AppDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAppDeploymentRequest) GetAppDeployment(ctx context.Context) (AppDeployment, bool) {
	var e AppDeployment
	if o.AppDeployment.IsNull() || o.AppDeployment.IsUnknown() {
		return e, false
	}
	var v []AppDeployment
	d := o.AppDeployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAppDeployment sets the value of the AppDeployment field in CreateAppDeploymentRequest.
func (o *CreateAppDeploymentRequest) SetAppDeployment(ctx context.Context, v AppDeployment) {
	vs := v.ToObjectValue(ctx)
	o.AppDeployment = vs
}

// Create an app
type CreateAppRequest struct {
	App types.Object `tfsdk:"app"`
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
func (a CreateAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app": reflect.TypeOf(App{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAppRequest
// only implements ToObjectValue() and Type().
func (o CreateAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":        o.App,
			"no_compute": o.NoCompute,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app":        App{}.Type(ctx),
			"no_compute": types.BoolType,
		},
	}
}

// GetApp returns the value of the App field in CreateAppRequest as
// a App value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAppRequest) GetApp(ctx context.Context) (App, bool) {
	var e App
	if o.App.IsNull() || o.App.IsUnknown() {
		return e, false
	}
	var v []App
	d := o.App.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetApp sets the value of the App field in CreateAppRequest.
func (o *CreateAppRequest) SetApp(ctx context.Context, v App) {
	vs := v.ToObjectValue(ctx)
	o.App = vs
}

// Delete an app
type DeleteAppRequest struct {
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
func (a DeleteAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAppRequest
// only implements ToObjectValue() and Type().
func (o DeleteAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get an app deployment
type GetAppDeploymentRequest struct {
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
func (a GetAppDeploymentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppDeploymentRequest
// only implements ToObjectValue() and Type().
func (o GetAppDeploymentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name":      o.AppName,
			"deployment_id": o.DeploymentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppDeploymentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name":      types.StringType,
			"deployment_id": types.StringType,
		},
	}
}

// Get app permission levels
type GetAppPermissionLevelsRequest struct {
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
func (a GetAppPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetAppPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name": o.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name": types.StringType,
		},
	}
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetAppPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionLevelsResponse) {
}

func (newState *GetAppPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionLevelsResponse) {
}

func (c GetAppPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetAppPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(AppPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetAppPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: AppPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetAppPermissionLevelsResponse as
// a slice of AppPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetAppPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]AppPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []AppPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetAppPermissionLevelsResponse.
func (o *GetAppPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []AppPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Get app permissions
type GetAppPermissionsRequest struct {
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
func (a GetAppPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppPermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetAppPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name": o.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name": types.StringType,
		},
	}
}

// Get an app
type GetAppRequest struct {
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
func (a GetAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppRequest
// only implements ToObjectValue() and Type().
func (o GetAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// List app deployments
type ListAppDeploymentsRequest struct {
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
func (a ListAppDeploymentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppDeploymentsRequest
// only implements ToObjectValue() and Type().
func (o ListAppDeploymentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name":   o.AppName,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAppDeploymentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name":   types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	AppDeployments types.List `tfsdk:"app_deployments"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListAppDeploymentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppDeploymentsResponse) {
}

func (newState *ListAppDeploymentsResponse) SyncEffectiveFieldsDuringRead(existingState ListAppDeploymentsResponse) {
}

func (c ListAppDeploymentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAppDeploymentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app_deployments": reflect.TypeOf(AppDeployment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppDeploymentsResponse
// only implements ToObjectValue() and Type().
func (o ListAppDeploymentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_deployments": o.AppDeployments,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAppDeploymentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_deployments": basetypes.ListType{
				ElemType: AppDeployment{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetAppDeployments returns the value of the AppDeployments field in ListAppDeploymentsResponse as
// a slice of AppDeployment values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAppDeploymentsResponse) GetAppDeployments(ctx context.Context) ([]AppDeployment, bool) {
	if o.AppDeployments.IsNull() || o.AppDeployments.IsUnknown() {
		return nil, false
	}
	var v []AppDeployment
	d := o.AppDeployments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAppDeployments sets the value of the AppDeployments field in ListAppDeploymentsResponse.
func (o *ListAppDeploymentsResponse) SetAppDeployments(ctx context.Context, v []AppDeployment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["app_deployments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AppDeployments = types.ListValueMust(t, vs)
}

// List apps
type ListAppsRequest struct {
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
func (a ListAppsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppsRequest
// only implements ToObjectValue() and Type().
func (o ListAppsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAppsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAppsResponse struct {
	Apps types.List `tfsdk:"apps"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListAppsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppsResponse) {
}

func (newState *ListAppsResponse) SyncEffectiveFieldsDuringRead(existingState ListAppsResponse) {
}

func (c ListAppsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAppsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(App{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppsResponse
// only implements ToObjectValue() and Type().
func (o ListAppsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            o.Apps,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAppsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: App{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetApps returns the value of the Apps field in ListAppsResponse as
// a slice of App values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAppsResponse) GetApps(ctx context.Context) ([]App, bool) {
	if o.Apps.IsNull() || o.Apps.IsUnknown() {
		return nil, false
	}
	var v []App
	d := o.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in ListAppsResponse.
func (o *ListAppsResponse) SetApps(ctx context.Context, v []App) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Apps = types.ListValueMust(t, vs)
}

type StartAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *StartAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartAppRequest) {
}

func (newState *StartAppRequest) SyncEffectiveFieldsDuringRead(existingState StartAppRequest) {
}

func (c StartAppRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartAppRequest
// only implements ToObjectValue() and Type().
func (o StartAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type StopAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *StopAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopAppRequest) {
}

func (newState *StopAppRequest) SyncEffectiveFieldsDuringRead(existingState StopAppRequest) {
}

func (c StopAppRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopAppRequest
// only implements ToObjectValue() and Type().
func (o StopAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StopAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Update an app
type UpdateAppRequest struct {
	App types.Object `tfsdk:"app"`
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
func (a UpdateAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app": reflect.TypeOf(App{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAppRequest
// only implements ToObjectValue() and Type().
func (o UpdateAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":  o.App,
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app":  App{}.Type(ctx),
			"name": types.StringType,
		},
	}
}

// GetApp returns the value of the App field in UpdateAppRequest as
// a App value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateAppRequest) GetApp(ctx context.Context) (App, bool) {
	var e App
	if o.App.IsNull() || o.App.IsUnknown() {
		return e, false
	}
	var v []App
	d := o.App.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetApp sets the value of the App field in UpdateAppRequest.
func (o *UpdateAppRequest) SetApp(ctx context.Context, v App) {
	vs := v.ToObjectValue(ctx)
	o.App = vs
}
