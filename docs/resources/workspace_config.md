# databricks_workspace_conf Resource

Wrapper resource for the workspace config map on a databricks workspace.  This is referenced throughout different feature sections.  Currently, only used in support of IP Access Lists.

## Example Usage

```hcl
resource "databricks_workspace_conf" "features" {
  enable_ip_access_lists = "true"
}
```
## Argument Reference

The following arguments are supported:

* `enable_ip_access_lists` -  **(Optional)** Boolean true or false to enable the [IP ACL Lists](https://docs.databricks.com/security/network/ip-access-list).html] feature.  If this is enabled and there are no IP ACL Lists defined, all IPs are allowed.
