---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_feature Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `full_name` (string, required) - The full three-part name (catalog, schema, name) of the feature
* `function` (Function, required) - The function by which the feature is computed
* `inputs` (list of string, required) - The input columns from which the feature is computed
* `source` (DataSource, required) - The data source of the feature
* `time_window` (TimeWindow, required) - The time window in which the feature is computed
* `description` (string, optional) - The description of the feature
* `filter_condition` (string, optional) - The filter condition applied to the source data before aggregation
* `lineage_context` (LineageContext, optional) - WARNING: This field is primarily intended for internal use by Databricks systems and
  is automatically populated when features are created through Databricks notebooks or jobs.
  Users should not manually set this field as incorrect values may lead to inaccurate lineage tracking or unexpected behavior.
  This field will be set by feature-engineering client and should be left unset by SDK and terraform users

### ColumnIdentifier
* `variant_expr_path` (string, required) - String representation of the column name or variant expression path. For nested fields, the leaf value is what will be present in materialized tables
  and expected to match at query time. For example, the leaf node of value:trip_details.location_details.pickup_zip is pickup_zip

### ContinuousWindow
* `window_duration` (string, required) - The duration of the continuous window (must be positive)
* `offset` (string, optional) - The offset of the continuous window (must be non-positive)

### DataSource
* `delta_table_source` (DeltaTableSource, optional)
* `kafka_source` (KafkaSource, optional)

### DeltaTableSource
* `entity_columns` (list of string, required) - The entity columns of the Delta table
* `full_name` (string, required) - The full three-part (catalog, schema, table) name of the Delta table
* `timeseries_column` (string, required) - The timeseries column of the Delta table

### Function
* `function_type` (string, required) - The type of the function. Possible values are: `APPROX_COUNT_DISTINCT`, `APPROX_PERCENTILE`, `AVG`, `COUNT`, `FIRST`, `LAST`, `MAX`, `MIN`, `STDDEV_POP`, `STDDEV_SAMP`, `SUM`, `VAR_POP`, `VAR_SAMP`
* `extra_parameters` (list of FunctionExtraParameter, optional) - Extra parameters for parameterized functions

### FunctionExtraParameter
* `key` (string, required) - The name of the parameter
* `value` (string, required) - The value of the parameter

### JobContext
* `job_id` (integer, optional) - The job ID where this API invoked
* `job_run_id` (integer, optional) - The job run ID where this API was invoked

### KafkaSource
* `entity_column_identifiers` (list of ColumnIdentifier, required) - The entity column identifiers of the Kafka source
* `name` (string, required) - Name of the Kafka source, used to identify it. This is used to look up the corresponding KafkaConfig object. Can be distinct from topic name
* `timeseries_column_identifier` (ColumnIdentifier, required) - The timeseries column identifier of the Kafka source

### LineageContext
* `job_context` (JobContext, optional) - Job context information including job ID and run ID
* `notebook_id` (integer, optional) - The notebook ID where this API was invoked

### SlidingWindow
* `slide_duration` (string, required) - The slide duration (interval by which windows advance, must be positive and less than duration)
* `window_duration` (string, required) - The duration of the sliding window

### TimeWindow
* `continuous` (ContinuousWindow, optional)
* `sliding` (SlidingWindow, optional)
* `tumbling` (TumblingWindow, optional)

### TumblingWindow
* `window_duration` (string, required) - The duration of each tumbling window (non-overlapping, fixed-duration windows)



## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "full_name"
  to = databricks_feature_engineering_feature.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_feature_engineering_feature.this "full_name"
```