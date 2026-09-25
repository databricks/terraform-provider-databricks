# NEXT CHANGELOG

## Release v1.135.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `string_value_wo` and `string_value_wo_version` attributes to `databricks_secret` resource ([#5480](https://github.com/databricks/terraform-provider-databricks/pull/5480)).

### Bug Fixes

* Fixed `databricks_registered_model` aliases not being applied to Unity Catalog by using dedicated `SetAlias`/`DeleteAlias` API calls ([#5448](https://github.com/databricks/terraform-provider-databricks/pull/5448)).

### Documentation

### Exporter

### Internal Changes
