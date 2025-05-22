# NEXT CHANGELOG

## Release v1.81.0

### Breaking Changes
* Removed `access_point` argument for `databricks_external_location` resource [#4735](https://github.com/databricks/terraform-provider-databricks/pull/4735)

### New Features and Improvements

* Allow to specify budget policy for `databricks_vector_search_endpoint` [#4707](https://github.com/databricks/terraform-provider-databricks/pull/4707)
* Added `account_network_policy` resource and data sources [#4735](https://github.com/databricks/terraform-provider-databricks/pull/4735)
* Added `recipient_federation_policy` resource and data sources [#4735](https://github.com/databricks/terraform-provider-databricks/pull/4735)
* Added `workspace_network_option` resource and data sources [#4735](https://github.com/databricks/terraform-provider-databricks/pull/4735)

### Bug Fixes

 * Don't fail delete when `databricks_system_schema` can be disabled only by Databricks [#4727](https://github.com/databricks/terraform-provider-databricks/pull/4727)
 * Fix debug logging for attributes used to configure the provider ([#4728](https://github.com/databricks/terraform-provider-databricks/pull/4728)).
 * Add missing external Id in trust relationship for `databricks_aws_unity_catalog_assume_role_policy` ([#4738](https://github.com/databricks/terraform-provider-databricks/pull/4738)).

### Documentation

 * Fix links to Delta Live Tables docs [#4732](https://github.com/databricks/terraform-provider-databricks/pull/4732)
 * Replaced `managed_policy_arns` with `aws_iam_role_policy_attachment` in AWS guides ([#4737](https://github.com/databricks/terraform-provider-databricks/pull/4737)).

### Exporter

### Internal Changes
* Bump Go SDK to v0.70.0 [#4735](https://github.com/databricks/terraform-provider-databricks/pull/4735)
