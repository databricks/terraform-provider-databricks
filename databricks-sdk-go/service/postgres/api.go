// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Use the Postgres API to create and manage Lakebase Autoscaling Postgres
// infrastructure, including projects, branches, compute endpoints, and roles.
//
// This API manages database infrastructure only. To query or modify data, use
// the Data API or direct SQL connections.
//
// **About resource IDs and names**
//
// Resources are identified by hierarchical resource names like
// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`. The
// `name` field on each resource contains this full path and is output-only.
// Note that `name` refers to this resource path, not the user-visible
// `display_name`.
package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/experimental/api"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type PostgresInterface interface {
	AcknowledgeAssessmentWarnings(ctx context.Context, request AcknowledgeAssessmentWarningsRequest) (AcknowledgeAssessmentWarningsOperationInterface, error)

	AssessMigration(ctx context.Context, request AssessMigrationRequest) (AssessMigrationOperationInterface, error)

	CreateBranch(ctx context.Context, request CreateBranchRequest) (CreateBranchOperationInterface, error)

	CreateCatalog(ctx context.Context, request CreateCatalogRequest) (CreateCatalogOperationInterface, error)

	CreateCdfConfig(ctx context.Context, request CreateCdfConfigRequest) (CreateCdfConfigOperationInterface, error)

	CreateDataApi(ctx context.Context, request CreateDataApiRequest) (CreateDataApiOperationInterface, error)

	CreateDatabase(ctx context.Context, request CreateDatabaseRequest) (CreateDatabaseOperationInterface, error)

	CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (CreateEndpointOperationInterface, error)

	CreateMigration(ctx context.Context, request CreateMigrationRequest) (CreateMigrationOperationInterface, error)

	CreateProject(ctx context.Context, request CreateProjectRequest) (CreateProjectOperationInterface, error)

	CreateRole(ctx context.Context, request CreateRoleRequest) (CreateRoleOperationInterface, error)

	CreateSnapshot(ctx context.Context, request CreateSnapshotRequest) (CreateSnapshotOperationInterface, error)

	CreateSyncedTable(ctx context.Context, request CreateSyncedTableRequest) (CreateSyncedTableOperationInterface, error)

	DeleteBranch(ctx context.Context, request DeleteBranchRequest) (DeleteBranchOperationInterface, error)

	DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) (DeleteCatalogOperationInterface, error)

	DeleteCdfConfig(ctx context.Context, request DeleteCdfConfigRequest) (DeleteCdfConfigOperationInterface, error)

	DeleteDataApi(ctx context.Context, request DeleteDataApiRequest) (DeleteDataApiOperationInterface, error)

	DeleteDatabase(ctx context.Context, request DeleteDatabaseRequest) (DeleteDatabaseOperationInterface, error)

	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) (DeleteEndpointOperationInterface, error)

	DeleteMigration(ctx context.Context, request DeleteMigrationRequest) (DeleteMigrationOperationInterface, error)

	DeleteProject(ctx context.Context, request DeleteProjectRequest) (DeleteProjectOperationInterface, error)

	DeleteRole(ctx context.Context, request DeleteRoleRequest) (DeleteRoleOperationInterface, error)

	DeleteSnapshot(ctx context.Context, request DeleteSnapshotRequest) (DeleteSnapshotOperationInterface, error)

	DeleteSyncedTable(ctx context.Context, request DeleteSyncedTableRequest) (DeleteSyncedTableOperationInterface, error)

	// Generate OAuth credentials for a Postgres database.
	GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error)

	// Gets a migration assessment run.
	GetAssessment(ctx context.Context, request GetAssessmentRequest) (*Assessment, error)

	// Retrieves information about the specified database branch.
	GetBranch(ctx context.Context, request GetBranchRequest) (*Branch, error)

	// Gets the catalog with the given name. The status reports the Postgres
	// database, project and branch that the catalog registers.
	GetCatalog(ctx context.Context, request GetCatalogRequest) (*Catalog, error)

	// Get a single Lakebase CDF configuration, including the source Postgres schema
	// and target Unity Catalog schema.
	GetCdfConfig(ctx context.Context, request GetCdfConfigRequest) (*CdfConfig, error)

	// Get the CDF status of a single table within a Lakebase CDF configuration,
	// including its current state and the last committed position in the feed.
	GetCdfStatus(ctx context.Context, request GetCdfStatusRequest) (*CdfStatus, error)

	// Get Data API configuration for a database.
	GetDataApi(ctx context.Context, request GetDataApiRequest) (*DataApi, error)

	// Retrieves information about the specified Postgres database.
	GetDatabase(ctx context.Context, request GetDatabaseRequest) (*Database, error)

	// Retrieves information about the specified compute endpoint, including its
	// connection details and operational state.
	GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error)

	// Gets a migration.
	GetMigration(ctx context.Context, request GetMigrationRequest) (*Migration, error)

	// Retrieves the status of a long-running operation.
	GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Retrieves information about the specified database project.
	GetProject(ctx context.Context, request GetProjectRequest) (*Project, error)

	// Retrieves information about the specified Postgres role, including its
	// authentication method and permissions.
	GetRole(ctx context.Context, request GetRoleRequest) (*Role, error)

	// Retrieves information about the specified snapshot.
	GetSnapshot(ctx context.Context, request GetSnapshotRequest) (*Snapshot, error)

	// Retrieves the snapshot schedule for a branch. A branch with no configured
	// schedule returns an empty schedule (not NOT_FOUND).
	GetSnapshotSchedule(ctx context.Context, request GetSnapshotScheduleRequest) (*SnapshotSchedule, error)

	// Get a Synced Table.
	GetSyncedTable(ctx context.Context, request GetSyncedTableRequest) (*SyncedTable, error)

	// Lists assessment runs for a migration.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAssessments(ctx context.Context, request ListAssessmentsRequest) listing.Iterator[Assessment]

	// Lists assessment runs for a migration.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAssessmentsAll(ctx context.Context, request ListAssessmentsRequest) ([]Assessment, error)

	// Returns a paginated list of database branches in the project.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListBranches(ctx context.Context, request ListBranchesRequest) listing.Iterator[Branch]

	// Returns a paginated list of database branches in the project.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListBranchesAll(ctx context.Context, request ListBranchesRequest) ([]Branch, error)

	// List all CDF configurations for a Lakebase database. Each configuration maps
	// a Postgres schema to a Unity Catalog schema where the change data feed is
	// materialized.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListCdfConfigs(ctx context.Context, request ListCdfConfigsRequest) listing.Iterator[CdfConfig]

	// List all CDF configurations for a Lakebase database. Each configuration maps
	// a Postgres schema to a Unity Catalog schema where the change data feed is
	// materialized.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListCdfConfigsAll(ctx context.Context, request ListCdfConfigsRequest) ([]CdfConfig, error)

	// List the per-table CDF statuses within a Lakebase CDF configuration. Each
	// status shows whether a table's change data feed is snapshotting, streaming,
	// or skipped.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListCdfStatuses(ctx context.Context, request ListCdfStatusesRequest) listing.Iterator[CdfStatus]

	// List the per-table CDF statuses within a Lakebase CDF configuration. Each
	// status shows whether a table's change data feed is snapshotting, streaming,
	// or skipped.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListCdfStatusesAll(ctx context.Context, request ListCdfStatusesRequest) ([]CdfStatus, error)

	// Returns a list of Postgres databases in the branch.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabases(ctx context.Context, request ListDatabasesRequest) listing.Iterator[Database]

	// Returns a list of Postgres databases in the branch.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabasesAll(ctx context.Context, request ListDatabasesRequest) ([]Database, error)

	// Lists the compute endpoints in the branch, including its read-write endpoint
	// and any read-only endpoints.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) listing.Iterator[Endpoint]

	// Lists the compute endpoints in the branch, including its read-write endpoint
	// and any read-only endpoints.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]Endpoint, error)

	// Lists migrations in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListMigrations(ctx context.Context, request ListMigrationsRequest) listing.Iterator[Migration]

	// Lists migrations in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListMigrationsAll(ctx context.Context, request ListMigrationsRequest) ([]Migration, error)

	// Returns a paginated list of database projects in the workspace that the user
	// has permission to access.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListProjects(ctx context.Context, request ListProjectsRequest) listing.Iterator[Project]

	// Returns a paginated list of database projects in the workspace that the user
	// has permission to access.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListProjectsAll(ctx context.Context, request ListProjectsRequest) ([]Project, error)

	// Returns a paginated list of Postgres roles in the branch.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListRoles(ctx context.Context, request ListRolesRequest) listing.Iterator[Role]

	// Returns a paginated list of Postgres roles in the branch.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListRolesAll(ctx context.Context, request ListRolesRequest) ([]Role, error)

	// Returns a paginated list of snapshots in the project.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSnapshots(ctx context.Context, request ListSnapshotsRequest) listing.Iterator[Snapshot]

	// Returns a paginated list of snapshots in the project.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSnapshotsAll(ctx context.Context, request ListSnapshotsRequest) ([]Snapshot, error)

	// Reveals the current password for a password-authenticated Postgres role.
	RevealRolePassword(ctx context.Context, request RevealRolePasswordRequest) (*RevealRolePasswordResponse, error)

	StartMigration(ctx context.Context, request StartMigrationRequest) (StartMigrationOperationInterface, error)

	UndeleteBranch(ctx context.Context, request UndeleteBranchRequest) (UndeleteBranchOperationInterface, error)

	UndeleteMigration(ctx context.Context, request UndeleteMigrationRequest) (UndeleteMigrationOperationInterface, error)

	UndeleteProject(ctx context.Context, request UndeleteProjectRequest) (UndeleteProjectOperationInterface, error)

	UpdateBranch(ctx context.Context, request UpdateBranchRequest) (UpdateBranchOperationInterface, error)

	UpdateDataApi(ctx context.Context, request UpdateDataApiRequest) (UpdateDataApiOperationInterface, error)

	UpdateDatabase(ctx context.Context, request UpdateDatabaseRequest) (UpdateDatabaseOperationInterface, error)

	UpdateEndpoint(ctx context.Context, request UpdateEndpointRequest) (UpdateEndpointOperationInterface, error)

	UpdateMigration(ctx context.Context, request UpdateMigrationRequest) (UpdateMigrationOperationInterface, error)

	UpdateProject(ctx context.Context, request UpdateProjectRequest) (UpdateProjectOperationInterface, error)

	UpdateRole(ctx context.Context, request UpdateRoleRequest) (UpdateRoleOperationInterface, error)

	UpdateSnapshotSchedule(ctx context.Context, request UpdateSnapshotScheduleRequest) (UpdateSnapshotScheduleOperationInterface, error)
}

func NewPostgres(client *client.DatabricksClient) *PostgresAPI {
	return &PostgresAPI{
		postgresImpl: postgresImpl{
			client: client,
		},
	}
}

// Use the Postgres API to create and manage Lakebase Autoscaling Postgres
// infrastructure, including projects, branches, compute endpoints, and roles.
//
// This API manages database infrastructure only. To query or modify data, use
// the Data API or direct SQL connections.
//
// **About resource IDs and names**
//
// Resources are identified by hierarchical resource names like
// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`. The
// `name` field on each resource contains this full path and is output-only.
// Note that `name` refers to this resource path, not the user-visible
// `display_name`.
type PostgresAPI struct {
	postgresImpl
}

func (a *PostgresAPI) AcknowledgeAssessmentWarnings(ctx context.Context, request AcknowledgeAssessmentWarningsRequest) (AcknowledgeAssessmentWarningsOperationInterface, error) {
	operation, err := a.postgresImpl.AcknowledgeAssessmentWarnings(ctx, request)
	if err != nil {
		return nil, err
	}
	return &acknowledgeAssessmentWarningsOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type AcknowledgeAssessmentWarningsOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*AcknowledgeAssessmentWarningsResponse, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*MigrationOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type acknowledgeAssessmentWarningsOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *acknowledgeAssessmentWarningsOperation) Wait(ctx context.Context, opts ...api.Option) (*AcknowledgeAssessmentWarningsResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *AcknowledgeAssessmentWarningsResponse
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var acknowledgeAssessmentWarningsResponse AcknowledgeAssessmentWarningsResponse
		err = json.Unmarshal(operation.Response, &acknowledgeAssessmentWarningsResponse)
		if err != nil {
			return fmt.Errorf("failed to unmarshal acknowledgeAssessmentWarningsResponse response: %w", err)
		}

		result = &acknowledgeAssessmentWarningsResponse

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
func (a *acknowledgeAssessmentWarningsOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *acknowledgeAssessmentWarningsOperation) Metadata() (*MigrationOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata MigrationOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *acknowledgeAssessmentWarningsOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) AssessMigration(ctx context.Context, request AssessMigrationRequest) (AssessMigrationOperationInterface, error) {
	operation, err := a.postgresImpl.AssessMigration(ctx, request)
	if err != nil {
		return nil, err
	}
	return &assessMigrationOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type AssessMigrationOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*AssessMigrationResponse, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*MigrationOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type assessMigrationOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *assessMigrationOperation) Wait(ctx context.Context, opts ...api.Option) (*AssessMigrationResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *AssessMigrationResponse
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var assessMigrationResponse AssessMigrationResponse
		err = json.Unmarshal(operation.Response, &assessMigrationResponse)
		if err != nil {
			return fmt.Errorf("failed to unmarshal assessMigrationResponse response: %w", err)
		}

		result = &assessMigrationResponse

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
func (a *assessMigrationOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *assessMigrationOperation) Metadata() (*MigrationOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata MigrationOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *assessMigrationOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateBranch(ctx context.Context, request CreateBranchRequest) (CreateBranchOperationInterface, error) {
	operation, err := a.postgresImpl.CreateBranch(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createBranchOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateBranchOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Branch, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*BranchOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createBranchOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createBranchOperation) Wait(ctx context.Context, opts ...api.Option) (*Branch, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Branch
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var branch Branch
		err = json.Unmarshal(operation.Response, &branch)
		if err != nil {
			return fmt.Errorf("failed to unmarshal branch response: %w", err)
		}

		result = &branch

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
func (a *createBranchOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createBranchOperation) Metadata() (*BranchOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata BranchOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createBranchOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateCatalog(ctx context.Context, request CreateCatalogRequest) (CreateCatalogOperationInterface, error) {
	operation, err := a.postgresImpl.CreateCatalog(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createCatalogOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateCatalogOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Catalog, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*CatalogOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createCatalogOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createCatalogOperation) Wait(ctx context.Context, opts ...api.Option) (*Catalog, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Catalog
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var catalog Catalog
		err = json.Unmarshal(operation.Response, &catalog)
		if err != nil {
			return fmt.Errorf("failed to unmarshal catalog response: %w", err)
		}

		result = &catalog

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
func (a *createCatalogOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createCatalogOperation) Metadata() (*CatalogOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata CatalogOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createCatalogOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateCdfConfig(ctx context.Context, request CreateCdfConfigRequest) (CreateCdfConfigOperationInterface, error) {
	operation, err := a.postgresImpl.CreateCdfConfig(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createCdfConfigOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateCdfConfigOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*CdfConfig, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*CdfConfigOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createCdfConfigOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createCdfConfigOperation) Wait(ctx context.Context, opts ...api.Option) (*CdfConfig, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *CdfConfig
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var cdfConfig CdfConfig
		err = json.Unmarshal(operation.Response, &cdfConfig)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cdfConfig response: %w", err)
		}

		result = &cdfConfig

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
func (a *createCdfConfigOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createCdfConfigOperation) Metadata() (*CdfConfigOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata CdfConfigOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createCdfConfigOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateDataApi(ctx context.Context, request CreateDataApiRequest) (CreateDataApiOperationInterface, error) {
	operation, err := a.postgresImpl.CreateDataApi(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createDataApiOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateDataApiOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*DataApi, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*DataApiOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createDataApiOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createDataApiOperation) Wait(ctx context.Context, opts ...api.Option) (*DataApi, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *DataApi
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var dataApi DataApi
		err = json.Unmarshal(operation.Response, &dataApi)
		if err != nil {
			return fmt.Errorf("failed to unmarshal dataApi response: %w", err)
		}

		result = &dataApi

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
func (a *createDataApiOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createDataApiOperation) Metadata() (*DataApiOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata DataApiOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createDataApiOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateDatabase(ctx context.Context, request CreateDatabaseRequest) (CreateDatabaseOperationInterface, error) {
	operation, err := a.postgresImpl.CreateDatabase(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createDatabaseOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateDatabaseOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Database, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*DatabaseOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createDatabaseOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createDatabaseOperation) Wait(ctx context.Context, opts ...api.Option) (*Database, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Database
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var database Database
		err = json.Unmarshal(operation.Response, &database)
		if err != nil {
			return fmt.Errorf("failed to unmarshal database response: %w", err)
		}

		result = &database

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
func (a *createDatabaseOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createDatabaseOperation) Metadata() (*DatabaseOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata DatabaseOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createDatabaseOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (CreateEndpointOperationInterface, error) {
	operation, err := a.postgresImpl.CreateEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createEndpointOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateEndpointOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Endpoint, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*EndpointOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createEndpointOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createEndpointOperation) Wait(ctx context.Context, opts ...api.Option) (*Endpoint, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Endpoint
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var endpoint Endpoint
		err = json.Unmarshal(operation.Response, &endpoint)
		if err != nil {
			return fmt.Errorf("failed to unmarshal endpoint response: %w", err)
		}

		result = &endpoint

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
func (a *createEndpointOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createEndpointOperation) Metadata() (*EndpointOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata EndpointOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createEndpointOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateMigration(ctx context.Context, request CreateMigrationRequest) (CreateMigrationOperationInterface, error) {
	operation, err := a.postgresImpl.CreateMigration(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createMigrationOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateMigrationOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Migration, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*MigrationOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createMigrationOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createMigrationOperation) Wait(ctx context.Context, opts ...api.Option) (*Migration, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Migration
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var migration Migration
		err = json.Unmarshal(operation.Response, &migration)
		if err != nil {
			return fmt.Errorf("failed to unmarshal migration response: %w", err)
		}

		result = &migration

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
func (a *createMigrationOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createMigrationOperation) Metadata() (*MigrationOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata MigrationOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createMigrationOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateProject(ctx context.Context, request CreateProjectRequest) (CreateProjectOperationInterface, error) {
	operation, err := a.postgresImpl.CreateProject(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createProjectOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateProjectOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Project, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*ProjectOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createProjectOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createProjectOperation) Wait(ctx context.Context, opts ...api.Option) (*Project, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Project
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var project Project
		err = json.Unmarshal(operation.Response, &project)
		if err != nil {
			return fmt.Errorf("failed to unmarshal project response: %w", err)
		}

		result = &project

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
func (a *createProjectOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createProjectOperation) Metadata() (*ProjectOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata ProjectOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createProjectOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateRole(ctx context.Context, request CreateRoleRequest) (CreateRoleOperationInterface, error) {
	operation, err := a.postgresImpl.CreateRole(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createRoleOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateRoleOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Role, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*RoleOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createRoleOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createRoleOperation) Wait(ctx context.Context, opts ...api.Option) (*Role, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Role
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var role Role
		err = json.Unmarshal(operation.Response, &role)
		if err != nil {
			return fmt.Errorf("failed to unmarshal role response: %w", err)
		}

		result = &role

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
func (a *createRoleOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createRoleOperation) Metadata() (*RoleOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata RoleOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createRoleOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateSnapshot(ctx context.Context, request CreateSnapshotRequest) (CreateSnapshotOperationInterface, error) {
	operation, err := a.postgresImpl.CreateSnapshot(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createSnapshotOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateSnapshotOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Snapshot, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*SnapshotOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createSnapshotOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createSnapshotOperation) Wait(ctx context.Context, opts ...api.Option) (*Snapshot, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Snapshot
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var snapshot Snapshot
		err = json.Unmarshal(operation.Response, &snapshot)
		if err != nil {
			return fmt.Errorf("failed to unmarshal snapshot response: %w", err)
		}

		result = &snapshot

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
func (a *createSnapshotOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createSnapshotOperation) Metadata() (*SnapshotOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata SnapshotOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createSnapshotOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) CreateSyncedTable(ctx context.Context, request CreateSyncedTableRequest) (CreateSyncedTableOperationInterface, error) {
	operation, err := a.postgresImpl.CreateSyncedTable(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createSyncedTableOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type CreateSyncedTableOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*SyncedTable, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*SyncedTableOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createSyncedTableOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createSyncedTableOperation) Wait(ctx context.Context, opts ...api.Option) (*SyncedTable, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *SyncedTable
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var syncedTable SyncedTable
		err = json.Unmarshal(operation.Response, &syncedTable)
		if err != nil {
			return fmt.Errorf("failed to unmarshal syncedTable response: %w", err)
		}

		result = &syncedTable

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
func (a *createSyncedTableOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createSyncedTableOperation) Metadata() (*SyncedTableOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata SyncedTableOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createSyncedTableOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteBranch(ctx context.Context, request DeleteBranchRequest) (DeleteBranchOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteBranch(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteBranchOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteBranchOperationInterface interface {

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
	Metadata() (*BranchOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteBranchOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteBranchOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteBranchOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteBranchOperation) Metadata() (*BranchOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata BranchOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteBranchOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) (DeleteCatalogOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteCatalog(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteCatalogOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteCatalogOperationInterface interface {

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
	Metadata() (*CatalogOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteCatalogOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteCatalogOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteCatalogOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteCatalogOperation) Metadata() (*CatalogOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata CatalogOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteCatalogOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteCdfConfig(ctx context.Context, request DeleteCdfConfigRequest) (DeleteCdfConfigOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteCdfConfig(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteCdfConfigOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteCdfConfigOperationInterface interface {

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
	Metadata() (*CdfConfigOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteCdfConfigOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteCdfConfigOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteCdfConfigOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteCdfConfigOperation) Metadata() (*CdfConfigOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata CdfConfigOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteCdfConfigOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteDataApi(ctx context.Context, request DeleteDataApiRequest) (DeleteDataApiOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteDataApi(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteDataApiOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteDataApiOperationInterface interface {

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
	Metadata() (*DataApiOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteDataApiOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteDataApiOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteDataApiOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteDataApiOperation) Metadata() (*DataApiOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata DataApiOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteDataApiOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteDatabase(ctx context.Context, request DeleteDatabaseRequest) (DeleteDatabaseOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteDatabase(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteDatabaseOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteDatabaseOperationInterface interface {

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
	Metadata() (*DatabaseOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteDatabaseOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteDatabaseOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteDatabaseOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteDatabaseOperation) Metadata() (*DatabaseOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata DatabaseOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteDatabaseOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) (DeleteEndpointOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteEndpointOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteEndpointOperationInterface interface {

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
	Metadata() (*EndpointOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteEndpointOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteEndpointOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteEndpointOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteEndpointOperation) Metadata() (*EndpointOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata EndpointOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteEndpointOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteMigration(ctx context.Context, request DeleteMigrationRequest) (DeleteMigrationOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteMigration(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteMigrationOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteMigrationOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Migration, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*MigrationOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteMigrationOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteMigrationOperation) Wait(ctx context.Context, opts ...api.Option) (*Migration, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Migration
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var migration Migration
		err = json.Unmarshal(operation.Response, &migration)
		if err != nil {
			return fmt.Errorf("failed to unmarshal migration response: %w", err)
		}

		result = &migration

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
func (a *deleteMigrationOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteMigrationOperation) Metadata() (*MigrationOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata MigrationOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteMigrationOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteProject(ctx context.Context, request DeleteProjectRequest) (DeleteProjectOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteProject(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteProjectOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteProjectOperationInterface interface {

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
	Metadata() (*ProjectOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteProjectOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteProjectOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteProjectOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteProjectOperation) Metadata() (*ProjectOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata ProjectOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteProjectOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteRole(ctx context.Context, request DeleteRoleRequest) (DeleteRoleOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteRole(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteRoleOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteRoleOperationInterface interface {

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
	Metadata() (*RoleOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteRoleOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteRoleOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteRoleOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteRoleOperation) Metadata() (*RoleOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata RoleOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteRoleOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteSnapshot(ctx context.Context, request DeleteSnapshotRequest) (DeleteSnapshotOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteSnapshot(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteSnapshotOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteSnapshotOperationInterface interface {

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
	Metadata() (*SnapshotOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteSnapshotOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteSnapshotOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteSnapshotOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteSnapshotOperation) Metadata() (*SnapshotOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata SnapshotOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteSnapshotOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) DeleteSyncedTable(ctx context.Context, request DeleteSyncedTableRequest) (DeleteSyncedTableOperationInterface, error) {
	operation, err := a.postgresImpl.DeleteSyncedTable(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteSyncedTableOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type DeleteSyncedTableOperationInterface interface {

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
	Metadata() (*SyncedTableOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteSyncedTableOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteSyncedTableOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *deleteSyncedTableOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteSyncedTableOperation) Metadata() (*SyncedTableOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata SyncedTableOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteSyncedTableOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) StartMigration(ctx context.Context, request StartMigrationRequest) (StartMigrationOperationInterface, error) {
	operation, err := a.postgresImpl.StartMigration(ctx, request)
	if err != nil {
		return nil, err
	}
	return &startMigrationOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type StartMigrationOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*StartMigrationResponse, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*MigrationOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type startMigrationOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *startMigrationOperation) Wait(ctx context.Context, opts ...api.Option) (*StartMigrationResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *StartMigrationResponse
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var startMigrationResponse StartMigrationResponse
		err = json.Unmarshal(operation.Response, &startMigrationResponse)
		if err != nil {
			return fmt.Errorf("failed to unmarshal startMigrationResponse response: %w", err)
		}

		result = &startMigrationResponse

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
func (a *startMigrationOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *startMigrationOperation) Metadata() (*MigrationOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata MigrationOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *startMigrationOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UndeleteBranch(ctx context.Context, request UndeleteBranchRequest) (UndeleteBranchOperationInterface, error) {
	operation, err := a.postgresImpl.UndeleteBranch(ctx, request)
	if err != nil {
		return nil, err
	}
	return &undeleteBranchOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UndeleteBranchOperationInterface interface {

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
	Metadata() (*BranchOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type undeleteBranchOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *undeleteBranchOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *undeleteBranchOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *undeleteBranchOperation) Metadata() (*BranchOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata BranchOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *undeleteBranchOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UndeleteMigration(ctx context.Context, request UndeleteMigrationRequest) (UndeleteMigrationOperationInterface, error) {
	operation, err := a.postgresImpl.UndeleteMigration(ctx, request)
	if err != nil {
		return nil, err
	}
	return &undeleteMigrationOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UndeleteMigrationOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*UndeleteMigrationResponse, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*MigrationOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type undeleteMigrationOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *undeleteMigrationOperation) Wait(ctx context.Context, opts ...api.Option) (*UndeleteMigrationResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *UndeleteMigrationResponse
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var undeleteMigrationResponse UndeleteMigrationResponse
		err = json.Unmarshal(operation.Response, &undeleteMigrationResponse)
		if err != nil {
			return fmt.Errorf("failed to unmarshal undeleteMigrationResponse response: %w", err)
		}

		result = &undeleteMigrationResponse

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
func (a *undeleteMigrationOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *undeleteMigrationOperation) Metadata() (*MigrationOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata MigrationOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *undeleteMigrationOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UndeleteProject(ctx context.Context, request UndeleteProjectRequest) (UndeleteProjectOperationInterface, error) {
	operation, err := a.postgresImpl.UndeleteProject(ctx, request)
	if err != nil {
		return nil, err
	}
	return &undeleteProjectOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UndeleteProjectOperationInterface interface {

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
	Metadata() (*ProjectOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type undeleteProjectOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *undeleteProjectOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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
func (a *undeleteProjectOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *undeleteProjectOperation) Metadata() (*ProjectOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata ProjectOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *undeleteProjectOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UpdateBranch(ctx context.Context, request UpdateBranchRequest) (UpdateBranchOperationInterface, error) {
	operation, err := a.postgresImpl.UpdateBranch(ctx, request)
	if err != nil {
		return nil, err
	}
	return &updateBranchOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UpdateBranchOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Branch, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*BranchOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type updateBranchOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *updateBranchOperation) Wait(ctx context.Context, opts ...api.Option) (*Branch, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Branch
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var branch Branch
		err = json.Unmarshal(operation.Response, &branch)
		if err != nil {
			return fmt.Errorf("failed to unmarshal branch response: %w", err)
		}

		result = &branch

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
func (a *updateBranchOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *updateBranchOperation) Metadata() (*BranchOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata BranchOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *updateBranchOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UpdateDataApi(ctx context.Context, request UpdateDataApiRequest) (UpdateDataApiOperationInterface, error) {
	operation, err := a.postgresImpl.UpdateDataApi(ctx, request)
	if err != nil {
		return nil, err
	}
	return &updateDataApiOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UpdateDataApiOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*DataApi, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*DataApiOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type updateDataApiOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *updateDataApiOperation) Wait(ctx context.Context, opts ...api.Option) (*DataApi, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *DataApi
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var dataApi DataApi
		err = json.Unmarshal(operation.Response, &dataApi)
		if err != nil {
			return fmt.Errorf("failed to unmarshal dataApi response: %w", err)
		}

		result = &dataApi

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
func (a *updateDataApiOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *updateDataApiOperation) Metadata() (*DataApiOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata DataApiOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *updateDataApiOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UpdateDatabase(ctx context.Context, request UpdateDatabaseRequest) (UpdateDatabaseOperationInterface, error) {
	operation, err := a.postgresImpl.UpdateDatabase(ctx, request)
	if err != nil {
		return nil, err
	}
	return &updateDatabaseOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UpdateDatabaseOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Database, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*DatabaseOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type updateDatabaseOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *updateDatabaseOperation) Wait(ctx context.Context, opts ...api.Option) (*Database, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Database
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var database Database
		err = json.Unmarshal(operation.Response, &database)
		if err != nil {
			return fmt.Errorf("failed to unmarshal database response: %w", err)
		}

		result = &database

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
func (a *updateDatabaseOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *updateDatabaseOperation) Metadata() (*DatabaseOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata DatabaseOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *updateDatabaseOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UpdateEndpoint(ctx context.Context, request UpdateEndpointRequest) (UpdateEndpointOperationInterface, error) {
	operation, err := a.postgresImpl.UpdateEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return &updateEndpointOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UpdateEndpointOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Endpoint, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*EndpointOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type updateEndpointOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *updateEndpointOperation) Wait(ctx context.Context, opts ...api.Option) (*Endpoint, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Endpoint
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var endpoint Endpoint
		err = json.Unmarshal(operation.Response, &endpoint)
		if err != nil {
			return fmt.Errorf("failed to unmarshal endpoint response: %w", err)
		}

		result = &endpoint

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
func (a *updateEndpointOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *updateEndpointOperation) Metadata() (*EndpointOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata EndpointOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *updateEndpointOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UpdateMigration(ctx context.Context, request UpdateMigrationRequest) (UpdateMigrationOperationInterface, error) {
	operation, err := a.postgresImpl.UpdateMigration(ctx, request)
	if err != nil {
		return nil, err
	}
	return &updateMigrationOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UpdateMigrationOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Migration, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*MigrationOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type updateMigrationOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *updateMigrationOperation) Wait(ctx context.Context, opts ...api.Option) (*Migration, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Migration
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var migration Migration
		err = json.Unmarshal(operation.Response, &migration)
		if err != nil {
			return fmt.Errorf("failed to unmarshal migration response: %w", err)
		}

		result = &migration

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
func (a *updateMigrationOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *updateMigrationOperation) Metadata() (*MigrationOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata MigrationOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *updateMigrationOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UpdateProject(ctx context.Context, request UpdateProjectRequest) (UpdateProjectOperationInterface, error) {
	operation, err := a.postgresImpl.UpdateProject(ctx, request)
	if err != nil {
		return nil, err
	}
	return &updateProjectOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UpdateProjectOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Project, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*ProjectOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type updateProjectOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *updateProjectOperation) Wait(ctx context.Context, opts ...api.Option) (*Project, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Project
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var project Project
		err = json.Unmarshal(operation.Response, &project)
		if err != nil {
			return fmt.Errorf("failed to unmarshal project response: %w", err)
		}

		result = &project

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
func (a *updateProjectOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *updateProjectOperation) Metadata() (*ProjectOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata ProjectOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *updateProjectOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UpdateRole(ctx context.Context, request UpdateRoleRequest) (UpdateRoleOperationInterface, error) {
	operation, err := a.postgresImpl.UpdateRole(ctx, request)
	if err != nil {
		return nil, err
	}
	return &updateRoleOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UpdateRoleOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*Role, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*RoleOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type updateRoleOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *updateRoleOperation) Wait(ctx context.Context, opts ...api.Option) (*Role, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *Role
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var role Role
		err = json.Unmarshal(operation.Response, &role)
		if err != nil {
			return fmt.Errorf("failed to unmarshal role response: %w", err)
		}

		result = &role

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
func (a *updateRoleOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *updateRoleOperation) Metadata() (*RoleOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata RoleOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *updateRoleOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *PostgresAPI) UpdateSnapshotSchedule(ctx context.Context, request UpdateSnapshotScheduleRequest) (UpdateSnapshotScheduleOperationInterface, error) {
	operation, err := a.postgresImpl.UpdateSnapshotSchedule(ctx, request)
	if err != nil {
		return nil, err
	}
	return &updateSnapshotScheduleOperation{
		impl:      &a.postgresImpl,
		operation: operation,
	}, nil
}

type UpdateSnapshotScheduleOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*SnapshotSchedule, error)

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*SnapshotScheduleOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type updateSnapshotScheduleOperation struct {
	impl      *postgresImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *updateSnapshotScheduleOperation) Wait(ctx context.Context, opts ...api.Option) (*SnapshotSchedule, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *SnapshotSchedule
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
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

		var snapshotSchedule SnapshotSchedule
		err = json.Unmarshal(operation.Response, &snapshotSchedule)
		if err != nil {
			return fmt.Errorf("failed to unmarshal snapshotSchedule response: %w", err)
		}

		result = &snapshotSchedule

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
func (a *updateSnapshotScheduleOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *updateSnapshotScheduleOperation) Metadata() (*SnapshotScheduleOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata SnapshotScheduleOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *updateSnapshotScheduleOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}
