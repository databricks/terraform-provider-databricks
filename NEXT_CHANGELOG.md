# NEXT CHANGELOG

## Release v1.130.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Default to a 600 second HTTP timeout for `databricks_repo` operations that run git commands inline (the clone on create and the branch/tag checkout on update), so larger repositories no longer fail with `request timed out after 1m5s of inactivity`. An explicitly configured `http_timeout_seconds` still takes precedence.

### Documentation

* Document the `auto_deploy` and `caller_credential_id` fields of the `databricks_app` `git_repository` block, and add a `git_source` block reference. Clarify that `auto_deploy` requires `git_source` to specify a `branch`.

### Exporter

* Add support for Lakebase autoscaling resources ([#5965](https://github.com/databricks/terraform-provider-databricks/pull/5965)).

### Internal Changes
