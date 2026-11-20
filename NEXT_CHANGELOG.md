# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_lakewatch_parser`.
* Add resource and data sources for `databricks_lakewatch_parser_version`.
* Add data sources for `databricks_lakewatch_builtin_parser`.
* Add data sources for `databricks_lakewatch_builtin_parser_version`.

* Added `databricks_user_by_external_id`, `databricks_group_by_external_id`, and `databricks_service_principal_by_external_id` data sources to resolve principals by the external ID assigned to them by the customer's Identity Provider ([#5932](https://github.com/databricks/terraform-provider-databricks/pull/5932)).

  These data sources work with both the account-level and workspace-level provider. They call the Databricks `resolve-by-external-id` APIs, which create the principal in the account if one with the given `external_id` does not already exist, and require the account to be onboarded to Automatic Identity Management (AIM).

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
