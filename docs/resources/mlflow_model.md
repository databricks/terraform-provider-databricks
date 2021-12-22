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

* `name` - (Required) Name of MLflow model.
* `description` - The description of the MLflow model.
* `tags` - Tags for the MLflow model.
