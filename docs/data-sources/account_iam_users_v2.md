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
* `parent` (string, required) - The parent resource name of the account whose users to list.
  Format: accounts/{account_id}
* `filter` (string, optional) - Optional. Allows filtering users by username or external id
* `page_size` (integer, optional) - The maximum number of users to return. The service may return fewer than this value.
  If not provided, defaults to 1000, which is also the maximum allowed. Requests for more than the maximum are clamped to 1000


## Attributes
This data source exports a single attribute, `users`. It is a list of resources, each with the following attributes:
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