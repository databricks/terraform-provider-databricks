# NEXT CHANGELOG

## Release v1.128.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add resource and data sources for `databricks_account_iam_direct_group_member_v2`.
* Add resource and data sources for `databricks_account_iam_group_v2`.
* Add resource and data sources for `databricks_account_iam_service_principal_v2`.
* Add resource and data sources for `databricks_account_iam_user_v2`.
* Add resource and data sources for `databricks_account_iam_workspace_assignment_v2`.
* Add resource and data sources for `databricks_workspace_iam_direct_group_member_v2`.
* Add resource and data sources for `databricks_workspace_iam_group_v2`.
* Add resource and data sources for `databricks_workspace_iam_service_principal_v2`.
* Add resource and data sources for `databricks_workspace_iam_user_v2`.
* Add resource and data sources for `databricks_workspace_iam_workspace_assignment_v2`.
* Add resource and data source for `databricks_workspace_iam_workspace_identity_detail_v2`.

* Resolve the workspace `workspace_id` once during provider configuration for workspace-level hosts, so no resource or data source operation issues a SCIM `/Me` call to resolve it ([#5939](https://github.com/databricks/terraform-provider-databricks/pull/5939)).

  Builds on the host-metadata resolution added previously: when a workspace host advertises `workspace_id` in `/.well-known/databricks-config` it is used directly, and when the metadata omits it (older control planes) the id is resolved eagerly via a single `/Me` at configuration time. A failure to resolve now surfaces at provider configuration rather than at the first resource operation.

### Bug Fixes

### Documentation

* Document Genie budgets on `databricks_budget` with shared vs per-user examples, and blocking usage ([#5946](https://github.com/databricks/terraform-provider-databricks/pull/5946))

### Exporter

* Automatically skip interactive mode when `-services` or `-listing` options is specified ([#5900](https://github.com/databricks/terraform-provider-databricks/issues/5900)).

### Internal Changes
