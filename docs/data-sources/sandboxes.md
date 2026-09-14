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
* `page_size` (integer, optional)
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `sandboxes`. It is a list of resources, each with the following attributes:
* `create_time` (string) - Output only. The creation time of the sandbox
* `display_name` (string) - Human-readable display label for the sandbox. At most 256 bytes
* `name` (string) - The AIP-compliant resource name, such as "sandboxes/my-sandbox"
* `spec` (SandboxSpec) - The desired configuration of the sandbox, supplied by the caller at creation time
* `status` (SandboxStatus) - The observed runtime state of the sandbox, populated by the server
* `update_time` (string) - Output only. The last update time of the sandbox metadata and spec

### ComputeSpec
* `inactivity_timeout` (string) - Idle duration after which the sandbox is automatically terminated

### SandboxSpec
* `compute` (ComputeSpec) - Compute configuration (size, inactivity timeout) requested for the sandbox

### SandboxStatus
* `state` (string) - Lifecycle state of the sandbox. Possible values are: `SANDBOX_STATE_PENDING`, `SANDBOX_STATE_RUNNING`, `SANDBOX_STATE_STOPPED`, `SANDBOX_STATE_STOPPING`