---
subcategory: "MLflow"
---
# databricks_mlflow_model Resource

This resource allows you to create MLflow models in Databricks.

## Example Usage

```hcl
resource "databricks_mlflow_model" "test" {
  name = "My MLflow Model"

  description = "My MLflow model description"

  tags {
    key   = "key1"
    value = "value1"
  }
  tags {
    key   = "key2"
    value = "value2"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of MLflow model. Change of name triggers new resource.
* `description` - The description of the MLflow model.
* `tags` - Tags for the MLflow model.

## Access Control

* [databricks_permissions](permissions.md#MLflow-Model-usage) can control which groups or individual users can *Read*, *Edit*, *Manage Staging Versions*, *Manage Production Versions*, and *Manage* individual models.
