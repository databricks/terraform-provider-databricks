+++
title = "instance_pool"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_instance_pool`

This resource allows you to manage instance pools on Databricks. 

An instance pool reduces cluster start and auto-scaling times by maintaining a set of idle, ready-to-use cloud instances. 
When a cluster attached to a pool needs an instance, it first attempts to allocate one of the pool’s idle instances. 
If the pool has no idle instances, it expands by allocating a new instance from the instance provider in order to 
accommodate the cluster’s request. When a cluster releases an instance, it returns to the pool and is free for another 
cluster to use. Only clusters attached to a pool can use that pool’s idle instances.


## Example Usage

.. Note:: It is important to know what that different cloud service providers have different `node_type_id`,
 `disk_specs` and potentially other configurations.  

### Instance Pool Example

{{< tabs groupId="instancePoolCloud" >}}
 {{% tab name="AWS" %}}
```hcl
resource "databricks_instance_pool" "my-pool" {
  instance_pool_name = "demo-terraform-pool"
  min_idle_instances = 0
  max_capacity = 5
  node_type_id = "i3.xlarge"
  aws_attributes {
    availability = "ON_DEMAND"
    zone_id = "us-east-1a"
    spot_bid_price_percent = "100"
  }
  idle_instance_autotermination_minutes = 10
  disk_spec {
    ebs_volume_type = "GENERAL_PURPOSE_SSD"
    disk_size = 80
    disk_count = 1
  }
  custom_tags = {
    "creator" = "Sriharsha Tikkireddy"
    "testChange" = "Sri Tikkireddy"
  }
}
```
 {{% /tab %}}
 {{% tab name="Azure" %}}
```hcl
resource "databricks_instance_pool" "my-pool" {
  instance_pool_name = "demo-terraform-pool"
  min_idle_instances = 0
  max_capacity = 5
  node_type_id = "Standard_DS3_v2"
  idle_instance_autotermination_minutes = 10
  disk_spec {
    azure_disk_volume_type = "PREMIUM_LRS"
    disk_size = 80
    disk_count = 1
  }
  custom_tags = {
    "creator" = "demo user"
    "testChange" = "demo user"
  }
}
```
{{% /tab %}}
{{< /tabs >}}

## Argument Reference

The following arguments are supported:

#### - `instance_pool_name`:
> **(Required)** The name of the instance pool. This 
is required for create and edit operations. It must be unique, non-empty, and less than 100 characters.

#### - `min_idle_instances`:
> **(Required)** 	The minimum number of idle 
instances maintained by the pool. This is in addition to any instances in use by active clusters.

#### - `max_capacity`:
> **(Required)** The maximum number of instances the pool can 
contain, including both idle instances and ones in use by clusters. Once the maximum capacity is reached, you cannot 
create new clusters from the pool and existing clusters cannot autoscale up until some instances are made idle in the 
pool via cluster termination or down-scaling.

#### - `idle_instance_autotermination_minutes`:
> **(Required)** 
The number of minutes that idle instances in excess of the min_idle_instances are maintained by the pool before being 
terminated. If not specified, excess idle instances are terminated automatically after a default timeout period. If 
specified, the time must be between 0 and 10000 minutes. If you specify 0, excess idle instances are removed as soon 
as possible.

#### - `aws_attributes` **(Optional)** :
> {{% chevron default=` Attributes related to instance pools running 
on Amazon Web Services. If not specified at creation time, a set of default values is used.  This block contains the 
following attributes of availability, zone_id, spot_bid_price_percent:` display="true" %}}

* `availability` - **(Optional)** Availability type used for all instances in the pool. Only `"ON_DEMAND"` and 
`"SPOT"` are supported.
* `zone_id` - **(Optional)** Identifier for the availability zone/datacenter in which the instance pool resides. 
This string is of a form like `"us-west-2a"`. The provided availability zone must be in the same region as the 
Databricks deployment. For example, `"us-west-2a"` is not a valid zone ID if the Databricks deployment resides 
in the `"us-east-1"` region. This is an optional field. If not specified, a default zone is used. 
You can find the list of available zones as well as the default value by using the 
[List Zones API](https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterservicelistavailablezones).
* `spot_bid_price_percent` - **(Optional)** The max price for AWS spot instances, as a percentage of the corresponding 
instance type’s on-demand price. For example, if this field is set to 50, and the instance pool needs a new i3.xlarge 
spot instance, then the max price is half of the price of on-demand i3.xlarge instances. Similarly, if this field 
is set to 200, the max price is twice the price of on-demand i3.xlarge instances. If not specified, the **default 
value is 100**. When spot instances are requested for this instance pool, only spot instances whose max price 
percentage matches this field are considered. *For safety, this field cannot be greater than 10000.*

.. Important:: **aws_attributes** will only work for instance pools in an AWS deployment of Databricks. They will **not** work 
    on Azure Databricks! 
{{% /chevron %}}

#### - `node_type_id`:
> **(Required)** The node type for the instances in the pool. All 
clusters attached to the pool inherit this node type and the pool’s idle instances are allocated based on this type. 
You can retrieve a list of available node types by using the 
[List Node Types API](https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterservicelistnodetypes) call.

#### - `custom_tags`:
> **(Optional)** Additional tags for instance pool resources. Databricks 
tags all pool resources (e.g. AWS & Azure instances and Disk volumes) with these tags in addition to default_tags. 
**Databricks allows at most 43 custom tags.**

#### - `enable_elastic_disk`:
> **(Optional)** Autoscaling Local Storage: when 
enabled, the instances in the pool dynamically acquire additional disk space when they are running low on disk space.

#### - `disk_spec` **(Optional)** :
> {{% notice note %}}
  For disk_spec make sure to use **ebs_volume_type** only on AWS deployment of Databricks and 
      **azure_disk_volume_type** only on a Azure deployment of Databricks.
  {{% /notice %}}
> {{% chevron default=`Defines the amount of initial remote storage attached 
to each instance in the pool. This block contains the following attributes of ebs_volume_type, 
azure_disk_volume_type, disk_count, disk_size:` display="true" %}}

* `ebs_volume_type` - **(Optional)** The EBS volume type to use. 
Options are: `"GENERAL_PURPOSE_SSD"` (Provision extra storage using AWS gp2 EBS volumes) or 
`"THROUGHPUT_OPTIMIZED_HDD"` (Provision extra storage using AWS st1 volumes.)

* `azure_disk_volume_type` - **(Optional)** The type of Azure disk to use. 
Options are: `"PREMIUM_LRS"` (Premium storage tier, backed by SSDs) 
or `"STANDARD_LRS"` (Standard storage tier, backed by HDDs.)

* `disk_count` - **(Optional)** The number of disks to attach to each instance:
    * This feature is only enabled for supported node types.
    * Users can choose up to the limit of the disks supported by the node type.
    * For node types with no local disk, at least one disk needs to be specified.

* `disk_size` - **(Optional)** The size of each disk (in GiB) to attach. Values must fall into the supported range 
for a particular instance type:
    * **AWS (ebs)**:
        * General Purpose SSD: `100 - 4096` GiB
        * Throughput Optimized HDD: `500 - 4096` GiB
    * **Azure (disk volume)**:
        * Premium LRS (SSD): `1 - 1023` GiB
        * Standard LRS (HDD): `1- 1023` GiB
    
{{% /chevron %}}

#### - `preloaded_spark_versions`:
> **(Optional)** 	A list with the 
runtime version the pool installs on each instance. Pool clusters that use a preloaded runtime version start faster 
as they do have to wait for the image to download. You can retrieve a list of available runtime versions by using the 
[Runtime Versions API](https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterservicelistsparkversions) 
call.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
* :ref:`id <r_instance_pool_id>` - The id for the instance pool object.

#### - `default_tags`:
> Tags that are added by Databricks regardless of any 
custom_tags, including:
>* **Vendor**: Databricks
>* **DatabricksInstancePoolCreatorId**: <create_user_id>
>* **DatabricksInstancePoolId**: <instance_pool_id>

#### - `state`:
> Current state of the instance pool.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
