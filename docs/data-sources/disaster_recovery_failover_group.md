---
subcategory: "Disaster Recovery"
---
# databricks_disaster_recovery_failover_group Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Fully qualified resource name in the format
  accounts/{account_id}/failover-groups/{failover_group_id}

## Attributes
The following attributes are exported:
* `create_time` (string) - Time at which this failover group was created
* `effective_primary_region` (string) - Current effective primary region. Replication flows FROM workspaces in this region.
  Changes after a successful failover
* `etag` (string) - Opaque version string for optimistic locking. Server-generated, returned in responses.
  Must be provided on Update requests to prevent concurrent modifications
* `initial_primary_region` (string) - Initial primary region. Used only in Create requests to set the starting
  primary region. Not returned in responses
* `name` (string) - Fully qualified resource name in the format
  accounts/{account_id}/failover-groups/{failover_group_id}
* `regions` (list of string) - List of all regions participating in this failover group
* `replication_point` (string) - The latest point in time to which data has been replicated
* `state` (string) - Aggregate state of the failover group. Possible values are: `ACTIVE`, `CREATING`, `CREATION_FAILED`, `DELETING`, `DELETION_FAILED`, `FAILING_OVER`, `FAILOVER_FAILED`, `INITIAL_REPLICATION`
* `unity_catalog_assets` (UcReplicationConfig) - Unity Catalog replication configuration
* `update_time` (string) - Time at which this failover group was last modified
* `workspace_sets` (list of WorkspaceSet) - Workspace sets, each containing workspaces that replicate to each other

### LocationMapping
* `name` (string) - Resource name for this location
* `uri_by_region` (list of LocationMappingEntry) - URI for each region. Each entry maps a region name to a storage URI

### LocationMappingEntry
* `region` (string) - The region name
* `uri` (string) - The storage URI for this region

### UcCatalog
* `name` (string) - The name of the UC catalog to replicate

### UcReplicationConfig
* `catalogs` (list of UcCatalog) - UC catalogs to replicate
* `data_replication_workspace_set` (string) - The workspace set whose workspaces will be used for data replication
  of all UC catalogs' underlying storage
* `location_mappings` (list of LocationMapping) - Location mappings - storage URI per region for each location

### WorkspaceSet
* `name` (string) - Resource name for this workspace set
* `replicate_workspace_assets` (boolean) - Whether to enable control plane DR (notebooks, jobs, clusters, etc.) for this set.
  Requires all workspaces in the set to be Mission Critical tier
* `stable_url_names` (list of string) - Resource names of stable URLs associated with this workspace set.
  Format: accounts/{account_id}/stable-urls/{stable_url_id}.
  The referenced stable URLs must already exist (via CreateStableUrl)
* `workspace_ids` (list of string) - Workspace IDs in this set. The system derives and validates regions.
  EA: exactly 2 workspaces (one per region)