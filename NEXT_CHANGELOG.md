# NEXT CHANGELOG

## Release v1.143.0

### Important Changes

* `databricks_app` now updates apps through the asynchronous Apps update API ([#6151](https://github.com/databricks/terraform-provider-databricks/pull/6151)).
  The synchronous update API ignored `compute_size`, `compute_min_instances`, and
  `compute_max_instances`, so changes to them never took effect; they are now applied. As
  before, an update replaces the app's updatable fields with the configuration, so fields
  omitted from it are cleared. That now also covers the fields the synchronous API skipped: an
  omitted `compute_size` resets to `MEDIUM`, omitted `compute_min_instances` and
  `compute_max_instances` disable horizontal scaling, and an omitted `git_source` clears the
  app's default git source. Updates now wait for the app update to reach `SUCCEEDED`.

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_assume_group_restricted_sharing_group`.
* Add `personal_access_token_wo` and `personal_access_token_wo_version` write-only attributes to `databricks_git_credential` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Add `options_wo` and `options_wo_version` write-only attributes to `databricks_connection` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Bug Fixes
* Keep the configured values of the `pem_key_passphrase`, `ssl_keystore_key`, `ssl_key_password` and `schema_registry_api_secret` options of `databricks_connection`, which the API doesn't return, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Ignore whitespace-only differences in PEM-encoded options of `databricks_connection`, such as `userProvidedServerCertificate`, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

* Support importing workspace-level resources through an account-level or unified-host provider by passing the workspace ID in the import block `identity` (Terraform 1.12 or later) ([#6149](https://github.com/databricks/terraform-provider-databricks/pull/6149)).

  Terraform does not pass `provider_config` to the provider during import. Importing with only `id` through such a provider now fails with an error that shows the `identity` import block, instead of reading the resource from the account API.

  On Terraform 1.12 and 1.13, state refreshed by this version cannot be read by older provider versions (`failed to decode identity`). To downgrade the provider, run `terraform apply -refresh-only` once with the older provider and Terraform 1.14 or later, or remove the resource from state and import it again with the older provider.

### Documentation

### Exporter

### Internal Changes
