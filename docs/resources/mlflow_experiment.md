---
subcategory: "MLflow"
---
# databricks_mlflow_experiment Resource

This resource allows you to manage [MLflow experiments](https://docs.databricks.com/data/data-sources/mlflow-experiment.html) in Databricks.

## Example Usage

```hcl
data "databricks_current_user" "me" {}

resource "databricks_mlflow_experiment" "this" {
  name              = "${data.databricks_current_user.me.home}/Sample"
  artifact_location = "dbfs:/tmp/my-experiment"
  description       = "My MLflow experiment description"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of MLflow experiment. It must be an absolute path within the Databricks workspace, e.g. `/Users/<some-username>/my-experiment`. For more information about changes to experiment naming conventions, see [mlflow docs](https://docs.databricks.com/applications/mlflow/experiments.html#experiment-migration).
* `artifact_location` - Path to dbfs:/ or s3:// artifact location of the MLflow experiment.
* `description` - The description of the MLflow experiment.

## Access Control

* [databricks_permissions](permissions.md#MLflow-Experiment-usage) can control which groups or individual users can *Read*, *Edit*, or *Manage* individual experiments.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_directory](directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_mlflow_experiment](mlflow_experiment.md) to manage [MLflow experiments](https://docs.databricks.com/data/data-sources/mlflow-experiment.html) in Databricks.
* [databricks_mlflow_model](mlflow_model.md) to create [MLflow models](https://docs.databricks.com/applications/mlflow/models.html) in Databricks.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_notebook](../data-sources/notebook.md) data to export a notebook from Databricks Workspace.
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
