---
subcategory: "Security"
---

# databricks_service_principal Resource

Directly manage [Service Principals](https://docs.databricks.com/administration-guide/users-groups/service-principals.html) that could be added to [databricks_group](group.md) in Databricks account or workspace.

-> This resource can be used with an account or workspace-level provider.

There are different types of service principals:

* Databricks-managed - exists only inside the Databricks platform (all clouds) and couldn't be used for accessing non-Databricks services.
* Azure-managed - existing Azure service principal (enterprise application) is registered inside Databricks.  It could be used to work with other Azure services.

-> To assign account level service principals to workspace use [databricks_mws_permission_assignment](mws_permission_assignment.md).

-> Entitlements, like, `allow_cluster_create`, `allow_instance_pool_create`, `databricks_sql_access`, `workspace_access` applicable only for workspace-level service principals. Use [databricks_entitlements](entitlements.md) resource to assign entitlements inside a workspace to account-level service principals.

The default behavior when deleting a `databricks_service_principal` resource depends on whether the provider is configured at the workspace-level or account-level. When the provider is configured at the workspace-level, the service principal will be deleted from the workspace. When the provider is configured at the account-level, the service principal will be deactivated but not deleted. When the provider is configured at the account level, to delete the service principal from the account when the resource is deleted, set `disable_as_user_deletion = false`. Conversely, when the provider is configured at the account-level, to deactivate the service principal when the resource is deleted, set `disable_as_user_deletion = true`.

## Example Usage

Creating regular Databricks-managed service principal:

```hcl
resource "databricks_service_principal" "sp" {
  display_name = "Admin SP"
}
```

Creating service principal with administrative permissions - referencing special `admins` [databricks_group](../data-sources/group.md) in [databricks_group_member](group_member.md) resource:

```hcl
data "databricks_group" "admins" {
  display_name = "admins"
}

resource "databricks_service_principal" "sp" {
  display_name = "Admin SP"
}

resource "databricks_group_member" "i-am-admin" {
  group_id  = data.databricks_group.admins.id
  member_id = databricks_service_principal.sp.id
}
```

Creating Azure-managed service principal with cluster create permissions:

```hcl
resource "databricks_service_principal" "sp" {
  application_id       = "00000000-0000-0000-0000-000000000000"
  display_name         = "Example service principal"
  allow_cluster_create = true
}
```

Creating Databricks-managed service principal in AWS Databricks account:

```hcl
// initialize provider at account-level
provider "databricks" {
  alias         = "account"
  host          = "https://accounts.cloud.databricks.com"
  account_id    = "00000000-0000-0000-0000-000000000000"
  client_id     = var.client_id
  client_secret = var.client_secret
}

resource "databricks_service_principal" "sp" {
  provider     = databricks.account
  display_name = "Automation-only SP"
}
```

Creating Azure-managed service principal in Azure Databricks account:

```hcl
// initialize provider at Azure account-level
provider "databricks" {
  alias      = "account"
  host       = "https://accounts.azuredatabricks.net"
  account_id = "00000000-0000-0000-0000-000000000000"
  auth_type  = "azure-cli"
}

resource "databricks_service_principal" "sp" {
  provider       = databricks.account
  application_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

-> `application_id` is required on Azure Databricks when using Azure-managed service principals and is not allowed for Databricks-managed service principals. `display_name` is required on all clouds when using Databricks-managed service principals, and optional for Azure Databricks.

The following arguments are available:

- `application_id` This is the Azure Application ID of the given Azure service principal and will be their form of access and identity. For Databricks-managed service principals this value is auto-generated.
- `display_name` - (Required for Databricks-managed service principals) This is an alias for the service principal and can be the full name of the service principal.
- `external_id` - (Optional) ID of the service principal in an external identity provider.
- `allow_cluster_create` - (Optional) Allow the service principal to have [cluster](cluster.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and `cluster_id` argument. Everyone without `allow_cluster_create` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within the boundaries of that specific policy.
- `allow_instance_pool_create` - (Optional) Allow the service principal to have [instance pool](instance_pool.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.
- `databricks_sql_access` - (Optional) This is a field to allow the group to have access to [Databricks SQL](https://databricks.com/product/databricks-sql) feature through [databricks_sql_endpoint](sql_endpoint.md).
- `workspace_access` - (Optional) This is a field to allow the group to have access to Databricks Workspace.
- `active` - (Optional) Either service principal is active or not. True by default, but can be set to false in case of service principal deactivation with preserving service principal assets.
- `force` - (Optional) Ignore `cannot create service principal: Service principal with application ID X already exists` errors and implicitly import the specified service principal into Terraform state, enforcing entitlements defined in the instance of resource. _This functionality is experimental_ and is designed to simplify corner cases, like Azure Active Directory synchronisation.
- `force_delete_repos` - (Optional) This flag determines whether the service principal's repo directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.
- `force_delete_home_dir` - (Optional) This flag determines whether the service principal's home directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.
- `disable_as_user_deletion` - (Optional) Deactivate the service principal when deleting the resource, rather than deleting the service principal entirely. Defaults to `true` when the provider is configured at the account-level and `false` when configured at the workspace-level. This flag is exclusive to force_delete_repos and force_delete_home_dir flags. 

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - Canonical unique identifier for the service principal (SCIM ID).
- `home` - Home folder of the service principal, e.g. `/Users/00000000-0000-0000-0000-000000000000`.
- `repos` - Personal Repos location of the service principal, e.g. `/Repos/00000000-0000-0000-0000-000000000000`.
- `acl_principal_id` - identifier for use in [databricks_access_control_rule_set](access_control_rule_set.md), e.g. `servicePrincipals/00000000-0000-0000-0000-000000000000`.

## Import

The resource scim service principal can be imported using its SCIM id, for example `2345678901234567`. To get the service principal ID, call [Get service principals](https://docs.databricks.com/dev-tools/api/latest/scim/scim-sp.html#get-service-principals).

```hcl
import {
  to = databricks_service_principal.me
  id = "<service-principal-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_service_principal.me <service-principal-id>
```

## Related Resources

The following resources are often used in the same context:

- [End to end workspace management](../guides/workspace-management.md) guide.
- [databricks_group](group.md) to manage [groups in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/groups.html) or [Account Console](https://accounts.cloud.databricks.com/) (for AWS deployments).
- [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
- [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
- [databricks_permissions](permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
- [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](<https://docs.databricks>.
- [databricks-service-principal-secret](service_principal_secret.md) to manage secrets for the service principal (only for AWS deployments)
