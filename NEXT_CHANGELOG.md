# NEXT CHANGELOG

## Release v1.118.0

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_ai_search_endpoint`.
* Add resource and data sources for `databricks_ai_search_index`.


### Bug Fixes

* Fix spurious `account_id` drift on `databricks_mws_ncc_private_endpoint_rule` ([#5347](https://github.com/databricks/terraform-provider-databricks/issues/5347)). The backend echoes `account_id` on read; the schema previously marked it as a plain `Optional` attribute, so once it landed in state (for example via `terraform import`) the next plan reported `account_id = "..." -> null` and a subsequent apply failed with `cannot update mws ncc private endpoint rule: Update mask must be specified.`. Marking `account_id` as `Computed` (matching the sibling `databricks_mws_network_connectivity_config` resource) preserves the server-provided value across refreshes and eliminates the spurious in-place update.
* Fixed `databricks_mws_workspaces` failing to update `private_access_settings_id` and other fields on GCP workspaces ([#5430](https://github.com/databricks/terraform-provider-databricks/issues/5430)).

### Documentation
* Added `disabled` field to `task` block in `databricks_job` resource, allowing individual tasks to be disabled ([#5767](https://github.com/databricks/terraform-provider-databricks/pull/5767)).

### Exporter

* Rewrote Exporter logging so it works with Databricks Go SDK logging ([#5805](https://github.com/databricks/terraform-provider-databricks/pull/5805)).

### Internal Changes
