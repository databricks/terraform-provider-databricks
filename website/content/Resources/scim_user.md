+++
title = "scim_user"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_scim_user`

This resource allows you to create users in Databricks and give them the proper level of access, as well as 
remove access for users (deprovision them) when they leave your organization or no longer need access to Databricks.

{{% notice note %}} 
You must be a Databricks administrator API token to use SCIM resources.
{{% /notice %}} 

{{% notice info %}} 
This resource is heavily reliant on inherited group information and the default_roles object, to determine deltas.
What this means is that, even if you change the roles field, if it is inherited it will ignore the change as it is 
inherited by parent group. It will only detect delta when it is a net new role or a net new delete not covered by 
inherited roles or default roles.  
{{% /notice %}} 

## Example Usage

```hcl
resource "databricks_scim_user" "my-user" {
  user_name = "testuser@databricks.com"
  display_name = "Test User"
  default_roles = []
  entitlements = [
    "allow-cluster-create",
  ]
}
```

## Argument Reference

The following arguments are supported:

#### - `user_name`:
> **(Required)** 

#### - `display_name`:
> **(Optional)** 

#### - `roles`:
> **(Optional)** 

#### - `entitlements`:
> **(Optional)** 

#### - `default_roles`:
> **(Required)** Set of roles that are assigned to the `all_users` group in Databricks. You can use the default_user_roles data source to fetch the values for this.

#### - `set_admin`:
> **(Required)** 


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the scim user object.

#### - `inherited_roles`:
> The list of roles inherited by parent and all_users groups. This is used to determine when there are no changes.


## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
