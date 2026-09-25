# NEXT CHANGELOG

## Release v1.135.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `string_value_wo` and `string_value_wo_version` attributes to `databricks_secret` resource ([#5480](https://github.com/databricks/terraform-provider-databricks/pull/5480)).

### Bug Fixes

* Fix `databricks_pipeline` so that `photon`, `serverless`, `continuous` and `development` set to `false` are sent in the create/update request if specified ([#5806](https://github.com/databricks/terraform-provider-databricks/pull/5806)).

### Documentation

### Exporter

### Internal Changes
