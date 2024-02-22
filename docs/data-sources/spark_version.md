---
subcategory: "Compute"
---
# databricks_spark_version Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Gets [Databricks Runtime (DBR)](https://docs.databricks.com/runtime/dbr.html) version that could be used for `spark_version` parameter in [databricks_cluster](../resources/cluster.md) and other resources that fits search criteria, like specific Spark or Scala version, ML or Genomics runtime, etc., similar to executing `databricks clusters spark-versions`, and filters it to return the latest version that matches criteria. Often used along [databricks_node_type](node_type.md) data source.

-> **Note** This is experimental functionality, which aims to simplify things. In case of wrong parameters given (e.g. together `ml = true` and `genomics = true`, or something like), data source will throw an error.  Similarly, if search returns multiple results, and `latest = false`, data source will throw an error.

## Example Usage

```hcl
data "databricks_node_type" "with_gpu" {
  local_disk  = true
  min_cores   = 16
  gb_per_core = 1
  min_gpus    = 1
}

data "databricks_spark_version" "gpu_ml" {
  gpu = true
  ml  = true
}

resource "databricks_cluster" "research" {
  cluster_name            = "Research Cluster"
  spark_version           = data.databricks_spark_version.gpu_ml.id
  node_type_id            = data.databricks_node_type.with_gpu.id
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 50
  }
}
```

## Argument Reference

Data source allows you to pick groups by the following attributes:

* `latest` - (boolean, optional) if we should return only the latest version if there is more than one result.  Default to `true`. If set to `false` and multiple versions are matching, throws an error.
* `long_term_support` - (boolean, optional) if we should limit the search only to LTS (long term support) & ESR (extended support) versions. Default to `false`.
* `ml` - (boolean, optional) if we should limit the search only to ML runtimes. Default to `false`.
* `genomics` - (boolean, optional)  if we should limit the search only to Genomics (HLS) runtimes. Default to `false`.
* `gpu` - (boolean, optional)  if we should limit the search only to runtimes that support GPUs. Default to `false`.
* `beta` - (boolean, optional) if we should limit the search only to runtimes that are in Beta stage. Default to `false`.
* `scala` - (string, optional) if we should limit the search only to runtimes that are based on specific Scala version. Default to `2.12`.
* `spark_version` - (string, optional) if we should limit the search only to runtimes that are based on specific Spark version. Default to empty string.  It could be specified as `3`, or `3.0`, or full version, like, `3.0.1`.
* `photon` - (boolean, optional)  if we should limit the search only to Photon runtimes. Default to `false`. *Deprecated with DBR 14.0 release. Specify `runtime_engine=\"PHOTON\"` in the cluster configuration instead!*
* `graviton` - (boolean, optional)  if we should limit the search only to runtimes supporting AWS Graviton CPUs. Default to `false`. _Deprecated with DBR 14.0 release. DBR version compiled for Graviton will be automatically installed when nodes with Graviton CPUs are specified in the cluster configuration._

## Attribute Reference

Data source exposes the following attributes:

* `id` - Databricks Runtime version, that can be used as `spark_version` field in [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md).

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_cluster](../resources/cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_cluster_policy](../resources/cluster_policy.md) to create a [databricks_cluster](../resources/cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_instance_pool](../resources/instance_pool.md) to manage [instance pools](https://docs.databricks.com/clusters/instance-pools/index.html) to reduce [cluster](../resources/cluster.md) start and auto-scaling times by maintaining a set of idle, ready-to-use instances.
* [databricks_job](../resources/job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](../resources/cluster.md).
