# NEXT CHANGELOG

## Release v1.141.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Require replacement of `databricks_budget` when its filter changes ([#6091](https://github.com/databricks/terraform-provider-databricks/pull/6091)).

  Replacement creates a new budget; accumulated usage and alert behavior should be reviewed before applying.

### Documentation

* Fix the Azure Databricks US Gov account console host in the provider authentication docs to `https://accounts.usgov.databricks.azure.us` ([#6071](https://github.com/databricks/terraform-provider-databricks/pull/6071)).

### Exporter

### Internal Changes
