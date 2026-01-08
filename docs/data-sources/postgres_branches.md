---
subcategory: "Postgres"
---
# databricks_postgres_branches Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `parent` (string, required) - The Project that owns this collection of branches.
  Format: projects/{project_id}
* `page_size` (integer, optional) - Upper bound for items returned


## Attributes
This data source exports a single attribute, `branches`. It is a list of resources, each with the following attributes:
* `create_time` (string) - A timestamp indicating when the branch was created
* `name` (string) - The resource name of the branch.
  Format: projects/{project_id}/branches/{branch_id}
* `parent` (string) - The project containing this branch.
  Format: projects/{project_id}
* `spec` (BranchSpec) - The desired state of a Branch
* `status` (BranchStatus) - The current status of a Branch
* `uid` (string) - System generated unique ID for the branch
* `update_time` (string) - A timestamp indicating when the branch was last updated

### BranchSpec
* `default` (boolean) - Whether the branch is the project's default branch
* `is_protected` (boolean) - Whether the branch is protected
* `source_branch` (string) - The name of the source branch from which this branch was created.
  Format: projects/{project_id}/branches/{branch_id}
* `source_branch_lsn` (string) - The Log Sequence Number (LSN) on the source branch from which this branch was created
* `source_branch_time` (string) - The point in time on the source branch from which this branch was created

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