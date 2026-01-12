---
subcategory: "Postgres"
---
# databricks_postgres_branch Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `parent` (string, required) - The project containing this branch.
  Format: projects/{project_id}
* `branch_id` (string, optional) - The ID to use for the Branch, which will become the final component of
  the branch's resource name.
  
  This value should be 4-63 characters, and valid characters are /[a-z][0-9]-/
* `spec` (BranchSpec, optional) - The desired state of a Branch

### BranchSpec
* `default` (boolean, optional) - Whether the branch is the project's default branch
* `is_protected` (boolean, optional) - Whether the branch is protected
* `source_branch` (string, optional) - The name of the source branch from which this branch was created.
  Format: projects/{project_id}/branches/{branch_id}
* `source_branch_lsn` (string, optional) - The Log Sequence Number (LSN) on the source branch from which this branch was created
* `source_branch_time` (string, optional) - The point in time on the source branch from which this branch was created

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the branch was created
* `name` (string) - The resource name of the branch.
  Format: projects/{project_id}/branches/{branch_id}
* `status` (BranchStatus) - The current status of a Branch
* `uid` (string) - System generated unique ID for the branch
* `update_time` (string) - A timestamp indicating when the branch was last updated

### BranchStatus
* `current_state` (string) - The branch's state, indicating if it is initializing, ready for use, or archived. Possible values are: `ARCHIVED`, `IMPORTING`, `INIT`, `READY`, `RESETTING`
* `default` (boolean) - Whether the branch is the project's default branch
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
