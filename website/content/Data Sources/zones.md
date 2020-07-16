+++
title = "zones"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Data Source: `databricks_zones`

This data source allows you to fetch all available availability zones on your aws workspace.

{{% notice note %}} 
This is only available on an AWS workspace.
{{% /notice %}} 

## Example Usage

```hcl
data "databricks_zones" "zones" {}
```
## Argument Reference

The are no arguments to this data source and only attributes that are computed.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the zone object.

#### - `default_zone`:
> This is the default zone that gets assigned to your workspace. This is the zone used by default for 
>clusters and instance pools.

#### - `zones`:
> This is a list of all the zones available for your subnets in your Databricks workspace.
