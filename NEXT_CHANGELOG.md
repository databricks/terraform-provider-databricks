# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

* Document list indexed paths for NCC egress config fields on `databricks_mws_network_connectivity_config`, including `egress_config[0].default_rules[0]` ([#6035](https://github.com/databricks/terraform-provider-databricks/pull/6035)).

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
