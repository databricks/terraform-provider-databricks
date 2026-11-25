// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"context"
)

// The Lakeflow Spark Declarative Pipelines API allows you to create, edit,
// delete, start, and view details about pipelines.
//
// Spark Declarative Pipelines is a framework for building reliable,
// maintainable, and testable data processing pipelines. You define the
// transformations to perform on your data, and Spark Declarative Pipelines
// manages task orchestration, cluster management, monitoring, data quality, and
// error handling.
//
// Instead of defining your data pipelines using a series of separate Apache
// Spark tasks, Spark Declarative Pipelines manages how your data is transformed
// based on a target schema you define for each processing step. You can also
// enforce data quality with Spark Declarative Pipelines expectations.
// Expectations allow you to define expected data quality and specify how to
// handle records that fail those expectations.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PipelinesService interface {

	// Applies the current pipeline environment onto the pipeline compute. The
	// environment applied can be used by subsequent dev-mode updates.
	ApplyEnvironment(ctx context.Context, request ApplyEnvironmentRequest) (*ApplyEnvironmentRequestResponse, error)

	// Creates a new pipeline using Unity Catalog from a pipeline using Hive
	// Metastore. This method returns the ID of the newly created clone.
	// Additionally, this method starts an update for the newly created
	// pipeline.
	Clone(ctx context.Context, request ClonePipelineRequest) (*ClonePipelineResponse, error)

	// Creates a new data processing pipeline based on the requested
	// configuration. If successful, this method returns the ID of the new
	// pipeline.
	Create(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error)

	// Deletes a pipeline. If the pipeline publishes to Unity Catalog, pipeline
	// deletion will cascade to all pipeline tables. Please reach out to
	// Databricks support for assistance to undo this action.
	Delete(ctx context.Context, request DeletePipelineRequest) error

	// Returns the details of a pipeline, including its settings, current state,
	// and the status of its most recent updates.
	Get(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetPipelinePermissionLevelsRequest) (*GetPipelinePermissionLevelsResponse, error)

	// Gets the permissions of a pipeline. Pipelines can inherit permissions
	// from their root object.
	GetPermissions(ctx context.Context, request GetPipelinePermissionsRequest) (*PipelinePermissions, error)

	// Gets an update from an active pipeline.
	GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error)

	// Retrieves events for a pipeline.
	ListPipelineEvents(ctx context.Context, request ListPipelineEventsRequest) (*ListPipelineEventsResponse, error)

	// Lists pipelines defined in the Spark Declarative Pipelines system.
	ListPipelines(ctx context.Context, request ListPipelinesRequest) (*ListPipelinesResponse, error)

	// List updates for an active pipeline.
	ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error)

	// Starts a new update for the pipeline. If there is already an active
	// update for the pipeline, the request will fail and the active update will
	// remain running.
	StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error)

	// Stops the pipeline by canceling the active update. If there is no active
	// update for the pipeline, this request is a no-op.
	Stop(ctx context.Context, request StopRequest) error

	// Updates a pipeline with the supplied configuration.
	Update(ctx context.Context, request EditPipeline) error

	// Updates the permissions on a pipeline. Pipelines can inherit permissions
	// from their root object.
	UpdatePermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error)
}
