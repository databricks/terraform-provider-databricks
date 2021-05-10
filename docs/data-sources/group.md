---
subcategory: "Security"
---
# databricks_group Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

Retrieves information about [databricks_group](../resources/group.md) members, entitlements and instance profiles.

## Example Usage

Adding user to administrative group

```hcl
data "databricks_group" "admins" {
    display_name = "admins"
}

resource "databricks_user" "me" {
  user_name    = "me@example.com"
}

resource "databricks_group_member" "my_member_a" {
  group_id = data.databricks_group.admins.id
  member_id = databricks_user.me.id
}
```

## Argument Reference

Data source allows you to pick groups by the following attributes

* `display_name` - (Required) Display name of the group. The group must exist before this resource can be planned.
* `recursive` - (Optional) Collect information for all nested groups. *Defaults to true.*

## Attribute Reference

Data source exposes the following attributes:

* `id` -  The id for the group object.
* `members` - Set of [user](../resources/user.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `groups` - Set of [group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `instance_profiles` - Set of [instance profile](../resources/instance_profile.md) ARNs, that can be modified by [databricks_group_instance_profile](../resources/group_instance_profile.md) resource.
* `allow_cluster_create` - True if group members can create [clusters](../resources/cluster.md)
* `allow_instance_pool_create` - True if group members can create [instance pools](../resources/instance_pool.md)
