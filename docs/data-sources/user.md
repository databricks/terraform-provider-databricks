---
subcategory: "Security"
---

# databricks_user Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves information about [databricks_user](../resources/user.md).

## Example Usage

Adding user to administrative group

```hcl
data "databricks_group" "admins" {
  display_name = "admins"
}

data "databricks_user" "me" {
  user_name = "me@example.com"
}

resource "databricks_group_member" "my_member_a" {
  group_id  = data.databricks_group.admins.id
  member_id = data.databricks_user.me.id
}
```

## Argument Reference

Data source allows you to pick groups by the following attributes

- `user_name` - (Optional) User name of the user. The user must exist before this resource can be planned.
- `user_id` - (Optional) ID of the user.

## Attribute Reference

Data source exposes the following attributes:

- `id` - The id of the user.
- `external_id` - ID of the user in an external identity provider.
- `user_name` - Name of the [user](../resources/user.md), e.g. `mr.foo@example.com`.
- `display_name` - Display name of the [user](../resources/user.md), e.g. `Mr Foo`.
- `home` - Home folder of the [user](../resources/user.md), e.g. `/Users/mr.foo@example.com`.
- `repos` - Personal Repos location of the [user](../resources/user.md), e.g. `/Repos/mr.foo@example.com`.
- `alphanumeric` - Alphanumeric representation of user local name. e.g. `mr_foo`.
- `active` - Whether the [user](../resources/user.md) is active.

* `acl_principal_id` - identifier for use in [databricks_access_control_rule_set](../resources/access_control_rule_set.md), e.g. `users/mr.foo@example.com`.

## Related Resources

The following resources are used in the same context:

- [End to end workspace management](../guides/workspace-management.md) guide.
- [databricks_current_user](current_user.md) data to retrieve information about [databricks_user](../resources/user.md) or [databricks_service_principal](../resources/service_principal.md), that is calling Databricks REST API.
- [databricks_group](../resources/group.md) to manage [groups in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/groups.html) or [Account Console](https://accounts.cloud.databricks.com/) (for AWS deployments).
- [databricks_group](group.md) data to retrieve information about [databricks_group](../resources/group.md) members, entitlements and instance profiles.
- [databricks_group_instance_profile](../resources/group_instance_profile.md) to attach [databricks_instance_profile](../resources/instance_profile.md) (AWS) to [databricks_group](../resources/group.md).
- [databricks_group_member](../resources/group_member.md) to attach [users](../resources/user.md) and [groups](../resources/group.md) as group members.
- [databricks_permissions](../resources/permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
- [databricks_user](../resources/user.md) to [manage users](https://docs.databricks.com/administration-guide/users-groups/users.html), that could be added to [databricks_group](../resources/group.md) within the workspace.
- [databricks_user_instance_profile](../resources/user_instance_profile.md) to attach [databricks_instance_profile](../resources/instance_profile.md) (AWS) to [databricks_user](../resources/user.md).
