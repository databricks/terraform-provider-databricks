# NEXT CHANGELOG

## Release v1.127.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add support for `model_service`, `mcp_service`, and `model_provider_service` securables in `databricks_grant` and `databricks_grants` ([#5941](https://github.com/databricks/terraform-provider-databricks/pull/5941)).

* Resolve the workspace `workspace_id` once during provider configuration for workspace-level hosts, so no resource or data source operation issues a SCIM `/Me` call to resolve it ([#5939](https://github.com/databricks/terraform-provider-databricks/pull/5939)).

  Builds on the host-metadata resolution added previously: when a workspace host advertises `workspace_id` in `/.well-known/databricks-config` it is used directly, and when the metadata omits it (older control planes) the id is resolved eagerly via a single `/Me` at configuration time. A failure to resolve now surfaces at provider configuration rather than at the first resource operation.

### Bug Fixes

### Documentation

* Document the `roles/group.assumer` role for group rule sets in `databricks_access_control_rule_set` ([#5924](https://github.com/databricks/terraform-provider-databricks/pull/5924)).

### Exporter

### Internal Changes
