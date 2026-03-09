---
subcategory: "Postgres"
---
# databricks_postgres_branch Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source retrieves a single Postgres branch.


## Example Usage
### Retrieve Branch by Name

```hcl
data "databricks_postgres_branch" "this" {
  name = "projects/my-project/branches/dev-branch"
}

output "branch_is_protected" {
  value = data.databricks_postgres_branch.this.status.is_protected
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Output only. The full resource path of the branch.
  Format: projects/{project_id}/branches/{branch_id}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string) - A timestamp indicating when the branch was created
* `name` (string) - Output only. The full resource path of the branch.
  Format: projects/{project_id}/branches/{branch_id}
* `parent` (string) - The project containing this branch (API resource hierarchy).
  Format: projects/{project_id}
  
  Note: This field indicates where the branch exists in the resource hierarchy.
  For point-in-time branching from another branch, see `status.source_branch`
* `spec` (BranchSpec) - The spec contains the branch configuration
* `status` (BranchStatus) - The current status of a Branch
* `uid` (string) - System-generated unique ID for the branch
* `update_time` (string) - A timestamp indicating when the branch was last updated

### BranchSpec
* `expire_time` (string) - Absolute expiration timestamp. When set, the branch will expire at this time
* `is_protected` (boolean) - When set to true, protects the branch from deletion and reset. Associated compute endpoints and the project cannot be deleted while the branch is protected
* `no_expiry` (boolean) - Explicitly disable expiration. When set to true, the branch will not expire.
  If set to false, the request is invalid; provide either ttl or expire_time instead
* `source_branch` (string) - The name of the source branch from which this branch was created (data lineage for point-in-time recovery).
  If not specified, defaults to the project's default branch.
  Format: projects/{project_id}/branches/{branch_id}
* `source_branch_lsn` (string) - The Log Sequence Number (LSN) on the source branch from which this branch was created
* `source_branch_time` (string) - The point in time on the source branch from which this branch was created
* `ttl` (string) - Relative time-to-live duration. When set, the branch will expire at creation_time + ttl

### BranchStatus
* `current_state` (string) - The branch's state, indicating if it is initializing, ready for use, or archived. Possible values are: `ARCHIVED`, `IMPORTING`, `INIT`, `READY`, `RESETTING`
* `default` (boolean) - Whether the branch is the project's default branch
* `expire_time` (string) - Absolute expiration time for the branch. Empty if expiration is disabled
* `is_protected` (boolean) - Whether the branch is protected
* `logical_size_bytes` (integer) - The logical size of the branch
* `pending_state` (string) - The pending state of the branch, if a state transition is in progress. Possible values are: `ARCHIVED`, `IMPORTING`, `INIT`, `READY`, `RESETTING`
* `source_branch` (string) - The name of the source branch from which this branch was created.
  Format: projects/{project_id}/branches/{branch_id}
* `source_branch_lsn` (string) - The Log Sequence Number (LSN) on the source branch from which this branch was created
* `source_branch_time` (string) - The point in time on the source branch from which this branch was created
* `state_change_time` (string) - A timestamp indicating when the `current_state` began