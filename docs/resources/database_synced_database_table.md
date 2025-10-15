---
subcategory: "Database Instances"
---
# databricks_database_synced_database_table Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

Lakebase Synced Database Tables are Postgres tables automatically synced from a source table inside Unity Catalog.
They can be used to serve realtime queries without the operational overhead of managing ETL pipelines. 

Synced Database Tables can be configured inside either Database Catalogs or Standard Catalogs. Multiple
Synced Database Tables can be bin packed inside a single pipeline to optimize costs.

## Example Usage
### Creating a Synced Database Table inside a Database Catalog

This example creates a Synced Database Table inside a Database Catalog.

```hcl
resource "databricks_database_synced_database_table" "this" {
  name = "my_database_catalog.public.synced_table"
  # logical_database_name is optional for synced tables created in
  # database catalogs, and is inferred from the catalog's database.
  # If specified, it must match the catalog's database.
  logical_database_name = "databricks_postgres"
  spec = {
    scheduling_policy = "SNAPSHOT"
    source_table_full_name = "source_delta.tpch.customer"
    primary_key_columns = ["c_custkey"]
    create_database_objects_if_missing = true
    new_pipeline_spec = {
      storage_catalog = "source_delta"
      storage_schema = "tpch"
    }
  }
}
```

### Creating a Synced Database Table inside a Standard Catalog

This example creates a Synced Database Table inside a Standard Catalog.

```hcl
resource "databricks_database_synced_database_table" "this" {
  name = "my_standard_catalog.public.synced_table"
  # logical_database_name is required for synced tables created in
  # standard catalogs.
  logical_database_name = "databricks_postgres"
  # database instnace name is required for synced tables created in
  # standard catalogs.
  database_instance_name = "my-database-instance"
  spec = {
    scheduling_policy = "SNAPSHOT"
    source_table_full_name = "source_delta.tpch.customer"
    primary_key_columns = ["c_custkey"]
    create_database_objects_if_missing = true
    new_pipeline_spec = {
      storage_catalog = "source_delta"
      storage_schema = "tpch"
    }
  }
}
```

### Creating multiple Synced Database Tables and bin packing them into a single pipeline

This example creates two Synced Database Tables. The first one specifies a new pipeline spec,
which generates a new pipeline. The second one utilizes the pipeline ID of the first table.

```hcl
resource "databricks_database_instance" "instance" {
  name = "my-database-instance"
  capacity = "CU_1"
}


resource "databricks_database_synced_database_table" "synced_table_1" {
  name = "my_standard_catalog.public.synced_table1"
  # logical_database_name is required for synced tables created in
  # standard catalogs.
  logical_database_name = "databricks_postgres"
  # database instnace name is required for synced tables created in
  # standard catalogs.
  database_instance_name = databricks_database_instance.instance.name
  spec = {
    scheduling_policy = "SNAPSHOT"
    source_table_full_name = "source_delta.tpch.customer"
    primary_key_columns = ["c_custkey"]
    create_database_objects_if_missing = true
    new_pipeline_spec = {
      storage_catalog = "source_delta"
      storage_schema = "tpch"
    }
  }
}

resource "databricks_database_synced_database_table" "synced_table_2" {
  name = "my_standard_catalog.public.synced_table2"
  # logical_database_name is required for synced tables created in
  # standard catalogs.
  logical_database_name = "databricks_postgres"
  # database instnace name is required for synced tables created in
  # standard catalogs.
  database_instance_name = databricks_database_instance.instance.name
  spec = {
    scheduling_policy = "SNAPSHOT"
    source_table_full_name = "source_delta.tpch.customer"
    primary_key_columns = ["c_custkey"]
    create_database_objects_if_missing = true
    existing_pipeline_id = databricks_database_synced_database_table.synced_table_1.data_synchronization_status.pipeline_id
  }
}
```



## Arguments
The following arguments are supported:
* `name` (string, required) - Full three-part (catalog, schema, table) name of the table
* `database_instance_name` (string, optional) - Name of the target database instance. This is required when creating synced database tables in standard catalogs.
  This is optional when creating synced database tables in registered catalogs. If this field is specified
  when creating synced database tables in registered catalogs, the database instance name MUST
  match that of the registered catalog (or the request will be rejected)
* `logical_database_name` (string, optional) - Target Postgres database object (logical database) name for this table.
  
  When creating a synced table in a registered Postgres catalog, the
  target Postgres database name is inferred to be that of the registered catalog.
  If this field is specified in this scenario, the Postgres database name MUST
  match that of the registered catalog (or the request will be rejected).
  
  When creating a synced table in a standard catalog, this field is required.
  In this scenario, specifying this field will allow targeting an arbitrary postgres database.
  Note that this has implications for the `create_database_objects_is_missing` field in `spec`
* `spec` (SyncedTableSpec, optional)
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource

### NewPipelineSpec
* `budget_policy_id` (string, optional) - Budget policy of this pipeline
* `storage_catalog` (string, optional) - This field needs to be specified if the destination catalog is a managed postgres catalog.
  
  UC catalog for the pipeline to store intermediate files (checkpoints, event logs etc).
  This needs to be a standard catalog where the user has permissions to create Delta tables
* `storage_schema` (string, optional) - This field needs to be specified if the destination catalog is a managed postgres catalog.
  
  UC schema for the pipeline to store intermediate files (checkpoints, event logs etc).
  This needs to be in the standard catalog where the user has permissions to create Delta tables

### SyncedTableSpec
* `create_database_objects_if_missing` (boolean, optional) - If true, the synced table's logical database and schema resources in PG
  will be created if they do not already exist
* `existing_pipeline_id` (string, optional) - At most one of existing_pipeline_id and new_pipeline_spec should be defined.
  
  If existing_pipeline_id is defined, the synced table will be bin packed into the existing pipeline
  referenced. This avoids creating a new pipeline and allows sharing existing compute.
  In this case, the scheduling_policy of this synced table must match the scheduling policy of the existing pipeline
* `new_pipeline_spec` (NewPipelineSpec, optional) - At most one of existing_pipeline_id and new_pipeline_spec should be defined.
  
  If new_pipeline_spec is defined, a new pipeline is created for this synced table. The location pointed to is used
  to store intermediate files (checkpoints, event logs etc). The caller must have write permissions to create Delta
  tables in the specified catalog and schema. Again, note this requires write permissions, whereas the source table
  only requires read permissions
* `primary_key_columns` (list of string, optional) - Primary Key columns to be used for data insert/update in the destination
* `scheduling_policy` (string, optional) - Scheduling policy of the underlying pipeline. Possible values are: `CONTINUOUS`, `SNAPSHOT`, `TRIGGERED`
* `source_table_full_name` (string, optional) - Three-part (catalog, schema, table) name of the source Delta table
* `timeseries_key` (string, optional) - Time series key to deduplicate (tie-break) rows with the same primary key

### SyncedTableStatus
* `continuous_update_status` (SyncedTableContinuousUpdateStatus, optional)
* `failed_status` (SyncedTableFailedStatus, optional)
* `provisioning_status` (SyncedTableProvisioningStatus, optional)
* `triggered_update_status` (SyncedTableTriggeredUpdateStatus, optional)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `data_synchronization_status` (SyncedTableStatus) - Synced Table data synchronization status
* `effective_database_instance_name` (string) - The name of the database instance that this table is registered to. This field is always returned, and for
  tables inside database catalogs is inferred database instance associated with the catalog
* `effective_logical_database_name` (string) - The name of the logical database that this table is registered to
* `table_serving_url` (string) - Data serving REST API URL for this table
* `unity_catalog_provisioning_state` (string) - The provisioning state of the synced table entity in Unity Catalog. This is distinct from the
  state of the data synchronization pipeline (i.e. the table may be in "ACTIVE" but the pipeline
  may be in "PROVISIONING" as it runs asynchronously). Possible values are: `ACTIVE`, `DEGRADED`, `DELETING`, `FAILED`, `PROVISIONING`, `UPDATING`

### DeltaTableSyncInfo
* `delta_commit_timestamp` (string) - The timestamp when the above Delta version was committed in the source Delta table.
  Note: This is the Delta commit time, not the time the data was written to the synced table
* `delta_commit_version` (integer) - The Delta Lake commit version that was last successfully synced

### SyncedTableContinuousUpdateStatus
* `initial_pipeline_sync_progress` (SyncedTablePipelineProgress) - Progress of the initial data synchronization
* `last_processed_commit_version` (integer) - The last source table Delta version that was successfully synced to the synced table
* `timestamp` (string) - The end timestamp of the last time any data was synchronized from the source table to the synced
  table. This is when the data is available in the synced table

### SyncedTableFailedStatus
* `last_processed_commit_version` (integer) - The last source table Delta version that was successfully synced to the synced table.
  The last source table Delta version that was synced to the synced table.
  Only populated if the table is still
  synced and available for serving
* `timestamp` (string) - The end timestamp of the last time any data was synchronized from the source table to the synced
  table. Only populated if the table is still synced and available for serving

### SyncedTablePipelineProgress
* `estimated_completion_time_seconds` (number) - The estimated time remaining to complete this update in seconds
* `latest_version_currently_processing` (integer) - The source table Delta version that was last processed by the pipeline. The pipeline may not
  have completely processed this version yet
* `provisioning_phase` (string) - The current phase of the data synchronization pipeline. Possible values are: `PROVISIONING_PHASE_INDEX_SCAN`, `PROVISIONING_PHASE_INDEX_SORT`, `PROVISIONING_PHASE_MAIN`
* `sync_progress_completion` (number) - The completion ratio of this update. This is a number between 0 and 1
* `synced_row_count` (integer) - The number of rows that have been synced in this update
* `total_row_count` (integer) - The total number of rows that need to be synced in this update. This number may be an estimate

### SyncedTablePosition
* `delta_table_sync_info` (DeltaTableSyncInfo)
* `sync_end_timestamp` (string) - The end timestamp of the most recent successful synchronization.
  This is the time when the data is available in the synced table
* `sync_start_timestamp` (string) - The starting timestamp of the most recent successful synchronization from the source table
  to the destination (synced) table.
  Note this is the starting timestamp of the sync operation, not the end time.
  E.g., for a batch, this is the time when the sync operation started

### SyncedTableProvisioningStatus
* `initial_pipeline_sync_progress` (SyncedTablePipelineProgress) - Details about initial data synchronization. Only populated when in the
  PROVISIONING_INITIAL_SNAPSHOT state

### SyncedTableStatus
* `detailed_state` (string) - The state of the synced table. Possible values are: `SYNCED_TABLED_OFFLINE`, `SYNCED_TABLE_OFFLINE_FAILED`, `SYNCED_TABLE_ONLINE`, `SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE`, `SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE`, `SYNCED_TABLE_ONLINE_PIPELINE_FAILED`, `SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE`, `SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES`, `SYNCED_TABLE_PROVISIONING`, `SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT`, `SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES`
* `last_sync` (SyncedTablePosition) - Summary of the last successful synchronization from source to destination.
  
  Will always be present if there has been a successful sync. Even if the most recent syncs have failed.
  
  Limitation:
  The only exception is if the synced table is doing a FULL REFRESH, then the last sync information
  will not be available until the full refresh is complete. This limitation will be addressed in a future version.
  
  This top-level field is a convenience for consumers who want easy access to last sync information
  without having to traverse detailed_status
* `message` (string) - A text description of the current state of the synced table
* `pipeline_id` (string) - ID of the associated pipeline. The pipeline ID may have been provided by the client
  (in the case of bin packing), or generated by the server (when creating a new pipeline)

### SyncedTableTriggeredUpdateStatus
* `last_processed_commit_version` (integer) - The last source table Delta version that was successfully synced to the synced table
* `timestamp` (string) - The end timestamp of the last time any data was synchronized from the source table to the synced
  table. This is when the data is available in the synced table
* `triggered_update_progress` (SyncedTablePipelineProgress) - Progress of the active data synchronization pipeline

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_database_synced_database_table.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_database_synced_database_table "name"
```