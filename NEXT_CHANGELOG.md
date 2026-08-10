# NEXT CHANGELOG

## Release v1.126.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix perpetual `databricks_share` plan diff and null `id` by restoring the resource `id` on read and update ([#XXXX](https://github.com/databricks/terraform-provider-databricks/pull/XXXX)).

  The plugin-framework `databricks_share` set its synthetic `id` only during create. Every subsequent refresh dropped `id` to null, which made `terraform plan` perpetually report a change, and left `databricks_share.<name>.id` reading back `null` for downstream references.

### Documentation

### Exporter

### Internal Changes
