---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_workspace_assignment_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Reads a single workspace assignment for an account principal by its resource name, `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`.


## Example Usage
Example usage:

Look up a single workspace assignment by workspace and principal ID.

```hcl
data "databricks_account_iam_workspace_assignment_v2" "this" {
  name = "accounts/00000000-0000-0000-0000-000000000000/workspaces/123456789/workspace-assignments/987654321"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the assignment, in one of these formats:
  - Account APIs:
  `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
  - Workspace APIs: `workspace-assignments/{principal_id}`

## Attributes
The following attributes are exported:
* `account_id` (string) - The ID of the account that contains the workspace
* `effective_entitlements` (list of string) - Every entitlement the principal has in the workspace, granted directly or through a group. List
  responses leave it empty
* `entitlements` (list of string) - The entitlements granted directly by this assignment. Create and update set exactly this list,
  even if the principal also gets some of these entitlements through a group. List responses
  leave it empty
* `name` (string) - The resource name of the assignment, in one of these formats:
  - Account APIs:
  `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
  - Workspace APIs: `workspace-assignments/{principal_id}`
* `parent` (string) - The resource name of the workspace, in the format
  `accounts/{account_id}/workspaces/{workspace_id}`. Workspace APIs return the same format
* `principal_id` (integer) - The Databricks ID of the assigned user, service principal, or group. It can't be changed after
  the assignment is created
* `principal_type` (string) - Whether the assigned principal is a user, service principal, or group. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `workspace_id` (integer) - The ID of the workspace