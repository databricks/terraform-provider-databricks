# NEXT CHANGELOG

## Release v1.132.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

### Bug Fixes
* `databricks_connection`: HTTP connections no longer show a perpetual `options` diff. The backend adds a server-managed `auth_scheme` option (e.g. `bearer`) that is not part of the configuration; it is now stripped on read, for both metastore- and schema-level connections ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).
* `databricks_connection`: `name` is now `ForceNew`. The resource never implemented rename (the update path did not send `new_name`), so changing `name` previously produced a silent no-op and a perpetual plan diff; a name change now correctly forces recreation ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).
* `databricks_connection`: schema-level connections no longer show a perpetual `environment_settings` diff. The backend returns an empty `environment_settings` object for schema-level connections, which the provider materialized as an empty block that always diffed against configurations omitting it; the empty object is now dropped on read ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).
* `databricks_app` can now manage `git_source`, `source_code_path`, and `git_repository.caller_credential_id`. These are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now calls the generated `SyncFields` reconciliation to preserve the configured value across reads, and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when any of these fields was set.

### Documentation

* Document the `auto_deploy` and `caller_credential_id` fields of the `databricks_app` `git_repository` block, the `git_source` block, and the top-level `source_code_path` argument. Clarify that `auto_deploy` requires `git_source` to specify a `branch`. These fields become manageable with [#5977](https://github.com/databricks/terraform-provider-databricks/pull/5977).

### Exporter

### Internal Changes
