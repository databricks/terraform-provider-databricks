# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `databricks_cluster` to preserve externally-set `spark_env_vars` (e.g. from cluster policies) during updates when not configured in Terraform. This also fixes `lifecycle { ignore_changes = [spark_env_vars] }` which previously failed to prevent deletion of externally-set values ([#1238](https://github.com/databricks/terraform-provider-databricks/issues/1238)).

### Documentation

### Exporter

### Internal Changes

* Host-agnostic cloud detection via node type patterns, replacing host-URL-based `IsAws()`/`IsAzure()`/`IsGcp()` checks.
