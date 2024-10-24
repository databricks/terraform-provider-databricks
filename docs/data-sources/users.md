---
subcategory: "Security"
---

# databricks_users Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves information about multiple [databricks_user](../resources/user.md) resources.

## Example Usage

Adding a subset of users to a group

```hcl
data "databricks_users" "company_users" {
  user_name_contains = "@domain.org"
}

resource "databricks_group" "data_users_group" {
  display_name = "Data Users"
}

resource "databricks_group_member" "add_users_to_group" {
  for_each = { for user in data.databricks_users.company_users.users : user.id => user }
  group_id  = databricks_group.data_users_group.id
  member_id = each.value.id
}
```

## Argument Reference

This data source allows you to filter the list of users using the following optional arguments: 

- `display_name_contains` - (Optional) A substring to filter users by their display name. Only users whose display names contain this substring will be included in the results.
- `user_name_contains` - (Optional) A substring to filter users by their username. Only users whose usernames contain this substring will be included in the results.

->**Note** You can specify **exactly one** of `display_name_contains` or `user_name_contains`. If neither is specified, all users will be returned.

## Attribute Reference

This data source exposes the following attributes:

- `users` - A list of users matching the specified criteria. Each user has the following attributes:
    - `id` - The ID of the user.
    - `user_name` - The username of the user.
    - `display_name` - The display name of the user. 

## Related Resources

The following resources are used in the same context:

- [**databricks_user**](../resources/user.md): Resource to manage individual users in Databricks.

- [**databricks_group**](../resources/group.md): Resource to manage groups in Databricks.

- [**databricks_group_member**](../resources/group_member.md): Resource to manage group memberships by adding users to groups.

- [**databricks_permissions**](../resources/permissions.md): Resource to manage access control in the Databricks workspace.

- [**databricks_current_user**](current_user.md): Data source to retrieve information about the user or service principal that is calling the Databricks REST API.
