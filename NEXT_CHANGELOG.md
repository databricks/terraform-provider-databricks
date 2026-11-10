# NEXT CHANGELOG

## Release v1.143.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `personal_access_token_wo` and `personal_access_token_wo_version` write-only attributes to `databricks_git_credential` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Add `options_wo` and `options_wo_version` write-only attributes to `databricks_connection` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Bug Fixes
* Keep the configured values of the `pem_key_passphrase`, `ssl_keystore_key`, `ssl_key_password` and `schema_registry_api_secret` options of `databricks_connection`, which the API doesn't return, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Ignore whitespace-only differences in PEM-encoded options of `databricks_connection`, such as `userProvidedServerCertificate`, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

 * Fixed `databricks_app` updates. The resource now uses the asynchronous update API with a field mask of the changed attributes, so changing a description, scopes, resources or sizing no longer fails with "Compute size updates are not supported in this update API" ([#6034](https://github.com/databricks/terraform-provider-databricks/pull/6034)).

### Documentation

### Exporter

### Internal Changes
