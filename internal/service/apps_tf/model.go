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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment []AppDeployment `tfsdk:"active_deployment" tf:"optional,object"`

	AppStatus []ApplicationStatus `tfsdk:"app_status" tf:"optional,object"`

	ComputeStatus []ComputeStatus `tfsdk:"compute_status" tf:"optional,object"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime          types.String `tfsdk:"create_time" tf:"optional"`
	EffectiveCreateTime types.String `tfsdk:"effective_create_time" tf:"computed,optional"`
	// The email of the user that created the app.
	Creator          types.String `tfsdk:"creator" tf:"optional"`
	EffectiveCreator types.String `tfsdk:"effective_creator" tf:"computed,optional"`
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
	PendingDeployment []AppDeployment `tfsdk:"pending_deployment" tf:"optional,object"`
	// Resources for the app.
	Resources []AppResource `tfsdk:"resources" tf:"optional"`

	ServicePrincipalId          types.Int64 `tfsdk:"service_principal_id" tf:"optional"`
	EffectiveServicePrincipalId types.Int64 `tfsdk:"effective_service_principal_id" tf:"computed,optional"`

	ServicePrincipalName          types.String `tfsdk:"service_principal_name" tf:"optional"`
	EffectiveServicePrincipalName types.String `tfsdk:"effective_service_principal_name" tf:"computed,optional"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime          types.String `tfsdk:"update_time" tf:"optional"`
	EffectiveUpdateTime types.String `tfsdk:"effective_update_time" tf:"computed,optional"`
	// The email of the user that last updated the app.
	Updater          types.String `tfsdk:"updater" tf:"optional"`
	EffectiveUpdater types.String `tfsdk:"effective_updater" tf:"computed,optional"`
	// The URL of the app once it is deployed.
	Url          types.String `tfsdk:"url" tf:"optional"`
	EffectiveUrl types.String `tfsdk:"effective_url" tf:"computed,optional"`
}

func (newState *App) SyncEffectiveFieldsDuringCreateOrUpdate(plan App) {
	newState.EffectiveCreateTime = newState.CreateTime
	newState.CreateTime = plan.CreateTime
	newState.EffectiveCreator = newState.Creator
	newState.Creator = plan.Creator
	newState.EffectiveServicePrincipalId = newState.ServicePrincipalId
	newState.ServicePrincipalId = plan.ServicePrincipalId
	newState.EffectiveServicePrincipalName = newState.ServicePrincipalName
	newState.ServicePrincipalName = plan.ServicePrincipalName
	newState.EffectiveUpdateTime = newState.UpdateTime
	newState.UpdateTime = plan.UpdateTime
	newState.EffectiveUpdater = newState.Updater
	newState.Updater = plan.Updater
	newState.EffectiveUrl = newState.Url
	newState.Url = plan.Url
}

func (newState *App) SyncEffectiveFieldsDuringRead(existingState App) {
	newState.EffectiveCreateTime = existingState.EffectiveCreateTime
	if existingState.EffectiveCreateTime.ValueString() == newState.CreateTime.ValueString() {
		newState.CreateTime = existingState.CreateTime
	}
	newState.EffectiveCreator = existingState.EffectiveCreator
	if existingState.EffectiveCreator.ValueString() == newState.Creator.ValueString() {
		newState.Creator = existingState.Creator
	}
	newState.EffectiveServicePrincipalId = existingState.EffectiveServicePrincipalId
	if existingState.EffectiveServicePrincipalId.ValueInt64() == newState.ServicePrincipalId.ValueInt64() {
		newState.ServicePrincipalId = existingState.ServicePrincipalId
	}
	newState.EffectiveServicePrincipalName = existingState.EffectiveServicePrincipalName
	if existingState.EffectiveServicePrincipalName.ValueString() == newState.ServicePrincipalName.ValueString() {
		newState.ServicePrincipalName = existingState.ServicePrincipalName
	}
	newState.EffectiveUpdateTime = existingState.EffectiveUpdateTime
	if existingState.EffectiveUpdateTime.ValueString() == newState.UpdateTime.ValueString() {
		newState.UpdateTime = existingState.UpdateTime
	}
	newState.EffectiveUpdater = existingState.EffectiveUpdater
	if existingState.EffectiveUpdater.ValueString() == newState.Updater.ValueString() {
		newState.Updater = existingState.Updater
	}
	newState.EffectiveUrl = existingState.EffectiveUrl
	if existingState.EffectiveUrl.ValueString() == newState.Url.ValueString() {
		newState.Url = existingState.Url
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

type AppAccessControlResponse struct {
	// All permissions.
	AllPermissions []AppPermission `tfsdk:"all_permissions" tf:"optional"`
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

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime          types.String `tfsdk:"create_time" tf:"optional"`
	EffectiveCreateTime types.String `tfsdk:"effective_create_time" tf:"computed,optional"`
	// The email of the user creates the deployment.
	Creator          types.String `tfsdk:"creator" tf:"optional"`
	EffectiveCreator types.String `tfsdk:"effective_creator" tf:"computed,optional"`
	// The deployment artifacts for an app.
	DeploymentArtifacts []AppDeploymentArtifacts `tfsdk:"deployment_artifacts" tf:"optional,object"`
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
	Status []AppDeploymentStatus `tfsdk:"status" tf:"optional,object"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime          types.String `tfsdk:"update_time" tf:"optional"`
	EffectiveUpdateTime types.String `tfsdk:"effective_update_time" tf:"computed,optional"`
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeployment) {
	newState.EffectiveCreateTime = newState.CreateTime
	newState.CreateTime = plan.CreateTime
	newState.EffectiveCreator = newState.Creator
	newState.Creator = plan.Creator
	newState.EffectiveUpdateTime = newState.UpdateTime
	newState.UpdateTime = plan.UpdateTime
}

func (newState *AppDeployment) SyncEffectiveFieldsDuringRead(existingState AppDeployment) {
	newState.EffectiveCreateTime = existingState.EffectiveCreateTime
	if existingState.EffectiveCreateTime.ValueString() == newState.CreateTime.ValueString() {
		newState.CreateTime = existingState.CreateTime
	}
	newState.EffectiveCreator = existingState.EffectiveCreator
	if existingState.EffectiveCreator.ValueString() == newState.Creator.ValueString() {
		newState.Creator = existingState.Creator
	}
	newState.EffectiveUpdateTime = existingState.EffectiveUpdateTime
	if existingState.EffectiveUpdateTime.ValueString() == newState.UpdateTime.ValueString() {
		newState.UpdateTime = existingState.UpdateTime
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

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	Message          types.String `tfsdk:"message" tf:"optional"`
	EffectiveMessage types.String `tfsdk:"effective_message" tf:"computed,optional"`
	// State of the deployment.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *AppDeploymentStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppDeploymentStatus) {
	newState.EffectiveMessage = newState.Message
	newState.Message = plan.Message
}

func (newState *AppDeploymentStatus) SyncEffectiveFieldsDuringRead(existingState AppDeploymentStatus) {
	newState.EffectiveMessage = existingState.EffectiveMessage
	if existingState.EffectiveMessage.ValueString() == newState.Message.ValueString() {
		newState.Message = existingState.Message
	}
}

type AppPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *AppPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermission) {
}

func (newState *AppPermission) SyncEffectiveFieldsDuringRead(existingState AppPermission) {
}

type AppPermissions struct {
	AccessControlList []AppAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *AppPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissions) {
}

func (newState *AppPermissions) SyncEffectiveFieldsDuringRead(existingState AppPermissions) {
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

type AppPermissionsRequest struct {
	AccessControlList []AppAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (newState *AppPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppPermissionsRequest) {
}

func (newState *AppPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState AppPermissionsRequest) {
}

type AppResource struct {
	// Description of the App Resource.
	Description types.String `tfsdk:"description" tf:"optional"`

	Job []AppResourceJob `tfsdk:"job" tf:"optional,object"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name" tf:""`

	Secret []AppResourceSecret `tfsdk:"secret" tf:"optional,object"`

	ServingEndpoint []AppResourceServingEndpoint `tfsdk:"serving_endpoint" tf:"optional,object"`

	SqlWarehouse []AppResourceSqlWarehouse `tfsdk:"sql_warehouse" tf:"optional,object"`
}

func (newState *AppResource) SyncEffectiveFieldsDuringCreateOrUpdate(plan AppResource) {
}

func (newState *AppResource) SyncEffectiveFieldsDuringRead(existingState AppResource) {
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

type ApplicationStatus struct {
	// Application status message
	Message          types.String `tfsdk:"message" tf:"optional"`
	EffectiveMessage types.String `tfsdk:"effective_message" tf:"computed,optional"`
	// State of the application.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *ApplicationStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApplicationStatus) {
	newState.EffectiveMessage = newState.Message
	newState.Message = plan.Message
}

func (newState *ApplicationStatus) SyncEffectiveFieldsDuringRead(existingState ApplicationStatus) {
	newState.EffectiveMessage = existingState.EffectiveMessage
	if existingState.EffectiveMessage.ValueString() == newState.Message.ValueString() {
		newState.Message = existingState.Message
	}
}

type ComputeStatus struct {
	// Compute status message
	Message          types.String `tfsdk:"message" tf:"optional"`
	EffectiveMessage types.String `tfsdk:"effective_message" tf:"computed,optional"`
	// State of the app compute.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *ComputeStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComputeStatus) {
	newState.EffectiveMessage = newState.Message
	newState.Message = plan.Message
}

func (newState *ComputeStatus) SyncEffectiveFieldsDuringRead(existingState ComputeStatus) {
	newState.EffectiveMessage = existingState.EffectiveMessage
	if existingState.EffectiveMessage.ValueString() == newState.Message.ValueString() {
		newState.Message = existingState.Message
	}
}

type CreateAppDeploymentRequest struct {
	// The name of the app.
	AppName types.String `tfsdk:"-"`
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
}

func (newState *CreateAppDeploymentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAppDeploymentRequest) {
}

func (newState *CreateAppDeploymentRequest) SyncEffectiveFieldsDuringRead(existingState CreateAppDeploymentRequest) {
}

type CreateAppRequest struct {
	// The description of the app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name" tf:""`
	// Resources for the app.
	Resources []AppResource `tfsdk:"resources" tf:"optional"`
}

func (newState *CreateAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAppRequest) {
}

func (newState *CreateAppRequest) SyncEffectiveFieldsDuringRead(existingState CreateAppRequest) {
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

// Get app permission levels
type GetAppPermissionLevelsRequest struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (newState *GetAppPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionLevelsRequest) {
}

func (newState *GetAppPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionLevelsRequest) {
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []AppPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetAppPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppPermissionLevelsResponse) {
}

func (newState *GetAppPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetAppPermissionLevelsResponse) {
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

// Get an app
type GetAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *GetAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAppRequest) {
}

func (newState *GetAppRequest) SyncEffectiveFieldsDuringRead(existingState GetAppRequest) {
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

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	AppDeployments []AppDeployment `tfsdk:"app_deployments" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListAppDeploymentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppDeploymentsResponse) {
}

func (newState *ListAppDeploymentsResponse) SyncEffectiveFieldsDuringRead(existingState ListAppDeploymentsResponse) {
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

type ListAppsResponse struct {
	Apps []App `tfsdk:"apps" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListAppsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAppsResponse) {
}

func (newState *ListAppsResponse) SyncEffectiveFieldsDuringRead(existingState ListAppsResponse) {
}

type StartAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *StartAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartAppRequest) {
}

func (newState *StartAppRequest) SyncEffectiveFieldsDuringRead(existingState StartAppRequest) {
}

type StopAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (newState *StopAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopAppRequest) {
}

func (newState *StopAppRequest) SyncEffectiveFieldsDuringRead(existingState StopAppRequest) {
}

type UpdateAppRequest struct {
	// The description of the app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name" tf:""`
	// Resources for the app.
	Resources []AppResource `tfsdk:"resources" tf:"optional"`
}

func (newState *UpdateAppRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAppRequest) {
}

func (newState *UpdateAppRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAppRequest) {
}
