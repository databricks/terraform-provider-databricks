---
subcategory: "Compute"
---
# databricks_pipeline Resource

Use `databricks_pipeline` to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html). 

## Example Usage

```hcl
resource "databricks_pipeline" "this" {
    name = "Pipeline Name"
    storage = "/test/first-pipeline"
    configuration = {
        key1 = "value1"
        key2 = "value2"
    }

    cluster {
        label = "default"
        num_workers = 2
        custom_tags = {
            cluster_type = "default"
        }
    }

    cluster {
        label = "maintenance"
        num_workers = 1
        custom_tags = {
            cluster_type = "maintenance"
        }
    }

    library {
        maven {
            coordinates = "com.microsoft.azure:azure-eventhubs-spark_2.11:2.3.7"
        }
    }

    filters {
        include = ["com.databricks.include"]
        exclude = ["com.databricks.exclude"]
    }

    continuous = false
}
```

## Argument Reference

The following arguments are required:

* `name` - A user-friendly name for this pipeline. The name can be used to identify pipeline jobs in the UI.
* `storage` - A location on DBFS or cloud storage where output data and metadata required for pipeline execution are stored. By default, tables are stored in a subdirectory of this location.
* `configuration` - An optional list of values to apply to the entire pipeline. Elements must be formatted as key:value pairs.
* `library` block - An array of notebooks containing the pipeline code and required artifacts. Syntax resembles [library](cluster.md#library-configuration-block) configuration block with the addition of special `notebook` type of library.
* `cluster` block - An array of specifications for the [clusters](cluster.md) to run the pipeline. If this is not specified, pipelines will automatically select a default cluster configuration for the pipeline.
* `continuous` - A flag indicating whether to run the pipeline continuously. The default value is `false`.
* `target` - The name of a database for persisting pipeline output data. Configuring the target setting allows you to view and query the pipeline output data from the Databricks UI.

## Import

The resource job can be imported using the id of the pipeline

```bash
$ terraform import databricks_pipeline.this <pipeline-id>
```
