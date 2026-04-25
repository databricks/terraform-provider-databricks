---
subcategory: "Environments"
---
# databricks_environments_default_workspace_base_environment Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

The Default Workspace Base Environment is a singleton resource that configures which workspace base environments are applied by default to new notebooks in the workspace. Defaults can be set separately for CPU and GPU compute.

Without a default configured, new notebooks do not use a workspace base environment by default.

### Unsetting Defaults

To unset the default for one compute type, omit the corresponding field. To unset both, omit both fields.
Note that the Terraform wildcard `*` for `update_mask` cannot be used to unset fields — you must specify field paths explicitly.

### Permissions

Only workspace admins can configure the default workspace base environment.


## Example Usage
### Set Default for CPU Compute

```hcl
resource "databricks_environments_default_workspace_base_environment" "this" {
  cpu_workspace_base_environment = "workspace-base-environments/my-base-env-12345"
}
```

### Set Defaults for Both CPU and GPU Compute

```hcl
resource "databricks_environments_workspace_base_environment" "cpu_env" {
  display_name = "my-cpu-environment"
  filepath     = "/Volumes/catalog/schema/volume/cpu-environment.yaml"
}

resource "databricks_environments_workspace_base_environment" "gpu_env" {
  display_name          = "my-gpu-environment"
  filepath              = "/Volumes/catalog/schema/volume/gpu-environment.yaml"
  base_environment_type = "GPU"
}

resource "databricks_environments_default_workspace_base_environment" "this" {
  cpu_workspace_base_environment = databricks_environments_workspace_base_environment.cpu_env.name
  gpu_workspace_base_environment = databricks_environments_workspace_base_environment.gpu_env.name
}
```

### Unset Both Defaults

```hcl
resource "databricks_environments_default_workspace_base_environment" "this" {}
```


## Arguments
The following arguments are supported:
* `cpu_workspace_base_environment` (string, optional) - The default workspace base environment for CPU compute.
  Format: workspace-base-environments/{workspace_base_environment}
* `gpu_workspace_base_environment` (string, optional) - The default workspace base environment for GPU compute.
  Format: workspace-base-environments/{workspace_base_environment}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `name` (string) - The resource name of this singleton resource.
  Format: default-workspace-base-environment

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_environments_default_workspace_base_environment.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_environments_default_workspace_base_environment.this "name"
```