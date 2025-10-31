---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_materialized_feature Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `materialized_feature_id` (string, required) - Unique identifier for the materialized feature

## Attributes
The following attributes are exported:
* `feature_name` (string) - The full name of the feature in Unity Catalog
* `last_materialization_time` (string) - The timestamp when the pipeline last ran and updated the materialized feature values.
  If the pipeline has not run yet, this field will be null
* `materialized_feature_id` (string) - Unique identifier for the materialized feature
* `offline_store_config` (OfflineStoreConfig)
* `online_store_config` (OnlineStore)
* `pipeline_schedule_state` (string) - The schedule state of the materialization pipeline. Possible values are: `ACTIVE`, `PAUSED`, `SNAPSHOT`
* `table_name` (string) - The fully qualified Unity Catalog path to the table containing the materialized feature (Delta table or Lakebase table). Output only

### OfflineStoreConfig
* `catalog_name` (string) - The Unity Catalog catalog name
* `schema_name` (string) - The Unity Catalog schema name
* `table_name_prefix` (string) - Prefix for Unity Catalog table name.
  The materialized feature will be stored in a table with this prefix and a generated postfix

### OnlineStore
* `capacity` (string) - The capacity of the online store. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"
* `creation_time` (string) - The timestamp when the online store was created
* `creator` (string) - The email of the creator of the online store
* `name` (string) - The name of the online store. This is the unique identifier for the online store
* `read_replica_count` (integer) - The number of read replicas for the online store. Defaults to 0
* `state` (string) - The current state of the online store. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`