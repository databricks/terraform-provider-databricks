---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_direct_group_member_v2 Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Manages a direct membership of a principal (user, service principal, or group) in an account group. Creating the resource adds the member; deleting it removes the member. Membership is immutable — changing the group or principal replaces the resource. Only directly-assigned, Databricks-managed memberships are managed here; memberships inherited through the identity provider are read-only.


## Example Usage
Example usage:

Adds a principal as a direct member of a group. `group_id` and `principal_id`
are the internal Databricks IDs of the group and the member. Membership is
immutable — changing either field replaces the resource.

```hcl
resource "databricks_account_iam_direct_group_member_v2" "this" {
  group_id     = 123456789
  principal_id = 987654321
}
```


## Arguments
The following arguments are supported:
* `principal_id` (integer, required) - Internal ID of the principal in Databricks

## Attributes
In addition to the above arguments, the following attributes are exported:
* `display_name` (string) - Display name of the principal
* `external_id` (string) - The external ID of the principal in Databricks
* `group_id` (integer) - The internal ID of the group this member belongs to
* `membership_source` (string) - The source of group membership (internal or from identity provider). Possible values are: `IDENTITY_PROVIDER`, `INTERNAL`
* `principal_type` (string) - The type of the principal (user/service principal/group). Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "group_id,principal_id"
  to = databricks_account_iam_direct_group_member_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_iam_direct_group_member_v2.this "group_id,principal_id"
```