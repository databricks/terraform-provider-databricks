---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_workspace_assignments_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Lists the principal assignments for the workspace, using a workspace-scoped provider.


## Example Usage
Example usage:

List all principal assignments for the workspace, using a workspace-scoped
provider (the workspace is taken from the provider context).

```hcl
data "databricks_workspace_iam_workspace_assignment_details_v2" "this" {
}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of assignments to return, up to 1000. The default is 1000. A page can
  contain fewer assignments
* `parent` (string, optional) - The resource name of the workspace, in the format
  `accounts/{account_id}/workspaces/{workspace_id}`. Defaults to this workspace. Any other
  workspace is rejected
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `workspace_assignments`. It is a list of resources, each with the following attributes:
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