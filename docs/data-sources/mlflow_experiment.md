---
subcategory: "MLflow"
---
# databricks_mlflow_experiment Data Source

Retrieves the settings of [databricks_mlflow_experiment](../resources/mlflow_experiment.md) by id or name.

-> This data source can only be used with a workspace-level provider!

## Example Usage

```hcl
data "databricks_mlflow_experiment" "this" {
  experiment_id = "1234567890"
}

data "databricks_mlflow_experiment" "this" {
  name = "/Users/databricks/my-experiment"
}
```

## Argument Reference

* `experiment_id` - (Required if `name` isn't specified) Unique identifier for the experiment.
* `name` - (Required if `experiment_id` isn't specified) Path to experiment.

## Attribute Reference

This data source exports the following attributes:

* `artifact_location` - Location where artifacts for the experiment are stored.
* `creation_time` - Creation time in unix time stamp.
* `experiment_id` - Unique identifier for the experiment. (same as `id`)
* `id` - Unique identifier for the experiment. (same as `experiment_id`)
* `last_update_time` - Last update time in unix time stamp.
* `lifecycle_stage` - Current life cycle stage of the experiment: `active` or `deleted`.
* `name` - Path to experiment.
* `tags` - Additional metadata key-value pairs.
