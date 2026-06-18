# NEXT CHANGELOG

## Release v1.119.0

### Important Changes

* [Backward Incompatible] Resource `databricks_mws_ncc_private_endpoint_rule` has been refactored in this version to fix multiple bugs in its implementation and behavior. Server-populated attributes are now strictly read-only (writes that previously produced perpetual plan drift now error at plan time), plan churn on previously-Optional output fields is eliminated, and the new Plugin Framework implementation (behind `DATABRICKS_TF_ENABLED_PF_RESOURCES`) polls the rule's `connection_state` during Create so a failed or unusable provisioning state surfaces at apply time instead of on the next plan. See `Breaking Changes` and `New Features and Improvements` below for specifics.

### Breaking Changes

* Tighten read-only attributes on `databricks_mws_ncc_private_endpoint_rule` ([#5819](https://github.com/databricks/terraform-provider-databricks/pull/5819)). The attributes `rule_id`, `account_id`, `endpoint_name`, `vpc_endpoint_id`, `connection_state`, `creation_time`, `updated_time`, `deactivated`, `deactivated_at`, `error_message`, and `gcp_endpoint.psc_endpoint_uri` are now computed-only and can no longer be set in HCL. They are populated by the server in every API response (and `psc_endpoint_uri` is unconditionally overwritten from the cloud-platform truth on every read), and previous releases accepted writes to these attributes silently while the API ignored or overwrote the values, producing perpetual drift on every plan. Configurations that explicitly assigned any of these attributes must remove the assignment; the value is still available in state after apply.

### New Features and Improvements

* A Plugin Framework implementation of `databricks_mws_ncc_private_endpoint_rule` is now available behind `DATABRICKS_TF_ENABLED_PF_RESOURCES=databricks_mws_ncc_private_endpoint_rule` ([#5819](https://github.com/databricks/terraform-provider-databricks/pull/5819)). The default remains the SDKv2 implementation; no HCL or state changes when opting in. Once opted in, `terraform apply` waits for the rule to leave `CREATING` before returning: `PENDING` and `ESTABLISHED` succeed, while a `CREATE_FAILED`, `REJECTED`, `DISCONNECTED`, or `EXPIRED` connection state surfaces as an apply-time error instead of on the next plan. This differs from the SDKv2 implementation, which returns immediately without polling.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
