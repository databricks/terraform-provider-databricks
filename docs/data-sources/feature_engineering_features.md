---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_features Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of results to return


## Attributes
This data source exports a single attribute, `features`. It is a list of resources, each with the following attributes:
* `description` (string) - The description of the feature
* `full_name` (string) - The full three-part name (catalog, schema, name) of the feature
* `function` (Function) - The function by which the feature is computed
* `inputs` (list of string) - The input columns from which the feature is computed
* `source` (DataSource) - The data source of the feature
* `time_window` (TimeWindow) - The time window in which the feature is computed

### DataSource
* `delta_table_source` (DeltaTableSource)

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

### TimeWindow
* `duration` (string) - The duration of the time window
* `offset` (string) - The offset of the time window