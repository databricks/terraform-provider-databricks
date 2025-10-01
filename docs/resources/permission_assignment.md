---
subcategory: "Security"
---
# databricks_permission_assignment Resource

This resource is used to assign account-level users, service principals and groups to a Databricks workspace. To configure additional entitlements such as cluster creation, please use [databricks_entitlements](entitlements.md)

-> This resource can only be used with a workspace-level provider!

## Example Usage

### Assign using `principal_id`

In workspace context, adding account-level user to a workspace:

```hcl
# Use the account provider
data "databricks_user" "me" {
  user_name = "me@example.com"
  provider  = databricks.account
}

resource "databricks_permission_assignment" "add_user" {
  principal_id = data.databricks_user.me.id
  permissions  = ["USER"]
  provider     = databricks.workspace
}
```

In workspace context, adding account-level service principal to a workspace:

```hcl
# Use the account provider
data "databricks_service_principal" "sp" {
  display_name = "Automation-only SP"
  provider     = databricks.account
}

resource "databricks_permission_assignment" "add_admin_spn" {
  principal_id = data.databricks_service_principal.sp.id
  permissions  = ["ADMIN"]
  provider     = databricks.workspace
}
```

In workspace context, adding account-level group to a workspace:

```hcl
# Use the account provider
data "databricks_group" "account_level" {
  display_name = "example-group"
  provider     = databricks.account
}

# Use the workspace provider
resource "databricks_permission_assignment" "this" {
  principal_id = data.databricks_group.account_level.id
  permissions  = ["USER"]
  provider     = databricks.workspace
}

data "databricks_group" "workspace_level" {
  display_name = "example-group"
  depends_on   = [databricks_permission_assignment.this]
  provider     = databricks.workspace
}

output "databricks_group_id" {
  value = data.databricks_group.workspace_level.id
}
```

### Assign using `user_name`, `group_name`, or `service_principal_name`

In workspace context, adding account-level user to a workspace:

```hcl
resource "databricks_permission_assignment" "add_user" {
  user_name   = "me@example.com"
  permissions = ["USER"]
  provider    = databricks.workspace
}
```

In workspace context, adding account-level service principal to a workspace:

```hcl
resource "databricks_permission_assignment" "add_admin_spn" {
  service_principal_name = "00000000-0000-0000-0000-000000000000"
  permissions  = ["ADMIN"]
  provider     = databricks.workspace
}
```

In workspace context, adding account-level group to a workspace:

```hcl
resource "databricks_permission_assignment" "this" {
  group_name  = "example-group"
  permissions = ["USER"]
  provider    = databricks.workspace
}
```

## Argument Reference

The following arguments are supported (exactly one of `principal_id`, `user_name`, `group_name`, or `service_principal_name` is required. Change of them triggers recreation):

* `principal_id` - Databricks ID of the user, service principal, or group. The principal ID can be retrieved using the account-level SCIM API, or using [databricks_user](../data-sources/user.md), [databricks_service_principal](../data-sources/service_principal.md) or [databricks_group](../data-sources/group.md) data sources with account API (and has to be an account admin). A more sensible approach is to retrieve the list of `principal_id` as outputs from another Terraform stack.
* `user_name` - the user name (email) to assign to a workspace.
* `service_principal_name` - the application ID of service principal to assign to a workspace.
* `group_name` - the group name to assign to a workspace.
* `permissions` (Required) - The list of workspace permissions to assign to the principal:
  * `"USER"` - Adds principal to the workspace `users` group. This gives basic workspace access.
  * `"ADMIN"` - Adds principal to the workspace `admins` group. This gives workspace admin privileges to manage users and groups, workspace configurations, and more.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the permission assignment - same as `principal_id`.
* `display_name` - the display name of the assigned principal.

## Import

The resource `databricks_permission_assignment` can be imported using the principal id:

```hcl
import {
  to = databricks_permission_assignment.this
  id = "<principal_id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_permission_assignment.this "<principal_id>"
```

## Related Resources

The following resources are used in the same context:

* [databricks_group](group.md) to manage [Account-level](https://docs.databricks.com/aws/en/admin/users-groups/groups) or [Workspace-level](https://docs.databricks.com/aws/en/admin/users-groups/workspace-local-groups) groups.
* [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_mws_permission_assignment](mws_permission_assignment.md) to manage permission assignment from an account context
