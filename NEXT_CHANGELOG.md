# NEXT CHANGELOG

## Release v1.71.0

### New Features and Improvements

 * Mark GKE-related fields for `databricks_mws_workspaces` and `databricks_mws_networks` as deprecated([#4531](https://github.com/databricks/terraform-provider-databricks/pull/4531)).

### Bug Fixes

 * Recreate `databricks_access_control_rule_set` when the `name` changes ([#4572](https://github.com/databricks/terraform-provider-databricks/pull/4572)).
 * Avoid timeouts during `databricks_mount` state refresh and creation ([#4590](https://github.com/databricks/terraform-provider-databricks/pull/4590)).

### Documentation

 * Improve documentation for `databricks_access_control_rule_set` ([#4580](https://github.com/databricks/terraform-provider-databricks/pull/4580)).
 * Correct `first_on_demand` documentation for `aws_attributes` in `databricks_cluster`.
 * Added file events permissions to GCP external location documentation. ([#4415](https://github.com/databricks/terraform-provider-databricks/pull/4415)).
 * Improve description of `metric` field in `databricks_job` resource [#4595](https://github.com/databricks/terraform-provider-databricks/pull/4595)

### Exporter

 * Add support for special selectors in `-listing` and `-services` [#4573](https://github.com/databricks/terraform-provider-databricks/pull/4573)
 * Fix incorrect reference to model serving endpoint [#4588](https://github.com/databricks/terraform-provider-databricks/pull/4588)
 * Allow the selective export of `databricks_mws_permission_assignment`, and change its service name to `idfed` instead of `access` ([#4571](https://github.com/databricks/terraform-provider-databricks/pull/4571))
  * Fix panic caused by incorrect values in the cluster policies ([#4585](https://github.com/databricks/terraform-provider-databricks/pull/4585))

### Internal Changes
