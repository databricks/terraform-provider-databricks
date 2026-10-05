# NEXT CHANGELOG

## Release v1.137.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

* Rewrite portable cluster policy attributes for the target cloud and omit attributes that cannot be safely converted ([#6040](https://github.com/databricks/terraform-provider-databricks/pull/6040)).
* Export dependencies for Databricks Apps telemetry tables and Unity Catalog securables ([#6036](https://github.com/databricks/terraform-provider-databricks/pull/6036)).
* Add support for exporting dependencies of the `ai_runtime_task`, `clean_rooms_notebook_task`, `dbt_cloud_task`, `dbt_platform_task`, `gen_ai_compute_task`, and `python_operator_task` job task types. ([#6027](https://github.com/databricks/terraform-provider-databricks/pull/6027)).

### Internal Changes

* Bump staticcheck from v0.6.0 to v0.8.1 and the minimum Go toolchain from 1.25.8 to 1.26.8, so `make lint` no longer panics with `unexpected expr: *ast.KeyValueExpr` when run with Go 1.27 ([#6043](https://github.com/databricks/terraform-provider-databricks/pull/6043)).
