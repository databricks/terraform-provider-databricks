# NEXT CHANGELOG

## Release v1.119.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

* Document that read-only workspace bindings aren't applicable for non-catalog objects ([#5611](https://github.com/databricks/terraform-provider-databricks/pull/5611))

### Exporter

### Internal Changes
* Run unit tests offline from a pre-warmed Go module cache for PRs that cannot authenticate to the internal Go module proxy (fork and Dependabot PRs), populated by the new "Warm Go Cache" workflow.
