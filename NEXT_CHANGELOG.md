# NEXT CHANGELOG

## Release v1.65.0

### New Features and Improvements

 * Add resources for AI/BI Dashboards embedding settings: `databricks_aibi_dashboard_embedding_access_policy_setting` and `databricks_aibi_dashboard_embedding_approved_domains_setting` ([#4213](https://github.com/databricks/terraform-provider-databricks/pull/4213)).
 * Update Databricks Go SDK to v0.57.0.

### Bug Fixes

 * Make removing the `config` attribute of `databricks_model_serving` a no-op ([#4446](https://github.com/databricks/terraform-provider-databricks/pull/4446)).
   
   This change allows integrations, such as DABs, to manage & update the configuration of a model serving endpoint independently of the lifecycle of the endpoint itself.
 * Correctly handle PAT and OBO tokens without expiration ([#4444](https://github.com/databricks/terraform-provider-databricks/pull/4444)).
 * Mark `task.spark_jar_task.run_as_repl` as `computed` to fix configuration drift ([#4452](https://github.com/databricks/terraform-provider-databricks/pull/4452)).

### Documentation

 * Fix attribute name in `databricks_instance_profile` examples ([#4426](https://github.com/databricks/terraform-provider-databricks/pull/4426)).
 * Remove mention that databricks_credential is storage-only on GCP ([#4460](https://github.com/databricks/terraform-provider-databricks/pull/4460)).
 * Officially document `databricks_table` as deprecated. Users of this resource should migrate to `databricks_sql_table`. See https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_table for more information.
 * Improve examples and fix headers layout for `databricks_job` ([#4455](https://github.com/databricks/terraform-provider-databricks/pull/4455)).

### Exporter


### Internal Changes

 * Refactored existing integration tests from the `internal/acceptance` package to the package corresponding to the resource under test. See `CONTRIBUTING.md` for more information on how to run the tests.
