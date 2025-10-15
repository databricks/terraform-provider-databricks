---
subcategory: "Machine Learning"
---
# databricks_materialized_features_feature_tags Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `feature_name` (string, required)
* `table_name` (string, required)
* `page_size` (integer, optional) - The maximum number of results to return
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource


## Attributes
This data source exports a single attribute, `feature_tags`. It is a list of resources, each with the following attributes:
* `key` (string)
* `value` (string)