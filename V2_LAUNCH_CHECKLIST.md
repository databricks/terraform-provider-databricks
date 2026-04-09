# V2 Launch Checklist

> This file tracks problems that need to be solved before launching the Databricks Terraform Provider v2. This list is not exhaustive and may evolve over time.

---

## TODO

- [ ] **Eliminate dual account/workspace resources and data sources** — Resources and data sources that currently work at both account and workspace level must be separated into distinct ones (e.g., `databricks_account_group` vs `databricks_workspace_group`).