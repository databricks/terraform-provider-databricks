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

locals {
  // dltp - databricks labs terraform provider
  prefix = "dltp${random_string.naming.result}"
  tags = {
    Environment = "Testing"
    Owner       = data.external.env.result.OWNER
    Epoch       = random_string.naming.result
  }
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
  source = "../modules/aws-mws-common"
  cidr_block = data.external.env.result.TEST_CIDR
  region = data.external.env.result.TEST_REGION
  prefix = local.prefix
  tags = local.tags
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
  value = data.external.env.result.TEST_REGION
}

output "databricks_account_id" {
  value = data.external.env.result.DATABRICKS_ACCOUNT_ID
  sensitive = true
}