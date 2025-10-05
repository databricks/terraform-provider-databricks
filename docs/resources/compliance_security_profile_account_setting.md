---
subcategory: "Settings"
---

# databricks_compliance_security_profile_account_setting Resource

The `databricks_compliance_security_profile_account_setting` resource manages the default compliance security profile
configuration that new workspaces inherit at workspace creation time. Use it with an account-level provider
configuration.

~> This resource controls whether new workspaces start with compliance security profile enforced; disabling it later only affects newly created workspaces.

-> This resource can only be used with an account-level provider!.

## Example Usage

```hcl
resource "databricks_compliance_security_profile_account_setting" "this" {
  csp_enablement_account {
    is_enforced = true
    compliance_standards = [
      "HIPAA",
      "FEDRAMP_MODERATE",
    ]
  }
}
```

## Argument Reference

The resource supports the following arguments:

- `csp_enablement_account` block with following attributes:
  - `is_enforced` - (Required) When `true`, new workspaces have the compliance security profile locked on and cannot
    opt out at the workspace level.
  - `compliance_standards` - (Required, list of strings) Compliance standards to apply by default for new workspaces.
    Accepts the same values as the workspace-level setting, e.g. `HIPAA`, `PCI_DSS`, `FEDRAMP_HIGH`, `NONE`, etc. (See [Go SDK documentation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceStandard) for the full list of supported values).

## Import

This resource can be imported by predefined name `global`:

```hcl
import {
  to = databricks_compliance_security_profile_account_setting.this
  id = "global"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_compliance_security_profile_account_setting.this global
```
