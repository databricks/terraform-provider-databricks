---
subcategory: "Compute"
---
# databricks_cluster Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves information about a [databricks_cluster](../resources/cluster.md) using its id. This could be retrieved programmatically using [databricks_clusters](../data-sources/clusters.md) data source.

## Example Usage

Retrieve attributes of each SQL warehouses in a workspace

```hcl
data "databricks_clusters" "all" {
}

data "databricks_cluster" "all" {
  for_each   = data.databricks_clusters.all.ids
  cluster_id = each.value
}

```

### Multiple clusters with the same name

When fetching a cluster whose name is not unique (including terminated but not permanently deleted clusters), you must use the `cluster_id` argument to uniquely identify the cluster. Combine this data source with `databricks_clusters` to get the `cluster_id` of the cluster you want to fetch.

```hcl
data "databricks_clusters" "my_cluster" {
  cluster_name_contains = "my-cluster"
  filter_by {
    cluster_states = ["RUNNING"]
    # Filter additionally on cluster sources if needed:
    # cluster_sources = ["API"] # if created by Terraform or another API-based tool
    # cluster_sources = ["UI"] # if created in the Databricks web interface
  }
}

data "databricks_cluster" "my_cluster" {
  cluster_id = data.databricks_clusters.my_cluster.ids[0]
}
```

## Argument Reference

* `cluster_id` - (Required if `cluster_name` isn't specified) The id of the cluster.
* `cluster_name` - (Required if `cluster_id` isn't specified) The exact name of the cluster to search. Can only be specified if there is exactly one cluster with the provided name.

## Attribute Reference

This data source exports the following attributes:

* `id` - cluster ID
* `cluster_info` block, consisting of following fields:
  * `cluster_name` - Cluster name, which doesn’t have to be unique.
  * `spark_version` - [Runtime version](https://docs.databricks.com/runtime/index.html) of the cluster.
  * `runtime_engine` - The type of runtime of the cluster
  * `driver_node_type_id` - The node type of the Spark driver.
  * `node_type_id` - Any supported [databricks_node_type](../data-sources/node_type.md) id.
  * `instance_pool_id` The [pool of idle instances](instance_pool.md) the cluster is attached to.
  * `driver_instance_pool_id` - similar to `instance_pool_id`, but for driver node.
  * `policy_id` - Identifier of [Cluster Policy](cluster_policy.md) to validate cluster and preset certain defaults.
  * `autotermination_minutes` - Automatically terminate the cluster after being inactive for this time in minutes. If specified, the threshold must be between 10 and 10000 minutes. You can also set this value to 0 to explicitly disable automatic termination.
  * `enable_elastic_disk` - Use autoscaling local storage.
  * `enable_local_disk_encryption` - Enable local disk encryption.
  * `data_security_mode` - Security features of the cluster. Unity Catalog requires `SINGLE_USER` or `USER_ISOLATION` mode. `LEGACY_PASSTHROUGH` for passthrough cluster and `LEGACY_TABLE_ACL` for Table ACL cluster. Default to `NONE`, i.e. no security feature enabled.
  * `single_user_name` - The optional user name of the user to assign to an interactive cluster. This field is required when using standard AAD Passthrough for Azure Data Lake Storage (ADLS) with a single-user cluster (i.e., not high-concurrency clusters).
  * `idempotency_token` - An optional token to guarantee the idempotency of cluster creation requests.
  * `ssh_public_keys` - SSH public key contents that will be added to each Spark node in this cluster.
  * `spark_env_vars` - Map with environment variable key-value pairs to fine-tune Spark clusters. Key-value pairs of the form (X,Y) are exported (i.e., X='Y') while launching the driver and workers.
  * `custom_tags` - Additional tags for cluster resources.
  * `spark_conf` - Map with key-value pairs to fine-tune Spark clusters.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_cluster](../resources/cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_cluster_policy](../resources/cluster_policy.md) to create a [databricks_cluster](../resources/cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_instance_pool](../resources/instance_pool.md) to manage [instance pools](https://docs.databricks.com/clusters/instance-pools/index.html) to reduce [cluster](../resources/cluster.md) start and auto-scaling times by maintaining a set of idle, ready-to-use instances.
* [databricks_job](../resources/job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](../resources/cluster.md).
* [databricks_library](../resources/library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](../resources/cluster.md).
* [databricks_pipeline](../resources/pipeline.md) to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html).
