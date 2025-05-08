---
subcategory: "Compute"
---
# databricks_jobs Data Source

Retrieves a list of [databricks_job](../resources/job.md) ids, that were created by Terraform or manually, so that special handling could be applied.

-> This data source can only be used with a workspace-level provider!

~> Data resource will error in case of jobs with duplicate names.

## Example Usage

Granting view [databricks_permissions](../resources/permissions.md) to all [databricks_job](../resources/job.md) within the workspace:

```hcl
data "databricks_jobs" "this" {}

data "databricks_jobs" "tests" {
  job_name_contains = "test"
}

resource "databricks_permissions" "everyone_can_view_all_jobs" {
  for_each = data.databricks_jobs.this.ids
  job_id   = each.value

  access_control {
    group_name       = "users"
    permission_level = "CAN_VIEW"
  }
}
```

Getting ID of specific [databricks_job](../resources/job.md) by name:

```hcl
data "databricks_jobs" "this" {}

output "x" {
  value     = "ID of `x` job is ${data.databricks_jobs.this.ids["x"]}"
  sensitive = false
}
```

## Argument Reference

* `job_name_contains` - (Optional) Only return [databricks_job](../resources/job.md#) ids that match the given name string (case-insensitive).

## Attribute Reference

This data source exports the following attributes:

* `ids` - map of [databricks_job](../resources/job.md) names to ids

## Related Resources

The following resources are used in the same context:

* [databricks_job](../resources/job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](../resources/cluster.md).
