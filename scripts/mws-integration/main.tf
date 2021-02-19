// get any env var to tf
data "external" "env" {
  program = ["python", "-c", "import sys,os,json;json.dump(dict(os.environ), sys.stdout)"]
}

provider "random" {}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

locals {
  // dltp - databricks labs terraform provider
  prefix = "dltp${random_string.naming.result}"
  region = data.external.env.result.TEST_REGION
  tags = {
    Environment = "Testing"
    Owner       = data.external.env.result.OWNER
    Epoch       = random_string.naming.result
  }
  pl_user_to_workspace = {
    "us-west-2" : "com.amazonaws.vpce.us-west-2.vpce-svc-0129f463fcfbc46c5",
    "us-east-1" : "com.amazonaws.vpce.us-east-1.vpce-svc-09143d1e626de2f04",
    "eu-west-1" : "com.amazonaws.vpce.eu-west-1.vpce-svc-0da6ebf1461278016",
  }
  pl_dataplane_to_controlplane = {
    "us-west-2" : "com.amazonaws.vpce.us-west-2.vpce-svc-0129f463fcfbc46c5"
    "us-east-1" : "com.amazonaws.vpce.us-east-1.vpce-svc-09143d1e626de2f04",
    # "eu-west-1": "com.amazonaws.vpce.eu-west-1.vpce-svc-0da6ebf1461278016",
    "eu-west-1" : "com.amazonaws.vpce.eu-west-1.vpce-svc-09b4eb2bc775f4e8c",
  }
}

provider "aws" {
  region = local.region
}

provider "databricks" {}

data "databricks_aws_assume_role_policy" "this" {
  external_id = data.external.env.result.DATABRICKS_ACCOUNT_ID
}

resource "aws_iam_role" "cross_account_role" {
  name               = "${local.prefix}-crossaccount"
  assume_role_policy = data.databricks_aws_assume_role_policy.this.json
  tags               = local.tags
}

data "databricks_aws_crossaccount_policy" "this" {
}

resource "aws_iam_role_policy" "test_policy" {
  name   = "test_policy"
  role   = aws_iam_role.cross_account_role.id
  policy = data.databricks_aws_crossaccount_policy.this.json
}

module "aws_common" {
  source     = "../modules/aws-mws-common"
  cidr_block = data.external.env.result.TEST_CIDR
  region     = local.region
  prefix     = local.prefix
  tags       = local.tags
}

resource "aws_s3_bucket" "logdelivery" {
  bucket = "${local.prefix}-logdelivery"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = local.tags
}

output "test_logdelivery_bucket" {
  value = aws_s3_bucket.logdelivery.bucket
}

resource "aws_s3_bucket_public_access_block" "logdelivery" {
  bucket             = aws_s3_bucket.logdelivery.id
  ignore_public_acls = true
}

data "databricks_aws_assume_role_policy" "logdelivery" {
  external_id      = data.external.env.result.DATABRICKS_ACCOUNT_ID
  for_log_delivery = true
}

resource "aws_iam_role" "logdelivery" {
  name               = "${local.prefix}-logdelivery"
  description        = "(${local.prefix}) UsageDelivery role"
  assume_role_policy = data.databricks_aws_assume_role_policy.logdelivery.json
  tags               = local.tags
}

output "test_logdelivery_arn" {
  value = aws_iam_role.logdelivery.arn
}

data "databricks_aws_bucket_policy" "logdelivery" {
  full_access_role = aws_iam_role.logdelivery.arn
  bucket           = aws_s3_bucket.logdelivery.bucket
}

resource "aws_s3_bucket_policy" "logdelivery" {
  bucket = aws_s3_bucket.logdelivery.id
  policy = data.databricks_aws_bucket_policy.logdelivery.json
}

resource "aws_vpc_endpoint" "relay" {
  service_name       = local.pl_dataplane_to_controlplane[local.region]
  vpc_id             = module.aws_common.vpc_id
  vpc_endpoint_type  = "Interface"
  security_group_ids = [module.aws_common.security_group]
  subnet_ids         = [module.aws_common.subnet_private]
}

output "test_relay_vpc_endpoint" {
  value = aws_vpc_endpoint.relay.id
}

resource "aws_vpc_endpoint" "rest_api" {
  service_name       = local.pl_dataplane_to_controlplane[local.region]
  vpc_id             = module.aws_common.vpc_id
  vpc_endpoint_type  = "Interface"
  security_group_ids = [module.aws_common.security_group]
  subnet_ids         = [module.aws_common.subnet_private]
}

data "aws_caller_identity" "current" {}

output "test_aws_account_id" {
  value = data.aws_caller_identity.current.account_id
}

output "test_rest_api_vpc_endpoint" {
  value = aws_vpc_endpoint.rest_api.id
}

output "cloud_env" {
  // needed to distinguish between azure, aws & mws tests
  value = "MWS"
}

output "test_root_bucket" {
  value = module.aws_common.root_bucket
}

output "test_crossaccount_arn" {
  value = aws_iam_role.cross_account_role.arn
}

output "test_vpc_id" {
  value = module.aws_common.vpc_id
}

output "test_subnet_public" {
  value = module.aws_common.subnet_public
}

output "test_subnet_private" {
  value = module.aws_common.subnet_private
}

output "test_security_group" {
  value = module.aws_common.security_group
}

output "test_kms_key_arn" {
  value = module.aws_common.kms_key_arn
}

output "test_kms_key_alias" {
  value = module.aws_common.kms_key_alias
}

output "test_prefix" {
  value = local.prefix
}

output "test_region" {
  value = local.region
}

output "databricks_account_id" {
  value     = data.external.env.result.DATABRICKS_ACCOUNT_ID
  sensitive = true
}
