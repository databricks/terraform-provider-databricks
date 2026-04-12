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
* `source` (DataSource, required) - The data source of the feature
* `description` (string, optional) - The description of the feature
* `entities` (list of EntityColumn, optional) - The entity columns for the feature, used as aggregation keys and for query-time lookup
* `filter_condition` (string, optional, deprecated) - Deprecated: Use DeltaTableSource.filter_condition or KafkaSource.filter_condition instead. Kept for backwards compatibility.
  The filter condition applied to the source data before aggregation
* `inputs` (list of string, optional, deprecated) - Deprecated: Use AggregationFunction.inputs instead. Kept for backwards compatibility.
  The input columns from which the feature is computed
* `lineage_context` (LineageContext, optional) - Lineage context information for this feature.
  WARNING: This field is primarily intended for internal use by Databricks systems and
  is automatically populated when features are created through Databricks notebooks or jobs.
  Users should not manually set this field as incorrect values may lead to inaccurate lineage tracking or unexpected behavior.
  This field will be set by feature-engineering client and should be left unset by SDK and terraform users
* `time_window` (TimeWindow, optional, deprecated) - Deprecated: Use Function.aggregation_function.time_window instead. Kept for backwards compatibility.
  The time window in which the feature is computed
* `timeseries_column` (TimeseriesColumn, optional) - Column recording time, used for point-in-time joins, backfills, and aggregations
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### AggregationFunction
* `approx_count_distinct` (ApproxCountDistinctFunction, optional)
* `approx_percentile` (ApproxPercentileFunction, optional)
* `avg` (AvgFunction, optional)
* `count_function` (CountFunction, optional)
* `first` (FirstFunction, optional)
* `last` (LastFunction, optional)
* `max` (MaxFunction, optional)
* `min` (MinFunction, optional)
* `stddev_pop` (StddevPopFunction, optional)
* `stddev_samp` (StddevSampFunction, optional)
* `sum` (SumFunction, optional)
* `time_window` (TimeWindow, optional) - The time window over which the aggregation is computed
* `var_pop` (VarPopFunction, optional)
* `var_samp` (VarSampFunction, optional)

### ApproxCountDistinctFunction
* `input` (string, required) - The input column from which the approximate count of distinct values is computed
* `relative_sd` (number, optional) - The maximum relative standard deviation allowed (default defined by Spark)

### ApproxPercentileFunction
* `input` (string, required) - The input column from which the approximate percentile is computed
* `percentile` (number, required) - The percentile value to compute (between 0 and 1)
* `accuracy` (integer, optional) - The accuracy parameter (higher is more accurate but slower)

### AvgFunction
* `input` (string, required) - The input column from which the average is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### ColumnIdentifier
* `variant_expr_path` (string, required) - String representation of the column name using dot-prefixed path notation. For nested fields, the leaf value is what will be present in materialized tables
  and expected to match at query time. For example, the leaf node of value.trip_details.location_details.pickup_zip is pickup_zip

### ColumnSelection
* `column` (string, required) - Column name from source to select as the feature value

### ContinuousWindow
* `window_duration` (string, required) - The duration of the continuous window (must be positive)
* `offset` (string, optional) - The offset of the continuous window (must be non-positive)

### CountFunction
* `input` (string, required) - The input column from which the count is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### DataSource
* `delta_table_source` (DeltaTableSource, optional)
* `kafka_source` (KafkaSource, optional)

### DeltaTableSource
* `full_name` (string, required) - The full three-part (catalog, schema, table) name of the Delta table
* `dataframe_schema` (string, optional) - Schema of the resulting dataframe after transformations, in Spark StructType JSON format (from df.schema.json()).
  Required if transformation_sql is specified.
  Example: {"type":"struct","fields":[{"name":"col_a","type":"integer","nullable":true,"metadata":{}},{"name":"col_c","type":"integer","nullable":true,"metadata":{}}]}
* `entity_columns` (list of string, optional, deprecated) - Deprecated: Use Feature.entity instead. Kept for backwards compatibility.
  The entity columns of the Delta table
* `filter_condition` (string, optional) - Single WHERE clause to filter delta table before applying transformations. Will be row-wise evaluated, so should only include conditionals and projections
* `timeseries_column` (string, optional, deprecated) - Deprecated: Use Feature.timeseries_column instead. Kept for backwards compatibility.
  The timeseries column of the Delta table
* `transformation_sql` (string, optional) - A single SQL SELECT expression applied after filter_condition.
  Should contains all the columns needed (eg. "SELECT *, col_a + col_b AS col_c FROM x.y.z WHERE col_a > 0" would have `transformation_sql` "*, col_a + col_b AS col_c")
  If transformation_sql is not provided, all columns of the delta table are present in the DataSource dataframe

### EntityColumn
* `name` (string, required) - The name of the entity column. For Kafka sources, use dot-prefixed path notation to reference
  fields within the key or value schema (e.g., "value.user_id", "key.partition_key"). For nested
  fields, the leaf node name (e.g., "user_id" from "value.trip_details.user_id") is what will
  be present in materialized tables and expected to match at query time.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:user_id") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### FirstFunction
* `input` (string, required) - The input column from which the first value is returned

### Function
* `aggregation_function` (AggregationFunction, optional) - An aggregation function applied over a time window
* `column_selection` (ColumnSelection, optional) - Selects the latest value of a single column in a data source
* `extra_parameters` (list of FunctionExtraParameter, optional, deprecated) - Deprecated: Use the function oneof with AggregationFunction instead. Kept for backwards compatibility.
  Extra parameters for parameterized functions
* `function_type` (string, optional, deprecated) - Deprecated: Use the function oneof with AggregationFunction instead. Kept for backwards compatibility.
  The type of the function. Possible values are: `APPROX_COUNT_DISTINCT`, `APPROX_PERCENTILE`, `AVG`, `COUNT`, `FIRST`, `LAST`, `MAX`, `MIN`, `STDDEV_POP`, `STDDEV_SAMP`, `SUM`, `VAR_POP`, `VAR_SAMP`

### FunctionExtraParameter
* `key` (string, required) - The name of the parameter
* `value` (string, required) - The value of the parameter

### JobContext
* `job_id` (integer, optional) - The job ID where this API invoked
* `job_run_id` (integer, optional) - The job run ID where this API was invoked

### KafkaSource
* `name` (string, required) - Name of the Kafka source, used to identify it. This is used to look up the corresponding KafkaConfig object. Can be distinct from topic name
* `entity_column_identifiers` (list of ColumnIdentifier, optional, deprecated) - Deprecated: Use Feature.entity instead. Kept for backwards compatibility.
  The entity column identifiers of the Kafka source
* `filter_condition` (string, optional) - The filter condition applied to the source data before aggregation
* `timeseries_column_identifier` (ColumnIdentifier, optional, deprecated) - Deprecated: Use Feature.timeseries_column instead. Kept for backwards compatibility.
  The timeseries column identifier of the Kafka source

### LastFunction
* `input` (string, required) - The input column from which the last value is returned

### LineageContext
* `job_context` (JobContext, optional) - Job context information including job ID and run ID
* `notebook_id` (integer, optional) - The notebook ID where this API was invoked

### MaxFunction
* `input` (string, required) - The input column from which the maximum is computed

### MinFunction
* `input` (string, required) - The input column from which the minimum is computed

### SlidingWindow
* `slide_duration` (string, required) - The slide duration (interval by which windows advance, must be positive and less than duration)
* `window_duration` (string, required) - The duration of the sliding window

### StddevPopFunction
* `input` (string, required) - The input column from which the population standard deviation is computed. For Kafka sources,
  use dot-prefixed path notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### StddevSampFunction
* `input` (string, required) - The input column from which the sample standard deviation is computed

### SumFunction
* `input` (string, required) - The input column from which the sum is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### TimeWindow
* `continuous` (ContinuousWindow, optional)
* `sliding` (SlidingWindow, optional)
* `tumbling` (TumblingWindow, optional)

### TimeseriesColumn
* `name` (string, required) - The name of the timeseries column. For Kafka sources, use dot-prefixed path notation to
  reference fields within the key or value schema (e.g., "value.event_timestamp"). For nested
  fields, the leaf node name (e.g., "event_timestamp" from "value.event_details.event_timestamp")
  is what will be present in materialized tables and expected to match at query time.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:event_timestamp") is supported for
  backwards compatibility but is deprecated; migrate to dot notation

### TumblingWindow
* `window_duration` (string, required) - The duration of each tumbling window (non-overlapping, fixed-duration windows)

### VarPopFunction
* `input` (string, required) - The input column from which the population variance is computed

### VarSampFunction
* `input` (string, required) - The input column from which the sample variance is computed



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