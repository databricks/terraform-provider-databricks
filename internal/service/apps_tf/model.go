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
	ActiveDeployment types.Object `tfsdk:"active_deployment" tf:"computed,object"`

	AppStatus types.Object `tfsdk:"app_status" tf:"computed,object"`

	ComputeStatus types.Object `tfsdk:"compute_status" tf:"computed,object"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time" tf:"computed"`
	// The email of the user that created the app.
	Creator types.String `tfsdk:"creator" tf:"computed"`
	// The default workspace file system path of the source code from which app
	// deployment are created. This field tracks the workspace source code path
	// of the last active deployment.
	DefaultSourceCodePath types.String `tfsdk:"default_source_code_path" tf:"computed"`
	// The description of the app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name" tf:""`
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	PendingDeployment types.Object `tfsdk:"pending_deployment" tf:"computed,object"`
	// Resources for the app.
	Resources types.List `tfsdk:"resources" tf:"optional"`

	ServicePrincipalClientId types.String `tfsdk:"service_principal_client_id" tf:"computed"`

	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id" tf:"computed"`

	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"computed"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed"`
	// The email of the user that last updated the app.
	Updater types.String `tfsdk:"updater" tf:"computed"`
	// The URL of the app once it is deployed.
	Url types.String `tfsdk:"url" tf:"computed"`
}

func (newState *App) SyncEffectiveFieldsDuringCreateOrUpdate(plan App) {
}

func (newState *App) SyncEffectiveFieldsDuringRead(existingState App) {
}

func (c App) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetComputed(append(path, "active_deployment")...)
	AppDeployment{}.ApplySchemaCustomizations(cs, append(path, "active_deployment")...)
	cs.SetComputed(append(path, "app_status")...)
	ApplicationStatus{}.ApplySchemaCustomizations(cs, append(path, "app_status")...)
	cs.SetComputed(append(path, "compute_status")...)
	ComputeStatus{}.ApplySchemaCustomizations(cs, append(path, "compute_status")...)
	cs.SetComputed(append(path, "create_time")...)
	cs.SetComputed(append(path, "creator")...)
	cs.SetComputed(append(path, "default_source_code_path")...)
	cs.SetRequired(append(path, "name")...)
	cs.SetComputed(append(path, "pending_deployment")...)
	AppDeployment{}.ApplySchemaCustomizations(cs, append(path, "pending_deployment")...)
	AppResource{}.ApplySchemaCustomizations(cs, append(path, "resources")...)
	cs.SetComputed(append(path, "service_principal_client_id")...)
	cs.SetComputed(append(path, "service_principal_id")...)
	cs.SetComputed(append(path, "service_principal_name")...)
	cs.SetComputed(append(path, "update_time")...)
	cs.SetComputed(append(path, "updater")...)
	cs.SetComputed(append(path, "url")...)

	return cs
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
		"active_deployment":  reflect.TypeOf(AppDeployment{}),
		"app_status":         reflect.TypeOf(ApplicationStatus{}),
		"compute_status":     reflect.TypeOf(ComputeStatus{}),
		"pending_deployment": reflect.TypeOf(AppDeployment{}),
		"resources":          reflect.TypeOf(AppResource{}),
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
			"compute_status":              o.ComputeStatus,
			"create_time":                 o.CreateTime,
			"creator":                     o.Creator,
			"default_source_code_path":    o.DefaultSourceCodePath,
			"description":                 o.Description,
			"name":                        o.Name,
			"pending_deployment":          o.PendingDeployment,
			"resources":                   o.Resources,
			"service_principal_client_id": o.ServicePrincipalClientId,
			"service_principal_id":        o.ServicePrincipalId,
			"service_principal_name":      o.ServicePrincipalName,
			"update_time":                 o.UpdateTime,
			"updater":                     o.Updater,
			"url":                         o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o App) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active_deployment":        AppDeployment{}.Type(ctx),
			"app_status":               ApplicationStatus{}.Type(ctx),
			"compute_status":           ComputeStatus{}.Type(ctx),
			"create_time":              types.StringType,
			"creator":                  types.StringType,
			"default_source_code_path": types.StringType,
			"description":              types.StringType,
			"name":                     types.StringType,
			"pending_deployment":       AppDeployment{}.Type(ctx),
			"resources": basetypes.ListType{
				ElemType: AppResource{}.Type(ctx),
			},
			"service_principal_client_id": types.StringType,
			"service_principal_id":        types.Int64Type,
			"service_principal_name":      types.StringType,
			"update_time":                 types.StringType,
			"updater":                     types.StringType,
			"url":                         types.StringType,
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

type AppAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *AppAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppAccessControlRequest) {
}

func (newState *AppAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState AppAccessControlRequest) {
}

func (c AppAccessControlRequest) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
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
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *AppAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppAccessControlResponse) {
}

func (newState *AppAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState AppAccessControlResponse) {
}

func (c AppAccessControlResponse) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AppPermission{}.ApplySchemaCustomizations(cs, append(path, "all_permissions")...)

	return cs
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
	CreateTime types.String `tfsdk:"create_time" tf:"computed"`
	// The email of the user creates the deployment.
	Creator types.String `tfsdk:"creator" tf:"computed"`
	// The deployment artifacts for an app.
	DeploymentArtifacts types.Object `tfsdk:"deployment_artifacts" tf:"computed,object"`
	// The unique id of the deployment.
	DeploymentId types.String `tfsdk:"deployment_id" tf:"optional"`
	// The mode of which the deployment will manage the source code.
	Mode types.String `tfsdk:"mode" tf:"optional"`
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	SourceCodePath types.String `tfsdk:"source_code_path" tf:"optional"`
	// Status and status message of the deployment
	Status types.Object `tfsdk:"status" tf:"computed,object"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed"`
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeployment) {
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringRead(existingState AppDeployment) {
}

func (c AppDeployment) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetComputed(append(path, "create_time")...)
	cs.SetComputed(append(path, "creator")...)
	cs.SetComputed(append(path, "deployment_artifacts")...)
	AppDeploymentArtifacts{}.ApplySchemaCustomizations(cs, append(path, "deployment_artifacts")...)
	cs.SetComputed(append(path, "status")...)
	AppDeploymentStatus{}.ApplySchemaCustomizations(cs, append(path, "status")...)
	cs.SetComputed(append(path, "update_time")...)

	return cs
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
	SourceCodePath types.String `tfsdk:"source_code_path" tf:"optional"`
}

func (newState *AppDeploymentArtifacts) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeploymentArtifacts) {
}

func (newState *AppDeploymentArtifacts) SyncEffectiveFieldsDuringRead(existingState AppDeploymentArtifacts) {
}

func (c AppDeploymentArtifacts) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
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
	Message types.String `tfsdk:"message" tf:"computed"`
	// State of the deployment.
	State types.String `tfsdk:"state" tf:"computed"`
}

func (newState *AppDeploymentStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeploymentStatus) {
}

func (newState *AppDeploymentStatus) SyncEffectiveFieldsDuringRead(existingState AppDeploymentStatus) {
}

func (c AppDeploymentStatus) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetComputed(append(path, "message")...)
	cs.SetComputed(append(path, "state")...)

	return cs
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
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *AppPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermission) {
}

func (newState *AppPermission) SyncEffectiveFieldsDuringRead(existingState AppPermission) {
}

func (c AppPermission) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
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
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *AppPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissions) {
}

func (newState *AppPermissions) SyncEffectiveFieldsDuringRead(existingState AppPermissions) {
}

func (c AppPermissions) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AppAccessControlResponse{}.ApplySchemaCustomizations(cs, append(path, "access_control_list")...)

	return cs
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
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *AppPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissionsDescription) {
}

func (newState *AppPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState AppPermissionsDescription) {
}

func (c AppPermissionsDescription) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
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
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (newState *AppPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissionsRequest) {
}

func (newState *AppPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState AppPermissionsRequest) {
}

func (c AppPermissionsRequest) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AppAccessControlRequest{}.ApplySchemaCustomizations(cs, append(path, "access_control_list")...)
	cs.SetRequired(append(path, "app_name")...)

	return cs
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
	Description types.String `tfsdk:"description" tf:"optional"`

	Job types.Object `tfsdk:"job" tf:"optional,object"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name" tf:""`

	Secret types.Object `tfsdk:"secret" tf:"optional,object"`

	ServingEndpoint types.Object `tfsdk:"serving_endpoint" tf:"optional,object"`

	SqlWarehouse types.Object `tfsdk:"sql_warehouse" tf:"optional,object"`
}

func (newState *AppResource) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResource) {
}

func (newState *AppResource) SyncEffectiveFieldsDuringRead(existingState AppResource) {
}

func (c AppResource) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AppResourceJob{}.ApplySchemaCustomizations(cs, append(path, "job")...)
	cs.SetRequired(append(path, "name")...)
	AppResourceSecret{}.ApplySchemaCustomizations(cs, append(path, "secret")...)
	AppResourceServingEndpoint{}.ApplySchemaCustomizations(cs, append(path, "serving_endpoint")...)
	AppResourceSqlWarehouse{}.ApplySchemaCustomizations(cs, append(path, "sql_warehouse")...)

	return cs
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
	Id types.String `tfsdk:"id" tf:""`
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission types.String `tfsdk:"permission" tf:""`
}

func (newState *AppResourceJob) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceJob) {
}

func (newState *AppResourceJob) SyncEffectiveFieldsDuringRead(existingState AppResourceJob) {
}

func (c AppResourceJob) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "id")...)
	cs.SetRequired(append(path, "permission")...)

	return cs
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
	Key types.String `tfsdk:"key" tf:""`
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission types.String `tfsdk:"permission" tf:""`
	// Scope of the secret to grant permission on.
	Scope types.String `tfsdk:"scope" tf:""`
}

func (newState *AppResourceSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceSecret) {
}

func (newState *AppResourceSecret) SyncEffectiveFieldsDuringRead(existingState AppResourceSecret) {
}

func (c AppResourceSecret) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "key")...)
	cs.SetRequired(append(path, "permission")...)
	cs.SetRequired(append(path, "scope")...)

	return cs
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
	Name types.String `tfsdk:"name" tf:""`
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission types.String `tfsdk:"permission" tf:""`
}

func (newState *AppResourceServingEndpoint) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceServingEndpoint) {
}

func (newState *AppResourceServingEndpoint) SyncEffectiveFieldsDuringRead(existingState AppResourceServingEndpoint) {
}

func (c AppResourceServingEndpoint) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "name")...)
	cs.SetRequired(append(path, "permission")...)

	return cs
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
	Id types.String `tfsdk:"id" tf:""`
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission types.String `tfsdk:"permission" tf:""`
}

func (newState *AppResourceSqlWarehouse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceSqlWarehouse) {
}

func (newState *AppResourceSqlWarehouse) SyncEffectiveFieldsDuringRead(existingState AppResourceSqlWarehouse) {
}

func (c AppResourceSqlWarehouse) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "id")...)
	cs.SetRequired(append(path, "permission")...)

	return cs
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
	Message types.String `tfsdk:"message" tf:"computed"`
	// State of the application.
	State types.String `tfsdk:"state" tf:"computed"`
}

func (newState *ApplicationStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApplicationStatus) {
}

func (newState *ApplicationStatus) SyncEffectiveFieldsDuringRead(existingState ApplicationStatus) {
}

func (c ApplicationStatus) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetComputed(append(path, "message")...)
	cs.SetComputed(append(path, "state")...)

	return cs
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
	Message types.String `tfsdk:"message" tf:"computed"`
	// State of the app compute.
	State types.String `tfsdk:"state" tf:"computed"`
}

func (newState *ComputeStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComputeStatus) {
}

func (newState *ComputeStatus) SyncEffectiveFieldsDuringRead(existingState ComputeStatus) {
}

func (c ComputeStatus) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetComputed(append(path, "message")...)
	cs.SetComputed(append(path, "state")...)

	return cs
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
	AppDeployment types.Object `tfsdk:"app_deployment" tf:"optional,object"`
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
	App types.Object `tfsdk:"app" tf:"optional,object"`
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
			"app": o.App,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app": App{}.Type(ctx),
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
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetAppPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionLevelsResponse) {
}

func (newState *GetAppPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionLevelsResponse) {
}

func (c GetAppPermissionLevelsResponse) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AppPermissionsDescription{}.ApplySchemaCustomizations(cs, append(path, "permission_levels")...)

	return cs
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
	AppDeployments types.List `tfsdk:"app_deployments" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListAppDeploymentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppDeploymentsResponse) {
}

func (newState *ListAppDeploymentsResponse) SyncEffectiveFieldsDuringRead(existingState ListAppDeploymentsResponse) {
}

func (c ListAppDeploymentsResponse) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AppDeployment{}.ApplySchemaCustomizations(cs, append(path, "app_deployments")...)

	return cs
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
	Apps types.List `tfsdk:"apps" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListAppsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppsResponse) {
}

func (newState *ListAppsResponse) SyncEffectiveFieldsDuringRead(existingState ListAppsResponse) {
}

func (c ListAppsResponse) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	App{}.ApplySchemaCustomizations(cs, append(path, "apps")...)

	return cs
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

func (c StartAppRequest) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "name")...)

	return cs
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

func (c StopAppRequest) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "name")...)

	return cs
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
	App types.Object `tfsdk:"app" tf:"optional,object"`
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
