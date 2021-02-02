# databricks_group Data Source

Retrieves information about [databricks_group](../resources/group.md) members, entitlements and instance profiles.

!> [Do not use](https://www.terraform.io/docs/configuration/data-sources.html#data-resource-dependencies) `depends_on` meta-argument within data sources, unless you explicitly want to have dependent resources updated each apply.

## Example Usage

Adding user to administrative group

```hcl
data "databricks_group" "admins" {
    display_name = "admins"
}

resource "databricks_user" "me" {
  user_name    = "me@example.com"
}

resource "databricks_group_member" "my_member_a" {
  group_id = data.databricks_group.admins.id
  member_id = databricks_user.me.id
}
```

## Argument Reference

Data source allows you to pick groups by the following attributes

* `display_name` - (Required) Display name of the group. The group must exist before this resource can be planned.
* `recursive` - (Optional) Collect information for all nested groups. *Defaults to true.*

## Attribute Reference

Data source exposes the following attributes:

* `id` -  The id for the group object.
* `members` - Set of [user](../resources/user.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `groups` - Set of [group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.
* `instance_profiles` - Set of [instance profile](../resources/instance_profile.md) ARNs, that can be modified by [databricks_group_instance_profile](../resources/group_instance_profile.md) resource.
* `allow_cluster_create` - True if group members can create [clusters](../resources/cluster.md)
* `allow_instance_pool_create` - True if group members can create [instance pools](../resources/instance_pool.md)
