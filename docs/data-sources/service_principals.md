---
subcategory: "Security"
---

# databricks_service_principals Data Source

Retrieves `application_ids` of all [databricks_service_principal](../resources/service_principal.md) based on their `display_name`

-> This data source can be used with an account or workspace-level provider.

## Example Usage

Adding all service principals of which display name contains `my-spn` to admin group

```hcl
data "databricks_group" "admins" {
  display_name = "admins"
}

data "databricks_service_principals" "spns" {
  display_name_contains = "my-spn"
}

data "databricks_service_principal" "spn" {
  for_each       = toset(data.databricks_service_principals.spns.application_ids)
  application_id = each.value
}

resource "databricks_group_member" "my_member_spn" {
  for_each  = toset(data.databricks_service_principals.spns.application_ids)
  group_id  = data.databricks_group.admins.id
  member_id = data.databricks_service_principal.spn[each.value].sp_id
}
```

## Argument Reference

Data source allows you to pick service principals by the following attributes

- `display_name_contains` - (Optional) Only return [databricks_service_principal](service_principal.md) display name that match the given name string

## Attribute Reference

Data source exposes the following attributes:

- `application_ids` - List of `application_ids` of service principals Individual service principal can be retrieved using [databricks_service_principal](service_principal.md) data source

## Related Resources

The following resources are used in the same context:

- [End to end workspace management](../guides/workspace-management.md) guide.
- [databricks_current_user](current_user.md) data to retrieve information about [databricks_user](../resources/user.md) or [databricks_service_principal](../resources/service_principal.md), that is calling Databricks REST API.
- [databricks_group](../resources/group.md) to manage [Account-level](https://docs.databricks.com/aws/en/admin/users-groups/groups) or [Workspace-level](https://docs.databricks.com/aws/en/admin/users-groups/workspace-local-groups) groups.
- [databricks_group](group.md) data to retrieve information about [databricks_group](../resources/group.md) members, entitlements and instance profiles.
- [databricks_group_instance_profile](../resources/group_instance_profile.md) to attach [databricks_instance_profile](../resources/instance_profile.md) (AWS) to [databricks_group](../resources/group.md).
- [databricks_group_member](../resources/group_member.md) to attach [users](../resources/user.md) and [groups](../resources/group.md) as group members.
- [databricks_permissions](../resources/permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
- [databricks_service principal](../resources/service_principal.md) to manage [service principals](../resources/service_principal.md)
