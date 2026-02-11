---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_materialized_feature Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `materialized_feature_id` (string, required) - Unique identifier for the materialized feature
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `cron_schedule` (string) - The quartz cron expression that defines the schedule of the materialization pipeline. The schedule is evaluated in the UTC timezone
* `feature_name` (string) - The full name of the feature in Unity Catalog
* `last_materialization_time` (string) - The timestamp when the pipeline last ran and updated the materialized feature values.
  If the pipeline has not run yet, this field will be null
* `materialized_feature_id` (string) - Unique identifier for the materialized feature
* `offline_store_config` (OfflineStoreConfig)
* `online_store_config` (OnlineStoreConfig)
* `pipeline_schedule_state` (string) - The schedule state of the materialization pipeline. Possible values are: `ACTIVE`, `PAUSED`, `SNAPSHOT`
* `table_name` (string) - The fully qualified Unity Catalog path to the table containing the materialized feature (Delta table or Lakebase table). Output only

### OfflineStoreConfig
* `catalog_name` (string) - The Unity Catalog catalog name
* `schema_name` (string) - The Unity Catalog schema name
* `table_name_prefix` (string) - Prefix for Unity Catalog table name.
  The materialized feature will be stored in a table with this prefix and a generated postfix

### OnlineStoreConfig
* `catalog_name` (string) - The Unity Catalog catalog name. This name is also used as the Lakebase logical database name
* `online_store_name` (string) - The name of the target online store
* `schema_name` (string) - The Unity Catalog schema name
* `table_name_prefix` (string) - Prefix for Unity Catalog table name.
  The materialized feature will be stored in a Lakebase table with this prefix and a generated postfix