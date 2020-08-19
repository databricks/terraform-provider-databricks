# databricks_cluster Resource

This resource allows you to create, update, and delete clusters.

```hcl
resource "databricks_cluster" "shared_autoscaling" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = "6.6.x-scala2.11"
  node_type_id            = "i3.xlarge"
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 50
  }
}
```

## Argument Reference

* `cluster_name` - (Optional) Cluster name. This doesn’t have to be unique. If not specified at creation, the cluster name will be an empty string.
* `spark_version` - (Required) [Runtime version](https://docs.databricks.com/runtime/index.html) of the cluster. A [list of available Spark versions](https://docs.databricks.com/release-notes/runtime/releases.html) can be retrieved by using the Runtime Versions API call or `databricks clusters spark-versions` CLI command. It is advised to use [Cluster Policies](cluster_policy.md) to restrict list of versions for simplicity, while maintaining enough of control.
* `driver_node_type_id` - (Optional) The node type of the Spark driver. This field is optional; if unset, the driver node type will be set as the same value as node_type_id defined above.
* `node_type_id` - (Required - optional if `instance_pool_id` is given) This field encodes, through a single value, the resources available to each of the Spark nodes in this cluster. For example, the Spark nodes can be provisioned and optimized for memory or compute intensive workloads A list of available node types can be retrieved by using the List Node Types API call. If `instance_pool_id` is specified, this field is not needed.
* `instance_pool_id` (Optional - required if `node_type_id` is not given) - To reduce cluster start time, you can attach a cluster to a predefined pool of idle instances. When attached to a pool, a cluster allocates its driver and worker nodes from the pool. If the pool does not have sufficient idle resources to accommodate the cluster’s request, the pool expands by allocating new instances from the instance provider. When an attached cluster is terminated, the instances it used are returned to the pool and can be reused by a different cluster.
* `policy_id` - (Optional) Idendifier of [Custer Policy](cluster_policy.md) to validate cluster and preset certail defaults. Cluster policy has bigger use when allowing users to create clusters, rather than automatically created one. Essentially, you can put all cluster configuration options into it.
* `autotermination_minutes` - (Optional) Automatically terminates the cluster after it is inactive for this time in minutes. If not set, this cluster will not be automatically terminated. If specified, the threshold must be between 10 and 10000 minutes. You can also set this value to 0 to explicitly 
disable automatic termination. _It is highly recommended to have this setting present for Interfactive/BI clusters._
* `enable_elastic_disk` - (Optional) If you don’t want to allocate a fixed number of EBS volumes at cluster creation time, use autoscaling local storage. With autoscaling local storage, Databricks monitors the amount of free disk space available on your cluster’s Spark workers. If a worker begins to run too low on disk, Databricks automatically attaches a new EBS volume to the worker before it runs out of disk space. EBS volumes are attached up to a limit of 5 TB of total disk space per instance (including the instance’s local storage). To scale down EBS usage, make sure you have `autotermination_minutes` and `autoscale` attributes set. More documentation available at [cluster configuration page](https://docs.databricks.com/clusters/configure.html#autoscaling-local-storage-1).
* `single_user_name` - (Optional) The optional user name of the user to assign to an interactive cluster. This is required when using standard AAD Passthrough for Azure Datalake Storage (ADLS) with a single-user cluster (i.e. not high-concurrency clusters).
* `idempotency_token` - (Optional) An optional token that can be used to guarantee the idempotency of cluster creation requests. If an active cluster with the provided token already exists, the request will not create a new cluster, but it will return the ID of the existing cluster instead. The existence of a cluster with the same token is not checked against terminated clusters. If you specify the idempotency token, upon failure you can retry until the request succeeds. Databricks will guarantee that exactly one cluster will be launched with that idempotency token. This token should have at most 64 characters.
* `ssh_public_keys` - (Optional) SSH public key contents that will be added to each Spark node in this cluster. The corresponding private keys can be used to login with the user name ubuntu on port 2200. Up to 10 keys can be specified.
TODO: add example
* `spark_env_vars` - (Optional) Map with environment variable key-value pairs to fine tune Spark clusters. Key-value pair of the form (X,Y) are exported as is (i.e., export X='Y') while launching the driver and workers. To specify an additional set of SPARK_DAEMON_JAVA_OPTS, we recommend appending them to $SPARK_DAEMON_JAVA_OPTS as shown in the example below. This ensures that all default databricks managed environmental variables are included as well. 
* `custom_tags` - (Optional) Additional tags for cluster resources. Databricks will tag all cluster resources (e.g., AWS instances and EBS volumes) with these tags in addition to `default_tags`.
* `spark_conf` - (Optional) Map with key-value pairs to fine tune Spark clusters, where you can provide custom pSpark configuration properties](https://spark.apache.org/docs/latest/configuration.html) in a cluster configuration. You can also pass in a string of extra JVM options to the driver and the executors via `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions` respectively. It is advised to keep all common configurations in [Cluster Policies](cluster_policy.md) to maintain control of the environments launched.

The following example demonstrates how to create an autoscaling cluster with [Delta Cache](https://docs.databricks.com/delta/optimizations/delta-cache.html) enabled:

```hcl
resource "databricks_cluster" "shared_autoscaling" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = "6.6.x-scala2.11"
  node_type_id            = "i3.xlarge"
  autotermination_minutes = 20
  autoscale {
      min_workers = 1
      max_workers = 50
  }
  spark_conf {
      "spark.databricks.io.cache.enabled": true,
      "spark.databricks.io.cache.maxDiskUsage": "50g",
      "spark.databricks.io.cache.maxMetaDataCache": "1g"
  }
}
```

## Fixed size or autoscaling cluster

When you [create a Databricks cluster](https://docs.databricks.com/clusters/configure.html#cluster-size-and-autoscaling), you can either provide a `num_workers` for the fixed size cluster or provide `min_workers` and/or `max_workers` for the cluster withing `autoscale` group. When you provide a fixed size cluster, Databricks ensures that your cluster has the specified number of workers. When you provide a range for the number of workers, Databricks chooses the appropriate number of workers required to run your job. This is referred to as autoscaling. With autoscaling, Databricks dynamically reallocates workers to account for the characteristics of your job. Certain parts of your pipeline may be more computationally demanding than others, and Databricks automatically adds additional workers during these phases of your job (and removes them when they’re no longer needed). It is advised to keep all common configurations in [Cluster Policies](cluster_policy.md) to maintain control of the environments launched.

* `num_workers` - (Optional) Number of worker nodes that this cluster should have. A cluster has one Spark Driver and num_workers Executors for a total of num_workers + 1 Spark node.

`autoscale` optional configuration block supports the following:

* `min_workers` - (Optional) The minimum number of workers to which the cluster can scale down when underutilized. It is also the initial number of workers the cluster will have after creation.
* `max_workers` - (Optional) The maximum number of workers to which the cluster can scale up when overloaded. max_workers must be strictly greater than min_workers.

### libraries Configuration Block

In order to install libraries, one must specify each library in own configuration block. Each different type of library has slightly different syntax. It's possible to specify only one type of library within one config block, otherwise plan will fail with error.

Installing JAR artifacts on a cluster. Location can be anyling, that is DBFS or mounted object store (s3, adls, ...)
```hcl
libraries {
  jar = "dbfs://FileStore/app-0.0.1.jar"
}
```

Installing Python EGG artifacts. Location can be anyling, that is DBFS or mounted object store (s3, adls, ...)
```hcl
libraries {
  egg = "dbfs://FileStore/foo.egg"
}
```

Installing Python Wheel artifacts. Location can be anyling, that is DBFS or mounted object store (s3, adls, ...)
```hcl
libraries {
  whl = "dbfs://FileStore/baz.whl"
}
```

Installing Python PyPI artifacts. You can also optionally also specify `repo` parameter for custom PyPI mirror, that should be accessible without any authentication for the network, that cluster runs in.
```hcl
libraries {
  pypi {
    package = "fbprophet==0.6"
    // repo can also be specified here
  }
}
```

Installing artifacts from Maven repository. You can also optionally also specify `repo` parameter for custom Maven-style repository, that should be accessible without any authentication for the network, that cluster runs in. It can even be properly configured [maven s3 wagon](https://github.com/seahen/maven-s3-wagon), [AWS CodeArtifact](https://aws.amazon.com/codeartifact/) or [Azure Artifacts](https://azure.microsoft.com/en-us/services/devops/artifacts/).
```hcl
libraries {
  maven {
    coordinates = "com.amazon.deequ:deequ:1.0.4"
    // exlusions block is optional
    exclusions = ["org.apache.avro:avro"]
  }
}
```

Installing artifacts from CRan. You can also optionally also specify `repo` parameter for custom cran mirror.
```hcl
libraries {
  cran {
    package = "rkeops"
  }
}
```

## cluster_log_conf

Example of pushing all cluster logs to DBFS:
```hcl
cluster_log_conf {
  dbfs {
    destination = "dbfs://cluster-logs"
  }
}
```

Example of pushing all cluster logs to S3:
```hcl
cluster_log_conf {
  s3 {
    destination = "s3a://acmecorp-main/cluster-logs"
    region = "us-east-1"
  }
}
```

There are few more advanced attributes for S3 log delivery:

* `destination` - S3 destination, e.g. `s3://my-bucket/some-prefix` You must configure the cluster with an instance profile and the instance profile must have write access to the destination. You cannot use AWS keys.  
* `region` - (Optional) S3 region, e.g. `us-west-2`. Either region or endpoint must be set. If both are set, endpoint is used.
* `endpoint` - (Optional) S3 endpoint, e.g. https://s3-us-west-2.amazonaws.com. Either region or endpoint needs to be set. If both are set, endpoint is used.
* `enable_encryption` - (Optional) Enable server side encryption, false by default.
* `encryption_type` - (Optional) The encryption type, it could be sse-s3 or sse-kms. It is used only when encryption is enabled and the default type is sse-s3.
* `kms_key` - (Optional) KMS key used if encryption is enabled and encryption type is set to sse-kms.
* `canned_acl` - (Optional) Set canned access control list, e.g. bucket-owner-full-control. If canned_cal is set, the cluster instance profile must have s3:PutObjectAcl permission on the destination bucket and prefix. The full list of possible canned ACL can be found [here](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl). By default only the object owner gets full control. If you are using cross account role for writing data, you may 
  want to set bucket-owner-full-control to make bucket owner able to read the logs.

## init_scripts

You can specify up to 10 different init scripts for cluster.

Example of taking init script from DBFS:
```hcl
init_scripts {
  dbfs {
    destination = "dbfs://init-scripts/install-elk.sh"
  }
}
```

Example of taking init script from S3:
```hcl
init_scripts {
  s3 {
    destination = "s3a://acmecorp-main/init-scripts/install-elk.sh"
    region = "us-east-1"
  }
}
```

Attributes are the same as for `cluster_log_conf` configuration block.

## aws_attributes

`aws_attributes` optional configuration block contains attributes related to [clusters running on Amazon Web Services](https://docs.databricks.com/clusters/configure.html#aws-configurations). If not specified at cluster creation, a set of default values will be used. It is advised to keep all common configurations in [Cluster Policies](cluster_policy.md) to maintain control of the environments launched.

Here is the example of shared autoscaling cluster with some of AWS options set:

```hcl
resource "databricks_cluster" "this" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = "6.6.x-scala2.11"
  node_type_id            = "i3.xlarge"
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 50
  }
  aws_attributes {
    availability            = "SPOT"
    zone_id                 = "us-east-1"
    first_on_demand         = 1
    spot_bid_price_percent  = 100
  }
}
```

The following options are available:

* `zone_id` - (Required) Identifier for the availability zone/datacenter in which the cluster resides. This string will be of a form like “us-west-2a”. The provided availability zone must be in the same region as the Databricks deployment. For example, “us-west-2a” is not a valid zone ID if the Databricks deployment resides in the “us-east-1” region. 
* `availability` - (Optional) Availability type used for all subsequent nodes past the `first_on_demand` ones. Valid values are `SPOT` and `ON_DEMAND`. Note: If `first_on_demand` is zero, this availability type will be used for the entire cluster. 
* `first_on_demand` - (Optional) The first first_on_demand nodes of the cluster will be placed on on-demand instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances. If this value is less than the current cluster size, first_on_demand nodes will be placed on on-demand instances and the remainder will be placed on availability instances. This value does not affect cluster size and cannot be mutated over the lifetime of a cluster.
* `spot_bid_price_percent` - (Optional) The max price for AWS spot instances, as a percentage of the corresponding instance type’s on-demand price. For example, if this field is set to 50, and the cluster needs a new `i3.xlarge` spot instance, then the max price is half of the price of on-demand `i3.xlarge` instances. Similarly, if this field is set to 200, the max price is twice the price of on-demand `i3.xlarge` instances. If not specified, the default value is `100`. When spot instances are requested for this cluster, only spot instances whose max price percentage matches this field will be considered. For safety, we enforce this field to be no more than `10000`.
* `instance_profile_arn` - (Optional) Nodes for this cluster will only be placed on AWS instances with this instance profile. Please see [databricks_instance_profile](instance_profile.md) resource documentation for extended examples on how to add a valid instance profile using Terraform.
* `ebs_volume_type` - (Optional) The type of EBS volumes that will be launched with this cluster. Valid values are `GENERAL_PURPOSE_SSD` or `THROUGHPUT_OPTIMIZED_HDD`. Use this option only if you're not picking _Delta Optinized `i3.*`_ node types.
* `ebs_volume_count` - (Optional) The number of volumes launched for each instance. You can choose up to 10 volumes. This feature is only enabled for supported node types. Legacy node types cannot specify custom EBS volumes. For node types with no instance store, at least one EBS volume needs to be specified; otherwise, cluster creation will fail. These EBS volumes will be mounted at /ebs0, /ebs1, and etc. Instance store volumes will be mounted at /local_disk0, /local_disk1, and etc. If EBS volumes are attached, Databricks will configure Spark to use only the EBS volumes for scratch storage because heterogeneously sized scratch devices can lead to inefficient disk utilization. If no EBS volumes are attached, Databricks will configure Spark to use instance store volumes. If EBS volumes are specified, then the Spark configuration spark.local.dir will be overridden.
* `ebs_volume_size` - (Optional) The size of each EBS volume (in GiB) launched for each instance. For general purpose SSD, this value must be within the range 100 - 4096. For throughput optimized HDD, this value must be within the range 500 - 4096. Custom EBS volumes cannot be specified for the legacy node types (memory-optimized and compute-optimized).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the cluster.
* `default_tags` - (map) Tags that are added by Databricks by default, regardless of any custom_tags that may have been added. These include: Vendor: Databricks, Creator: <username_of_creator>, ClusterName: <name_of_cluster>, ClusterId: <id_of_cluster>, Name: <Databricks internal use>
* `state` - (string) State of the cluster.


## Import

The resource cluster can be imported using cluster id

```bash
$ terraform import databricks_cluster.this <cluster-id>
```