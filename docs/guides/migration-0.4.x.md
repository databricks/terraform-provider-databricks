---
page_title: "Migration from 0.3.x to 0.4.x"
---
# Migration from 0.3.x to 0.4.x

Certain resources undergone changes in order to improve long-term maintainability. You can upgrade provider with `terraform init -upgrade`. If you're currently using v0.2.x of provider, please first complete the rewrites specified in [0.2.x to 0.3.x](migration-0.3.x.md) guide.

## provider

* Remove `azure_use_pat_for_spn`, `azure_use_pat_for_cli`, `azure_pat_token_duration_seconds` attributes.
* Remove deprecated `azure_workspace_name`, `azure_resource_group`, `azure_subscription_id` in favor of just using `azure_workspace_resource_id`.
* `DATABRICKS_AZURE_CLIENT_SECRET` environment variable is no longer having any effect in favor of just using `ARM_CLIENT_SECRET`.
* `DATABRICKS_AZURE_CLIENT_ID` environment variable is no longer having any effect in favor of just using `ARM_CLIENT_ID`.
* `DATABRICKS_AZURE_TENANT_ID` environment variable is no longer having any effect in favor of just using `ARM_TENANT_ID`.
* Rename `DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID` environment variable to `DATABRICKS_AZURE_RESOURCE_ID`.