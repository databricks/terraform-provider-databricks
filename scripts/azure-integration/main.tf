provider "azurerm" {
  version = "~> 2.14"
  features {}
}

provider "random" {
  version = "~> 2.2"
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

data "azurerm_client_config" "current" {
}

data "external" "me" {
  program = ["az", "account", "show", "--query", "user"]
}

locals {
  // dltp - databricks labs terraform provider
  prefix = "dltp${random_string.naming.result}"
  tags = {
    Environment = "Testing"
    Owner       = lookup(data.external.me.result, "name")
    Epoch       = random_string.naming.result
  }
}

resource "azurerm_resource_group" "example" {
  name     = "${local.prefix}-rg"
  location = "eastus"
  tags     = local.tags
}

resource "azurerm_databricks_workspace" "example" {
  name                        = "${local.prefix}-workspace"
  resource_group_name         = azurerm_resource_group.example.name
  location                    = azurerm_resource_group.example.location
  sku                         = "premium"
  managed_resource_group_name = "${local.prefix}-workspace-rg"
  tags                        = local.tags
}

# Create container in storage acc and container for use by adls gen2 mount tests
resource "azurerm_storage_account" "adlsaccount" {
  name                     = "${local.prefix}datalake"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
  account_kind             = "StorageV2"
  is_hns_enabled           = true
  tags                     = local.tags
}

# Create container in storage acc and container for use by blob mount tests
resource "azurerm_storage_account" "blobaccount" {
  name                     = "${local.prefix}blob"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  account_kind             = "StorageV2"
  tags                     = local.tags
}

//resource "azurerm_role_assignment" "datalake" {
//  scope = azurerm_storage_account.adlsaccount.id
//  #https://docs.microsoft.com/en-us/azure/role-based-access-control/built-in-roles#storage-blob-data-contributor
//  role_definition_name = "Storage Blob Data Contributor"
//  principal_id         = data.azurerm_client_config.current.object_id
//}

resource "azurerm_storage_container" "adlsexample" {
  name                  = "dev"
  storage_account_name  = azurerm_storage_account.adlsaccount.name
  container_access_type = "private"
}

resource "azurerm_storage_container" "blobexample" {
  name                  = "dev"
  storage_account_name  = azurerm_storage_account.blobaccount.name
  container_access_type = "private"
}

output "cloud_env" {
  value = "azure"
}

output "test_gen2_adal_name" {
  value = azurerm_storage_account.adlsaccount.name
}

output "test_storage_account_key" {
  value     = azurerm_storage_account.blobaccount.primary_access_key
  sensitive = true
}

output "test_storage_account_name" {
  value = azurerm_storage_account.blobaccount.name
}

output "azure_region" {
  value = azurerm_storage_account.adlsaccount.location
}

// todo: agree with Sri on the name
output "azure_databricks_workspace_resource_id" {
  // The ID of the Databricks Workspace in the Azure management plane.
  value = azurerm_databricks_workspace.example.id
}

output "databricks_host" {
  // The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
  value = "https://${azurerm_databricks_workspace.example.workspace_url}/"
}