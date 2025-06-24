---
subcategory: "Delta Sharing"
---
# databricks_shares Data Source

Retrieves a list of [databricks_share](../resources/share.md) name, that were created by Terraform or manually.

-> This data source can only be used with a workspace-level provider!

## Example Usage

Getting all existing shares in the metastore

```hcl
data "databricks_shares" "this" {}

output "share_name" {
  value     = data.databricks_shares.this.shares
  sensitive = false
}
```

## Attribute Reference

This data source exports the following attributes:

* `shares` - list of [databricks_share](../resources/share.md) names.

## Related Resources

The following resources are used in the same context:

* [databricks_share](../resources/share.md) to create Delta Sharing shares.
* [databricks_recipient](../resources/recipient.md) to create Delta Sharing recipients.
* [databricks_grants](../resources/grants.md) to manage Delta Sharing permissions.
