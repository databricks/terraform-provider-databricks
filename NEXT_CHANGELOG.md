# NEXT CHANGELOG

## Release v1.80.0

### New Features and Improvements

* Allow to specify budget policy for `databricks_vector_search_endpoint` [#4707](https://github.com/databricks/terraform-provider-databricks/pull/4707)
* Add `resource_model_serving_provisioned_throughput` for creation of [model serving provisioned throughput](https://docs.databricks.com/aws/en/machine-learning/foundation-model-apis/deploy-prov-throughput-foundation-model-apis) endpoints [#4701](https://github.com/databricks/terraform-provider-databricks/pull/4701)
* Replace DBFS with Unity Catalog resources in [index page](https://registry.terraform.io/providers/databricks/databricks/latest/docs) storage section [#4718](https://github.com/databricks/terraform-provider-databricks/pull/4718) 

### Bug Fixes

 * Support updating all attributes for `databricks_model_serving` ([#4575](https://github.com/databricks/terraform-provider-databricks/pull/4575)).
 * Fix reading of `external_id` for `databricks_service_principal` [#4712](https://github.com/databricks/terraform-provider-databricks/pull/4712)

### Documentation
* Added documentation for GITHUB OIDC authentication type [#4717] (https://github.com/databricks/terraform-provider-databricks/pull/4717)

### Exporter

 * Generate correct code for Databricks and Azure-managed service principals [#4715](https://github.com/databricks/terraform-provider-databricks/pull/4715)

### Internal Changes
