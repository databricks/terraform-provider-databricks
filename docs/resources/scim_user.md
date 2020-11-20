# databricks_scim_user Resource

!> **Warning** This resource is deprecated and going to be removed in 0.3. Please update all the code to [databricks_user](user.md) resource.

Directly creates a user within the databricks workspace. We're not recommending extensive use of this resource, because it's way more manageable to create few [databricks_group](group.md) instances with all related permissions to them and let Identity provider use SCIM provisioning to populate users into those groups:

* [Azure Active Directory](https://docs.microsoft.com/en-us/azure/databricks/administration-guide/users-groups/scim/aad)
* [Okta](https://docs.databricks.com/administration-guide/users-groups/scim/okta.html)
* [OneLogin](https://docs.databricks.com/administration-guide/users-groups/scim/onelogin.html)

This resource allows you to create users in Databricks and give them the proper level of access, as well as remove access for users (deprovision them) when they leave your organization or no longer need access to Databricks. This resource is heavily reliant on inherited group information and the default_roles object, to determine deltas. What this means is that, even if you change the roles field, if it is inherited it will ignore the change as it is inherited by the parent group. It will only detect delta when it is a net new role or a net new delete not covered by inherited roles or default roles.  


## Example Usage

```hcl
resource "databricks_scim_user" "admin" {
  user_name    = "me@example.com"
  display_name = "Example user"
  set_admin    = true
}
```

## Argument Reference

The following arguments are required:

* `user_name` - (Required) This is the username of the given user and will be their form of access and identity.
* `display_name` - (Optional) This is an alias for the username that can be the full name of the user.
* `roles` - (Optional) List of registered [instance profile](instance_profile.md) ARNs, that represent AWS IAM roles for the user to assume roles on AWS clusters. Has no effect on Azure.
* `entitlements` - (Optional) Entitlements for the user to be able to have the ability to create clusters and pools. Current options are: `"allow-cluster-create", "allow-instance-pool-create"`.
* `default_roles` - (Required) (Set) Set of roles that are assigned to the `all_users` group in Databricks. You can use the default_user_roles data source to fetch the values for this.
* `set_admin` - (Optional) (Bool) Setting this to true will patch this user to include the admin group id as a group item and if false, it will patch remove this user from the admin group.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the user.
* `inherited_roles` - (Set) The list of roles inherited by parent and all_users groups. This is used to determine when there are no changes.

## Import

The resource scim user can be imported using id:

```bash
$ terraform import databricks_scim_user.me <user-id>
```
