---
subcategory: "Sandbox"
---
# databricks_sandbox Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/sandbox)

Provisions and manages a Sandbox: an isolated, pre-configured, low-latency Serverless
compute environment for running code. Creating this resource provisions a sandbox with a
client-supplied id and optional compute and storage configuration; destroying it permanently
deletes the sandbox. Use the `spec.compute.inactivity_timeout` to have the sandbox
auto-terminate after an idle period.


## Example Usage
```hcl
resource "databricks_sandbox" "this" {
  sandbox_id   = "my-sandbox"
  display_name = "My sandbox"

  spec {
    compute {
      inactivity_timeout = "3600s"
    }
  }
}
```


## Arguments
The following arguments are supported:
* `sandbox_id` (string, required) - Client-supplied ID that becomes the final path segment of the resource name
* `display_name` (string, optional) - Human-readable display label for the sandbox. At most 256 bytes
* `spec` (SandboxSpec, optional) - The desired configuration of the sandbox, supplied by the caller at creation time
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### ComputeSpec
* `inactivity_timeout` (string, optional) - Idle duration after which the sandbox is automatically terminated

### SandboxSpec
* `compute` (ComputeSpec, optional) - Compute configuration (size, inactivity timeout) requested for the sandbox

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Output only. The creation time of the sandbox
* `name` (string) - The AIP-compliant resource name, such as "sandboxes/my-sandbox"
* `status` (SandboxStatus) - The observed runtime state of the sandbox, populated by the server
* `update_time` (string) - Output only. The last update time of the sandbox metadata and spec

### SandboxStatus
* `state` (string) - Lifecycle state of the sandbox. Possible values are: `SANDBOX_STATE_PENDING`, `SANDBOX_STATE_RUNNING`, `SANDBOX_STATE_STOPPED`, `SANDBOX_STATE_STOPPING`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_sandbox.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_sandbox.this "name"
```