---
subcategory: "Settings"
---

# databricks_disable_legacy_dbfs_setting Resource

The `databricks_disable_legacy_dbfs_setting` resource allows you to disable legacy dbfs features.
When this setting is on, access to DBFS root and DBFS mounts is disallowed (as well as creation of new mounts). When the setting is off, all DBFS functionality is enabled. This setting has no impact on workspace internal storage (WIS).

~> This setting is currently in private preview, and only available for enrolled customers.

-> This resource can only be used with a workspace-level provider!

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

Alternatively, when using `terraform` version 1.5 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_disable_legacy_dbfs_setting.this global
```
