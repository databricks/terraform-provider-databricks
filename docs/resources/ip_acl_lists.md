# databricks_ip_access_lists Resource

This resource allows you to create IP Access Lists in Databricks to control access to your workspace by IP. All IPs and CIDR ranges from each enabled list is put together applying type "BLACKLIST" first, then if still allowed type "WHITELIST" is checked.  Please see [IP Access List](https://docs.databricks.com/security/network/ip-access-list.html) for full feature documentation.

-> **Note** The total number of IP addresses and CIDR scopes provided across all ACL Lists in a workspace can not exceed 1000.  Refer to the docs above for specifics.

## Example Usage

```hcl
resource "databricks_ip_access_list" "allowed-list" {
  label = "allow_in"
  list_type = "WHITELIST"
  ip_addresses = [
    "1.2.3.0/24",
    "1.2.5.0/24"
  ]
  depends_on = [<Replace with terraform workspace id syntax>]
}
```
## Argument Reference

The following arguments are supported:

* `label` -  **(Optional)** This is the display name for the given IP ACL List.

* `list_type` -  **(Required)** Can only be "WHITELIST" or "BLACKLIST"

* `ip_addresses` -  **(Required)** This is a field to allow the group to have instance pool create priviliges.

* `enabled` - **(Optional)** Boolean `true` or `false` indicating whether this list should be active.  Defaults to `true`

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `list_id` - Canonical unique identifier for the IP Access List.

## Import

-> **Note** Importing this resource is not currently supported.