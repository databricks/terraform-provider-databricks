---
subcategory: "Security"
---

# databricks_user Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

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
