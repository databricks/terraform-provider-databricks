+++
title = "job"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_job`

The databricks_job resource allows you to create, edit, and delete jobs. 


## Example Usage

### Databricks Example

{{< tabs groupId="jobClusterTypes" >}}
{{% tab name="New Automated Cluster" %}}
```hcl
resource "databricks_job" "my_job3" {
  new_cluster {
    autoscale {
      min_workers = 2
      max_workers = 3
    }
    spark_version = "6.4.x-scala2.11"
    aws_attributes {
      zone_id = data.databricks_zones.my_wspace_zones.default_zone
      spot_bid_price_percent = "100"
    }
    node_type_id = "r3.xlarge"
  }
  notebook_path = "/Users/jane.doe@databricks.com/my-demo-notebook"
  name = "my-demo-notebook"
  timeout_seconds = 3600
  max_retries = 1
  max_concurrent_runs = 1
}
```
{{% /tab %}}
{{% tab name="Existing Cluster" %}}
```hcl
resource "databricks_job" "my_job3" {
  existing_cluster_id = "<Cluster ID>"
  notebook_path = "/Users/jane.doe@databricks.com/my-demo-notebook"
  name = "my-demo-notebook"
  timeout_seconds = 3600
  max_retries = 1
  max_concurrent_runs = 1
}
```
{{% /tab %}}
{{< /tabs >}}
    
## Argument Reference

The following arguments are supported:

#### - `existing_cluster_id`:
> **(Optional)** If existing_cluster_id, the ID of an existing 
cluster that will be used for all runs of this job. When running jobs on an existing cluster, you may need to manually 
restart the cluster if it stops responding. We suggest running jobs on new clusters for greater reliability.

#### - `new_cluster`:
> **(Optional)** A description of a cluster that will be created for each run. 
Please look [below](#new-cluster).

#### - `name`:
> **(Optional)** An optional name for the job. The default value is Untitled.

#### - `library_jar`:
> **(Optional)** URI of the JAR to be installed. 
DBFS and S3 URIs are supported. For example: "dbfs:/mnt/databricks/library.jar", "s3://my-bucket/library.jar". 
If S3 is used, make sure the cluster has read access on the library. You may need to launch the cluster with an 
instance profile to access the S3 URI.

#### - `library_egg`:
> **(Optional)**  URI of the egg to be installed. 
DBFS and S3 URIs are supported. For example: "dbfs:/my/egg", "s3://my-bucket/egg" }. 
If S3 is used, make sure the cluster has read access on the library. You may need to launch the cluster 
with an instance profile to access the S3 URI.

#### - `library_whl`:
> **(Optional)** If whl, URI of the wheel or zipped wheels to be installed. 
DBFS and S3 URIs are supported. For example: "dbfs:/my/whl", "s3://my-bucket/whl". If S3 is used, make sure the cluster 
has read access on the library. You may need to launch the cluster with an instance profile to access the S3 URI. 
Also the wheel file name needs to use the correct convention. If zipped wheels are to be installed, the file name 
suffix should be .wheelhouse.zip.

#### - `library_pypi` **(Optional)** :
> {{%chevron default="Specification of a PyPI library to be installed." display="true" %}}

* `package` - **(Required)**  The name of the PyPI package to install. An optional exact version specification 
is also supported. Examples: simplejson and simplejson==3.8.0. This field is required.

* `repo` - **(Optional)** The repository where the package can be found. If not specified, 
the default pip index is used.

{{% /chevron %}}

#### - `library_maven` **(Optional)** :
> {{%chevron default="Specification of a Maven library to be installed." display="true" %}}

* `coordinates` - **(Required)** Gradle-style Maven coordinates. For example: org.jsoup:jsoup:1.7.2. 
This field is required.

* `repo` - **(Optional)** Maven repo to install the Maven package from. If omitted, both Maven Central 
Repository and Spark Packages are searched.

* `exclusions` - **(Optional)** List of dependences to exclude. For example: 
("slf4j:slf4j", "*:hadoop-client"). 

{{% /chevron %}}
    
#### - `library_cran` **(Optional)** :
> {{%chevron default="Specification of a CRAN library to be installed." display="true" %}}

* `package` - **(Required)** The name of the CRAN package to install. This field is required.

* `repo` - **(Optional)** The repository where the package can be found. If not specified, 
the default CRAN repo is used.

{{% /chevron %}}

#### - `notebook_path`:
> **(Optional)** The absolute path of the notebook to be run in the 
Databricks Workspace. This path must begin with a slash. This field is required.

#### - `notebook_base_parameters`:
> **(Optional)** Base parameters to be used 
for each run of this job. If the run is initiated by a call to run-now with parameters specified, the two parameters 
maps will be merged. If the same key is specified in base_parameters and in run-now, the value from run-now will be used.

#### - `jar_uri`:
> **(Optional)** Deprecated since 04/2016. Provide a jar through the libraries 
field instead.

#### - `jar_main_class_name`:
> **(Optional)** The full name of the class containing the 
main method to be executed. This class must be contained in a JAR provided as a library. The code should use 
SparkContext.getOrCreate to obtain a Spark context; otherwise, runs of the job will fail.

#### - `jar_parameters`:
> **(Optional)** Parameters that will be passed to the main method.

#### - `python_file`:
> **(Optional)** The URI of the Python file to be executed. DBFS and S3 paths 
are supported. This field is required.

#### - `python_parameters`:
> **(Optional)** Command line parameters that will be passed to 
the Python file.

#### - `spark_submit_parameters`:
> **(Optional)** Command-line parameters passed 
to spark submit.

#### - `email_notifications` **(Optional)** :
> {{%chevron default=`An optional set of email addresses notified 
when runs of this job begin and complete and when this job is deleted. The default behavior is to not send any emails.` 
display="true" %}}

* `on_start` - **(Optional)** A list of email addresses to be notified when a run begins. 
If not specified upon job creation or reset, the list will be empty, i.e., no address will be notified.

* :`on_success` - **(Optional)** A list of email addresses to be notified when a run successfully completes. 
A run is considered to have completed successfully if it ends with a TERMINATED life_cycle_state and a SUCCESSFUL 
result_state. If not specified upon job creation or reset, the list will be empty, i.e., no address will be notified.

* `on_failure` - **(Optional)** A list of email addresses to be notified when a run unsuccessfully completes. 
A run is considered to have completed unsuccessfully if it ends with an INTERNAL_ERROR life_cycle_state or a 
SKIPPED, FAILED, or TIMED_OUT result_state. If not specified upon job creation or reset, the list will be empty, 
i.e., no address will be notified.

* `no_alert_for_skipped_runs` - **(Optional)** If true, do not send email to recipients specified in 
on_failure if the run is skipped.

{{% /chevron %}}

#### - `timeout_seconds`:
> **(Optional)** 	An optional timeout applied to each run of this job. 
The default behavior is to have no timeout.

#### - `max_retries`:
> **(Optional)** An optional maximum number of times to retry an unsuccessful 
run. A run is considered to be unsuccessful if it completes with a FAILED result_state or INTERNAL_ERROR 
life_cycle_state. The value -1 means to retry indefinitely and the value 0 means to never retry. The default behavior 
is to never retry.

#### - `min_retry_interval_millis`:
> **(Optional)** 	An optional minimal interval in 
milliseconds between the start of the failed run and the subsequent retry run. The default behavior is that 
unsuccessful runs are immediately retried.

#### - `retry_on_timeout`:
> **(Optional)** 	An optional policy to specify whether to retry 
a job when it times out. The default behavior is to not retry on timeout.

#### - `job_schedule` **(Optional)** :
> {{%chevron default=`An optional periodic schedule for this job. 
The default behavior is that the job runs when triggered by clicking Run Now in the Jobs UI or sending an API 
request to runNow.` display="true" %}}

* `quartz_cron_expression` - **(Optional)** 	A Cron expression using Quartz syntax that describes 
the schedule for a job. See [Cron Trigger](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) 
for details. This field is required.

* `timezone_id` - **(Optional)** A Java timezone ID. The schedule for a job will be resolved with respect 
to this timezone. See [Java TimeZone](https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html) for 
details. This field is required.

{{% /chevron %}}

#### - `max_concurrent_runs`:
> **(Optional)** An optional maximum allowed number of 
concurrent runs of the job. Set this value if you want to be able to execute multiple runs of the same job concurrently. 
This is useful for example if you trigger your job on a frequent schedule and want to allow consecutive runs to overlap
with each other, or if you want to trigger multiple runs which differ by their input parameters. This setting affects 
only new runs. For example, suppose the job’s concurrency is 4 and there are 4 concurrent active runs. Then setting 
the concurrency to 3 won’t kill any of the active runs. However, from then on, new runs are skipped unless there are 
fewer than 3 active runs. This value cannot exceed 150. Setting this value to 0 causes all new runs to be skipped. 
The default behavior is to allow only 1 concurrent run.


### New Cluster


#### - `num_workers` **(Optional)** :
> **(Optional)** Number of worker nodes that this cluster should have. 
A cluster has one Spark Driver and num_workers Executors for a total of num_workers + 1 Spark node

#### - `autoscale`  **(Optional)** :
> {{%chevron default="Parameters needed in order to automatically scale clusters up and down based on load." display="true" %}}

#### **Usage**

```hcl
resource "databricks_cluster" "my-cluster" {
...
    autoscale {
      min_workers = 2
      max_workers = 3
    }
...
}
```

* `min_workers` - **(Optional)** The minimum number of workers to which the cluster can scale down when 
underutilized. It is also the initial number of workers the cluster will have after creation.

* `max_workers` - **(Optional)** 	The maximum number of workers to which the cluster can scale up when 
overloaded. max_workers must be strictly greater than min_workers.
    
{{% /chevron %}}

#### - `cluster_name`:
> **(Optional)** Cluster name. This doesn’t have to be unique. 
If not specified at creation, the cluster name will be an empty string.

#### - `spark_version`:
> **(Optional)** 	The Spark version of the cluster. A list of available 
Spark versions can be retrieved by using the Runtime Versions API call. This field is required.

#### - `spark_conf`:
> **(Optional)** An object containing a set of optional, user-specified Spark 
configuration key-value pairs. You can also pass in a string of extra JVM options to the driver and the executors via 
spark.driver.extraJavaOptions and spark.executor.extraJavaOptions respectively.

#### - `aws_attributes` **(Optional)** :
> {{%chevron default="Attributes related to clusters running on Amazon Web Services. If not specified at cluster creation, a set of default values will be used." display="true" %}}

#### **Usage**

 ```hcl
resource "databricks_cluster" "my-cluster" {
...
    aws_attributes {
      zone_id = "us-east-1"
      availability = "SPOT"
      spot_bid_price_percent = 100
      instance_profile_arn = "arn:aws:iam::999999999999:instance-profile/custom-s3-access-instance-profile"
      first_on_demand = 1
      ebs_volume_type = "GENERAL_PURPOSE_SSD"
      ebs_volume_count = 1
      ebs_volume_size = 32
    }
...
}
```

* `zone_id` - **(Required)** Identifier for the availability zone/datacenter in which the cluster resides. 
This string will be of a form like “us-west-2a”. The provided availability zone must be in the same region as the 
Databricks deployment. For example, “us-west-2a” is not a valid zone ID if the Databricks deployment resides in the 
“us-east-1” region. 

* `availability` - **(Optional)** 	Availability type used for all subsequent nodes past the first_on_demand 
ones. Note: If first_on_demand is zero, this availability type will be used for the entire cluster.

* `spot_bid_price_percent` - **(Optional)** The max price for AWS spot instances, as a percentage of the 
corresponding instance type’s on-demand price. For example, if this field is set to 50, and the cluster needs a new 
i3.xlarge spot instance, then the max price is half of the price of on-demand i3.xlarge instances. Similarly, if 
this field is set to 200, the max price is twice the price of on-demand i3.xlarge instances. If not specified, the 
default value is 100. When spot instances are requested for this cluster, only spot instances whose max price 
percentage matches this field will be considered. For safety, we enforce this field to be no more than 10000.

* `instance_profile_arn` - **(Optional)** Nodes for this cluster will only be placed on AWS instances with 
this instance profile. If omitted, nodes will be placed on instances without an instance profile. The instance 
profile must have previously been added to the Databricks environment by an account administrator.

* `first_on_demand` - **(Optional)** The first first_on_demand nodes of the cluster will be placed on on-demand 
instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this 
value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances. If this 
value is less than the current cluster size, first_on_demand nodes will be placed on on-demand instances and the 
remainder will be placed on availability instances. This value does not affect cluster size and cannot be mutated 
over the lifetime of a cluster.

* `ebs_volume_type` - **(Optional)** The type of EBS volumes that will be launched with this cluster. 
GENERAL_PURPOSE_SSD or THROUGHPUT_OPTIMIZED_HDD

* `ebs_volume_count` - **(Optional)** The number of volumes launched for each instance. You can choose up to 
10 volumes. This feature is only enabled for supported node types. Legacy node types cannot specify custom EBS 
volumes. For node types with no instance store, at least one EBS volume needs to be specified; otherwise, cluster 
creation will fail. These EBS volumes will be mounted at /ebs0, /ebs1, and etc. Instance store volumes will 
be mounted at /local_disk0, /local_disk1, and etc. If EBS volumes are attached, Databricks will configure Spark 
to use only the EBS volumes for scratch storage because heterogeneously sized scratch devices can lead to 
inefficient disk utilization. If no EBS volumes are attached, Databricks will configure Spark to use instance 
store volumes. If EBS volumes are specified, then the Spark configuration spark.local.dir will be overridden.

* `ebs_volume_size` - **(Optional)** The size of each EBS volume (in GiB) launched for each instance. 
For general purpose SSD, this value must be within the range 100 - 4096. For throughput optimized HDD, this 
value must be within the range 500 - 4096. Custom EBS volumes cannot be specified for the legacy node types
(memory-optimized and compute-optimized).

{{% /chevron %}}

#### - `driver_node_type_id`:
> **(Optional)** The node type of the Spark driver. This field 
is optional; if unset, the driver node type will be set as the same value as node_type_id defined above.

#### - `node_type_id`:
> **(Optional)** This field encodes, through a single value, the resources 
available to each of the Spark nodes in this cluster. For example, the Spark nodes can be provisioned and optimized for 
memory or compute intensive workloads A list of available node types can be retrieved by using the List Node Types API 
call. This field is required.

#### - `ssh_public_keys`:
> **(Optional)** SSH public key contents that will be added to each 
Spark node in this cluster. The corresponding private keys can be used to login with the user name ubuntu on port 2200. 
Up to 10 keys can be specified.

#### - `custom_tags`:
> **(Optional)** Additional tags for cluster resources. Databricks will tag all 
cluster resources (e.g., AWS instances and EBS volumes) with these tags in addition to default_tags.


#### - `cluster_log_conf` **(Optional)**  :
> #### **Usage**
> {{< tabs groupId="jobStorageConfig" >}}
 {{% tab name="DBFS" %}}
 ```hcl
cluster_log_conf {
  dbfs_destination = "dbfs:/my/path/in/dbfs"
}
 ```
 {{% /tab %}}
 {{% tab name="S3" %}}
```hcl
cluster_log_conf {
  s3_destination = "dbfs:/my/path/in/dbfs"
  s3_region = "us-east-1"
  s3_endpoint = "https://s3-us-east-1.amazonaws.com."
  enable_encryption = true
  encryption_type = "sse-kms"
  kms_key = "my-kms-key-here"
  canned_acl = "bucket-owner-full-control"
}
```
{{% /tab %}}
{{< /tabs >}}
>
> {{% chevron default="The configuration for delivering Spark logs to a long-term storage destination. Only one destination can be specified for one cluster. If the conf is given, the logs will be delivered to the destination every 5 mins. The destination of driver logs is <destination>/<cluster-id>/driver, while the destination of executor logs is <destination>/<cluster-id>/executor." display="true" %}}


* `dbfs_destination` - **(Optional)** DBFS location of cluster log. destination must be provided. For example, 
"dbfs:/home/cluster_log"

* `s3_destination` - **(Optional)** S3 destination, e.g. s3://my-bucket/some-prefix You must configure the 
cluster with an instance profile and the instance profile must have write access to the destination. You cannot use 
AWS keys.

* `s3_region` - **(Optional)** S3 region, e.g. us-west-2. Either region or endpoint must be set. If both are 
set, endpoint is used.

* `s3_endpoint` - **(Optional)** S3 endpoint, e.g. https://s3-us-west-2.amazonaws.com. Either region or endpoint 
needs to be set. If both are set, endpoint is used.

* `s3_enable_encryption` - **(Optional)** Enable server side encryption, false by default.

* `s3_encryption_type` - **(Optional)** The encryption type, it could be sse-s3 or sse-kms. 
It is used only when encryption is enabled and the default type is sse-s3.

* `s3_kms_key` - **(Optional)** KMS key used if encryption is enabled and encryption type is set to sse-kms.

* `s3_canned_acl` - **(Optional)** Set canned access control list, e.g. bucket-owner-full-control. 
If canned_cal is set, the cluster instance profile must have s3:PutObjectAcl permission on the destination bucket 
and prefix. The full list of possible canned ACL can be found [here](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl). 
By default only the object owner gets full control. If you are using cross account role for writing data, you may 
want to set bucket-owner-full-control to make bucket owner able to read the logs.

{{% /chevron %}}

#### - `init_scripts` **(Optional)** :
> #### **Usage**
> {{< tabs groupId="jobStorageConfig" >}}
 {{% tab name="DBFS" %}}
 ```hcl
cluster_log_conf {
  dbfs_destination = "dbfs:/my/path/in/dbfs"
}
 ```
 {{% /tab %}}
 {{% tab name="S3" %}}
```hcl
init_scripts {
  s3_destination = "dbfs:/my/path/in/dbfs"
  s3_region = "us-east-1"
  s3_endpoint = "https://s3-us-east-1.amazonaws.com."
}
```
{{% /tab %}}
{{< /tabs >}}
>
> {{%chevron default="The configuration for storing init scripts. Any number of scripts can be specified. The scripts are executed sequentially in the order provided. If cluster_log_conf is specified, init script logs are sent to <destination>/<cluster-id>/init_scripts." display="true" %}}

* `dbfs_destination` - **(Optional)** DBFS location of init script. Destination must be provided. For example, 
"dbfs:/home/cluster_log"

* `s3_destination` - **(Optional)** S3 destination, e.g. s3://my-bucket/some-prefix You must configure the 
cluster with an instance profile and the instance profile must have write access to the destination. You cannot use 
AWS keys.

* `s3_region` - **(Optional)** S3 region, e.g. us-west-2. Either region or endpoint must be set. If both are 
set, endpoint is used.

* `s3_endpoint` - **(Optional)** S3 endpoint, e.g. https://s3-us-west-2.amazonaws.com. Either region or endpoint 
needs to be set. If both are set, endpoint is used.

{{% /chevron %}}

#### - `spark_env_vars`:
> **(Optional)** An object containing a set of optional, user-specified 
environment variable key-value pairs. Key-value pair of the form (X,Y) are exported as is (i.e., export X='Y') while 
launching the driver and workers. To specify an additional set of SPARK_DAEMON_JAVA_OPTS, we recommend appending them 
to $SPARK_DAEMON_JAVA_OPTS as shown in the example below. This ensures that all default databricks managed environmental 
variables are included as well. 

#### - `enable_elastic_disk`:
> **(Optional)** Autoscaling Local Storage: when enabled, 
this cluster dynamically acquires additional disk space when its Spark workers are running low on disk space. This 
feature requires specific AWS permissions to function correctly - refer to Autoscaling local storage for details.

#### - `instance_pool_id`:
> **(Optional)** The optional ID of the instance pool to which the 
cluster belongs. Refer to Instance Pools API for details.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the job object.

#### - `job_id`:
> The canonical identifier for the newly created job.

#### - `creator_user_name`:
> The creator user name. This field won’t be included in 
the response if the user has already been deleted.

#### - `created_time`:
> The time at which this job was created in epoch milliseconds 
(milliseconds since 1/1/1970 UTC).


## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
