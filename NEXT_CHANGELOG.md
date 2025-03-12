# NEXT CHANGELOG

## Release v1.70.0

### New Features and Improvements

### Bug Fixes

 * Fix creation of `databricks_storage_credential` and `databricks_credential` resources on GCP with isolation mode ([#4563](https://github.com/databricks/terraform-provider-databricks/pull/4563))
 * Handle auto-enabled errors with `databricks_system_schema` [#4547](https://github.com/databricks/terraform-provider-databricks/pull/4547)
 * Skip Read after Create in `databricks_secret_acl` to avoid errors([#4548](https://github.com/databricks/terraform-provider-databricks/pull/4548)).

### Documentation

 * Document `amazon_bedrock_config.instance_profile_arn` in `databricks_model_serving` ([#4549](https://github.com/databricks/terraform-provider-databricks/pull/4549))
 * Document management of permissions of `databricks_budget_policy` resource ([#4561](https://github.com/databricks/terraform-provider-databricks/pull/4561))
 * Document `budget_policy_id` in `databricks_app` resource and data sources ([#4557](https://github.com/databricks/terraform-provider-databricks/pull/4557))
 * Add a note on how `databricks_grants` work with `MANAGE` permission [#4546](https://github.com/databricks/terraform-provider-databricks/pull/4546)

### Exporter

 * Add support for special selectors in `-listing` and `-services` [#4573](https://github.com/databricks/terraform-provider-databricks/pull/4573)

### Internal Changes
