---
subcategory: "Settings"
---

# databricks_disable_legacy_features_setting Resource

-> This resource can only be used with an account-level provider!

The `databricks_disable_legacy_features_setting` resource allows you to disable legacy features on newly created workspaces.

~> Before disabling legacy features, make sure that default catalog for the workspace is set to value different than `hive_metastore`!  You can set it using the [databricks_default_namespace_setting](default_namespace_setting.md) resource.

When this setting is on, the following applies to new workspaces:

- Disables the use of DBFS root and mounts.
- Hive Metastore will not be provisioned.
- Disables the use of 'No-isolation clusters'.
- Disables Databricks Runtime versions prior to 13.3LTS

## Example Usage

```hcl
resource "databricks_disable_legacy_features_setting" "this" {
  disable_legacy_features {
    value = true
  }
}
```

## Argument Reference

The resource supports the following arguments:

- `disable_legacy_features` block with following attributes:
  - `value` - (Required) The boolean value for the setting.

## Import

This resource can be imported by predefined name `global`:

```hcl
import {
  to = databricks_disable_legacy_features_setting.this
  id = "global"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_disable_legacy_features_setting.this global
```

## Related Resources

The following resources are often used in the same context:

* [databricks_disable_legacy_access_setting](disable_legacy_access_setting.md) to disable legacy access, enabled by default when creating new workspaces with the `disable_legacy_features` account level setting turned on.
* [databricks_disable_legacy_dbfs_setting](disable_legacy_dbfs_setting.md) to disable legacy DBFS, enabled by default when creating new workspaces with the `disable_legacy_features` account level setting turned on.
