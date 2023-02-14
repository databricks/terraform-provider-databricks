---
subcategory: "Security"
---
# databricks_account_group Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

Retrieves information about account-level [databricks_group](../resources/group.md) members, entitlements and roles.

## Example Usage

Adding an account-level group to a workspace:

```hcl
resource "databricks_account_group" "group" {
  display_name = "dataeng-admin"
}

resource "databricks_permission_assignment" "add_acc_group" {
  principal_id = databricks_account_group.group.id
  permissions  = ["ADMIN"]
}
```

## Argument Reference

Data source allows you to pick groups by the following attributes

* `display_name` - (Required) Display name of the group. The group must exist before this resource can be planned.
* `recursive` - (Optional) Collect information for all nested groups. _Defaults to true._

## Attribute Reference

Data source exposes the following attributes:

* `id` -  The id for the group object.
* `external_id` - ID of the group in an external identity provider.
* `users` - Set of [databricks_user](../resources/user.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `service_principals` - Set of [databricks_service_principal](../resources/service_principal.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `child_groups` - Set of [databricks_group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `groups` - Set of [group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `roles` - Set of roles assigned to the group

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/passthrough-cluster-per-user.md) guide
* [databricks_cluster](../resources/cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_directory](../resources/directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_group_member](../resources/group_member.md) to attach [users](../resources/user.md) and [groups](../resources/group.md) as group members.
* [databricks_permissions](../resources/permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
* [databricks_user](../resources/user.md) to [manage users](https://docs.databricks.com/administration-guide/users-groups/users.html), that could be added to [databricks_group](../resources/group.md) within the workspace.
