+++
title = "group"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_group`

This resource allows you to create groups in Databricks. You can also associate Databricks users to the following groups. 
This is an alternative to `databricks_scim_group` and useful if you are using an application to sync users & groups with SCIM 
api.

{{% notice note %}} 
You must be a Databricks administrator API token to make SCIM api calls.
{{% /notice %}} 

## Example Usage

```hcl
resource "databricks_group" "my_group" {
  display_name = "group_name"
  allow_cluster_create = "true"
  allow_instance_pool_create = "true"
}
```
## Argument Reference

The following arguments are supported:

#### - `display_name`:
> **(Required)** This is the display name for the given group.

#### - `allow_cluster_create`:
> **(Optional)** This is a field to allow the group to have cluster create priviliges.

#### - `allow_instance_pool_create`:
> **(Optional)** This is a field to allow the group to have instance pool create priviliges.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the group object.

## Import

You can import a `databricks_group` resource with the name `my_group` like the following:

```bash
$ terraform import databricks_group.my_group <group_id>
```

You can use the SCIM api to fetch the `group_id`.
