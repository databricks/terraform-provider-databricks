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
* Add the `USE_SDK_GO_RESOURCES` environment variable, which makes the listed resources call the API through the modular `github.com/databricks/sdk-go` clients ([#6185](https://github.com/databricks/terraform-provider-databricks/pull/6185), [#6186](https://github.com/databricks/terraform-provider-databricks/pull/6186)).

  `USE_SDK_GO_RESOURCES=databricks_service_principal` makes `databricks_service_principal` read and write through the `github.com/databricks/sdk-go/scim` v1 client. The `/Me` lookup and the workspace deletes for `force_delete_repos` and `force_delete_home_dir` stay on the existing client. It is unset by default, and an unknown resource name fails provider configuration.
* Add resource and data sources for `databricks_assume_group_restricted_sharing_group`.
* Add `personal_access_token_wo` and `personal_access_token_wo_version` write-only attributes to `databricks_git_credential` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Add `options_wo` and `options_wo_version` write-only attributes to `databricks_connection` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Bug Fixes
* Keep the configured values of the `pem_key_passphrase`, `ssl_keystore_key`, `ssl_key_password` and `schema_registry_api_secret` options of `databricks_connection`, which the API doesn't return, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Ignore whitespace-only differences in PEM-encoded options of `databricks_connection`, such as `userProvidedServerCertificate`, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Documentation

### Exporter

### Internal Changes
