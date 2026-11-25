# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_mws_ncc_private_endpoint_rule` so that Create waits for the private endpoint to be provisioned on the cloud side before returning ([#XXXX](https://github.com/databricks/terraform-provider-databricks/pull/XXXX)).

  The NCC `CreatePrivateEndpointRule` API can return immediately with `connection_state=CREATING` and an empty `vpc_endpoint_id` / `endpoint_name`, breaking downstream resources that reference those fields. Create now polls `GetPrivateEndpointRule` until `connection_state` reaches `PENDING` or `ESTABLISHED`, surfaces `error_message` on `CREATE_FAILED`, and honours a configurable Create timeout (default 30 minutes; override with a `timeouts { create = "..." }` block).

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
