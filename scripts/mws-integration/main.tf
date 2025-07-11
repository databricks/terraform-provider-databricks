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
  region = data.external.env.result.AWS_REGION
  cidr = data.external.env.result.TEST_CIDR
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

data "aws_availability_zones" "available" {}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.2.0"

  name = local.prefix
  cidr = data.external.env.result.TEST_CIDR
  azs  = data.aws_availability_zones.available.names
  tags = local.tags

  enable_dns_hostnames = true
  enable_nat_gateway   = true
  single_nat_gateway   = true
  create_igw           = true

  public_subnets  = [cidrsubnet(local.cidr, 3, 0)]
  private_subnets = [cidrsubnet(local.cidr, 3, 1),
                     cidrsubnet(local.cidr, 3, 2)]

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
  }
  tags = local.tags
}

resource "aws_s3_bucket" "root_storage_bucket" {
  bucket = "${local.prefix}-root-bucket"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(local.tags, {
    Name = "${local.prefix}-root-bucket"
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
  vpc_id             = module.vpc.vpc_id
  vpc_endpoint_type  = "Interface"
  security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids         = module.vpc.private_subnets
}

output "test_relay_vpc_endpoint" {
  value = aws_vpc_endpoint.relay.id
}

resource "aws_vpc_endpoint" "rest_api" {
  service_name       = local.pl_dataplane_to_controlplane[local.region]
  vpc_id             = module.vpc.vpc_id
  vpc_endpoint_type  = "Interface"
  security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids         = module.vpc.private_subnets
}

variable "databricks_aws_account_id" {
  default = "414351767826"
}

resource "aws_kms_key" "customer_managed_key" {
}

resource "aws_kms_grant" "databricks-grant" {
  name              = "databricks-grant"
  key_id            = aws_kms_key.customer_managed_key.key_id
  grantee_principal = "arn:aws:iam::${var.databricks_aws_account_id}:root"

  operations = ["Encrypt", "Decrypt", "DescribeKey",
    "GenerateDataKey", "ReEncryptFrom", "ReEncryptTo",
    "GenerateDataKeyWithoutPlaintext"]
}

resource "aws_kms_alias" "customer_managed_key_alias" {
  name          = "alias/${local.prefix}-customer-key-alias"
  target_key_id = aws_kms_key.customer_managed_key.key_id
}

data "aws_caller_identity" "current" {}

output "aws_account_id" {
  value = data.aws_caller_identity.current.account_id
}

output "aws_region" {
  value = local.region
}

output "test_rest_api_vpc_endpoint" {
  value = aws_vpc_endpoint.rest_api.id
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
  value = module.vpc.vpc_id
}

output "test_subnet_public" {
  value = module.vpc.public_subnets[0]
}

output "test_subnet_private" {
  value = module.vpc.private_subnets[0]
}

output "test_subnet_private2" {
  value = module.vpc.private_subnets[1]
}

output "test_security_group" {
  value = module.vpc.default_security_group_id
}

output "test_managed_kms_key_arn" {
  value = aws_kms_key.customer_managed_key.arn
}

output "test_kms_key_alias" {
  value = aws_kms_alias.customer_managed_key_alias.name
}

output "test_prefix" {
  value = local.prefix
}

output "databricks_account_id" {
  value     = data.external.env.result.DATABRICKS_ACCOUNT_ID
  sensitive = true
}
