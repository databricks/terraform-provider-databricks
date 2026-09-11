# NEXT CHANGELOG

## Release v1.132.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* `databricks_app` can now manage `git_source`, `source_code_path`, and `git_repository.caller_credential_id`. These are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now calls the generated `SyncFields` reconciliation to preserve the configured value across reads, and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when any of these fields was set.

* Fix `databricks_mws_ncc_private_endpoint_rule` so that Create waits for the private endpoint to be provisioned on the cloud side before returning ([#XXXX](https://github.com/databricks/terraform-provider-databricks/pull/XXXX)).

  The NCC `CreatePrivateEndpointRule` API can return immediately with `connection_state=CREATING` and an empty `vpc_endpoint_id` / `endpoint_name`, breaking downstream resources that reference those fields. Create now polls `GetPrivateEndpointRule` until `connection_state` reaches `PENDING` or `ESTABLISHED`, surfaces `error_message` on `CREATE_FAILED`, and honours a configurable Create timeout (default 30 minutes; override with a `timeouts { create = "..." }` block).

### Documentation

* Document the `auto_deploy` and `caller_credential_id` fields of the `databricks_app` `git_repository` block, the `git_source` block, and the top-level `source_code_path` argument. Clarify that `auto_deploy` requires `git_source` to specify a `branch`. These fields become manageable with [#5977](https://github.com/databricks/terraform-provider-databricks/pull/5977).

### Exporter

### Internal Changes
