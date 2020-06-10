+++
title = "mws_credentials"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++


## Resource: `databricks_mws_credentials`

This resource to configure the credentials for the multiple workspaces api.

{{% notice warning %}}
It is important to understand that this will require you to configure your provider separately for the 
multiple workspaces resources
{{% /notice %}}

{{% notice note %}}
This will point to https://accounts.cloud.databricks.com for the HOST and it will use basic auth 
as that is the only authentication method available for multiple workspaces api.
{{% /notice %}}


## Example Usage

````hcl
provider "databricks" {
  host = "https://accounts.cloud.databricks.com"
  basic_auth {
    username = "username"
    password = "password"
  }
}
resource "databricks_mws_credentials" "my_mws_credentials" {
  account_id       = "my-mws-account-id"
  credentials_name = "my-cusotom-credentials"
  role_arn         = "arn:aws:iam::9999999999999:role/my-custom-cross-account-role"
}
````
## Argument Reference

The following arguments are supported:

#### - `account_id`:
> **(Required)** Databricks multi-workspace master account ID.

#### - `credentials_name`:
> **(Required)** The human-readable name of the credential configuration object.

#### - `role_arn`:
> **(Required)** This is the ARN of the role arn for the cross account role. 


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id of the resource which follows the format accountId/credentialsId.

#### - `creation_time`:
> Time in epoch milliseconds when the credential was created.

#### - `external_id`:
> The external ID that needs be trusted by the cross-account role. This is always the account_id, which is your Databricks multi-workspace master account ID.

#### - `credentials_id`:
> Databricks credential configuration ID.



## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
