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
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource

### DataSource
* `delta_table_source` (DeltaTableSource, optional)

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

### TimeWindow
* `duration` (string, required) - The duration of the time window
* `offset` (string, optional) - The offset of the time window

## Attributes
In addition to the above arguments, the following attributes are exported:

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
terraform import databricks_feature_engineering_feature "full_name"
```