// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Apps, Apps Settings, etc.
package apps

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/experimental/api"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type AppsInterface interface {

	// WaitGetAppActive repeatedly calls [AppsAPI.Get] and waits to reach ACTIVE state
	WaitGetAppActive(ctx context.Context, name string,
		timeout time.Duration, callback func(*App)) (*App, error)

	// WaitGetUpdateAppSucceeded repeatedly calls [AppsAPI.GetUpdate] and waits to reach SUCCEEDED state
	WaitGetUpdateAppSucceeded(ctx context.Context, appName string,
		timeout time.Duration, callback func(*AppUpdate)) (*AppUpdate, error)

	// WaitGetDeploymentAppSucceeded repeatedly calls [AppsAPI.GetDeployment] and waits to reach SUCCEEDED state
	WaitGetDeploymentAppSucceeded(ctx context.Context, appName string, deploymentId string,
		timeout time.Duration, callback func(*AppDeployment)) (*AppDeployment, error)

	// WaitGetAppStopped repeatedly calls [AppsAPI.Get] and waits to reach STOPPED state
	WaitGetAppStopped(ctx context.Context, name string,
		timeout time.Duration, callback func(*App)) (*App, error)

	// Creates a new app.
	Create(ctx context.Context, createAppRequest CreateAppRequest) (*WaitGetAppActive[App], error)

	// Calls [AppsAPIInterface.Create] and waits to reach ACTIVE state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[App](60*time.Minute) functional option.
	//
	// Deprecated: use [AppsAPIInterface.Create].Get() or [AppsAPIInterface.WaitGetAppActive]
	CreateAndWait(ctx context.Context, createAppRequest CreateAppRequest, options ...retries.Option[App]) (*App, error)

	CreateSpace(ctx context.Context, request CreateSpaceRequest) (CreateSpaceOperationInterface, error)

	// Creates an app update and starts the update process. The update process is
	// asynchronous and the status of the update can be checked with the
	// GetAppUpdate method.
	CreateUpdate(ctx context.Context, asyncUpdateAppRequest AsyncUpdateAppRequest) (*WaitGetUpdateAppSucceeded[AppUpdate], error)

	// Calls [AppsAPIInterface.CreateUpdate] and waits to reach SUCCEEDED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[AppUpdate](60*time.Minute) functional option.
	//
	// Deprecated: use [AppsAPIInterface.CreateUpdate].Get() or [AppsAPIInterface.WaitGetUpdateAppSucceeded]
	CreateUpdateAndWait(ctx context.Context, asyncUpdateAppRequest AsyncUpdateAppRequest, options ...retries.Option[AppUpdate]) (*AppUpdate, error)

	// Deletes an app and the dedicated service principal created for it.
	Delete(ctx context.Context, request DeleteAppRequest) (*App, error)

	// Deletes an app and the dedicated service principal created for it.
	DeleteByName(ctx context.Context, name string) (*App, error)

	// Deletes the thumbnail for an app.
	DeleteAppThumbnail(ctx context.Context, request DeleteAppThumbnailRequest) error

	DeleteSpace(ctx context.Context, request DeleteSpaceRequest) (DeleteSpaceOperationInterface, error)

	// Creates an app deployment for the app with the supplied name. The call
	// returns while the deployment is still `IN_PROGRESS`; poll the deployment
	// until it reaches a terminal state. The request fails with RESOURCE_CONFLICT
	// if another deployment of the same app is still pending and has not timed out.
	Deploy(ctx context.Context, createAppDeploymentRequest CreateAppDeploymentRequest) (*WaitGetDeploymentAppSucceeded[AppDeployment], error)

	// Calls [AppsAPIInterface.Deploy] and waits to reach SUCCEEDED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[AppDeployment](60*time.Minute) functional option.
	//
	// Deprecated: use [AppsAPIInterface.Deploy].Get() or [AppsAPIInterface.WaitGetDeploymentAppSucceeded]
	DeployAndWait(ctx context.Context, createAppDeploymentRequest CreateAppDeploymentRequest, options ...retries.Option[AppDeployment]) (*AppDeployment, error)

	// Retrieves information for the app with the supplied name.
	Get(ctx context.Context, request GetAppRequest) (*App, error)

	// Retrieves information for the app with the supplied name.
	GetByName(ctx context.Context, name string) (*App, error)

	// Retrieves a single security scan finding for an app deployment.
	GetAppDeploymentScanFinding(ctx context.Context, request GetAppDeploymentScanFindingRequest) (*AppDeploymentScanFinding, error)

	// Retrieves information for the app deployment with the supplied name and
	// deployment id.
	GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error)

	// Retrieves information for the app deployment with the supplied name and
	// deployment id.
	GetDeploymentByAppNameAndDeploymentId(ctx context.Context, appName string, deploymentId string) (*AppDeployment, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevelsByAppName(ctx context.Context, appName string) (*GetAppPermissionLevelsResponse, error)

	// Gets the permissions of an app. Apps can inherit permissions from their root
	// object.
	GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error)

	// Gets the permissions of an app. Apps can inherit permissions from their root
	// object.
	GetPermissionsByAppName(ctx context.Context, appName string) (*AppPermissions, error)

	// Retrieves information for the app space with the supplied name.
	GetSpace(ctx context.Context, request GetSpaceRequest) (*Space, error)

	// Gets the status of an app space update operation.
	GetSpaceOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Gets the status of an app update.
	GetUpdate(ctx context.Context, request GetAppUpdateRequest) (*AppUpdate, error)

	// Lists all apps in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListAppsRequest) listing.Iterator[App]

	// Lists all apps in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListAppsRequest) ([]App, error)

	// Lists the security scan findings for an app deployment.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAppDeploymentScanFindings(ctx context.Context, request ListAppDeploymentScanFindingsRequest) listing.Iterator[AppDeploymentScanFinding]

	// Lists the security scan findings for an app deployment.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAppDeploymentScanFindingsAll(ctx context.Context, request ListAppDeploymentScanFindingsRequest) ([]AppDeploymentScanFinding, error)

	// Lists all app deployments for the app with the supplied name. The history
	// holds only the app's most recent deployments.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) listing.Iterator[AppDeployment]

	// Lists all app deployments for the app with the supplied name. The history
	// holds only the app's most recent deployments.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDeploymentsAll(ctx context.Context, request ListAppDeploymentsRequest) ([]AppDeployment, error)

	// Lists all app deployments for the app with the supplied name. The history
	// holds only the app's most recent deployments.
	ListDeploymentsByAppName(ctx context.Context, appName string) (*ListAppDeploymentsResponse, error)

	// Lists all app spaces in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSpaces(ctx context.Context, request ListSpacesRequest) listing.Iterator[Space]

	// Lists all app spaces in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSpacesAll(ctx context.Context, request ListSpacesRequest) ([]Space, error)

	// Sets permissions on an object, replacing existing permissions if they exist.
	// Deletes all direct permissions if none are specified. Objects can inherit
	// permissions from their root object.
	SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error)

	// Start the last active deployment of the app in the workspace.
	Start(ctx context.Context, startAppRequest StartAppRequest) (*WaitGetAppActive[App], error)

	// Calls [AppsAPIInterface.Start] and waits to reach ACTIVE state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[App](60*time.Minute) functional option.
	//
	// Deprecated: use [AppsAPIInterface.Start].Get() or [AppsAPIInterface.WaitGetAppActive]
	StartAndWait(ctx context.Context, startAppRequest StartAppRequest, options ...retries.Option[App]) (*App, error)

	// Stops the active deployment of the app in the workspace.
	Stop(ctx context.Context, stopAppRequest StopAppRequest) (*WaitGetAppStopped[App], error)

	// Calls [AppsAPIInterface.Stop] and waits to reach STOPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[App](60*time.Minute) functional option.
	//
	// Deprecated: use [AppsAPIInterface.Stop].Get() or [AppsAPIInterface.WaitGetAppStopped]
	StopAndWait(ctx context.Context, stopAppRequest StopAppRequest, options ...retries.Option[App]) (*App, error)

	// Updates the app with the supplied name. This is a full replacement: fields
	// omitted from the request are cleared, so send the complete app.
	Update(ctx context.Context, request UpdateAppRequest) (*App, error)

	// Updates the triage status (resolved / ignored / in-progress) of a single scan
	// finding.
	UpdateAppDeploymentScanFinding(ctx context.Context, request UpdateAppDeploymentScanFindingRequest) (*AppDeploymentScanFinding, error)

	// Updates the thumbnail for an app.
	UpdateAppThumbnail(ctx context.Context, request UpdateAppThumbnailRequest) (*AppThumbnail, error)

	// Updates the permissions on an app. Apps can inherit permissions from their
	// root object.
	UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error)

	UpdateSpace(ctx context.Context, request UpdateSpaceRequest) (UpdateSpaceOperationInterface, error)
}

func NewApps(client *client.DatabricksClient) *AppsAPI {
	return &AppsAPI{
		appsImpl: appsImpl{
			client: client,
		},
	}
}

// Apps run directly on a customer's Databricks instance, integrate with their
// data, use and extend Databricks services, and enable users to interact
// through single sign-on.
type AppsAPI struct {
	appsImpl
}

func (a *AppsAPI) CreateSpace(ctx context.Context, request CreateSpaceRequest) (CreateSpaceOperationInterface, error) {
	operation, err := a.appsImpl.CreateSpace(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createSpaceOperation{
		impl:      &a.appsImpl,
		operation: operation,
	}, nil
}

type CreateSpaceOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Space, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*Space, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createSpaceOperation struct {
	impl      *appsImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createSpaceOperation) Wait(ctx context.Context, opts ...api.Option) (*Space, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Space
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetSpaceOperation(ctx, GetOperationRequest{
			Name: a.operation.Name,
		})
		if err != nil {
			return err
		}

		// Update local operation state
		a.operation = operation

		if !operation.Done {
			return errOperationInProgress
		}

		if operation.Error != nil {
			var errorMsg string
			if operation.Error.Message != "" {
				errorMsg = operation.Error.Message
			} else {
				errorMsg = "unknown error"
			}

			if operation.Error.ErrorCode != "" {
				errorMsg = fmt.Sprintf("[%s] %s", operation.Error.ErrorCode, errorMsg)
			}

			return fmt.Errorf("operation failed: %s", errorMsg)
		}

		// Operation completed successfully, unmarshal response
		if operation.Response == nil {
			return fmt.Errorf("operation completed but no response available")
		}

		var space Space
		err = json.Unmarshal(operation.Response, &space)
		if err != nil {
			return fmt.Errorf("failed to unmarshal space response: %w", err)
		}

		result = &space

		return nil
	}

	// Create a retrier that retries on errOperationInProgress with exponential backoff.
	retrier := api.RetryOn(api.BackoffPolicy{}, func(err error) bool {
		return errors.Is(err, errOperationInProgress)
	})

	// Add default retrier.
	defaultOpts := []api.Option{
		api.WithRetrier(func() api.Retrier { return retrier }),
	}
	allOpts := append(defaultOpts, opts...)

	err := api.Execute(ctx, call, allOpts...)

	if err != nil {
		return nil, err
	}
	return result, nil

}

// Name returns the name of the long-running operation. The name is assigned
// by the server and is unique within the service from which the operation is created.
func (a *createSpaceOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createSpaceOperation) Metadata() (*Space, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata Space
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createSpaceOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetSpaceOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *AppsAPI) DeleteSpace(ctx context.Context, request DeleteSpaceRequest) (DeleteSpaceOperationInterface, error) {
	operation, err := a.appsImpl.DeleteSpace(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteSpaceOperation{
		impl:      &a.appsImpl,
		operation: operation,
	}, nil
}

type DeleteSpaceOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) error

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*Space, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteSpaceOperation struct {
	impl      *appsImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteSpaceOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetSpaceOperation(ctx, GetOperationRequest{
			Name: a.operation.Name,
		})
		if err != nil {
			return err
		}

		// Update local operation state
		a.operation = operation

		if !operation.Done {
			return errOperationInProgress
		}

		if operation.Error != nil {
			var errorMsg string
			if operation.Error.Message != "" {
				errorMsg = operation.Error.Message
			} else {
				errorMsg = "unknown error"
			}

			if operation.Error.ErrorCode != "" {
				errorMsg = fmt.Sprintf("[%s] %s", operation.Error.ErrorCode, errorMsg)
			}

			return fmt.Errorf("operation failed: %s", errorMsg)
		}

		// Operation completed successfully, unmarshal response
		if operation.Response == nil {
			return fmt.Errorf("operation completed but no response available")
		}

		return nil
	}

	// Create a retrier that retries on errOperationInProgress with exponential backoff.
	retrier := api.RetryOn(api.BackoffPolicy{}, func(err error) bool {
		return errors.Is(err, errOperationInProgress)
	})

	// Add default retrier.
	defaultOpts := []api.Option{
		api.WithRetrier(func() api.Retrier { return retrier }),
	}
	allOpts := append(defaultOpts, opts...)

	err := api.Execute(ctx, call, allOpts...)

	return err

}

// Name returns the name of the long-running operation. The name is assigned
// by the server and is unique within the service from which the operation is created.
func (a *deleteSpaceOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteSpaceOperation) Metadata() (*Space, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata Space
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteSpaceOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetSpaceOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *AppsAPI) UpdateSpace(ctx context.Context, request UpdateSpaceRequest) (UpdateSpaceOperationInterface, error) {
	operation, err := a.appsImpl.UpdateSpace(ctx, request)
	if err != nil {
		return nil, err
	}
	return &updateSpaceOperation{
		impl:      &a.appsImpl,
		operation: operation,
	}, nil
}

type UpdateSpaceOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Space, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*SpaceUpdate, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type updateSpaceOperation struct {
	impl      *appsImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *updateSpaceOperation) Wait(ctx context.Context, opts ...api.Option) (*Space, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Space
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetSpaceOperation(ctx, GetOperationRequest{
			Name: a.operation.Name,
		})
		if err != nil {
			return err
		}

		// Update local operation state
		a.operation = operation

		if !operation.Done {
			return errOperationInProgress
		}

		if operation.Error != nil {
			var errorMsg string
			if operation.Error.Message != "" {
				errorMsg = operation.Error.Message
			} else {
				errorMsg = "unknown error"
			}

			if operation.Error.ErrorCode != "" {
				errorMsg = fmt.Sprintf("[%s] %s", operation.Error.ErrorCode, errorMsg)
			}

			return fmt.Errorf("operation failed: %s", errorMsg)
		}

		// Operation completed successfully, unmarshal response
		if operation.Response == nil {
			return fmt.Errorf("operation completed but no response available")
		}

		var space Space
		err = json.Unmarshal(operation.Response, &space)
		if err != nil {
			return fmt.Errorf("failed to unmarshal space response: %w", err)
		}

		result = &space

		return nil
	}

	// Create a retrier that retries on errOperationInProgress with exponential backoff.
	retrier := api.RetryOn(api.BackoffPolicy{}, func(err error) bool {
		return errors.Is(err, errOperationInProgress)
	})

	// Add default retrier.
	defaultOpts := []api.Option{
		api.WithRetrier(func() api.Retrier { return retrier }),
	}
	allOpts := append(defaultOpts, opts...)

	err := api.Execute(ctx, call, allOpts...)

	if err != nil {
		return nil, err
	}
	return result, nil

}

// Name returns the name of the long-running operation. The name is assigned
// by the server and is unique within the service from which the operation is created.
func (a *updateSpaceOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *updateSpaceOperation) Metadata() (*SpaceUpdate, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata SpaceUpdate
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *updateSpaceOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetSpaceOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

// WaitGetAppActive repeatedly calls [AppsAPI.Get] and waits to reach ACTIVE state
func (a *AppsAPI) WaitGetAppActive(ctx context.Context, name string,
	timeout time.Duration, callback func(*App)) (*App, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[App](ctx, timeout, func() (*App, *retries.Err) {
		app, err := a.Get(ctx, GetAppRequest{
			Name: name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(app)
		}
		status := app.ComputeStatus.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if app.ComputeStatus != nil {
			statusMessage = app.ComputeStatus.Message
		}
		switch status {
		case ComputeStateActive: // target state
			return app, nil
		case ComputeStateError, ComputeStateStopped:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ComputeStateActive, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetAppActive is a wrapper that calls [AppsAPI.WaitGetAppActive] and waits to reach ACTIVE state.
type WaitGetAppActive[R any] struct {
	Response *R
	Name     string `json:"name"`
	Poll     func(time.Duration, func(*App)) (*App, error)
	callback func(*App)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetAppActive[R]) OnProgress(callback func(*App)) *WaitGetAppActive[R] {
	w.callback = callback
	return w
}

// Get the App with the default timeout of 20 minutes.
func (w *WaitGetAppActive[R]) Get() (*App, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the App with custom timeout.
func (w *WaitGetAppActive[R]) GetWithTimeout(timeout time.Duration) (*App, error) {
	return w.Poll(timeout, w.callback)
}

// WaitGetUpdateAppSucceeded repeatedly calls [AppsAPI.GetUpdate] and waits to reach SUCCEEDED state
func (a *AppsAPI) WaitGetUpdateAppSucceeded(ctx context.Context, appName string,
	timeout time.Duration, callback func(*AppUpdate)) (*AppUpdate, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[AppUpdate](ctx, timeout, func() (*AppUpdate, *retries.Err) {
		appUpdate, err := a.GetUpdate(ctx, GetAppUpdateRequest{
			AppName: appName,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(appUpdate)
		}
		status := appUpdate.Status.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if appUpdate.Status != nil {
			statusMessage = appUpdate.Status.Message
		}
		switch status {
		case AppUpdateUpdateStatusUpdateStateSucceeded: // target state
			return appUpdate, nil
		case AppUpdateUpdateStatusUpdateStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				AppUpdateUpdateStatusUpdateStateSucceeded, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetUpdateAppSucceeded is a wrapper that calls [AppsAPI.WaitGetUpdateAppSucceeded] and waits to reach SUCCEEDED state.
type WaitGetUpdateAppSucceeded[R any] struct {
	Response *R
	AppName  string `json:"app_name"`
	Poll     func(time.Duration, func(*AppUpdate)) (*AppUpdate, error)
	callback func(*AppUpdate)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetUpdateAppSucceeded[R]) OnProgress(callback func(*AppUpdate)) *WaitGetUpdateAppSucceeded[R] {
	w.callback = callback
	return w
}

// Get the AppUpdate with the default timeout of 20 minutes.
func (w *WaitGetUpdateAppSucceeded[R]) Get() (*AppUpdate, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the AppUpdate with custom timeout.
func (w *WaitGetUpdateAppSucceeded[R]) GetWithTimeout(timeout time.Duration) (*AppUpdate, error) {
	return w.Poll(timeout, w.callback)
}

// WaitGetDeploymentAppSucceeded repeatedly calls [AppsAPI.GetDeployment] and waits to reach SUCCEEDED state
func (a *AppsAPI) WaitGetDeploymentAppSucceeded(ctx context.Context, appName string, deploymentId string,
	timeout time.Duration, callback func(*AppDeployment)) (*AppDeployment, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[AppDeployment](ctx, timeout, func() (*AppDeployment, *retries.Err) {
		appDeployment, err := a.GetDeployment(ctx, GetAppDeploymentRequest{
			AppName:      appName,
			DeploymentId: deploymentId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(appDeployment)
		}
		status := appDeployment.Status.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if appDeployment.Status != nil {
			statusMessage = appDeployment.Status.Message
		}
		switch status {
		case AppDeploymentStateSucceeded: // target state
			return appDeployment, nil
		case AppDeploymentStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				AppDeploymentStateSucceeded, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetDeploymentAppSucceeded is a wrapper that calls [AppsAPI.WaitGetDeploymentAppSucceeded] and waits to reach SUCCEEDED state.
type WaitGetDeploymentAppSucceeded[R any] struct {
	Response     *R
	AppName      string `json:"app_name"`
	DeploymentId string `json:"deployment_id"`
	Poll         func(time.Duration, func(*AppDeployment)) (*AppDeployment, error)
	callback     func(*AppDeployment)
	timeout      time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetDeploymentAppSucceeded[R]) OnProgress(callback func(*AppDeployment)) *WaitGetDeploymentAppSucceeded[R] {
	w.callback = callback
	return w
}

// Get the AppDeployment with the default timeout of 20 minutes.
func (w *WaitGetDeploymentAppSucceeded[R]) Get() (*AppDeployment, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the AppDeployment with custom timeout.
func (w *WaitGetDeploymentAppSucceeded[R]) GetWithTimeout(timeout time.Duration) (*AppDeployment, error) {
	return w.Poll(timeout, w.callback)
}

// WaitGetAppStopped repeatedly calls [AppsAPI.Get] and waits to reach STOPPED state
func (a *AppsAPI) WaitGetAppStopped(ctx context.Context, name string,
	timeout time.Duration, callback func(*App)) (*App, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[App](ctx, timeout, func() (*App, *retries.Err) {
		app, err := a.Get(ctx, GetAppRequest{
			Name: name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(app)
		}
		status := app.ComputeStatus.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if app.ComputeStatus != nil {
			statusMessage = app.ComputeStatus.Message
		}
		switch status {
		case ComputeStateStopped: // target state
			return app, nil
		case ComputeStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ComputeStateStopped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetAppStopped is a wrapper that calls [AppsAPI.WaitGetAppStopped] and waits to reach STOPPED state.
type WaitGetAppStopped[R any] struct {
	Response *R
	Name     string `json:"name"`
	Poll     func(time.Duration, func(*App)) (*App, error)
	callback func(*App)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetAppStopped[R]) OnProgress(callback func(*App)) *WaitGetAppStopped[R] {
	w.callback = callback
	return w
}

// Get the App with the default timeout of 20 minutes.
func (w *WaitGetAppStopped[R]) Get() (*App, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the App with custom timeout.
func (w *WaitGetAppStopped[R]) GetWithTimeout(timeout time.Duration) (*App, error) {
	return w.Poll(timeout, w.callback)
}

// Creates a new app.
func (a *AppsAPI) Create(ctx context.Context, createAppRequest CreateAppRequest) (*WaitGetAppActive[App], error) {
	app, err := a.appsImpl.Create(ctx, createAppRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetAppActive[App]{
		Response: app,
		Name:     app.Name,
		Poll: func(timeout time.Duration, callback func(*App)) (*App, error) {
			return a.WaitGetAppActive(ctx, app.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [AppsAPI.Create] and waits to reach ACTIVE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[App](60*time.Minute) functional option.
//
// Deprecated: use [AppsAPI.Create].Get() or [AppsAPI.WaitGetAppActive]
func (a *AppsAPI) CreateAndWait(ctx context.Context, createAppRequest CreateAppRequest, options ...retries.Option[App]) (*App, error) {
	wait, err := a.Create(ctx, createAppRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[App]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *App) {
		for _, o := range options {
			o(&retries.Info[App]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Creates an app update and starts the update process. The update process is
// asynchronous and the status of the update can be checked with the
// GetAppUpdate method.
func (a *AppsAPI) CreateUpdate(ctx context.Context, asyncUpdateAppRequest AsyncUpdateAppRequest) (*WaitGetUpdateAppSucceeded[AppUpdate], error) {
	appUpdate, err := a.appsImpl.CreateUpdate(ctx, asyncUpdateAppRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetUpdateAppSucceeded[AppUpdate]{
		Response: appUpdate,
		AppName:  asyncUpdateAppRequest.AppName,
		Poll: func(timeout time.Duration, callback func(*AppUpdate)) (*AppUpdate, error) {
			return a.WaitGetUpdateAppSucceeded(ctx, asyncUpdateAppRequest.AppName, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [AppsAPI.CreateUpdate] and waits to reach SUCCEEDED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[AppUpdate](60*time.Minute) functional option.
//
// Deprecated: use [AppsAPI.CreateUpdate].Get() or [AppsAPI.WaitGetUpdateAppSucceeded]
func (a *AppsAPI) CreateUpdateAndWait(ctx context.Context, asyncUpdateAppRequest AsyncUpdateAppRequest, options ...retries.Option[AppUpdate]) (*AppUpdate, error) {
	wait, err := a.CreateUpdate(ctx, asyncUpdateAppRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[AppUpdate]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *AppUpdate) {
		for _, o := range options {
			o(&retries.Info[AppUpdate]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Deletes an app and the dedicated service principal created for it.
func (a *AppsAPI) DeleteByName(ctx context.Context, name string) (*App, error) {
	return a.appsImpl.Delete(ctx, DeleteAppRequest{
		Name: name,
	})
}

// Creates an app deployment for the app with the supplied name. The call
// returns while the deployment is still `IN_PROGRESS`; poll the deployment
// until it reaches a terminal state. The request fails with RESOURCE_CONFLICT
// if another deployment of the same app is still pending and has not timed out.
func (a *AppsAPI) Deploy(ctx context.Context, createAppDeploymentRequest CreateAppDeploymentRequest) (*WaitGetDeploymentAppSucceeded[AppDeployment], error) {
	appDeployment, err := a.appsImpl.Deploy(ctx, createAppDeploymentRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetDeploymentAppSucceeded[AppDeployment]{
		Response:     appDeployment,
		AppName:      createAppDeploymentRequest.AppName,
		DeploymentId: appDeployment.DeploymentId,
		Poll: func(timeout time.Duration, callback func(*AppDeployment)) (*AppDeployment, error) {
			return a.WaitGetDeploymentAppSucceeded(ctx, createAppDeploymentRequest.AppName, appDeployment.DeploymentId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [AppsAPI.Deploy] and waits to reach SUCCEEDED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[AppDeployment](60*time.Minute) functional option.
//
// Deprecated: use [AppsAPI.Deploy].Get() or [AppsAPI.WaitGetDeploymentAppSucceeded]
func (a *AppsAPI) DeployAndWait(ctx context.Context, createAppDeploymentRequest CreateAppDeploymentRequest, options ...retries.Option[AppDeployment]) (*AppDeployment, error) {
	wait, err := a.Deploy(ctx, createAppDeploymentRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[AppDeployment]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *AppDeployment) {
		for _, o := range options {
			o(&retries.Info[AppDeployment]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Retrieves information for the app with the supplied name.
func (a *AppsAPI) GetByName(ctx context.Context, name string) (*App, error) {
	return a.appsImpl.Get(ctx, GetAppRequest{
		Name: name,
	})
}

// Retrieves information for the app deployment with the supplied name and
// deployment id.
func (a *AppsAPI) GetDeploymentByAppNameAndDeploymentId(ctx context.Context, appName string, deploymentId string) (*AppDeployment, error) {
	return a.appsImpl.GetDeployment(ctx, GetAppDeploymentRequest{
		AppName:      appName,
		DeploymentId: deploymentId,
	})
}

// Gets the permission levels that a user can have on an object.
func (a *AppsAPI) GetPermissionLevelsByAppName(ctx context.Context, appName string) (*GetAppPermissionLevelsResponse, error) {
	return a.appsImpl.GetPermissionLevels(ctx, GetAppPermissionLevelsRequest{
		AppName: appName,
	})
}

// Gets the permissions of an app. Apps can inherit permissions from their root
// object.
func (a *AppsAPI) GetPermissionsByAppName(ctx context.Context, appName string) (*AppPermissions, error) {
	return a.appsImpl.GetPermissions(ctx, GetAppPermissionsRequest{
		AppName: appName,
	})
}

// Lists all app deployments for the app with the supplied name. The history
// holds only the app's most recent deployments.
func (a *AppsAPI) ListDeploymentsByAppName(ctx context.Context, appName string) (*ListAppDeploymentsResponse, error) {
	return a.appsImpl.internalListDeployments(ctx, ListAppDeploymentsRequest{
		AppName: appName,
	})
}

// Start the last active deployment of the app in the workspace.
func (a *AppsAPI) Start(ctx context.Context, startAppRequest StartAppRequest) (*WaitGetAppActive[App], error) {
	app, err := a.appsImpl.Start(ctx, startAppRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetAppActive[App]{
		Response: app,
		Name:     app.Name,
		Poll: func(timeout time.Duration, callback func(*App)) (*App, error) {
			return a.WaitGetAppActive(ctx, app.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [AppsAPI.Start] and waits to reach ACTIVE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[App](60*time.Minute) functional option.
//
// Deprecated: use [AppsAPI.Start].Get() or [AppsAPI.WaitGetAppActive]
func (a *AppsAPI) StartAndWait(ctx context.Context, startAppRequest StartAppRequest, options ...retries.Option[App]) (*App, error) {
	wait, err := a.Start(ctx, startAppRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[App]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *App) {
		for _, o := range options {
			o(&retries.Info[App]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Stops the active deployment of the app in the workspace.
func (a *AppsAPI) Stop(ctx context.Context, stopAppRequest StopAppRequest) (*WaitGetAppStopped[App], error) {
	app, err := a.appsImpl.Stop(ctx, stopAppRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetAppStopped[App]{
		Response: app,
		Name:     app.Name,
		Poll: func(timeout time.Duration, callback func(*App)) (*App, error) {
			return a.WaitGetAppStopped(ctx, app.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [AppsAPI.Stop] and waits to reach STOPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[App](60*time.Minute) functional option.
//
// Deprecated: use [AppsAPI.Stop].Get() or [AppsAPI.WaitGetAppStopped]
func (a *AppsAPI) StopAndWait(ctx context.Context, stopAppRequest StopAppRequest, options ...retries.Option[App]) (*App, error) {
	wait, err := a.Stop(ctx, stopAppRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[App]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *App) {
		for _, o := range options {
			o(&retries.Info[App]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

type AppsSettingsInterface interface {

	// Creates a custom template.
	CreateCustomTemplate(ctx context.Context, request CreateCustomTemplateRequest) (*CustomTemplate, error)

	// Deletes the custom template with the specified name.
	DeleteCustomTemplate(ctx context.Context, request DeleteCustomTemplateRequest) (*CustomTemplate, error)

	// Gets the custom template with the specified name.
	GetCustomTemplate(ctx context.Context, request GetCustomTemplateRequest) (*CustomTemplate, error)

	// Lists all custom templates in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListCustomTemplates(ctx context.Context, request ListCustomTemplatesRequest) listing.Iterator[CustomTemplate]

	// Lists all custom templates in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListCustomTemplatesAll(ctx context.Context, request ListCustomTemplatesRequest) ([]CustomTemplate, error)

	// Updates the custom template with the specified name. Note that the template
	// name cannot be updated.
	UpdateCustomTemplate(ctx context.Context, request UpdateCustomTemplateRequest) (*CustomTemplate, error)
}

func NewAppsSettings(client *client.DatabricksClient) *AppsSettingsAPI {
	return &AppsSettingsAPI{
		appsSettingsImpl: appsSettingsImpl{
			client: client,
		},
	}
}

// Apps Settings manage the settings for the Apps service on a customer's
// Databricks instance.
type AppsSettingsAPI struct {
	appsSettingsImpl
}
