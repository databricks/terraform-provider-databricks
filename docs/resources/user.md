# databricks_user Resource

Directly creates user, that could be added to [databricks_group](group.md) within workspace. Upon user creation user will receive a password reset email.

## Example Usage

Creating regular user:

```hcl
resource "databricks_user" "me" {
  user_name = "me@example.com"
}
```

Creating user with administrative permissions - referencing special `admins` [databricks_group](../data-sources/group.md) in [databricks_group_member](group_member.md) resource:

```hcl
data "databricks_group" "admins" {
    display_name = "admins"
}

resource "databricks_user" "me" {
  user_name    = "me@example.com"
}

resource "databricks_group_member" "i-am-admin" {
  group_id = data.databricks_group.admins.id
  member_id = databricks_user.me.id
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
* `active` - (Optional) Either user is active or not. True by default, but can be set to false in case of user deactivation with preserving user assets.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the user.

## Import

The resource scim user can be imported using id:

```bash
$ terraform import databricks_user.me <user-id>
```