---
subcategory: "MLflow"
---
# databricks_mlflow_model Resource

This resource allows you to create [MLflow models](https://docs.databricks.com/applications/mlflow/models.html) in Databricks.

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

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_directory](directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_mlflow_experiment](mlflow_experiment.md) to manage [MLflow experiments](https://docs.databricks.com/data/data-sources/mlflow-experiment.html) in Databricks.
* [databricks_mlflow_model](mlflow_model.md) to create [MLflow models](https://docs.databricks.com/applications/mlflow/models.html) in Databricks.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_notebook](../data-sources/notebook.md) data to export a notebook from Databricks Workspace.
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
