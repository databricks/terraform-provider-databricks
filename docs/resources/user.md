---
subcategory: "Security"
---
# databricks_user Resource

This resource allows you to manage [users in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/users.html), [Databricks Account Console](https://accounts.cloud.databricks.com/) or [Azure Databricks Account Console](https://accounts.azuredatabricks.net). You can also [associate](group_member.md) Databricks users to [databricks_group](group.md). Upon user creation the user will receive a password reset email. You can also get information about caller identity using [databricks_current_user](../data-sources/current_user.md) data source.

To create users in the Databricks account, the provider must be configured with `host = "https://accounts.cloud.databricks.com"` on AWS deployments or `host = "https://accounts.azuredatabricks.net"` and authenticate using AAD tokens on Azure deployments

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
  user_name = "me@example.com"
}

resource "databricks_group_member" "i-am-admin" {
  group_id  = data.databricks_group.admins.id
  member_id = databricks_user.me.id
}
```

Creating user with cluster create permissions:

```hcl
resource "databricks_user" "me" {
  user_name            = "me@example.com"
  display_name         = "Example user"
  allow_cluster_create = true
}
```

Creating user in AWS Databricks account:
```hcl
// initialize provider at account-level
provider "databricks" {
  alias    = "mws"
  host     = "https://accounts.cloud.databricks.com"
  account_id = "00000000-0000-0000-0000-000000000000"
  username = var.databricks_account_username
  password = var.databricks_account_password
}

resource "databricks_user" "account_user" {
  provider     = databricks.mws
  user_name    = "me@example.com"
  display_name = "Example user"
}
```

Creating user in Azure Databricks account:
```hcl
// initialize provider at Azure account-level
provider "databricks" {
  alias      = "azure_account"
  host       = "https://accounts.azuredatabricks.net"
  account_id = "00000000-0000-0000-0000-000000000000"
  auth_type  = "azure-cli"
}

resource "databricks_user" "account_user" {
  provider     = databricks.mws
  user_name    = "me@example.com"
  display_name = "Example user"
}
```

## Argument Reference

The following arguments are available:

* `user_name` - (Required) This is the username of the given user and will be their form of access and identity.
* `display_name` - (Optional) This is an alias for the username that can be the full name of the user.
* `external_id` - (Optional) ID of the user in an external identity provider.
* `allow_cluster_create` -  (Optional) Allow the user to have [cluster](cluster.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and `cluster_id` argument. Everyone without `allow_cluster_create` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.
* `allow_instance_pool_create` -  (Optional) Allow the user to have [instance pool](instance_pool.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.
* `databricks_sql_access` - (Optional) This is a field to allow the group to have access to [Databricks SQL](https://databricks.com/product/databricks-sql) feature in User Interface and through [databricks_sql_endpoint](sql_endpoint.md).
* `active` - (Optional) Either user is active or not. True by default, but can be set to false in case of user deactivation with preserving user assets.
* `force` - (Optional) Ignore `cannot create user: User with username X already exists` errors and implicitly import the specific user into Terraform state, enforcing entitlements defined in the instance of resource. _This functionality is experimental_ and is designed to simplify corner cases, like Azure Active Directory synchronisation.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the user.

## Import

The resource scim user can be imported using id:

```bash
$ terraform import databricks_user.me <user-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_group](group.md) to manage [groups in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/groups.html) or [Account Console](https://accounts.cloud.databricks.com/) (for AWS deployments).
* [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
* [databricks_group_instance_profile](group_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_group](group.md).
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_user](../data-sources/user.md) data to retrieve information about [databricks_user](user.md).
