---
subcategory: "MLflow"
---
# databricks_mlflow_experiment Resource

[API Documentation](https://docs.databricks.com/api/workspace/experiments)

This resource allows you to manage [MLflow experiments](https://docs.databricks.com/data/data-sources/mlflow-experiment.html) in Databricks.

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
data "databricks_current_user" "me" {}

resource "databricks_mlflow_experiment" "this" {
  name              = "${data.databricks_current_user.me.home}/Sample"
  artifact_location = "s3://bucket/my-experiment"

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

```hcl
# Store this experiment's traces in a Unity Catalog schema
resource "databricks_mlflow_experiment" "with_uc_traces" {
  name = "${data.databricks_current_user.me.home}/uc-traces-experiment"

  trace_location {
    uc_trace_location {
      catalog      = "my_catalog"
      schema       = "my_schema"
      table_prefix = "my_experiment"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of MLflow experiment. It must be an absolute path within the Databricks workspace, e.g. `/Users/<some-username>/my-experiment`. For more information about changes to experiment naming conventions, see [mlflow docs](https://docs.databricks.com/applications/mlflow/experiments.html#experiment-migration).
* `artifact_location` - Path to artifact location of the MLflow experiment.
* `tags` - Tags for the MLflow experiment.
* `trace_location` - (Optional, Computed, Immutable) Unity Catalog location where the experiment's traces are stored. Cannot be changed after the experiment is created; changing it forces replacement of the experiment. Omitting the block for an experiment that already has a location leaves the existing location in place (it is read back from the server) rather than forcing replacement. This block consists of the following fields:
  * `uc_trace_location` - (Required) The Unity Catalog storage location. This block consists of the following fields:
    * `catalog` - (Required) Name of the Unity Catalog catalog.
    * `schema` - (Required) Name of the Unity Catalog schema within `catalog`.
    * `table_prefix` - (Optional) Prefix for the generated trace tables (named `{catalog}.{schema}.{table_prefix}_otel_*`). If omitted, the server generates a default prefix derived from the experiment ID; the field then stays empty and the resolved value is available in `effective_table_prefix`.
    * `effective_table_prefix` - (Computed) The trace-table prefix actually in effect: `table_prefix` if it was set on creation, otherwise the server-generated default.
* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the MLflow experiment.

## Access Control

* [databricks_permissions](permissions.md#MLflow-Experiment-usage) can control which groups or individual users can *Read*, *Edit*, or *Manage* individual experiments.

## Import

The experiment resource can be imported using the id of the experiment:

```hcl
import {
  to = databricks_mlflow_experiment.this
  id = "<experiment-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_mlflow_experiment.this "<experiment-id>"
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
