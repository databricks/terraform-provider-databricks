---
subcategory: "Environments"
---
# databricks_environments_default_workspace_base_environment Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source retrieves the current default Workspace Base Environment configuration for the workspace.


## Example Usage
```hcl
data "databricks_environments_default_workspace_base_environment" "current" {
  name = "default-workspace-base-environment"
}

output "default_cpu_environment" {
  value = data.databricks_environments_default_workspace_base_environment.current.cpu_workspace_base_environment
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of this singleton resource.
  Format: default-workspace-base-environment
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `cpu_workspace_base_environment` (string) - The default workspace base environment for CPU compute.
  Format: workspace-base-environments/{workspace_base_environment}
* `gpu_workspace_base_environment` (string) - The default workspace base environment for GPU compute.
  Format: workspace-base-environments/{workspace_base_environment}
* `name` (string) - The resource name of this singleton resource.
  Format: default-workspace-base-environment