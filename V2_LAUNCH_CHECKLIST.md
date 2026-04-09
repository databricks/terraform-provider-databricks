# V2 Launch Checklist

> This file tracks problems that need to be solved before launching the Databricks Terraform Provider v2. This list is not exhaustive and may evolve over time.

---

## TODO

- [ ] **Eliminate dual account/workspace resources and data sources** — Resources and data sources that currently work at both account and workspace level must be separated into distinct ones (e.g., `databricks_account_group` vs `databricks_workspace_group`). No single resource or data source should serve both levels. Examples: `databricks_group`, `databricks_user`, `databricks_service_principal`, `databricks_metastore`, `databricks_storage_credential`, ...
