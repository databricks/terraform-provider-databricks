# databricks_azure_blob_mount Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

This resource will mount your Azure Blob Storage bucket on `dbfs:/mnt/yourname`. It is important to understand that this will start up the [cluster](cluster.md) if the cluster is terminated. The read and refresh terraform command will require a cluster and may take some time to validate the mount. If cluster_id is not specified, it will create the smallest possible cluster called `terraform-mount` for the shortest possible amount of time. This resource will help you create, get and delete an azure blob storage mount using SAS token or storage account access keys.


## Example Usage

```hcl
resource "azurerm_storage_account" "blobaccount" {
  name                     = "${var.prefix}blob"
  resource_group_name      = var.resource_group_name
  location                 = var.resource_group_location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  account_kind             = "StorageV2"
}

resource "azurerm_storage_container" "marketing" {
  name                  = "marketing"
  storage_account_name  = azurerm_storage_account.blobaccount.name
  container_access_type = "private"
}

resource "databricks_secret_scope" "terraform" {
    name                     = "application"
    initial_manage_principal = "users"
}

resource "databricks_secret" "storage_key" {
    key          = "blob_storage_key"
    string_value = azurerm_storage_account.blobaccount.primary_access_key
    scope        = databricks_secret_scope.terraform.name
}

resource "databricks_azure_blob_mount" "marketing" {
    container_name       = azurerm_storage_container.marketing.name
    storage_account_name = azurerm_storage_account.blobaccount.name
    mount_name           = "marketing"
    auth_type            = "ACCESS_KEY"
    token_secret_scope   = databricks_secret_scope.terraform.name
    token_secret_key     = databricks_secret.storage_key.key
}
```

## Argument Reference

The following arguments are required:

* `auth_type` - (Required) (String) This is the auth type for blob storage. This can either be SAS tokens or account access keys.
* `token_secret_scope` - (Required) (String) This is the secret scope in which your auth type token is stored.
* `token_secret_key` - (Required) (String) This is the secret key in which your auth type token is stored.
* `container_name` - (Required) (String) The container in which the data is. This is what you are trying to mount.
* `storage_account_name` - (Required) (String) The name of the storage resource in which the data is.
* `cluster_id` - (Optional) (String) Cluster to use for mounting. If no cluster is specified, a new cluster will be created and will mount the bucket for all of the clusters in this workspace. If the cluster is not running - it's going to be started, so be aware to set auto-termination rules on it.
* `mount_name` - (Required) (String) Name, under which mount will be accessible in `dbfs:/mnt/<MOUNT_NAME>`.
* `directory` - (Computed) (String) This is optional if you want to add an additional directory that you wish to mount. This must start with a "/".

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - mount name
* `source` - (String) HDFS-compatible url `wasbs://<adlsv2-account>` 


## Import

The resource can be imported using it's mount name

```bash
$ terraform import databricks_azure_blob_mount.this <mount_name>
```