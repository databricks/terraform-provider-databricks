# Terraform V2 Tracker

> This file tracks problems that need to be solved before launching the Databricks Terraform Provider v2. This list is not exhaustive and may evolve over time.

---

## TODO

- [ ] **Eliminate dual account/workspace resources** — Resources like groups, SCIM, and others that currently work at both account and workspace level must be separated into distinct resources (e.g., `databricks_account_group` vs `databricks_workspace_group`). No single resource should serve both levels.
