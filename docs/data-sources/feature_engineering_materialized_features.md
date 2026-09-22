---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_materialized_features Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `feature_name` (string, optional) - Filter by feature name. If specified, only materialized features materialized from this feature will be returned
* `page_size` (integer, optional) - The maximum number of results to return. Defaults to 100 if not specified. Cannot be greater than 1000
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `materialized_features`. It is a list of resources, each with the following attributes:
* `budget_policy_id` (string) - The ID of the budget policy used to attribute the serverless compute cost of this
  materialization. If not specified, a default budget policy may be applied
* `cron_schedule` (string, deprecated)
* `cron_schedule_trigger` (CronSchedule) - A cron-based schedule trigger for the materialization pipeline
* `feature_name` (string) - The full name of the feature in Unity Catalog
* `is_online` (boolean) - True if this is an online materialized feature. False if it is an offline materialized feature
* `last_materialization_time` (string) - The timestamp when the pipeline last ran and updated the materialized feature values.
  If the pipeline has not run yet, this field will be null
* `latest_backfill_operation` (string) - Name of the latest backfill operation on this materialized feature. Format: operations/{operation_id}
* `materialized_feature_id` (string) - Server-assigned unique identifier for the materialized feature
* `offline_store_config` (OfflineStoreConfig) - Destination for writing feature values to an offline Delta table
* `online_store_config` (OnlineStoreConfig) - Destination for writing feature values to an online Lakebase table
* `pipeline_schedule_state` (string) - The schedule state of the materialization pipeline.
  Hidden from GraphQL: being deprecated, so not exposed to Catalog Explorer. Possible values are: `ACTIVE`, `PAUSED`, `SNAPSHOT`
* `streaming_mode` (StreamingMode) - The Structured Streaming trigger mode used for materialization. Real-time mode (RTM) targets
  sub-second latency for operational workloads; micro-batch mode (MBM) favors cost efficiency
  for ETL and analytics workloads
* `table_name` (string) - The fully qualified Unity Catalog path to the table containing the materialized feature (Delta table or Lakebase table). Output only
* `table_trigger` (TableTrigger) - A trigger that fires when the upstream source table changes
* `tags` (object) - Custom tags to associate with this materialization. They are applied to the materialization
  job (for batch features) or pipeline (for streaming features) and forwarded to the underlying
  compute as cluster tags, so materialization cost can be attributed in the billing system
  tables. These tags apply only to the materialization compute; they are not applied to the
  Unity Catalog Feature resource itself, whose tags are managed separately through the Unity
  Catalog tagging API. A maximum of 25 tags is supported; keys and values are subject to the
  same limitations as cluster tags

### CronSchedule
* `cron_expression` (string) - The cron expression defining the schedule (e.g., "0 0 * * *" for daily at midnight). The
  schedule is interpreted in the UTC time zone. Required when mode is MANUAL (or unset). Left
  empty when mode is DERIVED, where the service computes it (aligned to UTC) from the features'
  window timing and fills it in on the response
* `mode` (string) - How the schedule is determined. Defaults to MANUAL when unset. Possible values are: `DERIVED`, `MANUAL`

### OfflineStoreConfig
* `catalog_name` (string) - The Unity Catalog catalog name
* `schema_name` (string) - The Unity Catalog schema name
* `table_name_prefix` (string) - Prefix for Unity Catalog table name.
  The materialized feature will be stored in a table with this prefix and a generated postfix

### OnlineStoreConfig
* `catalog_name` (string) - The Unity Catalog catalog name. This name is also used as the Lakebase logical database name.
  Quoting is handled by the backend where needed, do not pre-quote it
* `online_store_name` (string) - The name of the target online store
* `schema_name` (string) - The Unity Catalog schema name. This name is also used as the Lakebase schema name under the database.
  Quoting is handled by the backend where needed, do not pre-quote it
* `table_name_prefix` (string) - Prefix for Unity Catalog table name.
  The materialized feature will be stored in a Lakebase table with this prefix and a generated postfix

### StreamingMode
* `freshness_target` (string) - The desired data freshness for feature materialization, expressed as a
  duration string (e.g. "1 minute")
* `mode` (string) - The type of streaming mode used by the materialization pipeline. Possible values are: `STREAMING_MODE_TYPE_MBM`, `STREAMING_MODE_TYPE_RTM`

### TableTrigger