---
subcategory: "Security"
---
# databricks_group Resource

This resource allows you to manage both [account groups and workspace-local groups](https://docs.databricks.com/administration-guide/users-groups/groups.html). You can use the [databricks_group_member resource](group_member.md) to assign Databricks users, [service principals](service_principal.md) as well as other groups as members of the group. This is useful if you are using an application to sync users & groups with SCIM API.

-> This resource can be used with an account or workspace-level provider.

-> To assign an account level group to a workspace use [databricks_mws_permission_assignment](mws_permission_assignment.md).

-> Entitlements, like, `allow_cluster_create`, `allow_instance_pool_create`, `databricks_sql_access`, `workspace_access` applicable only for workspace-level groups.  Use [databricks_entitlements](entitlements.md) resource to assign entitlements inside a workspace to account-level groups.

To create account groups in the Databricks account, the provider must be configured accordingly. On AWS deployment with `host = "https://accounts.cloud.databricks.com"` and `account_id = "00000000-0000-0000-0000-000000000000"`. On Azure deployments `host = "https://accounts.azuredatabricks.net"`, `account_id = "00000000-0000-0000-0000-000000000000"` and using [AAD tokens](https://registry.terraform.io/providers/databricks/databricks/latest/docs#special-configurations-for-azure) as authentication.

Recommended to use along with Identity Provider SCIM provisioning to populate users into those groups:

* [Azure Active Directory](https://docs.microsoft.com/en-us/azure/databricks/administration-guide/users-groups/scim/aad)
* [Okta](https://docs.databricks.com/administration-guide/users-groups/scim/okta.html)
* [OneLogin](https://docs.databricks.com/administration-guide/users-groups/scim/onelogin.html)

## Example Usage

Creating some group

```hcl
resource "databricks_group" "this" {
  display_name               = "Some Group"
  allow_cluster_create       = true
  allow_instance_pool_create = true
}
```

Adding [databricks_user](user.md) as [databricks_group_member](group_member.md) of some group

```hcl
resource "databricks_group" "this" {
  display_name               = "Some Group"
  allow_cluster_create       = true
  allow_instance_pool_create = true
}

resource "databricks_user" "this" {
  user_name = "someone@example.com"
}

resource "databricks_group_member" "vip_member" {
  group_id  = databricks_group.this.id
  member_id = databricks_user.this.id
}
```

Creating group in AWS Databricks account:

```hcl
// initialize provider at account-level
provider "databricks" {
  alias         = "mws"
  host          = "https://accounts.cloud.databricks.com"
  account_id    = "00000000-0000-0000-0000-000000000000"
  client_id     = var.client_id
  client_secret = var.client_secret
}

resource "databricks_group" "this" {
  provider     = databricks.mws
  display_name = "Some Group"
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

resource "databricks_group" "this" {
  provider     = databricks.azure_account
  display_name = "Some Group"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` -  (Required) This is the display name for the given group.
* `external_id` - (Optional) ID of the group in an external identity provider.
* `allow_cluster_create` -  (Optional) This is a field to allow the group to have [cluster](cluster.md) create privileges. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and [cluster_id](permissions.md#cluster_id) argument. Everyone without `allow_cluster_create` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.
* `allow_instance_pool_create` -  (Optional) This is a field to allow the group to have [instance pool](instance_pool.md) create privileges. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.
* `databricks_sql_access` - (Optional) This is a field to allow the group to have access to [Databricks SQL](https://databricks.com/product/databricks-sql) feature in User Interface and through [databricks_sql_endpoint](sql_endpoint.md).
* `workspace_access` - (Optional) This is a field to allow the group to have access to Databricks Workspace.
* `force` - (Optional) Ignore `cannot create group: Group with name X already exists.` errors and implicitly import the specific group into Terraform state, enforcing entitlements defined in the instance of resource. _This functionality is experimental_ and is designed to simplify corner cases, like Azure Active Directory synchronisation.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  The id for the group object.
* `acl_principal_id` - identifier for use in [databricks_access_control_rule_set](access_control_rule_set.md), e.g. `groups/Some Group`.

## Import

You can import a `databricks_group` resource with the name `my_group` like the following:

```bash
terraform import databricks_group.my_group <group_id>
```
