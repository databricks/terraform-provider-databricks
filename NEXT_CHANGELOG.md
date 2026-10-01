# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

* Added `databricks_user_by_external_id`, `databricks_group_by_external_id`, and `databricks_service_principal_by_external_id` data sources to resolve principals by the external ID assigned to them by the customer's Identity Provider ([#5932](https://github.com/databricks/terraform-provider-databricks/pull/5932)).

  These data sources work with both the account-level and workspace-level provider. They call the Databricks `resolve-by-external-id` APIs, which create the principal in the account if one with the given `external_id` does not already exist, and require the account to be onboarded to Automatic Identity Management (AIM).

### Bug Fixes

### Documentation

### Exporter

* Fixed a nil pointer panic in `emitRfaAccessRequestDestinations` when exporting workspace-level UC securables with an account-level provider ([#6028](https://github.com/databricks/terraform-provider-databricks/issues/6028)).

### Internal Changes
