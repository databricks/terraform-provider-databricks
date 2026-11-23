---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_user_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `name` (string, required) - The resource name of the user, in one of these formats:
  - Account APIs: `accounts/{account_id}/users/{user_id}`
  - Workspace APIs: `users/{user_id}`

## Attributes
The following attributes are exported:
* `account_id` (string) - The ID of the account that contains the user
* `account_user_status` (string) - Whether the user is active in the account. New users are active unless you set this to
  `INACTIVE`. Possible values are: `ACTIVE`, `INACTIVE`
* `external_id` (string) - The ID of the user in your identity provider (IdP)
* `full_name` (UserFullName) - The user's name
* `name` (string) - The resource name of the user, in one of these formats:
  - Account APIs: `accounts/{account_id}/users/{user_id}`
  - Workspace APIs: `users/{user_id}`
* `parent` (string) - The resource name of the account that contains the user, in the format
  `accounts/{account_id}`. Workspace APIs return the same format
* `user_id` (string) - The Databricks ID of the user
* `username` (string) - The user's username, which is their email address. It can't be changed after the user is
  created

### UserFullName
* `family_name` (string) - The user's family (last) name
* `given_name` (string) - The user's given (first) name