---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_group_v2 Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Manages an account-level group. Creating the resource creates the group; deleting it removes the group.

When Automatic Identity Management (AIM) is enabled for the account, this resource manages local groups only: a group cannot be created with `external_id` set. Groups that already have an `external_id` can still be read through this resource, but only their `external_id` may be updated — their other fields are sourced from the identity provider.


## Example Usage
Example usage:

Creates an account-level group. When AIM is enabled, the group cannot be created with `external_id` set.

```hcl
resource "databricks_account_iam_group_v2" "this" {
  group_name = "data-engineers"
}
```


## Arguments
The following arguments are supported:
* `external_id` (string, optional) - ExternalId of the group in the customer's IdP
* `group_name` (string, optional) - Display name of the group

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The parent account ID for group in Databricks
* `group_id` (string) - Internal group ID of the group in Databricks

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "group_id"
  to = databricks_account_iam_group_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_iam_group_v2.this "group_id"
```