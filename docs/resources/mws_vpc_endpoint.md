---
subcategory: "AWS"
---
# databricks_mws_vpc_endpoint Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Connects [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) with [databricks_mws_networks](mws_networks.md) through PrivateLink.

## Example Usage

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

```hcl
resource "aws_vpc_endpoint" "relay" {
  service_name       = local.private_link.relay_endpoint
  vpc_id             = aws_vpc.main.id
  vpc_endpoint_type  = "Interface"
  security_group_ids = [aws_security_group.this.id]
  subnet_ids         = [aws_subnet.databricks_endpoints.id]
  tags = {
    "Name" = "${var.workspace_name}-databricks-relay"
  }
}

resource "databricks_mws_vpc_endpoint" "relay" {
  account_id          = var.databricks_account_id
  aws_vpc_endpoint_id = aws_vpc_endpoint.relay.id
  vpc_endpoint_name   = "VPC Relay for ${aws_vpc.main.id}"
  region              = local.region
}

resource "databricks_mws_networks" "this" {
  # ...

  vpc_endpoints {
    dataplane_relay = [databricks_mws_vpc_endpoint.relay.vpc_endpoint_id]

    # rest_api VPC endpoint is created in a similar way
    rest_api = [databricks_mws_vpc_endpoint.rest.vpc_endpoint_id]
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
