# NEXT CHANGELOG

## Release v1.114.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Mark `effective_file_event_queue` as read-only in `databricks_external_location` to prevent Terraform drift.
* Retry `databricks_mws_ncc_private_endpoint_rule` creation on transient HTTP inactivity timeouts ([#5648](https://github.com/databricks/terraform-provider-databricks/pull/5648)). Fixes spurious "request timed out after 1m5s of inactivity" failures during `terraform apply`.
* Retry `databricks_mws_network_connectivity_config` deletion while the backend reports that private endpoint rules or workspaces are still attached ([#5648](https://github.com/databricks/terraform-provider-databricks/pull/5648)). Absorbs eventual-consistency delays during `terraform destroy`.
### Documentation

### Exporter

### Internal Changes

* Update Go SDK to v0.128.0.
* Bump minimum Go toolchain from 1.24.0 to 1.25.7 to pick up the `crypto/tls` TLS 1.3 session-resumption fix.
