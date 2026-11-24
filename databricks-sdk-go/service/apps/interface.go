// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"context"
)

// Apps run directly on a customer's Databricks instance, integrate with their
// data, use and extend Databricks services, and enable users to interact
// through single sign-on.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AppsService interface {

	// Creates a new app.
	Create(ctx context.Context, request CreateAppRequest) (*App, error)

	// Creates a new app space.
	CreateSpace(ctx context.Context, request CreateSpaceRequest) (*Operation, error)

	// Creates an app update and starts the update process. The update process
	// is asynchronous and the status of the update can be checked with the
	// GetAppUpdate method.
	CreateUpdate(ctx context.Context, request AsyncUpdateAppRequest) (*AppUpdate, error)

	// Deletes an app and the dedicated service principal created for it.
	Delete(ctx context.Context, request DeleteAppRequest) (*App, error)

	// Deletes the thumbnail for an app.
	DeleteAppThumbnail(ctx context.Context, request DeleteAppThumbnailRequest) error

	// Deletes an app space.
	DeleteSpace(ctx context.Context, request DeleteSpaceRequest) (*Operation, error)

	// Creates an app deployment for the app with the supplied name. The call
	// returns while the deployment is still `IN_PROGRESS`; poll the deployment
	// until it reaches a terminal state. The request fails with
	// RESOURCE_CONFLICT if another deployment of the same app is still pending
	// and has not timed out.
	Deploy(ctx context.Context, request CreateAppDeploymentRequest) (*AppDeployment, error)

	// Retrieves information for the app with the supplied name.
	Get(ctx context.Context, request GetAppRequest) (*App, error)

	// Retrieves a single security scan finding for an app deployment.
	GetAppDeploymentScanFinding(ctx context.Context, request GetAppDeploymentScanFindingRequest) (*AppDeploymentScanFinding, error)

	// Retrieves information for the app deployment with the supplied name and
	// deployment id.
	GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error)

	// Gets the permissions of an app. Apps can inherit permissions from their
	// root object.
	GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error)

	// Retrieves information for the app space with the supplied name.
	GetSpace(ctx context.Context, request GetSpaceRequest) (*Space, error)

	// Gets the status of an app space update operation.
	GetSpaceOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Gets the status of an app update.
	GetUpdate(ctx context.Context, request GetAppUpdateRequest) (*AppUpdate, error)

	// Lists all apps in the workspace.
	List(ctx context.Context, request ListAppsRequest) (*ListAppsResponse, error)

	// Lists the security scan findings for an app deployment.
	ListAppDeploymentScanFindings(ctx context.Context, request ListAppDeploymentScanFindingsRequest) (*ListAppDeploymentScanFindingsResponse, error)

	// Lists all app deployments for the app with the supplied name. The history
	// holds only the app's most recent deployments.
	ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error)

	// Lists all app spaces in the workspace.
	ListSpaces(ctx context.Context, request ListSpacesRequest) (*ListSpacesResponse, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error)

	// Start the last active deployment of the app in the workspace.
	Start(ctx context.Context, request StartAppRequest) (*App, error)

	// Stops the active deployment of the app in the workspace.
	Stop(ctx context.Context, request StopAppRequest) (*App, error)

	// Updates the app with the supplied name. This is a full replacement:
	// fields omitted from the request are cleared, so send the complete app.
	Update(ctx context.Context, request UpdateAppRequest) (*App, error)

	// Updates the triage status (resolved / ignored / in-progress) of a single
	// scan finding.
	UpdateAppDeploymentScanFinding(ctx context.Context, request UpdateAppDeploymentScanFindingRequest) (*AppDeploymentScanFinding, error)

	// Updates the thumbnail for an app.
	UpdateAppThumbnail(ctx context.Context, request UpdateAppThumbnailRequest) (*AppThumbnail, error)

	// Updates the permissions on an app. Apps can inherit permissions from
	// their root object.
	UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error)

	// Updates an app space. The update process is asynchronous and the status
	// of the update can be checked with the GetSpaceOperation method.
	UpdateSpace(ctx context.Context, request UpdateSpaceRequest) (*Operation, error)
}

// Apps Settings manage the settings for the Apps service on a customer's
// Databricks instance.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AppsSettingsService interface {

	// Creates a custom template.
	CreateCustomTemplate(ctx context.Context, request CreateCustomTemplateRequest) (*CustomTemplate, error)

	// Deletes the custom template with the specified name.
	DeleteCustomTemplate(ctx context.Context, request DeleteCustomTemplateRequest) (*CustomTemplate, error)

	// Gets the custom template with the specified name.
	GetCustomTemplate(ctx context.Context, request GetCustomTemplateRequest) (*CustomTemplate, error)

	// Lists all custom templates in the workspace.
	ListCustomTemplates(ctx context.Context, request ListCustomTemplatesRequest) (*ListCustomTemplatesResponse, error)

	// Updates the custom template with the specified name. Note that the
	// template name cannot be updated.
	UpdateCustomTemplate(ctx context.Context, request UpdateCustomTemplateRequest) (*CustomTemplate, error)
}
