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

There are no arguments to this data source and only attributes that are computed.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id for the zone object.
* `default_zone` - This is the default zone that gets assigned to your workspace. This is the zone used by default for clusters and instance pools.
* `zones` - This is a list of all the zones available for your subnets in your Databricks workspace.
