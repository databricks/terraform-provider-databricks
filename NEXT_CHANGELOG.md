# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_lakewatch_parser`.
* Add resource and data sources for `databricks_lakewatch_parser_version`.
* Add data sources for `databricks_lakewatch_builtin_parser`.
* Add data sources for `databricks_lakewatch_builtin_parser_version`.

### Bug Fixes

* Support importing workspace-level resources through an account-level or unified-host provider by passing the workspace ID in the import block `identity` (Terraform 1.12 or later) ([#6149](https://github.com/databricks/terraform-provider-databricks/pull/6149)).

  On Terraform 1.12 and 1.13, state refreshed by this version cannot be read by older provider versions (`failed to decode identity`). To downgrade the provider, run `terraform apply -refresh-only` once with the older provider and Terraform 1.14 or later, or remove the resource from state and import it again with the older provider.

### Documentation

### Exporter

### Internal Changes
