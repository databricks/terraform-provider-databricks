---
subcategory: "Compute"
---
# databricks_job Data Source

Retrieves the settings of [databricks_job](../resources/job.md) by name or by id. Complements the feature of the [databricks_jobs](jobs.md) data source.

-> This data source can only be used with a workspace-level provider!

## Example Usage

Getting the existing cluster id of specific [databricks_job](../resources/job.md) by name or by id:

```hcl
data "databricks_job" "this" {
  job_name = "My job"
}

output "job_num_workers" {
  value     = data.databricks_job.this.job_settings[0].settings[0].new_cluster[0].num_workers
  sensitive = false
}
```

## Attribute Reference

This data source exports the following attributes:

* `id` - the id of [databricks_job](../resources/job.md) if the resource was matched by name.
* `name` - the job name of [databricks_job](../resources/job.md) if the resource was matched by id.
* `job_settings` - the same fields as in [databricks_job](../resources/job.md).

## Related Resources

The following resources are used in the same context:

* [databricks_jobs](jobs.md) data to get all jobs and their names from a workspace.
* [databricks_job](../resources/job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](../resources/cluster.md).
