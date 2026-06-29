# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

### New Features and Improvements

* Add `databricks_genie_space` resource and `databricks_genie_spaces` data source, plus Genie Space support in `databricks_permissions` via the new `genie_space_id` attribute ([#5770](https://github.com/databricks/terraform-provider-databricks/pull/5770)).

  The resource normalizes `serialized_space` to suppress whitespace and key-order diffs and auto-creates a missing `parent_path` on first apply. Delete is trash-aware (treats an already-trashed space as a successful delete). Tags can be attached using the existing `databricks_workspace_entity_tag_assignment` resource with `entity_type = "geniespaces"`.

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).

### Documentation

### Exporter

### Internal Changes

* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
