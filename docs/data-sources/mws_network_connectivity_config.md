---
subcategory: "Deployment"
---
# databricks_mws_network_connectivity_config Data Source

Retrieves information about [databricks_mws_network_connectivity_config](../resources/mws_network_connectivity_config.md) in Databricks Account.

-> This data source can only be used with an account-level provider!

## Example Usage

Fetching information about a network connectivity configuration in Databricks Account

```hcl
provider "databricks" {
  // other configuration
  account_id = "<databricks account id>"
}

data "databricks_mws_network_connectivity_config" "this" {
  name = "ncc"
}

output "config" {
  value = data.databricks_mws_network_connectivity_config.this
}
```

## Argument Reference

* `name` - (Required) Name of the network connectivity configuration.

## Attribute Reference

* `account_id` - The Databricks account ID associated with this network configuration.
* `creation_time` - Time in epoch milliseconds when the network was created.
* `egress_config` - Array of egress configuration objects.
    * `default_rules` - Array of default rules.
        * `aws_stable_ip_rule` - The stable AWS IP CIDR blocks. You can use these to configure the firewall of your resources to allow traffic from your Databricks workspace.
            * `cidr_blocks` - The list of stable IP CIDR blocks from which Databricks network traffic originates when accessing your resources.
        * `azure_service_endpoint_rule` - Array of Azure service endpoint rules.
            * `subnets` - Array of strings representing the subnet IDs.
            * `target_region` - The target region for the service endpoint.
            * `target_services` - Array of target services.
            * `target_rules` - Array of target rules.
                * `azure_private_endpoint_rules` - Array of private endpoint rule objects.
                    * `rule_id` - The ID of a private endpoint rule.
                    * `network_connectivity_config_id` - The ID of a network connectivity configuration, which is the parent resource of this private endpoint rule object.
                    * `resource_id` - The Azure resource ID of the target resource.
                    * `group_id` - The sub-resource type (group ID) of the target resource.
                    * `endpoint_name` - The name of the Azure private endpoint resource.
                    * `connection_state` - The current status of this private endpoint.
                    * `deactivated` - Whether this private endpoint is deactivated.
                    * `deactivated_at` - Time in epoch milliseconds when this object was deactivated.
                    * `creation_time` - Time in epoch milliseconds when this object was created.
                    * `updated_time` - Time in epoch milliseconds when this object was updated.
* `name` - The name of the network connectivity configuration.
* `network_connectivity_config_id` - The Databricks network connectivity configuration ID.
* `region` - The region of the network connectivity configuration.
* `updated_time` - Time in epoch milliseconds when the network was updated.

## Related Resources

The following resources are used in the same context:

* [databricks_mws_network_connectivity_configs](./mws_network_connectivity_configs.md) to get names of all network connectivity configurations.
* [databricks_mws_network_connectivity_config](../resources/mws_network_connectivity_config.md) to manage network connectivity configuration.
