---
subcategory: "Compute"
---
# databricks_node_type Data Source

Gets the smallest node type for [databricks_cluster](../resources/cluster.md) that fits search criteria, like amount of RAM or number of cores. [AWS](https://databricks.com/product/aws-pricing/instance-types) or [Azure](https://azure.microsoft.com/en-us/pricing/details/databricks/). Internally data source fetches [node types](https://docs.databricks.com/dev-tools/api/latest/clusters.html#list-node-types) available per cloud, similar to executing `databricks clusters list-node-types`, and filters it to return the smallest possible node with criteria.

-> This data source can only be used with a workspace-level provider!

-> This is experimental functionality, which aims to simplify things. In case of wrong parameters given (e.g. `min_gpus = 876`) or no nodes matching, data source will return cloud-default node type, even though it doesn't match search criteria specified by data source arguments: [i3.xlarge](https://aws.amazon.com/ec2/instance-types/i3/) for AWS or [Standard_D3_v2](https://docs.microsoft.com/en-us/azure/cloud-services/cloud-services-sizes-specs#dv2-series) for Azure.

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

Data source allows you to pick groups by the following attributes

* `min_memory_gb` - (Optional) Minimum amount of memory per node in gigabytes. Defaults to _0_.
* `gb_per_core` - (Optional) Number of gigabytes per core available on instance. Conflicts with `min_memory_gb`. Defaults to _0_.
* `min_cores` - (Optional) Minimum number of CPU cores available on instance. Defaults to _0_.
* `min_gpus` - (Optional) Minimum number of GPU's attached to instance. Defaults to _0_.
* `local_disk` - (Optional) Pick only nodes with local storage. Defaults to _false_.
* `local_disk_min_size` - (Optional) Pick only nodes that have size local storage greater or equal to given value. Defaults to _0_.
* `category` - (Optional, case insensitive string) Node category, which can be one of (depending on the cloud environment, could be checked with `databricks clusters list-node-types -o json|jq '.node_types[]|.category'|sort |uniq`):
  * `General Purpose` (all clouds)
  * `General Purpose (HDD)` (Azure)
  * `Compute Optimized` (all clouds)
  * `Memory Optimized` (all clouds)
  * `Memory Optimized (Remote HDD)` (Azure)
  * `Storage Optimized` (AWS, Azure)
  * `GPU Accelerated` (AWS, Azure)
* `photon_worker_capable` - (Optional) Pick only nodes that can run Photon workers. Defaults to _false_.
* `photon_driver_capable` - (Optional) Pick only nodes that can run Photon driver. Defaults to _false_.
* `arm` - (boolean, optional)  if we should limit the search only to nodes with AWS Graviton or Azure Cobalt CPUs. Default to _false_.
* `fleet` - (boolean, optional)  if we should limit the search only to [AWS fleet instance types](https://docs.databricks.com/compute/aws-fleet-instances.html). Default to _false_.
* `is_io_cache_enabled` - (Optional) . Pick only nodes that have IO Cache. Defaults to _false_.
* `support_port_forwarding` - (Optional) Pick only nodes that support port forwarding. Defaults to _false_.
* `graviton` - (Deprecated. boolean, optional) if we should limit the search only to nodes with AWS Graviton or Azure Cobalt CPUs. Default to _false_. *Use `arm` instead!*

## Attribute Reference

Data source exposes the following attributes:

* `id` - node type, that can be used for [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md).

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_cluster](../resources/cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_cluster_policy](../resources/cluster_policy.md) to create a [databricks_cluster](../resources/cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_instance_pool](../resources/instance_pool.md) to manage [instance pools](https://docs.databricks.com/clusters/instance-pools/index.html) to reduce [cluster](../resources/cluster.md) start and auto-scaling times by maintaining a set of idle, ready-to-use instances.
* [databricks_job](../resources/job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](../resources/cluster.md).
