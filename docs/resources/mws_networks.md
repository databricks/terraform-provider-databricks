---
subcategory: "AWS"
---
# databricks_mws_networks Resource

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

Use this resource to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS. It is essential to understand that this will require you to configure your provider separately for the multiple workspaces resources.

* Databricks must have access to at least two subnets for each workspace, with each subnet in a different Availability Zone. You cannot specify more than one Databricks workspace subnet per Availability Zone in the Create network configuration API call. You can have more than one subnet per Availability Zone as part of your network setup, but you can choose only one subnet per Availability Zone for the Databricks workspace.
* Databricks assigns two IP addresses per node, one for management traffic and one for Spark applications. The total number of instances for each subnet is equal to half of the available IP addresses.
* Each subnet must have a netmask between /17 and /25.
* Subnets must be private.
* Subnets must have outbound access to the public network using a [aws_nat_gateway](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/nat_gateway) and [aws_internet_gateway](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/internet_gateway), or other similar customer-managed appliance infrastructure.
* The NAT gateway must be set up in its subnet that routes quad-zero (0.0.0.0/0) traffic to an internet gateway or other customer-managed appliance infrastructure.

Please follow this [complete runnable example](../guides/aws-workspace.md) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with `host="https://accounts.cloud.databricks.com/") and another for the workspace you've created with `databricks_mws_workspaces` resource. If you want both creations of workspaces & clusters within the same Terraform module (essentially the same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module to create workspace + PAT token and the rest in different modules.

## Example Usage

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the bottom left corner of https://accounts.cloud.databricks.com/"
}

data "aws_availability_zones" "available" {}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "2.70.0"

  name = local.prefix
  cidr = var.cidr_block
  azs  = data.aws_availability_zones.available.names
  tags = var.tags

  enable_dns_hostnames = true
  enable_nat_gateway   = true
  create_igw           = true

  private_subnets = [cidrsubnet(var.cidr_block, 3, 1),
  cidrsubnet(var.cidr_block, 3, 2)]

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
  account_id         = var.databricks_account_id
  network_name       = "${local.prefix}-network"
  security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids         = module.vpc.private_subnets
  vpc_id             = module.vpc.vpc_id
}
```

In order to create a VPC [that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) you would need to add the `vpc_endpoint_id` Attributes from [mws_vpc_endpoint](mws_vpc_endpoint.md) resources into the [databricks_mws_networks](databricks_mws_networks.md) resource. For example:

```hcl
 resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = var.databricks_account_id
  network_name       = "${local.prefix}-network"
  security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids         = module.vpc.private_subnets
  vpc_id             = module.vpc.vpc_id
  vpc_endpoints {
    dataplane_relay  = [databricks_mws_vpc_endpoint.relay.vpc_endpoint_id]
    rest_api         = [databricks_mws_vpc_endpoint.workspace.vpc_endpoint_id]
  }
  depends_on         = [aws_vpc_endpoint.workspace, aws_vpc_endpoint.relay]
}
  ```

## Modifying networks on running workspaces

Due to specifics of platform APIs, changing any attribute of network configuration would cause `databricks_mws_networks` to be re-created - deleted & added again with special case for running workspaces. Once network configuration is attached to a running [databricks_mws_workspaces](mws_workspaces.md), you cannot delete it and `terraform apply` would result in `INVALID_STATE: Unable to delete, Network is being used by active workspace X` error. In order to modify any attributes of a network, you have to perform three different `terraform apply` steps:

1. Create a new `databricks_mws_networks` resource.
2. Update the `databricks_mws_workspaces` to point to the new `network_id`.
3. Delete the old `databricks_mws_networks` resource.

## Argument Reference

The following arguments are available:

* `account_id` - Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)
* `network_name` - name under which this network is regisstered
* `vpc_id` - [aws_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) id
* `subnet_ids` - ids of [aws_subnet](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet)
* `security_group_ids` - ids of [aws_security_group](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)
* `vpc_endpoints` (Optional) - mapping of [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) for PrivateLink connections

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws networks.
* `network_id` - (String) id of network to be used for [databricks_mws_workspace](mws_workspaces.md) resource.
* `vpc_status` - (String) VPC attachment status
* `workspace_id` - (Integer) id of associated workspace
