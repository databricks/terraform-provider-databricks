---
subcategory: "AWS"
---
# databricks_mws_vpc_endpoint Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Enables you to register [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) resources with Databricks such that they can be used as part of a [databricks_mws_networks](mws_networks.md) configuration.

It is strongly recommended that customers read the [Enable Private Link](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) documentation before trying to leverage this resource.

## Example Usage

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

Before using this resource, you will need to create the necessary VPC Endpoints as per your [VPC endpoint requirements](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#vpc-endpoint-requirements). You can use the [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) resource for this, for example:

```hcl
resource "aws_vpc_endpoint" "workspace" {
  vpc_id             = module.vpc.vpc_id
  service_name       = local.private_link.workspace_service
  vpc_endpoint_type  = "Interface"
  security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids         = [aws_subnet.pl_subnet.id]
  depends_on         = [aws_subnet.pl_subnet]
}

resource "aws_vpc_endpoint" "relay" {
  vpc_id             = module.vpc.vpc_id
  service_name       = local.private_link.relay_service
  vpc_endpoint_type  = "Interface"
  security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids         = [aws_subnet.pl_subnet.id]
  depends_on         = [aws_subnet.pl_subnet]
}
```

Depending on your use case, you may need or chose to add VPC Endpoints for the AWS Services Databricks uses. See [Add VPC endpoints for other AWS services (recommended but optional)
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

Private DNS for the VPC Endpoints cannot be configured until the Endpoints have been accepted on both sides, which is after they have been registered via this resource. You can either do this [manually via the Account Console](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#step-4-enable-private-dns-names-on-aws-vpc-endpoints-using-the-aws-console) or by running a subsequent ```terraform apply``` after you have added ```private_dns_enabled = true``` to the configuration. For example:

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

Typically the next steps after this would be to create a [databricks_mws_private_access_settings](mws_private_access_settings.md) and [databricks_mws_networks](mws_networks.md) configuration, before passing the `databricks_mws_private_access_settings.pas.private_access_settings_id` and `databricks_mws_networks.this.network_id` into a [databricks_mws_workspaces](mws_workspaces.md) resource:

```hcl
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

## Argument Reference

The following arguments are required:

* `account_id` - Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)
* `aws_vpc_endpoint_id` - ID of configured [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint)
* `vpc_endpoint_name` - Name of VPC Endpoint in Databricks Account
* `aws_endpoint_service_id` - ID of Databricks VPC endpoint service to connect to. Please contact your Databricks representative to request mapping
* `region` - Region of AWS VPC

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `vpc_endpoint_id` - Canonical unique identifier of VPC Endpoint in Databricks Account
* `state` - State of VPC Endpoint

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on AWS](../guides/aws-workspace.md) guide.
* [Provisioning Databricks on AWS with PrivateLink](../guides/aws-private-link-workspace.md) guide.
* [Provisioning AWS Databricks E2 with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide.
* [databricks_mws_networks](mws_networks.md) to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS.
* [databricks_mws_private_access_settings](mws_private_access_settings.md) to create a [Private Access Setting](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#step-5-create-a-private-access-settings-configuration-using-the-databricks-account-api) that can be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
* [databricks_mws_workspaces](mws_workspaces.md) to set up [workspaces in E2 architecture on AWS](https://docs.databricks.com/getting-started/overview.html#e2-architecture-1).
