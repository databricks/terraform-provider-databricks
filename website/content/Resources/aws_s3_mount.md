+++
title = "aws_s3_mount"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++


## Resource: `databricks_aws_s3_mount`

This resource given a cluster id will help you create, get and delete a aws s3 mount.

{{% notice warning %}}
It is important to understand that this will start up the cluster if the cluster is terminated.
The read and refresh terraform command will require a cluster and make take some time to validate mount.
{{% /notice %}}

{{% notice note %}}
You can locate the mount in dbfs at `dbfs:/mnt/<mount_name>`
{{% /notice %}}


{{% notice note %}}
Currently only supports the use of IAM roles and not S3 keys and secrets.
{{% /notice %}}


## Example Usage

````hcl
resource "databricks_aws_s3_mount" "my_custom_mount4" {
  cluster_id = "####-######-hello###"
  s3_bucket_name = "my-s3-bucket-123"
  mount_name = "my_s3_bucket_mount"
}
````
## Argument Reference

The following arguments are supported:

#### - `cluster_id`:
> **(Required)** This is the cluster id in which the mount will be initalized from. 
>If the cluster is in a terminated state it will be started.

#### - `s3_bucket_name`:
> **(Required)** This is the S3 bucket that you are trying to 
mount.

#### - `mount_name`:
> **(Required)** This is the name of the mount that will represent 
where the data will be landed. 


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> Identifier for the mount.


## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
