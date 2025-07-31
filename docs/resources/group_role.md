---
subcategory: "Security"
---
# databricks_group_role Resource

This resource allows you to attach a role to [databricks_group](group.md). This role could be a pre-defined role such as account admin, or an instance profile ARN.

-> This resource can be used with an account or workspace-level provider.

## Example Usage

Attach an instance profile to a group

```hcl
resource "databricks_instance_profile" "instance_profile" {
  instance_profile_arn = "my_instance_profile_arn"
}

resource "databricks_group" "my_group" {
  display_name = "my_group_name"
}

resource "databricks_group_role" "my_group_instance_profile" {
  group_id = databricks_group.my_group.id
  role     = databricks_instance_profile.instance_profile.id
}
```

Attach account admin role to an account-level group

```hcl
provider "databricks" {
  host          = "https://accounts.cloud.databricks.com"
  account_id    = var.databricks_account_id
  client_id     = var.client_id
  client_secret = var.client_secret
}

resource "databricks_group" "my_group" {
  display_name = "my_group_name"
}

resource "databricks_group_role" "my_group_account_admin" {
  group_id = databricks_group.my_group.id
  role     = "account_admin"
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required) This is the id of the [group](group.md) resource.
* `role` - (Required) Either a role name or the ARN/ID of the [instance profile](instance_profile.md) resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id for the `databricks_group_role` object which is in the format `<group_id>|<role>`.

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_aws_bucket_policy](../data-sources/aws_bucket_policy.md) data to configure a simple access policy for AWS S3 buckets, so that Databricks can access data in it.
* [databricks_cluster_policy](cluster_policy.md) to create a [databricks_cluster](cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_group](group.md) to manage [Account-level](https://docs.databricks.com/aws/en/admin/users-groups/groups) or [Workspace-level](https://docs.databricks.com/aws/en/admin/users-groups/workspace-local-groups) groups.
* [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_instance_pool](instance_pool.md) to manage [instance pools](https://docs.databricks.com/clusters/instance-pools/index.html) to reduce [cluster](cluster.md) start and auto-scaling times by maintaining a set of idle, ready-to-use instances.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_user_instance_profile](user_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_user](user.md).
