---
page_title: "Enable Backend AWS PrivateLink for Databricks Workspace"
---

# Deploying pre-requisite resources and enabling PrivateLink connections (AWS Preview)

-> **Private Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Databricks PrivateLink support enables private connectivity between users and their Databricks workspaces and between clusters on the data plane and core services on the control plane within the Databricks workspace infrastructure. You can use Terraform to deploy the underlying cloud resources and the private access settings resources automatically, using a programmatic approach. This guide assumes you are deploying into an existing VPC and you have set up credentials and storage configurations as per prior examples, notably here.

![Private link backend](https://raw.githubusercontent.com/databrickslabs/terraform-provider-databricks/master/docs/images/aws-e2-private-link-backend.png)

This guide uses the following variables in configurations:

- `databricks_account_username`: The username an account-level admin uses to log in to  [https://accounts.cloud.databricks.com](https://accounts.cloud.databricks.com).
- `databricks_account_password`: The password for `databricks_account_username`.
- `databricks_account_id`: The numeric ID for your Databricks account. When you are logged in, it appears in the bottom left corner of the page.
- `vpc_id` - The ID for the AWS VPC 
- `region` - AWS region
- `security_group_id` - Security groups set up for the existing VPC
- `subnet_ids` - Existing subnets being used for the customer managed VPC
- `workspace_vpce_service` - Choose the region-specific service endpoint from this table.
- `relay_vpce_service` - Choose the region-specific service from this table.
- `vpce_subnet_cidr` - CIDR range for the subnet chosen for the VPC endpoint
- `credentials_id` - Databricks workspace credential ID - The credentials ID is one of the attributes which is created as a result of configuring the cross-account IAM role, which Databricks uses to orchestrate EC2 resources. The credentials are created via [databricks_mws_credentials](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/mws_credentials). 
- `storage_configuration_id` - Databricks workspace storage configuration ID - this is obtained from the [databricks_mws_storage_configurations](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/mws_storage_configurations) resource
- `tags` - tags for the Private Link backend setup
- `root_bucket_name` - AWS bucket name required for [storage mws resource](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/mws_storage_configurations) reference
- `cross_account_arn` - AWS EC2 role ARN required for [credentials mws resource](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/mws_credentials)

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
      version = "0.5.0"
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
variable "databricks_account_id" {}
variable "databricks_account_username" {}
variable "databricks_account_password" {}
variable "vpc_id" {}
variable "region" {}
variable "security_group_id" {}

// this input variable is of array type
variable "subnet_ids" {
  type = list(string)
}

variable "workspace_vpce_service" {}
variable "relay_vpce_service" {}
variable "vpce_subnet_cidr" {}

variable "private_dns_enabled" { default = false}
variable "tags" { default = {}}

// these resources (bucket and IAM role) are assumed created using your AWS provider and the examples here https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/mws_storage_configurations and https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/mws_credentials, respectively.
variable "root_bucket_name" {}
variable "cross_account_arn" {}

locals {
  prefix = "private-link-ws"
}
```

## Existing Storage Objects 
The following object is used in order to reference the storage configuration ID. 
```hcl
resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = var.databricks_account_id
  bucket_name                = var.root_bucket_name
  storage_configuration_name = "${local.prefix}-storage}"
}
```

## Existing IAM Role 
The following object is used in order to reference the credential configuration ID. 
```hcl
resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = var.databricks_account_id
  role_arn         = var.cross_account_arn
  credentials_name = "${local.prefix}-credentials"
}

```


## Configure AWS objects
The first step is to create the required AWS objects:
- A subnet dedicated to your VPC endpoints
- A security group dedicated to your VPC endpoints and satisfying required inbound/outbound TCP/HTTPS traffic rules on ports 443 and 6666, respectively.
- Lastly, creation of the private access settings and workspace.

```hcl
// this subnet houses the data plane VPC endpoints
resource "aws_subnet" "dataplane_vpce" {
  vpc_id     = var.vpc_id
  cidr_block = var.vpce_subnet_cidr

  tags = merge(
    data.aws_vpc.prod.tags,
    {
      Name = "${local.prefix}-${data.aws_vpc.prod.id}-pl-vpce"
    },
  )
}

resource "aws_route_table" "this" {
    vpc_id = var.vpc_id

    tags = merge(
    data.aws_vpc.prod.tags,
    {
      Name = "${local.prefix}-${data.aws_vpc.prod.id}-pl-local-route-tbl"
    },
  )
}

resource "aws_route_table_association" "dataplane_vpce_rtb" {
    subnet_id = aws_subnet.dataplane_vpce.id
    route_table_id = aws_route_table.this.id
}
```

```hcl
data "aws_subnet" "ws_vpc_subnets" {
  for_each = toset(var.subnet_ids)
  id       = each.value
}

locals {
  vpc_cidr_blocks = [
    for subnet in data.aws_subnet.ws_vpc_subnets :
      subnet.cidr_block
    ]
}

// security group for data plane VPC endpoints for backend/relay connections
resource "aws_security_group" "dataplane_vpce" {
  name        = "Data Plane VPC endpoint security group"
  description = "Security group shared with relay and workspace endpoints"
  vpc_id      = var.vpc_id

  ingress {
    description      = "Inbound rules"
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = concat([var.vpce_subnet_cidr], local.vpc_cidr_blocks)
  }

  ingress {
    description      = "Inbound rules"
    from_port        = 6666
    to_port          = 6666
    protocol         = "tcp"
    cidr_blocks      = concat([var.vpce_subnet_cidr], local.vpc_cidr_blocks)
  }

  egress {
    description      = "Outbound rules"
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = concat([var.vpce_subnet_cidr], local.vpc_cidr_blocks)
  }

  egress {
    description      = "Outbound rules"
    from_port        = 6666
    to_port          = 6666
    protocol         = "tcp"
    cidr_blocks      = concat([var.vpce_subnet_cidr], local.vpc_cidr_blocks)
  }

  tags = merge(
    data.aws_vpc.prod.tags,
    {
      Name = "${local.prefix}-${data.aws_vpc.prod.id}-pl-vpce-sg-rules"
    },
  )
}
```

```hcl
data "aws_vpc" "prod" {
  id    = var.vpc_id
}

resource "aws_vpc_endpoint" "backend_rest" {
  vpc_id             = var.vpc_id
  service_name       = var.workspace_vpce_service
  vpc_endpoint_type  = "Interface"
  security_group_ids = [aws_security_group.dataplane_vpce.id]
  subnet_ids         = [aws_subnet.dataplane_vpce.id]
  // run terraform apply twice when configuring PrivateLink - see this outstanding issue for understanding why this is required - https://github.com/hashicorp/terraform-provider-aws/issues/7148
  // Run 1 - comment the `private_dns_enabled` line
  // Run 2 - uncomment the `private_dns_enabled` line
  // private_dns_enabled = var.private_dns_enabled
  depends_on         = [aws_subnet.dataplane_vpce]
}

resource "aws_vpc_endpoint" "relay" {
  vpc_id             = var.vpc_id
  service_name       = var.relay_vpce_service
  vpc_endpoint_type  = "Interface"
  security_group_ids = [aws_security_group.dataplane_vpce.id]
  subnet_ids         = [aws_subnet.dataplane_vpce.id]
  // run terraform apply twice when configuring PrivateLink - see this outstanding issue for understanding why this is required - https://github.com/hashicorp/terraform-provider-aws/issues/7148
  // Run 1 - comment the `private_dns_enabled` line
  // Run 2 - uncomment the `private_dns_enabled` line
  // private_dns_enabled = var.private_dns_enabled
  depends_on         = [aws_subnet.dataplane_vpce]
}


resource "databricks_mws_vpc_endpoint" "backend_rest_vpce" {
  provider            = databricks.mws
  account_id          = var.databricks_account_id
  aws_vpc_endpoint_id = aws_vpc_endpoint.backend_rest.id
  vpc_endpoint_name   = "${local.prefix}-vpc-backend-${var.vpc_id}"
  region              = var.region
  depends_on          = [aws_vpc_endpoint.backend_rest]
}

resource "databricks_mws_vpc_endpoint" "relay" {
  provider            = databricks.mws
  account_id          = var.databricks_account_id
  aws_vpc_endpoint_id = aws_vpc_endpoint.relay.id
  vpc_endpoint_name   = "${local.prefix}-vpc-relay-${var.vpc_id}"
  region              = var.region
  depends_on          = [aws_vpc_endpoint.relay]
}

```

## Workspace creation

Once the VPC endpoints are created, they can be supplied in the `databricks_mws_networks` resource for workspace creation with AWS PrivateLink. After the terraform apply is run once (see the comment in the aws_vpc_endpoint resource above), run the terraform apply a second time with the line for private_dns_enabled set to true uncommented to set the proper DNS settings for PrivateLink. For understanding the reason that this needs to be applied twice, see this existing [issue](hashicorp/terraform-provider-aws#7148) in the underlying AWS provider.

The credentials ID which is referenced below is one of the attributes which is created as a result of configuring the cross-account IAM role, which Databricks uses to orchestrate EC2 resources. The credentials are created via [databricks_mws_credentials](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/mws_credentials). Similarly, the storage configuration ID is obtained from the [databricks_mws_storage_configurations](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/mws_storage_configurations) resource.

```hcl
// Inputs are 2 subnets and one security group from existing VPC that will be used for your Databricks workspace
resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = var.databricks_account_id
  network_name       = "${local.prefix}-network"
  security_group_ids = [var.security_group_id]
  subnet_ids         = var.subnet_ids
  vpc_id             = var.vpc_id
  vpc_endpoints {
    dataplane_relay = [databricks_mws_vpc_endpoint.relay.vpc_endpoint_id]
    rest_api        = [databricks_mws_vpc_endpoint.backend_rest_vpce.vpc_endpoint_id]
  }
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
  credentials_id             = databricks_mws_credentials.this.credentials_id
  storage_configuration_id   = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id                 = databricks_mws_networks.this.network_id
  private_access_settings_id = databricks_mws_private_access_settings.pas.private_access_settings_id
  pricing_tier               = "ENTERPRISE"
  depends_on                 = [databricks_mws_networks.this]
}
```

