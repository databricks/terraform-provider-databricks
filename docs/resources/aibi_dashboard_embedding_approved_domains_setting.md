---
subcategory: "Settings"
---

# databricks_aibi_dashboard_embedding_approved_domains_setting Resource

The `databricks_aibi_dashboard_embedding_approved_domains_setting` resource allows you to specify the list of domains allowed for  [embedding of AI/BI Dashboards](https://learn.microsoft.com/en-us/azure/databricks/dashboards/admin/#manage-dashboard-embedding) into other sites.

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_aibi_dashboard_embedding_access_policy_setting" "this" {
  aibi_dashboard_embedding_access_policy {
    access_policy_type = "ALLOW_APPROVED_DOMAINS"
  }
}

resource "databricks_aibi_dashboard_embedding_approved_domains_setting" "this" {
  aibi_dashboard_embedding_approved_domains {
    approved_domains = ["test.com"]
  }
  depends_on = [databricks_aibi_dashboard_embedding_access_policy_setting.this]
}
```

## Argument Reference

The resource supports the following arguments:

- `aibi_dashboard_embedding_approved_domains` block with following attributes:
  - `approved_domains` - (Required) the list of approved domains. To allow all subdomains for a given domain, use a wildcard symbol (`*`) before the domain name, i.e., `*.databricks.com` will allow to embed into any site under the `databricks.com`.

## Import

This resource can be imported by predefined name `global`:

```hcl
import {
  to = databricks_aibi_dashboard_embedding_approved_domains_setting.this
  id = "global"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_aibi_dashboard_embedding_approved_domains_setting.this global
```

## Related Resources

The following resources are often used in the same context:

- [databricks_aibi_dashboard_embedding_access_policy_setting](databricks_aibi_dashboard_embedding_access_policy_setting.md) is used to control embedding policy.
