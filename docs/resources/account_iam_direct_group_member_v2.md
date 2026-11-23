---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_direct_group_member_v2 Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Manages a direct membership of a principal (user, service principal, or group) in an account group. Creating the resource adds the member; deleting it removes the member. Membership is immutable — changing the group or principal replaces the resource. Only directly-assigned, Databricks-managed memberships are managed here; memberships inherited through the identity provider are read-only.


## Example Usage
Example usage:

Adds a principal as a direct member of a group. `group_id` and `principal_id`
are the internal Databricks IDs of the group and the member. Membership is
immutable — changing either field replaces the resource.

```hcl
resource "databricks_account_iam_direct_group_member_v2" "this" {
  parent       = "accounts/00000000-0000-0000-0000-000000000000/groups/123456789"
  principal_id = 987654321
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the group, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}`
  - Workspace APIs: `groups/{group_id}`
* `principal_id` (integer, required) - The Databricks ID of the member

## Attributes
In addition to the above arguments, the following attributes are exported:
* `display_name` (string) - The member's display name
* `external_id` (string) - The ID of the member in your identity provider (IdP)
* `group_id` (integer) - The Databricks ID of the group
* `membership_source` (string) - Where the membership comes from. Possible values are: `IDENTITY_PROVIDER`, `INTERNAL`
* `name` (string) - The resource name of the membership, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`
  - Workspace APIs: `groups/{group_id}/direct-members/{principal_id}`
* `principal_type` (string) - Whether the member is a user, service principal, or group. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_account_iam_direct_group_member_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_iam_direct_group_member_v2.this "name"
```