---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_workspace_assignments_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Lists the principal assignments for a workspace.


## Example Usage
Example usage:

List all principal assignments for a workspace.

```hcl
data "databricks_account_iam_workspace_assignment_details_v2" "this" {
  workspace_id = 123456789
}
```


## Arguments
The following arguments are supported:
* `workspace_id` (integer, required) - Required. The workspace ID for which the workspace assignments are being fetched
* `page_size` (integer, optional) - The maximum number of workspace assignments to return. The service may return fewer than this value


## Attributes
This data source exports a single attribute, `workspace_assignments`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The account ID parent of the workspace where the principal is assigned
* `effective_entitlements` (list of string) - Every entitlement the principal holds in this workspace, whether granted directly or through
  group membership. Get responses populate this field. List responses leave it empty
* `entitlements` (list of string) - Entitlements granted directly to the principal on this workspace. This is the only
  client-settable field. Create and update manage exactly this set, including entitlements the
  principal also holds through a group.
  List responses leave this field empty. Get a single principal to read its entitlements
* `principal_id` (integer) - The internal ID of the principal (user/sp/group) in Databricks
* `principal_type` (string) - The type of the principal (user/service principal/group) that is assigned. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `workspace_id` (integer) - The workspace ID where the principal is assigned