---
subcategory: "Unity Catalog"
---
# databricks_registered_model Data Source

-> This resource can only be used with a workspace-level provider!

This resource allows you to get information about [Model in Unity Catalog](https://docs.databricks.com/en/mlflow/models-in-uc.html) in Databricks.

## Example Usage

```hcl
data "databricks_registered_model" "this" {
  full_name = "main.default.my_model"
}
```

## Argument Reference

The following arguments are supported:

* `full_name` - (Required, String) The fully-qualified name of the registered model (`catalog_name.schema_name.name`).
* `include_aliases` - (Optional, Boolean) flag to specify if list of aliases should be included into output.

## Attribute Reference

The following attributes are exported:

* `model_info` - block with information about the model in Unity Catalog:
  * `aliases` - the list of aliases associated with this model. Each item is object consisting of following attributes:
    * `alias_name` - string with the name of alias
    * `version_num` - associated model version
  * `catalog_name` - The name of the catalog where the schema and the registered model reside.
  * `comment` - The comment attached to the registered model.
  * `created_at` - the Unix timestamp at the model's creation
  * `created_by` - the identifier of the user who created the model
  * `full_name` - The fully-qualified name of the registered model (`catalog_name.schema_name.name`).
  * `metastore_id` - the unique identifier of the metastore
  * `name` - The name of the registered model.
  * `owner` - Name of the registered model owner.
  * `schema_name` - The name of the schema where the registered model resides.
  * `storage_location` - The storage location under which model version data files are stored.
  * `updated_at` - the timestamp of the last time changes were made to the model
  * `updated_by` - the identifier of the user who updated the model last time

## Related Resources

The following resources are often used in the same context:

* [databricks_registered_model](../resources/schema.md) resource to manage models within Unity Catalog.
* [databricks_model_serving](../resources/model_serving.md) to serve this model on a Databricks serving endpoint.
* [databricks_mlflow_experiment](../resources/mlflow_experiment.md) to manage [MLflow experiments](https://docs.databricks.com/data/data-sources/mlflow-experiment.html) in Databricks.
