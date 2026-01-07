---
subcategory: "Postgres"
---
# databricks_postgres_projects Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Upper bound for items returned


## Attributes
This data source exports a single attribute, `projects`. It is a list of resources, each with the following attributes:
* `create_time` (string) - A timestamp indicating when the project was created
* `name` (string) - The resource name of the project.
  Format: projects/{project_id}
* `spec` (ProjectSpec) - The desired state of a Project
* `status` (ProjectStatus) - The current status of a Project
* `uid` (string) - System generated unique ID for the project
* `update_time` (string) - A timestamp indicating when the project was last updated

### ProjectDefaultEndpointSettings
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units
* `pg_settings` (object) - A raw representation of Postgres settings
* `pgbouncer_settings` (object) - A raw representation of PgBouncer settings
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended

### ProjectSettings
* `enable_logical_replication` (boolean) - Sets wal_level=logical for all compute endpoints in this project.
  All active endpoints will be suspended.
  Once enabled, logical replication cannot be disabled

### ProjectSpec
* `default_endpoint_settings` (ProjectDefaultEndpointSettings)
* `display_name` (string) - Human-readable project name
* `history_retention_duration` (string) - The number of seconds to retain the shared history for point in time recovery for all branches in this project
* `pg_version` (integer) - The major Postgres version number
* `settings` (ProjectSettings)

### ProjectStatus
* `branch_logical_size_limit_bytes` (integer) - The logical size limit for a branch
* `compute_last_active_time` (string) - The most recent time when any endpoint of this project was active
* `default_endpoint_settings` (ProjectDefaultEndpointSettings) - The effective default endpoint settings
* `display_name` (string) - The effective human-readable project name
* `history_retention_duration` (string) - The effective number of seconds to retain the shared history for point in time recovery
* `pg_version` (integer) - The effective major Postgres version number
* `settings` (ProjectSettings) - The effective project settings
* `synthetic_storage_size_bytes` (integer) - The current space occupied by the project in storage