# NEXT CHANGELOG

## Release v1.130.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Default to a 600 second HTTP timeout for `databricks_repo` operations that run git commands inline (the clone on create and the branch/tag checkout on update), so larger repositories no longer fail with `request timed out after 1m5s of inactivity`. An explicitly configured `http_timeout_seconds` still takes precedence.
* `databricks_app` can now manage `git_source` and `source_code_path`. Both are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now preserves the configured value across reads and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when either field was set.

### Documentation

### Exporter

* Add support for Lakebase autoscaling resources ([#5965](https://github.com/databricks/terraform-provider-databricks/pull/5965)).

### Internal Changes
