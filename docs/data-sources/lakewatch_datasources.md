---
subcategory: "Lakewatch"
---
# databricks_lakewatch_datasources Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)



## Example Usage


## Arguments
The following arguments are supported:
* `catalog_name` (string, required) - Name of the Unity Catalog catalog to list datasources from. Required; each request lists the
  datasources of a single catalog
* `filter` (string, optional) - Filter expression that restricts the results to datasources with a given name, for example
  `datasource_name = "my_source"`. Only a single exact-match comparison on `datasource_name`
  is supported, and it matches the datasource's current Unity Catalog name, which reads report
  in `effective_datasource_name`. Filtered results are returned in a single page
* `page_size` (integer, optional) - Maximum number of results to return in one page. This is a best-effort upper
  bound, not a guarantee: the server may return fewer items than requested (even
  zero) while more results still exist. If unset or 0, the server picks a default
* `schema_name` (string, optional) - Name of a Unity Catalog schema in `catalog_name`. Optional; if set, only the datasources in
  this schema are returned, and if omitted, the datasources in all schemas of the catalog are
  returned
* `view` (string, optional) - Which fields to return for each datasource. Defaults to `DATASOURCE_VIEW_FULL`.
  `DATASOURCE_VIEW_BASIC` returns only the identity and Unity Catalog location fields. Possible values are: `DATASOURCE_VIEW_BASIC`, `DATASOURCE_VIEW_FULL`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `datasources`. It is a list of resources, each with the following attributes:
* `annotations` (object) - Free-form key-value annotations on the datasource, managed by the client
* `autoloader` (AutoLoader)
* `catalog_name` (string) - Name of the Unity Catalog catalog that contains the datasource. Required on create and
  cannot be changed afterward. Reads return the value as supplied; the catalog's current name
  is reported in `effective_catalog_name`
* `compute_mode` (string) - Compute mode for the datasource's ingestion job. If unset, the `default_compute_mode` of
  the catalog's datasource configuration applies, or `DS_COMPUTE_PERFORMANCE_OPTIMIZED` if
  the catalog has none. Possible values are: `DS_COMPUTE_PERFORMANCE_OPTIMIZED`, `DS_COMPUTE_STANDARD`
* `compute_workspace_id` (string) - ID of the workspace whose compute runs this datasource's ingestion job and
  owns its associated resources, such as the storage the ingested data lands
  in. Server-assigned
* `create_time` (string)
* `created_by` (Principal)
* `datasource_id` (string) - Unity Catalog securable ID (a UUID) of the datasource. It never changes, including when the
  datasource is renamed in Unity Catalog, and is the final segment of `name`
* `datasource_name` (string) - Name of the datasource in Unity Catalog, unique within its schema. Required on create and
  cannot be changed afterward. Reads return the value as supplied; the datasource's current
  Unity Catalog name, which reflects renames, is reported in `effective_datasource_name`
* `description` (string) - Free-form description of the datasource
* `effective_catalog_name` (string) - Current name of the Unity Catalog catalog that contains the datasource, which reflects
  catalog renames without a change to `catalog_name`. Populated on get, create, and update
  responses; list responses omit it
* `effective_datasource_name` (string) - Current UC leaf name of the datasource (server-owned). Tracks UC renames so
  callers can observe the live name without having to update the client-owned
  `datasource_name`; unset until resolved
* `effective_run_as_principal` (string) - The identity the datasource runs as: the resolved `run_as_principal`, or the requester's
  identity if `run_as_principal` was omitted on create or cleared on update. Given as a user
  email or service principal application ID, or as the numeric principal ID if the identity
  can no longer be resolved. Populated on get, create, and update responses; list responses
  omit it
* `effective_schema_name` (string) - Current name of the Unity Catalog schema that contains the datasource, which reflects
  schema renames without a change to `schema_name`. Populated on get, create, and update
  responses; list responses omit it
* `etag` (string) - Opaque entity tag for optimistic concurrency control. Returned by get, create, and update
  responses and by list responses in the `DATASOURCE_VIEW_FULL` view. To make an update
  conditional on the datasource not having changed since it was read, pass this value as the
  update request's `etag`. Changes to `status` do not change the etag
* `existing_table` (ExistingTable)
* `name` (string) - Resource name of the datasource, in the format `datasources/{datasource_id}`. Unlike the
  datasource's Unity Catalog location (`catalog_name`, `schema_name`, and `datasource_name`),
  the name does not change when the datasource or its catalog or schema is renamed in Unity
  Catalog
* `owner` (Principal) - Unity Catalog owner of the datasource, assigned by the server. Create responses omit it
* `parsers` (list of ParserRef) - Parsers that schematize the ingested data. Each record is matched against the parsers in
  order and written to the parsed table of the first parser that it matches. Existing-table
  and Unity Catalog connection datasources require at least one parser; an Auto Loader
  datasource with no parsers writes only its raw table
* `run_as_principal` (string) - The user or service principal that the datasource runs as, given as a user email or
  service principal application ID. When it is supplied on create or update, the requester
  must be authorized to run the datasource as that principal, or the request is rejected.
  When it is omitted on create, the datasource runs as the requester. On update, a mask that
  includes `run_as_principal` without a value clears it, so the datasource runs as the
  requester; a mask that omits it keeps the current identity. Reads return the value as
  supplied; the identity the datasource runs as is reported in `effective_run_as_principal`
* `schedule` (Schedule) - Schedule on which the datasource's ingestion job runs. A datasource cannot be paused, so a
  request that sets `schedule.enabled` to `false` is rejected. Continuous schedules are not
  supported for Unity Catalog connection datasources
* `schema_name` (string) - Name of the Unity Catalog schema that contains the datasource. Required on create and
  cannot be changed afterward. Reads return the value as supplied; the schema's current name
  is reported in `effective_schema_name`
* `status` (DatasourceStatus) - Runtime status of the jobs and pipelines that run this datasource
* `unity_catalog_connection` (UnityCatalogConnection)
* `update_time` (string)
* `updated_by` (Principal)
* `version` (integer) - Version number of the datasource, incremented on each update. Informational only; use
  `etag` for optimistic concurrency control
* `warehouse_id` (string) - ID of the SQL warehouse that runs the checks Lakewatch performs to validate a create or update
  request, such as source and destination access checks. Required and non-blank on every create
  and update request; there is no default warehouse. The requester must be able to use the
  warehouse. Used only for those checks: it is not stored, is not returned on reads, and does
  not affect the compute that runs the ingestion job

### AutoLoader
* `aws_secret_scope` (string) - Name of the workspace secret scope that holds the AWS access key ID and secret access key,
  under the keys `awsAccessKeyId` and `awsSecretAccessKey`
* `effective_include_existing_files` (boolean) - Whether ingestion started with the files already present at `location`. Populated on get,
  create, and update responses
* `effective_raw_table` (TableName) - The table that the raw data is written to: `raw_table` in the datasource's catalog and
  schema, under their current Unity Catalog names. Populated on get, create, and update
  responses
* `envelope` (Envelope) - How to unwrap a transport envelope around the source records before parsing. Supported only
  for the JSON and JSONL formats. If unset, each source record is treated as one event
* `format` (string) - File format of the source files. Required. Possible values are: `AUTOLOADER_FORMAT_CSV`, `AUTOLOADER_FORMAT_JSON`, `AUTOLOADER_FORMAT_JSONL`, `AUTOLOADER_FORMAT_PARQUET`, `AUTOLOADER_FORMAT_XML`
* `format_options` (object) - Format-specific parsing options, passed to Auto Loader. Supported for the CSV, JSON, JSONL,
  and XML formats, each of which accepts a fixed set of options. The XML format requires the
  `rowTag` option
* `include_existing_files` (boolean) - Whether to ingest the files already present at `location` when ingestion
  first starts. Set true to ingest the existing files and then every new file.
  Set false, or leave unset, to ingest only files that arrive after the
  datasource starts. Defaults to false. Cannot be changed after creation.
  Reads return the value as supplied; the setting in effect is reported in
  `effective_include_existing_files`
* `location` (string) - URI of the source files: a path in a Unity Catalog volume or external location, or a
  cloud storage URL such as an Amazon S3 path. Required, including when new files are
  discovered through `queue_url`. Cannot be changed after creation
* `override_partition_columns` (list of string) - Hive-style partition columns to derive from the source file paths. If unset, no partition
  columns are derived
* `queue_url` (string) - URL of a customer-managed Amazon SQS queue that receives file-arrival notifications. Set it
  to discover new files through the queue. If unset, new files are discovered through the
  file events of the source's Unity Catalog external location when it has them configured,
  and by directory listing otherwise. Cannot be changed after creation
* `raw_table` (TableName) - Table that stores the raw data before parsers schematize it. Every Auto Loader datasource
  writes its raw rows to this table before any parser reads them. The table must be in the
  datasource's catalog and schema; an omitted `catalog_name` or `schema_name` defaults to
  them. Cannot be changed after creation. Reads return the value as supplied; the table the
  data is written to is reported in `effective_raw_table`
* `service_credentials` (string) - Name of the Unity Catalog service credential to use

### CloudWatchLogsEnvelope
* `header_retention` (boolean) - Whether to keep the envelope's `owner`, `logGroup`, and `logStream` values in
  the `data.lw_header` field of each unwrapped row. Must be set explicitly to
  `true` or `false`

### DatasourceStatus
* `ingestion_job` (ResourceStatus) - Status of the datasource's ingestion job
* `ingestion_pipeline` (ResourceStatus) - Status of the managed ingestion pipeline. Populated only when `unity_catalog_connection`
  is set, with `pipeline_id` identifying the pipeline

### Envelope
* `cloud_watch_logs` (CloudWatchLogsEnvelope)
* `json_array` (JsonArrayEnvelope)

### ExistingTable
* `catalog` (string) - Name of the catalog the existing table lives in. Optional on create and
  defaults to the datasource's own catalog. Cannot be changed. Reads return
  the value as supplied; the catalog in use is reported in `effective_catalog`
* `effective_catalog` (string) - Current Unity Catalog catalog name of the source table, resolved from its
  stable table ID, so a rename of the source table is reported here without
  changing `catalog`
* `effective_name` (string) - Current Unity Catalog name of the source table, resolved from its stable
  table ID
* `effective_schema` (string) - Current Unity Catalog schema name of the source table, resolved from its
  stable table ID
* `include_existing_rows` (boolean) - Whether to ingest the rows already present in the source table when
  ingestion first starts. Set true to ingest the existing rows and then every
  new commit. Set false, or leave unset, to ingest only rows written after the
  datasource starts. Defaults to false, and must be true for a Delta Sharing
  source table. Takes effect only at first start and cannot be changed. Reads
  return the value as supplied
* `name` (string) - Name of the source table or view. Required on create and cannot be changed.
  Reads return the value as supplied; the table's current name is reported in
  `effective_name`
* `schema` (string) - Name of the schema the existing table lives in. Optional, but required when
  `catalog` is set; supplying `catalog` without `schema` is rejected. Defaults
  to the datasource's own schema. Cannot be changed. Reads return the value as
  supplied; the schema in use is reported in `effective_schema`

### JsonArrayEnvelope
* `path` (string) - Name of the root-level key that holds the array of events, for example
  `Records`. Nested paths are not supported

### ParserRef
* `effective_target_table` (TableName) - The table this parser's data is written to: `target_table` when supplied, otherwise the
  table derived from the datasource's catalog and schema and the parser name. Fixed once
  resolved, so a later parser rename does not move it. Populated on get, create, and update
  responses with the resolved table
* `id` (string) - Identifier of the parser. Stable across renames, unlike a parser name
* `target_table` (TableName) - Table that this parser's data is written to. The table must be in the datasource's catalog
  and schema. If unset, the table name is derived from the parser name. Reads return the value
  as supplied; the table the data is written to is reported in `effective_target_table`
* `version` (string) - Version of the parser to use, given as an explicit version number

### Principal
* `display_name` (string) - Display name (user or SP name)
* `email` (string) - Email (regular users only)
* `id` (string) - User / service-principal ID
* `principal_type` (string) - Whether this principal is a regular user, service principal, or system principal. Possible values are: `REGULAR_USER`, `SERVICE_PRINCIPAL`, `SYSTEM_USER`

### ResourceStatus
* `create_time` (string)
* `job_id` (string) - Databricks Jobs API job ID. Unset for a pure ingestion pipeline
* `last_complete_time` (string) - When the most recent run completed. Pairs with `run_state = SUCCEEDED`
* `last_error_message` (string) - Most recent error message, if any. Pairs with `run_state = FAILED`
* `last_start_time` (string) - When the most recent run started. Pairs with `run_state = RUNNING`
* `last_status_change_time` (string)
* `lifecycle_state` (string) - Scheduling lifecycle of the resource's underlying job/pipeline: whether it
  is unscheduled, scheduled, stopped, or deleted. Distinct from `run_state`,
  which reports the outcome of the most recent run. Possible values are: `DELETED`, `SCHEDULED`, `STOPPED`, `UNSCHEDULED`
* `pipeline_id` (string) - DLT ingestion-pipeline ID. Set only when this status represents the
  datasource's ingestion pipeline (unset for job statuses)
* `run_state` (string) - Possible values are: `CANCELED`, `FAILED`, `RUNNING`, `SUCCEEDED`

### Schedule
* `at_least_every` (string) - Duration string defining the minimum cadence. May contain components
  for days, hours, minutes, seconds (e.g. `5d`, `5d3h`, `3d2h5m12s`,
  `12m16s`)
* `compute_group` (string) - `dedicated`, `inherited`, `automatic`, or a custom group name. Reads return
  the value as supplied. A rule whose schedule omits it runs on `dedicated`
* `continuous` (boolean) - Run continuously (sentinel — must be `true` when set in this branch)
* `cron_expression` (string) - Quartz cron expression
* `enabled` (boolean) - Whether the job is scheduled to run. Defaults to true

### TableName
* `catalog_name` (string) - Name of the parent Unity Catalog catalog
* `schema_name` (string) - Name of the parent Unity Catalog schema
* `table_name` (string) - Unity Catalog table name

### UnityCatalogConnection
* `configuration` (object) - Free-form key-value configuration for the managed ingestion pipeline
* `connection_name` (string) - Name of the Unity Catalog connection to ingest through
* `connector` (string) - Connector type of the connection, such as `SLACK_AUDIT_LOGS`. Resolved by the server from
  `connection_name` through Unity Catalog; not accepted as input
* `effective_start_time` (string) - The start time ingestion uses: the supplied `start_time`, or the creation time when it was
  omitted. Always populated
* `start_time` (string) - Earliest source event time to ingest. Leave unset to start at creation time, with no
  history backfilled. Rejected if in the future, beyond the connector's history window, or
  for a connector without start-time support. Takes effect only at first start and cannot
  be changed. Reads return the value as supplied; the start time ingestion uses is reported
  in `effective_start_time`