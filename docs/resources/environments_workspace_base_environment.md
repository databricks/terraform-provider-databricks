---
subcategory: "Environments"
---
# databricks_environments_workspace_base_environment Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

A Workspace Base Environment is a shareable specification that defines a serverless environment version and additional Python dependencies for serverless notebooks and jobs.

By pre-materializing environments, Databricks reduces startup time and ensures consistent, reproducible environments across notebooks and jobs within the workspace.

**Note:** Each workspace is limited to 10 workspace base environments.

### CPU and GPU Environments

Workspace base environments can be created for either CPU or GPU serverless compute. Set `base_environment_type` to target GPU workloads.

### Environment YAML

The `filepath` field points to a WSFS or UC Volumes path containing the environment YAML file. Example:

```yaml
environment_version: "4"
dependencies:
  - --index-url https://pypi.org/simple
  - -r "/Workspace/Shared/requirements.txt"
  - my-library==6.1
  - /Workspace/Shared/Path/To/simplejson-3.19.3-py3-none-any.whl
  - git+https://github.com/databricks/databricks-cli
```

### Permissions

Workspace admins can create and manage workspace base environments. All workspace users can access and use them.


## Example Usage
### Basic Example

This example creates a Workspace Base Environment referencing an environment YAML file stored in a UC Volume.

```hcl
resource "databricks_environments_workspace_base_environment" "this" {
  display_name = "my-environment"
  filepath     = "/Volumes/catalog/schema/volume/environment.yaml"
}
```

### Example with GPU Compute Type

This example creates a GPU-specific Workspace Base Environment.

```hcl
resource "databricks_environments_workspace_base_environment" "gpu_env" {
  display_name          = "my-gpu-environment"
  filepath              = "/Volumes/catalog/schema/volume/gpu-environment.yaml"
  base_environment_type = "GPU"
}
```


## Arguments
The following arguments are supported:
* `display_name` (string, required) - Human-readable display name for the workspace base environment
* `base_environment_type` (string, optional) - The type of base environment (CPU or GPU). Possible values are: `CPU`, `GPU`
* `filepath` (string, optional) - The WSFS or UC Volumes path to the environment YAML file
* `workspace_base_environment_id` (string, optional) - The ID to use for the workspace base environment, which will become the final component of
  the resource name.
  This value should be 4-63 characters, and valid characters are /[a-z][0-9]-/
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Timestamp when the environment was created
* `creator_user_id` (string) - User ID of the creator
* `is_default` (boolean) - Whether this is the default environment for the workspace
* `last_updated_user_id` (string) - User ID of the last user who updated the environment
* `message` (string) - Status message providing additional details about the environment status
* `name` (string) - The resource name of the workspace base environment.
  Format: workspace-base-environments/{workspace-base-environment}
* `status` (string) - The status of the materialized workspace base environment. Possible values are: `CREATED`, `EXPIRED`, `FAILED`, `INVALID`, `PENDING`, `REFRESHING`
* `update_time` (string) - Timestamp when the environment was last updated

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_environments_workspace_base_environment.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_environments_workspace_base_environment.this "name"
```