---
subcategory: "MLflow"
---
# databricks_mlflow_models Data Source

Retrieves a list of [databricks_mlflow_model](../resources/mlflow_model.md) objects, that were created by Terraform or manually, so that special handling could be applied.

-> This data source can only be used with a workspace-level provider!

## Example Usage

```hcl
data "databricks_mlflow_models" "this" {}

output "model" {
  value = data.databricks_mlflow_models.this
}
```

```hcl
data "databricks_mlflow_models" "this" {}

check "model_list_not_empty" {
  assert {
    condition     = length(data.databricks_mlflow_models.this.names) != 0
    error_message = "Model list is empty."
  }
}

check "model_list_contains_model" {
  assert {
    condition     = contains(data.databricks_mlflow_models.this.names, "model_1")
    error_message = "model_1 is missing in model list."
  }
}
```

## Attribute Reference

This data source exports the following attributes:

* `names` - List of names of [databricks_mlflow_model](./mlflow_model.md)
