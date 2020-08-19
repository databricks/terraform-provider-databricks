# databricks_group_instance_profile Resource

**This resource has evolving API, which may change in future versions of provider.**

This resource allows you to attach instance profiles to groups created by the `databricks_group` resource.

-> **Note** Please only use this resource in conjunction with the `databricks_group` resource and **not** the `databricks_scim_group` resource.

-> **Note** You must be a Databricks administrator API token to make SCIM api calls. 

## Example Usage

```hcl
resource "databricks_instance_profile" "instance_profile" {
  instance_profile_arn = "my_instance_profile_arn"
  skip_validation = true
}
resource "databricks_group" "my_group" {
  display_name = "my_group_name"
}
resource "databricks_group_instance_profile" "my_group_instance_profile" {
 group_id = databricks_group.my_group.id
 instance_profile_id = databricks_instance_profile.instance_profile.id
}
```
## Argument Reference

The following arguments are supported:

* `group_id` - **(Required)** This is the id of the `databricks_group` resource.

* `instance_profile_id` -  **(Required)** This is the id of the `databricks_instance_profile` resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

*  `id` - The id for the `databricks_group_instance_profile` object which is in the format `<group_id>|<instance_profile_id>`.

## Import

-> **Note** Importing this resource is not currently supported.
