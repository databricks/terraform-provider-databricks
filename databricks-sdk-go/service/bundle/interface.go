// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package bundle

import (
	"context"
)

// Service for managing bundle deployment metadata.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type BundleService interface {

	// Marks a version as complete and releases the deployment lock.
	//
	// The server atomically: 1. Sets the version status to the provided
	// terminal status. 2. Sets `complete_time` to the current server timestamp.
	// 3. Releases the lock on the parent deployment. 4. Updates the parent
	// deployment's `status` and `last_version_id`.
	CompleteVersion(ctx context.Context, request CompleteVersionRequest) (*Version, error)

	// Creates a new deployment in the workspace.
	//
	// The caller must provide a `deployment_id` which becomes the final
	// component of the deployment's resource name. If a deployment with the
	// same ID already exists, the server returns `ALREADY_EXISTS`.
	CreateDeployment(ctx context.Context, request CreateDeploymentRequest) (*Deployment, error)

	// Creates a resource operation under a version.
	//
	// The caller must provide a `resource_key` which becomes the final
	// component of the operation's name. If an operation with the same key
	// already exists under the version, the server returns `ALREADY_EXISTS`.
	//
	// On success the server also updates the corresponding deployment-level
	// Resource (creating it if this is the first operation for that
	// resource_key, or removing it if action_type is DELETE).
	CreateOperation(ctx context.Context, request CreateOperationRequest) (*Operation, error)

	// Creates a new version under a deployment.
	//
	// Creating a version acquires an exclusive lock on the deployment,
	// preventing concurrent deploys. The caller provides a `version_id` which
	// the server validates equals `last_version_id + 1` on the deployment.
	CreateVersion(ctx context.Context, request CreateVersionRequest) (*Version, error)

	// Deletes a deployment.
	//
	// The deployment is marked as deleted. It and all its children (versions
	// and their operations) will be permanently deleted after the retention
	// policy expires. If the deployment has an in-progress version, the server
	// returns `RESOURCE_CONFLICT`.
	DeleteDeployment(ctx context.Context, request DeleteDeploymentRequest) error

	// Retrieves a deployment by its resource name.
	GetDeployment(ctx context.Context, request GetDeploymentRequest) (*Deployment, error)

	// Retrieves a resource operation by its resource name.
	GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Retrieves a deployment resource by its resource name.
	GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error)

	// Retrieves a version by its resource name.
	GetVersion(ctx context.Context, request GetVersionRequest) (*Version, error)

	// Sends a heartbeat to renew the lock held by a version.
	//
	// The server validates that the version is the active (non-terminal)
	// version on the parent deployment and resets the lock expiry. If the lock
	// has already expired or the version is no longer active, the server
	// returns `ABORTED`.
	Heartbeat(ctx context.Context, request HeartbeatRequest) (*HeartbeatResponse, error)

	// Lists deployments in the workspace.
	ListDeployments(ctx context.Context, request ListDeploymentsRequest) (*ListDeploymentsResponse, error)

	// Lists resource operations under a version.
	ListOperations(ctx context.Context, request ListOperationsRequest) (*ListOperationsResponse, error)

	// Lists resources under a deployment.
	ListResources(ctx context.Context, request ListResourcesRequest) (*ListResourcesResponse, error)

	// Lists versions under a deployment, ordered by version_id descending (most
	// recent first).
	ListVersions(ctx context.Context, request ListVersionsRequest) (*ListVersionsResponse, error)
}
