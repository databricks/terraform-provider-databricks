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
* `default` (boolean, optional) - Whether the branch is the project's default branch. This field is only returned on create/update responses.
  See effective_default for the value that is actually applied to the branch
* `is_protected` (boolean, optional) - Whether the branch is protected
* `source_branch` (string, optional) - The name of the source branch from which this branch was created.
  Format: projects/{project_id}/branches/{branch_id}
* `source_branch_lsn` (string, optional) - The Log Sequence Number (LSN) on the source branch from which this branch was created
* `source_branch_time` (string, optional) - The point in time on the source branch from which this branch was created

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the branch was created
* `current_state` (string) - The branch's state, indicating if it is initializing, ready for use, or archived. Possible values are: `ARCHIVED`, `IMPORTING`, `INIT`, `READY`, `RESETTING`
* `effective_default` (boolean) - Whether the branch is the project's default branch
* `effective_is_protected` (boolean) - Whether the branch is protected
* `effective_source_branch` (string) - The name of the source branch from which this branch was created.
  Format: projects/{project_id}/branches/{branch_id}
* `effective_source_branch_lsn` (string) - The Log Sequence Number (LSN) on the source branch from which this branch was created
* `effective_source_branch_time` (string) - The point in time on the source branch from which this branch was created
* `logical_size_bytes` (integer) - The logical size of the branch
* `name` (string) - The resource name of the branch.
  Format: projects/{project_id}/branches/{branch_id}
* `pending_state` (string) - The pending state of the branch, if a state transition is in progress. Possible values are: `ARCHIVED`, `IMPORTING`, `INIT`, `READY`, `RESETTING`
* `state_change_time` (string) - A timestamp indicating when the `current_state` began
* `uid` (string) - System generated unique ID for the branch
* `update_time` (string) - A timestamp indicating when the branch was last updated

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