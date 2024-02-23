---
subcategory: "Compute"
---
# databricks_clusters Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves a list of [databricks_cluster](../resources/cluster.md#cluster_id) ids, that were created by Terraform or manually, with or without [databricks_cluster_policy](../resources/cluster_policy.md).

## Example Usage

Retrieve all clusters on this workspace on AWS or GCP:

```hcl
data "databricks_clusters" "all" {
  depends_on = [databricks_mws_workspaces.this]
}
```

Retrieve all clusters with "Shared" in their cluster name on this Azure Databricks workspace:

```hcl
data "databricks_clusters" "all_shared" {
  depends_on            = [azurerm_databricks_workspace.this]
  cluster_name_contains = "shared"
}
```

## Argument Reference

* `cluster_name_contains` - (Optional) Only return [databricks_cluster](../resources/cluster.md#cluster_id) ids that match the given name string.

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
* [databricks_pipeline](../resources/pipeline.md) to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html).
