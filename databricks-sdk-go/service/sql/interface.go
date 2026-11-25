// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"context"
)

// The alerts API can be used to perform CRUD operations on alerts. An alert is
// a Databricks SQL object that periodically runs a query, evaluates a condition
// of its result, and notifies one or more users and/or notification
// destinations if the condition was met. Alerts can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AlertsService interface {

	// Creates an alert.
	Create(ctx context.Context, request CreateAlertRequest) (*Alert, error)

	// Moves an alert to the trash. Trashed alerts immediately disappear from
	// searches and list views, and can no longer trigger. You can restore a
	// trashed alert through the UI. A trashed alert is permanently deleted
	// after 30 days.
	Delete(ctx context.Context, request TrashAlertRequest) error

	// Gets an alert.
	Get(ctx context.Context, request GetAlertRequest) (*Alert, error)

	// Gets a list of alerts accessible to the user, ordered by creation time.
	// **Warning:** Calling this API concurrently 10 or more times could result
	// in throttling, service degradation, or a temporary ban.
	List(ctx context.Context, request ListAlertsRequest) (*ListAlertsResponse, error)

	// Updates the fields of an alert named in `update_mask`. Changing
	// `owner_user_name` requires workspace admin privileges; changing any other
	// field requires CAN MANAGE permission on the alert.
	Update(ctx context.Context, request UpdateAlertRequest) (*Alert, error)
}

// The alerts API can be used to perform CRUD operations on alerts. An alert is
// a Databricks SQL object that periodically runs a query, evaluates a condition
// of its result, and notifies one or more users and/or notification
// destinations if the condition was met. Alerts can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
//
// **Warning**: This API is deprecated. Please see the latest version of the
// Databricks SQL API. [Learn more]
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type AlertsLegacyService interface {

	// Creates an alert. An alert is a Databricks SQL object that periodically
	// runs a query, evaluates a condition of its result, and notifies users or
	// notification destinations if the condition was met.
	//
	// **Warning**: This API is deprecated. Please use :method:alerts/create
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Create(ctx context.Context, request CreateAlert) (*LegacyAlert, error)

	// Deletes an alert. Deleted alerts are no longer accessible and cannot be
	// restored. **Note**: Unlike queries and dashboards, alerts cannot be moved
	// to the trash.
	//
	// **Warning**: This API is deprecated. Please use :method:alerts/delete
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Delete(ctx context.Context, request DeleteAlertsLegacyRequest) error

	// Gets an alert.
	//
	// **Warning**: This API is deprecated. Please use :method:alerts/get
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Get(ctx context.Context, request GetAlertsLegacyRequest) (*LegacyAlert, error)

	// Gets a list of alerts.
	//
	// **Warning**: This API is deprecated. Please use :method:alerts/list
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	List(ctx context.Context) ([]LegacyAlert, error)

	// Updates an alert.
	//
	// **Warning**: This API is deprecated. Please use :method:alerts/update
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Update(ctx context.Context, request EditAlert) error
}

// New version of SQL Alerts
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AlertsV2Service interface {

	// Creates an alert that runs `query_text` on the given schedule and
	// evaluates the result against the configured condition.
	CreateAlert(ctx context.Context, request CreateAlertV2Request) (*AlertV2, error)

	// Gets an alert.
	GetAlert(ctx context.Context, request GetAlertV2Request) (*AlertV2, error)

	// Gets a list of alerts accessible to the user, ordered by creation time.
	ListAlerts(ctx context.Context, request ListAlertsV2Request) (*ListAlertsV2Response, error)

	// Moves an alert to the trash. Trashed alerts immediately disappear from
	// list views, and can no longer trigger. You can restore a trashed alert
	// through the UI. A trashed alert is permanently deleted after 30 days.
	TrashAlert(ctx context.Context, request TrashAlertV2Request) error

	// Updates an alert. Only the fields listed in `update_mask` are changed;
	// all other fields keep their current values.
	UpdateAlert(ctx context.Context, request UpdateAlertV2Request) (*AlertV2, error)
}

// This is an evolving API that facilitates the addition and removal of widgets
// from existing dashboards within the Databricks Workspace. Data structures may
// change over time.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DashboardWidgetsService interface {

	// Adds a widget to a dashboard
	Create(ctx context.Context, request CreateWidget) (*Widget, error)

	// Removes a widget from a dashboard
	Delete(ctx context.Context, request DeleteDashboardWidgetRequest) error

	// Updates an existing widget
	Update(ctx context.Context, request UpdateWidgetRequest) (*Widget, error)
}

// In general, there is little need to modify dashboards using the API. However,
// it can be useful to use dashboard objects to look-up a collection of related
// query IDs. The API can also be used to duplicate multiple dashboards at once
// since you can get a dashboard definition with a GET request and then POST it
// to create a new one. Dashboards can be scheduled using the `sql_task` type of
// the Jobs API, e.g. :method:jobs/create.
//
// **Warning**: This API is deprecated. Please use the AI/BI Dashboards API
// instead. [Learn more]
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Learn more]: https://docs.databricks.com/en/dashboards/
type DashboardsService interface {

	// Moves a dashboard to the trash. Trashed dashboards do not appear in list
	// views or searches, and cannot be shared.
	//
	// **Warning**: This API is deprecated. Please use the AI/BI Dashboards API
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/dashboards/
	Delete(ctx context.Context, request DeleteDashboardRequest) error

	// Returns a JSON representation of a dashboard object, including its
	// visualization and query objects.
	//
	// **Warning**: This API is deprecated. Please use the AI/BI Dashboards API
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/dashboards/
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Fetch a paginated list of dashboard objects.
	//
	// **Warning**: Calling this API concurrently 10 or more times could result
	// in throttling, service degradation, or a temporary ban.
	//
	// **Warning**: This API is deprecated. Please use the AI/BI Dashboards API
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/dashboards/
	List(ctx context.Context, request ListDashboardsRequest) (*ListResponse, error)

	// A restored dashboard appears in list views and searches and can be
	// shared.
	//
	// **Warning**: This API is deprecated. Please use the AI/BI Dashboards API
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/dashboards/
	Restore(ctx context.Context, request RestoreDashboardRequest) error

	// Modify this dashboard definition. This operation only affects attributes
	// of the dashboard object. It does not add, modify, or remove widgets.
	//
	// **Note**: You cannot undo this operation.
	//
	// **Warning**: This API is deprecated. Please use the AI/BI Dashboards API
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/dashboards/
	Update(ctx context.Context, request DashboardEditContent) (*Dashboard, error)
}

// This API is provided to assist you in making new query objects. When creating
// a query object, you may optionally specify a `data_source_id` for the SQL
// warehouse against which it will run. If you don't already know the
// `data_source_id` for your desired SQL warehouse, this API will help you find
// it.
//
// This API does not support searches. It returns the full list of SQL
// warehouses in your workspace. We advise you to use any text editor, REST
// client, or `grep` to search the response from this API for the name of your
// SQL warehouse as it appears in Databricks SQL.
//
// **Warning**: This API is deprecated. Please see the latest version of the
// Databricks SQL API. [Learn more]
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type DataSourcesService interface {

	// Retrieves a full list of SQL warehouses available in this workspace. All
	// fields that appear in this API response are enumerated for clarity.
	// However, you need only a SQL warehouse's `id` to create new queries
	// against it.
	//
	// **Warning**: This API is deprecated. Please use :method:warehouses/list
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	List(ctx context.Context) ([]DataSource, error)
}

// The SQL Permissions API is similar to the endpoints of the
// :method:permissions/set. However, this exposes only one endpoint, which gets
// the Access Control List for a given object. You cannot modify any permissions
// using this API.
//
// There are three levels of permission:
//
// - `CAN_VIEW`: Allows read-only access
//
// - `CAN_RUN`: Allows read access and run access (superset of `CAN_VIEW`)
//
// - `CAN_MANAGE`: Allows all actions: read, run, edit, delete, modify
// permissions (superset of `CAN_RUN`)
//
// **Warning**: This API is deprecated. Please see the latest version of the
// Databricks SQL API. [Learn more]
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type DbsqlPermissionsService interface {

	// Gets a JSON representation of the access control list (ACL) for a
	// specified object.
	//
	// **Warning**: This API is deprecated. Please use
	// :method:workspace/getpermissions instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Get(ctx context.Context, request GetDbsqlPermissionRequest) (*GetResponse, error)

	// Sets the access control list (ACL) for a specified object. This operation
	// will complete rewrite the ACL.
	//
	// **Warning**: This API is deprecated. Please use
	// :method:workspace/setpermissions instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Set(ctx context.Context, request SetRequest) (*SetResponse, error)

	// Transfers ownership of a dashboard, query, or alert to an active user.
	// Requires an admin API key.
	//
	// **Warning**: This API is deprecated. For queries and alerts, please use
	// :method:queries/update and :method:alerts/update respectively instead.
	// [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error)
}

// The queries API can be used to perform CRUD operations on queries. A query is
// a Databricks SQL object that includes the target SQL warehouse, query text,
// name, description, tags, and parameters. Queries can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type QueriesService interface {

	// Creates a query. The caller becomes the owner of the new query.
	Create(ctx context.Context, request CreateQueryRequest) (*Query, error)

	// Moves a query to the trash. Trashed queries immediately disappear from
	// searches and list views, and cannot be used for alerts. You can restore a
	// trashed query through the UI. A trashed query is permanently deleted
	// after 30 days.
	Delete(ctx context.Context, request TrashQueryRequest) error

	// Gets a query.
	Get(ctx context.Context, request GetQueryRequest) (*Query, error)

	// Gets a list of queries accessible to the user, ordered by creation time.
	// **Warning:** Calling this API concurrently 10 or more times could result
	// in throttling, service degradation, or a temporary ban.
	List(ctx context.Context, request ListQueriesRequest) (*ListQueryObjectsResponse, error)

	// Gets a list of visualizations on a query.
	ListVisualizations(ctx context.Context, request ListVisualizationsForQueryRequest) (*ListVisualizationsForQueryResponse, error)

	// Updates the fields of a query named in `update_mask`. Changing
	// `owner_user_name` requires workspace admin privileges. Changing
	// `run_as_mode` requires CAN MANAGE permission on the query, and only the
	// query owner can switch it to `OWNER`.
	Update(ctx context.Context, request UpdateQueryRequest) (*Query, error)
}

// These endpoints are used for CRUD operations on query definitions. Query
// definitions include the target SQL warehouse, query text, name, description,
// tags, parameters, and visualizations. Queries can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
//
// **Warning**: This API is deprecated. Please see the latest version of the
// Databricks SQL API. [Learn more]
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type QueriesLegacyService interface {

	// Creates a new query definition. Queries created with this endpoint belong
	// to the authenticated user making the request.
	//
	// The `data_source_id` field specifies the ID of the SQL warehouse to run
	// this query against. You can use the Data Sources API to see a complete
	// list of available SQL warehouses. Or you can copy the `data_source_id`
	// from an existing query.
	//
	// **Note**: You cannot add a visualization until you create the query.
	//
	// **Warning**: This API is deprecated. Please use :method:queries/create
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Create(ctx context.Context, request QueryPostContent) (*LegacyQuery, error)

	// Moves a query to the trash. Trashed queries immediately disappear from
	// searches and list views, and they cannot be used for alerts. The trash is
	// deleted after 30 days.
	//
	// **Warning**: This API is deprecated. Please use :method:queries/delete
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Delete(ctx context.Context, request DeleteQueriesLegacyRequest) error

	// Retrieve a query object definition along with contextual permissions
	// information about the currently authenticated user.
	//
	// **Warning**: This API is deprecated. Please use :method:queries/get
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Get(ctx context.Context, request GetQueriesLegacyRequest) (*LegacyQuery, error)

	// Gets a list of queries. Optionally, this list can be filtered by a search
	// term.
	//
	// **Warning**: Calling this API concurrently 10 or more times could result
	// in throttling, service degradation, or a temporary ban.
	//
	// **Warning**: This API is deprecated. Please use :method:queries/list
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	List(ctx context.Context, request ListQueriesLegacyRequest) (*QueryList, error)

	// Restore a query that has been moved to the trash. A restored query
	// appears in list views and searches. You can use restored queries for
	// alerts.
	//
	// **Warning**: This API is deprecated. Please see the latest version.
	// [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Restore(ctx context.Context, request RestoreQueriesLegacyRequest) error

	// Modify this query definition.
	//
	// **Note**: You cannot undo this operation.
	//
	// **Warning**: This API is deprecated. Please use :method:queries/update
	// instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Update(ctx context.Context, request QueryEditContent) (*LegacyQuery, error)
}

// A service responsible for storing and retrieving the list of queries run
// against SQL endpoints and serverless compute.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type QueryHistoryService interface {

	// List the history of queries through SQL warehouses, and serverless
	// compute.
	//
	// You can filter by user ID, warehouse ID, status, and time range. Most
	// recently started queries are returned first (up to max_results in
	// request). The pagination token returned in response can be used to list
	// subsequent query statuses.
	List(ctx context.Context, request ListQueryHistoryRequest) (*ListQueriesResponse, error)
}

// This is an evolving API that facilitates the addition and removal of
// visualizations from existing queries in the Databricks Workspace. Data
// structures can change over time.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type QueryVisualizationsService interface {

	// Adds a visualization to a query.
	Create(ctx context.Context, request CreateVisualizationRequest) (*Visualization, error)

	// Removes a visualization.
	Delete(ctx context.Context, request DeleteVisualizationRequest) error

	// Updates a visualization.
	Update(ctx context.Context, request UpdateVisualizationRequest) (*Visualization, error)
}

// This is an evolving API that facilitates the addition and removal of
// visualizations from existing queries within the Databricks Workspace. Data
// structures may change over time.
//
// **Warning**: This API is deprecated. Please see the latest version of the
// Databricks SQL API. [Learn more]
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type QueryVisualizationsLegacyService interface {

	// Creates visualization in the query.
	//
	// **Warning**: This API is deprecated. Please use
	// :method:queryvisualizations/create instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Create(ctx context.Context, request CreateQueryVisualizationsLegacyRequest) (*LegacyVisualization, error)

	// Removes a visualization from the query.
	//
	// **Warning**: This API is deprecated. Please use
	// :method:queryvisualizations/delete instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Delete(ctx context.Context, request DeleteQueryVisualizationsLegacyRequest) error

	// Updates visualization in the query.
	//
	// **Warning**: This API is deprecated. Please use
	// :method:queryvisualizations/update instead. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
	Update(ctx context.Context, request LegacyVisualization) (*LegacyVisualization, error)
}

// Redash V2 service for workspace configurations (internal)
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type RedashConfigService interface {

	// Read workspace configuration for Redash-v2.
	GetConfig(ctx context.Context) (*ClientConfig, error)
}

// The Databricks SQL Statement Execution API can be used to execute SQL
// statements on a SQL warehouse and fetch the result.
//
// **Getting started**
//
// We suggest beginning with the [Databricks SQL Statement Execution API
// tutorial].
//
// **Lakehouse//RT warehouses**
//
// You can run statements on [Lakehouse//RT] warehouses with this API. For lower
// latency and higher concurrency, use version 2 of this API, Statement
// Execution (Lakehouse//RT), instead. It works the same way, except that
// `warehouse_id` is a path parameter.
//
// **Overview of statement execution and result fetching**
//
// Statement execution begins by issuing a
// :method:statementexecution/executeStatement request with a valid SQL
// statement and warehouse ID, along with optional parameters such as the data
// catalog and output format. If no other parameters are specified, the server
// will wait for up to 10s before returning a response. If the statement has
// completed within this timespan, the response will include the result data as
// a JSON array and metadata. Otherwise, if no result is available after the 10s
// timeout expired, the response will provide the statement ID that can be used
// to poll for results by using a :method:statementexecution/getStatement
// request.
//
// You can specify whether the call should behave synchronously, asynchronously
// or start synchronously with a fallback to asynchronous execution. This is
// controlled with the `wait_timeout` and `on_wait_timeout` settings. If
// `wait_timeout` is set between 5-50 seconds (default: 10s), the call waits for
// results up to the specified timeout; when set to `0s`, the call is
// asynchronous and responds immediately with a statement ID. The
// `on_wait_timeout` setting specifies what should happen when the timeout is
// reached while the statement execution has not yet finished. This can be set
// to either `CONTINUE`, to fallback to asynchronous mode, or it can be set to
// `CANCEL`, which cancels the statement.
//
// In summary: - **Synchronous mode** (`wait_timeout=30s` and
// `on_wait_timeout=CANCEL`): The call waits up to 30 seconds; if the statement
// execution finishes within this time, the result data is returned directly in
// the response. If the execution takes longer than 30 seconds, the execution is
// canceled and the call returns with a `CANCELED` state. - **Asynchronous
// mode** (`wait_timeout=0s` and `on_wait_timeout` is ignored): The call doesn't
// wait for the statement to finish but returns directly with a statement ID.
// The status of the statement execution can be polled by issuing
// :method:statementexecution/getStatement with the statement ID. Once the
// execution has succeeded, this call also returns the result and metadata in
// the response. - **[Default] Hybrid mode** (`wait_timeout=10s` and
// `on_wait_timeout=CONTINUE`): The call waits for up to 10 seconds; if the
// statement execution finishes within this time, the result data is returned
// directly in the response. If the execution takes longer than 10 seconds, a
// statement ID is returned. The statement ID can be used to fetch status and
// results in the same way as in the asynchronous mode.
//
// Depending on the size, the result can be split into multiple chunks. If the
// statement execution is successful, the statement response contains a manifest
// and the first chunk of the result. The manifest contains schema information
// and provides metadata for each chunk in the result. Result chunks can be
// retrieved by index with :method:statementexecution/getStatementResultChunkN
// which may be called in any order and in parallel. For sequential fetching,
// each chunk, apart from the last, also contains a `next_chunk_index` and
// `next_chunk_internal_link` that point to the next chunk.
//
// A statement can be canceled with :method:statementexecution/cancelExecution.
//
// **Fetching result data: format and disposition**
//
// To specify the format of the result data, use the `format` field, which can
// be set to one of the following options: `JSON_ARRAY` (JSON), `ARROW_STREAM`
// ([Apache Arrow Columnar]), or `CSV`.
//
// There are two ways to receive statement results, controlled by the
// `disposition` setting, which can be either `INLINE` or `EXTERNAL_LINKS`:
//
// - `INLINE`: In this mode, the result data is directly included in the
// response. It's best suited for smaller results. This mode can only be used
// with the `JSON_ARRAY` format.
//
// - `EXTERNAL_LINKS`: In this mode, the response provides links that can be
// used to download the result data in chunks separately. This approach is ideal
// for larger results and offers higher throughput. This mode can be used with
// all the formats: `JSON_ARRAY`, `ARROW_STREAM`, and `CSV`.
//
// By default, the API uses `format=JSON_ARRAY` and `disposition=INLINE`.
//
// **Limits and limitations**
//
// Note: The byte limit for INLINE disposition is based on internal storage
// metrics and will not exactly match the byte count of the actual payload.
//
// - Statements with `disposition=INLINE` are limited to 25 MiB and will fail
// when this limit is exceeded. - Statements with `disposition=EXTERNAL_LINKS`
// are limited to 100 GiB. Result sets larger than this limit will be truncated.
// Truncation is indicated by the `truncated` field in the result manifest. -
// The maximum query text size is 16 MiB. - Cancelation might silently fail. A
// successful response from a cancel request indicates that the cancel request
// was successfully received and sent to the processing engine. However, an
// outstanding statement might have already completed execution when the cancel
// request arrives. Polling for status until a terminal state is reached is a
// reliable way to determine the final state. - Wait timeouts are approximate,
// occur server-side, and cannot account for things such as caller delays and
// network latency from caller to service. - To keep a statement alive, poll for
// its status at least once every 15 minutes. Regular polling guarantees the
// statement stays alive. - To stop a statement you no longer need, cancel it
// explicitly with :method:statementexecution/cancelExecution. Stopping polling
// does not cancel a statement. Explicit cancellation is the only way to end
// statement execution on demand. - The results are only available for one hour
// after success; polling does not extend this. - The SQL Execution API must be
// used for the entire lifecycle of the statement. For example, you cannot use
// the Jobs API to execute the command, and then the SQL Execution API to cancel
// it.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Apache Arrow Columnar]: https://arrow.apache.org/overview/
// [Databricks SQL Statement Execution API tutorial]: https://docs.databricks.com/sql/api/sql-execution-tutorial.html
// [Lakehouse//RT]: https://docs.databricks.com/compute/sql-warehouse/real-time.html
type StatementExecutionService interface {

	// Requests that an executing statement be canceled. Callers must poll for
	// status to see the terminal state. Cancel response is empty; receiving
	// response indicates successful receipt.
	CancelExecution(ctx context.Context, request CancelExecutionRequest) error

	// Execute a SQL statement and optionally await its results for a specified
	// time.
	//
	// **Use case: small result sets with INLINE + JSON_ARRAY**
	//
	// For flows that generate small and predictable result sets (<= 25 MiB),
	// `INLINE` responses of `JSON_ARRAY` result data are typically the simplest
	// way to execute and fetch result data.
	//
	// **Use case: large result sets with EXTERNAL_LINKS**
	//
	// Using `EXTERNAL_LINKS` to fetch result data allows you to fetch large
	// result sets efficiently. The main differences from using `INLINE`
	// disposition are that the result data is accessed with short-lived
	// cloud-storage URLs, and that there are 3 supported formats: `JSON_ARRAY`,
	// `ARROW_STREAM` and `CSV` compared to only `JSON_ARRAY` with `INLINE`.
	//
	// **External-link URLs**
	//
	// External links point to data stored within your workspace's internal
	// cloud storage. The URLs are valid for only a short period, <= 15 minutes.
	// Alongside each `external_link` is an expiration field indicating the time
	// at which the URL is no longer valid. In `EXTERNAL_LINKS` mode, chunks can
	// be resolved and fetched multiple times and in parallel.
	//
	// ----
	//
	// ### **Warning: Databricks strongly recommends that you protect the URLs
	// that are returned by the `EXTERNAL_LINKS` disposition.**
	//
	// When you use the `EXTERNAL_LINKS` disposition, a short-lived
	// cloud-storage URL is generated to download the results. The URL contains
	// temporary access credentials, so protect it and do not set an
	// `Authorization` header in the download request.
	//
	// The `EXTERNAL_LINKS` disposition can be disabled upon request by creating
	// a [support case].
	//
	// See also [Security best practices].
	//
	// ----
	//
	// StatementResponse contains `statement_id` and `status`; other fields
	// might be absent or present depending on context. If the SQL warehouse
	// fails to execute the provided statement, a 200 response is returned with
	// `status.state` set to `FAILED` (in contrast to a failure when accepting
	// the request, which results in a non-200 response). Details of the error
	// can be found at `status.error` in case of execution failures.
	//
	// [Security best practices]: https://docs.databricks.com/sql/admin/sql-execution-tutorial.html#security-best-practices
	// [support case]: https://docs.databricks.com/resources/support.html
	ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error)

	// This request can be used to poll for the statement's status.
	// StatementResponse contains `statement_id` and `status`; other fields
	// might be absent or present depending on context. When the `status.state`
	// field is `SUCCEEDED` it will also return the result manifest and the
	// first chunk of the result data. When the statement is in the terminal
	// states `CANCELED`, `CLOSED` or `FAILED`, it returns HTTP 200 with the
	// state set. After at least 12 hours in terminal state, the statement is
	// removed from the warehouse and further calls will receive an HTTP 404
	// response.
	//
	// **NOTE** This call currently might take up to 5 seconds to get the latest
	// status and result.
	GetStatement(ctx context.Context, request GetStatementRequest) (*StatementResponse, error)

	// After the statement execution has `SUCCEEDED`, this request can be used
	// to fetch any chunk by index. Whereas the first chunk with `chunk_index=0`
	// is typically fetched with :method:statementexecution/executeStatement or
	// :method:statementexecution/getStatement, this request can be used to
	// fetch subsequent chunks. The response structure is identical to the
	// nested `result` element described in the
	// :method:statementexecution/getStatement request, and similarly includes
	// the `next_chunk_index` and `next_chunk_internal_link` fields for simple
	// iteration through the result set. Depending on `disposition`, the
	// response returns chunks of data either inline, or as links.
	GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error)
}

// A SQL warehouse is a compute resource that lets you run SQL commands on data
// objects within Databricks SQL. Compute resources are infrastructure resources
// that provide processing capabilities in the cloud.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WarehousesService interface {

	// Creates a new SQL warehouse.
	Create(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error)

	// Creates a new default warehouse override for a user. Users can create
	// their own override. Admins can create overrides for any user.
	CreateDefaultWarehouseOverride(ctx context.Context, request CreateDefaultWarehouseOverrideRequest) (*DefaultWarehouseOverride, error)

	// Deletes a SQL warehouse.
	Delete(ctx context.Context, request DeleteWarehouseRequest) error

	// Deletes the default warehouse override for a user. Users can delete their
	// own override. Admins can delete overrides for any user. After deletion,
	// the workspace default warehouse will be used.
	DeleteDefaultWarehouseOverride(ctx context.Context, request DeleteDefaultWarehouseOverrideRequest) error

	// Updates the configuration for a SQL warehouse.
	Edit(ctx context.Context, request EditWarehouseRequest) error

	// Gets the information for a single SQL warehouse.
	Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error)

	// Returns the default warehouse override for a user. Users can fetch their
	// own override. Admins can fetch overrides for any user. If no override
	// exists, the UI will fallback to the workspace default warehouse.
	GetDefaultWarehouseOverride(ctx context.Context, request GetDefaultWarehouseOverrideRequest) (*DefaultWarehouseOverride, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetWarehousePermissionLevelsRequest) (*GetWarehousePermissionLevelsResponse, error)

	// Gets the permissions of a SQL warehouse. SQL warehouses can inherit
	// permissions from their root object.
	GetPermissions(ctx context.Context, request GetWarehousePermissionsRequest) (*WarehousePermissions, error)

	// Gets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error)

	// Lists all SQL warehouses that a user has access to.
	List(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error)

	// Lists all default warehouse overrides in the workspace. Only workspace
	// administrators can list all overrides.
	ListDefaultWarehouseOverrides(ctx context.Context, request ListDefaultWarehouseOverridesRequest) (*ListDefaultWarehouseOverridesResponse, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error)

	// Sets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error

	// Starts a SQL warehouse.
	Start(ctx context.Context, request StartRequest) error

	// Stops a SQL warehouse.
	Stop(ctx context.Context, request StopRequest) error

	// Updates an existing default warehouse override for a user. Users can
	// update their own override. Admins can update overrides for any user.
	UpdateDefaultWarehouseOverride(ctx context.Context, request UpdateDefaultWarehouseOverrideRequest) (*DefaultWarehouseOverride, error)

	// Updates the permissions on a SQL warehouse. SQL warehouses can inherit
	// permissions from their root object.
	UpdatePermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error)
}
