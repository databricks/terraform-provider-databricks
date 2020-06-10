+++
title = "mws_storage_configurations"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++


## Resource: `databricks_mws_storage_configurations`

This resource to configure the root bucket for the multiple workspaces api.

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
resource "databricks_mws_storage_configurations" "my_mws_storage_configurations" {
  account_id = "my-mws-acct-id"
  storage_configuration_name = "storage-configuration-name"
  bucket_name         = "my-root-s3-bucket"
}
````
## Argument Reference

The following arguments are supported:

#### - `account_id`:
> **(Required)** Databricks multi-workspace master account ID.

#### - `storage_configuration_name`:
> **(Required)** The human-readable name of the storage configuration.
                 
#### - `bucket_name`:
> **(Required)** Root S3 bucket information. 


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id of the resource which follows the format accountId/storageConfigurationId.

#### - `creation_time`:
> Time in epoch milliseconds when the storage configuration was created

#### - `storage_configuration_id`:
> Databricks storage configuration ID.



## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
