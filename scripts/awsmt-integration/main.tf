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
  prefix     = "dltp${random_string.naming.result}"
  cidr_block = data.external.env.result.TEST_CIDR
  region     = data.external.env.result.TEST_REGION
  account_id = data.external.env.result.DATABRICKS_ACCOUNT_ID
  tags = {
    Environment = "Testing"
    Owner       = data.external.env.result.OWNER
    Epoch       = random_string.naming.result
  }
}

provider "aws" {
  region = local.region
}

// initialize provider in "MWS" mode to provision new workspace
provider "databricks" {
  alias = "mws"
  host  = "https://accounts.cloud.databricks.com"
}

data "databricks_aws_assume_role_policy" "this" {
  provider    = databricks.mws
  external_id = local.account_id
}

resource "aws_iam_role" "cross_account_role" {
  name               = "${local.prefix}-crossaccount"
  assume_role_policy = data.databricks_aws_assume_role_policy.this.json
  tags               = local.tags
}

data "databricks_aws_crossaccount_policy" "this" {
  provider   = databricks.mws
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
  account_id       = local.account_id
  role_arn         = aws_iam_role.cross_account_role.arn
  credentials_name = "${local.prefix}-creds"

  // not explicitly needed by this, but to make sure a smooth deployment
  depends_on = [aws_iam_role_policy.this]
}

module "this" {
  source     = "../modules/aws-mws-common"
  cidr_block = local.cidr_block
  region     = local.region
  prefix     = local.prefix
  tags       = local.tags
}

// register root bucket
resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = local.account_id
  bucket_name                = module.this.root_bucket
  storage_configuration_name = "${local.prefix}-storage"
}

// register VPC
data "aws_availability_zones" "available" {}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.2.0"

  name = local.prefix
  cidr = local.cidr_block
  azs  = data.aws_availability_zones.available.names
  tags = local.tags

  enable_dns_hostnames = true
  enable_nat_gateway   = true
  create_igw           = true

  public_subnets  = [cidrsubnet(local.cidr_block, 3, 0)]
  private_subnets = [cidrsubnet(local.cidr_block, 3, 1),
                     cidrsubnet(local.cidr_block, 3, 2)]

  manage_default_security_group = true
  default_security_group_name = "${local.prefix}-sg"

  default_security_group_egress = [{
    cidr_blocks = "0.0.0.0/0"
  }]

  default_security_group_ingress = [{
    description = "Allow all internal TCP and UDP"
    self        = true
  }]
}

module "vpc_endpoints" {
  source = "terraform-aws-modules/vpc/aws//modules/vpc-endpoints"
  version = "3.2.0"

  vpc_id             = module.vpc.vpc_id
  security_group_ids = [module.vpc.default_security_group_id]

  endpoints = {
    s3 = {
      service         = "s3"
      service_type    = "Gateway"
      route_table_ids = flatten([
        module.vpc.private_route_table_ids, 
        module.vpc.public_route_table_ids])
      tags            = {
        Name = "${local.prefix}-s3-vpc-endpoint"
      }
    },
    sts = {
      service             = "sts"
      private_dns_enabled = true
      subnet_ids          = module.vpc.private_subnets
      tags                = {
        Name = "${local.prefix}-sts-vpc-endpoint"
      }
    },
    kinesis-streams = {
      service             = "kinesis-streams"
      private_dns_enabled = true
      subnet_ids          = module.vpc.private_subnets
      tags                = {
        Name = "${local.prefix}-kinesis-vpc-endpoint"
      }
    },    
  }

  tags = local.tags
}

resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = local.account_id
  network_name       = "${local.prefix}-network"
  security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids         = module.vpc.private_subnets
  vpc_id             = module.vpc.vpc_id
}

// create workspace in given VPC with DBFS on root bucket
resource "databricks_mws_workspaces" "this" {
  provider        = databricks.mws
  account_id      = local.account_id
  aws_region      = local.region
  workspace_name  = local.prefix
  deployment_name = local.prefix

  credentials_id            = databricks_mws_credentials.this.credentials_id
  storage_configuration_id  = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id                = databricks_mws_networks.this.network_id
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
