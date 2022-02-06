---
page_title: "Enable AWS PrivateLink for Databricks Workspace"
---

# Deploying pre-requisite resources and enabling PrivateLink connections (AWS Preview)

-> **Private Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Databricks PrivateLink support enables private connectivity between users and their Databricks workspaces and between clusters on the data plane and core services on the control plane within the Databricks workspace infrastructure. You can use Terraform to deploy the underlying cloud resources and the private access settings resources automatically, using a programmatic approach. This guide assumes you are deploying into an existing VPC and you have set up credentials and storage configurations as per prior examples, notably here.

This guide uses the following variables in configurations:

- `databricks_account_username`: The username an account-level admin uses to log in to  [https://accounts.cloud.databricks.com](https://accounts.cloud.databricks.com).
- `databricks_account_password`: The password for `databricks_account_username`.
- `databricks_account_id`: The numeric ID for your Databricks account. When you are logged in, it appears in the bottom left corner of the page.
- `vpc_id` - The ID for the AWS VPC 
- `network_name` - Name for your Databricks-configured network
- `region` - AWS region
- `existing_network_sg` - Security groups set up for the existing VPC
- `existing_network_subnets` - Existing subnets being used for the customer managed VPC
- `workspace_vpce_service` - Choose the region-specific service endpoint from this table.
- `relay_vpce_service` - Choose the region-specific service from this table.
- `vpc_cidr_block` - CIDR range for the VPC being deployed into
- `vpce_cidr` - CIDR range for the subnet chosen for the VPC endpoint
- `credentials_id` - Databricks workspace credential ID 
- `storage_configuration_id` - Databricks workspace storage configuration ID

This guide is provided as-is and you can use this guide as the basis for your custom Terraform module.

To get started with AWS PrivateLink integration, this guide takes you throw the following high-level steps:
- Initialize the required providers
- Configure AWS objects
  - A subnet dedicated to your VPC relay and workspace endpoints
  - A security group dedicated to your VPC endpoints
  - Two AWS VPC endpoints
- Workspace Creation

## Provider initialization

Initialize [provider with `mws` alias](https://www.terraform.io/language/providers/configuration#alias-multiple-provider-configurations) to set up account-level resources. See [provider authentication](../index.md#authenticating-with-hostname,-username,-and-password) for more details.

```hcl
terraform {
 required_providers {
   databricks = {
     source  = "databrickslabs/databricks"
     version = "0.4.7"
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

// initialize provider in "MWS" mode for provisioning workspace with AWS PrivateLink
provider "databricks" {
 alias    = "mws"
 host     = "https://accounts.cloud.databricks.com"
 username = var.databricks_account_username
 password = var.databricks_account_password
}
```

Define the required variables

```hcl
variable databricks_account_id {}
variable databricks_account_username {}
variable databricks_account_password {}
variable vpc_id {}
variable network_name {}
variable region {}
variable existing_network_sg {}
variable existing_network_subnets {}

variable workspace_vpce_service {}
variable relay_vpce_service {}

variable vpc_cidr_block {}
variable vpce_cidr {}

// Use the Databricks Account API 2.0 to retrieve these two IDs below - https://docs.databricks.com/dev-tools/api/latest/account.html
variable credentials_id {}
variable storage_configuration_id {}

locals {
 prefix = "private-link-ws"
}

locals {
workspace_subnet_1 = (
 element(split(",", var.existing_network_subnets), 1)
)
workspace_subnet_2 = (
 element(split(",", var.existing_network_subnets), 2)
)
}
```

## Configure AWS objects
The first step is to create the required AWS objects:
- A subnet dedicated to your VPC endpoints
- A security group dedicated to your VPC endpoints and satisfying required inbound/outbound TCP/HTTPS traffic rules on ports 443 and 6666, respectively.
- Lastly, creation of the private access settings and workspace.

```hcl
resource "aws_subnet" "vpce" {
  vpc_id     = var.vpc_id
  cidr_block = var.vpce_cidr

  tags = {
    Name = "vpce subnet for workspace"
  }
}

resource "aws_route_table" "my_vpc_private" {
    vpc_id = var.vpc_id

    tags = {
        Name = "Local Route Table for Isolated Private Subnet"
    }
}

resource "aws_route_table_association" "my_vpc_us_east_1a_private" {
    subnet_id = aws_subnet.vpce.id
    route_table_id = aws_route_table.my_vpc_private.id
}
```

```hcl
resource "aws_security_group" "vpce_sg" {
 name        = "VPC endpoint security group"
 description = "Security group shared with relay and workspace endpoints"
 vpc_id      = var.vpc_id

 ingress {
   description      = "Inbound rules"
   from_port        = 443
   to_port          = 443
   protocol         = "tcp"
   cidr_blocks      = [var.vpc_cidr_block]
 }

 ingress {
   description      = "Outbound rules"
   from_port        = 6666
   to_port          = 6666
   protocol         = "tcp"
   cidr_blocks      = [var.vpc_cidr_block]
 }

 egress {
   from_port        = 0
   to_port          = 0
   protocol         = "-1"
   cidr_blocks      = ["0.0.0.0/0"]
   ipv6_cidr_blocks = ["::/0"]
 }

 tags = {
   Name = "vpce_rules"
 }
}
```

```hcl
resource "aws_vpc_endpoint" "workspace" {
 vpc_id             = var.vpc_id
 service_name       = var.workspace_vpce_service
 vpc_endpoint_type  = "Interface"
 security_group_ids = [aws_security_group.vpce_sg.id]
 subnet_ids         = [aws_subnet.vpce.id]
 // run terraform apply twice when configuring PrivateLink.
 // Run 1 - comment the `private_dns_enabled` line
 // Run 2 - uncomment the `private_dns_enabled` line
 // private_dns_enabled = true
 depends_on         = [aws_subnet.vpce]
}

resource "aws_vpc_endpoint" "relay" {
 vpc_id             = var.vpc_id
 service_name       = var.relay_vpce_service
 vpc_endpoint_type  = "Interface"
 security_group_ids = [aws_security_group.vpce_sg.id]
 subnet_ids         = [aws_subnet.vpce.id]
   // run terraform apply twice when configuring PrivateLink.
 // Run 1 - comment the `private_dns_enabled` line
 // Run 2 - uncomment the `private_dns_enabled` line
 // private_dns_enabled = true
 depends_on         = [aws_subnet.vpce]
}

resource "databricks_mws_vpc_endpoint" "workspace" {
 provider            = databricks.mws
 account_id          = var.databricks_account_id
 aws_vpc_endpoint_id = aws_vpc_endpoint.workspace.id
 vpc_endpoint_name   = "VPC Relay for ${var.vpc_id}"
 region              = var.region
 depends_on          = [aws_vpc_endpoint.workspace]
}

resource "databricks_mws_vpc_endpoint" "relay" {
 provider            = databricks.mws
 account_id          = var.databricks_account_id
 aws_vpc_endpoint_id = aws_vpc_endpoint.relay.id
 vpc_endpoint_name   = "VPC Relay for ${var.vpc_id}"
 region              = var.region
 depends_on          = [aws_vpc_endpoint.relay]
}
```

## Workspace creation

Once the VPC endpoints are created, they can be supplied in the `databricks_mws_networks` resource for workspace creation with AWS PrivateLink. After the terraform apply is run once (see the comment in the aws_vpc_endpoint resource above), run the terraform apply a second time with the line for private_dns_enabled set to true uncommented to set the proper DNS settings for PrivateLink.

```hcl
// Inputs are 2 subnets and one security group from existing VPC that will be used for your Databricks workspace
resource "databricks_mws_networks" "this" {
 provider           = databricks.mws
 account_id         = var.databricks_account_id
 network_name       = var.network_name
 security_group_ids = [var.existing_network_sg]
 subnet_ids         = [local.workspace_subnet_1, local.workspace_subnet_2]
 vpc_id             = var.vpc_id
 vpc_endpoints {
   dataplane_relay = [databricks_mws_vpc_endpoint.relay.vpc_endpoint_id]
   rest_api        = [databricks_mws_vpc_endpoint.workspace.vpc_endpoint_id]
 }
 depends_on = [aws_vpc_endpoint.workspace, aws_vpc_endpoint.relay]
}

resource "databricks_mws_private_access_settings" "pas" {
 provider                     = databricks.mws
 account_id                   = var.databricks_account_id
 private_access_settings_name = "Private Access Settings for ${local.prefix}"
 region                       = var.region
 public_access_enabled        = true
}

resource "databricks_mws_workspaces" "this" {
 provider                   = databricks.mws
 account_id                 = var.databricks_account_id
 aws_region                 = var.region
 workspace_name             = local.prefix
 deployment_name            = local.prefix
 credentials_id             = var.credentials_id
 storage_configuration_id   = var.storage_configuration_id
 network_id                 = databricks_mws_networks.this.network_id
 private_access_settings_id = databricks_mws_private_access_settings.pas.private_access_settings_id
 pricing_tier               = "ENTERPRISE"
 depends_on                 = [databricks_mws_networks.this]
}
```

