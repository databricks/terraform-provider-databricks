// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"context"
)

// Database Instances provide access to a database via REST API or direct SQL.
//
// For new workloads, Databricks recommends the [Postgres API], which supports
// the latest Lakebase features.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Postgres API]: https://docs.databricks.com/api/postgres/v1/postgres
type DatabaseService interface {

	// Create a Database Catalog.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Create a catalog].
	//
	// [Create a catalog]: https://docs.databricks.com/api/postgres/v1/create-catalog
	CreateDatabaseCatalog(ctx context.Context, request CreateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Create a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Create a project].
	//
	// [Create a project]: https://docs.databricks.com/api/postgres/v1/create-project
	CreateDatabaseInstance(ctx context.Context, request CreateDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Create a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Create a role].
	//
	// [Create a role]: https://docs.databricks.com/api/postgres/v1/create-role
	CreateDatabaseInstanceRole(ctx context.Context, request CreateDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Create a Database Table. This registers a table that already exists in a
	// Postgres database of a database instance as a table in Unity Catalog. The
	// Postgres database, schema and table must already exist; the request does
	// not create them. Repeating the request for the same table succeeds, and a
	// request for a name that is already registered from a different database
	// or instance fails. To create a Postgres table from a Unity Catalog source
	// table instead, use [Create a Synced Database Table].
	//
	// [Create a Synced Database Table]: https://docs.databricks.com/api/workspace/database/createsynceddatabasetable
	CreateDatabaseTable(ctx context.Context, request CreateDatabaseTableRequest) (*DatabaseTable, error)

	// Create a Synced Database Table.
	//
	// The table and its data synchronization pipeline are provisioned
	// asynchronously. The response reflects the initial state; use Get a Synced
	// Database Table to follow the data synchronization status until the table
	// is online.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Create a synced table].
	//
	// [Create a synced table]: https://docs.databricks.com/api/postgres/v1/create-synced-table
	CreateSyncedDatabaseTable(ctx context.Context, request CreateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// Delete a Database Catalog. This removes the catalog from Unity Catalog,
	// while the Postgres database and its tables remain in the instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Delete a catalog].
	//
	// [Delete a catalog]: https://docs.databricks.com/api/postgres/v1/delete-catalog
	DeleteDatabaseCatalog(ctx context.Context, request DeleteDatabaseCatalogRequest) error

	// Delete a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Delete a project].
	//
	// [Delete a project]: https://docs.databricks.com/api/postgres/v1/delete-project
	DeleteDatabaseInstance(ctx context.Context, request DeleteDatabaseInstanceRequest) error

	// Deletes a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Delete a role].
	//
	// [Delete a role]: https://docs.databricks.com/api/postgres/v1/delete-role
	DeleteDatabaseInstanceRole(ctx context.Context, request DeleteDatabaseInstanceRoleRequest) error

	// Delete a Database Table. This removes the table from Unity Catalog only;
	// the Postgres table and its data are kept. Tables in a database catalog
	// cannot be deleted with this method, and synced tables must be deleted
	// with [Delete a Synced Database Table].
	//
	// [Delete a Synced Database Table]: https://docs.databricks.com/api/workspace/database/deletesynceddatabasetable
	DeleteDatabaseTable(ctx context.Context, request DeleteDatabaseTableRequest) error

	// Delete a Synced Database Table.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Delete a synced table].
	//
	// [Delete a synced table]: https://docs.databricks.com/api/postgres/v1/delete-synced-table
	DeleteSyncedDatabaseTable(ctx context.Context, request DeleteSyncedDatabaseTableRequest) error

	// Find a Database Instance by its uid rather than by name.
	FindDatabaseInstanceByUid(ctx context.Context, request FindDatabaseInstanceByUidRequest) (*DatabaseInstance, error)

	// Generates a short-lived OAuth token for connecting to database instances.
	// Connect with the caller's Databricks identity as the Postgres user and
	// the returned token as the password. For details about OAuth token
	// authentication, see [About authentication].
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Generate a database credential].
	//
	// [About authentication]: https://docs.databricks.com/aws/en/oltp/projects/authentication
	// [Generate a database credential]: https://docs.databricks.com/api/postgres/v1/generate-database-credential
	GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error)

	// Get a Database Catalog.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Get a catalog].
	//
	// [Get a catalog]: https://docs.databricks.com/api/postgres/v1/get-catalog
	GetDatabaseCatalog(ctx context.Context, request GetDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Get a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Get a project].
	//
	// [Get a project]: https://docs.databricks.com/api/postgres/v1/get-project
	GetDatabaseInstance(ctx context.Context, request GetDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Gets a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Get a role].
	//
	// [Get a role]: https://docs.databricks.com/api/postgres/v1/get-role
	GetDatabaseInstanceRole(ctx context.Context, request GetDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Get a Database Table. The request fails if the name refers to a synced
	// table or to a table that is not a database table; for synced tables, use
	// [Get a Synced Database Table].
	//
	// [Get a Synced Database Table]: https://docs.databricks.com/api/workspace/database/getsynceddatabasetable
	GetDatabaseTable(ctx context.Context, request GetDatabaseTableRequest) (*DatabaseTable, error)

	// Get a Synced Database Table.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Get a synced table].
	//
	// [Get a synced table]: https://docs.databricks.com/api/postgres/v1/get-synced-table
	GetSyncedDatabaseTable(ctx context.Context, request GetSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	ListDatabaseCatalogs(ctx context.Context, request ListDatabaseCatalogsRequest) (*ListDatabaseCatalogsResponse, error)

	// List roles for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [List roles].
	//
	// [List roles]: https://docs.databricks.com/api/postgres/v1/list-roles
	ListDatabaseInstanceRoles(ctx context.Context, request ListDatabaseInstanceRolesRequest) (*ListDatabaseInstanceRolesResponse, error)

	// List Database Instances.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [List projects].
	//
	// [List projects]: https://docs.databricks.com/api/postgres/v1/list-projects
	ListDatabaseInstances(ctx context.Context, request ListDatabaseInstancesRequest) (*ListDatabaseInstancesResponse, error)

	// List synced database tables in a Database Instance.
	ListSyncedDatabaseTables(ctx context.Context, request ListSyncedDatabaseTablesRequest) (*ListSyncedDatabaseTablesResponse, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	UpdateDatabaseCatalog(ctx context.Context, request UpdateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Update a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalents:
	// [Update a project] and [Update an endpoint].
	//
	// [Update a project]: https://docs.databricks.com/api/postgres/v1/update-project
	// [Update an endpoint]: https://docs.databricks.com/api/postgres/v1/update-endpoint
	UpdateDatabaseInstance(ctx context.Context, request UpdateDatabaseInstanceRequest) (*DatabaseInstance, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	UpdateSyncedDatabaseTable(ctx context.Context, request UpdateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)
}
