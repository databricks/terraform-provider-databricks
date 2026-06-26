# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).

### Documentation

* Document additional custom configuration properties for resource `workspace_conf` ([#4859](https://github.com/databricks/terraform-provider-databricks/pull/4859))

### Exporter

### Internal Changes

* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
