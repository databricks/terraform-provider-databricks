---
subcategory: "Security"
---
# databricks_service_principal_role Resource

This resource allows you to attach a role or [databricks_instance_profile](instance_profile.md) (AWS) to a [databricks_service_principal](service_principal.md).

-> This resource can be used with an account or workspace-level provider.

## Example Usage

Granting a service principal access to an instance profile

```hcl
resource "databricks_instance_profile" "instance_profile" {
  instance_profile_arn = "my_instance_profile_arn"
}

resource "databricks_service_principal" "this" {
  display_name = "My Service Principal"
}

resource "databricks_service_principal_role" "my_service_principal_instance_profile" {
  service_principal_id = databricks_service_principal.this.id
  role                 = databricks_instance_profile.instance_profile.id
}
```

Granting a service principal the Account Admin role.

-> This can only be used with an account-level provider.

```hcl
resource "databricks_service_principal" "tf_admin" {
  display_name = "Terraform Admin"
}

resource "databricks_service_principal_role" "tf_admin_account" {
  service_principal_id = databricks_service_principal.tf_admin.id
  role                 = "account_admin"
}
```

## Argument Reference

The following arguments are supported:

* `service_principal_id` - (Required) This is the id of the [service principal](service_principal.md) resource.
* `role` -  (Required) This is the role name, role id, or [instance profile](instance_profile.md) resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id in the format `<service_principal_id>|<role>`.

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_user_role](user_instance_profile.md) to attach role or [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_user](user.md).
* [databricks_group_instance_profile](group_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_group](group.md).
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_access_control_rule_set](access_control_rule_set.md#grant_rules) to attach other roles to account level resources.
