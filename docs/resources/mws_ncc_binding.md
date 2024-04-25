---
subcategory: "Deployment"
---
# databricks_mws_ncc_binding Resource

-> **Note** Initialize provider with `alias = "account"`, `host = "https://accounts.azuredatabricks.net"` and use `provider = databricks.account` for all `databricks_mws_*` resources.

-> **Public Preview** This feature is available for AWS & Azure only, and is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html) in AWS.

Allows you to attach a [Network Connectivity Config](mws_network_connectivity_config) object to a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages serverless network connectivity configs](https://learn.microsoft.com/en-us/azure/databricks/sql/admin/serverless-firewall).

The NCC and workspace must be in the same region.

## Example Usage

```hcl
variable "region" {}
variable "prefix" {}

resource "databricks_mws_network_connectivity_config" "ncc" {
  provider = databricks.account
  name     = "Network Connectivity Config for ${var.prefix}"
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

* `network_connectivity_config_id` - Canonical unique identifier of Network Connectivity Config in Databricks Account.
* `workspace_id` - Identifier of the workspace to attach the NCC to. Change forces creation of a new resource.

## Related Resources

The following resources are used in the context:

* [databricks_mws_workspaces](mws_workspaces.md) to set up Databricks workspaces.
* [databricks_mws_network_connectivity_config](mws_network_connectivity_config.md) to create Network Connectivity Config objects.
