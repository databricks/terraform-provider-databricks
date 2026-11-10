---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_workspace_assignment_v2 Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Manages the direct assignment of a principal to the workspace, along with the entitlements that assignment grants, using a workspace-scoped provider. Creating the resource assigns the principal; deleting it removes the assignment. `entitlements` is the only updatable field.


## Example Usage
Example usage:

Assigns a principal to the workspace with a set of entitlements, using a
workspace-scoped provider (the workspace is taken from the provider context, so
only `principal_id` identifies the assignment). `entitlements` is the set
granted directly on it (at least one is required). `entitlements` is the only
updatable field — edit the list and re-apply to update in place. Changing
`principal_id` replaces the assignment.

```hcl
resource "databricks_workspace_iam_workspace_assignment_v2" "this" {
  principal_id = 987654321
  entitlements = ["WORKSPACE_ACCESS"]
}
```


## Arguments
The following arguments are supported:
* `principal_id` (integer, required) - The Databricks ID of the assigned user, service principal, or group. It can't be changed after
  the assignment is created
* `entitlements` (list of string, optional) - The entitlements granted directly by this assignment. Create and update set exactly this list,
  even if the principal also gets some of these entitlements through a group. List responses
  leave it empty
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The ID of the account that contains the workspace
* `effective_entitlements` (list of string) - Every entitlement the principal has in the workspace, granted directly or through a group. List
  responses leave it empty
* `name` (string) - The resource name of the assignment, in one of these formats:
  - Account APIs:
  `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
  - Workspace APIs: `workspace-assignments/{principal_id}`
* `parent` (string) - The resource name of the workspace, in the format
  `accounts/{account_id}/workspaces/{workspace_id}`. Workspace APIs return the same format
* `principal_type` (string) - Whether the assigned principal is a user, service principal, or group. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `workspace_id` (integer) - The ID of the workspace

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_workspace_iam_workspace_assignment_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_iam_workspace_assignment_v2.this "name"
```