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
  // dltp - databricks labs terraform provider
  prefix = "dltp${random_string.naming.result}"
  tags = {
    E2          = "yes"
    Environment = "Testing"
    Owner       = data.external.env.result.OWNER
    Epoch       = random_string.naming.result
  }
}

// initialize provider in "MWS" mode to provision new workspace
provider "databricks" {
  alias = "mws"
  host  = "https://accounts.cloud.databricks.com"
}

data "databricks_aws_assume_role_policy" "this" {
  external_id = data.external.env.result.DATABRICKS_ACCOUNT_ID
}

resource "aws_iam_role" "cross_account_role" {
  name               = "${local.prefix}-crossaccount"
  assume_role_policy = data.databricks_aws_assume_role_policy.this.json
  tags               = local.tags
}

data "databricks_aws_crossaccount_policy" "this" {
  pass_roles = [aws_iam_role.data_role.arn]
}

resource "aws_iam_role_policy" "this" {
  name   = "${local.prefix}-policy"
  role   = aws_iam_role.cross_account_role.id
  policy = data.databricks_aws_crossaccount_policy.this.json
}

// register cross-account ARN
resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = data.external.env.result.DATABRICKS_ACCOUNT_ID
  role_arn         = aws_iam_role.cross_account_role.arn
  credentials_name = "${local.prefix}-creds"

  // not explicitly needed by this, but to make sure a smooth deployment
  depends_on = [aws_iam_role_policy.this]
}

module "this" {
  source     = "../modules/aws-mws-common"
  cidr_block = data.external.env.result.TEST_CIDR
  region     = data.external.env.result.TEST_REGION
  prefix     = local.prefix
  tags       = local.tags
}

// register root bucket
resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = data.external.env.result.DATABRICKS_ACCOUNT_ID
  bucket_name                = module.this.root_bucket
  storage_configuration_name = "${local.prefix}-storage"
}

// register VPC
resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = data.external.env.result.DATABRICKS_ACCOUNT_ID
  network_name       = "${local.prefix}-network"
  subnet_ids         = [module.this.subnet_public, module.this.subnet_private]
  vpc_id             = module.this.vpc_id
  security_group_ids = [module.this.security_group]
}

// create workspace in given VPC with DBFS on root bucket
resource "databricks_mws_workspaces" "this" {
  provider        = databricks.mws
  account_id      = data.external.env.result.DATABRICKS_ACCOUNT_ID
  aws_region      = data.external.env.result.TEST_REGION
  workspace_name  = local.prefix
  deployment_name = local.prefix

  credentials_id            = databricks_mws_credentials.this.credentials_id
  storage_configuration_id  = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id                = databricks_mws_networks.this.network_id
  verify_workspace_runnning = true
}

// initialize provider in normal mode
provider "databricks" {
  // in normal scenario you won't have to give providers aliases
  alias = "created_workspace"

  host = databricks_mws_workspaces.this.workspace_url
}

// create PAT token to provision entities within workspace
resource "databricks_token" "pat" {
  provider = databricks.created_workspace
  comment  = "Terraform Provisioning"
  // 1 day token
  lifetime_seconds = 86400
}

output "cloud_env" {
  // needed to distinguish between azure, aws & mws tests
  value = "AWS"
}

// export host for integration tests to run on
output "databricks_host" {
  value = databricks_mws_workspaces.this.workspace_url
}

// export token for integraiton tests to run on
output "databricks_token" {
  value     = databricks_token.pat.token_value
  sensitive = true
}

// remove username from environment
output "databricks_username" {
  value = ""
}

// remove password from environment
output "databricks_password" {
  value = ""
}
