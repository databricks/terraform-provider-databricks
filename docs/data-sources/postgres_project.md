---
subcategory: "Postgres"
---
# databricks_postgres_project Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

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
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string) - A timestamp indicating when the project was created
* `delete_time` (string) - A timestamp indicating when the project was soft-deleted.
  Empty if the project is not deleted, otherwise set to a timestamp in the past
* `initial_branch_spec` (InitialBranchSpec) - Configuration for the initial default branch created as part of project creation.
  Allows overriding branch protection. These settings only apply at creation time
  and do not affect resources created after project creation
* `initial_endpoint_spec` (InitialEndpointSpec) - Configuration settings for the initial Read/Write endpoint created inside the initial branch for a newly
  created project. If omitted, the initial endpoint created will have default settings, without high availability
  configured. This field does not apply to any endpoints created after project creation. Use
  spec.default_endpoint_settings to configure default settings for endpoints created after project creation
* `name` (string) - Output only. The full resource path of the project.
  Format: projects/{project_id}
* `project_id` (string) - The project ID chosen when the project was created. It is the final segment of `name`
* `purge_time` (string) - A timestamp indicating when the project is scheduled for permanent deletion.
  Empty if the project is not deleted, otherwise set to a timestamp in the future
* `spec` (ProjectSpec) - The desired configuration of the project. The fields under `spec` are the ones that can be
  changed after creation, except `pg_version`, which is fixed when the project is created. See
  `status` for the configuration that is currently in effect.
  
  Read the resource's current values from `status`, rather than `spec`.
  `spec` tracks field ownership for infrastructure-as-code tools and may omit
  unmanaged fields
* `status` (ProjectStatus) - The configuration and state of the project that are currently in effect, as reported by the
  system
* `uid` (string) - A system-generated UUID that identifies the project. Unlike `name`, it is not derived from the
  user-chosen project ID
* `update_time` (string) - A timestamp indicating when the project was last updated

### EndpointGroupSpec
* `enable_readable_secondaries` (boolean) - Whether to allow read-only connections to read-write endpoints. Only relevant for read-write endpoints where
  `max` is greater than 1. When enabled, the endpoint also reports a read-only host in `status.hosts`
* `max` (integer) - The maximum number of computes in the endpoint group. Currently, this must be equal to min. Set to 1 for single
  compute endpoints, to disable HA. To manually suspend all computes in an endpoint group, set disabled to
  true on the EndpointSpec
* `min` (integer) - The minimum number of computes in the endpoint group. Currently, this must be equal to max. This must be greater
  than or equal to 1

### InitialBranchSpec
* `is_protected` (boolean) - Whether the initial default branch should be protected from deletion

### InitialEndpointSpec
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units for the initial endpoint. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units for the initial endpoint. Minimum value is 0.5, and it must
  not exceed autoscaling_limit_max_cu
* `group` (EndpointGroupSpec) - Settings for HA configuration of the endpoint
* `no_suspension` (boolean) - When set to true, explicitly disables automatic suspension (never suspend).
  Should be set to true when provided.
  Mutually exclusive with `suspend_timeout_duration`
* `suspend_timeout_duration` (string) - Duration of inactivity after which the initial endpoint is automatically suspended.
  If specified, should be between 60s and 604800s (1 minute to 1 week).
  Mutually exclusive with `no_suspension`

### ProjectCustomTag
* `key` (string) - The key of the custom tag
* `value` (string) - The value of the custom tag

### ProjectDefaultEndpointSettings
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units. Minimum value is 0.5
* `no_suspension` (boolean) - When set to true, explicitly disables automatic suspension (never suspend).
  Should be set to true when provided.
  Mutually exclusive with `suspend_timeout_duration`
* `pg_settings` (object) - A raw representation of Postgres settings
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended.
  If specified should be between 60s and 604800s (1 minute to 1 week).
  Mutually exclusive with `no_suspension`

### ProjectSpec
* `budget_policy_id` (string) - The desired budget policy to associate with the project.
  See status.budget_policy_id for the policy that is actually applied to the project
* `custom_tags` (list of ProjectCustomTag) - Custom tags to associate with the project, used for billing and cost tracking.
  To update tags, provide the new tag list and include "spec.custom_tags" in the update_mask.
  To clear all tags, provide an empty list and include "spec.custom_tags" in the update_mask.
  To preserve existing tags, omit this field from the update_mask (or use wildcard "*" which auto-excludes empty tags)
* `default_branch` (string) - The full resource path of the default branch of the project.
  Format: projects/{project_id}/branches/{branch_id}
  At creation time, the branch ID in this path names the initial branch, which is called
  `production` when the field is omitted. On update, the path must name an existing branch of this
  project, which then becomes the default branch
* `default_endpoint_settings` (ProjectDefaultEndpointSettings) - Settings applied by default to compute endpoints created in this project. The initial endpoint
  also takes these values for any setting that initial_endpoint_spec does not provide
* `display_name` (string) - Human-readable project name. Length should be between 1 and 256 characters
* `enable_pg_native_login` (boolean) - Whether to enable PG native password login on all endpoints in this project. Defaults to false
* `history_retention_duration` (string) - The number of seconds to retain the shared history for point in time recovery for all branches in this project. Value should be between 172800s (2 days) and 3024000s (35 days).
  Defaults to 604800s (7 days) when omitted
* `pg_version` (integer) - The major Postgres version number. Defaults to 17 when omitted. The set of versions that new
  projects can use may change over time

### ProjectStatus
* `branch_logical_size_limit_bytes` (integer) - The logical size limit for a branch
* `budget_policy_id` (string) - The budget policy that is applied to the project
* `compute_last_active_time` (string) - The most recent time when any endpoint of this project was active
* `custom_tags` (list of ProjectCustomTag) - The effective custom tags associated with the project
* `default_branch` (string) - The full resource path of the default branch of the project
* `default_endpoint_settings` (ProjectDefaultEndpointSettings) - The effective default endpoint settings
* `display_name` (string) - The effective human-readable project name
* `enable_pg_native_login` (boolean) - Whether to enable PG native password login on all endpoints in this project
* `history_retention_duration` (string) - The effective number of seconds to retain the shared history for point in time recovery
* `owner` (string) - The email of the project owner
* `pg_version` (integer) - The effective major Postgres version number
* `project_id` (string) - The project ID, which is the final segment of the project's resource name
* `synthetic_storage_size_bytes` (integer) - The current space occupied by the project in storage