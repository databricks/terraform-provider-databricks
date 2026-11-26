// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dataquality

import (
	"context"
)

// Manage the data quality of Unity Catalog objects (currently support `schema`
// and `table`)
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DataQualityService interface {

	// Cancels a data quality monitor refresh. Currently only supported for the
	// `table` `object_type`. The call must be made in the same workspace as
	// where the monitor was created.
	//
	// The caller must have either of the following sets of permissions: 1.
	// **MANAGE** and **USE_CATALOG** on the table's parent catalog. 2.
	// **USE_CATALOG** on the table's parent catalog, and **MANAGE** and
	// **USE_SCHEMA** on the table's parent schema. 3. **USE_CATALOG** on the
	// table's parent catalog, **USE_SCHEMA** on the table's parent schema, and
	// **MANAGE** on the table.
	CancelRefresh(ctx context.Context, request CancelRefreshRequest) (*CancelRefreshResponse, error)

	// Create a data quality monitor on a Unity Catalog object. The caller must
	// provide either `anomaly_detection_config` for a schema monitor or
	// `data_profiling_config` for a table monitor.
	//
	// For the `table` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// table's parent catalog, **USE_SCHEMA** on the table's parent schema, and
	// **SELECT** on the table 2. **USE_CATALOG** on the table's parent catalog,
	// **MANAGE** and **USE_SCHEMA** on the table's parent schema, and
	// **SELECT** on the table. 3. **USE_CATALOG** on the table's parent
	// catalog, **USE_SCHEMA** on the table's parent schema, and **MANAGE** and
	// **SELECT** on the table.
	//
	// Workspace assets, such as the dashboard, will be created in the workspace
	// where this call was made.
	//
	// For the `schema` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// schema's parent catalog. 2. **USE_CATALOG** on the schema's parent
	// catalog, and **MANAGE** and **USE_SCHEMA** on the schema.
	CreateMonitor(ctx context.Context, request CreateMonitorRequest) (*Monitor, error)

	// Creates a refresh. Currently only supported for the `table`
	// `object_type`. The call must be made in the same workspace as where the
	// monitor was created.
	//
	// The caller must have either of the following sets of permissions: 1.
	// **MANAGE** and **USE_CATALOG** on the table's parent catalog. 2.
	// **USE_CATALOG** on the table's parent catalog, and **MANAGE** and
	// **USE_SCHEMA** on the table's parent schema. 3. **USE_CATALOG** on the
	// table's parent catalog, **USE_SCHEMA** on the table's parent schema, and
	// **MANAGE** on the table.
	CreateRefresh(ctx context.Context, request CreateRefreshRequest) (*Refresh, error)

	// Delete a data quality monitor on a Unity Catalog object.
	//
	// For the `table` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// table's parent catalog. 2. **USE_CATALOG** on the table's parent catalog,
	// and **MANAGE** and **USE_SCHEMA** on the table's parent schema. 3.
	// **USE_CATALOG** on the table's parent catalog, **USE_SCHEMA** on the
	// table's parent schema, and **MANAGE** on the table.
	//
	// Note that the metric tables and dashboard will not be deleted as part of
	// this call; those assets must be manually cleaned up (if desired).
	//
	// For the `schema` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// schema's parent catalog. 2. **USE_CATALOG** on the schema's parent
	// catalog, and **MANAGE** and **USE_SCHEMA** on the schema.
	DeleteMonitor(ctx context.Context, request DeleteMonitorRequest) error

	// (Unimplemented) Delete a refresh
	DeleteRefresh(ctx context.Context, request DeleteRefreshRequest) error

	// Read a data quality monitor on a Unity Catalog object.
	//
	// For the `table` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// table's parent catalog. 2. **USE_CATALOG** on the table's parent catalog,
	// and **MANAGE** and **USE_SCHEMA** on the table's parent schema. 3.
	// **USE_CATALOG** on the table's parent catalog, **USE_SCHEMA** on the
	// table's parent schema, and **SELECT** on the table.
	//
	// For the `schema` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// schema's parent catalog. 2. **USE_CATALOG** on the schema's parent
	// catalog, and **USE_SCHEMA** on the schema.
	//
	// The returned information includes configuration values on the entity and
	// parent entity as well as information on assets created by the monitor.
	// Some information (e.g. dashboard) may be filtered out if the caller is in
	// a different workspace than where the monitor was created.
	GetMonitor(ctx context.Context, request GetMonitorRequest) (*Monitor, error)

	// Get data quality monitor refresh. The call must be made in the same
	// workspace as where the monitor was created.
	//
	// For the `table` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// table's parent catalog. 2. **USE_CATALOG** on the table's parent catalog,
	// and **MANAGE** and **USE_SCHEMA** on the table's parent schema. 3.
	// **USE_CATALOG** on the table's parent catalog, **USE_SCHEMA** on the
	// table's parent schema, and **SELECT** on the table.
	//
	// For the `schema` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// schema's parent catalog. 2. **USE_CATALOG** on the schema's parent
	// catalog, and **USE_SCHEMA** on the schema.
	GetRefresh(ctx context.Context, request GetRefreshRequest) (*Refresh, error)

	// (Unimplemented) List data quality monitors.
	ListMonitor(ctx context.Context, request ListMonitorRequest) (*ListMonitorResponse, error)

	// List data quality monitor refreshes. The call must be made in the same
	// workspace as where the monitor was created.
	//
	// For the `table` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// table's parent catalog. 2. **USE_CATALOG** on the table's parent catalog,
	// and **MANAGE** and **USE_SCHEMA** on the table's parent schema. 3.
	// **USE_CATALOG** on the table's parent catalog, **USE_SCHEMA** on the
	// table's parent schema, and **SELECT** on the table.
	//
	// For the `schema` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// schema's parent catalog. 2. **USE_CATALOG** on the schema's parent
	// catalog, and **USE_SCHEMA** on the schema.
	ListRefresh(ctx context.Context, request ListRefreshRequest) (*ListRefreshResponse, error)

	// Update a data quality monitor on a Unity Catalog object. Only the fields
	// named in `update_mask` are changed.
	//
	// For the `table` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// table's parent catalog. 2. **USE_CATALOG** on the table's parent catalog,
	// and **MANAGE** and **USE_SCHEMA** on the table's parent schema. 3.
	// **USE_CATALOG** on the table's parent catalog, **USE_SCHEMA** on the
	// table's parent schema, and **MANAGE** on the table.
	//
	// For the `schema` `object_type`, the caller must have either of the
	// following sets of permissions: 1. **MANAGE** and **USE_CATALOG** on the
	// schema's parent catalog. 2. **USE_CATALOG** on the schema's parent
	// catalog, and **MANAGE** and **USE_SCHEMA** on the schema.
	UpdateMonitor(ctx context.Context, request UpdateMonitorRequest) (*Monitor, error)

	// (Unimplemented) Update a refresh
	UpdateRefresh(ctx context.Context, request UpdateRefreshRequest) (*Refresh, error)
}
