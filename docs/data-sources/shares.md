---
subcategory: "Delta Sharing"
---
# databricks_shares Data Source

Retrieves a list of [databricks_share](../resources/share.md) name, that were created by Terraform or manually.

-> This data source can only be used with a workspace-level provider!

## Plugin Framework Migration

The shares data source has been migrated from sdkv2 to plugin framework. If you encounter any problem with this data source and suspect it is due to the migration, you can fallback to sdkv2 by setting the environment variable in the following way `export USE_SDK_V2_DATA_SOURCES="databricks_shares"`.

~> **Deprecation**: The SDKv2 fallback implementation, selectable via `USE_SDK_V2_DATA_SOURCES="databricks_shares"`, is **deprecated** and will be removed in the next major release of the provider. Setting the environment variable now emits a runtime warning; remove the override to use the default Plugin Framework implementation.

## Example Usage

Getting all existing shares in the metastore

```hcl
data "databricks_shares" "this" {}

output "share_name" {
  value     = data.databricks_shares.this.shares
  sensitive = false
}
```

## Argument Reference

* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

This data source exports the following attributes:

* `shares` - list of [databricks_share](../resources/share.md) names.

## Related Resources

The following resources are used in the same context:

* [databricks_share](../resources/share.md) to create Delta Sharing shares.
* [databricks_recipient](../resources/recipient.md) to create Delta Sharing recipients.
* [databricks_grants](../resources/grants.md) to manage Delta Sharing permissions.
