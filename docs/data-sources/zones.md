---
subcategory: "Deployment"
---
# databricks_zones Data Source

This data source allows you to fetch all available AWS availability zones on your workspace on AWS.

-> This data source can only be used with a workspace-level provider!

## Example Usage

```hcl
data "databricks_zones" "zones" {}
```

## Argument Reference

The following arguments are supported for this resource:
* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id for the zone object.
* `default_zone` - This is the default zone that gets assigned to your workspace. This is the zone used by default for clusters and instance pools.
* `zones` - This is a list of all the zones available for your subnets in your Databricks workspace.
