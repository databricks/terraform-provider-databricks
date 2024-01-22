variable "account_id" {
  type        = string
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
}

variable "username" {
  type        = string
  description = "Username to access https://accounts.cloud.databricks.com/"
}

variable "password" {
  type        = string
  description = "Password to access https://accounts.cloud.databricks.com/"
}

variable "credentials_id" {
  type        = string
  description = "credentials_id from databricks_mws_credentials"
}

variable "storage_configuration_id" {
  type        = string
  description = "storage_configuration_id from databricks_mws_storage_configurations"
}

variable "vpc_id" {
  type        = string
  description = "AWS VPC id"
}

variable "aws_security_group_ids" {
  description = "AWS VPC SG ids"
}

variable "region" {
  type        = string
  description = "AWS region name"
}

variable "deployment_name" {
  type        = string
  description = "Name of the workspace"
}

provider "databricks" {
  host     = "https://accounts.cloud.databricks.com/"
  username = var.username
  password = var.password
}

provider "aws" {
  region = var.region
}

locals {
  availability_zones = toset([for s in data.aws_subnet.other : s.availability_zone])
  availability_range = range(0, length(local.availability_zones))
  cidr_newbits       = 3
  private_subnets = zipmap(sort(local.availability_zones),
    [for i in range(0, length(local.availability_zones)) :
      cidrsubnet(data.aws_vpc.this.cidr_block, local.cidr_newbits,
      length(local.availability_zones) + 1 + i)
  ])
  route_tables_with_nat = [for rt in data.aws_route_table.vpc : rt.id if anytrue([
    for r in rt.routes : r.nat_gateway_id != ""
  ])]
  rtb_assoc = { for x in flatten([for rt in local.route_tables_with_nat :
    [for subnet in aws_subnet.private : {
      route_table_id = rt
      subnet_id      = subnet.id
  }]]) : "${x.route_table_id}-${x.subnet_id}" => x }
}

data "aws_vpc" "this" {
  id = var.vpc_id
}

data "aws_subnets" "other" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.this.id]
  }
}

data "aws_subnet" "other" {
  for_each = toset(data.aws_subnets.other.ids)
  id       = each.value
}

data "aws_route_tables" "vpc" {
  vpc_id = data.aws_vpc.this.id
}

data "aws_route_table" "vpc" {
  for_each       = data.aws_route_tables.vpc.ids
  route_table_id = each.value
}

resource "aws_subnet" "private" {
  for_each          = local.private_subnets
  cidr_block        = each.value
  availability_zone = each.key
  vpc_id            = data.aws_vpc.this.id
  tags = merge(data.aws_vpc.this.tags, {
    Name = "${var.deployment_name}-private-${each.key}"
  })
}

resource "aws_route_table_association" "private" {
  for_each       = local.rtb_assoc
  subnet_id      = each.value.subnet_id
  route_table_id = each.value.route_table_id
}

resource "databricks_mws_networks" "this" {
  account_id         = var.account_id
  network_name       = "${var.deployment_name}-network"
  vpc_id             = data.aws_vpc.this.id
  subnet_ids         = [for s in aws_subnet.private : s.id]
  security_group_ids = var.aws_security_group_ids
}

resource "databricks_mws_workspaces" "this" {
  account_id      = var.account_id
  aws_region      = var.region
  workspace_name  = var.deployment_name
  deployment_name = var.deployment_name

  credentials_id           = var.credentials_id
  storage_configuration_id = var.storage_configuration_id
  network_id               = databricks_mws_networks.this.network_id

  token {
  }
}

output "host" {
  value = databricks_mws_workspaces.this.workspace_url
}

output "token" {
  value     = databricks_mws_workspaces.this.token[0].token_value
  sensitive = true
}