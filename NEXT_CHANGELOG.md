# NEXT CHANGELOG

## Release v1.82.0

### Breaking Changes

### New Features and Improvements

 * Support configuration of file events in `databricks_external_location` [#4749](https://github.com/databricks/terraform-provider-databricks/pull/4749).
 * Improve support for new fields in `databricks_pipeline` [#4744](https://github.com/databricks/terraform-provider-databricks/pull/4744).

### Bug Fixes

 * Populate `partitions` when reading `databricks_sql_table` ([#4674](https://github.com/databricks/terraform-provider-databricks/pull/4674)).
 * Fail when creating `databricks_query` and `databricks_alert` with already existing names [#4697](https://github.com/databricks/terraform-provider-databricks/pull/4697).

### Documentation

 * Improve the landing page documentation to prioritize preferred authentication methods and provide better guidance on how to configure the `host` argument.
 * Mark GKE-related fields for `databricks_mws_workspaces` and `databricks_mws_networks` as deprecated([#4752](https://github.com/databricks/terraform-provider-databricks/pull/4752)).

### Exporter

### Internal Changes

* Use caching for group membership and user information retrieval to improve performance ([#4581](https://github.com/databricks/terraform-provider-databricks/pull/4581)).
