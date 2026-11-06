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
* `page_size` (integer, optional) - The maximum number of workspace assignments to return. The service may return fewer than this value.
  If not provided, defaults to 1000, which is also the maximum allowed. Requests for more than the maximum are clamped to 1000
* `parent` (string, optional) - The resource name of the workspace to list assignments of,
  `accounts/{account_id}/workspaces/{workspace_id}`. Optional: defaults to the calling workspace,
  and any other workspace is rejected
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `workspace_assignments`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The account ID parent of the workspace where the principal is assigned
* `effective_entitlements` (list of string) - Every entitlement the principal holds in this workspace, whether granted directly or through
  group membership. Get responses populate this field. List responses leave it empty
* `entitlements` (list of string) - Entitlements granted directly to the principal on this workspace. This is the only
  client-settable field. Create and update manage exactly this set, including entitlements the
  principal also holds through a group.
  List responses leave this field empty. Get a single principal to read its entitlements
* `name` (string) - The resource name of the workspace assignment. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}
  - Workspace-scoped: workspace-assignments/{principal_id}
* `parent` (string) - The resource name of the workspace that contains the assignment,
  `accounts/{account_id}/workspaces/{workspace_id}`. Both the account-scoped and workspace-scoped
  APIs return it
* `principal_id` (integer) - The internal ID of the principal (user/sp/group) in Databricks. Set when the assignment is
  created and cannot be changed afterwards; to assign a different principal, create a new
  assignment
* `principal_type` (string) - The type of the principal (user/service principal/group) that is assigned. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `workspace_id` (integer) - The workspace ID where the principal is assigned