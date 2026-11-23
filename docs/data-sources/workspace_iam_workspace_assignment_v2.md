---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_workspace_assignment_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Reads a single workspace assignment, scoped to the workspace, by its resource name, `workspace-assignments/{principal_id}`.


## Example Usage
Example usage:

Look up a single workspace assignment by principal ID, using a workspace-scoped
provider (the workspace is taken from the provider context).

```hcl
data "databricks_workspace_iam_workspace_assignment_v2" "this" {
  name = "workspace-assignments/987654321"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the assignment, in one of these formats:
  - Account APIs:
  `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
  - Workspace APIs: `workspace-assignments/{principal_id}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

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