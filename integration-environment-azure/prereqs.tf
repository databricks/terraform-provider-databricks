provider "azurerm" {
  version = "~> 2.3"
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

resource "azurerm_resource_group" "example" {
  name     = "inttest${random_string.naming.result}"
  location = "eastus"
}

resource "azurerm_databricks_workspace" "example" {
  name                        = "workspace${random_string.naming.result}"
  resource_group_name         = azurerm_resource_group.example.name
  location                    = azurerm_resource_group.example.location
  sku                         = "standard"
  managed_resource_group_name = "workspace${random_string.naming.result}"
}

# Create container in storage acc and container for use by mount tests
resource "azurerm_storage_account" "account" {
  name                     = "${random_string.naming.result}datalake"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
  account_kind             = "StorageV2"
  is_hns_enabled           = "true"
}

data "azurerm_client_config" "current" {
}

resource "azurerm_role_assignment" "datalake" {
  scope = azurerm_storage_account.account.id
  #https://docs.microsoft.com/en-us/azure/role-based-access-control/built-in-roles#storage-blob-data-contributor
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = data.azurerm_client_config.current.object_id
}

resource "azurerm_storage_container" "example" {
  name                  = "dev"
  storage_account_name  = azurerm_storage_account.account.name
  container_access_type = "private"
}

output "workspace_managed_rg_name" {
  value = "workspace${random_string.naming.result}"
}

output "workspace_name" {
  value = azurerm_databricks_workspace.example.name
}

output "gen2_adal_name" {
  value = azurerm_storage_account.account.name
}

output "location" {
  value = azurerm_storage_account.account.location
}

output "rg_name" {
  value = azurerm_resource_group.example.name
}
