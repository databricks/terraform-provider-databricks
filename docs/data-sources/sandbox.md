---
subcategory: "Sandbox"
---
# databricks_sandbox Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/sandbox)

Retrieves the configuration and runtime status of a single Sandbox by its resource name.
Use this data source to read an existing sandbox's spec and its current lifecycle state.


## Example Usage
```hcl
data "databricks_sandbox" "this" {
  name = "sandboxes/my-sandbox"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The AIP-compliant resource name, such as "sandboxes/my-sandbox"
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
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