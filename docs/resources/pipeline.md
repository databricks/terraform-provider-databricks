---
subcategory: "Compute"
---
# databricks_pipeline Resource

Use `databricks_pipeline` to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html). 

## Example Usage

```hcl
resource "databricks_notebook" "dlt_demo" {
  #...
}

resource "databricks_pipeline" "this" {
  name    = "Pipeline Name"
  storage = "/test/first-pipeline"
  configuration = {
    key1 = "value1"
    key2 = "value2"
  }

  cluster {
    label       = "default"
    num_workers = 2
    custom_tags = {
      cluster_type = "default"
    }
  }

  cluster {
    label       = "maintenance"
    num_workers = 1
    custom_tags = {
      cluster_type = "maintenance"
    }
  }

  library {
    notebook {
      path = databricks_notebook.dlt_demo.id
    }
  }

  continuous = false
}
```

## Argument Reference

The following arguments are supported:

* `name` - A user-friendly name for this pipeline. The name can be used to identify pipeline jobs in the UI.
* `storage` - A location on DBFS or cloud storage where output data and metadata required for pipeline execution are stored. By default, tables are stored in a subdirectory of this location. *Change of this parameter forces recreation of the pipeline.*
* `configuration` - An optional list of values to apply to the entire pipeline. Elements must be formatted as key:value pairs.
* `library` blocks - Specifies pipeline code and required artifacts. Syntax resembles [library](cluster.md#library-configuration-block) configuration block with the addition of a special `notebook` type of library that should have the `path` attribute. *Right now only the `notebook` type is supported.*
* `cluster` blocks - [Clusters](cluster.md) to run the pipeline. If none is specified, pipelines will automatically select a default cluster configuration for the pipeline. *Please note that DLT pipeline clusters are supporting only subset of attributes as described in [documentation](https://docs.databricks.com/data-engineering/delta-live-tables/delta-live-tables-api-guide.html#pipelinesnewcluster).*  Also, note that `autoscale` block is extended with the `mode` parameter that controls the autoscaling algorithm (possible values are `ENHANCED` for new, enhanced autoscaling algorithm, or `LEGACY` for old algorithm).
* `continuous` - A flag indicating whether to run the pipeline continuously. The default value is `false`.
* `development` - A flag indicating whether to run the pipeline in development mode. The default value is `true`.
* `photon` - A flag indicating whether to use Photon engine. The default value is `false`.
* `target` - The name of a database for persisting pipeline output data. Configuring the target setting allows you to view and query the pipeline output data from the Databricks UI.
* `edition` - optional name of the [product edition](https://docs.databricks.com/data-engineering/delta-live-tables/delta-live-tables-concepts.html#editions). Supported values are: `core`, `pro`, `advanced` (default).
* `channel` - optional name of the release channel for Spark version used by DLT pipeline.  Supported values are: `current` (default) and `preview`.

## Import

The resource job can be imported using the id of the pipeline

```bash
$ terraform import databricks_pipeline.this <pipeline-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
