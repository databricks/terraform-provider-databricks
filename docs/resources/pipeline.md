---
subcategory: "Compute"
---
# databricks_pipeline Resource

Use `databricks_pipeline` to deploy [Lakeflow Declarative Pipelines](https://docs.databricks.com/aws/en/dlt).

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_notebook" "ldp_demo" {
  #...
}

resource "databricks_repo" "ldp_demo" {
  #...
}

resource "databricks_pipeline" "this" {
  name    = "Pipeline Name"
  catalog = "main"
  schema  = "ldp_demo"

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
      path = databricks_notebook.ldp_demo.id
    }
  }

  library {
    file {
      path = "${databricks_repo.ldp_demo.path}/pipeline.sql"
    }
  }

  library {
    glob {
      include = "${databricks_repo.ldp_demo.path}/subfolder/**"
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
* `catalog` - The name of catalog in Unity Catalog. *Change of this parameter forces recreation of the pipeline.* (Conflicts with `storage`).
* `schema` - (Optional, String, Conflicts with `target`) The default schema (database) where tables are read from or published to. The presence of this attribute implies that the pipeline is in direct publishing mode.
* `storage` - A location on cloud storage where output data and metadata required for pipeline execution are stored. By default, tables are stored in a subdirectory of this location. *Change of this parameter forces recreation of the pipeline.* (Conflicts with `catalog`).
* `target` - (Optional, String, Conflicts with `schema`) The name of a database (in either the Hive metastore or in a UC catalog) for persisting pipeline output data. Configuring the target setting allows you to view and query the pipeline output data from the Databricks UI.
* `configuration` - An optional list of values to apply to the entire pipeline. Elements must be formatted as key:value pairs.
* `library` blocks - Specifies pipeline code.
* `root_path` - An optional string specifying the root path for this pipeline. This is used as the root directory when editing the pipeline in the Databricks user interface and it is added to `sys.path` when executing Python sources during pipeline execution.
* `cluster` blocks - [Clusters](cluster.md) to run the pipeline. If none is specified, pipelines will automatically select a default cluster configuration for the pipeline. *Please note that Lakeflow Declarative Pipeline clusters are supporting only subset of attributes as described in [documentation](https://docs.databricks.com/api/workspace/pipelines/create#clusters).*  Also, note that `autoscale` block is extended with the `mode` parameter that controls the autoscaling algorithm (possible values are `ENHANCED` for new, enhanced autoscaling algorithm, or `LEGACY` for old algorithm).
* `continuous` - A flag indicating whether to run the pipeline continuously. The default value is `false`.
* `development` - A flag indicating whether to run the pipeline in development mode. The default value is `false`.
* `photon` - A flag indicating whether to use Photon engine. The default value is `false`.
* `serverless` - An optional flag indicating if serverless compute should be used for this Lakeflow Declarative Pipeline.  Requires `catalog` to be set, as it could be used only with Unity Catalog.
* `edition` - optional name of the [product edition](https://docs.databricks.com/aws/en/dlt/configure-pipeline#choose-a-product-edition). Supported values are: `CORE`, `PRO`, `ADVANCED` (default).  Not required when `serverless` is set to `true`.
* `channel` - optional name of the release channel for Spark version used by Lakeflow Declarative Pipeline.  Supported values are: `CURRENT` (default) and `PREVIEW`.
* `budget_policy_id` - optional string specifying ID of the budget policy for this Lakeflow Declarative Pipeline.
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
  * `gateway_storage_name` - Required. The Unity Catalog-compatible naming for the gateway storage location. This is the destination to use for the data that is extracted by the gateway. Lakeflow Declarative Pipelines system will automatically create the storage location under the catalog and schema.
  * `gateway_storage_schema` - Required, Immutable. The name of the schema for the gateway pipelines's storage location.
* `event_log` - an optional block specifying a table where LDP Event Log will be stored.  Consists of the following fields:
  * `name` - (Required) The table name the event log is published to in UC.
  * `catalog` - (Optional, default to `catalog` defined on pipeline level) The UC catalog the event log is published under.
  * `schema` - (Optional, default to `schema` defined on pipeline level) The UC schema the event log is published under.
* `tags` - (Optional, map of strings) A map of tags associated with the pipeline. These are forwarded to the cluster as cluster tags, and are therefore subject to the same limitations. A maximum of 25 tags can be added to the pipeline.

### library block

Contains one of the blocks:

* `notebook` - specifies path to a Databricks Notebook to include as source. Actual path is specified as `path` attribute inside the block.
* `file` - specifies path to a file in Databricks Workspace to include as source. Actual path is specified as `path` attribute inside the block.
* `glob` - The unified field to include source code. Each entry should have the `include` attribute that can specify a notebook path, a file path, or a folder path that ends `/**` (to include everything from that folder). This field cannot be used together with `notebook` or `file`.

### environment block

Environment specification for the current pipeline used to install dependencies when running on serverless compute.  Consists of the following attributes:

* `dependencies` - (Required) a list of pip dependencies, as supported by the version of pip in this environment. Each dependency is a [pip requirement file line](https://pip.pypa.io/en/stable/reference/requirements-file-format/).  See [API docs](https://docs.databricks.com/api/azure/workspace/pipelines/create#environment-dependencies) for more information.

Example:

```hcl
resource "databricks_pipeline" "this" {
  name       = "Serverless demo"
  serverless = true
  catalog    = "main"
  schema     = "ldp_demo"

  # ...

  environment {
    dependencies = [
      "foo==0.0.1",
      "-r /Workspace/Users/user.name/my-pipeline/requirements.txt",
      "/Volumes/main/default/libs/my_lib.whl"
    ]
  }
}
```

### notification block

Lakeflow Declarative Pipeline allows to specify one or more notification blocks to get notifications about pipeline's execution.  This block consists of following attributes:

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

* `id` - Canonical unique identifier of the Lakeflow Declarative Pipeline.
* `url` - URL of the Lakeflow Declarative Pipeline on the given workspace.

## Import

The resource job can be imported using the id of the pipeline

```hcl
import {
  to = databricks_pipeline.this
  id = "<pipeline-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_pipeline.this <pipeline-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_pipelines](../data-sources/pipelines.md) to retrieve [Lakeflow Declarative Pipelines](https://docs.databricks.com/aws/en/dlt) data.
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
