# NEXT CHANGELOG

## Release v1.132.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add support for the `secret` securable in `databricks_grant` and `databricks_grants`. ([#5996](https://github.com/databricks/terraform-provider-databricks/pull/5996))

### Bug Fixes
* `databricks_app` can now manage `git_source`, `source_code_path`, and `git_repository.caller_credential_id`. These are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now calls the generated `SyncFields` reconciliation to preserve the configured value across reads, and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when any of these fields was set.

* Fixed AI Gateway rate limits not being sent when `calls` or `tokens` is explicitly set to `0` in `databricks_model_serving` resource ([#5333](https://github.com/databricks/terraform-provider-databricks/issues/5333)).

### Documentation

* Document the `auto_deploy` and `caller_credential_id` fields of the `databricks_app` `git_repository` block, the `git_source` block, and the top-level `source_code_path` argument. Clarify that `auto_deploy` requires `git_source` to specify a `branch`. These fields become manageable with [#5977](https://github.com/databricks/terraform-provider-databricks/pull/5977).

### Exporter

### Internal Changes
