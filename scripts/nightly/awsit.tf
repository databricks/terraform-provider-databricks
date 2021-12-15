terraform {
  required_providers {
    databricks = {
      source = "databrickslabs/databricks"
    }
  }
}

provider "aws" {
  region = local.region
}

output "aws_region" {
  value = local.region
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

output "databricks_account_id" {
  value = local.account_id
}

resource "aws_iam_role" "cross_account_role" {
  name               = "${local.prefix}-cx-terraform-it"
  assume_role_policy = data.databricks_aws_assume_role_policy.this.json
  tags               = local.tags
}

data "databricks_aws_crossaccount_policy" "this" {
  provider   = databricks.mws
  pass_roles = [aws_iam_role.data_role.arn]
}

resource "aws_iam_role_policy" "this" {
  name   = "${local.prefix}-cx-terraform-it"
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

output "aws_crossaccount_role" {
  value = aws_iam_role.cross_account_role.arn
}

resource "aws_s3_bucket" "root_storage_bucket" {
  bucket = "${local.prefix}-cx-terraform-it"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(local.tags, {
    Name = "${local.prefix}-cx-terraform-it"
  })
}

resource "aws_s3_bucket_public_access_block" "root_storage_bucket" {
  bucket             = aws_s3_bucket.root_storage_bucket.id
  ignore_public_acls = true
  depends_on         = [aws_s3_bucket.root_storage_bucket]
}

data "databricks_aws_bucket_policy" "this" {
  bucket = aws_s3_bucket.root_storage_bucket.bucket
}

resource "aws_s3_bucket_policy" "root_bucket_policy" {
  bucket = aws_s3_bucket.root_storage_bucket.id
  policy = data.databricks_aws_bucket_policy.this.json
}

// register root bucket
resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = local.account_id
  bucket_name                = aws_s3_bucket.root_storage_bucket.bucket
  storage_configuration_name = "${local.prefix}-cx-terraform-it"
}

output "aws_bucket" {
  value = aws_s3_bucket.root_storage_bucket.bucket
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
  single_nat_gateway   = true
  create_igw           = true

  public_subnets = [cidrsubnet(local.cidr_block, 3, 0)]
  private_subnets = [cidrsubnet(local.cidr_block, 3, 1),
  cidrsubnet(local.cidr_block, 3, 2)]

  manage_default_security_group = true
  default_security_group_name   = "${local.prefix}-sg"

  default_security_group_egress = [{
    cidr_blocks = "0.0.0.0/0"
  }]

  default_security_group_ingress = [{
    description = "Allow all internal TCP and UDP"
    self        = true
  }]
}

resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = local.account_id
  network_name       = "${local.prefix}-network"
  security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids         = module.vpc.private_subnets
  vpc_id             = module.vpc.vpc_id
}

output "aws_vpc_id" {
  value = module.vpc.vpc_id
}

output "aws_vpc_cidr" {
  value = local.cidr_block
}

output "aws_security_group_ids" {
  value = [module.vpc.default_security_group_id]
}

output "databricks_credentials_id" {
  value = databricks_mws_credentials.this.credentials_id
}

output "databricks_storage_configuration_id" {
  value = databricks_mws_storage_configurations.this.storage_configuration_id
}

// create workspace in given VPC with DBFS on root bucket
resource "databricks_mws_workspaces" "this" {
  provider        = databricks.mws
  account_id      = local.account_id
  aws_region      = local.region
  workspace_name  = local.prefix
  deployment_name = local.prefix

  credentials_id           = databricks_mws_credentials.this.credentials_id
  storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id               = databricks_mws_networks.this.network_id

  token {}
}

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

resource "azurerm_container_group" "aws" {
  name                = "${local.prefix}-aws-run"
  location            = azurerm_resource_group.this.location
  resource_group_name = azurerm_resource_group.this.name
  tags                = azurerm_resource_group.this.tags

  os_type            = "Linux"
  restart_policy     = "Never"
  ip_address_type    = "Private"
  network_profile_id = azurerm_network_profile.this.id

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.this.id]
  }

  container {
    name   = "acceptance"
    image  = "ghcr.io/databrickslabs/terraform-provider-it:master"
    cpu    = "2"
    memory = "2"
    environment_variables = {
      CLOUD_ENV                 = "AWS"
      TEST_FILTER               = "TestAcc"
      DATABRICKS_HOST           = databricks_mws_workspaces.this.workspace_url
      TEST_S3_BUCKET            = aws_s3_bucket.ds.bucket
      TEST_EC2_INSTANCE_PROFILE = aws_iam_instance_profile.this.arn
    }

    secure_environment_variables = {
      DATABRICKS_TOKEN = databricks_mws_workspaces.this.token[0].token_value
    }

    ports {
      port     = 443
      protocol = "TCP"
    }
  }
}