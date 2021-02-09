---
subcategory: "Workspace"
---
# databricks_workspace_conf Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

Manages workspace configuration for expert usage. Currently, more than one instance of resource can exist in Terraform state, though there's no deterministic behavior, when they manage the same property. We strongly recommend to use a single `databricks_workspace_conf` per workspace.

## Example Usage

Allows specification of custom configuration properties for expert usage:

 * `enableIpAccessLists` - enables the use of [databricks_ip_access_list](ip_access_list.md) resources

```hcl
resource "databricks_workspace_conf" "this" {
    custom_config = {
        "enableIpAccessLists": true
    }
}
```

## Argument Reference

The following arguments are available:

* `custom_config` - (Required) Key-value map of strings, that represent workspace configuration. Upon resource deletion, properties that start with `enable` or `enforce` will be reset to `false` value, regardless of initial default one.

## Import

This resource doesn't support import.
