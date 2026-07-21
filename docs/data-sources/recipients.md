---
subcategory: "Delta Sharing"
---
# databricks_recipients Data Source

Retrieves a list of [databricks_recipient](../resources/recipient.md) names, that were created by Terraform or manually.

-> This data source can only be used with a workspace-level provider!

## Example Usage

Getting all existing recipients in the metastore

```hcl
data "databricks_recipients" "this" {}

output "recipient_names" {
  value = data.databricks_recipients.this.recipients
}
```

## Argument Reference

* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

This data source exports the following attributes:

* `recipients` - list of [databricks_recipient](../resources/recipient.md) names.

## Related Resources

The following resources are used in the same context:

* [databricks_recipient](../resources/recipient.md) to create Delta Sharing recipients.
* [databricks_share](../resources/share.md) to create Delta Sharing shares.
* [databricks_shares](shares.md) to get a list of Delta Sharing shares.
* [databricks_grants](../resources/grants.md) to manage Delta Sharing permissions.
