---
subcategory: "Deployment"
---
# databricks_mws_vpc_endpoint Resource

-> **Note** Initialize provider with `alias = "mws"`, `host  = "https://accounts.cloud.databricks.com"` and use `provider = databricks.mws`

Enables you to register [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) resources or gcp vpc_endpoint resources with Databricks such that they can be used as part of a [databricks_mws_networks](mws_networks.md) configuration.

It is strongly recommended that customers read the [Enable AWS Private Link](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) or the [Enable GCP Private Service Connect](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html) documentation before trying to leverage this resource.

## Example Usage

### Databricks on AWS usage

Before using this resource, you will need to create the necessary VPC Endpoints as per your [VPC endpoint requirements](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#vpc-endpoint-requirements). You can use the [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) resource for this, for example:

```hcl
resource "aws_vpc_endpoint" "workspace" {
  vpc_id              = module.vpc.vpc_id
  service_name        = local.private_link.workspace_service
  vpc_endpoint_type   = "Interface"
  security_group_ids  = [module.vpc.default_security_group_id]
  subnet_ids          = [aws_subnet.pl_subnet.id]
  depends_on          = [aws_subnet.pl_subnet]
  private_dns_enabled = true
}

resource "aws_vpc_endpoint" "relay" {
  vpc_id              = module.vpc.vpc_id
  service_name        = local.private_link.relay_service
  vpc_endpoint_type   = "Interface"
  security_group_ids  = [module.vpc.default_security_group_id]
  subnet_ids          = [aws_subnet.pl_subnet.id]
  depends_on          = [aws_subnet.pl_subnet]
  private_dns_enabled = true
}
```

Depending on your use case, you may need or choose to add VPC Endpoints for the AWS Services Databricks uses. See [Add VPC endpoints for other AWS services (recommended but optional)
](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#step-9-add-vpc-endpoints-for-other-aws-services-recommended-but-optional) for more information. For example:

```hcl
resource "aws_vpc_endpoint" "s3" {
  vpc_id          = module.vpc.vpc_id
  route_table_ids = module.vpc.private_route_table_ids
  service_name    = "com.amazonaws.${var.region}.s3"
  depends_on      = [module.vpc]
}

resource "aws_vpc_endpoint" "sts" {
  vpc_id              = module.vpc.vpc_id
  service_name        = "com.amazonaws.${var.region}.sts"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = module.vpc.private_subnets
  security_group_ids  = [module.vpc.default_security_group_id]
  depends_on          = [module.vpc]
  private_dns_enabled = true
}

resource "aws_vpc_endpoint" "kinesis-streams" {
  vpc_id             = module.vpc.vpc_id
  service_name       = "com.amazonaws.${var.region}.kinesis-streams"
  vpc_endpoint_type  = "Interface"
  subnet_ids         = module.vpc.private_subnets
  security_group_ids = [module.vpc.default_security_group_id]
  depends_on         = [module.vpc]
}
```

Once you have created the necessary endpoints, you need to register each of them via *this* Terraform resource, which calls out to the [Databricks Account API](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#step-3-register-your-vpc-endpoint-ids-with-the-account-api)):

```hcl
resource "databricks_mws_vpc_endpoint" "workspace" {
  provider            = databricks.mws
  account_id          = var.databricks_account_id
  aws_vpc_endpoint_id = aws_vpc_endpoint.workspace.id
  vpc_endpoint_name   = "VPC Relay for ${module.vpc.vpc_id}"
  region              = var.region
  depends_on          = [aws_vpc_endpoint.workspace]
}

resource "databricks_mws_vpc_endpoint" "relay" {
  provider            = databricks.mws
  account_id          = var.databricks_account_id
  aws_vpc_endpoint_id = aws_vpc_endpoint.relay.id
  vpc_endpoint_name   = "VPC Relay for ${module.vpc.vpc_id}"
  region              = var.region
  depends_on          = [aws_vpc_endpoint.relay]
}
```

Typically the next steps after this would be to create a [databricks_mws_private_access_settings](mws_private_access_settings.md) and [databricks_mws_networks](mws_networks.md) configuration, before passing the `databricks_mws_private_access_settings.pas.private_access_settings_id` and `databricks_mws_networks.this.network_id` into a [databricks_mws_workspaces](mws_workspaces.md) resource:

```hcl
resource "databricks_mws_workspaces" "this" {
  provider                   = databricks.mws
  account_id                 = var.databricks_account_id
  aws_region                 = var.region
  workspace_name             = local.prefix
  credentials_id             = databricks_mws_credentials.this.credentials_id
  storage_configuration_id   = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id                 = databricks_mws_networks.this.network_id
  private_access_settings_id = databricks_mws_private_access_settings.pas.private_access_settings_id
  pricing_tier               = "ENTERPRISE"
  depends_on                 = [databricks_mws_networks.this]
}
```

### Databricks on GCP usage

Before using this resource, you will need to create the necessary Private Service Connect (PSC) connections on your Google Cloud VPC networks. You can see [Enable Private Service Connect for your workspace](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html) for more details.

Once you have created the necessary PSC connections, you need to register each of them via *this* Terraform resource, which calls out to the Databricks Account API.

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in https://accounts.gcp.databricks.com/"
}
variable "databricks_google_service_account" {}
variable "google_project" {}
variable "subnet_region" {}

provider "databricks" {
  alias = "mws"
  host  = "https://accounts.gcp.databricks.com"
}

resource "databricks_mws_vpc_endpoint" "workspace" {
  provider          = databricks.mws
  account_id        = var.databricks_account_id
  vpc_endpoint_name = "PSC Rest API endpoint"
  gcp_vpc_endpoint_info {
    project_id        = var.google_project
    psc_endpoint_name = "PSC Rest API endpoint"
    endpoint_region   = var.subnet_region
  }
}

resource "databricks_mws_vpc_endpoint" "relay" {
  provider          = databricks.mws
  account_id        = var.databricks_account_id
  vpc_endpoint_name = "PSC Relay endpoint"
  gcp_vpc_endpoint_info {
    project_id        = var.google_project
    psc_endpoint_name = "PSC Relay endpoint"
    endpoint_region   = var.subnet_region
  }
}
```

Typically the next steps after this would be to create a [databricks_mws_private_access_settings](mws_private_access_settings.md) and [databricks_mws_networks](mws_networks.md) configuration, before passing the `databricks_mws_private_access_settings.pas.private_access_settings_id` and `databricks_mws_networks.this.network_id` into a [databricks_mws_workspaces](mws_workspaces.md) resource:

```hcl
resource "databricks_mws_workspaces" "this" {
  provider       = databricks.mws
  account_id     = var.databricks_account_id
  workspace_name = "gcp workspace"
  location       = var.subnet_region
  cloud_resource_container {
    gcp {
      project_id = var.google_project
    }
  }
  gke_config {
    connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
    master_ip_range   = "10.3.0.0/28"
  }
  network_id                 = databricks_mws_networks.this.network_id
  private_access_settings_id = databricks_mws_private_access_settings.pas.private_access_settings_id
  pricing_tier               = "PREMIUM"
  depends_on                 = [databricks_mws_networks.this]
}
```

## Argument Reference

The following arguments are required:

* `account_id` - Account Id that could be found in the Accounts Console for [AWS](https://accounts.cloud.databricks.com/) or [GCP](https://accounts.gcp.databricks.com/)
* `aws_vpc_endpoint_id` - (AWS only) ID of configured [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint)
* `vpc_endpoint_name` - Name of VPC Endpoint in Databricks Account
* `region` - (AWS only) Region of AWS VPC
* `gcp_vpc_endpoint_info` - (GCP only) a block consists of Google Cloud specific information for this PSC endpoint. It has the following fields:
  * `project_id` - The Google Cloud project ID of the VPC network where the PSC connection resides.
  * `psc_endpoint_name` - The name of the PSC endpoint in the Google Cloud project.
  * `endpoint_region` - Region of the PSC endpoint.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - the ID of VPC Endpoint in form of `account_id/vpc_endpoint_id`
* `vpc_endpoint_id` - Canonical unique identifier of VPC Endpoint in Databricks Account
* `aws_endpoint_service_id` - (AWS Only) The ID of the Databricks endpoint service that this VPC endpoint is connected to. Please find the list of endpoint service IDs for each supported region in the [Databricks PrivateLink documentation](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html)
* `state` - (AWS Only) State of VPC Endpoint
* `gcp_vpc_endpoint_info`- (GCP only) a block consists of Google Cloud specific information for this PSC endpoint. It has the following fields exported:
  * `psc_connection_id` - The unique ID of this PSC connection.
  * `service_attachment_id` - The service attachment this PSC connection connects to.

## Import

-> **Note** Importing this resource is not currently supported.

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on AWS](../guides/aws-workspace.md) guide.
* [Provisioning Databricks on AWS with PrivateLink](../guides/aws-private-link-workspace.md) guide.
* [Provisioning AWS Databricks E2 with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide.
* [Provisioning Databricks workspaces on GCP with Private Service Connect](../guides/gcp-private-service-connect-workspace.md) guide.
* [databricks_mws_networks](mws_networks.md) to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS.
* [databricks_mws_private_access_settings](mws_private_access_settings.md) to create a [Private Access Setting](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#step-5-create-a-private-access-settings-configuration-using-the-databricks-account-api) that can be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
* [databricks_mws_workspaces](mws_workspaces.md) to set up [workspaces in E2 architecture on AWS](https://docs.databricks.com/getting-started/overview.html#e2-architecture-1).
