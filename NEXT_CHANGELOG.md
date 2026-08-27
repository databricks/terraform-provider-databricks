# NEXT CHANGELOG

## Release v1.130.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Allow `MINUTES` as a `unit` for the `databricks_job` `trigger.periodic` block ([#5972](https://github.com/databricks/terraform-provider-databricks/pull/5972)).

  The Jobs API accepts and schedules minute-based periodic triggers, but the provider rejected them during validation.

### Bug Fixes
* Default to a 600 second HTTP timeout for `databricks_repo` operations that run git commands inline (the clone on create and the branch/tag checkout on update), so larger repositories no longer fail with `request timed out after 1m5s of inactivity`. An explicitly configured `http_timeout_seconds` still takes precedence.

### Documentation

### Exporter

* Add support for Lakebase autoscaling resources ([#5965](https://github.com/databricks/terraform-provider-databricks/pull/5965)).

### Internal Changes
