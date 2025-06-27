---
subcategory: "Security"
---
# databricks_user_instance_profile Resource

-> **Deprecated** Please rewrite with [databricks_user_role](user_role.md). This resource will be removed in v0.5.x

This resource allows you to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_user](user.md).

## Example Usage

```hcl
resource "databricks_instance_profile" "instance_profile" {
  instance_profile_arn = "my_instance_profile_arn"
}

resource "databricks_user" "my_user" {
  user_name = "me@example.com"
}

resource "databricks_user_instance_profile" "my_user_instance_profile" {
  user_id             = databricks_user.my_user.id
  instance_profile_id = databricks_instance_profile.instance_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) This is the id of the [user](user.md) resource.
* `instance_profile_id` -  (Required) This is the id of the [instance profile](instance_profile.md) resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id in the format `<user_id>|<instance_profile_id>`.

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_group_instance_profile](group_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_group](group.md).
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_user](user.md) to [manage users](https://docs.databricks.com/administration-guide/users-groups/users.html), that could be added to [databricks_group](group.md) within the workspace.
* [databricks_user](../data-sources/user.md) data to retrieve information about [databricks_user](user.md).
