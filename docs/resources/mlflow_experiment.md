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

* `name` - (Required) Name of MLflow experiment. It must be an absolute path within the Databricks workspace, e.g. `/Users/<some-username>/my-experiment`. For more information about changes to experiment naming conventions, see [mlflow docs](https://docs.databricks.com/applications/mlflow/experiments.html#experiment-migration).
* `artifact_location` - Path to dbfs:/ or s3:// artifact location of the MLflow experiment.
* `tags` - Tags for the MLflow experiment.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the MLflow experiment.

## Access Control

* [databricks_permissions](permissions.md#MLflow-Experiment-usage) can control which groups or individual users can *Read*, *Edit*, or *Manage* individual experiments.

## Import

The experiment resource can be imported using the id of the experiment

```bash
terraform import databricks_mlflow_experiment.this <experiment-id>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_registered_model](registered_model.md) to create [Models in Unity Catalog](https://docs.databricks.com/en/mlflow/models-in-uc.html) in Databricks.
* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_directory](directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_mlflow_model](mlflow_model.md) to create models in the [workspace model registry](https://docs.databricks.com/en/mlflow/model-registry.html) in Databricks.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_notebook](../data-sources/notebook.md) data to export a notebook from Databricks Workspace.
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
