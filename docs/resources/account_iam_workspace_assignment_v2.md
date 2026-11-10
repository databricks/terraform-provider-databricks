---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_workspace_assignment_v2 Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Manages the direct assignment of an account principal to a workspace, along with the entitlements that assignment grants. Creating the resource assigns the principal to the workspace; deleting it removes the assignment. `entitlements` is the only updatable field.


## Example Usage
Example usage:

Assigns a principal to a workspace with a set of entitlements. `workspace_id`
and `principal_id` identify the assignment; `entitlements` is the set granted
directly on it (at least one is required). `entitlements` is the only updatable
field — edit the list and re-apply to update in place. Changing `workspace_id`
or `principal_id` replaces the assignment.

```hcl
resource "databricks_account_iam_workspace_assignment_v2" "this" {
  parent       = "accounts/00000000-0000-0000-0000-000000000000/workspaces/123456789"
  principal_id = 987654321
  entitlements = ["WORKSPACE_ACCESS"]
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the workspace, in the format
  `accounts/{account_id}/workspaces/{workspace_id}`. Workspace APIs return the same format
* `principal_id` (integer, required) - The Databricks ID of the assigned user, service principal, or group. It can't be changed after
  the assignment is created
* `entitlements` (list of string, optional) - The entitlements granted directly by this assignment. Create and update set exactly this list,
  even if the principal also gets some of these entitlements through a group. List responses
  leave it empty

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The ID of the account that contains the workspace
* `effective_entitlements` (list of string) - Every entitlement the principal has in the workspace, granted directly or through a group. List
  responses leave it empty
* `name` (string) - The resource name of the assignment, in one of these formats:
  - Account APIs:
  `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
  - Workspace APIs: `workspace-assignments/{principal_id}`
* `principal_type` (string) - Whether the assigned principal is a user, service principal, or group. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `workspace_id` (integer) - The ID of the workspace

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_account_iam_workspace_assignment_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_iam_workspace_assignment_v2.this "name"
```