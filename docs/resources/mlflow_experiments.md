---
subcategory: "MLflow"
---
# databricks_mlflow_experiment Resource

This resource allows you to create MLflow experiments in Databricks.

## Example Usage

```hcl
resource "databricks_mlflow_experiment" "test" {
  name = "My MLflow Experiment"

  description = "My MLflow experiment description"

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

* `name` - (Required) Name of MLflow experiment.
* `description` - The description of the MLflow experiment.
* `tags` - Tags for the MLflow experiment.