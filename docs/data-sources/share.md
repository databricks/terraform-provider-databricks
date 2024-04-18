---
subcategory: "Delta Sharing"
---
# databricks_share Data Source

Retrieves details about a [databricks_share](../resources/share.md) that were created by Terraform or manually.

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
