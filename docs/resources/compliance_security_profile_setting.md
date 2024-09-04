---
subcategory: "Settings"
---

# databricks_compliance_security_profile_workspace_setting Resource

-> **Note** This resource could be only used with workspace-level provider!

-> **Note** This setting can NOT be disabled once it is enabled.

The `databricks_compliance_security_profile_workspace_setting` resource allows you to control whether to enable the 
compliance security profile for the current workspace. Enabling it on a workspace is permanent. By default, it is 
turned off. This setting can NOT be disabled once it is enabled.

## Example Usage

```hcl
resource "databricks_compliance_security_profile_workspace_setting" "this" {
  compliance_security_profile_workspace {
    is_enabled           = true
    compliance_standards = ["HIPAA", "FEDRAMP_MODERATE"]
  }
}
```

## Argument Reference

The resource supports the following arguments:

- `compliance_security_profile_workspace` block with following attributes:
  - `is_enabled` - (Required) Enable the Compliance Security Profile on the workspace
  - `compliance_standards` - (Required) Enable one or more compliance standards on the workspace, e.g. `HIPAA`, `PCI_DSS`, `FEDRAMP_MODERATE`

## Import

This resource can be imported by predefined name `global`:

```bash
terraform import databricks_compliance_security_profile_workspace_setting.this global
```
