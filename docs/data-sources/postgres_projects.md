---
subcategory: "Postgres"
---
# databricks_postgres_projects Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Upper bound for items returned. Cannot be negative


## Attributes
This data source exports a single attribute, `projects`. It is a list of resources, each with the following attributes:
* `create_time` (string) - A timestamp indicating when the project was created
* `name` (string) - The resource name of the project. This field is output-only and constructed by the system.
  Format: `projects/{project_id}`
* `spec` (ProjectSpec) - The spec contains the project configuration, including display_name, pg_version (Postgres version), history_retention_duration, and default_endpoint_settings
* `status` (ProjectStatus) - The current status of a Project
* `uid` (string) - System-generated unique ID for the project
* `update_time` (string) - A timestamp indicating when the project was last updated

### ProjectDefaultEndpointSettings
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units. Minimum value is 0.5
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