---
subcategory: "Unity Catalog"
---

# databricks_connection Data Source

-> This resource can only be used with a workspace-level provider!

Retrieves details of an existing [databricks_connection](../resources/connection.md) by name. This is useful for referencing connections managed outside of the current Terraform configuration, e.g. to validate that a connection exists before creating a foreign catalog.

## Example Usage

```hcl
data "databricks_connection" "this" {
  name = "my_postgresql_connection"
}

output "connection_type" {
  value = data.databricks_connection.this.connection_info.0.connection_type
}
```

## Argument Reference

* `name` - (Required) The name of the connection.

## Attribute Reference

* `id` - The name of the connection.
* `connection_info` - block with information about the connection. It has the following fields:
  * `connection_type` - Connection type (e.g. `POSTGRESQL`, `MYSQL`, `SNOWFLAKE`, `REDSHIFT`).
  * `owner` - Owner of the connection.
  * `options` - Map of connection options (sensitive values are redacted).
  * `comment` - Free-form comment on the connection.
  * `read_only` - Whether the connection is read-only.
  * `full_name` - Full name of the connection.
  * `metastore_id` - Unique identifier of the parent metastore.
  * `created_at` - Time at which the connection was created, in epoch milliseconds.
  * `created_by` - Username of connection creator.
  * `updated_at` - Time at which the connection was last modified, in epoch milliseconds.
  * `updated_by` - Username of user who last modified the connection.

## Related Resources

* [databricks_connection](../resources/connection.md) resource to manage connections.
