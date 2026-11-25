# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation
* Document `provider_config` for the dual account/workspace resources `databricks_group`, `databricks_user`, `databricks_service_principal`, `databricks_group_member`, `databricks_group_role`, `databricks_user_role`, `databricks_service_principal_role`, `databricks_group_instance_profile`, `databricks_user_instance_profile`, and `databricks_access_control_rule_set` ([#6275](https://github.com/databricks/terraform-provider-databricks/pull/6275)).

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
