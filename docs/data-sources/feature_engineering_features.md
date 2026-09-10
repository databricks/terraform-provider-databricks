---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_features Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `catalog_name` (string, required) - Name of parent catalog for features of interest
* `schema_name` (string, required) - Name of parent schema relative to its parent catalog
* `page_size` (integer, optional) - The maximum number of results to return
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `features`. It is a list of resources, each with the following attributes:
* `catalog_name` (string) - Name of parent catalog
* `created_at` (string) - Time at which this feature was created
* `created_by` (string) - Username of the feature creator
* `description` (string) - The description of the feature
* `entities` (list of EntityColumn) - The entity columns for the feature, used as aggregation keys and for query-time lookup
* `full_name` (string) - The full three-part name (catalog, schema, name) of the feature. This is the
  feature's resource identifier; the catalog_name, schema_name, and name fields
  below are OUTPUT_ONLY decomposed views of this value
* `function` (Function) - The function by which the feature is computed
* `lineage_context` (LineageContext) - Lineage context information for this feature.
  WARNING: This field is primarily intended for internal use by Databricks systems and
  is automatically populated when features are created through Databricks notebooks or jobs.
  Users should not manually set this field as incorrect values may lead to inaccurate lineage tracking or unexpected behavior.
  This field will be set by feature-engineering client and should be left unset by SDK and terraform users
* `name` (string) - Name of the feature, extracted from the full three-part name (catalog.schema.name)
* `schema_name` (string) - Name of parent schema relative to its parent catalog
* `source` (DataSource) - The data source of the feature
* `timeseries_column` (TimeseriesColumn) - Column recording time, used for point-in-time joins, backfills, and aggregations

### AggregationFunction
* `approx_count_distinct` (ApproxCountDistinctFunction)
* `approx_percentile` (ApproxPercentileFunction)
* `avg` (AvgFunction)
* `count_function` (CountFunction)
* `first` (FirstFunction)
* `first_distinct` (FirstDistinctFunction)
* `first_n` (FirstNFunction)
* `last` (LastFunction)
* `last_distinct` (LastDistinctFunction)
* `last_n` (LastNFunction)
* `max` (MaxFunction)
* `min` (MinFunction)
* `stddev_pop` (StddevPopFunction)
* `stddev_samp` (StddevSampFunction)
* `sum` (SumFunction)
* `time_window` (TimeWindow) - The time window over which the aggregation is computed
* `var_pop` (VarPopFunction)
* `var_samp` (VarSampFunction)

### ApproxCountDistinctFunction
* `input` (string) - The input column from which the approximate count of distinct values is computed
* `relative_sd` (number) - The maximum relative standard deviation allowed (default defined by Spark)

### ApproxPercentileFunction
* `accuracy` (integer) - The accuracy parameter (higher is more accurate but slower)
* `input` (string) - The input column from which the approximate percentile is computed
* `percentile` (number) - The percentile value to compute (between 0 and 1)

### AvgFunction
* `input` (string) - The input column from which the average is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### ColumnSelection
* `column` (string) - Column name from source to select as the feature value

### CountFunction
* `input` (string) - The input column from which the count is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### CustomUdf
* `function_path` (string) - Fully qualified 3-part Unity Catalog path of the function to apply
* `input_bindings` (list of InputBinding) - Binds each UC function parameter to a source column.
  May be empty for zero-argument functions (e.g. a timestamp generator)

### DataSource
* `delta_table_source` (DeltaTableSource) - A Delta table data source
* `kafka_source` (KafkaSource) - A Kafka stream data source
* `lateness` (SourceLateness) - Completeness timing for this Feature's use of the source. This configuration is part of the
  Feature definition; it does not modify the underlying table or stream
* `request_source` (RequestSource) - A request-time data source
* `stream_source` (StreamSource) - A Stream data source

### DeltaTableSource
* `dataframe_schema` (string) - Schema of the resulting dataframe after transformations, in Spark StructType JSON format (from df.schema.json()).
  Required if transformation_sql is specified.
  Example: {"type":"struct","fields":[{"name":"col_a","type":"integer","nullable":true,"metadata":{}},{"name":"col_c","type":"integer","nullable":true,"metadata":{}}]}
* `filter_condition` (string) - Single WHERE clause to filter delta table before applying transformations. Will be row-wise evaluated, so should only include conditionals and projections
* `full_name` (string) - The full three-part (catalog, schema, table) name of the Delta table
* `transformation_sql` (string) - A single SQL SELECT expression applied after filter_condition.
  Should contains all the columns needed (eg. "SELECT *, col_a + col_b AS col_c FROM x.y.z WHERE col_a > 0" would have `transformation_sql` "*, col_a + col_b AS col_c")
  If transformation_sql is not provided, all columns of the delta table are present in the DataSource dataframe

### EntityColumn
* `name` (string) - The name of the entity column. For Kafka sources, use dot-prefixed path notation to reference
  fields within the key or value schema (e.g., "value.user_id", "key.partition_key"). For nested
  fields, the leaf node name (e.g., "user_id" from "value.trip_details.user_id") is what will
  be present in materialized tables and expected to match at query time.
  Colon-prefixed notation (e.g., "value:user_id") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### FieldDefinition
* `data_type` (string) - The scalar data type of the field. Possible values are: `BINARY`, `BOOLEAN`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INTEGER`, `LONG`, `SHORT`, `STRING`, `TIMESTAMP`
* `name` (string) - The name of the field

### FirstDistinctFunction
* `input` (string) - The input column from which the first N distinct values are returned
* `n` (integer) - The number of distinct values to return

### FirstFunction
* `input` (string) - The input column from which the first value is returned

### FirstNFunction
* `input` (string) - The input column from which the first N values are returned
* `n` (integer) - The number of values to return

### FlatSchema
* `fields` (list of FieldDefinition) - The list of fields in this schema

### Function
* `aggregation_function` (AggregationFunction) - An aggregation function applied over a time window
* `column_selection` (ColumnSelection) - Selects the latest value of a single column in a data source
* `custom_udf` (CustomUdf) - Applies a registered Unity Catalog function row-wise to source columns

### InputBinding
* `column` (string) - Source column whose value is passed for this parameter at execution time
* `parameter` (string) - Name of the UC function parameter

### JobContext
* `job_id` (integer) - The job ID where this API invoked
* `job_run_id` (integer) - The job run ID where this API was invoked

### KafkaSource
* `filter_condition` (string) - The filter condition applied to the source data before aggregation
* `name` (string) - Name of the Kafka source, used to identify it. This is used to look up the corresponding KafkaConfig object. Can be distinct from topic name

### LastDistinctFunction
* `input` (string) - The input column from which the last N distinct values are returned
* `n` (integer) - The number of distinct values to return

### LastFunction
* `input` (string) - The input column from which the last value is returned

### LastNFunction
* `input` (string) - The input column from which the last N values are returned
* `n` (integer) - The number of values to return

### LineageContext
* `job_context` (JobContext) - Job context information including job ID and run ID
* `notebook_id` (integer) - The notebook ID where this API was invoked

### MaxFunction
* `input` (string) - The input column from which the maximum is computed

### MinFunction
* `input` (string) - The input column from which the minimum is computed

### RequestSource
* `flat_schema` (FlatSchema) - A flat schema with scalar-typed fields only

### RollingWindow
* `delay` (string) - Non-negative analytic lag that evaluates the window this far in the past. Use this for timing
  variations unrelated to source lateness, such as a 30-day count as of one week ago. If unset,
  the analytic lag is zero. It composes with source.lateness when both are set
* `window_duration` (string) - The duration of the rolling window. Must be positive when set; absent means lifetime
  (aggregate over the entity's entire history)

### SawtoothWindow
* `delay` (string) - Delay is not currently supported for Sawtooth windows
* `window_duration` (string) - The duration of the window. Must be positive and span more than two days when set, so that both
  the batch (N-1 day) and stale-path (N-2 day) partial aggregates are well defined. The duration
  need not be a whole number of days (e.g. 3 days 15 minutes is allowed). Absent means lifetime
  (aggregate over the entity's entire history)

### SlidingWindow
* `delay` (string) - Non-negative analytic lag that evaluates the window this far in the past. Use this for timing
  variations unrelated to source lateness, such as a 30-day count as of one week ago. If unset,
  the analytic lag is zero. It composes with source.lateness when both are set
* `offset` (string) - Non-negative phase shift from the default midnight UTC alignment. For example, offset=22h on
  a 24h slide produces boundaries at 22:00 UTC (17:00 New York in standard time) instead of
  midnight UTC. If unset, the offset is zero. Must be shorter than slide_duration (and therefore
  window_duration)
* `slide_duration` (string) - The slide duration (interval by which windows advance, must be positive and less than duration)
* `window_duration` (string) - The duration of the sliding window. Must be positive when set; absent means lifetime
  (aggregate over the entity's entire history)

### SourceLateness
* `settling_delay` (string) - Non-negative time to wait after a window ends before treating its source data as complete.
  Training shifts the eligible evaluation time backwards by this duration so it does not join
  data that would still have been settling online. Materialization waits for the duration to
  elapse before publishing the window. If unset, source data is considered settled immediately

### StddevPopFunction
* `input` (string) - The input column from which the population standard deviation is computed. For Kafka sources,
  use dot-prefixed path notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### StddevSampFunction
* `input` (string) - The input column from which the sample standard deviation is computed

### StreamSource
* `dataframe_schema` (string) - Schema of the resulting dataframe after transformations, in Spark StructType
  JSON format (from df.schema.json()).
  Any subsequent functions operate against this dataframe
* `filter_condition` (string) - The filter condition applied to the source data before aggregation
* `full_name` (string) - Three-part full name of the Stream (catalog.schema.stream)
* `transformation_sql` (string) - The pipeline runs these SQL statements immediately after conversion into
  the schema specified on the Stream object

### SumFunction
* `input` (string) - The input column from which the sum is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### TimeWindow
* `rolling` (RollingWindow)
* `sawtooth` (SawtoothWindow) - A sawtooth window served via the hybrid batch + streaming path
* `sliding` (SlidingWindow)
* `start_time` (string) - Earliest event-time boundary at which the Feature may emit an output. This gates outputs, not
  the historical inputs read by a window. For example, a 365-day window with
  start_time=2026-01-01 begins emitting partial-window values on that date instead of waiting
  for 365 days of data; a lifetime window produces no output before start_time. If unset,
  tumbling and fixed-duration sliding windows first emit at an offset-aligned boundary after a
  full window can be formed. If unset, lifetime sliding windows and rolling windows emit as soon as
  eligible source data exists
* `tumbling` (TumblingWindow)

### TimeseriesColumn
* `name` (string) - The name of the timeseries column. For Kafka sources, use dot-prefixed path notation to
  reference fields within the key or value schema (e.g., "value.event_timestamp"). For nested
  fields, the leaf node name (e.g., "event_timestamp" from "value.event_details.event_timestamp")
  is what will be present in materialized tables and expected to match at query time.
  Colon-prefixed notation (e.g., "value:event_timestamp") is supported for
  backwards compatibility but is deprecated; migrate to dot notation

### TumblingWindow
* `delay` (string) - Non-negative analytic lag that evaluates the window this far in the past. Use this for timing
  variations unrelated to source lateness, such as a 30-day count as of one week ago. If unset,
  the analytic lag is zero. It composes with source.lateness when both are set
* `offset` (string) - Non-negative phase shift from the default midnight UTC alignment. For example, offset=22h on
  a 24h window produces boundaries at 22:00 UTC (17:00 New York in standard time) instead of
  midnight UTC. If unset, the offset is zero. Must be shorter than window_duration
* `window_duration` (string) - The duration of each tumbling window (non-overlapping, fixed-duration windows)

### VarPopFunction
* `input` (string) - The input column from which the population variance is computed

### VarSampFunction
* `input` (string) - The input column from which the sample variance is computed