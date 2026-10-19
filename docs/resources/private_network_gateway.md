---
subcategory: "Provisioning"
---
# databricks_private_network_gateway Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

Manages a private network gateway under a network connectivity configuration (NCC). A gateway connects serverless compute to destinations in your Azure VNet.

This Terraform resource is supported on Azure only. Use an Azure account-level provider with access to the private network gateway V1 API. Set `parent` to `accounts/{account_id}/network-connectivity-configs/{ncc_id}`. The service assigns the gateway ID and returns its canonical resource `name`.

Specify `azure_cloud_connection` for the gateway subnet. Changing the cloud connection requires replacing the gateway. Destinations must be nonempty for `SPECIFIC_DESTINATIONS` and empty for `ALL_TRAFFIC`.

### Creation and polling

Creation is asynchronous on the server. Terraform submits the create request, receives an operation name, and polls that operation automatically with exponential backoff. You do not need a separate Terraform resource, data source, or sleep to wait for creation.

While the operation reports `done: false`, Terraform continues waiting. When it reports `done: true`, Terraform either records the returned gateway in state or reports the operation's error. Resources that reference the gateway wait for its Terraform create step to finish.

The generated resource does not expose a `timeouts` block or set an explicit create timeout. Waiting respects cancellation of the provider request context. Interrupting Terraform does not cancel server-side provisioning; check whether the gateway was created before retrying. If it exists but was not recorded in Terraform state, import it by its canonical resource name before applying again.


## Example Usage
Create an Azure gateway in an existing NCC using an account-level provider. Use a provider release that includes this resource and an account with access to the private network gateway V1 API. Configure account-level authentication outside this file, for example through Databricks unified authentication environment variables.

Supply the input variables through a `.tfvars` file or `TF_VAR_` environment variables. The subnet must already exist. Choose the Azure bandwidth tier explicitly; it has no server-side default.

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

variable "subnet_resource_id" {
  description = "Full Azure resource ID of the existing gateway subnet."
  type        = string
}

variable "bandwidth_tier_gigabits_per_second" {
  description = "Supported Azure bandwidth tier to provision, in Gbps."
  type        = number
}

variable "destination_hostname" {
  description = "Private destination hostname, for example database.internal.example.com."
  type        = string
}

variable "dns_resolver_ip" {
  description = "Private DNS resolver IP address reachable from the gateway."
  type        = string
}

provider "databricks" {
  alias      = "account"
  host       = "https://accounts.azuredatabricks.net"
  account_id = var.databricks_account_id
}

resource "databricks_private_network_gateway" "azure" {
  provider = databricks.account

  parent                             = "accounts/${var.databricks_account_id}/network-connectivity-configs/${var.ncc_id}"
  display_name                       = "private-gateway"
  bandwidth_tier_gigabits_per_second = var.bandwidth_tier_gigabits_per_second
  traffic_mode                       = "SPECIFIC_DESTINATIONS"
  azure_cloud_connection = {
    gateway_subnet = {
      resource_id = var.subnet_resource_id
    }
  }
  destinations = [{
    destination_type = "DNS_NAME"
    value            = var.destination_hostname
  }]
  private_dns_resolvers = [{
    resolver_type = "IP_ADDRESS"
    value         = var.dns_resolver_ip
  }]
}

output "gateway_name" {
  value = databricks_private_network_gateway.azure.name
}

output "gateway_state" {
  value = databricks_private_network_gateway.azure.state
}
```


## Arguments
The following arguments are supported:
* `display_name` (string, required) - The human-readable name of the gateway
* `parent` (string, required) - The network connectivity configuration that will contain the gateway
* `traffic_mode` (string, required) - The traffic routed through this gateway. Possible values are: `ALL_TRAFFIC`, `SPECIFIC_DESTINATIONS`
* `aws_cloud_connection` (PrivateNetworkGatewayAwsCloudConnection, optional) - The AWS connection used by the gateway
* `azure_cloud_connection` (PrivateNetworkGatewayAzureCloudConnection, optional) - The Azure connection used by the gateway
* `bandwidth_tier_gigabits_per_second` (integer, optional) - The provisioned bandwidth tier for an Azure gateway, in gigabits per second.
  Required when creating an Azure gateway
* `destinations` (list of PrivateNetworkGatewayDestination, optional) - The destinations routed through this gateway
* `private_dns_resolvers` (list of PrivateNetworkGatewayPrivateDnsResolver, optional) - The DNS resolvers used for private name resolution

### PrivateNetworkGatewayAwsCloudConnection
* `cross_account_role` (PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole, required) - The IAM role that Databricks assumes to manage gateway resources
* `gateway_subnets` (list of PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet, required) - The subnets where the gateway establishes connectivity
* `security_group_ids` (list of string, required) - The security groups attached to the gateway network interface

### PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet
* `subnet_id` (string, required) - The AWS subnet ID

### PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole
* `role_arn` (string, required) - The ARN of the IAM role

### PrivateNetworkGatewayAzureCloudConnection
* `gateway_subnet` (PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet, required) - The subnet where the gateway establishes connectivity

### PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet
* `resource_id` (string, required) - The full Azure resource ID of the subnet

### PrivateNetworkGatewayDestination
* `destination_type` (string, required) - The destination type. Possible values are: `DNS_NAME`
* `value` (string, required) - The destination value

### PrivateNetworkGatewayPrivateDnsResolver
* `resolver_type` (string, required) - The resolver type. Possible values are: `IP_ADDRESS`
* `value` (string, required) - The resolver value

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - The time when the gateway was created
* `error_message` (string) - The failure reason when the gateway is in the FAILED state
* `name` (string) - The canonical resource name of the gateway, in the form
  `accounts/{account_id}/network-connectivity-configs/{ncc_id}/private-network-gateways/{gateway_id}`
* `state` (string) - The current lifecycle state of the gateway. Possible values are: `CREATING`, `DELETING`, `ESTABLISHED`, `FAILED`
* `update_time` (string) - The time when the gateway was last updated

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_private_network_gateway.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_private_network_gateway.this "name"
```