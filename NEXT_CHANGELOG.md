# NEXT CHANGELOG

## Release v1.129.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `databricks_pipeline` dropping `serverless = false` from create/update requests. Because `serverless` is `omitempty` in the SDK and was never force-sent, classic (non-serverless) ingestion pipelines failed with `cannot provide cluster settings when using serverless compute` (the API defaults an omitted `serverless` to `true` for ingestion pipelines, then rejects the `cluster` block).

### Documentation

### Exporter

* Skip system-managed jobs during export and add missing file references for job task parameters ([#5956](https://github.com/databricks/terraform-provider-databricks/issues/5956)).
* Added support for exporting `databricks_endpoint` resource ([#5951](https://github.com/databricks/terraform-provider-databricks/pull/5951)).
* Add an `exporter` dimension to the user agent ([#5954](https://github.com/databricks/terraform-provider-databricks/pull/5954)).
* Allow to generate named variables from references; introduce `databricks_account_id` variable for account-level exports; bug fixes ([#5952](https://github.com/databricks/terraform-provider-databricks/pull/5952)).
* Preserve zero `value` fields when exporting `databricks_workspace_setting_v2` and `databricks_account_setting_v2` ([#5955](https://github.com/databricks/terraform-provider-databricks/issues/5955)).
* Resolve references embedded in `databricks_cluster_policy` definitions instead of emitting hardcoded values ([#5953](https://github.com/databricks/terraform-provider-databricks/issues/5953)).

### Internal Changes
