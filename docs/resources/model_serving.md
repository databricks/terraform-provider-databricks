---
subcategory: "Serving"
---
# databricks_model_serving Resource

This resource allows you to manage [Model Serving](https://docs.databricks.com/machine-learning/model-serving/index.html) endpoints in Databricks.

## Example Usage

```hcl
resource "databricks_model_serving" "this" {
  name = "ads-serving-endpoint"
  config {
    served_models {
      name                  = "prod_model"
      model_name            = "ads-model"
      model_version         = "2"
      workload_size         = "Small"
      scale_to_zero_enabled = true
    }
    served_models {
      name                  = "candidate_model"
      model_name            = "ads-model"
      model_version         = "4"
      workload_size         = "Small"
      scale_to_zero_enabled = false
    }
    traffic_config {
      routes {
        served_model_name  = "prod_model"
        traffic_percentage = 90
      }
      routes {
        served_model_name  = "candidate_model"
        traffic_percentage = 10
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the model serving endpoint. This field is required and must be unique across a workspace. An endpoint name can consist of alphanumeric characters, dashes, and underscores. NOTE: Changing this name will delete the existing endpoint and create a new endpoint with the update name.
* `config` - (Required) The model serving endpoint configuration.

### config Configuration Block

* `served_models` - (Required) Each block represents a served model for the endpoint to serve. A model serving endpoint can have up to 10 served models.
* `traffic_config` - A single block represents the traffic split configuration amongst the served models.

### served_models Configuration Block

* `name` - The name of a served model. It must be unique across an endpoint. If not specified, this field will default to `modelname-modelversion`. A served model name can consist of alphanumeric characters, dashes, and underscores.
* `model_name` - (Required) The name of the model in Databricks Model Registry to be served.
* `model_version` - (Required) The version of the model in Databricks Model Registry to be served.
* `workload_size` - (Required) The workload size of the served model. The workload size corresponds to a range of provisioned concurrency that the compute will autoscale between. A single unit of provisioned concurrency can process one request at a time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency), "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64 provisioned concurrency).
* `scale_to_zero_enabled` - Whether the compute resources for the served model should scale down to zero. If scale-to-zero is enabled, the lower bound of the provisioned concurrency for each workload size will be 0. The default value is `true`.

### traffic_config Configuration Block

* `routes` - (Required) Each block represents a route that defines traffic to each served model. Each `served_models` block needs to have a corresponding `routes` block

### routes Configuration Block

* `served_model_name` - (Required) The name of the served model this route configures traffic for. This needs to match the name of a `served_models` block
* `traffic_percentage` - (Required) The percentage of endpoint traffic to send to this route. It must be an integer between 0 and 100 inclusive.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Equal to the `name` argument and used to identify the serving endpoint.
* `serving_endpoint_id` - Unique identifier of the serving endpoint primarily used to set permissions and refer to this instance for other operations.

## Access Control

* [databricks_permissions](permissions.md#model-serving-usage) can control which groups or individual users can *Manage*, *Query* or *View* individual serving endpoints.

## Timeouts

The `timeouts` block allows you to specify `create` and `update` timeouts. The default right now is 45 minutes for both operations.

```hcl
timeouts {
  create = "30m"
}
```

## Import

The model serving resource can be imported using the name of the endpoint.

```bash
$ terraform import databricks_model_serving.this <model-serving-endpoint-name>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_directory](directory.md) to manage directories in [Databricks Workspace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_mlflow_model](mlflow_model.md) to create [MLflow models](https://docs.databricks.com/applications/mlflow/models.html) in Databricks.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_notebook](../data-sources/notebook.md) data to export a notebook from Databricks Workspace.
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
