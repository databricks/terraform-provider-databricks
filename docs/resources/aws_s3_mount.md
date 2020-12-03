# databricks_aws_s3_mount Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

This resource will mount your S3 bucket on `dbfs:/mnt/yourname`. It is important to understand that this will start up the [cluster](cluster.md) if the cluster is terminated. The read and refresh terraform command will require a cluster and may take some time to validate the mount. If cluster_id is not specified, it will create the smallest possible cluster called `terraform-mount` for the shortest possible amount of time.

## Example Usage

Simple resource usage:

```hcl
// now you can do `%fs ls /mnt/experiments` in notebooks
resource "databricks_s3_mount" "this" {
    instance_profile = databricks_instance_profile.ds.id
    s3_bucket_name = aws_s3_bucket.this.bucket
    mount_name = "experiments"
}
```

Full end-to-end actions required to securely mount S3 bucket on all clusters with the same [instance profile](instance_profile.md):

```hcl
// Step 1: Create bucket
resource "aws_s3_bucket" "ds" {
  bucket = "${var.prefix}-ds"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(var.tags, {
    Name = "${var.prefix}-ds"
  })
}

// Step 2: Create IAM role for data access
resource "aws_iam_role" "data_role" {
  name               = "${var.prefix}-first-ec2s3"
  description        = "(${var.prefix}) EC2 Assume Role role for S3 access"
  assume_role_policy = data.aws_iam_policy_document.assume_role_for_ec2.json
  tags               = var.tags
}

// Step 3: Let it assume roles on EC2
data "aws_iam_policy_document" "assume_role_for_ec2" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      identifiers = ["ec2.amazonaws.com"]
      type        = "Service"
    }
  }
}

// Step 4: Create bucket policy that will give full access to this bucket
data "databricks_aws_bucket_policy" "ds" {
  provider         = databricks.mws
  full_access_role = aws_iam_role.data_role.arn
  bucket           = aws_s3_bucket.ds.bucket
}

// Step 5: Apply inline S3 bucket policy to newly created bucket
resource "aws_s3_bucket_policy" "ds" {
  bucket = aws_s3_bucket.ds.id
  policy = data.databricks_aws_bucket_policy.ds.json
}

// Step 6: Create cross-account policy, which allows Databricks to pass given list of data roles
data "databricks_aws_crossaccount_policy" "this" {
  pass_roles = [aws_iam_role.data_role.arn]
}

// Step 7: Apply created IAM inline policy
resource "aws_iam_policy" "cross_account_policy" {
  name   = "${var.prefix}-crossaccount-iam-policy"
  policy = data.databricks_aws_crossaccount_policy.this.json
}

// Step 8: Allow Databricks to perform actions within your account, given requests are with ExternalId you've received on the website.
data "databricks_aws_assume_role_policy" "this" {
    external_id = var.account_id
}

// Step 9: Grant Databricks full access to VPC resources
resource "aws_iam_role" "cross_account" {
  name               = "${var.prefix}-crossaccount-iam-role"
  assume_role_policy = data.databricks_aws_assume_role_policy.this.json
  description        = "Grants Databricks full access to VPC resources"
}

// Step 10: Attach cross-account policy to cross-account role
resource "aws_iam_role_policy_attachment" "cross_account" {
  policy_arn = aws_iam_policy.cross_account_policy.arn
  role       = aws_iam_role.cross_account.name
}

// Step 11: Register cross-account role for multi-workspace scenario (only if you're using multi-workspace setup)
resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = var.account_id
  credentials_name = "${var.prefix}-creds"
  role_arn         = aws_iam_role.cross_account.arn
}

// Step 12: Register your data role with instance profile
resource "aws_iam_instance_profile" "this" {
  name = "${var.prefix}-first-profile"
  role = aws_iam_role.data_role.name
}

// Step 13: Register instance profile at Databricks
resource "databricks_instance_profile" "ds" {
  instance_profile_arn = aws_iam_instance_profile.this.arn
  skip_validation      = false
}

// Step 14: now you can do `%fs ls /mnt/experiments` in notebooks
resource "databricks_s3_mount" "this" {
    instance_profile = databricks_instance_profile.ds.id
    s3_bucket_name = aws_s3_bucket.this.bucket
    mount_name = "experiments"
}
```

## Argument Reference

The following arguments are required:

* `cluster_id` - (Optional) (String) [Cluster](cluster.md) to use for mounting. If no cluster is specified, a new cluster will be created and will mount the bucket for all of the clusters in this workspace. If a cluster is specified, mount will be visible for all clusters with the same [instance profile](./instance_profile.md). If the cluster is not running - it's going to be started, so be aware to set auto-termination rules on it.
* `instance_profile` - (Optional) (String) ARN of registered [instance profile](instance_profile.md) for data access.
* `mount_name` - (Required) (String) Name, under which mount will be accessible in `dbfs:/mnt/<MOUNT_NAME>` or locally on each instance through FUSE mount `/dbfs/mnt/<MOUNT_NAME>`.
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
