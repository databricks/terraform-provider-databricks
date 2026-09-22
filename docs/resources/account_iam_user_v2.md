---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_user_v2 Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Manages an account-level user. Creating the resource provisions the user; deleting it removes the user. `username` is immutable — changing it replaces the resource.

When Automatic Identity Management (AIM) is enabled for the account, this resource manages local users only: a user cannot be created with `external_id` set. Users that already have an `external_id` can still be read through this resource, but only their `external_id` may be updated — their other fields are sourced from the identity provider.


## Example Usage
Example usage:

Creates an account-level user. When AIM is enabled, the user cannot be created with `external_id` set.

```hcl
resource "databricks_account_iam_user_v2" "this" {
  username            = "jane.doe@example.com"
  account_user_status = "ACTIVE"

  full_name = {
    given_name  = "Jane"
    family_name = "Doe"
  }
}
```


## Arguments
The following arguments are supported:
* `account_user_status` (string, required) - The activity status of a user in a Databricks account. Possible values are: `ACTIVE`, `INACTIVE`
* `full_name` (UserFullName, required)
* `username` (string, required) - Username/email of the user
* `external_id` (string, optional) - ExternalId of the user in the customer's IdP

### UserFullName
* `family_name` (string, optional)
* `given_name` (string, optional)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The accountId parent of the user in Databricks
* `user_id` (string) - Internal userId of the user in Databricks

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "user_id"
  to = databricks_account_iam_user_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_iam_user_v2.this "user_id"
```