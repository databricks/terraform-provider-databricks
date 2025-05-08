---
subcategory: "Security"
---
# databricks_entitlements Resource

This resource allows you to set entitlements to existing [databricks_users](user.md), [databricks_group](group.md) or [databricks_service_principal](service_principal.md).

-> This resource can only be used with a workspace-level provider!

-> You must define entitlements of a principal using either `databricks_entitlements` or directly within one of [databricks_users](user.md), [databricks_group](group.md) or [databricks_service_principal](service_principal.md). Having entitlements defined in both resources will result in non-deterministic behaviour.

## Example Usage

Setting entitlements for a regular user:

```hcl
data "databricks_user" "me" {
  user_name = "me@example.com"
}

resource "databricks_entitlements" "me" {
  user_id                    = data.databricks_user.me.id
  allow_cluster_create       = true
  allow_instance_pool_create = true
}
```

Setting entitlements for a service principal:

```hcl
data "databricks_service_principal" "this" {
  application_id = "11111111-2222-3333-4444-555666777888"
}

resource "databricks_entitlements" "this" {
  service_principal_id       = data.databricks_service_principal.this.sp_id
  allow_cluster_create       = true
  allow_instance_pool_create = true
}
```

Setting entitlements to all users in a workspace - referencing special `users` [databricks_group](../data-sources/group.md)

```hcl
data "databricks_group" "users" {
  display_name = "users"
}

resource "databricks_entitlements" "workspace-users" {
  group_id                   = data.databricks_group.users.id
  allow_cluster_create       = true
  allow_instance_pool_create = true
}
```

## Argument Reference

The following arguments are available to specify the identity you need to enforce entitlements. You must specify exactly one of those arguments otherwise resource creation will fail.

* `user_id` -  Canonical unique identifier for the user.
* `group_id` - Canonical unique identifier for the group.
* `service_principal_id` - Canonical unique identifier for the service principal.

The following entitlements are available.

* `allow_cluster_create` -  (Optional) Allow the principal to have [cluster](cluster.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and `cluster_id` argument. Everyone without `allow_cluster_create` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.
* `allow_instance_pool_create` -  (Optional) Allow the principal to have [instance pool](instance_pool.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.
* `databricks_sql_access` - (Optional) This is a field to allow the principal to have access to [Databricks SQL](https://databricks.com/product/databricks-sql) feature in User Interface and through [databricks_sql_endpoint](sql_endpoint.md).
* `workspace_access` - (Optional) This is a field to allow the principal to have access to Databricks Workspace.

## Import

The resource can be imported using a synthetic identifier. Examples of valid synthetic identifiers are:

* `user/user_id` - user `user_id`.
* `group/group_id` - group `group_id`.
* `spn/spn_id` - service principal `spn_id`.

```bash
terraform import databricks_entitlements.me user/<user-id>
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
