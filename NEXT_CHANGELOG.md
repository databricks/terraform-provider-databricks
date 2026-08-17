# NEXT CHANGELOG

## Release v1.127.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add support for `model_service`, `mcp_service`, and `model_provider_service` securables in `databricks_grant` and `databricks_grants` ([#5941](https://github.com/databricks/terraform-provider-databricks/pull/5941)).

### Bug Fixes
* Fixed `databricks_share` failing with `produced an unexpected new value: .comment` when a share's comment is set or removed outside of Terraform (e.g. in the UI). The `comment` attribute is now `Optional`+`Computed`, so an out-of-band value is adopted into state instead of erroring, while `comment = ""` still explicitly clears the description.

* Fixed AI Gateway rate limits not being sent when `calls` or `tokens` is explicitly set to `0` in `databricks_model_serving` resource ([#5333](https://github.com/databricks/terraform-provider-databricks/issues/5333)).

### Documentation

* Document the `roles/group.assumer` role for group rule sets in `databricks_access_control_rule_set` ([#5924](https://github.com/databricks/terraform-provider-databricks/pull/5924)).

### Exporter

### Internal Changes
