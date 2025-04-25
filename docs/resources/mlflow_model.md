---
subcategory: "MLflow"
---
# databricks_mlflow_model Resource

This resource allows you to create [MLflow models](https://docs.databricks.com/applications/mlflow/models.html) in Databricks.

-> This resource can only be used with a workspace-level provider!

-> This documentation covers the Workspace Model Registry. Databricks recommends using [Models in Unity Catalog](registered_model.md). Models in Unity Catalog provides centralized model governance, cross-workspace access, lineage, and deployment.

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

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the MLflow model, the same as `name`.

## Import

The model resource can be imported using the name

```bash
terraform import databricks_mlflow_model.this <name>
```

## Access Control

* [databricks_permissions](permissions.md#MLflow-Model-usage) can control which groups or individual users can *Read*, *Edit*, *Manage Staging Versions*, *Manage Production Versions*, and *Manage* individual models.

## Related Resources

The following resources are often used in the same context:

* [databricks_registered_model](registered_model.md) to create [Models in Unity Catalog](https://docs.databricks.com/en/mlflow/models-in-uc.html) in Databricks.
* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_model_serving](model_serving.md) to serve this model on a Databricks serving endpoint.
* [databricks_directory](directory.md) to manage directories in [Databricks Workspace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_mlflow_experiment](mlflow_experiment.md) to manage [MLflow experiments](https://docs.databricks.com/data/data-sources/mlflow-experiment.html) in Databricks.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_notebook](../data-sources/notebook.md) data to export a notebook from Databricks Workspace.
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
