# NEXT CHANGELOG

## Release v1.132.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add support for the `secret` securable in `databricks_grant` and `databricks_grants`. ([#5996](https://github.com/databricks/terraform-provider-databricks/pull/5996))

### Bug Fixes
* Fix `databricks_grant` and `databricks_grants` for the AI Gateway securables (`model_service`, `mcp_service`, `model_provider_service`). Wiring `databricks_ai_gateway_*.<x>.name` into the corresponding grant field previously failed on apply with `No API found for 'GET /unity-catalog/permissions/model_provider_service/model-provider-services/<full_name>'`, because that `name` attribute carries a resource-name prefix (e.g. `model-provider-services/`) that the permissions API does not expect. The prefix is now stripped before it reaches the API path and the resource ID, and a `DiffSuppressFunc` treats the prefixed and bare forms as equal, so both the `.name` reference and a bare full name apply cleanly with no perpetual diff.
* `databricks_app` can now manage `git_source`, `source_code_path`, and `git_repository.caller_credential_id`. These are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now calls the generated `SyncFields` reconciliation to preserve the configured value across reads, and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when any of these fields was set.

### Documentation

* Document the `auto_deploy` and `caller_credential_id` fields of the `databricks_app` `git_repository` block, the `git_source` block, and the top-level `source_code_path` argument. Clarify that `auto_deploy` requires `git_source` to specify a `branch`. These fields become manageable with [#5977](https://github.com/databricks/terraform-provider-databricks/pull/5977).

### Exporter

### Internal Changes
