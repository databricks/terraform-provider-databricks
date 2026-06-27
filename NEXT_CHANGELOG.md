# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix intermittent inconsistent final plan errors in `databricks_permissions` for SQL warehouse access control reads after permission updates ([#5837](https://github.com/databricks/terraform-provider-databricks/issues/5837)).
* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).

### Documentation

### Exporter

### Internal Changes

* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
