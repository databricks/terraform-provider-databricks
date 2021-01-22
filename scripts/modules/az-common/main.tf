terraform {
  required_providers {
    azurerm =  "~> 2.33"
    random = "~> 2.2"
  }
}

provider "azurerm" {
  features {}
}

variable "owner" {
  type    = string
  default = ""
}

variable "region" {
  type = string
  default = "westeurope"
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

data "azurerm_client_config" "current" {
}

locals {
  // dltp - databricks labs terraform provider
  prefix = "dltp${random_string.naming.result}"
  tags = {
    Environment = "Testing"
    Owner       = var.owner
    Epoch       = random_string.naming.result
  }
}

resource "azurerm_resource_group" "example" {
  name     = "${local.prefix}-rg"
  location = var.region
  tags     = local.tags
}

output "azure_region" {
  value = azurerm_resource_group.example.location
}

output "test_resource_group" {
  value = azurerm_resource_group.example.name
}

resource "azurerm_databricks_workspace" "example" {
  name                        = "${local.prefix}-workspace"
  resource_group_name         = azurerm_resource_group.example.name
  location                    = azurerm_resource_group.example.location
  sku                         = "premium"
  managed_resource_group_name = "${local.prefix}-workspace-rg"
  tags                        = local.tags
}

output "databricks_azure_workspace_resource_id" {
  // The ID of the Databricks Workspace in the Azure management plane.
  value = azurerm_databricks_workspace.example.id
}

output "databricks_host" {
  value = "https://${azurerm_databricks_workspace.example.workspace_url}/"
}

// ADLSv1 resource
resource "azurerm_data_lake_store" "gen1" {
  name                = "${local.prefix}adlsv1"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
}

output "test_data_lake_store_name" {
  value = azurerm_data_lake_store.gen1.name
}

# Create container in storage acc and container for use by blob mount tests
resource "azurerm_storage_account" "v2" {
  name                     = "${local.prefix}adlsv2"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  account_kind             = "StorageV2"
  tags                     = local.tags
}

output "test_storage_v2_account" {
  value = azurerm_storage_account.v2.name
}

output "test_storage_v2_key" {
  value     = azurerm_storage_account.v2.primary_access_key
  sensitive = true
}

resource "azurerm_storage_container" "wasbs" {
  name                  = "${local.prefix}-wasbs"
  storage_account_name  = azurerm_storage_account.v2.name
  container_access_type = "private"
}

output "test_storage_v2_wasbs" {
  value     = azurerm_storage_container.wasbs.name
}

data "azurerm_storage_account_blob_container_sas" "example" {
  connection_string = azurerm_storage_account.v2.primary_connection_string
  container_name    = azurerm_storage_container.wasbs.name
  https_only        = true
  start = "2020-02-01"
  expiry = "2021-12-31"
  permissions {
    read   = true
    add    = true
    create = true
    write  = true
    delete = true
    list   = true
  }
}

output "test_storage_v2_wasbs_sas" {
  value = data.azurerm_storage_account_blob_container_sas.example.sas
  sensitive = true
}

resource "azurerm_storage_blob" "example" {
  name                   = "main.tf"
  storage_account_name   = azurerm_storage_account.v2.name
  storage_container_name = azurerm_storage_container.wasbs.name
  type                   = "Block"
  source                 = "${path.module}/main.tf"
}

resource "azurerm_key_vault" "example" {
  name                     = "${local.prefix}-kv"
  location                 = azurerm_resource_group.example.location
  resource_group_name      = azurerm_resource_group.example.name
  tenant_id                = data.azurerm_client_config.current.tenant_id
  soft_delete_enabled      = false
  purge_protection_enabled = false
  sku_name                 = "standard"
  tags                     = local.tags

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id
    secret_permissions = [
      "delete", "get", "list", "set"
    ]
  }
}

output "test_key_vault_name" {
  value = azurerm_key_vault.example.name
}

output "test_key_vault_resource_id" {
  value = azurerm_key_vault.example.id
}

output "test_key_vault_dns_name" {
  value = azurerm_key_vault.example.vault_uri
}

// "Key Vault Administrator" role required for SP
resource "azurerm_key_vault_secret" "example" {
  name         = "answer"
  value        = "42"
  key_vault_id = azurerm_key_vault.example.id
  tags         = local.tags
}

output "test_key_vault_secret" {
  value = azurerm_key_vault_secret.example.name
}

output "test_key_vault_secret_value" {
  // this is for testing purposes only. 
  // must not be a practice for production.
  value = azurerm_key_vault_secret.example.value
  sensitive = true
}

output "cloud_env" {
  value = "azure"
}

output "test_node_type" {
  // or i3.xlarge for AWS
  value = "Standard_D3_v2"
}

output "test_prefix" {
  value = local.prefix
}
