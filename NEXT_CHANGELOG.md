# NEXT CHANGELOG

## Release v1.140.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `personal_access_token_wo` and `personal_access_token_wo_version` write-only attributes to `databricks_git_credential` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Add `options_wo` and `options_wo_version` write-only attributes to `databricks_connection` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Bug Fixes
* Treat `pem_key_passphrase`, `userProvidedServerCertificate`, `ssl_keystore_key`, `ssl_key_password`, `schema_registry_api_key` and `schema_registry_api_secret` options of `databricks_connection` as sensitive, preserving their configured values on read ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Documentation

### Exporter

### Internal Changes
