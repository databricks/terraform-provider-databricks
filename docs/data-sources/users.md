---
subcategory: "Security"
---

# databricks_users Data Source

-> This data source works with both the account-level and workspace-level provider. 

-> If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves information about multiple [databricks_user](../resources/user.md) resources.

## Example Usage

Adding a subset of users to a group

```hcl
data "databricks_users" "company_users" {
  filter = "userName co \"@domain.org\""
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

-> Attribute names and operators used in filters are case-insensitive. Find more information [here](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2).  

- `filter` - (Optional) Query by which the results have to be filtered. If not specified, all users will be returned. Supported operators are equals (`eq`), contains (`co`), starts with (`sw`), and not equals (`ne`). Additionally, simple expressions can be formed using logical operators `and` and `or`.

  **Examples:**
    - User whose `displayName` equals "john": 
    ```hcl
      filter = "displayName eq \"john\""
    ```
    - User whose `displayName` contains "john" or `userName` contains "@domain.org": 
    ```hcl
      filter = "displayName co \"john\" or userName co \"@domain.org\""
    ```

- `extra_attributes` - (Optional) A list of additional user attributes to include in the results. By default, the data source returns the following attributes: `id`, `userName`, `displayName`, and `externalId`. Use this argument to request additional attributes as needed. The list of all available attributes can be found in the [API reference](https://docs.databricks.com/api/workspace/users/list). 

## Attribute Reference

This data source exposes the following attributes:

- `users` - A list of users matching the specified criteria. Each user has the following attributes:
    - `id` - The ID of the user.
    - `userName` - The username of the user.
    - `emails` - All the emails associated with the Databricks user.
    - `name`
      - `givenName` - Given name of the Databricks user.
      - `familyName` - Family name of the Databricks user.
    - `displayName` - The display name of the user. 
    - `groups` - Indicates if the user is part of any groups. 
      - `$ref`
      - `value`
      - `display`
      - `primary`
      - `type`
    - `entitlements` - Entitlements assigned to the user. 
      - `$ref`
      - `value`
      - `display`
      - `primary`
      - `type`
    - `roles` - Indicates if the user has any associated roles.
      - `$ref`
      - `value`
      - `display`
      - `primary`
      - `type`
    - `schemas` - The schema of the user. 
    - `externalId` - Reserved for future use. 
    - `active` - Boolean that represents if this user is active. 

## Related Resources

The following resources are used in the same context:

- [**databricks_user**](../resources/user.md): Resource to manage individual users in Databricks.

- [**databricks_group**](../resources/group.md): Resource to manage groups in Databricks.

- [**databricks_group_member**](../resources/group_member.md): Resource to manage group memberships by adding users to groups.

- [**databricks_permissions**](../resources/permissions.md): Resource to manage access control in the Databricks workspace.

- [**databricks_current_user**](current_user.md): Data source to retrieve information about the user or service principal that is calling the Databricks REST API.
