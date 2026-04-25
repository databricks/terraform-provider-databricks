---
subcategory: "Postgres"
---
# databricks_postgres_synced_table Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Output only. The Full resource name of the synced table in Postgres
  where (catalog, schema, table) are the UC entity names.
  
  Format "synced_tables/{catalog}.{schema}.{table}"
  
  For the corresponding source table in the Unity catalog look for the "source_table_full_name" attribute
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string)
* `name` (string) - Output only. The Full resource name of the synced table in Postgres
  where (catalog, schema, table) are the UC entity names.
  
  Format "synced_tables/{catalog}.{schema}.{table}"
  
  For the corresponding source table in the Unity catalog look for the "source_table_full_name" attribute
* `spec` (SyncedTableSyncedTableSpec) - Configuration details of the synced table, such as the source table, scheduling policy, etc.
  This attribute is specified at creation time and most fields are returned as is on subsequent queries
* `status` (SyncedTableSyncedTableStatus) - Synced Table data synchronization status
* `uid` (string) - The Unity Catalog table ID for this synced table

### DeltaTableSyncInfo
* `delta_commit_time` (string) - The timestamp when the above Delta version was committed in the source Delta table.
  Note: This is the Delta commit time, not the time the data was written to the synced table
* `delta_commit_version` (integer) - The Delta Lake commit version that was last successfully synced

### NewPipelineSpec
* `budget_policy_id` (string) - Budget policy to set on the newly created pipeline
* `storage_catalog` (string) - UC catalog for the pipeline to store intermediate files (checkpoints, event logs etc).
  This needs to be a standard catalog where the user has permissions to create Delta tables
* `storage_schema` (string) - UC schema for the pipeline to store intermediate files (checkpoints, event logs etc).
  This needs to be in the standard catalog where the user has permissions to create Delta tables

### SyncedTablePipelineProgress
* `estimated_completion_time_seconds` (number) - The estimated time remaining to complete this update in seconds
* `latest_version_currently_processing` (integer) - The source table Delta version that was last processed by the pipeline. The pipeline may not
  have completely processed this version yet
* `sync_progress_completion` (number) - The completion ratio of this update. This is a number between 0 and 1
* `synced_row_count` (integer) - The number of rows that have been synced in this update
* `total_row_count` (integer) - The total number of rows that need to be synced in this update. This number may be an estimate

### SyncedTablePosition
* `delta_table_sync_info` (DeltaTableSyncInfo)
* `sync_end_time` (string) - The end timestamp of the most recent successful synchronization.
  This is the time when the data is available in the synced table
* `sync_start_time` (string) - The starting timestamp of the most recent successful synchronization from the source table
  to the destination (synced) table.
  Note this is the starting timestamp of the sync operation, not the end time.
  E.g., for a batch, this is the time when the sync operation started

### SyncedTableSyncedTableSpec
* `branch` (string) - The full resource name the branch associated with the table.
  
  Format: "projects/{project_id}/branches/{branch_id}"
* `create_database_objects_if_missing` (boolean) - If true, the synced table's logical database and schema resources in PG
  will be created if they do not already exist.
  The request will fail if this is false and the database/schema do not exist.
  
  Defaults to true if omitted
* `existing_pipeline_id` (string) - ID of an existing pipeline to bin-pack this synced table into.
  At most one of existing_pipeline_id and new_pipeline_spec should be defined.
  
  The pipeline used for the synced table is returned via the top level pipeline_id attribute
* `new_pipeline_spec` (NewPipelineSpec) - Specification for creating a new pipeline.
  At most one of existing_pipeline_id and new_pipeline_spec should be defined.
  
  The pipeline used for the synced table is returned via the top level pipeline_id attribute
* `postgres_database` (string) - The Postgres database name where the synced table will be created in.
  
  If this synced table is created inside a Lakebase Catalog, this attribute can be omitted on creation and is inferred
  from the postgres_database associated with the Lakebase Catalog. If specified when inside a Lakebase Catalog, the value must match.
  
  A value must be specified when creating a synced table inside a Standard Catalog
* `primary_key_columns` (list of string) - Primary Key columns to be used for data insert/update in the destination
* `scheduling_policy` (string) - Scheduling policy of the underlying pipeline. Possible values are: `CONTINUOUS`, `SNAPSHOT`, `TRIGGERED`
* `source_table_full_name` (string) - Three-part (catalog, schema, table) name of the source Delta table.
  
  For the corresponding destination table, use any of the two:
  
  * synced_table_id used at the creation of the SyncedTable
  * "name" consisting of "synced_tables/" prefix and the full name of the destination table
* `timeseries_key` (string) - Time series key to deduplicate (tie-break) rows with the same primary key

### SyncedTableSyncedTableStatus
* `detailed_state` (string) - The state of the synced table. Possible values are: `SYNCED_TABLE_OFFLINE`, `SYNCED_TABLE_OFFLINE_FAILED`, `SYNCED_TABLE_ONLINE`, `SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE`, `SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE`, `SYNCED_TABLE_ONLINE_PIPELINE_FAILED`, `SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE`, `SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES`, `SYNCED_TABLE_PROVISIONING`, `SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT`, `SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES`
* `last_processed_commit_version` (integer) - The last source table Delta version that was successfully synced to the synced table
* `last_sync` (SyncedTablePosition) - Summary of the last successful synchronization from source to destination
* `last_sync_time` (string) - The end timestamp of the last time any data was synchronized from the source table to the synced
  table. This is when the data is available in the synced table
* `message` (string) - A text description of the current state of the synced table
* `ongoing_sync_progress` (SyncedTablePipelineProgress)
* `pipeline_id` (string) - ID of the associated pipeline
* `project` (string) - The full resource name of the project associated with the table.
  
  Format: "projects/{project_id}"
* `provisioning_phase` (string) - The current phase of the data synchronization pipeline. Possible values are: `PROVISIONING_PHASE_INDEX_SCAN`, `PROVISIONING_PHASE_INDEX_SORT`, `PROVISIONING_PHASE_MAIN`
* `unity_catalog_provisioning_state` (string) - The provisioning state of the synced table entity in Unity Catalog. Possible values are: `ACTIVE`, `DEGRADED`, `DELETING`, `FAILED`, `PROVISIONING`, `UPDATING`