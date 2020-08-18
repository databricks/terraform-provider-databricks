+++
title = "azure_adls_gen1_mount"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_azure_adls_gen1_mount`

This resource given a cluster id will help you create, get and delete a azure data lake gen 1(ADLS gen 1) mount using a service 
principal/enterprise ad application which will provide you a client id and client secret to authenticate.

{{% notice warning %}}
It is important to understand that this will start up the cluster if the cluster is terminated.

The read and refresh terraform command will require a cluster and make take some time to validate mount.

If the cluster associated with the mount is deleted, then the mount will be re-created by terraform on next plan. It is important to note that provided the mount path and the storage account information remains the same, the mount will not actually get re-created inside the workspace.
{{% /notice %}}

{{% notice note %}}
You can locate the mount in dbfs at `dbfs:/mnt/<mount_name>`
{{% /notice %}}

## Example Usage

```hcl
resource "databricks_azure_adls_gen1_mount" "my_custom_mount3" {
  cluster_id = "####-######-pear###"
  storage_resource_name = "my_storage_resource_name"
  mount_name = "my_adls_gen1_mount"
  tenant_id = "????????-????-????-????-????????????"
  client_id = "????????-????-????-????-????????????"
  client_secret_scope = "my_adls_client_secret_scope"
  client_secret_key= "my_adls_client_secret_key"
}
```    

## Argument Reference

The following arguments are supported:

#### - `cluster_id`:
> **(Required)** This is the cluster id in which the mount will be initalized
from. If the cluster is in a terminated state it will be started.

#### - `storage_resource_name`:
> **(Required)** The name of the storage resource 
in which the data is for ADLS gen 1. This is what you are trying to mount.

#### - `spark_conf_prefix`:
> **(Required)** This is the spark configuration prefix
for adls gen 1 mount. The options are `fs.adl`, `dfs.adls`. Use `fs.adl` for runtime 6.0 and above for the clusters. 
Otherwise use `dfs.adls`. The default value is: `fs.adl`.

#### - `mount_directory`:
> **(Optional)** This is optional if you want to add an additional 
directory that you wish to mount. This must start with a "/"

#### - `mount_name`:
> **(Required)** The name of the folder that you want to mount to
in dbfs. You can access the data from `/mnt/<mount_name>` 

#### - `tenant_id`:
> **(Required)** This is your azure directory tenant id. This is 
required for creating the mount.

#### - `client_id`:
> **(Required)** This is the client_id for the enterprise application 
for the service principal. 

#### - `token_secret_scope`:
> **(Required)** This is the secret scope in which 
your service principal/enterprise app client secret will be stored.

#### - `token_secret_key`:
> **(Required)** This is the secret key in which 
your service principal/enterprise app client secret will be stored.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> Identifier for a adls gen 1 mount.


## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
