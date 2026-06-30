# NEXT CHANGELOG

## Release v1.120.0

### Important Changes

* [Backward Incompatible] Resource `databricks_mws_ncc_private_endpoint_rule` has been refactored in this version to fix multiple bugs in its implementation and behavior. Server-populated attributes are now strictly read-only (writes that previously produced perpetual plan drift now error at plan time), plan churn on previously-Optional output fields is eliminated, and the new Plugin Framework implementation (behind `DATABRICKS_TF_ENABLED_PF_RESOURCES`) polls the rule's `connection_state` during Create so a failed or unusable provisioning state surfaces at apply time instead of on the next plan. See `Breaking Changes` and `New Features and Improvements` below for specifics.

### Breaking Changes

* Tighten read-only attributes on `databricks_mws_ncc_private_endpoint_rule` ([#5819](https://github.com/databricks/terraform-provider-databricks/pull/5819)). The attributes `rule_id`, `account_id`, `endpoint_name`, `vpc_endpoint_id`, `connection_state`, `creation_time`, `updated_time`, `deactivated`, `deactivated_at`, `error_message`, and `gcp_endpoint.psc_endpoint_uri` are now computed-only and can no longer be set in HCL. They are populated by the server in every API response (and `psc_endpoint_uri` is unconditionally overwritten from the cloud-platform truth on every read), and previous releases accepted writes to these attributes silently while the API ignored or overwrote the values, producing perpetual drift on every plan. Configurations that explicitly assigned any of these attributes must remove the assignment; the value is still available in state after apply.

### New Features and Improvements

* A Plugin Framework implementation of `databricks_mws_ncc_private_endpoint_rule` is now available behind `DATABRICKS_TF_ENABLED_PF_RESOURCES=databricks_mws_ncc_private_endpoint_rule` ([#5819](https://github.com/databricks/terraform-provider-databricks/pull/5819)). The default remains the SDKv2 implementation; no HCL or state changes when opting in. Once opted in, `terraform apply` waits for the rule to leave `CREATING` before returning: `PENDING` and `ESTABLISHED` succeed, while a `CREATE_FAILED`, `REJECTED`, `DISCONNECTED`, or `EXPIRED` connection state surfaces as an apply-time error instead of on the next plan. This differs from the SDKv2 implementation, which returns immediately without polling.

  The create API does not accept `enabled` (rules are always created disabled), so a configuration that sets `enabled = true` reconciles the value through a follow-up update and may need a second `terraform apply` to converge. The backend rejects updates to `enabled` for Azure rules outright. Both behaviors also apply to the SDKv2 implementation.
* Deprecate the SDKv2 fallback implementations of `databricks_library`, `databricks_quality_monitor`, and `databricks_share` resources, and the `databricks_share`, `databricks_shares`, and `databricks_volumes` data sources. These resources have been served by the Plugin Framework by default since their migration; the SDKv2 implementations remain only as opt-in fallbacks via the `USE_SDK_V2_RESOURCES` / `USE_SDK_V2_DATA_SOURCES` environment variables. Setting either environment variable for any of these names now emits a runtime warning (visible with `TF_LOG=WARN` or higher), and the SDKv2 implementations will be removed in the next major release of the provider.

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).
* Allow setting `user_api_scopes = []` on `databricks_app` to disable OBO (On-Behalf-Of) user authorization ([#5834](https://github.com/databricks/terraform-provider-databricks/pull/5834)).

  The Apps API omits `user_api_scopes` from its response when OBO is inactive, so a configured empty list previously failed with `Provider produced inconsistent result after apply`. The provider now preserves a configured empty list in state, mirroring the reconciliation used by `databricks_app_space`.

### Documentation

### Exporter

### Internal Changes

* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
