# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

### Bug Fixes
* Default to a 600 second HTTP timeout for `databricks_repo` operations that run git commands inline (the clone on create and the branch/tag checkout on update), so larger repositories no longer fail with `request timed out after 1m5s of inactivity`. An explicitly configured `http_timeout_seconds` still takes precedence.
* Prevent repeated diffs for `databricks_mws_private_access_settings.allowed_vpc_endpoint_ids` when the Accounts API returns the same IDs in a different order ([#5970](https://github.com/databricks/terraform-provider-databricks/pull/5970)).
* `databricks_app` can now manage `git_source`, `source_code_path`, and `git_repository.caller_credential_id`. These are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now calls the generated `SyncFields` reconciliation to preserve the configured value across reads, and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when any of these fields was set.

### Documentation

### Exporter

* Fixed a nil pointer panic in `emitRfaAccessRequestDestinations` when exporting workspace-level UC securables with an account-level provider ([#6028](https://github.com/databricks/terraform-provider-databricks/issues/6028)).

### Internal Changes
