---
subcategory: "Provisioning"
---
# databricks_private_network_gateways Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

Lists private network gateways under a network connectivity configuration, following pagination automatically. Gateways being deleted are excluded.

This data source is supported on Azure only. Use an Azure account-level provider with access to the private network gateway V1 API.


## Example Usage
List gateways under an existing network connectivity configuration. Use a provider release that includes this data source and configure account-level authentication outside this file. Supply the account and NCC IDs through input variables.

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databricks/databricks"
    }
  }
}

variable "databricks_account_id" {
  type = string
}

variable "ncc_id" {
  type = string
}

provider "databricks" {
  alias      = "account"
  host       = "https://accounts.azuredatabricks.net"
  account_id = var.databricks_account_id
}

data "databricks_private_network_gateways" "all" {
  provider = databricks.account
  parent   = "accounts/${var.databricks_account_id}/network-connectivity-configs/${var.ncc_id}"
}

output "gateway_names" {
  value = [for gateway in data.databricks_private_network_gateways.all.private_network_gateways : gateway.name]
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The network connectivity configuration containing the gateways


## Attributes
This data source exports a single attribute, `private_network_gateways`. It is a list of resources, each with the following attributes:
* `aws_cloud_connection` (PrivateNetworkGatewayAwsCloudConnection) - The AWS connection used by the gateway
* `azure_cloud_connection` (PrivateNetworkGatewayAzureCloudConnection) - The Azure connection used by the gateway
* `bandwidth_tier_gigabits_per_second` (integer) - The provisioned bandwidth tier for an Azure gateway, in gigabits per second.
  Required when creating an Azure gateway
* `create_time` (string) - The time when the gateway was created
* `destinations` (list of PrivateNetworkGatewayDestination) - The destinations routed through this gateway
* `display_name` (string) - The human-readable name of the gateway
* `error_message` (string) - The failure reason when the gateway is in the FAILED state
* `name` (string) - The canonical resource name of the gateway, in the form
  `accounts/{account_id}/network-connectivity-configs/{ncc_id}/private-network-gateways/{gateway_id}`
* `private_dns_resolvers` (list of PrivateNetworkGatewayPrivateDnsResolver) - The DNS resolvers used for private name resolution
* `state` (string) - The current lifecycle state of the gateway. Possible values are: `CREATING`, `DELETING`, `ESTABLISHED`, `FAILED`
* `traffic_mode` (string) - The traffic routed through this gateway. Possible values are: `ALL_TRAFFIC`, `SPECIFIC_DESTINATIONS`
* `update_time` (string) - The time when the gateway was last updated

### PrivateNetworkGatewayAwsCloudConnection
* `cross_account_role` (PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) - The IAM role that Databricks assumes to manage gateway resources
* `gateway_subnets` (list of PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) - The subnets where the gateway establishes connectivity
* `security_group_ids` (list of string) - The security groups attached to the gateway network interface

### PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet
* `subnet_id` (string) - The AWS subnet ID

### PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole
* `role_arn` (string) - The ARN of the IAM role

### PrivateNetworkGatewayAzureCloudConnection
* `gateway_subnet` (PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) - The subnet where the gateway establishes connectivity

### PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet
* `resource_id` (string) - The full Azure resource ID of the subnet

### PrivateNetworkGatewayDestination
* `destination_type` (string) - The destination type. Possible values are: `DNS_NAME`
* `value` (string) - The destination value

### PrivateNetworkGatewayPrivateDnsResolver
* `resolver_type` (string) - The resolver type. Possible values are: `IP_ADDRESS`
* `value` (string) - The resolver value