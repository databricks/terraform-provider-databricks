---
subcategory: "MLflow"
---
# databricks_mlflow_experiment Resource

This resource allows you to create MLflow experiments in Databricks.

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
