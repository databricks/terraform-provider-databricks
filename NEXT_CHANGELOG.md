# NEXT CHANGELOG

## Release v1.117.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_external_location` so that creating a resource with `enable_file_events = false` is sent in the POST request. Previously the field was silently dropped (Go SDK marshals the bool with `omitempty`), so the server applied its `true` default and `effective_enable_file_events` came back `true`.

### Documentation

* In resource `access_control_rule_set` the field `grant_rule` is marked as optional to match the current behaviour.

### Exporter

### Internal Changes
