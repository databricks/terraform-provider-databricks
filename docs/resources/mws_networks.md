---
subcategory: "Deployment"
---
# databricks_mws_networks Resource

## Databricks on AWS usage

-> Initialize provider with `alias = "mws"`, `host  = "https://accounts.cloud.databricks.com"` and use `provider = databricks.mws`

Use this resource to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS. It is essential to understand that this will require you to configure your provider separately for the multiple workspaces resources.

* Databricks must have access to at least two subnets for each workspace, with each subnet in a different Availability Zone. You cannot specify more than one Databricks workspace subnet per Availability Zone in the Create network configuration API call. You can have more than one subnet per Availability Zone as part of your network setup, but you can choose only one subnet per Availability Zone for the Databricks workspace.
* Databricks assigns two IP addresses per node, one for management traffic and one for Spark applications. The total number of instances for each subnet is equal to half of the available IP addresses.
* Each subnet must have a netmask between /17 and /25.
* Subnets must be private.
* Subnets must have outbound access to the public network using a [aws_nat_gateway](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/nat_gateway) and [aws_internet_gateway](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/internet_gateway), or other similar customer-managed appliance infrastructure.
* The NAT gateway must be set up in its subnet (public_subnets in the example below) that routes quad-zero (0.0.0.0/0) traffic to an internet gateway or other customer-managed appliance infrastructure.

-> The NAT gateway needs only one IP address per AZ. Hence, the public subnet only needs two IP addresses. In order to limit the number of IP addresses in the public subnet, you can specify a secondary CIDR block (cidr_block_public) using the argument secondary_cidr_blocks then pass it to the public_subnets argument. Please review the [IPv4 CIDR block association restrictions](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) when choosing the secondary cidr block.

Please follow this [complete runnable example](../guides/aws-workspace.md) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with `host="https://accounts.cloud.databricks.com/"`) and another for the workspace you've created with `databricks_mws_workspaces` resource. If you want both creations of workspaces & clusters within the same Terraform module (essentially the same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module to create workspace + PAT token and the rest in different modules.

## Databricks on GCP usage

-> Initialize provider with `alias = "mws"`, `host  = "https://accounts.gcp.databricks.com"` and use `provider = databricks.mws`

Use this resource to [configure VPC](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/customer-managed-vpc.html) & subnet for new workspaces within GCP. It is essential to understand that this will require you to configure your provider separately for the multiple workspaces resources.

* Databricks must have access to a subnet in the same region as the workspace, of which IP range will be used to allocate your workspace’s GKE cluster nodes.
* The subnet must have a netmask between /29 and /9.
* Databricks must have access to 2 secondary IP ranges, one between /21 to /9 for workspace’s GKE cluster pods, and one between /27 to /16 for workspace’s GKE cluster services.
* Subnet must have outbound access to the public network using a [gcp_compute_router_nat](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_router_nat) or other similar customer-managed appliance infrastructure.

Please follow this [complete runnable example](../guides/gcp-workspace.md) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with `host="https://accounts.gcp.databricks.com/"`) and another for the workspace you've created with `databricks_mws_workspaces` resource. If you want both creations of workspaces & clusters within the same Terraform module (essentially the same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module to create workspace + PAT token and the rest in different modules.

## Example Usage

### Creating a Databricks on AWS workspace

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
}

data "aws_availability_zones" "available" {}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "2.70.0"

  name                  = local.prefix
  cidr                  = var.cidr_block
  secondary_cidr_blocks = [var.cidr_block_public]
  azs                   = data.aws_availability_zones.available.names
  tags                  = var.tags

  enable_dns_hostnames = true
  enable_nat_gateway   = true
  create_igw           = true

  public_subnets = [cidrsubnet(var.cidr_block_public, 6, 0)]
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

In order to create a VPC [that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) you would need to add the `vpc_endpoint_id` Attributes from [mws_vpc_endpoint](mws_vpc_endpoint.md) resources into the [databricks_mws_networks](mws_networks.md) resource. For example:

```hcl
resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = var.databricks_account_id
  network_name       = "${local.prefix}-network"
  security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids         = module.vpc.private_subnets
  vpc_id             = module.vpc.vpc_id
  vpc_endpoints {
    dataplane_relay = [databricks_mws_vpc_endpoint.relay.vpc_endpoint_id]
    rest_api        = [databricks_mws_vpc_endpoint.workspace.vpc_endpoint_id]
  }
  depends_on = [aws_vpc_endpoint.workspace, aws_vpc_endpoint.relay]
}
```

### Creating a Databricks on GCP workspace

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
}

resource "google_compute_network" "dbx_private_vpc" {
  project                 = var.google_project
  name                    = "tf-network-${random_string.suffix.result}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "test-dbx-${random_string.suffix.result}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.dbx_private_vpc.id
  secondary_ip_range {
    range_name    = "pods"
    ip_cidr_range = "10.1.0.0/16"
  }
  secondary_ip_range {
    range_name    = "svc"
    ip_cidr_range = "10.2.0.0/20"
  }
  private_ip_google_access = true
}

resource "google_compute_router" "router" {
  name    = "my-router-${random_string.suffix.result}"
  region  = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  network = google_compute_network.dbx_private_vpc.id
}

resource "google_compute_router_nat" "nat" {
  name                               = "my-router-nat-${random_string.suffix.result}"
  router                             = google_compute_router.router.name
  region                             = google_compute_router.router.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
}

resource "databricks_mws_networks" "this" {
  account_id   = var.databricks_account_id
  network_name = "test-demo-${random_string.suffix.result}"
  gcp_network_info {
    network_project_id    = var.google_project
    vpc_id                = google_compute_network.dbx_private_vpc.name
    subnet_id             = google_compute_subnetwork.network_with_private_secondary_ip_ranges.name
    subnet_region         = google_compute_subnetwork.network_with_private_secondary_ip_ranges.region
    pod_ip_range_name     = "pods"
    service_ip_range_name = "svc"
  }
}
```

In order to create a VPC [that leverages GCP Private Service Connect](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html) you would need to add the `vpc_endpoint_id` Attributes from [mws_vpc_endpoint](mws_vpc_endpoint.md) resources into the [databricks_mws_networks](mws_networks.md) resource. For example:

```hcl
resource "databricks_mws_networks" "this" {
  account_id   = var.databricks_account_id
  network_name = "test-demo-${random_string.suffix.result}"
  gcp_network_info {
    network_project_id    = var.google_project
    vpc_id                = google_compute_network.dbx_private_vpc.name
    subnet_id             = google_compute_subnetwork.network_with_private_secondary_ip_ranges.name
    subnet_region         = google_compute_subnetwork.network_with_private_secondary_ip_ranges.region
    pod_ip_range_name     = "pods"
    service_ip_range_name = "svc"
  }
  vpc_endpoints {
    dataplane_relay = [databricks_mws_vpc_endpoint.relay.vpc_endpoint_id]
    rest_api        = [databricks_mws_vpc_endpoint.workspace.vpc_endpoint_id]
  }
}
```

## Modifying networks on running workspaces (AWS only)

Due to specifics of platform APIs, changing any attribute of network configuration would cause `databricks_mws_networks` to be re-created - deleted & added again with special case for running workspaces. Once network configuration is attached to a running [databricks_mws_workspaces](mws_workspaces.md), you cannot delete it and `terraform apply` would result in `INVALID_STATE: Unable to delete, Network is being used by active workspace X` error. In order to modify any attributes of a network, you have to perform three different `terraform apply` steps:

1. Create a new `databricks_mws_networks` resource.
2. Update the `databricks_mws_workspaces` to point to the new `network_id`.
3. Delete the old `databricks_mws_networks` resource.

## Argument Reference

The following arguments are available:

* `account_id` - Account Id that could be found in the top right corner of [Accounts Console](https://accounts.cloud.databricks.com/)
* `network_name` - name under which this network is registered
* `vpc_id` - (AWS only) [aws_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) id
* `subnet_ids` - (AWS only) ids of [aws_subnet](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet)
* `security_group_ids` - (AWS only) ids of [aws_security_group](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)
* `vpc_endpoints` - (Optional) mapping of [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) for PrivateLink or Private Service Connect connections
* `gcp_network_info` - (GCP only) a block consists of Google Cloud specific information for this network, for example the VPC ID, subnet ID, and secondary IP ranges. It has the following fields:
  * `network_project_id` - The Google Cloud project ID of the VPC network.
  * `vpc_id` - The ID of the VPC associated with this network. VPC IDs can be used in multiple network configurations.
  * `subnet_id` - The ID of the subnet associated with this network.
  * `subnet_region` - The Google Cloud region of the workspace data plane. For example, `us-east4`.
  * `pod_ip_range_name` - The name of the secondary IP range for pods. A Databricks-managed GKE cluster uses this IP range for its pods. This secondary IP range can only be used by one workspace.
  * `service_ip_range_name` - The name of the secondary IP range for services. A Databricks-managed GKE cluster uses this IP range for its services. This secondary IP range can only be used by one workspace.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws networks.
* `network_id` - (String) id of network to be used for [databricks_mws_workspaces](mws_workspaces.md) resource.
* `vpc_status` - (String) VPC attachment status
* `workspace_id` - (Integer) id of associated workspace

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on AWS](../guides/aws-workspace.md) guide.
* [Provisioning Databricks on AWS with Private Link](../guides/aws-private-link-workspace.md) guide.
* [Provisioning AWS Databricks workspaces with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide.
* [Provisioning Databricks on GCP](../guides/gcp-workspace.md) guide.
* [Provisioning Databricks workspaces on GCP with Private Service Connect](../guides/gcp-private-service-connect-workspace.md) guide.
* [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) to register [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) resources with Databricks such that they can be used as part of a [databricks_mws_networks](mws_networks.md) configuration.
* [databricks_mws_private_access_settings](mws_private_access_settings.md) to create a Private Access Setting that can be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) or [GCP Private Service Connect](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html).
* [databricks_mws_workspaces](mws_workspaces.md) to set up [AWS and GCP workspaces](https://docs.databricks.com/getting-started/overview.html#e2-architecture-1).
