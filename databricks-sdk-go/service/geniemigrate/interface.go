// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package geniemigrate

import (
	"context"
)

// Use the Genie Migrate API to create and manage migration projects and the
// source files to convert.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type GenieMigrateService interface {

	// Adds source files to a migration project. Files that are already in the
	// project are skipped and keep their conversion results.
	AddFiles(ctx context.Context, request AddProjectFilesRequest) (*AddProjectFilesResponse, error)

	// Requests cancellation of a long-running operation. Requires `CAN_EDIT`
	// permission on the project. Cancellation is best effort: the operation can
	// complete before cancellation takes effect. Call `GetOperation` to see its
	// final state.
	CancelOperation(ctx context.Context, request CancelOperationRequest) error

	// Converts files of a migration project in a background job. Returns an
	// operation; poll it with `GetOperation`. The operation's metadata is the
	// current `ConversionRun`, with the run's state and the status of each
	// file. When the run completes, the operation's response is a
	// `ConvertProjectFilesResponse`. A failed or cancelled run sets the
	// operation's `error` instead. Files that fail to convert don't fail the
	// run.
	//
	// To stop a conversion, call `CancelOperation`. Cancellation is best
	// effort, and the conversion can complete first. `GetOperation` reports the
	// final outcome.
	ConvertProjectFiles(ctx context.Context, request ConvertProjectFilesRequest) (*Operation, error)

	// Creates a migration project in a workspace folder. The server assigns the
	// project ID. Add source files to the project with `AddProjectFiles`.
	Create(ctx context.Context, request CreateProjectRequest) (*Project, error)

	// Moves a migration project to the workspace trash. The project keeps its
	// ID until it's removed from the trash.
	Delete(ctx context.Context, request DeleteProjectRequest) error

	// Gets a migration project, including its conversion settings and whether
	// they're locked. To list the project's source files, call
	// `ListProjectFiles`. To track a conversion, call `GetOperation`.
	Get(ctx context.Context, request GetProjectRequest) (*Project, error)

	// Gets a file of a migration project, including the conversion details
	// stored with its converted file, when available. Requires `CAN_READ`
	// permission on the project and, once the file is converted, on the
	// converted file.
	GetFile(ctx context.Context, request GetProjectFileRequest) (*ProjectFile, error)

	// Gets the latest state of a long-running operation, such as a conversion
	// started by `ConvertProjectFiles`. Operation names have the format
	// `projects/{project_id}/operations/{operation_id}`. Use the full name that
	// the service returned, and treat the operation ID as opaque. Requires
	// `CAN_READ` permission on the project.
	GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Lists the migration projects in a workspace folder. Select the folder
	// with `filter`; without a filter, lists the caller's home folder. Listing
	// every project in the workspace isn't supported.
	//
	// Returns project summaries without conversion settings; call `GetProject`
	// for the full project. A page can be empty even when more pages follow, so
	// keep requesting pages until `next_page_token` is empty.
	List(ctx context.Context, request ListProjectsRequest) (*ListProjectsResponse, error)

	// Lists the files of a migration project with their latest conversion
	// statuses. To read a file's changes, to-do items, blockers, and notes,
	// call `GetProjectFile`. A page can have fewer files than requested, or
	// none, even when more pages follow; keep requesting pages until
	// `next_page_token` is empty.
	ListFiles(ctx context.Context, request ListProjectFilesRequest) (*ListProjectFilesResponse, error)

	// Removes files from a migration project by their resource names. The
	// source files and their converted output stay in the workspace.
	RemoveFiles(ctx context.Context, request RemoveProjectFilesRequest) (*RemoveProjectFilesResponse, error)

	// Marks files of a migration project as reviewed, accepting their converted
	// output. Each file's status becomes `FILE_STATUS_REVIEWED`, and reviewed
	// files aren't converted again. Files that are already reviewed don't
	// change. Returns the files in request order.
	ReviewProjectFiles(ctx context.Context, request ReviewProjectFilesRequest) (*ReviewProjectFilesResponse, error)

	// Removes the reviewed status from files of a migration project. Each file
	// returns to the status it had before it was reviewed. Files that aren't
	// reviewed don't change. Returns the files in request order.
	UnreviewProjectFiles(ctx context.Context, request UnreviewProjectFilesRequest) (*UnreviewProjectFilesResponse, error)

	// Updates a migration project. Only the fields listed in `update_mask`
	// change. An empty mask, or a mask with an unsupported field, returns
	// `INVALID_PARAMETER_VALUE`. When `changes_locked` is `true`, the source
	// dialect and target framework can't be changed.
	Update(ctx context.Context, request UpdateProjectRequest) (*Project, error)
}
