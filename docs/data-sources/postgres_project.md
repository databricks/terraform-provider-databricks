---
subcategory: "Postgres"
---
# databricks_postgres_project Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source retrieves a single Postgres project.


## Example Usage
### Retrieve Project by Name

```hcl
data "databricks_postgres_project" "this" {
  name = "projects/my-project"
}

output "project_pg_version" {
  value = data.databricks_postgres_project.this.status.pg_version
}

output "project_display_name" {
  value = data.databricks_postgres_project.this.status.display_name
}

output "project_history_retention" {
  value = data.databricks_postgres_project.this.status.history_retention_duration
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Output only. The full resource path of the project.
  Format: projects/{project_id}

## Attributes
The following attributes are exported:
* `create_time` (string) - A timestamp indicating when the project was created
* `name` (string) - Output only. The full resource path of the project.
  Format: projects/{project_id}
* `spec` (ProjectSpec) - The spec contains the project configuration, including display_name, pg_version (Postgres version), history_retention_duration, and default_endpoint_settings
* `status` (ProjectStatus) - The current status of a Project
* `uid` (string) - System-generated unique ID for the project
* `update_time` (string) - A timestamp indicating when the project was last updated

### ProjectDefaultEndpointSettings
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units. Minimum value is 0.5
* `no_suspension` (boolean) - When set to true, explicitly disables automatic suspension (never suspend).
  Should be set to true when provided
* `pg_settings` (object) - A raw representation of Postgres settings
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended.
  If specified should be between 60s and 604800s (1 minute to 1 week)

### ProjectSpec
* `default_endpoint_settings` (ProjectDefaultEndpointSettings)
* `display_name` (string) - Human-readable project name. Length should be between 1 and 256 characters
* `history_retention_duration` (string) - The number of seconds to retain the shared history for point in time recovery for all branches in this project. Value should be between 0s and 2592000s (up to 30 days)
* `pg_version` (integer) - The major Postgres version number. Supported versions are 16 and 17

### ProjectStatus
* `branch_logical_size_limit_bytes` (integer) - The logical size limit for a branch
* `default_endpoint_settings` (ProjectDefaultEndpointSettings) - The effective default endpoint settings
* `display_name` (string) - The effective human-readable project name
* `history_retention_duration` (string) - The effective number of seconds to retain the shared history for point in time recovery
* `owner` (string) - The email of the project owner
* `pg_version` (integer) - The effective major Postgres version number
* `synthetic_storage_size_bytes` (integer) - The current space occupied by the project in storage