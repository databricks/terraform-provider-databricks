---
subcategory: "Delta Sharing"
---
# databricks_share Data Source

Retrieves details about a [databricks_share](../resources/share.md) that were created by Terraform or manually.

-> This data source can only be used with a workspace-level provider!

## Plugin Framework Migration

The share data source has been migrated from sdkv2 to plugin framework. If you encounter any problem with this data source and suspect it is due to the migration, you can fallback to sdkv2 by setting the environment variable in the following way `export USE_SDK_V2_DATA_SOURCES="databricks_share"`.

~> **Deprecation**: The SDKv2 fallback implementation, selectable via `USE_SDK_V2_DATA_SOURCES="databricks_share"`, is **deprecated** and will be removed in the next major release of the provider. Setting the environment variable now emits a runtime warning; remove the override to use the default Plugin Framework implementation.

## Example Usage

Getting details of an existing share in the metastore

```hcl
data "databricks_share" "this" {
  name = "this"
}

output "created_by" {
  value     = data.databricks_share.this.created_by
  sensitive = false
}
```

## Argument Reference

* `name` - (Required) The name of the share
* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

This data source exports the following attributes:

* `created_at` - Time when the share was created.
* `created_by` - The principal that created the share.
* `object` - arrays containing details of each object in the share.
  * `name` - Full name of the object being shared.
  * `data_object_type` - Type of the object.
  * `comment` -  Description about the object.

## Related Resources

The following resources are used in the same context:

* [databricks_share](../resources/share.md) to create Delta Sharing shares.
* [databricks_recipient](../resources/recipient.md) to create Delta Sharing recipients.
* [databricks_grants](../resources/grants.md) to manage Delta Sharing permissions.
