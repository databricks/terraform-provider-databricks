---
subcategory: "Deployment"
---
# databricks_mws_ncc_private_endpoint_rule Resource

Allows you to create a private endpoint in a [Network Connectivity Config](mws_network_connectivity_config.md) that can be used to [configure private connectivity from serverless compute](https://learn.microsoft.com/en-us/azure/databricks/security/network/serverless-network-security/serverless-private-link).

-> This resource can only be used with an account-level provider!

-> This feature is available on Azure, and in Public Preview on AWS.

## Example Usage

Create a private endpoint to an Azure storage account

```hcl
variable "region" {}
variable "prefix" {}

resource "databricks_mws_network_connectivity_config" "ncc" {
  provider = databricks.account
  name     = "ncc-for-${var.prefix}"
  region   = var.region
}

resource "databricks_mws_ncc_private_endpoint_rule" "storage" {
  provider                       = databricks.account
  network_connectivity_config_id = databricks_mws_network_connectivity_config.ncc.network_connectivity_config_id
  resource_id                    = "/subscriptions/653bb673-1234-abcd-a90b-d064d5d53ca4/resourcegroups/example-resource-group/providers/Microsoft.Storage/storageAccounts/examplesa"
  group_id                       = "blob"
}
```

Create a private endpoint rule to an AWS VPC endpoint and to an S3 bucket

```hcl
variable "region" {}
variable "prefix" {}

resource "databricks_mws_network_connectivity_config" "ncc" {
  provider = databricks.account
  name     = "ncc-for-${var.prefix}"
  region   = var.region
}

resource "databricks_mws_ncc_private_endpoint_rule" "storage" {
  provider                       = databricks.account
  network_connectivity_config_id = databricks_mws_network_connectivity_config.ncc.network_connectivity_config_id
  resource_names                 = ["bucket"]
}

resource "databricks_mws_ncc_private_endpoint_rule" "vpce" {
  provider                       = databricks.account
  network_connectivity_config_id = databricks_mws_network_connectivity_config.ncc.network_connectivity_config_id
  endpoint_service               = "com.amazonaws.vpce.us-west-2.vpce-svc-xyz"
  domain_names                   = ["subdomain.internal.net"]
}
```

## Argument Reference

The following arguments are available:

* `network_connectivity_config_id` - Canonical unique identifier of Network Connectivity Config in Databricks Account. Change forces creation of a new resource.
* `resource_id` - (Azure only) The Azure resource ID of the target resource. Change forces creation of a new resource.
* `group_id` - (Azure only) The sub-resource type (group ID) of the target resource. Must be one of supported resource types (i.e., `blob`, `dfs`, `sqlServer` , etc. Consult the [Azure documentation](https://learn.microsoft.com/en-us/azure/private-link/private-endpoint-overview#private-link-resource) for full list of supported resources). Note that to connect to workspace root storage (root DBFS), you need two endpoints, one for `blob` and one for `dfs`. Change forces creation of a new resource.
* `domain_names` - (AWS only) Only used by private endpoints towards a VPC endpoint service behind a customer-managed VPC endpoint service. List of target AWS resource FQDNs accessible via the VPC endpoint service. Conflicts with `resource_names`.
* `endpoint_service` - (AWS only) Example `com.amazonaws.vpce.us-east-1.vpce-svc-123abcc1298abc123`. The full target AWS endpoint service name that connects to the destination resources of the private endpoint.
* `resource_names` - (AWS only) Only used by private endpoints towards AWS S3 service. List of globally unique S3 bucket names that will be accessed via the VPC endpoint. The bucket names must be in the same region as the NCC/endpoint service. Conflict with `domain_names`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `rule_id`- the ID of a private endpoint rule.
* `endpoint_name` - The name of the Azure private endpoint resource, e.g. "databricks-088781b3-77fa-4132-b429-1af0d91bc593-pe-3cb31234"
* `connection_state` - The current status of this private endpoint. The private endpoint rules are effective only if the connection state is ESTABLISHED. Remember that you must approve new endpoints on your resources in the Azure portal before they take effect.
The possible values are:
  * `PENDING`: The endpoint has been created and pending approval.
  * `ESTABLISHED`: The endpoint has been approved and is ready to be used in your serverless compute resources.
  * `REJECTED`: Connection was rejected by the private link resource owner.
  * `DISCONNECTED`: Connection was removed by the private link resource owner, the private endpoint becomes informative and should be deleted for clean-up.
* `deactivated` - Whether this private endpoint is deactivated.
* `deactivated_at` - Time in epoch milliseconds when this object was deactivated.
* `creation_time` - Time in epoch milliseconds when this object was created.
* `updated_time` - Time in epoch milliseconds when this object was updated.
* `enabled` - Activation status. Only used by private endpoints towards an AWS S3 service.
* `vpc_endpoint_id` - The AWS VPC endpoint ID. You can use this ID to identify the VPC endpoint created by Databricks.

## Import

This resource can be imported by Databricks account ID and Network Connectivity Config ID.

```hcl
import {
  to = databricks_mws_ncc_private_endpoint_rule.this
  id = "<network_connectivity_config_id>/<rule_id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```sh
terraform import databricks_mws_ncc_private_endpoint_rule.this "<network_connectivity_config_id>/<rule_id>"
```

## Related Resources

The following resources are used in the context:

* [databricks_mws_network_connectivity_config](mws_network_connectivity_config.md) to create Network Connectivity Config objects.
* [databricks_mws_ncc_binding](mws_ncc_binding.md) to attach an NCC to a workspace.
