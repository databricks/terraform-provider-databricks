---
subcategory: "Security"
---
# databricks_service_principal_instance_profile Resource

This resource allows you to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_service_principal](service_principal.md).

## Example Usage

```hcl
resource "databricks_instance_profile" "instance_profile" {
  instance_profile_arn = "my_instance_profile_arn"
}

resource "databricks_service_principal" "this" {
  display_name = "My Service Principal"
}

resource "databricks_service_principal_instance_profile" "my_service_principal_instance_profile" {
  service_principal_id      = databricks_service_principal.this.id
  role = databricks_instance_profile.instance_profile.id
}
```
## Argument Reference

The following arguments are supported:

* `application_id` - (Required) This is the id of the [service principal](service_principal.md) resource.
* `instance_profile_id` -  (Required) This is the id of the [instance profile](instance_profile.md) resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

*  `id` - The id in the format `<application_id>|<instance_profile_id>`.

## Import

-> **Note** Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_user_instance_profile](user_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_user](user.md).
* [databricks_group_instance_profile](group_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_group](group.md).
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
