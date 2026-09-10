---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_workspace_assignment_v2 Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
  workspace_id = 123456789
  principal_id = 987654321
  entitlements = ["WORKSPACE_ACCESS"]
}
```


## Arguments
The following arguments are supported:
* `principal_id` (integer, required) - The internal ID of the principal (user/sp/group) in Databricks
* `entitlements` (list of string, optional) - Entitlements granted directly to the principal on this workspace. This is the only
  client-settable field. Create and update manage exactly this set, including entitlements the
  principal also holds through a group.
  List responses leave this field empty. Get a single principal to read its entitlements

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The account ID parent of the workspace where the principal is assigned
* `effective_entitlements` (list of string) - Every entitlement the principal holds in this workspace, whether granted directly or through
  group membership. Get responses populate this field. List responses leave it empty
* `principal_type` (string) - The type of the principal (user/service principal/group) that is assigned. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `workspace_id` (integer) - The workspace ID where the principal is assigned

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "workspace_id,principal_id"
  to = databricks_account_iam_workspace_assignment_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_iam_workspace_assignment_v2.this "workspace_id,principal_id"
```