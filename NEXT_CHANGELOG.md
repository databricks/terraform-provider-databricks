# NEXT CHANGELOG

## Release v1.137.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

* Fix infinite loop in interactive mode when authentication keeps failing or standard input is closed: the exporter now gives up after 3 failed authentication attempts ([#6045](https://github.com/databricks/terraform-provider-databricks/pull/6045)).

### Internal Changes

* Bump staticcheck from v0.6.0 to v0.8.1 and the minimum Go toolchain from 1.25.8 to 1.26.8, so `make lint` no longer panics with `unexpected expr: *ast.KeyValueExpr` when run with Go 1.27 ([#6043](https://github.com/databricks/terraform-provider-databricks/pull/6043)).
