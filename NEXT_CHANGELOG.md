# NEXT CHANGELOG

## Release v1.138.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Retry deleting `databricks_mws_networks`, `databricks_mws_credentials`, `databricks_mws_storage_configurations` and `databricks_mws_customer_managed_keys` for up to 5 minutes while the account API still reports them as attached to a workspace that was just deleted, so `terraform destroy` doesn't need to be run twice ([#6062](https://github.com/databricks/terraform-provider-databricks/pull/6062)).

### Documentation

### Exporter

### Internal Changes
