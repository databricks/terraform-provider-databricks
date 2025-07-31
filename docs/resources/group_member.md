---
subcategory: "Security"
---
# databricks_group_member Resource

This resource allows you to attach [users](user.md), [service_principal](service_principal.md), and [groups](group.md) as group members.

-> This resource can be used with an account or workspace-level provider.

To attach members to groups in the Databricks account, the provider must be configured with `host = "https://accounts.cloud.databricks.com"` on AWS deployments or `host = "https://accounts.azuredatabricks.net"` and authenticate using [AAD tokens](https://registry.terraform.io/providers/databricks/databricks/latest/docs#special-configurations-for-azure) on Azure deployments

## Example Usage

After the following example, Bradley would have direct membership in group B and transitive membership in group A.

```hcl
resource "databricks_group" "a" {
  display_name = "A"
}

resource "databricks_group" "b" {
  display_name = "B"
}

resource "databricks_group_member" "ab" {
  group_id  = databricks_group.a.id
  member_id = databricks_group.b.id
}

resource "databricks_user" "bradley" {
  user_name = "bradley@example.com"
}

resource "databricks_group_member" "bb" {
  group_id  = databricks_group.b.id
  member_id = databricks_user.bradley.id
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required) This is the `id` attribute (SCIM ID) of the [group](group.md) resource.
* `member_id` - (Required) This is the `id` attribute (SCIM ID) of the [group](group.md), [service principal](service_principal.md), or [user](user.md).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id for the `databricks_group_member` object which is in the format `<group_id>|<member_id>`.

## Import

You can import a `databricks_group_member` resource with name `my_group_member` like the following:

```hcl
import {
  to = databricks_group_member.my_group_member
  id = "<group_id>|<member_id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_group_member.my_group_member "<group_id>|<member_id>"
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_group](group.md) to manage [Account-level](https://docs.databricks.com/aws/en/admin/users-groups/groups) or [Workspace-level](https://docs.databricks.com/aws/en/admin/users-groups/workspace-local-groups) groups.
* [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
* [databricks_group_instance_profile](group_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_group](group.md).
* [databricks_ip_access_list](ip_access_list.md) to allow access from [predefined IP ranges](https://docs.databricks.com/security/network/ip-access-list.html).
* [databricks_service_principal](service_principal.md) to grant access to a workspace to an automation tool or application.
* [databricks_user](user.md) to [manage users](https://docs.databricks.com/administration-guide/users-groups/users.html), that could be added to [databricks_group](group.md) within the workspace.
* [databricks_user](../data-sources/user.md) data to retrieve information about [databricks_user](user.md).
* [databricks_user_instance_profile](user_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_user](user.md).
