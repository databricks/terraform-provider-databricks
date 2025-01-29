# NEXT CHANGELOG

## Release v1.65.0

### New Features and Improvements

### Bug Fixes

 * Make removing the `config` attribute of `databricks_model_serving` a no-op ([#4446](https://github.com/databricks/terraform-provider-databricks/pull/4446)).
   
   This change allows integrations, such as DABs, to manage & update the configuration of a model serving endpoint independently of the lifecycle of the endpoint itself.

### Documentation

 * Fix attribute name in `databricks_instance_profile` examples ([#4426](https://github.com/databricks/terraform-provider-databricks/pull/4426)).

### Exporter

 * Additional tuning of references in databricks_job ([#4434](https://github.com/databricks/terraform-provider-databricks/pull/4434))

### Internal Changes

 * Started to use the new release framework for releases of the Terraform provider.
