---
subcategory: "Other"
---
# databricks_queryable_job Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

Retrieves the settings of [databricks_job](../resources/job.md) by name or by id. Complements the feature of the [databricks_jobs](./jobs.md) data source

## Example Usage

Getting the existing cluster id of specific [databricks_job](../resources/job.md) by name or by id:

```hcl
data "databricks_job" "this" {
  name = "My job"
}

output "cluster_id" {
  value     = data.databricks_job.job_settings.existing_cluster_id
  sensitive = false
}
```

## Attribute Reference

This data source exports the following attributes:

* `id` - the id of [databricks_job](../resources/job.md) if the resource was matched by name
* `name` - the job name of [databricks_job](../resources/job.md) if the resource was matched by id
* `job_settings` - the job settings of [databricks_job](../resources/job.md)

## Related Resources

The following resources are used in the same context:

* TODO: write me