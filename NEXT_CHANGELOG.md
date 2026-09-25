# NEXT CHANGELOG

## Release v1.135.0

### Important Changes

### Breaking Changes

* Require exactly one nonempty GCP target in `databricks_mws_ncc_private_endpoint_rule`. Omit unused targets instead of setting `all_vpc_sc_services = false` or `service_attachment = ""`; `google_api_endpoints.endpoints` must be nonempty ([#5973](https://github.com/databricks/terraform-provider-databricks/pull/5973)).

### New Features and Improvements
* Add `string_value_wo` and `string_value_wo_version` attributes to `databricks_secret` resource ([#5480](https://github.com/databricks/terraform-provider-databricks/pull/5480)).

* Add GCP targets to `databricks_mws_ncc_private_endpoint_rule` ([#5973](https://github.com/databricks/terraform-provider-databricks/pull/5973)).

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
