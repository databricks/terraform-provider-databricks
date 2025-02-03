# NEXT CHANGELOG

## Release v1.65.0

### New Features and Improvements

 * Add resources for AI/BI Dashboards embedding settings: `databricks_aibi_dashboard_embedding_access_policy_setting` and `databricks_aibi_dashboard_embedding_approved_domains_setting` ([#4213](https://github.com/databricks/terraform-provider-databricks/pull/4213)).

### Bug Fixes

 * Correctly handle PAT and OBO tokens without expiration ([#4444](https://github.com/databricks/terraform-provider-databricks/pull/4444)).
 * Mark `task.spark_jar_task.run_as_repl` as `computed` to fix configuration drift ([#4452](https://github.com/databricks/terraform-provider-databricks/pull/4452)).

### Documentation

 * Fix attribute name in `databricks_instance_profile` examples ([#4426](https://github.com/databricks/terraform-provider-databricks/pull/4426)).
 * Remove mention that databricks_credential is storage-only on GCP ([#4460](https://github.com/databricks/terraform-provider-databricks/pull/4460)).

### Exporter

 * Additional tuning of references in databricks_job ([#4434](https://github.com/databricks/terraform-provider-databricks/pull/4434))

### Internal Changes

 * Started to use the new release framework for releases of the Terraform provider.
 * Refactored existing integration tests from the `internal/acceptance` package to the package corresponding to the resource under test. See CONTRIBUTING.md for more information on how to run the tests.
