---
subcategory: "Data Quality Monitoring"
---
# databricks_data_quality_monitors Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to fetch the list of data quality monitors.

For the `table` `object_type`, the caller must either:
1. be an owner of the table's parent catalog
2. have **USE_CATALOG** on the table's parent catalog and be an owner of the table's parent schema.
3. have the following permissions:
    - **USE_CATALOG** on the table's parent catalog
    - **USE_SCHEMA** on the table's parent schema
    - **SELECT** privilege on the table.

-> **Note** This data source can only be used with a workspace-level provider!

## Example Usage
Getting a list of all data quality monitors:

```hcl
data "databricks_data_quality_monitors" "all" {
}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional)


## Attributes
This data source exports a single attribute, `monitors`. It is a list of resources, each with the following attributes:
* `anomaly_detection_config` (AnomalyDetectionConfig) - Anomaly Detection Configuration, applicable to `schema` object types
* `data_profiling_config` (DataProfilingConfig) - Data Profiling Configuration, applicable to `table` object types. Exactly one `Analysis Configuration`
  must be present
* `object_id` (string) - The UUID of the request object. It is `schema_id` for `schema`, and `table_id` for `table`.
  
  Find the `schema_id` from either:
  1. The [schema_id](https://docs.databricks.com/api/workspace/schemas/get#schema_id) of the `Schemas` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `schema` > go to the `Details` tab > the `Schema ID` field.
  
  Find the `table_id` from either:
  1. The [table_id](https://docs.databricks.com/api/workspace/tables/get#table_id) of the `Tables` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `table` > go to the `Details` tab > the `Table ID` field
* `object_type` (string) - The type of the monitored object. Can be one of the following: `schema` or `table`

### AnomalyDetectionConfig

### CronSchedule
* `pause_status` (string) - Read only field that indicates whether the schedule is paused or not. Possible values are: `CRON_SCHEDULE_PAUSE_STATUS_PAUSED`, `CRON_SCHEDULE_PAUSE_STATUS_UNPAUSED`
* `quartz_cron_expression` (string) - The expression that determines when to run the monitor. See [examples](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html)
* `timezone_id` (string) - A Java timezone id. The schedule for a job will be resolved with respect to this timezone.
  See `Java TimeZone <http://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html>`_ for details.
  The timezone id (e.g., ``America/Los_Angeles``) in which to evaluate the quartz expression

### DataProfilingConfig
* `assets_dir` (string) - Field for specifying the absolute path to a custom directory to store data-monitoring
  assets. Normally prepopulated to a default user location via UI and Python APIs
* `baseline_table_name` (string) - Baseline table name.
  Baseline data is used to compute drift from the data in the monitored `table_name`.
  The baseline table and the monitored table shall have the same schema
* `custom_metrics` (list of DataProfilingCustomMetric) - Custom metrics
* `dashboard_id` (string) - Id of dashboard that visualizes the computed metrics.
  This can be empty if the monitor is in PENDING state
* `drift_metrics_table_name` (string) - Table that stores drift metrics data. Format: `catalog.schema.table_name`
* `effective_warehouse_id` (string) - The warehouse for dashboard creation
* `inference_log` (InferenceLogConfig) - `Analysis Configuration` for monitoring inference log tables
* `latest_monitor_failure_message` (string) - The latest error message for a monitor failure
* `monitor_version` (integer) - Represents the current monitor configuration version in use. The version will be represented in a
  numeric fashion (1,2,3...). The field has flexibility to take on negative values, which can indicate corrupted
  monitor_version numbers
* `monitored_table_name` (string) - Unity Catalog table to monitor. Format: `catalog.schema.table_name`
* `notification_settings` (NotificationSettings) - Field for specifying notification settings
* `output_schema_id` (string) - ID of the schema where output tables are created
* `profile_metrics_table_name` (string) - Table that stores profile metrics data. Format: `catalog.schema.table_name`
* `schedule` (CronSchedule) - The cron schedule
* `skip_builtin_dashboard` (boolean) - Whether to skip creating a default dashboard summarizing data quality metrics
* `slicing_exprs` (list of string) - List of column expressions to slice data with for targeted analysis. The data is grouped by
  each expression independently, resulting in a separate slice for each predicate and its
  complements. For example `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the following
  slices: two slices for `col_2 > 10` (True and False), and one slice per unique value in
  `col1`. For high-cardinality columns, only the top 100 unique values by frequency will
  generate slices
* `snapshot` (SnapshotConfig) - `Analysis Configuration` for monitoring snapshot tables
* `status` (string) - The data profiling monitor status. Possible values are: `DATA_PROFILING_STATUS_ACTIVE`, `DATA_PROFILING_STATUS_DELETE_PENDING`, `DATA_PROFILING_STATUS_ERROR`, `DATA_PROFILING_STATUS_FAILED`, `DATA_PROFILING_STATUS_PENDING`
* `time_series` (TimeSeriesConfig) - `Analysis Configuration` for monitoring time series tables
* `warehouse_id` (string) - Optional argument to specify the warehouse for dashboard creation. If not specified, the first running
  warehouse will be used

### DataProfilingCustomMetric
* `definition` (string) - Jinja template for a SQL expression that specifies how to compute the metric. See [create metric definition](https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition)
* `input_columns` (list of string) - A list of column names in the input table the metric should be computed for.
  Can use ``":table"`` to indicate that the metric needs information from multiple columns
* `name` (string) - Name of the metric in the output tables
* `output_data_type` (string) - The output type of the custom metric
* `type` (string) - The type of the custom metric. Possible values are: `DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT`

### InferenceLogConfig
* `granularities` (list of string) - List of granularities to use when aggregating data into time windows based on their timestamp
* `label_column` (string) - Column for the label
* `model_id_column` (string) - Column for the model identifier
* `prediction_column` (string) - Column for the prediction
* `problem_type` (string) - Problem type the model aims to solve. Possible values are: `INFERENCE_PROBLEM_TYPE_CLASSIFICATION`, `INFERENCE_PROBLEM_TYPE_REGRESSION`
* `timestamp_column` (string) - Column for the timestamp

### NotificationDestination
* `email_addresses` (list of string) - The list of email addresses to send the notification to. A maximum of 5 email addresses is supported

### NotificationSettings
* `on_failure` (NotificationDestination) - Destinations to send notifications on failure/timeout

### SnapshotConfig

### TimeSeriesConfig
* `granularities` (list of string) - List of granularities to use when aggregating data into time windows based on their timestamp
* `timestamp_column` (string) - Column for the timestamp