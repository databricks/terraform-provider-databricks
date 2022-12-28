---
subcategory: "Security"
---
# databricks_permission_assignment Resource

These resources are invoked in the workspace context.

## Example Usage

In workspace context, adding account-level user to a workspace:

```hcl
resource "databricks_user" "me" {
  user_name = "me@example.com"
}

resource "databricks_permission_assignment" "add_user" {
  principal_id = databricks_user.me.id
  permissions  = ["USER"]
}
```

In workspace context, adding account-level service principal to a workspace:

```hcl
resource "databricks_service_principal" "sp" {
  display_name = "Automation-only SP"
}

resource "databricks_permission_assignment" "add_admin_spn" {
  principal_id = databricks_service_principal.sp.id
  permissions  = ["ADMIN"]
}
```

## Argument Reference

The following arguments are required:

* `principal_id` - Databricks ID of the user, service principal, or group. The principal ID can be retrieved using the SCIM API, or using [databricks_user](../data-sources/user.md), [databricks_service_principal](../data-sources/service_principal.md) or [databricks_group](../data-sources/group.md) data sources.
* `permissions` - The list of workspace permissions to assign to the principal:
  * `"USER"` - Can access the workspace with basic privileges.
  * `"ADMIN"` - Can access the workspace and has workspace admin privileges to manage users and groups, workspace configurations, and more.

## Import

The resource `databricks_permission_assignment` can be imported using the principal id

```bash
terraform import databricks_permission_assignment.this principal_id
```

## Related Resources

The following resources are used in the same context:

* [databricks_group](group.md) to manage [groups in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/groups.html) or [Account Console](https://accounts.cloud.databricks.com/) (for AWS deployments).
* [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_mws_permission_assignment](mws_permission_assignment.md) to manage permission assignment from an account context
