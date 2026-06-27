# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).

### Documentation

### Exporter

### Internal Changes

* Deprecate `instance_profiles` attribute, and replace it with `roles` in `databricks_group` data source ([#5086](https://github.com/databricks/terraform-provider-databricks/pull/5086))
* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
