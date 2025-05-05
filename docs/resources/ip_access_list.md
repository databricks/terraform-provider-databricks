---
subcategory: "Security"
---
# databricks_ip_access_list Resource

Security-conscious enterprises that use cloud SaaS applications need to restrict access to their own employees. Authentication helps to prove user identity, but that does not enforce network location of the users. Accessing a cloud service from an unsecured network can pose security risks to an enterprise, especially when the user may have authorized access to sensitive or personal data. Enterprise network perimeters apply security policies and limit access to external services (for example, firewalls, proxies, DLP, and logging), so access beyond these controls are assumed to be untrusted. Please see [IP Access List](https://docs.databricks.com/security/network/ip-access-list.html) for full feature documentation.

-> This resource can only be used with a workspace-level provider!

-> The total number of IP addresses and CIDR scopes provided across all ACL Lists in a workspace can not exceed 1000.  Refer to the docs above for specifics.

## Example Usage

```hcl
resource "databricks_workspace_conf" "this" {
  custom_config = {
    "enableIpAccessLists" = true
  }
}

resource "databricks_ip_access_list" "allowed-list" {
  label     = "allow_in"
  list_type = "ALLOW"
  ip_addresses = [
    "1.1.1.1",
    "1.2.3.0/24",
    "1.2.5.0/24"
  ]
  depends_on = [databricks_workspace_conf.this]
}
```

## Argument Reference

The following arguments are supported:

* `list_type` -  Can only be "ALLOW" or "BLOCK".
* `ip_addresses` - A string list of IP addresses and CIDR ranges.
* `label` -  This is the display name for the given IP ACL List.
* `enabled` - (Optional) Boolean `true` or `false` indicating whether this list should be active.  Defaults to `true`

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the IP Access List, same as `list_id`.
* `list_id` - Canonical unique identifier for the IP Access List.

## Import

The databricks_ip_access_list can be imported using id:

```bash
terraform import databricks_ip_access_list.this <list-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [Provisioning AWS Databricks workspaces with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide.
* [databricks_mws_networks](mws_networks.md) to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS.
* [databricks_mws_private_access_settings](mws_private_access_settings.md) to create a [Private Access Setting](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#step-5-create-a-private-access-settings-configuration-using-the-databricks-account-api) that can be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
* [databricks_permissions](permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
* [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html).
