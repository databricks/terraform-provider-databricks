---
subcategory: "Sandbox"
---
# databricks_sandboxes Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/sandbox)

Lists the Sandboxes visible to the caller in the workspace, returning each sandbox's spec and
current runtime status. Use this data source to enumerate existing sandboxes.


## Example Usage
```hcl
data "databricks_sandboxes" "all" {}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Maximum number of sandboxes to return in one page. When unset or 0 the server returns
  up to 1000 sandboxes; values above 5000 are reduced to 5000
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `sandboxes`. It is a list of resources, each with the following attributes:
* `create_time` (string) - Output only. The creation time of the sandbox
* `display_name` (string) - Human-readable display label for the sandbox. At most 256 characters
* `metadata` (SandboxMetadata) - Tags and other metadata associated with the sandbox
* `name` (string) - The sandbox resource name, in the form `sandboxes/{sandbox_id}`. Derived from
  `sandbox_id`; any value supplied in a create or update request body is ignored
* `spec` (SandboxSpec) - The desired configuration of the sandbox, supplied by the caller at creation time
* `status` (SandboxStatus) - The observed runtime state of the sandbox, populated by the server
* `update_time` (string) - Output only. The last update time of the sandbox metadata and spec

### ComputeSpec
* `inactivity_timeout` (string) - Idle time after which the sandbox is automatically stopped. Unset or 0 disables inactivity
  auto-stop. Negative values are rejected. Can be changed on an existing sandbox with
  UpdateSandbox

### EnvironmentSpec
* `image_uri` (string) - The Unity Catalog container image to run the sandbox from, in the form
  `catalog.schema.artifact:tag`, or `catalog.schema.artifact@sha256:<digest>` to pin an
  exact version. When omitted, the sandbox runs the default sandbox image. See
  [custom images](https://docs.databricks.com/en/compute/serverless/sandbox-custom-image)

### SandboxMetadata
* `custom_tags` (object) - Key-value tags associated with the sandbox. Supports up to 54 tags. Keys must be between
  1 and 512 characters, and values must not exceed 256 characters

### SandboxSpec
* `compute` (ComputeSpec) - Compute settings requested for the sandbox, such as the inactivity timeout
* `environment` (EnvironmentSpec) - Software environment for the sandbox, such as the container image it runs

### SandboxStatus
* `state` (string) - Lifecycle state of the sandbox.
  
  - `SANDBOX_STATE_PENDING`: The sandbox is being provisioned after a create or start request
  and cannot run commands yet.
  - `SANDBOX_STATE_RUNNING`: The sandbox is ready and commands can run in it.
  - `SANDBOX_STATE_STOPPING`: The sandbox is shutting down. It settles to
  `SANDBOX_STATE_STOPPED`, or disappears if it is being deleted.
  - `SANDBOX_STATE_STOPPED`: The sandbox is stopped, by a stop request, inactivity auto-stop,
  or a failed provisioning. Its home directory is kept and it can be started again with
  StartSandbox. Possible values are: `SANDBOX_STATE_PENDING`, `SANDBOX_STATE_RUNNING`, `SANDBOX_STATE_STOPPED`, `SANDBOX_STATE_STOPPING`