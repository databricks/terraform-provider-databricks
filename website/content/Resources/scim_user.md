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

## Example Usage

```hcl
resource "databricks_scim_user" "my-user" {
  user_name = "testuser@databricks.com"
  display_name = "Test User"
  entitlements = [
    "allow-cluster-create",
  ]
}
```

## Argument Reference

The following arguments are supported:

#### - `user_name`:
> **(Required)** This is the username of the given user and will be their form of access 
and identity.

#### - `display_name`:
> **(Optional)** This is an alias for the username can be the full name of the user.

#### - `user_roles`:
> **(Optional)** This is a list of roles assigned to the user, specific to the AWS environment for 
user to assume roles on clusters.

#### - `entitlements`:
> **(Optional)** Entitlements for the user to be able to have the ability to create 
clusters and pools. Current options are: `"allow-cluster-create", "allow-instance-pool-create"`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the scim user object.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
