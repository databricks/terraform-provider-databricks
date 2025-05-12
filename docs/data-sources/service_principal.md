---
subcategory: "Security"
---

# databricks_service_principal Data Source

Retrieves information about [databricks_service_principal](../resources/service_principal.md).

-> This data source can be used with an account or workspace-level provider.

## Example Usage

Adding service principal `11111111-2222-3333-4444-555666777888` to administrative group

```hcl
data "databricks_group" "admins" {
  display_name = "admins"
}

data "databricks_service_principal" "spn" {
  application_id = "11111111-2222-3333-4444-555666777888"
}

resource "databricks_group_member" "my_member_a" {
  group_id  = data.databricks_group.admins.id
  member_id = data.databricks_service_principal.spn.id
}
```

## Argument Reference

Data source allows you to pick service principals by one of the following attributes (only one of them):

- `application_id` - (Required if `display_name` isn't used) ID of the service principal. The service principal must exist before this resource can be retrieved.
- `display_name` - (Required if `application_id` isn't used) Exact display name of the service principal. The service principal must exist before this resource can be retrieved.  In case if there are several service principals with the same name, an error is thrown.

## Attribute Reference

Data source exposes the following attributes:

- `id` - The id of the service principal.
- `external_id` - ID of the service principal in an external identity provider.
- `display_name` - Display name of the [service principal](../resources/service_principal.md), e.g. `Foo SPN`.
- `home` - Home folder of the [service principal](../resources/service_principal.md), e.g. `/Users/11111111-2222-3333-4444-555666777888`.
- `repos` - Repos location of the [service principal](../resources/service_principal.md), e.g. `/Repos/11111111-2222-3333-4444-555666777888`.
- `active` - Whether service principal is active or not.

* `acl_principal_id` - identifier for use in [databricks_access_control_rule_set](../resources/access_control_rule_set.md), e.g. `servicePrincipals/00000000-0000-0000-0000-000000000000`.

## Related Resources

The following resources are used in the same context:

- [End to end workspace management](../guides/workspace-management.md) guide.
- [databricks_current_user](current_user.md) data to retrieve information about [databricks_user](../resources/user.md) or [databricks_service_principal](../resources/service_principal.md), that is calling Databricks REST API.
- [databricks_group](../resources/group.md) to manage [groups in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/groups.html) or [Account Console](https://accounts.cloud.databricks.com/) (for AWS deployments).
- [databricks_group](group.md) data to retrieve information about [databricks_group](../resources/group.md) members, entitlements and instance profiles.
- [databricks_group_instance_profile](../resources/group_instance_profile.md) to attach [databricks_instance_profile](../resources/instance_profile.md) (AWS) to [databricks_group](../resources/group.md).
- [databricks_group_member](../resources/group_member.md) to attach [users](../resources/user.md) and [groups](../resources/group.md) as group members.
- [databricks_permissions](../resources/permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
- [databricks_service principal](../resources/service_principal.md) to manage [service principals](../resources/service_principal.md)
