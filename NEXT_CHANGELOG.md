# NEXT CHANGELOG

## Release v1.135.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `string_value_wo` and `string_value_wo_version` attributes to `databricks_secret` resource ([#5480](https://github.com/databricks/terraform-provider-databricks/pull/5480)).

* Allow `MINUTES` as a `unit` for the `databricks_job` `trigger.periodic` block ([#5972](https://github.com/databricks/terraform-provider-databricks/pull/5972)).

  The Jobs API accepts and schedules minute-based periodic triggers, but the provider rejected them during validation.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
