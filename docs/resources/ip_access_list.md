# databricks_ip_access_list Resource

Security-conscious enterprises that use cloud SaaS applications need to restrict access to their own employees. Authentication helps to prove user identity, but that does not enforce network location of the users. Accessing a cloud service from an unsecured network can pose security risks to an enterprise, especially when the user may have authorized access to sensitive or personal data. Enterprise network perimeters apply security policies and limit access to external services (for example, firewalls, proxies, DLP, and logging), so access beyond these controls are assumed to be untrusted. Please see [IP Access List](https://docs.databricks.com/security/network/ip-access-list.html) for full feature documentation.

-> **Note** The total number of IP addresses and CIDR scopes provided across all ACL Lists in a workspace can not exceed 1000.  Refer to the docs above for specifics.

## Example Usage

```hcl
resource "databricks_workspace_conf" "this" {
  custom_config = {
    "enableIpAccessLists": true
  }
}

resource "databricks_ip_access_list" "allowed-list" {
  label = "allow_in"
  list_type = "ALLOW"
  ip_addresses = [
    "1.2.3.0/24",
    "1.2.5.0/24"
  ]
  depends_on = [databricks_workspace_conf.this]
}
```
## Argument Reference

The following arguments are supported:

* `list_type` -  Can only be "ALLOW" or "BLOCK"
* `ip_addresses` -  This is a field to allow the group to have instance pool create priviliges.
* `label` - (Optional) This is the display name for the given IP ACL List.
* `enabled` - (Optional) Boolean `true` or `false` indicating whether this list should be active.  Defaults to `true`

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `list_id` - Canonical unique identifier for the IP Access List.

## Import

Importing this resource is not currently supported.