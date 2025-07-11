---
subcategory: "Deployment"
---
# databricks_mws_network_connectivity_config Resource

Allows you to create a Network Connectivity Config that can be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages serverless network connectivity configs](https://learn.microsoft.com/en-us/azure/databricks/security/network/serverless-network-security/serverless-firewall).

-> This resource can only be used with an account-level provider!

## Example Usage

```hcl
variable "region" {}
variable "prefix" {}

resource "databricks_mws_network_connectivity_config" "ncc" {
  provider = databricks.account
  name     = "ncc-for-${var.prefix}"
  region   = var.region
}

resource "databricks_mws_ncc_binding" "ncc_binding" {
  provider                       = databricks.account
  network_connectivity_config_id = databricks_mws_network_connectivity_config.ncc.network_connectivity_config_id
  workspace_id                   = var.databricks_workspace_id
}
```

## Argument Reference

The following arguments are available:

* `name` - Name of the network connectivity configuration. The name can contain alphanumeric characters, hyphens, and underscores. The length must be between 3 and 30 characters. The name must match the regular expression `^[0-9a-zA-Z-_]{3,30}$`. Change forces creation of a new resource.
* `region` - Region of the Network Connectivity Config. NCCs can only be referenced by your workspaces in the same region. Change forces creation of a new resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - combination of `account_id` and `network_connectivity_config_id` separated by `/` character
* `network_connectivity_config_id` - Canonical unique identifier of Network Connectivity Config in Databricks Account
* `egress_config` - block containing information about network connectivity rules that apply to network traffic from your serverless compute resources. Consists of the following fields:
  * `default_rules` - block describing network connectivity rules that are applied by default without resource specific configurations.  Consists of the following fields:
    * `aws_stable_ip_rule` (AWS only) - block with information about stable AWS IP CIDR blocks. You can use these to configure the firewall of your resources to allow traffic from your Databricks workspace.  Consists of the following fields:
      * `cidr_blocks` - list of IP CIDR blocks.
    * `azure_service_endpoint_rule` (Azure only) - block with information about stable Azure service endpoints. You can configure the firewall of your Azure resources to allow traffic from your Databricks serverless compute resources.  Consists of the following fields:
      * `subnets` - list of subnets from which Databricks network traffic originates when accessing your Azure resources.
      * `target_region` - the Azure region in which this service endpoint rule applies.
      * `target_services` - the Azure services to which this service endpoint rule applies to.
  * `target_rules` - block describing network connectivity rules that configured for each destinations. These rules override default rules.  Consists of the following fields:
    * `azure_private_endpoint_rules` (Azure only) - list containing information about configure Azure Private Endpoints.
    * `aws_private_endpoint_rules` (AWS only) - list containing information about configure AWS Private Endpoints.
* `creation_time` - time in epoch milliseconds when this object was created.
* `updated_time` - time in epoch milliseconds when this object was updated.

## Import

This resource can be imported by Databricks account ID and Network Connectivity Config ID.

```hcl
import {
  to = databricks_mws_network_connectivity_config.this
  id = "<account_id>/<network_connectivity_config_id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_mws_network_connectivity_config.this "<account_id>/<network_connectivity_config_id>"
```

## Related Resources

The following resources are used in the context:

* [databricks_mws_workspaces](mws_workspaces.md) to set up Databricks workspaces.
* [databricks_mws_ncc_binding](mws_ncc_binding.md) to attach an NCC to a workspace.
* [databricks_mws_ncc_private_endpoint_rule](mws_ncc_private_endpoint_rule.md) to create a private endpoint rule.
