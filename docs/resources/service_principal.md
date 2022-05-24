---
subcategory: "Security"
---
# databricks_service_principal Resource

Directly manage [Service Principals](https://docs.databricks.com/administration-guide/users-groups/service-principals.html) that could be added to [databricks_group](group.md) in Databricks workspace or account.

To create service principals in the Databricks account, the provider must be configured with `host = "https://accounts.azuredatabricks.net"` on AWS deployments or `host = "https://accounts.azuredatabricks.net"` and `auth_type  = "azure-cli"` on Azure deployments

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
  application_id = "00000000-0000-0000-0000-000000000000"
}

resource "databricks_group_member" "i-am-admin" {
  group_id  = data.databricks_group.admins.id
  member_id = databricks_service_principal.sp.id
}
```

Creating service principal with cluster create permissions:

```hcl
resource "databricks_service_principal" "sp" {
  application_id       = "00000000-0000-0000-0000-000000000000"
  display_name         = "Example service principal"
  allow_cluster_create = true
}
```

Creating service principal in AWS Databricks account:
```hcl
// initialize provider at account-level
provider "databricks" {
  alias    = "mws"
  host     = "https://accounts.cloud.databricks.com"
  account_id = "00000000-0000-0000-0000-000000000000"
  username = var.databricks_account_username
  password = var.databricks_account_password
}

resource "databricks_service_principal" "sp" {
  provider     = databricks.mws
  display_name = "Automation-only SP"
}
```

Creating group in Azure Databricks account:
```hcl
// initialize provider at Azure account-level
provider "databricks" {
  alias      = "azure_account"
  host       = "https://accounts.azuredatabricks.net"
  account_id = "00000000-0000-0000-0000-000000000000"
  auth_type  = "azure-cli"
}

resource "databricks_service_principal" "sp" {
  provider       = databricks.azure_account
  application_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

-> `application_id` is required on Azure Databricks and is not allowed on other clouds. `display_name` is required on all clouds except Azure.

The following arguments are available:

* `application_id` - This is the application id of the given service principal and will be their form of access and identity. On other clouds than Azure this value is auto-generated.
* `display_name` - (Required) This is an alias for the service principal and can be the full name of the service principal.
* `external_id` - (Optional) ID of the service principal in an external identity provider.
* `allow_cluster_create` -  (Optional) Allow the service principal to have [cluster](cluster.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and `cluster_id` argument. Everyone without `allow_cluster_create` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within the boundaries of that specific policy.
* `allow_instance_pool_create` -  (Optional) Allow the service principal to have [instance pool](instance_pool.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.
* `databricks_sql_access` - (Optional) This is a field to allow the group to have access to [Databricks SQL](https://databricks.com/product/databricks-sql) feature through [databricks_sql_endpoint](sql_endpoint.md).
* `workspace_access` - (Optional) This is a field to allow the group to have access to Databricks Workspace.
* `active` - (Optional) Either service principal is active or not. True by default, but can be set to false in case of service principal deactivation with preserving service principal assets.
* `force` - (Optional) Ignore `cannot create service principal: Service principal with application ID X already exists` errors and implicitly import the specific service principal into Terraform state, enforcing entitlements defined in the instance of resource. _This functionality is experimental_ and is designed to simplify corner cases, like Azure Active Directory synchronisation.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the service principal.

## Import

The resource scim service principal can be imported using id:

```bash
$ terraform import databricks_service_principal.me <service-principal-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_group](group.md) to manage [groups in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/groups.html) or [Account Console](https://accounts.cloud.databricks.com/) (for AWS deployments).
* [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_permissions](permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
* [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.
