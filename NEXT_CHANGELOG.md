# NEXT CHANGELOG

## Release v1.137.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_mason_managed_memory_store`.
* Add resource and data sources for `databricks_mason_managed_memory_entry`.
* Add resource and data sources for `databricks_mason_session_store`.
* Add resource and data sources for `databricks_mason_session`.
* Add resource and data sources for `databricks_private_network_gateway`.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Bump staticcheck from v0.6.0 to v0.8.1 and the minimum Go toolchain from 1.25.8 to 1.26.8, so `make lint` no longer panics with `unexpected expr: *ast.KeyValueExpr` when run with Go 1.27 ([#6043](https://github.com/databricks/terraform-provider-databricks/pull/6043)).
