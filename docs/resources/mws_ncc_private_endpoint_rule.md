---
subcategory: "Deployment"
---
# databricks_mws_ncc_private_endpoint_rule Resource

-> **Note** Initialize provider with `alias = "account"`, `host = "https://accounts.azuredatabricks.net"` and use `provider = databricks.account` for all `databricks_mws_*` resources.

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html) in Azure.

Allows you to create a [Network Connectivity Config] that can be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages serverless network connectivity configs](https://learn.microsoft.com/en-us/azure/databricks/sql/admin/serverless-firewall).

## Example Usage

```hcl
variable "region" {}
variable "prefix" {}

resource "databricks_mws_network_connectivity_config" "ncc" {
  provider                     = databricks.account
  name                         = "Network Connectivity Config for ${var.prefix}"
  region                       = var.region
}

resource "databricks_mws_ncc_private_endpoint_rule" "storage" {
  provider                        = databricks.account
  network_connectivity_config_id  = databricks_mws_network_connectivity_config.ncc.id
  resource_id                     = "/subscriptions/653bb673-1234-abcd-a90b-d064d5d53ca4/resourcegroups/example-resource-group/providers/Microsoft.Storage/storageAccounts/examplesa"
  group_id                      = "blob"
}
```

## Argument Reference

The following arguments are available:

* `network_connectivity_config_id` - Canonical unique identifier of Network Connectivity Config in Databricks Account. Change forces creation of a new resource.
* `resource_id` - The Azure resource ID of the target resource. Change forces creation of a new resource.
* `group_id` - The sub-resource type (group ID) of the target resource. Must be one of `blob`, `dfs`, `sqlServer` or `mysqlServer`. Note that to connect to workspace root storage (root DBFS), you need two endpoints, one for blob and one for dfs. Change forces creation of a new resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `rule_id`- the ID of a private endpoint rule.
* `endpoint_name` - The name of the Azure private endpoint resource, e.g. "databricks-088781b3-77fa-4132-b429-1af0d91bc593-pe-3cb31234"
* `connection_state` - The current status of this private endpoint. The private endpoint rules are effective only if the connection state is ESTABLISHED. Remember that you must approve new endpoints on your resources in the Azure portal before they take effect.
The possible values are:
  * `INIT`: (deprecated) The endpoint has been created and pending approval.
  * `PENDING`: The endpoint has been created and pending approval.
  * `ESTABLISHED`: The endpoint has been approved and is ready to use in your serverless compute resources.
  * `REJECTED`: Connection was rejected by the private link resource owner.
  * `DISCONNECTED`: Connection was removed by the private link resource owner, the private endpoint becomes informative and should be deleted for clean-up.
* `deactivated` - Whether this private endpoint is deactivated.
* `deactivated_at` - Time in epoch milliseconds when this object was deactivated.
* `creation_time` - Time in epoch milliseconds when this object was created.
* `updated_time` - Time in epoch milliseconds when this object was updated.

## Import

This resource can be imported by Databricks account ID and Network Connectivity Config ID.

```hcl
terraform import databricks_mws_network_connectivity_config.ncc <accountID>/<NetworkConnectivityConfigID>
```

## Related Resources

The following resources are used in the context:

* [databricks_mws_network_connectivity_config](mws_network_connectivity_config.md) to create Network Connectivity Config objects
* [databricks_mws_ncc_binding](mws_ncc_binding.md) to attach an NCC to a workspace
