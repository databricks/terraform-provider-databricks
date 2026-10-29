# NEXT CHANGELOG

## Release v1.141.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Preserve `databricks_connection` options that the API doesn't return, such as `service_account_json` for `API_SOURCE` connections, to prevent permanent configuration drift ([#6116](https://github.com/databricks/terraform-provider-databricks/pull/6116)).

### Documentation

* Fix the Azure Databricks US Gov account console host in the provider authentication docs to `https://accounts.usgov.databricks.azure.us` ([#6071](https://github.com/databricks/terraform-provider-databricks/pull/6071)).

### Exporter

### Internal Changes
