---
subcategory: "Unity Catalog"
---
# databricks_quality_monitor Resource

This resource allows you to manage [Lakehouse Monitors](https://docs.databricks.com/en/lakehouse-monitoring/index.html) in Databricks. 

A `databricks_quality_monitor` is attached to a [databricks_sql_table](sql_table.md) and can be of type timeseries, snapshot or inference. 

## Plugin Framework Migration
The quality monitor resource has been migrated from sdkv2 to plugin framework。 If you encounter any problem with this resource and suspect it is due to the migration, you can fallback to sdkv2 by setting the environment variable in the following way `export USE_SDK_V2_RESOURCES="databricks_quality_monitor"`.

## Example Usage

```hcl
resource "databricks_catalog" "sandbox" {
  name    = "sandbox"
  comment = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_schema" "things" {
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this database is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_sql_table" "myTestTable" {
  catalog_name       = "main"
  schema_name        = databricks_schema.things.name
  name               = "bar"
  table_type         = "MANAGED"
  data_source_format = "DELTA"

  column {
    name     = "timestamp"
    type     = "int"
  }
}

resource "databricks_quality_monitor" "testTimeseriesMonitor" {
  table_name         = "${databricks_catalog.sandbox.name}.${databricks_schema.things.name}.${databricks_sql_table.myTestTable.name}"
  assets_dir         = "/Shared/provider-test/databricks_quality_monitoring/${databricks_sql_table.myTestTable.name}"
  output_schema_name = "${databricks_catalog.sandbox.name}.${databricks_schema.things.name}"
  time_series {
    granularities = ["1 hour"]
    timestamp_col = "timestamp"
  }
}
```

### Inference Monitor

```hcl
resource "databricks_quality_monitor" "testMonitorInference" {
  table_name         = "${databricks_catalog.sandbox.name}.${databricks_schema.things.name}.${databricks_table.myTestTable.name}"
  assets_dir         = "/Shared/provider-test/databricks_quality_monitoring/${databricks_table.myTestTable.name}"
  output_schema_name = "${databricks_catalog.sandbox.name}.${databricks_schema.things.name}"
  inference_log {
    granularities  = ["1 hour"]
    timestamp_col  = "timestamp"
    prediction_col = "prediction"
    model_id_col   = "model_id"
    problem_type   = "PROBLEM_TYPE_REGRESSION"
  }
}
```
### Snapshot Monitor
```hcl
resource "databricks_quality_monitor" "testMonitorInference" {
  table_name         = "${databricks_catalog.sandbox.name}.${databricks_schema.things.name}.${databricks_table.myTestTable.name}"
  assets_dir         = "/Shared/provider-test/databricks_quality_monitoring/${databricks_table.myTestTable.name}"
  output_schema_name = "${databricks_catalog.sandbox.name}.${databricks_schema.things.name}"
  snapshot {}
}
```

## Argument Reference

The following arguments are supported:

* `table_name` - (Required) - The full name of the table to attach the monitor too. Its of the format {catalog}.{schema}.{tableName}
* `assets_dir` - (Required) - The directory to store the monitoring assets (Eg. Dashboard and Metric Tables)
* `output_schema_name` - (Required) - Schema where output metric tables are created
* `baseline_table_name` - Name of the baseline table from which drift metrics are computed from.Columns in the monitored table should also be present in the baseline
table.
* `custom_metrics` - Custom metrics to compute on the monitored table. These can be aggregate metrics, derived metrics (from already computed aggregate metrics), or drift metrics (comparing metrics across time windows).
    * `definition` - [create metric definition](https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition)
    * `input_columns` - Columns on the monitored table to apply the custom metrics to.
    * `name` - Name of the custom metric.
    * `output_data_type` - The output type of the custom metric.
    * `type` - The type of the custom metric.
* `data_classification_config` - The data classification config for the monitor
* `inference_log` - Configuration for the inference log monitor
    * `granularities` -  List of granularities to use when aggregating data into time windows based on their timestamp.
    * `label_col` - Column of the model label
    * `model_id_col` - Column of the model id or version
    * `prediction_col` - Column of the model prediction
    * `prediction_proba_col` - Column of the model prediction probabilities
    * `problem_type` - Problem type the model aims to solve. Either `PROBLEM_TYPE_CLASSIFICATION` or `PROBLEM_TYPE_REGRESSION`
    * `timestamp_col` - Column of the timestamp of predictions
* `snapshot` - Configuration for monitoring snapshot tables.
* `time_series` - Configuration for monitoring timeseries tables.
    * `granularities` -  List of granularities to use when aggregating data into time windows based on their timestamp.
    * `timestamp_col` - Column of the timestamp of predictions
* `notifications` - The notification settings for the monitor.  The following optional blocks are supported, each consisting of the single string array field with name `email_addresses` containing a list of emails to notify:
    * `on_failure` - who to send notifications to on monitor failure.
    * `on_new_classification_tag_detected` - Who to send notifications to when new data classification tags are detected.
* `schedule` - The schedule for automatically updating and refreshing metric tables.  This block consists of following fields:
    * `quartz_cron_expression` - string expression that determines when to run the monitor. See [Quartz documentation](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) for examples.
    * `timezone_id` - string with timezone id (e.g., `PST`) in which to evaluate the Quartz expression.
* `skip_builtin_dashboard` - Whether to skip creating a default dashboard summarizing data quality metrics.
* `slicing_exprs` - List of column expressions to slice data with for targeted analysis. The data is grouped by each expression independently, resulting in a separate slice for each predicate and its complements. For high-cardinality columns, only the top 100 unique values by frequency will generate slices.
* `warehouse_id` - Optional argument to specify the warehouse for dashboard creation. If not specified, the first running warehouse will be used.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  ID of this monitor is the same as the full table name of the format `{catalog}.{schema_name}.{table_name}`
* `monitor_version` - The version of the monitor config (e.g. 1,2,3). If negative, the monitor may be corrupted
* `drift_metrics_table_name` - The full name of the drift metrics table. Format: __catalog_name__.__schema_name__.__table_name__.
* `profile_metrics_table_name` - The full name of the profile metrics table. Format: __catalog_name__.__schema_name__.__table_name__.
* `status` - Status of the Monitor 
* `dashboard_id` - The ID of the generated dashboard.

## Related Resources

The following resources are often used in the same context:

* [databricks_catalog](catalog.md)
* [databricks_schema](schema.md)
* [databricks_sql_table](sql_table.md)
