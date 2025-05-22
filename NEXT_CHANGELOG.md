# NEXT CHANGELOG

## Release v1.81.0

### New Features and Improvements

* Allow to specify budget policy for `databricks_vector_search_endpoint` [#4707](https://github.com/databricks/terraform-provider-databricks/pull/4707)

### Bug Fixes

 * Don't fail delete when `databricks_system_schema` can be disabled only by Databricks [#4727](https://github.com/databricks/terraform-provider-databricks/pull/4727)
 * Fix debug logging for attributes used to configure the provider ([#4728](https://github.com/databricks/terraform-provider-databricks/pull/4728)).

### Documentation

 * Replaced `managed_policy_arns` with `aws_iam_role_policy_attachment` in AWS guides ([#4737](https://github.com/databricks/terraform-provider-databricks/pull/4737)).

### Exporter

### Internal Changes
