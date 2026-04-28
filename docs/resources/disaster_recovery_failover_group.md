---
subcategory: "Disaster Recovery"
---
# databricks_disaster_recovery_failover_group Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `failover_group_id` (string, required) - Client-provided identifier for the failover group. Used to construct the
  resource name as {parent}/failover-groups/{failover_group_id}
* `initial_primary_region` (string, required) - Initial primary region. Used only in Create requests to set the starting
  primary region. Not returned in responses
* `parent` (string, required) - The parent resource. Format: accounts/{account_id}
* `regions` (list of string, required) - List of all regions participating in this failover group
* `workspace_sets` (list of WorkspaceSet, required) - Workspace sets, each containing workspaces that replicate to each other
* `etag` (string, optional) - Opaque version string for optimistic locking. Server-generated, returned in responses.
  Must be provided on Update requests to prevent concurrent modifications
* `unity_catalog_assets` (UcReplicationConfig, optional) - Unity Catalog replication configuration

### LocationMapping
* `name` (string, required) - Resource name for this location
* `uri_by_region` (list of LocationMappingEntry, required) - URI for each region. Each entry maps a region name to a storage URI

### LocationMappingEntry
* `region` (string, required) - The region name
* `uri` (string, required) - The storage URI for this region

### UcCatalog
* `name` (string, required) - The name of the UC catalog to replicate

### UcReplicationConfig
* `catalogs` (list of UcCatalog, required) - UC catalogs to replicate
* `data_replication_workspace_set` (string, required) - The workspace set whose workspaces will be used for data replication
  of all UC catalogs' underlying storage
* `location_mappings` (list of LocationMapping, optional) - Location mappings - storage URI per region for each location

### WorkspaceSet
* `name` (string, required) - Resource name for this workspace set
* `replicate_workspace_assets` (boolean, required) - Whether to enable control plane DR (notebooks, jobs, clusters, etc.) for this set.
  Requires all workspaces in the set to be Mission Critical tier
* `workspace_ids` (list of string, required) - Workspace IDs in this set. The system derives and validates regions.
  EA: exactly 2 workspaces (one per region)
* `stable_url_names` (list of string, optional) - Resource names of stable URLs associated with this workspace set.
  Format: accounts/{account_id}/stable-urls/{stable_url_id}.
  The referenced stable URLs must already exist (via CreateStableUrl)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Time at which this failover group was created
* `effective_primary_region` (string) - Current effective primary region. Replication flows FROM workspaces in this region.
  Changes after a successful failover
* `name` (string) - Fully qualified resource name in the format
  accounts/{account_id}/failover-groups/{failover_group_id}
* `replication_point` (string) - The latest point in time to which data has been replicated
* `state` (string) - Aggregate state of the failover group. Possible values are: `ACTIVE`, `CREATING`, `CREATION_FAILED`, `DELETING`, `DELETION_FAILED`, `FAILING_OVER`, `FAILOVER_FAILED`, `INITIAL_REPLICATION`
* `update_time` (string) - Time at which this failover group was last modified

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_disaster_recovery_failover_group.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_disaster_recovery_failover_group.this "name"
```