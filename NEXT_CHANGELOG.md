# NEXT CHANGELOG

## Release v1.125.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add optional `git_credential_id` attribute to `databricks_repo` resource to allow explicit credential selection ([#5877](https://github.com/databricks/terraform-provider-databricks/pull/5877)).

 * Added `user_agent_extra` provider configuration attribute to append products to the `User-Agent` header, equivalent to the `DATABRICKS_USER_AGENT_EXTRA` environment variable ([#5863](https://github.com/databricks/terraform-provider-databricks/pull/5863)).

   This lets Terraform modules built on top of the provider configure usage attribution in their `provider` block without requiring users to set environment variables.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
