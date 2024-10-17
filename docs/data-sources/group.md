---
subcategory: "Security"
---
# databricks_group Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves information about [databricks_group](../resources/group.md) members, entitlements and instance profiles.

## Example Usage

Adding user to administrative group

```hcl
data "databricks_group" "admins" {
  display_name = "admins"
}

resource "databricks_user" "me" {
  user_name = "me@example.com"
}

resource "databricks_group_member" "my_member_a" {
  group_id  = data.databricks_group.admins.id
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
* `external_id` - ID of the group in an external identity provider.
* `users` - Set of [databricks_user](../resources/user.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `service_principals` - Set of [databricks_service_principal](../resources/service_principal.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `child_groups` - Set of [databricks_group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `groups` - Set of [group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `instance_profiles` - Set of [instance profile](../resources/instance_profile.md) ARNs, that can be modified by [databricks_group_instance_profile](../resources/group_instance_profile.md) resource.
* `allow_cluster_create` - True if group members can create [clusters](../resources/cluster.md)
* `allow_instance_pool_create` - True if group members can create [instance pools](../resources/instance_pool.md)
* `acl_principal_id` - identifier for use in [databricks_access_control_rule_set](../resources/access_control_rule_set.md), e.g. `groups/Some Group`.

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide
* [databricks_groups](../data-sources/groups.md) to retrive [Groups](https://docs.databricks.com/en/admin/users-groups/groups.html) information.
* [databricks_cluster](../resources/cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_directory](../resources/directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_group_member](../resources/group_member.md) to attach [users](../resources/user.md) and [groups](../resources/group.md) as group members.
* [databricks_permissions](../resources/permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
* [databricks_user](../resources/user.md) to [manage users](https://docs.databricks.com/administration-guide/users-groups/users.html), that could be added to [databricks_group](../resources/group.md) within the workspace.
