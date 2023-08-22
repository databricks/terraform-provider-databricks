---
subcategory: "Unity Catalog"
---
# databricks_system_schema Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

-> **Notes**
  Unity Catalog APIs are accessible via **workspace-level APIs**. This design may change in the future.

Manages system tables enablement. System tables are a Databricks-hosted analytical store of your account’s operational data. System tables can be used for historical observability across your account. System tables must be enabled by an account admin. We strongly recommend to use a single `databricks_system_schema` per metastore.

## Example Usage

Enable the system schemas `access`, `billing`

```hcl
resource "databricks_system_schema" "this" {
  system_schema = ["access", "billing"]
}
```

## Argument Reference

The following arguments are available:

* `system_schema` - (Required) List of system schema to be enabled. Upon resource deletion, all schemas specified will be disabled, regardless of initial default state.

## Import

This resource can be imported by the metastore id

```bash
terraform import databricks_system_schema.this <metastore_id>
```
