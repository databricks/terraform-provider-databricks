# NEXT CHANGELOG

## Release v1.127.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add support for `model_service`, `mcp_service`, and `model_provider_service` securables in `databricks_grant` and `databricks_grants` ([#5941](https://github.com/databricks/terraform-provider-databricks/pull/5941)).

### Bug Fixes
* Fixed `databricks_sql_endpoint` resource failing when `min_num_clusters` was changed outside of Terraform by adding `Default: 1` to match `max_num_clusters` behavior ([#5294](https://github.com/databricks/terraform-provider-databricks/issues/5294)).

### Documentation

* Document the `roles/group.assumer` role for group rule sets in `databricks_access_control_rule_set` ([#5924](https://github.com/databricks/terraform-provider-databricks/pull/5924)).

### Exporter

### Internal Changes
