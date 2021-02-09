---
subcategory: "AWS"
---
# databricks_aws_bucket_policy Data Source

This datasource configures a simple access policy for AWS S3 buckets, so that Databricks can access data in it. 

## Example Usage

```hcl
resource "aws_s3_bucket" "this" {
  bucket = "things"
  region = "eu-west-1"
  acl    = "private"
  force_destroy = true
}

data "databricks_aws_bucket_policy" "stuff" {
  bucket_name = aws_s3_bucket.this.bucket
}

resource "aws_s3_bucket_policy" "this" {
  bucket     = aws_s3_bucket.this.id
  policy     = data.databricks_aws_bucket_policy.this.json
}
```

Bucket policy with full access:

```hcl
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

* `bucket` - (Required) AWS S3 Bucket name for which to generate the policy document.
* `full_access_role` - (Optional) Data access role that can have full access for this bucket

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - (Read-only) AWS IAM Policy JSON document to grant Databricks full access to bucket.
