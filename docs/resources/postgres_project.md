---
subcategory: "Postgres"
---
# databricks_postgres_project Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `project_id` (string, required) - The ID to use for the Project, which will become the final component of
  the project's resource name.
  
  This value should be 4-63 characters, and valid characters are /[a-z][0-9]-/
* `spec` (ProjectSpec, optional) - The desired state of a Project

### ProjectDefaultEndpointSettings
* `autoscaling_limit_max_cu` (number, optional) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number, optional) - The minimum number of Compute Units
* `pg_settings` (object, optional) - A raw representation of Postgres settings
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the compute endpoint is automatically suspended

### ProjectSettings
* `enable_logical_replication` (boolean, optional) - Sets wal_level=logical for all compute endpoints in this project.
  All active endpoints will be suspended.
  Once enabled, logical replication cannot be disabled

### ProjectSpec
* `default_endpoint_settings` (ProjectDefaultEndpointSettings, optional)
* `display_name` (string, optional) - Human-readable project name
* `history_retention_duration` (string, optional) - The number of seconds to retain the shared history for point in time recovery for all branches in this project
* `pg_version` (integer, optional) - The major Postgres version number
* `settings` (ProjectSettings, optional)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the project was created
* `name` (string) - The resource name of the project.
  Format: projects/{project_id}
* `status` (ProjectStatus) - The current status of a Project
* `uid` (string) - System generated unique ID for the project
* `update_time` (string) - A timestamp indicating when the project was last updated

### ProjectStatus
* `branch_logical_size_limit_bytes` (integer) - The logical size limit for a branch
* `compute_last_active_time` (string) - The most recent time when any endpoint of this project was active
* `default_endpoint_settings` (ProjectDefaultEndpointSettings) - The effective default endpoint settings
* `display_name` (string) - The effective human-readable project name
* `history_retention_duration` (string) - The effective number of seconds to retain the shared history for point in time recovery
* `owner` (string) - The email of the project owner
* `pg_version` (integer) - The effective major Postgres version number
* `settings` (ProjectSettings) - The effective project settings
* `synthetic_storage_size_bytes` (integer) - The current space occupied by the project in storage

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_project.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_project.this "name"
```