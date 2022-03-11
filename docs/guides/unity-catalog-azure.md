---
page_title: "Unity Catalog set up on Azure"
---

# Deploying pre-requisite resources and enabling Unity Catalog (Azure Preview)

-> **Public Preview** This feature is in [Public Preview](https://docs.microsoft.com/en-us/azure/databricks/data-governance/unity-catalog). Contact your Databricks representative to request access. 

Databricks Unity Catalog brings fine-grained governance and security to Lakehouse data using a familiar, open interface. You can use Terraform to deploy the underlying cloud resources and Unity Catalog objects automatically, using a programmatic approach.

This guide uses the following variables in configurations:

- `databricks_workspace_url`: Value of `workspace_url` attribute from [databricks_mws_workspaces](../resources/mws_workspaces.md#attribute-reference) resource.
- `tenant_id`: The Azure AD tenant ID should be used for the [azuread](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs) provider.
- `subscription_id`: The Azure subscription ID should be used for the [azurerm](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs) provider

This guide is provided as-is and you can use this guide as the basis for your custom Terraform module.

To get started with Unity Catalog, this guide takes you throw the following high-level steps:
- [Deploying pre-requisite resources and enabling Unity Catalog (Azure Preview)](#deploying-pre-requisite-resources-and-enabling-unity-catalog-azure-preview)
  - [Provider initialization](#provider-initialization)
  - [Configure Azure objects](#configure-azure-objects)
  - [Create a Unity Catalog metastore and link it to workspaces](#create-a-unity-catalog-metastore-and-link-it-to-workspaces)
  - [Create Unity Catalog objects in the metastore](#create-unity-catalog-objects-in-the-metastore)
  - [Configure external tables and credentials](#configure-external-tables-and-credentials)
  - [Configure Unity Catalog clusters](#configure-unity-catalog-clusters)

## Provider initialization

Initialize the 3 providers to set up the required resources. See [Databricks provider authentication](../index.md#authenticating-with-hostname,-username,-and-password), [Azure AD provider authentication](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs#authenticating-to-azure-active-directory) and [Azure provider authentication](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azure) for more details.

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databrickslabs/databricks"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>2.92.0"
    }
    azuread = {
      source  = "hashicorp/azuread"
      version = "~>2.15.0"
    }
  }
}

provider "azuread" {
  tenant_id = var.tenant_id
}

provider "azurerm" {
  subscription_id = var.subscription_id
  features {}
}

// initialize provider at workspace level, to create UC resources
provider "databricks" {
  alias    = "workspace"
  host     = var.databricks_workspace_url
}
```

Define the required variables

```hcl
variable "databricks_workspace_url" {}

variable "tags" {
  default = {}
}

variable "reuse_rg" {
  description = "Whether to reuse resource group, do not create a new resource group (enter true/false)"
  type        = bool
}

variable "rg_name" {
  description = "Enter the resource group name where the Azure objects are deployed in"
  type        = string
}

variable "location" {
  description = "Enter your location, i.e. West US or East US"
  type        = string
}

variable "tenant_id" {
  description = "Enter your tenant id from Azure Portal"
  type        = string
}

variable "subscription_id" {
  description = "Enter your subscription id from Azure Portal"
  type        = string
}

variable "databricks_workspace_ids" {
  description = <<EOT
  List of Databricks workspace IDs to be enabled with Unity Catalog.
  Enter with square brackets and double quotes
  e.g. ["111111111", "222222222"]
  EOT
  type        = list(string)
}

//generate a random string as the prefix for AWS resources, to ensure uniqueness
resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

variable "prefix" {
  description = "Enter a prefix to prepend to any created resources"
  type        = string
}

locals {
  prefix = format("%s%s", var.prefix, random_string.naming.result)
}
```

## Configure Azure objects
The first step is to create the required Azure objects:
- An Azure storage account, which is the default storage location for managed tables in Unity Catalog. Please use a dedicated account for each metastore.
- An AAD service principal that provides Unity Catalog permissions to access and manage data in the bucket.

```hcl
resource "azuread_application" "unity_catalog" {
  display_name = local.prefix
}

resource "azuread_application_password" "unity_catalog" {
  application_object_id = azuread_application.unity_catalog.object_id
}

resource "azuread_service_principal" "unity_catalog" {
  application_id               = azuread_application.unity_catalog.application_id
  app_role_assignment_required = false
}


resource "azurerm_resource_group" "unity_catalog" {
  count    = var.reuse_rg ? 0 : 1
  name     = var.rg_name
  location = var.location
}

data "azurerm_resource_group" "unity_catalog" {
  count = var.reuse_rg ? 1 : 0
  name  = var.rg_name
}

resource "azurerm_storage_account" "unity_catalog" {
  name                     = local.prefix
  resource_group_name      = var.reuse_rg ? data.azurerm_resource_group.unity_catalog[0].name : azurerm_resource_group.unity_catalog[0].name
  location                 = var.reuse_rg ? data.azurerm_resource_group.unity_catalog[0].location : azurerm_resource_group.unity_catalog[0].location
  account_tier             = "Standard"
  account_replication_type = "GRS"
  is_hns_enabled           = true

  tags = var.tags
}

resource "azurerm_storage_container" "unity_catalog" {
  name                  = local.prefix
  storage_account_name  = azurerm_storage_account.unity_catalog.name
  container_access_type = "private"
}

resource "azurerm_role_assignment" "example" {
  scope                = azurerm_storage_account.unity_catalog.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = azuread_service_principal.unity_catalog.object_id
}
```

## Create a Unity Catalog metastore and link it to workspaces

A [databricks_metastore](../resources/metastore.md) is the top level container for data in Unity Catalog. A single metastore can be shared across Databricks workspaces, and each linked workspace has a consistent view of the data and a single set of access policies. Databricks recommends using a small number of metastores, except when organizations wish to have hard isolation boundaries between data. Data cannot be easily joined/queried across metastores.

```hcl
resource "databricks_metastore" "this" {
  name = var.metastore_name
  storage_root = format("abfss://%s@%s.dfs.core.windows.net/",
    azurerm_storage_account.unity_catalog.name,
  azurerm_storage_container.unity_catalog.name)
  owner = var.metastore_owner
  // forcefully remove that auto-created
  // catalog we have no access to
  force_destroy = true
}

resource "databricks_metastore_data_access" "first" {
  metastore_id = databricks_metastore.this.id
  name         = "the-keys"
  azure_service_principal {
    directory_id   = var.tenant_id
    application_id = azuread_application.unity_catalog.application_id
    client_secret  = azuread_application_password.unity_catalog.value
  }

  is_default = true
}

resource "databricks_metastore_assignment" "this" {
  for_each             = toset(var.workspace_ids)
  workspace_id         = each.key
  metastore_id         = databricks_metastore.this.id
  default_catalog_name = "hive_metastore"
}
```

## Create Unity Catalog objects in the metastore

Each metastore exposes a 3-level namespace (catalog-schema-table) by which data can be organized. 

```hcl
resource "databricks_catalog" "sandbox" {
  metastore_id = databricks_metastore.this.id
  name         = "sandbox"
  comment      = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
  depends_on = [databricks_metastore_assignment.default_metastore]
}

resource "databricks_grants" "sandbox" {
  catalog  = databricks_catalog.sandbox.name
  grant {
    principal  = "Data Scientists"
    privileges = ["USAGE", "CREATE"]
  }
  grant {
    principal  = "Data Engineers"
    privileges = ["USAGE"]
  }
}

resource "databricks_schema" "things" {
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this database is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_grants" "things" {
  schema   = databricks_schema.things.id
  grant {
    principal  = "Data Engineers"
    privileges = ["USAGE"]
  }
}
```

## Configure external tables and credentials

To work with external tables, Unity Catalog introduces two new objects to access and work with external cloud storage:
- [databricks_storage_credential](../resources/storage_credential.md) represent authentication methods to access cloud storage (e.g. an IAM role for Amazon S3 or a service principal for Azure Storage). Storage credentials are access-controlled to determine which users can use the credential.
- [databricks_external_location](../resources/external_location.md) are objects that combine a cloud storage path with a Storage Credential that can be used to access the location. 

First, create the required objects in Azure.

```hcl
resource "azuread_application" "ext_cred" {
  display_name = "${local.prefix}-cred"
}

resource "azuread_application_password" "ext_cred" {
  application_object_id = azuread_application.ext_cred.object_id
}

resource "azuread_service_principal" "ext_cred" {
  application_id               = azuread_application.ext_cred.application_id
  app_role_assignment_required = false
}

resource "azurerm_storage_account" "ext_storage" {
  name                     = "${local.prefix}ext"
  resource_group_name      = var.reuse_rg ? data.azurerm_resource_group.unity_catalog[0].name : azurerm_resource_group.unity_catalog[0].name
  location                 = var.reuse_rg ? data.azurerm_resource_group.unity_catalog[0].location : azurerm_resource_group.unity_catalog[0].location
  account_tier             = "Standard"
  account_replication_type = "GRS"
  is_hns_enabled           = true

  tags = var.tags
}

resource "azurerm_storage_container" "ext_storage" {
  name                  = "${local.prefix}-ext"
  storage_account_name  = azurerm_storage_account.ext_storage.name
  container_access_type = "private"
}

resource "azurerm_role_assignment" "ext_storage" {
  scope                = azurerm_storage_account.ext_storage.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = azuread_service_principal.ext_cred.object_id
}
```

Then create the [databricks_storage_credential](../resources/storage_credential.md) and [databricks_external_location](../resources/external_location.md) in Unity Catalog.

```hcl

resource "databricks_storage_credential" "external" {
  name = azuread_application.ext_cred.display_name
  azure_service_principal {
    directory_id   = var.tenant_id
    application_id = azuread_application.ext_cred.application_id
    client_secret  = azuread_application_password.ext_cred.value
  }
  comment = "Managed by TF"
  depends_on = [
    databricks_metastore_assignment.this
  ]  
}

resource "databricks_grants" "external_creds" {
  storage_credential = databricks_storage_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE TABLE"]
  }
}

resource "databricks_external_location" "some" {
  name = "external"
  url = format("abfss://%s@%s.dfs.core.windows.net/",
    azurerm_storage_account.ext_storage.name,
  azurerm_storage_container.ext_storage.name)
  credential_name = databricks_storage_credential.external.id
  comment         = "Managed by TF"
  depends_on = [
    databricks_metastore_assignment.this
  ]
}

resource "databricks_grants" "some" {
  external_location = databricks_external_location.some.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE TABLE", "READ FILES"]
  }
}
```

## Configure Unity Catalog clusters

To ensure the integrity of ACLs, Unity Catalog data can be accessed only through compute resources configured with strong isolation guarantees and other security features. A Unity Catalog [databricks_cluster](../resources/cluster.md) has a  ‘Security Mode’ set to either **User Isolation** or **Single User**.

- **User Isolation** clusters can be shared by multiple users, but only SQL language is allowed. Some advanced cluster features such as library installation, init scripts and the DBFS Fuse mount are also disabled in this mode to ensure security isolation among cluster users.

```hcl
data "databricks_spark_version" "latest" {
}
data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_cluster" "unity_sql" {
  cluster_name            = "Unity SQL"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 60
  enable_elastic_disk     = false
  num_workers             = 2
  azure_attributes {
    availability = "SPOT"
  }
  data_security_mode = "USER_ISOLATION"
}
```

- To use those advanced cluster features or languages like Python, Scala and R with Unity Catalog, one must choose **Single User** Mode when launching the cluster. The cluster can only be used exclusively by a single user (by default the owner of the cluster); other users are not allowed to attach to the cluster.
This means a group of users, which is managed as a group through SCIM provisioning, will be a collection of single-user [databricks_cluster](../resources/cluster.md), which they should be able to restart. Terraform's `for_each` meta-attribute helps to do this easily.

First we use [databricks_group](../data-sources/group.md) and [databricks_user](../data-sources/user.md) data resources to get the list of user names, that belong to a group.

```hcl
data "databricks_group" "dev" {
  display_name = "dev-clusters"
}

data "databricks_user" "dev" {
  for_each = data.databricks_group.dev.members
  user_id  = each.key
}
```

Once we have a specific list of user resources, we could proceed creating single-user clusters and provide permissions with `for_each = data.databricks_user.dev` to ensure it's done for each user:

```hcl
resource "databricks_cluster" "dev" {
  for_each                = data.databricks_user.dev
  cluster_name            = "${each.value.display_name} unity cluster"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 10
  enable_elastic_disk     = false
  num_workers             = 2
  azure_attributes {
    availability = "SPOT"
  }
  data_security_mode = "SINGLE_USER"
  single_user_name   = each.value.user_name
}

resource "databricks_permissions" "dev_restart" {
  for_each   = data.databricks_user.dev
  cluster_id = databricks_cluster.dev[each.key].cluster_id
  access_control {
    user_name        = each.value.user_name
    permission_level = "CAN_RESTART"
  }
}
```
