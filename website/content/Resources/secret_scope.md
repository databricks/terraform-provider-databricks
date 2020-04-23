+++
title = "secret_scope"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_secret_scope`

This resource creates a Databricks-backed secret scope in which secrets are stored in Databricks-managed storage and 
encrypted with a cloud-based specific encryption key. 

The scope name:

* Must be unique within a workspace.
* Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.

## Example Usage

```hcl
resource "databricks_secret_scope" "my-scope" {
  name = "terraform-demo-scope"
  initial_manage_principal = "users"
}
```

## Argument Reference

The following arguments are supported:

#### - `name`:
> **(Required)** Scope name requested by the user. Scope names are unique. This field is required.

#### - `initial_manage_principal`:
> **(Optional)** The principal that is initially granted 
MANAGE permission to the created scope.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the secret scope object.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
