+++
title = "cluster"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_cluster`

This resource allows you to create, update, and delete clusters.  

## Example Usage

```hcl
resource "databricks_cluster" "my-cluster" {
  num_workers = "2"
  cluster_name = "sri-test"
  spark_version = "6.4.x-scala2.11"
  node_type_id = "i3.2xlarge"
  autotermination_minutes = 30
  aws_attributes {
    availability = "ON_DEMAND"
    zone_id = "us-east-1"
    instance_profile_arn = "arn:aws:iam::999999999999:instance-profile/s3-access"
  }
}
```  

>Note: For Azure, valid node_type_ids are the "Sizes" of the virtual machines to use as workers, e.g. Standard_DS3_v2.

## Argument Reference

The following arguments are supported:

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
> **(Required)** 	The Spark version of the cluster. A list of available 
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
> **(Optional - required if instance_pool_id is not given)** This field encodes, through a single value, the resources 
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
> {{< tabs groupId="storageConfig" >}}
 {{% tab name="DBFS" %}}
 ```hcl
cluster_log_conf {
  dbfs {
    destination = "dbfs:/my/path/in/dbfs"
  } 
}
 ```
 {{% /tab %}}
 {{% tab name="S3" %}}
```hcl
cluster_log_conf {
  s3 {
    destination = "s3:/my/path/in/dbfs"
    region = "us-east-1"
    endpoint = "https://s3-us-east-1.amazonaws.com"
    enable_encryption = true
    encryption_type = "sse-kms"
    kms_key = "my-kms-key-here"
    canned_acl = "bucket-owner-full-control" 
  }
}
```
{{% /tab %}}
{{< /tabs >}}
>
> {{% chevron default="The configuration for delivering Spark logs to a long-term storage destination. Only one destination can be specified for one cluster. If the conf is given, the logs will be delivered to the destination every 5 mins. The destination of driver logs is <destination>/<cluster-id>/driver, while the destination of executor logs is <destination>/<cluster-id>/executor." display="true" %}}


* `dbfs` - Configuration for the dbfs cluster logs configuration 

    * `destination` - **(Optional)** DBFS location of cluster log. destination must be provided. For example, 
    "dbfs:/home/cluster_log"

* `s3` - Configuration for the s3 cluster logs configuration
 
    * `destination` - **(Optional)** S3 destination, e.g. s3://my-bucket/some-prefix You must configure the 
    cluster with an instance profile and the instance profile must have write access to the destination. You cannot use 
    AWS keys.
    
    * `region` - **(Optional)** S3 region, e.g. us-west-2. Either region or endpoint must be set. If both are 
    set, endpoint is used.
    
    * `endpoint` - **(Optional)** S3 endpoint, e.g. https://s3-us-west-2.amazonaws.com. Either region or endpoint 
    needs to be set. If both are set, endpoint is used.
    
    * `enable_encryption` - **(Optional)** Enable server side encryption, false by default.
    
    * `encryption_type` - **(Optional)** The encryption type, it could be sse-s3 or sse-kms. 
    It is used only when encryption is enabled and the default type is sse-s3.
    
    * `kms_key` - **(Optional)** KMS key used if encryption is enabled and encryption type is set to sse-kms.
    
    * `canned_acl` - **(Optional)** Set canned access control list, e.g. bucket-owner-full-control. 
    If canned_cal is set, the cluster instance profile must have s3:PutObjectAcl permission on the destination bucket 
    and prefix. The full list of possible canned ACL can be found [here](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl). 
    By default only the object owner gets full control. If you are using cross account role for writing data, you may 
    want to set bucket-owner-full-control to make bucket owner able to read the logs.

{{% /chevron %}}

#### - `init_scripts` **(Optional)** :
> #### **Usage**
> {{< tabs groupId="storageConfig" >}}
 {{% tab name="DBFS" %}}
 ```hcl
init_scripts {
  dbfs {
    destination = "dbfs:/my/path/in/dbfs"
  }
}
 ```
 {{% /tab %}}
 {{% tab name="S3" %}}
```hcl
init_scripts {
  s3 {
    destination = "dbfs:/my/path/in/dbfs"
    region = "us-east-1"
    endpoint = "https://s3-us-east-1.amazonaws.com."
  }
}
```
{{% /tab %}}
{{< /tabs >}}
>
> {{%chevron default="The configuration for storing init scripts. Any number of scripts can be specified. The scripts are executed sequentially in the order provided. If cluster_log_conf is specified, init script logs are sent to <destination>/<cluster-id>/init_scripts." display="true" %}}

* `dbfs` - Configuration for the init scripts configuration

    * `destination` - **(Optional)** DBFS location of init script. Destination must be provided. For example, 
    "dbfs:/home/cluster_log"

* `s3` - Configuration for the s3 init scripts configuration

    * `destination` - **(Optional)** S3 destination, e.g. s3://my-bucket/some-prefix You must configure the 
    cluster with an instance profile and the instance profile must have write access to the destination. You cannot use 
    AWS keys.
    
    * `region` - **(Optional)** S3 region, e.g. us-west-2. Either region or endpoint must be set. If both are 
    set, endpoint is used.
    
    * `endpoint` - **(Optional)** S3 endpoint, e.g. https://s3-us-west-2.amazonaws.com. Either region or endpoint 
    needs to be set. If both are set, endpoint is used.
    
    * `enable_encryption` - **(Optional)** Enable server side encryption, false by default.
    
    * `encryption_type` - **(Optional)** The encryption type, it could be sse-s3 or sse-kms. 
    It is used only when encryption is enabled and the default type is sse-s3.
    
    * `kms_key` - **(Optional)** KMS key used if encryption is enabled and encryption type is set to sse-kms.
    
    * `canned_acl` - **(Optional)** Set canned access control list, e.g. bucket-owner-full-control. 
    If canned_cal is set, the cluster instance profile must have s3:PutObjectAcl permission on the destination bucket 
    and prefix. The full list of possible canned ACL can be found [here](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl). 
    By default only the object owner gets full control. If you are using cross account role for writing data, you may 
    want to set bucket-owner-full-control to make bucket owner able to read the logs.

{{% /chevron %}}

#### - `docker_image` **(Optional)** :
> #### **Usage**
>
>```hcl
>docker_image {
>    url = "https://hub.docker.com/_/alpine"
>    username = "my-user-name"
>    password = "password"
>}
>```
>
> {{%chevron default="Docker image for a custom container." display="true" %}}

* `url` - **(Required)** URL for the Docker image.

* `username` - **(Required)** User name for the Docker repository.

* `password` - **(Required)** 	Password for the Docker repository. (Sensitive field)

{{% /chevron %}}

#### - `spark_env_vars`:
> **(Optional)** An object containing a set of optional, user-specified 
environment variable key-value pairs. Key-value pair of the form (X,Y) are exported as is (i.e., export X='Y') while 
launching the driver and workers. To specify an additional set of SPARK_DAEMON_JAVA_OPTS, we recommend appending them 
to $SPARK_DAEMON_JAVA_OPTS as shown in the example below. This ensures that all default databricks managed environmental 
variables are included as well. 

#### - `autotermination_minutes`:
> **(Optional)** Automatically terminates the 
cluster after it is inactive for this time in minutes. If not set, this cluster will not be automatically terminated. 
If specified, the threshold must be between 10 and 10000 minutes. You can also set this value to 0 to explicitly 
disable automatic termination.

#### - `enable_elastic_disk`:
> **(Optional)** Autoscaling Local Storage: when enabled, 
this cluster dynamically acquires additional disk space when its Spark workers are running low on disk space. This 
feature requires specific AWS permissions to function correctly - refer to Autoscaling local storage for details.

#### - `instance_pool_id`:
> **(Optional - required if node_type_id is not given)** The optional ID of the instance pool to which the 
cluster belongs. Refer to Instance Pools API for details.

#### - `single_user_name`:
> **(Optional)** The optional user name of the user to assign to an interactive cluster. This is required when using standard AAD Passthrough for Azure Datalake Storage (ADLS) with a single-user cluster (i.e. not high-concurrency clusters).

#### - `idempotency_token`:
> **(Optional)** An optional token that can be 
used to guarantee the idempotency of cluster creation requests. If an active cluster with the provided token already 
exists, the request will not create a new cluster, but it will return the ID of the existing cluster instead. The 
existence of a cluster with the same token is not checked against terminated clusters. If you specify the idempotency 
token, upon failure you can retry until the request succeeds. Databricks will guarantee that exactly one cluster will 
be launched with that idempotency token. This token should have at most 64 characters.

### Libraries

Libraries are set objects within the Cluster resources, examples below.

{{< tabs groupId="libraries" >}}
{{% tab name="Jar" %}}
 ```hcl
library_jar {
     path = "dbfs:/my/path/in/dbfs/jar"
}
 ```
{{% /tab %}}
{{% tab name="Egg" %}}
```hcl
library_egg {
    path = "dbfs:/my/path/in/dbfs/egg"
}
```
{{% /tab %}}
{{% tab name="Whl" %}}
```hcl
library_whl {
    path = "dbfs:/my/path/in/dbfs/whl"
}
```
{{% /tab %}}
{{% tab name="PyPi" %}}
```hcl
library_pypi {
    package = "networkx"
    repo = "https://pypi.org"
}
```
{{% /tab %}}
{{% tab name="Maven" %}}
```hcl
library_maven {
    coordinates = "org.jsoup:jsoup:1.7.2"
    repo = "https://mavencentral.org"
    exclusions = ["slf4j:slf4j"]
}
```
{{% /tab %}}
{{% tab name="Cran" %}}
```hcl
library_cran {
    package = "ada"
    repo = "https://cran.us.r-project.org"
}
```
{{% /tab %}}
{{< /tabs >}}

#### - `library_jar` **(Optional)** :
> {{%chevron default="URI of the JAR to be installed. DBFS and S3 URIs are supported. For example: 'dbfs:/mnt/databricks/library.jar', 's3://my-bucket/library.jar'. If S3 is used, make sure the cluster has read access on the library. You may need to launch the cluster with an instance profile to access the S3 URI." display="true" %}}     

* `path` - **(Required)** Path of the jar in dbfs or in S3. For example: "dbfs:/mnt/databricks/library.jar", "s3://my-bucket/library.jar". 

* `messages` - **(Required)** Messages of the results of the library installation

* `status` - **(Computed)** The status of the library installation. Possible statuses are: PENDING, RESOLVING, 
INSTALLING, INSTALLED, FAILED, and UNINSTALL_ON_RESTART.
       
{{% /chevron %}}       

#### - `library_egg` **(Optional)** :
> {{%chevron default="URI of the egg to be installed. DBFS and S3 URIs are supported. For example: 'dbfs:/my/egg', 's3://my-bucket/egg'. If S3 is used, make sure the cluster has read access on the library. You may need to launch the cluster with an instance profile to access the S3 URI." display="true" %}}

* `path` - **(Required)** Path of the egg in dbfs or in S3. For example: "dbfs:/mnt/databricks/library.egg", 
"s3://my-bucket/library.egg". 

* `messages` - **(Required)** Messages of the results of the library installation

* `status` - **(Computed)** The status of the library installation. Possible statuses are: PENDING, RESOLVING, 
INSTALLING, INSTALLED, FAILED, and UNINSTALL_ON_RESTART.
    
{{% /chevron %}}

#### - `library_whl` **(Optional)** :

> {{%chevron default="If whl, URI of the wheel or zipped wheels to be installed. DBFS and S3 URIs are supported. For example: 'dbfs:/my/whl', 's3://my-bucket/whl'. If S3 is used, make sure the cluster has read access on the library. You may need to launch the cluster with an instance profile to access the S3 URI. Also the wheel file name needs to use the correct convention. If zipped wheels are to be installed, the file name suffix should be .wheelhouse.zip." display="true" %}}

* `path` - **(Required)** Path of the whl in dbfs or in S3. For example: "dbfs:/mnt/databricks/library.whl", 
"s3://my-bucket/library.whl". 

* `messages` - **(Required)** Messages of the results of the library installation

* `status` - **(Computed)** The status of the library installation. Possible statuses are: PENDING, RESOLVING, 
INSTALLING, INSTALLED, FAILED, and UNINSTALL_ON_RESTART.

{{% /chevron %}}

#### - `library_pypi` **(Optional)** :

> {{%chevron default="Specification of a PyPI library to be installed." display="true" %}}

* `package` - **(Required)**  The name of the PyPI package to install. An optional exact version specification 
is also supported. Examples: simplejson and simplejson==3.8.0. This field is required.

* `repo` - **(Optional)** The repository where the package can be found. If not specified, 
the default pip index is used.

* `messages` - **(Computed)** Messages of the results of the library installation
    
* `status` - **(Computed)** The status of the library installation. Possible statuses are: PENDING, RESOLVING, 
INSTALLING, INSTALLED, FAILED, and UNINSTALL_ON_RESTART.
    
{{% /chevron %}}

#### - `library_maven` **(Optional)** :

> {{%chevron default="Specification of a Maven library to be installed." display="true" %}}

* `coordinates` - **(Required)** Gradle-style Maven coordinates. For example: org.jsoup:jsoup:1.7.2. 
This field is required.

* `repo` - **(Optional)** Maven repo to install the Maven package from. If omitted, both Maven Central 
Repository and Spark Packages are searched.

* `exclusions` - **(Optional)** List of dependences to exclude. For example: 
("slf4j:slf4j", "*:hadoop-client"). 

* `messages` - **(Computed)** Messages of the results of the library installation
    
* `status` - **(Computed)** The status of the library installation. Possible statuses are: PENDING, RESOLVING, 
INSTALLING, INSTALLED, FAILED, and UNINSTALL_ON_RESTART.
    
{{% /chevron %}}
    
#### - `library_cran` **(Optional)** :

> {{%chevron default="Specification of a CRAN library to be installed." display="true" %}}

* `package` - **(Required)** The name of the CRAN package to install. This field is required.

* `repo` - **(Optional)** The repository where the package can be found. If not specified, 
the default CRAN repo is used.

* `messages` - **(Computed)** Messages of the results of the library installation
    
* `status` - **(Computed)** The status of the library installation. Possible statuses are: PENDING, RESOLVING, 
INSTALLING, INSTALLED, FAILED, and UNINSTALL_ON_RESTART.
    
{{% /chevron %}}

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the cluster object.

#### - `cluster_id`:
> Canonical identifier for the cluster.

#### - `default_tags`:
> Tags that are added by Databricks by default, regardless of any 
custom_tags that may have been added. These include: Vendor: Databricks, Creator: <username_of_creator>, ClusterName: 
<name_of_cluster>, ClusterId: <id_of_cluster>, Name: <Databricks internal use>

#### - `state`:
>	State of the cluster.

#### - `state_message`:
> A message associated with the most recent state transition 
(e.g., the reason why the cluster entered a TERMINATED state). This field is unstructured, and its exact format 
is subject to change.


## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
