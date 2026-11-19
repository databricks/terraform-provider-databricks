---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_groups_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Lists the groups in an account.


## Example Usage
Example usage:

List all account groups.

```hcl
data "databricks_account_iam_groups_v2" "this" {
  parent = "accounts/00000000-0000-0000-0000-000000000000"
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the account to list groups in, in the format `accounts/{account_id}`
* `filter` (string, optional) - Filters groups by `group_name` or `external_id`, for example `group_name eq "Data engineers"`.
  Put the value in quotes
* `page_size` (integer, optional) - The maximum number of groups to return, up to 1000. The default is 1000. A page can contain
  fewer groups


## Attributes
This data source exports a single attribute, `groups`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The ID of the account that contains the group
* `external_id` (string) - The ID of the group in your identity provider (IdP)
* `group_id` (string) - The Databricks ID of the group
* `group_name` (string) - The name of the group
* `name` (string) - The resource name of the group, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}`
  - Workspace APIs: `groups/{group_id}`
* `parent` (string) - The resource name of the account that contains the group, in the format
  `accounts/{account_id}`. Workspace APIs return the same format