# Resource: db_azure_blob_mount

This resource given a cluster id will help you create, get and delete a azure blob storage mount using SAS token or
storage account access keys.

.. Warning:: It is important to understand that this will start up the cluster if the cluster is terminated.
    The read and refresh terraform command will require a cluster and make take some time to validate mount.

.. Important:: You can locate the mount in dbfs at `dbfs:/mnt/<mount_name>`


## Example Usage

With the below resource the data gets mounted to `/mnt/my_cool_dbfs_mount`

.. code-block:: tf

    resource "db_azure_blob_mount" "my_custom_mount" {
      cluster_id = "####-######-pear###"
      container_name = "my_storage_container"
      directory = "/my_custom_folder"
      storage_account_name  = "mystorageaccountname"
      mount_name = "my_cool_blob_storage_mount"
      auth_type = "SAS"
      token_secret_scope = "my_secret_scope"
      token_secret_key = "my_secret_key"
    }

    
## Argument Reference

The following arguments are supported:

.. _r_azure_blob_mount_cluster_id:
* :ref:`cluster_id <r_azure_blob_mount_cluster_id>` - **(Required)** This is the cluster id in which the mount will be initalized
from. If the cluster is in a terminated state it will be started.

.. _r_azure_blob_mount_container_name:
* :ref:`container_name <r_azure_blob_mount_container_name>` - **(Required)** The container in which the data is. This 
is what you are trying to mount.

.. _r_azure_blob_mount_storage_account_name:
* :ref:`storage_account_name <r_azure_blob_mount_storage_account_name>` - **(Required)** The name of the storage account 
in which the data is. This is what you are trying to mount.

.. _r_azure_blob_mount_directory:
* :ref:`directory <r_azure_blob_mount_directory>` - **(Optional)** This is optional if you want to add an additional 
directory that you wish to mount. This must start with a "/"

.. _r_azure_blob_mount_mount_name:
* :ref:`mount_name <r_azure_blob_mount_mount_name>` - **(Required)** The name of the folder that you want to mount to
in dbfs. You can access the data from `/mnt/<mount_name>` 

.. _r_azure_blob_mount_auth_type:
* :ref:`auth_type <r_azure_blob_mount_auth_type>` - **(Required)** This is the auth type for blob storage. This can either 
be SAS tokens or account access keys.

.. _r_azure_blob_mount_token_secret_scope:
* :ref:`token_secret_scope <r_azure_blob_mount_token_secret_scope>` - **(Required)** This is the secret scope in which 
your auth type token exists in.

.. _r_azure_blob_mount_token_secret_key:
* :ref:`token_secret_key <r_azure_blob_mount_token_secret_key>` - **(Required)** This is the secret key in which 
your auth type token exists in.
 

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_azure_blob_mount_id:
* :ref:`id <r_azure_blob_mount_id>` - The id of the azure blob storage mount.


## Import

.. Note:: Importing this resource is not currently supported.