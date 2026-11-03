# NEXT CHANGELOG

## Release v1.142.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Support importing workspace-level resources through an account-level or unified-host provider by passing the workspace ID in the import block `identity` (Terraform 1.12 or later) ([#6149](https://github.com/databricks/terraform-provider-databricks/pull/6149)).

  Terraform does not pass `provider_config` to the provider during import. Importing with only `id` through such a provider now fails with an error that shows the `identity` import block, instead of reading the resource from the account API.

### Documentation

### Exporter

### Internal Changes
