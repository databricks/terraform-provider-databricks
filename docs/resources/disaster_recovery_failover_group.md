---
subcategory: "Disaster Recovery"
---
# databricks_disaster_recovery_failover_group Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/disasterrecovery)

A failover group coordinates Databricks Managed Disaster Recovery across one or more workspace sets, replicating data and (optionally) workspace assets from a primary region to a secondary region so you can fail over with minimal disruption.

Each workspace set groups the workspaces that replicate to each other across regions. Unity Catalog catalogs and their underlying storage can additionally be replicated by configuring `unity_catalog_assets`. After a successful failover, the group's effective primary region changes to the former secondary.

-> **Note** This resource can only be used with an account-level provider!


## Example Usage
Creating a failover group for a single workspace set replicating between two regions, with control plane (workspace asset) replication enabled:

```hcl
resource "databricks_disaster_recovery_failover_group" "this" {
  failover_group_id      = "accounting-failover-group"
  initial_primary_region = "us-east-1"
  regions                = ["us-east-1", "us-west-2"]

  workspace_sets = [{
    name                       = "accounting"
    workspace_ids              = ["1234567890123456", "6543210987654321"]
    replicate_workspace_assets = true
  }]
}
```


## Arguments
The following arguments are supported:
* `failover_group_id` (string, required) - Client-provided identifier for the failover group. Used to construct the
  resource name as {parent}/failover-groups/{failover_group_id}
* `initial_primary_region` (string, required) - Initial primary region. Used only in Create requests to set the starting
  primary region. Not returned in responses
* `parent` (string, required) - The parent resource. Format: accounts/{account_id}
* `regions` (list of string, required) - List of all regions participating in this failover group
* `workspace_sets` (list of WorkspaceSet, required) - Workspace sets, each containing workspaces that replicate to each other
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
* `workspace_ids` (list of string, required) - Workspace IDs in this set. The system derives and validates regions.
  All workspaces must be in the Mission Critical tier
* `replicate_workspace_assets` (boolean, optional) - Whether to enable control plane DR (notebooks, jobs, clusters, etc.) for this set.
  Defaults to false
* `stable_url_names` (list of string, optional) - Resource names of stable URLs associated with this workspace set.
  Format: accounts/{account_id}/stable-urls/{stable_url_id}.
  The referenced stable URLs must already exist (via CreateStableUrl)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Time at which this failover group was created
* `effective_primary_region` (string) - Current effective primary region. Replication flows FROM workspaces in this region.
  Changes after a successful failover
* `etag` (string) - Opaque version string for optimistic locking. Server-generated and returned in responses
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