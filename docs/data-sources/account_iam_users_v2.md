---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_users_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Lists the users in an account.


## Example Usage
Example usage:

List all account users.

```hcl
data "databricks_account_iam_users_v2" "this" {
  parent = "accounts/00000000-0000-0000-0000-000000000000"
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the account to list users in, in the format `accounts/{account_id}`
* `filter` (string, optional) - Filters users by `username` or `external_id`, for example
  `username eq "someone@example.com"`. Put the value in quotes
* `page_size` (integer, optional) - The maximum number of users to return, up to 1000. The default is 1000. A page can contain
  fewer users


## Attributes
This data source exports a single attribute, `users`. It is a list of resources, each with the following attributes:
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