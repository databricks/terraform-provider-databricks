---
subcategory: "Settings"
---

# databricks_disable_legacy_features_setting Resource

-> This resource can only be used with an account-level provider!

-> This setting is currently in private preview, and only available for enrolled customers.

The `databricks_disable_legacy_features_setting` resource allows you to disable legacy features on newly created workspaces.

When this setting is on, the following applies to new workspaces:
- Disables the use of DBFS root and mounts.
- Hive Metastore will not be provisioned.
- Disables the use of ‘No-isolation clusters’.
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

```bash
terraform import databricks_disable_legacy_features_setting.this global
```
