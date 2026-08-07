---
subcategory: "Postgres"
---
# databricks_postgres_synced_table Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

### Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics. For complete details on how spec/status fields work, drift detection behavior, and state management requirements, see the `databricks_postgres_project` resource documentation.

### Overview

A Postgres synced table replicates data from a Delta table in Unity Catalog into a Postgres table for low-latency serving. The synchronization is managed by an underlying Lakeflow pipeline.

### Hierarchy Context

Synced tables exist within the Lakebase Autoscaling resource hierarchy:
- A **synced table** belongs to a **catalog** within a **branch** and **project**
- Each synced table replicates exactly one source Delta table into Postgres
- The `synced_table_id` is the three-part Unity Catalog name: `catalog.schema.table`

### Use Cases

- **Low-latency serving**: Serve Delta table data through Postgres for millisecond read latency
- **Feature store**: Provide online feature serving for ML models backed by Delta tables
- **Reverse ETL**: Push lakehouse data into Postgres for consumption by external applications and APIs


## Example Usage
### Basic Synced Table with Snapshot Policy

```hcl
resource "databricks_postgres_project" "this" {
  project_id = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Project"
  }
}

resource "databricks_postgres_branch" "main" {
  branch_id = "main"
  parent    = databricks_postgres_project.this.name
  spec = {
    no_expiry = true
  }
}

resource "databricks_postgres_catalog" "this" {
  catalog_id = "app_catalog"
  spec = {
    postgres_database          = "app_db"
    create_database_if_missing = true
    branch                     = databricks_postgres_branch.main.name
  }
}

resource "databricks_postgres_synced_table" "this" {
  synced_table_id = "app_catalog.default.users_synced"
  spec = {
    source_table_full_name            = "main.default.users"
    primary_key_columns               = ["user_id"]
    scheduling_policy                 = "SNAPSHOT"
    postgres_database                 = "app_db"
    branch                            = databricks_postgres_branch.main.name
    create_database_objects_if_missing = true
    new_pipeline_spec = {
      storage_catalog = "main"
      storage_schema  = "default"
    }
  }
}
```

### Synced Table with Triggered Policy

Use `TRIGGERED` for on-demand updates. Requires Change Data Feed (CDF) enabled on the source table.

```hcl
resource "databricks_postgres_synced_table" "triggered" {
  synced_table_id = "app_catalog.default.orders_synced"
  spec = {
    source_table_full_name            = "main.default.orders"
    primary_key_columns               = ["order_id"]
    scheduling_policy                 = "TRIGGERED"
    postgres_database                 = "app_db"
    branch                            = databricks_postgres_branch.main.name
    create_database_objects_if_missing = true
    new_pipeline_spec = {
      storage_catalog = "main"
      storage_schema  = "default"
    }
  }
}
```

### Synced Table with Existing Pipeline

Bin-pack into an existing pipeline instead of creating a new one:

```hcl
resource "databricks_postgres_synced_table" "this" {
  synced_table_id = "app_catalog.default.products_synced"
  spec = {
    source_table_full_name            = "main.default.products"
    primary_key_columns               = ["product_id"]
    scheduling_policy                 = "SNAPSHOT"
    postgres_database                 = "app_db"
    branch                            = databricks_postgres_branch.main.name
    create_database_objects_if_missing = true
    existing_pipeline_id              = "abc123-def456"
  }
}
```


## Arguments
The following arguments are supported:
* `synced_table_id` (string, required) - The part of the name, chosen by the user when the resource was created
* `spec` (SyncedTableSyncedTableSpec, optional) - Configuration details of the synced table, such as the source table, scheduling policy, etc.
  This attribute is specified at creation time and most fields are returned as is on subsequent queries
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### NewPipelineSpec
* `budget_policy_id` (string, optional) - Budget policy to set on the newly created pipeline
* `storage_catalog` (string, optional) - UC catalog for the pipeline to store intermediate files (checkpoints, event logs etc).
  This needs to be a standard catalog where the user has permissions to create Delta tables
* `storage_schema` (string, optional) - UC schema for the pipeline to store intermediate files (checkpoints, event logs etc).
  This needs to be in the standard catalog where the user has permissions to create Delta tables

### SyncedTableSyncedTableSpec
* `accelerated_sync` (boolean, optional) - When true, enables accelerated sync mode for the initial data load.
  This significantly improves performance for large tables.
  Requires workspace-level enablement through Lakebase Accelerated Sync preview
* `branch` (string, optional) - The full resource name the branch associated with the table.
  
  Format: "projects/{project_id}/branches/{branch_id}"
* `create_database_objects_if_missing` (boolean, optional) - If true, the synced table's logical database and schema resources in PG
  will be created if they do not already exist.
  The request will fail if this is false and the database/schema do not exist.
  
  Defaults to true if omitted
* `existing_pipeline_id` (string, optional) - ID of an existing pipeline to bin-pack this synced table into.
  At most one of existing_pipeline_id and new_pipeline_spec should be defined.
  
  The pipeline used for the synced table is returned via the top level pipeline_id attribute
* `extra_columns` (list of SyncedTableSyncedTableSpecExtraColumn, optional) - Extra PostgreSQL-only columns to add to the synced table
* `new_pipeline_spec` (NewPipelineSpec, optional) - Specification for creating a new pipeline.
  At most one of existing_pipeline_id and new_pipeline_spec should be defined.
  
  The pipeline used for the synced table is returned via the top level pipeline_id attribute
* `postgres_database` (string, optional) - The Postgres database name where the synced table will be created in.
  
  If this synced table is created inside a Lakebase Catalog, this attribute can be omitted on creation and is inferred
  from the postgres_database associated with the Lakebase Catalog. If specified when inside a Lakebase Catalog, the value must match.
  
  A value must be specified when creating a synced table inside a Standard Catalog
* `primary_key_columns` (list of string, optional) - Primary Key columns to be used for data insert/update in the destination
* `scheduling_policy` (string, optional) - Scheduling policy of the underlying pipeline. Possible values are: `CONTINUOUS`, `SNAPSHOT`, `TRIGGERED`
* `source_table_full_name` (string, optional) - Three-part (catalog, schema, table) name of the source Delta table.
  
  For the corresponding destination table, use any of the two:
  
  * synced_table_id used at the creation of the SyncedTable
  * "name" consisting of "synced_tables/" prefix and the full name of the destination table
* `timeseries_key` (string, optional) - Time series key to deduplicate (tie-break) rows with the same primary key
* `type_overrides` (list of SyncedTableSyncedTableSpecTypeOverride, optional) - Override the default Delta->PG type mapping for specific columns.
  A TypeOverride with PG_SPECIFIC_TYPE_UNSPECIFIED is rejected; a valid pg_type must be set

### SyncedTableSyncedTableSpecExtraColumn
* `column_name` (string, required) - Name of the column
* `column_type` (string, required) - PostgreSQL type of the column, for example "tsvector" or "vector(1024)"
* `compute` (string, optional) - SQL expression used to compute the column's value, for example
  "to_tsvector('english', content)"
* `maintenance` (string, optional) - Possible values are: `STORED_GENERATED`

### SyncedTableSyncedTableSpecTypeOverride
* `column_name` (string, required) - Name of the source column whose target PostgreSQL type should be overridden
* `pg_type` (string, required) - PostgreSQL-specific target type to use for the column. Possible values are: `PG_SPECIFIC_TYPE_VECTOR`
* `size` (integer, optional) - Size parameter for the target type, for types that take one (e.g. vector
  dimension, varchar length). Required when the chosen pg_type needs a size

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string)
* `name` (string) - Output only. The Full resource name of the synced table in Postgres
  where (catalog, schema, table) are the UC entity names.
  
  Format "synced_tables/{catalog}.{schema}.{table}"
  
  For the corresponding source table in the Unity catalog look for the "source_table_full_name" attribute
* `status` (SyncedTableSyncedTableStatus) - Synced Table data synchronization status
* `uid` (string) - The Unity Catalog table ID for this synced table

### DeltaTableSyncInfo
* `delta_commit_time` (string) - The timestamp when the above Delta version was committed in the source Delta table.
  Note: This is the Delta commit time, not the time the data was written to the synced table
* `delta_commit_version` (integer) - The Delta Lake commit version that was last successfully synced

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

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_synced_table.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_synced_table.this "name"
```