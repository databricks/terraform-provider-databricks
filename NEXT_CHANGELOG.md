# NEXT CHANGELOG

## Release v1.118.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_pipeline` so that `photon`, `serverless`, `continuous` and `development` set to `false` are sent in the create/update request. Previously these fields were silently dropped (Go SDK marshals the bool with `omitempty`), so the server applied its own default instead of the configured `false` ([#5806](https://github.com/databricks/terraform-provider-databricks/pull/5806)).

### Documentation

### Exporter

### Internal Changes
