---
subcategory: "Unity Catalog"
---
# databricks_system_schema Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

-> **Note** This resource can only be used with a workspace-level provider!

Manages system tables enablement. System tables are a Databricks-hosted analytical store of your account’s operational data. System tables can be used for historical observability across your account. System tables must be enabled by an account admin.

## Example Usage

Enable the system schema `access`

```hcl
resource "databricks_system_schema" "this" {
  schema = "access"
}
```

## Argument Reference

The following arguments are available:

* `schema` - (Required) name of the system schema.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - the ID of system schema in form of `metastore_id|schema_name`.
* `state` - The current state of enablement for the system schema.
* `full_name` - the full name of the system schema, in form of `system.<schema>`.

## Import

This resource can be imported by the metastore id and schema name

```bash
terraform import databricks_system_schema.this '<metastore_id>|<schema_name>'
```
