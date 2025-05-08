---
subcategory: "Compute"
---
# databricks_pipeline Resource

Use `databricks_pipeline` to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html).

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_notebook" "dlt_demo" {
  #...
}

resource "databricks_repo" "dlt_demo" {
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

  library {
    file {
      path = "${databricks_repo.dlt_demo.path}/pipeline.sql"
    }
  }

  continuous = false

  notification {
    email_recipients = ["user@domain.com", "user1@domain.com"]
    alerts = [
      "on-update-failure",
      "on-update-fatal-failure",
      "on-update-success",
      "on-flow-failure"
    ]
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - A user-friendly name for this pipeline. The name can be used to identify pipeline jobs in the UI.
* `storage` - A location on DBFS or cloud storage where output data and metadata required for pipeline execution are stored. By default, tables are stored in a subdirectory of this location. *Change of this parameter forces recreation of the pipeline.* (Conflicts with `catalog`).
* `configuration` - An optional list of values to apply to the entire pipeline. Elements must be formatted as key:value pairs.
* `library` blocks - Specifies pipeline code and required artifacts. Syntax resembles [library](cluster.md#library-configuration-block) configuration block with the addition of a special `notebook` & `file` library types that should have the `path` attribute. *Right now only the `notebook` & `file` types are supported.*
* `cluster` blocks - [Clusters](cluster.md) to run the pipeline. If none is specified, pipelines will automatically select a default cluster configuration for the pipeline. *Please note that DLT pipeline clusters are supporting only subset of attributes as described in [documentation](https://docs.databricks.com/data-engineering/delta-live-tables/delta-live-tables-api-guide.html#pipelinesnewcluster).*  Also, note that `autoscale` block is extended with the `mode` parameter that controls the autoscaling algorithm (possible values are `ENHANCED` for new, enhanced autoscaling algorithm, or `LEGACY` for old algorithm).
* `continuous` - A flag indicating whether to run the pipeline continuously. The default value is `false`.
* `development` - A flag indicating whether to run the pipeline in development mode. The default value is `false`.
* `photon` - A flag indicating whether to use Photon engine. The default value is `false`.
* `serverless` - An optional flag indicating if serverless compute should be used for this DLT pipeline.  Requires `catalog` to be set, as it could be used only with Unity Catalog.
* `catalog` - The name of catalog in Unity Catalog. *Change of this parameter forces recreation of the pipeline.* (Conflicts with `storage`).
* `target` - (Optional, String, Conflicts with `schema`) The name of a database (in either the Hive metastore or in a UC catalog) for persisting pipeline output data. Configuring the target setting allows you to view and query the pipeline output data from the Databricks UI.
* `schema` - (Optional, String, Conflicts with `target`) The default schema (database) where tables are read from or published to. The presence of this attribute implies that the pipeline is in direct publishing mode. 
* `edition` - optional name of the [product edition](https://docs.databricks.com/data-engineering/delta-live-tables/delta-live-tables-concepts.html#editions). Supported values are: `CORE`, `PRO`, `ADVANCED` (default).  Not required when `serverless` is set to `true`.
* `channel` - optional name of the release channel for Spark version used by DLT pipeline.  Supported values are: `CURRENT` (default) and `PREVIEW`.
* `budget_policy_id` - optional string specifying ID of the budget policy for this DLT pipeline.
* `allow_duplicate_names` - Optional boolean flag. If false, deployment will fail if name conflicts with that of another pipeline. default is `false`.
* `deployment` - Deployment type of this pipeline. Supports following attributes:
  * `kind` - The deployment method that manages the pipeline.
  * `metadata_file_path` - The path to the file containing metadata about the deployment.
* `filters` - Filters on which Pipeline packages to include in the deployed graph.  This block consists of following attributes:
  * `include` - Paths to include.
  * `exclude` - Paths to exclude.
* `gateway_definition` - The definition of a gateway pipeline to support CDC. Consists of following attributes:
  * `connection_id` - Immutable. The Unity Catalog connection this gateway pipeline uses to communicate with the source.
  * `gateway_storage_catalog` - Required, Immutable. The name of the catalog for the gateway pipeline's storage location.
  * `gateway_storage_name` - Required. The Unity Catalog-compatible naming for the gateway storage location. This is the destination to use for the data that is extracted by the gateway. Delta Live Tables system will automatically create the storage location under the catalog and schema.
  * `gateway_storage_schema` - Required, Immutable. The name of the schema for the gateway pipelines's storage location.
* `event_log` - an optional block specifying a table where DLT Event Log will be stored.  Consists of the following fields:
  * `name` - (Required) The table name the event log is published to in UC.
  * `catalog` - (Optional, default to `catalog` defined on pipeline level) The UC catalog the event log is published under.
  * `schema` - (Optional, default to `schema` defined on pipeline level) The UC schema the event log is published under.

### notification block

DLT allows to specify one or more notification blocks to get notifications about pipeline's execution.  This block consists of following attributes:

* `email_recipients` (Required) non-empty list of emails to notify.
* `alerts` (Required) non-empty list of alert types. Right now following alert types are supported, consult documentation for actual list
  * `on-update-success` - a pipeline update completes successfully.
  * `on-update-failure` - a pipeline update fails with a retryable error.
  * `on-update-fatal-failure` - a pipeline update fails with a non-retryable (fatal) error.
  * `on-flow-failure` - a single data flow fails.

### ingestion_definition block

The configuration for a managed ingestion pipeline. These settings cannot be used with the `library`, `target` or `catalog` settings. This block consists of following attributes:

* `connection_name` - Immutable. The Unity Catalog connection this ingestion pipeline uses to communicate with the source. Specify either ingestion_gateway_id or connection_name.
* `ingestion_gateway_id` - Immutable. Identifier for the ingestion gateway used by this ingestion pipeline to communicate with the source. Specify either ingestion_gateway_id or connection_name.
* `objects` - Required. Settings specifying tables to replicate and the destination for the replicated tables.
* `table_configuration` - Configuration settings to control the ingestion of tables. These settings are applied to all tables in the pipeline.



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier of the DLT pipeline.
* `url` - URL of the DLT pipeline on the given workspace.

## Import

The resource job can be imported using the id of the pipeline

```hcl
import {
  to = databricks_pipeline.this
  id = "<pipeline-id>"
}
```

Alternatively, when using `terraform` version 1.5 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_pipeline.this <pipeline-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_pipelines](../data-sources/pipelines.md) to retrieve [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html) pipeline data.
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
