---
subcategory: "Settings"
---

# databricks_aibi_dashboard_embedding_access_policy_setting Resource

The `databricks_aibi_dashboard_embedding_access_policy_setting` resource allows you to control [embedding of AI/BI Dashboards](https://learn.microsoft.com/en-us/azure/databricks/dashboards/admin/#manage-dashboard-embedding) into other sites.

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_aibi_dashboard_embedding_access_policy_setting" "this" {
  aibi_dashboard_embedding_access_policy {
    access_policy_type = "ALLOW_APPROVED_DOMAINS"
  }
}
```

## Argument Reference

The resource supports the following arguments:

- `aibi_dashboard_embedding_access_policy` block with following attributes:
  - `access_policy_type` - (Required) Configured embedding policy. Possible values are `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`.

## Import

This resource can be imported by predefined name `global`:

```bash
terraform import databricks_aibi_dashboard_embedding_access_policy_setting.this global
```

## Related Resources

The following resources are often used in the same context:

- [databricks_aibi_dashboard_embedding_approved_domains_setting](aibi_dashboard_embedding_approved_domains_setting.md) is used to control approved domains.
