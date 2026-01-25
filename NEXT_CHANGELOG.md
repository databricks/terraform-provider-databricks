# NEXT CHANGELOG

## Release v1.105.0

### Breaking Changes

* Return empty string for `data_source_id` in `databricks_sql_warehouse` and `databricks_sql_endpoint` if data source API call fails ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### New Features and Improvements

* Added support for `azure-openai` provider in `databricks_model_serving` resource with a dedicated `azure_openai_config` block. This allows for explicit configuration of Azure OpenAI endpoints, including `openai_api_base`, `openai_api_version`, and `openai_deployment_name` ([#5310](https://github.com/databricks/terraform-provider-databricks/pull/5310)).

### Bug Fixes

* Correctly transform Azure OpenAI endpoints during import in `databricks_model_serving` to prevent immediate drift ([#5310](https://github.com/databricks/terraform-provider-databricks/pull/5310)).
* Added deprecation warning when using legacy `openai_config` with Azure OpenAI settings in `databricks_model_serving` ([#5310](https://github.com/databricks/terraform-provider-databricks/pull/5310)).

### Documentation

* Mark `data_source_id` as deprecated in `databricks_sql_warehouse` and `databricks_sql_endpoint` ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### Exporter

### Internal Changes
