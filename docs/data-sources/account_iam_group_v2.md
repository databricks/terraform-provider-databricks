---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_group_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Reads a single account group by its internal ID.


## Example Usage
Example usage:

Look up an account group by its internal ID.

```hcl
data "databricks_account_iam_group_v2" "this" {
  group_id = "123456789"
}
```


## Arguments
The following arguments are supported:
* `group_id` (string, required) - Internal group ID of the group in Databricks

## Attributes
The following attributes are exported:
* `account_id` (string) - The parent account ID for group in Databricks
* `external_id` (string) - ExternalId of the group in the customer's IdP
* `group_id` (string) - Internal group ID of the group in Databricks
* `group_name` (string) - Display name of the group