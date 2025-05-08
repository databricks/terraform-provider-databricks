---
subcategory: "Unity Catalog"
---
# databricks_registered_model Resource

This resource allows you to create [Models in Unity Catalog](https://docs.databricks.com/en/mlflow/models-in-uc.html) in Databricks.

-> This resource can only be used with a workspace-level provider!

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

* `name` - (Required) The name of the registered model.  *Change of this parameter forces recreation of the resource.*
* `catalog_name` - (Required) The name of the catalog where the schema and the registered model reside. *Change of this parameter forces recreation of the resource.*
* `schema_name` - (Required) The name of the schema where the registered model resides. *Change of this parameter forces recreation of the resource.*
* `owner` - (Optional) Name of the registered model owner.
* `comment` - (Optional) The comment attached to the registered model.
* `storage_location` - (Optional) The storage location under which model version data files are stored. *Change of this parameter forces recreation of the resource.*

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Equal to the full name of the model (`catalog_name.schema_name.name`) and used to identify the model uniquely across the metastore.

## Access Control

* [databricks_grants](grants.md#registered-model-grants) can be used to grant principals `ALL_PRIVILEGES`, `APPLY_TAG`, and `EXECUTE` privileges.

## Import

The registered model resource can be imported using the full (3-level) name of the model.

```hcl
import {
  to = databricks_registered_model.this
  id = "<catalog_name>.<schema_name>.<model_name>"
}
```

Alternatively, when using `terraform` version 1.5 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_registered_model.this <catalog_name>.<schema_name>.<model_name>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_model_serving](model_serving.md) to serve this model on a Databricks serving endpoint.
* [databricks_mlflow_experiment](mlflow_experiment.md) to manage [MLflow experiments](https://docs.databricks.com/data/data-sources/mlflow-experiment.html) in Databricks.
* [databricks_schema](schema.md) to manage schemas within Unity Catalog.
* [databricks_catalog](catalog.md) to manage catalogs within Unity Catalog.
