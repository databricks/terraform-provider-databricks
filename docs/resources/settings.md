---
subcategory: "Settings"
---

# databricks_namespace_settings Resource

The `databricks_namespace_settings` resource allows you to set the settings for the default namespace.

## Example Usage

-> **Note** `setting_name` field is currently mandatory, but only accepts the `default` value. 

```hcl
resource "databricks_namespace_settings" "this" {
  	setting_name = "default"
		namespace {
			value = "namespace_value"
		}
}
```

## Argument Reference

The resource supports the following arguments:

* `setting_name` - (Required) The setting name. Currently, only `default` setting is available.
* `namespace` - (Required) The configuration details.
* `value` - (Required) The value for the setting.


