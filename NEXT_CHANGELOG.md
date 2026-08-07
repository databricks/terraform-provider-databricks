# NEXT CHANGELOG

## Release v1.125.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add optional `git_credential_id` attribute to `databricks_repo` resource to allow explicit credential selection ([#5877](https://github.com/databricks/terraform-provider-databricks/pull/5877)).

* For workspace-level hosts, resolve and validate the provider's `workspace_id` from the host's `/.well-known/databricks-config` discovery metadata instead of a SCIM `/Me` call ([#5922](https://github.com/databricks/terraform-provider-databricks/pull/5922)).

  This removes the authenticated `/Me` request on workspace hosts (avoiding false failures for service principals that can manage resources but cannot call `/Me`), and makes a `workspace_id` that disagrees with the host fail at plan time instead of at apply. Hosts whose metadata does not advertise a `workspace_id` fall back to the previous `/Me` behavior.

### Bug Fixes

* Fix `databricks_mws_ncc_private_endpoint_rule` so that Create waits for the private endpoint to be provisioned on the cloud side before returning ([#XXXX](https://github.com/databricks/terraform-provider-databricks/pull/XXXX)).

  The NCC `CreatePrivateEndpointRule` API can return immediately with `connection_state=CREATING` and an empty `vpc_endpoint_id` / `endpoint_name`, breaking downstream resources that reference those fields. Create now polls `GetPrivateEndpointRule` until `connection_state` reaches `PENDING` or `ESTABLISHED`, surfaces `error_message` on `CREATE_FAILED`, and honours a configurable Create timeout (default 30 minutes; override with a `timeouts { create = "..." }` block).

### Documentation

### Exporter

### Internal Changes
