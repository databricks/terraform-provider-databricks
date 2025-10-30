---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_materialized_feature Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `feature_name` (string, required) - The full name of the feature in Unity Catalog
* `offline_store_config` (OfflineStoreConfig, optional)
* `online_store_config` (OnlineStore, optional)
* `pipeline_schedule_state` (string, optional) - The schedule state of the materialization pipeline. Possible values are: `ACTIVE`, `PAUSED`, `SNAPSHOT`

### OfflineStoreConfig
* `catalog_name` (string, required) - The Unity Catalog catalog name
* `schema_name` (string, required) - The Unity Catalog schema name
* `table_name_prefix` (string, required) - Prefix for Unity Catalog table name.
  The materialized feature will be stored in a table with this prefix and a generated postfix

### OnlineStore
* `capacity` (string, required) - The capacity of the online store. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"
* `name` (string, required) - The name of the online store. This is the unique identifier for the online store
* `read_replica_count` (integer, optional) - The number of read replicas for the online store. Defaults to 0

## Attributes
In addition to the above arguments, the following attributes are exported:
* `last_materialization_time` (string) - The timestamp when the pipeline last ran and updated the materialized feature values.
  If the pipeline has not run yet, this field will be null
* `materialized_feature_id` (string) - Unique identifier for the materialized feature
* `table_name` (string) - The fully qualified Unity Catalog path to the table containing the materialized feature (Delta table or Lakebase table). Output only

### OnlineStore
* `creation_time` (string) - The timestamp when the online store was created
* `creator` (string) - The email of the creator of the online store
* `state` (string) - The current state of the online store. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "materialized_feature_id"
  to = databricks_feature_engineering_materialized_feature.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_feature_engineering_materialized_feature.this "materialized_feature_id"
```