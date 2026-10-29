---
subcategory: "Data Quality Monitoring"
---
# databricks_data_quality_monitor Resource
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/dataquality)

This resource allows you to set up data quality monitoring checks for Unity Catalog objects, currently schema and table. 

For the `table` `object_type`, you must either:
1. be an owner of the table's parent catalog, have **USE_SCHEMA** on the table's parent schema, and have **SELECT** access on the table
2. have **USE_CATALOG** on the table's parent catalog, be an owner of the table's parent schema, and have **SELECT** access on the table.
3. have the following permissions:
   - **USE_CATALOG** on the table's parent catalog
   - **USE_SCHEMA** on the table's parent schema
   - be an owner of the table.

For a `schema` monitor, `anomaly_detection_config.notification_settings` manages the schema's alert, the same alert shown under **Manage alerts** in the Data Quality Monitoring UI. Because the provider updates the whole `anomaly_detection_config`, a monitor without `notification_settings` deletes an alert that was set up in the UI. To stop alerts, remove the `notification_settings` block. Don't set it to an empty value: the API accepts it and deletes the alert, but the apply then fails because the API returns no `notification_settings`.

-> **Note** This resource can only be used with a workspace-level provider!


## Example Usage
```hcl
resource "databricks_schema" "this" {
  catalog_name = "my_catalog"
  name = "my_schema"
}
resource "databricks_data_quality_monitor" "this" {
  object_type = "schema"
  object_id = databricks_schema.this.schema_id
  anomaly_detection_config = {}
}
```

To alert workspace users or notification destinations when anomaly detection finds unhealthy tables in the schema, set `notification_settings`:

```hcl
resource "databricks_data_quality_monitor" "with_alerts" {
  object_type = "schema"
  object_id = databricks_schema.this.schema_id
  anomaly_detection_config = {
    notification_settings = {
      on_unhealthy_table = {
        destination_user_ids = ["<user-id>"]
        destination_ids = ["<notification-destination-id>"]
      }
    }
  }
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
* `anomaly_detection_config` (AnomalyDetectionConfig, optional) - Anomaly Detection Configuration, applicable to `schema` object types
* `data_profiling_config` (DataProfilingConfig, optional) - Data Profiling Configuration, applicable to `table` object types. Exactly one `Analysis Configuration`
  must be present
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### AnomalyDetectionConfig
* `excluded_table_full_names` (list of string, optional) - List of fully qualified table names to exclude from anomaly detection
* `notification_settings` (NotificationSettings, optional) - Who to alert when anomaly detection finds unhealthy tables in this schema. Applicable to
  `schema` monitors only. This is the same schema alert you manage from **Manage alerts** in the
  Data Quality Monitoring UI, so a change made in either place shows in both.
  
  Set the recipients in `on_unhealthy_table`; `on_failure` is not supported here. Only
  `destination_ids` and `destination_user_ids` are supported.
  
  On update, an update mask of `*`, `anomaly_detection_config`, or
  `anomaly_detection_config.notification_settings` replaces the alert's recipients with the ones
  in the request; omitting the field, or leaving both recipient lists empty, removes the alert.
  Other masks leave the alert unchanged

### CronSchedule
* `quartz_cron_expression` (string, required) - The expression that determines when to run the monitor. See [examples](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html)
* `timezone_id` (string, required) - A Java timezone id. The schedule for a job will be resolved with respect to this timezone.
  See `Java TimeZone <http://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html>`_ for details.
  The timezone id (e.g., ``America/Los_Angeles``) in which to evaluate the quartz expression

### DataProfilingConfig
* `output_schema_id` (string, required) - ID of the schema where output tables are created
* `assets_dir` (string, optional) - Field for specifying the absolute path to a custom directory to store data-monitoring
  assets. Normally prepopulated to a default user location via UI and Python APIs
* `baseline_table_name` (string, optional) - Baseline table name.
  Baseline data is used to compute drift from the data in the monitored `table_name`.
  The baseline table and the monitored table shall have the same schema
* `custom_metrics` (list of DataProfilingCustomMetric, optional) - Custom metrics
* `inference_log` (InferenceLogConfig, optional) - `Analysis Configuration` for monitoring inference log tables
* `notification_settings` (NotificationSettings, optional) - Field for specifying notification settings
* `schedule` (CronSchedule, optional) - The cron schedule
* `skip_builtin_dashboard` (boolean, optional) - Whether to skip creating a default dashboard summarizing data quality metrics
* `slicing_exprs` (list of string, optional) - List of column expressions to slice data with for targeted analysis. The data is grouped by
  each expression independently, resulting in a separate slice for each predicate and its
  complements. For example `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the following
  slices: two slices for `col_2 > 10` (True and False), and one slice per unique value in
  `col1`. For high-cardinality columns, only the top 100 unique values by frequency will
  generate slices
* `snapshot` (SnapshotConfig, optional) - `Analysis Configuration` for monitoring snapshot tables
* `time_series` (TimeSeriesConfig, optional) - `Analysis Configuration` for monitoring time series tables
* `warehouse_id` (string, optional) - Optional argument to specify the warehouse for dashboard creation. If not specified, the first running
  warehouse will be used

### DataProfilingCustomMetric
* `definition` (string, required) - Jinja template for a SQL expression that specifies how to compute the metric. See [create metric definition](https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition)
* `input_columns` (list of string, required) - A list of column names in the input table the metric should be computed for.
  Can use ``":table"`` to indicate that the metric needs information from multiple columns
* `name` (string, required) - Name of the metric in the output tables
* `output_data_type` (string, required) - The output type of the custom metric
* `type` (string, required) - The type of the custom metric. Possible values are: `DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT`

### InferenceLogConfig
* `granularities` (list of string, required) - List of granularities to use when aggregating data into time windows based on their timestamp
* `model_id_column` (string, required) - Column for the model identifier
* `prediction_column` (string, required) - Column for the prediction
* `problem_type` (string, required) - Problem type the model aims to solve. Possible values are: `INFERENCE_PROBLEM_TYPE_CLASSIFICATION`, `INFERENCE_PROBLEM_TYPE_REGRESSION`
* `timestamp_column` (string, required) - Column for the timestamp
* `label_column` (string, optional) - Column for the label

### NotificationDestination
* `destination_ids` (list of string, optional) - The notification-destination IDs (Slack / Teams / webhook) to send the notification to. The
  destinations must already exist in the workspace; a workspace admin creates them. A maximum of
  25 destination IDs is supported.
  Only supported on schema monitors, via `anomaly_detection_config`; rejected on table monitors
* `destination_user_ids` (list of integer, optional) - The IDs of the workspace users to notify by email. A maximum of 25 user IDs is supported.
  Only supported on schema monitors, via `anomaly_detection_config`; rejected on table monitors
* `email_addresses` (list of string, optional) - The list of email addresses to send the notification to. A maximum of 5 email addresses is supported

### NotificationSettings
* `on_failure` (NotificationDestination, optional) - Destinations to send notifications on failure/timeout
* `on_unhealthy_table` (NotificationDestination, optional) - Recipients to notify when anomaly detection finds unhealthy tables in the schema. Each
  recipient gets one notification per important unhealthy table. Only supported on schema
  monitors, via `anomaly_detection_config`

### TimeSeriesConfig
* `granularities` (list of string, required) - List of granularities to use when aggregating data into time windows based on their timestamp
* `timestamp_column` (string, required) - Column for the timestamp



## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "object_type,object_id"
  to = databricks_data_quality_monitor.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_data_quality_monitor.this "object_type,object_id"
```