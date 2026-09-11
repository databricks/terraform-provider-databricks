# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

### Bug Fixes
* `databricks_connection`: `name` is now `ForceNew`. The resource never implemented rename (the update path did not send `new_name`), so changing `name` previously produced a silent no-op and a perpetual plan diff; a name change now correctly forces recreation ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).
* `databricks_connection`: schema-level connections no longer show a perpetual `environment_settings` diff. The backend returns an empty `environment_settings` object for schema-level connections, which the provider materialized as an empty block that always diffed against configurations omitting it; the empty object is now dropped on read ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

### Documentation

### Exporter

### Internal Changes
