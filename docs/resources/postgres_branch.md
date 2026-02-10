---
subcategory: "Postgres"
---
# databricks_postgres_branch Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

### Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics. For complete details on how spec/status fields work, drift detection behavior, and state management requirements, see the `databricks_postgres_project` resource documentation.

### Overview

A Postgres branch is an independent database environment within a project that shares storage with its parent branch through copy-on-write. Branches allow you to create isolated development and testing environments, test applications against realistic data sets, and perform point-in-time recovery operations.

### Hierarchy Context

Branches exist within the Lakebase Autoscaling resource hierarchy:
- Each **branch** belongs to a **project**
- A **branch** can contain multiple **endpoints**, **databases**, and **roles**
- Branches can be created from other branches for point-in-time recovery

### Use Cases

- **Development environments**: Create isolated branches for feature development without affecting production
- **Testing**: Spin up temporary branches for testing changes before applying to production
- **Point-in-time recovery**: Create a branch from a specific point in time on another branch
- **Data exploration**: Safely query and analyze data without risking production workloads


## Example Usage
### Basic Branch Creation

```hcl
resource "databricks_postgres_project" "this" {
  project_id = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Project"
  }
}

resource "databricks_postgres_branch" "dev" {
  branch_id = "dev-branch"
  parent    = databricks_postgres_project.this.name
  spec = {
    no_expiry = true
  }
}
```

### Protected Branch

```hcl
resource "databricks_postgres_branch" "production" {
  branch_id = "production"
  parent    = databricks_postgres_project.this.name
  spec = {
    is_protected = true
    no_expiry = true
  }
}
```

### Branch with Expiration (TTL)

```hcl
resource "databricks_postgres_branch" "temporary" {
  branch_id = "temp-feature-test"
  parent    = databricks_postgres_project.this.name
  spec = {
    ttl = "604800s"  # 7 days
  }
}
```


## Arguments
The following arguments are supported:
* `branch_id` (string, required) - The ID to use for the Branch. This becomes the final component of the branch's resource name.
  The ID is required and must be 1-63 characters long, start with a lowercase letter, and contain only lowercase letters, numbers, and hyphens.
  For example, `development` becomes `projects/my-app/branches/development`
* `parent` (string, required) - The project containing this branch (API resource hierarchy).
  Format: projects/{project_id}
  
  Note: This field indicates where the branch exists in the resource hierarchy.
  For point-in-time branching from another branch, see `status.source_branch`
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
* `name` (string) - Output only. The full resource path of the branch.
  Format: projects/{project_id}/branches/{branch_id}
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