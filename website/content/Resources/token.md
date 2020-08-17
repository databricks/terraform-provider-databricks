+++
title = "token"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_token`

This resource creates an api token that can be used to create Databricks resources. 

{{% notice warning %}}
This will create an API token for the user that has authenticated on the provider. So if you have used an 
admin user to setup the provider then you will be making API tokens for that admin user. 
{{% /notice %}}

## Example Usage

```hcl
resource "databricks_token" "my-token" {
  lifetime_seconds = 6000
  comment = "Testing terraform v2"
}
```
## Argument Reference

The following arguments are supported:

#### - `lifetime_seconds`:
> **(Optional) (Numeric)** The lifetime of the token, in seconds. If no lifetime is specified, the token remains valid indefinitely.

#### - `comment`:
> **(Optional)** Optional description to attach to the token.


## Attribute Reference
In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the token object.

#### - `creation_time`:
> Server time (in epoch milliseconds) when the token was created.

#### - `token_value`:
> **Sensitive** The value of the newly-created token. 

#### - `expiry_time`:
> Server time (in epoch milliseconds) when the token will expire, or -1 if not applicable.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
