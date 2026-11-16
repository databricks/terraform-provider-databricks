---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_group_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Reads a single account group by its resource name, `accounts/{account_id}/groups/{group_id}`.


## Example Usage
Example usage:

Look up an account group by its internal ID.

```hcl
data "databricks_account_iam_group_v2" "this" {
  name = "accounts/00000000-0000-0000-0000-000000000000/groups/123456789"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the group, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}`
  - Workspace APIs: `groups/{group_id}`

## Attributes
The following attributes are exported:
* `account_id` (string) - The ID of the account that contains the group
* `external_id` (string) - The ID of the group in your identity provider (IdP)
* `group_id` (string) - The Databricks ID of the group
* `group_name` (string) - The name of the group
* `name` (string) - The resource name of the group, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}`
  - Workspace APIs: `groups/{group_id}`
* `parent` (string) - The resource name of the account that contains the group, in the format
  `accounts/{account_id}`. Workspace APIs return the same format