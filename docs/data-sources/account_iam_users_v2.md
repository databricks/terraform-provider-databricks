---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_users_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Lists the users in an account.


## Example Usage
Example usage:

List all account users.

```hcl
data "databricks_account_iam_users_v2" "this" {
}
```


## Arguments
The following arguments are supported:
* `filter` (string, optional) - Optional. Allows filtering users by username or external id
* `page_size` (integer, optional) - The maximum number of users to return. The service may return fewer than this value


## Attributes
This data source exports a single attribute, `users`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The accountId parent of the user in Databricks
* `account_user_status` (string) - The activity status of a user in a Databricks account. Possible values are: `ACTIVE`, `INACTIVE`
* `external_id` (string) - ExternalId of the user in the customer's IdP
* `full_name` (UserFullName)
* `user_id` (string) - Internal userId of the user in Databricks
* `username` (string) - Username/email of the user

### UserFullName
* `family_name` (string)
* `given_name` (string)