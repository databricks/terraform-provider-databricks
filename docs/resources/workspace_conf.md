---
subcategory: "Workspace"
---

# databricks_workspace_conf Resource

~> This resource has an evolving API, which may change in future versions of the provider.

Manages workspace configuration for expert usage. Currently, more than one instance of resource can exist in Terraform state, though there's no deterministic behavior, when they manage the same property. We strongly recommend to use a single `databricks_workspace_conf` per workspace.

-> Deleting `databricks_workspace_conf` resources may fail depending on the configuration properties set, including but not limited to `enableIpAccessLists`, `enableGp3`, and `maxTokenLifetimeDays`. The provider will print a warning if this occurs. You can verify the workspace configuration by reviewing [the workspace settings in the UI](https://docs.databricks.com/en/admin/workspace-settings/index.html).

## Example Usage

Allows specification of custom configuration properties for expert usage:

- `enableIpAccessLists` - enables the use of [databricks_ip_access_list](ip_access_list.md) resources
- `maxTokenLifetimeDays` - (string) Maximum token lifetime of new tokens in days, as an integer. If zero, new tokens are permitted to have no lifetime limit. Negative numbers are unsupported. **WARNING:** This limit only applies to new tokens, so there may be tokens with lifetimes longer than this value, including unlimited lifetime. Such tokens may have been created before the current maximum token lifetime was set.
- `enableTokensConfig` - (boolean) Enable or disable personal access tokens for this workspace.
- `enableDeprecatedClusterNamedInitScripts` - (boolean) Enable or disable [legacy cluster-named init scripts](https://docs.databricks.com/clusters/init-scripts.html#disable-legacy-cluster-named-init-scripts-for-a-workspace) for this workspace.
- `enableDeprecatedGlobalInitScripts` - (boolean) Enable or disable [legacy global init scripts](https://docs.databricks.com/clusters/init-scripts.html#migrate-legacy-scripts) for this workspace.

```hcl
resource "databricks_workspace_conf" "this" {
  custom_config = {
    "enableIpAccessLists" : true
  }
}
```

## Argument Reference

The following arguments are available:

- `custom_config` - (Required) Key-value map of strings that represent workspace configuration. Upon resource deletion, properties that start with `enable` or `enforce` will be reset to `false` value, regardless of initial default one.

## Import

!> Importing this resource is not currently supported.
