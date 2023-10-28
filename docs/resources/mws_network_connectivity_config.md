---
subcategory: "Deployment"
---
# databricks_mws_network_connectivity_config Resource

-> **Note** Initialize provider with `alias = "mws"`, `host = "https://accounts.azuredatabricks.net"` and use `provider = databricks.mws` for all `databricks_mws_*` resources.

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html) in Azure. This feature is currently not available in other clouds.

Allows you to create a [Network Connectivity Config] that can be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages serverless network connectivity configs](https://learn.microsoft.com/en-us/azure/databricks/sql/admin/serverless-firewall).

## Example Usage

```hcl
variable "region" {}
variable "databricks_account_id" {}
variable "databricks_workspace_id" {}
variable "prefix" {}

resource "databricks_mws_network_connectivity_config" "ncc" {
  provider                     = databricks.mws
  account_id                   = var.databricks_account_id
  name                         = "Network Connectivity Config for ${var.prefix}"
  region                       = var.region
}
```

The `databricks_mws_network_connectivity_config.ncc.network_connectivity_config_id` can then be used as part of a [databricks_mws_workspaces](databricks_mws_workspaces.md) resource:

```hcl
resource "databricks_mws_workspaces" "this" {
  provider                       = databricks.mws
  account_id                     = var.databricks_account_id
  location                       = var.databricks_location
  deployment_name                = var.databricks_workspace_deployment_name
  workspace_name                 = var.databricks_workspace_name
  network_connectivity_config_id = databricks_mws_network_connectivity_config.ncc.network_connectivity_config_id
}
```

-> **Note** The workspace resource here is different from the workspace resource provided via `azurerm` provider. To configure the serverless network connetivity configs, use `databricks.mws` as the provider to configure your workspace in your Databricks Account.

## Argument Reference

The following arguments are available:

* `account_id` - Account Id that could be found in the [Accounts Console](https://learn.microsoft.com/en-us/azure/databricks/administration-guide/account-settings/#--locate-your-account-id).
* `name` - Name of Network Connectivity Config in Databricks Account
* `region` - Region of the Network Connectivity Config. NCCs can only be referenced by your workspaces in the same region.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `network_connectivity_config_id` - Canonical unique identifier of Network Connectivity Config in Databricks Account
* `default_rules.azure_service_endpoint_rule` - This provides a list of subnets. These subnets need to be allowed in your Azure resources in order for Databricks to access. See `default_rules.azure_service_endpoint_rule.target_services` for the supported Azure services.

## Import

This resource can be imported by Databricks account ID and Network Connectivity Config ID.

```hcl
terraform import databricks_mws_network_connectivity_config.ncc <accountID>/<NetworkConnectivityConfigID>
```

## Related Resources

The following resources are used in the context:

* [databricks_mws_workspaces](mws_workspaces.md) to set up Databricks workspaces.
