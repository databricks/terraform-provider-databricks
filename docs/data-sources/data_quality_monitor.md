---
subcategory: "Data Quality Monitoring"
---
# databricks_data_quality_monitor Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/dataquality)

This data source can be used to fetch a data quality monitor.

For the `table` `object_type`, the caller must either:
1. be an owner of the table's parent catalog
2. have **USE_CATALOG** on the table's parent catalog and be an owner of the table's parent schema.
3. have the following permissions:
   - **USE_CATALOG** on the table's parent catalog
   - **USE_SCHEMA** on the table's parent schema
   - **SELECT** privilege on the table.

-> **Note** This data source can only be used with a workspace-level provider!


## Example Usage
Getting a data quality monitor by Unity Catalog object type (currently supports `schema` and `table`) and object id:

```hcl
data "databricks_schema" "this" {
  name = "my_catalog.my_schema"
}
data "databricks_data_quality_monitor" "this" {
  object_type = "schema"
  object_id = data.databricks_schema.this.schema_info.schema_id
}
```


## Arguments
The following arguments are supported:
* `object_id` (string, required) - The UUID of the request object. It is `schema_id` for `schema`, and `table_id` for `table`.
  
  Find the `schema_id` from either:
  1. The [schema_id](https://docs.databricks.com/api/workspace/schemas/get#schema_id) of the `Schemas` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `schema` > go to the `Details` tab > the `Schema ID` field.
  
  Find the `table_id` from either:
  1. The [table_id](https://docs.databricks.com/api/workspace/tables/get#table_id) of the `Tables` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `table` > go to the `Details` tab > the `Table ID` field
* `object_type` (string, required) - The type of the monitored object. Can be one of the following: `schema` or `table`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `anomaly_detection_config` (AnomalyDetectionConfig) - Anomaly detection settings. Required when `object_type` is `schema`. When reading a `table` monitor, this holds
  the settings of the parent schema's monitor if the schema has one
* `data_profiling_config` (DataProfilingConfig) - Data profiling settings. Required when `object_type` is `table`. Exactly one of `inference_log`, `time_series`
  or `snapshot` must be set
* `object_id` (string) - The UUID of the request object. It is `schema_id` for `schema`, and `table_id` for `table`.
  
  Find the `schema_id` from either:
  1. The [schema_id](https://docs.databricks.com/api/workspace/schemas/get#schema_id) of the `Schemas` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `schema` > go to the `Details` tab > the `Schema ID` field.
  
  Find the `table_id` from either:
  1. The [table_id](https://docs.databricks.com/api/workspace/tables/get#table_id) of the `Tables` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `table` > go to the `Details` tab > the `Table ID` field
* `object_type` (string) - The type of the monitored object. Can be one of the following: `schema` or `table`

### AnomalyDetectionConfig
* `excluded_table_full_names` (list of string) - List of fully qualified table names to exclude from anomaly detection. Format: `catalog.schema.table_name`
* `notification_settings` (NotificationSettings) - Who to alert when anomaly detection finds unhealthy tables in this schema. Applicable to
  `schema` monitors only. This is the same schema alert you manage from **Manage alerts** in the
  Data Quality Monitoring UI, so a change made in either place shows in both.
  
  Set the recipients in `on_unhealthy_table`; `on_failure` is not supported here. Only
  `destination_ids` and `destination_user_ids` are supported.
  
  On update, an update mask of `*`, `anomaly_detection_config`, or
  `anomaly_detection_config.notification_settings` replaces the alert's recipients with the ones
  in the request; omitting the field, or leaving both recipient lists empty, removes the alert.
  Other masks leave the alert unchanged

### CronSchedule
* `pause_status` (string) - Read only field that indicates whether the schedule is paused or not. Possible values are: `CRON_SCHEDULE_PAUSE_STATUS_PAUSED`, `CRON_SCHEDULE_PAUSE_STATUS_UNPAUSED`
* `quartz_cron_expression` (string) - The expression that determines when to run the monitor. See [examples](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html)
* `timezone_id` (string) - The Java timezone ID in which the cron expression is evaluated. See
  [Java TimeZone](https://docs.oracle.com/javase/8/docs/api/java/util/TimeZone.html) for the supported IDs

### DataProfilingConfig
* `assets_dir` (string) - The absolute workspace path of the directory where the monitor's assets, such as
  the dashboard, are stored. The UI and the Python APIs normally prepopulate a default user location
* `baseline_table_name` (string) - The fully qualified name of a baseline table. Format: `catalog.schema.table_name`.
  Drift metrics compare the monitored table against the baseline data, so the two tables must have the same
  schema
* `custom_metrics` (list of DataProfilingCustomMetric) - Custom metrics to compute in addition to the built-in ones. See
  [custom metrics](https://docs.databricks.com/aws/en/data-governance/unity-catalog/data-quality-monitoring/data-profiling/custom-metrics)
* `dashboard_id` (string) - The ID of the dashboard that visualizes the computed metrics. It is empty while
  the monitor is pending, when the dashboard was skipped with `skip_builtin_dashboard`, or when the caller is in
  a different workspace from the one where the monitor was created
* `drift_metrics_table_name` (string) - Table that stores drift metrics data. Format: `catalog.schema.table_name`
* `effective_warehouse_id` (string) - The warehouse for dashboard creation
* `inference_log` (InferenceLogConfig) - Profiles an inference log table, which stores the requests and predictions of a model. Metrics are
  computed per time window and include model quality metrics
* `latest_monitor_failure_message` (string) - The latest error message for a monitor failure
* `monitor_version` (integer) - Represents the current monitor configuration version in use. The version will be represented in a
  numeric fashion (1,2,3...). The field has flexibility to take on negative values, which can indicate corrupted
  monitor_version numbers
* `monitored_table_name` (string) - The fully qualified name of the monitored table. Format: `catalog.schema.table_name`
* `notification_settings` (NotificationSettings) - Who to notify about this monitor's refreshes. Only `on_failure` with
  `email_addresses` is supported here
* `output_schema_id` (string) - The UUID of the Unity Catalog schema where the profile and drift metrics tables
  are created
* `profile_metrics_table_name` (string) - Table that stores profile metrics data. Format: `catalog.schema.table_name`
* `schedule` (CronSchedule) - The schedule on which the monitor refreshes its metrics. Refreshes can also be
  triggered on demand through the Refresh resource
* `skip_builtin_dashboard` (boolean) - Whether to skip creating a default dashboard summarizing data quality metrics
* `slicing_exprs` (list of string) - List of column expressions to slice data with for targeted analysis. The data is grouped by
  each expression independently, resulting in a separate slice for each predicate and its
  complements. For example `slicing_exprs=["col_1", "col_2 > 10"]` will generate the following
  slices: two slices for `col_2 > 10` (True and False), and one slice per unique value in
  `col_1`. For high-cardinality columns, only the top 100 unique values by frequency will
  generate slices
* `snapshot` (SnapshotConfig) - Profiles the full contents of the table. Metrics are computed over all rows at each refresh
* `status` (string) - The status of the monitor. `DATA_PROFILING_STATUS_PENDING` means the monitor's
  assets are being created or updated. `DATA_PROFILING_STATUS_ACTIVE` means the monitor is ready and can be
  refreshed. `DATA_PROFILING_STATUS_FAILED` and `DATA_PROFILING_STATUS_ERROR` mean the monitor is not
  operational, and `latest_monitor_failure_message` describes the cause when one was recorded.
  `DATA_PROFILING_STATUS_DELETE_PENDING` means the monitor is being deleted. Possible values are: `DATA_PROFILING_STATUS_ACTIVE`, `DATA_PROFILING_STATUS_DELETE_PENDING`, `DATA_PROFILING_STATUS_ERROR`, `DATA_PROFILING_STATUS_FAILED`, `DATA_PROFILING_STATUS_PENDING`
* `time_series` (TimeSeriesConfig) - Profiles a time series table. Metrics are computed per time window, as defined by `timestamp_column`
  and `granularities`
* `warehouse_id` (string) - The ID of the SQL warehouse used to create the dashboard. If not specified, the first running warehouse the
  caller can access is used

### DataProfilingCustomMetric
* `definition` (string) - Jinja template for a SQL expression that specifies how to compute the metric. See [create metric definition](https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition)
* `input_columns` (list of string) - A list of column names in the input table the metric should be computed for.
  Can use `":table"` to indicate that the metric needs information from multiple columns
* `name` (string) - Name of the metric in the output tables
* `output_data_type` (string) - Specifies the Spark SQL data type of the metric value, either as a plain type name or as a JSON-encoded
  StructField
* `type` (string) - The type of the custom metric. `DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE` metrics depend only on the
  columns of the input table, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED` metrics depend only on previously
  computed aggregate metrics, and `DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT` metrics depend on previously computed
  aggregate or derived metrics. Possible values are: `DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT`

### InferenceLogConfig
* `granularities` (list of string) - List of granularities to use when aggregating data into time windows based on their timestamp
* `label_column` (string) - Column containing the ground truth for the model's predictions
* `model_id_column` (string) - Column containing the ID of the model that produced each prediction. A slice is created automatically for
  each distinct model ID
* `prediction_column` (string) - Column containing the model's predicted values
* `problem_type` (string) - The type of problem the model solves. It determines which model quality metrics are computed. Possible values are: `INFERENCE_PROBLEM_TYPE_CLASSIFICATION`, `INFERENCE_PROBLEM_TYPE_REGRESSION`
* `timestamp_column` (string) - Column containing the timestamp of each inference request. Its type must be TIMESTAMP or convertible to a
  timestamp

### NotificationDestination
* `destination_ids` (list of string) - The notification-destination IDs (Slack / Teams / webhook / Asana / ServiceNow / AWS
  EventBridge) to send the notification to. The destinations must already exist in the
  workspace; a workspace admin creates them. A maximum of 25 destination IDs is supported.
  Only supported on schema monitors, via `anomaly_detection_config`; rejected on table monitors
* `destination_user_ids` (list of integer) - The IDs of the workspace users to notify by email. A maximum of 25 user IDs is supported.
  Only supported on schema monitors, via `anomaly_detection_config`; rejected on table monitors
* `email_addresses` (list of string) - The list of email addresses to send the notification to. A maximum of 5 email addresses is supported.
  Only supported on table monitors, via `data_profiling_config`; rejected on schema monitors

### NotificationSettings
* `on_failure` (NotificationDestination) - Recipients to notify when a refresh of the monitor fails or times out. Only supported on table
  monitors, via `data_profiling_config`, and only `email_addresses` is accepted there
* `on_unhealthy_table` (NotificationDestination) - Recipients to notify when anomaly detection finds unhealthy tables in the schema. Each
  recipient gets one notification per important unhealthy table. Only supported on schema
  monitors, via `anomaly_detection_config`

### SnapshotConfig

### TimeSeriesConfig
* `granularities` (list of string) - List of granularities to use when aggregating data into time windows based on their timestamp
* `timestamp_column` (string) - Column containing the timestamp of each row. Its type must be TIMESTAMP or convertible to a timestamp