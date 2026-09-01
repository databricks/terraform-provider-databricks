---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_feature Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `full_name` (string, required) - The full three-part name (catalog, schema, name) of the feature. This is the
  feature's resource identifier; the catalog_name, schema_name, and name fields
  below are OUTPUT_ONLY decomposed views of this value
* `function` (Function, required) - The function by which the feature is computed
* `source` (DataSource, required) - The data source of the feature
* `description` (string, optional) - The description of the feature
* `entities` (list of EntityColumn, optional) - The entity columns for the feature, used as aggregation keys and for query-time lookup
* `lineage_context` (LineageContext, optional) - Lineage context information for this feature.
  WARNING: This field is primarily intended for internal use by Databricks systems and
  is automatically populated when features are created through Databricks notebooks or jobs.
  Users should not manually set this field as incorrect values may lead to inaccurate lineage tracking or unexpected behavior.
  This field will be set by feature-engineering client and should be left unset by SDK and terraform users
* `timeseries_column` (TimeseriesColumn, optional) - Column recording time, used for point-in-time joins, backfills, and aggregations
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### AggregationFunction
* `approx_count_distinct` (ApproxCountDistinctFunction, optional)
* `approx_percentile` (ApproxPercentileFunction, optional)
* `avg` (AvgFunction, optional)
* `count_function` (CountFunction, optional)
* `first` (FirstFunction, optional)
* `first_distinct` (FirstDistinctFunction, optional)
* `first_n` (FirstNFunction, optional)
* `last` (LastFunction, optional)
* `last_distinct` (LastDistinctFunction, optional)
* `last_n` (LastNFunction, optional)
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
  Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### ColumnSelection
* `column` (string, required) - Column name from source to select as the feature value

### CountFunction
* `input` (string, required) - The input column from which the count is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### CustomUdf
* `function_path` (string, required) - Fully qualified 3-part Unity Catalog path of the function to apply
* `input_bindings` (list of InputBinding, optional) - Binds each UC function parameter to a source column.
  May be empty for zero-argument functions (e.g. a timestamp generator)

### DataSource
* `delta_table_source` (DeltaTableSource, optional) - A Delta table data source
* `kafka_source` (KafkaSource, optional) - A Kafka stream data source
* `lateness` (SourceLateness, optional) - Completeness timing for this Feature's use of the source. This configuration is part of the
  Feature definition; it does not modify the underlying table or stream
* `request_source` (RequestSource, optional) - A request-time data source
* `stream_source` (StreamSource, optional) - A Stream data source

### DeltaTableSource
* `full_name` (string, required) - The full three-part (catalog, schema, table) name of the Delta table
* `dataframe_schema` (string, optional) - Schema of the resulting dataframe after transformations, in Spark StructType JSON format (from df.schema.json()).
  Required if transformation_sql is specified.
  Example: {"type":"struct","fields":[{"name":"col_a","type":"integer","nullable":true,"metadata":{}},{"name":"col_c","type":"integer","nullable":true,"metadata":{}}]}
* `filter_condition` (string, optional) - Single WHERE clause to filter delta table before applying transformations. Will be row-wise evaluated, so should only include conditionals and projections
* `transformation_sql` (string, optional) - A single SQL SELECT expression applied after filter_condition.
  Should contains all the columns needed (eg. "SELECT *, col_a + col_b AS col_c FROM x.y.z WHERE col_a > 0" would have `transformation_sql` "*, col_a + col_b AS col_c")
  If transformation_sql is not provided, all columns of the delta table are present in the DataSource dataframe

### EntityColumn
* `name` (string, required) - The name of the entity column. For Kafka sources, use dot-prefixed path notation to reference
  fields within the key or value schema (e.g., "value.user_id", "key.partition_key"). For nested
  fields, the leaf node name (e.g., "user_id" from "value.trip_details.user_id") is what will
  be present in materialized tables and expected to match at query time.
  Colon-prefixed notation (e.g., "value:user_id") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### FieldDefinition
* `data_type` (string, required) - The scalar data type of the field. Possible values are: `BINARY`, `BOOLEAN`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INTEGER`, `LONG`, `SHORT`, `STRING`, `TIMESTAMP`
* `name` (string, required) - The name of the field

### FirstDistinctFunction
* `input` (string, required) - The input column from which the first N distinct values are returned
* `n` (integer, required) - The number of distinct values to return

### FirstFunction
* `input` (string, required) - The input column from which the first value is returned

### FirstNFunction
* `input` (string, required) - The input column from which the first N values are returned
* `n` (integer, required) - The number of values to return

### FlatSchema
* `fields` (list of FieldDefinition, required) - The list of fields in this schema

### Function
* `aggregation_function` (AggregationFunction, optional) - An aggregation function applied over a time window
* `column_selection` (ColumnSelection, optional) - Selects the latest value of a single column in a data source
* `custom_udf` (CustomUdf, optional) - Applies a registered Unity Catalog function row-wise to source columns

### InputBinding
* `column` (string, required) - Source column whose value is passed for this parameter at execution time
* `parameter` (string, required) - Name of the UC function parameter

### JobContext
* `job_id` (integer, optional) - The job ID where this API invoked
* `job_run_id` (integer, optional) - The job run ID where this API was invoked

### KafkaSource
* `name` (string, required) - Name of the Kafka source, used to identify it. This is used to look up the corresponding KafkaConfig object. Can be distinct from topic name
* `filter_condition` (string, optional) - The filter condition applied to the source data before aggregation

### LastDistinctFunction
* `input` (string, required) - The input column from which the last N distinct values are returned
* `n` (integer, required) - The number of distinct values to return

### LastFunction
* `input` (string, required) - The input column from which the last value is returned

### LastNFunction
* `input` (string, required) - The input column from which the last N values are returned
* `n` (integer, required) - The number of values to return

### LineageContext
* `job_context` (JobContext, optional) - Job context information including job ID and run ID
* `notebook_id` (integer, optional) - The notebook ID where this API was invoked

### MaxFunction
* `input` (string, required) - The input column from which the maximum is computed

### MinFunction
* `input` (string, required) - The input column from which the minimum is computed

### RequestSource
* `flat_schema` (FlatSchema, optional) - A flat schema with scalar-typed fields only

### RollingWindow
* `delay` (string, optional) - Non-negative analytic lag that evaluates the window this far in the past. Use this for timing
  variations unrelated to source lateness, such as a 30-day count as of one week ago. If unset,
  the analytic lag is zero. It composes with source.lateness when both are set
* `window_duration` (string, optional) - The duration of the rolling window. Must be positive when set; absent means lifetime
  (aggregate over the entity's entire history)

### SawtoothWindow
* `delay` (string, optional) - Delay is not currently supported for Sawtooth windows
* `window_duration` (string, optional) - The duration of the window. Must be positive and span more than two days when set, so that both
  the batch (N-1 day) and stale-path (N-2 day) partial aggregates are well defined. The duration
  need not be a whole number of days (e.g. 3 days 15 minutes is allowed). Absent means lifetime
  (aggregate over the entity's entire history)

### SlidingWindow
* `slide_duration` (string, required) - The slide duration (interval by which windows advance, must be positive and less than duration)
* `delay` (string, optional) - Non-negative analytic lag that evaluates the window this far in the past. Use this for timing
  variations unrelated to source lateness, such as a 30-day count as of one week ago. If unset,
  the analytic lag is zero. It composes with source.lateness when both are set
* `offset` (string, optional) - Non-negative phase shift from the default midnight UTC alignment. For example, offset=22h on
  a 24h slide produces boundaries at 22:00 UTC (17:00 New York in standard time) instead of
  midnight UTC. If unset, the offset is zero. Must be shorter than slide_duration (and therefore
  window_duration)
* `window_duration` (string, optional) - The duration of the sliding window. Must be positive when set; absent means lifetime
  (aggregate over the entity's entire history)

### SourceLateness
* `settling_delay` (string, optional) - Non-negative time to wait after a window ends before treating its source data as complete.
  Training shifts the eligible evaluation time backwards by this duration so it does not join
  data that would still have been settling online. Materialization waits for the duration to
  elapse before publishing the window. If unset, source data is considered settled immediately

### StddevPopFunction
* `input` (string, required) - The input column from which the population standard deviation is computed. For Kafka sources,
  use dot-prefixed path notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### StddevSampFunction
* `input` (string, required) - The input column from which the sample standard deviation is computed

### StreamSource
* `full_name` (string, required) - Three-part full name of the Stream (catalog.schema.stream)
* `dataframe_schema` (string, optional) - Schema of the resulting dataframe after transformations, in Spark StructType
  JSON format (from df.schema.json()).
  Any subsequent functions operate against this dataframe
* `filter_condition` (string, optional) - The filter condition applied to the source data before aggregation
* `transformation_sql` (string, optional) - The pipeline runs these SQL statements immediately after conversion into
  the schema specified on the Stream object

### SumFunction
* `input` (string, required) - The input column from which the sum is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### TimeWindow
* `rolling` (RollingWindow, optional)
* `sawtooth` (SawtoothWindow, optional) - A sawtooth window served via the hybrid batch + streaming path
* `sliding` (SlidingWindow, optional)
* `start_time` (string, optional) - Earliest event-time boundary at which the Feature may emit an output. This gates outputs, not
  the historical inputs read by a window. For example, a 365-day window with
  start_time=2026-01-01 begins emitting partial-window values on that date instead of waiting
  for 365 days of data; a lifetime window produces no output before start_time. If unset,
  tumbling and fixed-duration sliding windows first emit at an offset-aligned boundary after a
  full window can be formed. If unset, lifetime sliding windows and rolling windows emit as soon as
  eligible source data exists
* `tumbling` (TumblingWindow, optional)

### TimeseriesColumn
* `name` (string, required) - The name of the timeseries column. For Kafka sources, use dot-prefixed path notation to
  reference fields within the key or value schema (e.g., "value.event_timestamp"). For nested
  fields, the leaf node name (e.g., "event_timestamp" from "value.event_details.event_timestamp")
  is what will be present in materialized tables and expected to match at query time.
  Colon-prefixed notation (e.g., "value:event_timestamp") is supported for
  backwards compatibility but is deprecated; migrate to dot notation

### TumblingWindow
* `window_duration` (string, required) - The duration of each tumbling window (non-overlapping, fixed-duration windows)
* `delay` (string, optional) - Non-negative analytic lag that evaluates the window this far in the past. Use this for timing
  variations unrelated to source lateness, such as a 30-day count as of one week ago. If unset,
  the analytic lag is zero. It composes with source.lateness when both are set
* `offset` (string, optional) - Non-negative phase shift from the default midnight UTC alignment. For example, offset=22h on
  a 24h window produces boundaries at 22:00 UTC (17:00 New York in standard time) instead of
  midnight UTC. If unset, the offset is zero. Must be shorter than window_duration

### VarPopFunction
* `input` (string, required) - The input column from which the population variance is computed

### VarSampFunction
* `input` (string, required) - The input column from which the sample variance is computed

## Attributes
In addition to the above arguments, the following attributes are exported:
* `catalog_name` (string) - Name of parent catalog
* `created_at` (string) - Time at which this feature was created
* `created_by` (string) - Username of the feature creator
* `name` (string) - Name of the feature, extracted from the full three-part name (catalog.schema.name)
* `schema_name` (string) - Name of parent schema relative to its parent catalog

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