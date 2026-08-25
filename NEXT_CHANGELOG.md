# NEXT CHANGELOG

## Release v1.130.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Default to a 600 second HTTP timeout for `databricks_repo` operations that run git commands inline (the clone on create and the branch/tag checkout on update), so larger repositories no longer fail with `request timed out after 1m5s of inactivity`. An explicitly configured `http_timeout_seconds` still takes precedence.
* Prevent repeated diffs for `databricks_mws_private_access_settings.allowed_vpc_endpoint_ids` when the Accounts API returns the same IDs in a different order ([#5970](https://github.com/databricks/terraform-provider-databricks/pull/5970)).

### Documentation

### Exporter

* Add support for Lakebase autoscaling resources ([#5965](https://github.com/databricks/terraform-provider-databricks/pull/5965)).

### Internal Changes
