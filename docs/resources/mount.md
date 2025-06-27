---
subcategory: "Storage"
---
# databricks_mount Resource

This resource will [mount your cloud storage](https://docs.databricks.com/data/databricks-file-system.html#mount-object-storage-to-dbfs) on `dbfs:/mnt/name`. Right now it supports mounting AWS S3, Azure (Blob Storage, ADLS Gen1 & Gen2), Google Cloud Storage.  It is important to understand that this will start up the [cluster](cluster.md) if the cluster is terminated. The read and refresh terraform command will require a cluster and may take some time to validate the mount.

-> This resource can only be used with a workspace-level provider!

-> When `cluster_id` is not specified, it will create the smallest possible cluster in the default availability zone with name equal to or starting with `terraform-mount` for the shortest possible amount of time. To avoid mount failure due to potentially quota or capacity issues with the default cluster, we recommend specifying a cluster to use for mounting.

-> CRUD operations on a databricks mount require a running cluster. Due to limitations of terraform and the databricks mounts APIs, if the cluster the mount was most recently created / updated using no longer exists AND the mount is destroyed as a part of a terraform apply, we mark it as deleted without cleaning it up from the workspace.

This resource provides two ways of mounting a storage account:

1. Use a storage-specific configuration block - this could be used for the most cases, as it will fill most of the necessary details. Currently we support following configuration blocks:

* `s3` - to [mount AWS S3](https://docs.databricks.com/data/data-sources/aws/amazon-s3.html)
* `gs` - to [mount Google Cloud Storage](https://docs.gcp.databricks.com/data/data-sources/google/gcs.html)
* `abfs` - to [mount ADLS Gen2](https://docs.microsoft.com/en-us/azure/databricks/data/data-sources/azure/adls-gen2/) using Azure Blob Filesystem (ABFS) driver
* `adl` - to [mount ADLS Gen1](https://docs.microsoft.com/en-us/azure/databricks/data/data-sources/azure/azure-datalake) using Azure Data Lake (ADL) driver
* `wasb`  - to [mount Azure Blob Storage](https://docs.microsoft.com/en-us/azure/databricks/data/data-sources/azure/azure-storage) using Windows Azure Storage Blob (WASB) driver

1. Use generic arguments - you have a responsibility for providing all necessary parameters that are required to mount specific storage. This is most flexible option

## Common arguments

* `cluster_id` - (Optional, String) Cluster to use for mounting. If no cluster is specified, a new cluster will be created and will mount the bucket for all of the clusters in this workspace. If the cluster is not running - it's going to be started, so be aware to set auto-termination rules on it.
* `name` - (Optional, String) Name, under which mount will be accessible in `dbfs:/mnt/<MOUNT_NAME>`. If not specified, provider will try to infer it from depending on the resource type:
  * `bucket_name` for AWS S3 and Google Cloud Storage
  * `container_name` for ADLS Gen2 and Azure Blob Storage
  * `storage_resource_name` for ADLS Gen1
* `uri` - (Optional, String) the URI for accessing specific storage (`s3a://....`, `abfss://....`, `gs://....`, etc.)
* `extra_configs` - (Optional, String map) configuration parameters that are necessary for mounting of specific storage
* `resource_id` - (Optional, String) resource ID for a given storage account. Could be used to fill defaults, such as storage account & container names on Azure.
* `encryption_type` - (Optional, String) encryption type. Currently used only for [AWS S3 mounts](https://docs.databricks.com/data/data-sources/aws/amazon-s3.html#encrypt-data-in-s3-buckets)

### Example mounting ADLS Gen2 using uri and extra_configs

```hcl
locals {
  tenant_id    = "00000000-1111-2222-3333-444444444444"
  client_id    = "55555555-6666-7777-8888-999999999999"
  secret_scope = "some-kv"
  secret_key   = "some-sp-secret"
  container    = "test"
  storage_acc  = "lrs"
}

resource "databricks_mount" "this" {
  name = "tf-abfss"

  uri = "abfss://${local.container}@${local.storage_acc}.dfs.core.windows.net"
  extra_configs = {
    "fs.azure.account.auth.type" : "OAuth",
    "fs.azure.account.oauth.provider.type" : "org.apache.hadoop.fs.azurebfs.oauth2.ClientCredsTokenProvider",
    "fs.azure.account.oauth2.client.id" : local.client_id,
    "fs.azure.account.oauth2.client.secret" : "{{secrets/${local.secret_scope}/${local.secret_key}}}",
    "fs.azure.account.oauth2.client.endpoint" : "https://login.microsoftonline.com/${local.tenant_id}/oauth2/token",
    "fs.azure.createRemoteFileSystemDuringInitialization" : "false",
  }
}
```

### Example mounting ADLS Gen2 with AAD passthrough

-> AAD passthrough is considered a legacy data access pattern. Use Unity Catalog for fine-grained data access control.

-> Mounts using AAD passthrough cannot be created using a service principal.

To mount ALDS Gen2 with Azure Active Directory Credentials passthrough we need to execute the mount commands using the cluster configured with AAD Credentials passthrough & provide necessary configuration parameters (see [documentation](https://docs.microsoft.com/en-us/azure/databricks/security/credential-passthrough/adls-passthrough#--mount-azure-data-lake-storage-to-dbfs-using-credential-passthrough) for more details).

```hcl
provider "azurerm" {
  features {}
}

variable "resource_group" {
  type        = string
  description = "Resource group for Databricks Workspace"
}

variable "workspace_name" {
  type        = string
  description = "Name of the Databricks Workspace"
}

data "azurerm_databricks_workspace" "this" {
  name                = var.workspace_name
  resource_group_name = var.resource_group
}

# it works only with AAD token!
provider "databricks" {
  host = data.azurerm_databricks_workspace.this.workspace_url
}

data "databricks_node_type" "smallest" {
  local_disk = true
}

data "databricks_spark_version" "latest" {
}

resource "databricks_cluster" "shared_passthrough" {
  cluster_name            = "Shared Passthrough for mount"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 10
  num_workers             = 1

  spark_conf = {
    "spark.databricks.cluster.profile" : "serverless",
    "spark.databricks.repl.allowedLanguages" : "python,sql",
    "spark.databricks.passthrough.enabled" : "true",
    "spark.databricks.pyspark.enableProcessIsolation" : "true"
  }

  custom_tags = {
    "ResourceClass" : "Serverless"
  }
}

variable "storage_acc" {
  type        = string
  description = "Name of the ADLS Gen2 storage container"
}

variable "container" {
  type        = string
  description = "Name of container inside storage account"
}

resource "databricks_mount" "passthrough" {
  name       = "passthrough-test"
  cluster_id = databricks_cluster.shared_passthrough.id

  uri = "abfss://${var.container}@${var.storage_acc}.dfs.core.windows.net"
  extra_configs = {
    "fs.azure.account.auth.type" : "CustomAccessToken",
    "fs.azure.account.custom.token.provider.class" : "{{sparkconf/spark.databricks.passthrough.adls.gen2.tokenProviderClassName}}",
  }
}
```

## s3 block

This block allows specifying parameters for mounting of the ADLS Gen2. The following arguments are required inside the `s3` block:

* `instance_profile` - (Optional) (String) ARN of registered [instance profile](instance_profile.md) for data access.  If it's not specified, then the `cluster_id` should be provided, and the cluster should have an instance profile attached to it. If both `cluster_id` & `instance_profile` are specified, then `cluster_id` takes precedence.
* `bucket_name` - (Required) (String) S3 bucket name to be mounted.

### Example of mounting S3

```hcl
// now you can do `%fs ls /mnt/experiments` in notebooks
resource "databricks_mount" "this" {
  name = "experiments"
  s3 {
    instance_profile = databricks_instance_profile.ds.id
    bucket_name      = aws_s3_bucket.this.bucket
  }
}
```

## abfs block

This block allows specifying parameters for mounting of the ADLS Gen2. The following arguments are required inside the `abfs` block:

* `client_id` - (Required) (String) This is the client_id (Application Object ID) for the enterprise application for the service principal.
* `tenant_id` - (Optional) (String) This is your azure directory tenant id. It is required for creating the mount. (Could be omitted if Azure authentication is used, and we can extract `tenant_id` from it).
* `client_secret_key` - (Required) (String) This is the secret key in which your service principal/enterprise app client secret will be stored.
* `client_secret_scope` - (Required) (String) This is the secret scope in which your service principal/enterprise app client secret will be stored.
* `container_name` - (Required) (String) ADLS gen2 container name. (Could be omitted if `resource_id` is provided)
* `storage_account_name` - (Required) (String) The name of the storage resource in which the data is. (Could be omitted if `resource_id` is provided)
* `directory` - (Computed) (String) This is optional if you don't want to add an additional directory that you wish to mount. This must start with a "/".
* `initialize_file_system` - (Required) (Bool) either or not initialize FS for the first use

### Creating mount for ADLS Gen2 using abfs block

In this example, we're using Azure authentication, so we can omit some parameters (`tenant_id`, `storage_account_name`, and `container_name`) that will be detected automatically.

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

resource "databricks_mount" "marketing" {
  name        = "marketing"
  resource_id = azurerm_storage_container.this.resource_manager_id
  abfs {
    client_id              = data.azurerm_client_config.current.client_id
    client_secret_scope    = databricks_secret_scope.terraform.name
    client_secret_key      = databricks_secret.service_principal_key.key
    initialize_file_system = true
  }
}
```

## gs block

This block allows specifying parameters for mounting of the Google Cloud Storage. The following arguments are required inside the `gs` block:

* `service_account` - (Optional) (String) email of registered [Google Service Account](https://docs.gcp.databricks.com/data/data-sources/google/gcs.html#step-1-set-up-google-cloud-service-account-using-google-cloud-console) for data access.  If it's not specified, then the `cluster_id` should be provided, and the cluster should have a Google service account attached to it.
* `bucket_name` - (Required) (String) GCS bucket name to be mounted.

### Example mounting Google Cloud Storage

```hcl
resource "databricks_mount" "this_gs" {
  name = "gs-mount"
  gs {
    service_account = "acc@company.iam.gserviceaccount.com"
    bucket_name     = "mybucket"
  }
}
```

## adl block

This block allows specifying parameters for mounting of the ADLS Gen1. The following arguments are required inside the `adl` block:

* `client_id` - (Required) (String) This is the client_id for the enterprise application for the service principal.
* `tenant_id` - (Optional) (String) This is your azure directory tenant id. It is required for creating the mount. (Could be omitted if Azure authentication is used, and we can extract `tenant_id` from it)
* `client_secret_key` - (Required) (String) This is the secret key in which your service principal/enterprise app client secret will be stored.
* `client_secret_scope` - (Required) (String) This is the secret scope in which your service principal/enterprise app client secret will be stored.

* `storage_resource_name` - (Required) (String) The name of the storage resource in which the data is for ADLS gen 1. This is what you are trying to mount. (Could be omitted if `resource_id` is provided)
* `spark_conf_prefix` - (Optional) (String) This is the spark configuration prefix for adls gen 1 mount. The options are `fs.adl`, `dfs.adls`. Use `fs.adl` for runtime 6.0 and above for the clusters. Otherwise use `dfs.adls`. The default value is: `fs.adl`.
* `directory` - (Computed) (String) This is optional if you don't want to add an additional directory that you wish to mount. This must start with a "/".

### Example mounting ADLS Gen1

```hcl
resource "databricks_mount" "mount" {
  name = "{var.RANDOM}"
  adl {
    storage_resource_name = "{env.TEST_STORAGE_ACCOUNT_NAME}"
    tenant_id             = data.azurerm_client_config.current.tenant_id
    client_id             = data.azurerm_client_config.current.client_id
    client_secret_scope   = databricks_secret_scope.terraform.name
    client_secret_key     = databricks_secret.service_principal_key.key
    spark_conf_prefix     = "fs.adl"
  }
}
```

## wasb block

This block allows specifying parameters for mounting of the Azure Blob Storage. The following arguments are required inside the `wasb` block:

* `auth_type` - (Required) (String) This is the auth type for blob storage. This can either be SAS tokens (`SAS`) or account access keys (`ACCESS_KEY`).
* `token_secret_scope` - (Required) (String) This is the secret scope in which your auth type token is stored.
* `token_secret_key` - (Required) (String) This is the secret key in which your auth type token is stored.
* `container_name` - (Required) (String) The container in which the data is. This is what you are trying to mount. (Could be omitted if `resource_id` is provided)
* `storage_account_name` - (Required) (String) The name of the storage resource in which the data is. (Could be omitted if `resource_id` is provided)
* `directory` - (Computed) (String) This is optional if you don't want to add an additional directory that you wish to mount. This must start with a "/".

### Example mounting Azure Blob Storage

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

resource "databricks_mount" "marketing" {
  name = "marketing"
  wasb {
    container_name       = azurerm_storage_container.marketing.name
    storage_account_name = azurerm_storage_account.blobaccount.name
    auth_type            = "ACCESS_KEY"
    token_secret_scope   = databricks_secret_scope.terraform.name
    token_secret_key     = databricks_secret.storage_key.key
  }
}
```

## Migration from other mount resources

Migration from the specific mount resource is straightforward:

* rename `mount_name` to `name`
* wrap storage-specific settings (`container_name`, ...) into corresponding block (`adl`, `abfs`, `s3`, `wasbs`)
* for S3 mounts, rename `s3_bucket_name` to `bucket_name`

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - mount name
* `source` - (String) HDFS-compatible url

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_aws_bucket_policy](../data-sources/aws_bucket_policy.md) data to configure a simple access policy for AWS S3 buckets, so that Databricks can access data in it.
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_dbfs_file](../data-sources/dbfs_file.md) data to get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file_paths](../data-sources/dbfs_file_paths.md) data to get list of file names from get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file](dbfs_file.md) to manage relatively small files on [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_library](library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md).
