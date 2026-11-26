// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Database Instances provide access to a database via REST API or direct SQL.
//
// For new workloads, Databricks recommends the [Postgres API], which supports
// the latest Lakebase features.
//
// [Postgres API]: https://docs.databricks.com/api/postgres/v1/postgres
package database

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type DatabaseInterface interface {

	// WaitGetDatabaseInstanceDatabaseAvailable repeatedly calls [DatabaseAPI.GetDatabaseInstance] and waits to reach AVAILABLE state
	WaitGetDatabaseInstanceDatabaseAvailable(ctx context.Context, name string,
		timeout time.Duration, callback func(*DatabaseInstance)) (*DatabaseInstance, error)

	// Create a Database Catalog.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Create
	// a catalog].
	//
	// [Create a catalog]: https://docs.databricks.com/api/postgres/v1/create-catalog
	CreateDatabaseCatalog(ctx context.Context, request CreateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Create a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Create
	// a project].
	//
	// [Create a project]: https://docs.databricks.com/api/postgres/v1/create-project
	CreateDatabaseInstance(ctx context.Context, createDatabaseInstanceRequest CreateDatabaseInstanceRequest) (*WaitGetDatabaseInstanceDatabaseAvailable[DatabaseInstance], error)

	// Calls [DatabaseAPIInterface.CreateDatabaseInstance] and waits to reach AVAILABLE state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[DatabaseInstance](60*time.Minute) functional option.
	//
	// Deprecated: use [DatabaseAPIInterface.CreateDatabaseInstance].Get() or [DatabaseAPIInterface.WaitGetDatabaseInstanceDatabaseAvailable]
	CreateDatabaseInstanceAndWait(ctx context.Context, createDatabaseInstanceRequest CreateDatabaseInstanceRequest, options ...retries.Option[DatabaseInstance]) (*DatabaseInstance, error)

	// Create a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Create
	// a role].
	//
	// [Create a role]: https://docs.databricks.com/api/postgres/v1/create-role
	CreateDatabaseInstanceRole(ctx context.Context, request CreateDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Create a Database Table. This registers a table that already exists in a
	// Postgres database of a database instance as a table in Unity Catalog. The
	// Postgres database, schema and table must already exist; the request does not
	// create them. Repeating the request for the same table succeeds, and a request
	// for a name that is already registered from a different database or instance
	// fails. To create a Postgres table from a Unity Catalog source table instead,
	// use [Create a Synced Database Table].
	//
	// [Create a Synced Database Table]: https://docs.databricks.com/api/workspace/database/createsynceddatabasetable
	CreateDatabaseTable(ctx context.Context, request CreateDatabaseTableRequest) (*DatabaseTable, error)

	// Create a Synced Database Table.
	//
	// The table and its data synchronization pipeline are provisioned
	// asynchronously. The response reflects the initial state; use Get a Synced
	// Database Table to follow the data synchronization status until the table is
	// online.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Create
	// a synced table].
	//
	// [Create a synced table]: https://docs.databricks.com/api/postgres/v1/create-synced-table
	CreateSyncedDatabaseTable(ctx context.Context, request CreateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// Delete a Database Catalog. This removes the catalog from Unity Catalog, while
	// the Postgres database and its tables remain in the instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
	// a catalog].
	//
	// [Delete a catalog]: https://docs.databricks.com/api/postgres/v1/delete-catalog
	DeleteDatabaseCatalog(ctx context.Context, request DeleteDatabaseCatalogRequest) error

	// Delete a Database Catalog. This removes the catalog from Unity Catalog, while
	// the Postgres database and its tables remain in the instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
	// a catalog].
	//
	// [Delete a catalog]: https://docs.databricks.com/api/postgres/v1/delete-catalog
	DeleteDatabaseCatalogByName(ctx context.Context, name string) error

	// Delete a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
	// a project].
	//
	// [Delete a project]: https://docs.databricks.com/api/postgres/v1/delete-project
	DeleteDatabaseInstance(ctx context.Context, request DeleteDatabaseInstanceRequest) error

	// Delete a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
	// a project].
	//
	// [Delete a project]: https://docs.databricks.com/api/postgres/v1/delete-project
	DeleteDatabaseInstanceByName(ctx context.Context, name string) error

	// Deletes a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
	// a role].
	//
	// [Delete a role]: https://docs.databricks.com/api/postgres/v1/delete-role
	DeleteDatabaseInstanceRole(ctx context.Context, request DeleteDatabaseInstanceRoleRequest) error

	// Deletes a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
	// a role].
	//
	// [Delete a role]: https://docs.databricks.com/api/postgres/v1/delete-role
	DeleteDatabaseInstanceRoleByInstanceNameAndName(ctx context.Context, instanceName string, name string) error

	// Delete a Database Table. This removes the table from Unity Catalog only; the
	// Postgres table and its data are kept. Tables in a database catalog cannot be
	// deleted with this method, and synced tables must be deleted with [Delete a
	// Synced Database Table].
	//
	// [Delete a Synced Database Table]: https://docs.databricks.com/api/workspace/database/deletesynceddatabasetable
	DeleteDatabaseTable(ctx context.Context, request DeleteDatabaseTableRequest) error

	// Delete a Database Table. This removes the table from Unity Catalog only; the
	// Postgres table and its data are kept. Tables in a database catalog cannot be
	// deleted with this method, and synced tables must be deleted with [Delete a
	// Synced Database Table].
	//
	// [Delete a Synced Database Table]: https://docs.databricks.com/api/workspace/database/deletesynceddatabasetable
	DeleteDatabaseTableByName(ctx context.Context, name string) error

	// Delete a Synced Database Table.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
	// a synced table].
	//
	// [Delete a synced table]: https://docs.databricks.com/api/postgres/v1/delete-synced-table
	DeleteSyncedDatabaseTable(ctx context.Context, request DeleteSyncedDatabaseTableRequest) error

	// Delete a Synced Database Table.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
	// a synced table].
	//
	// [Delete a synced table]: https://docs.databricks.com/api/postgres/v1/delete-synced-table
	DeleteSyncedDatabaseTableByName(ctx context.Context, name string) error

	// Find a Database Instance by its uid rather than by name.
	FindDatabaseInstanceByUid(ctx context.Context, request FindDatabaseInstanceByUidRequest) (*DatabaseInstance, error)

	// Generates a short-lived OAuth token for connecting to database instances.
	// Connect with the caller's Databricks identity as the Postgres user and the
	// returned token as the password. For details about OAuth token authentication,
	// see [About authentication].
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Generate a database credential].
	//
	// [About authentication]: https://docs.databricks.com/aws/en/oltp/projects/authentication
	// [Generate a database credential]: https://docs.databricks.com/api/postgres/v1/generate-database-credential
	GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error)

	// Get a Database Catalog.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
	// catalog].
	//
	// [Get a catalog]: https://docs.databricks.com/api/postgres/v1/get-catalog
	GetDatabaseCatalog(ctx context.Context, request GetDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Get a Database Catalog.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
	// catalog].
	//
	// [Get a catalog]: https://docs.databricks.com/api/postgres/v1/get-catalog
	GetDatabaseCatalogByName(ctx context.Context, name string) (*DatabaseCatalog, error)

	// Get a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
	// project].
	//
	// [Get a project]: https://docs.databricks.com/api/postgres/v1/get-project
	GetDatabaseInstance(ctx context.Context, request GetDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Get a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
	// project].
	//
	// [Get a project]: https://docs.databricks.com/api/postgres/v1/get-project
	GetDatabaseInstanceByName(ctx context.Context, name string) (*DatabaseInstance, error)

	// Gets a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
	// role].
	//
	// [Get a role]: https://docs.databricks.com/api/postgres/v1/get-role
	GetDatabaseInstanceRole(ctx context.Context, request GetDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Gets a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
	// role].
	//
	// [Get a role]: https://docs.databricks.com/api/postgres/v1/get-role
	GetDatabaseInstanceRoleByInstanceNameAndName(ctx context.Context, instanceName string, name string) (*DatabaseInstanceRole, error)

	// Get a Database Table. The request fails if the name refers to a synced table
	// or to a table that is not a database table; for synced tables, use [Get a
	// Synced Database Table].
	//
	// [Get a Synced Database Table]: https://docs.databricks.com/api/workspace/database/getsynceddatabasetable
	GetDatabaseTable(ctx context.Context, request GetDatabaseTableRequest) (*DatabaseTable, error)

	// Get a Database Table. The request fails if the name refers to a synced table
	// or to a table that is not a database table; for synced tables, use [Get a
	// Synced Database Table].
	//
	// [Get a Synced Database Table]: https://docs.databricks.com/api/workspace/database/getsynceddatabasetable
	GetDatabaseTableByName(ctx context.Context, name string) (*DatabaseTable, error)

	// Get a Synced Database Table.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
	// synced table].
	//
	// [Get a synced table]: https://docs.databricks.com/api/postgres/v1/get-synced-table
	GetSyncedDatabaseTable(ctx context.Context, request GetSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// Get a Synced Database Table.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
	// synced table].
	//
	// [Get a synced table]: https://docs.databricks.com/api/postgres/v1/get-synced-table
	GetSyncedDatabaseTableByName(ctx context.Context, name string) (*SyncedDatabaseTable, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseCatalogs(ctx context.Context, request ListDatabaseCatalogsRequest) listing.Iterator[DatabaseCatalog]

	// This API is currently unimplemented, but exposed for Terraform support.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseCatalogsAll(ctx context.Context, request ListDatabaseCatalogsRequest) ([]DatabaseCatalog, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	ListDatabaseCatalogsByInstanceName(ctx context.Context, instanceName string) (*ListDatabaseCatalogsResponse, error)

	// List roles for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [List
	// roles].
	//
	// [List roles]: https://docs.databricks.com/api/postgres/v1/list-roles
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseInstanceRoles(ctx context.Context, request ListDatabaseInstanceRolesRequest) listing.Iterator[DatabaseInstanceRole]

	// List roles for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [List
	// roles].
	//
	// [List roles]: https://docs.databricks.com/api/postgres/v1/list-roles
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseInstanceRolesAll(ctx context.Context, request ListDatabaseInstanceRolesRequest) ([]DatabaseInstanceRole, error)

	// List roles for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [List
	// roles].
	//
	// [List roles]: https://docs.databricks.com/api/postgres/v1/list-roles
	ListDatabaseInstanceRolesByInstanceName(ctx context.Context, instanceName string) (*ListDatabaseInstanceRolesResponse, error)

	// List Database Instances.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [List
	// projects].
	//
	// [List projects]: https://docs.databricks.com/api/postgres/v1/list-projects
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseInstances(ctx context.Context, request ListDatabaseInstancesRequest) listing.Iterator[DatabaseInstance]

	// List Database Instances.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent: [List
	// projects].
	//
	// [List projects]: https://docs.databricks.com/api/postgres/v1/list-projects
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseInstancesAll(ctx context.Context, request ListDatabaseInstancesRequest) ([]DatabaseInstance, error)

	// List synced database tables in a Database Instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSyncedDatabaseTables(ctx context.Context, request ListSyncedDatabaseTablesRequest) listing.Iterator[SyncedDatabaseTable]

	// List synced database tables in a Database Instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSyncedDatabaseTablesAll(ctx context.Context, request ListSyncedDatabaseTablesRequest) ([]SyncedDatabaseTable, error)

	// List synced database tables in a Database Instance.
	ListSyncedDatabaseTablesByInstanceName(ctx context.Context, instanceName string) (*ListSyncedDatabaseTablesResponse, error)

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

func NewDatabase(client *client.DatabricksClient) *DatabaseAPI {
	return &DatabaseAPI{
		databaseImpl: databaseImpl{
			client: client,
		},
	}
}

// Database Instances provide access to a database via REST API or direct SQL.
//
// For new workloads, Databricks recommends the [Postgres API], which supports
// the latest Lakebase features.
//
// [Postgres API]: https://docs.databricks.com/api/postgres/v1/postgres
type DatabaseAPI struct {
	databaseImpl
}

// WaitGetDatabaseInstanceDatabaseAvailable repeatedly calls [DatabaseAPI.GetDatabaseInstance] and waits to reach AVAILABLE state
func (a *DatabaseAPI) WaitGetDatabaseInstanceDatabaseAvailable(ctx context.Context, name string,
	timeout time.Duration, callback func(*DatabaseInstance)) (*DatabaseInstance, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[DatabaseInstance](ctx, timeout, func() (*DatabaseInstance, *retries.Err) {
		databaseInstance, err := a.GetDatabaseInstance(ctx, GetDatabaseInstanceRequest{
			Name: name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(databaseInstance)
		}
		status := databaseInstance.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case DatabaseInstanceStateAvailable: // target state
			return databaseInstance, nil
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetDatabaseInstanceDatabaseAvailable is a wrapper that calls [DatabaseAPI.WaitGetDatabaseInstanceDatabaseAvailable] and waits to reach AVAILABLE state.
type WaitGetDatabaseInstanceDatabaseAvailable[R any] struct {
	Response *R
	Name     string `json:"name"`
	Poll     func(time.Duration, func(*DatabaseInstance)) (*DatabaseInstance, error)
	callback func(*DatabaseInstance)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetDatabaseInstanceDatabaseAvailable[R]) OnProgress(callback func(*DatabaseInstance)) *WaitGetDatabaseInstanceDatabaseAvailable[R] {
	w.callback = callback
	return w
}

// Get the DatabaseInstance with the default timeout of 20 minutes.
func (w *WaitGetDatabaseInstanceDatabaseAvailable[R]) Get() (*DatabaseInstance, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the DatabaseInstance with custom timeout.
func (w *WaitGetDatabaseInstanceDatabaseAvailable[R]) GetWithTimeout(timeout time.Duration) (*DatabaseInstance, error) {
	return w.Poll(timeout, w.callback)
}

// Create a Database Instance.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [Create
// a project].
//
// [Create a project]: https://docs.databricks.com/api/postgres/v1/create-project
func (a *DatabaseAPI) CreateDatabaseInstance(ctx context.Context, createDatabaseInstanceRequest CreateDatabaseInstanceRequest) (*WaitGetDatabaseInstanceDatabaseAvailable[DatabaseInstance], error) {
	databaseInstance, err := a.databaseImpl.CreateDatabaseInstance(ctx, createDatabaseInstanceRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetDatabaseInstanceDatabaseAvailable[DatabaseInstance]{
		Response: databaseInstance,
		Name:     databaseInstance.Name,
		Poll: func(timeout time.Duration, callback func(*DatabaseInstance)) (*DatabaseInstance, error) {
			return a.WaitGetDatabaseInstanceDatabaseAvailable(ctx, databaseInstance.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [DatabaseAPI.CreateDatabaseInstance] and waits to reach AVAILABLE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[DatabaseInstance](60*time.Minute) functional option.
//
// Deprecated: use [DatabaseAPI.CreateDatabaseInstance].Get() or [DatabaseAPI.WaitGetDatabaseInstanceDatabaseAvailable]
func (a *DatabaseAPI) CreateDatabaseInstanceAndWait(ctx context.Context, createDatabaseInstanceRequest CreateDatabaseInstanceRequest, options ...retries.Option[DatabaseInstance]) (*DatabaseInstance, error) {
	wait, err := a.CreateDatabaseInstance(ctx, createDatabaseInstanceRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[DatabaseInstance]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *DatabaseInstance) {
		for _, o := range options {
			o(&retries.Info[DatabaseInstance]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Delete a Database Catalog. This removes the catalog from Unity Catalog, while
// the Postgres database and its tables remain in the instance.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
// a catalog].
//
// [Delete a catalog]: https://docs.databricks.com/api/postgres/v1/delete-catalog
func (a *DatabaseAPI) DeleteDatabaseCatalogByName(ctx context.Context, name string) error {
	return a.databaseImpl.DeleteDatabaseCatalog(ctx, DeleteDatabaseCatalogRequest{
		Name: name,
	})
}

// Delete a Database Instance.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
// a project].
//
// [Delete a project]: https://docs.databricks.com/api/postgres/v1/delete-project
func (a *DatabaseAPI) DeleteDatabaseInstanceByName(ctx context.Context, name string) error {
	return a.databaseImpl.DeleteDatabaseInstance(ctx, DeleteDatabaseInstanceRequest{
		Name: name,
	})
}

// Deletes a role for a Database Instance.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
// a role].
//
// [Delete a role]: https://docs.databricks.com/api/postgres/v1/delete-role
func (a *DatabaseAPI) DeleteDatabaseInstanceRoleByInstanceNameAndName(ctx context.Context, instanceName string, name string) error {
	return a.databaseImpl.DeleteDatabaseInstanceRole(ctx, DeleteDatabaseInstanceRoleRequest{
		InstanceName: instanceName,
		Name:         name,
	})
}

// Delete a Database Table. This removes the table from Unity Catalog only; the
// Postgres table and its data are kept. Tables in a database catalog cannot be
// deleted with this method, and synced tables must be deleted with [Delete a
// Synced Database Table].
//
// [Delete a Synced Database Table]: https://docs.databricks.com/api/workspace/database/deletesynceddatabasetable
func (a *DatabaseAPI) DeleteDatabaseTableByName(ctx context.Context, name string) error {
	return a.databaseImpl.DeleteDatabaseTable(ctx, DeleteDatabaseTableRequest{
		Name: name,
	})
}

// Delete a Synced Database Table.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [Delete
// a synced table].
//
// [Delete a synced table]: https://docs.databricks.com/api/postgres/v1/delete-synced-table
func (a *DatabaseAPI) DeleteSyncedDatabaseTableByName(ctx context.Context, name string) error {
	return a.databaseImpl.DeleteSyncedDatabaseTable(ctx, DeleteSyncedDatabaseTableRequest{
		Name: name,
	})
}

// Get a Database Catalog.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
// catalog].
//
// [Get a catalog]: https://docs.databricks.com/api/postgres/v1/get-catalog
func (a *DatabaseAPI) GetDatabaseCatalogByName(ctx context.Context, name string) (*DatabaseCatalog, error) {
	return a.databaseImpl.GetDatabaseCatalog(ctx, GetDatabaseCatalogRequest{
		Name: name,
	})
}

// Get a Database Instance.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
// project].
//
// [Get a project]: https://docs.databricks.com/api/postgres/v1/get-project
func (a *DatabaseAPI) GetDatabaseInstanceByName(ctx context.Context, name string) (*DatabaseInstance, error) {
	return a.databaseImpl.GetDatabaseInstance(ctx, GetDatabaseInstanceRequest{
		Name: name,
	})
}

// Gets a role for a Database Instance.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
// role].
//
// [Get a role]: https://docs.databricks.com/api/postgres/v1/get-role
func (a *DatabaseAPI) GetDatabaseInstanceRoleByInstanceNameAndName(ctx context.Context, instanceName string, name string) (*DatabaseInstanceRole, error) {
	return a.databaseImpl.GetDatabaseInstanceRole(ctx, GetDatabaseInstanceRoleRequest{
		InstanceName: instanceName,
		Name:         name,
	})
}

// Get a Database Table. The request fails if the name refers to a synced table
// or to a table that is not a database table; for synced tables, use [Get a
// Synced Database Table].
//
// [Get a Synced Database Table]: https://docs.databricks.com/api/workspace/database/getsynceddatabasetable
func (a *DatabaseAPI) GetDatabaseTableByName(ctx context.Context, name string) (*DatabaseTable, error) {
	return a.databaseImpl.GetDatabaseTable(ctx, GetDatabaseTableRequest{
		Name: name,
	})
}

// Get a Synced Database Table.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [Get a
// synced table].
//
// [Get a synced table]: https://docs.databricks.com/api/postgres/v1/get-synced-table
func (a *DatabaseAPI) GetSyncedDatabaseTableByName(ctx context.Context, name string) (*SyncedDatabaseTable, error) {
	return a.databaseImpl.GetSyncedDatabaseTable(ctx, GetSyncedDatabaseTableRequest{
		Name: name,
	})
}

// This API is currently unimplemented, but exposed for Terraform support.
func (a *DatabaseAPI) ListDatabaseCatalogsByInstanceName(ctx context.Context, instanceName string) (*ListDatabaseCatalogsResponse, error) {
	return a.databaseImpl.internalListDatabaseCatalogs(ctx, ListDatabaseCatalogsRequest{
		InstanceName: instanceName,
	})
}

// List roles for a Database Instance.
//
// For new workloads, Databricks recommends the Postgres API equivalent: [List
// roles].
//
// [List roles]: https://docs.databricks.com/api/postgres/v1/list-roles
func (a *DatabaseAPI) ListDatabaseInstanceRolesByInstanceName(ctx context.Context, instanceName string) (*ListDatabaseInstanceRolesResponse, error) {
	return a.databaseImpl.internalListDatabaseInstanceRoles(ctx, ListDatabaseInstanceRolesRequest{
		InstanceName: instanceName,
	})
}

// List synced database tables in a Database Instance.
func (a *DatabaseAPI) ListSyncedDatabaseTablesByInstanceName(ctx context.Context, instanceName string) (*ListSyncedDatabaseTablesResponse, error) {
	return a.databaseImpl.internalListSyncedDatabaseTables(ctx, ListSyncedDatabaseTablesRequest{
		InstanceName: instanceName,
	})
}
