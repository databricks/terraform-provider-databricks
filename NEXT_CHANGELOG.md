# NEXT CHANGELOG

## Release v1.125.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add optional `git_credential_id` attribute to `databricks_repo` resource to allow explicit credential selection ([#5877](https://github.com/databricks/terraform-provider-databricks/pull/5877)).

* For workspace-level hosts, resolve and validate the provider's `workspace_id` from the host's `/.well-known/databricks-config` discovery metadata instead of a SCIM `/Me` call ([#5922](https://github.com/databricks/terraform-provider-databricks/pull/5922)).

  This removes the authenticated `/Me` request on workspace hosts (avoiding false failures for service principals that can manage resources but cannot call `/Me`), and makes a `workspace_id` that disagrees with the host fail at plan time instead of at apply. Hosts whose metadata does not advertise a `workspace_id` fall back to the previous `/Me` behavior.

* Add `provider_config` support for manual plugin framework resources and data sources([#5127](https://github.com/databricks/terraform-provider-databricks/pull/5127))

* Added support for custom instance profiles on instance pools on AWS ([#5144](https://github.com/databricks/terraform-provider-databricks/pull/5144))

* Added `deployment_names` attribute to `databricks_mws_workspaces` data block ([#5100](https://github.com/databricks/terraform-provider-databricks/pull/5100))

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
