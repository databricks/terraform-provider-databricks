+++
title = "secret_acl"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_secret_acl`

Create or overwrite the ACL associated with the given principal (user or group) on the specified scope point. 
In general, a user or group will use the most powerful permission available to them, and 
permissions are ordered as follows:

* **MANAGE** - Allowed to change ACLs, and read and write to this secret scope.
* **WRITE** - Allowed to read and write to this secret scope.
* **READ** - Allowed to read this secret scope and list what secrets are available.

## Example Usage

```hcl
resource "databricks_secret_scope" "my-scope" {
  name = "terraform-demo-scope"
}

resource "databricks_secret_acl" "my-acl" {
  principal = "USERS"
  permission = "READ"
  scope = "${databricks_secret_scope.my-scope.name}"
}
```
## Argument Reference

The following arguments are supported:

#### - `scope`:
> **(Required)** The name of the scope to remove permissions from. This field is required. 
(MB4) form.

#### - `principal`:
> **(Required)** The principal to remove an existing ACL from. The principal is a user 
or group name corresponding to an existing Databricks principal to be granted or revoked access. This field is required. 

#### - `permission`:
> **(Required)** The permission level applied to the principal. 
Options are: `"READ", "WRITE", "MANAGE"`. This field is required.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the secret scope acl object.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
