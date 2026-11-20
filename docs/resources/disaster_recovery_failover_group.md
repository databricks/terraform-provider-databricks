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
  primary region, and must be one of the regions in `regions`. Not returned in responses
* `parent` (string, required) - The parent resource. Format: accounts/{account_id}
* `regions` (list of string, required) - The two cloud regions that participate in this failover group, the primary and the
  secondary. Each workspace set contains one workspace in each region
* `workspace_sets` (list of WorkspaceSet, required) - Workspace sets, each containing workspaces that replicate to each other. A failover group
  currently supports exactly one workspace set
* `unity_catalog_assets` (UcReplicationConfig, optional) - Unity Catalog replication configuration: the catalogs to replicate and the storage
  location mappings for their data. Omit it to replicate workspace assets only

### LocationMapping
* `name` (string, required) - Identifier of this location mapping, unique within the failover group
* `uri_by_region` (list of LocationMappingEntry, required) - The storage URI of this location in each region, with exactly one entry per region of the
  failover group

### LocationMappingEntry
* `region` (string, required) - Cloud region this URI applies to. Must be one of the failover group's regions
* `uri` (string, required) - Cloud storage URI of the location in this region. End the URI with `/*` to map every path
  under it as a prefix; within one mapping, either all URIs or none carry the `/*` suffix

### UcCatalog
* `name` (string, required) - Name of the Unity Catalog catalog to replicate

### UcReplicationConfig
* `catalogs` (list of UcCatalog, required) - Unity Catalog catalogs to replicate, including their metadata and managed table data.
  Adding a catalog to an existing failover group returns the group to INITIAL_REPLICATION
  until the catalog has been replicated
* `data_replication_workspace_set` (string, required) - Name of the workspace set whose workspaces run the data replication for the catalogs'
  underlying storage. Must match the name of a workspace set in this failover group
* `location_mappings` (list of LocationMapping, optional) - Storage location mappings. Each mapping pairs a storage path used by the replicated
  catalogs in one region with the corresponding path in the other region, so that
  replicated data is written to the matching storage in the secondary region

### WorkspaceSet
* `name` (string, required) - Identifier of the workspace set, unique within the failover group and referenced by
  unity_catalog_assets.data_replication_workspace_set. Cannot be changed after creation
* `workspace_ids` (list of string, required) - IDs of the two workspaces in this set, one in each of the failover group's regions. The
  system derives and validates each workspace's region. Both workspaces must belong to the
  account and have the Mission Critical add-on enabled, and they cannot be changed after
  creation
* `replicate_workspace_assets` (boolean, optional) - Whether to replicate workspace assets (notebooks, jobs, clusters, SQL warehouses and so
  on) and their ACLs for this set. Defaults to false. Can be turned on or off with
  UpdateFailoverGroup
* `stable_url_names` (list of string, optional) - Resource names of stable URLs associated with this workspace set.
  Format: accounts/{account_id}/stable-urls/{stable_url_id}.
  The referenced stable URLs must already exist (via CreateStableUrl), and the list can only
  be set when replicate_workspace_assets is true. While workspace asset replication stays
  enabled, an update can attach or detach one stable URL at a time

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Time at which this failover group was created
* `effective_primary_region` (string) - Current effective primary region. Replication flows FROM workspaces in this region.
  Changes after a successful failover
* `etag` (string) - Opaque version string for optimistic locking. Server-generated and returned in responses
* `has_replication_point_blocking_errors` (boolean) - True when replication errors are preventing the replication point from advancing.
  Populated only while the failover group is INITIAL_REPLICATION or ACTIVE; the blocking
  errors are listed in the system.replication.states system table
* `initial_replication_progress` (InitialReplicationProgress) - Initial replication progress of all in-scope assets. Set only when the FailoverGroup is in INITIAL_REPLICATION state
* `name` (string) - Fully qualified resource name in the format
  accounts/{account_id}/failover-groups/{failover_group_id}
* `replication_point` (string) - The most recent point in time up to which all in-scope assets have been replicated to the
  secondary region. Data written after this point might not exist in the secondary yet and
  can be lost in an unplanned failover. When both Unity Catalog and workspace assets are
  replicated, this is the older of the two replication points. Not populated until the
  first replication point exists, and not populated in ListFailoverGroups responses
* `state` (string) - Aggregate state of the failover group across Unity Catalog and workspace asset
  replication. The group is INITIAL_REPLICATION until it has a replication point, ACTIVE
  while replication is in steady state and failover is available, and FAILING_OVER while a
  failover is in progress; the `*_FAILED` values mean that the last operation did not
  complete. Not populated in ListFailoverGroups responses. Possible values are: `ACTIVE`, `CREATING`, `CREATION_FAILED`, `DELETING`, `DELETION_FAILED`, `FAILING_OVER`, `FAILOVER_FAILED`, `INITIAL_REPLICATION`
* `update_time` (string) - Time at which this failover group was last modified

### InitialReplicationProgress
* `completed` (integer) - Percentage of in-scope assets replicated so far, from 0 to 100. Stays below 100 while
  errors block the first replication point

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