---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_user_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Reads a single account user by its internal ID.


## Example Usage
Example usage:

Look up an account user by its internal ID.

```hcl
data "databricks_account_iam_user_v2" "this" {
  user_id = "123456789"
}
```


## Arguments
The following arguments are supported:
* `user_id` (string, required) - Internal userId of the user in Databricks

## Attributes
The following attributes are exported:
* `account_id` (string) - The accountId parent of the user in Databricks
* `account_user_status` (string) - The activity status of a user in a Databricks account. Possible values are: `ACTIVE`, `INACTIVE`
* `external_id` (string) - ExternalId of the user in the customer's IdP
* `full_name` (UserFullName)
* `user_id` (string) - Internal userId of the user in Databricks
* `username` (string) - Username/email of the user

### UserFullName
* `family_name` (string)
* `given_name` (string)