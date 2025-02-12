# NEXT CHANGELOG

## Release v1.66.0

### New Features and Improvements

 * Support updating `options` in `databricks_catalog` ([#4476](https://github.com/databricks/terraform-provider-databricks/pull/4476)).
 * Increase `databricks_library` timeout from 15m to 30m.

### Bug Fixes
 * Fixed an issue where reordering objects in a (pluginfw) Share wouldn’t update properly unless other changes were made ([#4481](https://github.com/databricks/terraform-provider-databricks/pull/4481)).

 * Suppress `options.pem_private_key_expiration_epoch_sec` attribute for databricks_connection ([#4474](https://github.com/databricks/terraform-provider-databricks/pull/4474)).

### Documentation

 * Add an example for Databricks Apps permissions ([#4475](https://github.com/databricks/terraform-provider-databricks/pull/4475)).
 * Add explanation of timeouts to the troubleshooting guide ([#4482](https://github.com/databricks/terraform-provider-databricks/pull/4482)).
 * Clarify that `databricks_token` and `databricks_obo_token` could be used only with workspace-level provider ([#4480](https://github.com/databricks/terraform-provider-databricks/pull/4480)).

### Exporter

 * Refactor UC, SQL and SCIM objects into separate files ([#4477](https://github.com/databricks/terraform-provider-databricks/pull/4477)).

### Internal Changes

 * Remove incorrectly working integration test `TestAccLibraryUpdateTransitionFromPluginFw` ([#4487](https://github.com/databricks/terraform-provider-databricks/pull/4487)).
