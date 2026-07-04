# NEXT CHANGELOG

## Release v1.121.0

### Important Changes

* [Backward Incompatible] Resource `databricks_mws_ncc_private_endpoint_rule` has been refactored in this version to fix multiple bugs in its implementation and behavior. Server-populated attributes are now strictly read-only (writes that previously produced perpetual plan drift now error at plan time), plan churn on previously-Optional output fields is eliminated, and the new Plugin Framework implementation (behind `DATABRICKS_TF_ENABLED_PF_RESOURCES`) polls the rule's `connection_state` during Create so a failed or unusable provisioning state surfaces at apply time instead of on the next plan. See `Breaking Changes` and `New Features and Improvements` below for specifics.

### Breaking Changes

* Tighten read-only attributes on `databricks_mws_ncc_private_endpoint_rule` ([#5819](https://github.com/databricks/terraform-provider-databricks/pull/5819)). The attributes `rule_id`, `account_id`, `endpoint_name`, `vpc_endpoint_id`, `connection_state`, `creation_time`, `updated_time`, `deactivated`, `deactivated_at`, `error_message`, and `gcp_endpoint.psc_endpoint_uri` are now computed-only and can no longer be set in HCL. They are populated by the server in every API response (and `psc_endpoint_uri` is unconditionally overwritten from the cloud-platform truth on every read), and previous releases accepted writes to these attributes silently while the API ignored or overwrote the values, producing perpetual drift on every plan. Configurations that explicitly assigned any of these attributes must remove the assignment; the value is still available in state after apply.

### New Features and Improvements

* A Plugin Framework implementation of `databricks_mws_ncc_private_endpoint_rule` is now available behind `DATABRICKS_TF_ENABLED_PF_RESOURCES=databricks_mws_ncc_private_endpoint_rule` ([#5819](https://github.com/databricks/terraform-provider-databricks/pull/5819)). Once opted in, `terraform apply` waits for the rule to leave `CREATING` before returning: `PENDING` and `ESTABLISHED` succeed, while a `CREATE_FAILED`, `REJECTED`, `DISCONNECTED`, or `EXPIRED` connection state surfaces as an apply-time error instead of on the next plan. This differs from the SDKv2 implementation, which returns immediately without polling. The Plugin Framework implementation will replace the SDKv2 implementation in a follow-up version.
* Deprecate the SDKv2 fallback implementations of `databricks_library`, `databricks_quality_monitor`, and `databricks_share` resources, and the `databricks_share`, `databricks_shares`, and `databricks_volumes` data sources. These resources have been served by the Plugin Framework by default since their migration; the SDKv2 implementations remain only as opt-in fallbacks via the `USE_SDK_V2_RESOURCES` / `USE_SDK_V2_DATA_SOURCES` environment variables. Setting either environment variable for any of these names now emits a runtime warning (visible with `TF_LOG=WARN` or higher), and the SDKv2 implementations will be removed in the next major release of the provider.
* Add resource and data sources for `databricks_ai_search_endpoint`.
* Add resource and data sources for `databricks_ai_search_index`.
* Add `clear_cloud_attributes_on_remove` to `databricks_cluster` ([#5812](https://github.com/databricks/terraform-provider-databricks/pull/5812)). When set to `true`, removing a cloud attributes block (`aws_attributes`, `azure_attributes`, `gcp_attributes`) from the configuration clears it instead of the removal being silently suppressed. The flag defaults to `false`, preserving the existing diff-suppression behavior that prevents perpetual drift from platform-returned cloud attribute defaults. Keeping a block, even partially specified, is unaffected; only removing the whole block clears.

### Bug Fixes

* Fix import for jobs with >100 tasks ([#5417](https://github.com/databricks/terraform-provider-databricks/pull/5417)).

### Documentation

* Added an example to `databricks_budget` for creating budgets to control Genie usage costs ([#5817](https://github.com/databricks/terraform-provider-databricks/pull/5817)).

### Exporter

* Generate code in `import.sh` more safely ([#5848](hattps://github.com/databricks/terraform-provider-databricks/issues/5848)).

### Internal Changes
