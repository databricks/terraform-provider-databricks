---
subcategory: "Unity Catalog"
---
# databricks_registered_model Resource

-> **Note** This resource could be only used with workspace-level provider!

This resource allows you to create [Models in Unity Catalog](https://docs.databricks.com/en/mlflow/models-in-uc.html) in Databricks.

## Example Usage

```hcl
resource "databricks_registered_model" "this" {
  name         = "my_model"
  catalog_name = "main"
  schema_name  = "default"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the registered model.
* `catalog_name` - (Required) The name of the catalog where the schema and the registered model reside.
* `schema_name` - (Required) The name of the schema where the registered model resides.
* `comment` - The comment attached to the registered model.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Equal to the full name of the model (`catalog_name.schema_name.name`) and used to identify the model uniquely across the metastore.

## Access Control

* [databricks_grants](grants.md#registered-model-grants) can be used to grant principals `ALL_PRIVILEGES`, `APPLY_TAG`, and `EXECUTE` privileges.

## Import

The registered model resource can be imported using the full (3-level) name of the model.

```bash
$ terraform import databricks_registered_model.this <catalog_name.schema_name.model_name>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_model_serving](model_serving.md) to serve this model on a Databricks serving endpoint.
* [databricks_mlflow_experiment](mlflow_experiment.md) to manage [MLflow experiments](https://docs.databricks.com/data/data-sources/mlflow-experiment.html) in Databricks.
* [databricks_table](tables.md) data to manage tables within Unity Catalog.
* [databricks_schema](schemas.md) data to manage schemas within Unity Catalog.
* [databricks_catalog](catalogs.md) data to manage catalogs within Unity Catalog.
