# databricks_aws_s3_mount Resource

This resource will mount your S3 bucket on `dbfs:/mnt/yourname`. It is important to understand that this will start up the cluster if the cluster is terminated. The read and refresh terraform command will require a cluster and make take some time to validate mount. If cluster_id is not specified, it will create the smallest possible cluster called `terraform-mount` for shortest possible amount of time. 

## Example Usage

```hcl
// will create AWS S3 bucket
resource "aws_s3_bucket" "this" {
  bucket = "${var.prefix}-rootbucket"
  acl    = "private"
  versioning {
    enabled = false
  }
}

// now you can do `%fs ls /mnt/experiments` in notebooks
resource "databricks_s3_mount" "this" {
    s3_bucket_name = aws_s3_bucket.this.bucket
    mount_name = "experiments"
}
```

## Argument Reference

The following arguments are required:

* `cluster_id` - (Optional) (String) Cluster to use for mounting. If no cluster is specified, new cluster will be created and will mount the bucket for all of the clusters in this workspace. If cluster is specified, mount will be visible for all clusters with the same [instance profile](./instance_profile.md). If cluster is not running - it's going to be started, so be aware to set autotermination rules on it.
* `mount_name` - (Required) (String) Name, under which mount will be accessible in `dbfs:/mnt/<MOUNT_NAME>`.
* `s3_bucket_name` - (Required) (String) S3 bucket name to be mounted.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - mount name
* `source` - (String) HDFS-compatible S3 bucket url `s3a://<s3_bucket_name>` 


## Import

The resource aws s3 mount can be imported using it's mount name

```bash
$ terraform import databricks_aws_s3_mount.this <mount_name>
```