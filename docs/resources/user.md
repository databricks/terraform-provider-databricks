# databricks_user Resource

Directly creates user, that could be added to [databricks_group](group.md) within workspace. Upon user creation user will receive a password reset email.

## Example Usage

Creating regular user:

```hcl
resource "databricks_user" "me" {
  user_name = "me@example.com"
}
```

Creating user with cluster create permissions:

```hcl
resource "databricks_user" "me" {
  user_name    = "me@example.com"
  display_name = "Example user"
  allow_cluster_create = true
}
```

## Argument Reference

The following arguments are available:

* `user_name` - (Required) This is the username of the given user and will be their form of access and identity.
* `display_name` - (Optional) This is an alias for the username can be the full name of the user.
* `allow_cluster_create` -  (Optional) Allow the user to have [cluster](cluster.md) create priviliges. Defaults to false.
* `allow_instance_pool_create` -  (Optional) Allow the user to have [instance pool](instance_pool.md) create priviliges. Defaults to false.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the user.

## Import

The resource scim user can be imported using id:

```bash
$ terraform import databricks_user.me <user-id>
```