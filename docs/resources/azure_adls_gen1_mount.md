---
subcategory: "Azure"
---
# databricks_azure_adls_gen1_mount Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

This resource will mount your ADLS v1 bucket on `dbfs:/mnt/yourname`. It is important to understand that this will start up the [cluster](cluster.md) if the cluster is terminated. The read and refresh terraform command will require a cluster and may take some time to validate the mount. If cluster_id is not specified, it will create the smallest possible cluster called `terraform-mount` for the shortest possible amount of time.


## Example Usage

```hcl
resource "azurerm_data_lake_store" "gen1" {
  name                = "${local.prefix}gen1"
  resource_group_name = var.resource_group_name
  location            = var.resource_group_location
}

// azurerm_data_lake_store.gen1.name

resource "databricks_secret_scope" "this" {
    name                     = "application"
    initial_manage_principal = "users"
}

resource "databricks_secret" "service_principal_key" {
    key          = "service_principal_key"
    string_value = "{env.TEST_STORAGE_V2_ACCOUNT_KEY}"
    scope        = databricks_secret_scope.terraform.name
}

resource "databricks_azure_adls_gen1_mount" "mount" {
    container_name       = "dev"
    storage_account_name = "{env.TEST_STORAGE_ACCOUNT_NAME}"
    mount_name           = "{var.RANDOM}"
    auth_type            = "ACCESS_KEY"
    token_secret_scope   = databricks_secret_scope.terraform.name
    token_secret_key     = databricks_secret.service_principal_key.key
}

```

## Argument Reference

The following arguments are required:

* `client_id` - (Required) (String) This is the client_id for the enterprise application for the service principal. 
* `tenant_id` - (Required) (String) This is your azure directory tenant id. This is required for creating the mount.
* `client_secret_key` - (Required) (String) This is the secret key in which your service principal/enterprise app client secret will be stored.
* `client_secret_scope` - (Required) (String) This is the secret scope in which your service principal/enterprise app client secret will be stored.

* `cluster_id` - (Optional) (String) Cluster to use for mounting. If no cluster is specified, a new cluster will be created and will mount the bucket for all of the clusters in this workspace. If the cluster is not running - it's going to be started, so be aware to set auto-termination rules on it.
* `mount_name` - (Required) (String) Name, under which mount will be accessible in `dbfs:/mnt/<MOUNT_NAME>`.
* `storage_resource_name` - (Required) (String) The name of the storage resource in which the data is for ADLS gen 1. This is what you are trying to mount.
* `spark_conf_prefix` - (Optional) (String) This is the spark configuration prefix for adls gen 1 mount. The options are `fs.adl`, `dfs.adls`. Use `fs.adl` for runtime 6.0 and above for the clusters. Otherwise use `dfs.adls`. The default value is: `fs.adl`.
* `directory` - (Computed) (String) This is optional if you want to add an additional directory that you wish to mount. This must start with a "/".



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - mount name
* `source` - (String) HDFS-compatible url `adl://<adlsv1-account>` 


## Import

The resource can be imported using it's mount name

```bash
$ terraform import databricks_azure_adls_gen1_mount.this <mount_name>
```