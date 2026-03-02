---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_materialized_feature Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `feature_name` (string, required) - The full name of the feature in Unity Catalog
* `cron_schedule` (string, optional) - The quartz cron expression that defines the schedule of the materialization pipeline. The schedule is evaluated in the UTC timezone
* `offline_store_config` (OfflineStoreConfig, optional)
* `online_store_config` (OnlineStoreConfig, optional)
* `pipeline_schedule_state` (string, optional) - The schedule state of the materialization pipeline. Possible values are: `ACTIVE`, `PAUSED`, `SNAPSHOT`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### OfflineStoreConfig
* `catalog_name` (string, required) - The Unity Catalog catalog name
* `schema_name` (string, required) - The Unity Catalog schema name
* `table_name_prefix` (string, required) - Prefix for Unity Catalog table name.
  The materialized feature will be stored in a table with this prefix and a generated postfix

### OnlineStoreConfig
* `catalog_name` (string, required) - The Unity Catalog catalog name. This name is also used as the Lakebase logical database name.
  Quoting is handled by the backend where needed, do not pre-quote it
* `online_store_name` (string, required) - The name of the target online store
* `schema_name` (string, required) - The Unity Catalog schema name. This name is also used as the Lakebase schema name under the database.
  Quoting is handled by the backend where needed, do not pre-quote it
* `table_name_prefix` (string, required) - Prefix for Unity Catalog table name.
  The materialized feature will be stored in a Lakebase table with this prefix and a generated postfix

## Attributes
In addition to the above arguments, the following attributes are exported:
* `last_materialization_time` (string) - The timestamp when the pipeline last ran and updated the materialized feature values.
  If the pipeline has not run yet, this field will be null
* `materialized_feature_id` (string) - Unique identifier for the materialized feature
* `table_name` (string) - The fully qualified Unity Catalog path to the table containing the materialized feature (Delta table or Lakebase table). Output only

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