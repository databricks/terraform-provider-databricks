---
subcategory: "Settings"
---

# databricks_enhanced_security_monitoring_workspace_setting Resource

-> This resource can only be used with a workspace-level provider!

~> On Azure you need to use [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace#enhanced_security_monitoring_enabled-1) resource to configure this setting.


The `databricks_enhanced_security_monitoring_workspace_setting` resource allows you to control whether enhanced security monitoring 
is enabled for the current workspace. If the compliance security profile is enabled, this is automatically enabled. By default, 
it is disabled. However, if the compliance security profile is enabled, this is automatically enabled. If the compliance security 
profile is disabled, you can enable or disable this setting and it is not permanent.

## Example Usage

```hcl
resource "databricks_enhanced_security_monitoring_workspace_setting" "this" {
  enhanced_security_monitoring_workspace {
    is_enabled = true
  }
}
```

## Argument Reference

The resource supports the following arguments:

 - `enhanced_security_monitoring_workspace` block with following attributes:
   - `is_enabled` - (Required) Enable the Enhanced Security Monitoring on the workspace

## Import

This resource can be imported by predefined name `global`:

```bash
terraform import databricks_enhanced_security_monitoring_workspace_setting.this global
```
