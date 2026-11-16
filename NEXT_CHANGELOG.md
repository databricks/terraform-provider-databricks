# NEXT CHANGELOG

## Release v1.144.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Support `statement_timeout` for all SQL warehouse types; omitting or removing the argument preserves the existing warehouse timeout, while explicit `0` clears the override ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Bug Fixes

* Support importing workspace-level resources through an account-level or unified-host provider by passing the workspace ID in the import block `identity` (Terraform 1.12 or later) ([#6149](https://github.com/databricks/terraform-provider-databricks/pull/6149)).

  On Terraform 1.12 and 1.13, state refreshed by this version cannot be read by older provider versions (`failed to decode identity`). To downgrade the provider, run `terraform apply -refresh-only` once with the older provider and Terraform 1.14 or later, or remove the resource from state and import it again with the older provider.

### Documentation

### Exporter

* Preserve `statement_timeout`, including `0`, in exports of all SQL warehouse types ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Internal Changes
