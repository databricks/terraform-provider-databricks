---
subcategory: "Environments"
---
# databricks_environments_workspace_base_environment Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source retrieves a single Workspace Base Environment by its resource name.


## Example Usage
```hcl
data "databricks_environments_workspace_base_environment" "my_env" {
  name = "workspace-base-environments/my-environment"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the workspace base environment.
  Format: workspace-base-environments/{workspace-base-environment}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `base_environment_type` (string) - The type of base environment (CPU or GPU). Possible values are: `CPU`, `GPU`
* `create_time` (string) - Timestamp when the environment was created
* `creator_user_id` (string) - User ID of the creator
* `display_name` (string) - Human-readable display name for the workspace base environment
* `filepath` (string) - The WSFS or UC Volumes path to the environment YAML file
* `is_default` (boolean) - Whether this is the default environment for the workspace
* `last_updated_user_id` (string) - User ID of the last user who updated the environment
* `message` (string) - Status message providing additional details about the environment status
* `name` (string) - The resource name of the workspace base environment.
  Format: workspace-base-environments/{workspace-base-environment}
* `status` (string) - The status of the materialized workspace base environment. Possible values are: `CREATED`, `EXPIRED`, `FAILED`, `INVALID`, `PENDING`, `REFRESHING`
* `update_time` (string) - Timestamp when the environment was last updated