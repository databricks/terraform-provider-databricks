# NEXT CHANGELOG

## Release v1.66.0

### New Features and Improvements

 * Add multipart permissions to `databricks_aws_unity_catalog_policy` data source ([#4440](https://github.com/databricks/terraform-provider-databricks/pull/4440)).

### Bug Fixes
 * Fixed an issue where reordering objects in a (pluginfw) Share wouldn’t update properly unless other changes were made ([#4481](https://github.com/databricks/terraform-provider-databricks/pull/4481)).

### Documentation

 * Add an example for Databricks Apps permissions ([#4475](https://github.com/databricks/terraform-provider-databricks/pull/4475)).
 * Add explanation of timeouts to the troubleshooting guide ([#4482](https://github.com/databricks/terraform-provider-databricks/pull/4482)).

### Exporter

 * Refactor UC, SQL and SCIM objects into separate files ([#4477](https://github.com/databricks/terraform-provider-databricks/pull/4477)).

### Internal Changes
