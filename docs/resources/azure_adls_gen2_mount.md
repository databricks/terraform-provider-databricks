---
subcategory: "Azure"
---
# databricks_azure_adls_gen2_mount Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

This resource will mount your ADLS v2 bucket on `dbfs:/mnt/yourname`. It is important to understand that this will start up the [cluster](cluster.md) if the cluster is terminated. The read and refresh terraform command will require a cluster and may take some time to validate the mount. If cluster_id is not specified, it will create the smallest possible cluster called `terraform-mount` for the shortest possible amount of time.

## Example Usage

```hcl
resource "databricks_secret_scope" "terraform" {
    name                     = "application"
    initial_manage_principal = "users"
}

resource "databricks_secret" "service_principal_key" {
    key          = "service_principal_key"
    string_value = "${var.ARM_CLIENT_SECRET}"
    scope        = databricks_secret_scope.terraform.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_storage_account" "this" {
  name                     = "${var.prefix}datalake"
  resource_group_name      = var.resource_group_name
  location                 = var.resource_group_location
  account_tier             = "Standard"
  account_replication_type = "GRS"
  account_kind             = "StorageV2"
  is_hns_enabled           = true
}

resource "azurerm_role_assignment" "this" {
  scope                = azurerm_storage_account.this.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = data.azurerm_client_config.current.object_id
}

resource "azurerm_storage_container" "this" {
  name                  = "marketing"
  storage_account_name  = azurerm_storage_account.this.name
  container_access_type = "private"
}

resource "databricks_azure_adls_gen2_mount" "marketing" {
    container_name         = azurerm_storage_container.this.name
    storage_account_name   = azurerm_storage_account.this.name
    mount_name             = "marketing"
    tenant_id              = data.azurerm_client_config.current.tenant_id
    client_id              = data.azurerm_client_config.current.client_id
    client_secret_scope    = databricks_secret_scope.terraform.name
    client_secret_key      = databricks_secret.service_principal_key.key
    initialize_file_system = true
}
```

## Argument Reference

The following arguments are required:

* `client_id` - (Required) (String) This is the client_id for the enterprise application for the service principal. 
* `tenant_id` - (Required) (String) This is your azure directory tenant id. This is required for creating the mount.
* `client_secret_key` - (Required) (String) This is the secret key in which your service principal/enterprise app client secret will be stored.
* `client_secret_scope` - (Required) (String) This is the secret scope in which your service principal/enterprise app client secret will be stored.

* `cluster_id` - (Optional) (String) Cluster to use for mounting. If no cluster is specified, a new cluster will be created and will mount the bucket for all of the clusters in this workspace. If the cluster is not running - it's going to be started, so be aware to set auto-termination rules on it.

* `container_name` - (Required) (String) ADLS gen2 container name
* `storage_account_name` - (Required) (String) The name of the storage resource in which the data is.
* `mount_name` - (Required) (String) Name, under which mount will be accessible in `dbfs:/mnt/<MOUNT_NAME>`.
* `directory` - (Computed) (String) This is optional if you want to add an additional directory that you wish to mount. This must start with a "/".
* `initialize_file_system` - (Required) (Bool) either or not initialize FS for the first use

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - mount name
* `source` - (String) HDFS-compatible url `abfss://<adlsv2-account>` 


## Import

The resource can be imported using it's mount name

```bash
$ terraform import databricks_azure_adls_gen2_mount.this <mount_name>
```