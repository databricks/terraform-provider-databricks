// create bucket for mounting
resource "aws_s3_bucket" "ds" {
  bucket = "${local.prefix}-ds"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(local.tags, {
    Name = "${local.prefix}-ds"
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
  name               = "${local.prefix}-first-ec2s3"
  description        = "(${local.prefix}) EC2 Assume Role role for S3 access"
  assume_role_policy = data.aws_iam_policy_document.assume_role_for_ec2.json
  tags               = local.tags
}

resource "aws_iam_instance_profile" "this" {
  name = "${local.prefix}-first-profile"
  role = aws_iam_role.data_role.name
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

// block all public access to created bucket
resource "aws_s3_bucket_public_access_block" "this" {
  bucket             = aws_s3_bucket.ds.id
  ignore_public_acls = true
}

resource "aws_s3_bucket_object" "this" {
  key    = "/dummy-${aws_s3_bucket_public_access_block.this.bucket}/main.tf"
  bucket = aws_s3_bucket.ds.id
  source = "${path.module}/main.tf"
  tags   = local.tags
}

output "test_s3_bucket" {
  value = aws_s3_bucket.ds.bucket
}

output "test_ec2_instance_profile" {
  value = aws_iam_instance_profile.this.arn
}