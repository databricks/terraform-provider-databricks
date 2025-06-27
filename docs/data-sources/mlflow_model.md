---
subcategory: "MLflow"
---
# databricks_mlflow_model Data Source

Retrieves the settings of [databricks_mlflow_model](../resources/mlflow_model.md) by name.

-> This data source can only be used with a workspace-level provider!

## Example Usage

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

```hcl
data "databricks_mlflow_model" "this" {
  name = "My MLflow Model with multiple versions"
}

resource "databricks_model_serving" "this" {
  name = "model-serving-endpoint"
  config {
    served_models {
      name                  = "model_serving_prod"
      model_name            = data.databricks_mlflow_model.this.name
      model_version         = data.databricks_mlflow_model.this.latest_versions[0].version
      workload_size         = "Small"
      scale_to_zero_enabled = true
    }
  }
}
```

## Argument Reference

* `name` - (Required) Name of the registered model.

## Attribute Reference

This data source exports the following attributes:

* `model` - Model object
  * `description` - User-specified description for the object.
  * `id` - Unique identifier for the object.
  * `latest_versions` - Array of model versions, each the latest version for its stage.
  * `name` - Name of the model.
  * `permission_level` - Permission level of the requesting user on the object. For what is allowed at each level, see MLflow Model permissions.
  * `tags` - Array of tags associated with the model.
  * `user_id` - The username of the user that created the object.
