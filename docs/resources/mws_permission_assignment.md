---
subcategory: "Security"
---
# databricks_mws_permission_assignment Resource

These resources are invoked in the account context. Permission Assignment Account API endpoints are restricted to account admins. Provider must have `account_id` attribute configured. Account Id that could be found in the bottom left corner of Accounts Console

## Example Usage

In account context, adding account-level group to a workspace:

```hcl
provider "databricks" {
  // <other properties>
  account_id = "<databricks account id>"
}

resource "databricks_group" "data_eng" {
  display_name = "Data Engineering"
}

resource "databricks_mws_permission_assignment" "add_admin_group" {
  workspace_id = databricks_mws_workspaces.this.workspace_id
  principal_id = databricks_group.data_eng.id
  permissions  = ["ADMIN"]
}
```

In account context, adding account-level user to a workspace:

```hcl
provider "databricks" {
  // <other properties>
  account_id = "<databricks account id>"
}

resource "databricks_user" "me" {
  user_name = "me@example.com"
}

resource "databricks_mws_permission_assignment" "add_user" {
  workspace_id = databricks_mws_workspaces.this.workspace_id
  principal_id = databricks_user.me.id
  permissions  = ["USER"]
}
```

In account context, adding account-level service principal to a workspace:

```hcl
provider "databricks" {
  // <other properties>
  account_id = "<databricks account id>"
}

resource "databricks_service_principal" "sp" {
  display_name = "Automation-only SP"
}

resource "databricks_mws_permission_assignment" "add_admin_spn" {
  workspace_id = databricks_mws_workspaces.this.workspace_id
  principal_id = databricks_service_principal.sp.id
  permissions  = ["ADMIN"]
}
```

## Argument Reference

The following arguments are required:

* `workspace_id` - Databricks workspace ID
* `principal_id` - Databricks ID of the user, service principal, or group. The principal ID can be retrieved using the SCIM API.
* `permissions` - The list of workspace permissions to assign to the principal:
  * `“USER”` - Can access the workspace with basic privileges.
  * `“ADMIN”` - Can access the workspace and has workspace admin privileges to manage users and groups, workspace configurations, and more.

## Import

The resource `databricks_mws_permission_assignment` can be imported using the workspace id and principal id

```bash
terraform import databricks_mws_permission_assignment.this workspace_id|principal_id
```

## Related Resources

The following resources are used in the same context:

* [databricks_group](group.md) to manage [groups in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/groups.html) or [Account Console](https://accounts.cloud.databricks.com/) (for AWS deployments).
* [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_permission_assignment](permission_assignment.md) to manage permission assignment from a workspace context
