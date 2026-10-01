# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

 * Added `user_agent_extra` provider configuration attribute to append products to the `User-Agent` header, equivalent to the `DATABRICKS_USER_AGENT_EXTRA` environment variable ([#5863](https://github.com/databricks/terraform-provider-databricks/pull/5863)).

   This lets Terraform modules built on top of the provider configure usage attribution in their `provider` block without requiring users to set environment variables.

### Bug Fixes

### Documentation

### Exporter

* Fixed a nil pointer panic in `emitRfaAccessRequestDestinations` when exporting workspace-level UC securables with an account-level provider ([#6028](https://github.com/databricks/terraform-provider-databricks/issues/6028)).

### Internal Changes
