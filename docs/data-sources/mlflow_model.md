---
subcategory: "MLflow"
---
# databricks_mlflow_model Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves the settings of [databricks_mlflow_model](../resources/mlflow_model.md) by name.

## Example Usage

```hcl
data "databricks_mlflow_model" "this" {
  name = "My MLflow Model"
}

data "databricks_mlflow_model" "this" {
  name = "My MLflow Model"
  version = "1"
}
```

```hcl
resource "databricks_mlflow_model" "this" {
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

data "databricks_mlflow_model" "this" {
  depends_on = [databricks_mlflow_model.this]
  name       = "My MLflow Model"
}

output "model" {
  value = data.databricks_mlflow_model.this
}
```

## Argument Reference

* `name` - (Required) Name of the registered model.
* `version` - (Optional) Model version number.

## Attribute Reference

This data source exports the following attributes:

* `name` - Name of the registered model.
* `version` - Model version number.
* `model_versions` - Array of all model versions.