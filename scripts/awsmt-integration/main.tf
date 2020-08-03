module "this" {
  source = "../mws-integration"
}

data "external" "env" {
  program = ["python", "-c", "import sys,os,json;json.dump(dict(os.environ), sys.stdout)"]
}

locals {
  prefix = module.this.test_prefix
  account_id = module.this.databricks_account_id
  tags = {
    Environment = "Testing"
    Owner       = data.external.env.result.OWNER
    Epoch       = module.this.test_prefix
  }
}

// initialize provider in "MWS" mode to provision new workspace
provider "databricks" {
  alias = "mws"
  host  = "https://accounts.cloud.databricks.com"
}

// register cross-account ARN
resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = local.account_id
  credentials_name = "${local.prefix}-creds"
  role_arn         = module.this.test_crossaccount_arn
}

// register root bucket
resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = local.account_id
  storage_configuration_name = "${local.prefix}-storage"
  bucket_name                = module.this.test_root_bucket
}

// register VPC
resource "databricks_mws_networks" "this" {
  provider     = databricks.mws
  account_id   = local.account_id
  network_name = "${local.prefix}-network"
  vpc_id       = module.this.test_vpc_id
  subnet_ids = [module.this.test_subnet_public, module.this.test_subnet_private]
  security_group_ids = [module.this.test_security_group]
}

// create workspace in given VPC with DBFS on root bucket
resource "databricks_mws_workspaces" "this" {
  provider        = databricks.mws
  account_id      = local.account_id
  workspace_name  = local.prefix
  deployment_name = local.prefix
  aws_region      = module.this.test_region

  credentials_id            = databricks_mws_credentials.this.credentials_id
  storage_configuration_id  = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id                = databricks_mws_networks.this.network_id
  verify_workspace_runnning = true
}

// initialize provider in normal mode
provider "databricks" {
  // in normal scenario you won't have to give providers aliases
  alias = "created_workspace" 
  
  host  = databricks_mws_workspaces.this.workspace_url
}

// create PAT token to provision entities within workspace
resource "databricks_token" "pat" {
  provider = databricks.created_workspace
  comment  = "Terraform Provisioning"
  // 1 day token
  lifetime_seconds = 86400
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

// block all public access to created bucket
resource "aws_s3_bucket_public_access_block" "this" {
  bucket              = aws_s3_bucket.this.id
  ignore_public_acls  = true
}

output "cloud_env" {
  // needed to distinguish between azure, aws & mws tests
  value = "AWS"
}

// export bucket name to test mounting
output "test_bucket" {
  value = aws_s3_bucket.this.bucket
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
