# NEXT CHANGELOG

## Release v1.141.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Make `resource_group` and `subscription_id` optional in the `file_event_queue.managed_aqs` block of `databricks_external_location`, so they can be inferred from an Azure managed identity storage credential.

### Bug Fixes

### Documentation

* Fix the Azure Databricks US Gov account console host in the provider authentication docs to `https://accounts.usgov.databricks.azure.us` ([#6071](https://github.com/databricks/terraform-provider-databricks/pull/6071)).

### Exporter

### Internal Changes
