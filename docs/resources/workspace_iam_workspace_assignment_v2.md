---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_workspace_assignment_v2 Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `principal_id` (integer, required) - The internal ID of the principal (user/sp/group) in Databricks
* `entitlements` (list of string, optional) - Entitlements granted directly to the principal on this workspace. This is the only
  client-settable field. Create and update manage exactly this set, including entitlements the
  principal also holds through a group.
  List responses leave this field empty. Get a single principal to read its entitlements
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

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
  id = "principal_id"
  to = databricks_workspace_iam_workspace_assignment_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_iam_workspace_assignment_v2.this "principal_id"
```