---
subcategory: "Settings"
---

# databricks_disable_legacy_access_setting Resource

-> This resource can only be used with a workspace-level provider!

The `databricks_disable_legacy_access_setting` resource allows you to disable legacy access. It has the following impact:
1. Disables direct access to Hive Metastores from the workspace. However, you can still access a Hive Metastore through Hive Metastore federation.
2. Disables Fallback Mode on any External Location access from the workspace.
3. Disables Databricks Runtime versions prior to 13.3LTS.

It may take 5 minutes to take effect and requires a restart of clusters and SQL warehouses.
Please also set the default namespace to any value other than hive_metastore to avoid potential issues.
## Example Usage

```hcl
resource "databricks_disable_legacy_access_setting" "this" {
  disable_legacy_access {
 	value = "true"
  }
}
```

## Argument Reference

The resource supports the following arguments:

* `disable_legacy_access` - (Required) The configuration details.
* `value` - (Required) The boolean value for the setting.

## Import

This resource can be imported by predefined name `global`:

```bash
terraform import databricks_disable_legacy_access_setting.this global
```
