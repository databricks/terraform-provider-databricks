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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment types.List `tfsdk:"active_deployment" tf:"optional,object"`

	AppStatus types.List `tfsdk:"app_status" tf:"optional,object"`

	ComputeStatus types.List `tfsdk:"compute_status" tf:"optional,object"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time" tf:"computed,optional"`
	// The email of the user that created the app.
	Creator types.String `tfsdk:"creator" tf:"computed,optional"`
	// The default workspace file system path of the source code from which app
	// deployment are created. This field tracks the workspace source code path
	// of the last active deployment.
	DefaultSourceCodePath types.String `tfsdk:"default_source_code_path" tf:"optional"`
	// The description of the app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name" tf:""`
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	PendingDeployment types.List `tfsdk:"pending_deployment" tf:"optional,object"`
	// Resources for the app.
	Resources types.List `tfsdk:"resources" tf:"optional"`

	ServicePrincipalClientId types.String `tfsdk:"service_principal_client_id" tf:"computed,optional"`

	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id" tf:"computed,optional"`

	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"computed,optional"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed,optional"`
	// The email of the user that last updated the app.
	Updater types.String `tfsdk:"updater" tf:"computed,optional"`
	// The URL of the app once it is deployed.
	Url types.String `tfsdk:"url" tf:"computed,optional"`
}

func (newState *App) SyncEffectiveFieldsDuringCreateOrUpdate(plan App) {
}

func (newState *App) SyncEffectiveFieldsDuringRead(existingState App) {
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

// ToObjectType returns the representation of App in the Terraform plugin framework type
// system.
func (a App) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active_deployment": basetypes.ListType{
				ElemType: AppDeployment{}.ToObjectType(ctx),
			},
			"app_status": basetypes.ListType{
				ElemType: ApplicationStatus{}.ToObjectType(ctx),
			},
			"compute_status": basetypes.ListType{
				ElemType: ComputeStatus{}.ToObjectType(ctx),
			},
			"create_time":              types.StringType,
			"creator":                  types.StringType,
			"default_source_code_path": types.StringType,
			"description":              types.StringType,
			"name":                     types.StringType,
			"pending_deployment": basetypes.ListType{
				ElemType: AppDeployment{}.ToObjectType(ctx),
			},
			"resources": basetypes.ListType{
				ElemType: AppResource{}.ToObjectType(ctx),
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

// ToObjectType returns the representation of AppAccessControlRequest in the Terraform plugin framework type
// system.
func (a AppAccessControlRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of AppAccessControlResponse in the Terraform plugin framework type
// system.
func (a AppAccessControlResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: AppPermission{}.ToObjectType(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time" tf:"computed,optional"`
	// The email of the user creates the deployment.
	Creator types.String `tfsdk:"creator" tf:"computed,optional"`
	// The deployment artifacts for an app.
	DeploymentArtifacts types.List `tfsdk:"deployment_artifacts" tf:"optional,object"`
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
	Status types.List `tfsdk:"status" tf:"optional,object"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed,optional"`
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeployment) {
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringRead(existingState AppDeployment) {
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

// ToObjectType returns the representation of AppDeployment in the Terraform plugin framework type
// system.
func (a AppDeployment) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"creator":     types.StringType,
			"deployment_artifacts": basetypes.ListType{
				ElemType: AppDeploymentArtifacts{}.ToObjectType(ctx),
			},
			"deployment_id":    types.StringType,
			"mode":             types.StringType,
			"source_code_path": types.StringType,
			"status": basetypes.ListType{
				ElemType: AppDeploymentStatus{}.ToObjectType(ctx),
			},
			"update_time": types.StringType,
		},
	}
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

// ToObjectType returns the representation of AppDeploymentArtifacts in the Terraform plugin framework type
// system.
func (a AppDeploymentArtifacts) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"source_code_path": types.StringType,
		},
	}
}

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	Message types.String `tfsdk:"message" tf:"computed,optional"`
	// State of the deployment.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *AppDeploymentStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeploymentStatus) {
}

func (newState *AppDeploymentStatus) SyncEffectiveFieldsDuringRead(existingState AppDeploymentStatus) {
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

// ToObjectType returns the representation of AppDeploymentStatus in the Terraform plugin framework type
// system.
func (a AppDeploymentStatus) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of AppPermission in the Terraform plugin framework type
// system.
func (a AppPermission) ToObjectType(ctx context.Context) types.ObjectType {
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

type AppPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *AppPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissions) {
}

func (newState *AppPermissions) SyncEffectiveFieldsDuringRead(existingState AppPermissions) {
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

// ToObjectType returns the representation of AppPermissions in the Terraform plugin framework type
// system.
func (a AppPermissions) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AppAccessControlResponse{}.ToObjectType(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
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

// ToObjectType returns the representation of AppPermissionsDescription in the Terraform plugin framework type
// system.
func (a AppPermissionsDescription) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of AppPermissionsRequest in the Terraform plugin framework type
// system.
func (a AppPermissionsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AppAccessControlRequest{}.ToObjectType(ctx),
			},
			"app_name": types.StringType,
		},
	}
}

type AppResource struct {
	// Description of the App Resource.
	Description types.String `tfsdk:"description" tf:"optional"`

	Job types.List `tfsdk:"job" tf:"optional,object"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name" tf:""`

	Secret types.List `tfsdk:"secret" tf:"optional,object"`

	ServingEndpoint types.List `tfsdk:"serving_endpoint" tf:"optional,object"`

	SqlWarehouse types.List `tfsdk:"sql_warehouse" tf:"optional,object"`
}

func (newState *AppResource) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResource) {
}

func (newState *AppResource) SyncEffectiveFieldsDuringRead(existingState AppResource) {
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

// ToObjectType returns the representation of AppResource in the Terraform plugin framework type
// system.
func (a AppResource) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"job": basetypes.ListType{
				ElemType: AppResourceJob{}.ToObjectType(ctx),
			},
			"name": types.StringType,
			"secret": basetypes.ListType{
				ElemType: AppResourceSecret{}.ToObjectType(ctx),
			},
			"serving_endpoint": basetypes.ListType{
				ElemType: AppResourceServingEndpoint{}.ToObjectType(ctx),
			},
			"sql_warehouse": basetypes.ListType{
				ElemType: AppResourceSqlWarehouse{}.ToObjectType(ctx),
			},
		},
	}
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

// ToObjectType returns the representation of AppResourceJob in the Terraform plugin framework type
// system.
func (a AppResourceJob) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of AppResourceSecret in the Terraform plugin framework type
// system.
func (a AppResourceSecret) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of AppResourceServingEndpoint in the Terraform plugin framework type
// system.
func (a AppResourceServingEndpoint) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of AppResourceSqlWarehouse in the Terraform plugin framework type
// system.
func (a AppResourceSqlWarehouse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"permission": types.StringType,
		},
	}
}

type ApplicationStatus struct {
	// Application status message
	Message types.String `tfsdk:"message" tf:"computed,optional"`
	// State of the application.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *ApplicationStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApplicationStatus) {
}

func (newState *ApplicationStatus) SyncEffectiveFieldsDuringRead(existingState ApplicationStatus) {
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

// ToObjectType returns the representation of ApplicationStatus in the Terraform plugin framework type
// system.
func (a ApplicationStatus) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type ComputeStatus struct {
	// Compute status message
	Message types.String `tfsdk:"message" tf:"computed,optional"`
	// State of the app compute.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *ComputeStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComputeStatus) {
}

func (newState *ComputeStatus) SyncEffectiveFieldsDuringRead(existingState ComputeStatus) {
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

// ToObjectType returns the representation of ComputeStatus in the Terraform plugin framework type
// system.
func (a ComputeStatus) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

// Create an app deployment
type CreateAppDeploymentRequest struct {
	AppDeployment types.List `tfsdk:"app_deployment" tf:"optional,object"`
	// The name of the app.
	AppName types.String `tfsdk:"-"`
}

func (newState *CreateAppDeploymentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAppDeploymentRequest) {
}

func (newState *CreateAppDeploymentRequest) SyncEffectiveFieldsDuringRead(existingState CreateAppDeploymentRequest) {
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

// ToObjectType returns the representation of CreateAppDeploymentRequest in the Terraform plugin framework type
// system.
func (a CreateAppDeploymentRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_deployment": basetypes.ListType{
				ElemType: AppDeployment{}.ToObjectType(ctx),
			},
			"app_name": types.StringType,
		},
	}
}

// Create an app
type CreateAppRequest struct {
	App types.List `tfsdk:"app" tf:"optional,object"`
}

func (newState *CreateAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAppRequest) {
}

func (newState *CreateAppRequest) SyncEffectiveFieldsDuringRead(existingState CreateAppRequest) {
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

// ToObjectType returns the representation of CreateAppRequest in the Terraform plugin framework type
// system.
func (a CreateAppRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app": basetypes.ListType{
				ElemType: App{}.ToObjectType(ctx),
			},
		},
	}
}

// Delete an app
type DeleteAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAppRequest) {
}

func (newState *DeleteAppRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAppRequest) {
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

// ToObjectType returns the representation of DeleteAppRequest in the Terraform plugin framework type
// system.
func (a DeleteAppRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

func (newState *GetAppDeploymentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppDeploymentRequest) {
}

func (newState *GetAppDeploymentRequest) SyncEffectiveFieldsDuringRead(existingState GetAppDeploymentRequest) {
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

// ToObjectType returns the representation of GetAppDeploymentRequest in the Terraform plugin framework type
// system.
func (a GetAppDeploymentRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

func (newState *GetAppPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionLevelsRequest) {
}

func (newState *GetAppPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionLevelsRequest) {
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

// ToObjectType returns the representation of GetAppPermissionLevelsRequest in the Terraform plugin framework type
// system.
func (a GetAppPermissionLevelsRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of GetAppPermissionLevelsResponse in the Terraform plugin framework type
// system.
func (a GetAppPermissionLevelsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: AppPermissionsDescription{}.ToObjectType(ctx),
			},
		},
	}
}

// Get app permissions
type GetAppPermissionsRequest struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (newState *GetAppPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionsRequest) {
}

func (newState *GetAppPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionsRequest) {
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

// ToObjectType returns the representation of GetAppPermissionsRequest in the Terraform plugin framework type
// system.
func (a GetAppPermissionsRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

func (newState *GetAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppRequest) {
}

func (newState *GetAppRequest) SyncEffectiveFieldsDuringRead(existingState GetAppRequest) {
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

// ToObjectType returns the representation of GetAppRequest in the Terraform plugin framework type
// system.
func (a GetAppRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

func (newState *ListAppDeploymentsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppDeploymentsRequest) {
}

func (newState *ListAppDeploymentsRequest) SyncEffectiveFieldsDuringRead(existingState ListAppDeploymentsRequest) {
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

// ToObjectType returns the representation of ListAppDeploymentsRequest in the Terraform plugin framework type
// system.
func (a ListAppDeploymentsRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of ListAppDeploymentsResponse in the Terraform plugin framework type
// system.
func (a ListAppDeploymentsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_deployments": basetypes.ListType{
				ElemType: AppDeployment{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// List apps
type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListAppsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppsRequest) {
}

func (newState *ListAppsRequest) SyncEffectiveFieldsDuringRead(existingState ListAppsRequest) {
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

// ToObjectType returns the representation of ListAppsRequest in the Terraform plugin framework type
// system.
func (a ListAppsRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of ListAppsResponse in the Terraform plugin framework type
// system.
func (a ListAppsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: App{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

type StartAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *StartAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartAppRequest) {
}

func (newState *StartAppRequest) SyncEffectiveFieldsDuringRead(existingState StartAppRequest) {
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

// ToObjectType returns the representation of StartAppRequest in the Terraform plugin framework type
// system.
func (a StartAppRequest) ToObjectType(ctx context.Context) types.ObjectType {
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

// ToObjectType returns the representation of StopAppRequest in the Terraform plugin framework type
// system.
func (a StopAppRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Update an app
type UpdateAppRequest struct {
	App types.List `tfsdk:"app" tf:"optional,object"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"-"`
}

func (newState *UpdateAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAppRequest) {
}

func (newState *UpdateAppRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAppRequest) {
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

// ToObjectType returns the representation of UpdateAppRequest in the Terraform plugin framework type
// system.
func (a UpdateAppRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app": basetypes.ListType{
				ElemType: App{}.ToObjectType(ctx),
			},
			"name": types.StringType,
		},
	}
}
