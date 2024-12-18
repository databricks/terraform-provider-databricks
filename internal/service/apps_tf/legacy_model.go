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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type App_SdkV2 struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment types.List `tfsdk:"active_deployment" tf:"computed,object"`

	AppStatus types.List `tfsdk:"app_status" tf:"computed,object"`

	ComputeStatus types.List `tfsdk:"compute_status" tf:"computed,object"`
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
	PendingDeployment types.List `tfsdk:"pending_deployment" tf:"computed,object"`
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

func (newState *App_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan App_SdkV2) {
}

func (newState *App_SdkV2) SyncEffectiveFieldsDuringRead(existingState App_SdkV2) {
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
		"active_deployment":  reflect.TypeOf(AppDeployment_SdkV2{}),
		"app_status":         reflect.TypeOf(ApplicationStatus_SdkV2{}),
		"compute_status":     reflect.TypeOf(ComputeStatus_SdkV2{}),
		"pending_deployment": reflect.TypeOf(AppDeployment_SdkV2{}),
		"resources":          reflect.TypeOf(AppResource_SdkV2{}),
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
func (o App_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active_deployment": basetypes.ListType{
				ElemType: AppDeployment{}.Type(ctx),
			},
			"app_status": basetypes.ListType{
				ElemType: ApplicationStatus{}.Type(ctx),
			},
			"compute_status": basetypes.ListType{
				ElemType: ComputeStatus{}.Type(ctx),
			},
			"create_time":              types.StringType,
			"creator":                  types.StringType,
			"default_source_code_path": types.StringType,
			"description":              types.StringType,
			"name":                     types.StringType,
			"pending_deployment": basetypes.ListType{
				ElemType: AppDeployment{}.Type(ctx),
			},
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

type AppAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *AppAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppAccessControlRequest_SdkV2) {
}

func (newState *AppAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppAccessControlRequest_SdkV2) {
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

func (newState *AppAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppAccessControlResponse_SdkV2) {
}

func (newState *AppAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppAccessControlResponse_SdkV2) {
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
				ElemType: AppPermission{}.Type(ctx),
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
	CreateTime types.String `tfsdk:"create_time" tf:"computed"`
	// The email of the user creates the deployment.
	Creator types.String `tfsdk:"creator" tf:"computed"`
	// The deployment artifacts for an app.
	DeploymentArtifacts types.List `tfsdk:"deployment_artifacts" tf:"computed,object"`
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
	Status types.List `tfsdk:"status" tf:"computed,object"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed"`
}

func (newState *AppDeployment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeployment_SdkV2) {
}

func (newState *AppDeployment_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppDeployment_SdkV2) {
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
				ElemType: AppDeploymentArtifacts{}.Type(ctx),
			},
			"deployment_id":    types.StringType,
			"mode":             types.StringType,
			"source_code_path": types.StringType,
			"status": basetypes.ListType{
				ElemType: AppDeploymentStatus{}.Type(ctx),
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
	SourceCodePath types.String `tfsdk:"source_code_path" tf:"optional"`
}

func (newState *AppDeploymentArtifacts_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeploymentArtifacts_SdkV2) {
}

func (newState *AppDeploymentArtifacts_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppDeploymentArtifacts_SdkV2) {
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
	Message types.String `tfsdk:"message" tf:"computed"`
	// State of the deployment.
	State types.String `tfsdk:"state" tf:"computed"`
}

func (newState *AppDeploymentStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeploymentStatus_SdkV2) {
}

func (newState *AppDeploymentStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppDeploymentStatus_SdkV2) {
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

type AppPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *AppPermission_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermission_SdkV2) {
}

func (newState *AppPermission_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppPermission_SdkV2) {
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
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *AppPermissions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissions_SdkV2) {
}

func (newState *AppPermissions_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppPermissions_SdkV2) {
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
				ElemType: AppAccessControlResponse{}.Type(ctx),
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
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *AppPermissionsDescription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissionsDescription_SdkV2) {
}

func (newState *AppPermissionsDescription_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppPermissionsDescription_SdkV2) {
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
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (newState *AppPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissionsRequest_SdkV2) {
}

func (newState *AppPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppPermissionsRequest_SdkV2) {
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
				ElemType: AppAccessControlRequest{}.Type(ctx),
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
	// Description of the App Resource.
	Description types.String `tfsdk:"description" tf:"optional"`

	Job types.List `tfsdk:"job" tf:"optional,object"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name" tf:""`

	Secret types.List `tfsdk:"secret" tf:"optional,object"`

	ServingEndpoint types.List `tfsdk:"serving_endpoint" tf:"optional,object"`

	SqlWarehouse types.List `tfsdk:"sql_warehouse" tf:"optional,object"`
}

func (newState *AppResource_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResource_SdkV2) {
}

func (newState *AppResource_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppResource_SdkV2) {
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
		"job":              reflect.TypeOf(AppResourceJob_SdkV2{}),
		"secret":           reflect.TypeOf(AppResourceSecret_SdkV2{}),
		"serving_endpoint": reflect.TypeOf(AppResourceServingEndpoint_SdkV2{}),
		"sql_warehouse":    reflect.TypeOf(AppResourceSqlWarehouse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResource_SdkV2
// only implements ToObjectValue() and Type().
func (o AppResource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AppResource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"job": basetypes.ListType{
				ElemType: AppResourceJob{}.Type(ctx),
			},
			"name": types.StringType,
			"secret": basetypes.ListType{
				ElemType: AppResourceSecret{}.Type(ctx),
			},
			"serving_endpoint": basetypes.ListType{
				ElemType: AppResourceServingEndpoint{}.Type(ctx),
			},
			"sql_warehouse": basetypes.ListType{
				ElemType: AppResourceSqlWarehouse{}.Type(ctx),
			},
		},
	}
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

type AppResourceJob_SdkV2 struct {
	// Id of the job to grant permission on.
	Id types.String `tfsdk:"id" tf:""`
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission types.String `tfsdk:"permission" tf:""`
}

func (newState *AppResourceJob_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceJob_SdkV2) {
}

func (newState *AppResourceJob_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppResourceJob_SdkV2) {
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
	Key types.String `tfsdk:"key" tf:""`
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission types.String `tfsdk:"permission" tf:""`
	// Scope of the secret to grant permission on.
	Scope types.String `tfsdk:"scope" tf:""`
}

func (newState *AppResourceSecret_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceSecret_SdkV2) {
}

func (newState *AppResourceSecret_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppResourceSecret_SdkV2) {
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
	Name types.String `tfsdk:"name" tf:""`
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission types.String `tfsdk:"permission" tf:""`
}

func (newState *AppResourceServingEndpoint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceServingEndpoint_SdkV2) {
}

func (newState *AppResourceServingEndpoint_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppResourceServingEndpoint_SdkV2) {
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
	Id types.String `tfsdk:"id" tf:""`
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission types.String `tfsdk:"permission" tf:""`
}

func (newState *AppResourceSqlWarehouse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResourceSqlWarehouse_SdkV2) {
}

func (newState *AppResourceSqlWarehouse_SdkV2) SyncEffectiveFieldsDuringRead(existingState AppResourceSqlWarehouse_SdkV2) {
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

type ApplicationStatus_SdkV2 struct {
	// Application status message
	Message types.String `tfsdk:"message" tf:"computed"`
	// State of the application.
	State types.String `tfsdk:"state" tf:"computed"`
}

func (newState *ApplicationStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApplicationStatus_SdkV2) {
}

func (newState *ApplicationStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState ApplicationStatus_SdkV2) {
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
	Message types.String `tfsdk:"message" tf:"computed"`
	// State of the app compute.
	State types.String `tfsdk:"state" tf:"computed"`
}

func (newState *ComputeStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComputeStatus_SdkV2) {
}

func (newState *ComputeStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState ComputeStatus_SdkV2) {
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

// Create an app deployment
type CreateAppDeploymentRequest_SdkV2 struct {
	AppDeployment types.List `tfsdk:"app_deployment" tf:"optional,object"`
	// The name of the app.
	AppName types.String `tfsdk:"-"`
}

func (newState *CreateAppDeploymentRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAppDeploymentRequest_SdkV2) {
}

func (newState *CreateAppDeploymentRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateAppDeploymentRequest_SdkV2) {
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
				ElemType: AppDeployment{}.Type(ctx),
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

// Create an app
type CreateAppRequest_SdkV2 struct {
	App types.List `tfsdk:"app" tf:"optional,object"`
}

func (newState *CreateAppRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAppRequest_SdkV2) {
}

func (newState *CreateAppRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateAppRequest_SdkV2) {
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
			"app": o.App,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAppRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app": basetypes.ListType{
				ElemType: App{}.Type(ctx),
			},
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

// Delete an app
type DeleteAppRequest_SdkV2 struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteAppRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAppRequest_SdkV2) {
}

func (newState *DeleteAppRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteAppRequest_SdkV2) {
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

// Get an app deployment
type GetAppDeploymentRequest_SdkV2 struct {
	// The name of the app.
	AppName types.String `tfsdk:"-"`
	// The unique id of the deployment.
	DeploymentId types.String `tfsdk:"-"`
}

func (newState *GetAppDeploymentRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppDeploymentRequest_SdkV2) {
}

func (newState *GetAppDeploymentRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetAppDeploymentRequest_SdkV2) {
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

// Get app permission levels
type GetAppPermissionLevelsRequest_SdkV2 struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (newState *GetAppPermissionLevelsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionLevelsRequest_SdkV2) {
}

func (newState *GetAppPermissionLevelsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionLevelsRequest_SdkV2) {
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
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetAppPermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionLevelsResponse_SdkV2) {
}

func (newState *GetAppPermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionLevelsResponse_SdkV2) {
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
				ElemType: AppPermissionsDescription{}.Type(ctx),
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

// Get app permissions
type GetAppPermissionsRequest_SdkV2 struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (newState *GetAppPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionsRequest_SdkV2) {
}

func (newState *GetAppPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionsRequest_SdkV2) {
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

// Get an app
type GetAppRequest_SdkV2 struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *GetAppRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppRequest_SdkV2) {
}

func (newState *GetAppRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetAppRequest_SdkV2) {
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

// List app deployments
type ListAppDeploymentsRequest_SdkV2 struct {
	// The name of the app.
	AppName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListAppDeploymentsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppDeploymentsRequest_SdkV2) {
}

func (newState *ListAppDeploymentsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAppDeploymentsRequest_SdkV2) {
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
	AppDeployments types.List `tfsdk:"app_deployments" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListAppDeploymentsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppDeploymentsResponse_SdkV2) {
}

func (newState *ListAppDeploymentsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAppDeploymentsResponse_SdkV2) {
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
				ElemType: AppDeployment{}.Type(ctx),
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

// List apps
type ListAppsRequest_SdkV2 struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListAppsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppsRequest_SdkV2) {
}

func (newState *ListAppsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAppsRequest_SdkV2) {
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
	Apps types.List `tfsdk:"apps" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListAppsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppsResponse_SdkV2) {
}

func (newState *ListAppsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAppsResponse_SdkV2) {
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
				ElemType: App{}.Type(ctx),
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

type StartAppRequest_SdkV2 struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *StartAppRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartAppRequest_SdkV2) {
}

func (newState *StartAppRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState StartAppRequest_SdkV2) {
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

func (newState *StopAppRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopAppRequest_SdkV2) {
}

func (newState *StopAppRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState StopAppRequest_SdkV2) {
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

// Update an app
type UpdateAppRequest_SdkV2 struct {
	App types.List `tfsdk:"app" tf:"optional,object"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"-"`
}

func (newState *UpdateAppRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAppRequest_SdkV2) {
}

func (newState *UpdateAppRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateAppRequest_SdkV2) {
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
				ElemType: App{}.Type(ctx),
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
