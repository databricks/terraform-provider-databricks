# NEXT CHANGELOG

## Release v1.114.0

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_supervisor_agent`.
* Add resource and data sources for `databricks_supervisor_agent_tool`.
* Add resource and data sources for `databricks_secret_uc`.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Update Go SDK to v0.128.0.
* Bump minimum Go toolchain from 1.24.0 to 1.25.7 to pick up the `crypto/tls` TLS 1.3 session-resumption fix.
* Fail at plan time with "please set api to account or workspace" for dual workspace/account resources when the provider is configured against a unified host and the resource's `api` field is not set.
* Mark `effective_file_event_queue` as computed with diff suppression in `databricks_external_location` to prevent Terraform drift when the Unity Catalog backend returns the server-populated field.
