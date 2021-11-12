---
page_title: "Provisioning AWS Databricks E2 with a Hub & Spoke firewall for data exfiltration protection"
---

# Provisioning AWS Databricks E2 with a Hub & Spoke firewall for data exfiltration protection

You can provision multiple Databricks workspaces with Terraform and where many Databricks workspaces are deployed, we recommend a hub and spoke topology reference architecture, powered by AWS Transit Gateway. The hub will consist of a central inspection and egress virtual private cloud (VPC), while the Spoke VPC houses federated Databricks workspaces for different business units or segregated teams. In this way, you create your own version of a centralized deployment model for your egress architecture, as is recommended for large enterprises. For more information please visit [Data Exfiltration Protection With Databricks on AWS](https://databricks.com/blog/2021/02/02/data-exfiltration-protection-with-databricks-on-aws.html).

![Data Exfiltration](https://databricks.com/wp-content/uploads/2021/02/aws-exfiltration-replace-1.jpg)

## Provider initialization for E2 workspaces

This guide assumes you have `databricks_account_username` and `databricks_account_password` for [https://accounts.cloud.databricks.com](https://accounts.cloud.databricks.com) and can find `databricks_account_id` in the bottom left corner of the page, once you're logged in. This guide is provided as is and assumes you'll use it as the basis for your setup. If you are using AWS Firewall to block most traffic but allow the URLs that Databricks needs to connect to please update the configuration based on your region. You can get the configuration details for your region from [Firewall Appliance](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html#firewall-appliance-infrastructure) document.

```hcl
variable "databricks_account_username" {}
variable "databricks_account_password" {}
variable "databricks_account_id" {}

variable "tags" {
  default = {}
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}
        
variable "spoke_cidr_block" {
  default = "10.173.0.0/16"
}
variable "hub_cidr_block" {
  default = "10.10.0.0/16"
}
variable "region" {
  default = "eu-central-1"
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

locals {
  prefix                           = "demo${random_string.naming.result}"
  spoke_db_private_subnets_cidr    = [cidrsubnet(var.spoke_cidr_block, 3, 0), cidrsubnet(var.spoke_cidr_block, 3, 1)]
  spoke_tgw_private_subnets_cidr   = [cidrsubnet(var.spoke_cidr_block, 3, 2), cidrsubnet(var.spoke_cidr_block, 3, 3)]
  hub_tgw_private_subnets_cidr     = [cidrsubnet(var.hub_cidr_block, 3, 0)]
  hub_nat_public_subnets_cidr      = [cidrsubnet(var.hub_cidr_block, 3, 1)]
  hub_firewall_subnets_cidr        = [cidrsubnet(var.hub_cidr_block, 3, 2)]
  sg_egress_ports                  = [443, 3306, 6666]
  sg_ingress_protocol              = ["tcp", "udp"]
  sg_egress_protocol               = ["tcp", "udp"]
  whitelisted_urls                 = [".pypi.org", ".pythonhosted.org", ".cran.r-project.org"]
  db_web_app                       = "frankfurt.cloud.databricks.com"
  db_tunnel                        = "tunnel.eu-central-1.cloud.databricks.com"
  db_rds                           = "mdv2llxgl8lou0.ceptxxgorjrc.eu-central-1.rds.amazonaws.com"
  db_control_plane                 = "18.159.44.32/28"
  availability_zones               = ["${var.region}a", "${var.region}b"]
}
```

Before [managing workspace](workspace-management.md), you have to create:
  - [VPC](#vpc)
  - [AWS Firewall](#aws-firewall)
  - [Root bucket](aws-workspace.md#root-bucket)
  - [Cross-account role](aws-workspace.md#cross-account-iam-role)
  - [Databricks E2 workspace](aws-workspace.md#databricks-e2-workspace)
  - [Host and Token outputs](aws-workspace.md#provider-configuration) 

> Initializing provider with `alias = "mws"` and using `provider = databricks.mws` for all `databricks_mws_*` resources. We require all `databricks_mws_*` resources to be created within it's own dedicated terraform module of your environment. Usually this module creates VPC and IAM roles as well.

```hcl
terraform {
  required_providers {
    databricks = {
      source  = "databrickslabs/databricks"
      version = "0.3.9"
    }
    aws = {
      source = "hashicorp/aws"
      version = "3.49.0"
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
  username = var.databricks_account_username
  password = var.databricks_account_password
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
  tags               = var.tags
}

data "databricks_aws_crossaccount_policy" "this" {
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

The very first step is VPC creation with necessary firewall rules. Please consult [main documetation page](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) for **the most complete and up-to-date details on networking**. AWS VPC is registered as [databricks_mws_networks](../resources/mws_networks.md) resource. 

```hcl
data "aws_availability_zones" "available" {}

/* Create VPC */
resource "aws_vpc" "spoke_vpc" {
  cidr_block           = var.spoke_cidr_block
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = merge(var.tags, {
    Name = "${local.prefix}-spoke-vpc"
  })
}

/* Spoke private subnet for dataplane cluster */
resource "aws_subnet" "spoke_db_private_subnet" {
  vpc_id                  = aws_vpc.spoke_vpc.id
  count                   = length(local.spoke_db_private_subnets_cidr)
  cidr_block              = element(local.spoke_db_private_subnets_cidr, count.index)
  availability_zone       = element(local.availability_zones, count.index)
  map_public_ip_on_launch = false
  tags = merge(var.tags, {
    Name = "${local.prefix}-spoke-db-private-${element(local.availability_zones, count.index)}"
  })
}

/* Spoke private subnet for dataplane cluster */
resource "aws_subnet" "spoke_tgw_private_subnet" {
  vpc_id                  = aws_vpc.spoke_vpc.id
  count                   = length(local.spoke_tgw_private_subnets_cidr)
  cidr_block              = element(local.spoke_tgw_private_subnets_cidr, count.index)
  availability_zone       = element(local.availability_zones, count.index)
  map_public_ip_on_launch = false
  tags = merge(var.tags, {
    Name = "${local.prefix}-spoke-tgw-private-${element(local.availability_zones, count.index)}"
  })
}

/* Routing table for spoke private subnet */
resource "aws_route_table" "spoke_db_private_rt" {
  vpc_id = aws_vpc.spoke_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-spoke-db-private-rt"
  })
}

/* Manage the main routing table for VPC  */
resource "aws_main_route_table_association" "spoke-set-worker-default-rt-assoc" {
  vpc_id         = aws_vpc.spoke_vpc.id
  route_table_id = aws_route_table.spoke_db_private_rt.id
}

/* Routing table associations for spoke */
resource "aws_route_table_association" "spoke_db_private_rta" {
  count          = length(local.spoke_db_private_subnets_cidr)
  subnet_id      = element(aws_subnet.spoke_db_private_subnet.*.id, count.index)
  route_table_id = aws_route_table.spoke_db_private_rt.id
}

/* VPC's Default Security Group */
resource "aws_security_group" "default_sg" {
  name        = "${local.prefix}-default-sg"
  description = "Default security group to allow inbound/outbound from the VPC"
  vpc_id      = aws_vpc.db_vpc.id
  depends_on  = [aws_vpc.db_vpc]

  dynamic "ingress" {
    for_each = local.sg_ingress_protocol
    content {
      from_port = 0
      to_port   = 65535
      protocol  = ingress.value
      self      = true
    }
  }

  dynamic "egress" {
    for_each = local.sg_egress_protocol
    content {
      from_port = 0
      to_port   = 65535
      protocol  = egress.value
      self      = true
    }
  }

  dynamic "egress" {
    for_each = local.sg_egress_ports
    content {
      from_port   = egress.value
      to_port     = egress.value
      protocol    = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  tags = {
    Environment = var.environment
  }
}

resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = var.databricks_account_id
  network_name       = "${local.prefix}-network"
  security_group_ids = [aws_security_group.default_spoke_sg.id]
  subnet_ids         = aws_subnet.spoke_db_private_subnet[*].id
  vpc_id             = aws_vpc.spoke_vpc.id
}

// Hub VPC
/* Create VPC */
resource "aws_vpc" "hub_vpc" {
  cidr_block           = var.hub_cidr_block
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = merge(var.tags, {
    Name = "${local.prefix}-hub-vpc"
  })
}

/* Private subnet for Hub TGW Databricks */
resource "aws_subnet" "hub_tgw_private_subnet" {
  vpc_id                  = aws_vpc.hub_vpc.id
  count                   = length(local.hub_tgw_private_subnets_cidr)
  cidr_block              = element(local.hub_tgw_private_subnets_cidr, count.index)
  availability_zone       = element(local.availability_zones, count.index)
  map_public_ip_on_launch = false
  tags = merge(var.tags, {
    Name = "${local.prefix}-hub-tgw-private-${element(local.availability_zones, count.index)}"
  })
}

/* NAT Public subnet */
resource "aws_subnet" "hub_nat_public_subnet" {
  vpc_id                  = aws_vpc.hub_vpc.id
  count                   = length(local.hub_nat_public_subnets_cidr)
  cidr_block              = element(local.hub_nat_public_subnets_cidr, count.index)
  availability_zone       = element(local.availability_zones, count.index)
  map_public_ip_on_launch = true
  tags = merge(var.tags, {
    Name = "${local.prefix}-hub-nat-public-${element(local.availability_zones, count.index)}"
  })
}

/* Firewall subnet */
resource "aws_subnet" "hub_firewall_subnet" {
  vpc_id                  = aws_vpc.hub_vpc.id
  count                   = length(local.hub_firewall_subnets_cidr)
  cidr_block              = element(local.hub_firewall_subnets_cidr, count.index)
  availability_zone       = element(local.availability_zones, count.index)
  map_public_ip_on_launch = false
    tags = merge(var.tags, {
    Name = "${local.prefix}-hub-firewall-public-${element(local.availability_zones, count.index)}"
  })
}

/* Internet gateway for the public subnet */
resource "aws_internet_gateway" "hub_igw" {
  vpc_id = aws_vpc.hub_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-hub-igw"
  })
}

/* Elastic IP for NAT */
resource "aws_eip" "hub_nat_eip" {
  vpc        = true
  depends_on = [aws_internet_gateway.hub_igw]
}

/* Hub NAT Gateway */
resource "aws_nat_gateway" "hub_nat" {
  allocation_id = aws_eip.hub_nat_eip.id
  subnet_id     = element(aws_subnet.hub_nat_public_subnet.*.id, 0)
  depends_on    = [aws_internet_gateway.hub_igw]
  tags = merge(var.tags, {
    Name = "${local.prefix}-hub-nat"
  })
}

/* Routing table for hub private subnet */
resource "aws_route_table" "hub_tgw_private_rt" {
  vpc_id = aws_vpc.hub_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-hub-tgw-private-rt"
  })
}

/* Routing table for hub nat public subnet */
resource "aws_route_table" "hub_nat_public_rt" {
  vpc_id = aws_vpc.hub_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-hub-nat-rt"
  })
}

/* Routing table for spoke nat public subnet */
resource "aws_route_table" "hub_firewall_rt" {
  vpc_id = aws_vpc.hub_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-hub-firewall-rt"
  })
}

/* Routing table for internet gateway */
resource "aws_route_table" "hub_igw_rt" {
  vpc_id = aws_vpc.hub_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-hub-igw-rt"
  })
}

/* Routing table associations for hub tgw */
resource "aws_route_table_association" "hub_tgw_rta" {
  count          = length(local.hub_tgw_private_subnets_cidr)
  subnet_id      = element(aws_subnet.hub_tgw_private_subnet.*.id, count.index)
  route_table_id = aws_route_table.hub_tgw_private_rt.id
}

resource "aws_route_table_association" "hub_nat_rta" {
  count          = length(local.hub_nat_public_subnets_cidr)
  subnet_id      = element(aws_subnet.hub_nat_public_subnet.*.id, count.index)
  route_table_id = aws_route_table.hub_nat_public_rt.id
}

resource "aws_route_table_association" "hub_firewall_rta" {
  count          = length(local.hub_firewall_subnets_cidr)
  subnet_id      = element(aws_subnet.hub_firewall_subnet.*.id, count.index)
  route_table_id = aws_route_table.hub_firewall_rt.id
}

resource "aws_route_table_association" "hub_igw_rta" {
  gateway_id     = aws_internet_gateway.hub_igw.id
  route_table_id = aws_route_table.hub_igw_rt.id
}

/* Adding routes to route tables */
resource "aws_route" "db_private_nat_gtw" {
  route_table_id         = aws_route_table.hub_tgw_private_rt.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.hub_nat.id
}

resource "aws_route" "db_firewall_public_gtw" {
  route_table_id         = aws_route_table.hub_firewall_rt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.hub_igw.id
}

/* Manage the main routing table for VPC  */
resource "aws_main_route_table_association" "set-worker-default-rt-assoc" {
  vpc_id         = aws_vpc.hub_vpc.id
  route_table_id = aws_route_table.hub_firewall_rt.id
}

/* VPC's Default Security Group */
resource "aws_security_group" "default_hub_sg" {
  name        = "${local.prefix}-default-hub-sg"
  description = "Default security group to allow inbound/outbound from the VPC"
  vpc_id      = aws_vpc.hub_vpc.id
  depends_on  = [aws_vpc.hub_vpc]

  ingress {
    from_port = "0"
    to_port   = "0"
    protocol  = "-1"
    self      = true
  }

  egress {
    from_port   = "0"
    to_port     = "0"
    cidr_blocks = ["0.0.0.0/0"]
    protocol    = "-1"
    self        = false
  }
  tags = var.tags
}

/* Create VPC Endpoint */
module "vpc_endpoints" {
  source  = "terraform-aws-modules/vpc/aws//modules/vpc-endpoints"
  version = "3.2.0"

  vpc_id             = aws_vpc.spoke_vpc.id
  security_group_ids = [aws_security_group.default_spoke_sg.id]

  endpoints = {
    s3 = {
      service         = "s3"
      service_type    = "Gateway"
      route_table_ids = flatten([
        aws_route_table.spoke_db_private_rt.id
      ])
      tags            = {
        Name = "${local.prefix}-s3-vpc-endpoint"
      }
    },
    sts = {
      service             = "sts"
      private_dns_enabled = true
      subnet_ids          = aws_subnet.spoke_db_private_subnet[*].id
      tags = {
        Name = "${local.prefix}-sts-vpc-endpoint"
      }
    },
    kinesis-streams = {
      service             = "kinesis-streams"
      private_dns_enabled = true
      subnet_ids          = aws_subnet.spoke_db_private_subnet[*].id
      tags = {
        Name = "${local.prefix}-kinesis-vpc-endpoint"
      }
    },

  }

  tags = var.tags
}
```
## AWS Firewall

Once [VPC](#vpc) is ready, create AWS Firewall for your VPC that restricts outbound http/s traffic to an approved set of Fully Qualified Domain Names (FQDNs).

```hcl
/*Firewall Rule group for accessing hive metastore and public repositories*/
resource "aws_networkfirewall_rule_group" "databricks_fqdns_rg" {
  capacity = 100
  name     = "${local.prefix}-databricks-fqdns-rg"
  type     = "STATEFUL"
  rule_group {
    rules_source {
      rules_source_list {
        generated_rules_type = "ALLOWLIST"
        target_types         = ["TLS_SNI", "HTTP_HOST"]
        targets              = concat([local.db_web_app, local.db_tunnel, local.db_rds], local.whitelisted_urls)
      }
    }
    rule_variables {
      ip_sets {
        key = "HOME_NET"
        ip_set {
          definition = [var.spoke_cidr_block,var.hub_cidr_block]
        }
      }
    }
  }
  tags = var.tags
}

resource "aws_networkfirewall_rule_group" "allow_db_cpl_protocols_rg" {
  capacity    = 100
  description = "Allows control plane traffic traffic from source"
  name        = "${local.prefix}-allow-db-cpl-protocols-rg"
  type        = "STATEFUL"
  rule_group {
    rule_variables {
      ip_sets {
        key = "HOME_NET"
        ip_set {
          definition = [var.spoke_cidr_block,var.hub_cidr_block]
        }
      }
    }
    rules_source {
      dynamic "stateful_rule" {
        for_each = local.protocols_control_plane
        content {
          action = "PASS"
          header {
            destination      = local.db_control_plane
            destination_port = "443"
            protocol         = stateful_rule.value
            direction        = "ANY"
            source_port      = "ANY"
            source           = "ANY"
          }
          rule_option {
            keyword = "sid:${stateful_rule.key + 1}"
          }
        }
      }
    }
  }
  tags = var.tags
}

/* Firewall Rule group for dropping ICMP, FTP, SSH*/
resource "aws_networkfirewall_rule_group" "deny_protocols_rg" {
  capacity    = 100
  description = "Drops FTP,ICMP, SSH traffic from source"
  name        = "${local.prefix}-deny-protocols-rg"
  type        = "STATEFUL"
  rule_group {
    rule_variables {
      ip_sets {
        key = "HOME_NET"
        ip_set {
          definition = [var.spoke_cidr_block,var.hub_cidr_block]
        }
      }
    }
    rules_source {
      dynamic "stateful_rule" {
        for_each = local.protocols
        content {
          action = "DROP"
          header {
            destination      = "ANY"
            destination_port = "ANY"
            protocol         = stateful_rule.value
            direction        = "ANY"
            source_port      = "ANY"
            source           = "ANY"
          }
          rule_option {
            keyword = "sid:${stateful_rule.key + 1}"
          }
        }
      }
    }
  }

  tags = var.tags
}

locals {
  protocols               = ["ICMP", "FTP", "SSH"]
  protocols_control_plane = ["TCP"]
}

/* Firewall Policy */
resource "aws_networkfirewall_firewall_policy" "egress_policy" {
  name = "${local.prefix}-egress-policy"
  firewall_policy {
    stateless_default_actions          = ["aws:forward_to_sfe"]
    stateless_fragment_default_actions = ["aws:forward_to_sfe"]
    stateful_rule_group_reference {
      resource_arn = aws_networkfirewall_rule_group.databricks_fqdns_rg.arn
    }
    stateful_rule_group_reference {
      resource_arn = aws_networkfirewall_rule_group.deny_protocols_rg.arn
    }
    stateful_rule_group_reference {
      resource_arn = aws_networkfirewall_rule_group.allow_db_cpl_protocols_rg.arn
    }
  }
  tags = var.tags
}

/* Create Firewall*/
resource "aws_networkfirewall_firewall" "exfiltration_firewall" {
  name                = "${local.prefix}-fw"
  firewall_policy_arn = aws_networkfirewall_firewall_policy.egress_policy.arn
  vpc_id              = aws_vpc.hub_vpc.id
  dynamic "subnet_mapping" {
    for_each = aws_subnet.hub_firewall_subnet[*].id
    content {
      subnet_id = subnet_mapping.value
    }
  }
  tags = var.tags
}

/* Get Firewall Endpoint*/
data "aws_vpc_endpoint" "firewall" {
  vpc_id = aws_vpc.hub_vpc.id

  tags = {
    "AWSNetworkFirewallManaged" = "true"
    "Firewall"                  = aws_networkfirewall_firewall.exfiltration_firewall.arn
  }

  depends_on = [aws_networkfirewall_firewall.exfiltration_firewall]
}

/* Add Route from Nat Gateway to Firewall */
resource "aws_route" "db_nat_firewall" {
  route_table_id         = aws_route_table.hub_nat_public_rt.id
  destination_cidr_block = "0.0.0.0/0"
  vpc_endpoint_id        = data.aws_vpc_endpoint.firewall.id
}

/* Add Route from Internet Gateway to Firewall */
resource "aws_route" "db_igw_nat_firewall" {
  route_table_id         = aws_route_table.hub_igw_rt.id
  count                  = length(local.hub_nat_public_subnets_cidr)
  destination_cidr_block = element(local.hub_nat_public_subnets_cidr, count.index)
  vpc_endpoint_id        = data.aws_vpc_endpoint.firewall.id
}
```