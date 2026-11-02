---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_group_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `name` (string, required) - The resource name of the group. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/groups/{group_id}
  - Workspace-scoped: groups/{group_id}

## Attributes
The following attributes are exported:
* `account_id` (string) - The parent account ID for group in Databricks
* `external_id` (string) - ExternalId of the group in the customer's IdP
* `group_id` (string) - Internal group ID of the group in Databricks
* `group_name` (string) - Display name of the group
* `name` (string) - The resource name of the group. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/groups/{group_id}
  - Workspace-scoped: groups/{group_id}
* `parent` (string) - The resource name of the account that contains the group, `accounts/{account_id}`. Both the
  account-scoped and workspace-scoped APIs return it