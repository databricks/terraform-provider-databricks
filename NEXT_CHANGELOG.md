# NEXT CHANGELOG

## Release v1.138.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `personal_access_token_wo` and `personal_access_token_wo_version` write-only attributes to `databricks_git_credential` resource ([#4752](https://github.com/databricks/terraform-provider-databricks/issues/4752)).
* Add `options_wo` and `options_wo_version` write-only attributes to `databricks_connection` resource ([#4752](https://github.com/databricks/terraform-provider-databricks/issues/4752)).

### Bug Fixes
* Treat `pem_key_passphrase` and `userProvidedServerCertificate` options of `databricks_connection` as sensitive, preserving their configured values on read ([#4752](https://github.com/databricks/terraform-provider-databricks/issues/4752)).

### Documentation

### Exporter

### Internal Changes
