---
subcategory: "Settings"
---

# databricks_disable_legacy_dbfs_setting Resource

The `databricks_disable_legacy_dbfs_setting` resource allows you to disable legacy DBFS.

-> This resource can only be used with a workspace-level provider!

Disabling legacy DBFS has the following implications:

1. Access to DBFS root and DBFS mounts is disallowed (as well as the creation of new mounts). 
2. Disables Databricks Runtime versions prior to 13.3LTS.

When the setting is off, all DBFS functionality is enabled and no restrictions are imposed on Databricks Runtime versions. This setting can take up to 20 minutes to take effect and requires a manual restart of all-purpose compute clusters and SQL warehouses.

Refer to official docs for more details:

- [Azure](https://learn.microsoft.com/azure/databricks/dbfs/disable-dbfs-root-mounts)
- [AWS](https://docs.databricks.com/aws/dbfs/disable-dbfs-root-mounts)
- [GCP](https://docs.gcp.databricks.com/dbfs/disable-dbfs-root-mounts)

## Example Usage

```hcl
resource "databricks_disable_legacy_dbfs_setting" "this" {
  disable_legacy_dbfs {
    value = true
  }
}
```

## Argument Reference

The resource supports the following arguments:

- `disable_legacy_dbfs` block with following attributes:
  - `value` - (Required) The boolean value for the setting.

## Import

This resource can be imported by predefined name `global`:

```hcl
import {
  to = databricks_disable_legacy_dbfs_setting.this
  id = "global"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_disable_legacy_dbfs_setting.this global
```
