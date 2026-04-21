# NEXT CHANGELOG

## Release v1.114.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Mark `effective_file_event_queue` as read-only in `databricks_external_location` to prevent Terraform drift.
### Documentation

### Exporter

### Internal Changes

* Update Go SDK to v0.128.0.
* Bump minimum Go toolchain from 1.24.0 to 1.25.7 to pick up the `crypto/tls` TLS 1.3 session-resumption fix.
