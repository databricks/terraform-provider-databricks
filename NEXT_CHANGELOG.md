# NEXT CHANGELOG

## Release v1.65.0

### New Features and Improvements

 * Add multipart permissions to `databricks_aws_unity_catalog_policy` data source ([#4440](https://github.com/databricks/terraform-provider-databricks/pull/4440)).

### Bug Fixes

 * Correctly handle PAT and OBO tokens without expiration ([#4444](https://github.com/databricks/terraform-provider-databricks/pull/4444)).
 * Mark `task.spark_jar_task.run_as_repl` as `computed` to fix configuration drift ([#4452](https://github.com/databricks/terraform-provider-databricks/pull/4452)).

### Documentation

 * Fix attribute name in `databricks_instance_profile` examples ([#4426](https://github.com/databricks/terraform-provider-databricks/pull/4426)).

### Exporter

 * Additional tuning of references in databricks_job ([#4434](https://github.com/databricks/terraform-provider-databricks/pull/4434))

### Internal Changes

 * Started to use the new release framework for releases of the Terraform provider.
