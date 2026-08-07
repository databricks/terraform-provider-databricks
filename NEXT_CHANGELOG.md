# NEXT CHANGELOG

## Release v1.125.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add optional `git_credential_id` attribute to `databricks_repo` resource to allow explicit credential selection ([#5877](https://github.com/databricks/terraform-provider-databricks/pull/5877)).

* For workspace-level hosts, resolve and validate the provider's `workspace_id` from the host's `/.well-known/databricks-config` discovery metadata instead of a SCIM `/Me` call ([#5922](https://github.com/databricks/terraform-provider-databricks/pull/5922)).

  This removes the authenticated `/Me` request on workspace hosts (avoiding false failures for service principals that can manage resources but cannot call `/Me`), and makes a `workspace_id` that disagrees with the host fail at plan time instead of at apply. Hosts whose metadata does not advertise a `workspace_id` fall back to the previous `/Me` behavior.

* Add `databricks_genie_space` resource and `databricks_genie_spaces` data source, plus Genie Space support in `databricks_permissions` via the new `genie_space_id` attribute ([#5770](https://github.com/databricks/terraform-provider-databricks/pull/5770)).

  The resource normalizes `serialized_space` to suppress whitespace and key-order diffs and auto-creates a missing `parent_path` on first apply. Delete is trash-aware (treats an already-trashed space as a successful delete). Tags can be attached using the existing `databricks_workspace_entity_tag_assignment` resource with `entity_type = "geniespaces"`.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
