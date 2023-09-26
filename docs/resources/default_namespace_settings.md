---
subcategory: "Settings"
---

# databricks_namespace_settings Resource

The `databricks_default_namespace_settings` resource allows you to set the settings for the default namespace.

## Example Usage

```hcl
resource "databricks_default_namespace_settings" "this" {
		namespace {
			value = "namespace_value"
		}
}
```

## Argument Reference

The resource supports the following arguments:

* `namespace` - (Required) The configuration details.
* `value` - (Required) The value for the setting.


