+++
title = "azure_blob_mount"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_azure_blob_mount`

This resource given a cluster id will help you create, get and delete a azure blob storage mount using SAS token or
storage account access keys.

{{% notice warning %}}
It is important to understand that this will start up the cluster if the cluster is terminated.

The read and refresh terraform command will require a cluster and make take some time to validate mount.

If the cluster associated with the mount is deleted, then the mount will be re-created by terraform on next plan. It is important to note that provided the mount path and the storage account information remains the same, the mount will not actually get re-created inside the workspace.
{{% /notice %}}

{{% notice note %}}
You can locate the mount in dbfs at `dbfs:/mnt/<mount_name>`
{{% /notice %}}


## Example Usage

With the below resource the data gets mounted to `/mnt/my_cool_dbfs_mount`

```hcl
resource "databricks_azure_blob_mount" "my_custom_mount" {
  cluster_id = "####-######-pear###"
  container_name = "my_storage_container"
  directory = "/my_custom_folder"
  storage_account_name  = "mystorageaccountname"
  mount_name = "my_cool_blob_storage_mount"
  auth_type = "SAS"
  token_secret_scope = "my_secret_scope"
  token_secret_key = "my_secret_key"
}
```
    
## Argument Reference

The following arguments are supported:

#### - `cluster_id`:
> **(Required)** This is the cluster id in which the mount will be initalized
from. If the cluster is in a terminated state it will be started.

#### - `container_name`:
> **(Required)** The container in which the data is. This 
is what you are trying to mount.

#### - `storage_account_name`:
> **(Required)** The name of the storage account 
in which the data is. This is what you are trying to mount.

#### - `directory`:
> **(Optional)** This is optional if you want to add an additional 
directory that you wish to mount. This must start with a "/"

#### - `mount_name`:
> **(Required)** The name of the folder that you want to mount to
in dbfs. You can access the data from `/mnt/<mount_name>` 

#### - `auth_type`:
> **(Required)** This is the auth type for blob storage. This can either 
be SAS tokens or account access keys.

#### - `token_secret_scope`:
> **(Required)** This is the secret scope in which 
your auth type token exists in.

#### - `token_secret_key`:
> **(Required)** This is the secret key in which 
your auth type token exists in.
 

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id of the azure blob storage mount.


## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
