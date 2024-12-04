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
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment types.Object `tfsdk:"active_deployment" tf:"optional,object"`

	AppStatus types.Object `tfsdk:"app_status" tf:"optional,object"`

	ComputeStatus types.Object `tfsdk:"compute_status" tf:"optional,object"`
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
	PendingDeployment types.Object `tfsdk:"pending_deployment" tf:"optional,object"`
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

func (a *App) GetNestedTypes() map[string]any {
	return map[string]any{
		"active_deployment": AppDeployment{},
	}
}

func (newState *App) SyncEffectiveFieldsDuringCreateOrUpdate(plan App) {
}

func (newState *App) SyncEffectiveFieldsDuringRead(existingState App) {
}

func (a App) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ActiveDeployment":  reflect.TypeOf(AppDeployment{}),
		"AppStatus":         reflect.TypeOf(ApplicationStatus{}),
		"ComputeStatus":     reflect.TypeOf(ComputeStatus{}),
		"PendingDeployment": reflect.TypeOf(AppDeployment{}),
		"Resources":         reflect.TypeOf(AppResource{}),
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

func (a AppAccessControlRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a AppAccessControlResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AllPermissions": reflect.TypeOf(AppPermission{}),
	}
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time" tf:"computed,optional"`
	// The email of the user creates the deployment.
	Creator types.String `tfsdk:"creator" tf:"computed,optional"`
	// The deployment artifacts for an app.
	DeploymentArtifacts types.Object `tfsdk:"deployment_artifacts" tf:"optional,object"`
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
	Status types.Object `tfsdk:"status" tf:"optional,object"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed,optional"`
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeployment) {
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringRead(existingState AppDeployment) {
}

func (a AppDeployment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DeploymentArtifacts": reflect.TypeOf(AppDeploymentArtifacts{}),
		"Status":              reflect.TypeOf(AppDeploymentStatus{}),
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

func (a AppDeploymentArtifacts) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a AppDeploymentStatus) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a AppPermission) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"InheritedFromObject": reflect.TypeOf(""),
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

func (a AppPermissions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(AppAccessControlResponse{}),
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

func (a AppPermissionsDescription) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a AppPermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(AppAccessControlRequest{}),
	}
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

func (a AppResource) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Job":             reflect.TypeOf(AppResourceJob{}),
		"Secret":          reflect.TypeOf(AppResourceSecret{}),
		"ServingEndpoint": reflect.TypeOf(AppResourceServingEndpoint{}),
		"SqlWarehouse":    reflect.TypeOf(AppResourceSqlWarehouse{}),
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

func (a AppResourceJob) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a AppResourceSecret) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a AppResourceServingEndpoint) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a AppResourceSqlWarehouse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a ApplicationStatus) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a ComputeStatus) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Create an app deployment
type CreateAppDeploymentRequest struct {
	AppDeployment types.Object `tfsdk:"app_deployment" tf:"optional,object"`
	// The name of the app.
	AppName types.String `tfsdk:"-"`
}

func (newState *CreateAppDeploymentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAppDeploymentRequest) {
}

func (newState *CreateAppDeploymentRequest) SyncEffectiveFieldsDuringRead(existingState CreateAppDeploymentRequest) {
}

func (a CreateAppDeploymentRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AppDeployment": reflect.TypeOf(AppDeployment{}),
	}
}

// Create an app
type CreateAppRequest struct {
	App types.Object `tfsdk:"app" tf:"optional,object"`
}

func (newState *CreateAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAppRequest) {
}

func (newState *CreateAppRequest) SyncEffectiveFieldsDuringRead(existingState CreateAppRequest) {
}

func (a CreateAppRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"App": reflect.TypeOf(App{}),
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

func (a DeleteAppRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a GetAppDeploymentRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a GetAppPermissionLevelsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetAppPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionLevelsResponse) {
}

func (newState *GetAppPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionLevelsResponse) {
}

func (a GetAppPermissionLevelsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PermissionLevels": reflect.TypeOf(AppPermissionsDescription{}),
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

func (a GetAppPermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a GetAppRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a ListAppDeploymentsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a ListAppDeploymentsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AppDeployments": reflect.TypeOf(AppDeployment{}),
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

func (a ListAppsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
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

func (a ListAppsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Apps": reflect.TypeOf(App{}),
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

func (a StartAppRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type StopAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *StopAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopAppRequest) {
}

func (newState *StopAppRequest) SyncEffectiveFieldsDuringRead(existingState StopAppRequest) {
}

func (a StopAppRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Update an app
type UpdateAppRequest struct {
	App types.Object `tfsdk:"app" tf:"optional,object"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"-"`
}

func (newState *UpdateAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAppRequest) {
}

func (newState *UpdateAppRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAppRequest) {
}

func (a UpdateAppRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"App": reflect.TypeOf(App{}),
	}
}
