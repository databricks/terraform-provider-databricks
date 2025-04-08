---
subcategory: "Compute"
---
# databricks_cluster resource

This resource allows you to manage [Databricks Clusters](https://docs.databricks.com/clusters/index.html).

-> In case of [`Cannot access cluster ####-######-####### that was terminated or unpinned more than 30 days ago`](https://github.com/databricks/terraform-provider-databricks/issues/1197#issuecomment-1069386670) errors, please upgrade to v0.5.5 or later. If for some reason you cannot upgrade the version of provider, then the other viable option to unblock the apply pipeline is [`terraform state rm path.to.databricks_cluster.resource`](https://www.terraform.io/cli/commands/state/rm) command.

```hcl
data "databricks_node_type" "smallest" {
  local_disk = true
}

data "databricks_spark_version" "latest_lts" {
  long_term_support = true
}

resource "databricks_cluster" "shared_autoscaling" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = data.databricks_spark_version.latest_lts.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 50
  }
}
```

## Argument Reference

* `num_workers` - (Optional) Number of worker nodes that this cluster should have. A cluster has one Spark driver and `num_workers` executors for a total of `num_workers` + 1 Spark nodes.
* `cluster_name` - (Optional) Cluster name, which doesn’t have to be unique. If not specified at creation, the cluster name will be an empty string.
* `spark_version` - (Required) [Runtime version](https://docs.databricks.com/runtime/index.html) of the cluster. Any supported [databricks_spark_version](../data-sources/spark_version.md) id.  We advise using [Cluster Policies](cluster_policy.md) to restrict the list of versions for simplicity while maintaining enough control.
* `runtime_engine` - (Optional) The type of runtime engine to use. If not specified, the runtime engine type is inferred based on the spark_version value. Allowed values include: `PHOTON`, `STANDARD`.
* `use_ml_runtime` - (Optional, Boolean, only with `kind`) Whenever ML runtime should be selected or not.  Actual runtime is determined by `spark_version` (DBR release), this field `use_ml_runtime`, and whether `node_type_id` is GPU node or not.
* `is_single_node` - (Optional, Boolean, only with `kind`) When set to true, Databricks will automatically set single node related `custom_tags`, `spark_conf`, and `num_workers`.
* `driver_node_type_id` - (Optional) The node type of the Spark driver. This field is optional; if unset, API will set the driver node type to the same value as `node_type_id` defined above.
* `node_type_id` - (Required - optional if `instance_pool_id` is given) Any supported [databricks_node_type](../data-sources/node_type.md) id. If `instance_pool_id` is specified, this field is not needed.
* `instance_pool_id` (Optional - required if `node_type_id` is not given) - To reduce cluster start time, you can attach a cluster to a [predefined pool of idle instances](instance_pool.md). When attached to a pool, a cluster allocates its driver and worker nodes from the pool. If the pool does not have sufficient idle resources to accommodate the cluster’s request, it expands by allocating new instances from the instance provider. When an attached cluster changes its state to `TERMINATED`, the instances it used are returned to the pool and reused by a different cluster.
* `driver_instance_pool_id` (Optional) - similar to `instance_pool_id`, but for driver node. If omitted, and `instance_pool_id` is specified, then the driver will be allocated from that pool.
* `policy_id` - (Optional) Identifier of [Cluster Policy](cluster_policy.md) to validate cluster and preset certain defaults. *The primary use for cluster policies is to allow users to create policy-scoped clusters via UI rather than sharing configuration for API-created clusters.* For example, when you specify `policy_id` of [external metastore](https://docs.databricks.com/administration-guide/clusters/policies.html#external-metastore-policy) policy, you still have to fill in relevant keys for `spark_conf`.  If relevant fields aren't filled in, then it will cause the configuration drift detected on each plan/apply, and Terraform will try to apply the detected changes.
* `apply_policy_default_values` - (Optional) Whether to use policy default values for missing cluster attributes.
* `autotermination_minutes` - (Optional) Automatically terminate the cluster after being inactive for this time in minutes. If specified, the threshold must be between 10 and 10000 minutes. You can also set this value to 0 to explicitly disable automatic termination. Defaults to `60`.  *We highly recommend having this setting present for Interactive/BI clusters.*
* `enable_elastic_disk` - (Optional) If you don’t want to allocate a fixed number of EBS volumes at cluster creation time, use autoscaling local storage. With autoscaling local storage, Databricks monitors the amount of free disk space available on your cluster’s Spark workers. If a worker begins to run too low on disk, Databricks automatically attaches a new EBS volume to the worker before it runs out of disk space. EBS volumes are attached up to a limit of 5 TB of total disk space per instance (including the instance’s local storage). To scale down EBS usage, make sure you have `autotermination_minutes` and `autoscale` attributes set. More documentation available at [cluster configuration page](https://docs.databricks.com/clusters/configure.html#autoscaling-local-storage-1).
* `enable_local_disk_encryption` - (Optional) Some instance types you use to run clusters may have locally attached disks. Databricks may store shuffle data or temporary data on these locally attached disks. To ensure that all data at rest is encrypted for all storage types, including shuffle data stored temporarily on your cluster’s local disks, you can enable local disk encryption. When local disk encryption is enabled, Databricks generates an encryption key locally unique to each cluster node and uses it to encrypt all data stored on local disks. The scope of the key is local to each cluster node and is destroyed along with the cluster node itself. During its lifetime, the key resides in memory for encryption and decryption and is stored encrypted on the disk. *Your workloads may run more slowly because of the performance impact of reading and writing encrypted data to and from local volumes. This feature is not available for all Azure Databricks subscriptions. Contact your Microsoft or Databricks account representative to request access.*
* `kind` - (Optional, enum) The kind of compute described by this compute specification.  Possible values (see [API docs](https://docs.databricks.com/api/workspace/clusters/create#kind) for full list): `CLASSIC_PREVIEW` (if corresponding public preview is enabled).
* `data_security_mode` - (Optional) Select the security features of the cluster (see [API docs](https://docs.databricks.com/api/workspace/clusters/create#data_security_mode) for full list of values). [Unity Catalog requires](https://docs.databricks.com/data-governance/unity-catalog/compute.html#create-clusters--sql-warehouses-with-unity-catalog-access) `SINGLE_USER` or `USER_ISOLATION` mode. `LEGACY_PASSTHROUGH` for passthrough cluster and `LEGACY_TABLE_ACL` for Table ACL cluster. If omitted, default security features are enabled. To disable security features use `NONE` or legacy mode `NO_ISOLATION`.  If `kind` is specified, then the following options are available:
  * `DATA_SECURITY_MODE_AUTO`: Databricks will choose the most appropriate access mode depending on your compute configuration.
  * `DATA_SECURITY_MODE_STANDARD`: Alias for `USER_ISOLATION`.
  * `DATA_SECURITY_MODE_DEDICATED`: Alias for `SINGLE_USER`.
* `single_user_name` - (Optional) The optional user name of the user (or group name if `kind` if specified) to assign to an interactive cluster. This field is required when using `data_security_mode` set to `SINGLE_USER` or AAD Passthrough for Azure Data Lake Storage (ADLS) with a single-user cluster (i.e., not high-concurrency clusters).
* `idempotency_token` - (Optional) An optional token to guarantee the idempotency of cluster creation requests. If an active cluster with the provided token already exists, the request will not create a new cluster, but it will return the existing running cluster's ID instead. If you specify the idempotency token, upon failure, you can retry until the request succeeds. Databricks platform guarantees to launch exactly one cluster with that idempotency token. This token should have at most 64 characters.
* `ssh_public_keys` - (Optional) SSH public key contents that will be added to each Spark node in this cluster. The corresponding private keys can be used to login with the user name ubuntu on port 2200. You can specify up to 10 keys.
* `spark_env_vars` - (Optional) Map with environment variable key-value pairs to fine-tune Spark clusters. Key-value pairs of the form (X,Y) are exported (i.e., X='Y') while launching the driver and workers.
* `custom_tags` - (Optional) Additional tags for cluster resources. Databricks will tag all cluster resources (e.g., AWS EC2 instances and EBS volumes) with these tags in addition to `default_tags`. If a custom cluster tag has the same name as a default cluster tag, the custom tag is prefixed with an `x_` when it is propagated.
* `spark_conf` - (Optional) Map with key-value pairs to fine-tune Spark clusters, where you can provide custom [Spark configuration properties](https://spark.apache.org/docs/latest/configuration.html) in a cluster configuration.
* `is_pinned` - (Optional) boolean value specifying if the cluster is pinned (not pinned by default). You must be a Databricks administrator to use this.  The pinned clusters' maximum number is [limited to 100](https://docs.databricks.com/clusters/clusters-manage.html#pin-a-cluster), so `apply` may fail if you have more than that (this number may change over time, so check Databricks documentation for actual number).
* `no_wait` - (Optional) If true, the provider will not wait for the cluster to reach `RUNNING` state when creating the cluster, allowing cluster creation and library installation to continue asynchronously. Defaults to false (the provider will wait for cluster creation and library installation to succeed).

The following example demonstrates how to create an autoscaling cluster with [Delta Cache](https://docs.databricks.com/delta/optimizations/delta-cache.html) enabled:

```hcl
data "databricks_node_type" "smallest" {
  local_disk = true
}

data "databricks_spark_version" "latest_lts" {
  long_term_support = true
}

resource "databricks_cluster" "shared_autoscaling" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = data.databricks_spark_version.latest_lts.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 50
  }
  spark_conf = {
    "spark.databricks.io.cache.enabled" : true,
    "spark.databricks.io.cache.maxDiskUsage" : "50g",
    "spark.databricks.io.cache.maxMetaDataCache" : "1g"
  }
}
```

### Fixed size or autoscaling cluster

When you [create a Databricks cluster](https://docs.databricks.com/clusters/configure.html#cluster-size-and-autoscaling), you can either provide a `num_workers` for the fixed-size cluster or provide `min_workers` and/or `max_workers` for the cluster within the `autoscale` group. When you give a fixed-sized cluster, Databricks ensures that your cluster has a specified number of workers. When you provide a range for the number of workers, Databricks chooses the appropriate number of workers required to run your job - also known as "autoscaling." With autoscaling, Databricks dynamically reallocates workers to account for the characteristics of your job. Certain parts of your pipeline may be more computationally demanding than others, and Databricks automatically adds additional workers during these phases of your job (and removes them when they’re no longer needed).

`autoscale` optional configuration block supports the following:

* `min_workers` - (Optional) The minimum number of workers to which the cluster can scale down when underutilized. It is also the initial number of workers the cluster will have after creation.
* `max_workers` - (Optional) The maximum number of workers to which the cluster can scale up when overloaded. max_workers must be strictly greater than min_workers.

When using a [Single Node cluster](https://docs.databricks.com/clusters/single-node.html), `num_workers` needs to be `0`. It can be set to `0` explicitly, or simply not specified, as it defaults to `0`.  When `num_workers` is `0`, provider checks for presence of the required Spark configurations:

* `spark.master` must have prefix `local`, like `local[*]`
* `spark.databricks.cluster.profile` must have value `singleNode`

and also `custom_tag` entry:

* `"ResourceClass" = "SingleNode"`

The following example demonstrates how to create an single node cluster:

```hcl
data "databricks_node_type" "smallest" {
  local_disk = true
}

data "databricks_spark_version" "latest_lts" {
  long_term_support = true
}

resource "databricks_cluster" "single_node" {
  cluster_name            = "Single Node"
  spark_version           = data.databricks_spark_version.latest_lts.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 20

  spark_conf = {
    # Single-node
    "spark.databricks.cluster.profile" : "singleNode"
    "spark.master" : "local[*]"
  }

  custom_tags = {
    "ResourceClass" = "SingleNode"
  }
}
```

### (Legacy) High-Concurrency clusters

~> This is a legacy cluster type, not related to the real serverless compute. See [Clusters UI changes and cluster access modes](https://docs.databricks.com/archive/compute/cluster-ui-preview.html#legacy) for information on what access mode to use when creating new clusters.

To create High-Concurrency cluster, following settings should be provided:

* `spark_conf` should have following items:
  * `spark.databricks.repl.allowedLanguages` set to a list of supported languages, for example: `python,sql`, or `python,sql,r`.  Scala is not supported!
  * `spark.databricks.cluster.profile` set to `serverless`
* `custom_tags` should have tag `ResourceClass` set to value `Serverless`

For example:

```hcl
resource "databricks_cluster" "cluster_with_table_access_control" {
  cluster_name            = "Shared High-Concurrency"
  spark_version           = data.databricks_spark_version.latest_lts.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 20

  spark_conf = {
    "spark.databricks.repl.allowedLanguages" : "python,sql",
    "spark.databricks.cluster.profile" : "serverless"
  }

  custom_tags = {
    "ResourceClass" = "Serverless"
  }
}
```

### library Configuration Block

To install libraries, one must specify each library in a separate configuration block. Each different type of library has a slightly different syntax. It's possible to set only one type of library within one config block. Otherwise, the plan will fail with an error.

-> Please consider using [databricks_library](library.md) resource for a more flexible setup.

Installing JAR artifacts on a cluster. Location can be anything, that is DBFS or mounted object store (s3, adls, ...)

```hcl
library {
  jar = "dbfs:/FileStore/app-0.0.1.jar"
}
```

Installing Python EGG artifacts. Location can be anything, that is DBFS or mounted object store (s3, adls, ...)

```hcl
library {
  egg = "dbfs:/FileStore/foo.egg"
}
```

Installing Python Wheel artifacts. Location can be anything, that is DBFS or mounted object store (s3, adls, ...)

```hcl
library {
  whl = "dbfs:/FileStore/baz.whl"
}
```

Installing Python PyPI artifacts. You can optionally also specify the `repo` parameter for a custom PyPI mirror, which should be accessible without any authentication for the network that cluster runs in.

```hcl
library {
  pypi {
    package = "fbprophet==0.6"
    // repo can also be specified here
  }
}
```

Installing Python libraries listed in the `requirements.txt` file.  Only Workspace paths and Unity Catalog Volumes paths are supported.  Requires a cluster with DBR 15.0+.

```hcl
library {
  requirements = "/Workspace/path/to/requirements.txt"
}
```

Installing artifacts from Maven repository. You can also optionally specify a `repo` parameter for a custom Maven-style repository, that should be accessible without any authentication. Maven libraries are resolved in Databricks Control Plane, so repo should be accessible from it. It can even be properly configured [maven s3 wagon](https://github.com/seahen/maven-s3-wagon), [AWS CodeArtifact](https://aws.amazon.com/codeartifact/) or [Azure Artifacts](https://azure.microsoft.com/en-us/services/devops/artifacts/).

```hcl
library {
  maven {
    coordinates = "com.amazon.deequ:deequ:1.0.4"
    // exlusions block is optional
    exclusions = ["org.apache.avro:avro"]
  }
}
```

Installing artifacts from CRan. You can also optionally specify a `repo` parameter for a custom cran mirror.

```hcl
library {
  cran {
    package = "rkeops"
  }
}
```

### cluster_log_conf

Example of pushing all cluster logs to DBFS:

```hcl
cluster_log_conf {
  dbfs {
    destination = "dbfs:/cluster-logs"
  }
}
```

Example of pushing all cluster logs to S3:

```hcl
cluster_log_conf {
  s3 {
    destination = "s3://acmecorp-main/cluster-logs"
    region      = "us-east-1"
  }
}
```

Example of pushing all cluster logs to UC Volumes:

```hcl
cluster_log_conf {
  volumes {
    destination = "/Volumes/catalog/schema/cluster_logs_volume/"
  }
}
```

There are a few more advanced attributes for S3 log delivery:

* `destination` - S3 destination, e.g., `s3://my-bucket/some-prefix` You must configure the cluster with an instance profile, and the instance profile must have write access to the destination. You cannot use AWS keys.
* `region` - (Optional) S3 region, e.g. `us-west-2`. Either `region` or `endpoint` must be set. If both are set, the endpoint is used.
* `endpoint` - (Optional) S3 endpoint, e.g. <https://s3-us-west-2.amazonaws.com>. Either `region` or `endpoint` needs to be set. If both are set, the endpoint is used.
* `enable_encryption` - (Optional) Enable server-side encryption, false by default.
* `encryption_type` - (Optional) The encryption type, it could be `sse-s3` or `sse-kms`. It is used only when encryption is enabled, and the default type is `sse-s3`.
* `kms_key` - (Optional) KMS key used if encryption is enabled and encryption type is set to `sse-kms`.
* `canned_acl` - (Optional) Set canned access control list, e.g. `bucket-owner-full-control`. If `canned_cal` is set, the cluster instance profile must have `s3:PutObjectAcl` permission on the destination bucket and prefix. The full list of possible canned ACLs can be found [here](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl). By default, only the object owner gets full control. If you are using a cross-account role for writing data, you may want to set `bucket-owner-full-control` to make bucket owners able to read the logs.

### init_scripts

To run a particular init script on all clusters within the same workspace, both automated/job and interactive/all-purpose cluster types, please consider the [databricks_global_init_script](global_init_script.md) resource.

It is possible to specify up to 10 different cluster-scoped init scripts per cluster.  Init scripts support DBFS, cloud storage locations, and workspace files.

Example of using a Databricks workspace file as init script:

```hcl
init_scripts {
  workspace {
    destination = "/Users/user@domain/install-elk.sh"
  }
}
```

Example of using a file from Unity Catalog Volume as init script:

```hcl
init_scripts {
  volumes {
    destination = "/Volumes/Catalog/default/init-scripts/init-script.sh"
  }
}
```

Example of taking init script from DBFS (deprecated):

```hcl
init_scripts {
  dbfs {
    destination = "dbfs:/init-scripts/install-elk.sh"
  }
}
```

Example of taking init script from S3:

```hcl
init_scripts {
  s3 {
    destination = "s3://acmecorp-main/init-scripts/install-elk.sh"
    region      = "us-east-1"
  }
}
```

Similarly, for an init script stored in GCS:

```hcl
init_scripts {
  gcs {
    destination = "gs://init-scripts/install-elk.sh"
  }
}
```

Similarly, for an init script stored in ADLS:

```hcl
init_scripts {
  abfss {
    destination = "abfss://container@storage.dfs.core.windows.net/install-elk.sh"
  }
}
```

Please note that you need to provide Spark Hadoop configuration (`spark.hadoop.fs.azure...`) to authenticate to ADLS to get access to the init script.

Clusters with [custom Docker containers](https://docs.databricks.com/clusters/custom-containers.html) also allow a local file location for init scripts as follows:

```hcl
init_scripts {
  file {
    destination = "file:/my/local/file.sh"
  }
}
```

### aws_attributes

`aws_attributes` optional configuration block contains attributes related to [clusters running on Amazon Web Services](https://docs.databricks.com/clusters/configure.html#aws-configurations).

Here is the example of shared autoscaling cluster with some of AWS options set:

```hcl
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" {
  local_disk = true
}
resource "databricks_cluster" "this" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 50
  }
  aws_attributes {
    availability           = "SPOT"
    zone_id                = "us-east-1"
    first_on_demand        = 1
    spot_bid_price_percent = 100
  }
}
```

The following options are available:

* `zone_id` - (Required) Identifier for the availability zone/datacenter in which the cluster resides. This string will be of a form like `us-west-2a`. The provided availability zone must be in the same region as the Databricks deployment. For example, `us-west-2a` is not a valid zone ID if the Databricks deployment resides in the `us-east-1` region. Enable automatic availability zone selection ("Auto-AZ"), by setting the value `auto`. Databricks selects the AZ based on available IPs in the workspace subnets and retries in other availability zones if AWS returns insufficient capacity errors.
* `availability` - (Optional) Availability type used for all subsequent nodes past the `first_on_demand` ones. Valid values are `SPOT`, `SPOT_WITH_FALLBACK` and `ON_DEMAND`. Note: If `first_on_demand` is zero, this availability type will be used for the entire cluster. Backend default value is `SPOT_WITH_FALLBACK` and could change in the future
* `first_on_demand` - (Optional) The first `first_on_demand` nodes of the cluster will be placed on on-demand instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances. If this value is less than the current cluster size, `first_on_demand` nodes will be placed on on-demand instances, and the remainder will be placed on availability instances. This value does not affect cluster size and cannot be mutated over the lifetime of a cluster. If unspecified, the default value is 0.
* `spot_bid_price_percent` - (Optional) The max price for AWS spot instances, as a percentage of the corresponding instance type’s on-demand price. For example, if this field is set to 50, and the cluster needs a new `i3.xlarge` spot instance, then the max price is half of the price of on-demand `i3.xlarge` instances. Similarly, if this field is set to 200, the max price is twice the price of on-demand `i3.xlarge` instances. If not specified, the default value is `100`. When spot instances are requested for this cluster, only spot instances whose max price percentage matches this field will be considered. For safety, we enforce this field to be no more than `10000`.
* `instance_profile_arn` - (Optional) Nodes for this cluster will only be placed on AWS instances with this instance profile. Please see [databricks_instance_profile](instance_profile.md) resource documentation for extended examples on adding a valid instance profile using Terraform.
* `ebs_volume_type` - (Optional) The type of EBS volumes that will be launched with this cluster. Valid values are `GENERAL_PURPOSE_SSD` or `THROUGHPUT_OPTIMIZED_HDD`. Use this option only if you're not picking *Delta Optimized `i3.*`* node types.
* `ebs_volume_count` - (Optional) The number of volumes launched for each instance. You can choose up to 10 volumes. This feature is only enabled for supported node types. Legacy node types cannot specify custom EBS volumes. For node types with no instance store, at least one EBS volume needs to be specified; otherwise, cluster creation will fail. These EBS volumes will be mounted at /ebs0, /ebs1, and etc. Instance store volumes will be mounted at /local_disk0, /local_disk1, and etc. If EBS volumes are attached, Databricks will configure Spark to use only the EBS volumes for scratch storage because heterogeneously sized scratch devices can lead to inefficient disk utilization. If no EBS volumes are attached, Databricks will configure Spark to use instance store volumes. If EBS volumes are specified, then the Spark configuration spark.local.dir will be overridden.
* `ebs_volume_size` - (Optional) The size of each EBS volume (in GiB) launched for each instance. For general purpose SSD, this value must be within the range 100 - 4096. For throughput optimized HDD, this value must be within the range 500 - 4096. Custom EBS volumes cannot be specified for the legacy node types (memory-optimized and compute-optimized).

### azure_attributes

`azure_attributes` optional configuration block contains attributes related to [clusters running on Azure](https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/clusters#--azureattributes).

Here is the example of shared autoscaling cluster with some of Azure options set:

```hcl
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" {
  local_disk = true
}
resource "databricks_cluster" "this" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 50
  }
  azure_attributes {
    availability       = "SPOT_WITH_FALLBACK_AZURE"
    first_on_demand    = 1
    spot_bid_max_price = 100
  }
}
```

The following options are [available](https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/clusters#--azureattributes):

* `availability` - (Optional) Availability type used for all subsequent nodes past the `first_on_demand` ones. Valid values are `SPOT_AZURE`, `SPOT_WITH_FALLBACK_AZURE`, and `ON_DEMAND_AZURE`. Note: If `first_on_demand` is zero, this availability type will be used for the entire cluster.
* `first_on_demand` - (Optional) The first `first_on_demand` nodes of the cluster will be placed on on-demand instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances. If this value is less than the current cluster size, `first_on_demand` nodes will be placed on on-demand instances, and the remainder will be placed on availability instances. This value does not affect cluster size and cannot be mutated over the lifetime of a cluster.
* `spot_bid_max_price` - (Optional) The max bid price used for Azure spot instances. You can set this to greater than or equal to the current spot price. You can also set this to `-1`, which specifies that the instance cannot be evicted on the basis of price. The price for the instance will be the current price for spot instances or the price for a standard instance.

### gcp_attributes

`gcp_attributes` optional configuration block contains attributes related to [clusters running on GCP](https://docs.gcp.databricks.com/dev-tools/api/latest/clusters.html#clustergcpattributes).

Here is the example of shared autoscaling cluster with some of GCP options set:

```hcl
resource "databricks_cluster" "this" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 50
  }
  gcp_attributes {
    availability = "PREEMPTIBLE_WITH_FALLBACK_GCP"
    zone_id      = "AUTO"
  }
}
```

The following options are available:

* `use_preemptible_executors` - (Optional, bool) if we should use preemptible executors ([GCP documentation](https://cloud.google.com/compute/docs/instances/preemptible)). *Warning: this field is deprecated in favor of `availability`, and will be removed soon.*
* `google_service_account` - (Optional, string) Google Service Account email address that the cluster uses to authenticate with Google Identity. This field is used for authentication with the GCS and BigQuery data sources.
* `availability` - (Optional) Availability type used for all nodes. Valid values are `PREEMPTIBLE_GCP`, `PREEMPTIBLE_WITH_FALLBACK_GCP` and `ON_DEMAND_GCP`, default: `ON_DEMAND_GCP`.
* `boot_disk_size` (optional, int) Boot disk size in GB
* `local_ssd_count` (optional, int) Number of local SSD disks (each is 375GB in size) that will be attached to each node of the cluster.
* `zone_id` (optional)  Identifier for the availability zone in which the cluster resides. This can be one of the following:
  * `HA` (default): High availability, spread nodes across availability zones for a Databricks deployment region.
  * `AUTO`: Databricks picks an availability zone to schedule the cluster on.
  * name of a GCP availability zone: pick one of the available zones from the [list of available availability zones](https://cloud.google.com/compute/docs/regions-zones#available).

### docker_image

[Databricks Container Services](https://docs.databricks.com/clusters/custom-containers.html) lets you specify a Docker image when you create a cluster. You need to enable Container Services in *Admin Console /  Advanced* page in the user interface. By enabling this feature, you acknowledge and agree that your usage of this feature is subject to the [applicable additional terms](http://www.databricks.com/product-specific-terms).

`docker_image` configuration block has the following attributes:

* `url` - URL for the Docker image
* `basic_auth` - (Optional) `basic_auth.username` and `basic_auth.password` for Docker repository. Docker registry credentials are encrypted when they are stored in Databricks internal storage and when they are passed to a registry upon fetching Docker images at cluster launch. However, other authenticated and authorized API users of this workspace can access the username and password.

Example usage with [azurerm_container_registry](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/container_registry) and [docker_registry_image](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/registry_image), that you can adapt to your specific use-case:

```hcl
resource "docker_registry_image" "this" {
  name = "${azurerm_container_registry.this.login_server}/sample:latest"
  build {
    # ...
  }
}

resource "databricks_cluster" "this" {
  # ...
  docker_image {
    url = docker_registry_image.this.name
    basic_auth {
      username = azurerm_container_registry.this.admin_username
      password = azurerm_container_registry.this.admin_password
    }
  }
}
```

### cluster_mount_info blocks (experimental)

~> The underlying API is experimental and may change in the future.

It's possible to mount NFS (Network File System) resources into the Spark containers inside the cluster.  You can specify one or more `cluster_mount_info` blocks describing the mount. This block has following attributes:

* `network_filesystem_info` - block specifying connection. It consists of:
  * `server_address` - (Required) host name.
  * `mount_options` - (Optional) string that will be passed as options passed to the `mount` command.
* `remote_mount_dir_path` - (Optional) string specifying path to mount on the remote service.
* `local_mount_dir_path` - (Required) path inside the Spark container.

For example, you can mount Azure Data Lake Storage container using the following code:

```hcl
locals {
  storage_account   = "ewfw3ggwegwg"
  storage_container = "test"
}

resource "databricks_cluster" "with_nfs" {
  # ...
  cluster_mount_info {
    network_filesystem_info {
      server_address = "${local.storage_account}.blob.core.windows.net"
      mount_options  = "sec=sys,vers=3,nolock,proto=tcp"
    }
    remote_mount_dir_path = "${local.storage_account}/${local.storage_container}"
    local_mount_dir_path  = "/mnt/nfs-test"
  }
}
```

### workload_type block

It's possible to restrict which workloads may run on the given cluster - notebooks and/or jobs. It's done by defining a `workload_type` block that consists of a single block `clients` with following attributes:

* `notebooks` - (Optional) boolean flag defining if it's possible to run notebooks on this cluster. Default: `true`.
* `jobs` - (Optional) boolean flag defining if it's possible to run Databricks Jobs on this cluster. Default: `true`.

```hcl
resource "databricks_cluster" "with_nfs" {
  # ...
  workload_type {
    clients {
      jobs      = false
      notebooks = true
    }
  }
}
```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the cluster.
* `default_tags` - (map) Tags that are added by Databricks by default, regardless of any `custom_tags` that may have been added. These include: Vendor: Databricks, Creator: <username_of_creator>, ClusterName: <name_of_cluster>, ClusterId: <id_of_cluster>, Name: <Databricks internal use>, and any workspace and pool tags.
* `state` - (string) State of the cluster.

## Access Control

* [databricks_group](group.md#allow_cluster_create) and [databricks_user](user.md#allow_cluster_create) can control which groups or individual users can create clusters.
* [databricks_cluster_policy](cluster_policy.md) can control which kinds of clusters users can create.
* Users, who have access to Cluster Policy, but do not have an `allow_cluster_create` argument set would still be able to create clusters, but within the boundary of the policy.
* [databricks_permissions](permissions.md#Cluster-usage) can control which groups or individual users can *Manage*, *Restart* or *Attach to* individual clusters.
* `instance_profile_arn` *(AWS only)* can control which data a given cluster can access through cloud-native controls.

## Import

The resource cluster can be imported using cluster id.

```bash
terraform import databricks_cluster.this <cluster-id>
```

## Related Resources

The following resources are often used in the same context:

* [Dynamic Passthrough Clusters for a Group](../guides/workspace-management.md) guide.
* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_clusters](../data-sources/clusters.md) data to retrieve a list of [databricks_cluster](cluster.md) ids.
* [databricks_cluster_policy](cluster_policy.md) to create a [databricks_cluster](cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_current_user](../data-sources/current_user.md) data to retrieve information about [databricks_user](user.md) or [databricks_service_principal](service_principal.md), that is calling Databricks REST API.
* [databricks_global_init_script](global_init_script.md) to manage [global init scripts](https://docs.databricks.com/clusters/init-scripts.html#global-init-scripts), which are run on all [databricks_cluster](cluster.md#init_scripts) and [databricks_job](job.md#new_cluster).
* [databricks_instance_pool](instance_pool.md) to manage [instance pools](https://docs.databricks.com/clusters/instance-pools/index.html) to reduce [cluster](cluster.md) start and auto-scaling times by maintaining a set of idle, ready-to-use instances.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).
* [databricks_library](library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md).
* [databricks_mount](mount.md) to [mount your cloud storage](https://docs.databricks.com/data/databricks-file-system.html#mount-object-storage-to-dbfs) on `dbfs:/mnt/name`.
* [databricks_node_type](../data-sources/node_type.md) data to get the smallest node type for [databricks_cluster](cluster.md) that fits search criteria, like amount of RAM or number of cores.
* [databricks_pipeline](pipeline.md) to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html).
* [databricks_spark_version](../data-sources/spark_version.md) data to get [Databricks Runtime (DBR)](https://docs.databricks.com/runtime/dbr.html) version that could be used for `spark_version` parameter in [databricks_cluster](cluster.md) and other resources.
* [databricks_zones](../data-sources/zones.md) data to fetch all available AWS availability zones on your workspace on AWS.
