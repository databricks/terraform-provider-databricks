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
	// The active deployment of the app.
	ActiveDeployment *AppDeployment `tfsdk:"active_deployment" tf:"optional"`

	AppStatus *ApplicationStatus `tfsdk:"app_status" tf:"optional"`

	ComputeStatus *ComputeStatus `tfsdk:"compute_status" tf:"optional"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// The email of the user that created the app.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// The default workspace file system path of the source code from which app
	// deployment are created. This field tracks the workspace source code path
	// of the last active deployment.
	DefaultSourceCodePath types.String `tfsdk:"default_source_code_path" tf:"optional"`
	// The description of the app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name" tf:""`
	// The pending deployment of the app.
	PendingDeployment *AppDeployment `tfsdk:"pending_deployment" tf:"optional"`
	// Resources for the app.
	Resources []AppResource `tfsdk:"resources" tf:"optional"`

	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id" tf:"optional"`

	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
	// The email of the user that last updated the app.
	Updater types.String `tfsdk:"updater" tf:"optional"`
	// The URL of the app once it is deployed.
	Url types.String `tfsdk:"url" tf:"optional"`
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

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// The email of the user creates the deployment.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// The deployment artifacts for an app.
	DeploymentArtifacts *AppDeploymentArtifacts `tfsdk:"deployment_artifacts" tf:"optional"`
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
	Status *AppDeploymentStatus `tfsdk:"status" tf:"optional"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

type AppDeploymentArtifacts struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	SourceCodePath types.String `tfsdk:"source_code_path" tf:"optional"`
}

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	Message types.String `tfsdk:"message" tf:"optional"`
	// State of the deployment.
	State types.String `tfsdk:"state" tf:"optional"`
}

type AppPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

type AppPermissions struct {
	AccessControlList []AppAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type AppPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

type AppPermissionsRequest struct {
	AccessControlList []AppAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

type AppResource struct {
	// Description of the App Resource.
	Description types.String `tfsdk:"description" tf:"optional"`

	Job *AppResourceJob `tfsdk:"job" tf:"optional"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name" tf:""`

	Secret *AppResourceSecret `tfsdk:"secret" tf:"optional"`

	ServingEndpoint *AppResourceServingEndpoint `tfsdk:"serving_endpoint" tf:"optional"`

	SqlWarehouse *AppResourceSqlWarehouse `tfsdk:"sql_warehouse" tf:"optional"`
}

type AppResourceJob struct {
	// Id of the job to grant permission on.
	Id types.String `tfsdk:"id" tf:""`
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission types.String `tfsdk:"permission" tf:""`
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

type AppResourceServingEndpoint struct {
	// Name of the serving endpoint to grant permission on.
	Name types.String `tfsdk:"name" tf:""`
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission types.String `tfsdk:"permission" tf:""`
}

type AppResourceSqlWarehouse struct {
	// Id of the SQL warehouse to grant permission on.
	Id types.String `tfsdk:"id" tf:""`
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission types.String `tfsdk:"permission" tf:""`
}

type ApplicationStatus struct {
	// Application status message
	Message types.String `tfsdk:"message" tf:"optional"`
	// State of the application.
	State types.String `tfsdk:"state" tf:"optional"`
}

type ComputeStatus struct {
	// Compute status message
	Message types.String `tfsdk:"message" tf:"optional"`
	// State of the app compute.
	State types.String `tfsdk:"state" tf:"optional"`
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

type CreateAppRequest struct {
	// The description of the app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name" tf:""`
	// Resources for the app.
	Resources []AppResource `tfsdk:"resources" tf:"optional"`
}

// Delete an app
type DeleteAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

// Get an app deployment
type GetAppDeploymentRequest struct {
	// The name of the app.
	AppName types.String `tfsdk:"-"`
	// The unique id of the deployment.
	DeploymentId types.String `tfsdk:"-"`
}

// Get app permission levels
type GetAppPermissionLevelsRequest struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []AppPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get app permissions
type GetAppPermissionsRequest struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

// Get an app
type GetAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
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

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	AppDeployments []AppDeployment `tfsdk:"app_deployments" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List apps
type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-"`
}

type ListAppsResponse struct {
	Apps []App `tfsdk:"apps" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

type StartAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

type StopAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
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
