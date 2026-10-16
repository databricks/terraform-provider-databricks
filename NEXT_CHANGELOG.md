# NEXT CHANGELOG

## Release v1.138.0

### Important Changes

### Breaking Changes

* Remove the `databricks_mason_managed_memory_entry`, `databricks_mason_managed_memory_store`, `databricks_mason_session`, and `databricks_mason_session_store` resources and their data sources ([#6087](https://github.com/databricks/terraform-provider-databricks/pull/6087)).

  The Databricks Go SDK renamed the Mason API to AgentKit.

* Remove the `databricks_feature_engineering_kafka_config` and `databricks_feature_engineering_materialized_feature` resources and their data sources ([#6087](https://github.com/databricks/terraform-provider-databricks/pull/6087)).

  The API no longer exposes these resources to Terraform.

### New Features and Improvements
* Add `personal_access_token_wo` and `personal_access_token_wo_version` write-only attributes to `databricks_git_credential` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Add `options_wo` and `options_wo_version` write-only attributes to `databricks_connection` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Bug Fixes
* Treat `pem_key_passphrase`, `userProvidedServerCertificate`, `ssl_keystore_key`, `ssl_key_password`, `schema_registry_api_key` and `schema_registry_api_secret` options of `databricks_connection` as sensitive, preserving their configured values on read ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Documentation

### Exporter

### Internal Changes
