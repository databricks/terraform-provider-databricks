---
page_title: "Provisioning AWS Databricks workspace with a AWS Firewall"
---

# Provisioning AWS Databricks workspace with a AWS Firewall

-> **Note** Refer to the [Databricks Terraform Registry modules](https://registry.terraform.io/modules/databricks/examples/databricks/latest) for Terraform modules and examples to deploy Azure Databricks resources.

You can provision multiple Databricks workspaces with Terraform. This example shows how to deploy a Databricks workspace into a VPC, which uses an AWS Network firewall to manage egress out to the public network. For smaller Databricks deployments, this is our recommended configuration; for larger deployments, see [Provisioning AWS Databricks workspace with a Hub & Spoke firewall for data exfiltration protection](aws-e2-firewall-hub-and-spoke.md).

For more information, please visit [Data Exfiltration Protection With Databricks on AWS](https://databricks.com/blog/2021/02/02/data-exfiltration-protection-with-databricks-on-aws.html).

![Data Exfiltration_Workspace](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/images/aws-e2-firewall-workspace.png)

## Provider initialization for AWS workspaces

This guide assumes you have the `client_id`, which is the `application_id` of the [Service Principal](resources/service_principal.md), `client_secret`, which is its secret, and `databricks_account_id`, which can be found in the top right corner of the [Account Console](https://accounts.cloud.databricks.com). (see [instruction](https://docs.databricks.com/dev-tools/authentication-oauth.html#step-2-create-an-oauth-secret-for-a-service-principal)). This guide is provided as is and assumes you will use it as the basis for your setup. If you are using AWS Firewall to block most traffic but allow the URLs that Databricks needs to connect to, please update the configuration based on your region. You can get the configuration details for your region from [Firewall Appliance](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html#firewall-appliance-infrastructure) document.

```hcl
variable "client_id" {}
variable "client_secret" {}
variable "databricks_account_id" {}

variable "tags" {
  default = {}
}

variable "cidr_block" {
  default = "10.4.0.0/16"
}

variable "region" {
  default = "eu-west-2"
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

variable "whitelisted_urls" {
  default = [".pypi.org", ".pythonhosted.org", ".cran.r-project.org"]
}

variable "db_web_app" {
  default = "london.cloud.databricks.com"
}

variable "db_tunnel" {
  default = "tunnel.eu-west-2.cloud.databricks.com"
}

variable "db_rds" {
  default = "mdio2468d9025m.c6fvhwk6cqca.eu-west-2.rds.amazonaws.com"
}

variable "db_control_plane" {
  default = "18.134.65.240/28"
}

variable "prefix" {
  default = "demo"
}

locals {
  prefix                       = "${var.prefix}${random_string.naming.result}"
  private_subnets_cidr         = [cidrsubnet(var.cidr_block, 3, 0), cidrsubnet(var.cidr_block, 3, 1)]
  nat_public_subnets_cidr      = [cidrsubnet(var.cidr_block, 3, 2), cidrsubnet(var.cidr_block, 3, 3)]
  firewall_public_subnets_cidr = [cidrsubnet(var.cidr_block, 3, 4)]
  sg_egress_ports              = [443, 3306, 6666]
  sg_ingress_protocol          = ["tcp", "udp"]
  sg_egress_protocol           = ["tcp", "udp"]
  availability_zones           = ["${var.region}a", "${var.region}b"]
  db_root_bucket               = "${var.prefix}${random_string.naming.result}-rootbucket.s3.amazonaws.com"
}
```

Before [managing workspace](workspace-management.md), you have to create:

- [VPC](#vpc)
- [AWS Network Firewall](#aws-network-firewall)
- [Root bucket](aws-workspace.md#root-bucket)
- [Cross-account role](aws-workspace.md#cross-account-iam-role)
- [Databricks workspace](aws-workspace.md#databricks-workspace)
- [Host and Token outputs](aws-workspace.md#provider-configuration)

> Initializing provider with `alias = "mws"` and using `provider = databricks.mws` for all `databricks_mws_*` resources. We require all `databricks_mws_*` resources to be created within its own dedicated terraform module of your environment. Usually, this module creates VPC and IAM roles as well.

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databricks/databricks"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.15.0"
    }
  }
}

provider "aws" {
  region = var.region
}

// initialize provider in "MWS" mode to provision new workspace
provider "databricks" {
  alias         = "mws"
  host          = "https://accounts.cloud.databricks.com"
  account_id    = var.databricks_account_id
  client_id     = var.client_id
  client_secret = var.client_secret
}
```

## VPC

The very first step is VPC creation. Please consult [main documentation page](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) for **the most complete and up-to-date details on networking**. AWS VPC is registered as [databricks_mws_networks](../resources/mws_networks.md) resource.

```hcl
data "aws_availability_zones" "available" {}

/* Create VPC */
resource "aws_vpc" "db_vpc" {
  cidr_block           = var.cidr_block
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-vpc"
  })
}

/* Private subnet for Databricks */
resource "aws_subnet" "db_private_subnet" {
  vpc_id                  = aws_vpc.db_vpc.id
  count                   = length(local.private_subnets_cidr)
  cidr_block              = element(local.private_subnets_cidr, count.index)
  availability_zone       = element(local.availability_zones, count.index)
  map_public_ip_on_launch = false
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-private-${element(local.availability_zones, count.index)}"
  })
}

/* NAT Public subnet */
resource "aws_subnet" "db_nat_public_subnet" {
  vpc_id                  = aws_vpc.db_vpc.id
  count                   = length(local.nat_public_subnets_cidr)
  cidr_block              = element(local.nat_public_subnets_cidr, count.index)
  availability_zone       = element(local.availability_zones, count.index)
  map_public_ip_on_launch = true
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-nat-public-${element(local.availability_zones, count.index)}"
  })
}

/* Firewall Public subnet */
resource "aws_subnet" "db_firewall_subnet" {
  vpc_id                  = aws_vpc.db_vpc.id
  count                   = length(local.firewall_public_subnets_cidr)
  cidr_block              = element(local.firewall_public_subnets_cidr, count.index)
  availability_zone       = element(local.availability_zones, count.index)
  map_public_ip_on_launch = false
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-firewall-public-${element(local.availability_zones, count.index)}"
  })
}

/* Internet gateway for the public subnet */
resource "aws_internet_gateway" "db_igw" {
  vpc_id = aws_vpc.db_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-igw"
  })
}

/* Elastic IP for NAT */
resource "aws_eip" "db_nat_eip" {
  vpc        = true
  depends_on = [aws_internet_gateway.db_igw]
}

/* NAT Gateway */
resource "aws_nat_gateway" "db_nat" {
  allocation_id = aws_eip.db_nat_eip.id
  subnet_id     = element(aws_subnet.db_nat_public_subnet.*.id, 0)
  depends_on    = [aws_internet_gateway.db_igw]
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-nat"
  })
}

```

### Security Group

Databricks must have access to at least one AWS security group and no more than five security groups. You can reuse existing security groups rather than create new ones.
Security groups must have the following rules:

***Egress (outbound):***

- Allow all TCP and UDP access to the workspace security group (for internal traffic)
- Allow TCP access to 0.0.0.0/0 for these ports:
  - 443: for Databricks infrastructure, cloud data sources, and library repositories
  - 3306: for the metastore
  - 6666: only required if you use PrivateLink

***Ingress (inbound):*** Required for all workspaces (these can be separate rules or combined into one):

- Allow TCP on all ports when the traffic source uses the same security group
- Allow UDP on all ports when the traffic source uses the same security group

```hcl
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

  tags = var.tags
}
```

### Register AWS VPC as the databricks_mws_networks resource

Now, we configure VPC & subnets for new workspaces within AWS.

```hcl
resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = var.databricks_account_id
  network_name       = "${local.prefix}-network"
  security_group_ids = [aws_security_group.default_sg.id]
  subnet_ids         = aws_subnet.db_private_subnet[*].id
  vpc_id             = aws_vpc.db_vpc.id
}
```

### Route Tables

Next, we will create route tables for VPC subnets, NAT gateway, and Internet Gateway and add some routes.

```hcl
/* Routing table for private subnet */
resource "aws_route_table" "db_private_rt" {
  vpc_id = aws_vpc.db_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-private-rt"
  })
}

/* Routing table for nat public subnet */
resource "aws_route_table" "db_nat_public_rt" {
  vpc_id = aws_vpc.db_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-nat-rt"
  })
}

/* Routing table for nat public subnet */
resource "aws_route_table" "db_firewall_rt" {
  vpc_id = aws_vpc.db_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-firewall-rt"
  })
}

/* Routing table for internet gateway */
resource "aws_route_table" "db_igw_rt" {
  vpc_id = aws_vpc.db_vpc.id
  tags = merge(var.tags, {
    Name = "${local.prefix}-db-igw-rt"
  })
}

/* Routing table associations */
resource "aws_route_table_association" "db_private" {
  count          = length(local.private_subnets_cidr)
  subnet_id      = element(aws_subnet.db_private_subnet.*.id, count.index)
  route_table_id = aws_route_table.db_private_rt.id
}

resource "aws_route_table_association" "db_nat" {
  count          = length(local.nat_public_subnets_cidr)
  subnet_id      = element(aws_subnet.db_nat_public_subnet.*.id, count.index)
  route_table_id = aws_route_table.db_nat_public_rt.id
}

resource "aws_route_table_association" "db_firewall" {
  count          = length(local.firewall_public_subnets_cidr)
  subnet_id      = element(aws_subnet.db_firewall_subnet.*.id, count.index)
  route_table_id = aws_route_table.db_firewall_rt.id
}

resource "aws_route_table_association" "db_igw" {
  gateway_id     = aws_internet_gateway.db_igw.id
  route_table_id = aws_route_table.db_igw_rt.id
}

/* Adding routes to route tables */
resource "aws_route" "db_private_nat_gtw" {
  route_table_id         = aws_route_table.db_private_rt.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.db_nat.id
}

resource "aws_route" "db_firewall_gtw" {
  route_table_id         = aws_route_table.db_firewall_rt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.db_igw.id
}

/* Manage the main routing table for VPC  */
resource "aws_main_route_table_association" "set-worker-default-rt-assoc" {
  vpc_id         = aws_vpc.db_vpc.id
  route_table_id = aws_route_table.db_firewall_rt.id
}
```

### VPC Endpoints

For STS, S3, and Kinesis, it's important to create VPC gateway or interface endpoints such that the relevant in-region traffic from clusters could transit over the secure AWS backbone rather than the public network for more direct connections and reduced cost compared to AWS global endpoints.

```hcl
module "vpc_endpoints" {
  source  = "terraform-aws-modules/vpc/aws//modules/vpc-endpoints"
  version = "3.11.0"

  vpc_id             = aws_vpc.db_vpc.id
  security_group_ids = [aws_security_group.default_sg.id]

  endpoints = {
    s3 = {
      service      = "s3"
      service_type = "Gateway"
      route_table_ids = flatten([
        aws_route_table.db_private_rt.id
      ])
      tags = {
        Name = "${local.prefix}-s3-vpc-endpoint"
      }
    },
    sts = {
      service             = "sts"
      private_dns_enabled = true
      subnet_ids          = aws_subnet.db_private_subnet[*].id
      tags = {
        Name = "${local.prefix}-sts-vpc-endpoint"
      }
    },
    kinesis-streams = {
      service             = "kinesis-streams"
      private_dns_enabled = true
      subnet_ids          = aws_subnet.db_private_subnet[*].id
      tags = {
        Name = "${local.prefix}-kinesis-vpc-endpoint"
      }
    },

  }

  tags = var.tags
}

```

## AWS Network Firewall

Once [VPC](#vpc) is ready, create an AWS Network Firewall for your VPC that restricts outbound http/s traffic to an approved set of Fully Qualified Domain Names (FQDNs).

### AWS Firewall Rule Groups

First, we will create a Firewall Rule group for accessing hive metastore and public repositories.

```hcl
resource "aws_networkfirewall_rule_group" "databricks_fqdns_rg" {
  capacity = 100
  name     = "${local.prefix}-databricks-fqdns-rg"
  type     = "STATEFUL"
  rule_group {
    rules_source {
      rules_source_list {
        generated_rules_type = "ALLOWLIST"
        target_types         = ["TLS_SNI", "HTTP_HOST"]
        targets              = concat([var.db_web_app, var.db_tunnel, var.db_rds, local.db_root_bucket], var.whitelisted_urls)
      }
    }
    rule_variables {
      ip_sets {
        key = "HOME_NET"
        ip_set {
          definition = [var.cidr_block]
        }
      }
    }
  }
  tags = var.tags
}
```

As the next step, we will create a Firewall Rule group that allows control plane traffic from the VPC.

```hcl
resource "aws_networkfirewall_rule_group" "allow_db_cpl_protocols_rg" {
  capacity    = 100
  description = "Allows control plane traffic from source"
  name        = "${local.prefix}-allow-db-cpl-protocols-rg"
  type        = "STATEFUL"
  rule_group {
    rule_variables {
      ip_sets {
        key = "HOME_NET"
        ip_set {
          definition = [var.cidr_block]
        }
      }
    }
    rules_source {
      dynamic "stateful_rule" {
        for_each = local.protocols_control_plane
        content {
          action = "PASS"
          header {
            destination      = var.db_control_plane
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
locals {
  protocols               = ["ICMP", "FTP", "SSH"]
  protocols_control_plane = ["TCP"]
}
```

Finally, we will add some basic deny rules to cater for common firewall scenarios, such as preventing the use of protocols like SSH/SFTP, FTP, and ICMP.

```hcl
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
          definition = [var.cidr_block]
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
```

### AWS Network Firewall Policy

First, we will create an AWS Firewall Policy and include stateful firewall rule groups created in previous steps.

```hcl
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
```

### AWS Firewall

As the next step, we will create an AWS Network Firewall with the Firewall Policy we defined in the previous step.

```hcl
resource "aws_networkfirewall_firewall" "exfiltration_firewall" {
  name                = "${local.prefix}-fw"
  firewall_policy_arn = aws_networkfirewall_firewall_policy.egress_policy.arn
  vpc_id              = aws_vpc.db_vpc.id
  dynamic "subnet_mapping" {
    for_each = aws_subnet.db_firewall_subnet[*].id
    content {
      subnet_id = subnet_mapping.value
    }
  }
  tags = var.tags
}

/* Get Firewall Endpoint*/
data "aws_vpc_endpoint" "firewall" {
  vpc_id = aws_vpc.db_vpc.id

  tags = {
    "AWSNetworkFirewallManaged" = "true"
    "Firewall"                  = aws_networkfirewall_firewall.exfiltration_firewall.arn
  }

  depends_on = [aws_networkfirewall_firewall.exfiltration_firewall]
}

```

Finally, the AWS Network Firewall is now deployed and configured - all you need to do now is route traffic to it.

```hcl
/* Add Route from Nat Gateway to Firewall */
resource "aws_route" "db_nat_firewall" {
  route_table_id         = aws_route_table.db_nat_public_rt.id
  destination_cidr_block = "0.0.0.0/0"
  vpc_endpoint_id        = data.aws_vpc_endpoint.firewall.id
}

/* Add Route from Internet Gateway to Firewall */
resource "aws_route" "db_igw_nat_firewall" {
  route_table_id         = aws_route_table.db_igw_rt.id
  count                  = length(local.nat_public_subnets_cidr)
  destination_cidr_block = element(local.nat_public_subnets_cidr, count.index)
  vpc_endpoint_id        = data.aws_vpc_endpoint.firewall.id
}
```

## Troubleshooting

If the Databricks clusters cannot reach DBFS or VPC endpoints do not work as intended, for example if your data sources are inaccessible or if the traffic is bypassing the endpoints please visit [Troubleshoot regional endpoints](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html#troubleshoot-regional-endpoints)
