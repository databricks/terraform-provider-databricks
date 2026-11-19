# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Roll back the `databricks_volume` owner change when an apply that also changes other fields is cancelled while their update is being retried, instead of leaving the new owner on the volume ([#6240](https://github.com/databricks/terraform-provider-databricks/pull/6240)).
* Roll back the owner change of `databricks_catalog`, `databricks_external_location`, `databricks_credential`, `databricks_metastore`, `databricks_registered_model`, `databricks_recipient` and the SDKv2 `databricks_share` (selected with `USE_SDK_V2_RESOURCES`) when an apply that also changes other fields is cancelled while their update is being retried, instead of leaving the new owner on the resource ([#6249](https://github.com/databricks/terraform-provider-databricks/pull/6249)).

### Documentation

### Exporter

### Internal Changes
