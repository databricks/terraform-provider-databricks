# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Allow `MINUTES` as a `unit` for the `databricks_job` `trigger.periodic` block ([#5972](https://github.com/databricks/terraform-provider-databricks/pull/5972)).

  The Jobs API accepts and schedules minute-based periodic triggers, but the provider rejected them during validation.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
