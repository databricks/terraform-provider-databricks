---
subcategory: "Compute"
---
# databricks_clusters Data Source

Retrieves a list of [databricks_cluster](../resources/cluster.md#cluster_id) ids, that were created by Terraform or manually, with or without [databricks_cluster_policy](../resources/cluster_policy.md).

-> This data source can only be used with a workspace-level provider!

## Example Usage

Retrieve cluster IDs for all clusters:

```hcl
data "databricks_clusters" "all" {
}
```

Retrieve cluster IDs for all clusters having "Shared" in the cluster name:

```hcl
data "databricks_clusters" "all_shared" {
  cluster_name_contains = "shared"
}
```

### Filtering clusters

Listing clusters can be slow for workspaces containing many clusters. Use filters to limit the number of clusters returned for better performance. You can filter clusters by state, source, policy, or pinned status:

```hcl
data "databricks_clusters" "all_running_clusters" {
  filter_by {
    cluster_states = ["RUNNING"]
  }
}

data "databricks_clusters" "all_clusters_with_policy" {
  filter_by {
    policy_id = "1234-5678-9012"
  }
}

data "databricks_clusters" "all_api_clusters" {
  filter_by {
    cluster_sources = ["API"]
  }
}

data "databricks_clusters" "all_pinned_clusters" {
  filter_by {
    is_pinned = true
  }
}
```

## Argument Reference

* `cluster_name_contains` - (Optional) Only return [databricks_cluster](../resources/cluster.md#cluster_id) ids that match the given name string.
* `filter_by` - (Optional) Filters to apply to the listed clusters. See [filter_by Configuration Block](#filter_by-configuration-block) below for details.

### filter_by Configuration Block

The `filter_by` block controls the filtering of the listed clusters. It supports the following arguments:

* `cluster_sources` - (Optional) List of cluster sources to filter by. Possible values are `API`, `JOB`, `MODELS`, `PIPELINE`, `PIPELINE_MAINTENANCE`, `SQL`, and `UI`.
* `cluster_states` - (Optional) List of cluster states to filter by. Possible values are `RUNNING`, `PENDING`, `RESIZING`, `RESTARTING`, `TERMINATING`, `TERMINATED`, `ERROR`, and `UNKNOWN`.
* `is_pinned` - (Optional) Whether to filter by pinned clusters.
* `policy_id` - (Optional) Filter by [databricks_cluster_policy](../resources/cluster_policy.md) id.

## Attribute Reference

This data source exports the following attributes:

* `ids` - list of [databricks_cluster](../resources/cluster.md#cluster_id) ids

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_cluster](../resources/cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_cluster_policy](../resources/cluster_policy.md) to create a [databricks_cluster](../resources/cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_instance_pool](../resources/instance_pool.md) to manage [instance pools](https://docs.databricks.com/clusters/instance-pools/index.html) to reduce [cluster](../resources/cluster.md) start and auto-scaling times by maintaining a set of idle, ready-to-use instances.
* [databricks_job](../resources/job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](../resources/cluster.md).
* [databricks_library](../resources/library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](../resources/cluster.md).
* [databricks_pipeline](../resources/pipeline.md) to deploy [Lakeflow Declarative Pipelines](https://docs.databricks.com/aws/en/dlt).
