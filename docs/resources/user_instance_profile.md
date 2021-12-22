---
subcategory: "Security"
---
# databricks_user_instance_profile Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

This resource allows you to attach instance profiles to users.

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

*  `id` - The id in the format `<user_id>|<instance_profile_id>`.

## Import

-> **Note** Importing this resource is not currently supported.
