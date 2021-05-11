---
subcategory: "AWS"
---
# databricks_mws_vpc_endpoint Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Connects [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) with [databricks_mws_networks](mws_networks.md) through PrivateLink.

## Example Usage

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

In order to use this resource, first you will need to create the required VPC Endpoints as your [VPC endpoint requirements](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#vpc-endpoint-requirements) using the [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) resource. For example:

```hcl
rresource "aws_vpc_endpoint" "workspace" {
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

Once you have created the necessary endpoints, you need to register them via this resource (which calls out to the [Databricks Account API](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#step-3-register-your-vpc-endpoint-ids-with-the-account-api)):

```hcl
resource "databricks_mws_vpc_endpoint" "workspace" {
  account_id          = var.databricks_account_id
  aws_vpc_endpoint_id = aws_vpc_endpoint.workspace.id
  vpc_endpoint_name   = "VPC Relay for ${module.vpc.vpc_id}"
  region              = var.region
  depends_on          = [aws_vpc_endpoint.workspace]
}

resource "databricks_mws_vpc_endpoint" "relay" {
  account_id          = var.databricks_account_id
  aws_vpc_endpoint_id = aws_vpc_endpoint.relay.id
  vpc_endpoint_name   = "VPC Relay for ${module.vpc.vpc_id}"
  region              = var.region
  depends_on          = [aws_vpc_endpoint.relay]
}
```

And then pass in the vpc_endpoints to the [databricks_mws_networks](mws_networks.md) resource:

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
    rest_api = [databricks_mws_vpc_endpoint.workspace.vpc_endpoint_id]
  }
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
