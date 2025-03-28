# NEXT CHANGELOG

## Release v1.71.0

### New Features and Improvements

 * Mark GKE-related fields for `databricks_mws_workspaces` and `databricks_mws_networks` as deprecated([#4531](https://github.com/databricks/terraform-provider-databricks/pull/4531)).
 * Add support for `aws-us-gov-dod` (AWS Govcloud DoD shard) ([#4594](https://github.com/databricks/terraform-provider-databricks/pull/4594/commits/5ac01118f546070ae5b8938f06807c8325d0f5d7))
 * Add a new settings resources disable_legacy_access ([#4578](https://github.com/databricks/terraform-provider-databricks/pull/4578)).

### Bug Fixes

 * Recreate `databricks_access_control_rule_set` when the `name` changes ([#4572](https://github.com/databricks/terraform-provider-databricks/pull/4572)).

### Documentation

 * Improve documentation for `databricks_access_control_rule_set` ([#4580](https://github.com/databricks/terraform-provider-databricks/pull/4580)).
 * Correct `first_on_demand` documentation for `aws_attributes` in `databricks_cluster`.
 * Added file events permissions to GCP external location documentation. ([#4415](https://github.com/databricks/terraform-provider-databricks/pull/4415)).

### Exporter

 * Add support for special selectors in `-listing` and `-services` [#4573](https://github.com/databricks/terraform-provider-databricks/pull/4573)
 * Allow the selective export of `databricks_mws_permission_assignment`, and change its service name to `idfed` instead of `access` ([#4571](https://github.com/databricks/terraform-provider-databricks/pull/4571))

### Internal Changes
