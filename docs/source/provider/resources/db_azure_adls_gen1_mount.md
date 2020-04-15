# Resource: db_azure_adls_gen1_mount

This resource given a cluster id will help you create, get and delete a azure data lake gen 1(ADLS gen 1) mount using a service 
principal/enterprise ad application which will provide you a client id and client secret to authenticate.

.. Warning:: It is important to understand that this will start up the cluster if the cluster is terminated.
    The read and refresh terraform command will require a cluster and make take some time to validate mount.

.. Important:: You can locate the mount in dbfs at `dbfs:/mnt/<mount_name>`

## Example Usage

.. code-block:: tf

    resource "db_azure_adls_gen1_mount" "my_custom_mount3" {
      cluster_id = "####-######-pear###"
      storage_resource_name = "my_storage_resource_name"
      mount_name = "my_adls_gen1_mount"
      tenant_id = "????????-????-????-????-????????????"
      client_id = "????????-????-????-????-????????????"
      client_secret_scope = "my_adls_client_secret_scope"
      client_secret_key= "my_adls_client_secret_key"
    }
    
## Argument Reference

The following arguments are supported:

.. _r_azure_adls_gen1_mount_cluster_id:
* :ref:`cluster_id <r_azure_adls_gen1_mount_cluster_id>` - **(Required)** This is the cluster id in which the mount will be initalized
from. If the cluster is in a terminated state it will be started.

.. _r_azure_adls_gen1_mount_storage_resource_name:
* :ref:`storage_resource_name <r_azure_adls_gen1_mount_storage_resource_name>` - **(Required)** The name of the storage resource 
in which the data is for ADLS gen 1. This is what you are trying to mount.

.. _r_azure_adls_gen1_mount_spark_conf_prefix:
* :ref:`spark_conf_prefix <r_azure_adls_gen1_mount_spark_conf_prefix>` - **(Required)** This is the spark configuration prefix
for adls gen 1 mount. The options are `"fs.adl", "dfs.adls"`. Use `"fs.adl"` for runtime 6.0 and above for the clusters. 
Otherwise use `"dfs.adls"`. The default value is: `"fs.adl"`.

.. _r_azure_adls_gen1_mount_directory:
* :ref:`directory <r_azure_adls_gen1_mount_directory>` - **(Optional)** This is optional if you want to add an additional 
directory that you wish to mount. This must start with a "/"

.. _r_azure_adls_gen1_mount_mount_name:
* :ref:`mount_name <r_azure_adls_gen1_mount_mount_name>` - **(Required)** The name of the folder that you want to mount to
in dbfs. You can access the data from `/mnt/<mount_name>` 

.. _r_azure_adls_gen1_mount_tenant_id:
* :ref:`tenant_id <r_azure_adls_gen1_mount_tenant_id>` - **(Required)** This is your azure directory tenant id. This is 
required for creating the mount.

.. _r_azure_adls_gen1_mount_client_id:
* :ref:`client_id <r_azure_adls_gen1_mount_client_id>` - **(Required)** This is the client_id for the enterprise application 
for the service principal. 

.. _r_azure_adls_gen1_mount_token_secret_scope:
* :ref:`token_secret_scope <r_azure_adls_gen1_mount_token_secret_scope>` - **(Required)** This is the secret scope in which 
your service principal/enterprise app client secret will be stored.

.. _r_azure_adls_gen1_mount_token_secret_key:
* :ref:`token_secret_key <r_azure_adls_gen1_mount_token_secret_key>` - **(Required)** This is the secret key in which 
your service principal/enterprise app client secret will be stored.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_azure_adls_gen1_mount_id:
* :ref:`id <r_azure_adls_gen1_mount_id>` - Identifier for a adls gen 1 mount.


## Import

.. Note:: Importing this resource is not currently supported.