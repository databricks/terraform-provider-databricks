provider "aws" {
}

provider "random" {
  version = "~> 2.2"
}

// get any env var to tf
data "external" "env" {
  program = ["python", "-c", "import sys,os,json;json.dump(dict(os.environ), sys.stdout)"]
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

variable "databricks_aws_acct_id" {
  // public info - https://docs.databricks.com/administration-guide/account-settings/aws-accounts.html
  default = "414351767826"
  type    = string
}

locals {
  // dltp - databricks labs terraform provider
  prefix = "dltp${random_string.naming.result}"
  tags = {
    Environment = "Testing"
    Owner       = data.external.env.result.OWNER
    Epoch       = random_string.naming.result
  }
}

data "template_file" "cross_account_role_policy" {
  template = file("${path.module}/templates/cross_account_role_policy.tpl")
}

data "template_file" "cross_account_role_assume_policy" {
  template = file("${path.module}/templates/cross_account_role_assume_policy.tpl")
  vars = {
    databricks_app_external_id = data.external.env.result.DATABRICKS_ACCOUNT_ID
    databricks_aws_account_id = var.databricks_aws_acct_id
  }
}

resource "aws_iam_role" "cross_account_role" {
  name = "${local.prefix}-crossaccount"
  assume_role_policy = data.template_file.cross_account_role_assume_policy.rendered
  tags               = local.tags
}

resource "aws_iam_policy" "cross_account_role_policy" {
  name = "${local.prefix}-policy"
  description = "E2 Workspace Cross account role policy policy"
  policy = data.template_file.cross_account_role_policy.rendered
}

resource "aws_iam_role_policy_attachment" "cross_account_role_policy_attach" {
  role       = aws_iam_role.cross_account_role.name
  policy_arn = aws_iam_policy.cross_account_role_policy.arn
}

data "template_file" "storage_bucket_policy" {
  template = file("${path.module}/templates/storage_bucket_policy.tpl")
  vars = {
    bucket_name = aws_s3_bucket.root_storage_bucket.bucket
    databricks_aws_account_id = var.databricks_aws_acct_id
  }
}

resource "aws_s3_bucket" "root_storage_bucket" {
  bucket = "${local.prefix}-rootbucket"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(local.tags, {
    Name = "${local.prefix}-rootbucket"
  })
}

resource "aws_s3_bucket_public_access_block" "root_storage_bucket" {
  bucket              = aws_s3_bucket.root_storage_bucket.id
  ignore_public_acls  = true
}

resource "aws_s3_bucket_policy" "root_bucket_policy" {
  bucket = aws_s3_bucket.root_storage_bucket.id
  policy = data.template_file.storage_bucket_policy.rendered
}

resource "aws_vpc" "main" {
  cidr_block           = data.external.env.result.TEST_CIDR
  enable_dns_hostnames = true

  tags = merge(local.tags, {
    Name = "${local.prefix}-vpc"
  })
}

resource "aws_subnet" "public" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(aws_vpc.main.cidr_block, 3, 0)
  availability_zone = "${data.external.env.result.TEST_REGION}b"

  tags = merge(local.tags, {
    Name = "${local.prefix}-public-sn"
  })
}

resource "aws_subnet" "private" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(aws_vpc.main.cidr_block, 3, 1)
  availability_zone = "${data.external.env.result.TEST_REGION}a"

  tags = merge(local.tags, {
    Name = "${local.prefix}-private-sn"
  })
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id
  tags = merge(local.tags, {
    Name = "${local.prefix}-igw"
  })
}

resource "aws_route" "r" {
  route_table_id            = aws_vpc.main.default_route_table_id
  destination_cidr_block    = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.gw.id

  depends_on = [aws_internet_gateway.gw, aws_vpc.main]
}

resource "aws_security_group" "test_sg" {
  name        = "all all"
  description = "Allow inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "All"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = merge(local.tags, {
    Name = "${local.prefix}-sg"
  })
}

output "cloud_env" {
  // needed to distinguish between azure, aws & mws tests
  value = "MWS"
}

output "test_root_bucket" {
  value = aws_s3_bucket.root_storage_bucket.bucket
}

output "test_crossaccount_arn" {
  value = aws_iam_role.cross_account_role.arn
}

output "test_vpc_id" {
  value = aws_vpc.main.id
}

output "test_subnet_public" {
  value = aws_subnet.public.id
}

output "test_subnet_private" {
  value = aws_subnet.private.id
}

output "test_security_group" {
  value = aws_security_group.test_sg.id
}

output "test_prefix" {
  value = local.prefix
}

output "test_region" {
  value = data.external.env.result.TEST_REGION
}

output "databricks_account_id" {
  value = data.external.env.result.DATABRICKS_ACCOUNT_ID
}