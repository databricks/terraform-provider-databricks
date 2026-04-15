---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_features Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of results to return
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `features`. It is a list of resources, each with the following attributes:
* `description` (string) - The description of the feature
* `entities` (list of EntityColumn) - The entity columns for the feature, used as aggregation keys and for query-time lookup
* `filter_condition` (string, deprecated) - Deprecated: Use DeltaTableSource.filter_condition or KafkaSource.filter_condition instead. Kept for backwards compatibility.
  The filter condition applied to the source data before aggregation
* `full_name` (string) - The full three-part name (catalog, schema, name) of the feature
* `function` (Function) - The function by which the feature is computed
* `inputs` (list of string, deprecated) - Deprecated: Use AggregationFunction.inputs instead. Kept for backwards compatibility.
  The input columns from which the feature is computed
* `lineage_context` (LineageContext) - Lineage context information for this feature.
  WARNING: This field is primarily intended for internal use by Databricks systems and
  is automatically populated when features are created through Databricks notebooks or jobs.
  Users should not manually set this field as incorrect values may lead to inaccurate lineage tracking or unexpected behavior.
  This field will be set by feature-engineering client and should be left unset by SDK and terraform users
* `source` (DataSource) - The data source of the feature
* `time_window` (TimeWindow, deprecated) - Deprecated: Use Function.aggregation_function.time_window instead. Kept for backwards compatibility.
  The time window in which the feature is computed
* `timeseries_column` (TimeseriesColumn) - Column recording time, used for point-in-time joins, backfills, and aggregations

### AggregationFunction
* `approx_count_distinct` (ApproxCountDistinctFunction)
* `approx_percentile` (ApproxPercentileFunction)
* `avg` (AvgFunction)
* `count_function` (CountFunction)
* `first` (FirstFunction)
* `last` (LastFunction)
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
  TODO(FS-939): Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### ColumnIdentifier
* `variant_expr_path` (string) - String representation of the column name using dot-prefixed path notation. For nested fields, the leaf value is what will be present in materialized tables
  and expected to match at query time. For example, the leaf node of value.trip_details.location_details.pickup_zip is pickup_zip

### ColumnSelection
* `column` (string) - Column name from source to select as the feature value

### ContinuousWindow
* `offset` (string) - The offset of the continuous window (must be non-positive)
* `window_duration` (string) - The duration of the continuous window (must be positive)

### CountFunction
* `input` (string) - The input column from which the count is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### DataSource
* `delta_table_source` (DeltaTableSource) - A Delta table data source
* `kafka_source` (KafkaSource) - A Kafka stream data source
* `request_source` (RequestSource) - A request-time data source

### DeltaTableSource
* `dataframe_schema` (string) - Schema of the resulting dataframe after transformations, in Spark StructType JSON format (from df.schema.json()).
  Required if transformation_sql is specified.
  Example: {"type":"struct","fields":[{"name":"col_a","type":"integer","nullable":true,"metadata":{}},{"name":"col_c","type":"integer","nullable":true,"metadata":{}}]}
* `entity_columns` (list of string, deprecated) - Deprecated: Use Feature.entity instead. Kept for backwards compatibility.
  The entity columns of the Delta table
* `filter_condition` (string) - Single WHERE clause to filter delta table before applying transformations. Will be row-wise evaluated, so should only include conditionals and projections
* `full_name` (string) - The full three-part (catalog, schema, table) name of the Delta table
* `timeseries_column` (string, deprecated) - Deprecated: Use Feature.timeseries_column instead. Kept for backwards compatibility.
  The timeseries column of the Delta table
* `transformation_sql` (string) - A single SQL SELECT expression applied after filter_condition.
  Should contains all the columns needed (eg. "SELECT *, col_a + col_b AS col_c FROM x.y.z WHERE col_a > 0" would have `transformation_sql` "*, col_a + col_b AS col_c")
  If transformation_sql is not provided, all columns of the delta table are present in the DataSource dataframe

### EntityColumn
* `name` (string) - The name of the entity column. For Kafka sources, use dot-prefixed path notation to reference
  fields within the key or value schema (e.g., "value.user_id", "key.partition_key"). For nested
  fields, the leaf node name (e.g., "user_id" from "value.trip_details.user_id") is what will
  be present in materialized tables and expected to match at query time.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:user_id") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### FieldDefinition
* `data_type` (string) - The scalar data type of the field. Possible values are: `BINARY`, `BOOLEAN`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INTEGER`, `LONG`, `SHORT`, `STRING`, `TIMESTAMP`
* `name` (string) - The name of the field

### FirstFunction
* `input` (string) - The input column from which the first value is returned

### FlatSchema
* `fields` (list of FieldDefinition) - The list of fields in this schema

### Function
* `aggregation_function` (AggregationFunction) - An aggregation function applied over a time window
* `column_selection` (ColumnSelection) - Selects the latest value of a single column in a data source
* `extra_parameters` (list of FunctionExtraParameter, deprecated) - Deprecated: Use the function oneof with AggregationFunction instead. Kept for backwards compatibility.
  Extra parameters for parameterized functions
* `function_type` (string, deprecated) - Deprecated: Use the function oneof with AggregationFunction instead. Kept for backwards compatibility.
  The type of the function. Possible values are: `APPROX_COUNT_DISTINCT`, `APPROX_PERCENTILE`, `AVG`, `COUNT`, `FIRST`, `LAST`, `MAX`, `MIN`, `STDDEV_POP`, `STDDEV_SAMP`, `SUM`, `VAR_POP`, `VAR_SAMP`

### FunctionExtraParameter
* `key` (string) - The name of the parameter
* `value` (string) - The value of the parameter

### JobContext
* `job_id` (integer) - The job ID where this API invoked
* `job_run_id` (integer) - The job run ID where this API was invoked

### KafkaSource
* `entity_column_identifiers` (list of ColumnIdentifier, deprecated) - Deprecated: Use Feature.entity instead. Kept for backwards compatibility.
  The entity column identifiers of the Kafka source
* `filter_condition` (string) - The filter condition applied to the source data before aggregation
* `name` (string) - Name of the Kafka source, used to identify it. This is used to look up the corresponding KafkaConfig object. Can be distinct from topic name
* `timeseries_column_identifier` (ColumnIdentifier, deprecated) - Deprecated: Use Feature.timeseries_column instead. Kept for backwards compatibility.
  The timeseries column identifier of the Kafka source

### LastFunction
* `input` (string) - The input column from which the last value is returned

### LineageContext
* `job_context` (JobContext) - Job context information including job ID and run ID
* `notebook_id` (integer) - The notebook ID where this API was invoked

### MaxFunction
* `input` (string) - The input column from which the maximum is computed

### MinFunction
* `input` (string) - The input column from which the minimum is computed

### RequestSource
* `flat_schema` (FlatSchema) - A flat schema with scalar-typed fields only

### SlidingWindow
* `slide_duration` (string) - The slide duration (interval by which windows advance, must be positive and less than duration)
* `window_duration` (string) - The duration of the sliding window

### StddevPopFunction
* `input` (string) - The input column from which the population standard deviation is computed. For Kafka sources,
  use dot-prefixed path notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### StddevSampFunction
* `input` (string) - The input column from which the sample standard deviation is computed

### SumFunction
* `input` (string) - The input column from which the sum is computed. For Kafka sources, use dot-prefixed path
  notation (e.g., "value.amount"). For nested fields, the leaf node name is used.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:amount") is supported for backwards
  compatibility but is deprecated; migrate to dot notation

### TimeWindow
* `continuous` (ContinuousWindow)
* `sliding` (SlidingWindow)
* `tumbling` (TumblingWindow)

### TimeseriesColumn
* `name` (string) - The name of the timeseries column. For Kafka sources, use dot-prefixed path notation to
  reference fields within the key or value schema (e.g., "value.event_timestamp"). For nested
  fields, the leaf node name (e.g., "event_timestamp" from "value.event_details.event_timestamp")
  is what will be present in materialized tables and expected to match at query time.
  TODO(FS-939): Colon-prefixed notation (e.g., "value:event_timestamp") is supported for
  backwards compatibility but is deprecated; migrate to dot notation

### TumblingWindow
* `window_duration` (string) - The duration of each tumbling window (non-overlapping, fixed-duration windows)

### VarPopFunction
* `input` (string) - The input column from which the population variance is computed

### VarSampFunction
* `input` (string) - The input column from which the sample variance is computed