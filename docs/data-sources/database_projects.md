---
subcategory: "Database Instances"
---
# databricks_database_projects Data Source
[![Development](https://img.shields.io/badge/Release_Stage-Development-red)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Upper bound for items returned
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource


## Attributes
This data source exports a single attribute, `database_projects`. It is a list of resources, each with the following attributes:
* `branch_logical_size_limit_bytes` (integer) - The logical size limit for a branch
* `budget_policy_id` (string) - The desired budget policy to associate with the instance. This field is only returned on create/update responses,
  and represents the customer provided budget policy. See effective_budget_policy_id for the policy that is actually
  applied to the instance
* `compute_last_active_time` (string) - The most recent time when any endpoint of this project was active
* `create_time` (string) - A timestamp indicating when the project was created
* `custom_tags` (list of DatabaseProjectCustomTag) - Custom tags associated with the instance
* `default_endpoint_settings` (DatabaseProjectDefaultEndpointSettings)
* `display_name` (string) - Human-readable project name
* `effective_budget_policy_id` (string) - The policy that is applied to the instance
* `history_retention_duration` (string) - The number of seconds to retain the shared history for point in time recovery for all branches in this project
* `pg_version` (integer) - The major Postgres version number.
  NOTE: fields could be either user-set or server-set.
  we can't have fields that are optionally user-provided and server-set to default value.
  TODO: this needs an effective variant or make REQUIRED
* `project_id` (string)
* `settings` (DatabaseProjectSettings)
* `synthetic_storage_size_bytes` (integer) - The current space occupied by the project in storage. Synthetic storage size combines the logical data size and Write-Ahead Log (WAL) size for all branches in a project
* `update_time` (string) - A timestamp indicating when the project was last updated

### DatabaseProjectCustomTag
* `key` (string) - The key of the custom tag
* `value` (string) - The value of the custom tag

### DatabaseProjectDefaultEndpointSettings
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units
* `pg_settings` (object) - A raw representation of Postgres settings
* `pgbouncer_settings` (object) - A raw representation of PgBouncer settings
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended

### DatabaseProjectSettings
* `enable_logical_replication` (boolean) - Sets wal_level=logical for all compute endpoints in this project.
  All active endpoints will be suspended.
  Once enabled, logical replication cannot be disabled