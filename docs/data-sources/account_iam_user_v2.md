---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_user_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Reads a single account user by its resource name, `accounts/{account_id}/users/{user_id}`.


## Example Usage
Example usage:

Look up an account user by its internal ID.

```hcl
data "databricks_account_iam_user_v2" "this" {
  name = "accounts/00000000-0000-0000-0000-000000000000/users/123456789"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the user. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/users/{user_id}
  - Workspace-scoped: users/{user_id}

## Attributes
The following attributes are exported:
* `account_id` (string) - The accountId parent of the user in Databricks
* `account_user_status` (string) - The activity status of a user in a Databricks account. Possible values are: `ACTIVE`, `INACTIVE`
* `external_id` (string) - ExternalId of the user in the customer's IdP
* `full_name` (UserFullName)
* `name` (string) - The resource name of the user. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/users/{user_id}
  - Workspace-scoped: users/{user_id}
* `parent` (string) - The resource name of the account that contains the user, `accounts/{account_id}`. Both the
  account-scoped and workspace-scoped APIs return it
* `user_id` (string) - Internal userId of the user in Databricks
* `username` (string) - Username/email of the user

### UserFullName
* `family_name` (string)
* `given_name` (string)