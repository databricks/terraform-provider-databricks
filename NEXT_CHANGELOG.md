# NEXT CHANGELOG

## Release v1.132.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add `databricks_genie_space` resource and `databricks_genie_spaces` data source, plus Genie Space support in `databricks_permissions` via the new `genie_space_id` attribute ([#5770](https://github.com/databricks/terraform-provider-databricks/pull/5770)).

  The resource normalizes `serialized_space` to suppress whitespace and key-order diffs and auto-creates a missing `parent_path` on first apply. Delete is trash-aware (treats an already-trashed space as a successful delete). Tags can be attached using the existing `databricks_workspace_entity_tag_assignment` resource with `entity_type = "geniespaces"`.

### Bug Fixes
* `databricks_app` can now manage `git_source`, `source_code_path`, and `git_repository.caller_credential_id`. These are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now calls the generated `SyncFields` reconciliation to preserve the configured value across reads, and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when any of these fields was set.

### Documentation

* Document the `auto_deploy` and `caller_credential_id` fields of the `databricks_app` `git_repository` block, the `git_source` block, and the top-level `source_code_path` argument. Clarify that `auto_deploy` requires `git_source` to specify a `branch`. These fields become manageable with [#5977](https://github.com/databricks/terraform-provider-databricks/pull/5977).

### Exporter

### Internal Changes
