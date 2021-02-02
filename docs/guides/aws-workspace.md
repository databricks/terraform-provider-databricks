# Provisioning AWS Databricks E2

## Provider initialization for E2 workspaces

This guide assumes you have `databricks_account_username` and `databricks_account_password` for [https://accounts.cloud.databricks.com](https://accounts.cloud.databricks.com) and can find `databricks_account_id` in the top right corner of the page, once you're logged in. This guide is provided as is and assumes you'll use it as the basis for your setup.

```hcl
variable "databricks_account_username" {}
variable "databricks_account_password" {}
variable "databricks_account_id" {}

variable "cidr_block" {
    default = "10.4.0.0/16"
}

variable "region" {
  default = "eu-west-1"
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

locals {
  prefix = "databricks${random_string.naming.result}"
  tags = {
    Environment = "Demo"
    Owner       = var.databricks_account_username
  }
}
```

Before [managing workspace](workspace-management.md), you have to create [VPC](#vpc), [root bucket](#root-bucket), [cross-account role](#cross-account-iam-role), [Databricks E2 workspace](#databricks-e2-workspace), and [host and token outputs](#provider-configuration). Initializing provider with `alias = "mws"` and using `provider = databricks.mws` for all `databricks_mws_*` resources is ultimately important. We require all `databricks_mws_*` resources to be created within it's own dedicated terraform module of your environment. Usually this module creates VPC and IAM roles as well.

```hcl
terraform {
  required_providers {
    databricks = {
      source  = "databrickslabs/databricks"
      version = "0.3.0"
    }
  }
}

provider "aws" {
  region = var.region
}

// initialize provider in "MWS" mode to provision new workspace
provider "databricks" {
  alias    = "mws"
  host     = "https://accounts.cloud.databricks.com"
  databricks_account_username = var.databricks_account_username
  databricks_account_password = var.databricks_account_password
}
```

## Cross-account IAM Role

Cross-account IAM role is registered with [databricks_mws_credentials](../resources/mws_credentials.md) resource.

```hcl

data "databricks_aws_assume_role_policy" "this" {
  external_id = var.databricks_account_id
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

resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = var.databricks_account_id
  role_arn         = aws_iam_role.cross_account_role.arn
  credentials_name = "${local.prefix}-creds"
  depends_on       = [aws_iam_role_policy.this]
}
```

## VPC

The very first step is VPC creation with necessary firewall rules. Please consult [main documetation page](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) for **the most complete and up-to-date details on networking**. AWS VPS is registered as [databricks_mws_networks](../resources/mws_networks.md) resource.

```hcl
locals {
  allows = {
    "us-east-1" : [
      "52.27.216.188/32",
      "54.156.226.103/32",
    ],
    "us-east-2" : [
      "52.27.216.188/32",
      "18.221.200.169/32",
    ],
    "us-west-2" : [
      "10.200.0.0/16"
    ],
    "eu-west-1" : [
      "46.137.47.49/32",
    ]
  }
}

resource "aws_vpc" "main" {
  cidr_block           = var.cidr_block
  enable_dns_hostnames = true

  tags = merge(var.tags, {
    Name = "${local.prefix}-vpc"
  })
}

resource "aws_subnet" "public" {
  vpc_id     = aws_vpc.main.id
  cidr_block = cidrsubnet(aws_vpc.main.cidr_block, 3, 0)
  tags = merge(var.tags, {
    Name = "${local.prefix}-public-sn"
  })
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-igw"
  })
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id
  route {
    gateway_id = aws_internet_gateway.gw.id
    cidr_block = "0.0.0.0/0"
  }

  tags = merge(var.tags, {
    Name = "${local.prefix}-public-rt"
  })
}

resource "aws_route_table_association" "public" {
  route_table_id = aws_route_table.public.id
  subnet_id      = aws_subnet.public.id
}

resource "aws_eip" "nat" {
  vpc        = true
  depends_on = [aws_internet_gateway.gw]
  tags = merge(var.tags, {
    Name = "${local.prefix}-eip"
  })
}

resource "aws_nat_gateway" "gw" {
  allocation_id = aws_eip.nat.id
  subnet_id     = aws_subnet.public.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-nat"
  })
}

resource "aws_subnet" "private" {
  vpc_id     = aws_vpc.main.id
  cidr_block = cidrsubnet(aws_vpc.main.cidr_block, 3, 1)
  tags = merge(var.tags, {
    Name = "${local.prefix}-private-sn"
  })
}

resource "aws_subnet" "private_secondary" {
  vpc_id     = aws_vpc.main.id
  cidr_block = cidrsubnet(aws_vpc.main.cidr_block, 3, 2)
  tags = merge(var.tags, {
    Name = "${local.prefix}-private-sn"
  })
}

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id
  route {
    nat_gateway_id = aws_nat_gateway.gw.id
    cidr_block     = "0.0.0.0/0"
  }
  tags = merge(var.tags, {
    Name = "${local.prefix}-private-rt"
  })
}

resource "aws_route_table_association" "private" {
  route_table_id = aws_route_table.private.id
  subnet_id      = aws_subnet.private.id
}

resource "aws_route_table_association" "private_secondary" {
  route_table_id = aws_route_table.private.id
  subnet_id      = aws_subnet.private_secondary.id
}

resource "aws_security_group" "this" {
  name        = "all all"
  description = "Allow inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "All Internal TCP"
    protocol    = "tcp"
    from_port   = 0
    to_port     = 65535
    self        = true
  }

  ingress {
    description = "All Internal UDP"
    protocol    = "udp"
    from_port   = 0
    to_port     = 65535
    self        = true
  }

  ingress {
    cidr_blocks = local.allows[var.region]
    description = "ssh"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
  }

  ingress {
    cidr_blocks = local.allows[var.region]
    description = "proxy"
    from_port   = 5557
    to_port     = 5557
    protocol    = "tcp"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = merge(var.tags, {
    Name = "${local.prefix}-sg"
  })
}

resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = var.databricks_account_id
  network_name       = "${local.prefix}-network"
  subnet_ids         = [aws_subnet.private.id, aws_subnet.private_secondary.id]
  vpc_id             = aws_vpc.main.id
  security_group_ids = [aws_security_group.this.id]
}
```

## Root bucket

Once [VPC](#vpc) is ready, you need to create AWS S3 bucket for DBFS workspace storage, which is commonly referred to as **root bucket**. This provider has [databricks_aws_bucket_policy](../data-sources/aws_bucket_policy.md) with the necessary IAM policy template. AWS S3 bucket has to be registered through [databricks_mws_storage_configurations](../resources/mws_storage_configurations.md).

```hcl
resource "aws_s3_bucket" "root_storage_bucket" {
  bucket = "${local.prefix}-rootbucket"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(var.tags, {
    Name = "${local.prefix}-rootbucket"
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

resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = var.databricks_account_id
  bucket_name                = aws_s3_bucket.root_storage_bucket.bucket
  storage_configuration_name = "${local.prefix}-storage"
}
```

## Databricks E2 Workspace

Once you have [VPC](#vpc), [cross-account role](#cross-account-iam-role), and [root bucket](#root-bucket) setup, you can create Databricks AWS E2 workspace through [databricks_mws_workspaces](../resources/mws_workspaces.md) resource. Code, that creates workpaces and code that [manages workspaces](workspace-management.md) must be in separate terraform modules to avoid common confusion between `provider = databricks.mws` and `provider = databricks.created_workspace`. This is we specify `databricks_host` and `databricks_token` outputs, that have to be used in the latter modules.

```hcl
resource "databricks_mws_workspaces" "this" {
  provider        = databricks.mws
  account_id      = var.databricks_account_id
  aws_region      = var.region
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

// export host for integration tests to run on
output "databricks_host" {
  value = databricks_mws_workspaces.this.workspace_url
}

// export token for integraiton tests to run on
output "databricks_token" {
  value     = databricks_token.pat.token_value
  sensitive = true
}
```

## Provider configuration

In [the next step](workspace-management.md), please use the following configuration for the provider:

```hcl
provider "databricks" {
  host = module.ai.databricks_host
  token = module.ai.databricks_token
}
```