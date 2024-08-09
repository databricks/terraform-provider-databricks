---
subcategory: "MLflow"
---
# databricks_mlflow_models Data Source

-> **Note** This data source could be only used with workspace-level provider!

Retrieves a list of [databricks_mlflow_model](../resources/mlflow_model.md) objects, that were created by Terraform or manually, so that special handling could be applied.

## Example Usage

```hcl
data "databricks_mlflow_models" "this" {}

output "model" {
  value = data.databricks_mlflow_models.this
}
```

## Attribute Reference

This data source exports the following attributes:

* `names` - List of names of [databricks_mlflow_model](./mlflow_model.md)