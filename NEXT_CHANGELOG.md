# NEXT CHANGELOG

## Release v1.125.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add optional `git_credential_id` attribute to `databricks_repo` resource to allow explicit credential selection ([#5877](https://github.com/databricks/terraform-provider-databricks/pull/5877)).

* For workspace-level hosts, resolve and validate the provider's `workspace_id` from the host's `/.well-known/databricks-config` discovery metadata instead of a SCIM `/Me` call ([#5922](https://github.com/databricks/terraform-provider-databricks/pull/5922)).

  This removes the authenticated `/Me` request on workspace hosts (avoiding false failures for service principals that can manage resources but cannot call `/Me`), and makes a `workspace_id` that disagrees with the host fail at plan time instead of at apply. Hosts whose metadata does not advertise a `workspace_id` fall back to the previous `/Me` behavior.

### Bug Fixes

* Fix `databricks_pipeline` with `ingestion_definition` failing when `serverless = false` is set alongside a `clusters` block ([#5783](https://github.com/databricks/terraform-provider-databricks/pull/5783)). The Go SDK marshals `serverless` with `omitempty`, so `false` was dropped and the platform defaulted ingestion pipelines to serverless, then rejected the cluster settings. The provider now force-sends `serverless` when it is explicitly set in the configuration.

### Documentation

### Exporter

### Internal Changes
