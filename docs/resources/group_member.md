# databricks_group_member Resource

This resource allows you to attach members to groups created by the [group](group.md) resource.

-> **Note** Please only use this resource in conjunction with the [group](group.md) resource and **not** the `databricks_scim_group` resource.

## Example Usage

```hcl
resource "databricks_group" "my_group" {
  display_name = "my_group_1"
}
resource "databricks_group" "my_sub_group_a" {
  display_name = "my_sub_group_a"
}
resource "databricks_group" "my_sub_group_b" {
  display_name = "my_sub_group_b"
}
resource "databricks_group_member" "my_member_a" {
 group_id = databricks_group.my_group.id
 member_id = databricks_group.my_sub_group_a.id
}
resource "databricks_group_member" "my_member_b" {
 group_id = databricks_group.my_group.id
 member_id = databricks_group.my_sub_group_b.id
}
```
## Argument Reference

The following arguments are supported:

* `group_id` - (Required) This is the id of the [group](group.md) resource.
* `member_id` - (Required) This is the id of the [group](group.md) or `databricks_scim_user` resource. 
>Members can be groups or users created by the scim api.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id for the `databricks_group_member` object which is in the format `<group_id>|<member_id>`.

## Import

-> **Note** Importing this resource is not currently supported.
