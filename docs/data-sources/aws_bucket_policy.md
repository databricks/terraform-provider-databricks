---
subcategory: "Deployment"
---
# databricks_aws_bucket_policy Data Source

This datasource configures a simple access policy for AWS S3 buckets, so that Databricks can access data in it.

-> This data source can be used with an account or workspace-level provider.

## Example Usage

```hcl
resource "aws_s3_bucket" "this" {
  bucket        = "<unique_bucket_name>"
  force_destroy = true
}

data "databricks_aws_bucket_policy" "this" {
  bucket = aws_s3_bucket.this.bucket
}

resource "aws_s3_bucket_policy" "this" {
  bucket = aws_s3_bucket.this.id
  policy = data.databricks_aws_bucket_policy.this.json
}
```

Bucket policy with full access:

```hcl
resource "aws_s3_bucket" "ds" {
  bucket        = "${var.prefix}-ds"
  force_destroy = true
  tags = merge(var.tags, {
    Name = "${var.prefix}-ds"
  })
}

resource "aws_s3_bucket_versioning" "ds_versioning" {
  bucket = aws_s3_bucket.ds.id
  versioning_configuration {
    status = "Disabled"
  }
}

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

resource "aws_iam_role" "data_role" {
  name               = "${var.prefix}-first-ec2s3"
  description        = "(${var.prefix}) EC2 Assume Role role for S3 access"
  assume_role_policy = data.aws_iam_policy_document.assume_role_for_ec2.json
  tags               = var.tags
}

data "databricks_aws_bucket_policy" "ds" {
  provider         = databricks.mws
  full_access_role = aws_iam_role.data_role.arn
  bucket           = aws_s3_bucket.ds.bucket
}

// allow databricks to access this bucket
resource "aws_s3_bucket_policy" "ds" {
  bucket = aws_s3_bucket.ds.id
  policy = data.databricks_aws_bucket_policy.ds.json
}
```

## Argument Reference

* `bucket` - (Required) AWS S3 Bucket name for which to generate the policy document. The name must follow the [S3 bucket naming rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
* `aws_partition` - (Optional) AWS partition. The options are `aws`, `aws-us-gov`, or `aws-us-gov-dod`. Defaults to `aws`
* `full_access_role` - (Optional) Data access role that can have full access for this bucket
* `databricks_e2_account_id` - (Optional) Your Databricks account ID. Used to generate  restrictive IAM policies that will increase the security of your root bucket

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - (Read-only) AWS IAM Policy JSON document to grant Databricks full access to bucket.

## Related Resources

The following resources are used in the same context:

* [Provisioning AWS Databricks workspaces with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide.
* [End to end workspace management](../guides/workspace-management.md) guide
* [databricks_instance_profile](../resources/instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](../resources/cluster.md) and access data, like [databricks_mount](../resources/mount.md).
* [databricks_mount](../resources/mount.md) to [mount your cloud storage](https://docs.databricks.com/data/databricks-file-system.html#mount-object-storage-to-dbfs) on `dbfs:/mnt/name`.
