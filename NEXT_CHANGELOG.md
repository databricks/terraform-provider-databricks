# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).

### Documentation

### Exporter

### Internal Changes

* Move the integration-test inert-path allowlist into a single shared classifier (`.github/scripts/classify_inert.mjs`) consumed by both the `detect-changes` and `carry-forward` jobs, replacing the duplicated `dorny/paths-filter` definition and inline glob matcher so the two classifiers can no longer drift.
