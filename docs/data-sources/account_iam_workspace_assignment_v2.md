---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_workspace_assignment_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Reads a single workspace assignment for an account principal, identified by workspace and principal ID.


## Example Usage
Example usage:

Look up a single workspace assignment by workspace and principal ID.

```hcl
data "databricks_account_iam_workspace_assignment_v2" "this" {
  workspace_id = 123456789
  principal_id = 987654321
}
```


## Arguments
The following arguments are supported:
* `principal_id` (integer, required) - The internal ID of the principal (user/sp/group) in Databricks
* `workspace_id` (integer, required) - The workspace ID where the principal is assigned

## Attributes
The following attributes are exported:
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