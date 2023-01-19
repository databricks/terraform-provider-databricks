---
subcategory: "Security"
---
# databricks_group_instance_profile Resource

-> **Deprecated** Please migrate to [databricks_group_role](group_role.md).

This resource allows you to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_group](group.md).

## Example Usage

```hcl
resource "databricks_instance_profile" "instance_profile" {
  instance_profile_arn = "my_instance_profile_arn"
}

resource "databricks_group" "my_group" {
  display_name = "my_group_name"
}

resource "databricks_group_instance_profile" "my_group_instance_profile" {
  group_id            = databricks_group.my_group.id
  instance_profile_id = databricks_instance_profile.instance_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required) This is the id of the [group](group.md) resource.
* `instance_profile_id` -  (Required) This is the id of the [instance profile](instance_profile.md) resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id in the format `<group_id>|<instance_profile_id>`.

## Import

-> **Note** Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_aws_bucket_policy](../data-sources/aws_bucket_policy.md) data to configure a simple access policy for AWS S3 buckets, so that Databricks can access data in it.
* [databricks_cluster_policy](cluster_policy.md) to create a [databricks_cluster](cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_group](group.md) to manage [groups in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/groups.html) or [Account Console](https://accounts.cloud.databricks.com/) (for AWS deployments).
* [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_instance_pool](instance_pool.md) to manage [instance pools](https://docs.databricks.com/clusters/instance-pools/index.html) to reduce [cluster](cluster.md) start and auto-scaling times by maintaining a set of idle, ready-to-use instances.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_user_instance_profile](user_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_user](user.md).
