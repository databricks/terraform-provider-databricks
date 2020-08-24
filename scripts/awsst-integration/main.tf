data "external" "env" {
  program = ["python", "-c", "import sys,os,json;json.dump(dict(os.environ), sys.stdout)"]
}

provider "aws" {
  region = data.external.env.result.TEST_REGION
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

locals {
  prefix = "dltp${random_string.naming.result}"
  tags = {
    Environment = "Testing"
    Owner       = data.external.env.result.OWNER
    Epoch       = random_string.naming.result
  }
}

// create bucket for mounting
resource "aws_s3_bucket" "this" {
  bucket = "${local.prefix}-test"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(local.tags, {
    Name = "${local.prefix}-test"
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

resource "aws_iam_role" "role_for_s3_access" {
  name               = "${local.prefix}-ec2s3"
  description        = "(${local.prefix}) EC2 Assume Role role for S3 access"
  assume_role_policy = data.aws_iam_policy_document.assume_role_for_ec2.json
  tags               = local.tags
}

resource "aws_iam_instance_profile" "this" {
  name = "${local.prefix}-ip"
  role = aws_iam_role.role_for_s3_access.name
}

data "aws_iam_policy_document" "bucket_access" {
  statement {
    effect = "Allow"
    actions = ["s3:GetObject",
      "s3:GetObjectVersion",
      "s3:ListBucket",
      "s3:GetBucketLocation",
      "s3:PutObject",
    "s3:DeleteObject"]
    principals {
      identifiers = [aws_iam_role.role_for_s3_access.arn]
      type        = "AWS"
    }
    resources = [
      "${aws_s3_bucket.this.arn}/*",
    aws_s3_bucket.this.arn]
  }
}

// allow databricks to access this bucket
resource "aws_s3_bucket_policy" "this" {
  bucket = aws_s3_bucket.this.id
  policy = data.aws_iam_policy_document.bucket_access.json
}

// block all public access to created bucket
resource "aws_s3_bucket_public_access_block" "this" {
  bucket             = aws_s3_bucket.this.id
  ignore_public_acls = true
}

resource "aws_s3_bucket_object" "this" {
  key    = "/dummy-${aws_s3_bucket_public_access_block.this.bucket}/main.tf"
  bucket = aws_s3_bucket.this.id
  source = "${path.module}/main.tf"
  tags   = local.tags
}


output "cloud_env" {
  // needed to distinguish between azure, aws & mws tests
  value = "AWS"
}

output "test_s3_bucket" {
  value = aws_s3_bucket.this.bucket
}

output "test_ec2_instance_profile" {
  value = aws_iam_instance_profile.this.arn
}