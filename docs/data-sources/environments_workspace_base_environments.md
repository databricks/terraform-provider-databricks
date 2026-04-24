---
subcategory: "Environments"
---
# databricks_environments_workspace_base_environments Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source retrieves the list of all Workspace Base Environments in the workspace.
The list can be accessed via the data object's `workspace_base_environments` field.


## Example Usage
```hcl
data "databricks_environments_workspace_base_environments" "all" {}

output "all_environments" {
  value = data.databricks_environments_workspace_base_environments.all.workspace_base_environments
}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of environments to return per page.
  Default is 1000
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `workspace_base_environments`. It is a list of resources, each with the following attributes:
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