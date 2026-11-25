// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres

import (
	"context"
)

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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PostgresService interface {

	// Acknowledges non-blocking warnings on a specific assessment.
	AcknowledgeAssessmentWarnings(ctx context.Context, request AcknowledgeAssessmentWarningsRequest) (*Operation, error)

	// Creates an immutable assessment run for the migration's current source
	// configuration.
	AssessMigration(ctx context.Context, request AssessMigrationRequest) (*Operation, error)

	// Creates a new database branch in the project.
	CreateBranch(ctx context.Context, request CreateBranchRequest) (*Operation, error)

	// Registers a Postgres database from a Lakebase project branch as a catalog
	// in Unity Catalog, so that its schemas, tables and views are governed and
	// discoverable through Unity Catalog. The catalog is read-only; the data
	// stays managed through Postgres. Returns a long-running operation. Poll it
	// until done is true; its response holds the new catalog. For an overview,
	// see https://docs.databricks.com/aws/en/oltp/projects/register-uc.
	CreateCatalog(ctx context.Context, request CreateCatalogRequest) (*Operation, error)

	// Create a CDF configuration that materializes the change data feed for all
	// tables in a Postgres schema as open-format Delta tables in Unity Catalog.
	// Once created, each table's change history is continuously written to its
	// corresponding Lakehouse table.
	CreateCdfConfig(ctx context.Context, request CreateCdfConfigRequest) (*Operation, error)

	// Enable Data API for a database.
	CreateDataApi(ctx context.Context, request CreateDataApiRequest) (*Operation, error)

	// Create a Database.
	//
	// Creates a database in the specified branch. A branch can have multiple
	// databases.
	CreateDatabase(ctx context.Context, request CreateDatabaseRequest) (*Operation, error)

	// Creates a compute endpoint in the branch. Each branch gets its read-write
	// endpoint when the branch is created, so use this method to add read-only
	// endpoints (read replicas) or, with `replace_existing`, to reconfigure the
	// existing read-write endpoint.
	CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Operation, error)

	// Creates a migration.
	CreateMigration(ctx context.Context, request CreateMigrationRequest) (*Operation, error)

	// Creates a new Lakebase Autoscaling Postgres database project, which
	// contains branches and compute endpoints. The new project starts with a
	// single branch, a read-write compute endpoint on that branch, a Postgres
	// database named `databricks_postgres`, and a Postgres role for the calling
	// identity.
	CreateProject(ctx context.Context, request CreateProjectRequest) (*Operation, error)

	// Creates a new Postgres role in the branch.
	CreateRole(ctx context.Context, request CreateRoleRequest) (*Operation, error)

	// Creates a snapshot, an immutable point-in-time copy of a branch's data,
	// within the project. Returns a long-running operation. Poll it until done
	// is true; its response holds the new snapshot.
	CreateSnapshot(ctx context.Context, request CreateSnapshotRequest) (*Operation, error)

	// Create a Synced Table. Creates the synced table in Unity Catalog and the
	// Postgres table on the branch, and starts the pipeline that loads the
	// initial data from the source table.
	CreateSyncedTable(ctx context.Context, request CreateSyncedTableRequest) (*Operation, error)

	// Deletes the specified database branch.
	DeleteBranch(ctx context.Context, request DeleteBranchRequest) (*Operation, error)

	// Deletes the catalog from Unity Catalog. The Postgres database and its
	// tables remain in the branch, and the database can be registered again
	// afterwards. Returns a long-running operation. Poll it until done is true.
	DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) (*Operation, error)

	// Delete a CDF configuration and stop materializing the change data feed.
	// When force=true, also drops the Delta tables in Unity Catalog. When
	// force=false (default), the existing tables are preserved at their last
	// state.
	DeleteCdfConfig(ctx context.Context, request DeleteCdfConfigRequest) (*Operation, error)

	// Disable Data API for a database.
	DeleteDataApi(ctx context.Context, request DeleteDataApiRequest) (*Operation, error)

	// Deletes the specified Postgres database.
	DeleteDatabase(ctx context.Context, request DeleteDatabaseRequest) (*Operation, error)

	// Deletes the specified compute endpoint. Deletion is permanent. Only
	// read-only endpoints can be deleted this way, because a branch's
	// read-write endpoint is deleted together with the branch. To stop the
	// compute costs of a read-write endpoint without deleting the branch, set
	// `spec.disabled` to true.
	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) (*Operation, error)

	// Soft-deletes a migration.
	DeleteMigration(ctx context.Context, request DeleteMigrationRequest) (*Operation, error)

	// Deletes the specified database project. By default the project is
	// soft-deleted and can be restored with Undelete during its retention
	// period; set `purge` to delete it permanently.
	DeleteProject(ctx context.Context, request DeleteProjectRequest) (*Operation, error)

	// Deletes the specified Postgres role.
	DeleteRole(ctx context.Context, request DeleteRoleRequest) (*Operation, error)

	// Deletes the specified snapshot. A deleted snapshot cannot be recovered.
	// Returns a long-running operation. Poll it until done is true.
	DeleteSnapshot(ctx context.Context, request DeleteSnapshotRequest) (*Operation, error)

	// Delete a Synced Table. Drops the Postgres table, removes the synced table
	// from its pipeline, and deletes it from Unity Catalog. The pipeline itself
	// is deleted when no other synced table uses it.
	DeleteSyncedTable(ctx context.Context, request DeleteSyncedTableRequest) (*Operation, error)

	// Generate OAuth credentials for a Postgres database.
	GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error)

	// Gets a migration assessment run.
	GetAssessment(ctx context.Context, request GetAssessmentRequest) (*Assessment, error)

	// Retrieves information about the specified database branch.
	GetBranch(ctx context.Context, request GetBranchRequest) (*Branch, error)

	// Gets the catalog with the given name. The status reports the Postgres
	// database, project and branch that the catalog registers.
	GetCatalog(ctx context.Context, request GetCatalogRequest) (*Catalog, error)

	// Get a single Lakebase CDF configuration, including the source Postgres
	// schema and target Unity Catalog schema.
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
	ListAssessments(ctx context.Context, request ListAssessmentsRequest) (*ListAssessmentsResponse, error)

	// Returns a paginated list of database branches in the project.
	ListBranches(ctx context.Context, request ListBranchesRequest) (*ListBranchesResponse, error)

	// List all CDF configurations for a Lakebase database. Each configuration
	// maps a Postgres schema to a Unity Catalog schema where the change data
	// feed is materialized.
	ListCdfConfigs(ctx context.Context, request ListCdfConfigsRequest) (*ListCdfConfigsResponse, error)

	// List the per-table CDF statuses within a Lakebase CDF configuration. Each
	// status shows whether a table's change data feed is snapshotting,
	// streaming, or skipped.
	ListCdfStatuses(ctx context.Context, request ListCdfStatusesRequest) (*ListCdfStatusesResponse, error)

	// Returns a list of Postgres databases in the branch.
	ListDatabases(ctx context.Context, request ListDatabasesRequest) (*ListDatabasesResponse, error)

	// Lists the compute endpoints in the branch, including its read-write
	// endpoint and any read-only endpoints.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error)

	// Lists migrations in the workspace.
	ListMigrations(ctx context.Context, request ListMigrationsRequest) (*ListMigrationsResponse, error)

	// Returns a paginated list of database projects in the workspace that the
	// user has permission to access.
	ListProjects(ctx context.Context, request ListProjectsRequest) (*ListProjectsResponse, error)

	// Returns a paginated list of Postgres roles in the branch.
	ListRoles(ctx context.Context, request ListRolesRequest) (*ListRolesResponse, error)

	// Returns a paginated list of snapshots in the project.
	ListSnapshots(ctx context.Context, request ListSnapshotsRequest) (*ListSnapshotsResponse, error)

	// Reveals the current password for a password-authenticated Postgres role.
	RevealRolePassword(ctx context.Context, request RevealRolePasswordRequest) (*RevealRolePasswordResponse, error)

	// Starts the assessed migration named in the request, moving it from
	// READY_TO_START into execution. Its first phase, SCHEMA_LOAD, creates the
	// target databases, roles, and schema on the migration's target Lakebase
	// branch without copying any table data, ahead of INITIAL_DATA_LOAD.
	StartMigration(ctx context.Context, request StartMigrationRequest) (*Operation, error)

	// Undeletes the specified database branch.
	UndeleteBranch(ctx context.Context, request UndeleteBranchRequest) (*Operation, error)

	// Restores a soft-deleted migration before it is permanently purged.
	UndeleteMigration(ctx context.Context, request UndeleteMigrationRequest) (*Operation, error)

	// Restores a soft-deleted project together with its branches and their
	// data, compute endpoints, Postgres databases and roles, project settings,
	// and permissions. A project can be restored until its `purge_time`; after
	// that it is permanently deleted and cannot be recovered.
	UndeleteProject(ctx context.Context, request UndeleteProjectRequest) (*Operation, error)

	// Updates the specified database branch. You can set this branch as the
	// project's default branch, or protect/unprotect it.
	UpdateBranch(ctx context.Context, request UpdateBranchRequest) (*Operation, error)

	// Update Data API configuration for a database.
	UpdateDataApi(ctx context.Context, request UpdateDataApiRequest) (*Operation, error)

	// Updates the specified Postgres database. Only the fields named in
	// `update_mask` change.
	UpdateDatabase(ctx context.Context, request UpdateDatabaseRequest) (*Operation, error)

	// Updates the configuration of the specified compute endpoint. Only the
	// `spec` fields named in `update_mask` change, and `spec.endpoint_type`
	// cannot be changed. The endpoint keeps its hostnames across updates, so
	// existing connection strings stay valid.
	UpdateEndpoint(ctx context.Context, request UpdateEndpointRequest) (*Operation, error)

	// Updates user-owned migration configuration.
	UpdateMigration(ctx context.Context, request UpdateMigrationRequest) (*Operation, error)

	// Updates the configuration of the specified database project. Only the
	// fields named in `update_mask` change.
	UpdateProject(ctx context.Context, request UpdateProjectRequest) (*Operation, error)

	// Updates the mutable fields of a Postgres role for a branch.
	//
	// The fields `spec.identity_type`, `spec.auth_method`, and
	// `spec.postgres_role` are fixed when the role is created. UpdateRole
	// ignores values supplied for these fields in the request body and retains
	// their existing values. Explicitly naming one of these fields in
	// `update_mask` returns an INVALID_PARAMETER_VALUE error.
	UpdateRole(ctx context.Context, request UpdateRoleRequest) (*Operation, error)

	// Sets the snapshot schedule for a branch. The `schedule` field is replaced
	// wholesale; an empty schedule disables automatic snapshots.
	UpdateSnapshotSchedule(ctx context.Context, request UpdateSnapshotScheduleRequest) (*Operation, error)
}
