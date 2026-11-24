---
subcategory: "Disaster Recovery"
---
# databricks_disaster_recovery_failover_groups Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/disasterrecovery)

This data source can be used to fetch the list of failover groups in the account.

-> **Note** This data source can only be used with an account-level provider!


## Example Usage
Getting a list of all failover groups in the account:

```hcl
data "databricks_disaster_recovery_failover_groups" "all" {
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The parent resource. Format: accounts/{account_id}
* `page_size` (integer, optional) - Maximum number of failover groups to return per page:
  - when set to a value greater than 0, the page length is the minimum of this value
  and a server configured value;
  - when set to 0 or unset, the page length is set to a server configured value
  (recommended);
  - when set to a value less than 0, an invalid parameter error is returned


## Attributes
This data source exports a single attribute, `failover_groups`. It is a list of resources, each with the following attributes:
* `create_time` (string) - Time at which this failover group was created
* `effective_primary_region` (string) - Current effective primary region. Replication flows FROM workspaces in this region.
  Changes after a successful failover
* `etag` (string) - Opaque version string for optimistic locking. Server-generated and returned in responses
* `has_replication_point_blocking_errors` (boolean) - True when replication errors are preventing the replication point from advancing.
  Populated only while the failover group is INITIAL_REPLICATION or ACTIVE; the blocking
  errors are listed in the system.replication.states system table
* `initial_primary_region` (string) - Initial primary region. Used only in Create requests to set the starting
  primary region, and must be one of the regions in `regions`. Not returned in responses
* `initial_replication_progress` (InitialReplicationProgress) - Initial replication progress of all in-scope assets. Set only when the FailoverGroup is in INITIAL_REPLICATION state
* `name` (string) - Fully qualified resource name in the format
  accounts/{account_id}/failover-groups/{failover_group_id}
* `regions` (list of string) - The two cloud regions that participate in this failover group, the primary and the
  secondary. Each workspace set contains one workspace in each region
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
* `unity_catalog_assets` (UcReplicationConfig) - Unity Catalog replication configuration: the catalogs to replicate and the storage
  location mappings for their data. Omit it to replicate workspace assets only
* `update_time` (string) - Time at which this failover group was last modified
* `workspace_sets` (list of WorkspaceSet) - Workspace sets, each containing workspaces that replicate to each other. A failover group
  currently supports exactly one workspace set

### InitialReplicationProgress
* `completed` (integer) - Percentage of in-scope assets replicated so far, from 0 to 100. Stays below 100 while
  errors block the first replication point

### LocationMapping
* `name` (string) - Identifier of this location mapping, unique within the failover group
* `uri_by_region` (list of LocationMappingEntry) - The storage URI of this location in each region, with exactly one entry per region of the
  failover group

### LocationMappingEntry
* `region` (string) - Cloud region this URI applies to. Must be one of the failover group's regions
* `uri` (string) - Cloud storage URI of the location in this region. End the URI with `/*` to map every path
  under it as a prefix; within one mapping, either all URIs or none carry the `/*` suffix

### UcCatalog
* `name` (string) - Name of the Unity Catalog catalog to replicate

### UcReplicationConfig
* `catalogs` (list of UcCatalog) - Unity Catalog catalogs to replicate, including their metadata and managed table data.
  Adding a catalog to an existing failover group returns the group to INITIAL_REPLICATION
  until the catalog has been replicated
* `data_replication_workspace_set` (string) - Name of the workspace set whose workspaces run the data replication for the catalogs'
  underlying storage. Must match the name of a workspace set in this failover group
* `location_mappings` (list of LocationMapping) - Storage location mappings. Each mapping pairs a storage path used by the replicated
  catalogs in one region with the corresponding path in the other region, so that
  replicated data is written to the matching storage in the secondary region

### WorkspaceSet
* `name` (string) - Identifier of the workspace set, unique within the failover group and referenced by
  unity_catalog_assets.data_replication_workspace_set. Cannot be changed after creation
* `replicate_workspace_assets` (boolean) - Whether to replicate workspace assets (notebooks, jobs, clusters, SQL warehouses and so
  on) and their ACLs for this set. Defaults to false. Can be turned on or off with
  UpdateFailoverGroup
* `stable_url_names` (list of string) - Resource names of stable URLs associated with this workspace set.
  Format: accounts/{account_id}/stable-urls/{stable_url_id}.
  The referenced stable URLs must already exist (via CreateStableUrl), and the list can only
  be set when replicate_workspace_assets is true. While workspace asset replication stays
  enabled, an update can attach or detach one stable URL at a time
* `workspace_ids` (list of string) - IDs of the two workspaces in this set, one in each of the failover group's regions. The
  system derives and validates each workspace's region. Both workspaces must belong to the
  account and have the Mission Critical add-on enabled, and they cannot be changed after
  creation