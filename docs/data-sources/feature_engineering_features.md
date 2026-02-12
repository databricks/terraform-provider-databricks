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
* `filter_condition` (string) - The filter condition applied to the source data before aggregation
* `full_name` (string) - The full three-part name (catalog, schema, name) of the feature
* `function` (Function) - The function by which the feature is computed
* `inputs` (list of string) - The input columns from which the feature is computed
* `lineage_context` (LineageContext) - WARNING: This field is primarily intended for internal use by Databricks systems and
  is automatically populated when features are created through Databricks notebooks or jobs.
  Users should not manually set this field as incorrect values may lead to inaccurate lineage tracking or unexpected behavior.
  This field will be set by feature-engineering client and should be left unset by SDK and terraform users
* `source` (DataSource) - The data source of the feature
* `time_window` (TimeWindow) - The time window in which the feature is computed

### ColumnIdentifier
* `variant_expr_path` (string) - String representation of the column name or variant expression path. For nested fields, the leaf value is what will be present in materialized tables
  and expected to match at query time. For example, the leaf node of value:trip_details.location_details.pickup_zip is pickup_zip

### ContinuousWindow
* `offset` (string) - The offset of the continuous window (must be non-positive)
* `window_duration` (string) - The duration of the continuous window (must be positive)

### DataSource
* `delta_table_source` (DeltaTableSource)
* `kafka_source` (KafkaSource)

### DeltaTableSource
* `entity_columns` (list of string) - The entity columns of the Delta table
* `full_name` (string) - The full three-part (catalog, schema, table) name of the Delta table
* `timeseries_column` (string) - The timeseries column of the Delta table

### Function
* `extra_parameters` (list of FunctionExtraParameter) - Extra parameters for parameterized functions
* `function_type` (string) - The type of the function. Possible values are: `APPROX_COUNT_DISTINCT`, `APPROX_PERCENTILE`, `AVG`, `COUNT`, `FIRST`, `LAST`, `MAX`, `MIN`, `STDDEV_POP`, `STDDEV_SAMP`, `SUM`, `VAR_POP`, `VAR_SAMP`

### FunctionExtraParameter
* `key` (string) - The name of the parameter
* `value` (string) - The value of the parameter

### JobContext
* `job_id` (integer) - The job ID where this API invoked
* `job_run_id` (integer) - The job run ID where this API was invoked

### KafkaSource
* `entity_column_identifiers` (list of ColumnIdentifier) - The entity column identifiers of the Kafka source
* `name` (string) - Name of the Kafka source, used to identify it. This is used to look up the corresponding KafkaConfig object. Can be distinct from topic name
* `timeseries_column_identifier` (ColumnIdentifier) - The timeseries column identifier of the Kafka source

### LineageContext
* `job_context` (JobContext) - Job context information including job ID and run ID
* `notebook_id` (integer) - The notebook ID where this API was invoked

### SlidingWindow
* `slide_duration` (string) - The slide duration (interval by which windows advance, must be positive and less than duration)
* `window_duration` (string) - The duration of the sliding window

### TimeWindow
* `continuous` (ContinuousWindow)
* `sliding` (SlidingWindow)
* `tumbling` (TumblingWindow)

### TumblingWindow
* `window_duration` (string) - The duration of each tumbling window (non-overlapping, fixed-duration windows)