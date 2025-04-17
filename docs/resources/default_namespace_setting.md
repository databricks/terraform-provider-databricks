---
subcategory: "Settings"
---

# databricks_default_namespace_setting Resource

The `databricks_default_namespace_setting` resource allows you to operate the setting configuration for the default namespace in the Databricks workspace.

-> This resource can only be used with a workspace-level provider!

Setting the default catalog for the workspace determines the catalog that is used when queries do not reference
a fully qualified 3 level name. For example, if the default catalog is set to 'retail_prod' then a query
'SELECT * FROM myTable' would reference the object 'retail_prod.default.myTable'
(the schema 'default' is always assumed).
This setting requires a restart of clusters and SQL warehouses to take effect. Additionally, the default namespace only applies when using Unity Catalog-enabled compute.

## Example Usage

```hcl
resource "databricks_default_namespace_setting" "this" {
  namespace {
    value = "namespace_value"
  }
}
```

## Argument Reference

The resource supports the following arguments:

* `namespace` - (Required) The configuration details.
* `value` - (Required) The value for the setting.

## Import

This resource can be imported by predefined name `global`:

```bash
terraform import databricks_default_namespace_setting.this global
```
