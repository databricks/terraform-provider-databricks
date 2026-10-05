# NEXT CHANGELOG

## Release v1.137.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Bump staticcheck from v0.6.0 to v0.8.1 and the minimum Go toolchain from 1.25.8 to 1.26.8, so `make lint` no longer panics with `unexpected expr: *ast.KeyValueExpr` when run with Go 1.27 ([#6043](https://github.com/databricks/terraform-provider-databricks/pull/6043)).
* Run `make` targets with the Go version from `go.mod`, so `make fmt` and `make lint` produce the same results as CI regardless of the locally installed Go version.
