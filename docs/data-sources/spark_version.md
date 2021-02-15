---
subcategory: "Compute"
---
# databricks_spark_version Data Source

Gets Databricks Runtime (DBR) version that could be used for `spark_version` parameter in [databricks_cluster](../resources/cluster.md) and other resources that fits search criteria, like specific Spark or Scala version, ML or Genomics runtime, etc., similar to executing `databricks clusters spark-versions`, and filters it to return the latest version that matches criteria. Often used along [databricks_node_type](node_type.md) data source.

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
  ml = true
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

* `latest` - (boolean, optional) if we should return only the latest version if there is more than one result.  Default to `true`. If set to `false` and multiple versions are matching, throws an error
* `long_term_support` - (boolean, optional) if we should limit the search only to LTS (long term support) versions. Default to `false`
* `ml` - (boolean, optional) if we should limit the search only to ML runtimes. Default to `false`
* `genomics` - (boolean, optional)  if we should limit the search only to Genomics (HLS) runtimes. Default to `false`
* `gpu` - (boolean, optional)  if we should limit the search only to runtimes that support GPUs. Default to `false`
* `beta` - (boolean, optional) if we should limit the search only to runtimes that are in Beta stage. Default to `false`
* `scala` - (boolean, optional) if we should limit the search only to runtimes that are based on specific Scala version. Default to `2.12`
* `spark_version` - (string, optional) if we should limit the search only to runtimes that are based on specific Spark version. Default to empty string.  It could be specified as `3`, or `3.0`, or full version, like, `3.0.1`

## Attribute Reference

Data source exposes the following attributes:

* `id` - Databricks Runtime version, that can be used as `spark_version` field in [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md).
