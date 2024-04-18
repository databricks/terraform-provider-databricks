---
subcategory: "Compute"
---
# databricks_pipelines Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

Retrieves a list of all [databricks_pipeline](../resources/pipeline.md) ([Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html)) ids deployed in a workspace, or those matching the provided search term. Maximum 100 results.

## Example Usage

Get all Delta Live Tables pipelines:

```hcl
data "databricks_pipelines" "all" {}

output "all_pipelines" {
  value = data.databricks_pipelines.all.ids
}
```

Filter Delta Live Tables pipelines by name (exact match):

```hcl
data "databricks_pipelines" "this" {
  pipeline_name = "my_pipeline"
}

output "my_pipeline" {
  value = data.databricks_pipelines.this.ids
}
```

Filter Delta Live Tables pipelines by name (wildcard search):

```hcl
data "databricks_pipelines" "this" {
  pipeline_name = "%pipeline%"
}

output "wildcard_pipelines" {
  value = data.databricks_pipelines.this.ids
}
```

## Argument Reference

This data source exports the following attributes:

* `pipeline_name` - (Optional) Filter Delta Live Tables pipelines by name for a given search term. `%` is the supported wildcard operator.
  
## Attribute Reference

This data source exports the following attributes:

* `ids` - List of ids for [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html) pipelines matching the provided search criteria.

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_pipeline](../resources/pipeline.md) to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html).
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
