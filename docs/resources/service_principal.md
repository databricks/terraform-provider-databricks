# databricks_service_principal Resource

Directly creates service principal, that could be added to [databricks_group](group.md) within workspace.

## Example Usage

Creating regular service principal:

```hcl
resource "databricks_service_principal" "sp" {
  application_id = "00000000-0000-0000-0000-000000000000"
}
```

Creating service principal with administrative permissions - referencing special `admins` [databricks_group](../data-sources/group.md) in [databricks_group_member](group_member.md) resource:

```hcl
data "databricks_group" "admins" {
    display_name = "admins"
}

resource "databricks_service_principal" "sp" {
  application_id    = "00000000-0000-0000-0000-000000000000"
}

resource "databricks_group_member" "i-am-admin" {
  group_id = data.databricks_group.admins.id
  member_id = databricks_service_principal.sp.id
}
```

Creating service principal with cluster create permissions:

```hcl
resource "databricks_service_principal" "sp" {
  application_id    = "00000000-0000-0000-0000-000000000000"
  display_name = "Example service principal"
  allow_cluster_create = true
}
```

## Argument Reference

The following arguments are available:

* `application_id` - (Required) This is the application id of the given service principal and will be their form of access and identity.
* `display_name` - (Optional) This is an alias for the service principal can be the full name of the service principal.
* `allow_cluster_create` -  (Optional) Allow the service principal to have [cluster](cluster.md) create priviliges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and `cluster_id` argument. Everyone without `allow_cluster_create` arugment set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.
* `allow_instance_pool_create` -  (Optional) Allow the service principal to have [instance pool](instance_pool.md) create priviliges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.
* `active` - (Optional) Either service principal is active or not. True by default, but can be set to false in case of service principal deactivation with preserving service principal assets.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the service principal.

## Import

The resource scim service principal can be imported using id:

```bash
$ terraform import databricks_service_principal.me <service-principal-id>
```
