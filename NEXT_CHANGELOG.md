# NEXT CHANGELOG

## Release v1.114.3

### Bug Fixes

* Fix `provider_config` block triggering destroy-and-recreate on upgrade from v1.114.0 / v1.114.2 for SDKv2 resources without an `Update` method (`databricks_token`, `databricks_secret`, `databricks_secret_scope`, `databricks_secret_acl`, `databricks_mount`, `databricks_dbfs_file`, `databricks_obo_token`, `databricks_service_principal_secret`, `databricks_user_role`, `databricks_group_role`, `databricks_user_instance_profile`, `databricks_group_instance_profile`, `databricks_service_principal_role`, `databricks_group_member`, `databricks_metastore_data_access`, `databricks_online_table`, `databricks_workspace_binding`, `databricks_catalog_workspace_binding`, `databricks_vector_search_index`). The auto-ForceNew sweep in `common.Resource.ToResource` now skips `provider_config` and supplies a no-op `Update` wrapper, so an in-place change to `provider_config` is honored without recreating the resource.

## Release v1.114.2

### Breaking Changes

### New Features and Improvements

* Support adopting pre-existing `databricks_postgres_branch` and `databricks_postgres_endpoint` resources via `replace_existing = true` argument.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Mark `effective_file_event_queue` as computed with diff suppression in `databricks_external_location` to prevent Terraform drift when the Unity Catalog backend returns the server-populated field.
