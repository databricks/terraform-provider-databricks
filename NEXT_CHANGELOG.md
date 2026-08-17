# NEXT CHANGELOG

## Release v1.127.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add support for `model_service`, `mcp_service`, and `model_provider_service` securables in `databricks_grant` and `databricks_grants` ([#5941](https://github.com/databricks/terraform-provider-databricks/pull/5941)).

 * Added `user_agent_extra` provider configuration attribute to append products to the `User-Agent` header, equivalent to the `DATABRICKS_USER_AGENT_EXTRA` environment variable ([#5863](https://github.com/databricks/terraform-provider-databricks/pull/5863)).

   This lets Terraform modules built on top of the provider configure usage attribution in their `provider` block without requiring users to set environment variables.

### Bug Fixes
* Fixed `databricks_share` failing with `produced an unexpected new value: .comment` when a share's comment is set or removed outside of Terraform (e.g. in the UI). The `comment` attribute is now `Optional`+`Computed`, so an out-of-band value is adopted into state instead of erroring, while `comment = ""` still explicitly clears the description.

### Documentation

* Document the `roles/group.assumer` role for group rule sets in `databricks_access_control_rule_set` ([#5924](https://github.com/databricks/terraform-provider-databricks/pull/5924)).

### Exporter

### Internal Changes
