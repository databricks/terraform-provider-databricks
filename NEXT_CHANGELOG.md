# NEXT CHANGELOG

## Release v1.126.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add support for `model_service`, `mcp_service`, and `model_provider_service` securables in `databricks_grant` and `databricks_grants` ([#5941](https://github.com/databricks/terraform-provider-databricks/pull/5941)).

### Bug Fixes

* Fix perpetual `databricks_share` plan diff and null `id` by restoring the resource `id` on read and update ([#5934](https://github.com/databricks/terraform-provider-databricks/pull/5934)).

  The plugin-framework `databricks_share` set its synthetic `id` only during create. Every subsequent refresh dropped `id` to null, which made `terraform plan` perpetually report a change, and left `databricks_share.<name>.id` reading back `null` for downstream references.

### Documentation

### Exporter

### Internal Changes
