---
page_title: "Provisioning AWS Databricks workspace with a Hub & Spoke firewall for data exfiltration protection"
---

# Provisioning AWS Databricks workspace with a Hub & Spoke firewall for data exfiltration protection

-> **Note** Refer to the [Databricks Terraform Registry modules](https://registry.terraform.io/modules/databricks/examples/databricks/latest) for Terraform modules and examples to deploy Azure Databricks resources.

You can provision multiple Databricks workspaces with Terraform, and where many Databricks workspaces are deployed, we recommend a hub and spoke topology reference architecture powered by AWS Transit Gateway. The hub will consist of a central inspection and egress virtual private cloud (VPC), while the Spoke VPC houses federated Databricks workspaces for different business units or segregated teams. In this way, you create your version of a centralized deployment model for your egress architecture, as is recommended for large enterprises. For more information, please visit [Data Exfiltration Protection With Databricks on AWS](https://databricks.com/blog/2021/02/02/data-exfiltration-protection-with-databricks-on-aws.html).

![Data Exfiltration](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/images/aws-exfiltration-replace-1.png)

## Provider initialization for AWS workspaces

This guide assumes you have the `client_id`, which is the `application_id` of the [Service Principal](resources/service_principal.md), `client_secret`, which is its secret, and `databricks_account_id`, which can be found in the top right corner of the [Account Console](https://accounts.cloud.databricks.com). (see [instruction](https://docs.databricks.com/dev-tools/authentication-oauth.html#step-2-create-an-oauth-secret-for-a-service-principal)). This guide is provided as is and assumes you will use it as the basis for your setup. If you use AWS Firewall to block most traffic but allow the URLs to which Databricks needs to connect, please update the configuration based on your region. You can get the configuration details for your region from [Firewall Appliance](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html#firewall-appliance-infrastructure) document.

```hcl
variable "client_id" {}
variable "client_secret" {}
variable "databricks_account_id" {}

variable "tags" {
  default = {}
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
variable "whitelisted_urls" {
  default = [".pypi.org", ".pythonhosted.org", ".cran.r-project.org"]
}

variable "db_web_app" {
  default = "frankfurt.cloud.databricks.com"
}

variable "db_tunnel" {
  default = "tunnel.eu-central-1.cloud.databricks.com"
}

variable "db_rds" {
  default = "mdv2llxgl8lou0.ceptxxgorjrc.eu-central-1.rds.amazonaws.com"
}

variable "db_control_plane" {
  default = "18.159.44.32/28"
}

variable "prefix" {
  default = "demo"
}

locals {
  prefix                         = "${var.prefix}${random_string.naming.result}"
  spoke_db_private_subnets_cidr  = [cidrsubnet(var.spoke_cidr_block, 3, 0), cidrsubnet(var.spoke_cidr_block, 3, 1)]
  spoke_tgw_private_subnets_cidr = [cidrsubnet(var.spoke_cidr_block, 3, 2), cidrsubnet(var.spoke_cidr_block, 3, 3)]
  hub_tgw_private_subnets_cidr   = [cidrsubnet(var.hub_cidr_block, 3, 0)]
  hub_nat_public_subnets_cidr    = [cidrsubnet(var.hub_cidr_block, 3, 1)]
  hub_firewall_subnets_cidr      = [cidrsubnet(var.hub_cidr_block, 3, 2)]
  sg_egress_ports                = [443, 3306, 6666]
  sg_ingress_protocol            = ["tcp", "udp"]
  sg_egress_protocol             = ["tcp", "udp"]
  availability_zones             = ["${var.region}a", "${var.region}b"]
  db_root_bucket                 = "${var.prefix}${random_string.naming.result}-rootbucket.s3.amazonaws.com"
}
```

Before [managing workspace](workspace-management.md), you have to create:

- [VPC](#vpc)
- [AWS Transit Gateway](#aws-transit-gateway)
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

The very first step is Hub & Spoke VPC creation. Please consult [main documentation page](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) for **the most complete and up-to-date details on networking**. AWS VPC is registered as [databricks_mws_networks](../resources/mws_networks.md) resource.

### Spoke VPC for Databricks Workspace

The first step is to create a Spoke VPC, which houses federated Databricks workspaces for different business units or segregated teams.

![SpokeVPC](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/images/aws-e2-firewall-spoke-vpc.png)

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
```

### Security Group for Spoke

Databricks must have access to at least one AWS security group and no more than five security groups. You can reuse existing security groups rather than create new ones.
Security groups must have the following rules:

***Egress (outbound):***

- Allow all TCP and UDP access to the workspace security group (for internal traffic)
- Allow TCP access to 0.0.0.0/0 for these ports:
  - 443: for Databricks infrastructure, cloud data sources, and library repositories
  - 3306: for the metastore
  - 6666: only required if you use PrivateLink

***Ingress (inbound):***:

- Allow TCP on all ports when the traffic source uses the same security group
- Allow UDP on all ports when the traffic source uses the same security group

```hcl
/* VPC's Default Security Group */
resource "aws_security_group" "default_spoke_sg" {
  name        = "${local.prefix}-default_spoke_sg"
  description = "Default security group to allow inbound/outbound from the VPC"
  vpc_id      = aws_vpc.spoke_vpc.id
  depends_on  = [aws_vpc.spoke_vpc]

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
  security_group_ids = [aws_security_group.default_spoke_sg.id]
  subnet_ids         = aws_subnet.spoke_db_private_subnet[*].id
  vpc_id             = aws_vpc.spoke_vpc.id
}
```

### VPC Endpoint for Spoke VPC

For STS, S3, and Kinesis, it's important to create a VPC gateway or interface endpoints such that the relevant in-region traffic from clusters could transit over the secure AWS backbone rather than the public network for more direct connections and reduced cost compared to AWS global endpoints.

```hcl
/* Create VPC Endpoint */
module "vpc_endpoints" {
  source  = "terraform-aws-modules/vpc/aws//modules/vpc-endpoints"
  version = "3.11.0"

  vpc_id             = aws_vpc.spoke_vpc.id
  security_group_ids = [aws_security_group.default_spoke_sg.id]

  endpoints = {
    s3 = {
      service      = "s3"
      service_type = "Gateway"
      route_table_ids = flatten([
        aws_route_table.spoke_db_private_rt.id
      ])
      tags = {
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

### Hub VPC

The hub will consist of a central inspection and egress virtual private cloud (VPC). We're going to create a central inspection/egress VPC, which, once we’ve finished, should look like this:

![HubVPC](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/images/aws-e2-firewall-hub-vpc.png)

```hcl
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
```

### Route Tables for Hub

Next, we will create route tables for Hub VPC subnets, NAT gateway, and Internet Gateway and add some routes.

```hcl
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
```

## AWS Transit Gateway

Now that our spoke and inspection/egress VPCs are ready to go, all you need to do is link them all together, and AWS Transit Gateway is the perfect solution for that.
First, we will create a Transit Gateway and link our Databricks data plane via TGW subnets.
All of the logic that determines what routes are going via a Transit Gateway is encapsulated within Transit Gateway Route Tables. We will create some TGW route tables for our Hub & Spoke networks.

![TransitGateway](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/images/aws-e2-firewall-tgw.png)

```hcl
//Create transit gateway
resource "aws_ec2_transit_gateway" "tgw" {
  description                     = "Transit Gateway for Hub/Spoke"
  auto_accept_shared_attachments  = "enable"
  default_route_table_association = "enable"
  default_route_table_propagation = "enable"
  tags = merge(var.tags, {
    Name = "${local.prefix}-tgw"
  })
}

# Attach Transit VPC to Transit Gateway
resource "aws_ec2_transit_gateway_vpc_attachment" "hub" {
  subnet_ids         = aws_subnet.hub_tgw_private_subnet[*].id
  transit_gateway_id = aws_ec2_transit_gateway.tgw.id
  vpc_id             = aws_vpc.hub_vpc.id
  dns_support        = "enable"

  transit_gateway_default_route_table_association = true
  transit_gateway_default_route_table_propagation = true
  tags = merge(var.tags, {
    Name    = "${local.prefix}-hub"
    Purpose = "Transit Gateway Attachment - Hub VPC"
  })
}

//# Attach Spoke VPC to Transit Gateway
resource "aws_ec2_transit_gateway_vpc_attachment" "spoke" {
  subnet_ids         = aws_subnet.spoke_tgw_private_subnet[*].id
  transit_gateway_id = aws_ec2_transit_gateway.tgw.id
  vpc_id             = aws_vpc.spoke_vpc.id
  dns_support        = "enable"

  transit_gateway_default_route_table_association = true
  transit_gateway_default_route_table_propagation = true
  tags = merge(var.tags, {
    Name    = "${local.prefix}-spoke"
    Purpose = "Transit Gateway Attachment - Spoke VPC"
  })
}

```

### Route Table Configurations for Transit Gateway

The Transit Gateway should be set up and ready to go. Now, all that needs to be done is update the route tables in each subnet so traffic flows through it.

```hcl
# Create Route to Internet
resource "aws_ec2_transit_gateway_route" "spoke_to_hub" {
  destination_cidr_block         = "0.0.0.0/0"
  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.hub.id
  transit_gateway_route_table_id = aws_ec2_transit_gateway.tgw.association_default_route_table_id
}

# Add route for spoke db to tgw
resource "aws_route" "spoke_db_to_tgw" {
  route_table_id         = aws_route_table.spoke_db_private_rt.id
  destination_cidr_block = "0.0.0.0/0"
  transit_gateway_id     = aws_ec2_transit_gateway.tgw.id
  depends_on             = [aws_vpc.spoke_vpc, aws_vpc.hub_vpc]
}

resource "aws_route" "hub_tgw_private_subnet_to_tgw" {
  route_table_id         = aws_route_table.hub_tgw_private_rt.id
  destination_cidr_block = var.spoke_cidr_block
  transit_gateway_id     = aws_ec2_transit_gateway.tgw.id
  depends_on             = [aws_vpc.spoke_vpc, aws_vpc.hub_vpc]
}

resource "aws_route" "hub_nat_to_tgw" {
  route_table_id         = aws_route_table.hub_nat_public_rt.id
  destination_cidr_block = var.spoke_cidr_block
  transit_gateway_id     = aws_ec2_transit_gateway.tgw.id
  depends_on             = [aws_vpc.spoke_vpc, aws_vpc.hub_vpc]
}
```

## AWS Network Firewall

Once [VPC](#vpc) is ready, we're going to create AWS Network Firewall for your VPC that restricts outbound http/s traffic to an approved set of Fully Qualified Domain Names (FQDNs).

![AWS Network Firewall](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/images/aws-e2-firewall-config.png)

### AWS Firewall Rule Groups

First, we will create a Firewall Rule group for accessing hive metastore and public repositories.

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
        targets              = concat([var.db_web_app, var.db_tunnel, var.db_rds, local.db_root_bucket], var.whitelisted_urls)
      }
    }
    rule_variables {
      ip_sets {
        key = "HOME_NET"
        ip_set {
          definition = [var.spoke_cidr_block, var.hub_cidr_block]
        }
      }
    }
  }
  tags = var.tags
}

```

As the next step, we will create a Firewall Rule group that allows control plane traffic from the VPC.

```hcl
locals {
  protocols               = ["ICMP", "FTP", "SSH"]
  protocols_control_plane = ["TCP"]
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
          definition = [var.spoke_cidr_block, var.hub_cidr_block]
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

```

Next, we will create basic deny rules to cater for common firewall scenarios, such as preventing the use of protocols like SSH/SFTP, FTP, and ICMP.

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
          definition = [var.spoke_cidr_block, var.hub_cidr_block]
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

Now, we can create an AWS Firewall Policy and include stateful firewall rule groups created in previous steps.

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

The next step is to create an AWS Network Firewall with the Firewall Policy we defined in the previous step.

```hcl
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

```

Finally, the AWS Network Firewall is now deployed and configured; all you need to do now is route traffic to it.

```hcl
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

## Troubleshooting

If the Databricks clusters cannot reach DBFS or VPC endpoints do not work as intended, for example if your data sources are inaccessible or if the traffic is bypassing the endpoints please visit [Troubleshoot regional endpoints](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html#troubleshoot-regional-endpoints)
