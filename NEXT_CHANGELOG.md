# NEXT CHANGELOG

## Release v1.140.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix perpetual diffs and `cannot be updated` errors in `databricks_budget` when the API returns alerts or actions in a different order than configured ([#6100](https://github.com/databricks/terraform-provider-databricks/pull/6100)).

  Alerts, actions, and principal overrides are now matched by content. Plans also ignore the empty filter that the API returns for budgets without one, and decimal formatting of `override_threshold`.

### Documentation

### Exporter

### Internal Changes
