---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_groups_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Lists the groups in an account.


## Example Usage
Example usage:

List all account groups.

```hcl
data "databricks_account_iam_groups_v2" "this" {
}
```


## Arguments
The following arguments are supported:
* `filter` (string, optional) - Optional. Allows filtering groups by group name or external id
* `page_size` (integer, optional) - The maximum number of groups to return. The service may return fewer than this value


## Attributes
This data source exports a single attribute, `groups`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The parent account ID for group in Databricks
* `external_id` (string) - ExternalId of the group in the customer's IdP
* `group_id` (string) - Internal group ID of the group in Databricks
* `group_name` (string) - Display name of the group