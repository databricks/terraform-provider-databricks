---
subcategory: "Postgres"
---
# databricks_postgres_branch Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `branch_id` (string, required) - The ID to use for the Branch. This becomes the final component of the branch's resource name.
  The ID must be 1-63 characters long, start with a lowercase letter, and contain only lowercase letters, numbers, and hyphens (RFC 1123).
  Examples:
  - With custom ID: `staging` → name becomes `projects/{project_id}/branches/staging`
  - Without custom ID: system generates slug → name becomes `projects/{project_id}/branches/br-example-name-x1y2z3a4`
* `parent` (string, required) - The project containing this branch (API resource hierarchy).
  Format: projects/{project_id}
  
  Note: This field indicates where the branch exists in the resource hierarchy.
  For point-in-time branching from another branch, see `spec.source_branch`
* `spec` (BranchSpec, optional) - The spec contains the branch configuration

### BranchSpec
* `expire_time` (string, optional) - Absolute expiration timestamp. When set, the branch will expire at this time
* `is_protected` (boolean, optional) - When set to true, protects the branch from deletion and reset. Associated compute endpoints and the project cannot be deleted while the branch is protected
* `no_expiry` (boolean, optional) - Explicitly disable expiration. When set to true, the branch will not expire.
  If set to false, the request is invalid; provide either ttl or expire_time instead
* `source_branch` (string, optional) - The name of the source branch from which this branch was created (data lineage for point-in-time recovery).
  If not specified, defaults to the project's default branch.
  Format: projects/{project_id}/branches/{branch_id}
* `source_branch_lsn` (string, optional) - The Log Sequence Number (LSN) on the source branch from which this branch was created
* `source_branch_time` (string, optional) - The point in time on the source branch from which this branch was created
* `ttl` (string, optional) - Relative time-to-live duration. When set, the branch will expire at creation_time + ttl

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the branch was created
* `name` (string) - The resource name of the branch. This field is output-only and constructed by the system.
  Format: `projects/{project_id}/branches/{branch_id}`
* `status` (BranchStatus) - The current status of a Branch
* `uid` (string) - System-generated unique ID for the branch
* `update_time` (string) - A timestamp indicating when the branch was last updated

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

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_branch.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_branch.this "name"
```