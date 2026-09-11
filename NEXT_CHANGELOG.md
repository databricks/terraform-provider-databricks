# NEXT CHANGELOG

## Release v1.132.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Added `databricks_user_by_external_id`, `databricks_group_by_external_id`, and `databricks_service_principal_by_external_id` data sources to resolve principals by the external ID assigned to them by the customer's Identity Provider ([#5932](https://github.com/databricks/terraform-provider-databricks/pull/5932)).

  These data sources work with both the account-level and workspace-level provider. They call the Databricks `resolve-by-external-id` APIs, which create the principal in the account if one with the given `external_id` does not already exist, and require the account to be onboarded to Automatic Identity Management (AIM).

### Bug Fixes
* `databricks_app` can now manage `git_source`, `source_code_path`, and `git_repository.caller_credential_id`. These are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now calls the generated `SyncFields` reconciliation to preserve the configured value across reads, and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when any of these fields was set.

### Documentation

* Document the `auto_deploy` and `caller_credential_id` fields of the `databricks_app` `git_repository` block, the `git_source` block, and the top-level `source_code_path` argument. Clarify that `auto_deploy` requires `git_source` to specify a `branch`. These fields become manageable with [#5977](https://github.com/databricks/terraform-provider-databricks/pull/5977).

### Exporter

### Internal Changes
