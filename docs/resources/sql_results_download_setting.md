---
subcategory: "Settings"
---

# databricks_sql_results_download_setting Resource

The `databricks_sql_results_download_setting` resource allows you to control whether users within the workspace are allowed to download results from the SQL Editor and AI/BI Dashboards UIs. By default, this setting is enabled (set to true).

-> This resource can only be used with a workspace-level provider!

When the setting is off, users cannot download results from the SQL Editor or the AI/BI Dashboards UIs.

When this resource is deleted, it resets the SQL Results Download feature to its original default value: true.

Refer to official docs for more details:

- [Databricks SDK for Python](https://databricks-sdk-py.readthedocs.io/en/stable/workspace/settings/sql_results_download.html)

## Example Usage

```hcl
resource "databricks_sql_results_download_setting" "this" {
  boolean_val {
    value = false
  }
}
```

## Argument Reference

The resource supports the following arguments:

- `boolean_val` block with following attributes:
  - `value` - (Required) The boolean value for the setting.

## Import

This resource can be imported by predefined name `global`:

```hcl
import {
  to = databricks_sql_results_download_setting.this
  id = "global"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_sql_results_download_setting.this global
```
