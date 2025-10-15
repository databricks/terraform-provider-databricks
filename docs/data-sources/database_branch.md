---
subcategory: "Database Instances"
---
# databricks_database_branch Data Source
[![Development](https://img.shields.io/badge/Release_Stage-Development-red)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `branch_id` (string, required)
* `project_id` (string, required)
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource

## Attributes
The following attributes are exported:
* `branch_id` (string)
* `create_time` (string) - A timestamp indicating when the branch was created
* `current_state` (string) - The branch’s state, indicating if it is initializing, ready for use, or archived
* `default` (boolean) - Whether the branch is the project's default branch. This field is only returned on create/update responses.
  See effective_default for the value that is actually applied to the database branch
* `effective_default` (boolean) - Whether the branch is the project's default branch
* `logical_size_bytes` (integer) - The logical size of the branch
* `parent_id` (string) - The id of the parent branch
* `parent_lsn` (string) - The Log Sequence Number (LSN) on the parent branch from which this branch was created.
  When restoring a branch using the Restore Database Branch endpoint,
  this value isn’t finalized until all operations related to the restore have completed successfully
* `parent_time` (string) - The point in time on the parent branch from which this branch was created
* `pending_state` (string)
* `project_id` (string)
* `protected` (boolean) - Whether the branch is protected
* `state_change_time` (string) - A timestamp indicating when the `current_state` began
* `update_time` (string) - A timestamp indicating when the branch was last updated