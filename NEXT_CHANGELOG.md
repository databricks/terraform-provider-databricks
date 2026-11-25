# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add `serverless_compute_id`, `default_serverless_compute`, `compute_policy_id` and `default_compute_policy` to [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) to manage access to serverless compute, including the Default Interactive and Default Automated Compute ([#6215](https://github.com/databricks/terraform-provider-databricks/pull/6215)).

### Bug Fixes

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
